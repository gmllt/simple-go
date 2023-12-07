package web

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	namespace      = "simple_go"
	gTotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests.",
		},
		[]string{"path"},
	)

	gResponseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_response_status",
			Help:      "Total number of HTTP response status.",
		},
		[]string{"path", "status"},
	)

	gHTTPDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request duration distribution.",
		},
		[]string{"path"},
	)
)

func init() {
	prometheus.MustRegister(gTotalRequests)
	prometheus.MustRegister(gResponseStatus)
}
