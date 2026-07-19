# 044 - Cryptography

> **Document ID:** 044
>
> **Version:** 1.0.0
>
> **Status:** Draft
>
> **Last Updated:** 2026-07-17
>
> **Authors:** BeaconLink Engineering Team
>
> **Reviewers:** TBD
>
> **Related Documents**
>
> - 040 Security Overview
> - 041 Zero Trust
> - 042 Authentication
> - 045 Key Management
> - 046 Certificate Management
> - 030 BeaconLink Tunnel Protocol
> - 047 Secure Updates
> - ADR-008 Approved Cryptographic Standards

---

# Table of Contents

1. Introduction
2. Purpose
3. Cryptographic Objectives
4. Cryptographic Principles
5. Approved Algorithms
6. Cryptographic Architecture
7. Encryption
8. Digital Signatures
9. Hash Functions
10. Random Number Generation
11. Key Exchange
12. Password Handling
13. Cryptographic Agility
14. Security Considerations
15. Future Evolution
16. Summary

---

# 1. Introduction

Cryptography provides the foundation for confidentiality, integrity,
authentication, and authenticity throughout the BeaconLink platform.

BeaconLink exclusively uses modern, standardized cryptographic algorithms
that have undergone extensive public review.

Custom cryptographic algorithms are not permitted.

---

# 2. Purpose

The Cryptography architecture defines:

- Approved cryptographic algorithms
- Encryption standards
- Signature algorithms
- Hash functions
- Random number generation
- Key exchange mechanisms
- Cryptographic design principles

Key lifecycle management is defined separately in **045 Key Management**.

---

# 3. Cryptographic Objectives

BeaconLink cryptography is designed to ensure:

## Confidentiality

Protect sensitive data from unauthorized disclosure.

---

## Integrity

Detect unauthorized modification of data.

---

## Authenticity

Verify the origin of messages and software.

---

## Forward Secrecy

Prevent compromise of historical sessions if long-term keys are exposed.

---

## Replay Protection

Prevent reuse of previously valid messages.

---

## Algorithm Agility

Allow cryptographic algorithms to evolve without redesigning the platform.

---

# 4. Cryptographic Principles

BeaconLink follows these principles.

## Use Proven Standards

Only standardized and publicly reviewed algorithms are permitted.

---

## No Custom Cryptography

BeaconLink shall not implement proprietary encryption algorithms.

---

## Encrypt by Default

Sensitive communication shall always be encrypted.

---

## Minimize Secret Exposure

Secrets should exist in memory only as long as necessary.

---

## Secure Defaults

Weak algorithms shall not be supported.

---

## Cryptographic Agility

Algorithms may be upgraded as security recommendations evolve.

---

# 5. Approved Algorithms

BeaconLink currently approves the following algorithms.

| Purpose                            | Approved Algorithm      |
| ---------------------------------- | ----------------------- |
| Transport Security                 | TLS 1.3                 |
| Symmetric Encryption               | AES-256-GCM             |
| Symmetric Encryption (Alternative) | ChaCha20-Poly1305       |
| Key Exchange                       | X25519                  |
| Digital Signatures                 | Ed25519                 |
| Hashing                            | SHA-256                 |
| Message Authentication             | HMAC-SHA-256            |
| Password Hashing                   | Argon2id                |
| Random Numbers                     | Operating System CSPRNG |

Additional algorithms may be approved following security review.

---

# 6. Cryptographic Architecture

```
Application Data
        │
        ▼
BeaconLink Tunnel Protocol (ATP)
        │
        ▼
TLS 1.3
        │
        ▼
Secure Transport
```

Multiple cryptographic layers provide defense in depth.

---

# 7. Encryption

BeaconLink uses authenticated encryption.

Approved symmetric encryption:

- AES-256-GCM
- ChaCha20-Poly1305

Encryption shall provide:

- Confidentiality
- Integrity
- Authentication

Unencrypted transport of sensitive data is prohibited.

---

# 8. Digital Signatures

Digital signatures provide proof of origin and integrity.

BeaconLink uses:

- Ed25519

Digital signatures are used for:

- Agent authentication
- Software update verification
- Configuration integrity
- Message authenticity

Private signing keys must never leave secure storage.

---

# 9. Hash Functions

Approved cryptographic hash function:

- SHA-256

Hashing is used for:

- Integrity verification
- Message digests
- Certificate fingerprints
- Update verification

Hash functions shall not be used directly for password storage.

---

# 10. Random Number Generation

All security-sensitive randomness shall originate from a
Cryptographically Secure Pseudo-Random Number Generator (CSPRNG).

Examples include:

- `/dev/urandom`
- `getrandom()`
- Windows CNG
- Other operating system CSPRNG implementations

Pseudo-random generators not intended for cryptographic use are prohibited.

---

# 11. Key Exchange

BeaconLink uses modern elliptic-curve key exchange.

Approved algorithm:

- X25519

Key exchange shall provide:

- Forward secrecy
- Session isolation
- Strong confidentiality

Long-term private keys shall not be used directly as session keys.

---

# 12. Password Handling

Where passwords are used, BeaconLink shall:

- Never store plaintext passwords
- Hash passwords using Argon2id
- Generate unique salts
- Apply appropriate work factors
- Support future parameter upgrades

Password verification shall occur using constant-time comparisons.

---

# 13. Cryptographic Agility

Cryptographic algorithms inevitably become obsolete.

BeaconLink is designed to support future algorithm replacement without major
architectural changes.

Algorithm negotiation shall be version-aware where appropriate.

Deprecated algorithms should be phased out through planned migration.

---

# 14. Security Considerations

BeaconLink cryptography shall ensure:

- Strong key sizes
- Secure random generation
- Constant-time cryptographic operations where applicable
- Secure memory handling
- Replay resistance
- Forward secrecy
- Authenticated encryption

Weak or deprecated algorithms shall not be supported.

---

# 15. Future Evolution

Future cryptographic enhancements may include:

- Post-quantum cryptography
- Hybrid key exchange
- Hardware Security Module (HSM) integration
- TPM-backed key protection
- Secure enclave support
- Modern hash algorithm upgrades
- Automated algorithm migration

Future enhancements should preserve interoperability whenever possible.

---

# 16. Summary

Cryptography provides the security foundation for BeaconLink.

By relying on modern, standardized cryptographic algorithms and avoiding
custom implementations, BeaconLink protects communication, identities,
software integrity, and sensitive data against known attack classes.

> **"Do not invent cryptography. Apply it correctly."**

---

# Appendix A — Cryptographic Stack

```
Applications
      │
      ▼
BeaconLink Tunnel Protocol
      │
      ▼
TLS 1.3
      │
      ▼
X25519 Key Exchange
      │
      ▼
AES-256-GCM / ChaCha20-Poly1305
```

---

# Appendix B — Algorithm Reference

| Category               | Algorithm         |
| ---------------------- | ----------------- |
| Key Exchange           | X25519            |
| Digital Signature      | Ed25519           |
| Symmetric Encryption   | AES-256-GCM       |
| Alternative Encryption | ChaCha20-Poly1305 |
| Hash Function          | SHA-256           |
| Password Hash          | Argon2id          |
| MAC                    | HMAC-SHA-256      |
| Transport              | TLS 1.3           |

---

# Appendix C — Cryptographic Usage Matrix

| Component        | Encryption | Signature  | Hashing  |
| ---------------- | ---------- | ---------- | -------- |
| Agent ↔ Relay    | TLS 1.3    | Ed25519    | SHA-256  |
| Software Updates | —          | Ed25519    | SHA-256  |
| ATP Messages     | TLS 1.3    | Optional\* | SHA-256  |
| Configuration    | Optional   | Ed25519    | SHA-256  |
| Password Storage | —          | —          | Argon2id |

\*ATP relies on TLS for transport security. Message-level signatures may be introduced for specific use cases if end-to-end authenticity beyond the transport layer is required.

---

# Appendix D — Cryptographic Lifecycle

```
Generate Keys
      │
      ▼
Store Securely
      │
      ▼
Use
      │
      ▼
Rotate
      │
      ▼
Revoke
      │
      ▼
Destroy
```

---

# Appendix E — Approved vs. Prohibited

| Approved          | Prohibited                    |
| ----------------- | ----------------------------- |
| TLS 1.3           | SSL, TLS 1.0, TLS 1.1         |
| AES-256-GCM       | ECB mode                      |
| ChaCha20-Poly1305 | RC4                           |
| Ed25519           | DSA                           |
| X25519            | Custom key exchange           |
| SHA-256           | MD5, SHA-1                    |
| Argon2id          | Plain SHA-256 password hashes |
| OS CSPRNG         | `rand()`, `Math.random()`     |
