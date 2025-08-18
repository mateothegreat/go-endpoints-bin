// Package config - configures the application.
package config

import (
	"log"
	"os"

	"github.com/mateothegreat/go-config/config"
	"github.com/mateothegreat/go-config/plugins/sources"
	"github.com/mateothegreat/go-multilog/multilog"
)

// PrometheusConfig - configures the prometheus server.
type PrometheusConfig struct {
	Enabled bool `yaml:"enabled" default:"true"`
	Port    int  `yaml:"port" default:"9100"`
}

// TracingConfig - configures the monitoring tracing.
type TracingConfig struct {
	Enabled   bool   `yaml:"enabled" default:"true"`
	Collector string `yaml:"collector" default:"jaeger"`
}

// MonitoringConfig - configures the monitoring.
type MonitoringConfig struct {
	Prometheus PrometheusConfig `yaml:"prometheus" default:"true"`
	Tracing    TracingConfig    `yaml:"tracing" default:"true"`
}

// HTTPConfig - configures the HTTP server.
type HTTPConfig struct {
	Enabled bool `yaml:"enabled" default:"true"`
	Port    int  `yaml:"port" default:"8080"`
}

// BaseConfig - configures the application.
type BaseConfig struct {
	HTTP       HTTPConfig       `yaml:"http"`
	Monitoring MonitoringConfig `yaml:"monitoring"`
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

	err = config.LoadWithPlugins(
		config.FromYAML(sources.YAMLOpts{Path: "/config.yaml"}),
		config.FromEnv(sources.EnvOpts{Prefix: "SIMPLE"}),
	).Build(cfg)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	multilog.Info("config", "setup", map[string]any{
		"config": cfg,
		"cwd":    cwd,
	})

	return cfg
}
