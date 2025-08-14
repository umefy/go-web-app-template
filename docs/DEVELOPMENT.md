# Development Guide

This document covers the development workflow, API development, database operations, and how to extend the application with new domains.

## 🚀 Development Workflow

### Infrastructure Setup

The project includes Docker Compose for local development:

```bash
# Start infrastructure services
make docker_compose_up

# Stop infrastructure services
make docker_compose_down

# Start development with auto-infrastructure setup
make dev  # This automatically starts Docker services and runs migrations
```

**Services Available:**

- **PostgreSQL**: `localhost:5433` (user: `test_user`, password: `test_password`, database: `goWebapp_test`)
- **Jaeger**: `http://localhost:16686` (UI), `localhost:4317` (OTLP gRPC), `localhost:4318` (OTLP HTTP)

### Protocol Selection

The application supports multiple transport protocols with configuration-driven selection:

```yaml
# configs/app-dev.yaml
env: dev
version: '' # inject git hash from .envrc
http_server:
  enabled: true # Enable HTTP (REST + GraphQL)
  port: 8082

grpc_server:
  enabled: false # Disable gRPC
  port: 30082

tracing:
  enabled: true # Enable OpenTelemetry tracing
  jaeger_endpoint: 'localhost:4318'
  service_name: 'Server'
  service_version: '' # inject git hash from .envrc

logging:
  level: debug # Development logging level
  add_source: true # Enable source file tracking
  use_json: false # Plain text for development

database:
  logger:
    level: info # Database logging level
    show_sql_params: true # Show SQL parameters in dev
    slow_threshold_in_seconds: 1 # Slow query detection
```

- **HTTP Protocol**: Serves both OpenAPI (REST) and GraphQL APIs
- **gRPC Protocol**: Separate gRPC services (when enabled)
- **Tracing**: OpenTelemetry tracing with Jaeger backend (configurable)
- **Logging**: Configurable application and database logging with source tracking
- **Configuration**: Easy to enable/disable protocols, tracing, and logging via YAML config

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

### Code Generation

```bash
# Generate all required files
make generate

# Individual generation commands
make regen_openapi    # Generate Go models from OpenAPI specification
make regen_graphql    # Generate GraphQL resolvers and models
make regen_proto      # Generate Go code from proto files
make regen_gorm       # Generate database models and queries
make mockery          # Generate testing mocks
```

## 🌐 API Development (Multi-Protocol)

### OpenAPI (REST) Development

1. **Update API Specification**

   - Edit `./openapi/docs/api.yaml` with new API definitions
   - Follow OpenAPI 3.0 specification standards

2. **Generate Go Models**

   ```bash
   make regen_openapi
   ```

3. **Implement Handlers**

   - Create handlers in `internal/delivery/restful/openapi/v1/[domain]/`
   - Follow the shared handler pattern for consistency
   - Implement proper error handling and validation

4. **Add Routing**
   - Update router in `internal/delivery/restful/openapi/v1/router.go`
   - Add middleware as needed

### GraphQL Development

1. **Update Schema**

   - Edit GraphQL schema files in `graphql/` directory
   - Add new types, queries, and mutations

2. **Generate Resolvers**

   ```bash
   make regen_graphql
   ```

3. **Implement Resolvers**
   - Add resolver logic in `internal/delivery/graphql/`
   - Follow the existing resolver patterns
   - Access GraphQL playground at `/graphql/playground` in development

### gRPC Development

1. **Update Protocol Buffers**

   - Edit `.proto` files in `proto/` directory
   - Define new services and message types

2. **Generate Go Code**

   ```bash
   make regen_proto
   ```

3. **Implement Handlers**
   - Add handlers in `internal/delivery/grpc/handler/`
   - Enable gRPC server in configuration when needed

## 🗄️ Database Development

### Creating Migrations

```bash
# Create a new migration
make migration_create migration_name=[MigrationName]

# Example
make migration_create migration_name=AddUserProfile
```

### Writing Migration SQL

```sql
-- Example migration: Add user profile fields
ALTER TABLE users ADD COLUMN profile JSONB;
ALTER TABLE users ADD COLUMN preferences JSONB DEFAULT '{}';
```

### Running Migrations

```bash
# Apply all pending migrations
make migration_up

# Rollback last migration (if supported)
make migration_down
```

### Generating Models

```bash
# Generate GORM models and queries
make regen_gorm
```

### Database Seeding

```bash
# Seed database with sample data
make seed_database
```

**What Gets Seeded:**

- **Users**: 100 users with realistic email addresses and ages
- **Orders**: 1000 orders linked to random users with realistic amounts
- **Data Generation**: Uses `gofakeit` for realistic test data

**Customization:**

```go
// Adjust seeding quantities in cmd/seed/database/seed_users.go
const userCount = 100
const orderCount = 1000

// Customize data generation
users[i] = &dbModel.User{
    Email: null.ValueFrom(gofakeit.Email()),
    Age:   null.ValueFrom(gofakeit.IntRange(0, 60)),
}
```

## 🆕 Adding New Domains

### 1. Create Domain Structure

Create the domain directory structure:

```bash
mkdir -p internal/domain/[domain]/
mkdir -p internal/domain/[domain]/repo/
mkdir -p internal/domain/[domain]/error/
```

**Example for a `product` domain:**

```go
// internal/domain/product/product.go
package product

import (
    "time"
    "github.com/umefy/go-web-app-template/pkg/null"
)

type Product struct {
    ID          int64       `json:"id"`
    Name        string      `json:"name"`
    Description string      `json:"description"`
    Price       int64       `json:"price_cents"`
    Version     int64       `json:"version"` // For optimistic locking
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

// internal/domain/product/repo/repo.go
package repo

import (
    "context"
    "github.com/umefy/go-web-app-template/internal/domain/product"
)

type ProductRepo interface {
    Create(ctx context.Context, product *product.Product) error
    GetByID(ctx context.Context, id int64) (*product.Product, error)
    Update(ctx context.Context, product *product.Product) error
    Delete(ctx context.Context, id int64) error
    List(ctx context.Context, limit, offset int) ([]*product.Product, error)
}
```

### 2. Implement Business Logic

Create the service layer:

```bash
mkdir -p internal/service/[domain]/
```

**Example service implementation:**

```go
// internal/service/product/service.go
package product

import (
    "context"
    "github.com/umefy/go-web-app-template/internal/domain/product"
    "github.com/umefy/go-web-app-template/internal/domain/product/repo"
)

type Service interface {
    CreateProduct(ctx context.Context, input CreateProductInput) (*product.Product, error)
    GetProduct(ctx context.Context, id int64) (*product.Product, error)
    UpdateProduct(ctx context.Context, id int64, input UpdateProductInput) (*product.Product, error)
    DeleteProduct(ctx context.Context, id int64) error
    ListProducts(ctx context.Context, limit, offset int) ([]*product.Product, error)
}

type service struct {
    repo repo.ProductRepo
}

func NewService(repo repo.ProductRepo) Service {
    return &service{repo: repo}
}

// Implement service methods...
```

### 3. Add Transport Layer

#### REST API

```bash
mkdir -p internal/delivery/restful/openapi/v1/[domain]/
```

**Example handler:**

```go
// internal/delivery/restful/openapi/v1/product/handler.go
package product

import (
    "net/http"
    "github.com/umefy/go-web-app-template/internal/delivery/restful/handler"
    "github.com/umefy/go-web-app-template/internal/service/product"
)

type Handler interface {
    CreateProduct(w http.ResponseWriter, r *http.Request)
    GetProduct(w http.ResponseWriter, r *http.Request)
    UpdateProduct(w http.ResponseWriter, r *http.Request)
    DeleteProduct(w http.ResponseWriter, r *http.Request)
    ListProducts(w http.ResponseWriter, r *http.Request)
}

type handler struct {
    handler.DefaultHandler
    service product.Service
}

func NewHandler(service product.Service) Handler {
    return &handler{service: service}
}

// Implement handler methods...
```

#### GraphQL

Update GraphQL schema and implement resolvers:

```graphql
# graphql/Product.graphqls
type Product {
  id: ID!
  name: String!
  description: String!
  priceCents: Int!
  version: Int!
  createdAt: Time!
  updatedAt: Time!
}

type Query {
  product(id: ID!): Product
  products(limit: Int, offset: Int): [Product!]!
}

type Mutation {
  createProduct(input: CreateProductInput!): Product!
  updateProduct(id: ID!, input: UpdateProductInput!): Product!
  deleteProduct(id: ID!): Boolean!
}
```

### 4. Implement Data Access

Create the repository implementation:

```bash
mkdir -p internal/infrastructure/database/gorm/repo/
```

**Example repository:**

```go
// internal/infrastructure/database/gorm/repo/product_repo.go
package repo

import (
    "context"
    "github.com/umefy/go-web-app-template/internal/domain/product"
    "github.com/umefy/go-web-app-template/internal/domain/product/repo"
    "gorm.io/gorm"
)

type productRepo struct {
    db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repo.ProductRepo {
    return &productRepo{db: db}
}

// Implement repository methods...
```

### 5. Update FX Configuration

Add new services and repositories to appropriate FX modules:

```go
// internal/service/fx.go
func Module() fx.Option {
    return fx.Options(
        fx.Provide(
            user.NewService,
            order.NewService,
            product.NewService, // Add this line
        ),
    )
}

// internal/infrastructure/database/fx.go
func Module() fx.Option {
    return fx.Options(
        fx.Provide(
            gorm.NewUserRepo,
            gorm.NewOrderRepo,
            gorm.NewProductRepo, // Add this line
        ),
    )
}
```

### 6. Add Seeding (Optional)

Add seeding logic in `cmd/seed/database/`:

```go
// cmd/seed/database/seed_products.go
package database

func (s *Seeder) seedProducts() error {
    // Implementation for seeding products
    return nil
}
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific test package
go test ./internal/service/user/
```

### Testing Patterns

#### Unit Tests

```go
// internal/service/user/service_test.go
func TestUserService_CreateUser(t *testing.T) {
    // Mock repository
    mockRepo := &mocks.UserRepo{}

    // Create service with mock
    service := user.NewService(mockRepo)

    // Test implementation
    // ...
}
```

#### Integration Tests

```go
// internal/infrastructure/database/gorm/repo/user_repo_test.go
func TestUserRepo_Integration(t *testing.T) {
    // Use test database
    db := setupTestDB(t)
    repo := NewUserRepo(db)

    // Test with real database
    // ...
}
```

### Concurrent Testing

Test optimistic locking behavior:

```bash
# Test concurrent updates
go run cmd/concurrent/concurrent_user_update.go
```

## 🔧 Configuration Management

### Environment Variables

Update `.envrc` files for different environments:

```bash
# .envrc
export DATABASE_URL="postgres://test_user:test_password@localhost:5433/goWebapp_test?sslmode=disable"
export ENV="dev"

# .envrc.test
export DATABASE_URL="postgres://test_user:test_password@localhost:5433/goWebapp_test?sslmode=disable"
export ENV="test"
```

### Configuration Files

Environment-specific configurations in `configs/`:

- `app-dev.yaml` - Development settings
- `app-prod.yaml` - Production settings

## 📊 Monitoring & Debugging

### Logging

Configure logging levels and output:

```yaml
# configs/app-dev.yaml
logging:
  level: debug
  writer: stdout
  use_json: false
  add_source: true
  source_key: 'source'
```

### Tracing

View traces in Jaeger UI:

```bash
# Access Jaeger UI
open http://localhost:16686
```

### Profiling

Access debug profiler:

```bash
# Development profiler endpoint
open http://localhost:8082/debug
```

## 🚀 Deployment

### Building

```bash
# Build for production
make build

# Build with specific environment
make build ENV=prod
```

### Docker

```bash
# Build Docker image
docker build -t go-web-app .

# Run with Docker Compose
docker-compose up -d
```

## 📚 Best Practices

### Code Organization

- Follow clean architecture principles
- Keep domain logic pure and infrastructure-agnostic
- Use interfaces for dependency injection
- Organize by business domains rather than technical concerns

### Error Handling

- Use domain-specific error types
- Implement proper error wrapping
- Log errors with context
- Return appropriate HTTP status codes

### Testing

- Write tests for all business logic
- Use mocks for external dependencies
- Test error scenarios
- Maintain high test coverage

### Performance

- Use database transactions appropriately
- Implement proper connection pooling
- Monitor slow queries
- Use optimistic locking for concurrent updates

### Security

- Validate all inputs
- Use parameterized queries
- Implement proper CORS policies
- Configure rate limiting
