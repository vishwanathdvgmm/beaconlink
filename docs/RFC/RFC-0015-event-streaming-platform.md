# RFC-0015: Event Streaming Platform

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing an **Event Streaming Platform** that
provides a centralized, durable, and scalable event backbone for BeaconLink.

Rather than relying solely on synchronous service-to-service
communication, BeaconLink services publish and consume domain events through
a shared event streaming infrastructure, enabling loose coupling,
improved scalability, and resilient asynchronous processing.

The objective is to establish events as a first-class architectural
primitive while preserving synchronous APIs where appropriate.

---

# Motivation

BeaconLink already contains numerous components that naturally generate
events, including:

- Deployments
- Agent lifecycle
- Runtime operations
- Policy evaluation
- Secret management
- Health monitoring
- Audit logging
- Workflow execution
- Authentication
- Configuration reconciliation

Currently, many of these interactions occur through direct service
communication.

As BeaconLink grows, this approach increases:

- Service coupling
- Deployment dependencies
- Retry complexity
- Failure propagation
- Operational bottlenecks
- Integration effort

A centralized event streaming platform enables services to communicate
through immutable event streams rather than direct dependencies.

---

# Goals

The proposed platform should:

- Support asynchronous communication
- Decouple platform services
- Provide durable event storage
- Support multiple event consumers
- Enable replayable event streams
- Preserve event ordering where required
- Scale horizontally
- Integrate with observability

---

# Non-Goals

This RFC does **not** propose:

- Replacing synchronous APIs
- Becoming the primary data store
- Implementing Event Sourcing
- Eliminating request-response communication
- Replacing audit logging
- Executing business workflows
- Providing message transformation
- Becoming an enterprise integration platform

---

# Proposed Design

BeaconLink introduces a centralized Event Streaming Platform shared across
all major platform components.

```
               BeaconLink Services
                      │
      ┌───────────────┼───────────────┐
      ▼               ▼               ▼
 Deployment        Agent         API Gateway
      │               │               │
      └───────────────┼───────────────┘
                      ▼
           Event Streaming Platform
                      │
      ┌───────────────┼───────────────┐
      ▼               ▼               ▼
 Observability    Workflow     Remediation
                  Engine         Engine
                      │
                      ▼
                Event Consumers
```

Services publish domain events without requiring knowledge of event
subscribers.

Consumers independently subscribe to events relevant to their domain.

---

# Event Model

Events represent immutable facts describing completed activities.

Examples include:

- DeploymentStarted
- DeploymentCompleted
- AgentConnected
- AgentDisconnected
- PolicyEvaluated
- SecretRotated
- WorkflowStarted
- WorkflowCompleted
- RuntimeProvisioned
- SiteUpdated

Events should be:

- Immutable
- Timestamped
- Versioned
- Authenticated
- Serializable
- Independently consumable

---

# Event Lifecycle

Each event progresses through:

1. Generation
2. Validation
3. Publication
4. Persistence
5. Distribution
6. Consumption
7. Acknowledgement
8. Retention
9. Expiration

Event publication should remain reliable even during temporary consumer
failures.

---

# Event Consumers

Potential consumers include:

- Observability Platform
- Workflow Engine
- Autonomous Remediation
- Audit subsystem
- Notification services
- Deployment engine
- Analytics services
- Plugin extensions
- Future integrations

Multiple consumers may independently process the same event without
affecting one another.

---

# Delivery Semantics

The platform should support:

- Durable delivery
- Ordered delivery within streams
- Consumer isolation
- Replay capability
- Configurable retention
- Retry mechanisms
- Dead-letter handling
- Backpressure management

The exact delivery guarantees remain implementation-specific.

---

# Event Schema

Every event should contain common metadata including:

- Event identifier
- Event type
- Version
- Timestamp
- Source service
- Correlation identifier
- Tenant identifier (where applicable)
- Site identifier
- Payload

Schemas should evolve using backward-compatible versioning whenever
possible.

---

# Security Considerations

Events may contain operationally sensitive information.

The platform should:

- Authenticate publishers
- Authenticate consumers
- Authorize subscriptions
- Encrypt event transport
- Validate event schemas
- Audit event publication
- Prevent unauthorized event injection
- Protect tenant isolation

Sensitive secret values must never be published within event payloads.

---

# Operational Impact

Potential benefits include:

- Reduced service coupling
- Improved scalability
- Better fault isolation
- Replayable operational history
- Simplified integrations
- Enhanced observability
- Independent service evolution

Potential challenges include:

- Event schema governance
- Stream management
- Consumer lag
- Event retention policies
- Operational monitoring
- Capacity planning

---

# Compatibility

Existing BeaconLink APIs remain unchanged.

Services may gradually adopt event-driven communication while preserving
existing synchronous interfaces.

The Event Streaming Platform complements, rather than replaces,
traditional service APIs.

---

# Alternatives Considered

## Direct Service Communication

Continue using synchronous APIs between services.

Pros:

- Simple architecture
- Familiar programming model

Cons:

- Tight coupling
- Reduced resilience
- Difficult scalability
- Increased dependency management

---

## Traditional Message Queue

Adopt a point-to-point messaging system.

Pros:

- Simple implementation
- Reliable messaging

Cons:

- Limited replay
- Weak event history
- Reduced support for multiple consumers
- Less suitable for event-driven architecture

---

## Event Sourcing

Persist all application state exclusively as events.

Rejected because:

- Significant architectural complexity
- Higher storage requirements
- Difficult migration
- Broader platform impact
- Outside BeaconLink architectural scope

---

# Open Questions

- Which streaming technology should BeaconLink adopt?
- What event retention periods should be supported?
- How should schema evolution be managed?
- Should events support cross-region replication?
- Which services should publish events first?
- Should external consumers be supported?

---

# Related Documents

- ADR-015 Event-Driven Internal Architecture
- ADR-016 Declarative Desired State
- ADR-018 Observability Strategy
- RFC-0009 Policy Engine
- RFC-0013 Autonomous Remediation
- RFC-0014 Workflow Engine

---

# Recommendation

BeaconLink should introduce a centralized **Event Streaming Platform** as
the backbone for asynchronous communication between platform services.

The preferred architecture establishes immutable domain events as the
primary integration mechanism for internal services while preserving
existing synchronous APIs for request-response interactions.

Initial implementation should focus on core operational domains—
including deployments, agent lifecycle, workflow execution, and policy
evaluation—providing a scalable foundation for future automation,
observability, analytics, and ecosystem integrations without increasing
service coupling.
