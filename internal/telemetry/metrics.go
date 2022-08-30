package telemetry

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	TotalPublish   prometheus.Counter
	TotalReceive   prometheus.Counter
	TimeoutReceive prometheus.Counter
	ReceiveTime    prometheus.Summary
}

const (
	namespace = "carrot"
	subsystem = "carrot"
)

func NewMetrics() Metrics {
	return Metrics{
		TotalPublish: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "total_publish",
		}),
		TotalReceive: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "total_receives",
		}),
		TimeoutReceive: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "timeout_receives",
		}),
		ReceiveTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "receive_time",
		}),
	}
}
