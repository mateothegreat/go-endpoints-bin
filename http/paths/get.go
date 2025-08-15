// Package paths - provides the GET http path handler.
package paths

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/http/types"
	"github.com/mateothegreat/go-mock-endpoints/http/utils"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Get handles the GET request.
//
// Arguments:
// - c: The echo context.
//
// Returns:
// - The error.
func Get(c echo.Context) error {
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
