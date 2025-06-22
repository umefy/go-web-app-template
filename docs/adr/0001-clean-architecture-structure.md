# ADR-001: Clean Architecture Implementation

## Status

Accepted

## Context

We needed to choose an architecture pattern for our Go web application template that would provide:

- Clear separation of concerns
- Testability
- Maintainability
- Scalability
- Flexibility for different infrastructure choices

## Decision

We chose Clean Architecture with the following structure:

```bash
internal/
├── domain/                    # Pure business logic (no external dependencies)
│   ├── user/                 # User domain
│   │   ├── repository/       # Repository interfaces only
│   │   ├── service/          # Business logic services
│   │   └── error/            # Domain-specific errors
│   ├── config/               # Configuration domain
│   ├── greeter/              # Greeter domain
│   └── logger/               # Logger domain
├── infrastructure/            # External concerns & implementations
│   ├── database/             # Database implementations
│   ├── http/                 # HTTP infrastructure
│   ├── grpc/                 # gRPC infrastructure
│   ├── config/               # Configuration loading
│   └── logger/               # Logger setup
└── app/                      # Application composition & DI
    ├── app.go                # Main application struct
    ├── wire.go               # Dependency injection
    └── wire_gen.go           # Generated wire code
```

## Consequences

### Positive

- ✅ **Clear separation of concerns**: Domain logic is isolated from infrastructure
- ✅ **Testability**: Business logic can be tested without external dependencies
- ✅ **Flexibility**: Easy to swap infrastructure implementations (e.g., different databases)
- ✅ **Maintainability**: Related code is co-located and easy to find
- ✅ **Scalability**: Easy to add new domains without affecting existing code
- ✅ **Dependency direction**: Domain doesn't know about infrastructure

### Negative

- ⚠️ **More files**: More complex file structure
- ⚠️ **Learning curve**: New developers need to understand the architecture
- ⚠️ **Boilerplate**: More code required for simple operations
- ⚠️ **Overhead**: May be overkill for very small applications

## Related Decisions

- ADR-002: Repository Pattern Implementation
- ADR-003: Dependency Injection with Wire
- ADR-004: Domain-Driven Handler Organization
