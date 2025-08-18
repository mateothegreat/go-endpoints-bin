// Package main - the main package for the application.
package main

import (
	"fmt"

	"github.com/mateothegreat/go-mock-endpoints/config"
	"github.com/mateothegreat/go-mock-endpoints/http"
	"github.com/mateothegreat/go-mock-endpoints/monitoring"
	"github.com/mateothegreat/go-mock-endpoints/setup"
	"github.com/mateothegreat/go-multilog/multilog"
)

func main() {
	conf := config.Setup()
	_ = setup.Setup()
	fmt.Printf("conf: %+v\n", conf)
	monitoring.Start()

	if conf.HTTP.Enabled {
		if err := http.Start(); err != nil {
			multilog.Fatal("main", "failed to start http server", map[string]any{
				"error": err,
			})
		}
	}

}
