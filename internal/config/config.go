package config

// Config is the root BeaconLink configuration.
//
// All application configuration should be represented here.
// Loading, defaults, validation, and environment overrides are
// implemented elsewhere in this package.
type Config struct {
	// Environment identifies the current runtime environment.
	//
	// Examples:
	//   development
	//   staging
	//   production
	Environment string `koanf:"environment" yaml:"environment"`

	// Logging controls application logging.
	Logging LoggingConfig `koanf:"logging" yaml:"logging"`

	// Version controls version command behavior.
	Version VersionConfig `koanf:"version" yaml:"version"`
}

// LoggingConfig defines structured logging behavior.
type LoggingConfig struct {
	// Level is the minimum log level.
	//
	// Supported values:
	//   debug
	//   info
	//   warn
	//   error
	Level string `koanf:"level" yaml:"level"`

	// Format specifies the output formatter.
	//
	// Supported values:
	//   text
	//   json
	Format string `koanf:"format" yaml:"format"`

	// AddSource includes source file and line information in log output.
	//
	// This is typically enabled during development and disabled in
	// production.
	AddSource bool `koanf:"add_source" yaml:"add_source"`
}

// VersionConfig controls version information output.
type VersionConfig struct {
	// Enabled determines whether version information is available
	// through the CLI.
	Enabled bool `koanf:"enabled" yaml:"enabled"`
}
