# RFC-0006: High Availability Relays

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing **High Availability (HA) Relay Clusters**
to eliminate single-relay failure domains within a site or region.

While the existing BeaconLink architecture supports multiple distributed
relays, individual relay deployments may still represent localized
points of failure. This RFC introduces active-active relay clusters that
operate as a single logical endpoint, allowing agent connections,
routing, and session management to continue seamlessly when individual
relay instances become unavailable.

The objective is to increase platform availability without changing the
existing relay-first networking model.

---

# Motivation

BeaconLink Relays are responsible for:

- Agent connectivity
- Authentication
- Session management
- Message routing
- Event forwarding
- Health reporting

Although multiple relay deployments can exist across regions, an
individual site may still rely on a small number of relay instances.

Relay outages may result from:

- Hardware failures
- Operating system failures
- Network interruptions
- Software defects
- Planned maintenance
- Infrastructure upgrades

Without local redundancy, agents connected to an unavailable relay must
reconnect elsewhere, potentially increasing recovery time.

High Availability Relay Clusters reduce service interruption while
improving operational resilience.

---

# Goals

The proposed architecture should:

- Eliminate single relay failure domains
- Support active-active relay operation
- Allow seamless failover
- Preserve persistent agent connections
- Scale horizontally
- Support rolling upgrades
- Minimize operational downtime
- Remain transparent to agents

---

# Non-Goals

This RFC does **not** propose:

- Replacing relay federation
- Eliminating relay-first networking
- Global consensus across all relays
- Stateful distributed databases
- Session migration across organizations
- Changes to agent authentication
- New networking protocols
- Modifying Site Manifest behavior

---

# Proposed Design

A relay deployment consists of multiple relay instances operating behind
a shared logical endpoint.

Agents establish outbound connections using the cluster endpoint rather
than individual relay addresses.

```
               Agents
                  │
                  ▼
         Logical Relay Endpoint
                  │
      ┌───────────┼───────────┐
      ▼           ▼           ▼
   Relay A     Relay B     Relay C
      │           │           │
      └───────────┼───────────┘
                  ▼
         Shared Routing State
```

Relay instances cooperate to provide continuous service while remaining
independently replaceable.

---

# Cluster Responsibilities

The relay cluster should provide:

- Connection distribution
- Session routing
- Health monitoring
- Traffic balancing
- Failure detection
- Rolling maintenance
- Capacity scaling
- Metrics aggregation

Clients interact only with the logical relay endpoint.

---

# Session Management

The relay cluster should support resilient session handling.

Possible approaches include:

- Stateless routing with reconnect
- Replicated session metadata
- Shared distributed state
- Sticky session routing

The implementation should minimize disruption while avoiding excessive
coordination overhead.

Session management strategy remains an implementation decision.

---

# Failure Handling

The cluster should automatically detect:

- Relay failures
- Network partitions
- Health degradation
- Resource exhaustion

Upon failure:

1. Traffic is redirected.
2. Unhealthy relays are removed from service.
3. Existing sessions recover automatically where possible.
4. New connections avoid failed instances.
5. Recovered relays rejoin after validation.

No manual intervention should be required for routine failures.

---

# Scaling

Relay clusters should support horizontal scaling by adding or removing
relay instances without interrupting service.

Scaling operations should support:

- Capacity expansion
- Maintenance windows
- Infrastructure replacement
- Regional growth
- Rolling upgrades

Agents should not require reconfiguration when cluster membership
changes.

---

# Security Considerations

All relay instances remain subject to BeaconLink Zero Trust principles.

Each relay must:

- Possess a unique identity
- Authenticate with peer relays
- Encrypt relay communication
- Validate routing information
- Record audit events
- Enforce authorization policies

High availability must not weaken security guarantees.

---

# Operational Impact

Potential benefits include:

- Improved service availability
- Reduced downtime
- Seamless maintenance
- Horizontal scalability
- Better fault tolerance
- Improved operational resilience
- Simplified infrastructure upgrades

Potential challenges include:

- Cluster coordination
- Session synchronization
- Additional infrastructure
- Health monitoring complexity
- Operational diagnostics

---

# Compatibility

The proposal is fully compatible with the existing BeaconLink architecture.

Single-relay deployments remain supported.

Relay clustering is optional and may be enabled only where high
availability requirements justify the additional operational complexity.

---

# Alternatives Considered

## Single Relay Per Site

Continue operating one relay instance.

Pros:

- Simple deployment
- Minimal infrastructure

Cons:

- Single failure domain
- Maintenance downtime
- Limited scalability

---

## Active-Passive Relays

Maintain standby relay instances.

Pros:

- Simpler synchronization
- Predictable failover

Cons:

- Idle capacity
- Slower recovery
- Limited horizontal scaling

---

## External Load Balancer Only

Use infrastructure load balancers without relay coordination.

Rejected because:

- Limited session awareness
- No application-level health
- Weak routing intelligence
- Reduced operational visibility

---

# Open Questions

- Should relay clusters replicate session metadata?
- What level of state sharing is necessary?
- How should split-brain scenarios be handled?
- Which load-balancing strategy should be preferred?
- Should relay clusters support automatic scaling?
- How should rolling upgrades preserve active sessions?

---

# Related Documents

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-005 Multi-Site Routing
- ADR-007 Zero Trust
- ADR-012 Logical Distributed Relay Architecture
- RFC-0002 Global Relay Network

---

# Recommendation

BeaconLink should introduce High Availability Relay Clusters as an optional
deployment model for environments requiring increased resilience and
continuous availability.

The preferred architecture is an **active-active relay cluster**
presenting a single logical endpoint, allowing relay instances to be
added, removed, upgraded, or replaced without disrupting platform
operations.

This proposal builds upon the existing logical distributed relay
architecture by improving local fault tolerance while remaining fully
compatible with relay federation, relay-first networking, and the
platform's Zero Trust security model.
