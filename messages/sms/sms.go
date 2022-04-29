package sms

import "awesomeProject1/messages"

type Sms struct {
}

func New() *Sms {
	return &Sms{}
}

func (s *Sms) Send(m messages.Payload) {}
