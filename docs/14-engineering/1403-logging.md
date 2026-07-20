# 1403 - Logging Standards

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the logging philosophy, conventions, implementation guidelines, and operational practices used throughout the BeaconLink platform.

Logging serves several critical purposes:

- Operational monitoring
- Troubleshooting
- Incident response
- Auditing
- Performance analysis
- Security investigations

All BeaconLink components must produce consistent, structured, machine-readable logs.

---

# 2. Design Principles

BeaconLink logging follows these principles:

- Structured by default
- Human-readable in development
- Machine-readable in production
- Context-rich
- Low overhead
- Consistent across all services

Logging should aid operations without becoming noise.

---

# 3. Logging Library

The official logging implementation is Go's standard library:

```go
log/slog
```

No alternative logging framework should be introduced unless approved through an Architecture Decision Record (ADR).

---

# 4. Centralized Logging Package

All services must use:

```text
internal/logging
```

Packages must never configure loggers directly.

The logging package is responsible for:

- Logger creation
- Formatting
- Output configuration
- Log levels
- Context propagation

---

# 5. Logger Construction

Services should receive loggers through dependency injection.

Preferred:

```go
logger := logging.New(cfg)
```

Avoid:

```go
log.Default()
```

or

```go
slog.New(...)
```

inside application packages.

---

# 6. Logger Ownership

Each component owns its own logger instance.

Examples:

```text
API Server

Relay

Agent

Deployment Engine

Runtime

CLI
```

Child components should derive loggers from their parent.

---

# 7. Structured Logging

All logs must use structured fields.

Preferred:

```go
logger.Info(
    "agent registered",
    slog.String("agent_id", id),
    slog.String("relay", relayID),
)
```

Avoid:

```go
logger.Info("Agent " + id + " registered")
```

Structured fields enable efficient searching and aggregation.

---

# 8. Log Levels

BeaconLink defines five logging levels.

| Level | Purpose                                            |
| ----- | -------------------------------------------------- |
| DEBUG | Diagnostic information                             |
| INFO  | Normal operation                                   |
| WARN  | Recoverable issues                                 |
| ERROR | Failed operations                                  |
| FATAL | Process termination (application entry point only) |

Application packages must never terminate the process directly.

---

# 9. DEBUG

Used for:

- Development
- Diagnostics
- Internal state
- Protocol tracing
- Temporary troubleshooting

DEBUG logs should normally be disabled in production.

---

# 10. INFO

Used for significant operational events.

Examples:

- Service started
- Service stopped
- Agent connected
- Deployment created
- Configuration loaded
- Certificate renewed

INFO should describe expected system behavior.

---

# 11. WARN

WARN indicates abnormal but recoverable situations.

Examples:

- Retry initiated
- Slow response
- Configuration fallback
- Temporary network failure
- Certificate nearing expiration

WARN does not indicate application failure.

---

# 12. ERROR

ERROR indicates failed operations.

Examples:

- Deployment failed
- Authentication failed
- Runtime unavailable
- Database unavailable
- Relay connection lost

Errors should include sufficient context for diagnosis.

---

# 13. FATAL

Fatal logging is reserved exclusively for application entry points.

Example:

```go
func main() {
    if err := run(); err != nil {
        logger.Error("startup failed", slog.Any("error", err))
        os.Exit(1)
    }
}
```

Library packages must never call:

- `os.Exit()`
- `log.Fatal()`
- `panic()` for operational failures

---

# 14. Required Fields

Every log entry should include enough context to identify:

- Component
- Operation
- Outcome

Typical fields include:

```text
component
operation
request_id
trace_id
duration
error
```

---

# 15. Component Field

Every logger should include:

```text
component
```

Examples:

```text
api

relay

agent

deployment

runtime

security
```

This should be attached when the logger is created.

---

# 16. Correlation IDs

Request-scoped logs should include:

- request_id
- trace_id
- span_id (when available)

This enables end-to-end tracing across services.

---

# 17. Common Fields

Frequently used fields include:

```text
component

service

version

request_id

trace_id

agent_id

deployment_id

runtime

host

duration

error
```

Use consistent names across all services.

---

# 18. Message Style

Messages should:

- be concise
- describe completed or attempted actions
- avoid punctuation
- begin with lowercase

Good:

```text
deployment completed

relay connected

certificate renewed
```

Avoid:

```text
Deployment Completed Successfully!

ERROR: Deployment Failed
```

---

# 19. Errors

Errors should be logged as structured values.

Preferred:

```go
logger.Error(
    "deployment failed",
    slog.Any("error", err),
)
```

Avoid embedding errors inside strings.

---

# 20. Sensitive Information

Logs must never contain:

- Passwords
- API keys
- JWTs
- Bearer tokens
- Certificates
- Private keys
- Secrets
- Session tokens

Sensitive values must always be redacted.

---

# 21. Performance Logging

Long-running operations should record duration.

Example:

```go
logger.Info(
    "deployment completed",
    slog.Duration("duration", elapsed),
)
```

Performance metrics complement—but do not replace—Prometheus metrics.

---

# 22. Startup Logging

Startup should log:

- Version
- Commit
- Build date
- Configuration source
- Runtime
- Listening address

This information aids operational support.

---

# 23. Shutdown Logging

Shutdown should log:

- Shutdown initiated
- Active connections
- Cleanup completion
- Total shutdown duration

Graceful shutdown should be observable.

---

# 24. HTTP Logging

Each request should capture:

- Method
- Path
- Status code
- Duration
- Request ID
- Remote address
- User agent (optional)

Bodies should not be logged by default.

---

# 25. Deployment Logging

Deployment operations should include:

- Deployment ID
- Version
- Target agent
- Runtime
- Duration
- Result

Deployment history should remain easy to reconstruct.

---

# 26. Security Logging

Security events should include:

- Login attempts
- Authentication failures
- Authorization failures
- Certificate rotation
- Token expiration
- Administrative actions

Security logs should also generate audit events where applicable.

---

# 27. Audit Logging

Audit logs differ from operational logs.

Audit events should be:

- Immutable
- Timestamped
- Complete
- Attributable

Audit logging requirements are defined separately in the security documentation.

---

# 28. Log Volume

Avoid excessive logging inside:

- Tight loops
- Heartbeats
- High-frequency polling
- Packet processing

Repeated events should be summarized where practical.

---

# 29. Configuration

Logging configuration should support:

- Level
- Format
- Output destination

Example:

```yaml
logging:
  level: info
  format: json
  output: stdout
```

Configuration is managed centrally by the `config` package.

---

# 30. Output Formats

Development:

```text
Text
```

Production:

```text
JSON
```

JSON is required for integration with centralized logging systems.

---

# 31. Observability

Logging complements:

- Metrics
- Traces
- Health checks

Logs should not duplicate information already provided more effectively through metrics or tracing.

---

# 32. Testing

Tests should verify:

- Correct log level
- Required structured fields
- Absence of sensitive data
- Expected log messages where behavior depends on logging

Logging itself should not become the primary assertion unless necessary.

---

# 33. Anti-Patterns

Avoid:

- String interpolation instead of structured fields
- Logging secrets
- Duplicate logging across layers
- Excessive DEBUG logging
- Logging every successful loop iteration
- Using logs instead of metrics
- Using `fmt.Println()` or `log.Printf()` in production code
- Creating loggers inside individual packages

---

# 34. Example

```go
logger.Info(
    "deployment completed",
    slog.String("component", "deployment"),
    slog.String("deployment_id", id),
    slog.String("agent_id", agentID),
    slog.Duration("duration", elapsed),
)
```

This example demonstrates the preferred use of structured fields and contextual information.

---

# 35. Compliance

All BeaconLink services must use the centralized logging package and follow these standards. Code reviews should verify that logging is structured, context-rich, free of sensitive information, and uses consistent field names and log levels across the platform.

---

# 36. Summary

BeaconLink standardizes on Go's `log/slog` package to provide structured, contextual, and consistent logging across all services. Loggers are centrally configured, injected into components, and produce machine-readable output suitable for production observability. By enforcing common field names, appropriate log levels, and strict handling of sensitive information, the platform delivers logs that support debugging, monitoring, auditing, and incident response while remaining predictable and maintainable.
