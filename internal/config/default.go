package config

import (
	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
)

func Default() Config {
	return Config{
		Queue: "master",
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
			Host: "",
		},
	}
}
