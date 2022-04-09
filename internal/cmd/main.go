package cmd

import (
	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/config"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
	"github.com/amirhnajafiz/playful-rabbit/internal/test"
)

func Execute() {
	c := config.Load()
	tests := test.Generate(20)

	{
		r, _ := rabbitMQT.Init(c.Rabbit)
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		_ = cli.Listen()
	}
	{
		r, _ := rabbitMQT.Init(c.Rabbit)
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		for _, t := range tests {
			_ = cli.Push(t.Id + " Brear " + t.Content)
		}
	}
}
