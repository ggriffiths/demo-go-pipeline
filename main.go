package main

import (
	"log"
)

func main() {
	subscriber, err := newKafkaSubscriber([]string{"kafka-1"}, "demo-group", []string{"demo-topic"})
	if err != nil {
		log.Fatalf("Failed to create new Kafka Subscriber: %s", err.Error())
	}

	sink, err := newCassandraSink()
	if err != nil {
		log.Fatalf("Failed to create new Cassandra Sink: %s", err.Error())
	}

	for {
		select {
		case msg := <-subscriber.Messages():
			event, err := Transform(msg)
			if err != nil {
				log.Printf("Failed to transform message to an event: %s", err.Error())
			}

			if err := sink.Write(event); err != nil {
				log.Printf("Failed to write event to Cassandra: %s", err.Error())
			}
		}
	}
}
