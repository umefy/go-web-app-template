# Architecture Decision Records (ADRs)

This directory contains Architecture Decision Records (ADRs) for the Go Web App Template project. ADRs document important architectural decisions made during the project's development.

## What are ADRs?

Architecture Decision Records are documents that capture important architectural decisions made during a project's development. They serve as a historical record of why certain technical choices were made.

## ADR Index

| ADR                                                     | Title                              | Status   | Date       |
| ------------------------------------------------------- | ---------------------------------- | -------- | ---------- |
| [ADR-001](./0001-clean-architecture-structure.md)       | Clean Architecture Implementation  | Accepted | 2025-01-21 |
| [ADR-002](./0002-repository-pattern-implementation.md)  | Repository Pattern Implementation  | Accepted | 2025-01-21 |
| [ADR-003](./0003-dependency-injection-with-wire.md)     | Dependency Injection with Wire     | Accepted | 2025-01-21 |
| [ADR-004](./0004-domain-driven-handler-organization.md) | Domain-Driven Handler Organization | Accepted | 2025-01-21 |
| [ADR-005](./0005-transaction-handling-strategy.md)      | Transaction Handling Strategy      | Accepted | 2025-01-21 |

## ADR Template

When creating new ADRs, use the following template:

```markdown
# ADR-XXX: [Title]

## Status

[Proposed | Accepted | Rejected | Deprecated | Superseded]

## Context

[Describe the context and problem statement that led to this decision]

## Decision

[Describe the decision that was made]

## Consequences

### Positive

- ✅ [Positive consequence 1]
- ✅ [Positive consequence 2]

### Negative

- ⚠️ [Negative consequence 1]
- ⚠️ [Negative consequence 2]

## Alternatives Considered

- [Alternative 1]: [Why it was rejected]
- [Alternative 2]: [Why it was rejected]

## Related Decisions

- [Related ADR-XXX]: [Description]
```

## Benefits of ADRs

1. **Historical Context**: Future developers understand why decisions were made
2. **Knowledge Transfer**: New team members can quickly understand the architecture
3. **Decision Tracking**: Avoid repeating past mistakes or re-discussing settled decisions
4. **Documentation**: Living documentation that evolves with the project

## When to Create an ADR

Create an ADR when making decisions about:

- Architecture patterns
- Technology choices
- Design patterns
- Infrastructure decisions
- Major refactoring decisions
- Performance optimizations
- Security implementations

## References

- [ADR GitHub Repository](https://github.com/joelparkerhenderson/architecture_decision_record)
- [ADR Wikipedia](https://en.wikipedia.org/wiki/Architecture_decision_record)
- [ADR by Michael Nygard](http://thinkrelevance.com/blog/2011/11/15/documenting-architecture-decisions)
