package bridge

import "fmt"

type INotification interface {
	SetMsgSender(sender IMsgSender)
	Notify(msg string)
}

type IMsgSender interface {
	Send(msg string)
}

func NewNormalNotification() INotification {
	return &NormalNotification{}
}

type NormalNotification struct {
	sender IMsgSender
}

func (n *NormalNotification) SetMsgSender(sender IMsgSender) {
	n.sender = sender
}

func (n *NormalNotification) Notify(msg string) {
	n.sender.Send(msg)
}

func NewEmailMsgSender() IMsgSender {
	return &EmailMsgSender{}
}

type EmailMsgSender struct {
}

func (m *EmailMsgSender) Send(msg string) {
	fmt.Printf("Email '%s' sending\n", msg)
}



