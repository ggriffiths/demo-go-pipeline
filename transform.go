package main

import (
	"github.com/Shopify/sarama"
)

// Event represents a transformation event
type Event struct {
}

// Transform take a sarama Message and transforms it based
// on our business logic
func Transform(msg *sarama.ConsumerMessage) (*Event, error) {
	return &Event{}, nil
}
