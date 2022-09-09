package senders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Send(t *testing.T) {
	t.Run("", func(t *testing.T) {
		client, err := GetMessageClient("your first queue name")
		assert.Nil(t, err)

		message := NewMessageBuilder().
			SetTitle("title").
			SetAuthor("auther").
			SetBody("this is body")
		err = client.Send(message)
		assert.Nil(t, err)
	})
}
