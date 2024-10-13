package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)

	HttpRequestPath = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_path",
			Help: "HTTP requests with path",
		},
		[]string{"path"},
	)

	HttpErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"path", "method", "status"},
	)

	// PROMQL => rate(http_request_duration_seconds_sum{}[5m]) / rate(http_request_duration_seconds_count{}[5m])
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Average response time of HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(HttpRequestCount)
	prometheus.MustRegister(HttpRequestPath)
	prometheus.MustRegister(HttpErrorCount)
	prometheus.MustRegister(HttpRequestDuration)
}
