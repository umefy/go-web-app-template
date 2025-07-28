# Go Web App Template

[![Go Report Card](https://goreportcard.com/badge/github.com/umefy/go-web-app-template)](https://goreportcard.com/report/github.com/umefy/go-web-app-template)
[![Go Version](https://img.shields.io/github/go-mod/go-version/umefy/go-web-app-template)](https://go.dev/)
[![License](https://img.shields.io/github/license/umefy/go-web-app-template)](LICENSE)

A production-ready Go web application template following clean architecture principles with HTTP/gRPC servers, database integration, and comprehensive tooling.

## 1. Quick Start

1. **Setup Environment**

   - Run `./scripts/local_setup.sh` to setup the tools required by the project
   - Update `.envrc` and `.envrc.test` based on your own needs (especially `DATABASE_URL`)

2. **Generate Code**

   - Run `make generate` to generate all required files (OpenAPI models, GraphQL resolvers, proto code, GORM models, Wire DI, and mocks)
   - Or run individual commands:
     - `make regen_openapi` - Generate Go models from OpenAPI specification
     - `make regen_graphql` - Generate GraphQL resolvers and models
     - `make regen_proto` - Generate Go code from proto files
     - `make regen_gorm` - Generate database models and queries
     - `make wire` - Generate dependency injection files
     - `make mockery` - Generate testing mocks

3. **Database Setup**

   - Run `make migration_create migration_name=[MigrationName]` to create a new migration
   - Write SQL in the generated migration file
   - Run `make migration_up` to apply all migrations

4. **Start Development**
   - Run `make` or `make dev` to start the project in development mode 🚀
   - Access APIs:
     - REST API: `http://localhost:8082/api/v1/`
     - GraphQL: `http://localhost:8082/graphql/` (playground available in dev mode)
     - Health check: `http://localhost:8082/health`

**Note**: For testing, run `make generate ENVRC_FILE=.envrc.test` to use test environment configuration.

## 2. Project Structure

This project follows **Clean Architecture** principles with a clear separation of concerns:

```bash
go-web-app-template/
├── cmd/                           # Application entry points
│   └── server/                   # HTTP/gRPC server startup
├── internal/                      # Private application code (Go-enforced privacy)
│   ├── domain/                    # Pure business logic & interfaces (no external dependencies)
│   │   ├── user/                 # User domain
│   │   │   ├── repo/             # Repository interfaces
│   │   │   ├── user.go           # Domain models
│   │   │   ├── user_with_order.go # Domain models with relationships
│   │   │   └── error/            # Domain-specific errors
│   │   ├── order/                # Order domain
│   │   │   ├── repo/             # Repository interfaces
│   │   │   └── order.go          # Domain models
│   │   └── error/                # Shared domain errors
│   ├── service/                   # Business logic implementation
│   │   ├── user/                 # User business logic
│   │   │   ├── service.go        # User service implementation
│   │   │   ├── wire.go           # Service wire set
│   │   │   └── input.go          # Service input/output models
│   │   ├── order/                # Order business logic
│   │   │   ├── service.go        # Order service implementation
│   │   │   └── wire.go           # Service wire set
│   │   └── greeter/              # Greeter business logic
│   │       ├── service.go        # Greeter service implementation
│   │       └── wire.go           # Service wire set
│   ├── delivery/                  # Transport layer (HTTP/gRPC/GraphQL)
│   │   ├── restful/              # HTTP REST API
│   │   │   ├── handler/          # Shared handler utilities
│   │   │   │   ├── handler.go    # Handler interface
│   │   │   │   ├── default_handler.go # Default handler with error handling
│   │   │   │   └── middleware/   # HTTP middleware
│   │   │   └── openapi/          # OpenAPI REST endpoints
│   │   │       └── v1/           # API version 1
│   │   │           ├── router.go # OpenAPI router
│   │   │           └── user/     # User REST handlers
│   │   │               ├── handler.go      # User handler interface
│   │   │               ├── create_user.go  # Create user endpoint
│   │   │               ├── get_user.go     # Get user endpoint
│   │   │               ├── get_users.go    # Get users endpoint
│   │   │               ├── update_user.go  # Update user endpoint
│   │   │               ├── router.go       # User routing
│   │   │               └── mapping/        # Data mapping
│   │   ├── graphql/              # GraphQL API
│   │   │   ├── generated.go      # gqlgen generated code
│   │   │   ├── User.resolvers.go # User GraphQL resolvers
│   │   │   ├── Order.resolvers.go # Order GraphQL resolvers
│   │   │   ├── router.go         # GraphQL router with playground
│   │   │   └── model/            # GraphQL models
│   │   ├── grpc/                 # gRPC API
│   │   │   ├── handler/          # gRPC handlers
│   │   │   │   └── greeter/      # Greeter gRPC service
│   │   │   └── server.go         # gRPC server
│   │   └── errutil/              # Error handling utilities
│   ├── infrastructure/            # External concerns & implementations
│   │   ├── database/             # Database infrastructure
│   │   │   ├── wire.go           # Database wire set
│   │   │   ├── with_tx.go        # Transaction utilities
│   │   │   ├── ctx/              # Database context utilities
│   │   │   └── gorm/             # GORM database implementation
│   │   │       ├── setup.go      # Database connection setup
│   │   │       ├── with_tx.go    # GORM transaction utilities
│   │   │       ├── generated/    # GORM generated models
│   │   │       └── repo/         # Repository implementations
│   │   │           ├── user_repo.go    # User repository implementation
│   │   │           ├── order_repo.go   # Order repository implementation
│   │   │           └── mapping/        # Database mapping utilities
│   │   ├── server/               # Server infrastructure
│   │   │   ├── http/             # HTTP server setup
│   │   │   │   ├── router.go     # Main HTTP router
│   │   │   │   ├── server.go     # HTTP server
│   │   │   │   └── middleware/   # HTTP middleware
│   │   │   └── grpc/             # gRPC server setup
│   │   │       ├── handler/      # gRPC handlers
│   │   │       └── server.go     # gRPC server
│   │   └── logger/               # Logger setup
│   ├── core/                      # Core shared components
│   │   └── config/               # Configuration management
│   │       ├── config.go         # Main configuration struct
│   │       ├── setup.go          # Configuration loading
│   │       ├── app_config.go     # Application configuration
│   │       ├── db_config.go      # Database configuration
│   │       ├── http_server_config.go # HTTP server configuration
│   │       ├── grpc_server_config.go # gRPC server configuration
│   │       └── logging_config.go # Logging configuration
│   └── app/                      # Application composition & DI
│       ├── app.go                # Main application struct
│       ├── wire.go               # Dependency injection
│       └── wire_gen.go           # Generated wire code
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
├── migrations/                    # Database migrations
├── proto/                         # Protocol buffer definitions
├── gorm/                          # GORM generated code
├── scripts/                       # Build and deployment scripts
├── bruno/                         # API testing
└── ... (other config files)
```

### Architecture Principles

- **Domain Layer**: Pure business logic with interfaces and models only
- **Service Layer**: Business logic implementation using domain interfaces
- **Delivery Layer**: Transport concerns (HTTP, GraphQL, gRPC)
- **Infrastructure Layer**: External implementations (database, server, logger)
- **Core Layer**: Shared core components (configuration)
- **App Layer**: Dependency injection and composition
- **Dependency Direction**: Domain ← Service ← Delivery ← Infrastructure (Domain doesn't know about external concerns)

### Key Architectural Decisions

#### 1. Clean Architecture with Clear Layer Separation

- **Decision**: Follow clean architecture with distinct layers for domain, service, delivery, and infrastructure
- **Why**: Ensures testability, maintainability, and independence from external concerns
- **Structure**: Domain (interfaces) → Service (business logic) → Delivery (transport) → Infrastructure (implementations)

#### 2. Multi-Protocol API Support

- **Decision**: HTTP (REST + GraphQL) and gRPC with configuration-driven selection
- **Why**: Flexibility to serve different client types and deployment scenarios
- **Implementation**: HTTP serves both OpenAPI and GraphQL, gRPC is separate protocol

#### 3. Repository Pattern in Infrastructure

- **Decision**: Repository interfaces in domain, implementations in infrastructure
- **Why**: Maintains clean architecture while providing data access abstraction
- **Benefit**: Easy to swap database implementations and test with mocks

#### 4. Shared Handler Architecture

- **Decision**: Common handler interface and default implementation for all HTTP handlers
- **Why**: Consistent error handling and middleware application across all endpoints
- **Benefit**: Reduces code duplication and ensures consistent behavior

### Key Features

- ✅ **Clean Architecture**: Clear separation of concerns with domain, service, delivery, and infrastructure layers
- ✅ **Dependency Injection**: Wire-based DI with proper bindings
- ✅ **Transaction Support**: Full transaction handling in services and repositories
- ✅ **Domain-Driven Design**: Organized by business domains with clear boundaries
- ✅ **Multi-Protocol Support**: HTTP (REST + GraphQL) and gRPC with configuration-driven selection
- ✅ **Database Integration**: GORM with migrations and generated queries
- ✅ **API-First Development**: OpenAPI-driven development with Go model generation
- ✅ **GraphQL Support**: gqlgen-based GraphQL server with type-safe resolvers and playground
- ✅ **Comprehensive Testing**: Mockery for mocking with comprehensive test coverage
- ✅ **Health Checks**: Built-in health check endpoints with chi middleware
- ✅ **Rate Limiting**: Request throttling with httprate (600 req/min global, 100 req/min per IP)
- ✅ **Logging**: Structured logging with request ID tracking and context-aware logging
- ✅ **Input Validation**: Comprehensive validation with custom validation rules
- ✅ **Content Type Validation**: JSON content type enforcement
- ✅ **Request Timeout**: 60-second request timeout
- ✅ **Error Recovery**: Panic recovery with logging and graceful error handling
- ✅ **GitHub Actions**: CI/CD workflows for linting and testing
- ✅ **Profiling**: Built-in debug profiler endpoint for performance analysis

## 3. Development Workflow

### Protocol Selection

The application supports multiple transport protocols with configuration-driven selection:

```yaml
# configs/app-dev.yaml
http_server:
  enabled: true # Enable HTTP (REST + GraphQL)
  port: 8082

grpc_server:
  enabled: false # Disable gRPC
  port: 30082
```

- **HTTP Protocol**: Serves both OpenAPI (REST) and GraphQL APIs
- **gRPC Protocol**: Separate gRPC services (when enabled)
- **Configuration**: Easy to enable/disable protocols via YAML config

### API Development (Multi-Protocol)

#### OpenAPI (REST) Development

- Update `./openapi/docs/api.yaml` with new API definitions
- Run `make regen_openapi` to generate Go models
- Implement handlers in `internal/delivery/restful/openapi/v1/[domain]/`

#### GraphQL Development

- Update GraphQL schema files in `graphql/` directory
- Run `make regen_graphql` to generate resolvers and models
- Implement resolvers in `internal/delivery/graphql/`
- Access GraphQL playground at `/graphql/playground` in development

#### gRPC Development

- Update protocol buffer definitions in `proto/` directory
- Run `make regen_proto` to generate Go code
- Implement handlers in `internal/delivery/grpc/handler/`
- Enable gRPC server in configuration when needed

### Database Development

- Create migration: `make migration_create migration_name=[MigrationName]`
- Write SQL in the generated migration file
- Run migration: `make migration_up`
- Generate models: `make regen_gorm`

### Adding New Domains

1. **Create Domain Structure**: `internal/domain/[domain]/`

   - Define domain models (e.g., `user.go`)
   - Create repository interface in `repo/repo.go`
   - Add domain-specific errors in `error/`

2. **Implement Business Logic**: `internal/service/[domain]/`

   - Create service interface and implementation
   - Add service wire set
   - Define input/output models

3. **Add Transport Layer**:

   - **REST**: Add handlers in `internal/delivery/restful/openapi/v1/[domain]/`
   - **GraphQL**: Add resolvers in `internal/delivery/graphql/`
   - **gRPC**: Add handlers in `internal/delivery/grpc/handler/` (if needed)

4. **Implement Data Access**: `internal/infrastructure/database/gorm/repo/`

   - Create repository implementation
   - Add database mapping utilities

5. **Update Wire Configuration**: Add new services and repositories to DI

## 4. Future Enhancements

### Planned Features

- [ ] **Authentication & Authorization**: JWT, OAuth2, RBAC
- [ ] **Caching Layer**: Redis integration
- [ ] **Event System**: Domain events and messaging
- [ ] **Metrics & Observability**: Prometheus, OpenTelemetry
- [ ] **Advanced API Versioning**: Multiple version coexistence, deprecation policies, migration guides
- [ ] **Background Jobs**: Task queue integration
- [ ] **File Upload**: Multipart file handling
- [ ] **Email Integration**: SMTP/email service

### Infrastructure Improvements

- [ ] **Docker Support**: Multi-stage builds
- [ ] **Kubernetes**: Deployment manifests
- [ ] **Enhanced CI/CD**: Additional GitHub Actions workflows
- [ ] **Monitoring**: Enhanced health checks, metrics dashboard
- [ ] **Security**: Security headers, enhanced CORS configuration
- [ ] **Performance**: Connection pooling, caching strategies

### Development Experience

- [ ] **Enhanced Testing**: Integration tests, performance tests, test coverage reporting
- [ ] **Development Tools**: Additional development utilities and scripts
- [ ] **Performance Monitoring**: Development-time performance insights
- [ ] **Database Tools**: Database schema visualization, query optimization tools

## 5. License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### What MIT License Means for You

- ✅ **Use freely**: Use this template for any purpose (personal, commercial, etc.)
- ✅ **Modify freely**: Customize and adapt the code to your needs
- ✅ **Distribute freely**: Share your modified versions
- ✅ **Minimal requirements**: Just include the original license and copyright notice
- ✅ **No warranty**: The software is provided "as is" without warranties

## 6. Contributing

This is a template project designed for rapid development of Go web applications. Feel free to fork and customize for your specific needs.

### Best Practices

- Follow clean architecture principles with clear layer separation
- Write tests for all business logic in the service layer
- Use dependency injection for all dependencies
- Keep domain logic pure and infrastructure-agnostic
- Document APIs with OpenAPI specifications
- Use transactions for data consistency
- Implement proper error handling and logging
- Use the shared handler architecture for consistent HTTP handling
- Organize code by business domains rather than technical concerns
