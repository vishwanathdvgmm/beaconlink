# 101 - Entity Model

> **Document ID:** 101
>
> **Version:** 1.0.0
>
> **Status:** Draft
>
> **Last Updated:** 2026-07-17
>
> **Authors:** BeaconLink Engineering Team
>
> **Reviewers:** TBD
>
> **Related Documents**
>
> - 030 Domain Model
> - 050 Agent Overview
> - 080 Deployment Overview
> - 081 Site Manifest
> - 100 Database Design
> - 102 Entity Relationships
> - ADR-071 Domain Entity Model

---

# Table of Contents

1. Introduction
2. Purpose
3. Entity Model Objectives
4. Design Principles
5. Entity Architecture
6. Core Entities
7. Aggregate Boundaries
8. Entity Lifecycle
9. Identity Model
10. Future Evolution
11. Summary

---

# 1. Introduction

The BeaconLink Entity Model defines the core business entities that represent
the platform's operational and administrative domains.

Entities encapsulate business identity, lifecycle, and behavior while
remaining independent of storage technologies and transport protocols.

The entity model serves as the canonical representation of BeaconLink domain
objects.

---

# 2. Purpose

The entity model provides:

- Canonical business entities
- Stable resource identities
- Aggregate boundaries
- Domain ownership
- Lifecycle definitions
- Shared terminology
- Persistence independence
- API consistency

The entity model forms the foundation of the BeaconLink domain architecture.

---

# 3. Entity Model Objectives

The entity model is designed to provide:

## Consistency

Represent business concepts uniformly across the platform.

---

## Stability

Maintain durable identities throughout an entity's lifecycle.

---

## Encapsulation

Group related business rules within aggregate boundaries.

---

## Extensibility

Allow new entities without disrupting existing domains.

---

## Independence

Remain independent of APIs, databases, and user interfaces.

---

# 4. Design Principles

BeaconLink entities follow these principles.

## Business Driven

Entities represent business concepts rather than implementation details.

---

## Stable Identity

Every entity possesses a globally unique, immutable identifier.

---

## Aggregate Ownership

Each entity belongs to exactly one aggregate boundary.

---

## Explicit Lifecycle

Entity creation, modification, and retirement follow defined lifecycle
rules.

---

## Persistence Agnostic

Entities do not depend on database schemas or storage engines.

---

# 5. Entity Architecture

```
                 Organization
                      │
        ┌─────────────┴─────────────┐
        ▼                           ▼
      Users                      Sites
                                    │
          ┌─────────────────────────┼────────────────────────┐
          ▼                         ▼                        ▼
       Devices                Deployments               Domains
          │                         │                        │
          └───────────────┬─────────┴───────────────┬────────┘
                          ▼                         ▼
                    Runtime State             Monitoring
```

Business entities are organized into logical aggregates with clear
ownership boundaries.

---

# 6. Core Entities

BeaconLink defines the following primary entities.

| Entity       | Purpose                      |
| ------------ | ---------------------------- |
| Organization | Administrative boundary      |
| User         | Platform identity            |
| Role         | Permission assignment        |
| Site         | Operational boundary         |
| Device       | Managed infrastructure       |
| Deployment   | Desired deployment state     |
| Application  | Deployable workload          |
| Domain       | Network endpoint             |
| Runtime      | Execution environment        |
| Health       | Operational status           |
| Event        | Recorded platform occurrence |
| Audit Record | Historical activity          |

Each entity owns its own lifecycle and business identity.

---

# 7. Aggregate Boundaries

Entities are grouped into aggregates.

| Aggregate       | Root Entity  |
| --------------- | ------------ |
| Identity        | Organization |
| Access Control  | Role         |
| Site Management | Site         |
| Deployment      | Deployment   |
| Runtime         | Device       |
| Monitoring      | Site         |
| Audit           | Audit Record |

Aggregate roots enforce consistency within their boundaries.

---

# 8. Entity Lifecycle

All entities progress through a lifecycle.

```
Created
    ↓
Active
    ↓
Updated
    ↓
Archived
    ↓
Deleted
```

Some entities may omit lifecycle stages where business rules require
permanent retention.

---

# 9. Identity Model

Every entity includes a durable identity.

Common identity attributes include:

- Entity ID
- Entity Type
- Created Timestamp
- Updated Timestamp
- Version
- Owner
- Status

Entity identity remains stable regardless of storage location or API
representation.

---

# 10. Future Evolution

Future enhancements may include:

- Additional domain entities
- Dynamic entity extensions
- Domain-specific metadata
- Custom resource definitions
- Entity versioning
- Soft deletion policies
- Cross-domain projections
- Federated identity relationships

Future evolution should preserve existing entity identities and
aggregate boundaries.

---

# 11. Summary

The BeaconLink Entity Model defines the canonical business entities used
throughout the platform.

By separating business identity from persistence and transport concerns,
BeaconLink establishes a consistent domain language that supports APIs,
deployments, monitoring, and administration across all platform
services.

> **"Entities model the business; aggregates protect its consistency."**

---

# Appendix A — Entity Hierarchy

```
Organization
│
├── Users
├── Roles
└── Sites
     │
     ├── Devices
     ├── Deployments
     ├── Applications
     ├── Domains
     ├── Runtime
     └── Monitoring
```

---

# Appendix B — Aggregate Structure

```
Identity
│
├── Organization
├── User
└── Role

Site
│
├── Site
├── Device
├── Domain
└── Application
```

---

# Appendix C — Common Entity Attributes

| Attribute  | Description                    |
| ---------- | ------------------------------ |
| ID         | Immutable unique identifier    |
| Type       | Entity classification          |
| Name       | Human-readable identifier      |
| Status     | Current lifecycle state        |
| Version    | Optimistic concurrency version |
| Created At | Creation timestamp             |
| Updated At | Last modification timestamp    |

---

# Appendix D — Entity Lifecycle

```
Create
    ↓
Validate
    ↓
Persist
    ↓
Operate
    ↓
Archive
    ↓
Retire
```

---

# Appendix E — Component Responsibilities

| Component      | Responsibility          |
| -------------- | ----------------------- |
| Domain Model   | Business entities       |
| Aggregate Root | Consistency enforcement |
| Repository     | Persistence abstraction |
| Service Layer  | Business operations     |
| API Layer      | Resource representation |
| Database       | Durable storage         |
