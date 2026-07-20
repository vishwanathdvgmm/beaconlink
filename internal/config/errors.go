package config

import "errors"

var (
	// ErrInvalidConfig indicates that the configuration failed validation.
	ErrInvalidConfig = errors.New("invalid configuration")

	// ErrConfigNotFound indicates that the configuration file could not be found.
	ErrConfigNotFound = errors.New("configuration file not found")

	// ErrConfigLoad indicates that the configuration could not be loaded.
	ErrConfigLoad = errors.New("failed to load configuration")
)
