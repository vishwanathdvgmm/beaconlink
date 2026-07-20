# 1402 - Error Handling Guidelines

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the error handling philosophy, conventions, and best practices used throughout the BeaconLink codebase.

Consistent error handling improves:

- Reliability
- Observability
- Debuggability
- API consistency
- Operational visibility
- Maintainability

Every package should follow these guidelines when producing, wrapping, propagating, and handling errors.

---

# 2. Design Principles

BeaconLink adopts the standard Go error model.

Error handling should be:

- Explicit
- Predictable
- Contextual
- Observable
- Recoverable where appropriate

Errors are values and should be treated as part of the application's control flow.

---

# 3. Error Philosophy

Errors represent exceptional conditions.

They should:

- Describe what failed
- Preserve the underlying cause
- Include useful context
- Avoid exposing sensitive information

Errors should never be ignored.

---

# 4. Error Ownership

A package owns the errors it produces.

Example:

```go
package deployment

var ErrDeploymentNotFound = errors.New("deployment not found")
```

Packages should not define errors for other packages.

---

# 5. Sentinel Errors

Sentinel errors are appropriate when callers need to identify specific failure conditions.

Example:

```go
var (
    ErrNotFound      = errors.New("not found")
    ErrAlreadyExists = errors.New("already exists")
    ErrUnauthorized  = errors.New("unauthorized")
)
```

Sentinel errors should be used sparingly.

---

# 6. Typed Errors

When additional context is required, prefer typed errors.

Example:

```go
type ValidationError struct {
    Field string
    Reason string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}
```

Typed errors provide structured information without string parsing.

---

# 7. Error Wrapping

Always preserve the original cause.

Preferred:

```go
return fmt.Errorf("loading configuration: %w", err)
```

Avoid:

```go
return errors.New("configuration failed")
```

Wrapping enables callers to inspect the original error.

---

# 8. Error Inspection

Use the standard library.

Preferred:

```go
errors.Is(err, ErrNotFound)
```

```go
errors.As(err, &validationErr)
```

Avoid string comparisons.

Incorrect:

```go
if err.Error() == "not found" {
}
```

---

# 9. Error Context

Every layer should add meaningful context.

Example:

```text
open configuration
↓

parse configuration
↓

validate configuration
↓

start API server
```

The resulting error chain should describe the complete failure path.

---

# 10. Error Messages

Error messages should:

- begin with lowercase
- avoid punctuation
- describe the failure
- avoid redundancy

Good:

```text
failed to load configuration
```

Bad:

```text
Error: Failed To Load Configuration!
```

---

# 11. Sensitive Information

Errors must never expose:

- passwords
- API tokens
- certificates
- private keys
- secrets
- internal credentials

Sensitive information belongs only in secure storage.

---

# 12. API Errors

Public APIs should return structured error responses.

Example:

```json
{
  "error": {
    "code": "deployment_not_found",
    "message": "deployment does not exist",
    "request_id": "req-12345"
  }
}
```

Internal implementation details must not be exposed.

---

# 13. Error Codes

Externally visible APIs should expose stable error codes.

Examples:

```text
invalid_request

unauthorized

forbidden

deployment_not_found

runtime_unavailable

internal_error
```

Codes are stable.

Messages may evolve.

---

# 14. HTTP Status Mapping

Recommended mappings:

| Condition               | Status |
| ----------------------- | -----: |
| Validation failure      |    400 |
| Authentication required |    401 |
| Permission denied       |    403 |
| Resource not found      |    404 |
| Conflict                |    409 |
| Rate limited            |    429 |
| Internal failure        |    500 |
| Service unavailable     |    503 |

Status codes should accurately reflect the failure.

---

# 15. Logging Errors

Errors should be logged once.

Preferred:

```go
logger.Error(
    "deployment failed",
    slog.Any("error", err),
)
```

Avoid logging the same error at multiple layers.

---

# 16. Return vs Log

General rule:

- Produce error → return it.
- Boundary layer → log it.

Business logic should rarely log errors.

The application boundary owns logging.

---

# 17. Recoverable Errors

Recoverable failures should return errors.

Examples:

- network timeout
- configuration error
- validation failure
- unavailable runtime

The caller decides how to respond.

---

# 18. Fatal Errors

Fatal conditions belong only at application boundaries.

Examples:

```go
func main() {
    if err := run(); err != nil {
        logger.Error(...)
        os.Exit(1)
    }
}
```

Libraries must never terminate the process.

Never call:

```go
log.Fatal()
os.Exit()
panic()
```

inside reusable packages.

---

# 19. Panic Usage

`panic` is reserved for:

- impossible states
- programmer errors
- unrecoverable initialization failures

Examples:

- invalid generated code
- violated internal invariants

Operational failures should return errors.

---

# 20. Validation Errors

Validation failures are expected.

Example:

```go
if req.Name == "" {
    return ErrInvalidName
}
```

Validation errors should not be logged as system failures.

---

# 21. Context Cancellation

Context cancellation is not an error condition.

Example:

```go
if errors.Is(err, context.Canceled) {
    return err
}
```

Avoid logging expected cancellation events as failures.

---

# 22. Timeouts

Timeouts should be distinguishable.

Example:

```go
errors.Is(err, context.DeadlineExceeded)
```

Timeouts often indicate infrastructure issues rather than application bugs.

---

# 23. Retryable Errors

Retryable failures include:

- temporary network errors
- service unavailable
- timeout
- connection reset

Non-retryable failures include:

- validation errors
- authentication failures
- malformed requests

Retry behavior should be explicit.

---

# 24. Error Translation

Subsystems may translate low-level errors into domain errors.

Example:

```text
sql.ErrNoRows

↓

deployment.ErrNotFound
```

Translation should preserve the underlying cause.

---

# 25. Database Errors

Database implementations should not leak driver-specific errors.

Instead:

```text
pq.Error

↓

storage.ErrConflict
```

Business logic should remain database-agnostic.

---

# 26. External Dependencies

Do not expose third-party error types through public APIs.

Wrap external errors with BeaconLink domain errors where appropriate.

This reduces coupling to implementation details.

---

# 27. Concurrent Operations

Concurrent tasks should collect and aggregate errors.

Preferred approaches include:

- `errors.Join`
- explicit aggregation

Avoid silently discarding failures from goroutines.

---

# 28. Deferred Cleanup

Cleanup errors should be handled.

Example:

```go
defer func() {
    if err := file.Close(); err != nil {
        logger.Warn("failed to close file", slog.Any("error", err))
    }
}()
```

Cleanup failures should not hide the primary error.

---

# 29. Testing Errors

Tests should verify:

- returned error
- wrapped error
- sentinel matching
- typed error
- error message (only when appropriate)

Prefer:

```go
require.ErrorIs(t, err, ErrNotFound)
```

over string comparisons.

---

# 30. Documentation

Public packages should document:

- exported sentinel errors
- typed errors
- expected failure conditions
- retry behavior
- recovery expectations

Consumers should understand how to handle package errors.

---

# 31. Anti-Patterns

Avoid:

- Ignoring returned errors
- String comparison
- Logging every error at every layer
- Returning generic `errors.New("failed")`
- Exposing secrets
- Using `panic` for operational failures
- Calling `os.Exit()` from libraries
- Hiding the original cause
- Swallowing errors silently

---

# 32. Example Error Flow

```text
Database

↓

storage.ErrNotFound

↓

deployment.ErrDeploymentNotFound

↓

HTTP 404

↓

{
  "code": "deployment_not_found",
  "message": "deployment does not exist"
}
```

Each layer adds context while preserving the underlying cause.

---

# 33. Compliance

All Go packages within BeaconLink must follow these error handling guidelines. Code reviews should verify that errors are wrapped correctly, sensitive information is protected, logging occurs only at appropriate boundaries, and callers can reliably inspect failures using the standard `errors` package.

---

# 34. Summary

BeaconLink adopts idiomatic Go error handling based on explicit error returns, contextual wrapping, structured domain errors, and centralized logging at application boundaries. Errors should preserve their causes, communicate meaningful context, avoid exposing sensitive information, and remain easy to inspect using `errors.Is` and `errors.As`. Following these practices ensures consistent behavior across the platform while improving reliability, observability, and long-term maintainability.
