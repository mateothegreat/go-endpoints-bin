// Package types - provides types for the application.
package types

import (
	"net/http"
	"net/url"
)

// Request contains the request data.
type Request struct {
	URL        string         `json:"url"`
	Method     string         `json:"method"`
	Header     http.Header    `json:"header,omitempty"`
	Body       string         `json:"body,omitempty"`
	Params     []string       `json:"params,omitempty"`
	Query      url.Values     `json:"query,omitempty"`
	Form       url.Values     `json:"form,omitempty"`
	Cookies    []*http.Cookie `json:"cookies,omitempty"`
	RemoteAddr string         `json:"remote_addr,omitempty"`
	UserAgent  string         `json:"user_agent,omitempty"`
	Referer    string         `json:"referer,omitempty"`
	Host       string         `json:"host,omitempty"`
	Protocol   string         `json:"protocol,omitempty"`
	Scheme     string         `json:"scheme,omitempty"`
	Search     *Search        `json:"search,omitempty"`
}

// Response contains the request and response data.
//
// Arguments:
// - TRequest: The type of the request.
// - TResponse: The type of the response.
type Response[TRequest any, TResponse any] struct {
	Request  Request `json:"request"`
	Response any     `json:"response"`
}
