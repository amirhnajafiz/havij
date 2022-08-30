package client

import (
	"log"
	"strings"

	"github.com/amirhnajafiz/carrot/internal/test"
	"github.com/streadway/amqp"
)

// Client manages the connection to rabbitMQ server.
type Client struct {
	Cfg Config

	Queue      string
	Provider   bool
	Connection *amqp.Connection
}

// Initialize creates a new queue over rabbitMQ.
func (c *Client) Initialize() error {
	// open channel
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	// declare queue
	_, err = ch.QueueDeclare(
		c.Queue,
		c.Cfg.Durable,
		c.Cfg.AutoDelete,
		c.Cfg.Exclusive,
		!c.Cfg.Wait,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

// Publish over rabbitMQ.
func (c *Client) Publish(s string) error {
	// open channel
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	// publish over a channel
	err = ch.Publish(
		"",
		c.Queue,
		c.Cfg.Mandatory,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(s),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// Subscribe over topic.
func (c *Client) Subscribe(timeout int) error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	messages, err := ch.Consume(
		c.Queue,
		"",
		c.Cfg.AutoAck,
		c.Cfg.Exclusive,
		!c.Cfg.Local,
		!c.Cfg.Wait,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range messages {
			parts := strings.Split(string(d.Body), " Brear ")

			r, du := test.Done(parts[0], timeout)

			log.Printf("[test %s][duration %s][timeout %t]: %s \n", parts[0], du, !r, parts[1])
		}
	}()

	<-forever

	return nil
}
