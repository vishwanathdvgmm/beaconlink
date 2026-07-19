# ADR-008: Container Runtime Strategy

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Workload Execution Architecture

---

## Context

BeaconLink deploys and manages applications across diverse infrastructure,
including cloud platforms, edge environments, virtual machines, and
bare-metal servers.

Containerization has become the standard mechanism for packaging and
isolating modern workloads, but organizations may use different
container runtimes depending on operational requirements, security
policies, or existing infrastructure.

The platform requires a deployment strategy that:

- Supports industry-standard container technologies
- Avoids dependency on a single runtime
- Enables portable application deployment
- Supports immutable deployments
- Integrates with declarative configuration
- Allows native execution where containers are unavailable
- Supports future runtime technologies

Containerization should improve workload portability without becoming an
architectural dependency.

---

## Decision

BeaconLink adopts a **container-first, runtime-agnostic deployment strategy**.

Containers are the preferred packaging and execution model for managed
applications.

The platform interacts with container runtimes exclusively through the
Runtime Abstraction Layer, allowing workloads to execute on multiple
supported runtimes without modifying platform services.

Where container execution is unavailable or inappropriate, BeaconLink may
execute workloads through supported native runtime adapters using the
same deployment lifecycle.

---

## Alternatives Considered

### Docker-Only Strategy

Require Docker for all BeaconLink deployments.

**Rejected because:**

- Vendor dependency
- Limited runtime flexibility
- Reduced portability
- Incompatible with some enterprise environments
- Difficult future migration

---

### Native Execution Only

Deploy applications directly as operating system services.

**Rejected because:**

- Reduced workload isolation
- Inconsistent deployment model
- Lower portability
- More complex dependency management
- Environment-specific behavior

---

### Runtime-Specific Implementations

Implement deployment logic separately for each runtime.

**Rejected because:**

- Code duplication
- Increased maintenance
- Inconsistent behavior
- Higher testing effort
- Difficult feature evolution

---

## Consequences

### Positive

- Portable application deployment
- Runtime independence
- Improved workload isolation
- Immutable deployment support
- Simplified packaging
- Consistent deployment lifecycle
- Future runtime extensibility
- Better operational consistency

### Negative

- Runtime abstraction layer required
- Runtime capability differences
- Additional adapter maintenance
- Container image lifecycle management
- Native execution requires separate adapters

---

## Architecture

```
             BeaconLink Platform
                    │
                    ▼
         Deployment Orchestrator
                    │
                    ▼
       Runtime Abstraction Layer
                    │
      ┌─────────────┼─────────────┐
      ▼             ▼             ▼
 Docker        Podman      Native Runtime
 Runtime       Runtime        Adapter
      │             │             │
      ▼             ▼             ▼
 Container     Container     Process/Service
```

Deployment behavior remains consistent regardless of the underlying
runtime.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-010 Update System
- ADR-011 Modular Monolith Agent
- ADR-020 Immutable Deployments

---

## Implementation Notes

The container strategy should:

- Prefer OCI-compliant container images
- Support multiple container runtimes
- Use immutable deployment artifacts
- Detect runtime capabilities automatically
- Isolate runtime-specific behavior within adapters
- Support native execution when required
- Expose consistent lifecycle operations
- Maintain compatibility across supported runtime versions

Specific container engines, image registries, orchestration platforms,
and packaging formats remain implementation decisions and are outside
the scope of this ADR.

---

## Rationale

A container-first strategy provides BeaconLink with a consistent, portable,
and isolated execution model while avoiding dependence on any single
container runtime.

By combining OCI-compatible containers with a runtime abstraction layer,
BeaconLink supports heterogeneous deployment environments, simplifies
application packaging, and enables future runtime technologies to be
introduced with minimal architectural impact. Native execution remains
available for environments where containers are not practical, ensuring
the platform can operate across a broad range of infrastructure while
maintaining a unified deployment experience.
