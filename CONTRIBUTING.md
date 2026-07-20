# Contributing to BeaconLink

First, thank you for your interest in contributing to **BeaconLink**.

BeaconLink is an open-source platform for secure application networking, deployment, and edge connectivity. The project emphasizes simplicity, maintainability, reliability, and engineering excellence over rapid feature development.

Please read this guide before opening an issue or submitting a pull request.

---

# Table of Contents

- Code of Conduct
- Ways to Contribute
- Before You Start
- Development Environment
- Repository Structure
- Development Workflow
- Coding Standards
- Documentation Standards
- Testing Requirements
- Commit Guidelines
- Pull Request Process
- Reporting Bugs
- Requesting Features
- Security Issues
- Community Expectations
- License

---

# Code of Conduct

This project follows the guidelines described in the project's **Code of Conduct**.

By participating, you agree to follow those guidelines and contribute respectfully and professionally.

---

# Ways to Contribute

Contributions are welcome in many forms, including:

- Bug fixes
- New features
- Documentation improvements
- Performance improvements
- Security improvements
- Testing
- Developer tooling
- Examples
- Tutorials
- Architecture discussions
- RFC proposals

Not every contribution requires writing Go code.

---

# Before You Start

Before starting significant work:

1. Search existing issues.
2. Search existing pull requests.
3. Read the relevant documentation.
4. Open an issue for discussion if the change is substantial.

Large features should generally begin with discussion before implementation.

---

# Development Environment

## Requirements

Install the latest stable versions of:

- Go
- Node.js
- Git
- Docker or Podman (optional)
- Make

Refer to the project README for installation instructions.

---

## Clone the Repository

```bash
git clone https://github.com/vishwanathdvgmm/beaconlink.git

cd beaconlink
```

---

## Bootstrap

```bash
make bootstrap
```

or

```bash
./scripts/bootstrap.sh
```

This installs project dependencies and prepares the local development environment.

---

# Repository Structure

The repository follows a documentation-first architecture.

```text
cmd/            Application entry points
internal/       Internal packages
pkg/            Public packages
web/            Frontend application
configs/        Configuration files
deployments/    Deployment assets
docs/           Project documentation
scripts/        Development utilities
test/           Test assets
```

Please place new code in the appropriate package rather than creating new top-level directories.

---

# Development Workflow

BeaconLink follows a trunk-based development model.

Typical workflow:

1. Fork the repository.
2. Create a feature branch.
3. Implement the change.
4. Write tests.
5. Update documentation.
6. Run all checks locally.
7. Submit a Pull Request.

Example:

```bash
git checkout -b feature/add-runtime-health-check
```

Keep branches focused on a single logical change.

---

# Coding Standards

All code should follow the engineering handbook.

General expectations include:

- Write idiomatic Go.
- Keep functions focused.
- Keep packages cohesive.
- Prefer composition over inheritance.
- Avoid global state.
- Use dependency injection.
- Return errors rather than panicking.
- Follow project logging conventions.
- Write readable code before clever code.

Run formatting before committing:

```bash
gofmt ./...
```

and

```bash
goimports ./...
```

---

# Documentation Standards

Documentation is considered part of the implementation.

Update documentation whenever changes affect:

- APIs
- Configuration
- Architecture
- Deployment
- Runtime behavior
- User workflows

Documentation should remain synchronized with the codebase.

---

# Testing Requirements

Every functional change should include appropriate tests.

Depending on the change, this may include:

- Unit tests
- Integration tests
- End-to-end tests
- Regression tests
- Benchmarks

Run tests before opening a Pull Request.

```bash
go test ./...
```

---

# Commit Guidelines

BeaconLink follows the Conventional Commits specification.

Examples:

```text
feat: add runtime scheduler

fix: prevent relay reconnect loop

docs: improve deployment guide

test: add protocol integration tests

refactor: simplify configuration loading

ci: improve release workflow
```

Commit messages should clearly describe the change.

---

# Pull Request Process

Before submitting a Pull Request:

- Ensure CI passes locally where practical.
- Add tests.
- Update documentation.
- Resolve merge conflicts.
- Rebase onto the latest `main`.

A Pull Request should include:

- Summary
- Motivation
- Testing performed
- Documentation updates
- Breaking changes (if applicable)

Small, focused Pull Requests are preferred over large changes.

---

# Code Review

Every Pull Request is reviewed before merging.

Review criteria include:

- Correctness
- Maintainability
- Readability
- Performance
- Security
- Testing
- Documentation

Requested changes should be addressed before merge.

---

# Reporting Bugs

When reporting a bug, include:

- BeaconLink version
- Operating system
- Go version (if applicable)
- Steps to reproduce
- Expected behavior
- Actual behavior
- Relevant logs
- Configuration (redacted if necessary)

Reproducible reports are significantly easier to investigate.

---

# Feature Requests

Feature requests should explain:

- The problem
- Proposed solution
- Alternative approaches
- Expected benefits

Feature requests should focus on user value rather than implementation details.

---

# Security Issues

Do **not** disclose security vulnerabilities publicly.

Please follow the process described in **SECURITY.md**.

Responsible disclosure helps protect users while issues are investigated and resolved.

---

# Community Expectations

We aim to foster a collaborative and professional engineering community.

Contributors are expected to:

- Be respectful.
- Be constructive.
- Accept feedback professionally.
- Review others' work thoughtfully.
- Help improve documentation.
- Prioritize long-term maintainability.

Technical disagreements are expected and should be resolved through discussion and evidence.

---

# Design Philosophy

BeaconLink values:

- Simplicity
- Reliability
- Security
- Observability
- Maintainability
- Reproducibility
- Explicitness over magic
- Documentation-first development

Contributions should align with these principles.

---

# License

By contributing to BeaconLink, you agree that your contributions will be licensed under the project's license.

See the `LICENSE` file for details.

---

# Questions

If you have questions about contributing, please:

- Open a GitHub Discussion (when available).
- Open an Issue for clarification.
- Review the project documentation.

Thank you for helping improve BeaconLink.
