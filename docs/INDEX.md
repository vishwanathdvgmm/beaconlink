# BeaconLink Documentation Index

Welcome to the BeaconLink documentation.

This documentation is organized into a series of logical sections that progress from high-level concepts to detailed architectural specifications, implementation guidance, architectural decisions, and future platform proposals.

For readers new to BeaconLink, it is recommended to follow the documentation in the order presented below.

---

# Documentation Overview

```
BeaconLink Documentation

├── README
├── Manifesto
├── Glossary
│
├── Foundation
├── Architecture
├── Networking
├── Protocol
├── Security
├── Agent
├── Relay
├── Console
├── Deployment
├── API
├── Data
├── Quality
├── Development
│
├── Architecture Decision Records (ADR)
├── Request for Comments (RFC)
│
└── Diagrams
```

---

# Getting Started

## README

Project overview, architecture summary, repository structure, and navigation.

**File**

- `README.md`

---

## BeaconLink Manifesto

The long-term vision, philosophy, architectural values, and design principles behind BeaconLink.

**File**

- `BeaconLink_MANIFESTO.md`

---

## Glossary

Definitions of terminology used throughout the documentation.

**File**

- `GLOSSARY.md`

---

# Core Documentation

## 00 Foundation

Introduces the fundamental concepts that guide every architectural decision.

Topics include:

- Vision
- Goals
- Principles
- Design philosophy
- System assumptions
- Constraints

Directory

```
00-foundation/
```

---

## 01 Architecture

Provides the overall platform architecture.

Topics include:

- System overview
- Component architecture
- Control plane
- Data plane
- Layered architecture
- Design principles

Directory

```
01-architecture/
```

---

## 02 Networking

Describes the networking model used throughout BeaconLink.

Topics include:

- Relay-first networking
- Multi-site routing
- Network topology
- Connectivity
- High availability
- Federation

Directory

```
02-networking/
```

---

## 03 Protocol

Documents communication protocols between BeaconLink components.

Topics include:

- Agent protocol
- Relay protocol
- Message format
- Authentication
- Session lifecycle
- Version negotiation

Directory

```
03-protocol/
```

---

## 04 Security

Describes the BeaconLink security architecture.

Topics include:

- Zero Trust
- Authentication
- Authorization
- Encryption
- Certificates
- Audit logging
- Secrets
- Identity

Directory

```
04-security/
```

---

## 05 Agent

Documents the BeaconLink Agent architecture.

Topics include:

- Agent lifecycle
- Runtime management
- Health monitoring
- Configuration
- Inventory
- Updates

Directory

```
05-agent/
```

---

## 06 Relay

Documents the Relay architecture.

Topics include:

- Persistent connections
- Routing
- Session management
- High availability
- Federation
- Scalability

Directory

```
06-relay/
```

---

## 07 Console

Documents the BeaconLink web console.

Topics include:

- User interface
- Authentication
- Administration
- Dashboards
- Operations
- User experience

Directory

```
07-console/
```

---

## 08 Deployment

Documents deployment architecture.

Topics include:

- Desired state
- Reconciliation
- Runtime abstraction
- Immutable deployments
- Rollouts
- Deployment lifecycle

Directory

```
08-deployment/
```

---

## 09 API

Documents public platform APIs.

Topics include:

- REST API
- WebSocket API
- Authentication
- Versioning
- Error handling
- Client interactions

Directory

```
09-api/
```

---

## 10 Data

Documents data architecture.

Topics include:

- Storage
- Persistence
- Data models
- Event storage
- Audit records
- Metadata

Directory

```
10-data/
```

---

## 11 Quality

Documents quality engineering.

Topics include:

- Testing
- Reliability
- Performance
- Observability
- Monitoring
- Validation

Directory

```
11-quality/
```

---

## 12 Development

Documents engineering practices.

Topics include:

- Coding standards
- Repository organization
- Build system
- CI/CD
- Release process
- Contribution guidelines

Directory

```
12-development/
```

---

# Architecture Decision Records

Architecture Decision Records (ADRs) capture significant architectural decisions, including their context, rationale, alternatives, and consequences.

Current ADRs include:

| ADR     | Topic                                  |
| ------- | -------------------------------------- |
| ADR-001 | Relay-First Networking                 |
| ADR-002 | Persistent Connections                 |
| ADR-003 | Public Key Authentication              |
| ADR-004 | Runtime Abstraction                    |
| ADR-005 | Multi-Site Routing                     |
| ADR-006 | Site Manifest                          |
| ADR-007 | Zero Trust                             |
| ADR-008 | Container Strategy                     |
| ADR-009 | Protocol Versioning                    |
| ADR-010 | Immutable Update System                |
| ADR-011 | Modular Monolith Agent                 |
| ADR-012 | Logical Distributed Relay Architecture |
| ADR-013 | Monorepo Strategy                      |
| ADR-014 | Polyglot Persistence                   |
| ADR-015 | Event-Driven Internal Architecture     |
| ADR-016 | Declarative Desired State              |
| ADR-017 | API Gateway Pattern                    |
| ADR-018 | Observability Strategy                 |
| ADR-019 | Trunk-Based Development                |
| ADR-020 | Immutable Deployments                  |

Directory

```
ADR/
```

---

# Request for Comments

RFCs describe future architectural proposals and platform evolution.

Current RFCs include:

| RFC      | Proposal                     |
| -------- | ---------------------------- |
| RFC-0001 | Peer-to-Peer Networking      |
| RFC-0002 | Global Relay Network         |
| RFC-0003 | Plugin System                |
| RFC-0004 | Mobile Agent                 |
| RFC-0005 | Edge Cache                   |
| RFC-0006 | High Availability Relays     |
| RFC-0007 | Multi-Tenancy                |
| RFC-0008 | Service Mesh Integration     |
| RFC-0009 | Policy Engine                |
| RFC-0010 | Secret Management            |
| RFC-0011 | GitOps Integration           |
| RFC-0012 | Kubernetes Operator          |
| RFC-0013 | Autonomous Remediation       |
| RFC-0014 | Workflow Engine              |
| RFC-0015 | Event Streaming Platform     |
| RFC-0016 | WebAssembly Runtime          |
| RFC-0017 | AI-Assisted Operations       |
| RFC-0018 | Distributed Scheduler        |
| RFC-0019 | Cross-Site Disaster Recovery |
| RFC-0020 | Marketplace                  |

Directory

```
RFC/
```

---

# Architecture Diagrams

The `diagrams/` directory contains visual representations of the BeaconLink architecture.

```
diagrams/

├── architecture/
├── database/
├── deployment/
├── networking/
├── security/
├── sequence/
├── state/
└── ui/
```

These diagrams complement the written documentation and illustrate key architectural concepts, workflows, state machines, deployment models, communication protocols, and user interface layouts.

---

# Recommended Reading Order

Readers new to BeaconLink should follow this sequence:

1. README
2. BeaconLink Manifesto
3. Glossary
4. Foundation
5. Architecture
6. Networking
7. Protocol
8. Security
9. Agent
10. Relay
11. Console
12. Deployment
13. API
14. Data
15. Quality
16. Development
17. Architecture Decision Records
18. Request for Comments
19. Architecture Diagrams

---

# Documentation Conventions

Throughout this documentation:

- **ADR** refers to an Architecture Decision Record documenting accepted architectural decisions.
- **RFC** refers to a Request for Comments describing proposed future capabilities.
- **Site Manifest** defines the desired state of managed infrastructure.
- **Relay** provides secure outbound communication between agents and the control plane.
- **Agent** executes workloads and manages infrastructure on target systems.

Additional terminology is defined in `GLOSSARY.md`.

---

# Document Status

This documentation represents the current architectural vision for BeaconLink.

As the platform evolves, documents will continue to be refined while preserving the architectural principles established in the ADRs and long-term direction described by the RFCs.
