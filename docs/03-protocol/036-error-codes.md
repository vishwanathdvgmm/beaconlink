# 036 - Error Codes

> **Document ID:** 036
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
> - 035 Version Negotiation
> - ADR-009 Protocol Versioning

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Principles
4. Error Model
5. Error Categories
6. Connection Errors
7. Authentication Errors
8. Session Errors
9. Tunnel Errors
10. Stream Errors
11. Protocol Errors
12. Internal Errors
13. Error Processing
14. Security Considerations
15. Future Evolution
16. Summary

---

# 1. Introduction

The BeaconLink Tunnel Protocol (ATP) defines a standardized error model to
ensure consistent behavior across all protocol implementations.

Errors communicate abnormal conditions between BeaconLink Agents and Relay
Nodes while allowing automated recovery whenever possible.

Every ATP error consists of:

- Error Code
- Error Name
- Error Category
- Human-readable description

---

# 2. Purpose

ATP Error Codes provide:

- Consistent protocol behavior
- Predictable recovery
- Easier debugging
- Structured logging
- Cross-language compatibility

Implementations should never invent undocumented protocol error codes.

---

# 3. Design Principles

ATP errors shall be:

- Stable
- Machine-readable
- Human-readable
- Versioned
- Extensible

Error codes should never change meaning once published.

---

# 4. Error Model

Every ATP error contains:

```
Error Code
    ↓
Error Name
    ↓
Category
    ↓
Description
    ↓
Optional Metadata
```

Example:

```
0x1002

AUTHENTICATION_FAILED

Authentication

Device authentication failed.
```

---

# 5. Error Categories

ATP groups errors into logical ranges.

| Range         | Category       |
| ------------- | -------------- |
| 0x0000–0x00FF | General        |
| 0x0100–0x01FF | Connection     |
| 0x0200–0x02FF | Authentication |
| 0x0300–0x03FF | Session        |
| 0x0400–0x04FF | Tunnel         |
| 0x0500–0x05FF | Stream         |
| 0x0600–0x06FF | Protocol       |
| 0x0700–0x07FF | Internal       |
| 0xF000–0xFFFF | Reserved       |

Reserved ranges shall not be used in ATP v1.

---

# 6. Connection Errors

|   Code | Name               | Description                  |
| -----: | ------------------ | ---------------------------- |
| 0x0101 | CONNECTION_TIMEOUT | Connection timed out         |
| 0x0102 | CONNECTION_REFUSED | Connection refused           |
| 0x0103 | CONNECTION_RESET   | Connection reset             |
| 0x0104 | CONNECTION_CLOSED  | Connection closed            |
| 0x0105 | CONNECTION_LIMIT   | Maximum connections exceeded |

---

# 7. Authentication Errors

|   Code | Name                  | Description                  |
| -----: | --------------------- | ---------------------------- |
| 0x0201 | AUTHENTICATION_FAILED | Authentication failed        |
| 0x0202 | INVALID_CREDENTIALS   | Credentials invalid          |
| 0x0203 | TOKEN_EXPIRED         | Authentication token expired |
| 0x0204 | DEVICE_NOT_REGISTERED | Unknown device               |
| 0x0205 | AUTHORIZATION_DENIED  | Permission denied            |

---

# 8. Session Errors

|   Code | Name              | Description       |
| -----: | ----------------- | ----------------- |
| 0x0301 | SESSION_NOT_FOUND | Session missing   |
| 0x0302 | SESSION_EXPIRED   | Session expired   |
| 0x0303 | SESSION_INVALID   | Invalid session   |
| 0x0304 | SESSION_CONFLICT  | Duplicate session |
| 0x0305 | SESSION_CLOSED    | Session closed    |

---

# 9. Tunnel Errors

|   Code | Name                   | Description            |
| -----: | ---------------------- | ---------------------- |
| 0x0401 | TUNNEL_NOT_FOUND       | Tunnel missing         |
| 0x0402 | TUNNEL_CLOSED          | Tunnel closed          |
| 0x0403 | TUNNEL_BUSY            | Tunnel unavailable     |
| 0x0404 | TUNNEL_TIMEOUT         | Tunnel timeout         |
| 0x0405 | TUNNEL_RECOVERY_FAILED | Tunnel recovery failed |

---

# 10. Stream Errors

|   Code | Name             | Description              |
| -----: | ---------------- | ------------------------ |
| 0x0501 | STREAM_NOT_FOUND | Stream missing           |
| 0x0502 | STREAM_CLOSED    | Stream closed            |
| 0x0503 | STREAM_RESET     | Stream reset             |
| 0x0504 | STREAM_LIMIT     | Maximum streams exceeded |
| 0x0505 | STREAM_TIMEOUT   | Stream timeout           |

---

# 11. Protocol Errors

|   Code | Name             | Description                  |
| -----: | ---------------- | ---------------------------- |
| 0x0601 | INVALID_PACKET   | Packet malformed             |
| 0x0602 | INVALID_HEADER   | Header invalid               |
| 0x0603 | INVALID_VERSION  | Unsupported protocol version |
| 0x0604 | UNKNOWN_MESSAGE  | Unknown message type         |
| 0x0605 | INVALID_STATE    | Invalid state transition     |
| 0x0606 | INVALID_SEQUENCE | Sequence mismatch            |
| 0x0607 | INVALID_PAYLOAD  | Payload validation failed    |

---

# 12. Internal Errors

|   Code | Name                  | Description              |
| -----: | --------------------- | ------------------------ |
| 0x0701 | INTERNAL_ERROR        | Generic internal failure |
| 0x0702 | DATABASE_ERROR        | Storage unavailable      |
| 0x0703 | CONFIGURATION_ERROR   | Invalid configuration    |
| 0x0704 | RESOURCE_EXHAUSTED    | Resources unavailable    |
| 0x0705 | FEATURE_NOT_SUPPORTED | Feature unavailable      |

Internal errors should not expose sensitive implementation details.

---

# 13. Error Processing

When an error occurs:

```
Error Detected
    ↓
Assign Error Code
    ↓
Generate ERROR Message
    ↓
Transmit
    ↓
Log Event
    ↓
Recovery Decision
```

Recoverable errors should not immediately terminate the connection.

---

## Recovery Classification

| Error Type     | Action           |
| -------------- | ---------------- |
| Recoverable    | Retry            |
| Authentication | Disconnect       |
| Protocol       | Reject packet    |
| Fatal          | Close connection |

---

# 14. Security Considerations

ATP error handling shall:

- Avoid information leakage
- Validate error messages
- Log security-related failures
- Prevent replay abuse
- Never expose secrets

Authentication failures should avoid revealing whether a device exists.

---

# 15. Future Evolution

Future ATP versions may introduce:

- Vendor-specific errors
- Extension error ranges
- Localized descriptions
- Diagnostic metadata
- Recovery hints

Existing error codes shall remain stable.

---

# 16. Summary

ATP Error Codes define a standardized mechanism for communicating
failures between BeaconLink Agents and Relay Nodes.

By using structured numeric codes, symbolic names, and categorized
recovery behavior, ATP implementations remain predictable,
interoperable, and easy to diagnose.

Error handling is considered a first-class part of the protocol rather
than an implementation detail.

> **"A reliable protocol defines failures as carefully as successes."**

---

# Appendix A — Error Processing Flow

```
Receive Packet
      │
      ▼
Validate
      │
 ┌────┴─────┐
 │          │
 ▼          ▼
Valid     Invalid
 │          │
 ▼          ▼
Process   Generate Error
 │          │
 ▼          ▼
Complete  Send ERROR Message
```

---

# Appendix B — Error Category Overview

```
ATP Errors
│
├── General
├── Connection
├── Authentication
├── Session
├── Tunnel
├── Stream
├── Protocol
└── Internal
```

---

# Appendix C — Error Message Example

```
Packet Header
    ↓
Message Type = ERROR
    ↓
Error Code = 0x0203
    ↓
Error Name = TOKEN_EXPIRED
    ↓
Metadata = {
    RetryAfter: 30
}
```

---

# Appendix D — Error Severity Levels

| Severity | Description                | Typical Action   |
| -------- | -------------------------- | ---------------- |
| INFO     | Informational event        | Continue         |
| WARNING  | Recoverable issue          | Retry            |
| ERROR    | Operation failed           | Reject operation |
| FATAL    | Connection cannot continue | Disconnect       |

---

# Appendix E — Reserved Error Ranges

| Range         | Purpose                  |
| ------------- | ------------------------ |
| 0x0800–0x0FFF | Future ATP Core          |
| 0x1000–0xEFFF | Extension Specifications |
| 0xF000–0xFFFF | Vendor / Private Use     |
