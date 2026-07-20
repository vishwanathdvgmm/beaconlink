# 1407 - Versioning Policy

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the versioning strategy used throughout the BeaconLink project.

A consistent versioning policy enables:

- Predictable releases
- Stable APIs
- Reliable upgrades
- Dependency compatibility
- Clear communication of change

All BeaconLink software components must follow this policy.

---

# 2. Objectives

The versioning system should provide:

- Clear release progression
- Backward compatibility expectations
- Upgrade confidence
- Reproducible builds
- Traceability between source code and releases

Version numbers communicate the stability and compatibility of a release.

---

# 3. Versioning Standard

BeaconLink adopts:

```text
Semantic Versioning 2.0.0
```

Version format:

```text
MAJOR.MINOR.PATCH
```

Example:

```text
v1.4.2
```

---

# 4. Version Components

Semantic versions consist of three components.

| Component | Meaning                          |
| --------- | -------------------------------- |
| MAJOR     | Breaking changes                 |
| MINOR     | New backward-compatible features |
| PATCH     | Backward-compatible bug fixes    |

Each component has a distinct purpose and should only change under the appropriate conditions.

---

# 5. Major Releases

Increment the major version when:

- Public APIs change incompatibly
- Protocol compatibility is broken
- Configuration becomes incompatible
- User-visible breaking behavior is introduced

Examples:

```text
v1.0.0

↓

v2.0.0
```

Major releases should be planned and documented through an RFC or ADR when architectural impact exists.

---

# 6. Minor Releases

Increment the minor version when:

- New features are added
- New APIs are introduced
- New CLI commands are added
- New runtimes are supported
- Existing behavior remains compatible

Examples:

```text
v1.3.0

↓

v1.4.0
```

Minor releases must preserve backward compatibility.

---

# 7. Patch Releases

Increment the patch version when:

- Bugs are fixed
- Performance improves
- Documentation is corrected
- Internal refactoring occurs
- Security fixes remain compatible

Examples:

```text
v1.4.1

↓

v1.4.2
```

Patch releases must not introduce breaking behavior.

---

# 8. Initial Development

Before the first stable release, versions begin with:

```text
v0.x.y
```

During the `0.x` series:

- APIs may evolve
- Breaking changes are permitted
- Rapid iteration is expected

However, breaking changes should still be documented.

---

# 9. Stable Release

The first stable production release is:

```text
v1.0.0
```

Version `1.0.0` indicates:

- Stable APIs
- Stable protocol
- Stable CLI
- Production readiness
- Long-term compatibility expectations

---

# 10. Pre-Release Versions

Pre-release identifiers follow Semantic Versioning.

Examples:

```text
v1.0.0-alpha.1

v1.0.0-alpha.2

v1.0.0-beta.1

v1.0.0-beta.2

v1.0.0-rc.1

v1.0.0-rc.2
```

Pre-release versions are not considered production-ready.

---

# 11. Build Metadata

Build metadata may be appended when necessary.

Example:

```text
v1.2.0+20260720

v1.2.0+git.abcd123
```

Build metadata identifies a specific build but does not affect version precedence.

---

# 12. Git Tags

Every release must have a Git tag.

Format:

```text
v1.0.0
```

Avoid:

```text
1.0

release1

latest

stable
```

Git tags represent immutable release points.

---

# 13. Release Branches

Recommended branching strategy:

```text
main

release/v1.2

hotfix/v1.2.1
```

Long-lived release branches should exist only when maintenance of previous versions is required.

---

# 14. Version Source

The authoritative version source is:

```text
Git tag
```

Build systems derive version information from Git rather than hard-coded values.

---

# 15. Embedded Version Information

Every binary should expose:

- Version
- Commit hash
- Build date
- Go version

Example:

```bash
beacon version
```

Output:

```text
Version:     v1.2.0
Commit:      abcd123
Build Date:  2026-07-20T12:00:00Z
Go Version:  go1.25
```

This information should be embedded during the build process.

---

# 16. Version Package

Version information belongs in:

```text
internal/version
```

The package should expose immutable build metadata.

Example:

```go
version.Version
version.Commit
version.BuildDate
version.GoVersion
```

Values should be injected during the build.

---

# 17. Go Modules

The Go module version follows Git tags.

Example:

```text
v1.3.0
```

The module version should always correspond to the tagged release.

---

# 18. Public APIs

Public APIs under `pkg/` must follow semantic versioning.

Breaking API changes require a major version increment.

Internal packages are exempt because they are not public interfaces.

---

# 19. Protocol Versioning

BeaconLink Protocol (BLP) uses an independent protocol version.

Example:

```text
BLP v1
```

Application version and protocol version are related but distinct.

Protocol compatibility rules are documented separately.

---

# 20. API Versioning

HTTP APIs should use explicit version prefixes.

Example:

```text
/api/v1/

ws://.../v1
```

Future incompatible APIs should introduce a new version rather than modifying existing endpoints.

---

# 21. Configuration Compatibility

Configuration changes should remain backward compatible whenever practical.

Examples:

Acceptable:

- New optional fields
- Additional configuration sections

Breaking:

- Renamed required fields
- Removed configuration keys
- Changed field semantics

Breaking configuration changes require a major release.

---

# 22. Database Schema Versioning

Database schema changes should use ordered migrations.

Migration versions should be monotonic and immutable.

Applied migrations must never be modified after release.

---

# 23. Dependency Versioning

Dependencies should:

- Prefer stable releases
- Avoid floating versions
- Be updated intentionally
- Be reviewed during dependency upgrades

Major dependency upgrades should receive architectural review when they introduce breaking changes.

---

# 24. Compatibility Policy

BeaconLink aims to preserve:

- API compatibility
- CLI compatibility
- Configuration compatibility
- Deployment compatibility

Within a major version.

Breaking changes should be minimized.

---

# 25. Deprecation Policy

Deprecated functionality should:

- Be documented
- Produce migration guidance
- Remain available for at least one minor release when practical
- Be removed only in the next major version

Deprecation notices should be clearly communicated in release notes.

---

# 26. Release Notes

Every release should include:

- Version
- Summary
- New features
- Bug fixes
- Security fixes
- Breaking changes
- Upgrade instructions

Release notes form part of the public release documentation.

---

# 27. Upgrade Path

Major releases should include:

- Migration documentation
- Compatibility notes
- Configuration changes
- Removed features
- Upgrade recommendations

Users should have a clear upgrade path.

---

# 28. Development Builds

Development builds may use versions such as:

```text
v1.3.0-dev

v1.3.0-dev.abcd123
```

Development builds are not considered stable releases.

---

# 29. Nightly Builds

Nightly builds should include unique identifiers.

Example:

```text
v1.4.0-nightly.20260720
```

Nightly builds are intended for testing and should not be used in production.

---

# 30. Version Validation

CI should verify:

- Valid semantic version format
- Matching Git tag
- Embedded version metadata
- Release artifact consistency

Invalid versions should fail the release pipeline.

---

# 31. Documentation

Documentation should reference released versions where applicable.

Examples:

- API documentation
- CLI documentation
- Configuration examples
- Upgrade guides

Documentation should remain synchronized with released software.

---

# 32. Anti-Patterns

Avoid:

- Floating version numbers
- Manual version editing in source code
- Reusing release tags
- Mutating released artifacts
- Breaking changes in patch releases
- Breaking changes in minor releases
- Unversioned public APIs
- Ambiguous release identifiers

---

# 33. Compliance

All BeaconLink releases must follow Semantic Versioning 2.0.0. Release artifacts, Git tags, embedded build metadata, documentation, and API compatibility expectations must remain consistent with the version assigned to each release. Code reviews and release automation should enforce compliance with this policy.

---

# 34. Summary

BeaconLink adopts Semantic Versioning as the authoritative versioning strategy for all releases. Git tags serve as the canonical source of version information, while binaries embed version metadata generated during the build process. Stable APIs, CLI behavior, configuration formats, and deployment workflows are preserved within a major release, with breaking changes reserved for major version increments. This policy provides predictable upgrades, reliable release management, and clear compatibility guarantees for users and contributors alike.
