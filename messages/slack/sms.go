package slack

import "awesomeProject1/messages"

type Slack struct {
}

func New() *Slack {
	return &Slack{}
}

func (s *Slack) Send(m messages.Payload) {}
