package logger

import (
	"log"
	"os"
)

func CreateLogFile(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(f)
}
