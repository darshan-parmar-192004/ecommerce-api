# Ecommerce API Backend

A production-ready RESTful e-commerce backend API built with Go Fiber, featuring PostgreSQL for data persistence, Redis for caching and session management, and JWT-based authentication.

## Business Context

This API powers an e-commerce platform enabling:
- Product catalog management with categories
- Customer registration and authentication
- Order processing and tracking
- Inventory management with analytics
- Role-based access control (customers, admins)

## Features Implemented

### Authentication & Authorization
- JWT-based authentication with HS256 signing
- Redis session storage for token revocation
- Password hashing with bcrypt
- Role-based access control (customer, admin)
- HTTP-only secure cookies + Bearer token support

### Products
- Full CRUD operations
- Pagination, filtering, and search
- Redis caching (5 min TTL)

### Categories
- Hierarchical parent-child structure
- Category tree retrieval
- Products by category

### Customers
- Registration and profile management
- Order history retrieval
- Lifetime value calculation

### Orders
- Order creation with items
- Ownership-validated retrieval

### Inventory (Admin)
- Stock tracking by warehouse
- Customer lifetime value analytics
- Top sellers and category tree reports

## Technology Stack

| Component | Technology |
|-----------|------------|
| Framework | Go Fiber v3 |
| Database | PostgreSQL 16 |
| Cache | Redis 7 |
| Auth | JWT (HS256) + bcrypt |
| Container | Docker/Podman Compose |
| Testing | Go test framework |

## Prerequisites

- Go 1.21+
- Docker or Podman
- PostgreSQL 16+ (or use container)
- Redis 7+ (or use container)

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api/backend

# Install dependencies
go mod download
```

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | API server port | `8080` |
| `APP_ENV` | Environment (development/production) | `development` |
| `JWT_SECRET` | Secret key for JWT signing | (required in production) |
| `BLUEPRINT_DB_HOST` | PostgreSQL host | `localhost` |
| `BLUEPRINT_DB_PORT` | PostgreSQL port | `5432` |
| `BLUEPRINT_DB_DATABASE` | Database name | `ecommerce` |
| `BLUEPRINT_DB_USERNAME` | Database user | `postgres` |
| `BLUEPRINT_DB_PASSWORD` | Database password | (required) |
| `BLUEPRINT_DB_SCHEMA` | Database schema | `public` |
| `REDIS_HOST` | Redis host | `localhost` |
| `REDIS_PORT` | Redis port | `6379` |

### Example .env file

```bash
PORT=8080
APP_ENV=development
JWT_SECRET=your-secret-key-change-in-production

BLUEPRINT_DB_HOST=localhost
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_DATABASE=ecommerce
BLUEPRINT_DB_USERNAME=postgres
BLUEPRINT_DB_PASSWORD=your-password
BLUEPRINT_DB_SCHEMA=public

REDIS_HOST=localhost
REDIS_PORT=6379
```

## Running Locally (Development)

### Option 1: With Docker/Podman (Recommended)

```bash
# Start PostgreSQL and Redis containers
cd ..
docker-compose up -d psql_bp redis

# Or use podman-compose
podman-compose up -d psql_bp redis
```

### Option 2: Local Database

Ensure PostgreSQL and Redis are running locally, then:

```bash
# Run database migrations
cd ../backend
go run cmd/migrate/main.go up

# Start the API server
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

### Using Make Commands

```bash
# Build and test
make all

# Build binary
make build

# Run application
make run

# Run tests
make test

# Run integration tests (requires database)
make itest

# Live reload during development
make watch

# Clean build artifacts
make clean
```

## Running with Podman/Docker

```bash
# From project root
cd ..
docker-compose up --build

# Or with podman
podman-compose up --build
```

This starts:
- API server on port 8080
- PostgreSQL on port 5432
- Redis on port 6379

## Running Tests

```bash
# All tests
make test

# Integration tests (requires database)
make itest

# Specific test
go test ./internal/test/unit/product_unit_test.go -v
go test -run TestProduct ./... -v
```

## API Documentation

The OpenAPI 3.0 specification is available at:
- `/docs/openapi.yaml` - OpenAPI spec file
- Swagger UI: `http://localhost:8080/api/docs` (if enabled)

### Authentication

Include JWT token in Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

Or use the `auth_token` cookie (set automatically on login).

### Rate Limiting

- Login endpoint: 5 requests/minute per IP
- Other endpoints: 100 requests/minute

### Example Requests

```bash
# Register
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Password123!","name":"John"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Password123!"}'

# Get products (with token)
curl -X GET http://localhost:8080/products \
  -H "Authorization: Bearer <token>"
```

## Deployment

### Production Checklist

1. Set `APP_ENV=production`
2. Generate strong `JWT_SECRET`
3. Configure PostgreSQL with SSL
4. Enable Redis password
5. Set up reverse proxy (nginx)
6. Configure SSL/TLS
7. Set up monitoring and logging

### Container Deployment

```bash
# Build and run
docker-compose up -d --build

# Check logs
docker-compose logs -f app

# Stop services
docker-compose down
```

### Health Check

```bash
curl http://localhost:8080/health
```

## Contributing

Contributions are welcome! Please see the root [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

### Quick Start

```bash
# Fork and clone
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api/backend

# Create branch
git checkout -b feature/your-feature

# Make changes and test
make test

# Commit and push
git commit -m "Add: your feature description"
git push origin feature/your-feature
```

## License

MIT License - see [LICENSE](../LICENSE) for details.

## Related Documentation

- [API Design](./docs/architecture/api-design.md)
- [Database Schema](./docs/architecture/database-schema.md)
- [System Overview](./docs/architecture/system-overview.md)
- [OpenAPI Spec](./docs/openapi.yaml)
- [Deployment Runbook](./docs/runbooks/deployment.md)
- [ADRs](./docs/adrs/)