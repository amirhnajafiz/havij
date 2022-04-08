package test

import "time"

var tests []*Test

type Test struct {
	Id      string
	Content string
	Start   time.Time
}

func New(test *Test) {
	tests = append(tests, test)
}

func Get(Id string) *Test {
	for _, test := range tests {
		if test.Id == Id {
			return test
		}
	}

	return nil
}
