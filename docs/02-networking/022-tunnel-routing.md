# 022 - Tunnel Routing

> **Document ID:** 022
>
> **Version:** 1.0.0
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
> - 013 Relay Architecture
> - 016 Network Architecture
> - 020 Networking Overview
> - 021 Relay Network
> - 023 Domain Routing
> - 030 BeaconLink Tunnel Protocol
> - ADR-001 Relay First Networking

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Objectives
4. Routing Philosophy
5. Routing Model
6. Routing Pipeline
7. Tunnel Lifecycle
8. Tunnel Multiplexing
9. Session Routing
10. Health-Based Routing
11. Routing State Machine
12. Error Handling
13. Performance Requirements
14. Security Considerations
15. Future Evolution
16. Summary

---

# 1. Introduction

Tunnel Routing defines how BeaconLink forwards traffic from the Relay Network
to user applications through persistent encrypted tunnels.

Rather than exposing user devices directly to the Internet, BeaconLink
establishes outbound tunnels between the BeaconLink Agent and the Relay
Network.

Incoming traffic is forwarded through these tunnels based on routing
metadata rather than physical network addresses.

---

# 2. Purpose

Tunnel Routing provides:

- Device discovery
- Tunnel selection
- Session forwarding
- Request multiplexing
- Traffic isolation
- Automatic recovery

The routing system should remain transparent to end users.

---

# 3. Design Objectives

The routing system should be:

- Deterministic
- Fast
- Reliable
- Secure
- Stateless whenever practical
- Horizontally scalable

Routing decisions should rely only on routing metadata and tunnel state.

---

# 4. Routing Philosophy

BeaconLink routes requests using logical identities rather than IP addresses.

The Relay should never need to know:

- Local IP addresses
- Internal network topology
- Router configuration

Instead, routing is based on:

```
Domain
    ↓
Site
    ↓
Device
    ↓
Tunnel
    ↓
Application
```

This abstraction keeps the routing layer independent of physical
networks.

---

# 5. Routing Model

Each hosted application is associated with a unique routing path.

```
Custom Domain
    ↓
Site Identifier
    ↓
Device Identifier
    ↓
Tunnel Identifier
    ↓
Runtime
    ↓
Application
```

Each level has a single responsibility.

| Layer       | Responsibility               |
| ----------- | ---------------------------- |
| Domain      | Public entry point           |
| Site        | Logical application          |
| Device      | Hosting machine              |
| Tunnel      | Secure communication channel |
| Runtime     | Application execution        |
| Application | User workload                |

---

# 6. Routing Pipeline

Every incoming request follows the same routing pipeline.

```
Visitor
    ↓
DNS Resolution
    ↓
Relay
    ↓
TLS Validation
    ↓
Domain Lookup
    ↓
Site Lookup
    ↓
Device Lookup
    ↓
Tunnel Lookup
    ↓
Tunnel Validation
    ↓
ATP Forward
    ↓
Agent
    ↓
Runtime
    ↓
Application
```

Responses travel through the same tunnel in reverse order.

---

# 7. Tunnel Lifecycle

Each tunnel progresses through a defined lifecycle.

```
Disconnected
    ↓
Connecting
    ↓
Authenticating
    ↓
Established
    ↓
Active
    ↓
Idle
    ↓
Closing
    ↓
Disconnected
```

Only **Established**, **Active**, and **Idle** tunnels are eligible for
traffic.

---

# 8. Tunnel Multiplexing

Each BeaconLink Agent maintains one persistent ATP tunnel.

Multiple applications may share this tunnel.

```
                ATP Tunnel
                     │
      ┌──────────────┼──────────────┐
      ▼              ▼              ▼
   Site A         Site B         Site C
      ▼              ▼              ▼
  Runtime A      Runtime B      Runtime C
```

Advantages:

- Lower connection overhead
- Faster request handling
- Simplified firewall traversal
- Reduced resource usage

---

# 9. Session Routing

Every request is associated with a session.

Session metadata includes:

- Session ID
- Device ID
- Tunnel ID
- Site ID
- Authentication state
- Creation timestamp
- Last activity

Sessions are ephemeral and automatically expire after inactivity.

---

# 10. Health-Based Routing

Traffic is routed only to healthy applications.

Health verification includes:

- Active tunnel
- Connected Agent
- Healthy runtime
- Responsive application

If any stage fails, routing is aborted.

```
Tunnel Healthy?
    ↓
YES → Continue
    ↓
NO → Return Service Unavailable
```

Future versions may support automatic failover.

---

# 11. Routing State Machine

```
Request Received
    ↓
Validate
    ↓
Lookup Domain
    ↓
Lookup Site
    ↓
Lookup Device
    ↓
Lookup Tunnel
    ↓
Tunnel Available?
    ↓
YES
    ↓
Forward Request
    ↓
Receive Response
    ↓
Return Response
    ↓
Complete
```

Failure path:

```
Tunnel Missing
    ↓
Retry Lookup
    ↓
Fail
    ↓
Return Error
```

---

# 12. Error Handling

Typical routing failures include:

## Unknown Domain

Return:

```
404 Not Found
```

---

## Invalid Certificate

Return:

```
TLS Handshake Failure
```

---

## Offline Agent

Return:

```
503 Service Unavailable
```

---

## Missing Tunnel

Attempt:

```
Retry
    ↓
Reconnect
    ↓
Fail
```

Errors should be logged with sufficient context for diagnostics.

---

# 13. Performance Requirements

Tunnel Routing should achieve:

## Low Lookup Latency

Routing decisions should be completed quickly.

---

## Efficient Multiplexing

Support multiple concurrent requests per tunnel.

---

## High Throughput

Process large numbers of concurrent requests efficiently.

---

## Low Resource Usage

Maintain minimal CPU and memory overhead.

---

# 14. Security Considerations

Tunnel Routing must provide:

- Mutual authentication
- Tunnel validation
- Session verification
- Replay protection
- Request integrity
- Least privilege

Routing metadata should never expose sensitive internal network
information.

---

# 15. Future Evolution

Future routing enhancements may include:

- Multi-tunnel Agents
- Intelligent routing
- Geo-aware routing
- Load-aware routing
- Automatic tunnel migration
- Edge Relay optimization
- Peer-to-peer routing
- Adaptive traffic steering

Future capabilities should remain compatible with existing routing
semantics.

---

# 16. Summary

Tunnel Routing is responsible for mapping incoming Internet traffic to
applications hosted on user-owned devices.

By routing requests through logical identities rather than physical
network addresses, BeaconLink simplifies self-hosting while maintaining
security, scalability, and flexibility.

The routing architecture is designed to evolve incrementally without
requiring changes to application deployments.

> **"Applications are reached through identity, not location."**

---

# Appendix A — Routing Metadata

| Field      | Description                   |
| ---------- | ----------------------------- |
| Domain ID  | Public hostname               |
| Site ID    | Hosted application identifier |
| Device ID  | Registered BeaconLink Agent        |
| Tunnel ID  | Active ATP tunnel             |
| Runtime ID | Runtime instance              |
| Session ID | Current request session       |

---

# Appendix B — Routing Decision Flow

```
Incoming Request
        │
        ▼
Validate TLS
        │
        ▼
Resolve Domain
        │
        ▼
Find Site
        │
        ▼
Locate Device
        │
        ▼
Verify Tunnel
        │
        ▼
Check Health
        │
        ▼
Forward via ATP
        │
        ▼
Receive Response
        │
        ▼
Return to Client
```

---

# Appendix C — Tunnel Multiplexing Diagram

```
                 BeaconLink Relay
                      │
          Persistent ATP Tunnel
                      │
            ┌─────────┬─────────┐
            ▼         ▼         ▼
          Site A     Site B     Site C
            │         │         │
        Runtime A  Runtime B  Runtime C
            │         │         │
          App A     App B     App C
```

---

# Appendix D — Sequence Diagram

```
Visitor        Relay        Agent       Runtime      Application
   |             |             |             |              |
   | HTTPS Req   |             |             |              |
   |------------>|             |             |              |
   |             | Lookup Site |             |              |
   |             |------------>|             |              |
   |             | ATP Forward |             |              |
   |             |------------>|             |              |
   |             |             | Execute     |              |
   |             |             |-----------> |              |
   |             |             |             | Process Req  |
   |             |             |             |------------->|
   |             |             |             |<-------------|
   |             |<------------| Response    |              |
   |<------------|             |             |              |
```
