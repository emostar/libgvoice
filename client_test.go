package libgvoice_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/beeper/libgvoice"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var (
	cookies string
	logger  *zap.Logger
)

func init() {
	panic("Set your cookie in init and remove this panic")
	// cookies = ""
	// logger, _ = zap.NewDevelopment()
}

func TestClient_GetAccountInfo(t *testing.T) {
	gv := libgvoice.NewGoogleVoiceClient(logger.Sugar())
	gv.SetAuth(cookies)

	info, err := gv.GetAccountInfo()
	assert.NoError(t, err)
	assert.NotNil(t, info)
	fmt.Println(info.PrimaryDID)
}

func TestFetchInbox(t *testing.T) {
	gv := libgvoice.NewGoogleVoiceClient(logger.Sugar())
	gv.SetAuth(cookies)

	// Fetch inbox
	threads, err := gv.FetchInbox("", false)
	assert.Nil(t, err)
	assert.NotEmpty(t, threads, "No threads found in inbox")
	assert.NotEmpty(t, threads[0].ID, "No ID found in first thread")
	assert.NotEmpty(t, threads[0].Messages, "No messages found in first thread")
	assert.NotEmpty(t, threads[0].Messages[0].ID, "No ID found in first message")

	logger.Sugar().Infof("Thread count: %d", len(threads))
	logger.Sugar().Infof("First thread: %#v", threads[0])
	logger.Sugar().Infof("First message: %#v", threads[0].Messages[0])
}

func TestFetchThread(t *testing.T) {

}

func TestSendMessage(t *testing.T) {
	gv := libgvoice.NewGoogleVoiceClient(logger.Sugar())
	gv.SetAuth(cookies)

	// https://7sim.org/free-phone-number-Ambr41nA
	msgResponse, err := gv.SendSMS("t.+12087444182", "Hey there!")
	if err != nil {
		t.Fatal(err.Error())
	}
	logger.Sugar().Infoln("Sent message ID:", msgResponse.ID)
}

func TestNewMessageListener(t *testing.T) {
	gv := libgvoice.NewGoogleVoiceClient(logger.Sugar())
	gv.SetAuth(cookies)

	// Start listening for new messages
	_, _ = gv.FetchInbox("", false)
	go gv.StartEventListener()

	// Wait for 1 minute
	time.Sleep(1 * time.Minute)
}
