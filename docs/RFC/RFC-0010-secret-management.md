# RFC-0010: Secret Management

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing a **Secret Management Framework** that
enables BeaconLink to securely store, retrieve, distribute, and rotate
sensitive information without embedding secrets directly into Site
Manifests, deployment artifacts, or application configuration.

The framework introduces a provider abstraction that supports both
BeaconLink-managed secrets and external secret management systems while
maintaining the platform's Zero Trust security architecture.

---

# Motivation

BeaconLink manages infrastructure that frequently requires sensitive
information such as:

- API tokens
- Database credentials
- TLS certificates
- Private keys
- SSH keys
- OAuth credentials
- Container registry credentials
- Cloud provider credentials
- Encryption keys
- Service account tokens

Embedding secrets directly into configuration introduces significant
security risks including:

- Accidental disclosure
- Version control exposure
- Difficult rotation
- Duplicate secret storage
- Weak auditing
- Inconsistent access control

Many organizations already operate enterprise secret management
platforms and expect infrastructure tooling to integrate with them.

---

# Goals

The proposed framework should:

- Eliminate plaintext secrets from configuration
- Support secret rotation
- Integrate with external providers
- Preserve Zero Trust principles
- Support least-privilege access
- Audit secret operations
- Support multiple secret backends
- Maintain runtime abstraction

---

# Non-Goals

This RFC does **not** propose:

- Creating a proprietary enterprise vault
- Replacing existing secret management platforms
- Storing secrets in Site Manifests
- Distributing secrets without authorization
- Managing user passwords
- Performing key generation for external systems
- Encrypting arbitrary application data
- Changing BeaconLink authentication architecture

---

# Proposed Design

BeaconLink introduces a **Secret Provider** abstraction.

Platform components request secrets through the abstraction rather than
accessing storage systems directly.

```
               BeaconLink Services
                      │
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
     Agent        Deployment      API
                      │
                      ▼
              Secret Provider API
                      │
      ┌───────────────┼────────────────┐
      ▼               ▼                ▼
 BeaconLink Store      HashiCorp Vault   Cloud Providers
                                       │
                         AWS • Azure • GCP
```

Application services remain independent of specific secret management
implementations.

---

# Secret Providers

Potential providers include:

- BeaconLink-managed encrypted storage
- HashiCorp Vault
- AWS Secrets Manager
- Azure Key Vault
- Google Secret Manager
- Kubernetes Secrets
- Hardware Security Modules (HSM)
- Future provider plugins

The provider abstraction allows additional implementations without
modifying application logic.

---

# Secret Lifecycle

Secrets progress through the following lifecycle:

1. Creation
2. Storage
3. Authorization
4. Retrieval
5. Distribution
6. Rotation
7. Revocation
8. Deletion

BeaconLink should never permanently cache sensitive secrets beyond what is
operationally necessary.

---

# Secret References

Site Manifests and deployment specifications reference secrets by
logical identifier rather than embedding secret values.

Example:

```
database.password:
    secretRef: production/database/password
```

During deployment, BeaconLink resolves the reference through the configured
Secret Provider.

Secret values are never stored within deployment manifests or version
control repositories.

---

# Secret Rotation

The framework should support:

- Manual rotation
- Scheduled rotation
- Provider-driven rotation
- Certificate renewal
- Credential replacement
- Version tracking

Applications should receive updated secrets without unnecessary service
interruption whenever supported by the underlying runtime.

---

# Security Considerations

Secret Management is a critical security capability.

The framework should:

- Encrypt secrets at rest
- Encrypt secrets in transit
- Authenticate every request
- Authorize every retrieval
- Record audit events
- Support least-privilege access
- Avoid plaintext logging
- Prevent unauthorized export

Secret values should remain visible only to authorized runtime
components.

---

# Operational Impact

Potential benefits include:

- Improved security
- Simplified credential rotation
- Better compliance
- Reduced secret duplication
- Centralized governance
- Provider flexibility
- Improved auditability

Potential challenges include:

- External provider availability
- Provider compatibility
- Secret lifecycle management
- Operational monitoring
- Integration complexity

---

# Compatibility

Existing BeaconLink deployments remain compatible.

Deployments without external secret providers may use BeaconLink-managed
encrypted storage.

Migration from embedded secrets should be incremental and non-breaking.

---

# Alternatives Considered

## Embedded Configuration Secrets

Store secrets directly within Site Manifests.

Pros:

- Simple implementation
- No external dependencies

Cons:

- High security risk
- Difficult rotation
- Poor auditability
- Version control exposure

---

## Environment Variables Only

Distribute secrets exclusively through environment variables.

Pros:

- Broad runtime compatibility
- Simple deployment

Cons:

- Weak lifecycle management
- Limited auditing
- Difficult rotation
- Potential process exposure

---

## Provider-Specific Integrations

Implement direct integrations for each secret platform.

Rejected because:

- Vendor lock-in
- Duplicate engineering effort
- Inconsistent interfaces
- Difficult maintenance

---

# Open Questions

- Which secret providers should be supported initially?
- Should BeaconLink-managed secrets support automatic rotation?
- How should temporary credentials be handled?
- Should secret leasing be supported?
- How should secret access be monitored?
- Should secret replication across regions be configurable?

---

# Related Documents

- ADR-007 Zero Trust
- ADR-010 Immutable Update System
- ADR-014 Polyglot Persistence
- ADR-016 Declarative Desired State
- ADR-017 API Gateway Pattern
- RFC-0003 Plugin System
- RFC-0009 Policy Engine

---

# Recommendation

BeaconLink should introduce a provider-based Secret Management Framework that
treats secrets as first-class managed resources while remaining
independent of any particular secret management product.

The preferred architecture uses logical secret references throughout the
platform, resolving values only when required through authenticated,
authorized, and audited provider interfaces.

This approach enables seamless integration with enterprise secret
management systems, simplifies credential rotation, strengthens
regulatory compliance, and preserves the security and portability
principles established throughout the BeaconLink architecture.
