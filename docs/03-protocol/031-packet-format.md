# 031 - Packet Format

> **Document ID:** 031
>
> **Protocol:** BeaconLink Tunnel Protocol (ATP)
>
> **Version:** 1.0
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
> - 032 Message Types
> - 033 State Machine
> - 034 Session Management
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Objectives
4. Packet Philosophy
5. Packet Structure
6. Header Format
7. Payload Format
8. Packet Types
9. Packet Lifecycle
10. Fragmentation
11. Validation
12. Security
13. Performance
14. Future Evolution
15. Summary

---

# 1. Introduction

Every communication within the BeaconLink Tunnel Protocol (ATP) is transmitted
using packets.

A packet is the smallest transferable protocol unit exchanged between an
BeaconLink Agent and an BeaconLink Relay.

Packets encapsulate protocol metadata together with an application or
control payload.

---

# 2. Purpose

The ATP packet format provides:

- Protocol identification
- Version compatibility
- Routing information
- Session association
- Stream multiplexing
- Payload transport
- Integrity validation

Every ATP message is transported inside one or more packets.

---

# 3. Design Objectives

The packet format should be:

- Compact
- Efficient
- Extensible
- Versioned
- Easy to parse
- Transport independent

Header fields should remain fixed whenever possible.

---

# 4. Packet Philosophy

ATP packets follow several principles.

## Fixed Header

Every packet begins with a fixed-length header.

---

## Variable Payload

Payload size depends on message type.

---

## Explicit Versioning

Every packet declares its protocol version.

---

## Independent Parsing

Each packet can be parsed without requiring previous packets.

---

## Forward Compatibility

Unknown optional fields should be ignored when supported by protocol
rules.

---

# 5. Packet Structure

An ATP packet consists of four logical sections.

```
+------------------------------------------------+
| Fixed Header                                   |
+------------------------------------------------+
| Optional Extension Header                      |
+------------------------------------------------+
| Payload                                        |
+------------------------------------------------+
| Integrity Trailer (Future Extension)           |
+------------------------------------------------+
```

Only the Fixed Header is mandatory.

---

# 6. Header Format

ATP v1 defines the following logical header.

| Field           | Size     | Description               |
| --------------- | -------- | ------------------------- |
| Magic Number    | 4 bytes  | ATP protocol identifier   |
| Version         | 2 bytes  | Protocol version          |
| Header Length   | 2 bytes  | Header size               |
| Packet Type     | 2 bytes  | Packet classification     |
| Flags           | 2 bytes  | Protocol flags            |
| Session ID      | 16 bytes | Active session identifier |
| Stream ID       | 8 bytes  | Logical stream            |
| Sequence Number | 8 bytes  | Packet ordering           |
| Payload Length  | 4 bytes  | Payload size              |
| Reserved        | 16 bytes | Future compatibility      |

The exact binary layout is defined by the reference implementation.

---

## Example Logical Header

```
Magic
    ↓
Version
    ↓
Header Length
    ↓
Packet Type
    ↓
Flags
    ↓
Session ID
    ↓
Stream ID
    ↓
Sequence Number
    ↓
Payload Length
    ↓
Reserved
```

---

# 7. Payload Format

The payload contains protocol-specific data.

Examples:

- Authentication
- Heartbeat
- HTTP Request
- HTTP Response
- Tunnel Control
- Monitoring
- Configuration

Payload interpretation depends on Packet Type.

---

# 8. Packet Types

ATP defines multiple packet categories.

| Type           | Purpose               |
| -------------- | --------------------- |
| Control        | Connection management |
| Authentication | Identity verification |
| Management     | Configuration         |
| Data           | Application traffic   |
| Monitoring     | Metrics               |
| Heartbeat      | Keepalive             |
| Error          | Protocol failures     |

Additional packet types may be introduced in future protocol versions.

---

# 9. Packet Lifecycle

Every packet follows the same lifecycle.

```
Create
    ↓
Serialize
    ↓
Encrypt
    ↓
Transmit
    ↓
Receive
    ↓
Decrypt
    ↓
Validate
    ↓
Deserialize
    ↓
Process
```

Packets failing validation must be discarded.

---

# 10. Fragmentation

ATP prefers transmitting complete packets.

Large payloads may require fragmentation.

```
Large Message

↓

Fragment

↓

Packet 1

Packet 2

Packet 3

↓

Transmit

↓

Reassemble

↓

Process
```

Fragmentation should remain transparent to higher protocol layers.

---

# 11. Validation

Before processing, every packet should pass validation.

Validation includes:

- Magic Number
- Version
- Header Length
- Payload Length
- Session ID
- Stream ID
- Packet Type

Invalid packets should never reach application logic.

---

# 12. Security

Packet processing should include:

- Integrity verification
- Session validation
- Replay protection
- Authentication checks
- Bounds checking

Malformed packets should be rejected immediately.

---

# 13. Performance

Packet processing should:

- Minimize allocations
- Support zero-copy parsing where practical
- Avoid unnecessary serialization
- Process packets efficiently

Packet parsing should not become the protocol bottleneck.

---

# 14. Future Evolution

Future protocol versions may introduce:

- Optional extension headers
- Compression metadata
- Stream priorities
- Datagram packets
- Packet signatures
- Checksum extensions

New fields should preserve backward compatibility whenever possible.

---

# 15. Summary

The ATP packet format defines the standard transport unit used by the
BeaconLink Tunnel Protocol.

Its fixed header, version-aware design, and extensible payload structure
provide a stable foundation for protocol evolution while maintaining
efficient communication between BeaconLink Agents and Relay Nodes.

Every ATP message ultimately becomes one or more ATP packets.

> **"Packets carry bytes. Protocols carry meaning."**

---

# Appendix A — Logical Packet Layout

```
+------------------------------------------------------------+
| Magic Number                                               |
+------------------------------------------------------------+
| Version | Header Length | Packet Type | Flags              |
+------------------------------------------------------------+
| Session ID                                                 |
+------------------------------------------------------------+
| Stream ID                                                  |
+------------------------------------------------------------+
| Sequence Number                                            |
+------------------------------------------------------------+
| Payload Length                                             |
+------------------------------------------------------------+
| Reserved                                                   |
+------------------------------------------------------------+
| Payload                                                    |
+------------------------------------------------------------+
```

---

# Appendix B — Packet Processing Pipeline

```
Receive Packet
    ↓
Read Header
    ↓
Validate Header
    ↓
Validate Session
    ↓
Read Payload
    ↓
Deserialize
    ↓
Dispatch Message
    ↓
Application Logic
```

---

# Appendix C — Packet State Machine

```
CREATED
    ↓
SERIALIZED
    ↓
ENCRYPTED
    ↓
TRANSMITTED
    ↓
RECEIVED
    ↓
DECRYPTED
    ↓
VALIDATED
    ↓
PROCESSED
    ↓
DESTROYED
```

---

# Appendix D — Packet Relationships

```
Application Message
    ↓
ATP Message
    ↓
ATP Packet
    ↓
TLS Record
    ↓
TCP / QUIC Frame
    ↓
Network
```

---

# Appendix E — Reserved Header Fields

The Reserved field exists to allow protocol evolution without changing
the fixed header layout.

Future protocol versions may use this space for:

- Compression flags
- Priority hints
- Routing extensions
- Datagram support
- Experimental features

Unused reserved bytes shall be transmitted as zero and ignored by
receivers unless defined by a future ATP specification.
