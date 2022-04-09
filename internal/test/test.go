package test

import "time"

var tests []*Test

type Test struct {
	Id       string
	Content  string
	Start    time.Time
	Duration time.Duration
}

func New(test *Test) {
	tests = append(tests, test)
}

func Find(Id string) *Test {
	for _, test := range tests {
		if test.Id == Id {
			return test
		}
	}

	return nil
}

func Done(Id string) bool {
	for _, test := range tests {
		if test.Id == Id {
			test.Duration = test.Start.Sub(time.Now())

			return true
		}
	}

	return false
}
