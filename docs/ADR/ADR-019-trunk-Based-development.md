# ADR-019: Trunk-Based Development

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Source Control Workflow

---

## Context

BeaconLink is developed as a single platform comprising multiple
interdependent components, including:

- Agent
- Relay
- API services
- Console
- Shared libraries
- Deployment tooling
- Documentation

These components evolve together and frequently require coordinated
changes across multiple modules.

Long-lived feature branches introduce several challenges:

- Merge conflicts
- Integration delays
- Divergent implementations
- Delayed feedback
- Reduced deployment confidence
- Complex release management

The platform requires a development workflow that promotes continuous
integration, rapid feedback, and a consistently releasable codebase.

---

## Decision

BeaconLink adopts **Trunk-Based Development**.

The primary branch represents the authoritative source of the platform
and is maintained in a continuously releasable state.

Development is performed through:

- Small, focused changes
- Short-lived feature branches
- Frequent integration into the trunk
- Mandatory automated validation
- Peer code review before merge

Incomplete functionality is isolated using feature flags or equivalent
mechanisms rather than long-running branches.

Every change merged into the trunk must satisfy the platform's quality
standards and be eligible for inclusion in a release.

---

## Alternatives Considered

### GitFlow

Separate long-lived development, release, and hotfix branches.

**Rejected because:**

- Increased branching complexity
- Delayed integration
- Larger merge conflicts
- Slower delivery
- Additional release overhead

---

### Long-Lived Feature Branches

Develop features over extended periods before integration.

**Rejected because:**

- Merge complexity
- Integration risk
- Delayed testing
- Divergent codebases
- Reduced collaboration

---

### Component-Specific Branching

Maintain independent branching strategies for different platform
components.

**Rejected because:**

- Inconsistent workflows
- Difficult cross-component changes
- Fragmented engineering practices
- Higher maintenance burden
- Reduced platform cohesion

---

## Consequences

### Positive

- Continuous integration
- Faster feedback
- Smaller code reviews
- Reduced merge conflicts
- Consistently releasable trunk
- Simplified release process
- Improved collaboration
- Higher deployment confidence

### Negative

- Strong CI discipline required
- Frequent integration expected
- Feature flags may be necessary
- Rapid review turnaround required
- Broken trunk must be corrected immediately

---

## Architecture

```
          Short-Lived Branch
                  │
                  ▼
          Automated Validation
                  │
                  ▼
            Code Review
                  │
                  ▼
              Merge to Trunk
                  │
                  ▼
         Continuously Releasable
             Main Branch
```

Development flows rapidly into the trunk through automated validation
and peer review, ensuring the primary branch remains deployable at all
times.

---

## Related Decisions

- ADR-013 Monorepo Strategy
- ADR-017 API Gateway Pattern
- ADR-018 Observability Strategy
- ADR-020 Immutable Deployments

---

## Implementation Notes

The development workflow should:

- Keep feature branches short-lived
- Require automated builds and tests before merge
- Enforce code review for all changes
- Protect the main branch from direct modification
- Integrate changes frequently
- Use feature flags for incomplete functionality
- Reject changes that fail quality gates
- Restore trunk health immediately after failures

The version control platform, branching permissions, merge strategy,
review policies, and CI implementation remain implementation decisions
and are outside the scope of this ADR.

---

## Rationale

Trunk-Based Development enables BeaconLink to maintain a continuously
integrated and releasable codebase while minimizing the operational
overhead associated with complex branching models.

By encouraging frequent integration, automated validation, and small,
incremental changes, the platform reduces merge conflicts, accelerates
feedback, and improves software quality.

This decision complements the BeaconLink monorepo strategy, continuous
delivery architecture, and immutable deployment model, providing an
efficient and predictable engineering workflow that supports rapid and
reliable platform evolution.
