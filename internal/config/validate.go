package config

import (
	"fmt"
	"strings"

	"github.com/vishwanathdvgmm/beaconlink/internal/logging"
)

// Validate verifies that the configuration is valid.
func (c Config) Validate() error {
	// Environment
	if strings.TrimSpace(c.Environment) == "" {
		return fmt.Errorf("%w: environment must not be empty", ErrInvalidConfig)
	}

	// Logging level
	if _, err := logging.ParseLevel(c.Logging.Level); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}

	// Logging format
	switch strings.ToLower(strings.TrimSpace(c.Logging.Format)) {
	case "text", "json":
		// valid

	default:
		return fmt.Errorf(
			"%w: logging format must be either \"text\" or \"json\"",
			ErrInvalidConfig,
		)
	}

	return nil
}
