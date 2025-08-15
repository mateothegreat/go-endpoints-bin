// Package http - provides the http server for the endpoints service.
package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mateothegreat/go-mock-endpoints/config"
	"github.com/mateothegreat/go-mock-endpoints/http/paths"
	"github.com/mateothegreat/go-mock-endpoints/monitoring"
)

// Start - starts the http server.
func Start() error {
	conf := config.Setup()

	e := echo.New()
	e.Use(monitoring.GetMiddleware)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
			"http://localhost:*",
			"https://bin.matthewdavis.io",
		},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderXRequestedWith,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	if conf.Monitoring.Prometheus.Enabled {
		e.Use(monitoring.AddPrometheusMiddleware())
	}

	paths.Router(e)

	return e.Start(fmt.Sprintf(":%d", conf.HTTP.Port))
}
