package cmd

import (
	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
)

func Execute() {
	{
		r, _ := rabbitMQT.Init(rabbitMQT.Config{})
		cli := client.Client{
			Connection: r,
		}

		_ = cli.Listen()
	}
	{
		r, _ := rabbitMQT.Init(rabbitMQT.Config{})
		cli := client.Client{
			Connection: r,
		}

		_ = cli.Push()
	}
}
