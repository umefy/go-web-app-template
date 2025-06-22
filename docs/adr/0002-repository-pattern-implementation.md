# ADR-002: Repository Pattern Implementation

## Status

Accepted

## Context

We needed to decide how to handle data access in our application. The options were:

1. Direct database access in services
2. Repository pattern with interfaces in domain and implementations in infrastructure
3. Repository pattern with both interfaces and implementations in domain

## Decision

We chose **Repository Pattern with interfaces in domain and implementations in infrastructure**:

### Domain Layer (Interfaces Only)

```go
// internal/domain/user/repository/repository.go
type Repository interface {
    GetUser(ctx context.Context, id int) (*model.User, error)
    GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*model.User, error)
    GetUsers(ctx context.Context) ([]*model.User, error)
    GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*model.User, error)
    CreateUser(ctx context.Context, user *model.User, tx *query.QueryTx) (*model.User, error)
    UpdateUser(ctx context.Context, id int, user *model.User, tx *query.QueryTx) (*model.User, error)
}
```

### Infrastructure Layer (Implementations)

```go
// internal/infrastructure/database/user_repository.go
type userRepository struct {
    loggerService loggerSrv.Service
    dbQuery       *query.Query
}

func NewUserRepository(dbQuery *query.Query, loggerService loggerSrv.Service) userRepo.Repository {
    return &userRepository{loggerService: loggerService, dbQuery: dbQuery}
}
```

### Wire Binding

```go
// internal/app/wire.go
wire.Bind(new(userRepo.Repository), new(*database.UserRepository))
```

## Consequences

### Positive

- ✅ **True Clean Architecture**: Domain stays pure without infrastructure dependencies
- ✅ **Easy Testing**: Can mock repository interfaces for unit tests
- ✅ **Flexibility**: Easy to swap database implementations (GORM, SQLx, etc.)
- ✅ **Transaction Support**: Clear transaction handling with separate Tx methods
- ✅ **Dependency Inversion**: Domain depends on abstractions, not concretions

### Negative

- ⚠️ **More Files**: Separate interface and implementation files
- ⚠️ **Complexity**: More complex dependency injection setup
- ⚠️ **Boilerplate**: Need to implement both regular and transaction methods

## Alternatives Considered

### Alternative 1: Repository in Domain

```bash
internal/domain/user/repository/
├── repository.go          # Interface
├── repository_impl.go     # Implementation (same package)
└── wire.go
```

**Rejected**: Violates clean architecture principles, domain knows about database details.

### Alternative 2: Direct Database Access

```go
// Service directly uses database
func (s *userService) GetUser(ctx context.Context, id string) (*model.User, error) {
    return s.db.User.WithContext(ctx).Where(s.db.User.ID.Eq(id)).First()
}
```

**Rejected**: Hard to test, violates dependency inversion principle.

## Related Decisions

- ADR-001: Clean Architecture Implementation
- ADR-003: Dependency Injection with Wire
- ADR-005: Transaction Handling Strategy
