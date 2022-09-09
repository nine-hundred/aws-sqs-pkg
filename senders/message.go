package senders

type Message struct {
	author string
	title  string
	body   string
}

func NewMessageBuilder() *Message {
	return &Message{}
}
func (m *Message) SetTitle(title string) *Message {
	m.title = title
	return m
}

func (m *Message) SetAuthor(author string) *Message {
	m.author = author
	return m
}

func (m *Message) SetBody(body string) *Message {
	m.body = body
	return m
}

func (m *Message) GetMessage() *Message {
	return m
}
