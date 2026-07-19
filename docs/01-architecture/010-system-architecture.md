# 010 - System Architecture

> **Document ID:** 010
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
> - 000 Engineering Principles
> - 001 Project Vision
> - 007 Non Functional Requirements
> - ADR-001 Relay First Networking

---

# Table of Contents

1. Overview
2. Architecture Goals
3. Design Philosophy
4. System Overview
5. High-Level Components
6. Component Responsibilities
7. Data Flow
8. Communication Flow
9. Trust Boundaries
10. Deployment Model
11. Design Decisions
12. Scalability
13. Failure Handling
14. Future Evolution
15. Summary

---

# 1. Overview

BeaconLink is a distributed self-hosting platform.

Unlike traditional hosting providers, BeaconLink does not host customer
applications.

Instead, BeaconLink provides the infrastructure necessary to securely expose
applications running on hardware owned by the user.

BeaconLink consists of several independent components working together to
provide a cloud-like deployment experience while preserving user
ownership.

---

# 2. Architecture Goals

The architecture is designed to satisfy the following objectives.

- Simplicity
- Security
- Reliability
- Extensibility
- Modularity
- Cross-platform compatibility

Every architectural decision should improve at least one of these
qualities without significantly weakening another.

---

# 3. Design Philosophy

BeaconLink follows several architectural principles.

## User-Owned Infrastructure

Applications always execute on hardware owned by the user.

BeaconLink never executes user applications on BeaconLink-managed servers.

---

## Stateless Cloud

Cloud services should only coordinate communication.

Whenever possible they should avoid storing application-specific state.

---

## Modular Components

Every major component should be independently replaceable.

Changes in one component should require minimal changes elsewhere.

---

## Interface-Based Communication

Components communicate only through documented protocols.

No component should depend upon another component's internal
implementation.

---

# 4. System Overview

BeaconLink consists of four primary systems.

```text
                      Internet
                         │
                         ▼
                   Custom Domain
                         │
                         ▼
                  BeaconLink Edge Relay
                         │
                BeaconLink Tunnel Protocol
                         │
                         ▼
                    BeaconLink Agent
                         │
                         ▼
                  User Application
```

---

# 5. High-Level Components

## BeaconLink Agent

Runs on the user's machine.

Responsibilities:

- Runtime detection
- Application lifecycle
- Tunnel maintenance
- Monitoring
- Logging
- Health checks

---

## BeaconLink Relay

Runs within BeaconLink infrastructure.

Responsibilities:

- Device authentication
- Tunnel management
- Request routing
- TLS termination
- Domain mapping

The Relay never executes user applications.

---

## BeaconLink Console

Web application used by users.

Responsibilities:

- Device management
- Domain management
- Deployment management
- Monitoring
- Configuration

---

## BeaconLink API

Provides a stable interface between:

- Console
- Agent
- Relay

All communication occurs through documented APIs.

---

# 6. Component Responsibilities

| Component      | Responsibility          |
| -------------- | ----------------------- |
| Agent          | Execute applications    |
| Relay          | Route traffic           |
| Console        | User interface          |
| API            | Service communication   |
| DNS            | Domain verification     |
| Update Service | Secure software updates |

Each component has a single primary responsibility.

---

# 7. Data Flow

Application traffic flows as follows.

```text
Visitor
    ↓
DNS
    ↓
Relay
    ↓
Encrypted Tunnel
    ↓
BeaconLink Agent
    ↓
Application
    ↓
BeaconLink Agent
    ↓
Tunnel
    ↓
Relay
    ↓
Visitor
```

Management traffic follows a separate path.

```text
 Console
    ↓
   API
    ↓
  Relay
    ↓
  Agent
```

Separating management traffic from application traffic simplifies
security and troubleshooting.

---

# 8. Communication Flow

BeaconLink uses persistent encrypted connections.

Communication categories include.

- Authentication
- Heartbeats
- Tunnel traffic
- Configuration
- Logging
- Monitoring
- Software updates

Each category uses the BeaconLink Tunnel Protocol.

---

# 9. Trust Boundaries

BeaconLink defines clear trust boundaries.

```text
User
──────────────

BeaconLink Console

──────────────

BeaconLink Infrastructure

──────────────

BeaconLink Agent

──────────────

User Application
```

Every boundary requires authentication.

No implicit trust exists between components.

---

# 10. Deployment Model

BeaconLink supports deployments on:

- Desktop
- Workstation
- Mini PC
- Home Server
- Rack Server
- Virtual Machine

Future versions may support container orchestration environments.

---

# 11. Design Decisions

The following architectural decisions influence the platform.

| Decision                  | ADR     |
| ------------------------- | ------- |
| Relay First Networking    | ADR-001 |
| Persistent Connections    | ADR-002 |
| Public Key Authentication | ADR-003 |
| Runtime Abstraction       | ADR-004 |
| Zero Trust                | ADR-007 |

Future decisions should be documented through Architecture Decision
Records.

---

# 12. Scalability

The architecture supports horizontal growth.

Examples include.

- Multiple relay regions
- Multiple agents
- Multiple domains
- Multiple applications
- Multiple users

Scaling should occur through component replication rather than
architectural redesign.

---

# 13. Failure Handling

BeaconLink assumes failures are inevitable.

The platform should detect and recover from.

- Agent failures
- Application failures
- Network interruptions
- Relay outages
- Authentication failures

Recovery should occur automatically whenever possible.

---

# 14. Future Evolution

The architecture intentionally leaves room for.

- Peer-to-peer networking
- Edge relays
- Plugin system
- Additional runtimes
- Enterprise deployments
- High availability
- Advanced monitoring

These capabilities should integrate without requiring architectural
redesign.

---

# 15. Summary

BeaconLink is built around a modular distributed architecture that separates
user applications from BeaconLink infrastructure.

Applications remain under user ownership while BeaconLink provides secure
networking, deployment automation, and management capabilities.

This architecture provides the foundation upon which every future BeaconLink
component will be built.

> "BeaconLink simplifies infrastructure without taking ownership away from the
> user."
