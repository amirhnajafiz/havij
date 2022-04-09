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

			return true
		}
	}

	return false
}

func Generate(number int) []*Test {
	for i := 0; i < number; i++ {
		temp := Test{
			Id:      time.UnixMilli(int64(i)).String(),
			Content: "", // random text
			Start:   time.Now(),
		}

		New(&temp)
	}

	return tests
}
