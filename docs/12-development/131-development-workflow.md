# 131 - Development Workflow

- **Status:** Approved
- **Version:** 1.0
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the standard development workflow for the BeaconLink project.

The objective is to provide a consistent development process that ensures:

- High code quality
- Stable releases
- Predictable development
- Easy onboarding
- Continuous integration
- Minimal technical debt

All contributors should follow this workflow.

---

# 2. Development Philosophy

BeaconLink follows several core engineering principles.

- Small incremental changes
- Continuous integration
- Test before merge
- Documentation first
- Code review before release
- Keep the repository releasable at all times

---

# 3. Development Lifecycle

Every feature should progress through the following lifecycle.

```text
Idea
    │
    ▼
Documentation
    │
    ▼
Architecture Review
    │
    ▼
Implementation
    │
    ▼
Testing
    │
    ▼
Code Review
    │
    ▼
Merge
    │
    ▼
Release
```

Documentation should always precede implementation.

---

# 4. Feature Development Workflow

Every new feature should follow this order.

1. Define the requirement.
2. Update documentation.
3. Create or update an ADR if required.
4. Design the implementation.
5. Write tests.
6. Implement the feature.
7. Run quality checks.
8. Submit for review.
9. Merge.
10. Release.

---

# 5. Local Development

Typical development workflow.

```text
Pull latest changes
        │
        ▼
Create feature branch
        │
        ▼
Implement feature
        │
        ▼
Run tests
        │
        ▼
Run formatting
        │
        ▼
Commit changes
        │
        ▼
Push branch
        │
        ▼
Open Pull Request
```

---

# 6. Daily Development Cycle

A normal development session should follow this sequence.

```text
git pull
    ↓
Implement
    ↓
go fmt
    ↓
go vet
    ↓
go test ./...
    ↓
Commit
    ↓
Push
```

Developers should avoid committing unformatted or untested code.

---

# 7. Branching Strategy

BeaconLink follows **Trunk-Based Development**.

Permanent branches:

```text
main
```

Temporary branches:

```text
feature/*
bugfix/*
hotfix/*
docs/*
refactor/*
```

Feature branches should be short-lived.

---

# 8. Commit Workflow

Every commit should represent one logical change.

Avoid mixing:

- Refactoring
- Bug fixes
- Features
- Documentation

into the same commit.

---

# 9. Pull Request Workflow

A pull request should contain:

- Clear description
- Related issue (if applicable)
- Test results
- Documentation updates
- Screenshots (UI changes)

Large pull requests should be avoided.

---

# 10. Documentation Workflow

Documentation is part of the implementation.

Whenever code changes:

Evaluate whether documentation must also change.

Possible updates include:

- Architecture
- API
- Configuration
- Deployment
- User documentation
- Examples

Documentation should never lag behind implementation.

---

# 11. Code Review Workflow

Every pull request should be reviewed for:

- Correctness
- Readability
- Maintainability
- Architecture compliance
- Testing
- Documentation

Reviewers should focus on long-term maintainability rather than personal coding preferences.

---

# 12. Testing Workflow

Before opening a pull request, execute:

```bash
go test ./...
```

In addition:

- Unit tests
- Integration tests
- Benchmark tests (where applicable)

New functionality should include corresponding tests.

---

# 13. Formatting Workflow

Before committing:

```bash
go fmt ./...
```

Formatting is mandatory.

Manual formatting differences should never appear in pull requests.

---

# 14. Static Analysis

Run:

```bash
go vet ./...
```

Future versions may include:

- golangci-lint
- staticcheck

Warnings should be addressed before merging.

---

# 15. Dependency Management

When introducing a dependency:

- Verify it is actively maintained.
- Review its license.
- Evaluate its security history.
- Confirm there is no standard library alternative.
- Document the reason for adoption if significant.

Dependencies should be added conservatively.

---

# 16. Version Control Workflow

Recommended workflow:

```text
Pull
    ↓
Implement
    ↓
Test
    ↓
Commit
    ↓
Push
    ↓
Review
    ↓
Merge
```

Avoid long-running feature branches.

---

# 17. Build Workflow

Developers should verify builds locally.

```bash
go build ./...
```

Every executable should compile successfully before submission.

---

# 18. Continuous Integration

Every push should trigger automated checks.

Typical pipeline:

```text
Checkout
    ↓
Restore Dependencies
    ↓
Format Check
    ↓
Static Analysis
    ↓
Unit Tests
    ↓
Integration Tests
    ↓
Build
    ↓
Package
```

Any failing stage blocks merging.

---

# 19. Release Workflow

Release process:

```text
Development
    ↓
Merge
    ↓
Tag
    ↓
CI Release
    ↓
GoReleaser
    ↓
GitHub Release
```

Releases should be automated.

---

# 20. Bug Fix Workflow

Bug fixes should follow:

```text
Reproduce
    ↓
Write Test
    ↓
Fix
    ↓
Verify
    ↓
Merge
```

Whenever practical, create a failing test before implementing the fix.

---

# 21. Refactoring Workflow

Refactoring should:

- Preserve behavior.
- Improve readability.
- Reduce complexity.
- Increase maintainability.

Refactoring should not introduce unrelated functional changes.

---

# 22. Documentation-First Development

Major features should begin with documentation.

Recommended sequence:

```text
RFC (optional)
    ↓
ADR (if required)
    ↓
Documentation
    ↓
Implementation
    ↓
Tests
```

Architecture should drive implementation—not the reverse.

---

# 23. Definition of Done

A task is complete when:

- Code compiles.
- Tests pass.
- Documentation is updated.
- Code is reviewed.
- CI succeeds.
- No known regressions exist.
- Feature meets acceptance criteria.

Anything less is considered incomplete.

---

# 24. Development Checklist

Before merging:

- [ ] Code builds successfully
- [ ] All tests pass
- [ ] Code formatted
- [ ] Static analysis passes
- [ ] Documentation updated
- [ ] No debugging code remains
- [ ] No commented-out code
- [ ] No unnecessary dependencies
- [ ] Pull request reviewed

---

# 25. Engineering Principles

Developers should strive to:

- Write readable code.
- Keep functions small.
- Avoid premature optimization.
- Prefer composition over inheritance.
- Minimize global state.
- Keep APIs stable.
- Make behavior explicit.
- Reduce complexity whenever possible.

---

# 26. Repository Hygiene

The repository should always remain:

- Buildable
- Testable
- Documented
- Releasable

Broken builds should be treated as high-priority issues.

---

# 27. Continuous Improvement

The development workflow should evolve as the project grows.

Improvements should:

- Reduce developer friction.
- Increase automation.
- Improve reliability.
- Shorten feedback loops.
- Preserve code quality.

Workflow changes should be documented and communicated to all contributors.

---

# 28. Summary

The BeaconLink development workflow is designed to keep the project stable, maintainable, and scalable throughout its lifecycle.

Every contribution should follow a consistent path:

```text
Plan
    ↓
Document
    ↓
Implement
    ↓
Test
    ↓
Review
    ↓
Merge
    ↓
ssRelease
```

Following this workflow ensures that BeaconLink remains production-ready while enabling rapid and predictable development.
