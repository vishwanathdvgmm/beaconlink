# ADR-004: Runtime Abstraction

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Runtime Architecture

---

## Context

BeaconLink is designed to deploy and manage workloads across heterogeneous
execution environments.

Target environments may include:

- Docker
- Podman
- containerd
- CRI-O
- Native operating system services
- Future runtime technologies

Coupling the platform directly to a single runtime would reduce
portability, complicate future adoption of emerging technologies, and
increase implementation complexity across the codebase.

The platform therefore requires a consistent execution model independent
of the underlying runtime implementation.

---

## Decision

BeaconLink adopts a **runtime abstraction layer** between platform services
and workload execution environments.

Platform components interact with a common runtime interface rather than
directly with runtime-specific APIs.

Concrete runtime adapters translate the abstract runtime operations into
the appropriate implementation for the target execution environment.

Runtime selection is determined by deployment configuration and runtime
capability detection rather than application logic.

---

## Alternatives Considered

### Docker-Only Architecture

Implement all workload management directly against Docker APIs.

**Rejected because:**

- Vendor-specific dependency
- Limited portability
- Difficult future migration
- Incompatible with non-container workloads
- Increased long-term technical debt

---

### Runtime-Specific Implementations

Implement separate deployment logic for every supported runtime.

**Rejected because:**

- Significant code duplication
- Higher maintenance cost
- Inconsistent behavior
- Increased testing complexity
- Difficult feature evolution

---

### Direct Operating System Integration

Manage workloads directly through operating system facilities.

**Rejected because:**

- Limited container support
- Reduced portability
- Platform-specific implementations
- Difficult workload isolation
- Increased operational complexity

---

## Consequences

### Positive

- Runtime independence
- Improved portability
- Simplified testing
- Reduced code duplication
- Easier platform evolution
- Support for future runtimes
- Consistent deployment behavior
- Cleaner architecture

### Negative

- Additional abstraction layer
- Adapter maintenance required
- Lowest-common-denominator constraints
- Runtime capability differences require handling
- Slight implementation overhead

---

## Architecture

```
          BeaconLink Platform
                 │
                 ▼
      Runtime Abstraction Layer
                 │
     ┌───────────┼───────────┐
     ▼           ▼           ▼
 Docker      Podman     Native Runtime
 Adapter      Adapter       Adapter
     │           │             │
     ▼           ▼             ▼
 Execution   Execution    Execution
 Environment Environment Environment
```

Platform components depend only on the runtime abstraction, while
runtime-specific behavior is encapsulated within adapters.

---

## Related Decisions

- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-010 Update System
- ADR-011 Modular Monolith Agent
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The runtime abstraction should:

- Provide a stable execution interface
- Support capability discovery
- Handle runtime-specific differences internally
- Normalize common operations
- Expose consistent lifecycle management
- Support health monitoring
- Allow new runtime adapters without modifying platform services
- Enable comprehensive automated testing through mock implementations

Specific adapter implementations remain outside the scope of this ADR.

---

## Rationale

A runtime abstraction layer separates BeaconLink platform logic from
runtime-specific implementations, allowing workloads to execute across
multiple environments without changing higher-level services.

This approach improves portability, reduces vendor lock-in, simplifies
testing, and enables future runtime technologies to be supported through
additional adapters rather than platform-wide architectural changes. The
result is a flexible execution model that aligns with BeaconLink's goals of
modularity, extensibility, and long-term maintainability.
