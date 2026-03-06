// Package config provides application configuration management.
//
// This package centralizes all configuration settings for the caracal application,
// including version information, application name, and other runtime settings.
//
// Example usage:
//
//	cfg := config.Load()
//	fmt.Printf("Running %s v%s\n", cfg.AppName, cfg.Version)
package config

// Config holds all application configuration settings.
type Config struct {
	// AppName is the name of the application.
	AppName string

	// Version is the application version.
	Version string

	// Environment specifies the runtime environment (development, staging, production).
	Environment string

	// LogLevel specifies the logging level (debug, info, warn, error).
	LogLevel string
}

// Load returns a Config struct with all application settings.
// This function loads configuration from environment variables or default values.
func Load() *Config {
	return &Config{
		AppName:     "caracal",
		Version:     "0.1.0",
		Environment: "development",
		LogLevel:    "info",
	}
}

// GetAppInfo returns a formatted string with application name and version.
func (c *Config) GetAppInfo() string {
	return c.AppName + " v" + c.Version
}
