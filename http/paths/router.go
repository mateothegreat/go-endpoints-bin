// Package paths - provides the paths router for the endpoints service.
package paths

import (
	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/monitoring"
)

// Router - the echo router for the paths.
//
// Arguments:
// - e: the echo instance.
//
// Returns:
// - the echo group.
func Router(e *echo.Echo) *echo.Group {
	g := e.Group("*", monitoring.GetMiddleware)

	g.GET("", Get)
	g.POST("", Post)
	g.GET("/search", Search)
	// g.GET("/:id", Get)
	// g.DELETE("/:id", Delete)
	// g.PUT("/:id/information", UpdateInformation)

	return g
}
