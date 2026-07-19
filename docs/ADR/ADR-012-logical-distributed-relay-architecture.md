# ADR-012: Logical Distributed Relay Architecture

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Distributed Relay Architecture

---

## Context

BeaconLink is intended to manage infrastructure across geographically
distributed environments, including:

- Enterprise data centers
- Regional cloud deployments
- Branch offices
- Edge computing sites
- Air-gapped environments
- Customer-managed infrastructure

A single centralized relay would introduce several challenges:

- Geographic latency
- Single operational bottleneck
- Reduced fault tolerance
- Limited horizontal scalability
- Regional service disruptions
- Operational concentration

At the same time, exposing multiple independent relay clusters directly
to clients would complicate routing, identity management, and
administration.

The platform therefore requires an architecture that provides
distributed deployment while presenting a unified communication model.

---

## Decision

BeaconLink adopts a **Logical Distributed Relay Architecture**.

Multiple relay instances may be deployed across independent sites and
regions, but collectively they operate as a single logical relay
platform.

Relays cooperate by exchanging routing information, forwarding traffic,
and maintaining awareness of reachable sites.

Clients and agents connect to any available relay without requiring
knowledge of the overall relay topology.

Routing decisions are based on logical identities and site information
rather than physical network location.

---

## Alternatives Considered

### Single Central Relay

Deploy one relay for the entire platform.

**Rejected because:**

- Single point of failure
- Limited scalability
- Higher latency
- Geographic dependency
- Operational bottleneck

---

### Independent Relay Islands

Deploy completely isolated relay instances.

**Rejected because:**

- No inter-site routing
- Fragmented management
- Difficult federation
- Poor resource sharing
- Limited scalability

---

### Full Mesh Connectivity

Require every relay to maintain direct communication with every other
relay.

**Rejected because:**

- Poor scalability
- Increased synchronization complexity
- Higher operational overhead
- Difficult large-scale deployments
- Complex topology management

---

## Consequences

### Positive

- Horizontal scalability
- Regional deployment flexibility
- Improved resilience
- Lower communication latency
- Simplified client configuration
- Better fault isolation
- Supports multi-site expansion
- Consistent networking model

### Negative

- Distributed routing coordination
- Relay discovery required
- Routing metadata synchronization
- Additional operational monitoring
- More complex distributed diagnostics

---

## Architecture

```
                  BeaconLink Control Plane
                          │
                          ▼
                Logical Relay Network
        ┌─────────────────┼─────────────────┐
        ▼                 ▼                 ▼
    Relay A           Relay B           Relay C
        │                 │                 │
        ├────────────┬────┴────┬────────────┤
        ▼            ▼         ▼            ▼
     Site A      Site B     Site C      Site D
        │            │          │            │
        ▼            ▼          ▼            ▼
      Agents       Agents     Agents       Agents
```

The relay layer operates as one logical communication fabric regardless
of the number or location of physical relay instances.

---

## Related Decisions

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-005 Multi-Site Routing
- ADR-006 Site Manifest
- ADR-007 Zero Trust
- ADR-009 Protocol Versioning

---

## Implementation Notes

The relay architecture should:

- Support horizontal scaling
- Route traffic using logical identities
- Exchange routing metadata between relays
- Detect relay availability automatically
- Support relay failover
- Minimize routing latency
- Preserve secure end-to-end communication
- Expose health and routing metrics

The relay synchronization protocol, topology optimization algorithms,
service discovery mechanism, and deployment model remain implementation
decisions and are outside the scope of this ADR.

---

## Rationale

A logical distributed relay architecture enables BeaconLink to scale beyond a
single deployment while preserving a simple and consistent communication
model.

By treating multiple relay instances as a unified logical platform,
BeaconLink achieves high availability, geographic distribution, improved
fault tolerance, and efficient routing without exposing deployment
complexity to agents or users.

This architecture complements the relay-first networking model,
persistent connections, multi-site routing, and zero-trust security
principles, providing a resilient and scalable communication foundation
for the BeaconLink platform as it grows across regions, organizations, and
deployment environments.
