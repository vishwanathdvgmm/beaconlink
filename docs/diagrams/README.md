# BeaconLink Architecture Diagrams

This directory contains the complete architecture documentation for the **BeaconLink Platform** using **Mermaid** diagrams.

The documentation is organized into logical domains covering system architecture, deployment workflows, networking, security, data model, runtime behavior, operational sequences, state machines, and user interface design.

All diagrams are designed to be rendered by Mermaid-compatible documentation platforms such as GitHub, GitLab, MkDocs, Docusaurus, Obsidian, and Mermaid Live Editor.

---

# Directory Structure

```
diagrams/
├── architecture/
├── database/
├── deployment/
├── networking/
├── security/
├── sequence/
├── state/
├── ui/
└── README.md
```

---

# Documentation Overview

## architecture/

High-level platform architecture.

| File                      | Description                              |
| ------------------------- | ---------------------------------------- |
| 01-system-overview        | Overall BeaconLink platform architecture |
| 02-control-plane          | BeaconLink Control Plane components      |
| 03-agent-architecture     | BeaconLink Agent internals               |
| 04-runtime-manager        | Runtime abstraction architecture         |
| 05-plugin-framework       | Plugin Framework architecture            |
| 06-event-driven           | Event-driven architecture                |
| 07-control-flow           | Control plane execution flow             |
| 08-data-flow              | Platform data flow                       |
| 09-high-availability      | High availability architecture           |
| 10-reference-architecture | Complete enterprise deployment           |

---

## database/

Platform data model.

| File                   | Description                 |
| ---------------------- | --------------------------- |
| 01-erd                 | Entity relationship diagram |
| 02-agent-schema        | Agent data model            |
| 03-deployment-schema   | Deployment model            |
| 04-workflow-schema     | Workflow model              |
| 05-policy-schema       | Policy model                |
| 06-event-schema        | Event model                 |
| 07-audit-schema        | Audit model                 |
| 08-complete-data-model | Complete platform schema    |

---

## deployment/

Deployment lifecycle.

| File                    | Description                     |
| ----------------------- | ------------------------------- |
| 01-deployment-lifecycle | End-to-end deployment lifecycle |
| 02-planning             | Deployment planning             |
| 03-scheduling           | Capability-based scheduling     |
| 04-runtime-deployment   | Runtime deployment process      |
| 05-reconciliation       | Continuous reconciliation       |
| 06-rollback             | Rollback workflow               |
| 07-upgrade              | Platform upgrade flow           |
| 08-blue-green           | Blue/Green deployment           |
| 09-canary               | Canary deployment               |

---

## networking/

Relay-first networking architecture.

| File                          | Description                 |
| ----------------------------- | --------------------------- |
| 01-network-overview           | Network architecture        |
| 02-relay-network              | Relay Network topology      |
| 03-agent-connectivity         | Agent connectivity          |
| 04-control-plane-connectivity | Control Plane communication |
| 05-mtls                       | Mutual TLS architecture     |
| 06-multi-region               | Multi-region networking     |
| 07-failover                   | Network failover            |
| 08-load-balancing             | Traffic distribution        |
| 09-network-security           | Secure networking model     |

---

## security/

Security architecture.

| File                      | Description                    |
| ------------------------- | ------------------------------ |
| 01-zero-trust             | Zero Trust architecture        |
| 02-authentication         | Authentication flow            |
| 03-authorization          | Authorization model            |
| 04-policy-engine          | Policy Engine architecture     |
| 05-secret-management      | Secret lifecycle               |
| 06-certificate-management | Certificate lifecycle          |
| 07-audit-security         | Security auditing              |
| 08-security-overview      | Complete security architecture |

---

## sequence/

Runtime interaction sequences.

| File                  | Description               |
| --------------------- | ------------------------- |
| 01-agent-bootstrap    | Agent bootstrap           |
| 02-agent-registration | Agent registration        |
| 03-agent-heartbeat    | Heartbeat synchronization |
| 04-deployment-request | Deployment request        |
| 05-runtime-deployment | Runtime deployment        |
| 06-secret-fetch       | Secret retrieval          |
| 07-policy-evaluation  | Policy evaluation         |
| 08-plugin-loading     | Plugin loading            |
| 09-workflow-execution | Workflow execution        |
| 10-failover-sequence  | Platform failover         |

---

## state/

Platform state machines.

| File                           | Description                  |
| ------------------------------ | ---------------------------- |
| 01-agent-state-machine         | BeaconLink Agent lifecycle   |
| 02-deployment-state-machine    | Deployment lifecycle         |
| 03-relay-state-machine         | Relay lifecycle              |
| 04-workflow-state-machine      | Workflow execution lifecycle |
| 05-job-state-machine           | Job execution lifecycle      |
| 06-session-state-machine       | Session lifecycle            |
| 07-control-plane-state-machine | Control Plane lifecycle      |

---

## ui/

User interface documentation.

| File                 | Description        |
| -------------------- | ------------------ |
| 01-dashboard-layout  | Dashboard layout   |
| 02-site-management   | Site management    |
| 03-agent-details     | Agent details      |
| 04-deployment-view   | Deployment details |
| 05-relay-management  | Relay management   |
| 06-policy-management | Policy management  |
| 07-workflow-view     | Workflow execution |
| 08-marketplace       | Marketplace        |
| 09-settings          | Platform settings  |

---

# Architectural Principles

BeaconLink is built around several core architectural principles:

- Runtime Independent execution
- BeaconLink Control Plane
- BeaconLink Agent
- Relay Network
- Runtime Manager
- Runtime Adapter abstraction
- Plugin Framework
- Workflow Engine orchestration
- Deployment Planner
- Distributed Scheduler
- Capability-Based Scheduling
- Desired State management
- Continuous Reconciliation
- Immutable Revisions
- Zero Trust security
- Event Streaming Platform
- Inventory Service
- Policy Engine governance

---

# Mermaid Diagram Types

The repository contains multiple Mermaid diagram types including:

- Flowcharts
- Sequence diagrams
- State diagrams
- Entity Relationship Diagrams
- Component diagrams
- Deployment diagrams
- Network topology diagrams
- UI layout diagrams

---

# Documentation Conventions

The diagrams follow consistent naming conventions across the repository.

## Naming

- BeaconLink Control Plane
- BeaconLink Agent
- Relay Network
- Runtime Manager
- Runtime Adapter
- Deployment Planner
- Distributed Scheduler
- Workflow Engine
- Policy Engine
- Inventory Service
- Event Streaming Platform

## Architectural Concepts

- Runtime Independent
- Zero Trust
- Capability-Based Scheduling
- Continuous Reconciliation
- Desired State
- Immutable Revisions
- Event-Driven Architecture
- Plugin Framework

---

# Intended Audience

These diagrams are intended for:

- Platform Architects
- Software Engineers
- Infrastructure Engineers
- DevOps Engineers
- SRE Teams
- Security Engineers
- Solution Architects
- Technical Writers
- Operations Teams

---

# Rendering

These diagrams can be rendered using:

- GitHub
- GitLab
- Mermaid Live Editor
- MkDocs Material
- Docusaurus
- Obsidian
- VS Code Mermaid extensions

---

# Repository Summary

| Category     | Diagrams |
| ------------ | -------: |
| Architecture |       10 |
| Database     |        8 |
| Deployment   |        9 |
| Networking   |        9 |
| Security     |        8 |
| Sequence     |       10 |
| State        |        7 |
| UI           |        9 |
| **Total**    |   **70** |

---

# License

This documentation is intended as the architectural reference for the BeaconLink platform. The diagrams describe the logical design, operational behavior, and enterprise architecture of BeaconLink using Mermaid for portability and maintainability.
