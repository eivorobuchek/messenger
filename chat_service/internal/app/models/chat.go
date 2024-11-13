package models

type Message struct {
	ID         string
	SenderID   string
	ReceiverID string
	Content    string
	Timestamp  int64
}

type Chat struct {
	Messages []*Message
}

func (c *Chat) AddMessage(msg *Message) {
	c.Messages = append(c.Messages, msg)
}
