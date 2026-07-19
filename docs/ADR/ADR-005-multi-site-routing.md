# ADR-005: Multi-Site Routing

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Distributed Networking Architecture

---

## Context

BeaconLink is designed to manage infrastructure deployed across multiple
independent sites, including:

- Data centers
- Branch offices
- Edge locations
- Cloud regions
- Private clouds
- Customer-managed environments

Each site may have independent networking policies, address spaces,
security boundaries, and relay deployments.

The platform requires a routing architecture that enables communication
between services and agents without assuming global network
connectivity or flat network topology.

Routing decisions should remain independent of physical network layout
and support future expansion without requiring architectural changes.

---

## Decision

BeaconLink adopts a **logical multi-site routing architecture**.

Every managed resource belongs to a logical site, and all communication
is routed using site identity rather than physical network location.

Relay services exchange routing information that allows requests to be
forwarded to the appropriate destination site while preserving security
boundaries and routing policies.

Site-aware routing becomes the authoritative communication model
throughout the platform.

---

## Alternatives Considered

### Flat Global Network

Treat every managed component as part of a single network.

**Rejected because:**

- Poor scalability
- Difficult security isolation
- Increased routing complexity
- Unsuitable for disconnected environments
- Operationally inflexible

---

### Manual Route Configuration

Administrators configure communication paths manually.

**Rejected because:**

- High operational overhead
- Error-prone configuration
- Difficult to maintain
- Limited scalability
- Poor automation support

---

### Network-Based Routing

Route traffic using IP topology and network addressing.

**Rejected because:**

- Tight coupling to infrastructure
- Difficult across NAT boundaries
- Inconsistent addressing
- Limited portability
- Reduced architectural flexibility

---

## Consequences

### Positive

- Scalable multi-site architecture
- Clear administrative boundaries
- Improved routing flexibility
- Network-independent communication
- Better fault isolation
- Simplified expansion
- Consistent routing behavior
- Supports hybrid infrastructure

### Negative

- Routing metadata required
- Additional routing layer
- Route synchronization required
- Increased relay coordination
- More complex operational visibility

---

## Architecture

```
                  BeaconLink Console
                        │
                        ▼
                 BeaconLink Control Plane
                        │
                        ▼
                  Global Routing
                        │
        ┌───────────────┼───────────────┐
        ▼               ▼               ▼
     Site A          Site B          Site C
        │               │               │
     Relay A         Relay B         Relay C
        │               │               │
        ▼               ▼               ▼
      Agents          Agents          Agents
```

Communication is routed according to logical site identity rather than
physical network topology.

---

## Related Decisions

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-006 Site Manifest
- ADR-007 Zero Trust
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The routing architecture should:

- Assign every managed resource to a site
- Support dynamic route discovery
- Route using logical identities
- Preserve site isolation
- Handle relay failover
- Support route propagation
- Detect unreachable destinations
- Expose routing metrics and health information

Specific routing algorithms, protocols, and optimization strategies are
implementation decisions and remain outside the scope of this ADR.

---

## Rationale

A logical multi-site routing architecture enables BeaconLink to manage
distributed infrastructure without depending on physical network
topology or direct connectivity.

By treating sites as first-class architectural entities and routing
traffic according to logical identities, BeaconLink achieves scalable,
secure, and resilient communication across geographically distributed
environments. This approach aligns with the relay-first networking
model, supports hybrid and disconnected deployments, and provides a
stable foundation for future expansion without introducing network-
specific assumptions into higher-level platform services.
