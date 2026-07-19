# ADR-017: API Gateway Pattern

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** External API Architecture

---

## Context

BeaconLink exposes multiple APIs serving different consumers, including:

- Web Console
- CLI
- Automation tools
- External integrations
- SDKs
- Future third-party applications

These consumers interact with multiple platform capabilities such as:

- Authentication
- Site management
- Deployment
- Inventory
- Health monitoring
- Configuration
- Events
- Audit information

Allowing clients to communicate directly with internal services would
introduce several challenges:

- Increased client complexity
- Tight coupling to internal architecture
- Inconsistent authentication
- Duplicate authorization logic
- Fragmented APIs
- Difficult API evolution
- Expanded attack surface

The platform requires a unified entry point that presents a consistent,
stable, and secure interface while isolating clients from internal
implementation details.

---

## Decision

BeaconLink adopts the **API Gateway Pattern**.

All external API requests are routed through a single logical API
Gateway.

The gateway is responsible for cross-cutting concerns including:

- Authentication
- Authorization
- Request validation
- Routing
- Protocol translation
- Rate limiting
- Audit logging
- API version negotiation

Business logic remains within backend services.

The gateway acts solely as a coordination and policy enforcement layer
and does not become a repository for domain logic.

---

## Alternatives Considered

### Direct Service Access

Expose every backend service directly to clients.

**Rejected because:**

- Tight client coupling
- Multiple authentication mechanisms
- Inconsistent APIs
- Larger attack surface
- Difficult client evolution

---

### Backend-for-Frontend Only

Create separate APIs for every client application.

**Rejected because:**

- Duplicate gateway functionality
- Increased maintenance
- Fragmented external APIs
- More deployment complexity
- Reduced consistency

---

### Service Mesh as External Gateway

Expose the service mesh directly to clients.

**Rejected because:**

- Service meshes target service-to-service communication
- Limited API lifecycle management
- Weak client-facing abstractions
- Poor developer experience
- Increased operational complexity

---

## Consequences

### Positive

- Single API entry point
- Consistent authentication
- Centralized authorization
- Simplified clients
- Stable external contracts
- Easier API evolution
- Improved security
- Unified observability

### Negative

- Additional infrastructure component
- Gateway availability becomes critical
- Potential routing bottleneck
- Requires horizontal scaling
- Gateway policies require governance

---

## Architecture

```
                 External Clients
                        │
        ┌───────────────┼────────────────┐
        ▼               ▼                ▼
      Console          CLI             SDKs
            \           │             /
             \          │            /
              ▼         ▼           ▼
               Logical API Gateway
                      │
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
 Authentication   API Services   Event Services
        │             │             │
        └─────────────┼─────────────┘
                      ▼
               Platform Services
```

Clients communicate exclusively with the logical API Gateway, which
applies platform-wide policies before routing requests to backend
services.

---

## Related Decisions

- ADR-006 Site Manifest
- ADR-007 Zero Trust
- ADR-009 Protocol Versioning
- ADR-015 Event-Driven Internal Architecture
- ADR-016 Declarative Desired State

---

## Implementation Notes

The API Gateway should:

- Authenticate every request
- Enforce authorization policies
- Route requests to appropriate services
- Validate request structure
- Support API version negotiation
- Apply rate limiting and quotas
- Record audit and access logs
- Expose unified health information

The gateway implementation, deployment topology, routing engine,
authentication providers, and rate-limiting algorithms remain
implementation decisions and are outside the scope of this ADR.

---

## Rationale

A logical API Gateway provides a consistent and secure interface between
external consumers and the BeaconLink platform.

By centralizing authentication, authorization, routing, and other
cross-cutting concerns, BeaconLink reduces client complexity, strengthens
security, and preserves the independence of internal services.

This decision supports long-term API evolution by allowing backend
implementations to change without affecting external consumers, while
aligning with the platform's zero-trust security model, protocol
versioning strategy, and modular architecture.
