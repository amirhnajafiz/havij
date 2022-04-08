package cmd

import (
	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/config"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
)

func Execute() {
	c := config.Load()

	{
		r, _ := rabbitMQT.Init(c.Rabbit)
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
		}

		_ = cli.Listen()
	}
	{
		r, _ := rabbitMQT.Init(c.Rabbit)
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
		}

		_ = cli.Push()
	}
}
