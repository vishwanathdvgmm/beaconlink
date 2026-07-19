# 128 - Go Monorepo

- **Status:** Approved
- **Version:** 1.0
- **Last Updated:** 2026-07-20

---

# 1. Purpose

This document defines the official repository organization for the BeaconLink codebase.

BeaconLink is implemented as a **single Go monorepo**, containing all platform components, shared libraries, documentation, SDKs, tooling, and deployment assets.

This document establishes the repository layout, package organization, dependency boundaries, and engineering guidelines for contributors.

---

# 2. Goals

The repository structure should:

- Be easy to navigate.
- Scale to millions of lines of code.
- Minimize coupling.
- Encourage modular design.
- Support independent executables.
- Enable code reuse without circular dependencies.
- Keep implementation details private.
- Provide a predictable developer experience.

---

# 3. Repository Strategy

BeaconLink uses a **single Git repository**.

All major components evolve together under a single version history.

Benefits include:

- Unified versioning
- Shared Go modules
- Easier refactoring
- Simplified dependency management
- Single CI/CD pipeline
- Shared engineering standards
- Atomic changes across components

---

# 4. Repository Layout

```text
BeaconLink/
├── 📁 api
├── 📁 cmd
│   ├── 📁 agent
│   ├── 📁 beacon
│   ├── 📁 installer
│   └── 📁 relay
├── 📁 configs
├── 📁 deployments
├── 📁 docs
├── 📁 examples
├── 📁 internal
├── 📁 pkg
├── 📁 scripts
├── 📁 sdk
├── 📁 test
├── 📁 web
├── ⚙️ .gitignore
├── 📄 LICENSE
├── 📝 README.md
├── 📄 go.mod
└── 📄 go.sum
```

---

# 5. Root Directory Responsibilities

| Directory   | Purpose                         |
| ----------- | ------------------------------- |
| cmd         | Application entry points        |
| internal    | Private implementation packages |
| pkg         | Public reusable libraries       |
| web         | React frontend                  |
| sdk         | Generated SDKs                  |
| api         | API specifications              |
| docs        | Documentation                   |
| configs     | Default configuration           |
| deployments | Deployment manifests            |
| examples    | Example projects                |
| scripts     | Development scripts             |
| test        | Integration and E2E tests       |

---

# 6. Executables (`cmd`)

Every executable lives inside the `cmd` directory.

Example:

```text
cmd/

agent/
relay/
api/
beacon/
installer/
```

Each executable contains only application bootstrap code.

Typical contents:

```go
func main() {
    app := relay.New()
    app.Run()
}
```

Business logic must never live inside `cmd`.

---

# 7. Internal Packages (`internal`)

The `internal` directory contains implementation packages that are private to BeaconLink.

External applications must not import packages from `internal`.

Example layout:

```text
internal/

agent/
relay/
api/

auth/
config/
database/
deployment/
events/
inventory/
logging/
metrics/
policy/
runtime/
scheduler/
security/
telemetry/
workflow/
version/
```

Packages should be organized by **business capability**, not technical layer.

---

# 8. Public Packages (`pkg`)

Packages inside `pkg` are intended for external consumption.

Examples:

```text
pkg/

client/
protocol/
sdk/
types/
version/
```

Only stable APIs belong here.

Internal implementation should never leak into `pkg`.

---

# 9. Web Application

The web frontend is implemented independently.

```text
web/

src/
public/
package.json
vite.config.ts
```

Technology stack:

- React
- TypeScript
- Vite
- Tailwind CSS

The web application communicates exclusively through documented APIs.

---

# 10. SDK Directory

The SDK directory contains generated client libraries.

Example:

```text
sdk/

go/
typescript/
python/
```

SDKs are generated from API specifications whenever possible.

Manual implementations should be minimized.

---

# 11. API Directory

The `api` directory contains API contracts.

Examples include:

- OpenAPI
- JSON Schemas
- Protocol definitions

Implementation code does not belong here.

---

# 12. Configuration

```text
configs/
```

Contains:

- Development configuration
- Sample configuration
- Default settings

Secrets must never be committed.

---

# 13. Deployment Assets

```text
deployments/
```

Contains deployment artifacts including:

- Docker
- Docker Compose
- systemd
- Kubernetes (future)

Deployment manifests remain separate from application code.

---

# 14. Examples

```text
examples/
```

Contains:

- Sample applications
- SDK examples
- API examples
- Configuration examples

Examples serve as living documentation.

---

# 15. Scripts

```text
scripts/
```

Contains helper scripts for:

- Development
- CI
- Releases
- Code generation

Scripts should remain platform-independent whenever practical.

---

# 16. Test Directory

```text
test/
```

Contains:

- Integration tests
- End-to-end tests
- Test fixtures
- Test data

Unit tests should remain alongside source files.

---

# 17. Go Module Strategy

BeaconLink uses **one Go module**.

```text
go.mod
```

located at the repository root.

Multiple modules are intentionally avoided.

Advantages:

- Easier dependency management
- Simpler refactoring
- Single dependency graph
- Unified tooling

---

# 18. Package Design Principles

Packages should:

- Have a single responsibility.
- Be cohesive.
- Minimize exported symbols.
- Hide implementation details.
- Avoid cyclic dependencies.

Smaller packages are preferred over large utility packages.

---

# 19. Domain-Oriented Organization

Packages should be organized around business capabilities.

Preferred:

```text
deployment/
workflow/
runtime/
inventory/
policy/
scheduler/
relay/
agent/
```

Avoid organizing by technical layers.

Avoid:

```text
controllers/
repositories/
services/
helpers/
utils/
```

---

# 20. Shared Code

Shared functionality belongs in dedicated packages.

Examples include:

- configuration
- logging
- authentication
- metrics
- protocol
- versioning

Shared code should have well-defined ownership.

---

# 21. Dependency Rules

Dependencies must flow inward.

```text
cmd
 ↓
internal
 ↓
pkg
```

Never reverse this flow.

---

## Allowed

```text
cmd
  ↓
internal
```

```text
internal
  ↓
pkg
```

```text
cmd
  ↓
pkg
```

---

## Forbidden

```text
pkg
 ↓
internal
```

```text
internal
 ↓
cmd
```

```text
cmd
 ↓
cmd
```

Circular dependencies are strictly prohibited.

---

# 22. Import Guidelines

Prefer explicit imports.

Avoid large "utility" packages.

Dependencies should remain obvious from package names.

Imports should reflect architectural boundaries.

---

# 23. Naming Conventions

Directories:

- lowercase
- singular
- descriptive

Examples:

```text
agent
relay
runtime
workflow
policy
config
```

Avoid:

```text
Helpers
CommonStuff
Utilities
Misc
```

---

# 24. File Organization

Files should remain focused.

Recommended size:

- 100–400 lines

Very large files should be decomposed.

---

# 25. Generated Code

Generated code should be clearly separated.

Examples:

- sqlc output
- SDK generation
- API generation

Generated files must never be manually modified.

---

# 26. Future Expansion

The repository is designed to accommodate future additions including:

- Plugin system
- Marketplace
- Kubernetes operator
- Additional SDKs
- Distributed scheduler
- AI services

New top-level directories require architectural review.

---

# 27. Repository Principles

The BeaconLink repository should remain:

- Predictable
- Modular
- Discoverable
- Scalable
- Easy to maintain
- Friendly to new contributors

Every directory should have a clearly defined responsibility.

---

# 28. Repository Governance

Changes to repository organization should be infrequent.

New top-level directories require justification.

Significant structural changes should be documented through an Architecture Decision Record (ADR).

Maintaining a stable repository structure is essential for long-term maintainability and contributor productivity.
