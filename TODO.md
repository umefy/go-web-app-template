# TODO - Future Enhancements & Tasks

This file tracks planned features and tasks for the Go Web App Template, organized by priority.

## 🚨 High Priority

### Core Features

- [ ] **API Pagination**: Add pagination support to getAll users endpoint and other list endpoints
- [ ] **Authentication & Authorization**: JWT, OAuth2, RBAC

### Infrastructure

- [ ] **Security**: Security headers, enhanced CORS configuration, security scanning
- [ ] **Performance**: Connection pooling, caching strategies, performance testing
- [ ] **Health Checks**: Add more comprehensive health check endpoints

## 🔶 Medium Priority

### API & Data

- [ ] **API Rate Limiting**: Per-user rate limiting and advanced throttling
- [ ] **API Caching**: Response caching with ETags and conditional requests

### Development Experience

- [ ] **Testing**: Add more unit tests for edge cases
- [ ] **Development Tools**: Additional development utilities and scripts

### Observability

- [ ] **Metrics**: Prometheus integration, enhanced OpenTelemetry metrics
- [ ] **Tracing**: Custom trace attributes, sampling strategies, trace correlation
- [ ] **Monitoring**: Enhanced health checks, metrics dashboard, alerting

## 🔵 Low Priority

### Features

- [ ] **Caching Layer**: Redis integration
- [ ] **Event System**: Domain events and messaging
- [ ] **Background Jobs**: Task queue integration
- [ ] **File Upload**: Multipart file handling
- [ ] **Email Integration**: SMTP/email service

### Infrastructure

- [ ] **Docker Support**: Multi-stage builds, production Docker images
- [ ] **Kubernetes**: Deployment manifests, Helm charts
- [ ] **CI/CD**: Additional GitHub Actions workflows

### Development

- [ ] **Documentation**: Add more code examples and tutorials
- [ ] **Performance**: Add performance benchmarks
- [ ] **Load Testing**: Automated load testing and performance benchmarks

## 🔒 Optimistic Locking Enhancements

- [ ] **Conflict Resolution**: Automatic retry mechanisms, conflict resolution policies
- [ ] **Version History**: Track version changes for audit purposes
- [ ] **Performance Monitoring**: Metrics on lock conflicts and resolution times

## 🌱 Database Seeding Enhancements

- [ ] **Incremental Seeding**: Add data without clearing existing data
- [ ] **Performance Testing Data**: Large datasets for performance testing

---

## 📝 Notes

- **High Priority**: Must-have features for production readiness
- **Medium Priority**: Important improvements for better user experience
- **Low Priority**: Nice-to-have features for future releases
- **Focus**: Start with High Priority items and work your way down
