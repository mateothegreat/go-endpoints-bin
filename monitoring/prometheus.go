package monitoring

import (
	"net/http"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// RequestsCounter is a counter for the number of requests.
	RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "rest_http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"path", "code", "method"})
)

// StartPrometheus starts the prometheus server.
func StartPrometheus() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9100", nil)

	reg := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
}

// AddPrometheusMiddleware adds the prometheus related middleware for echo server.
//
// Arguments:
// - next: The next handler to call.
//
// Returns:
// - The middleware function.
func AddPrometheusMiddleware() echo.MiddlewareFunc {
	conf := config.Setup()

	if conf.Monitoring.Prometheus.Enabled {
		return echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
			Namespace: "go-mock-endpoints",
			Subsystem: "http",
		})
	}

	return nil
}
