package main

import (
	"github.com/gocql/gocql"
)

// SinkConfig represents a configuration for the CassandraSink struct
type SinkConfig struct {
}

// CassandraSink represents a medium to write to Cassandra
type CassandraSink struct {
	cfg      SinkConfig
	Sesssion *gocql.Session
}

func newCassandraSink() (CassandraSink, error) {
	return CassandraSink{}, nil
}

// Write persists an event to Cassandra
func (cs *CassandraSink) Write(e *Event) error {
	return nil
}
