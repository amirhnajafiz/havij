package cmd

import (
	"log"
	"sync"

	"github.com/amirhnajafiz/carrot/internal/client"
	"github.com/amirhnajafiz/carrot/internal/config"
	"github.com/amirhnajafiz/carrot/internal/logger"
	"github.com/amirhnajafiz/carrot/internal/rabbit"
	"github.com/amirhnajafiz/carrot/internal/storage"
)

func Execute() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	c := config.Load()
	s := storage.NewStorage()

	if err := logger.CreateLogFile(); err != nil {
		panic(err)
	}

	{
		r, err := rabbit.Connect(c.Rabbit)
		if err != nil {
			log.Fatalf("Rabbit connection failed %v\n", err)
		}

		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
			Storage:    s,
		}

		go func() {
			err := cli.Subscribe(c.Timeout)
			if err != nil {
				panic(err)
			}

			wg.Done()
		}()
	}
	{
		r, err := rabbit.Connect(c.Rabbit)
		if err != nil {
			log.Fatalf("Rabbit connection failed %v\n", err)
		}

		cli := client.Client{
			Cfg:        c.Client,
			Connection: r,
			Queue:      c.Queue,
			Storage:    s,
		}

		if err := cli.Initialize(); err != nil {
			log.Fatalln(err)
		}

		for i := 0; i < c.Number; i++ {
			if err := cli.Publish(); err != nil {
				log.Println(err)
			}
		}
	}

	wg.Wait()
}
