package senders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Send(t *testing.T) {
	t.Run("send test", func(t *testing.T) {
		client, err := GetMessageSender("your first queue name")
		assert.Nil(t, err)

		err = client.SetMessage(
			NewMessageBuilder().
				SetTitle("this is title").
				SetAuthor("this is author").
				SetBody("this is body"),
		).Send()
		assert.Nil(t, err)
	})
}
