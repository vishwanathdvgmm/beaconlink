package config

// loader contains internal configuration loading options.
//
// It is intentionally unexported so callers interact only through
// the Option type and Load() API.
type loader struct {
	envPrefix string
}

// Option configures the behavior of the configuration loader.
type Option func(*loader)

// defaultLoader returns the default loader configuration.
func defaultLoader() *loader {
	return &loader{
		envPrefix: "BEACONLINK",
	}
}

// WithEnvPrefix overrides the default environment variable prefix.
//
// Example:
//
//	BEACONLINK_LOGGING_LEVEL=debug
//
// can become:
//
//	MYAPP_LOGGING_LEVEL=debug
func WithEnvPrefix(prefix string) Option {
	return func(l *loader) {
		if prefix != "" {
			l.envPrefix = prefix
		}
	}
}
