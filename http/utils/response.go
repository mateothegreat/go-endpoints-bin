// Package utils - provides the utilities for handling responses.
package utils

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

// GenericResponse is a generic response type.
type GenericResponse struct {
	Message string `json:"message" faker:"paragraph"`
}

// NewResponse takes in the response object and returns a new response object.
//
// Arguments:
// - input: The response object.
//
// Returns:
// - The new response object.
func NewResponse[T any](input *T) *T {
	err := faker.FakeData(input, options.WithTagName("faker"))
	if err != nil {
		fmt.Println(err)
	}

	return input
}
