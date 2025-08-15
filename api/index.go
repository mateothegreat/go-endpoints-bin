package api

import (
	_http "net/http"

	"github.com/mateothegreat/go-mock-endpoints/http"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Handler is the handler if running in vercel functions.
//
// Arguments:
// - w: The response writer.
// - r: The request.
func Handler(w _http.ResponseWriter, r *_http.Request) {
	multilog.Info("main", "handler", map[string]any{
		"path":   r.URL.Path,
		"method": r.Method,
	})

	e := http.Create()
	e.ServeHTTP(w, r)
}
