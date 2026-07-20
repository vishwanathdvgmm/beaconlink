# Security Policy

Thank you for helping keep **BeaconLink** secure.

The security of our users, contributors, and deployments is a top priority. We appreciate responsible disclosure of security vulnerabilities and will make every reasonable effort to investigate and address valid reports promptly.

---

# Supported Versions

Security updates are provided only for supported releases.

| Version |      Supported      |
| ------- | :-----------------: |
| 1.x     |         ✅          |
| 0.x     | ❌ Development only |

Pre-release versions (`alpha`, `beta`, `rc`) should not be considered production-ready and may not receive security updates.

---

# Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub Issues or Pull Requests.**

Instead, report vulnerabilities privately by contacting the project maintainers.

Until a dedicated security contact is established, please use GitHub's **Private Vulnerability Reporting** feature if it is enabled for the repository.

If private reporting is unavailable, contact the maintainers through the project's published contact information.

---

# What to Include

A good vulnerability report should include:

- Description of the issue
- Affected component(s)
- Affected version(s)
- Steps to reproduce
- Proof of concept (if available)
- Potential impact
- Suggested mitigation (optional)
- Relevant logs or screenshots (if applicable)

Providing complete information helps us investigate more quickly.

---

# Response Process

Every report is reviewed by the maintainers.

The general process is:

1. Acknowledge receipt of the report.
2. Verify the vulnerability.
3. Assess severity and impact.
4. Develop and test a fix.
5. Coordinate a release.
6. Publish a security advisory if appropriate.

The exact timeline depends on the complexity and severity of the issue.

---

# Responsible Disclosure

We ask security researchers to:

- Allow reasonable time for investigation.
- Avoid public disclosure before a fix is available.
- Avoid accessing, modifying, or deleting user data.
- Avoid disrupting production systems.
- Report vulnerabilities in good faith.

Responsible disclosure helps protect BeaconLink users while remediation is underway.

---

# Scope

This policy covers vulnerabilities affecting the BeaconLink project, including:

- Agent
- Relay
- API
- Web Console
- BeaconLink Protocol (BLP)
- Authentication
- Authorization
- Cryptography
- Deployment system
- Configuration handling
- Secret management
- Update mechanism

Third-party dependencies should generally be reported to their respective maintainers unless the vulnerability results from BeaconLink's integration or configuration.

---

# Out of Scope

The following are generally considered out of scope:

- Vulnerabilities in unsupported versions
- Denial-of-service attacks requiring unrealistic resources
- Social engineering
- Physical attacks
- Missing security headers in local development environments
- Issues requiring local administrator access without privilege escalation
- Reports based solely on outdated dependencies without a demonstrated impact on BeaconLink
- Feature requests or general hardening recommendations

These issues may still be discussed through normal project channels.

---

# Security Principles

BeaconLink is designed around several core security principles:

- Zero Trust networking
- Mutual authentication
- Least privilege
- Defense in depth
- Secure-by-default configuration
- Explicit authorization
- Encrypted communication
- Immutable deployments where practical
- Comprehensive audit logging
- Observability and monitoring

These principles guide both architecture and implementation.

---

# Security Updates

Security fixes are released as soon as practical after validation.

Critical vulnerabilities may result in:

- Hotfix releases
- Security advisories
- Updated documentation
- Additional mitigation guidance

Security releases follow the project's standard release process.

---

# Dependency Security

The project aims to:

- Keep dependencies up to date.
- Minimize third-party dependencies.
- Review new dependencies before adoption.
- Monitor known vulnerabilities.
- Replace unmaintained libraries when appropriate.

Dependency updates may be released independently of feature releases.

---

# Secure Development

Contributors are expected to:

- Follow the project's engineering standards.
- Avoid introducing unnecessary attack surface.
- Validate all external input.
- Handle secrets securely.
- Use parameterized database queries.
- Avoid logging sensitive information.
- Write tests for security-sensitive functionality.

Security should be considered throughout the development lifecycle rather than only during review.

---

# Cryptography

BeaconLink uses well-established cryptographic algorithms and libraries.

Contributors should:

- Avoid implementing custom cryptography.
- Prefer standard library implementations where appropriate.
- Use modern, widely reviewed cryptographic primitives.
- Follow current best practices for key management and certificate handling.

---

# Secrets

Never commit:

- API keys
- Passwords
- Access tokens
- Private certificates
- Private keys
- Database credentials
- Cloud credentials
- Encryption keys

Configuration examples should use placeholder values.

---

# Security Testing

Security testing may include:

- Static analysis
- Dependency scanning
- Configuration validation
- Authentication testing
- Authorization testing
- Fuzz testing
- Penetration testing
- Threat modeling

Security testing is integrated into the project's overall quality process.

---

# Security Advisories

When appropriate, confirmed vulnerabilities will be communicated through GitHub Security Advisories and documented in the release notes.

Advisories will typically include:

- Affected versions
- Severity
- Description
- Mitigation
- Upgrade guidance
- References

---

# Acknowledgements

We appreciate responsible security research and will acknowledge contributors who responsibly disclose valid vulnerabilities, unless anonymity is requested.

Recognition is provided at the discretion of the project maintainers.

---

# Questions

For general security questions or best practices, please open a GitHub Discussion or Issue if the topic does not involve a confidential vulnerability.

For suspected vulnerabilities, always use the private reporting process described above.

---

# License

Submitting a vulnerability report does not grant ownership of the project or create any contractual relationship. All project code remains licensed under the terms of the project's `LICENSE` file.

---

# Thank You

Security is a shared responsibility. We appreciate everyone who helps improve the safety, reliability, and resilience of BeaconLink through responsible disclosure and thoughtful security practices.
