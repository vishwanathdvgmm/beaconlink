# 030 - BeaconLink Tunnel Protocol (ATP)

> **Document ID:** 030
>
> **Protocol Name:** BeaconLink Tunnel Protocol
>
> **Protocol Abbreviation:** ATP
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
> - 016 Network Architecture
> - 020 Networking Overview
> - 021 Relay Network
> - 022 Tunnel Routing
> - 031 Packet Format
> - 032 Message Types
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Scope
4. Design Goals
5. Protocol Philosophy
6. Protocol Stack
7. Core Concepts
8. Communication Model
9. Connection Lifecycle
10. Session Model
11. Multiplexing Model
12. Reliability
13. Security
14. Version Compatibility
15. Performance Goals
16. Future Evolution
17. Summary

---

# 1. Introduction

The BeaconLink Tunnel Protocol (ATP) is the primary communication protocol
used between BeaconLink Relay Nodes and BeaconLink Agents.

ATP establishes authenticated, encrypted, persistent tunnels that carry
management traffic, control traffic, monitoring information, and user
application requests.

ATP is designed specifically for BeaconLink and is not intended as a
general-purpose networking protocol.

---

# 2. Purpose

ATP provides a secure and efficient communication channel between BeaconLink
infrastructure and user-owned devices.

The protocol is responsible for:

- Device authentication
- Tunnel establishment
- Session management
- Request forwarding
- Multiplexing
- Health monitoring
- Heartbeats
- Recovery
- Version negotiation

---

# 3. Scope

ATP defines:

- Connection establishment
- Packet structure
- Message types
- Session lifecycle
- Stream multiplexing
- Error handling
- Protocol negotiation

ATP does **not** define:

- HTTP semantics
- TLS implementation
- DNS behavior
- Application execution

Those responsibilities belong to other subsystems.

---

# 4. Design Goals

ATP should be:

- Secure
- Lightweight
- Persistent
- Extensible
- Efficient
- Reliable
- Versioned

The protocol should remain transport-independent wherever practical.

---

# 5. Protocol Philosophy

ATP follows several core principles.

## Long-Lived Connections

Connections remain open for extended periods.

---

## One Tunnel Per Agent

Each Agent maintains one persistent ATP connection.

Multiple applications share this tunnel.

---

## Multiplex Everything

Management traffic and application traffic share the same secure
connection through logical streams.

---

## Zero Trust

Every message originates from an authenticated session.

---

## Backward Compatibility

Protocol evolution should preserve compatibility whenever possible.

---

# 6. Protocol Stack

ATP operates above the secure transport layer.

```
Application
    ↓
HTTP / WebSocket
    ↓
BeaconLink Tunnel Protocol (ATP)
    ↓
TLS 1.3
    ↓
TCP / QUIC
    ↓
Internet
    ↓
Operating System
    ↓
Physical Network
```

ATP assumes the transport layer already provides encrypted communication.

---

# 7. Core Concepts

ATP introduces several protocol concepts.

## Tunnel

A long-lived encrypted communication channel.

---

## Session

Represents an authenticated Agent connection.

---

## Stream

An independent logical communication channel inside a tunnel.

---

## Packet

The smallest protocol unit transmitted over ATP.

---

## Message

A structured protocol payload contained within one or more packets.

---

# 8. Communication Model

ATP uses a persistent bidirectional communication model.

```
BeaconLink Agent
    ⇅
Persistent ATP Tunnel
    ⇅
BeaconLink Relay
```

Both endpoints may initiate protocol messages.

---

# 9. Connection Lifecycle

Every ATP connection follows the same lifecycle.

```
Initialize
    ↓
Connect
    ↓
Authenticate
    ↓
Negotiate Version
    ↓
Create Session
    ↓
Create Tunnel
    ↓
Active
    ↓
Heartbeat
    ↓
Disconnect
    ↓
Reconnect
```

Connections should recover automatically after temporary failures.

---

# 10. Session Model

Each connection contains exactly one authenticated session.

Session responsibilities include:

- Identity
- Authorization
- Heartbeats
- Tunnel ownership
- Stream management

A session terminates when authentication expires or the connection is
closed.

---

# 11. Multiplexing Model

ATP supports multiple independent logical streams.

```
ATP Tunnel
│
├── Control Stream
├── Management Stream
├── Monitoring Stream
├── HTTP Stream #1
├── HTTP Stream #2
├── HTTP Stream #3
└── Future Streams
```

Failure of one stream should not interrupt other streams.

---

# 12. Reliability

ATP provides:

- Ordered delivery
- Duplicate detection
- Connection recovery
- Stream isolation
- Automatic reconnection

Reliability should remain independent from application behavior.

---

# 13. Security

ATP requires:

- Mutual authentication
- TLS 1.3
- Device identity verification
- Session validation
- Replay protection
- Message integrity
- Forward secrecy

ATP never transmits secrets in plaintext.

---

# 14. Version Compatibility

Every connection begins with protocol negotiation.

```
Agent
    ↓
Supported Versions
    ↓
Relay
    ↓
Common Version
    ↓
Session Established
```

Connections should fail gracefully if no compatible version exists.

---

# 15. Performance Goals

ATP should support:

- Persistent tunnels
- Thousands of concurrent streams
- Low protocol overhead
- Fast recovery
- Efficient multiplexing
- Minimal memory usage

Protocol efficiency should scale with the number of connected Agents.

---

# 16. Future Evolution

Future ATP capabilities may include:

- QUIC-native transport
- Datagram support
- Adaptive compression
- Stream prioritization
- Connection migration
- Multi-path networking
- Peer-to-peer extensions

Future versions should preserve protocol compatibility whenever
possible.

---

# 17. Summary

The BeaconLink Tunnel Protocol (ATP) is the communication backbone of the
BeaconLink platform.

It enables secure, authenticated, persistent communication between
BeaconLink Relay Nodes and BeaconLink Agents while supporting multiplexed traffic,
automatic recovery, and protocol evolution.

ATP provides the foundation upon which every networking capability in
BeaconLink is built.

> **"One secure tunnel. Many independent conversations."**

---

# Appendix A — ATP Stack

```
                User Application
                       │
                 HTTP / WebSocket
                       │
              BeaconLink Tunnel Protocol
                       │
                   TLS 1.3
                       │
                 TCP / QUIC
                       │
                  Internet
```

---

# Appendix B — ATP Lifecycle

```
START
    ↓
CONNECTING
    ↓
AUTHENTICATING
    ↓
NEGOTIATING
    ↓
CONNECTED
    ↓
ACTIVE
    ↓
HEARTBEAT
    ↓
RECOVERING
    ↓
ACTIVE
    or
    ↓
DISCONNECTED
```

---

# Appendix C — Tunnel Structure

```
                ATP Tunnel

┌───────────────────────────────────┐
│ Session                           │
│ ├── Control Stream                │
│ ├── Management Stream             │
│ ├── Monitoring Stream             │
│ ├── HTTP Stream                   │
│ ├── HTTP Stream                   │
│ └── HTTP Stream                   │
└───────────────────────────────────┘
```

---

# Appendix D — ATP Responsibilities

| Responsibility      |  ATP  |
| ------------------- | :---: |
| Authentication      |  ✅   |
| Session Management  |  ✅   |
| Multiplexing        |  ✅   |
| Heartbeats          |  ✅   |
| Request Forwarding  |  ✅   |
| Version Negotiation |  ✅   |
| Packet Definition   | ➜ 031 |
| Message Types       | ➜ 032 |
| State Machine       | ➜ 033 |
| Error Codes         | ➜ 036 |
