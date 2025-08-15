// Package config - configures the application.
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mateothegreat/go-config/config"
	"github.com/mateothegreat/go-config/plugins/sources"
	"github.com/mateothegreat/go-multilog/multilog"
)

// PrometheusConfig - configures the prometheus server.
type PrometheusConfig struct {
	Enabled bool `yaml:"enabled" env:"ENABLED"`
	Port    int  `yaml:"port" env:"PORT" env-default:"9100"`
}

// TracingConfig - configures the monitoring tracing.
type TracingConfig struct {
	Enabled   bool   `yaml:"enabled" env:"ENABLED"`
	Collector string `yaml:"collector" env:"COLLECTOR"`
}

// MonitoringConfig - configures the monitoring.
type MonitoringConfig struct {
	Prometheus PrometheusConfig `yaml:"prometheus" env-prefix:"PROMETHEUS_"`
	Tracing    TracingConfig    `yaml:"tracing" env-prefix:"TRACING_"`
}

// HTTPConfig - configures the HTTP server.
type HTTPConfig struct {
	Enabled bool `yaml:"enabled" env:"ENABLED"`
	Port    int  `yaml:"port" env:"PORT" env-default:"8080"`
}

// BaseConfig - configures the application.
type BaseConfig struct {
	HTTP       HTTPConfig       `yaml:"http" env-prefix:"HTTP_"`
	Monitoring MonitoringConfig `yaml:"monitoring" env-prefix:"MONITORING_"`
}

// cfg - the configuration.
var cfg *BaseConfig

// Setup - sets up the configuration.
func Setup() *BaseConfig {
	if cfg != nil {
		return cfg
	}

	cfg = &BaseConfig{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	filelisting, err := os.ReadDir(cwd + "/..")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range filelisting {
		fmt.Println(file.Name())
	}

	multilog.Info("config", "setup", map[string]any{
		"config":      cfg,
		"cwd":         cwd,
		"filelisting": filelisting,
	})

	err = config.LoadWithPlugins(
		config.FromYAML(sources.YAMLOpts{Path: cwd + "/config.yaml"}),
		config.FromEnv(sources.EnvOpts{Prefix: "SIMPLE"}),
	).Build(cfg)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	multilog.Info("config", "setup", map[string]any{
		"config":      cfg,
		"cwd":         cwd,
		"filelisting": filelisting,
	})

	return cfg
}
