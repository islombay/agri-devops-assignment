package http

import (
	"net/http"
	"time"

	"github.com/islombay/agri-devops-assignment/internal/metrics"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call next handler
		next.ServeHTTP(w, r)

		// Record metrics
		duration := time.Since(start).Seconds()
		metrics.HttpRequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()
		metrics.HttpRequestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
	})
}
