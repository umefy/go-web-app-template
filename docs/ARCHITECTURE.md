# Architecture Guide

This document provides a comprehensive overview of the architecture principles, design decisions, and technical implementation details.

## 🏗️ Clean Architecture Principles

This project follows **Clean Architecture** principles with a clear separation of concerns and uses **Uber FX** for dependency injection. Each major component has its own `fx.go` file that defines its dependencies and provides its services to the application.

### Architecture Layers

- **Domain Layer**: Pure business logic with interfaces and models only
- **Service Layer**: Business logic implementation using domain interfaces
- **Delivery Layer**: Transport concerns (HTTP, GraphQL, gRPC)
- **Infrastructure Layer**: External implementations (database, server, logger, tracing)
- **Core Layer**: Shared core components (configuration)
- **Entry Point**: `cmd/server/main.go` - FX module composition and dependency injection

### Dependency Direction

```
Domain ← Service ← Delivery ← Infrastructure
```

The domain layer doesn't know about external concerns, ensuring testability and maintainability.

## 🔧 FX Dependency Injection Architecture

The project uses [Uber FX](https://github.com/uber-go/fx) for dependency injection, providing a clean and modular approach to managing application dependencies.

### Module Composition

The dependency injection is composed directly in `cmd/server/main.go`:

```go
app := fx.New(
    fx.Supply(args),
    fx.Provide(func() context.Context {
        return context.Background()
    }),
    config.Module,        // Configuration management
    database.Module,      // Database connections and repositories
    logger.Module,        // Logging infrastructure
    tracing.Module,       // OpenTelemetry tracing
    http.Module,          // HTTP server and REST/GraphQL routers
    grpc.Module,          // gRPC server and handlers
    service.Module,       // Business logic services
    fx.Invoke(start),     // Application startup
)
```

### Module Structure

Each layer provides its own FX module:

- **Config Module** (`internal/core/config/fx.go`): Configuration management
- **Database Module** (`internal/infrastructure/database/fx.go`): Database connections and repositories
- **Logger Module** (`internal/infrastructure/logger/fx.go`): Logging infrastructure
- **Tracing Module** (`internal/infrastructure/tracing/fx.go`): OpenTelemetry tracing
- **HTTP Server Module** (`internal/infrastructure/server/http/fx.go`): HTTP server and REST/GraphQL routers
- **gRPC Server Module** (`internal/infrastructure/server/grpc/fx.go`): gRPC server and handlers
- **Service Module** (`internal/service/fx.go`): Business logic services
- **GraphQL Module** (`internal/delivery/graphql/fx.go`): GraphQL resolvers and router
- **API V1 Module** (`internal/delivery/restful/openapi/v1/fx.go`): REST API handlers

### Benefits

- **Modular Design**: Each component is self-contained with its own dependencies
- **Lazy Loading**: Dependencies are only created when needed
- **Lifecycle Management**: Automatic startup and shutdown of services
- **Error Handling**: Graceful error handling during dependency resolution
- **Testing**: Easy to mock and test individual modules

## 📁 Project Structure

```
go-web-app-template/
├── cmd/                           # Application entry points
│   ├── server/                   # HTTP/gRPC server startup
│   ├── seed/                     # Database seeding utilities
│   │   └── database/             # Database seeding implementation
│   │       ├── main.go           # Main seeding entry point
│   │       ├── seed_users.go     # User seeding logic
│   │       └── seed_orders.go    # Order seeding logic
│   └── concurrent/               # Concurrent testing utilities
│       └── concurrent_user_update.go # Optimistic locking test
├── internal/                      # Private application code (Go-enforced privacy)
│   ├── domain/                    # Pure business logic & interfaces (no external dependencies)
│   │   ├── user/                 # User domain
│   │   │   ├── repo/             # Repository interfaces
│   │   │   ├── user.go           # Domain models with optimistic locking
│   │   │   ├── user_with_order.go # Domain models with relationships
│   │   │   └── error/            # Domain-specific errors including optimistic lock conflicts
│   │   ├── order/                # Order domain
│   │   │   ├── repo/             # Repository interfaces
│   │   │   └── order.go          # Domain models
│   │   └── error/                # Shared domain errors
│   ├── service/                   # Business logic implementation
│   │   ├── fx.go                 # Service FX module
│   │   ├── user/                 # User business logic
│   │   │   ├── service.go        # User service implementation with optimistic locking
│   │   │   └── input.go          # Service input/output models
│   │   ├── order/                # Order business logic
│   │   │   ├── service.go        # Order service implementation
│   │   └── greeter/              # Greeter business logic
│   │       ├── service.go        # Greeter service implementation
│   ├── delivery/                  # Transport layer (HTTP/gRPC/GraphQL)
│   │   ├── restful/              # HTTP REST API
│   │   │   ├── handler/          # Shared handler utilities
│   │   │   │   ├── handler.go    # Handler interface
│   │   │   │   ├── default_handler.go # Default handler with error handling
│   │   │   │   └── middleware/   # HTTP middleware
│   │   │   └── openapi/          # OpenAPI REST endpoints
│   │   │       └── v1/           # API version 1
│   │   │           ├── fx.go     # API V1 FX module
│   │   │           ├── router.go # OpenAPI router
│   │   │           └── user/     # User REST handlers
│   │   │               ├── handler.go      # User handler interface
│   │   │               ├── create_user.go  # Create user endpoint
│   │   │               ├── get_user.go     # Get user endpoint
│   │   │               ├── get_users.go    # Get users endpoint
│   │   │               ├── update_user.go  # Update user endpoint with optimistic locking
│   │   │               ├── router.go       # User routing
│   │   │               └── mapping/        # Data mapping
│   │   ├── graphql/              # GraphQL API
│   │   │   ├── fx.go             # GraphQL FX module
│   │   │   ├── generated.go      # gqlgen generated code
│   │   │   ├── User.resolvers.go # User GraphQL resolvers
│   │   │   ├── Order.resolvers.go # Order GraphQL resolvers
│   │   │   ├── router.go         # GraphQL router with playground
│   │   │   └── model/            # GraphQL models
│   │   ├── grpc/                 # gRPC API
│   │   │   ├── fx.go             # gRPC handler FX module
│   │   │   ├── handler/          # gRPC handlers
│   │   │   │   └── greeter/      # Greeter gRPC service
│   │   │   └── server.go         # gRPC server
│   │   └── errutil/              # Error handling utilities
│   ├── infrastructure/            # External concerns & implementations
│   │   ├── database/             # Database infrastructure
│   │   │   ├── fx.go             # Database FX module
│   │   │   ├── with_tx.go        # Transaction utilities
│   │   │   ├── ctx/              # Database context utilities
│   │   │   └── gorm/             # GORM database implementation
│   │   │       ├── setup.go      # Database connection setup
│   │   │       ├── with_tx.go    # GORM transaction utilities
│   │   │       ├── generated/    # GORM generated models with optimistic locking
│   │   │       └── repo/         # Repository implementations
│   │   │           ├── user_repo.go    # User repository with optimistic locking
│   │   │           ├── order_repo.go   # Order repository implementation
│   │   │           └── mapping/        # Database mapping utilities
│   │   ├── server/               # Server infrastructure
│   │   │   ├── fx.go             # Server FX module
│   │   │   ├── http/             # HTTP server setup
│   │   │   │   ├── fx.go         # HTTP server FX module
│   │   │   │   ├── router.go     # Main HTTP router
│   │   │   │   ├── server.go     # HTTP server
│   │   │   │   └── middleware/   # HTTP middleware
│   │   │   └── grpc/             # gRPC server setup
│   │   │       ├── fx.go         # gRPC server FX module
│   │   │       ├── handler/      # gRPC handlers
│   │   │       └── server.go     # gRPC server
│   │   ├── logger/               # Logger setup
│   │   │   └── fx.go             # Logger FX module
│   │   └── tracing/              # OpenTelemetry tracing setup
│   │       ├── opentelemetry/    # OpenTelemetry implementation
│   │       │   └── setup.go      # Tracing setup and configuration
│   │       └── fx.go             # Tracing FX module
│   ├── core/                      # Core shared components
│   │   └── config/               # Configuration management
│   │       ├── fx.go             # Config FX module
│   │       ├── config.go         # Main configuration struct
│   │       ├── setup.go          # Configuration loading
│   │       ├── app_config.go     # Application configuration
│   │       ├── db_config.go      # Database configuration
│   │       ├── http_server_config.go # HTTP server configuration
│   │       ├── grpc_server_config.go # gRPC server configuration
│   │       ├── logging_config.go # Logging configuration
│   │       └── tracing_config.go # Tracing configuration
│   └── cmd/                      # Application entry point
│       └── server/               # Main server application
│           └── main.go           # Server entry point with FX DI
├── pkg/                           # Public reusable packages
├── openapi/                       # OpenAPI specifications & generated code
│   ├── docs/                     # OpenAPI specification files
│   │   └── api.yaml              # Main API specification
│   ├── generated/                # Generated Go code from OpenAPI
│   │   └── go/openapi/           # Generated Go models and utilities
│   └── openapi_generator_config.yml # OpenAPI generator configuration
├── graphql/                       # GraphQL schema definitions
│   ├── User.graphqls             # User GraphQL schema
│   └── Order.graphqls            # Order GraphQL schema
├── gqlgen.yml                    # gqlgen configuration
├── configs/                       # Configuration files
│   ├── app-dev.yaml              # Development configuration
│   └── app-prod.yaml             # Production configuration
├── migrations/                    # Database migrations with optimistic locking support
├── proto/                         # Protocol buffer definitions
├── gorm/                          # GORM generated code with optimistic locking
├── scripts/                       # Build and deployment scripts
├── bruno/                         # API testing
├── docker-compose.yml             # Docker Compose for local development
└── ... (other config files)
```

## 🎯 Key Architectural Decisions

### 1. Clean Architecture with Clear Layer Separation

- **Decision**: Follow clean architecture with distinct layers for domain, service, delivery, and infrastructure
- **Why**: Ensures testability, maintainability, and independence from external concerns
- **Structure**: Domain (interfaces) → Service (business logic) → Delivery (transport) → Infrastructure (implementations)

### 2. Multi-Protocol API Support

- **Decision**: HTTP (REST + GraphQL) and gRPC with configuration-driven selection
- **Why**: Flexibility to serve different client types and deployment scenarios
- **Implementation**: HTTP serves both OpenAPI and GraphQL, gRPC is separate protocol

### 3. Repository Pattern in Infrastructure

- **Decision**: Repository interfaces in domain, implementations in infrastructure
- **Why**: Maintains clean architecture while providing data access abstraction
- **Benefit**: Easy to swap database implementations and test with mocks

### 4. Shared Handler Architecture

- **Decision**: Common handler interface and default implementation for all HTTP handlers
- **Why**: Consistent error handling and middleware application across all endpoints
- **Benefit**: Reduces code duplication and ensures consistent behavior

### 5. Observability with OpenTelemetry

- **Decision**: OpenTelemetry tracing with Jaeger backend for distributed tracing
- **Why**: Provides visibility into request flows across services and infrastructure
- **Benefit**: Better debugging, performance monitoring, and operational insights

### 6. Optimistic Locking for Data Consistency

- **Decision**: Implement optimistic locking using GORM's optimistic lock plugin
- **Why**: Prevents data corruption in concurrent update scenarios without performance penalties of pessimistic locking
- **Implementation**: Version field in database tables, automatic version checking in updates
- **Benefit**: Better performance, handles concurrent updates gracefully, prevents lost updates

### 7. Database Seeding for Development

- **Decision**: Comprehensive database seeding system with realistic test data
- **Why**: Provides consistent development environment and realistic data for testing
- **Implementation**: Separate seeding command with configurable data generation
- **Benefit**: Faster development setup, better testing scenarios, consistent demo data

## 🔄 Data Flow

### Request Flow

1. **HTTP Request** → HTTP Server (chi router)
2. **Middleware** → Request ID, CORS, rate limiting, logging
3. **Handler** → Domain-specific handler implementation
4. **Service** → Business logic with domain models
5. **Repository** → Data access through GORM
6. **Database** → PostgreSQL with optimistic locking
7. **Response** → Structured response with proper HTTP status codes

### Dependency Flow

1. **Entry Point** (`cmd/server/main.go`) → Composes all FX modules
2. **FX Modules** → Each layer provides its dependencies
3. **Module Dependencies** → Dependencies flow through module hierarchy
4. **Service Startup** → `fx.Invoke(start)` triggers application startup

### Transaction Flow

1. **Service Method** → Begins transaction
2. **Repository Operations** → Execute within transaction context
3. **Commit/Rollback** → Based on business logic success/failure
4. **Error Handling** → Proper error propagation and logging

## 🧪 Testing Strategy

### Unit Testing

- **Domain Logic**: Pure functions with no external dependencies
- **Service Layer**: Mocked repositories for isolated testing
- **Handlers**: Mocked services for HTTP layer testing

### Integration Testing

- **Repository Tests**: Real database with test data
- **Service Tests**: Real repositories with mocked external dependencies
- **End-to-End**: Full request flow through all layers

### Concurrent Testing

- **Optimistic Locking**: Test concurrent update scenarios
- **Race Conditions**: Verify data consistency under load
- **Performance**: Measure response times and throughput

## 📊 Monitoring & Observability

### Logging

- **Structured Logging**: JSON format with consistent field names
- **Log Levels**: Configurable per environment (debug, info, warn, error)
- **Source Tracking**: File and line numbers for debugging
- **Context**: Request ID, user ID, correlation IDs

### Tracing

- **OpenTelemetry**: Distributed tracing across services
- **Jaeger Backend**: Trace visualization and analysis
- **Custom Spans**: Business logic tracing and performance monitoring
- **Error Tracking**: Automatic error correlation with traces

### Metrics

- **Request Counts**: API endpoint usage statistics
- **Response Times**: Performance monitoring and alerting
- **Error Rates**: Error tracking and alerting
- **Database Metrics**: Query performance and connection pool status
