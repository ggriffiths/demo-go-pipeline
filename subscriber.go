package main

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// SubscriberConfig represents a configuration for subscribing to Kafka
type SubscriberConfig struct {
}

func newKafkaSubscriber(addrs []string, groupID string, topics []string) (*cluster.Consumer, error) {
	config := cluster.Config{}
	config.Version = sarama.V1_0_0_0
	return cluster.NewConsumer(addrs, groupID, topics, &config)
}
