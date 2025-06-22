# ADR-005: Transaction Handling Strategy

## Status

Accepted

## Context

We needed to decide how to handle database transactions in our application. The key questions were:

1. How to pass transactions from middleware to services and repositories
2. How to support both transactional and non-transactional operations
3. How to make transaction handling explicit and testable

## Decision

We chose **Explicit Transaction Passing** with the following approach:

### Repository Interface with Transaction Support

```go
// internal/domain/user/repository/repository.go
type Repository interface {
    // Regular operations (non-transactional)
    GetUser(ctx context.Context, id int) (*model.User, error)
    GetUsers(ctx context.Context) ([]*model.User, error)

    // Transactional operations
    GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*model.User, error)
    GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*model.User, error)
    CreateUser(ctx context.Context, user *model.User, tx *query.QueryTx) (*model.User, error)
    UpdateUser(ctx context.Context, id int, user *model.User, tx *query.QueryTx) (*model.User, error)
}
```

### Service Layer with Transaction Parameters

```go
// internal/domain/user/service/service.go
type Service interface {
    GetUsers(ctx context.Context) ([]*model.User, error)
    GetUser(ctx context.Context, id string) (*model.User, error)
    CreateUser(ctx context.Context, user *model.User, tx *query.QueryTx) (*model.User, error)
    UpdateUser(ctx context.Context, id string, user *model.User, tx *query.QueryTx) (*model.User, error)
}
```

### Middleware for Transaction Management

```go
// internal/infrastructure/http/middleware/transaction.go
func Transaction(dbQuery *query.Query, loggerService loggerSrv.Service) middleware.Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            tx := dbQuery.Begin()
            ctx := context.WithValue(r.Context(), TransactionCtxKey, tx)

            next.ServeHTTP(w, r.WithContext(ctx))

            if r.Context().Err() != nil {
                tx.Rollback()
            } else {
                tx.Commit()
            }
        })
    }
}
```

### Handler Implementation

```go
// internal/infrastructure/http/handler/user/update_user.go
func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
    ctx := r.Context()

    // Get transaction from middleware
    tx, err := middleware.GetTransaction(ctx)
    if err != nil {
        return err
    }

    // Pass transaction to service
    user, err = h.userService.UpdateUser(ctx, userID, user, tx)
    if err != nil {
        return err
    }

    // ... rest of handler
}
```

## Consequences

### Positive

- ✅ **Explicit Dependencies**: Transaction requirement is clear in interfaces
- ✅ **Testability**: Easy to mock transactions in tests
- ✅ **Flexibility**: Services can be used with or without transactions
- ✅ **Clean Architecture**: Service doesn't know about middleware implementation
- ✅ **Type Safety**: Compile-time checking of transaction usage
- ✅ **Clear Separation**: Regular vs transactional operations are distinct

### Negative

- ⚠️ **More Parameters**: Service methods need transaction parameters
- ⚠️ **Boilerplate**: Need to pass transactions through multiple layers
- ⚠️ **Complexity**: More complex method signatures

## Alternatives Considered

### Alternative 1: Context-Based Transaction Access

```go
// Service gets transaction from context
func (s *userService) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
    tx := ctx.Value("transaction").(*query.QueryTx)
    return s.userRepository.UpdateUser(ctx, id, user, tx)
}
```

**Rejected**: Service depends on middleware implementation details, hard to test.

### Alternative 2: Transaction Manager

```go
// Centralized transaction management
type TransactionManager interface {
    Begin(ctx context.Context) (Transaction, error)
    Commit(tx Transaction) error
    Rollback(tx Transaction) error
}
```

**Rejected**: Adds unnecessary abstraction layer, more complex than needed.

### Alternative 3: Repository-Level Transaction Management

```go
// Repository manages its own transactions
func (r *userRepository) UpdateUserWithTx(ctx context.Context, id int, user *model.User) (*model.User, error) {
    tx := r.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // ... implementation
    return user, tx.Commit()
}
```

**Rejected**: Hard to coordinate transactions across multiple repositories.

## Related Decisions

- ADR-001: Clean Architecture Implementation
- ADR-002: Repository Pattern Implementation
- ADR-004: Domain-Driven Handler Organization
