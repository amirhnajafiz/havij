package cmd

import (
	"log"
	"sync"

	"github.com/amirhnajafiz/carrot/internal/client"
	"github.com/amirhnajafiz/carrot/internal/config"
	"github.com/amirhnajafiz/carrot/internal/logger"
	"github.com/amirhnajafiz/carrot/internal/rabbitMQT"
	"github.com/amirhnajafiz/carrot/internal/test"
)

func Execute() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	c := config.Load()
	tests := test.Generate(c.Test.Number)

	logger.CreateLogFile("logs.txt")
	log.Println("start testing")

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

		go func() {
			err := cli.Listen(c.Test.Timeout)
			if err != nil {
				panic(err)
			}

			wg.Done()
		}()
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
			err := cli.Push(t.Id + c.Prefix + t.Content)
			if err != nil {
				panic(err)
			}
		}
	}

	wg.Wait()
}
