package telemetry

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}
