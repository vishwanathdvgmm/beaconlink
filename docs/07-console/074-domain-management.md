# 074 - Domain Management

> **Document ID:** 074
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
> - 064 Domain Registry
> - 070 Console Overview
> - 073 Site Management
> - 100 Deployment
> - 110 Certificate Management
> - ADR-044 Domain Management

---

# Table of Contents

1. Introduction
2. Purpose
3. Domain Objectives
4. Design Principles
5. Domain Architecture
6. Domain Model
7. Domain Lifecycle
8. Domain Operations
9. Domain Assignment
10. DNS Validation
11. Security Considerations
12. Future Evolution
13. Summary

---

# 1. Introduction

Domain Management provides the administrative interface for managing
Internet domains within the BeaconLink platform.

It enables organizations to register, validate, assign, and manage
domains that expose applications through the BeaconLink Relay network.

While the Relay resolves domains during request routing, Domain
Management maintains the authoritative administrative configuration.

---

# 2. Purpose

Domain Management provides:

- Domain registration
- Ownership validation
- DNS verification
- Site assignment
- Application mapping
- Certificate association
- Lifecycle management

The service acts as the control-plane authority for domain
configuration.

---

# 3. Domain Objectives

The Domain Management system is designed to provide:

## Verified Ownership

Ensure only verified domains can be managed.

---

## Consistent Configuration

Maintain consistent domain configuration across the platform.

---

## Flexible Assignment

Allow domains to be associated with Sites and applications.

---

## Secure Exposure

Support secure publication of Internet-facing services.

---

## Operational Simplicity

Provide straightforward domain administration for operators.

---

# 4. Design Principles

BeaconLink Domain Management follows these principles.

## Ownership Before Usage

Domains shall be validated before assignment.

---

## Separation of Concerns

Administrative configuration is separated from runtime routing.

---

## Stable Identity

A domain remains the same resource regardless of routing changes.

---

## Declarative Configuration

Desired state is stored and propagated to dependent services.

---

## API First

All domain operations shall be available through documented APIs.

---

# 5. Domain Architecture

```
Administrator
      │
      ▼
Domain Management
      │
      ▼
Domain Registry (Control Plane)
      │
 ┌────┼──────────────┐
 ▼    ▼              ▼
DNS  Certificates   Sites
                     │
                     ▼
               Relay Registry
```

Administrative changes are propagated to the Relay after validation.

---

# 6. Domain Model

Each domain maintains structured metadata.

Typical attributes include:

- Domain ID
- Domain name
- Organization
- Site
- Application
- Deployment
- Certificate
- Validation status
- DNS status
- Creation timestamp
- Last updated timestamp
- Labels
- Tags

Domain identifiers remain stable throughout their lifecycle.

---

# 7. Domain Lifecycle

Domains progress through a defined lifecycle.

```
Created
    ↓
Pending Validation
    ↓
Verified
    ↓
Assigned
    ↓
Active
    ↓
Suspended
    ↓
Archived
    ↓
Removed
```

Lifecycle events shall be retained for auditing.

---

# 8. Domain Operations

Supported operations include:

| Operation  | Purpose                              |
| ---------- | ------------------------------------ |
| Register   | Create a new domain                  |
| Verify     | Validate ownership                   |
| Assign     | Associate with a Site or application |
| Update     | Modify configuration                 |
| Suspend    | Disable routing                      |
| Reactivate | Restore service                      |
| Archive    | Preserve inactive domain             |
| Delete     | Remove configuration                 |

Operations shall be validated before becoming effective.

---

# 9. Domain Assignment

A domain may be associated with:

- One organization
- One Site
- One application
- One active deployment
- One certificate

Assignments shall maintain referential integrity across platform
services.

---

# 10. DNS Validation

Ownership validation may use:

- TXT records
- CNAME records
- HTTP challenge
- DNS challenge
- Certificate validation

A domain shall not become active until ownership verification succeeds.

---

# 11. Security Considerations

Domain Management shall:

- Verify domain ownership
- Protect configuration changes
- Require authenticated administration
- Enforce organization boundaries
- Audit all lifecycle events
- Validate certificate associations

Unauthorized domain reassignment shall be prevented.

---

# 12. Future Evolution

Future capabilities may include:

- Wildcard domains
- Automatic certificate provisioning
- Multi-region domain routing
- Domain templates
- Bulk management
- DNS provider integrations
- Traffic policies
- Blue/green domain switching

Future enhancements should preserve domain identity and ownership.

---

# 13. Summary

Domain Management provides the administrative framework for managing
Internet domains within BeaconLink.

By separating configuration from runtime routing, BeaconLink enables secure,
consistent, and scalable domain administration while allowing the Relay
to optimize request resolution independently.

> **"Domains are configured by the control plane and resolved by the data plane."**

---

# Appendix A — Domain Lifecycle

```
Created
     │
     ▼
Pending Validation
     │
     ▼
Verified
     │
     ▼
Assigned
     │
     ▼
Active
     │
     ▼
Suspended
     │
     ▼
Archived
     │
     ▼
Removed
```

---

# Appendix B — Domain Relationships

```
Organization
      │
      ▼
Domain
      │
 ┌────┼────┐
 ▼    ▼    ▼
Site App Certificate
      │
      ▼
 Deployment
```

---

# Appendix C — Domain Metadata

| Attribute         | Description                 |
| ----------------- | --------------------------- |
| Domain ID         | Unique identifier           |
| Domain Name       | Fully qualified domain name |
| Organization      | Owning tenant               |
| Site              | Assigned Site               |
| Application       | Target application          |
| Certificate       | TLS certificate             |
| Validation Status | Ownership verification      |
| DNS Status        | DNS health                  |
| Created At        | Creation timestamp          |

---

# Appendix D — Domain Events

| Event            | Description              |
| ---------------- | ------------------------ |
| DomainRegistered | Domain created           |
| DomainVerified   | Ownership confirmed      |
| DomainAssigned   | Linked to Site           |
| DomainUpdated    | Configuration changed    |
| DomainSuspended  | Routing disabled         |
| DomainArchived   | Preserved for future use |
| DomainDeleted    | Removed from BeaconLink  |

---

# Appendix E — Component Responsibilities

| Component              | Responsibility           |
| ---------------------- | ------------------------ |
| Domain Management      | Administrative lifecycle |
| DNS Validation Service | Ownership verification   |
| Certificate Service    | TLS association          |
| Site Management        | Resource assignment      |
| Relay Domain Registry  | Runtime resolution       |
| Audit Service          | Event history            |
