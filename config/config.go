// Package config provides application configuration management.
//
// This package centralizes all configuration settings for the caracal application,
// including version information, application name, and other runtime settings.
//
// The configuration loading strategy:
// 1. Load base config.yml as default
// 2. Load environment-specific config (config.dev.yml, config.staging.yml, config.prod.yml)
// 3. Environment-specific values override base values
//
// Example usage:
//
//	cfg := config.Load()
//	fmt.Printf("Running %s v%s\n", cfg.App.Name, cfg.App.Version)
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

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

// Load returns a Config struct by loading base config.yml and overriding with environment-specific config.
// It reads config.yml first (base defaults), then loads config.{environment}.yml and merges values.
// If no environment is specified, defaults to "development".
func Load() *Config {
	// Start with the default config
	cfg := &Config{
		App: AppConfig{
			Name:        "caracal",
			Version:     "0.1.0",
			Description: "Caracal - A high-performance Go application",
			Environment: "dev",
		},
	}

	// Load base config
	baseConfig := loadBaseConfig()
	if baseConfig != nil {
		cfg = baseConfig
	}

	// Determine environment from env var or use config value
	env := os.Getenv("CARACAL_ENV")
	if env == "" {
		env = cfg.App.Environment
		if env == "" {
			env = "dev"
		}
	}

	// Update environment in config
	cfg.App.Environment = env

	// Load and merge environment-specific config
	envConfig := loadEnvironmentConfig(env)
	if envConfig != nil {
		cfg = mergeConfigs(cfg, envConfig)
	}

	return cfg
}

// loadBaseConfig loads the base config.yml file.
func loadBaseConfig() *Config {
	configPath := filepath.Join("config", "config.yml")
	return loadConfigFile(configPath)
}

// loadEnvironmentConfig loads the environment-specific config file (e.g., config.dev.yml).
func loadEnvironmentConfig(environment string) *Config {
	configPath := filepath.Join("config", fmt.Sprintf("config.%s.yml", environment))
	return loadConfigFile(configPath)
}

// loadConfigFile loads a YAML config file and returns a Config struct.
// Uses gopkg.in/yaml.v3 for robust YAML parsing.
func loadConfigFile(filePath string) *Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil
	}

	// Only return config if at least one field was set
	if cfg.App.Name != "" || cfg.App.Version != "" || cfg.App.Environment != "" {
		return cfg
	}

	return nil
}

// mergeConfigs merges environment-specific config into base config.
// Non-empty values from envConfig override baseConfig values.
func mergeConfigs(baseConfig, envConfig *Config) *Config {
	if envConfig.App.Name != "" {
		baseConfig.App.Name = envConfig.App.Name
	}
	if envConfig.App.Version != "" {
		baseConfig.App.Version = envConfig.App.Version
	}
	if envConfig.App.Description != "" {
		baseConfig.App.Description = envConfig.App.Description
	}
	if envConfig.App.Environment != "" {
		baseConfig.App.Environment = envConfig.App.Environment
	}

	return baseConfig
}

// GetAppInfo returns a formatted string with the application name and version.
func (c *Config) GetAppInfo() string {
	return c.App.Name + " v" + c.App.Version
}
