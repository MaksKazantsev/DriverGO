package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type Metrics interface {
	Increment(statusCode int, method string)
	Latency(t float64, path string)
}

type metrics struct {
	totalRequest    *prometheus.CounterVec
	requestsLatency *prometheus.HistogramVec
}

func (m metrics) Latency(t float64, path string) {
	m.requestsLatency.WithLabelValues(path).Observe(t)
}

func (m metrics) Increment(statusCode int, method string) {
	m.totalRequest.WithLabelValues(http.StatusText(statusCode), method).Inc()
}

var _ Metrics = &metrics{}

func NewMetrics(colls *[]prometheus.Collector) *metrics {
	totalReq := prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "driverGO", Name: "request_counter", Help: "Total number of HTTP requests", Subsystem: "http"}, []string{"status_code", "method"})
	latency := prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "driverGO", Name: "requests_latency", Help: "Measure latency of requests", Buckets: []float64{0.15, 0.2, 0.3, 0.5}}, []string{"requests"})

	*colls = append(*colls, totalReq, latency)

	return &metrics{totalRequest: totalReq, requestsLatency: latency}
}
