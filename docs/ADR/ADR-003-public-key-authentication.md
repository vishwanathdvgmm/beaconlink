# ADR-003: Public-Key Authentication

- **Status:** Accepted
- **Date:** 2026-07-17
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Platform Authentication Architecture

---

## Context

BeaconLink manages infrastructure across multiple organizations, sites, and
networks where communication occurs over potentially untrusted
transport networks.

The platform requires an authentication mechanism that:

- Verifies the identity of every communicating component
- Eliminates reliance on shared secrets
- Supports automated provisioning
- Enables secure key rotation
- Scales to thousands of managed agents
- Integrates with zero-trust networking
- Supports mutual authentication

Traditional username/password authentication and shared API keys are
insufficient for long-lived infrastructure services because they are
difficult to rotate, increase operational risk, and provide limited
identity assurance.

---

## Decision

BeaconLink adopts **public-key cryptography as the primary authentication
mechanism** for all platform components.

Every security principal—including agents, relays, API services, and
other infrastructure components—possesses a unique asymmetric key pair.

Authentication is performed by proving possession of the corresponding
private key while identities are represented by trusted public keys or
certificates.

No component authenticates using reusable passwords or shared secrets.

---

## Alternatives Considered

### Username and Password Authentication

Components authenticate using usernames and passwords.

**Rejected because:**

- Difficult to manage at scale
- Requires secure password storage
- Higher credential compromise risk
- Unsuitable for automated infrastructure
- Weak identity assurance

---

### Shared API Keys

Components authenticate using long-lived API keys.

**Rejected because:**

- Shared secret distribution
- Difficult key rotation
- Limited identity granularity
- Increased operational risk
- Poor support for mutual authentication

---

### Token-Only Authentication

Components authenticate exclusively using bearer tokens.

**Rejected because:**

- Requires trusted token issuer
- Tokens remain transferable if compromised
- Does not establish cryptographic identity
- Better suited for authorization than primary identity

---

## Consequences

### Positive

- Strong cryptographic identity
- Mutual authentication
- No shared secrets
- Simplified credential rotation
- Scalable identity management
- Improved security posture
- Supports zero-trust architecture
- Suitable for automated infrastructure

### Negative

- Public key infrastructure required
- Secure private key storage required
- Certificate lifecycle management
- Additional onboarding complexity
- Revocation mechanisms required

---

## Architecture

```
          Agent
      Private Key
           │
           ▼
 Signs Authentication
           │
           ▼
        Relay/API
           │
           ▼
Verify Signature
Using Public Key
           │
           ▼
Identity Confirmed
```

Authentication is based on cryptographic proof of private key
possession rather than knowledge of a shared secret.

---

## Related Decisions

- ADR-001 Relay-First Networking
- ADR-002 Persistent Connections
- ADR-007 Zero Trust
- ADR-009 Protocol Versioning
- ADR-012 Logical Distributed Relay Architecture

---

## Implementation Notes

The authentication system should:

- Generate unique key pairs for every component
- Protect private keys using secure storage
- Support automated key rotation
- Support certificate renewal
- Validate certificate chains where applicable
- Maintain identity revocation capabilities
- Record authentication events for auditing
- Integrate with authorization services

Specific cryptographic algorithms, certificate formats, and trust models
remain implementation decisions and are outside the scope of this ADR.

---

## Rationale

Public-key authentication provides a scalable and secure identity model
for distributed infrastructure management.

By assigning every platform component a unique cryptographic identity,
BeaconLink eliminates reliance on shared secrets, enables mutual
authentication, simplifies credential lifecycle management, and aligns
with the platform's zero-trust security architecture. This approach
provides a strong foundation for secure communication across diverse
networks while supporting future enhancements in authorization,
certificate management, and automated provisioning.
