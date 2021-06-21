package common

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"code", "path"})

	HttpRequestsErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_err_total",
		Help: "Count of all err HTTP requests",
	}, []string{"code", "path"})

	HttpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of all HTTP requests",
	}, []string{"code", "path"})
)


func Registry()  {
	prometheus.MustRegister(HttpRequestDuration, HttpRequestsErrorTotal, HttpRequestsTotal)

}