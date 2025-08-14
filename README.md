# Go Web App Template

[![Go Report Card](https://goreportcard.com/badge/github.com/umefy/go-web-app-template)](https://goreportcard.com/report/github.com/umefy/go-web-app-template)
[![Go Version](https://img.shields.io/github/go-mod/go-version/umefy/go-web-app-template)](https://go.dev/)
[![License](https://img.shields.io/github/license/umefy/go-web-app-template)](LICENSE)

A production-ready Go web application template with clean architecture, multi-protocol APIs, and comprehensive tooling.

## 🚀 Quick Start

1. **Setup Environment**

   ```bash
   ./scripts/local_setup.sh
   ```

   **Note**: Create a `.envrc` file based on `.envrc.example` to enable git hash injection for version tracking.

2. **Start Services**

   ```bash
   make docker_compose_up
   ```

3. **Generate Code**

   ```bash
   make generate
   ```

4. **Run Migrations**

   ```bash
   make migration_up
   ```

5. **Start Development**

   ```bash
   make dev
   ```

**Access your app:**

- REST API: `http://localhost:8082/api/v1/`
- GraphQL: `http://localhost:8082/graphql/`
- Health check: `http://localhost:8082/health`
- Jaeger UI: `http://localhost:16686`

## ✨ Key Features

- **🏗️ Clean Architecture** with FX dependency injection
- **🌐 Multi-Protocol APIs** - HTTP (REST + GraphQL) + gRPC
- **🗄️ Database** - GORM + migrations + optimistic locking
- **📊 Observability** - OpenTelemetry + structured logging
- **🐳 Development** - Docker Compose + seeding + testing tools
- **🔒 Security** - Input validation + rate limiting + CORS
- **🧪 Testing** - Comprehensive testing + mocking + concurrent testing

## 📁 Project Structure

```bash
cmd/
└── server/         # Application entry point & FX composition

internal/
├── domain/          # Business logic & interfaces
├── service/         # Business logic implementation
├── delivery/        # HTTP/GraphQL/gRPC handlers
├── infrastructure/  # Database, server, logging
└── core/           # Configuration & shared code
```

## 📚 Documentation

- **[🏗️ Architecture Guide](docs/ARCHITECTURE.md)** - Clean architecture & FX dependency injection
- **[🛠️ Development Guide](docs/DEVELOPMENT.md)** - Workflow & API development
- **[✨ Features Guide](docs/FEATURES.md)** - Detailed feature explanations
- **[📖 API Reference](openapi/docs/api.yaml)** - OpenAPI specification

## 🛠️ Development Commands

```bash
make dev              # Start development server
make generate         # Generate all code
make test            # Run tests
make seed_database   # Seed with sample data
make migration_up    # Run database migrations
make docker_compose_up   # Start infrastructure services
```

## 🔧 Configuration

Environment-specific configurations in `configs/`:

- `app-dev.yaml` - Development settings
- `app-prod.yaml` - Production settings

**Git Hash Injection**: The application automatically injects git commit hashes into configuration values marked with `""` (empty string) for automatic version tracking.

## 🐳 Infrastructure Services

- **PostgreSQL**: `localhost:5433`
- **Jaeger**: `http://localhost:16686`

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Need help?** Check the [documentation](docs/) or [open an issue](https://github.com/umefy/go-web-app-template/issues).
