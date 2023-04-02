package libgvoice_test

import (
	"testing"

	"github.com/beeper/libgvoice"
)

func TestBrowserChannel(t *testing.T) {
	eventChannel := make(chan libgvoice.BrowserChannelEvent)
	gv := libgvoice.NewBrowserChannel(eventChannel, logger.Sugar())
	gv.SetAuth(cookies)

	go gv.StartEventListener()

	var lastEvent libgvoice.BrowserChannelEvent

	for {
		select {
		case event := <-eventChannel:
			logger.Sugar().Infof("**** EVENT: %d", event)
			if lastEvent == libgvoice.NoopEvent && event == libgvoice.NoopEvent {
				logger.Sugar().Infof("2 noops in a row, resetting connection")
				gv.ResetData()
			}
			lastEvent = event
		}
	}
}
