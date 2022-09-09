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

type MessageClient struct {
	QueueConfig
}

var messageClientFactories = make(map[string]*MessageClient)

func Register(client *MessageClient) {
	if client == nil {
		log.Printf("MessageClientFactory %s does not exist.", client.QueueName)
		return
	}
	_, registered := messageClientFactories[client.QueueName]
	if registered {
		log.Printf("MessageClientFactory %s already registered. Ignoring.", client.QueueName)
	}
	messageClientFactories[client.QueueName] = client
}

func GetMessageClient(name string) (*MessageClient, error) {
	factory, ok := messageClientFactories[name]
	if !ok {
		availableMessageClients := make([]string, 0)
		for messageFactoryName, _ := range messageClientFactories {
			availableMessageClients = append(availableMessageClients, messageFactoryName)
		}
		return nil, errors.New(fmt.Sprintf("Invalid MessageClient name. Must be one of: %s",
			strings.Join(availableMessageClients, ", ")))
	}
	return factory, nil
}

func (c *MessageClient) Send(message *Message) error {
	_, err := sqs.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))).SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"title": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message.title),
			},
			"author": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(message.author),
			},
		},
		MessageBody: aws.String(message.body),
		QueueUrl:    &c.QueueUrl,
	})
	if err != nil {
		return err
	}
	return nil
}
