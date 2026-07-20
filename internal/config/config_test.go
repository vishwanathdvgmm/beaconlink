package config

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestDefault(t *testing.T) {
	cfg := Default()

	if cfg.Environment != "development" {
		t.Fatalf("expected environment %q, got %q", "development", cfg.Environment)
	}

	if cfg.Logging.Level != "info" {
		t.Fatalf("expected log level %q, got %q", "info", cfg.Logging.Level)
	}

	if cfg.Logging.Format != "text" {
		t.Fatalf("expected log format %q, got %q", "text", cfg.Logging.Format)
	}

	if !cfg.Logging.AddSource {
		t.Fatal("expected AddSource to be true")
	}

	if !cfg.Version.Enabled {
		t.Fatal("expected Version.Enabled to be true")
	}
}

func TestLoadValidConfig(t *testing.T) {
	path := filepath.Join("testdata", "valid.yaml")

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Environment != "development" {
		t.Fatalf("expected environment %q, got %q", "development", cfg.Environment)
	}

	if cfg.Logging.Level != "debug" {
		t.Fatalf("expected log level %q, got %q", "debug", cfg.Logging.Level)
	}

	if cfg.Logging.Format != "json" {
		t.Fatalf("expected format %q, got %q", "json", cfg.Logging.Format)
	}

	if cfg.Logging.AddSource {
		t.Fatal("expected AddSource to be false")
	}
}

func TestLoadInvalidConfig(t *testing.T) {
	path := filepath.Join("testdata", "invalid.yaml")

	_, err := Load(path)
	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, ErrInvalidConfig) {
		t.Fatalf("expected ErrInvalidConfig, got %v", err)
	}
}

func TestLoadMissingConfig(t *testing.T) {
	_, err := Load("does-not-exist.yaml")
	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, ErrConfigNotFound) {
		t.Fatalf("expected ErrConfigNotFound, got %v", err)
	}
}

func TestEnvironmentOverride(t *testing.T) {
	t.Setenv("BEACONLINK_LOGGING_LEVEL", "error")

	path := filepath.Join("testdata", "valid.yaml")

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Logging.Level != "error" {
		t.Fatalf("expected environment override %q, got %q", "error", cfg.Logging.Level)
	}
}

func TestWithEnvPrefix(t *testing.T) {
	t.Setenv("TESTAPP_LOGGING_LEVEL", "warn")

	path := filepath.Join("testdata", "valid.yaml")

	cfg, err := Load(path, WithEnvPrefix("TESTAPP"))
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.Logging.Level != "warn" {
		t.Fatalf("expected log level %q, got %q", "warn", cfg.Logging.Level)
	}
}

func TestValidateDefault(t *testing.T) {
	cfg := Default()

	if err := cfg.Validate(); err != nil {
		t.Fatalf("default configuration should be valid: %v", err)
	}
}

func TestLoadInvalidPath(t *testing.T) {
	path := filepath.Join("testdata", "missing.yaml")

	_, err := Load(path)
	if err == nil {
		t.Fatal("expected error")
	}

	if !os.IsNotExist(err) && !errors.Is(err, ErrConfigNotFound) {
		t.Fatalf("unexpected error: %v", err)
	}
}
