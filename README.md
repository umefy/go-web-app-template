# Go Web App Template

A production-ready Go web application template following clean architecture principles with HTTP/gRPC servers, database integration, and comprehensive tooling.

## 1. Quick Start

- Run `./scripts/local_setup.sh` to setup the tools required by the project. After setup, please update `.envrc` and `.envrc.test` based on your own needs.
- Run `make openapi_to_proto` to **generate** the API required proto file.
- Run `make regen_proto` to **regenerate** all the golang code from `proto` files.
- Run `go mod tidy` to get all the dependencies.
- Check `configs` folder and `.envrc` file, especially `.envrc` file, it contains several env vars, you will need to modify it such as `DATABASE_URL` to your own config.
- Run `make migration_create migration_name=[MigrationName]` to create migration
- Run `make migration_up` to do all the migrations for the database.
- Run `make regen_gorm` to **generate** all the database models and query.
- Run `make wire` to **generate** all the required dependency injection files.
- Run `make mockery` to **generate** all the testing required mockery package.
- All above generated commands can be combined with `make generate`.
- For testing, can run `make generate ENVRC_FILE=.envrc.test` to specify we use `.envrc.test` instead of `.envrc`.
- Run `make` to start the project in dev env. 🚀

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
│   │   ├── http/                 # HTTP infrastructure
│   │   │   ├── handler/          # Domain-specific handlers
│   │   │   │   └── user/         # User HTTP handlers
│   │   │   ├── middleware/       # HTTP middleware
│   │   │   ├── router.go         # Main router setup
│   │   │   └── server.go         # HTTP server
│   │   ├── grpc/                 # gRPC infrastructure
│   │   ├── config/               # Configuration loading
│   │   └── logger/               # Logger setup
│   └── app/                      # Application composition & DI
│       ├── app.go                # Main application struct
│       ├── wire.go               # Dependency injection
│       └── wire_gen.go           # Generated wire code
├── pkg/                           # Public reusable packages
│   ├── config/                   # Generic config utilities
│   ├── db/                       # Database utilities
│   ├── server/                   # Server utilities
│   ├── validation/               # Validation utilities
│   └── null/                     # Null value utilities
├── migrations/                    # Database migrations
├── proto/                         # Protocol buffer definitions
├── openapi/                       # OpenAPI specifications
├── gorm/                          # GORM generated code
├── scripts/                       # Build and deployment scripts
├── bruno/                         # API testing
├── configs/                       # Configuration files
```

### Architecture Principles

- **Domain Layer**: Pure business logic with interfaces only
- **Infrastructure Layer**: External implementations (database, HTTP, gRPC)
- **App Layer**: Dependency injection and composition
- **Dependency Direction**: Domain ← Infrastructure (Domain doesn't know about infrastructure)

### Key Features

- ✅ **Clean Architecture**: Clear separation of concerns
- ✅ **Dependency Injection**: Wire-based DI with proper bindings
- ✅ **Transaction Support**: Full transaction handling in services and repositories
- ✅ **Domain-Driven Design**: Organized by business domains
- ✅ **HTTP/gRPC Support**: Both transport protocols supported
- ✅ **Database Integration**: GORM with migrations
- ✅ **API-First Development**: OpenAPI-driven development
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

### API Development (OpenAPI First)

- Update `./openapi/docs/api.yaml` with new API definitions
- Run `make openapi_to_proto` to generate proto files
- Run `make regen_proto` to generate Go code
- Implement handlers in `internal/infrastructure/http/handler/`

### Database Development

- Create migration: `make migration_create migration_name=[MigrationName]`
- Write SQL in the generated migration file
- Run migration: `make migration_up`
- Generate models: `make regen_gorm`

### Adding New Domains

1. Create domain structure: `internal/domain/[domain]/`
2. Define repository interface and service
3. Implement repository in `internal/infrastructure/database/`
4. Add HTTP handlers in `internal/infrastructure/http/handler/[domain]/`
5. Update wire configuration

## 4. Future Enhancements

### Planned Features

- [ ] **Authentication & Authorization**: JWT, OAuth2, RBAC
- [ ] **Caching Layer**: Redis integration
- [ ] **Event System**: Domain events and messaging
- [ ] **Metrics & Observability**: Prometheus, OpenTelemetry
- [ ] **API Versioning**: Proper API version management
- [ ] **Background Jobs**: Task queue integration
- [ ] **File Upload**: Multipart file handling
- [ ] **Email Integration**: SMTP/email service
- [ ] **WebSocket Support**: Real-time communication

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
