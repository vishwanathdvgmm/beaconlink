# ADR-010: Immutable Update System

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Platform Update Architecture

---

## Context

BeaconLink consists of multiple independently deployable components,
including agents, relays, API services, console services, and supporting
platform infrastructure.

These components require regular updates for:

- Feature delivery
- Security patches
- Bug fixes
- Performance improvements
- Protocol evolution
- Operational enhancements

Updating distributed infrastructure presents several challenges:

- Partial deployment failures
- Mixed-version environments
- Rollback requirements
- Interrupted updates
- Configuration consistency
- Service availability

The platform requires an update mechanism that minimizes operational
risk while supporting reliable, repeatable software delivery.

---

## Decision

BeaconLink adopts an **immutable update system**.

Every software release is packaged as an immutable deployment artifact.

Updates replace existing software with a validated release artifact
rather than modifying software in place.

The update process consists of:

1. Downloading a verified artifact
2. Validating integrity and compatibility
3. Deploying the new version
4. Performing health verification
5. Activating the release
6. Rolling back automatically if validation fails

Update execution is declarative and driven by the desired platform state
defined within the Site Manifest.

---

## Alternatives Considered

### In-Place Updates

Modify installed software incrementally.

**Rejected because:**

- Difficult rollback
- Configuration drift
- Partial update failures
- Non-repeatable deployments
- Higher operational risk

---

### Manual Updates

Require administrators to perform updates manually.

**Rejected because:**

- Poor scalability
- Error-prone
- Operational inconsistency
- Difficult auditing
- Increased downtime

---

### Package Manager Updates

Rely exclusively on operating system package managers.

**Rejected because:**

- Platform dependency
- Inconsistent environments
- Limited deployment control
- Difficult cross-platform support
- Reduced deployment portability

---

## Consequences

### Positive

- Deterministic deployments
- Reliable rollback
- Simplified recovery
- Immutable release history
- Improved deployment consistency
- Reduced configuration drift
- Better auditability
- Automated deployment workflow

### Negative

- Artifact storage requirements
- Version retention management
- Additional validation logic
- Temporary storage during updates
- Health verification required

---

## Architecture

```
        Desired Version
               │
               ▼
      Artifact Repository
               │
               ▼
      Download & Verify
               │
               ▼
     Compatibility Check
               │
               ▼
          Deploy Update
               │
               ▼
      Health Verification
         │           │
         │ Success   │ Failure
         ▼           ▼
     Activate     Rollback
```

Every deployment progresses through a validated update lifecycle before
becoming active.

---

## Related Decisions

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-009 Protocol Versioning
- ADR-020 Immutable Deployments

---

## Implementation Notes

The update system should:

- Deploy immutable release artifacts
- Verify artifact integrity
- Validate software compatibility
- Support staged deployments
- Perform automated health checks
- Support automatic rollback
- Maintain update audit history
- Resume interrupted updates safely

Artifact distribution mechanisms, signing algorithms, update scheduling,
and rollout policies remain implementation decisions and are outside the
scope of this ADR.

---

## Rationale

An immutable update system provides a predictable and recoverable
deployment process suitable for distributed infrastructure management.

By replacing software with validated release artifacts rather than
modifying existing installations, BeaconLink minimizes deployment risk,
simplifies rollback, and ensures consistent behavior across diverse
environments.

This decision complements the platform's declarative configuration
model, runtime abstraction, and continuous delivery architecture,
providing a reliable foundation for long-term software lifecycle
management.
