package config

// Default returns the default BeaconLink configuration.
//
// These defaults are intended for local development and are applied
// before configuration files and environment variables are loaded.
//
// Precedence:
//
//	Defaults
//	    ↓
//	YAML Configuration
//	    ↓
//	Environment Variables
func Default() Config {
	return Config{
		Environment: "development",

		Logging: LoggingConfig{
			Level:     "info",
			Format:    "text",
			AddSource: true,
		},

		Version: VersionConfig{
			Enabled: true,
		},
	}
}
