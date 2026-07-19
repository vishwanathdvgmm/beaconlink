# RFC-0008: Service Mesh Integration

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes optional integration between BeaconLink and **Service Mesh**
platforms to improve service-to-service communication, security,
observability, and traffic management within Kubernetes and other
orchestrated environments.

The objective is to allow BeaconLink-managed workloads to leverage existing
service mesh capabilities without making a service mesh a mandatory
platform dependency.

---

# Motivation

Modern cloud-native deployments increasingly adopt service meshes such
as:

- Istio
- Linkerd
- Kuma
- Consul Connect

These platforms provide capabilities including:

- Mutual TLS
- Traffic routing
- Service discovery
- Retry policies
- Circuit breaking
- Distributed tracing
- Traffic shaping
- Service authorization

BeaconLink currently manages deployment and runtime configuration but does
not directly integrate with service mesh control planes.

Organizations already operating service meshes may wish to coordinate
BeaconLink deployments with existing networking and security policies.

---

# Goals

The proposed integration should:

- Remain optional
- Support multiple service mesh implementations
- Preserve BeaconLink networking architecture
- Improve service observability
- Support mesh-aware deployments
- Respect existing mesh policies
- Avoid vendor lock-in
- Maintain runtime abstraction

---

# Non-Goals

This RFC does **not** propose:

- Replacing BeaconLink Relay networking
- Making a service mesh mandatory
- Managing the service mesh itself
- Implementing a proprietary service mesh
- Replacing API Gateway functionality
- Altering BeaconLink authentication
- Changing Site Manifest semantics
- Introducing service mesh dependencies outside supported environments

---

# Proposed Design

BeaconLink introduces an optional **Service Mesh Adapter** responsible for
communicating with supported service mesh platforms.

```
             BeaconLink Control Plane
                     │
                     ▼
             Deployment Engine
                     │
          Service Mesh Adapter
                     │
     ┌───────────────┼───────────────┐
     ▼               ▼               ▼
   Istio         Linkerd        Consul Connect
                     │
                     ▼
            Managed Workloads
```

The adapter translates BeaconLink deployment intent into mesh-specific
configuration where supported.

---

# Integration Scope

Potential integration areas include:

- Service registration
- Traffic policies
- Mutual TLS configuration
- Canary deployments
- Blue-green deployments
- Distributed tracing
- Retry policies
- Circuit breaking
- Service identity
- Network policy synchronization

BeaconLink remains the authoritative deployment platform.

The service mesh remains responsible for runtime traffic management.

---

# Deployment Workflow

During deployment BeaconLink may:

1. Deploy application workloads.
2. Register services with the mesh.
3. Configure routing policies.
4. Verify service health.
5. Update traffic allocation.
6. Complete rollout.

Mesh-specific operations remain optional.

Deployments continue normally when no mesh is present.

---

# Security Considerations

BeaconLink and the service mesh should complement one another.

BeaconLink continues to provide:

- Identity management
- Agent authentication
- Deployment authorization
- Infrastructure management

The service mesh may additionally provide:

- Mutual TLS
- Service identities
- Traffic encryption
- Runtime authorization
- Service-level policy enforcement

BeaconLink security decisions remain authoritative.

---

# Operational Impact

Potential benefits include:

- Improved service security
- Rich traffic management
- Better observability
- Safer deployments
- Canary rollout support
- Enhanced resilience
- Reduced application complexity

Potential challenges include:

- Multiple mesh implementations
- Vendor-specific capabilities
- Additional operational complexity
- Compatibility testing
- Adapter maintenance

---

# Compatibility

Service mesh integration is optional.

Existing BeaconLink deployments remain unchanged.

Deployments without a service mesh continue using standard BeaconLink
networking and deployment workflows.

The Runtime Abstraction Layer remains independent of service mesh
implementations.

---

# Alternatives Considered

## No Service Mesh Integration

Continue treating service meshes as external infrastructure.

Pros:

- Simpler architecture
- Fewer dependencies

Cons:

- Limited automation
- Duplicate configuration
- Reduced deployment awareness

---

## Mesh-Specific Implementations

Implement direct support for each service mesh.

Rejected because:

- Vendor lock-in
- Duplicate engineering effort
- Difficult maintenance
- Inconsistent capabilities

---

## BeaconLink-Native Service Mesh

Develop a proprietary service mesh.

Rejected because:

- Extremely large engineering effort
- Existing mature ecosystems
- Increased maintenance burden
- Outside BeaconLink platform scope

---

# Open Questions

- Which service meshes should be supported initially?
- Which mesh capabilities should be standardized?
- How should unsupported mesh features be handled?
- Should BeaconLink automatically detect service mesh availability?
- How should traffic policies integrate with Site Manifests?
- Should mesh telemetry be integrated into BeaconLink observability?

---

# Related Documents

- ADR-004 Runtime Abstraction
- ADR-007 Zero Trust
- ADR-012 Logical Distributed Relay Architecture
- ADR-017 API Gateway Pattern
- ADR-018 Observability Strategy
- ADR-020 Immutable Deployments

---

# Recommendation

BeaconLink should provide an **optional Service Mesh Integration Layer**
through a vendor-neutral adapter architecture rather than tightly
coupling the platform to any specific implementation.

This approach enables organizations to leverage existing service mesh
investments while preserving BeaconLink's runtime abstraction, relay-first
networking model, and declarative deployment architecture.

The initial implementation should focus on common capabilities such as
service registration, traffic management, mutual TLS integration, and
deployment coordination, allowing advanced mesh-specific features to be
added incrementally without compromising platform portability.
