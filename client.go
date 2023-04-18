package libgvoice

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/beeper/libgvoice/api"
	gvEvents "github.com/beeper/libgvoice/events"
	"github.com/beeper/libgvoice/models"
	"github.com/beeper/libgvoice/util"
	gvLog "github.com/beeper/libgvoice/util/log"
	"github.com/bitly/go-simplejson"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

// GoogleVoiceClient interfaces with the Google Voice API over HTTPS and works
// similar to how the web client at https://voice.google.com works.
type GoogleVoiceClient struct {
	http *http.Client

	baseURL string
	apiKey  string

	apiSIDHash  string
	cookieParam string

	log gvLog.Logger

	// When a message is received, it is added to this map. Ideally this would
	// be backed up to a database, but for now it's just in memory.
	seenMessages map[string]bool

	// Channel to send events to the caller of the client
	eventChannel chan interface{}

	// XXX WhatsApp Fork Code
	// Remove after switch from whatsapp fork is complete

	Store *store.Device
	Log   waLog.Logger

	EnableAutoReconnect   bool
	LastSuccessfulConnect time.Time
	AutoReconnectErrors   int

	// EmitAppStateEventsOnFullSync can be set to true if you want to get app state events emitted
	// even when re-syncing the whole state.
	EmitAppStateEventsOnFullSync bool

	// GetMessageForRetry is used to find the source message for handling retry receipts
	// when the message is not found in the recently sent message cache.
	GetMessageForRetry func(requester, to types.JID, id types.MessageID) *waProto.Message
	// PreRetryCallback is called before a retry receipt is accepted.
	// If it returns false, the accepting will be cancelled and the retry receipt will be ignored.
	PreRetryCallback func(receipt *events.Receipt, id types.MessageID, retryCount int, msg *waProto.Message) bool

	// PrePairCallback is called before pairing is completed. If it returns false, the pairing will be cancelled and
	// the client will disconnect.
	PrePairCallback func(jid types.JID, platform, businessName string) bool

	// Should untrusted identity errors be handled automatically? If true, the stored identity and existing signal
	// sessions will be removed on untrusted identity errors, and an events.IdentityChange will be dispatched.
	// If false, decrypting a message from untrusted devices will fail.
	AutoTrustIdentity bool

	// Should sending to own devices be skipped when sending broadcasts?
	// This works around a bug in the WhatsApp android app where it crashes if you send a status message from a linked device.
	DontSendSelfBroadcast bool

	// Should SubscribePresence return an error if no privacy token is stored for the user?
	ErrorOnSubscribePresenceWithoutToken bool
}

// NewGoogleVoiceClient creates a new GoogleVoiceClient to interact with the API
func NewGoogleVoiceClient(log gvLog.Logger, eventChannel chan interface{}) *GoogleVoiceClient {
	return &GoogleVoiceClient{
		http:         &http.Client{},
		baseURL:      "https://clients6.google.com/voice/v1/voiceclient",
		apiKey:       "AIzaSyDTYc1N4xiODyrQYK0Kl6g_y279LjYkrBg",
		log:          log,
		seenMessages: make(map[string]bool),
		eventChannel: eventChannel,
	}
}

// SetAuth takes the full list of cookies from the browser and creates the
// necessary auth headers for the API.
func (c *GoogleVoiceClient) SetAuth(cookieParam string) {
	c.cookieParam = cookieParam
	c.apiSIDHash = util.ExtractSID(cookieParam)
}

// Connect will take the cookies, call SetAuth, verify the cookies are valid,
// and start listening for new messages via the browser channel.
func (c *GoogleVoiceClient) Connect(cookies string) error {
	c.log.Debugf("Connecting to Google Voice...")
	c.SetAuth(cookies)
	info, err := c.GetAccountInfo()
	if err != nil {
		return err
	}

	// TODO Start the browser channel

	c.eventChannel <- &gvEvents.ConnectedEvent{PrimaryDID: info.PrimaryDID}

	return nil
}

// IsConnected will return true if Auth has been setup, otherwise false
func (c *GoogleVoiceClient) IsConnected() bool {
	c.log.Debugf("libgvoice: IsConnected called")
	return c.apiSIDHash != "" && c.cookieParam != ""
}

// IsLoggedIn will return true if Auth has been setup, otherwise false
func (c *GoogleVoiceClient) IsLoggedIn() bool {
	c.log.Debugf("libgvoice: IsLoggedIn called")
	return c.IsConnected()
}

// GetAccountInfo is used to get all the account details of the current user.
func (c *GoogleVoiceClient) GetAccountInfo() (*api.AccountInfo, error) {
	protobufMessage := "[null,1]"

	u := fmt.Sprintf("%s/account/get?alt=json&key=%s", c.baseURL, c.apiKey)
	json, err := c.doRequest("POST", u, protobufMessage)
	if err != nil {
		return nil, err
	}

	resp := &api.AccountInfo{
		PrimaryDID: json.GetPath("account", "primaryDid").MustString(),
	}

	return resp, nil
}

// SendSMS sends a text message to the given thread. For a single recipient,
// the thread will always be in the format of "t.+1XXXXXXXXXX". For a group,
// there is a different format which is not yet supported. If you have an
// existing thread for a group message, you can send a message to that thread
// and it will work properly.
func (c *GoogleVoiceClient) SendSMS(threadID, msg string) (*api.MessageResponse, error) {
	protobufMessage := fmt.Sprintf(
		"[null,null,null,null,\"%s\",\"%s\",[],null,[%d]]",
		msg, threadID, rand.Int63(),
	)
	c.log.Debugf("protobufMessage: %s", protobufMessage)

	u := fmt.Sprintf("%s/api2thread/sendsms?alt=json&key=%s", c.baseURL, c.apiKey)
	json, err := c.doRequest("POST", u, protobufMessage)
	if err != nil {
		return nil, err
	}

	timestampStr, err := json.Get("timestampMs").String()
	if err != nil {
		return nil, err
	}
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return nil, err
	}

	resp := &api.MessageResponse{
		ID:        json.Get("threadItemId").MustString(),
		Timestamp: time.UnixMilli(timestamp),
	}
	return resp, nil
}

func (c *GoogleVoiceClient) GetThread(thread *models.Thread) ([]models.Message, error) {
	protobufMessage := fmt.Sprintf(
		"[\"%s\",40,\"\",[null,true,true]]", thread.ID,
	)
	c.log.Debugf("protobufMessage: %s", protobufMessage)

	u := fmt.Sprintf("%s/api2thread/get?alt=json&key=%s", c.baseURL, c.apiKey)
	json, err := c.doRequest("POST", u, protobufMessage)
	if err != nil {
		return nil, err
	}

	jsonMessages := json.Get("thread").Get("item").MustArray()
	messages := make([]models.Message, 0, len(jsonMessages))

	for idx := range jsonMessages {
		jsonMessage := json.Get("thread").Get("item").GetIndex(idx)
		messageMap := jsonMessage.MustMap()

		message := models.Message{
			ID:         messageMap["id"].(string),
			Thread:     thread,
			SenderE164: jsonMessage.Get("contact").Get("phoneNumber").MustString(),
			Body:       messageMap["messageText"].(string),
		}
		timestamp, _ := strconv.ParseInt(messageMap["startTime"].(string), 10, 64)
		message.Timestamp = time.UnixMilli(timestamp)
		if messageMap["type"].(string) == "smsIn" {
			message.Direction = models.DirectionInbound
		} else {
			message.Direction = models.DirectionOutbound
		}

		messages = append(messages, message)
	}

	return messages, nil
}

// FetchInbox fetches the inbox for the current user. The fetchPage parameter is
// for pagination, which is not implemented yet. The alertNew parameter is set
// to false on the first call to cache the messages. Subsequent calls should
// be true, so new messages can be detected.
//
// In the future we should support inbox type, thread count, and max recent
// parameters, that get set in the protobufMessage.
func (c *GoogleVoiceClient) FetchInbox(fetchPage string, alertNew bool) ([]models.Thread, error) {
	// Format:
	// Inbox type (2 = SMS, 3 = CallHistory)
	// Thread count
	// Max recent messages, can be all in one thread or spread out
	protobufMessage := fmt.Sprintf(
		"[2,20,30,\"%s\",null,[null,true,true]]",
		fetchPage,
	)
	c.log.Debugf("outbound protobufMessage: %s", protobufMessage)

	u := fmt.Sprintf("%s/api2thread/list?alt=json&key=%s", c.baseURL, c.apiKey)
	json, err := c.doRequest("POST", u, protobufMessage)
	if err != nil {
		return nil, err
	}

	jsonThreads := json.Get("thread").MustArray()
	threads := make([]models.Thread, 0, len(jsonThreads))

	for idx := range jsonThreads {
		threadObj := json.Get("thread").GetIndex(idx).MustMap()
		jsonItems := json.Get("thread").GetIndex(idx).Get("item").MustArray()
		jsonContacts := json.Get("thread").GetIndex(idx).Get("contact").MustArray()

		thread := models.Thread{
			ID:       threadObj["id"].(string),
			IsRead:   threadObj["read"].(bool),
			Messages: make([]models.Message, 0, len(jsonItems)),
			Contacts: make([]models.Contact, 0, len(jsonContacts)),
		}

		for contactIndex := range jsonContacts {
			contactObj := json.Get("thread").GetIndex(idx).Get("contact").GetIndex(contactIndex).MustMap()
			contact := models.Contact{
				Name:        contactObj["name"].(string),
				PhoneNumber: contactObj["phoneNumber"].(string),
			}
			thread.Contacts = append(thread.Contacts, contact)
		}

		for itemIndex := range jsonItems {
			itemObj := json.Get("thread").GetIndex(idx).Get("item").GetIndex(itemIndex).MustMap()
			message := models.Message{
				ID:         itemObj["id"].(string),
				Timestamp:  util.StringMilliTimestampToTime(itemObj["startTime"].(string)),
				SenderE164: itemObj["did"].(string),
				Status:     itemObj["status"].(string),
				Body:       itemObj["messageText"].(string),
				Thread:     &thread,
			}
			if itemObj["type"].(string) == "smsOut" {
				message.Direction = models.DirectionOutbound
			} else {
				message.Direction = models.DirectionInbound
			}
			if itemObj["messageId"] != nil {
				message.MessageID = itemObj["messageId"].(string)
			}

			// Check if we should alert that we have a new inbound message
			if alertNew {
				if _, ok := c.seenMessages[message.ID]; !ok {
					c.log.Infof("new message: %s %#v", message.ID, message)
					c.seenMessages[message.ID] = true
				}
			}

			// Save that we have seen this message
			c.seenMessages[message.ID] = true

			thread.Messages = append(thread.Messages, message)
		}

		threads = append(threads, thread)
	}

	// TODO - handle pagination
	nextPageToken := json.Get("paginationToken").MustString()
	_ = nextPageToken

	return threads, nil
}

// StartEventListener is meant to be run in a goroutine. It will listen for
// events from the browser channel. Currently, it only prints out any new
// incoming messages. In the future we should support pushing new inbound
// messages over a channel, so the caller can be alerted in real-time of new
// messages.
func (c *GoogleVoiceClient) StartEventListener() {
	eventChannel := make(chan BrowserChannelEvent)
	bc := NewBrowserChannel(eventChannel, c.log)
	bc.SetAuth(c.cookieParam)

	go bc.StartEventListener()

	var lastEvent BrowserChannelEvent

	for {
		select {
		case event := <-eventChannel:
			c.log.Infof("**** EVENT: %d", event)
			if lastEvent == NoopEvent && event == NoopEvent {
				c.log.Infof("2 noops in a row, resetting connection")
				bc.ResetData()
			}
			lastEvent = event

			if event == RefreshInboxEvent {
				_, _ = c.FetchInbox("", true)
			}
		}
	}
}

// buildRequest builds a new HTTP request, adding the required headers.
func (c *GoogleVoiceClient) buildRequest(method, rawURL string, body io.Reader, contentType string) (*http.Request, error) {
	if c.apiSIDHash == "" {
		return nil, errors.New("missing SID hash")
	}

	if c.cookieParam == "" {
		return nil, errors.New("missing auth cookie")
	}

	var bodyw io.Reader
	if contentType != "" {
		bodyw = body
	}
	req, err := http.NewRequest(method, rawURL, bodyw)
	if err != nil {
		return nil, err
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// If we want to change the User-Agent and then encode it into
	// X-ClientDetails, as well as update the User-Agent header

	if parsedURL.Host == "clients6.google.com" {
		req.Header.Add(
			"X-ClientDetails",
			"appVersion=5.0%20(X11%3B%20Ubuntu)&platform=Linux%20x86_64&userAgent=Mozilla%2F5.0%20(X11%3B%20Ubuntu%3B%20Linux%20x86_64%3B%20rv%3A109.0)%20Gecko%2F20100101%20Firefox%2F109.0",
		)
		req.Header.Add("X-Requested-With", "XMLHttpRequest")
		req.Header.Add("X-JavaScript-User-Agent", "google-api-javascript-client/1.1.0")
		req.Header.Add("X-Client-Version", "512793257")
		req.Header.Add("X-Origin", "https://voice.google.com")
		req.Header.Add("X-Referer", "https://voice.google.com")
		req.Header.Add("X-Goog-Encode-Response-If-Executable", "base64")
		req.Header.Add("Origin", "https://clients6.google.com")
		req.Header.Add("Referer", "https://clients6.google.com/static/proxy.html?usegapi=1")
		req.Header.Add("Sec-Fetch-Site", "same-origin")
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Authorization", c.apiSIDHash)
	req.Header.Add("X-Goog-AuthUser", "0")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-GPC", "1")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("TE", "trailers")
	req.Header.Add("Cookie", c.cookieParam)
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	return req, nil
}

// doRequest performs an HTTP request, and returns the response as a JSON
func (c *GoogleVoiceClient) doRequest(method, rawURL, body string) (*simplejson.Json, error) {
	req, err := c.buildRequest(method, rawURL, bytes.NewBufferString(body), "application/json+protobuf")
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp.Body = http.MaxBytesReader(nil, resp.Body, 1<<20) // 1MB max
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	json, err := simplejson.NewJson(respBody)
	if err != nil {
		return nil, err
	}

	if apiError, ok := json.CheckGet("error"); ok {
		return nil, errors.New(apiError.Get("message").MustString())
	}

	return json, nil
}
