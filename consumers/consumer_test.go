package consumers

import (
	"aws-sqs-pkg/payloads"
	"github.com/aws/aws-sdk-go/service/sqs"
	"testing"
)

func Test_Consume(t *testing.T) {
	t.Run("consume test", func(t *testing.T) {
		c := Consumer{
			MessageCh:   make(chan *sqs.Message, 2),
			ProcessorCh: make(chan payloads.Payload, 2),
		}
		c.Consume()
	})
}
