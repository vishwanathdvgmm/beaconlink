# RFC-0018: Distributed Scheduler

- **Status:** Draft
- **Author:** BeaconLink Engineering Team
- **Created:** 2026-07-19
- **Target Release:** Future
- **Discussion:** TBD

---

# Summary

This RFC proposes introducing a **Distributed Scheduler** that enables
BeaconLink to coordinate scheduled operations across multiple sites, relays,
agents, and runtime environments.

The scheduler provides reliable execution of time-based and
event-triggered operations while ensuring high availability,
fault tolerance, and consistent execution across distributed BeaconLink
deployments.

The objective is to centralize scheduling capabilities without coupling
scheduled execution to any individual platform component.

---

# Motivation

Infrastructure platforms frequently require coordinated scheduled
operations, including:

- Deployment windows
- Certificate renewal
- Secret rotation
- Backup operations
- Health verification
- Maintenance tasks
- Compliance scans
- Data synchronization
- Cleanup operations
- Periodic reconciliation

Embedding scheduling logic within individual services creates several
problems:

- Duplicate scheduling implementations
- Inconsistent execution
- Difficult failover
- Poor observability
- Resource contention
- Operational complexity

A centralized scheduler provides consistent scheduling behavior across
the entire BeaconLink platform.

---

# Goals

The proposed scheduler should:

- Support distributed execution
- Eliminate duplicate scheduling logic
- Ensure reliable execution
- Support recurring schedules
- Support one-time schedules
- Integrate with workflows
- Remain highly available
- Preserve platform modularity

---

# Non-Goals

This RFC does **not** propose:

- Replacing the Workflow Engine
- Becoming a CI/CD platform
- Executing arbitrary user scripts
- Replacing Kubernetes CronJobs
- Acting as a distributed database
- Managing application scheduling
- Performing business process automation
- Replacing external schedulers

---

# Proposed Design

BeaconLink introduces a centralized Distributed Scheduler responsible for
coordinating scheduled work across the platform.

```
              BeaconLink Console / API
                      │
              Scheduled Request
                      │
                      ▼
           Distributed Scheduler
                      │
        Schedule Persistence Store
                      │
      ┌───────────────┼───────────────┐
      ▼               ▼               ▼
 Workflow        Deployment      Maintenance
  Engine            Engine          Services
      │               │               │
      └───────────────┼───────────────┘
                      ▼
               Managed Resources
```

The scheduler determines **when** work executes.

Individual platform services determine **how** work is performed.

---

# Scheduling Model

Supported scheduling types may include:

- One-time execution
- Fixed intervals
- Cron expressions
- Maintenance windows
- Delayed execution
- Recurring schedules
- Event-triggered scheduling
- Policy-controlled scheduling

Schedules should be declarative and versioned.

---

# Scheduler Lifecycle

Each scheduled task progresses through:

1. Creation
2. Validation
3. Policy evaluation
4. Persistence
5. Scheduling
6. Execution
7. Completion
8. Audit recording
9. Retirement

Execution state should survive temporary service interruptions.

---

# Distributed Coordination

The scheduler should support:

- Leader election
- Failover
- Task ownership
- Cluster coordination
- Execution deduplication
- Load balancing
- High availability
- Recovery after failure

Only one scheduler instance should execute an individual scheduled task
at a given time.

---

# Execution Model

When a schedule becomes eligible:

1. Validate schedule
2. Verify policy
3. Acquire execution ownership
4. Trigger target service
5. Monitor execution
6. Record results
7. Publish events

Execution responsibilities remain delegated to existing BeaconLink services.

---

# Integration Points

The Distributed Scheduler integrates with:

- Workflow Engine
- Deployment Engine
- Policy Engine
- Secret Management
- Observability Platform
- Event Streaming Platform
- Autonomous Remediation
- Notification services

The scheduler remains independent of service implementation details.

---

# Security Considerations

Scheduled execution must preserve BeaconLink Zero Trust principles.

The scheduler should:

- Authenticate schedule creation
- Authorize scheduled operations
- Audit executions
- Validate execution ownership
- Protect schedule integrity
- Enforce policy evaluation
- Prevent duplicate execution
- Respect tenant isolation

Scheduled tasks should execute using delegated least-privilege
credentials.

---

# Operational Impact

Potential benefits include:

- Centralized scheduling
- Reliable execution
- High availability
- Simplified operations
- Improved observability
- Better governance
- Consistent scheduling behavior

Potential challenges include:

- Cluster coordination
- Leader election
- Clock synchronization
- Schedule persistence
- Failure recovery
- Capacity planning

---

# Compatibility

The Distributed Scheduler is optional.

Existing BeaconLink deployments continue operating without centralized
scheduling.

Platform components may gradually migrate existing scheduled operations
to the shared scheduler.

---

# Alternatives Considered

## Service-Level Scheduling

Allow each service to manage its own schedules.

Pros:

- Independent implementation
- Simple architecture

Cons:

- Duplicate logic
- Inconsistent behavior
- Difficult coordination
- Increased maintenance

---

## External Scheduler

Delegate scheduling entirely to external platforms.

Pros:

- Mature tooling
- Existing infrastructure

Cons:

- Reduced BeaconLink awareness
- Limited policy integration
- Fragmented observability
- Additional operational dependency

---

## Operating System Cron

Use operating system schedulers for background tasks.

Pros:

- Simple implementation
- Widely available

Cons:

- No distributed coordination
- No failover
- Poor visibility
- Difficult multi-node management

---

# Open Questions

- Which scheduling syntax should BeaconLink standardize on?
- How should missed executions be handled after outages?
- Should schedules support dependency ordering?
- How should maintenance windows interact with scheduling?
- Should schedules be tenant-scoped?
- How should scheduler clusters perform leader election?

---

# Related Documents

- ADR-015 Event-Driven Internal Architecture
- ADR-016 Declarative Desired State
- ADR-018 Observability Strategy
- RFC-0009 Policy Engine
- RFC-0014 Workflow Engine
- RFC-0015 Event Streaming Platform
- RFC-0019 Cross-Site Disaster Recovery

---

# Recommendation

BeaconLink should introduce a **Distributed Scheduler** that provides a
centralized, highly available scheduling service for operational tasks
while preserving clear separation between scheduling, orchestration, and
execution.

The preferred architecture treats scheduling as shared platform
infrastructure that integrates with the Workflow Engine, Policy Engine,
and Event Streaming Platform while delegating actual work to existing
BeaconLink services.

Initial implementation should prioritize recurring maintenance,
deployment windows, certificate and secret rotation, periodic
reconciliation, and compliance operations, establishing a reliable
foundation for coordinated automation across distributed BeaconLink
deployments.
