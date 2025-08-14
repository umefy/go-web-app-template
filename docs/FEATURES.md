# Features Guide

This document provides detailed explanations of all features, configuration options, and best practices.

## ✨ Key Features Overview

- **🏗️ Clean Architecture** with FX dependency injection
- **🌐 Multi-Protocol APIs** - HTTP (REST + GraphQL) + gRPC
- **🗄️ Database** - GORM + migrations + optimistic locking
- **📊 Observability** - OpenTelemetry + structured logging
- **🐳 Development** - Docker Compose + seeding + testing tools
- **🔒 Security** - Input validation + rate limiting + CORS
- **🧪 Testing** - Comprehensive testing + mocking + concurrent testing

## 🏗️ Clean Architecture

### Architecture Layers

- **Domain Layer**: Pure business logic with interfaces and models only
- **Service Layer**: Business logic implementation using domain interfaces
- **Delivery Layer**: Transport concerns (HTTP, GraphQL, gRPC)
- **Infrastructure Layer**: External implementations (database, server, logger, tracing)
- **Core Layer**: Shared core components (configuration)
- **App Layer**: Dependency injection and composition

### Benefits

- **Testability**: Easy to test business logic in isolation
- **Maintainability**: Clear separation of concerns
- **Flexibility**: Easy to swap implementations
- **Scalability**: Modular design supports growth

## 🌐 Multi-Protocol API Support

### HTTP (REST + GraphQL)

The application serves both OpenAPI (REST) and GraphQL APIs through the same HTTP server:

- **REST API**: OpenAPI 3.0 specification with automatic Go model generation
- **GraphQL**: gqlgen-based server with playground for development
- **Shared Middleware**: CORS, rate limiting, logging, tracing

### gRPC

Separate gRPC server for high-performance inter-service communication:

- **Protocol Buffers**: Type-safe message definitions
- **Code Generation**: Automatic Go code generation
- **Configuration**: Enable/disable via YAML configuration

### Protocol Selection

```yaml
# configs/app-dev.yaml
http_server:
  enabled: true # Enable HTTP (REST + GraphQL)
  port: 8082

grpc_server:
  enabled: false # Disable gRPC
  port: 30082
```

## 🗄️ Database Features

### GORM Integration

- **ORM**: Full-featured Go ORM with database agnostic design
- **Migrations**: Version-controlled database schema changes
- **Generated Queries**: Type-safe query building with code generation
- **Transactions**: Full transaction support with context

### Optimistic Locking

Prevents data corruption in concurrent update scenarios:

**How It Works:**

1. **Version Field**: Each entity has a `version` field that increments on each update
2. **Update Validation**: Updates check that the current version matches the expected version
3. **Conflict Detection**: If versions don't match, the update fails with a conflict error
4. **Automatic Handling**: GORM automatically increments the version field on successful updates

**Database Schema:**

```sql
-- Users table with optimistic locking
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    age INT NOT NULL,
    version BIGINT NOT NULL DEFAULT 0,  -- Optimistic lock version
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Orders table with optimistic locking
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount_cents BIGINT NOT NULL,
    version BIGINT NOT NULL DEFAULT 0,  -- Optimistic lock version
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

**Error Handling:**

```go
// User update conflict error
var UserUpdateConflict = appError.NewError(
    "userService_1003",
    "user update conflict - version mismatch",
    http.StatusConflict,
)
```

**Testing Concurrent Updates:**

```bash
# Test optimistic locking with concurrent updates
go run cmd/concurrent/concurrent_user_update.go
```

### Database Seeding

Comprehensive seeding system for development and testing:

**Seeding Command:**

```bash
# Seed database with sample data
make seed_database
```

**What Gets Seeded:**

- **Users**: 100 users with realistic email addresses and ages
- **Orders**: 1000 orders linked to random users with realistic amounts
- **Data Generation**: Uses `gofakeit` for realistic test data

**Seeding Process:**

1. **Database Reset**: Clears existing data and runs migrations
2. **User Creation**: Creates 100 users with fake data
3. **Order Creation**: Creates 1000 orders linked to users
4. **Batch Processing**: Uses efficient batch inserts for performance

**Customization:**

```go
// Adjust seeding quantities
const userCount = 100
const orderCount = 1000

// Customize data generation
users[i] = &dbModel.User{
    Email: null.ValueFrom(gofakeit.Email()),
    Age:   null.ValueFrom(gofakeit.IntRange(0, 60)),
}
```

**Benefits:**

- **Consistent Environment**: Same data across all development instances
- **Realistic Testing**: Test with realistic data volumes and relationships
- **Quick Setup**: Fast development environment initialization
- **Demo Ready**: Immediate demonstration of application capabilities

## 📊 Observability Features

### Advanced Logging Configuration

The application includes comprehensive logging with multiple configuration options:

#### Application Logging

- **Log Levels**: Configurable levels (debug, info, warn, error)
- **Output Format**: JSON or plain text formatting
- **Source Tracking**: Optional file and line number logging for debugging
- **Process ID**: Automatic PID inclusion in all log entries
- **Output Writers**: Configurable output destinations (currently stdout)

#### Database Logging

- **SQL Logging**: Configurable GORM SQL query logging
- **Slow Query Detection**: Configurable threshold for slow query identification
- **Parameter Visibility**: Optional SQL parameter logging for debugging
- **Colorful Output**: Enhanced readability in development environments

#### Configuration Examples

```yaml
# Development logging (configs/app-dev.yaml)
logging:
  level: debug
  writer: stdout
  use_json: false
  add_source: true
  source_key: "source"

database:
  logger:
    level: info
    writer: stdout
    show_sql_params: true
    slow_threshold_in_seconds: 1

# Production logging (configs/app-prod.yaml)
logging:
  level: info
  writer: stdout
  use_json: true
  add_source: true
  source_key: "source"

database:
  logger:
    level: info
    writer: stdout
    show_sql_params: false
    slow_threshold_in_seconds: 1
```

### OpenTelemetry Tracing

Comprehensive tracing support with Jaeger backend:

- **Tracing**: OpenTelemetry integration with Jaeger backend
- **Jaeger UI**: Access traces at `http://localhost:16686`
- **Configuration**: Tracing can be enabled/disabled per environment
- **Service Context**: Automatic service name, version, and tracer configuration

### Git Hash Version Injection

The application automatically injects git commit hashes into configuration values marked with `""` (empty string). This provides automatic version tracking for both the application and tracing services.

**Configuration Values with Git Hash Injection:**

```yaml
# configs/app-dev.yaml
version: '' # inject git hash from .envrc
tracing:
  service_version: '' # inject git hash from .envrc
```

**Environment Variable Setup:**

```bash
# .envrc (or .envrc.example)
export VERSION=$(git rev-parse --short HEAD)
```

**How It Works:**

1. **Environment Variable**: `VERSION` is set to the short git commit hash
2. **Configuration Loading**: Viper automatically replaces empty strings with environment variables
3. **Automatic Updates**: Version updates automatically with each git commit
4. **Tracing Integration**: Service version is automatically set for OpenTelemetry traces

**Benefits:**

- **Automatic Versioning**: No manual version updates required
- **Trace Correlation**: Easy to correlate traces with specific code versions
- **Deployment Tracking**: Know exactly which code version is running
- **Debugging**: Quickly identify which commit introduced issues

**Configuration:**

```yaml
# configs/app-dev.yaml
env: dev
version: '' # inject git hash from .envrc
tracing:
  enabled: true
  jaeger_endpoint: 'localhost:4318'
  service_name: 'Server'
  service_version: '' # inject git hash from .envrc
```

## 🔒 Security Features

### Input Validation

Comprehensive validation with custom validation rules:

- **Struct Validation**: Automatic validation of request structs
- **Custom Rules**: Domain-specific validation logic
- **Error Messages**: Clear, user-friendly validation errors
- **Type Safety**: Strong typing with Go's type system

### Rate Limiting

Request throttling to prevent abuse:

- **Global Limit**: 600 requests per minute across all endpoints
- **Per-IP Limit**: 100 requests per minute per IP address
- **Configurable**: Easy to adjust limits via configuration
- **Graceful Handling**: Proper HTTP 429 responses

### CORS Configuration

Cross-Origin Resource Sharing support:

- **Development**: Permissive CORS for local development
- **Production**: Restrictive CORS for security
- **Configurable**: Easy to customize allowed origins
- **Preflight Support**: Full OPTIONS request handling

### Content Type Validation

JSON content type enforcement:

- **Strict Validation**: Only accepts `application/json` for POST/PUT requests
- **Clear Errors**: Proper HTTP 415 responses for invalid content types
- **Security**: Prevents content type confusion attacks

## 🧪 Testing Features

### Comprehensive Testing

- **Unit Tests**: Isolated testing of business logic
- **Integration Tests**: Database and service integration testing
- **End-to-End Tests**: Full request flow testing
- **Mockery**: Automatic mock generation for dependencies

### Concurrent Testing

Utilities for testing optimistic locking behavior:

- **Race Condition Testing**: Verify data consistency under load
- **Performance Testing**: Measure response times and throughput
- **Stress Testing**: High-load scenario validation

### Testing Commands

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific test package
go test ./internal/service/user/

# Generate mocks
make mockery
```

## 🐳 Development Tools

### Docker Compose

Local development environment:

- **PostgreSQL**: Database with persistent data
- **Jaeger**: Distributed tracing backend
- **Networking**: Proper service discovery and communication
- **Volumes**: Persistent data across container restarts

### Make Commands

Comprehensive build and development commands:

```bash
make dev              # Start development server
make generate         # Generate all code
make test            # Run tests
make seed_database   # Seed with sample data
make migration_up    # Run database migrations
make docker_compose_up   # Start infrastructure services
make build           # Build for production
make clean           # Clean generated files
```

### Code Generation

Automatic code generation for various components:

- **OpenAPI**: Go models from API specifications
- **GraphQL**: Resolvers and models from schemas
- **Protocol Buffers**: Go code from proto definitions
- **GORM**: Database models and queries
- **Mocks**: Testing mocks with Mockery

## 📚 Best Practices

### Code Organization

- **Clean Architecture**: Follow established architectural patterns
- **Domain-Driven Design**: Organize by business domains
- **Interface Segregation**: Small, focused interfaces
- **Dependency Inversion**: Depend on abstractions, not concretions

### Error Handling

- **Domain Errors**: Use domain-specific error types
- **Error Wrapping**: Proper error context preservation
- **HTTP Status Codes**: Return appropriate status codes
- **Logging**: Log errors with full context

### Performance

- **Database Transactions**: Use transactions appropriately
- **Connection Pooling**: Proper database connection management
- **Query Optimization**: Monitor and optimize slow queries
- **Caching**: Implement caching where appropriate

### Security

- **Input Validation**: Validate all external inputs
- **SQL Injection Prevention**: Use parameterized queries
- **Authentication**: Implement proper authentication (when needed)
- **Authorization**: Role-based access control

### Testing

- **Test Coverage**: Maintain high test coverage
- **Mock Dependencies**: Use mocks for external dependencies
- **Test Data**: Use realistic test data
- **Performance Testing**: Test under load

## 🔧 Configuration Management

### Environment Variables

```bash
# .envrc
export DATABASE_URL="postgres://test_user:test_password@localhost:5433/goWebapp_test?sslmode=disable"
export ENV="dev"

# .envrc.test
export DATABASE_URL="postgres://test_user:test_password@localhost:5433/goWebapp_test?sslmode=disable"
export ENV="test"
```

### Configuration Files

Environment-specific configurations:

- **Development**: `configs/app-dev.yaml`
- **Production**: `configs/app-prod.yaml`

### Configuration Validation

All configuration is validated at startup:

- **Required Fields**: Ensure all required config is present
- **Value Validation**: Validate configuration values
- **Type Safety**: Strong typing for configuration structs
- **Default Values**: Sensible defaults for optional settings

## 🚀 Deployment Features

### Building

```bash
# Build for production
make build

# Build with specific environment
make build ENV=prod
```

### Docker Support

```bash
# Build Docker image
docker build -t go-web-app .

# Run with Docker Compose
docker-compose up -d
```

### Health Checks

Built-in health check endpoints:

- **Health Check**: `/health` endpoint for load balancers
- **Readiness Probe**: Service readiness verification
- **Liveness Probe**: Service liveness verification

## 📊 Monitoring & Debugging

### Profiling

Built-in debug profiler endpoint:

```bash
# Development profiler endpoint
open http://localhost:8082/debug
```

### Metrics

Application metrics for monitoring:

- **Request Counts**: API endpoint usage statistics
- **Response Times**: Performance monitoring
- **Error Rates**: Error tracking and alerting
- **Database Metrics**: Query performance monitoring

### Logging

Structured logging with context:

- **Request ID**: Track requests across the system
- **User Context**: User identification and context
- **Performance**: Request timing and performance data
- **Errors**: Detailed error information with stack traces
