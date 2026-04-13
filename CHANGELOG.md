# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0-alpha] - 2024-01-15

### Added
- **Authentication**
  - JWT-based authentication with HS256 signing
  - Redis session storage for token revocation
  - Password hashing with bcrypt (cost 10)
  - Password validation (min 8 chars, uppercase, lowercase, digit, special)
  - Role-based access control (customer, admin)
  - HTTP-only, secure cookies for web clients
  - Bearer token support for API clients

- **Products**
  - CRUD operations (Create, Read, Update, Delete)
  - Pagination support (page, limit)
  - Filtering (category, price range, search)
  - Redis caching (5 min TTL)

- **Categories**
  - Hierarchical category system (parent-child)
  - Category tree retrieval
  - Products by category endpoint

- **Customers**
  - User registration and profile management
  - Customer orders retrieval
  - Customer lifetime value calculation

- **Orders**
  - Order creation with items
  - Order retrieval with ownership validation
  - Order total calculation

- **Inventory**
  - Stock tracking by warehouse
  - Customer lifetime value statistics
  - Top sellers analytics
  - Category tree for inventory

- **Infrastructure**
  - PostgreSQL database with migrations
  - Redis for caching and sessions
  - Container support (Docker/Podman)
  - Health check endpoint
  - Cache statistics endpoint

- **API Features**
  - Rate limiting (general + login-specific)
  - CORS support
  - Request logging
  - Security headers
  - Recovery middleware

### Documentation
- OpenAPI 3.0 specification
- Swagger UI at /api/docs
- Architecture documentation (system overview, database schema, API design)
- Architecture Decision Records (ADRs)
- Operational runbooks (deployment, backup, troubleshooting, scaling)

## [0.1.0] - 2023-12-01

### Added
- Initial project setup
- Basic Go Fiber server
- PostgreSQL connection
- Redis connection
- Docker compose configuration
- Basic health check endpoint

### Migration Guide

#### From 0.1.0 to 1.0.0-alpha

1. **Database Migrations**: Run new migrations to add roles and password_hash columns
2. **Environment**: Add JWT_SECRET to environment variables (min 32 characters)
3. **Authentication**: Update client code to handle JWT tokens in cookies or Authorization header
4. **Dependencies**: No breaking changes to API contract

## Version History

| Version | Date | Status |
|---------|------|--------|
| 1.0.0-alpha | 2024-01-15 | Current |
| 0.1.0 | 2023-12-01 | Deprecated |

## Deprecation Notices

None at this time.

## Security Vulnerabilities

None at this time.