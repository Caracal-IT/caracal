package config

import (
	"strings"
	"testing"
)

// TestLoad verifies that the Load function returns a valid Config.
func TestLoad(t *testing.T) {
	cfg := Load()

	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}

	if cfg.AppName != "caracal" {
		t.Errorf("Load() AppName = %v, want %v", cfg.AppName, "caracal")
	}

	if cfg.Version != "0.1.0" {
		t.Errorf("Load() Version = %v, want %v", cfg.Version, "0.1.0")
	}

	if cfg.Environment != "development" {
		t.Errorf("Load() Environment = %v, want %v", cfg.Environment, "development")
	}

	if cfg.LogLevel != "info" {
		t.Errorf("Load() LogLevel = %v, want %v", cfg.LogLevel, "info")
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
				AppName: "caracal",
				Version: "0.1.0",
			},
			wantStr: "caracal v0.1.0",
		},
		{
			name: "custom version",
			cfg: &Config{
				AppName: "caracal",
				Version: "1.0.0",
			},
			wantStr: "caracal v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cfg.GetAppInfo()
			if !strings.Contains(got, tt.wantStr) {
				t.Errorf("GetAppInfo() = %v, want substring %v", got, tt.wantStr)
			}
		})
	}
}
