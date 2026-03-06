package config

import (
	"testing"
)

// TestLoad verifies that the Load function returns a valid Config with AppConfig.
func TestLoad(t *testing.T) {
	cfg := Load()

	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}

	// Test App configuration
	if cfg.App.Name != "caracal" {
		t.Errorf("App.Name = %v, want %v", cfg.App.Name, "caracal")
	}
	if cfg.App.Version != "0.1.0" {
		t.Errorf("App.Version = %v, want %v", cfg.App.Version, "0.1.0")
	}
	if cfg.App.Environment != "development" {
		t.Errorf("App.Environment = %v, want %v", cfg.App.Environment, "development")
	}
}

// TestGetAppInfo verifies that GetAppInfo returns correctly formatted string.
func TestGetAppInfo(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		wantStr string
	}{
		{
			name: "default config",
			cfg: &Config{
				App: AppConfig{
					Name:    "caracal",
					Version: "0.1.0",
				},
			},
			wantStr: "caracal v0.1.0",
		},
		{
			name: "custom version",
			cfg: &Config{
				App: AppConfig{
					Name:    "caracal",
					Version: "1.0.0",
				},
			},
			wantStr: "caracal v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cfg.GetAppInfo()
			if got != tt.wantStr {
				t.Errorf("GetAppInfo() = %v, want %v", got, tt.wantStr)
			}
		})
	}
}
