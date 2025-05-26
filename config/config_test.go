package config

import (
	"os"
	"testing"
)

func TestLoadServiceConfig_WithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVICE_HOST", "127.0.0.1")
	os.Setenv("SERVICE_PORT", "9090")

	defer func() {
		// Clean up after test
		os.Unsetenv("SERVICE_HOST")
		os.Unsetenv("SERVICE_PORT")
	}()

	cfg := LoadServiceConfig()

	if cfg.Server.Host != "127.0.0.1" {
		t.Errorf("Expected Host to be '127.0.0.1', got '%s'", cfg.Server.Host)
	}
	if cfg.Server.Port != "9090" {
		t.Errorf("Expected Port to be '9090', got '%s'", cfg.Server.Port)
	}
	if cfg.Server.ReadTimeout != 10 {
		t.Errorf("Expected ReadTimeout to be 10, got %d", cfg.Server.ReadTimeout)
	}
	if cfg.Server.WriteTimeout != 10 {
		t.Errorf("Expected WriteTimeout to be 10, got %d", cfg.Server.WriteTimeout)
	}
	if cfg.Server.IdleTimeout != 60 {
		t.Errorf("Expected IdleTimeout to be 60, got %d", cfg.Server.IdleTimeout)
	}
}

func TestLoadServiceConfig_WithDefaults(t *testing.T) {
	// Ensure environment variables are not set
	os.Unsetenv("SERVICE_HOST")
	os.Unsetenv("SERVICE_PORT")

	cfg := LoadServiceConfig()

	if cfg.Server.Host != "0.0.0.0" {
		t.Errorf("Expected default Host to be '0.0.0.0', got '%s'", cfg.Server.Host)
	}
	if cfg.Server.Port != "8088" {
		t.Errorf("Expected default Port to be '8088', got '%s'", cfg.Server.Port)
	}
}
