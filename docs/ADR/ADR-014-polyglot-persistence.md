# ADR-014: Polyglot Persistence

- **Status:** Accepted
- **Date:** 2026-07-18
- **Decision Makers:** BeaconLink Engineering Team
- **Technical Story:** Data Persistence Architecture

---

## Context

BeaconLink stores multiple categories of data with significantly different
characteristics, including:

- Configuration
- Site manifests
- Inventory
- Deployment metadata
- Audit logs
- Metrics
- Events
- Caches
- Operational state

These workloads differ in:

- Access patterns
- Query complexity
- Consistency requirements
- Retention policies
- Scalability characteristics
- Performance expectations

A single database technology is unlikely to provide an optimal solution
for every workload without introducing compromises in performance,
maintainability, or operational efficiency.

The platform therefore requires a persistence strategy that allows
storage technologies to be selected according to workload
characteristics while preserving a consistent application architecture.

---

## Decision

BeaconLink adopts a **Polyglot Persistence** strategy.

Different persistence technologies may be used for different categories
of data based on their functional and operational requirements.

Application services access data exclusively through repository
interfaces and domain abstractions.

Business logic remains independent of the underlying storage
implementation.

Storage technologies may evolve independently without affecting domain
services or application workflows.

---

## Alternatives Considered

### Single Relational Database

Store all platform data within one relational database.

**Rejected because:**

- Different workloads have conflicting requirements
- Time-series data performs poorly
- Event storage becomes inefficient
- Cache workloads are unsuitable
- Scalability constraints

---

### Single NoSQL Database

Use one document or key-value database for all persistence.

**Rejected because:**

- Weak relational capabilities
- Complex transactional workflows
- Reduced query flexibility
- Reporting limitations
- Less suitable for structured configuration

---

### Runtime-Selected Storage

Allow application modules to choose storage technologies directly.

**Rejected because:**

- Tight application coupling
- Difficult testing
- Reduced portability
- Inconsistent persistence patterns
- Higher maintenance cost

---

## Consequences

### Positive

- Storage optimized per workload
- Improved scalability
- Better performance
- Technology flexibility
- Simplified future migrations
- Reduced vendor lock-in
- Cleaner architecture
- Independent storage evolution

### Negative

- Multiple storage technologies to operate
- More complex backup strategy
- Additional operational expertise
- Repository abstraction required
- Cross-storage transactions require careful design

---

## Architecture

```
            Domain Services
                   │
                   ▼
         Repository Interfaces
                   │
     ┌─────────────┼─────────────┐
     ▼             ▼             ▼
Relational     Document     Time-Series
 Storage         Storage       Storage
     │             │             │
     └─────────────┼─────────────┘
                   ▼
           Persistent Data
```

Application services depend only on repository interfaces, while
storage implementations remain interchangeable.

---

## Related Decisions

- ADR-006 Site Manifest
- ADR-013 Monorepo Strategy
- ADR-015 Event-Driven Internal Architecture
- ADR-016 Declarative Desired State
- ADR-020 Immutable Deployments

---

## Implementation Notes

The persistence architecture should:

- Separate business logic from persistence
- Select storage based on workload characteristics
- Use repository abstractions
- Support schema evolution
- Enable independent storage migration
- Implement consistent backup policies
- Expose storage health metrics
- Maintain clear ownership of stored data

Specific database products, storage engines, indexing strategies, and
replication mechanisms remain implementation decisions and are outside
the scope of this ADR.

---

## Rationale

Polyglot persistence allows BeaconLink to optimize storage for each category
of platform data without coupling business logic to specific database
technologies.

By introducing repository abstractions between domain services and
storage implementations, BeaconLink maintains a clean architectural
separation while allowing persistence technologies to evolve over time.
This approach improves scalability, performance, and operational
flexibility while reducing vendor lock-in and supporting future
technology adoption without widespread application changes.
