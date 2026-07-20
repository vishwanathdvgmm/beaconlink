# 1406 - Go Code Style Guide

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the coding conventions and style guidelines used throughout the BeaconLink codebase.

Consistent code style improves:

- Readability
- Maintainability
- Code reviews
- Collaboration
- Long-term evolution

These guidelines apply to all Go source code in the repository.

---

# 2. Philosophy

BeaconLink follows idiomatic Go.

The project prioritizes:

- Clarity over cleverness
- Simplicity over abstraction
- Consistency over personal preference
- Explicitness over magic
- Readability over brevity

If a style decision conflicts with standard Go conventions, standard Go wins.

---

# 3. Formatting

All source code must be formatted using:

```bash
gofmt
```

Formatting is mandatory.

No manual formatting should override `gofmt`.

---

# 4. Import Formatting

Imports should be managed using:

```bash
goimports
```

Imports should be grouped into:

1. Standard library
2. Internal packages
3. External packages

Example:

```go
import (
    "context"
    "time"

    "github.com/beaconlink/beaconlink/internal/config"
    "github.com/beaconlink/beaconlink/internal/logging"

    "github.com/go-chi/chi/v5"
)
```

---

# 5. Linting

All code must pass:

```text
golangci-lint
```

Lint warnings should be treated as defects unless explicitly justified.

Repository lint configuration is authoritative.

---

# 6. File Organization

Files should group related responsibilities.

Example:

```text
deployment/

service.go
planner.go
validator.go
repository.go
errors.go
types.go
```

Avoid excessively large source files.

---

# 7. File Size

As guidance:

- Prefer fewer than 500 lines per file.
- Consider refactoring files exceeding 800 lines.
- Files exceeding 1,000 lines require architectural review.

These are guidelines rather than strict limits.

---

# 8. Function Size

Functions should perform one logical task.

General guidance:

- Prefer fewer than 50 lines.
- Consider refactoring functions exceeding 100 lines.

Extract helper functions when doing so improves readability.

---

# 9. Function Naming

Function names should describe behavior.

Good:

```go
LoadConfig()

Validate()

RegisterAgent()

Start()

Shutdown()
```

Avoid vague names:

```go
Handle()

Do()

Process()

ExecuteStuff()
```

---

# 10. Variable Naming

Variable names should be concise and meaningful.

Good:

```go
cfg

ctx

logger

deployment

agent
```

Avoid:

```go
temp

obj

thing

foo

bar
```

Single-letter variables are acceptable for short-lived loop indices and receivers.

---

# 11. Receiver Names

Receiver names should be short.

Preferred:

```go
func (s *Server)

func (r *Relay)

func (c *Client)
```

Avoid:

```go
func (server *Server)
```

Receiver names should remain consistent within a type.

---

# 12. Constants

Constants should describe meaning rather than values.

Preferred:

```go
const DefaultPort = 8080
```

Avoid:

```go
const Port = 8080
```

Magic numbers should be eliminated whenever practical.

---

# 13. Enumerations

Use typed constants.

Example:

```go
type Runtime string

const (
    RuntimeDocker Runtime = "docker"
    RuntimePodman Runtime = "podman"
    RuntimeNative Runtime = "native"
)
```

---

# 14. Comments

Code should explain _why_, not _what_.

Good comments document:

- Design decisions
- Constraints
- Trade-offs
- Non-obvious behavior

Avoid comments that merely restate the code.

Bad:

```go
// Increment i
i++
```

---

# 15. Package Documentation

Every exported package should include:

```text
doc.go
```

Package documentation should summarize the package's responsibility.

---

# 16. Exported Identifiers

All exported identifiers should have documentation comments.

Example:

```go
// Deployment manages deployment execution.
type Deployment struct{}
```

Documentation should begin with the identifier name.

---

# 17. Error Handling

Errors must be checked immediately.

Preferred:

```go
result, err := load()

if err != nil {
    return err
}
```

Avoid deferred or ignored error handling.

---

# 18. Error Messages

Error messages should:

- begin with lowercase
- avoid punctuation
- provide context

Preferred:

```go
return fmt.Errorf("loading deployment: %w", err)
```

---

# 19. Context Usage

`context.Context` should be the first parameter.

Example:

```go
func (s *Server) Deploy(
    ctx context.Context,
    req Request,
) error
```

Never store `context.Context` in structs.

---

# 20. Interfaces

Interfaces should be:

- Small
- Focused
- Behavior-oriented

Preferred:

```go
type Reader interface {
    Read(context.Context) error
}
```

Avoid large interfaces that combine unrelated behavior.

---

# 21. Constructors

Constructors should follow Go conventions.

Preferred:

```go
func New(...)
```

Avoid:

```go
Create()

Build()

Initialize()
```

---

# 22. Dependency Injection

Dependencies should be explicit.

Preferred:

```go
func New(
    cfg Config,
    logger *slog.Logger,
    store Store,
)
```

Avoid:

- Global variables
- Service locators
- Hidden dependencies

---

# 23. Global State

Global mutable state is prohibited.

Acceptable globals include:

- Constants
- Sentinel errors
- Immutable lookup tables

Everything else should be injected.

---

# 24. Panic

`panic` should be reserved for programmer errors or impossible states.

Operational failures should return errors.

Library packages must never use `panic` for normal error handling.

---

# 25. Logging

Always use structured logging.

Preferred:

```go
logger.Info(
    "deployment started",
    slog.String("deployment_id", id),
)
```

Never use:

```go
fmt.Println()

log.Printf()
```

for production code.

---

# 26. Control Flow

Prefer early returns.

Preferred:

```go
if err != nil {
    return err
}
```

Avoid deeply nested conditionals.

---

# 27. Boolean Expressions

Prefer positive conditions.

Good:

```go
if authenticated {
}
```

Less clear:

```go
if !unauthenticated {
}
```

Complex boolean expressions should be decomposed.

---

# 28. Switch Statements

Prefer `switch` over long `if-else` chains when evaluating multiple values.

Example:

```go
switch runtime {
case RuntimeDocker:
case RuntimePodman:
default:
}
```

---

# 29. Type Assertions

Avoid unchecked assertions.

Preferred:

```go
v, ok := value.(MyType)
if !ok {
    return ErrInvalidType
}
```

Unchecked assertions should only be used when correctness is guaranteed.

---

# 30. Goroutines

Every goroutine should have a defined lifecycle.

The creator is responsible for:

- Cancellation
- Cleanup
- Error propagation

Never launch unmanaged goroutines.

---

# 31. Channels

Channels should express ownership clearly.

General guidance:

- Sender closes the channel.
- Receiver never closes the channel.
- Buffered channels require justification.

Prefer simple synchronization primitives where sufficient.

---

# 32. Time

Use the `time` package consistently.

Avoid hard-coded durations.

Preferred:

```go
const DefaultTimeout = 30 * time.Second
```

---

# 33. Testing Style

Tests should:

- Be table-driven where appropriate
- Use descriptive names
- Verify observable behavior
- Avoid implementation details

Tests should be easy to read and maintain.

---

# 34. TODO Comments

TODOs should include sufficient context.

Preferred:

```go
// TODO(beaconlink): Support runtime hot reload.
```

Avoid anonymous TODO comments with no ownership or explanation.

---

# 35. Deprecated Code

Deprecated APIs should use Go's standard format.

Example:

```go
// Deprecated: Use NewRuntime instead.
```

Deprecated code should have a documented removal plan.

---

# 36. Generated Code

Generated files should not be edited manually.

Generated code should include the standard notice:

```go
// Code generated ... DO NOT EDIT.
```

---

# 37. Third-Party Packages

Use external packages only when they provide clear value.

New dependencies should:

- Be actively maintained
- Have permissive licenses
- Be well documented
- Avoid unnecessary transitive dependencies

Adding a new dependency should be justified during code review.

---

# 38. Readability

Readable code is preferred over clever code.

Avoid:

- Excessive chaining
- Deep nesting
- Reflection without necessity
- Premature optimization
- Unnecessary abstraction

Future maintainers should be able to understand code quickly.

---

# 39. Anti-Patterns

Avoid:

- Hungarian notation
- Global mutable state
- Package aliases without need
- Overly clever code
- Excessive comments
- Large functions
- Deep nesting
- Ignored errors
- Unchecked type assertions
- Hidden side effects
- Excessive use of `interface{}`
- Duplicate code caused by avoiding small abstractions

---

# 40. Code Review Checklist

Every review should verify:

- Code is formatted with `gofmt`.
- Imports are managed with `goimports`.
- Lint checks pass.
- Functions are cohesive.
- Errors are handled correctly.
- Dependencies are explicit.
- Public APIs are documented.
- Tests accompany new functionality.
- No unnecessary complexity has been introduced.

---

# 41. Compliance

All Go code committed to the BeaconLink repository must conform to this style guide. Code reviews and automated CI checks should enforce formatting, linting, documentation, and consistency requirements before changes are merged.

---

# 42. Summary

BeaconLink follows idiomatic Go coding practices with an emphasis on simplicity, readability, and consistency. Code should be automatically formatted, linted, well documented, and organized into small, cohesive functions and packages. Explicit dependencies, structured logging, robust error handling, and predictable control flow are preferred over clever abstractions or hidden behavior. These standards establish a uniform codebase that remains approachable, maintainable, and scalable as the project grows.
