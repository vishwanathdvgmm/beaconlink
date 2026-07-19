# RFC-0005: Edge Cache

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing an **Edge Cache** capability that allows
BeaconLink to cache frequently accessed artifacts and metadata closer to
managed sites.

The Edge Cache is intended to reduce bandwidth consumption, improve
deployment performance, increase resiliency during intermittent network
connectivity, and reduce dependency on centralized artifact repositories
for repetitive operations.

The cache operates as an optimization layer and does not become the
authoritative source of platform data.

---

# Motivation

BeaconLink deployments frequently require repeated distribution of:

- Container images
- Application packages
- Runtime binaries
- Configuration bundles
- Deployment manifests
- Plugins
- Software updates
- Policy definitions

In geographically distributed environments, repeatedly downloading the
same artifacts from centralized repositories can result in:

- Increased WAN bandwidth usage
- Higher deployment latency
- Slower recovery
- Regional bottlenecks
- Greater dependency on upstream availability
- Increased infrastructure costs

An Edge Cache can improve deployment efficiency by serving previously
retrieved artifacts from a nearby location.

---

# Goals

The proposed Edge Cache should:

- Reduce repeated downloads
- Improve deployment performance
- Minimize WAN bandwidth usage
- Support offline and intermittent environments
- Integrate transparently with BeaconLink deployments
- Preserve artifact integrity
- Respect cache policies
- Operate independently of application logic

---

# Non-Goals

This RFC does **not** propose:

- Replacing authoritative artifact repositories
- Permanent artifact storage
- Modifying deployment workflows
- Caching mutable runtime state
- Distributed database replication
- General-purpose CDN functionality
- Application-level data caching
- Changing the declarative deployment model

---

# Proposed Design

The Edge Cache acts as an intermediary between managed sites and
upstream artifact sources.

When an artifact is requested:

1. The cache checks for a local copy.
2. If available and valid, the cached artifact is served.
3. Otherwise, the artifact is retrieved from the upstream repository.
4. The retrieved artifact is validated.
5. The artifact is stored locally according to cache policy.
6. Future requests use the cached copy.

```
          Artifact Repository
                   │
                   ▼
             Edge Cache Node
          ┌────────┼────────┐
          ▼        ▼        ▼
       Agent A  Agent B  Agent C
```

The cache is transparent to agents and deployment workflows.

---

# Cacheable Content

Potential cacheable resources include:

- Container images
- Deployment packages
- Runtime binaries
- Plugins
- Configuration bundles
- Update artifacts
- Policy definitions
- Static documentation

Dynamic operational state remains outside the scope of the cache.

---

# Cache Management

The Edge Cache should support configurable policies for:

- Size limits
- Retention periods
- Expiration
- Eviction
- Integrity verification
- Preloading
- Cache warming
- Cleanup

Artifacts remain immutable once cached.

---

# Security Considerations

The Edge Cache must maintain BeaconLink security guarantees.

Cached artifacts should:

- Be integrity verified
- Be cryptographically authenticated where supported
- Remain immutable
- Preserve version information
- Prevent unauthorized modification
- Respect authorization policies
- Maintain audit records for cache operations

The cache must never bypass BeaconLink authentication or deployment
authorization.

---

# Operational Impact

Potential benefits include:

- Faster deployments
- Lower bandwidth consumption
- Improved regional performance
- Better resilience during network outages
- Reduced repository load
- Improved recovery speed
- Lower operational costs

Potential challenges include:

- Cache storage management
- Cache invalidation
- Additional infrastructure
- Monitoring cache health
- Cache consistency policies

---

# Compatibility

The Edge Cache is an optional deployment optimization.

Existing BeaconLink deployments continue operating without modification.

If no cache is available, agents retrieve artifacts directly from the
configured upstream repository.

---

# Alternatives Considered

## Central Repository Only

Always retrieve artifacts from centralized repositories.

Pros:

- Simple architecture
- Single source of artifacts

Cons:

- Higher latency
- Greater bandwidth usage
- Reduced resilience
- Increased repository load

---

## Local Agent Cache

Each agent maintains its own independent artifact cache.

Pros:

- No shared infrastructure
- Simple implementation

Cons:

- Duplicate storage
- Repeated downloads
- Lower cache efficiency
- Limited sharing between agents

---

## External CDN

Use a third-party Content Delivery Network.

Rejected because:

- External dependency
- Limited support for private infrastructure
- Less control over deployment policies
- Additional operational considerations

---

# Open Questions

- Should caches proactively preload commonly used artifacts?
- How should cache eviction policies be configured?
- Should cache replication between sites be supported?
- How should cache health influence deployment decisions?
- Should administrators be able to pin artifacts indefinitely?
- How should storage quotas be enforced?

---

# Related Documents

- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-010 Immutable Update System
- ADR-014 Polyglot Persistence
- ADR-016 Declarative Desired State
- ADR-020 Immutable Deployments

---

# Recommendation

The Edge Cache should be introduced as an optional optimization layer
that accelerates deployments while preserving BeaconLink's immutable artifact
and declarative deployment principles.

By caching validated release artifacts closer to managed infrastructure,
BeaconLink can significantly reduce bandwidth usage, improve deployment
performance, and increase resilience in geographically distributed and
bandwidth-constrained environments.

The initial implementation should focus on immutable artifact caching,
policy-driven lifecycle management, and transparent integration with
existing deployment workflows, ensuring that the cache enhances platform
efficiency without becoming a source of operational complexity.
