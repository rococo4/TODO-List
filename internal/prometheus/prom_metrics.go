package meth

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RequestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests.",
})
var HttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_duration_seconds",
	Help:    "Duration of HTTP requests.",
	Buckets: prometheus.DefBuckets,
},
	[]string{"method", "status_code"})
