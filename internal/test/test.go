package test

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

var tests []*Test

type Test struct {
	Id      string
	Content string

	Missed bool

	Start    time.Time
	Duration time.Duration
}

func create(test *Test) {
	tests = append(tests, test)
}

func find(Id string) bool {
	for _, test := range tests {
		if test.Id == Id {
			return true
		}
	}

	return false
}

func Done(Id string, timeout int) (bool, time.Duration) {
	if !find(Id) {
		return false, 0
	}

	for _, test := range tests {
		if test.Id == Id {
			test.Duration = time.Now().Sub(test.Start)

			if test.Duration > time.Duration(timeout) {
				test.Missed = true
			}

			return test.Missed, test.Duration
		}
	}

	return false, 0
}

func Generate(number int) []*Test {
	for i := 0; i < number; i++ {
		create(&Test{
			Id:      time.UnixMilli(int64(i)).String(),
			Content: faker.Sentence(), // random text
			Start:   time.Now(),
			Missed:  false,
		})
	}

	return tests
}
