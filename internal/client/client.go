package client

import (
	"log"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

// Storage manages the tests for client.
type Storage interface {
	Done(string, int) (bool, time.Duration)
	Generate() (string, string)
}

// Client manages the connection to rabbitMQ server.
type Client struct {
	Cfg Config

	Prefix     string
	Queue      string
	Provider   bool
	Connection *amqp.Connection

	Storage Storage
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
func (c *Client) Publish() error {
	// open channel
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	// create a new message
	id, content := c.Storage.Generate()
	s := id + c.Prefix + content

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

			r, du := c.Storage.Done(parts[0], timeout)

			log.Printf("[storage %s][duration %s][timeout %t]: %s \n", parts[0], du, !r, parts[1])
		}
	}()

	<-forever

	return nil
}
