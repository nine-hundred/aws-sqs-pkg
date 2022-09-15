package senders

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Send(t *testing.T) {
	t.Run("", func(t *testing.T) {
		client, err := GetMessageSender("your first queue name")
		assert.Nil(t, err)

		err = client.SetMessage(
			NewMessageBuilder().
				SetTitle("").
				SetAuthor(":").
				SetBody(""),
		).Send()
		assert.Nil(t, err)
	})
}
