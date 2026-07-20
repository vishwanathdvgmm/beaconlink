package logging

import (
	"fmt"
	"log/slog"
	"strings"
)

// ParseLevel converts a string into a slog.Level.
func ParseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug, nil

	case "info":
		return slog.LevelInfo, nil

	case "warn", "warning":
		return slog.LevelWarn, nil

	case "error":
		return slog.LevelError, nil

	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level %q", level)
	}
}
