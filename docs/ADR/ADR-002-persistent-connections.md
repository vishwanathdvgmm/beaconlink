# ADR-002: Persistent Connections

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Agent Communication Architecture

---

## Context

BeaconLink manages infrastructure through continuous communication between
the control plane, relay services, and managed agents.

Many management operations require low-latency bidirectional
communication, including:

- Configuration reconciliation
- Deployment orchestration
- Health monitoring
- Event streaming
- Log forwarding
- Remote execution
- Heartbeats
- Status updates

Opening a new network connection for every request introduces additional
latency, increases resource consumption, and complicates communication
across NAT and firewall boundaries.

The platform therefore requires an efficient communication model that
supports continuous interaction between components.

---

## Decision

BeaconLink adopts **persistent, long-lived outbound connections** between
agents and relay services.

Each agent establishes one or more authenticated connections during
startup and maintains them throughout normal operation.

These connections provide a bidirectional communication channel that
supports both request-response interactions and asynchronous event
delivery.

Connections should automatically recover from transient network failures
without requiring operator intervention.

---

## Alternatives Considered

### Short-Lived Request Connections

Agents establish a new connection for every request.

**Rejected because:**

- Higher latency
- Increased TLS negotiation overhead
- Poor scalability
- Inefficient resource usage
- Increased connection churn

---

### Periodic Polling

Agents periodically poll the server for work.

**Rejected because:**

- Higher latency
- Increased unnecessary traffic
- Delayed event delivery
- Inefficient resource utilization
- Poor user experience

---

### Message Queue Only

All communication occurs through an external messaging system.

**Rejected because:**

- Additional infrastructure dependency
- Increased operational complexity
- Less suitable for interactive management
- More difficult request-response coordination

---

## Consequences

### Positive

- Low-latency communication
- Reduced connection overhead
- Efficient resource utilization
- Immediate event delivery
- Better scalability
- Improved user responsiveness
- NAT-friendly networking
- Simplified session management

### Negative

- Long-lived connection management
- Heartbeat implementation required
- Automatic reconnection logic required
- Relay capacity planning becomes important
- Additional monitoring of connection health

---

## Architecture

```
          BeaconLink Control Plane
                  │
                  ▼
            BeaconLink Relay
                  │
      ╔═══════════╧═══════════╗
      ║                       ║
      ▼                       ▼
   Agent A                 Agent B
      ▲                       ▲
      ║                       ║
      ╚═══════════════════════╝

Persistent authenticated connections
remain established during normal
operation.
```

Connections remain active for the lifetime of the agent and carry all
management traffic.

---

## Related Decisions

- ADR-001 Relay-First Networking
- ADR-003 Public-Key Authentication
- ADR-005 Multi-Site Routing
- ADR-009 Protocol Versioning
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

Persistent connections should:

- Use encrypted transport
- Support bidirectional messaging
- Detect connection failures
- Reconnect automatically
- Preserve authentication state where appropriate
- Implement heartbeat mechanisms
- Support protocol version negotiation
- Expose operational metrics

Specific transport protocols remain implementation decisions and are not
defined by this ADR.

---

## Rationale

Persistent connections provide the communication model best suited to a
modern infrastructure management platform.

By maintaining authenticated long-lived sessions, BeaconLink minimizes network
overhead, enables near real-time communication, and supports reliable
management of distributed infrastructure while remaining compatible with
relay-first networking and zero-trust security principles.
