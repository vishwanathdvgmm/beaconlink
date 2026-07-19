# RFC-0001: Peer-to-Peer Networking

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing optional **Peer-to-Peer (P2P)
communication** between BeaconLink Agents.

Currently, all communication flows through the BeaconLink Relay network,
which provides identity, routing, security, auditing, and NAT traversal.

While the relay-first architecture remains the default networking model,
certain workloads could benefit from direct agent-to-agent
communication.

This RFC proposes allowing authenticated agents to establish secure
peer-to-peer connections whenever policy, topology, and network
conditions permit.

---

# Motivation

The relay-first architecture intentionally prioritizes simplicity,
security, and operational consistency.

However, some workloads generate high volumes of traffic between agents,
including:

- Large artifact distribution
- Site replication
- Backup synchronization
- Log aggregation
- Database replication
- Edge data exchange
- Cluster communication

Routing all traffic through relays may introduce:

- Additional latency
- Increased relay bandwidth
- Higher infrastructure costs
- Regional bottlenecks
- Reduced throughput

Optional peer-to-peer communication could significantly reduce relay
load while improving transfer performance.

---

# Goals

The proposed architecture should:

- Preserve relay-first networking
- Support optional direct communication
- Maintain Zero Trust security
- Authenticate every peer
- Operate behind NAT when possible
- Fall back to relays automatically
- Require no manual network configuration
- Preserve existing APIs

---

# Non-Goals

This RFC does **not** propose:

- Removing relays
- Decentralizing BeaconLink
- Replacing relay routing
- Allowing anonymous peers
- Introducing mesh networking
- Exposing inbound agent ports
- Changing Site Manifest behavior
- Modifying existing APIs

---

# Proposed Design

The relay remains responsible for:

- Authentication
- Identity verification
- Peer discovery
- Session negotiation
- Policy validation

Once two agents are authorized to communicate, the relay may coordinate
direct connection establishment.

Possible connection methods include:

1. Direct local routing
2. NAT traversal
3. IPv6 connectivity
4. Relay-assisted negotiation

If direct communication cannot be established, traffic continues through
the relay transparently.

```
        Agent A
           │
           ▼
        Relay
           │
    Peer Discovery
           │
           ▼
      Session Setup
      /           \
     /             \
Direct Link      Relay Path
     │               │
     ▼               ▼
   Agent B       Agent B
```

The relay remains the authoritative control plane regardless of the
selected data path.

---

# Authentication

Peer-to-peer communication must preserve the BeaconLink Zero Trust model.

Every peer connection must:

- Use mutual authentication
- Validate public identities
- Negotiate encrypted sessions
- Verify authorization policies
- Support certificate or key rotation
- Prevent impersonation
- Support audit logging

No peer communication may bypass BeaconLink identity verification.

---

# Routing

Routing decisions remain controlled by the relay.

The relay determines:

- Whether direct communication is permitted
- Which peers may communicate
- Connection preferences
- Session lifetime
- Routing fallback

Agents never independently discover or trust arbitrary peers.

---

# Security Considerations

The proposal maintains existing security guarantees by ensuring:

- Relay-authorized session establishment
- End-to-end encryption
- Mutual authentication
- Authorization before connection
- Session auditing
- Policy enforcement
- Automatic session expiration

The relay remains responsible for network governance.

---

# Operational Impact

Potential benefits include:

- Lower relay bandwidth
- Reduced latency
- Faster file transfers
- Improved regional performance
- Better scalability
- Lower operational costs

Potential challenges include:

- NAT traversal complexity
- Firewall compatibility
- Session negotiation
- Additional diagnostics
- More complex networking

---

# Compatibility

The proposal is fully backward compatible.

Existing deployments continue operating without modification.

Peer-to-peer networking remains optional.

If unsupported by either endpoint, communication automatically continues
through the relay.

---

# Alternatives Considered

## Relay-Only Communication

Continue routing all traffic through relays.

Pros:

- Simpler architecture
- Easier diagnostics
- Predictable routing

Cons:

- Higher bandwidth costs
- Increased latency
- Relay bottlenecks

---

## Full Mesh Networking

Allow agents to communicate directly without relay coordination.

Rejected because:

- Weak governance
- Complex discovery
- Increased attack surface
- Difficult authorization
- Poor operational visibility

---

## VPN-Based Networking

Require agents to participate in a VPN.

Rejected because:

- Operational complexity
- Platform dependency
- Manual configuration
- Poor scalability

---

# Open Questions

- Which NAT traversal techniques should be supported?
- Should direct communication require administrator approval?
- How should relay-assisted fallback be prioritized?
- Should peer sessions be persistent or ephemeral?
- What telemetry should be collected for peer links?
- Should peer communication be site-local by default?

---

# Related Documents

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-005 Multi-Site Routing
- ADR-007 Zero Trust
- ADR-012 Logical Distributed Relay Architecture

---

# Recommendation

The relay-first networking model should remain the default architecture
for BeaconLink because it provides the strongest operational consistency,
security, and manageability.

Peer-to-peer networking should be introduced only as an **optional data
plane optimization**, with relays continuing to own identity,
authorization, routing decisions, and policy enforcement.

This approach improves performance for bandwidth-intensive workloads
while preserving the architectural principles established by the BeaconLink
ADR series.
