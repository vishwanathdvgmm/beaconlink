# RFC-0007: Multi-Tenancy

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing **Multi-Tenancy** into the BeaconLink
platform, enabling multiple independent organizations, business units,
or customers to share a single BeaconLink deployment while maintaining strict
isolation of data, identities, resources, and administrative
boundaries.

The objective is to evolve BeaconLink from a single-organization platform
into a secure multi-tenant control plane suitable for managed services,
large enterprises, and SaaS deployments.

---

# Motivation

The current BeaconLink architecture assumes a single administrative domain.

As deployments expand, organizations may require:

- Managed service offerings
- Shared enterprise platforms
- Departmental isolation
- Customer isolation
- Regional administration
- Delegated management
- Independent operational policies
- Resource usage reporting

Deploying a separate BeaconLink instance for every organization introduces:

- Higher infrastructure costs
- Duplicate management
- Operational overhead
- Fragmented monitoring
- Increased maintenance
- More complex upgrades

A multi-tenant architecture enables efficient infrastructure sharing
without compromising security or operational isolation.

---

# Goals

The proposed architecture should:

- Support multiple isolated tenants
- Enforce strong tenant boundaries
- Isolate identities and resources
- Support delegated administration
- Enable tenant-specific policies
- Preserve platform scalability
- Minimize operational overhead
- Remain compatible with existing deployments

---

# Non-Goals

This RFC does **not** propose:

- Sharing resources across tenants by default
- Weakening Zero Trust security
- Multi-master administration
- Cross-tenant authentication
- Distributed billing systems
- Public cloud marketplace integration
- Eliminating site isolation
- Changing the relay-first networking model

---

# Proposed Design

A new **Tenant** becomes the highest logical ownership boundary within
the BeaconLink platform.

Every managed object belongs to exactly one tenant.

```
                BeaconLink Platform
                      │
      ┌───────────────┼───────────────┐
      ▼                               ▼
   Tenant A                       Tenant B
      │                               │
      ▼                               ▼
    Sites                           Sites
      │                               │
      ▼                               ▼
    Agents                         Agents
```

Tenants remain logically isolated while sharing the same BeaconLink control
plane.

---

# Tenant Model

Each tenant owns its own:

- Users
- Roles
- Sites
- Agents
- Deployments
- Policies
- Secrets
- Audit records
- API tokens
- Configuration

Platform administrators retain visibility across tenants according to
their assigned privileges.

---

# Identity and Access Control

Authentication remains centralized.

Authorization becomes tenant-aware.

Every authenticated request includes tenant context.

Authorization decisions evaluate:

- User identity
- Tenant membership
- Assigned roles
- Resource ownership
- Administrative policies

Cross-tenant access is denied unless explicitly authorized through
future federation mechanisms.

---

# Data Isolation

Tenant isolation applies to:

- Databases
- API responses
- Event streams
- Audit logs
- Metrics
- Configuration
- Secrets
- Inventory

Data isolation should be enforced within application services rather
than relying solely on user interface restrictions.

---

# Networking

The relay architecture remains unchanged.

Agents register with a tenant during provisioning.

Relay routing continues to operate using:

- Tenant identity
- Site identity
- Agent identity

Tenant isolation is preserved throughout the communication pipeline.

---

# Operational Impact

Potential benefits include:

- Shared infrastructure
- Lower operational cost
- Tenant isolation
- Simplified management
- Improved scalability
- SaaS deployment support
- Delegated administration

Potential challenges include:

- Authorization complexity
- Data partitioning
- Resource quotas
- Tenant lifecycle management
- Operational monitoring

---

# Security Considerations

Multi-tenancy must preserve BeaconLink Zero Trust principles.

The platform should:

- Enforce tenant isolation
- Validate tenant ownership
- Protect tenant secrets
- Separate audit records
- Prevent cross-tenant information leakage
- Apply tenant-aware authorization
- Record administrative actions

Isolation failures must be treated as security defects.

---

# Compatibility

Existing BeaconLink deployments remain compatible.

Single-tenant deployments become a special case consisting of one
default tenant.

Migration to multi-tenancy should require minimal operational changes.

---

# Alternatives Considered

## Separate BeaconLink Deployment Per Organization

Deploy one BeaconLink instance for every tenant.

Pros:

- Complete isolation
- Simple authorization

Cons:

- Higher infrastructure cost
- Duplicate operations
- Difficult fleet management
- Upgrade overhead

---

## Namespace-Based Isolation Only

Separate resources using logical namespaces.

Rejected because:

- Weak ownership boundaries
- Limited administrative separation
- Difficult policy enforcement
- Increased risk of isolation failures

---

## Shared Administration

Allow unrestricted visibility across organizations.

Rejected because:

- Poor security
- Weak governance
- Regulatory concerns
- Reduced customer trust

---

# Open Questions

- Should tenant-specific resource quotas be supported?
- How should tenant lifecycle management operate?
- Should tenant administrators manage relay deployments?
- How should billing or usage reporting integrate with tenants?
- Should tenants support custom authentication providers?
- How should platform-wide maintenance affect tenant operations?

---

# Related Documents

- ADR-005 Multi-Site Routing
- ADR-006 Site Manifest
- ADR-007 Zero Trust
- ADR-017 API Gateway Pattern
- ADR-018 Observability Strategy
- RFC-0002 Global Relay Network

---

# Recommendation

BeaconLink should evolve toward a **tenant-aware architecture** that enables
multiple independent organizations to securely share a single platform
deployment.

The preferred implementation introduces **Tenant** as the highest
logical ownership boundary, ensuring consistent isolation across
identity, authorization, networking, configuration, telemetry, and
resource management.

By making tenant awareness a core architectural concept rather than an
application-layer feature, BeaconLink can support enterprise and SaaS
deployments while preserving the security, scalability, and operational
simplicity established by the platform's existing architectural
principles.
