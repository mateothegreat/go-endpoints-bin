// Package paths - provides the POST http path handler.
package paths

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/http/types"
	"github.com/mateothegreat/go-mock-endpoints/http/utils"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Post handles the POST request.
//
// Arguments:
// - c: The echo context.
//
// Returns:
// - The error.
func Post(c echo.Context) error {
	// id := c.Param("id")
	multilog.Info("paths", "get", map[string]any{
		"path": c.Path(),
	})

	multilog.Info("paths", "path", map[string]any{
		"path":   c.Path(),
		"params": c.ParamNames(),
		"values": c.ParamValues(),
	})

	response := &types.Response[any, types.Simple]{
		Request:  *utils.NewRequest(c),
		Response: utils.NewResponse(&types.Simple{}),
	}

	err := c.JSON(http.StatusOK, response)
	if err != nil {
		multilog.Error("paths", "get", map[string]any{
			"error": err,
		})
	}

	return err
}
