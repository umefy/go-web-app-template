# ADR-004: Domain-Driven Handler Organization

## Status

Accepted

## Context

We needed to organize HTTP handlers in a way that would:

- Scale well as the application grows
- Keep related endpoints together
- Make it easy to find and maintain specific functionality
- Support multiple domains (user, order, greeter, etc.)

## Decision

We chose **Domain-Driven Handler Organization** with the following structure:

```bash
internal/infrastructure/http/
├── server.go              # HTTP server initialization
├── router.go              # Main router setup
├── middleware/            # Shared middlewares
│   └── transaction.go     # Transaction middleware
└── handler/               # Domain-specific handlers
    ├── user/              # User domain handlers
    │   ├── handler.go     # UserHandler struct & methods
    │   ├── create_user.go # POST /users
    │   ├── get_user.go    # GET /users/{id}
    │   ├── get_users.go   # GET /users
    │   ├── update_user.go # PUT /users/{id}
    │   └── mapping/       # Request/response mapping
    │       ├── user_mapping.go
    │       └── order_mapping.go
    ├── default_handler.go # Base handler interface
    └── handler.go         # Common handler utilities
```

### Handler Implementation

```go
// internal/infrastructure/http/handler/user/handler.go
type Handler interface {
    handler.Handler
    GetUsers(w http.ResponseWriter, r *http.Request) error
    GetUser(w http.ResponseWriter, r *http.Request) error
    CreateUser(w http.ResponseWriter, r *http.Request) error
    UpdateUser(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
    *handler.DefaultHandler
    userService   userSrv.Service
    loggerService loggerSrv.Service
}
```

### Router Setup

```go
// internal/infrastructure/http/router.go
func NewUserRouter(app *app.App) http.Handler {
    r := router.NewRouter()
    h := userHandler.NewHandler(app.UserService, app.LoggerService)

    r.Get("/", h.Handle(h.GetUsers))
    r.Get("/{id}", h.Handle(h.GetUser))
    r.Post("/", h.Handle(h.CreateUser))
    r.Patch("/{id}", h.Handle(h.ApplyMiddlewares(
        h.UpdateUser,
        middleware.Transaction(app.DbQuery, app.LoggerService),
    )))
    return r
}
```

## Consequences

### Positive

- ✅ **Scalability**: Easy to add new domains without cluttering existing code
- ✅ **Maintainability**: Related endpoints are grouped together
- ✅ **Clear API Structure**: URL structure mirrors the folder structure
- ✅ **Domain Separation**: Each domain has its own handler package
- ✅ **Middleware Integration**: Easy to apply domain-specific middleware
- ✅ **Testability**: Can test each domain's handlers independently

### Negative

- ⚠️ **More Files**: More complex file structure
- ⚠️ **Boilerplate**: Need to create handler struct for each domain
- ⚠️ **Learning Curve**: New developers need to understand the pattern

## Alternatives Considered

### Alternative 1: Single Handler File

```bash
internal/infrastructure/http/handler/
├── user_handler.go
├── order_handler.go
└── greeter_handler.go
```

**Rejected**: Becomes unwieldy as handlers grow, hard to find specific endpoints.

### Alternative 2: Function-Based Handlers

```go
// No struct, just functions
func GetUsers(w http.ResponseWriter, r *http.Request) error {
    // handler logic
}
```

**Rejected**: Hard to share dependencies, no common error handling.

### Alternative 3: RESTful Resource Organization

```bash
internal/infrastructure/http/handler/
├── users/
│   ├── create.go
│   ├── read.go
│   ├── update.go
│   └── delete.go
```

**Rejected**: Too granular, doesn't align with domain boundaries.

## Related Decisions

- ADR-001: Clean Architecture Implementation
- ADR-003: Dependency Injection with Wire
- ADR-005: Transaction Handling Strategy
