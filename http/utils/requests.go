// Package utils - provides the utilities for handling requests.
package utils

import (
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mateothegreat/go-mock-endpoints/http/types"
	"github.com/mateothegreat/go-multilog/multilog"
)

// NewRequest takes in the echo context and returns a new request object.
//
// Arguments:
// - c: The echo context.
//
// Returns:
// - The new request object.
func NewRequest(c echo.Context) *types.Request {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		multilog.Error("paths", "get", map[string]any{
			"error": err,
		})
	}

	request := &types.Request{
		URL:        c.Request().URL.String(),
		Method:     c.Request().Method,
		Header:     c.Request().Header,
		Body:       string(body),
		Params:     strings.Split(c.Request().URL.Path, "/")[1:],
		Query:      c.QueryParams(),
		Form:       c.Request().Form,
		Cookies:    c.Request().Cookies(),
		RemoteAddr: c.Request().RemoteAddr,
		UserAgent:  c.Request().UserAgent(),
		Referer:    c.Request().Referer(),
		Host:       c.Request().Host,
		Protocol:   c.Request().Proto,
		Scheme:     c.Request().URL.Scheme,
	}
	return request
}
