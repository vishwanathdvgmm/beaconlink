# 064 - Domain Registry

> **Document ID:** 064
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
> - 020 Networking Overview
> - 030 BeaconLink Tunnel Protocol
> - 060 Relay Overview
> - 061 Request Routing
> - 062 Tunnel Manager
> - 063 Session Manager
> - 065 Relay Discovery
> - 067 Traffic Flow
> - ADR-031 Resource Resolution

---

# Table of Contents

1. Introduction
2. Purpose
3. Registry Objectives
4. Design Principles
5. Registry Architecture
6. Domain Model
7. Resolution Process
8. Registry Updates
9. Consistency
10. Failure Handling
11. Security Considerations
12. Future Evolution
13. Summary

---

# 1. Introduction

The Domain Registry is the authoritative mapping service that resolves
incoming domain names to BeaconLink resources.

Rather than routing directly to network endpoints, BeaconLink resolves a
domain into the appropriate hosted application through a structured
resolution process.

The Domain Registry enables flexible deployments while abstracting
infrastructure details from clients.

---

# 2. Purpose

The Domain Registry provides:

- Domain registration
- Domain lookup
- Domain validation
- Resource mapping
- Registry synchronization
- Resolution metadata

The registry provides the first step in request routing.

---

# 3. Registry Objectives

The Domain Registry is designed to provide:

## Deterministic Resolution

A domain should always resolve to the same resource given identical
configuration.

---

## Fast Lookup

Resolution should introduce minimal latency.

---

## Scalability

The registry should support millions of registered domains.

---

## High Availability

Registry failures should not interrupt active traffic whenever possible.

---

## Consistency

Every Relay should observe the same domain mappings.

---

# 4. Design Principles

BeaconLink follows several principles.

## Identity Before Address

Domains resolve to BeaconLink resources rather than IP addresses.

---

## Hierarchical Resolution

Resolution proceeds through multiple logical stages.

---

## Immutable Identifiers

Internal resource identifiers remain stable even if domains change.

---

## Cached Reads

Frequently used mappings may be cached to improve performance.

---

## Event-Driven Updates

Registry changes propagate through the Relay infrastructure as events.

---

# 5. Registry Architecture

```
Incoming Request
        │
        ▼
Domain Registry
        │
        ▼
Domain Record
        │
        ▼
Site
        │
        ▼
Application
        │
        ▼
Deployment
        │
        ▼
Routing Engine
```

The Domain Registry resolves the requested hostname before routing
begins.

---

# 6. Domain Model

Each domain record represents an externally accessible endpoint.

Typical attributes include:

- Domain name
- Site identifier
- Application identifier
- Deployment identifier
- Status
- Creation timestamp
- Last updated timestamp

The registry stores logical mappings rather than network addresses.

---

# 7. Resolution Process

Each request follows the same resolution pipeline.

```
Receive Request
    ↓
Extract Host Header
    ↓
Lookup Domain
    ↓
Resolve Site
    ↓
Resolve Application
    ↓
Resolve Deployment
    ↓
Pass to Routing Engine
```

Only after successful resolution does routing continue.

---

# 8. Registry Updates

The registry changes whenever:

- Domains are added
- Domains are removed
- Applications are deployed
- Deployments change
- Sites migrate
- Administrative updates occur

Updates should propagate to all Relay nodes with minimal delay.

---

# 9. Consistency

All Relay nodes should maintain a consistent view of domain mappings.

Consistency mechanisms may include:

- Event propagation
- Incremental synchronization
- Version tracking
- Snapshot recovery
- Conflict detection

Temporary inconsistency should be minimized.

---

# 10. Failure Handling

Possible failures include:

- Unknown domain
- Missing application
- Invalid deployment
- Registry synchronization failure
- Corrupted registry entry
- Cache inconsistency

Recovery actions include:

- Retry lookup
- Refresh cached data
- Synchronize registry
- Return appropriate client error
- Log diagnostic information

Registry failures should not affect unrelated domains.

---

# 11. Security Considerations

The Domain Registry shall:

- Validate administrative updates
- Restrict registry modification
- Audit registry changes
- Protect registry integrity
- Prevent unauthorized mappings

Clients shall never influence registry contents directly.

---

# 12. Future Evolution

Future enhancements may include:

- Wildcard domains
- Regional domains
- Vanity domains
- Internationalized domain names (IDN)
- Weighted domain routing
- Policy-based domain resolution
- Multi-region deployments

Future capabilities should preserve existing domain mappings whenever
possible.

---

# 13. Summary

The Domain Registry provides the authoritative mapping between public
domain names and BeaconLink resources.

By separating domain resolution from routing and networking, BeaconLink
supports flexible deployments, scalable routing, and infrastructure
abstraction while maintaining deterministic request resolution.

> **"Domains identify services; the registry identifies where those services live."**

---

# Appendix A — Resolution Pipeline

```
Domain
    ↓
Site
    ↓
Application
    ↓
Deployment
    ↓
Agent
    ↓
Tunnel
```

---

# Appendix B — Domain Record

| Attribute      | Description        |
| -------------- | ------------------ |
| Domain         | Public hostname    |
| Site ID        | Target site        |
| Application ID | Hosted application |
| Deployment ID  | Active deployment  |
| Status         | Active or inactive |
| Version        | Registry revision  |

---

# Appendix C — Registry Events

| Event             | Description               |
| ----------------- | ------------------------- |
| DomainAdded       | New domain registered     |
| DomainUpdated     | Existing mapping modified |
| DomainRemoved     | Mapping deleted           |
| DeploymentChanged | Deployment updated        |
| RegistrySynced    | Synchronization completed |
| RegistryRecovered | Registry restored         |

---

# Appendix D — Resolution States

```
Received
    ↓
Resolved
    ↓
Validated
    ↓
Forwarded
```

Failure path:

```
Lookup
    ↓
Not Found
    ↓
Reject Request
```

---

# Appendix E — Component Responsibilities

| Component       | Responsibility           |
| --------------- | ------------------------ |
| Domain Registry | Resolve domains          |
| Routing Engine  | Resolve destinations     |
| Session Manager | Validate session context |
| Tunnel Manager  | Locate ATP tunnel        |
| Relay Discovery | Maintain Relay topology  |
| Traffic Flow    | Forward packets          |
