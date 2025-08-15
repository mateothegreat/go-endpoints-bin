// Package setup - sets up the logger and the configuration.
package setup

import "github.com/mateothegreat/go-multilog/multilog"

// Setup sets up the logger and the configuration.
func Setup() error {
	multilog.RegisterLogger(
		multilog.LogMethod("console"),
		multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
			Level:  multilog.DEBUG,
			Format: multilog.FormatText,
		}),
	)

	return nil
}
