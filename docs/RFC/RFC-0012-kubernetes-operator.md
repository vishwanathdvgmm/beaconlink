# RFC-0012: Kubernetes Operator

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing a **Kubernetes Operator** that enables
BeaconLink to manage Kubernetes-native resources through Custom Resource
Definitions (CRDs) and the Kubernetes reconciliation model.

The Operator would integrate BeaconLink into Kubernetes clusters, allowing
cluster administrators to manage BeaconLink resources using standard
Kubernetes APIs while preserving BeaconLink as the authoritative
infrastructure management platform.

The objective is to provide a cloud-native operational experience
without coupling BeaconLink exclusively to Kubernetes.

---

# Motivation

Kubernetes has become the standard orchestration platform for modern
containerized applications.

Organizations operating Kubernetes often prefer managing
infrastructure through native Kubernetes resources because they provide:

- Declarative configuration
- Kubernetes reconciliation
- Native RBAC
- Standard tooling
- GitOps integration
- Kubernetes events
- API consistency
- Existing operational workflows

BeaconLink currently supports Kubernetes workloads through its runtime
abstraction but does not expose Kubernetes-native management resources.

A Kubernetes Operator would improve integration with existing cloud-
native ecosystems while maintaining BeaconLink platform independence.

---

# Goals

The proposed Operator should:

- Integrate with Kubernetes APIs
- Support declarative resource management
- Preserve BeaconLink reconciliation
- Remain optional
- Support GitOps workflows
- Integrate with Kubernetes RBAC
- Follow Kubernetes Operator best practices
- Avoid Kubernetes-specific platform coupling

---

# Non-Goals

This RFC does **not** propose:

- Rewriting BeaconLink as a Kubernetes application
- Making Kubernetes mandatory
- Replacing the BeaconLink API
- Replacing Site Manifests
- Managing Kubernetes itself
- Supporting every Kubernetes resource
- Creating a Kubernetes distribution
- Eliminating BeaconLink Agents

---

# Proposed Design

The BeaconLink Operator runs inside a Kubernetes cluster and watches BeaconLink
Custom Resources.

```
              Kubernetes API Server
                      │
          Custom Resource Definitions
                      │
                      ▼
               BeaconLink Operator
                      │
          BeaconLink API / Control Plane
                      │
                      ▼
          Reconciliation Engine
                      │
                      ▼
          Managed Infrastructure
```

The Operator translates Kubernetes resources into BeaconLink desired state
without changing BeaconLink's internal architecture.

---

# Custom Resources

Potential Custom Resources include:

- BeaconLinkSite
- BeaconLinkAgent
- BeaconLinkDeployment
- BeaconLinkPolicy
- BeaconLinkSecretReference
- BeaconLinkRuntime
- BeaconLinkRepository
- BeaconLinkRelay

The exact resource model should evolve independently of BeaconLink internal
implementation.

---

# Reconciliation

The Operator follows the standard Kubernetes reconciliation pattern.

For each managed resource it:

1. Observe resource state
2. Validate specification
3. Synchronize with BeaconLink
4. Monitor reconciliation
5. Update resource status
6. Retry when necessary

BeaconLink remains responsible for infrastructure reconciliation.

The Operator reconciles Kubernetes resources with BeaconLink.

---

# Status Reporting

Custom Resources should expose status including:

- Synchronization state
- Validation results
- Health information
- Deployment progress
- Last reconciliation
- Error conditions
- Observed generation

Status should reflect BeaconLink operational state while conforming to
Kubernetes conventions.

---

# Security Considerations

The Operator must preserve BeaconLink Zero Trust architecture.

The implementation should:

- Authenticate with BeaconLink
- Use least-privilege Kubernetes permissions
- Respect Kubernetes RBAC
- Validate resource ownership
- Audit synchronization operations
- Protect credentials
- Verify API compatibility

The Operator should not bypass BeaconLink authorization or policy
enforcement.

---

# Operational Impact

Potential benefits include:

- Native Kubernetes experience
- Simplified GitOps workflows
- Kubernetes tooling compatibility
- Declarative management
- Familiar operational model
- Automated reconciliation
- Improved cloud-native adoption

Potential challenges include:

- CRD lifecycle management
- Kubernetes API compatibility
- Version migration
- Operator maintenance
- Additional deployment complexity

---

# Compatibility

The Kubernetes Operator is entirely optional.

Existing BeaconLink deployments continue operating without Kubernetes.

Organizations not using Kubernetes experience no architectural or
operational changes.

---

# Alternatives Considered

## BeaconLink API Only

Continue using the BeaconLink API directly.

Pros:

- Platform simplicity
- Runtime independence

Cons:

- Less Kubernetes integration
- Separate operational workflow
- Reduced cloud-native experience

---

## Kubernetes Provider Only

Treat Kubernetes as another runtime without native resources.

Pros:

- Minimal engineering effort
- Runtime abstraction preserved

Cons:

- No Kubernetes-native management
- Limited GitOps integration
- Reduced ecosystem compatibility

---

## Kubernetes-Native BeaconLink

Redesign BeaconLink around Kubernetes APIs.

Rejected because:

- Kubernetes dependency
- Reduced portability
- Limited deployment flexibility
- Conflicts with runtime abstraction
- Excludes non-Kubernetes environments

---

# Open Questions

- Which Custom Resources should be introduced first?
- Should the Operator support multiple BeaconLink control planes?
- How should CRD version migration be managed?
- Should BeaconLink automatically register CRDs?
- How should Operator upgrades coordinate with BeaconLink upgrades?
- Should Operator status expose BeaconLink health metrics?

---

# Related Documents

- ADR-004 Runtime Abstraction
- ADR-006 Site Manifest
- ADR-008 Container Strategy
- ADR-016 Declarative Desired State
- ADR-020 Immutable Deployments
- RFC-0011 GitOps Integration

---

# Recommendation

BeaconLink should provide an **optional Kubernetes Operator** that exposes
BeaconLink resources through Kubernetes-native APIs while preserving BeaconLink as
the authoritative control plane for infrastructure management.

The Operator should act as an integration layer—not a replacement for
BeaconLink APIs or reconciliation—allowing Kubernetes users to manage BeaconLink
through familiar declarative workflows and GitOps practices.

This approach strengthens BeaconLink's cloud-native capabilities while
maintaining portability across virtual machines, bare metal, edge
systems, containers, and non-Kubernetes environments, ensuring that
Kubernetes remains a supported platform rather than a platform
dependency.
