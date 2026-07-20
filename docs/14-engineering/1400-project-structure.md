# 1400 - Project Structure

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the official repository layout, package organization, dependency boundaries, and engineering conventions used throughout the BeaconLink codebase.

A consistent project structure improves:

- Discoverability
- Maintainability
- Testability
- Code ownership
- Dependency management
- Long-term scalability

Every contributor should follow this document when introducing new packages or modules.

---

# 2. Design Principles

The BeaconLink repository follows several guiding principles.

## 2.1 Simplicity

Directory layouts should be obvious and predictable.

Developers should immediately understand where new code belongs.

---

## 2.2 Modular Design

Each package should have a single responsibility.

Avoid large packages that implement unrelated functionality.

---

## 2.3 Explicit Dependencies

Dependencies should always point inward.

Higher-level components depend on lower-level components.

Lower-level packages must never depend on higher-level services.

---

## 2.4 Internal First

Unless functionality is intentionally reusable outside BeaconLink, it belongs under `internal/`.

---

## 2.5 Stable Public APIs

Only packages intended for external consumption belong under `pkg/`.

Everything else remains internal.

---

# 3. Repository Layout

```text
beaconlink/

├── 📁 .github
│   ├── 📁 ISSUE_TEMPLATE
│   ├── 📁 workflows
│   ├── 📄 CODEOWNERS
│   └── 📝 PULL_REQUEST_TEMPLATE.md
├── 📁 api
├── 📁 cmd
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 beacon
│   └── 📁 relay
├── 📁 configs
├── 📁 deployments
│   ├── 📁 compose
│   ├── 📁 docker
│   ├── 📁 kubernetes
│   └── 📁 systemd
├── 📁 docs
├── 📁 examples
├── 📁 internal
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 auth
│   ├── 📁 config
│   ├── 📁 deployment
│   ├── 📁 generated
│   ├── 📁 logging
│   ├── 📁 metrics
│   ├── 📁 protocol
│   ├── 📁 relay
│   ├── 📁 runtime
│   ├── 📁 security
│   ├── 📁 storage
│   ├── 📁 telemetry
│   └── 📁 version
├── 📁 migrations
├── 📁 pkg
│   ├── 📁 blp
│   └── 📁 sdk
├── 📁 scripts
├── 📁 sdk
├── 📁 tools
├── 📁 test
│   ├── 📁 certificates
│   ├── 📁 configs
│   ├── 📁 fixtures
│   └── 📁 golden
├── 📁 web
│   ├── 📁 public
│   ├── 📁 src
│   ├── ⚙️ package.json
│   └── 📄 vite.config.ts
├── ⚙️ .editorconfig
├── ⚙️ .gitattributes
├── ⚙️ .gitignore
├── 📄 CODEOWNERS
├── 📝 CODE_OF_CONDUCT.md
├── 📝 CONTRIBUTING.md
├── 📄 LICENSE
├── 📄 Makefile
├── 📄 NOTICE
├── 📝 README.md
├── 📝 SECURITY.md
└── 📄 go.mod
```

---

# 4. Top-Level Directories

## `.github/`

Contains GitHub configuration.

```text
├── 📁 .github
│   ├── 📁 ISSUE_TEMPLATE
│   ├── 📁 workflows
│   ├── 📄 CODEOWNERS
│   └── 📝 PULL_REQUEST_TEMPLATE.md
```

Responsibilities:

- CI
- Release automation
- Templates
- Repository configuration

---

## `configs/`

Contains example configuration files.

Examples:

```text
├── 📁 configs
│   ├── ⚙️ agent.yaml
│   ├── ⚙️ development.yaml
│   ├── ⚙️ production.yaml
│   └── ⚙️ relay.yaml
```

Configuration examples must never contain secrets.

---

## `deployments/`

Deployment manifests.

Examples:

```text
├── 📁 deployments
│   ├── 📁 compose
│   ├── 📁 docker
│   ├── 📁 kubernetes (future)
│   └── 📁 systemd
```

---

## `docs/`

Project documentation.

Organized by numbered sections.

Documentation is considered part of the source code.

---

## `examples/`

Example applications and sample configurations.

Examples should compile and remain functional.

---

## `scripts/`

Repository automation.

Examples:

```text
├── 📁 scripts
│   ├── 📄 bootstrap.sh
│   ├── 📄 generate.sh
│   ├── 📄 lint.sh
│   └── 📄 release.sh
```

Scripts should remain platform-independent where practical.

---

## `test/`

Shared test assets.

Examples:

```text
├── 📁 test
│   ├── 📁 certificates
│   ├── 📁 configs
│   ├── 📁 fixtures
│   └── 📁 golden
```

No production code belongs here.

---

## `web/`

React frontend.

Contains:

```text
├── 📁 web
│   ├── 📁 public
│   ├── 📁 src
│   ├── ⚙️ package.json
│   └── 📄 vite.config.ts
```

The frontend is maintained independently from Go services.

---

# 5. cmd/

Every executable application has its own directory.

Example:

```text
├── 📁 cmd
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 beacon
│   └── 📁 relay
```

Each directory contains only:

- main.go
- startup wiring
- dependency construction

Business logic must never live inside `cmd/`.

---

# 6. internal/

The majority of BeaconLink lives here.

```text
├── 📁 internal
```

Packages under `internal/` are private.

They may not be imported by external projects.

---

# 7. internal Package Layout

```text
├── 📁 internal
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 auth
│   ├── 📁 config
│   ├── 📁 deployment
│   ├── 📁 generated
│   ├── 📁 logging
│   ├── 📁 metrics
│   ├── 📁 protocol
│   ├── 📁 relay
│   ├── 📁 runtime
│   ├── 📁 security
│   ├── 📁 storage
│   ├── 📁 telemetry
│   └── 📁 version
```

Each package owns one subsystem.

---

# 8. pkg/

Packages intended for external consumption belong here.

Examples include:

```text
├── 📁 pkg
│   ├── 📁 blp
│   └── 📁 sdk
```

Rules:

- Stable APIs
- Documentation required
- Backward compatibility expected

If external reuse is uncertain, use `internal/`.

---

# 9. Package Responsibilities

Every package should answer one question:

> "What responsibility does this package own?"

If multiple unrelated answers exist, the package should be split.

---

# 10. Package Size

Large packages become difficult to maintain.

General guidance:

- Small interfaces
- Small files
- Focused responsibilities

Avoid "utility" packages.

---

# 11. Dependency Direction

Dependencies always point toward foundational packages.

Example:

```text
   cmd
    │
    ▼
   api
  relay
  agent
    ▼
deployment
    ▼
 runtime
    ▼
  config
 logging
 metrics
    ▼
 stdlib
```

Never reverse this direction.

---

# 12. Circular Dependencies

Circular imports are prohibited.

If packages require each other:

- introduce interfaces
- move shared abstractions
- extract common packages

Do not work around circular dependencies.

---

# 13. Shared Packages

Shared functionality belongs in dedicated packages.

Examples:

```text
logging
config
metrics
errors
validation
telemetry
```

Shared packages should remain generic.

---

# 14. Feature Packages

Feature packages encapsulate business logic.

Examples:

```text
relay
agent
deployment
runtime
api
```

Feature packages should avoid depending on one another unnecessarily.

---

# 15. Interfaces

Interfaces belong where they are consumed.

Do not create "interfaces" packages.

Small interfaces are preferred.

Example:

```go
type Store interface {
    Save(ctx context.Context, d Deployment) error
}
```

---

# 16. Models

Avoid global model packages.

Models should remain close to the subsystem that owns them.

Example:

```text
deployment/
    deployment.go

runtime/
    container.go

agent/
    registration.go
```

---

# 17. Configuration

Configuration should be centralized.

```text
internal/config
```

All services should use the same configuration loader.

---

# 18. Logging

Logging is centralized.

```text
internal/logging
```

Packages should never configure loggers themselves.

---

# 19. Metrics

Metrics belong under:

```text
internal/metrics
```

Subsystem-specific metrics should remain close to the owning package.

---

# 20. Errors

Avoid global error constants.

Prefer:

- wrapped errors
- typed errors
- contextual messages

Error ownership remains with the package that produces them.

---

# 21. Tests

Every package owns its tests.

```text
deployment/

deployment.go
deployment_test.go
```

Integration tests belong under:

```text
test/
```

---

# 22. Naming Conventions

Directory names:

- lowercase
- singular
- descriptive

Examples:

Good:

```text
runtime
logging
config
deployment
protocol
```

Avoid:

```text
utils
helpers
common
misc
temp
```

---

# 23. File Organization

Typical package:

```text
deployment/

deployment.go

service.go

repository.go

validator.go

errors.go

types.go

doc.go

*_test.go
```

Files should be organized by responsibility rather than size.

---

# 24. Generated Code

Generated files belong alongside the owning package.

Examples:

```text
internal/storage/sqlc/

internal/api/openapi/

internal/protocol/generated/
```

Generated files must never be edited manually.

---

# 25. Third-Party Dependencies

External libraries should be isolated.

Avoid exposing third-party types through package APIs whenever possible.

This simplifies future dependency replacement.

---

# 26. Import Rules

Imports should be grouped as:

```go
import (
    "context"
    "errors"

    "github.com/.../internal/config"

    "github.com/go-chi/chi/v5"
)
```

Order:

1. Standard library
2. Internal packages
3. External dependencies

---

# 27. Ownership

Each top-level subsystem should have clear ownership.

Examples:

| Package    | Responsibility                 |
| ---------- | ------------------------------ |
| api        | HTTP interface                 |
| relay      | Communication                  |
| agent      | Host execution                 |
| deployment | Deployment engine              |
| runtime    | Runtime abstraction            |
| protocol   | BeaconLink Protocol            |
| config     | Configuration                  |
| logging    | Logging                        |
| security   | Authentication & authorization |

---

# 28. Evolution

Repository organization should evolve deliberately.

Large structural changes require:

- ADR
- Documentation updates
- Migration plan

Avoid unnecessary package churn.

---

# 29. Anti-Patterns

Avoid introducing:

- God packages
- Circular imports
- Global mutable state
- Generic utility packages
- Deep package nesting
- Multiple responsibilities per package
- Hidden dependencies

---

# 30. Example Repository

```text
beaconlink/
├── 📁 .github
│   ├── 📁 ISSUE_TEMPLATE
│   ├── 📁 workflows
│   ├── 📄 CODEOWNERS
│   └── 📝 PULL_REQUEST_TEMPLATE.md
├── 📁 api
├── 📁 cmd
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 beacon
│   └── 📁 relay
├── 📁 configs
├── 📁 deployments
│   ├── 📁 compose
│   ├── 📁 docker
│   ├── 📁 kubernetes
│   └── 📁 systemd
├── 📁 docs
├── 📁 examples
├── 📁 internal
│   ├── 📁 agent
│   ├── 📁 api
│   ├── 📁 auth
│   ├── 📁 config
│   ├── 📁 deployment
│   ├── 📁 logging
│   ├── 📁 metrics
│   ├── 📁 protocol
│   ├── 📁 relay
│   ├── 📁 runtime
│   ├── 📁 security
│   ├── 📁 storage
│   ├── 📁 telemetry
│   └── 📁 version
├── 📁 pkg
│   ├── 📁 blp
│   └── 📁 sdk
├── 📁 scripts
├── 📁 sdk
├── 📁 test
│   ├── 📁 certificates
│   ├── 📁 configs
│   ├── 📁 fixtures
│   └── 📁 golden
├── 📁 web
│   ├── 📁 public
│   ├── 📁 src
│   ├── ⚙️ package.json
│   └── 📄 vite.config.ts
├── ⚙️ .editorconfig
├── ⚙️ .gitattributes
├── ⚙️ .gitignore
├── 📄 CODEOWNERS
├── 📝 CODE_OF_CONDUCT.md
├── 📝 CONTRIBUTING.md
├── 📄 LICENSE
├── 📄 Makefile
├── 📄 NOTICE
├── 📝 README.md
├── 📝 SECURITY.md
└── 📄 go.mod
```

---

# 31. Compliance

All new packages, directories, and modules introduced into the BeaconLink repository must conform to this structure unless an approved Architecture Decision Record (ADR) explicitly documents an exception.

Consistency in repository organization is essential for maintainability, onboarding, and long-term evolution of the platform.

---

# 32. Summary

The BeaconLink repository is organized around clear ownership boundaries, modular package design, and explicit dependency direction. Executables remain thin, business logic resides within focused internal packages, reusable APIs are isolated under `pkg/`, and documentation evolves alongside the codebase. Adhering to these conventions ensures that the project remains scalable, understandable, and maintainable as new contributors and features are added over time.
