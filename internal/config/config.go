package config

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/havij/internal/client"
	"github.com/amirhnajafiz/havij/internal/rabbit"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	Queue  string `koanf:"queue"`
	Prefix string `koanf:"prefix"`

	Timeout   int `koanf:"timeout"`
	Providers int `koanf:"number"`
	Consumers int `koanf:"consumers"`

	Client client.Config `koanf:"client"`
	Rabbit rabbit.Config `koanf:"rabbit"`
}

func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		_ = fmt.Errorf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		_ = fmt.Errorf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
