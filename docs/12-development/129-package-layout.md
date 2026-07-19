# 129 - Package Layout

- **Status:** Approved
- **Version:** 1.0
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the package organization of the BeaconLink Go codebase.

The objective is to ensure every package has a clear responsibility, well-defined ownership, and predictable dependencies.

This document should be used whenever creating new packages or refactoring existing code.

---

# 2. Design Principles

Packages should follow these principles.

- Single Responsibility Principle
- High cohesion
- Low coupling
- Explicit dependencies
- Small public APIs
- Hide implementation details
- Business-domain organization
- No circular imports

---

# 3. Package Categories

BeaconLink packages are divided into three categories.

| Category | Purpose                  |
| -------- | ------------------------ |
| cmd      | Executable entry points  |
| internal | Private implementation   |
| pkg      | Public reusable packages |

---

# 4. Internal Package Layout

```text
internal/

agent/
api/
auth/
config/
database/
deployment/
events/
inventory/
logging/
metrics/
policy/
protocol/
relay/
runtime/
scheduler/
security/
telemetry/
version/
workflow/
```

Each package owns a specific business capability.

---

# 5. Agent

```text
internal/agent/
```

Responsible for the Beacon Agent.

Responsibilities

- Registration
- Heartbeats
- Configuration
- Runtime control
- Health monitoring
- Deployment execution
- Auto recovery
- Update management

Does **not** contain runtime implementations.

---

# 6. Relay

```text
internal/relay/
```

Responsible for the Relay Server.

Responsibilities

- Session management
- Connection lifecycle
- Routing
- Tunnel management
- Domain routing
- Edge networking
- Traffic forwarding

---

# 7. API

```text
internal/api/
```

Responsible for HTTP and WebSocket APIs.

Responsibilities

- REST handlers
- WebSocket handlers
- Middleware
- Validation
- Authentication integration
- Response formatting

Contains no business logic.

---

# 8. Authentication

```text
internal/auth/
```

Responsibilities

- JWT
- mTLS validation
- Identity resolution
- Session authentication
- Token validation

---

# 9. Security

```text
internal/security/
```

Responsibilities

- Certificates
- Cryptography
- Trust management
- Key management
- Secure storage
- Security utilities

Authentication and cryptography remain separate concerns.

---

# 10. Configuration

```text
internal/config/
```

Responsibilities

- Configuration loading
- Validation
- Environment variables
- Defaults
- Configuration watching

No package should read configuration files directly.

---

# 11. Database

```text
internal/database/
```

Responsibilities

- Database initialization
- Connection pools
- sqlc integration
- Transactions
- Migrations
- Query execution

Business logic belongs elsewhere.

---

# 12. Deployment

```text
internal/deployment/
```

Responsibilities

- Desired state
- Deployment plans
- Rollouts
- Rollbacks
- Reconciliation
- Deployment orchestration

---

# 13. Runtime

```text
internal/runtime/
```

Responsible for runtime abstraction.

Responsibilities

- Runtime interface
- Runtime detection
- Runtime selection
- Runtime registry

Implementations

```text
runtime/

docker/
podman/
containerd/
native/
```

Higher-level packages should never depend on a specific runtime.

---

# 14. Workflow

```text
internal/workflow/
```

Responsibilities

- Workflow engine
- Task execution
- State machine
- Scheduling
- Retry logic

Future RFCs extend this package.

---

# 15. Scheduler

```text
internal/scheduler/
```

Responsibilities

- Background jobs
- Timers
- Reconciliation loops
- Periodic tasks

Separate from workflow execution.

---

# 16. Inventory

```text
internal/inventory/
```

Responsible for managing infrastructure inventory.

Responsibilities

- Sites
- Agents
- Applications
- Resources
- Metadata

---

# 17. Policy

```text
internal/policy/
```

Responsibilities

- Policy evaluation
- Authorization rules
- Compliance
- Constraints
- Future policy engine

---

# 18. Protocol

```text
internal/protocol/
```

Responsibilities

- Message encoding
- Packet decoding
- Protocol validation
- Handshake
- Session negotiation

Implements the BeaconLink Protocol (BLP).

---

# 19. Events

```text
internal/events/
```

Responsible for asynchronous communication.

Responsibilities

- Event publishing
- Event subscriptions
- Internal message bus
- Event dispatching

Current implementation uses Go channels.

Future versions may support NATS.

---

# 20. Logging

```text
internal/logging/
```

Responsibilities

- slog initialization
- Structured logging
- Context logging
- Log formatting

No package should create custom loggers.

---

# 21. Metrics

```text
internal/metrics/
```

Responsibilities

- Prometheus metrics
- Counters
- Histograms
- Health metrics

---

# 22. Telemetry

```text
internal/telemetry/
```

Responsibilities

- OpenTelemetry
- Distributed tracing
- Trace propagation

---

# 23. Version

```text
internal/version/
```

Responsibilities

- Build version
- Commit hash
- Build date
- Feature flags

---

# 24. Public Packages

```text
pkg/

client/
protocol/
sdk/
types/
version/
```

These packages may be imported by external applications.

They must maintain backward compatibility.

---

# 25. Package Ownership

Each package should have a clearly defined owner.

A package should never become a dumping ground for unrelated functionality.

When responsibilities grow, create a new package.

---

# 26. Package Size

Guidelines

Package

- Prefer fewer than 20 source files.
- Prefer a focused responsibility.

Source File

- Prefer 100–400 lines.
- Split large files by concern.

Function

- Prefer fewer than 50 lines.
- Longer functions require justification.

---

# 27. Public API Guidelines

Export only what is necessary.

Preferred

```go
type Manager struct {}

func New(...) *Manager
```

Avoid exporting implementation details.

---

# 28. Internal Communication

Packages should communicate through interfaces.

Example

```text
deployment
      │
      ▼
runtime.Interface

runtime/

docker
podman
native
containerd
```

High-level packages depend on abstractions.

Low-level packages implement them.

---

# 29. Dependency Flow

```text
cmd
 │
 ▼
internal/api
 │
 ▼
deployment
 │
 ▼
runtime
 │
 ▼
database
```

Dependencies should always move toward lower-level services.

Reverse dependencies are prohibited.

---

# 30. Forbidden Patterns

Avoid:

- God packages
- Circular imports
- Global mutable state
- Utility packages containing unrelated helpers
- Cross-domain dependencies
- Deep package nesting

---

# 31. Package Review Checklist

Before introducing a new package, ask:

- Does an existing package already own this responsibility?
- Is the responsibility cohesive?
- Does the package expose only necessary APIs?
- Does it introduce circular dependencies?
- Can another package own this functionality instead?

If the answer suggests unnecessary fragmentation, do not create the package.

---

# 32. Evolution

Packages should evolve carefully.

Major restructuring should occur only when:

- Responsibilities change significantly.
- A package becomes difficult to understand.
- Coupling increases.
- Performance or maintainability suffers.

Repository stability is preferred over frequent structural changes.

---

# 33. Summary

BeaconLink packages are organized around **business capabilities rather than technical layers**.

This approach provides:

- Better maintainability
- Clear ownership
- Reduced coupling
- Improved scalability
- Easier onboarding
- Cleaner architecture

Every package should have one reason to change and one clearly defined responsibility.
