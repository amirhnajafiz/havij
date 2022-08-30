package config

import (
	"github.com/amirhnajafiz/carrot/internal/client"
	"github.com/amirhnajafiz/carrot/internal/rabbitMQT"
	"github.com/amirhnajafiz/carrot/internal/test"
)

func Default() Config {
	return Config{
		Queue:  "master",
		Prefix: " Brear ",
		Client: client.Config{
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			Wait:       true,
			Mandatory:  false,
			Local:      false,
			AutoAck:    true,
		},
		Rabbit: rabbitMQT.Config{
			Host: "amqp://guest:guest@localhost:5672/",
		},
		Test: test.Config{
			Number:  10,
			Timeout: 5,
		},
	}
}
