# ADR-013: Monorepo Strategy

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Source Code Organization

---

## Context

BeaconLink is a platform composed of multiple closely related components,
including:

- Control Plane
- Relay Services
- Agent
- Console
- APIs
- Shared libraries
- SDKs
- Deployment tooling
- Documentation

These components evolve together, share common domain models, APIs,
protocols, and engineering standards.

Maintaining separate repositories for each component would introduce
additional complexity through:

- Version synchronization
- Cross-repository dependency management
- Coordinated releases
- Duplicate tooling
- Fragmented documentation
- Complex pull request workflows

The platform requires a repository strategy that supports consistent
development while enabling independent modular evolution.

---

## Decision

BeaconLink adopts a **Monorepo** strategy.

All first-party platform components are maintained within a single
version-controlled repository.

The repository is organized into clearly defined modules with explicit
architectural boundaries and ownership.

Although components may be independently built, tested, packaged, and
deployed, they share a common repository, engineering workflow, and
development lifecycle.

Repository organization reflects business domains rather than
programming languages or team structures.

---

## Alternatives Considered

### Multiple Independent Repositories

Maintain each component in a separate repository.

**Rejected because:**

- Cross-repository dependency management
- Version coordination
- More complex refactoring
- Fragmented documentation
- Increased CI/CD complexity
- Difficult atomic platform changes

---

### Hybrid Repository Model

Maintain shared libraries separately while keeping applications together.

**Rejected because:**

- Artificial repository boundaries
- Dependency synchronization
- Duplicate release management
- Increased maintenance
- Limited operational benefit

---

### Service-Owned Repositories

Each service owns its own repository.

**Rejected because:**

- Reduced architectural consistency
- Difficult platform-wide changes
- Duplicate tooling
- Increased engineering overhead
- Harder protocol evolution

---

## Consequences

### Positive

- Single source of truth
- Simplified dependency management
- Atomic cross-component changes
- Unified CI/CD pipeline
- Consistent engineering standards
- Easier refactoring
- Shared documentation
- Improved developer experience

### Negative

- Larger repository size
- Longer clone times
- More complex build optimization
- Repository permissions require careful management
- CI scalability requires optimization

---

## Architecture

```
                 BeaconLink Repository
                        │
 ┌──────────────────────┼──────────────────────┐
 ▼                      ▼                      ▼
Applications       Shared Libraries      Infrastructure
 │                      │                      │
 ├──────────────┐        │        ┌────────────┤
 ▼              ▼        ▼        ▼            ▼
Agent        Relay     SDKs     Deployment  Documentation
```

All platform components share a common repository while maintaining
clear module boundaries.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-009 Protocol Versioning
- ADR-011 Modular Monolith Agent
- ADR-019 Trunk-Based Development

---

## Implementation Notes

The repository should:

- Organize code by business capability
- Maintain explicit module boundaries
- Prevent cyclic dependencies
- Support incremental builds
- Enable independent testing
- Centralize shared tooling
- Share engineering standards
- Maintain a consistent directory structure

Repository layout, build tooling, package management, and CI
implementation remain implementation decisions and are outside the scope
of this ADR.

---

## Rationale

A monorepo provides the best balance between architectural consistency,
developer productivity, and long-term maintainability for the BeaconLink
platform.

By maintaining all first-party components in a single repository, BeaconLink
simplifies dependency management, enables atomic platform-wide changes,
and ensures that shared contracts, APIs, and documentation evolve
together.

This strategy complements BeaconLink's modular architecture by allowing
independent evolution of components without sacrificing consistency,
traceability, or engineering efficiency.
