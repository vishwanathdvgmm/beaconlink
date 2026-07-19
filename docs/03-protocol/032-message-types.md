# 032 - Message Types

> **Document ID:** 032
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
> - 033 State Machine
> - 034 Session Management
> - 036 Error Codes
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Principles
4. Message Architecture
5. Message Categories
6. Connection Messages
7. Authentication Messages
8. Tunnel Messages
9. Application Messages
10. Monitoring Messages
11. Management Messages
12. Error Messages
13. Message Lifecycle
14. Security
15. Future Evolution
16. Summary

---

# 1. Introduction

Message Types define the operations performed by the BeaconLink Tunnel
Protocol.

Every ATP packet contains exactly one protocol message.

The receiving endpoint determines the required action by interpreting
the Message Type.

---

# 2. Purpose

ATP messages provide a standardized method for communication between
BeaconLink Agents and Relay Nodes.

They define:

- Connection establishment
- Authentication
- Tunnel management
- Request forwarding
- Monitoring
- Configuration
- Error reporting

---

# 3. Design Principles

ATP messages follow these principles.

## Explicit

Every message performs one well-defined operation.

---

## Stateless

Whenever practical, messages should contain all information required for
processing.

---

## Versioned

Messages evolve through protocol versioning.

---

## Extensible

New message types should not break existing implementations.

---

# 4. Message Architecture

Every ATP message contains:

```
Packet
    ↓
Header
    ↓
Message Type
    ↓
Message Payload
    ↓
Process
```

Message payload interpretation depends entirely on the Message Type.

---

# 5. Message Categories

ATP groups messages into logical categories.

| Category       | Purpose               |
| -------------- | --------------------- |
| Connection     | Connection management |
| Authentication | Identity verification |
| Tunnel         | Tunnel operations     |
| Application    | User traffic          |
| Monitoring     | Health reporting      |
| Management     | Configuration         |
| Error          | Failure reporting     |

---

# 6. Connection Messages

Connection messages manage the protocol lifecycle.

## CONNECT

Initiates a new ATP connection.

---

## CONNECT_ACK

Confirms successful connection establishment.

---

## DISCONNECT

Gracefully terminates a connection.

---

## DISCONNECT_ACK

Acknowledges connection termination.

---

## KEEPALIVE

Maintains an active connection.

---

# 7. Authentication Messages

Authentication messages verify identity.

## AUTH_REQUEST

Requests Agent authentication.

---

## AUTH_RESPONSE

Returns authentication credentials.

---

## AUTH_SUCCESS

Authentication completed successfully.

---

## AUTH_FAILURE

Authentication failed.

---

## TOKEN_REFRESH

Renews session credentials.

---

# 8. Tunnel Messages

Tunnel messages manage ATP tunnels.

## TUNNEL_OPEN

Create a tunnel.

---

## TUNNEL_READY

Tunnel established.

---

## TUNNEL_CLOSE

Terminate tunnel.

---

## TUNNEL_STATUS

Report tunnel health.

---

## TUNNEL_RECOVER

Recover an interrupted tunnel.

---

# 9. Application Messages

Application messages transport user traffic.

## REQUEST

Forward an application request.

---

## RESPONSE

Return an application response.

---

## STREAM_OPEN

Create a logical stream.

---

## STREAM_CLOSE

Close a logical stream.

---

## STREAM_RESET

Abort a stream.

---

# 10. Monitoring Messages

Monitoring messages provide operational visibility.

## HEARTBEAT

Verify endpoint availability.

---

## HEALTH_REPORT

Transmit runtime health.

---

## METRICS

Report operational metrics.

---

## LOG_EVENT

Forward structured log entries.

---

# 11. Management Messages

Management messages control platform behaviour.

## CONFIG_UPDATE

Update configuration.

---

## CONFIG_RESPONSE

Configuration acknowledgement.

---

## DEPLOY

Deploy an application.

---

## UPDATE

Software update notification.

---

## RESTART

Restart Agent components.

---

# 12. Error Messages

Error handling uses dedicated protocol messages.

## ERROR

Generic protocol error.

---

## WARNING

Recoverable condition.

---

## REJECT

Message rejected.

---

## RETRY

Request retransmission.

---

Specific error codes are defined in
**036-error-codes.md**.

---

# 13. Message Lifecycle

Every ATP message follows the same lifecycle.

```
Create
    ↓
Serialize
    ↓
Packetize
    ↓
Encrypt
    ↓
Transmit
    ↓
Receive
    ↓
Validate
    ↓
Deserialize
    ↓
Dispatch
    ↓
Process
    ↓
Complete
```

---

# 14. Security

All ATP messages must satisfy:

- Session validation
- Authentication verification
- Message integrity
- Replay protection
- Authorization

Messages failing validation shall be discarded.

---

# 15. Future Evolution

Future protocol versions may introduce additional message categories.

Examples:

- Peer-to-peer negotiation
- Plugin communication
- Service discovery
- Edge synchronization
- Cluster coordination
- AI diagnostics

New messages should preserve compatibility with previous ATP versions.

---

# 16. Summary

ATP Message Types define every operation performed between BeaconLink Agents
and Relay Nodes.

Each message performs one clearly defined responsibility, enabling ATP to
remain modular, extensible, and easy to implement.

By separating transport from semantics, BeaconLink ensures protocol evolution
without unnecessary complexity.

> **"Packets move data. Messages express intent."**

---

# Appendix A — Complete Message Catalog

| Category       | Message         |
| -------------- | --------------- |
| Connection     | CONNECT         |
| Connection     | CONNECT_ACK     |
| Connection     | DISCONNECT      |
| Connection     | DISCONNECT_ACK  |
| Connection     | KEEPALIVE       |
| Authentication | AUTH_REQUEST    |
| Authentication | AUTH_RESPONSE   |
| Authentication | AUTH_SUCCESS    |
| Authentication | AUTH_FAILURE    |
| Authentication | TOKEN_REFRESH   |
| Tunnel         | TUNNEL_OPEN     |
| Tunnel         | TUNNEL_READY    |
| Tunnel         | TUNNEL_CLOSE    |
| Tunnel         | TUNNEL_STATUS   |
| Tunnel         | TUNNEL_RECOVER  |
| Application    | REQUEST         |
| Application    | RESPONSE        |
| Application    | STREAM_OPEN     |
| Application    | STREAM_CLOSE    |
| Application    | STREAM_RESET    |
| Monitoring     | HEARTBEAT       |
| Monitoring     | HEALTH_REPORT   |
| Monitoring     | METRICS         |
| Monitoring     | LOG_EVENT       |
| Management     | CONFIG_UPDATE   |
| Management     | CONFIG_RESPONSE |
| Management     | DEPLOY          |
| Management     | UPDATE          |
| Management     | RESTART         |
| Error          | ERROR           |
| Error          | WARNING         |
| Error          | REJECT          |
| Error          | RETRY           |

---

# Appendix B — Communication Flow

```
Agent                         Relay

CONNECT
──────────────────────────────►

CONNECT_ACK
◄──────────────────────────────

AUTH_REQUEST
◄──────────────────────────────

AUTH_RESPONSE
──────────────────────────────►

AUTH_SUCCESS
◄──────────────────────────────

TUNNEL_OPEN
──────────────────────────────►

TUNNEL_READY
◄──────────────────────────────

REQUEST
◄──────────────────────────────

RESPONSE
──────────────────────────────►
```

---

# Appendix C — Message Classification

```
ATP Messages
│
├── Connection
├── Authentication
├── Tunnel
├── Application
├── Monitoring
├── Management
└── Error
```

---

# Appendix D — Message Processing Pipeline

```
Receive Packet
    ↓
Decode Header
    ↓
Identify Message Type
    ↓
Validate Session
    ↓
Dispatch Handler
    ↓
Execute
    ↓
Generate Response
```
