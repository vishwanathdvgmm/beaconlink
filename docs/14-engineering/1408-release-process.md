# 1408 - Release Process

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers, Release Managers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the standardized release process for the BeaconLink platform.

A consistent release process ensures:

- Predictable releases
- Reproducible builds
- High software quality
- Reliable deployments
- Traceable artifacts
- Operational confidence

Every official BeaconLink release must follow this process.

---

# 2. Objectives

The release process should provide:

- Repeatable automation
- Minimal manual intervention
- Verified artifacts
- Comprehensive validation
- Secure publishing
- Complete documentation

Releases should be deterministic and fully reproducible from source.

---

# 3. Release Principles

Every release must be:

- Versioned
- Tested
- Tagged
- Reproducible
- Signed (future)
- Documented
- Traceable

No release should depend on undocumented manual steps.

---

# 4. Release Types

BeaconLink supports the following release types.

| Type              | Purpose                   |
| ----------------- | ------------------------- |
| Development       | Local testing             |
| Nightly           | Automated validation      |
| Alpha             | Early feature preview     |
| Beta              | Feature complete testing  |
| Release Candidate | Final validation          |
| Stable            | Production release        |
| Hotfix            | Critical production fixes |

Each release type follows the same pipeline with different publication policies.

---

# 5. Release Prerequisites

Before creating a release:

- All planned work is complete.
- Documentation is updated.
- CI passes.
- Required approvals are complete.
- Release notes are prepared.
- Version number is finalized.

No release proceeds with unresolved blocking issues.

---

# 6. Version Selection

Versions follow Semantic Versioning.

Examples:

```text
v0.8.0

v0.9.2

v1.0.0
```

Version selection must comply with the Versioning Policy.

---

# 7. Code Freeze

Stable releases should begin with a code freeze.

During the freeze:

- Only release blockers may be merged.
- New features are deferred.
- Documentation updates remain permitted.
- CI stability is prioritized.

The freeze duration depends on release complexity.

---

# 8. Release Branch

When appropriate, create a release branch.

Example:

```text
release/v1.0
```

Release branches isolate stabilization work from ongoing development.

---

# 9. Final Validation

Prior to release, verify:

- Unit tests
- Integration tests
- End-to-end tests
- Security tests
- Protocol compatibility
- API compatibility
- Documentation completeness

All mandatory checks must succeed.

---

# 10. Continuous Integration

CI should execute:

```bash
go test ./...
```

Additionally:

- Static analysis
- Linting
- Race detection
- Build verification
- Dependency validation
- License checks (future)

Release creation depends on successful CI.

---

# 11. Build Process

Official releases are produced using:

```text
GoReleaser
```

GoReleaser is responsible for:

- Cross compilation
- Packaging
- Checksums
- Archive generation
- Release publication

Manual builds are not official releases.

---

# 12. Supported Platforms

Official binaries should be produced for supported platforms.

Examples:

```text
Linux amd64

Linux arm64

macOS amd64

macOS arm64

Windows amd64
```

Platform support evolves according to the project roadmap.

---

# 13. Embedded Metadata

Every release binary should include:

- Version
- Commit hash
- Build date
- Go version

Metadata is injected during the build process.

---

# 14. Git Tags

Each release is identified by an annotated Git tag.

Example:

```text
v1.2.0
```

Tags are immutable.

Released tags must never be modified or reused.

---

# 15. Release Artifacts

Typical release artifacts include:

```text
Binary archives

Checksums

Release notes
```

Future releases may additionally include:

- SBOM
- Software signatures
- Provenance attestations

---

# 16. Checksums

Every published artifact must include checksums.

Recommended algorithm:

```text
SHA-256
```

Checksums allow users to verify download integrity.

---

# 17. Artifact Naming

Artifacts should follow a consistent naming convention.

Example:

```text
beaconlink_linux_amd64.tar.gz

beaconlink_linux_arm64.tar.gz

beaconlink_darwin_arm64.tar.gz

beaconlink_windows_amd64.zip
```

Names should clearly identify platform and architecture.

---

# 18. Release Notes

Each release must include:

- Overview
- Highlights
- New features
- Bug fixes
- Security fixes
- Breaking changes
- Upgrade guidance
- Known limitations

Release notes should be generated from merged changes and reviewed before publication.

---

# 19. Changelog

The project maintains a cumulative changelog.

Each release updates:

```text
CHANGELOG.md
```

Entries should summarize user-visible changes.

---

# 20. Documentation Review

Before publishing:

- CLI documentation is current.
- API documentation matches implementation.
- Configuration examples are valid.
- Installation guides are verified.
- Migration guides are updated.

Documentation is considered part of the release.

---

# 21. Security Review

Release validation should confirm:

- No known critical vulnerabilities
- Dependency review completed
- Secrets excluded from artifacts
- TLS assets excluded
- Security documentation updated if required

Security issues may delay a release.

---

# 22. Dependency Review

Dependencies should be reviewed for:

- Version updates
- License compatibility
- Known vulnerabilities
- Maintenance status

New dependencies require justification.

---

# 23. Publishing

Stable releases are published through:

- GitHub Releases
- Release artifacts generated by GoReleaser

Publishing should be fully automated from tagged releases.

---

# 24. Verification

After publication, verify:

- Release page
- Artifact availability
- Checksums
- Binary execution
- Version output

Example:

```bash
beacon version
```

Should report the expected release metadata.

---

# 25. Rollback

If a release contains critical defects:

- Halt further distribution.
- Document the issue.
- Prepare a hotfix.
- Publish corrected release notes.

Released Git tags remain immutable.

Do not replace published artifacts.

---

# 26. Hotfix Releases

Hotfixes are reserved for:

- Critical bugs
- Security vulnerabilities
- Production failures

Hotfixes should:

- Contain minimal changes
- Undergo expedited review
- Follow the same validation process

---

# 27. Release Automation

Release automation should perform:

- Build
- Test verification
- Cross compilation
- Packaging
- Checksum generation
- Artifact upload
- Release publication

Manual intervention should be minimized.

---

# 28. Observability

After release, monitor:

- Deployment success
- Crash reports
- Error rates
- Performance regressions
- User feedback

Operational monitoring begins immediately after publication.

---

# 29. Release Ownership

The release manager is responsible for:

- Coordinating the release
- Confirming readiness
- Reviewing release notes
- Verifying published artifacts
- Communicating release status

Release responsibility should be clearly assigned.

---

# 30. Emergency Releases

Emergency releases may bypass:

- Feature freeze timing
- Normal release cadence

They must **not** bypass:

- CI validation
- Critical testing
- Security review
- Versioning policy

Quality standards remain unchanged.

---

# 31. Release Lifecycle

Typical release flow:

```text
Feature Complete
        │
        ▼
Documentation Complete
        │
        ▼
Code Freeze
        │
        ▼
Release Branch
        │
        ▼
CI Validation
        │
        ▼
GoReleaser Build
        │
        ▼
Git Tag
        │
        ▼
GitHub Release
        │
        ▼
Post-Release Verification
        │
        ▼
Operational Monitoring
```

---

# 32. Compliance

Before publication, every release should satisfy the following checklist.

| Requirement            | Required |
| ---------------------- | :------: |
| Version assigned       |    ✓     |
| Git tag created        |    ✓     |
| CI passed              |    ✓     |
| Tests passed           |    ✓     |
| Lint passed            |    ✓     |
| Documentation updated  |    ✓     |
| Release notes prepared |    ✓     |
| Artifacts generated    |    ✓     |
| Checksums generated    |    ✓     |
| Metadata embedded      |    ✓     |
| Artifacts verified     |    ✓     |

Releases that do not satisfy all mandatory requirements must not be published.

---

# 33. Anti-Patterns

Avoid:

- Manual release builds
- Publishing untested binaries
- Reusing Git tags
- Editing published artifacts
- Skipping release notes
- Incomplete changelogs
- Unverified binaries
- Hidden release steps
- Releasing with failing CI
- Publishing without reproducible builds

---

# 34. Future Enhancements

Future releases may incorporate:

- Artifact signing (Sigstore/Cosign)
- SBOM generation
- SLSA provenance
- Supply-chain attestations
- Automated CVE reporting
- Container image signing
- Release health dashboards

These enhancements should integrate with the existing release pipeline without changing its overall workflow.

---

# 35. Summary

BeaconLink releases are fully automated, reproducible, and driven by Git tags using GoReleaser. Every release follows a consistent workflow encompassing validation, documentation, packaging, publication, and post-release verification. Stable releases require successful testing, updated documentation, immutable versioning, and verified artifacts before publication. By minimizing manual intervention and enforcing repeatable quality gates, the release process provides reliable, traceable, and production-ready software deliveries.
