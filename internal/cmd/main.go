package cmd

import (
	"log"
	"time"

	"github.com/amirhnajafiz/carrot/internal/client"
	"github.com/amirhnajafiz/carrot/internal/config"
	"github.com/amirhnajafiz/carrot/internal/logger"
	"github.com/amirhnajafiz/carrot/internal/rabbit"
	"github.com/amirhnajafiz/carrot/internal/storage"
	"github.com/amirhnajafiz/carrot/internal/telemetry"
)

func Execute() {
	c := config.Load()
	s := storage.NewStorage()

	if err := logger.CreateLogFile(); err != nil {
		panic(err)
	}

	// initialize queue
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
			panic(err)
		}

		_ = r.Close()
	}

	// consumers
	for i := 0; i < c.Consumers; i++ {
		consumer(s, c)
	}

	// providers
	for i := 0; i < c.Providers; i++ {
		provider(s, c)
	}
}

func consumer(s client.Storage, cfg config.Config) {
	r, err := rabbit.Connect(cfg.Rabbit)
	if err != nil {
		log.Fatalf("Rabbit connection failed %v\n", err)
	}

	cli := client.Client{
		Cfg:        cfg.Client,
		Connection: r,
		Queue:      cfg.Queue,
		Storage:    s,
	}

	go func() {
		if err := cli.Subscribe(cfg.Timeout); err != nil {
			panic(err)
		}
	}()
}

func provider(s client.Storage, cfg config.Config) {
	r, err := rabbit.Connect(cfg.Rabbit)
	if err != nil {
		log.Fatalf("Rabbit connection failed %v\n", err)
	}

	cli := client.Client{
		Cfg:        cfg.Client,
		Connection: r,
		Queue:      cfg.Queue,
		Storage:    s,
	}

	go func() {
		for {
			if err := cli.Publish(); err != nil {
				break
			}

			time.Sleep(2 * time.Second)
		}
	}()

	// starting prometheus server
	telemetry.NewServer()
}
