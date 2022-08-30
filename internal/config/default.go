package config

import (
	"github.com/amirhnajafiz/carrot/internal/client"
	"github.com/amirhnajafiz/carrot/internal/rabbit"
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
		Rabbit: rabbit.Config{
			Host: "amqp://guest:guest@localhost:5672/",
		},
		Number:  10,
		Timeout: 5,
	}
}
