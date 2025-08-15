// Package paths - provides the search (GET) http path handler.
package paths

import (
	"fmt"
	"net/http"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/http/types"
	"github.com/mateothegreat/go-mock-endpoints/http/utils"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Search handles search (GET) requests.
//
// Arguments:
// - c: The echo context.
//
// Returns:
// - The error.
func Search(c echo.Context) error {
	search := types.NewSearch(c)
	request := utils.NewRequest(c)
	request.Search = search

	results := []types.User{}

	if search.Size == nil {
		size := 3
		search.Size = &size
	}

	for i := 0; i < *search.Size; i++ {
		user := &types.User{}
		err := faker.FakeData(user, options.WithTagName("faker"))
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, *user)
	}

	response := &types.Response[any, []types.User]{
		Request:  *request,
		Response: results,
	}

	err := c.JSON(http.StatusOK, response)
	if err != nil {
		multilog.Error("paths", "get", map[string]any{
			"error": err,
		})
	}

	return err
}
