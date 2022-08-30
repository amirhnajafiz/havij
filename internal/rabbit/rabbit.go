package rabbit

import (
	"github.com/streadway/amqp"
)

// Connect to rabbitMQ server.
func Connect(cfg Config) (*amqp.Connection, error) {
	return amqp.Dial(cfg.Host)
}
