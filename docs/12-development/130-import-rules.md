# 130 - Import Rules

- **Status:** Approved
- **Version:** 1.0
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the architectural dependency rules for the BeaconLink codebase.

Its purpose is to prevent architectural erosion by establishing clear package boundaries, dependency directions, and import policies.

Every package must follow these rules.

---

# 2. Guiding Principles

BeaconLink follows four fundamental dependency principles.

- Dependencies flow inward.
- High-level modules depend on abstractions.
- Low-level modules never depend on application logic.
- Circular dependencies are never allowed.

These principles apply to every package in the repository.

---

# 3. Dependency Hierarchy

The permitted dependency flow is:

```text
Applications (cmd)
        │
        ▼
Application Services (internal)
        │
        ▼
Shared Libraries (pkg)
```

Imports may only move downward.

---

# 4. Layer Responsibilities

| Layer    | Responsibility            |
| -------- | ------------------------- |
| cmd      | Bootstrap applications    |
| internal | Business logic            |
| pkg      | Shared reusable libraries |

---

# 5. Import Direction

The only valid dependency graph is:

```text
cmd
 │
 ▼
internal
 │
 ▼
pkg
```

Never reverse this direction.

---

# 6. Allowed Imports

## cmd

May import

- internal
- pkg
- Go standard library

---

## internal

May import

- internal
- pkg
- Go standard library

---

## pkg

May import

- pkg
- Go standard library

---

# 7. Forbidden Imports

The following imports are prohibited.

## pkg → internal

❌

```text
pkg/client
        │
        ▼
internal/database
```

Public packages must never depend on implementation packages.

---

## internal → cmd

❌

```text
internal/runtime
        │
        ▼
cmd/agent
```

Applications are entry points only.

---

## cmd → cmd

❌

Applications must never import one another.

Shared functionality belongs in internal packages.

---

## Circular Dependencies

Forbidden.

```text
agent
 │
 ▼
deployment
 │
 ▼
runtime
 │
 ▼
agent
```

Go already prevents circular imports.

Architecturally they are never acceptable.

---

# 8. Domain Dependencies

Business domains should remain independent whenever possible.

Preferred

```text
deployment
      │
      ▼
runtime
```

Avoid

```text
deployment
      │
      ▼
relay
      │
      ▼
agent
      │
      ▼
deployment
```

---

# 9. Cross-Domain Communication

Cross-domain interaction should occur through interfaces.

Example

```go
type Runtime interface {
    Deploy(...)
    Stop(...)
}
```

Concrete implementations remain hidden.

---

# 10. Interface Placement

Interfaces belong where they are consumed.

Preferred

```text
deployment/

runtime.go
```

```go
type Runtime interface {
    Deploy(...)
}
```

Implemented by

```text
runtime/docker
runtime/podman
runtime/native
```

Do not place interfaces in arbitrary "interfaces" packages.

---

# 11. Dependency Inversion

High-level packages depend on abstractions.

Example

```text
deployment
        │
        ▼
Runtime Interface
        ▲
        │
docker
podman
native
```

Concrete implementations never control application flow.

---

# 12. Package Visibility

Use the smallest possible visibility.

Prefer

```go
type manager struct {}
```

instead of

```go
type Manager struct {}
```

Export only when required.

---

# 13. Constructors

Every package should expose constructors instead of global variables.

Preferred

```go
func New(cfg Config) *Manager
```

Avoid

```go
var DefaultManager = ...
```

---

# 14. Global State

Global mutable state is prohibited.

Avoid

- package globals
- singleton managers
- hidden caches

Exceptions

- immutable constants
- version information
- build metadata

---

# 15. Configuration Access

Configuration is loaded once.

Only the configuration package reads configuration files.

Other packages receive configuration through dependency injection.

Preferred

```text
config
      │
      ▼
deployment
```

Not

```text
deployment
      │
      ▼
config.yaml
```

---

# 16. Database Access

Database access belongs exclusively in the database package.

Preferred

```text
deployment
      │
      ▼
database
```

Avoid

```text
deployment
      │
      ▼
sql.DB
```

Application packages should not construct SQL queries directly.

---

# 17. Logging

Logging is centralized.

Preferred

```text
logging
      │
      ▼
deployment
```

Avoid

```go
log.New(...)
```

inside every package.

---

# 18. Event Publishing

Event publishers should expose interfaces.

Example

```text
deployment
      │
      ▼
events.Publisher
```

Packages should never know the implementation details.

---

# 19. Runtime Independence

Deployment code must never depend on Docker directly.

Preferred

```text
deployment
      │
      ▼
runtime.Interface
```

Not

```text
deployment
      │
      ▼
docker.Client
```

---

# 20. Protocol Independence

Business packages should not serialize protocol messages.

Preferred

```text
relay
      │
      ▼
protocol
```

Protocol encoding remains isolated.

---

# 21. Error Handling

Errors should move upward.

Lower layers return errors.

Upper layers decide how to respond.

Example

```text
database
      │
      ▼
deployment
      │
      ▼
api
```

The API layer converts errors into HTTP responses.

---

# 22. Context Propagation

All long-running operations must accept a context.

Preferred

```go
func Deploy(ctx context.Context, ...)
```

Avoid

```go
func Deploy(...)
```

for operations involving:

- network I/O
- filesystem
- database
- runtime operations

---

# 23. Utility Packages

Avoid generic utility packages.

Do not create

```text
utils/
helpers/
common/
misc/
```

Instead

```text
logging/
config/
runtime/
security/
```

Utilities should belong to their owning domain.

---

# 24. Package Size

If a package becomes difficult to understand, split it by responsibility.

Signs include

- too many exported types
- unrelated responsibilities
- excessive file count
- frequent merge conflicts

---

# 25. External Dependencies

External packages should be imported only where needed.

Do not wrap the entire application around a framework.

Prefer the Go standard library whenever practical.

---

# 26. Public APIs

Packages should expose only stable APIs.

Preferred

```go
type Manager struct {}

func New(...)
func Start(...)
func Stop(...)
```

Avoid exposing implementation details.

---

# 27. Testing Imports

Tests may import

- internal
- pkg
- testing utilities

Production packages must never depend on testing packages.

---

# 28. Build Dependencies

Development tooling must never become runtime dependencies.

Examples

- linters
- generators
- benchmarking tools

These belong in the development workflow only.

---

# 29. Architectural Review

Before adding a new import, ask:

- Does this violate the dependency direction?
- Can an interface remove this dependency?
- Does this create tighter coupling?
- Does another package already own this responsibility?
- Will this make testing more difficult?

If the answer is yes, redesign before importing.

---

# 30. Dependency Diagram

```text
                    cmd
         ┌──────────┼──────────┐
         │          │          │
         ▼          ▼          ▼
      agent      relay       api
         │          │          │
         └──────┬───┴──────────┘
                ▼
            deployment
                │
      ┌─────────┼──────────┐
      ▼         ▼          ▼
   runtime   database   workflow
      │         │          │
      └─────────┼──────────┘
                ▼
              pkg
```

Dependencies always move downward.

---

# 31. Repository Health

A healthy repository exhibits:

- No circular imports
- Small focused packages
- Stable public APIs
- Clear ownership
- Minimal coupling
- High cohesion
- Predictable dependency flow

---

# 32. Governance

Import rules are architectural constraints.

Violations should be corrected immediately rather than accepted as technical debt.

Changes to these rules require an Architecture Decision Record (ADR).

Maintaining strict dependency boundaries is essential to preserving the long-term maintainability, modularity, and scalability of the BeaconLink platform.
