# 013 - Relay Architecture

> **Document ID:** 013
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
> - 010 System Architecture
> - 011 Component Architecture
> - 012 Agent Architecture
> - 016 Network Architecture
> - 030 BeaconLink Tunnel Protocol
> - ADR-001 Relay First Networking

---

# Table of Contents

1. Overview
2. Purpose
3. Design Goals
4. Responsibilities
5. Internal Architecture
6. Core Modules
7. Request Lifecycle
8. Tunnel Management
9. Domain Routing
10. Security
11. Scalability
12. Failure Recovery
13. Monitoring
14. Future Evolution
15. Summary

---

# 1. Overview

The BeaconLink Relay is the public entry point of the BeaconLink platform.

It accepts incoming HTTPS requests, authenticates BeaconLink Agents,
maintains secure tunnels, and forwards traffic to the appropriate
user-owned device.

The Relay never executes user applications.

It exists solely to route traffic securely and efficiently.

---

# 2. Purpose

The Relay solves several networking challenges that commonly prevent
users from self-hosting.

These include:

- NAT traversal
- Dynamic IP addresses
- HTTPS termination
- Domain routing
- Tunnel management
- Device discovery

The Relay hides these complexities from users.

---

# 3. Design Goals

The Relay should be:

- Stateless whenever practical
- Horizontally scalable
- Highly available
- Secure
- Observable
- Efficient

The Relay should avoid storing persistent application state.

---

# 4. Responsibilities

The Relay is responsible for:

## Traffic Routing

- Receive HTTPS requests
- Identify destination
- Forward requests through ATP

---

## Tunnel Management

- Register tunnels
- Maintain tunnel state
- Detect disconnects
- Remove inactive tunnels

---

## Device Authentication

- Verify Agent identity
- Validate sessions
- Reject unauthorized devices

---

## Domain Resolution

- Resolve custom domains
- Associate domains with Device IDs
- Verify ownership

---

## Session Management

- Track active sessions
- Manage heartbeats
- Remove expired sessions

---

# 5. Internal Architecture

```
                 BeaconLink Relay

┌─────────────────────────────────────┐
│ Edge Listener                       │
│ TLS Manager                         │
│ Authentication Service              │
│ Tunnel Manager                      │
│ Routing Engine                      │
│ Session Manager                     │
│ Domain Registry                     │
│ Health Monitor                      │
│ Metrics Engine                      │
│ Logging Engine                      │
│ Configuration Service               │
└─────────────────────────────────────┘
```

Every module owns one responsibility.

---

# 6. Core Modules

## Edge Listener

Accepts:

- HTTPS
- HTTP/3
- ATP connections

Routes traffic internally.

---

## TLS Manager

Responsible for:

- TLS termination
- Certificate selection
- Secure negotiation

Supports modern TLS standards only.

---

## Authentication Service

Responsible for:

- Device authentication
- Session validation
- Token verification

Every Agent must authenticate before traffic is accepted.

---

## Tunnel Manager

Responsible for:

- Tunnel creation
- Tunnel recovery
- Tunnel health
- Tunnel multiplexing

The Relay maintains exactly one tunnel per connected Agent.

---

## Routing Engine

Maps:

```
Domain
    ↓
Device
    ↓
Tunnel
    ↓
Agent
```

The Routing Engine never knows how the application works.

It only forwards traffic.

---

## Session Manager

Tracks:

- Connected devices
- Active sessions
- Heartbeats
- Connection lifetime

---

## Domain Registry

Stores:

- Domain mappings
- Device associations
- Certificate status

The Registry stores metadata only.

---

## Health Monitor

Continuously verifies:

- Relay health
- Tunnel health
- Internal service status

---

## Metrics Engine

Collects:

- Requests
- Tunnel count
- Active Agents
- Latency
- Errors

Future versions may export metrics to external monitoring systems.

---

## Logging Engine

Produces:

- Access logs
- Security logs
- Error logs
- Audit logs

Logs should be structured and searchable.

---

# 7. Request Lifecycle

```
Visitor
    ↓
HTTPS Request
    ↓
Edge Listener
    ↓
TLS
    ↓
Routing Engine
    ↓
Tunnel Manager
    ↓
BeaconLink Agent
    ↓
Application
    ↓
Agent
    ↓
Tunnel
    ↓
Relay
    ↓
Visitor
```

The Relay never executes user code.

---

# 8. Tunnel Management

Each Agent establishes one persistent ATP tunnel.

The Relay:

- verifies authentication
- tracks heartbeats
- reconnects sessions
- cleans stale tunnels

Tunnel multiplexing allows multiple applications to share one connection.

---

# 9. Domain Routing

```
example.com
    ↓
Domain Registry
    ↓
Device ID
    ↓
Tunnel
    ↓
BeaconLink Agent
    ↓
Application
```

Domains are resolved without exposing the user's network topology.

---

# 10. Security

The Relay follows Zero Trust principles.

Requirements:

- Mutual authentication
- TLS 1.3+
- Signed sessions
- Replay protection
- Rate limiting
- Request validation

The Relay never stores Agent private keys.

---

# 11. Scalability

The Relay should support horizontal scaling.

Future architecture:

```
          DNS
           │
Load Balancer
   │      │      │
Relay  Relay  Relay
   │      │      │
 Shared Session Store
```

Additional Relay instances should increase capacity without requiring
architectural redesign.

---

# 12. Failure Recovery

Failures should be isolated.

Examples:

## Agent Disconnect

```
Heartbeat Lost
    ↓
Tunnel Closed
    ↓
Routing Disabled
    ↓
Wait for Reconnect
```

---

## Relay Restart

```
Restart
    ↓
Recover State
    ↓
Accept New Tunnels
```

---

## Routing Failure

```
Routing Error
    ↓
Log
    ↓
Return Friendly Error
```

---

# 13. Monitoring

The Relay should expose operational metrics.

Examples:

- Active tunnels
- Connected Agents
- Request latency
- Authentication failures
- Error rate
- CPU usage
- Memory usage

Future versions may support distributed tracing.

---

# 14. Future Evolution

Future improvements include:

- Multi-region relays
- Relay clustering
- Intelligent routing
- Edge relays
- Peer-to-peer optimization
- Geo-aware routing
- Automatic failover

These enhancements should preserve compatibility with the existing ATP.

---

# 15. Summary

The BeaconLink Relay serves as the networking backbone of the platform.

It securely connects visitors to applications running on user-owned
hardware while hiding networking complexity and preserving user
ownership.

Its architecture prioritizes security, scalability, modularity, and
reliability.

> "The Relay routes traffic—not applications."
