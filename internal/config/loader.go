package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Load loads a BeaconLink configuration from disk.
//
// Loading precedence:
//
//	Default values
//	    ↓
//	YAML configuration file
//	    ↓
//	Environment variables
func Load(path string, opts ...Option) (Config, error) {
	l := defaultLoader()

	for _, opt := range opts {
		opt(l)
	}

	// Initialize the configuration with default values.
	// Values loaded from the YAML file and environment variables
	// overwrite these defaults during unmarshalling.
	cfg := Default()

	k := koanf.New(".")

	// Load configuration from the YAML file.
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("%w: %s", ErrConfigNotFound, path)
		}

		return Config{}, fmt.Errorf("%w: %w", ErrConfigLoad, err)
	}

	// Load configuration from environment variables.
	if err := k.Load(
		env.Provider(
			l.envPrefix+"_",
			".",
			func(s string) string {
				s = strings.TrimPrefix(s, l.envPrefix+"_")
				s = strings.ToLower(s)
				s = strings.ReplaceAll(s, "_", ".")
				return s
			},
		),
		nil,
	); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrConfigLoad, err)
	}

	// Merge loaded values into the default configuration.
	if err := k.Unmarshal("", &cfg); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrConfigLoad, err)
	}

	// Validate the final configuration.
	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
