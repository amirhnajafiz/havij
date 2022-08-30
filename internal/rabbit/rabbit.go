package rabbit

import (
	"github.com/streadway/amqp"
)

func Init(cfg Config) (*amqp.Connection, error) {
	return amqp.Dial(cfg.Host)
}
