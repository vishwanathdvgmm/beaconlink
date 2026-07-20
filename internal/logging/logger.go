package logging

import (
	"io"
	"log/slog"
	"os"
)

// New creates a configured logger.
func New(cfg Config) (*slog.Logger, error) {

	level, err := ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	var handler slog.Handler

	switch cfg.Format {

	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)

	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler), nil
}

// NewWithWriter is primarily intended for testing.
func NewWithWriter(cfg Config, w io.Writer) (*slog.Logger, error) {

	level, err := ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	var handler slog.Handler

	switch cfg.Format {

	case "json":
		handler = slog.NewJSONHandler(w, opts)

	default:
		handler = slog.NewTextHandler(w, opts)
	}

	return slog.New(handler), nil
}
