package senders

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"strings"
)

var sqsClient *sqs.SQS

func init() {
	sqsClient = sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}

type MessageSender struct {
	QueueConfig
	Message *Message
}

var messageSenderFactories = make(map[string]*MessageSender)

func Register(sender *MessageSender) {
	if sender == nil {
		log.Printf("MessageSenderFactory %s does not exist.", sender.QueueName)
		return
	}
	_, registered := messageSenderFactories[sender.QueueName]
	if registered {
		log.Printf("MessageSenderFactory %s already registered. Ignoring.", sender.QueueName)
	}
	messageSenderFactories[sender.QueueName] = sender
}

func GetMessageSender(name string) (*MessageSender, error) {
	factory, ok := messageSenderFactories[name]
	if !ok {
		availableMessageSenders := make([]string, 0)
		for messageFactoryName, _ := range messageSenderFactories {
			availableMessageSenders = append(availableMessageSenders, messageFactoryName)
		}
		return nil, errors.New(fmt.Sprintf("Invalid MessageSender name. Must be one of: %s",
			strings.Join(availableMessageSenders, ", ")))
	}
	return factory, nil
}

func (m *MessageSender) SetMessage(message *Message) *MessageSender {
	m.Message = message
	return m
}

func (c *MessageSender) Send() error {
	if c.Message == nil {
		return errors.New("message field is empty, can't send message")
	}

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(c.Message.title),
			},
			"author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(c.Message.author),
			},
		},
		MessageBody: aws.String(c.Message.body),
		QueueUrl:    &c.QueueUrl,
	})
	if err != nil {
		return err
	}
	return nil
}
