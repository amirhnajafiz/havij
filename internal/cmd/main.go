package cmd

import (
	"fmt"
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
	log.Println("start testing")

	r, err := rabbitMQT.Init(c.Rabbit)
	if err != nil {
		log.Fatalf("Rabbit connection failed %v\n", err)
	}

	{
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		go cli.Listen(c.Test.Timeout)
	}
	{
		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
		}

		for _, t := range tests {
			err := cli.Push(t.Id + c.Prefix + t.Content)
			if err != nil {
				_ = fmt.Errorf(err.Error())
			}
		}
	}
}
