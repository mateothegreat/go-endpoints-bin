// Package types - provides reusable types for http responses.
package types

// User is the user type.
type User struct {
	ID       string `json:"id" faker:"uuid_hyphenated"`
	Name     string `json:"name" faker:"first_name"`
	Email    string `json:"email" faker:"email"`
	UserName string `json:"username" faker:"username"`
	Phone    string `json:"phone" faker:"phone_number"`
}
