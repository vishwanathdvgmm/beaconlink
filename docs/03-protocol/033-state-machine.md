# 033 - State Machine

> **Document ID:** 033
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
> - 034 Session Management
> - 036 Error Codes
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Principles
4. State Machine Overview
5. Connection States
6. Tunnel States
7. Stream States
8. State Transitions
9. Recovery States
10. Invalid Transitions
11. Timeout Handling
12. Implementation Requirements
13. Future Evolution
14. Summary

---

# 1. Introduction

The BeaconLink Tunnel Protocol (ATP) is a stateful protocol.

Every connection, tunnel, and stream progresses through a well-defined
series of states during its lifetime.

The ATP State Machine defines these states, the transitions between
them, and the conditions required for each transition.

Implementations shall follow this specification to ensure consistent
behavior across all BeaconLink components.

---

# 2. Purpose

The ATP State Machine provides:

- Predictable protocol behavior
- Consistent implementations
- Automatic recovery
- Connection lifecycle management
- Tunnel lifecycle management
- Stream lifecycle management

No ATP component shall invent additional runtime states outside this
specification without protocol version changes.

---

# 3. Design Principles

The ATP State Machine follows these principles.

## Deterministic

The same event shall always produce the same transition.

---

## Observable

Every state transition should be logged for diagnostics.

---

## Recoverable

Recoverable failures should enter RECOVERING instead of DISCONNECTED.

---

## Explicit

No hidden or implicit states are permitted.

---

## Minimal

Only essential states are defined.

---

# 4. State Machine Overview

ATP defines three independent state machines.

```
Connection State
    ↓
Tunnel State
    ↓
Stream State
```

Each operates independently while remaining coordinated.

---

# 5. Connection States

A connection represents the transport channel between an BeaconLink Agent
and an BeaconLink Relay.

## States

```
INITIALIZED
    ↓
CONNECTING
    ↓
CONNECTED
    ↓
AUTHENTICATING
    ↓
NEGOTIATING
    ↓
READY
    ↓
ACTIVE
    ↓
RECOVERING
    ↓
DISCONNECTING
    ↓
DISCONNECTED
```

### INITIALIZED

The connection object has been created but no network activity has
started.

---

### CONNECTING

The Agent is attempting to establish a transport connection.

---

### CONNECTED

A secure transport (TLS/QUIC) has been established.

ATP communication has not yet begun.

---

### AUTHENTICATING

The Agent is proving its identity to the Relay.

---

### NEGOTIATING

Protocol version and capabilities are exchanged.

---

### READY

Authentication and negotiation have completed successfully.

The connection is ready to create tunnels.

---

### ACTIVE

Normal protocol operation.

All ATP messages may be exchanged.

---

### RECOVERING

Temporary failure detected.

Automatic recovery procedures are in progress.

---

### DISCONNECTING

Graceful shutdown initiated.

Outstanding streams should complete where possible.

---

### DISCONNECTED

Connection terminated.

No ATP communication exists.

---

# 6. Tunnel States

Each ATP session owns one logical tunnel.

```
CREATED
    ↓
OPENING
    ↓
ESTABLISHED
    ↓
ACTIVE
    ↓
IDLE
    ↓
RECOVERING
    ↓
CLOSING
    ↓
CLOSED
```

### CREATED

Tunnel metadata allocated.

---

### OPENING

Tunnel establishment messages exchanged.

---

### ESTABLISHED

Tunnel successfully created.

---

### ACTIVE

Traffic currently flowing.

---

### IDLE

Tunnel connected but temporarily inactive.

---

### RECOVERING

Tunnel restoration in progress.

---

### CLOSING

Graceful shutdown initiated.

---

### CLOSED

Tunnel no longer exists.

---

# 7. Stream States

Each application request executes inside an ATP stream.

```
NEW
    ↓
OPEN
    ↓
ACTIVE
    ↓
HALF_CLOSED_LOCAL
    ↓
HALF_CLOSED_REMOTE
    ↓
CLOSING
    ↓
CLOSED
    ↓
RESET
```

### NEW

Stream allocated.

---

### OPEN

Ready to exchange messages.

---

### ACTIVE

Application data flowing.

---

### HALF_CLOSED_LOCAL

Local endpoint finished sending.

Receiving remains allowed.

---

### HALF_CLOSED_REMOTE

Remote endpoint finished sending.

Local endpoint may still send.

---

### CLOSING

Graceful stream shutdown.

---

### CLOSED

Stream resources released.

---

### RESET

Abnormal termination.

---

# 8. State Transitions

## Connection Lifecycle

```
INITIALIZED
    ↓
CONNECTING
    ↓
CONNECTED
    ↓
AUTHENTICATING
    ↓
NEGOTIATING
    ↓
READY
    ↓
ACTIVE
    ↓
DISCONNECTING
    ↓
DISCONNECTED
```

---

## Recovery Path

```
ACTIVE
    ↓
Network Failure
    ↓
RECOVERING
    ↓
Reconnect
    ↓
CONNECTED
    ↓
AUTHENTICATING
    ↓
NEGOTIATING
    ↓
READY
    ↓
ACTIVE
```

---

## Tunnel Lifecycle

```
CREATED
    ↓
OPENING
    ↓
ESTABLISHED
    ↓
ACTIVE
    ↓
IDLE
    ↓
ACTIVE
    ↓
CLOSING
    ↓
CLOSED
```

---

# 9. Recovery States

ATP should distinguish temporary failures from permanent failures.

Recoverable examples:

- Packet loss
- Network interruption
- Relay restart
- Temporary authentication timeout

Recovery flow:

```
ACTIVE
    ↓
Failure
    ↓
RECOVERING
    ↓
Recovery Successful?
    ↓
YES
    ↓
ACTIVE
    ↓
NO
    ↓
DISCONNECTED
```

---

# 10. Invalid Transitions

The following transitions are prohibited.

```
INITIALIZED
    ↓
ACTIVE
```

Authentication cannot be skipped.

---

```
CONNECTED
    ↓
ACTIVE
```

Version negotiation must occur first.

---

```
CLOSED
    ↓
ACTIVE
```

Closed streams cannot be reused.

---

```
DISCONNECTED
    ↓
READY
```

A new connection must be established.

Invalid transitions shall generate protocol errors.

---

# 11. Timeout Handling

Each state may define timeout limits.

Examples:

| State          |  Example Timeout |
| -------------- | ---------------: |
| CONNECTING     |             30 s |
| AUTHENTICATING |             15 s |
| NEGOTIATING    |             10 s |
| ACTIVE         | Heartbeat-driven |
| RECOVERING     |     Configurable |

Timeout values may be configurable by implementation.

---

# 12. Implementation Requirements

ATP implementations shall:

- Validate every state transition
- Reject invalid transitions
- Preserve state consistency
- Log state changes
- Prevent race conditions
- Recover automatically when possible

Implementations should avoid duplicated state logic.

---

# 13. Future Evolution

Future ATP versions may introduce:

- Multi-path connections
- Parallel tunnels
- Priority streams
- Datagram streams
- Peer-to-peer states
- Cluster synchronization states

New states should preserve backward compatibility.

---

# 14. Summary

The ATP State Machine defines the legal lifecycle of every connection,
tunnel, and stream within the BeaconLink Tunnel Protocol.

By standardizing protocol behavior through explicit states and
deterministic transitions, ATP implementations remain predictable,
interoperable, and resilient to failure.

Every ATP implementation shall treat this state machine as the canonical
definition of protocol behavior.

> **"A protocol is reliable only when every implementation follows the
> same state machine."**

---

# Appendix A — Complete Connection State Machine

```
INITIALIZED
      │
      ▼
CONNECTING
      │
      ▼
CONNECTED
      │
      ▼
AUTHENTICATING
      │
      ▼
NEGOTIATING
      │
      ▼
READY
      │
      ▼
ACTIVE
      │
 ┌────┴────┐
 │         │
 ▼         ▼
RECOVERING DISCONNECTING
 │         │
 ▼         ▼
ACTIVE DISCONNECTED
```

---

# Appendix B — Tunnel State Machine

```
CREATED
    │
    ▼
OPENING
    │
    ▼
ESTABLISHED
    │
    ▼
ACTIVE
    │
 ┌──┴──┐
 │     │
 ▼     ▼
IDLE RECOVERING
 │     │
 ▼     ▼
ACTIVE ACTIVE
 │
 ▼
CLOSING
 │
 ▼
CLOSED
```

---

# Appendix C — Stream State Machine

```
NEW
 │
 ▼
OPEN
 │
 ▼
ACTIVE
 │
 ├─────────────┐
 ▼             ▼
HALF_LOCAL HALF_REMOTE
 │             │
 └──────┬──────┘
        ▼
     CLOSING
        │
        ▼
      CLOSED

Error
  │
  ▼
RESET
```

---

# Appendix D — State Transition Matrix

| Current State  | Event            | Next State     |
| -------------- | ---------------- | -------------- |
| INITIALIZED    | Connect          | CONNECTING     |
| CONNECTING     | Transport Ready  | CONNECTED      |
| CONNECTED      | AUTH_REQUEST     | AUTHENTICATING |
| AUTHENTICATING | AUTH_SUCCESS     | NEGOTIATING    |
| NEGOTIATING    | Version Agreed   | READY          |
| READY          | Tunnel Created   | ACTIVE         |
| ACTIVE         | Failure          | RECOVERING     |
| RECOVERING     | Recovery Success | ACTIVE         |
| RECOVERING     | Recovery Failed  | DISCONNECTED   |
| ACTIVE         | Disconnect       | DISCONNECTING  |
| DISCONNECTING  | Complete         | DISCONNECTED   |

---

# Appendix E — Implementation Notes

- State transitions should be atomic.
- Concurrent transitions for the same object must be serialized.
- Every transition should emit an internal event for logging and observability.
- Recovery logic should be idempotent where practical.
- State changes should be exposed through metrics to aid monitoring and debugging.
