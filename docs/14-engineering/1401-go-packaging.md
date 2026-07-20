# 1401 - Go Packaging Guidelines

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the Go package design principles, naming conventions, visibility rules, dependency boundaries, and architectural practices used throughout the BeaconLink codebase.

Consistent package design improves:

- Maintainability
- Testability
- Discoverability
- API clarity
- Dependency management
- Long-term scalability

These guidelines apply to every Go package in the repository.

---

# 2. Design Philosophy

BeaconLink follows standard Go idioms while emphasizing modular architecture.

Package design should prioritize:

- Simplicity
- Single responsibility
- Explicit dependencies
- Small public APIs
- High cohesion
- Low coupling

The repository should feel familiar to experienced Go developers.

---

# 3. Package Ownership

Every package owns exactly one responsibility.

Examples:

| Package    | Responsibility         |
| ---------- | ---------------------- |
| config     | Configuration loading  |
| logging    | Logging infrastructure |
| protocol   | BeaconLink Protocol    |
| relay      | Relay service          |
| deployment | Deployment engine      |
| runtime    | Runtime abstraction    |
| security   | Security services      |

A package should never become a collection of unrelated functionality.

---

# 4. Package Layout

Typical package:

```text
deployment/

doc.go

types.go

service.go

repository.go

validator.go

errors.go

deployment_test.go
```

Files should be grouped by responsibility rather than size.

---

# 5. Package Naming

Package names must be:

- lowercase
- singular
- descriptive
- short

Examples:

```text
agent
relay
runtime
logging
config
metrics
storage
```

Avoid:

```text
utilities
helpers
common
misc
shared
manager
serviceutils
```

---

# 6. One Package, One Concept

Each package should answer a single question.

Example:

```text
runtime

├── discovery
├── lifecycle
├── health
```

All belong to runtime because they relate to runtime management.

Do not mix unrelated concerns.

---

# 7. Package Size

Large packages are difficult to maintain.

When a package exceeds roughly:

- 8–15 source files
- multiple unrelated domains
- several independent abstractions

consider splitting it into subpackages.

Example:

```text
security/

authentication/
authorization/
certificates/
audit/
```

---

# 8. Package Visibility

Go visibility rules should be used intentionally.

Public:

```go
type Deployment struct {}

func New() {}
```

Private:

```go
type deployment struct {}

func validate() {}
```

Only export identifiers that form the package API.

Everything else remains private.

---

# 9. Public APIs

Every exported identifier should satisfy one of the following:

- Required by another package
- Required by tests
- Part of the package contract

Avoid exporting implementation details.

---

# 10. Constructors

Constructors should follow Go conventions.

Preferred:

```go
func New(...) *Service
```

Alternative:

```go
func NewManager(...)
```

Avoid:

```go
CreateService()
BuildService()
InitService()
```

---

# 11. Package Documentation

Every exported package should contain:

```text
doc.go
```

Example:

```go
// Package runtime provides the runtime abstraction
// layer used by BeaconLink.
package runtime
```

Public packages require package documentation.

---

# 12. Internal Packages

Everything not intended for external reuse belongs under:

```text
internal/
```

External projects should never depend on internal implementation.

---

# 13. pkg/

Only stable APIs belong in:

```text
pkg/
```

Candidates include:

```text
pkg/blp
pkg/sdk
```

These packages require:

- Documentation
- Versioning
- Backward compatibility

---

# 14. Avoid Generic Utility Packages

Do not create:

```text
util
utils
common
shared
helpers
```

Instead:

```text
validation
logging
config
errors
metrics
```

Packages should communicate intent.

---

# 15. Interfaces

Interfaces belong where they are consumed.

Correct:

```go
type Store interface {
    Save(context.Context, Deployment) error
}
```

Incorrect:

```text
interfaces/
```

Avoid dedicated interface packages.

---

# 16. Interface Design

Interfaces should be:

- small
- focused
- behavior-oriented

Preferred:

```go
type Reader interface {
    Read(context.Context) error
}
```

Avoid "god interfaces."

---

# 17. Concrete Types

Return concrete types unless abstraction is necessary.

Preferred:

```go
func New() *Relay
```

Avoid:

```go
func New() RelayInterface
```

Return interfaces only when required.

---

# 18. Dependency Injection

Dependencies should be explicit.

Preferred:

```go
func New(
    logger *Logger,
    store Store,
    cfg Config,
)
```

Avoid:

- globals
- service locators
- package initialization

---

# 19. Package Initialization

Avoid:

```go
func init() {}
```

Except for:

- registration
- unavoidable library initialization

Initialization should be explicit.

---

# 20. Configuration

Packages should never read configuration directly.

Incorrect:

```go
os.Getenv(...)
```

Preferred:

```go
func New(cfg Config)
```

Configuration enters through constructors.

---

# 21. Logging

Packages should receive loggers.

Incorrect:

```go
log.Println(...)
```

Preferred:

```go
type Service struct {
    logger *slog.Logger
}
```

Logging configuration belongs elsewhere.

---

# 22. Error Ownership

Errors belong to the package that produces them.

Example:

```go
var ErrDeploymentNotFound = errors.New(...)
```

Avoid central error registries.

---

# 23. Data Models

Keep models close to their owner.

Preferred:

```text
deployment/

deployment.go
```

Avoid:

```text
models/
entities/
types/
```

shared across unrelated domains.

---

# 24. Cross-Package Communication

Packages communicate through:

- interfaces
- method calls
- events
- messages

Avoid exposing internal implementation.

---

# 25. Cyclic Dependencies

Circular imports are prohibited.

If encountered:

- introduce interfaces
- move abstractions
- reorganize packages

Never bypass Go's import restrictions.

---

# 26. Context Usage

Public methods that perform:

- I/O
- network operations
- database operations
- long-running work

should accept:

```go
context.Context
```

Context should be the first parameter.

Example:

```go
func (s *Service) Deploy(
    ctx context.Context,
    req Request,
) error
```

---

# 27. Goroutines

Packages that start goroutines should also own:

- cancellation
- cleanup
- shutdown

Never leak goroutines.

---

# 28. Concurrency

Shared mutable state should be minimized.

Synchronization primitives should remain internal.

Avoid exposing:

- mutexes
- channels
- wait groups

through public APIs.

---

# 29. Package APIs

Public APIs should remain:

- stable
- minimal
- unsurprising

Avoid exposing third-party library types unless necessary.

Example:

Prefer:

```go
type Config struct {}
```

instead of exposing library-specific configuration structures.

---

# 30. Imports

Imports should be grouped:

```go
import (
    "context"
    "errors"
    "time"

    "github.com/beaconlink/beaconlink/internal/config"
    "github.com/beaconlink/beaconlink/internal/logging"

    "github.com/go-chi/chi/v5"
)
```

Order:

1. Standard library
2. Internal packages
3. External packages

Imports should be formatted with `go fmt`.

---

# 31. Package Tests

Tests belong beside production code.

```text
runtime/

runtime.go
runtime_test.go
```

External integration tests belong under:

```text
test/
```

---

# 32. Mocking

Prefer interfaces over mocking frameworks.

Example:

```go
type Store interface {}
```

Use lightweight test implementations.

Avoid excessive generated mocks unless they provide significant value.

---

# 33. Package Evolution

Public packages evolve carefully.

Breaking API changes require:

- ADR (if architectural)
- RFC (if externally visible)
- Documentation updates
- Versioning considerations

Internal packages may evolve more freely but should still preserve clean boundaries.

---

# 34. Anti-Patterns

Avoid:

- God packages
- Global state
- Deep package hierarchies
- Hidden dependencies
- Generic helper packages
- Circular imports
- Large exported APIs
- Excessive abstractions
- Premature interfaces

---

# 35. Example Package Tree

```text
internal/

config/
logging/
metrics/
telemetry/
security/

protocol/

relay/
    session.go
    registry.go
    router.go

agent/
    daemon.go
    heartbeat.go

deployment/
    service.go
    planner.go
    executor.go

runtime/
    docker/
    podman/
    native/

storage/
    postgres/
    redis/

version/
```

---

# 36. Compliance

All Go packages introduced into the BeaconLink repository must follow these packaging guidelines unless an approved Architecture Decision Record explicitly defines an exception.

Code reviews should verify adherence to these conventions before changes are merged.

---

# 37. Summary

BeaconLink adopts idiomatic Go packaging centered on single-responsibility packages, explicit dependency injection, minimal public APIs, and clear ownership boundaries. Most code resides under `internal/`, while only carefully designed, stable interfaces are exposed through `pkg/`. By avoiding generic utility packages, circular dependencies, global state, and unnecessary abstractions, the project maintains a codebase that is predictable, maintainable, and scalable as it evolves.
