package cmd

import (
	"log"

	"github.com/amirhnajafiz/playful-rabbit/internal/client"
	"github.com/amirhnajafiz/playful-rabbit/internal/config"
	"github.com/amirhnajafiz/playful-rabbit/internal/logger"
	"github.com/amirhnajafiz/playful-rabbit/internal/rabbitMQT"
	"github.com/amirhnajafiz/playful-rabbit/internal/test"
)

func Execute() {
	c := config.Load()
	tests := test.Generate(c.Test.Number)

	logger.CreateLogFile("logs.txt")

	{
		r, err := rabbitMQT.Init(c.Rabbit)
		if err != nil {
			log.Fatalf("Rabbit connection failed %v\n", err)
		}

		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		_ = cli.Listen(c.Test.Timeout)
	}
	{
		r, err := rabbitMQT.Init(c.Rabbit)
		if err != nil {
			log.Fatalf("Rabbit connection failed %v\n", err)
		}

		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		for _, t := range tests {
			_ = cli.Push(t.Id + c.Prefix + t.Content)
		}
	}
}
