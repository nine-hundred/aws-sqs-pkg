package consumers

import (
	"aws-sqs-pkg/payloads"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sqsClient *sqs.SQS
var qURL = "your queue url"

func init() {
	sqsClient = sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}

type Consumer struct {
	MessageCh   chan *sqs.Message
	ProcessorCh chan payloads.Payload
}

func (c *Consumer) receiveMessage() ([]*sqs.Message, error) {
	result, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(60), // 60 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})
	if err != nil {
		return nil, err
	}

	if len(result.Messages) == 0 {
		return nil, errors.New("received no message")
	}

	for _, message := range result.Messages {
		c.MessageCh <- message
	}

	return result.Messages, nil
}

func deleteMessage(message *sqs.Message) {
	sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: message.ReceiptHandle,
	})
}

func (c *Consumer) Consume() error {
	messages, err := c.receiveMessage()
	if err != nil {
		return err
	}
	deleteMessage(messages[0])
	return nil
}
