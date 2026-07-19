# ADR-009: Protocol Versioning

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Communication Protocol Evolution

---

## Context

BeaconLink consists of independently deployable components including:

- Agents
- Relays
- API services
- Console services
- Deployment services
- Future platform components

These components communicate using well-defined protocols across
potentially different software versions.

Since upgrades may occur incrementally across distributed environments,
different versions of BeaconLink components will temporarily coexist during
deployment, rollback, disaster recovery, and staged upgrades.

The platform therefore requires a protocol versioning strategy that
supports compatibility while allowing the communication protocol to
evolve over time.

---

## Decision

BeaconLink adopts **explicit protocol versioning** for all externally visible
communication protocols.

Every protocol implementation shall expose its supported protocol
version(s) during connection establishment.

Communication partners negotiate a mutually supported protocol version
before exchanging operational messages.

Protocol evolution shall preserve backward compatibility whenever
practical, with incompatible changes introduced only through new major
protocol versions.

Protocol versions are independent of software release versions.

---

## Alternatives Considered

### No Protocol Versioning

Assume every deployed component runs the same software version.

**Rejected because:**

- Prevents rolling upgrades
- Breaks mixed-version deployments
- Complicates rollback
- Reduces operational flexibility
- Poor support for distributed environments

---

### Software Version Matching

Use application version numbers as protocol versions.

**Rejected because:**

- Couples protocol evolution to releases
- Limits independent protocol improvements
- Makes compatibility difficult
- Complicates long-term maintenance
- Increases upgrade constraints

---

### Multiple Independent Protocols

Create a new communication protocol for every major release.

**Rejected because:**

- High maintenance burden
- Fragmented implementations
- Increased testing effort
- Poor interoperability
- Difficult migration

---

## Consequences

### Positive

- Rolling upgrades supported
- Backward compatibility
- Forward protocol evolution
- Simplified rollback
- Stable communication contracts
- Independent software releases
- Improved interoperability
- Reduced deployment risk

### Negative

- Compatibility testing required
- Multiple protocol versions may coexist
- Version negotiation logic required
- Deprecated versions require lifecycle management
- Increased protocol documentation

---

## Architecture

```
        Agent
          │
          ▼
 Advertise Supported Versions
          │
          ▼
        Relay
          │
          ▼
Version Negotiation
          │
          ▼
Selected Protocol Version
          │
          ▼
Operational Communication
```

Protocol negotiation occurs before operational communication begins.

---

## Related Decisions

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-003 Public-Key Authentication
- ADR-010 Update System
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The protocol architecture should:

- Define explicit protocol versions
- Support version negotiation
- Allow multiple supported versions where practical
- Maintain documented compatibility guarantees
- Clearly identify deprecated protocol versions
- Validate negotiated capabilities
- Reject unsupported protocol combinations
- Record protocol versions for operational diagnostics

Specific message formats, serialization mechanisms, negotiation
algorithms, and transport protocols remain implementation decisions and
are outside the scope of this ADR.

---

## Rationale

Explicit protocol versioning enables BeaconLink components to evolve
independently while maintaining reliable communication across mixed
software versions.

By separating protocol compatibility from software release cycles,
BeaconLink supports rolling upgrades, staged deployments, rollback
operations, and long-term platform evolution without requiring every
component to be upgraded simultaneously. This approach improves
operational flexibility while preserving stable communication contracts
throughout the distributed platform.
