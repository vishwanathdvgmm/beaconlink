# ADR-016: Declarative Desired State

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Desired State Reconciliation

---

## Context

BeaconLink manages distributed infrastructure across multiple sites and
runtime environments.

Infrastructure continuously changes due to:

- Software deployments
- Configuration updates
- Runtime failures
- Manual intervention
- Node replacement
- Network interruptions
- System recovery
- Platform upgrades

Managing infrastructure through imperative operations alone would
require the platform to track every action taken, making recovery from
partial failures increasingly difficult.

Imperative workflows also introduce challenges including:

- Configuration drift
- Non-repeatable deployments
- Inconsistent recovery
- Complex rollback
- Operational uncertainty
- Hidden system state

The platform requires an operational model where the intended platform
configuration is explicitly defined and continuously enforced.

---

## Decision

BeaconLink adopts a **Declarative Desired State** architecture.

The desired configuration of every managed site is expressed as a
declarative specification, represented by the Site Manifest.

Agents continuously reconcile the observed system state with the desired
state.

Rather than executing imperative sequences of operations, reconciliation
identifies differences between the current and desired states and
performs only the actions necessary to eliminate those differences.

Reconciliation is continuous, idempotent, and repeatable.

---

## Alternatives Considered

### Imperative Deployment Model

Execute predefined deployment scripts or commands.

**Rejected because:**

- Difficult recovery
- Configuration drift
- Poor repeatability
- Complex rollback
- Hidden operational state

---

### Manual State Management

Require operators to maintain infrastructure manually.

**Rejected because:**

- Error-prone
- Poor scalability
- Operational inconsistency
- Limited automation
- Increased maintenance effort

---

### Event-Only Synchronization

React exclusively to individual configuration change events.

**Rejected because:**

- Missed events create inconsistency
- Recovery becomes difficult
- No complete desired state exists
- Drift detection is limited
- Full reconciliation is impossible

---

## Consequences

### Positive

- Continuous drift correction
- Idempotent operations
- Predictable recovery
- Simplified automation
- Repeatable deployments
- Consistent infrastructure
- Easier rollback
- Improved operational resilience

### Negative

- Reconciliation engine required
- Drift detection overhead
- State comparison logic
- Eventual consistency model
- Resource consumption for reconciliation

---

## Architecture

```
          Site Manifest
      (Desired State)
              │
              ▼
     Reconciliation Engine
              │
      Compare Desired
      vs Current State
              │
              ▼
      Execute Required
          Operations
              │
              ▼
       Observed State
              │
              └───────────────┐
                              │
                    Continuous Loop
```

The reconciliation engine continuously aligns the observed system state
with the declarative desired state.

---

## Related Decisions

- ADR-006 Site Manifest
- ADR-010 Immutable Update System
- ADR-011 Modular Monolith Agent
- ADR-014 Polyglot Persistence
- ADR-015 Event-Driven Internal Architecture

---

## Implementation Notes

The reconciliation engine should:

- Continuously evaluate desired state
- Detect configuration drift
- Execute idempotent operations
- Minimize unnecessary changes
- Support partial reconciliation
- Recover automatically after interruptions
- Report reconciliation status
- Preserve audit history of state transitions

The reconciliation interval, scheduling strategy, conflict resolution
policies, and execution algorithms remain implementation decisions and
are outside the scope of this ADR.

---

## Rationale

A declarative desired-state architecture provides a reliable and
predictable model for managing distributed infrastructure.

By defining what the platform should look like rather than prescribing
how to reach that state, BeaconLink enables continuous reconciliation,
automatic drift correction, repeatable deployments, and resilient
recovery from failures.

This decision aligns with the Site Manifest as the authoritative source
of configuration and complements the platform's immutable deployment,
event-driven architecture, and modular agent design, providing a
consistent operational model across all managed environments.
