package client

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Client struct {
	Queue      string
	Connection *amqp.Connection
}

func (c *Client) Push() error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	defer func() {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = ch.QueueDeclare(
		c.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		c.Queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello"),
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Listen() error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	defer func() {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}()

	messages, err := ch.Consume(
		c.Queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("%v \n", d)
		}
	}()

	<-forever

	return nil
}
