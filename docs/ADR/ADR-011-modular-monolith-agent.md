# ADR-011: Modular Monolith Agent

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Agent Architecture

---

## Context

The BeaconLink Agent is responsible for executing the desired state of a
managed site. Its responsibilities include:

- Deployment management
- Runtime management
- Configuration reconciliation
- Health monitoring
- Inventory collection
- Update management
- Network communication
- Event reporting

These capabilities are tightly related and share the same operational
context, lifecycle, and local system resources.

Splitting the agent into multiple independently deployed services would
introduce unnecessary operational complexity through:

- Inter-process communication
- Additional deployment artifacts
- Increased resource consumption
- Complex service orchestration
- Distributed debugging
- Version synchronization

The platform requires an architecture that remains modular while
minimizing operational overhead on managed systems.

---

## Decision

BeaconLink adopts a **Modular Monolith** architecture for the Agent.

The agent is deployed as a single executable while internally organized
into well-defined modules with explicit interfaces and responsibilities.

Modules communicate through in-process interfaces and domain events
rather than network-based APIs.

The modular boundaries should allow future extraction into independent
services if operational requirements change, but service decomposition
is not an architectural objective.

---

## Alternatives Considered

### Microservices

Split the agent into multiple independent services.

**Rejected because:**

- Increased operational complexity
- Higher memory consumption
- IPC overhead
- More deployment artifacts
- Difficult debugging
- Version coordination required

---

### Layered Monolith

Organize the agent primarily into technical layers.

**Rejected because:**

- Weak business boundaries
- Higher coupling
- Difficult feature evolution
- Encourages shared mutable state
- Reduced modularity

---

### Plugin-Based Architecture

Implement every capability as dynamically loaded plugins.

**Rejected because:**

- Increased complexity
- Plugin lifecycle management
- Dependency management challenges
- More difficult testing
- Limited operational benefit

---

## Consequences

### Positive

- Single deployment artifact
- Low operational overhead
- High performance
- In-process communication
- Simplified debugging
- Easier testing
- Clear module boundaries
- Future extraction remains possible

### Negative

- Single process boundary
- Entire agent upgraded together
- Module discipline must be maintained
- Larger executable
- Requires architectural governance

---

## Architecture

```
               BeaconLink Agent
                     │
 ┌───────────────────┼────────────────────┐
 ▼                   ▼                    ▼
Networking      Deployment          Monitoring
 Module            Module              Module
 │                   │                    │
 ├──────────┐        │        ┌───────────┤
 ▼          ▼        ▼        ▼           ▼
Runtime  Configuration Inventory     Update
Module      Module        Module      Module
                     │
                     ▼
              Shared Domain
```

The agent executes as a single process while maintaining clear modular
boundaries internally.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-010 Update System
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The agent architecture should:

- Be deployed as a single executable
- Organize functionality into cohesive modules
- Maintain explicit module interfaces
- Minimize coupling between modules
- Communicate using domain events where appropriate
- Prevent cyclic dependencies
- Allow independent module testing
- Preserve clear ownership boundaries

The programming language, dependency injection framework, module loading
mechanism, and internal package organization remain implementation
decisions and are outside the scope of this ADR.

---

## Rationale

A modular monolith provides the best balance between architectural
modularity and operational simplicity for infrastructure agents.

By deploying a single executable with clearly defined internal modules,
BeaconLink minimizes resource usage, simplifies deployment, improves
performance through in-process communication, and reduces operational
complexity on managed systems.

This architecture preserves clean separation of responsibilities while
keeping the option to extract individual modules into independent
services if future scalability or operational requirements justify such
a transition. Until that point, the modular monolith offers the lowest
complexity with the highest engineering productivity.
