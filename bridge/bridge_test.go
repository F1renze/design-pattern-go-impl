package bridge

import "testing"

func TestBridge(t *testing.T) {
	n := NewNormalNotification()

	n.SetMsgSender(NewEmailMsgSender())
	n.Notify("notify")
}
