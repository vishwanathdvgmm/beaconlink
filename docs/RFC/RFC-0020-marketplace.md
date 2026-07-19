# RFC-0020: Marketplace

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing an **BeaconLink Marketplace** that enables
organizations to discover, install, publish, and manage reusable BeaconLink
extensions, integrations, templates, policies, workflows, and runtime
providers.

The Marketplace provides a standardized distribution mechanism for
extending the BeaconLink platform without modifying the core system,
supporting both official and community-contributed content through a
secure, versioned ecosystem.

The objective is to foster an extensible platform while preserving
BeaconLink's modular architecture, security model, and runtime independence.

---

# Motivation

As BeaconLink adoption grows, organizations will increasingly require
integrations with external platforms and reusable operational assets.

Examples include:

- Runtime providers
- Notification integrations
- Secret providers
- Authentication providers
- Deployment templates
- Policy libraries
- Workflow templates
- Monitoring integrations
- Infrastructure blueprints
- Community plugins

Without a centralized distribution model, organizations must maintain
custom integrations independently, leading to:

- Duplicate engineering effort
- Inconsistent quality
- Difficult upgrades
- Fragmented documentation
- Poor discoverability
- Limited ecosystem growth

A Marketplace establishes a common mechanism for distributing and
maintaining reusable BeaconLink extensions.

---

# Goals

The proposed Marketplace should:

- Support secure extension distribution
- Enable versioned packages
- Simplify extension discovery
- Support dependency management
- Preserve platform modularity
- Verify package integrity
- Support official and community content
- Remain optional

---

# Non-Goals

This RFC does **not** propose:

- Replacing package managers
- Executing untrusted code automatically
- Hosting arbitrary software
- Replacing Plugin APIs
- Eliminating manual installation
- Becoming a software repository
- Providing commercial licensing
- Replacing enterprise software distribution

---

# Proposed Design

BeaconLink introduces a centralized Marketplace Service responsible for
discovering and distributing reusable platform extensions.

```
             BeaconLink Marketplace
                     │
          Package Repository
                     │
      ┌──────────────┼──────────────┐
      ▼              ▼              ▼
 Plugins      Templates      Integrations
      │              │              │
      └──────────────┼──────────────┘
                     ▼
            BeaconLink Installation
                     │
                     ▼
            Platform Components
```

The Marketplace distributes packages.

BeaconLink remains responsible for installation, validation, and execution.

---

# Marketplace Content

Potential Marketplace packages include:

- Runtime providers
- Secret providers
- Authentication providers
- Notification integrations
- Deployment templates
- Site templates
- Workflow templates
- Policy libraries
- Monitoring integrations
- Dashboard packages

Each package should declare supported BeaconLink versions and dependencies.

---

# Package Lifecycle

Packages progress through the following lifecycle:

1. Authoring
2. Validation
3. Publishing
4. Discovery
5. Installation
6. Upgrade
7. Verification
8. Removal
9. Retirement

Package lifecycle management should remain independent of BeaconLink release
cycles.

---

# Version Management

Marketplace packages should support:

- Semantic versioning
- Compatibility metadata
- Dependency resolution
- Upgrade recommendations
- Rollback support
- Deprecated versions
- Version history
- Release notes

BeaconLink should verify compatibility before installation.

---

# Security Model

Marketplace packages should be treated as untrusted until verified.

The platform should support:

- Package signing
- Publisher verification
- Integrity validation
- Malware scanning
- Dependency verification
- Permission declarations
- Audit logging
- Policy-controlled installation

Installation should require appropriate administrative authorization.

---

# Extension Model

Marketplace packages extend BeaconLink through stable extension points,
including:

- Plugin APIs
- Runtime providers
- Secret providers
- Authentication providers
- Workflow templates
- Policy libraries
- Observability integrations
- Deployment templates

Packages should not depend upon BeaconLink internal implementation details.

---

# Operational Impact

Potential benefits include:

- Faster platform adoption
- Reduced custom development
- Improved ecosystem growth
- Reusable operational assets
- Easier integration
- Community contributions
- Simplified extension management

Potential challenges include:

- Package governance
- Compatibility testing
- Security review
- Publisher verification
- Ecosystem maintenance
- Version management

---

# Compatibility

The Marketplace is optional.

Organizations may continue installing extensions manually or operating
BeaconLink without any Marketplace integration.

Existing Plugin APIs remain unchanged.

---

# Alternatives Considered

## Manual Extension Installation

Continue distributing extensions independently.

Pros:

- Simple implementation
- Full organizational control

Cons:

- Poor discoverability
- Duplicate effort
- Difficult upgrades
- Fragmented ecosystem

---

## Vendor-Specific Repository

Provide only official BeaconLink extensions.

Pros:

- Simplified governance
- Consistent quality

Cons:

- Limited ecosystem growth
- Reduced community innovation
- Higher maintenance burden

---

## Unrestricted Community Repository

Allow unrestricted publication of extensions.

Pros:

- Rapid ecosystem growth
- Broad community participation

Cons:

- Higher security risk
- Quality inconsistency
- Difficult governance
- Increased operational risk

---

# Open Questions

- Should Marketplace packages be reviewed before publication?
- Which package types should be supported initially?
- How should enterprise-private repositories be supported?
- Should packages support automatic updates?
- How should package trust levels be represented?
- Should Marketplace analytics be available to publishers?

---

# Related Documents

- ADR-004 Runtime Abstraction
- ADR-007 Zero Trust
- ADR-016 Declarative Desired State
- RFC-0003 Plugin System
- RFC-0009 Policy Engine
- RFC-0010 Secret Management
- RFC-0014 Workflow Engine

---

# Recommendation

BeaconLink should introduce an **optional Marketplace** that provides a
secure, versioned ecosystem for distributing reusable platform
extensions, templates, integrations, and operational assets while
preserving the platform's modular architecture and Zero Trust security
model.

The preferred implementation builds upon the existing Plugin System and
provider abstractions, using signed packages, compatibility validation,
and policy-governed installation to ensure reliable and secure
extension management.

Initial Marketplace support should prioritize official runtime
providers, secret providers, notification integrations, workflow
templates, policy libraries, and deployment templates, creating a
foundation for a broader community ecosystem without coupling BeaconLink to a
centralized distribution service.
