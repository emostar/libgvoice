package libgvoice

import (
	"context"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/appstate"
	waBinary "go.mau.fi/whatsmeow/binary"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/socket"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

// Add functions from whatsmeow client, so we can compile in gvoice bridge

func (c *GoogleVoiceClient) BuildEdit(chat types.JID, id types.MessageID, newContent *waProto.Message) *waProto.Message {
	c.log.Debugf("whatsmeow: BuildEdit called")
	return nil
}
func (c *GoogleVoiceClient) BuildPollCreation(name string, optionNames []string, selectableOptionCount int) *waProto.Message {
	c.log.Debugf("whatsmeow: BuildPollCreation called")
	return nil
}
func (c *GoogleVoiceClient) BuildPollVote(pollInfo *types.MessageInfo, optionNames []string) (*waProto.Message, error) {
	c.log.Debugf("whatsmeow: BulidPollVote called")
	return nil, nil
}
func (c *GoogleVoiceClient) BuildRevoke(chat, sender types.JID, id types.MessageID) *waProto.Message {
	c.log.Debugf("whatsmeow: BuildRevoke called")
	return nil
}
func (c *GoogleVoiceClient) CheckUpdate() (respData whatsmeow.CheckUpdateResponse, err error) {
	c.log.Debugf("whatsmeow: CheckUpdate called")
	return whatsmeow.CheckUpdateResponse{}, nil
}
func (c *GoogleVoiceClient) CreateGroup(req whatsmeow.ReqCreateGroup) (*types.GroupInfo, error) {
	c.log.Debugf("whatsmeow: CreateGroup called")
	return nil, nil
}
func (c *GoogleVoiceClient) DangerousInternals() *whatsmeow.DangerousInternalClient {
	c.log.Debugf("whatsmeow: DangerousInternals called")
	return nil
}
func (c *GoogleVoiceClient) DecryptPollVote(vote *events.Message) (*waProto.PollVoteMessage, error) {
	c.log.Debugf("whatsmeow: DecryptPollVote called")
	return nil, nil
}
func (c *GoogleVoiceClient) DecryptReaction(reaction *events.Message) (*waProto.ReactionMessage, error) {
	c.log.Debugf("whatsmeow: DecryptReaction called")
	return nil, nil

}
func (c *GoogleVoiceClient) Disconnect() {
	c.log.Debugf("whatsmeow: Disconnect called")
}
func (c *GoogleVoiceClient) Download(msg whatsmeow.DownloadableMessage) ([]byte, error) {
	c.log.Debugf("whatsmeow: Download called")
	return nil, nil
}
func (c *GoogleVoiceClient) DownloadAny(msg *waProto.Message) (data []byte, err error) {
	c.log.Debugf("whatsmeow: DownloadAny called")
	return nil, nil
}
func (c *GoogleVoiceClient) DownloadMediaWithPath(directPath string, encFileHash, fileHash, mediaKey []byte, fileLength int, mediaType whatsmeow.MediaType, _ string) (data []byte, err error) {
	c.log.Debugf("whatsmeow: DownloadMediaWithPath called")
	return nil, nil
}
func (c *GoogleVoiceClient) DownloadThumbnail(msg whatsmeow.DownloadableThumbnail) ([]byte, error) {
	c.log.Debugf("whatsmeow: DownloadThumbnail called")
	return nil, nil
}
func (c *GoogleVoiceClient) EncryptPollVote(pollInfo *types.MessageInfo, vote *waProto.PollVoteMessage) (*waProto.PollUpdateMessage, error) {
	c.log.Debugf("whatsmeow: EncryptPollVote called")
	return nil, nil
}
func (c *GoogleVoiceClient) FetchAppState(name appstate.WAPatchName, fullSync, onlyIfNotSynced bool) error {
	c.log.Debugf("whatsmeow: FetchAppState called")
	return nil
}
func (c *GoogleVoiceClient) GetContactQRLink(revoke bool) (string, error) {
	c.log.Debugf("whatsmeow: GetContactQRLink called")
	return "", nil
}
func (c *GoogleVoiceClient) GetGroupInfo(jid types.JID) (*types.GroupInfo, error) {
	c.log.Debugf("whatsmeow: GetGroupInfo called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetGroupInfoFromInvite(jid, inviter types.JID, code string, expiration int64) (*types.GroupInfo, error) {
	c.log.Debugf("whatsmeow: GetGroupInfoFromInvite called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetGroupInfoFromLink(code string) (*types.GroupInfo, error) {
	c.log.Debugf("whatsmeow: GetGroupInfoFromLink called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetGroupInviteLink(jid types.JID, reset bool) (string, error) {
	c.log.Debugf("whatsmeow: GetGroupInviteLink called")
	return "", nil
}
func (c *GoogleVoiceClient) GetJoinedGroups() ([]*types.GroupInfo, error) {
	c.log.Debugf("whatsmeow: GetJoinedGroups called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetLinkedGroupsParticipants(community types.JID) ([]types.JID, error) {
	c.log.Debugf("whatsmeow: GetLinkedGroupsParticipants called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetPrivacySettings() (settings types.PrivacySettings) {
	c.log.Debugf("whatsmeow: GetPrivacySettings called")
	return types.PrivacySettings{}
}
func (c *GoogleVoiceClient) GetProfilePictureInfo(jid types.JID, params *whatsmeow.GetProfilePictureParams) (*types.ProfilePictureInfo, error) {
	c.log.Debugf("whatsmeow: GetProfilePictureInfo called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetQRChannel(ctx context.Context) (<-chan whatsmeow.QRChannelItem, error) {
	c.log.Debugf("whatsmeow: GetQRChannel called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetStatusPrivacy() ([]types.StatusPrivacy, error) {
	c.log.Debugf("whatsmeow: GetStatusPrivacy called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetSubGroups(community types.JID) ([]*types.GroupLinkTarget, error) {
	c.log.Debugf("whatsmeow: GetSubGroups called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetUserDevices(jids []types.JID) ([]types.JID, error) {
	c.log.Debugf("whatsmeow: GetUserDevices called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetUserDevicesContext(ctx context.Context, jids []types.JID) ([]types.JID, error) {
	c.log.Debugf("whatsmeow: GetUserDevicesContext called")
	return nil, nil
}
func (c *GoogleVoiceClient) GetUserInfo(jids []types.JID) (map[types.JID]types.UserInfo, error) {
	c.log.Debugf("whatsmeow: GetUserInfo called")
	return nil, nil
}
func (c *GoogleVoiceClient) IsOnWhatsApp(phones []string) ([]types.IsOnWhatsAppResponse, error) {
	c.log.Debugf("whatsmeow: IsOnWhatsApp called")
	return nil, nil
}
func (c *GoogleVoiceClient) JoinGroupWithInvite(jid, inviter types.JID, code string, expiration int64) error {
	c.log.Debugf("whatsmeow: JoinGroupWithInvite called")
	return nil
}
func (c *GoogleVoiceClient) JoinGroupWithLink(code string) (types.JID, error) {
	c.log.Debugf("whatsmeow: JoinGroupWithLink called")
	return types.JID{}, nil
}
func (c *GoogleVoiceClient) LeaveGroup(jid types.JID) error {
	c.log.Debugf("whatsmeow: LeaveGroup called")
	return nil
}
func (c *GoogleVoiceClient) LinkGroup(parent, child types.JID) error {
	c.log.Debugf("whatsmeow: LinkGroup called")
	return nil
}
func (c *GoogleVoiceClient) Logout() error {
	c.log.Debugf("whatsmeow: Logout called")
	return nil
}
func (c *GoogleVoiceClient) MarkRead(ids []types.MessageID, timestamp time.Time, chat, sender types.JID) error {
	c.log.Debugf("whatsmeow: MarkRead called")
	return nil
}
func (c *GoogleVoiceClient) ParseWebMessage(chatJID types.JID, webMsg *waProto.WebMessageInfo) (*events.Message, error) {
	c.log.Debugf("whatsmeow: ParseWebMessage called")
	return nil, nil
}
func (c *GoogleVoiceClient) RemoveEventHandler(id uint32) bool {
	c.log.Debugf("whatsmeow: RemoveEventHandler called")
	return false
}
func (c *GoogleVoiceClient) RemoveEventHandlers() {
	c.log.Debugf("whatsmeow: RemoveEventHandlers called")
}
func (c *GoogleVoiceClient) ResolveBusinessMessageLink(code string) (*types.BusinessMessageLinkTarget, error) {
	c.log.Debugf("whatsmeow: ResolveBusinessMessageLink called")
	return nil, nil
}
func (c *GoogleVoiceClient) ResolveContactQRLink(code string) (*types.ContactQRLinkTarget, error) {
	c.log.Debugf("whatsmeow: ResolveContactQRLink called")
	return nil, nil
}
func (c *GoogleVoiceClient) RevokeMessage(chat types.JID, id types.MessageID) (whatsmeow.SendResponse, error) {
	c.log.Debugf("whatsmeow: RevokeMessage called")
	return whatsmeow.SendResponse{}, nil
}
func (c *GoogleVoiceClient) SendChatPresence(jid types.JID, state types.ChatPresence, media types.ChatPresenceMedia) error {
	c.log.Debugf("whatsmeow: SendChatPresence called")
	return nil
}
func (c *GoogleVoiceClient) SendMediaRetryReceipt(message *types.MessageInfo, mediaKey []byte) error {
	c.log.Debugf("whatsmeow: SendMediaRetryReceipt called")
	return nil
}
func (c *GoogleVoiceClient) SendMessage(ctx context.Context, to types.JID, message *waProto.Message, extra ...whatsmeow.SendRequestExtra) (resp whatsmeow.SendResponse, err error) {
	c.log.Debugf("whatsmeow: SendMessage called")
	return whatsmeow.SendResponse{}, nil
}
func (c *GoogleVoiceClient) SendPresence(state types.Presence) error {
	c.log.Debugf("whatsmeow: SendPresence called")
	return nil
}
func (c *GoogleVoiceClient) SetDisappearingTimer(chat types.JID, timer time.Duration) (err error) {
	c.log.Debugf("whatsmeow: SetDisappearingTimer called")
	return nil
}
func (c *GoogleVoiceClient) SetForceActiveDeliveryReceipts(active bool) {
	c.log.Debugf("whatsmeow: SetForceActiveDeliveryReceipts called")
}
func (c *GoogleVoiceClient) SetGroupAnnounce(jid types.JID, announce bool) error {
	c.log.Debugf("whatsmeow: SetGroupAnnounce called")
	return nil
}
func (c *GoogleVoiceClient) SetGroupLocked(jid types.JID, locked bool) error {
	c.log.Debugf("whatsmeow: SetGroupLocked called")
	return nil
}
func (c *GoogleVoiceClient) SetGroupName(jid types.JID, name string) error {
	c.log.Debugf("whatsmeow: SetGroupName called")
	return nil
}
func (c *GoogleVoiceClient) SetGroupPhoto(jid types.JID, avatar []byte) (string, error) {
	c.log.Debugf("whatsmeow: SetGroupPhoto called")
	return "", nil
}
func (c *GoogleVoiceClient) SetGroupTopic(jid types.JID, previousID, newID, topic string) error {
	c.log.Debugf("whatsmeow: SetGroupTopic called")
	return nil
}
func (c *GoogleVoiceClient) SetPassive(passive bool) error {
	c.log.Debugf("whatsmeow: SetPassive called")
	return nil
}
func (c *GoogleVoiceClient) SetProxy(proxy socket.Proxy) {
	c.log.Debugf("whatsmeow: SetProxy called")
}
func (c *GoogleVoiceClient) SetProxyAddress(addr string) error {
	c.log.Debugf("whatsmeow: SetProxyAddress called")
	return nil
}
func (c *GoogleVoiceClient) SetStatusMessage(msg string) error {
	c.log.Debugf("whatsmeow: SetStatusMessage called")
	return nil
}
func (c *GoogleVoiceClient) SubscribePresence(jid types.JID) error {
	c.log.Debugf("whatsmeow: SubscribePresence called")
	return nil
}
func (c *GoogleVoiceClient) TryFetchPrivacySettings(ignoreCache bool) (*types.PrivacySettings, error) {
	c.log.Debugf("whatsmeow: TryFetchPrivacySettings called")
	return nil, nil
}
func (c *GoogleVoiceClient) UnlinkGroup(parent, child types.JID) error {
	c.log.Debugf("whatsmeow: UnlinkGroup called")
	return nil
}
func (c *GoogleVoiceClient) UpdateGroupParticipants(jid types.JID, participantChanges map[types.JID]whatsmeow.ParticipantChange) (*waBinary.Node, error) {
	c.log.Debugf("whatsmeow: UpdateGroupParticipants called")
	return nil, nil
}
func (c *GoogleVoiceClient) Upload(ctx context.Context, plaintext []byte, appInfo whatsmeow.MediaType) (resp whatsmeow.UploadResponse, err error) {
	c.log.Debugf("whatsmeow: Upload called")
	return whatsmeow.UploadResponse{}, nil
}
func (c *GoogleVoiceClient) WaitForConnection(timeout time.Duration) bool {
	c.log.Debugf("whatsmeow: WaitForConnection called")
	return false
}
