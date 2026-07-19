# ADR-006: Site Manifest as the Source of Truth

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Configuration Management Architecture

---

## Context

BeaconLink manages infrastructure across multiple sites containing services,
agents, networks, deployments, domains, policies, and operational
configuration.

Traditional infrastructure management often distributes configuration
across multiple files, databases, environment variables, and runtime
settings. This fragmentation leads to:

- Configuration drift
- Inconsistent deployments
- Difficult auditing
- Complex recovery procedures
- Operational uncertainty
- Reduced automation

The platform requires a single authoritative representation of a site's
desired state that can be versioned, validated, distributed, and
reconciled consistently across the platform.

---

## Decision

BeaconLink adopts the **Site Manifest** as the authoritative description of a
site's desired state.

The Site Manifest defines all managed resources, configuration,
relationships, deployment intent, networking, policies, and runtime
requirements for a site.

Platform components reconcile the actual runtime state toward the
desired state described by the Site Manifest rather than executing
imperative configuration changes.

The Site Manifest becomes the single source of truth for site
configuration.

---

## Alternatives Considered

### Database-Driven Configuration

Store configuration exclusively within internal databases.

**Rejected because:**

- Difficult version control
- Poor portability
- Limited auditability
- Complicated disaster recovery
- Reduced transparency

---

### Imperative Configuration

Execute configuration through commands and scripts.

**Rejected because:**

- Configuration drift
- Non-repeatable deployments
- Difficult rollback
- Poor state visibility
- Increased operational complexity

---

### Multiple Configuration Files

Maintain separate configuration for each subsystem.

**Rejected because:**

- Fragmented configuration
- Difficult validation
- Higher maintenance effort
- Increased inconsistency
- Poor dependency visibility

---

## Consequences

### Positive

- Single source of truth
- Declarative infrastructure
- Improved auditability
- Simplified disaster recovery
- Predictable deployments
- Reduced configuration drift
- Better validation
- Easier automation

### Negative

- Manifest schema evolution required
- Validation complexity increases
- Large manifests require careful organization
- Reconciliation engine required
- Manifest compatibility must be maintained

---

## Architecture

```
            Site Manifest
                  │
                  ▼
        Validation & Schema
                  │
                  ▼
      Desired State Repository
                  │
                  ▼
      Reconciliation Engine
                  │
      ┌───────────┼───────────┐
      ▼           ▼           ▼
   Services     Agents     Network
                  │
                  ▼
           Actual Runtime
```

All platform components derive their desired configuration from the Site
Manifest.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-005 Multi-Site Routing
- ADR-010 Update System
- ADR-011 Modular Monolith Agent
- ADR-016 Declarative Desired State

---

## Implementation Notes

The Site Manifest should:

- Be declarative
- Be version controlled
- Support schema evolution
- Be fully validated before application
- Describe desired rather than current state
- Support incremental reconciliation
- Be portable across environments
- Allow backward-compatible schema evolution

The physical representation, serialization format, and storage
mechanism remain implementation decisions and are outside the scope of
this ADR.

---

## Rationale

The Site Manifest provides BeaconLink with a single authoritative definition
of infrastructure intent.

By adopting a declarative configuration model, BeaconLink separates desired
state from implementation details, enabling deterministic deployments,
continuous reconciliation, simplified recovery, and consistent
operations across distributed environments.

This decision establishes the Site Manifest as the architectural
foundation for deployment, configuration management, orchestration, and
lifecycle management throughout the BeaconLink platform.
