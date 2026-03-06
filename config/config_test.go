package config

import (
	"testing"
)

// TestLoad verifies that the Load function returns a valid Config with AppConfig.
// This test ensures default values are properly initialized.
func TestLoad(t *testing.T) {
	cfg := Load()

	if cfg == nil {
		t.Fatal("Load() returned nil config")
	}

	// Test App configuration - verify default values
	if cfg.App.Name != "caracal" {
		t.Errorf("App.Name = %v, want %v", cfg.App.Name, "caracal")
	}
	if cfg.App.Version != "0.1.0" {
		t.Errorf("App.Version = %v, want %v", cfg.App.Version, "0.1.0")
	}
	// Environment should default to "dev" when no CARACAL_ENV is set
	if cfg.App.Environment == "" {
		t.Errorf("App.Environment is empty, want non-empty value")
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

// BenchmarkLoad benchmarks the Load function.
func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load()
	}
}

// BenchmarkLoadBaseConfig benchmarks loading the base config file.
func BenchmarkLoadBaseConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loadBaseConfig()
	}
}

// BenchmarkMergeConfigs benchmarks the mergeConfigs function.
func BenchmarkMergeConfigs(b *testing.B) {
	baseConfig := &Config{
		App: AppConfig{
			Name:        "caracal",
			Version:     "0.1.0",
			Description: "Caracal - A high-performance Go application",
			Environment: "development",
		},
	}
	envConfig := &Config{
		App: AppConfig{
			Name:        "caracal",
			Version:     "1.0.0",
			Description: "Caracal - Updated",
			Environment: "production",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergeConfigs(baseConfig, envConfig)
	}
}
