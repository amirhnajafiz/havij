package config

import (
	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
	"github.com/amirhnajafiz/playful-rabbit/internal/test"
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
			Immediate:  true,
			Local:      false,
			AutoAck:    true,
		},
		Rabbit: rabbitMQT.Config{
			Host: "amqp://127.0.0.1:15672",
		},
		Test: test.Config{
			Number:  10,
			Timeout: 5,
		},
	}
}
