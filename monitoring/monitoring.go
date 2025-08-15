// Package monitoring - provides monitoring for the application.
package monitoring

import (
	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/config"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Start starts the monitoring services.
func Start() {
	conf := config.Setup()

	// Start the prometheus server if enabled.
	if conf.Monitoring.Prometheus.Enabled {
		multilog.Info("monitoring", "starting prometheus...", map[string]any{
			"config": conf.Monitoring.Prometheus,
		})
		go StartPrometheus()
	}

	// Start the tracing service if enabled.
	if conf.Monitoring.Tracing.Enabled {
		multilog.Info("monitoring", "starting tracing...", map[string]any{
			"config": conf.Monitoring.Tracing,
		})
		go StartTracer()
	}
}

// GetMiddleware - gets the monitoring related middleware for echo server.
//
// Arguments:
// - next: The next handler to call.
//
// Returns:
// - The middleware function.
func GetMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	conf := config.Setup()

	if conf.Monitoring.Tracing.Enabled {
		return GetOtelMiddleware(next)
	}

	return next
}
