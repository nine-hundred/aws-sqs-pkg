package senders

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

type QueueConfig struct {
	QueueName string
	QueueUrl  string
}

var QueueConfigs []QueueConfig

func init() {
	absDir, _ := filepath.Abs("./")
	viper.SetConfigFile(absDir + "/config.json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Unable to read config file:", err)
	}

	viper.UnmarshalKey("QueueConfigs", &QueueConfigs)

	for _, config := range QueueConfigs {
		Register(&MessageClient{config})
	}
}
