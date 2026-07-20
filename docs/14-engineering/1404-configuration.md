# 1404 - Configuration Management

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the configuration philosophy, loading process, validation rules, and management practices used throughout the BeaconLink platform.

A consistent configuration system improves:

- Deployment consistency
- Operational reliability
- Environment portability
- Security
- Testability
- Maintainability

Every BeaconLink service must use the centralized configuration package.

---

# 2. Design Principles

Configuration should be:

- Explicit
- Typed
- Validated
- Centralized
- Immutable after startup
- Environment-aware

Configuration is considered part of the application's public operational interface.

---

# 3. Configuration Package

All configuration management is implemented under:

```text
internal/config
```

This package is responsible for:

- Loading
- Merging
- Validation
- Defaults
- Environment overrides
- Configuration decoding

No other package should implement its own configuration loader.

---

# 4. Configuration Library

BeaconLink standardizes on:

```text
koanf
```

Koanf provides:

- YAML support
- Environment variables
- Multiple providers
- Nested configuration
- Type-safe decoding

No additional configuration libraries should be introduced without an approved ADR.

---

# 5. Configuration Sources

Configuration is loaded in the following order:

1. Built-in defaults
2. Configuration file
3. Environment variables
4. CLI flags

Each source overrides the previous one.

```text
Defaults
      │
      ▼
Configuration File
      │
      ▼
Environment Variables
      │
      ▼
CLI Flags
```

---

# 6. Configuration Files

Primary configuration format:

```yaml
server:
  address: ":8080"

logging:
  level: info

security:
  tls: true
```

YAML is the canonical configuration format.

---

# 7. File Naming

Recommended file names:

```text
development.yaml

production.yaml

agent.yaml

relay.yaml

api.yaml
```

Configuration files should clearly indicate their purpose.

---

# 8. Default Location

Services should search for configuration in the following order:

```text
./config.yaml

./configs/config.yaml

/etc/beaconlink/config.yaml
```

A custom path may be supplied using CLI flags.

---

# 9. Environment Variables

Environment variables override file values.

Naming convention:

```text
BEACONLINK_SERVER_ADDRESS

BEACONLINK_LOGGING_LEVEL

BEACONLINK_DATABASE_HOST
```

Rules:

- Uppercase
- Underscore-separated
- Common prefix
- Reflect configuration hierarchy

---

# 10. CLI Flags

CLI flags have the highest precedence.

Example:

```bash
beacon api \
    --config production.yaml \
    --logging.level debug \
    --server.address :9090
```

Flags should mirror the configuration hierarchy whenever practical.

---

# 11. Configuration Structure

Configuration should be organized by subsystem.

Example:

```yaml
server:

logging:

database:

redis:

relay:

agent:

runtime:

security:

metrics:

telemetry:
```

Avoid large, flat configuration files.

---

# 12. Configuration Types

All configuration should be represented by typed Go structures.

Example:

```go
type Config struct {
    Server   ServerConfig
    Logging  LoggingConfig
    Security SecurityConfig
}
```

Avoid dynamic maps for application configuration.

---

# 13. Defaults

Reasonable defaults should exist whenever possible.

Example:

```yaml
logging:
  level: info
```

Required values without sensible defaults should fail validation.

---

# 14. Required Configuration

Examples of required values:

- Database credentials
- TLS certificates
- Private keys
- API secrets

Application startup should fail if required configuration is missing.

---

# 15. Validation

Every configuration object must be validated after loading.

Validation includes:

- Required fields
- Port ranges
- URL syntax
- File existence
- Duration values
- Numeric ranges
- Enum values

Applications must never start with invalid configuration.

---

# 16. Immutability

Configuration becomes immutable after startup.

Packages should receive configuration through constructors.

Configuration should not be modified at runtime unless explicitly supported.

---

# 17. Dependency Injection

Configuration should be injected.

Preferred:

```go
service := deployment.New(cfg, logger)
```

Avoid:

```go
config.Get()
```

or package-level global configuration.

---

# 18. Secrets

Secrets include:

- Passwords
- JWT signing keys
- API tokens
- TLS private keys
- Database credentials

Secrets must never:

- Be committed
- Be logged
- Be printed
- Be exposed through APIs

---

# 19. Secret Sources

Production secrets should originate from:

- Environment variables
- Secret files
- External secret managers (future)

Configuration files should not contain production secrets.

---

# 20. Logging Configuration

Logging configuration includes:

```yaml
logging:
  level: info
  format: json
  output: stdout
```

Supported levels:

- debug
- info
- warn
- error

---

# 21. Server Configuration

Typical server configuration:

```yaml
server:
  address: ":8080"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 60s
```

Timeouts should always be explicitly configurable.

---

# 22. Security Configuration

Example:

```yaml
security:
  tls: true
  cert_file: cert.pem
  key_file: key.pem
```

Security defaults should favor secure operation.

---

# 23. Runtime Configuration

Example:

```yaml
runtime:
  default: docker
```

Runtime-specific options belong under dedicated subsections.

---

# 24. Metrics Configuration

Example:

```yaml
metrics:
  enabled: true
  address: ":9091"
```

Metrics should be configurable independently of the API server.

---

# 25. Telemetry Configuration

Example:

```yaml
telemetry:
  tracing: true
  exporter: otlp
```

Tracing configuration should remain isolated from logging configuration.

---

# 26. Environment Profiles

Recommended profiles:

```text
development

testing

staging

production
```

Profiles should primarily influence default values rather than application behavior.

---

# 27. Hot Reloading

Configuration is static for v1.0.

Runtime configuration reloading is not supported unless explicitly implemented.

Future support requires:

- Validation
- Atomic replacement
- Rollback strategy

---

# 28. Error Handling

Configuration errors should clearly identify:

- Source
- Field
- Reason

Example:

```text
invalid configuration: server.address: invalid port
```

Error messages should enable rapid diagnosis.

---

# 29. Testing

Configuration tests should verify:

- Default loading
- File loading
- Environment overrides
- CLI overrides
- Validation failures
- Missing required values
- Invalid types

Every configuration package should include unit tests.

---

# 30. Documentation

Every configuration option should document:

- Purpose
- Type
- Default value
- Whether it is required
- Valid range or accepted values

Configuration changes require corresponding documentation updates.

---

# 31. Anti-Patterns

Avoid:

- Global configuration
- Reading environment variables throughout the codebase
- Runtime mutation of configuration
- Unvalidated configuration
- Logging secrets
- Large flat configuration files
- Magic defaults without documentation
- Duplicate configuration loaders

---

# 32. Example Directory

```text
configs/

development.yaml

production.yaml

agent.yaml

relay.yaml

api.yaml
```

Example package layout:

```text
internal/config/

config.go

defaults.go

loader.go

validator.go

types.go

config_test.go
```

---

# 33. Compliance

All BeaconLink services must use the centralized `internal/config` package for loading, validating, and accessing configuration. Code reviews should verify that configuration is strongly typed, validated before use, injected through constructors, and never read directly from environment variables outside the configuration package.

---

# 34. Summary

BeaconLink centralizes all configuration management within a dedicated package built on Koanf. Configuration is loaded from layered sources, validated before startup, represented by strongly typed structures, and injected into services rather than accessed globally. By enforcing immutable, well-documented, and security-conscious configuration practices, the platform achieves predictable deployments, consistent behavior across environments, and a maintainable operational model.
