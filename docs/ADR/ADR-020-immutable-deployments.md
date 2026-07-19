# ADR-020: Immutable Deployments

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Deployment Architecture

---

## Context

BeaconLink is deployed across diverse environments, including:

- Enterprise data centers
- Cloud infrastructure
- Edge locations
- Customer-managed systems
- Development environments
- Test environments

Deployments must be:

- Predictable
- Repeatable
- Auditable
- Recoverable
- Consistent
- Safe

Traditional mutable deployments modify existing installations by:

- Updating files in place
- Changing runtime configuration
- Installing incremental patches
- Executing imperative deployment scripts

These approaches introduce operational challenges including:

- Configuration drift
- Inconsistent environments
- Difficult rollback
- Partial deployment failures
- Hidden system state
- Non-repeatable deployments

The platform requires a deployment model that guarantees consistency
across all environments while minimizing operational risk.

---

## Decision

BeaconLink adopts an **Immutable Deployment** architecture.

Every deployment consists of a complete, versioned, immutable release
artifact.

Deployments replace an existing running version with a new validated
artifact rather than modifying the existing installation.

Configuration is supplied externally and remains separate from the
deployment artifact.

Deployments progress through a controlled lifecycle:

1. Retrieve release artifact
2. Verify integrity
3. Validate compatibility
4. Deploy new artifact
5. Perform health verification
6. Activate deployment
7. Roll back automatically upon failure

Once deployed, release artifacts are never modified.

Any change requires creating and deploying a new immutable release.

---

## Alternatives Considered

### Mutable In-Place Deployments

Update the existing installation directly.

**Rejected because:**

- Configuration drift
- Difficult rollback
- Partial updates
- Non-repeatable deployments
- Hidden operational changes

---

### Patch-Based Deployments

Deploy only incremental software changes.

**Rejected because:**

- Patch dependency chains
- Complex recovery
- Version inconsistency
- Difficult validation
- Increased operational risk

---

### Manual Installation

Deploy software manually on each target system.

**Rejected because:**

- Error-prone
- Poor scalability
- Inconsistent environments
- Difficult auditing
- Reduced automation

---

## Consequences

### Positive

- Deterministic deployments
- Reliable rollback
- Consistent environments
- Simplified recovery
- Improved auditability
- Reduced configuration drift
- Easier disaster recovery
- Predictable release management

### Negative

- Larger deployment artifacts
- Increased storage requirements
- Complete artifact rebuilds required
- Artifact lifecycle management
- Deployment validation required

---

## Architecture

```
         Immutable Release
              Artifact
                  │
                  ▼
        Integrity Verification
                  │
                  ▼
      Compatibility Validation
                  │
                  ▼
         Deploy New Version
                  │
                  ▼
        Health Verification
             │         │
             │ Pass    │ Fail
             ▼         ▼
         Activate   Rollback
```

Each deployment installs a complete immutable release artifact and
either becomes active after successful validation or automatically
reverts to the previously validated version.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-010 Immutable Update System
- ADR-016 Declarative Desired State

---

## Implementation Notes

The deployment architecture should:

- Build immutable release artifacts
- Keep configuration external to artifacts
- Verify artifact integrity before deployment
- Support automated rollback
- Validate application health before activation
- Preserve deployment history
- Support staged and progressive rollouts
- Ensure deployments are idempotent

The artifact format, deployment orchestration platform, rollout
strategy, health verification mechanism, and release packaging remain
implementation decisions and are outside the scope of this ADR.

---

## Rationale

Immutable deployments provide a predictable and repeatable deployment
model that minimizes operational risk while improving platform
reliability.

By replacing complete release artifacts instead of modifying existing
installations, BeaconLink eliminates configuration drift, simplifies
rollback, and ensures consistent behavior across all deployment
environments.

This decision complements the platform's declarative desired-state
architecture, immutable update system, runtime abstraction, and
container-first deployment strategy, providing a robust foundation for
continuous delivery and long-term operational stability.
