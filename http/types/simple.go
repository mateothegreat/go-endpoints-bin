// Package types - provides reusable types for http responses.
package types

// Simple is but a simple type.
type Simple struct {
	Foo string `json:"foo" faker:"paragraph"`
}
