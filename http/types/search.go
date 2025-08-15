// Package types - provides reusable types for http responses.
package types

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Search is the search request type.
type Search struct {
	Terms *string `json:"terms,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Size  *int    `json:"size,omitempty"`
}

// NewSearch creates a new search request.
func NewSearch(c echo.Context) *Search {
	search := &Search{}

	if terms := c.QueryParam("terms"); terms != "" {
		search.Terms = &terms
	}

	if page, err := strconv.Atoi(c.QueryParam("page")); err == nil {
		search.Page = &page
	}

	if size, err := strconv.Atoi(c.QueryParam("size")); err == nil {
		search.Size = &size
	}

	return search
}
