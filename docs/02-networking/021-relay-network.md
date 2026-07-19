# 021 - Relay Network

> **Document ID:** 021
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
> - 018 Scalability Architecture
> - 020 Networking Overview
> - 022 Tunnel Routing
> - 024 Load Balancing
> - 030 BeaconLink Tunnel Protocol
> - ADR-012 Logical Distributed Relay Architecture
> - ADR-015 Triple Network Architecture

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Objectives
4. Relay Network Philosophy
5. Relay Network Model
6. Relay Components
7. Relay Topologies
8. Relay Lifecycle
9. Request Processing
10. Session Management
11. State Synchronization
12. Security Model
13. Performance Requirements
14. High Availability
15. Failure Scenarios
16. Monitoring
17. Future Evolution
18. Summary

---

# 1. Introduction

The Relay Network forms the public networking layer of BeaconLink.

Every request originating from the Internet enters the BeaconLink platform
through one or more Relay Nodes before being securely forwarded to the
appropriate BeaconLink Agent.

The Relay Network provides secure routing, session management, tunnel
coordination, and traffic forwarding while ensuring that user-owned
devices remain protected from direct Internet exposure.

---

# 2. Purpose

The Relay Network exists to provide:

- Public network entry points
- Secure traffic routing
- Tunnel management
- Session coordination
- Device discovery
- Horizontal scalability
- High availability

The Relay Network never executes user applications or stores user
application data.

---

# 3. Design Objectives

The Relay Network should be:

- Highly available
- Horizontally scalable
- Stateless where practical
- Fault tolerant
- Secure
- Observable
- Region-aware

Growth should require additional Relay Nodes rather than architectural
changes.

---

# 4. Relay Network Philosophy

BeaconLink adopts a **Relay-First** model.

All inbound traffic follows the same logical path:

```
Internet
    ↓
BeaconLink Relay Network
    ↓
Secure Tunnel
    ↓
BeaconLink Agent
    ↓
User Application
```

This approach eliminates the need for:

- Port forwarding
- Public IP addresses
- Manual firewall configuration
- Reverse proxy configuration

Future peer-to-peer communication may supplement—but not replace—the
Relay Network.

---

# 5. Relay Network Model

The Relay Network consists of one or more Relay Nodes operating as a
logical platform.

```
                    Internet
                         │
                         ▼
                 Global DNS Service
                         │
                         ▼
                  Load Balancer
                         │
        ┌────────────────┼────────────────┐
        ▼                ▼                ▼
   Relay Node A     Relay Node B     Relay Node C
        │                │                │
        └──────────── Shared Services ────┘
                         │
             ┌───────────┼───────────┐
             ▼           ▼           ▼
      Session Store  Domain Registry  Metrics
```

Users interact with the Relay Network as a single logical service,
regardless of the number of Relay Nodes.

---

# 6. Relay Components

The Relay Network consists of the following logical components.

## Edge Gateway

Receives all public network traffic.

Responsibilities:

- Accept HTTPS requests
- Accept ATP connections
- Forward requests internally

---

## Relay Node

Responsible for:

- Authentication
- Tunnel management
- Request routing
- Session handling

Relay Nodes are designed to be interchangeable.

---

## Session Store

Maintains:

- Active sessions
- Tunnel metadata
- Heartbeat timestamps

Persistent application data is never stored here.

---

## Domain Registry

Maps:

```
Domain
    ↓
Site
    ↓
Device
    ↓
Tunnel
```

The registry stores routing metadata only.

---

## Metrics Service

Collects:

- Request counts
- Tunnel counts
- Active sessions
- Latency
- Error rates

---

# 7. Relay Topologies

## Single Relay

Suitable for development and small deployments.

```
Internet
    ↓
Relay
    ↓
Agent
```

---

## Clustered Relay

Recommended for production.

```
Internet
    ↓
Load Balancer
    ↓
Relay Cluster
    ↓
Agents
```

---

## Multi-Region Relay

Future deployment model.

```
Asia Relay

Europe Relay

North America Relay

↓

Shared Services
```

Users should automatically connect to the optimal region.

---

# 8. Relay Lifecycle

Each Relay Node follows the same lifecycle.

```
Start
    ↓
Load Configuration
    ↓
Initialize Services
    ↓
Register Health
    ↓
Accept Connections
    ↓
Authenticate Agents
    ↓
Route Traffic
    ↓
Monitor
    ↓
Shutdown
```

Graceful shutdown should allow existing requests to complete.

---

# 9. Request Processing

Every incoming request follows the same pipeline.

```
Visitor
    ↓
DNS Resolution
    ↓
Load Balancer
    ↓
Relay Node
    ↓
TLS Validation
    ↓
Domain Lookup
    ↓
Tunnel Lookup
    ↓
ATP Forwarding
    ↓
BeaconLink Agent
    ↓
Runtime
    ↓
Application
    ↓
Response
```

Relay Nodes never inspect or modify application logic.

---

# 10. Session Management

Each connected Agent maintains an active session.

Session state includes:

- Device ID
- Tunnel ID
- Authentication status
- Last heartbeat
- Relay assignment

Sessions remain valid only while heartbeats continue.

---

# 11. State Synchronization

Relay Nodes share only operational metadata.

Examples include:

- Active sessions
- Domain mappings
- Certificate status
- Health information

Application state remains exclusively on user devices.

Relay Nodes should avoid unnecessary shared state.

---

# 12. Security Model

The Relay Network follows Zero Trust principles.

Security controls include:

- Mutual authentication
- TLS 1.3
- ATP encryption
- Replay protection
- Rate limiting
- Session validation
- Certificate verification

Private keys are never stored by Relay Nodes.

---

# 13. Performance Requirements

The Relay Network should achieve:

## Low Latency

Minimize request forwarding overhead.

---

## High Throughput

Support large numbers of concurrent requests.

---

## Efficient Resource Usage

Maintain low CPU and memory overhead during normal operation.

---

## Fast Recovery

Restore sessions quickly after temporary failures.

---

# 14. High Availability

Relay availability should improve through redundancy.

```
Internet
    ↓
Global Load Balancer
    ↓
Relay Cluster
    ↓
Shared Session Store
```

Failure of a single Relay Node should not interrupt platform operation.

---

# 15. Failure Scenarios

## Relay Node Failure

```
Node Offline
    ↓
Health Check Failed
    ↓
Load Balancer Removes Node
    ↓
Traffic Redirected
    ↓
New Connections Accepted
```

---

## Agent Disconnect

```
Heartbeat Missing
    ↓
Session Expired
    ↓
Tunnel Closed
    ↓
Await Reconnection
```

---

## Shared Service Failure

```
Service Unavailable
    ↓
Retry
    ↓
Fallback
    ↓
Alert Operator
```

Failures should remain isolated whenever possible.

---

# 16. Monitoring

Relay Network observability should include:

Infrastructure Metrics

- CPU utilization
- Memory usage
- Network bandwidth

Operational Metrics

- Active tunnels
- Connected Agents
- Requests per second
- Session count
- Authentication failures
- Error rate
- Average latency

Monitoring should support proactive capacity planning.

---

# 17. Future Evolution

Future enhancements may include:

- Relay federation
- Intelligent regional routing
- Anycast networking
- Edge Relay Nodes
- Dynamic traffic optimization
- Automatic session migration
- Self-healing clusters

These improvements should preserve compatibility with the existing Relay
Network architecture.

---

# 18. Summary

The Relay Network provides the public networking foundation of BeaconLink.

By combining stateless Relay Nodes, shared operational metadata,
persistent tunnels, and secure request routing, BeaconLink delivers reliable
Internet connectivity while keeping applications on user-owned
infrastructure.

The Relay Network is designed to grow from a single Relay Node to a
globally distributed platform without requiring fundamental
architectural redesign.

> **"The Relay Network connects the Internet to BeaconLink—not the Internet to
> your computer."**

---

# Appendix A — Relay Node Responsibilities

| Responsibility          | Relay Node | Agent |
| ----------------------- | ---------: | ----: |
| Accept Internet Traffic |         ✅ |    ❌ |
| Authenticate Devices    |         ✅ |    ❌ |
| Maintain Tunnel         |         ✅ |    ✅ |
| Execute Applications    |         ❌ |    ✅ |
| Store Application Data  |         ❌ |    ✅ |
| Route Requests          |         ✅ |    ❌ |
| Monitor Local Processes |         ❌ |    ✅ |

---

# Appendix B — C4 Context Diagram (ASCII)

```
                    +----------------------+
                    |      Internet        |
                    +----------+-----------+
                               |
                               v
                    +----------------------+
                    |   BeaconLink Relay Net    |
                    +----------+-----------+
                               |
                      BeaconLink Tunnel Protocol
                               |
                               v
                    +----------------------+
                    |     BeaconLink Agent      |
                    +----------+-----------+
                               |
                               v
                    +----------------------+
                    |  User Applications   |
                    +----------------------+
```

---

# Appendix C — Relay Network Sequence Diagram

```
Visitor          Relay         Agent        Runtime      Application
   |               |             |             |              |
   | HTTPS Request |             |             |              |
   |-------------->|             |             |              |
   |               | ATP Forward |             |              |
   |               |-----------> |             |              |
   |               |             | Execute     |              |
   |               |             |-----------> |              |
   |               |             |             | Run Request  |
   |               |             |             |------------->|
   |               |             |             |<-------------|
   |               |<----------- | Response    |              |
   |<--------------|             |             |              |
```
