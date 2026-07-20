# 1409 - GitHub Workflow

- **Status:** Approved
- **Category:** Engineering Standards
- **Audience:** Contributors, Maintainers
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the GitHub workflow used throughout the BeaconLink project.

A consistent workflow ensures:

- High code quality
- Predictable development
- Clean Git history
- Reliable automation
- Efficient collaboration
- Safe releases

All contributions must follow this workflow.

---

# 2. Objectives

The workflow should provide:

- Clear development practices
- Automated quality gates
- Reproducible builds
- Traceable changes
- Efficient code reviews
- Stable main branch

The `main` branch should remain releasable at all times.

---

# 3. Branching Strategy

BeaconLink follows a trunk-based development model.

Primary branches:

```text
main
```

Temporary branches:

```text
feature/

bugfix/

hotfix/

release/
```

Long-lived feature branches should be avoided.

---

# 4. Branch Naming

Branch names should be descriptive.

Examples:

```text
feature/add-agent-registration

feature/runtime-podman

bugfix/fix-deployment-timeout

hotfix/cve-2026-001

release/v1.0
```

Use lowercase words separated by hyphens.

---

# 5. Protected Branches

The following branches should be protected:

```text
main
```

Recommended protections:

- Pull requests required
- Passing CI required
- Review required
- Linear history
- No force pushes
- No direct commits

---

# 6. Development Workflow

Typical workflow:

```text
main
    │
    ▼
Create Feature Branch
    │
    ▼
Implement Changes
    │
    ▼
Run Local Tests
    │
    ▼
Open Pull Request
    │
    ▼
CI Validation
    │
    ▼
Code Review
    │
    ▼
Merge
```

All development follows this process.

---

# 7. Commits

Commits should be:

- Small
- Focused
- Atomic
- Buildable
- Easy to review

Avoid large "everything" commits.

---

# 8. Commit Messages

BeaconLink adopts the Conventional Commits specification.

Examples:

```text
feat: add runtime discovery

fix: prevent deployment deadlock

docs: update engineering handbook

refactor: simplify relay routing

test: improve runtime coverage

chore: update dependencies

ci: improve release pipeline
```

Commit messages should describe intent rather than implementation.

---

# 9. Pull Requests

Every change should be introduced through a Pull Request.

A PR should:

- Solve one logical problem
- Be reviewable
- Include tests
- Update documentation if necessary

Large PRs should be split into smaller changes.

---

# 10. Pull Request Title

PR titles should also follow Conventional Commits.

Examples:

```text
feat: implement deployment scheduler

fix: handle runtime shutdown correctly
```

This improves release note generation.

---

# 11. Pull Request Description

A Pull Request should include:

- Summary
- Motivation
- Testing performed
- Documentation updates
- Breaking changes (if any)

Template example:

```text
## Summary

## Testing

## Documentation

## Breaking Changes
```

---

# 12. Code Reviews

Every Pull Request requires review.

Reviewers should verify:

- Correctness
- Readability
- Maintainability
- Tests
- Documentation
- Security
- Performance

Reviews should focus on long-term maintainability.

---

# 13. Review Expectations

Reviews should be:

- Constructive
- Specific
- Respectful
- Actionable

Comments should explain _why_ changes are requested.

---

# 14. Merge Strategy

Preferred merge strategy:

```text
Squash Merge
```

Benefits include:

- Clean history
- One commit per feature
- Easier rollback
- Better changelog generation

Merge commits should generally be avoided.

---

# 15. Rebasing

Feature branches should be rebased regularly.

Preferred:

```bash
git fetch origin

git rebase origin/main
```

Avoid unnecessary merge commits from `main`.

---

# 16. Continuous Integration

Every Pull Request should execute:

- Build
- Unit tests
- Integration tests
- Linting
- Static analysis
- Race detection (where applicable)

Pull requests with failing CI must not be merged.

---

# 17. GitHub Actions

Automation is implemented using:

```text
GitHub Actions
```

Primary workflows include:

- Build
- Test
- Lint
- Release
- Dependency updates (future)

All workflows should be version controlled.

---

# 18. Status Checks

Required status checks include:

- Build
- Unit tests
- Lint
- Static analysis

Future additions may include:

- Security scanning
- Coverage reporting
- SBOM generation

---

# 19. Documentation Updates

Documentation must accompany:

- New features
- API changes
- Configuration changes
- CLI changes
- Architecture changes

Documentation is part of the definition of done.

---

# 20. Testing Requirements

New functionality should include:

- Unit tests
- Integration tests (where applicable)
- Updated fixtures
- Updated golden files (if applicable)

Bug fixes should include regression tests.

---

# 21. Issue Tracking

Work should generally originate from GitHub Issues.

Issues should include:

- Problem description
- Acceptance criteria
- Priority
- Labels

Issues improve traceability.

---

# 22. Labels

Recommended labels:

```text
bug

feature

documentation

enhancement

security

performance

testing

ci

good first issue

help wanted
```

Labels should remain consistent across the repository.

---

# 23. Milestones

Milestones should align with roadmap releases.

Examples:

```text
v0.2.0

v0.3.0

v0.4.0

v1.0.0
```

Milestones provide release visibility.

---

# 24. Project Boards

GitHub Projects may be used to track:

- Backlog
- In Progress
- Review
- Done

Project management should remain lightweight.

---

# 25. Dependency Updates

Dependency updates should:

- Be isolated
- Pass CI
- Include release notes where relevant

Automated dependency tooling may be adopted in the future.

---

# 26. Security Workflow

Security vulnerabilities should:

- Be reported privately
- Be reviewed promptly
- Receive priority
- Avoid public disclosure before remediation

The security policy defines the responsible disclosure process.

---

# 27. Release Workflow

Tagged releases automatically trigger:

- Build
- Packaging
- Artifact generation
- Release publication

Release automation is documented separately.

---

# 28. Branch Cleanup

Merged branches should be deleted.

Temporary branches should not accumulate.

Repository hygiene improves navigation.

---

# 29. Force Pushes

Force pushes are:

Allowed:

- Personal feature branches

Prohibited:

- `main`
- Release branches
- Shared branches

History should remain stable after review begins.

---

# 30. Repository Settings

Recommended repository configuration:

- Branch protection
- Required reviews
- Required status checks
- Signed commits (future)
- Secret scanning
- Dependency graph
- Code scanning

Repository settings should reinforce engineering standards.

---

# 31. Automation

GitHub Actions should automate:

- Formatting validation
- Linting
- Tests
- Builds
- Releases
- Dependency validation

Manual verification should be minimized.

---

# 32. Code Ownership

A `CODEOWNERS` file should define review ownership.

Example:

```text
/docs/          @maintainers

/internal/api/  @backend-team

/internal/ui/   @frontend-team
```

Ownership improves review consistency.

---

# 33. Repository Files

Recommended repository files:

```text
.github/

CODEOWNERS

ISSUE_TEMPLATE/

PULL_REQUEST_TEMPLATE.md

workflows/

SECURITY.md

CONTRIBUTING.md
```

These files standardize collaboration.

---

# 34. Workflow Diagram

```text
Issue
   │
   ▼
Feature Branch
   │
   ▼
Development
   │
   ▼
Local Tests
   │
   ▼
Pull Request
   │
   ▼
GitHub Actions
   │
   ▼
Code Review
   │
   ▼
Squash Merge
   │
   ▼
main
   │
   ▼
Tagged Release
```

---

# 35. Anti-Patterns

Avoid:

- Direct commits to `main`
- Force pushing shared branches
- Skipping code review
- Merging with failing CI
- Oversized Pull Requests
- Mixing unrelated changes
- Missing documentation
- Missing tests
- Ambiguous commit messages
- Long-lived feature branches

---

# 36. Compliance Checklist

Before merging a Pull Request, verify:

| Requirement               | Required |
| ------------------------- | :------: |
| Branch up to date         |    ✓     |
| CI passed                 |    ✓     |
| Lint passed               |    ✓     |
| Tests passed              |    ✓     |
| Documentation updated     |    ✓     |
| Review approved           |    ✓     |
| No merge conflicts        |    ✓     |
| Conventional Commit title |    ✓     |

Changes failing mandatory requirements must not be merged.

---

# 37. Future Enhancements

As the project matures, the workflow may incorporate:

- Signed commits
- Signed releases
- Automated dependency updates
- CodeQL analysis
- Supply chain security (SLSA)
- SBOM generation
- Automatic release notes
- Merge queue support

These enhancements should integrate into the existing workflow without changing its core principles.

---

# 38. Summary

BeaconLink follows a trunk-based GitHub workflow centered on short-lived feature branches, protected `main`, Conventional Commits, mandatory pull requests, automated CI, and squash merging. Every contribution is expected to include appropriate tests, documentation, and successful quality checks before review and merge. By combining lightweight collaboration practices with comprehensive automation, the workflow keeps the `main` branch stable, releasable, and suitable for continuous development.
