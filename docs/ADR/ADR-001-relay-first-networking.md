# ADR-001: Relay-First Networking

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Core Networking Architecture

---

## Context

BeaconLink is designed to manage infrastructure distributed across multiple
sites, networks, cloud providers, and isolated environments.

Many managed environments operate behind firewalls, NAT devices, carrier
networks, or restrictive security policies that prevent inbound
connections. Requiring direct connectivity to every managed node would
increase deployment complexity, reduce reliability, and limit platform
adoption.

The networking architecture therefore requires a communication model
that:

- Works across NAT and firewall boundaries
- Avoids inbound port exposure
- Supports geographically distributed deployments
- Enables secure communication between isolated sites
- Scales to large numbers of managed agents
- Maintains centralized control while minimizing operational overhead

---

## Decision

BeaconLink adopts a **relay-first networking architecture**.

All managed agents establish outbound authenticated connections to one
or more BeaconLink Relay services. Communication between the management
platform and agents flows through these relay connections rather than
through direct inbound connectivity.

Direct peer-to-peer communication is considered an optimization rather
than a fundamental architectural requirement.

The relay layer becomes the primary communication fabric for the BeaconLink
platform.

---

## Alternatives Considered

### Direct Agent Connectivity

Every agent exposes an API endpoint accessible from the management
platform.

**Rejected because:**

- Requires inbound firewall rules
- Difficult across NAT
- Complicates multi-site deployments
- Increases operational burden
- Larger attack surface

---

### VPN-Based Connectivity

All managed systems participate in a shared VPN.

**Rejected because:**

- Operationally complex
- Requires additional infrastructure
- Difficult for customer-managed environments
- Higher deployment friction

---

### Hybrid Direct + Relay

Attempt direct connectivity and fall back to relays.

**Rejected because:**

- Introduces multiple networking models
- Complicates troubleshooting
- Increases implementation complexity
- Makes security policy inconsistent

---

## Consequences

### Positive

- Works in restrictive networks
- No inbound firewall requirements
- Simplified deployment
- Consistent communication model
- Improved scalability
- Easier multi-site support
- Better security posture
- Simplified routing architecture

### Negative

- Relay infrastructure becomes critical
- Additional network hop
- Relay availability impacts connectivity
- Increased relay operational responsibility

---

## Architecture

```
          BeaconLink Console
                 │
                 ▼
          BeaconLink API Services
                 │
                 ▼
          ┌───────────────┐
          │ BeaconLink Relay   │
          └───────┬───────┘
                  │
      ┌───────────┼───────────┐
      ▼           ▼           ▼
   Agent A     Agent B     Agent C
```

All agents maintain outbound authenticated connections to the relay.

---

## Related Decisions

- ADR-002 Persistent Connections
- ADR-003 Public-Key Authentication
- ADR-005 Multi-Site Routing
- ADR-007 Zero Trust
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The relay implementation should:

- Maintain long-lived outbound connections
- Support horizontal scaling
- Remain stateless where practical
- Route traffic using authenticated identities
- Support protocol evolution
- Integrate with the BeaconLink Site Manifest
- Expose operational metrics and health information

Implementation details remain outside the scope of this ADR.

---

## Rationale

A relay-first architecture provides the most reliable networking model
for heterogeneous enterprise environments while minimizing deployment
complexity and maximizing security.

Although direct connectivity may reduce latency in specific scenarios,
the operational simplicity, consistent security model, and universal
deployability of relay-first networking provide significantly greater
long-term architectural value for the BeaconLink platform.
