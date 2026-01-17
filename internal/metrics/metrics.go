package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	// Tracks total HTTP requests by path and method
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)

	// Tracks request duration in seconds
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

// Register all metrics
func Register() {
	prometheus.MustRegister(HttpRequestsTotal, HttpRequestDuration)
}
