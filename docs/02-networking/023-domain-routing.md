# 023 - Domain Routing

> **Document ID:** 023
>
> **Version:** 1.0.0
>
> **Status:** Draft
>
> **Last Updated:** 2026-07-16
>
> **Authors:** BeaconLink Engineering Team
>
> **Reviewers:** TBD
>
> **Related Documents**
>
> - 020 Networking Overview
> - 021 Relay Network
> - 022 Tunnel Routing
> - 030 BeaconLink Tunnel Protocol
> - 041 Certificate Management
> - 043 DNS Management
> - ADR-001 Relay First Networking

---

# Table of Contents

1. Introduction
2. Purpose
3. Design Objectives
4. Domain Routing Philosophy
5. Domain Routing Architecture
6. Domain Lifecycle
7. DNS Resolution
8. Domain Verification
9. Certificate Management
10. Domain Registry
11. Routing Pipeline
12. Multi-Domain Support
13. Wildcard Domains
14. Security
15. Failure Handling
16. Performance
17. Future Evolution
18. Summary

---

# 1. Introduction

Domain Routing is responsible for mapping public Internet domain names
to applications hosted on BeaconLink Agents.

Unlike traditional hosting providers where domains point directly to
servers, BeaconLink domains point to the BeaconLink Relay Network.

The Relay then resolves the destination application using BeaconLink routing
metadata.

---

# 2. Purpose

Domain Routing provides:

- Domain registration
- Domain verification
- DNS validation
- Certificate management
- Site mapping
- Device association
- Traffic routing

Domains are treated as public identifiers rather than infrastructure
addresses.

---

# 3. Design Objectives

The Domain Routing system should be:

- Secure
- Fast
- Deterministic
- Highly available
- Extensible
- Easy to configure

Users should be able to connect a domain without understanding routing
internals.

---

# 4. Domain Routing Philosophy

BeaconLink separates **public identity** from **physical location**.

Traditional hosting:

```
Domain
    ↓
Server IP
    ↓
Application
```

BeaconLink:

```
Domain
    ↓
BeaconLink Relay
    ↓
Site ID
    ↓
Device ID
    ↓
Tunnel ID
    ↓
Application
```

Applications may move between devices without requiring DNS changes.

---

# 5. Domain Routing Architecture

```
                Internet
                    │
              DNS Resolver
                    │
                    ▼
               BeaconLink Relay
                    │
             Domain Registry
                    │
              Site Registry
                    │
             Device Lookup
                    │
             Tunnel Lookup
                    │
              BeaconLink Agent
                    │
             User Application
```

Each component performs one routing responsibility.

---

# 6. Domain Lifecycle

Every domain follows the same lifecycle.

```
Add Domain
    ↓
DNS Configuration
    ↓
Ownership Verification
    ↓
Certificate Request
    ↓
Certificate Issued
    ↓
Domain Activated
    ↓
Traffic Enabled
    ↓
Monitoring
    ↓
Removal
```

Only verified domains may receive traffic.

---

# 7. DNS Resolution

Users configure DNS records pointing to BeaconLink.

Typical flow:

```
example.com
    ↓
DNS Lookup
    ↓
BeaconLink Relay Address
    ↓
Relay Processing
```

Supported record types:

- A
- AAAA
- CNAME
- TXT (verification)

Additional record types may be supported in future versions.

---

# 8. Domain Verification

BeaconLink verifies that users control a domain before activation.

Supported verification methods:

## DNS TXT Record

```
_BeaconLink.example.com
    ↓
Verification Token
    ↓
Validation
```

---

## HTTP Challenge

Future support.

---

## ACME Challenge

Used during certificate issuance.

Verification must succeed before certificates are requested.

---

# 9. Certificate Management

BeaconLink automatically provisions TLS certificates.

Certificate lifecycle:

```
Verification
    ↓
ACME Request
    ↓
Certificate Issued
    ↓
Certificate Installed
    ↓
Automatic Renewal
```

Certificates should renew automatically before expiration.

Supported initially:

- Let's Encrypt
- ACME-compatible providers

Future support:

- Custom certificates
- Enterprise CAs

---

# 10. Domain Registry

The Domain Registry stores routing metadata.

Example record:

| Field          | Description                 |
| -------------- | --------------------------- |
| Domain ID      | Unique domain identifier    |
| Domain Name    | Public hostname             |
| Site ID        | Hosted site                 |
| Device ID      | Hosting Agent               |
| Certificate ID | TLS certificate             |
| Status         | Active / Pending / Disabled |

The registry stores **metadata only**.

Application content remains on the user's device.

---

# 11. Routing Pipeline

Incoming request flow:

```
Visitor
    ↓
DNS
    ↓
BeaconLink Relay
    ↓
TLS Handshake
    ↓
Certificate Selection
    ↓
Domain Registry
    ↓
Site Lookup
    ↓
Device Lookup
    ↓
Tunnel Lookup
    ↓
BeaconLink Agent
    ↓
Runtime
    ↓
Application
    ↓
Response
```

Every stage must complete successfully before traffic is forwarded.

---

# 12. Multi-Domain Support

One application may expose multiple domains.

Example:

```
example.com

example.org

api.example.com

↓

Same Site

↓

Same Application
```

Each domain maintains its own certificate.

---

# 13. Wildcard Domains

Future versions may support wildcard domains.

Example:

```
*.example.com
    ↓
BeaconLink Relay
    ↓
Dynamic Site Resolution
```

Possible use cases:

- Preview deployments
- Multi-tenant applications
- Customer subdomains

Wildcard support should preserve existing security guarantees.

---

# 14. Security

Domain Routing must provide:

- Domain ownership verification
- Certificate validation
- Secure TLS negotiation
- Replay protection
- Request validation

Domains must never be assigned without successful verification.

Private keys should remain protected throughout the certificate lifecycle.

---

# 15. Failure Handling

## Invalid DNS

```
Verification Failed
    ↓
Retry
    ↓
Display Configuration Guidance
```

---

## Certificate Failure

```
Certificate Request Failed
    ↓
Retry
    ↓
Alert User
```

---

## Unknown Domain

```
Incoming Request
    ↓
Domain Not Found
    ↓
Return 404
```

---

## Disabled Domain

```
Request
    ↓
Domain Disabled
    ↓
Return 503
```

Failures should be logged and surfaced through the BeaconLink Console.

---

# 16. Performance

Domain Routing should provide:

- Fast DNS resolution
- Low-latency registry lookups
- Efficient certificate selection
- Minimal routing overhead

Frequently accessed routing metadata may be cached where appropriate.

---

# 17. Future Evolution

Future capabilities may include:

- Geo-aware domain routing
- Regional Relay selection
- Domain aliases
- Automatic failover domains
- Traffic steering
- DNS analytics
- Split-horizon DNS
- Intelligent certificate providers

Future enhancements should remain compatible with existing domain
configurations.

---

# 18. Summary

The BeaconLink Domain Routing system connects public Internet identities to
applications running on user-owned infrastructure.

By separating domain management from infrastructure addressing, BeaconLink
allows applications to move between devices, Relay Nodes, and regions
without requiring users to modify DNS records.

The system provides secure, automated, and extensible domain management
while preserving the simplicity expected by BeaconLink users.

> **"Domains identify applications—not infrastructure."**

---

# Appendix A — Domain Routing Metadata

| Field          | Example     |
| -------------- | ----------- |
| Domain ID      | dom_01HX... |
| Domain Name    | example.com |
| Site ID        | site_001    |
| Device ID      | dev_123     |
| Tunnel ID      | tun_456     |
| Certificate ID | cert_789    |
| Status         | Active      |

---

# Appendix B — Domain State Machine

```
UNREGISTERED
       │
       ▼
REGISTERED
       │
       ▼
VERIFYING
       │
       ▼
VERIFIED
       │
       ▼
CERTIFICATE_PENDING
       │
       ▼
ACTIVE
       │
       ▼
RENEWING
       │
       ▼
ACTIVE

Failure Path

VERIFYING
      │
      ▼
FAILED
      │
      ▼
RETRY
```

---

# Appendix C — Sequence Diagram

```
User          DNS          Relay       Registry      Agent
 |             |             |             |            |
 | Request     |             |             |            |
 |------------>|             |             |            |
 |             | Resolve     |             |            |
 |             |------------>|             |            |
 |             |             | Lookup      |            |
 |             |             |-----------> |            |
 |             |             |             | Find Site  |
 |             |             |             |----------->|
 |             |             |<------------|            |
 |<------------|             |             |            |
```

---

# Appendix D — C4 Container View (ASCII)

```
+----------------------+
|     DNS Provider     |
+----------+-----------+
           |
           v
+----------------------+
|    BeaconLink Relay       |
+----------+-----------+
           |
           v
+----------------------+
|   Domain Registry    |
+----------+-----------+
           |
           v
+----------------------+
|    BeaconLink Agent       |
+----------+-----------+
           |
           v
+----------------------+
| User Application     |
+----------------------+
```
