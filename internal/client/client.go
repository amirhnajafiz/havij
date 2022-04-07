package client

import "github.com/streadway/amqp"

type Client struct {
	Connection *amqp.Connection
}
