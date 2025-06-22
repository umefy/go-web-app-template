# ADR-003: Dependency Injection with Wire

## Status

Accepted

## Context

We needed to choose a dependency injection solution for our Go application. The options were:

1. Manual dependency injection
2. Wire (Google's dependency injection tool)
3. Dig (Uber's dependency injection framework)
4. fx (Uber's application framework)

## Decision

We chose **Google Wire** for dependency injection with the following structure:

### Wire Configuration

```go
// internal/app/wire.go
var WireSet = wire.NewSet(
    logger.NewLogger,
    database.WireSet,
    wire.Bind(new(loggerSvc.Logger), new(*dashLogger.Logger)),
    config.LoadConfig,
    configSvc.WireSet,
    userSvc.WireSet,
    greeterSvc.WireSet,
    loggerSvc.WireSet,
    wire.Bind(new(userRepo.Repository), new(*database.UserRepository)),
    wire.Struct(new(App), "*"),
)
```

### Domain Wire Sets

```go
// internal/domain/user/service/wire.go
var WireSet = wire.NewSet(
    NewService,
)

// internal/infrastructure/database/wire.go
var WireSet = wire.NewSet(
    NewDB,
    NewDBQuery,
    NewUserRepository,
)
```

### Generated Code

```go
// internal/app/wire_gen.go (auto-generated)
func InitializeApp(configOptions config.Options) (*App, error) {
    appConfig, err := config.LoadConfig(configOptions)
    // ... dependency resolution
    userRepository := database.NewUserRepository(query, loggerService)
    userService := service3.NewService(loggerService, userRepository)
    // ... rest of initialization
}
```

## Consequences

### Positive

- ✅ **Compile-time Safety**: Wire validates dependencies at compile time
- ✅ **Code Generation**: No runtime reflection, better performance
- ✅ **Clear Dependencies**: Explicit dependency graph
- ✅ **Modular**: Each layer has its own wire set
- ✅ **Type Safety**: Go's type system ensures correct bindings
- ✅ **Google Backed**: Well-maintained and stable

### Negative

- ⚠️ **Learning Curve**: Developers need to understand Wire concepts
- ⚠️ **Code Generation**: Need to regenerate after changes
- ⚠️ **Debugging**: Generated code can be harder to debug
- ⚠️ **Complexity**: More complex than manual DI for simple cases

## Alternatives Considered

### Alternative 1: Manual Dependency Injection

```go
func NewApp() *App {
    config := config.New()
    logger := logger.New(config)
    db := database.New(config)
    userRepo := user.NewRepository(db, logger)
    userService := user.NewService(userRepo, logger)
    return &App{userService: userService}
}
```

**Rejected**: Becomes unwieldy as application grows, hard to maintain.

### Alternative 2: Dig (Uber)

```go
container := dig.New()
container.Provide(user.NewService)
container.Provide(user.NewRepository)
```

**Rejected**: Runtime reflection, less type safety, more complex error handling.

### Alternative 3: fx (Uber)

```go
app := fx.New(
    fx.Provide(user.NewService),
    fx.Provide(user.NewRepository),
)
```

**Rejected**: More opinionated, runtime reflection, overkill for our needs.

## Related Decisions

- ADR-001: Clean Architecture Implementation
- ADR-002: Repository Pattern Implementation
- ADR-004: Domain-Driven Handler Organization
