# 035 - Version Negotiation

> **Document ID:** 035
>
> **Protocol:** BeaconLink Tunnel Protocol (ATP)
>
> **Protocol Version:** 1.0
>
> **Document Version:** 1.0.0
>
> **Status:** Draft
>
> **Last Updated:** 2026-07-16
>
> **Authors:** BeaconLink Engineering Team
>
> **Reviewers:** TBD
>
> **Related Documents**
>
> - 030 BeaconLink Tunnel Protocol
> - 031 Packet Format
> - 032 Message Types
> - 033 State Machine
> - 034 Session Management
> - 036 Error Codes
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Principles
4. Versioning Strategy
5. Version Format
6. Negotiation Workflow
7. Compatibility Rules
8. Capability Negotiation
9. Negotiation Failure
10. Security Considerations
11. Performance Considerations
12. Future Evolution
13. Summary

---

# 1. Introduction

BeaconLink Tunnel Protocol (ATP) is expected to evolve over time.

New protocol versions may introduce additional message types, packet
fields, capabilities, or behaviors.

To maintain interoperability between different software releases,
every ATP connection begins with a Version Negotiation phase.

This process determines the highest protocol version supported by both
the BeaconLink Agent and the BeaconLink Relay.

---

# 2. Purpose

Version Negotiation provides:

- Protocol compatibility
- Backward compatibility
- Controlled protocol evolution
- Feature discovery
- Graceful upgrade paths
- Safe protocol rejection

No ATP session may become active before successful version negotiation.

---

# 3. Design Principles

ATP version negotiation follows these principles.

## Explicit

Every endpoint declares its supported protocol versions.

---

## Deterministic

Both endpoints shall always reach the same negotiated result.

---

## Backward Compatible

New protocol versions should preserve compatibility whenever practical.

---

## Fail Safely

Unsupported protocol combinations shall fail cleanly.

---

## Extensible

Future protocol capabilities shall not require redesign of the
negotiation mechanism.

---

# 4. Versioning Strategy

ATP follows Semantic Versioning principles.

```
MAJOR.MINOR.PATCH
```

Example:

```
1.0.0

1.1.0

1.2.4

2.0.0
```

Meaning:

| Component | Description                      |
| --------- | -------------------------------- |
| MAJOR     | Breaking protocol changes        |
| MINOR     | New backward-compatible features |
| PATCH     | Bug fixes and clarifications     |

Only MAJOR version differences may introduce incompatible behavior.

---

# 5. Version Format

Each endpoint advertises:

- Supported protocol versions
- Preferred protocol version
- Optional capabilities

Example:

```
Agent

Supported Versions

1.0

1.1

1.2
```

```
Relay

Supported Versions

1.1

1.2

2.0
```

Negotiated Version:

```
1.2
```

---

# 6. Negotiation Workflow

Version negotiation occurs immediately after authentication.

```
Transport Connected
    ↓
Authentication
    ↓
VERSION_REQUEST
    ↓
Supported Versions
    ↓
VERSION_RESPONSE
    ↓
Highest Compatible Version
    ↓
Capability Negotiation
    ↓
Session Active
```

If negotiation fails, the connection shall be terminated.

---

# 7. Compatibility Rules

ATP uses the following compatibility policy.

### Patch Versions

Always compatible.

Example:

```
1.0.0

↓

1.0.3
```

---

### Minor Versions

Expected to remain backward compatible.

Example:

```
Agent

1.1

Relay

1.3

↓

Negotiated

1.1
```

---

### Major Versions

Compatibility is not guaranteed.

Example:

```
Agent

1.x

Relay

2.x

↓

Negotiation Failure
```

The Relay may optionally provide upgrade guidance.

---

# 8. Capability Negotiation

Version numbers alone do not describe every protocol feature.

After selecting a protocol version, endpoints exchange supported
capabilities.

Examples:

- Compression
- Datagram support
- Stream priorities
- Connection migration
- Experimental extensions

Example:

```
Protocol Version

↓

1.2

↓

Capabilities

Compression

Heartbeat v2

Multiplexing

↓

Common Capability Set

↓

Session Active
```

Only mutually supported capabilities may be enabled.

---

# 9. Negotiation Failure

Negotiation fails when:

- No compatible protocol version exists.
- Mandatory capabilities are unsupported.
- Protocol information is malformed.
- Negotiation times out.

Failure sequence:

```
VERSION_REQUEST
    ↓
No Compatible Version
    ↓
VERSION_REJECT
    ↓
ERROR
    ↓
DISCONNECT
```

Connections shall fail before any ATP application traffic is exchanged.

---

# 10. Security Considerations

Version negotiation must prevent protocol downgrade attacks.

Requirements include:

- Perform negotiation inside the encrypted tunnel.
- Authenticate both endpoints before negotiation.
- Reject unsupported protocol versions.
- Validate capability advertisements.
- Log negotiation failures.

Endpoints shall never silently downgrade to an insecure protocol version.

---

# 11. Performance Considerations

Version negotiation should:

- Require minimal round trips.
- Exchange compact metadata.
- Complete before tunnel establishment.
- Cache compatible versions where appropriate.

Negotiation should contribute negligible overhead to connection setup.

---

# 12. Future Evolution

Future ATP versions may support:

- Capability extensions
- Experimental protocol flags
- Feature deprecation notices
- Multi-version interoperability
- Vendor-specific extensions
- Dynamic capability updates

Future enhancements should preserve the negotiation framework defined by
this document.

---

# 13. Summary

Version Negotiation enables BeaconLink Tunnel Protocol implementations to
communicate safely across multiple software releases.

By combining semantic versioning, capability discovery, and explicit
compatibility rules, ATP can evolve without sacrificing interoperability
or security.

Every ATP connection shall complete successful version negotiation before
processing application traffic.

> **"Negotiate first. Communicate second."**

---

# Appendix A — Negotiation Sequence Diagram

```
Agent                          Relay
  |                              |
  | CONNECT                      |
  |----------------------------->|
  |                              |
  | AUTH_RESPONSE                |
  |----------------------------->|
  |                              |
  | VERSION_REQUEST              |
  |----------------------------->|
  |                              |
  | VERSION_RESPONSE             |
  |<-----------------------------|
  |                              |
  | CAPABILITY_NEGOTIATION       |
  |<============================>|
  |                              |
  | SESSION_ACTIVE               |
```

---

# Appendix B — Negotiation State Machine

```
CONNECTED
     │
     ▼
AUTHENTICATED
     │
     ▼
NEGOTIATING_VERSION
     │
 ┌───┴────────────┐
 │                │
 ▼                ▼
SUCCESS         FAILURE
 │                │
 ▼                ▼
NEGOTIATING    DISCONNECT
CAPABILITIES
 │
 ▼
ACTIVE
```

---

# Appendix C — Compatibility Matrix

| Agent | Relay | Result          |
| ----- | ----- | --------------- |
| 1.0   | 1.0   | ✅ Compatible   |
| 1.0   | 1.1   | ✅ Use 1.0      |
| 1.1   | 1.3   | ✅ Use 1.1      |
| 1.2   | 2.0   | ❌ Incompatible |
| 2.0   | 2.1   | ✅ Use 2.0      |

---

# Appendix D — Negotiation Decision Flow

```
Receive Supported Versions
            │
            ▼
Find Highest Common Version
            │
     ┌──────┴──────┐
     │             │
     ▼             ▼
Found?           Not Found
     │             │
     ▼             ▼
Exchange        VERSION_REJECT
Capabilities        │
     │              ▼
     ▼         DISCONNECT
All Required?
     │
 ┌───┴───┐
 │       │
 ▼       ▼
Yes      No
 │       │
 ▼       ▼
ACTIVE  ERROR
```

---

# Appendix E — Version Number Examples

| Protocol Release       | Version | Compatibility       |
| ---------------------- | ------- | ------------------- |
| Initial Release        | 1.0.0   | Baseline            |
| Heartbeat Improvements | 1.1.0   | Backward compatible |
| Compression Support    | 1.2.0   | Backward compatible |
| Protocol Redesign      | 2.0.0   | Breaking change     |
| Bug Fixes              | 2.0.1   | Backward compatible |
