# RFC-0011: GitOps Integration

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing **GitOps Integration** to make Git
repositories an optional source of truth for BeaconLink configuration and
deployment workflows.

Instead of administrators interacting directly with the BeaconLink API or
Console, desired infrastructure state can be managed through version-
controlled repositories. BeaconLink continuously reconciles repository
contents with the deployed infrastructure while preserving its existing
declarative reconciliation model.

The objective is to integrate BeaconLink with established GitOps workflows
without replacing the Site Manifest or existing management interfaces.

---

# Motivation

Many organizations already use Git as the authoritative repository for
infrastructure configuration.

Common Git-based workflows include:

- Infrastructure as Code
- Pull Request reviews
- Change approvals
- Version history
- Compliance auditing
- Rollback through Git history
- Automated CI/CD pipelines
- Multi-environment promotion

Although BeaconLink already uses declarative Site Manifests, configuration is
not currently synchronized directly from Git repositories.

GitOps integration would improve collaboration, traceability, and
deployment automation.

---

# Goals

The proposed integration should:

- Support Git as an optional configuration source
- Preserve declarative reconciliation
- Detect repository changes automatically
- Support branch-based workflows
- Integrate with CI/CD pipelines
- Maintain complete auditability
- Support rollback through Git history
- Remain independent of Git hosting providers

---

# Non-Goals

This RFC does **not** propose:

- Replacing Site Manifests
- Replacing the BeaconLink API
- Requiring Git for every deployment
- Executing arbitrary repository code
- Managing Git repositories
- Replacing CI/CD systems
- Supporting imperative deployments
- Modifying BeaconLink networking architecture

---

# Proposed Design

BeaconLink introduces a **Git Synchronization Service** responsible for
monitoring configured repositories and reconciling approved
configuration into the BeaconLink control plane.

```
            Git Repository
                  │
          Commit / Pull Request
                  │
                  ▼
      Git Synchronization Service
                  │
          Configuration Validation
                  │
                  ▼
            Site Manifest
                  │
                  ▼
      Reconciliation Engine
                  │
                  ▼
         Managed Infrastructure
```

Git remains the configuration source.

BeaconLink remains responsible for reconciliation and deployment.

---

# Repository Structure

Repositories may contain:

- Site Manifests
- Deployment specifications
- Environment configuration
- Policy definitions
- Runtime configuration
- Documentation
- Templates
- Shared configuration

Repository organization remains an implementation decision.

---

# Synchronization Workflow

The synchronization process may include:

1. Detect repository changes
2. Retrieve updated configuration
3. Validate syntax
4. Validate policies
5. Verify compatibility
6. Update BeaconLink desired state
7. Trigger reconciliation
8. Record deployment history

Invalid configuration must never be applied automatically.

---

# Change Management

Git naturally provides:

- Version history
- Pull request reviews
- Branch protection
- Commit signatures
- Approval workflows
- Rollback through revision history

BeaconLink should integrate with these workflows without duplicating Git
functionality.

---

# Security Considerations

Git integration must preserve BeaconLink Zero Trust architecture.

The platform should:

- Authenticate repository access
- Verify repository integrity
- Support signed commits where available
- Validate configuration before reconciliation
- Audit synchronization events
- Prevent unauthorized repository modification
- Restrict repository permissions

Git repositories remain external dependencies and must not become
implicit trust boundaries.

---

# Operational Impact

Potential benefits include:

- Infrastructure as Code workflows
- Complete configuration history
- Simplified collaboration
- Easier rollback
- Improved compliance
- Automated deployments
- Consistent change management

Potential challenges include:

- Repository availability
- Merge conflict management
- Synchronization delays
- Validation complexity
- Repository governance

---

# Compatibility

GitOps integration is optional.

Existing BeaconLink deployments continue operating through the Console, API,
or other management interfaces.

Organizations may adopt Git synchronization incrementally.

---

# Alternatives Considered

## BeaconLink as the Only Source of Truth

Continue storing configuration exclusively within BeaconLink.

Pros:

- Simpler architecture
- No external dependencies

Cons:

- Reduced collaboration
- Limited version control
- Separate operational workflows

---

## CI/CD Push Model

Allow CI/CD systems to push configuration directly into BeaconLink.

Pros:

- Simple automation
- Existing tooling

Cons:

- Reduced reconciliation visibility
- Difficult drift detection
- Less resilient synchronization

---

## Git as the Primary Database

Treat Git as the platform database.

Rejected because:

- Limited operational querying
- Poor runtime state representation
- Git is optimized for version control, not operational state
- Increased implementation complexity

---

# Open Questions

- Which Git hosting platforms should be supported initially?
- Should synchronization be event-driven, polling-based, or both?
- How should merge conflicts affect reconciliation?
- Should BeaconLink support multiple repositories per deployment?
- How should environment promotion workflows be modeled?
- Should Git commit signatures be mandatory?

---

# Related Documents

- ADR-006 Site Manifest
- ADR-010 Immutable Update System
- ADR-013 Monorepo Strategy
- ADR-016 Declarative Desired State
- ADR-020 Immutable Deployments
- RFC-0009 Policy Engine

---

# Recommendation

BeaconLink should provide **optional GitOps Integration** that treats Git as
an external source of declarative configuration while preserving the
existing Site Manifest and reconciliation architecture.

The preferred design introduces a dedicated synchronization service that
continuously validates and imports approved repository changes before
triggering BeaconLink reconciliation.

By separating configuration management from deployment execution, BeaconLink
can integrate seamlessly with established GitOps workflows while
maintaining its principles of declarative infrastructure, immutable
deployments, auditability, and operational consistency.
