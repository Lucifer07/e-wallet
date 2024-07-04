package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
            Name: "http_request_total",
            Help: "http request total",
        },
        []string{"path", "method","status"},
	)
	HttpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
            Help:    "http request duration",
            Buckets: prometheus.DefBuckets,
        },
		[]string{"path", "method","status"},
	)

	)