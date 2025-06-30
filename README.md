# Go Web App Template

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
│   ├── domain/                    # Pure business logic (no external dependencies)
│   │   ├── user/                 # User domain
│   │   │   ├── repository/       # Repository interfaces
│   │   │   ├── service/          # Business logic services
│   │   │   └── error/            # Domain-specific errors
│   │   ├── config/               # Configuration domain
│   │   ├── greeter/              # Greeter domain
│   │   └── logger/               # Logger domain
│   ├── infrastructure/            # External concerns & implementations
│   │   ├── database/             # Database implementations
│   │   │   ├── setup.go          # Database connection
│   │   │   ├── user_repository.go # Repository implementations
│   │   │   └── wire.go           # Database wire set
│   │   ├── http/                 # HTTP infrastructure (REST + GraphQL)
│   │   │   ├── handler/          # Default HTTP handlers
│   │   │   │   ├── default_handler.go # Default handler
│   │   │   │   ├── handler.go    # Handler interface
│   │   │   │   └── middleware/   # HTTP middleware
│   │   │   ├── graphql/          # GraphQL infrastructure
│   │   │   │   ├── generated.go  # gqlgen generated code
│   │   │   │   ├── User.resolvers.go # User GraphQL resolvers
│   │   │   │   ├── Order.resolvers.go # Order GraphQL resolvers
│   │   │   │   ├── router.go     # GraphQL router with playground
│   │   │   │   └── model/        # GraphQL models
│   │   │   ├── openapi/          # OpenAPI infrastructure
│   │   │   │   └── v1/           # API version 1
│   │   │   │       ├── router.go # OpenAPI router
│   │   │   │       └── handler/  # OpenAPI handlers
│   │   │   │           └── user/ # User OpenAPI handlers
│   │   │   │               ├── create_user.go
│   │   │   │               ├── get_user.go
│   │   │   │               ├── get_users.go
│   │   │   │               ├── update_user.go
│   │   │   │               ├── handler.go
│   │   │   │               └── mapping/ # Data mapping
│   │   │   ├── middleware/       # HTTP middleware
│   │   │   ├── router.go         # Main router setup
│   │   │   └── server.go         # HTTP server
│   │   ├── grpc/                 # gRPC infrastructure
│   │   │   ├── handler/          # gRPC handlers
│   │   │   │   └── greeter/      # Greeter gRPC service
│   │   │   └── server.go         # gRPC server
│   │   ├── config/               # Configuration loading
│   │   └── logger/               # Logger setup
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

- **Domain Layer**: Pure business logic with interfaces only
- **Infrastructure Layer**: External implementations (database, HTTP, gRPC)
- **App Layer**: Dependency injection and composition
- **Dependency Direction**: Domain ← Infrastructure (Domain doesn't know about infrastructure)

### Key Architectural Decisions

#### 1. Clean Architecture Implementation

- **Decision**: Follow clean architecture with clear layer separation
- **Why**: Ensures testability, maintainability, and independence from external concerns
- **Structure**: Domain (business logic) → Infrastructure (implementations) → App (composition)

#### 2. Multi-Protocol API Support

- **Decision**: HTTP (REST + GraphQL) and gRPC with configuration-driven selection
- **Why**: Flexibility to serve different client types and deployment scenarios
- **Implementation**: HTTP serves both OpenAPI and GraphQL, gRPC is separate protocol

#### 3. Repository Pattern in Infrastructure

- **Decision**: Repository interfaces in domain, implementations in infrastructure
- **Why**: Maintains clean architecture while providing data access abstraction
- **Benefit**: Easy to swap database implementations and test with mocks

### Key Features

- ✅ **Clean Architecture**: Clear separation of concerns
- ✅ **Dependency Injection**: Wire-based DI with proper bindings
- ✅ **Transaction Support**: Full transaction handling in services and repositories
- ✅ **Domain-Driven Design**: Organized by business domains
- ✅ **Multi-Protocol Support**: HTTP (REST + GraphQL) and gRPC with configuration-driven selection
- ✅ **Database Integration**: GORM with migrations
- ✅ **Multi-Protocol API**: OpenAPI (REST) + GraphQL support on HTTP
- ✅ **API-First Development**: OpenAPI-driven development with Go model generation
- ✅ **GraphQL Support**: gqlgen-based GraphQL server with type-safe resolvers and playground
- ✅ **Comprehensive Testing**: Mockery for mocking
- ✅ **Health Checks**: Built-in health check endpoints with chi middleware
- ✅ **Rate Limiting**: Request throttling with httprate (600 req/min global, 100 req/min per IP)
- ✅ **Logging**: Structured logging with request ID tracking
- ✅ **Input Validation**: Comprehensive validation with custom validation rules
- ✅ **Content Type Validation**: JSON content type enforcement
- ✅ **Request Timeout**: 60-second request timeout
- ✅ **Error Recovery**: Panic recovery with logging
- ✅ **GitHub Actions**: CI/CD workflows for linting and testing
- ✅ **Profiling**: Built-in debug profiler endpoint

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
- Implement handlers in `internal/infrastructure/http/openapi/v1/handler/[domain]/`

#### GraphQL Development

- Update GraphQL schema files in `graphql/` directory
- Run `make regen_graphql` to generate resolvers and models
- Implement resolvers in `internal/infrastructure/http/graphql/`
- Access GraphQL playground at `/graphql/playground` in development

#### gRPC Development

- Update protocol buffer definitions in `proto/` directory
- Run `make regen_proto` to generate Go code
- Implement handlers in `internal/infrastructure/grpc/handler/`
- Enable gRPC server in configuration when needed

### Database Development

- Create migration: `make migration_create migration_name=[MigrationName]`
- Write SQL in the generated migration file
- Run migration: `make migration_up`
- Generate models: `make regen_gorm`

### Adding New Domains

1. Create domain structure: `internal/domain/[domain]/`
2. Define repository interface and service
3. Implement repository in `internal/infrastructure/database/`
4. Add REST handlers in `internal/infrastructure/http/openapi/v1/handler/[domain]/`
5. Add GraphQL resolvers in `internal/infrastructure/http/graphql/`
6. Add gRPC handlers in `internal/infrastructure/grpc/handler/` (if needed)
7. Update wire configuration

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

## 5. Contributing

This is a template project designed for rapid development of Go web applications. Feel free to fork and customize for your specific needs.

### Best Practices

- Follow clean architecture principles
- Write tests for all business logic
- Use dependency injection for all dependencies
- Keep domain logic pure and infrastructure-agnostic
- Document APIs with OpenAPI specifications
- Use transactions for data consistency
- Implement proper error handling and logging
