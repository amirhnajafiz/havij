package logger

import (
	"log"
	"os"
)

const (
	// logs file name.
	name = "logs.txt"
)

// CreateLogFile
// saves the logs into a file.
func CreateLogFile() error {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(f)

	return nil
}
