package storage

import (
	"time"

	"github.com/amirhnajafiz/havij/internal/telemetry"
)

// storage is for holding tests.
type storage struct {
	metrics telemetry.Metrics
	tests   []*test
}

// test is a single message that
// is sent over rabbitMQ.
type test struct {
	Id       string
	Content  string
	Missed   bool
	Start    time.Time
	Duration time.Duration
}

// create a new test.
func (s *storage) create() *test {
	t := &test{
		Id:      time.Now().String(),
		Content: "string for content",
		Start:   time.Now(),
		Missed:  false,
	}

	s.tests = append(s.tests, t)

	return t
}

// find a test.
func (s *storage) find(Id string) bool {
	for _, t := range s.tests {
		if t.Id == Id {
			return true
		}
	}

	return false
}

// Done
// manage to finish a test.
func (s *storage) Done(id string, timeout int) (bool, time.Duration) {
	s.metrics.TotalReceive.Add(1)

	if !s.find(id) {
		return false, 0
	}

	for _, t := range s.tests {
		if t.Id == id {
			t.Duration = time.Now().Sub(t.Start)

			if t.Duration > time.Duration(timeout) {
				t.Missed = true

				s.metrics.TimeoutReceive.Add(1)
			}

			s.metrics.ReceiveTime.Observe(float64(t.Duration))

			return t.Missed, t.Duration
		}
	}

	return false, 0
}

// Generate
// manage to create a new test.
func (s *storage) Generate() (string, string) {
	t := s.create()

	s.metrics.TotalPublish.Add(1)

	return t.Id, t.Content
}

// NewStorage
// generates a new store.
func NewStorage() *storage {
	return &storage{
		metrics: telemetry.NewMetrics(),
	}
}
