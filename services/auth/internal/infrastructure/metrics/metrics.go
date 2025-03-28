package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// HTTP metrics
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// Authentication metrics
	loginAttempts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "login_attempts_total",
			Help: "Total number of login attempts",
		},
		[]string{"status"},
	)

	// Token metrics
	tokenOperations = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "token_operations_total",
			Help: "Total number of token operations",
		},
		[]string{"operation", "status"},
	)

	// Database metrics
	dbOperations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_operation_duration_seconds",
			Help:    "Database operation duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(loginAttempts)
	prometheus.MustRegister(tokenOperations)
	prometheus.MustRegister(dbOperations)
}

// HTTP metrics
func RecordHTTPRequest(method, path string, status int, duration time.Duration) {
	httpRequestsTotal.WithLabelValues(method, path, string(status)).Inc()
	httpRequestDuration.WithLabelValues(method, path).Observe(duration.Seconds())
}

// Authentication metrics
func RecordLoginAttempt(success bool) {
	status := "failure"
	if success {
		status = "success"
	}
	loginAttempts.WithLabelValues(status).Inc()
}

// Token metrics
func RecordTokenOperation(operation string, success bool) {
	status := "failure"
	if success {
		status = "success"
	}
	tokenOperations.WithLabelValues(operation, status).Inc()
}

// Database metrics
func RecordDBOperation(operation string, duration time.Duration) {
	dbOperations.WithLabelValues(operation).Observe(duration.Seconds())
}
