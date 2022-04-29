package main

import (
	"awesomeProject1/messages"
	"awesomeProject1/messages/sms"
)

// CONSUMER
type MessageSender interface {
	Send(payload messages.Payload)
}

func main() {
	s := sms.New()
	p := messages.Payload{}

	go SendMessage(s, p)
}

func SendMessage(sender MessageSender, payload messages.Payload) {
	sender.Send(payload)
}
