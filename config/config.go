// Package config provides application configuration management.
//
// This package centralizes all configuration settings for the caracal application,
// including version information, application name, and other runtime settings.
//
// Example usage:
//
//	cfg := config.Load()
//	fmt.Printf("Running %s v%s\n", cfg.App.Name, cfg.App.Version)
package config

// Config holds all application configuration settings.
type Config struct {
	App AppConfig `yaml:"app"`
}

// AppConfig holds application metadata and settings.
type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Environment string `yaml:"environment"`
}

// Load returns a Config struct with all application settings.
// This function loads configuration from environment variables or default values.
func Load() *Config {
	return &Config{
		App: AppConfig{
			Name:        "caracal",
			Version:     "0.1.0",
			Description: "Caracal - A high-performance Go application",
			Environment: "development",
		},
	}
}

// GetAppInfo returns a formatted string with the application name and version.
func (c *Config) GetAppInfo() string {
	return c.App.Name + " v" + c.App.Version
}
