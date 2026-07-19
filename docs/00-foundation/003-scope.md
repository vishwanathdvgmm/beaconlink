# 003 - Scope

> **Document ID:** 003
>
> **Version:** 1.0.0
>
> **Status:** Stable
>
> **Last Updated:** 2026-07-16
>
> **Authors:** BeaconLink Engineering Team
>
> **Reviewers:** TBD
>
> **Related Documents:**
>
> - 001-project-vision.md
> - 002-product-goals.md
> - 004-target-users.md
> - 005-project-roadmap.md

---

# Table of Contents

1. Purpose
2. Scope Statement
3. In Scope
4. Out of Scope
5. Future Scope
6. Scope Boundaries
7. Feature Acceptance Criteria
8. Summary

---

# 1. Purpose

This document defines the functional boundaries of BeaconLink.

Its purpose is to clearly establish what BeaconLink is intended to do, what it is not intended to do, and which capabilities may be considered in the future.

Maintaining a well-defined scope helps prevent unnecessary complexity and ensures that development remains aligned with the project's vision.

---

# 2. Scope Statement

BeaconLink is a self-hosting platform that enables users to securely deploy and manage applications on infrastructure they own.

BeaconLink provides the tooling, automation, and networking required to make self-hosting simple.

BeaconLink does not provide hosted compute resources.

---

# 3. In Scope

The following capabilities are considered part of BeaconLink.

## Self-Hosting

- Host applications on user-owned hardware.
- Support desktop, workstation, server, and home lab deployments.
- Support multiple applications on a single device.

---

## Secure Connectivity

- Secure encrypted tunnels.
- Automatic tunnel management.
- Automatic reconnection.
- HTTPS by default.
- Secure device authentication.

---

## Domain Management

- Custom domain support.
- DNS verification.
- Automatic TLS certificate provisioning.
- Domain-to-application mapping.

---

## Application Management

- Runtime detection.
- Deployment automation.
- Process lifecycle management.
- Health monitoring.
- Automatic restart.
- Environment variable management.

---

## User Experience

- Web-based management console.
- Device management.
- Site management.
- Deployment monitoring.
- Application logs.
- Configuration management.

---

## Cross-Platform Support

- Windows
- Linux
- macOS

---

## Extensibility

- Plugin-ready architecture.
- Runtime abstraction.
- Modular component design.

---

# 4. Out of Scope

BeaconLink intentionally does not provide the following services.

## Infrastructure Hosting

BeaconLink does not sell or rent compute resources.

---

## Domain Registration

BeaconLink is not a domain registrar.

Users purchase domains from registrars of their choice.

---

## Email Hosting

BeaconLink does not provide:

- email accounts
- SMTP hosting
- mailbox services

---

## Database Hosting

BeaconLink does not operate managed database services.

Users may self-host databases independently if desired.

---

## CDN Services

BeaconLink is not intended to replace content delivery networks.

---

## Public Cloud Platform

BeaconLink is not a competitor to:

- AWS
- Azure
- Google Cloud

Its purpose is to simplify self-hosting.

---

## Virtual Machine Hosting

BeaconLink does not provide virtual machine infrastructure.

---

## Container Orchestration

BeaconLink is not intended to replace:

- Kubernetes
- Docker Swarm
- Nomad

Although BeaconLink may support Docker containers, it is not a container orchestration platform.

---

# 5. Future Scope

The following capabilities may be introduced in future versions.

- Peer-to-peer networking.
- Multiple relay regions.
- Plugin marketplace.
- Mobile management application.
- High-availability deployments.
- Team collaboration.
- Enterprise administration.
- Advanced analytics.
- Backup integration.
- Edge caching.

Future capabilities must remain consistent with BeaconLink Engineering Principles.

---

# 6. Scope Boundaries

BeaconLink is responsible for:

- deployment
- connectivity
- security
- monitoring
- management

BeaconLink is not responsible for:

- application development
- application business logic
- user content
- infrastructure ownership

Users remain responsible for their own applications.

---

# 7. Feature Acceptance Criteria

Every proposed feature should satisfy the following questions.

### Does it simplify self-hosting?

If not, reconsider the feature.

---

### Does it preserve user ownership?

If not, reject the feature.

---

### Does it improve security?

If it weakens security, redesign it.

---

### Does it fit within BeaconLink's architectural principles?

If not, an ADR is required.

---

### Can it be implemented without compromising modularity?

If not, redesign the architecture first.

---

# 8. Summary

BeaconLink has a deliberately focused scope.

Its responsibility is to simplify secure self-hosting while preserving ownership and maintaining a high-quality deployment experience.

BeaconLink will continue evolving, but it will not expand beyond its core purpose.

Features that do not directly support this purpose should remain outside the project.

> **Focus creates quality. BeaconLink succeeds by doing a small number of things exceptionally well.**
