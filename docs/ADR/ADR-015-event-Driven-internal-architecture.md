# ADR-015: Event-Driven Internal Architecture

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Internal Component Communication

---

## Context

BeaconLink consists of multiple modules that collaborate to manage platform
state, including:

- Deployment
- Runtime management
- Health monitoring
- Inventory collection
- Configuration reconciliation
- Update management
- Networking
- Audit logging

Many platform operations naturally produce events, such as:

- Deployment completed
- Container started
- Health status changed
- Inventory updated
- Configuration reconciled
- Agent connected
- Update installed
- Policy violated

Direct synchronous communication between modules would create:

- Tight coupling
- Complex dependency graphs
- Reduced extensibility
- Limited observability
- Difficult testing
- Cascading failures

The platform requires an internal communication model that allows
modules to evolve independently while remaining coordinated.

---

## Decision

BeaconLink adopts an **Event-Driven Internal Architecture**.

Modules communicate primarily through domain events that describe
significant state changes within the system.

Event producers publish immutable events without knowledge of event
consumers.

Consumers subscribe only to events relevant to their responsibilities.

Events represent completed business actions rather than commands or
requests.

Synchronous method calls remain appropriate for direct query operations,
validation, and transactional workflows where immediate responses are
required.

---

## Alternatives Considered

### Direct Module Calls

Modules invoke one another through synchronous interfaces.

**Rejected because:**

- Tight coupling
- Dependency proliferation
- Difficult extensibility
- Reduced modularity
- Harder testing

---

### Shared State Communication

Modules coordinate through shared mutable data structures.

**Rejected because:**

- Hidden dependencies
- Race conditions
- Poor traceability
- Difficult debugging
- Weak ownership boundaries

---

### Message Broker for All Communication

Route every internal interaction through an external messaging system.

**Rejected because:**

- Unnecessary operational complexity
- Additional infrastructure
- Increased latency
- Higher resource usage
- Excessive complexity for in-process communication

---

## Consequences

### Positive

- Loose coupling
- Improved modularity
- Easier feature extension
- Better observability
- Simplified testing
- Clear ownership boundaries
- Independent module evolution
- Natural audit capabilities

### Negative

- Event ordering considerations
- Event schema evolution
- Additional event definitions
- Asynchronous debugging complexity
- Event handling discipline required

---

## Architecture

```
             Deployment Module
                     │
        DeploymentCompleted
                     │
                     ▼
              Event Dispatcher
     ┌───────────────┼────────────────┐
     ▼               ▼                ▼
Inventory      Health Monitor     Audit Log
 Module            Module          Module
     │               │                │
     └───────────────┼────────────────┘
                     ▼
              Additional Events
```

Modules publish domain events without depending on the internal
implementation of event consumers.

---

## Related Decisions

- ADR-006 Site Manifest
- ADR-011 Modular Monolith Agent
- ADR-014 Polyglot Persistence
- ADR-016 Declarative Desired State
- ADR-018 Observability Strategy

---

## Implementation Notes

The event architecture should:

- Publish immutable domain events
- Maintain versioned event contracts
- Avoid cyclic event flows
- Support multiple subscribers
- Preserve event ordering where required
- Record events for diagnostics
- Isolate event handling failures
- Distinguish domain events from commands

The event dispatcher implementation, serialization format, event bus
technology, and delivery guarantees remain implementation decisions and
are outside the scope of this ADR.

---

## Rationale

An event-driven internal architecture enables BeaconLink modules to
communicate through explicit business events rather than direct
dependencies.

By separating event producers from consumers, BeaconLink improves
maintainability, extensibility, and testability while allowing new
capabilities to be introduced with minimal impact on existing modules.

This decision complements the modular monolith architecture by providing
loose coupling within a single process, while preserving the option to
evolve individual modules into distributed services in the future with
minimal architectural change.
