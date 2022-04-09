package test

import (
	"time"

	"github.com/bxcodec/faker/v3"
)

var tests []*Test
var timeout int

type Test struct {
	Id      string
	Content string

	Missed bool

	Start    time.Time
	Duration time.Duration
}

func New(test *Test) {
	tests = append(tests, test)
}

func Find(Id string) bool {
	for _, test := range tests {
		if test.Id == Id {
			return true
		}
	}

	return false
}

func Done(Id string) bool {
	for _, test := range tests {
		if test.Id == Id {
			test.Duration = test.Start.Sub(time.Now())

			if test.Duration > time.Duration(timeout) {
				test.Missed = true
			}

			return true
		}
	}

	return false
}

func Generate(number int) []*Test {
	for i := 0; i < number; i++ {
		temp := Test{
			Id:      time.UnixMilli(int64(i)).String(),
			Content: faker.Sentence(), // random text
			Start:   time.Now(),
			Missed:  false,
		}

		New(&temp)
	}

	return tests
}
