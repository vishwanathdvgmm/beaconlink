# 127 - Technology Stack

- **Status:** Approved
- **Version:** 1.0
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the official implementation technology stack for the BeaconLink platform.

It serves as the single source of truth for all programming languages, frameworks, libraries, tooling, databases, build systems, and supporting technologies used throughout the project.

Technology selections are made with the following priorities:

1. Reliability
2. Maintainability
3. Simplicity
4. Performance
5. Long-term stability

BeaconLink intentionally favors mature, production-proven technologies over rapidly changing ecosystems.

---

# 2. Design Principles

Technology choices should follow these principles.

- Prefer the Go standard library whenever practical.
- Minimize external dependencies.
- Prefer stable technologies over experimental ones.
- Keep the platform cloud agnostic.
- Design for portability across operating systems.
- Optimize for operational simplicity.
- Build modular components with clear responsibilities.
- Favor explicit code over framework magic.

---

# 3. Core Programming Language

## Selected Technology

**Go**

### Purpose

Primary implementation language for all backend services.

### Used By

- Agent
- Relay
- API Server
- CLI
- Installer
- Shared libraries

### Rationale

- Excellent concurrency model
- High-performance networking
- Static binaries
- Cross-platform support
- Fast compilation
- Excellent standard library
- Mature ecosystem
- Cloud-native ecosystem adoption

---

# 4. Repository Strategy

## Selected Technology

**Go Monorepo**

### Structure

A single Git repository containing all BeaconLink components.

### Components

- Agent
- Relay
- API
- CLI
- Installer
- Web Console
- SDKs
- Documentation

### Benefits

- Single version history
- Shared modules
- Easier refactoring
- Unified CI/CD
- Consistent coding standards
- Simplified dependency management

---

# 5. Frontend

## Framework

React

## Language

TypeScript

## Build Tool

Vite

## Styling

Tailwind CSS

## Icons

Lucide

## State Management

- TanStack Query
- Zustand

### Purpose

Implements the Beacon Console web application.

---

# 6. Backend HTTP Framework

## Selected Technology

Chi

### Purpose

HTTP routing and middleware.

### Rationale

- Lightweight
- Idiomatic Go
- Standard library compatible
- Excellent middleware ecosystem
- Minimal abstraction

---

# 7. API Technologies

## REST API

Used for:

- Resource management
- Authentication
- Administrative operations
- Configuration

---

## WebSocket

Used for:

- Live dashboard
- Event streaming
- Deployment progress
- Log streaming
- Real-time updates

---

## Future

Possible adoption of gRPC where internal service communication requires high throughput.

---

# 8. Database

## Primary Database

PostgreSQL

### Purpose

Persistent storage for platform data.

### Stores

- Sites
- Agents
- Deployments
- Users
- Policies
- Audit logs
- Metadata

### Rationale

- Mature
- ACID compliant
- Reliable
- Excellent tooling
- Widely supported

---

# 9. Cache

## Selected Technology

Redis

### Usage

- Session storage
- Distributed locks
- Rate limiting
- Temporary state
- Short-lived cache

Redis is not considered the system of record.

---

# 10. Event Bus

## Initial Implementation

Go Channels

### Purpose

Internal asynchronous communication.

---

## Future

NATS

Planned when multiple distributed services require external event streaming.

Kafka is intentionally not part of the initial architecture.

---

# 11. Database Access

## Driver

database/sql

## Query Generation

sqlc

### Rationale

- Compile-time query validation
- Type-safe generated code
- High performance
- SQL-first development
- No runtime ORM overhead

Object-relational mappers (ORMs) are intentionally avoided.

---

# 12. Database Migrations

## Selected Technology

golang-migrate

### Purpose

Version-controlled database schema management.

---

# 13. Configuration

## Selected Technology

koanf

### Sources

- YAML
- Environment Variables
- Command-line Flags
- JSON
- Configuration Files

### Rationale

Flexible and modular configuration management.

---

# 14. Logging

## Selected Technology

Go log/slog

### Format

Structured JSON logging.

### Features

- Structured fields
- Log levels
- Context propagation
- Standard library support

---

# 15. Command Line

## Selected Technology

Cobra

### Purpose

Beacon CLI implementation.

### Responsibilities

- Deployment
- Site management
- Agent management
- Diagnostics
- Administration

---

# 16. Testing

## Frameworks

- Go testing package
- testify
- httptest

### Testing Levels

- Unit Testing
- Integration Testing
- End-to-End Testing
- Benchmark Testing

Heavy testing frameworks are intentionally avoided.

---

# 17. Security

## Authentication

- Mutual TLS (mTLS)
- JWT

## Future

OpenID Connect (OIDC)

---

## Cryptography

Go standard library

External cryptographic libraries should only be introduced when absolutely necessary.

---

# 18. Serialization

## Current

JSON

### Usage

- REST APIs
- Configuration
- Metadata
- Administrative interfaces

---

## Future

Binary protocol for BeaconLink Protocol (BLP).

---

# 19. Observability

## Metrics

Prometheus

---

## Tracing

OpenTelemetry

---

## Profiling

Go pprof

---

## Health Monitoring

Native HTTP health endpoints.

---

# 20. Runtime Support

BeaconLink is runtime-independent.

Supported runtimes include:

- Docker
- Podman
- containerd
- Native Processes

Additional runtimes may be added in future releases.

---

# 21. Plugin System

## Strategy

RPC-based plugin architecture.

### Rationale

Go plugins are intentionally avoided due to portability limitations.

---

# 22. Build System

## Primary

Go Toolchain

Commands include:

- go build
- go test
- go generate
- go fmt
- go vet

---

## Future

Taskfile may be introduced to simplify developer workflows.

---

# 23. Continuous Integration

## Platform

GitHub Actions

### Responsibilities

- Build verification
- Testing
- Linting
- Security scanning
- Release automation

---

# 24. Release Management

## Selected Technology

GoReleaser

### Responsibilities

- Cross-platform binaries
- Checksums
- Release archives
- GitHub Releases

---

# 25. Containerization

Supported technologies include:

- Docker
- Podman
- OCI Images

Container images should remain runtime-independent.

---

# 26. Dependency Management

## Package Manager

Go Modules

### Policy

- Minimize third-party dependencies.
- Prefer actively maintained projects.
- Remove unused dependencies promptly.
- Review new dependencies before adoption.

---

# 27. Technology Decision Matrix

| Area                 | Selected Technology                   |
| -------------------- | ------------------------------------- |
| Programming Language | Go                                    |
| Repository           | Go Monorepo                           |
| Frontend             | React + TypeScript                    |
| UI Styling           | Tailwind CSS                          |
| Build Tool           | Vite                                  |
| HTTP Router          | Chi                                   |
| API                  | REST + WebSocket                      |
| Database             | PostgreSQL                            |
| Cache                | Redis                                 |
| Event Bus            | Go Channels → NATS                    |
| Database Layer       | database/sql + sqlc                   |
| Database Migrations  | golang-migrate                        |
| Configuration        | koanf                                 |
| Logging              | log/slog                              |
| CLI                  | Cobra                                 |
| Testing              | testing + testify                     |
| Metrics              | Prometheus                            |
| Tracing              | OpenTelemetry                         |
| Profiling            | pprof                                 |
| Authentication       | mTLS + JWT                            |
| Serialization        | JSON                                  |
| Runtime Support      | Docker / Podman / containerd / Native |
| Plugin System        | RPC                                   |
| Build                | Go Toolchain                          |
| CI/CD                | GitHub Actions                        |
| Releases             | GoReleaser                            |

---

# 28. Future Technology Considerations

The following technologies are under evaluation for future releases and are not part of the approved implementation stack.

- NATS
- OpenID Connect (OIDC)
- WebAssembly Runtime
- Multi-region federation
- Plugin marketplace
- Distributed scheduler
- AI-assisted operations
- GitOps integrations
- Kubernetes Operator

These technologies require an approved Architecture Decision Record (ADR) before adoption.

---

# 29. Technology Governance

Technology changes shall follow the project's Architecture Decision Record (ADR) process.

No production dependency shall be introduced without documented justification.

The approved technology stack should remain stable across major releases to ensure maintainability, predictability, and long-term support.

---

| Area                 | Selected Technology | Status      |
| -------------------- | ------------------- | ----------- |
| Programming Language | Go                  | ✅ Approved |
| Frontend             | React + TypeScript  | ✅ Approved |
| Router               | Chi                 | ✅ Approved |
| Database             | PostgreSQL          | ✅ Approved |
| Cache                | Redis               | ✅ Approved |
| Event Bus            | Go Channels → NATS  | ✅ Approved |
| SQL Layer            | sqlc + database/sql | ✅ Approved |
| Configuration        | koanf               | ✅ Approved |
| Logging              | slog                | ✅ Approved |
| CLI                  | Cobra               | ✅ Approved |
| Metrics              | Prometheus          | ✅ Approved |
| Tracing              | OpenTelemetry       | ✅ Approved |
| CI/CD                | GitHub Actions      | ✅ Approved |
| Releases             | GoReleaser          | ✅ Approved |
