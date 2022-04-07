package client

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Client struct {
	Connection *amqp.Connection
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

	msgs, err := ch.Consume("Queue", "", true, false, false, false, nil)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("%v \n", d)
		}
	}()

	<-forever

	return nil
}
