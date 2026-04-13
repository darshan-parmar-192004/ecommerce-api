# Ecommerce API

A production-ready RESTful e-commerce backend API built with Go Fiber, featuring PostgreSQL for data persistence, Redis for caching and session management, and JWT-based authentication.

## Project Overview

This is a complete e-commerce backend API that provides:
- **Product Management** - CRUD operations for products with category associations
- **Category System** - Hierarchical categories with tree navigation
- **Customer Accounts** - User registration, authentication, and profile management
- **Order Processing** - Order creation with ownership validation
- **Inventory Management** - Stock tracking, warehouse management, and analytics
- **Analytics** - Top sellers, customer lifetime value, category hierarchies
- **Caching** - Redis-powered query caching for improved performance

## Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.21+ |
| **Web Framework** | Fiber v3 |
| **Database** | PostgreSQL 16+ |
| **Cache/Session** | Redis 7+ |
| **ORM/Query Builder** | goqu v9 |
| **Authentication** | JWT (golang-jwt/jwt/v5) |
| **Password Hashing** | bcrypt |
| **Migrations** | golang-migrate |
| **Frontend** | React + Vite + Tailwind (optional) |

## Features Implemented

### Authentication & Authorization
- JWT-based stateless authentication
- Redis session storage
- Role-based access control (customer, admin)
- Password validation (min 8 chars, uppercase, lowercase, digit, special char)
- Rate limiting on login endpoints

### API Endpoints

#### Products (Public Read, Admin Write)
```
GET    /products              - List all products (paginated, filtered)
GET    /products/:id          - Get product by ID
POST   /products               - Create product (admin)
PUT    /products/:id          - Update product (admin)
DELETE /products/:id          - Delete product (admin)
```

#### Categories (Public)
```
GET    /categories                        - List all categories
GET    /categories/:id/products           - Get products in category
GET    /categories/hierarchy              - Get category tree
```

#### Customers (Authenticated)
```
GET    /customers/me           - Get current user profile
PUT    /customers/me           - Update current user profile
GET    /customers/:id/orders   - Get customer orders
GET    /customers/:id/lifetime-value - Get customer CLV (admin)
```

#### Orders (Authenticated)
```
GET    /orders/:id            - Get order (ownership validated)
POST   /orders                - Create new order
```

#### Inventory (Admin Only)
```
GET    /inventory                     - Get all inventory
GET    /inventory/stock               - Get stock levels
GET    /inventory/customer-lifetime-value - Get CLV stats
GET    /inventory/hierarchy           - Get category tree
GET    /inventory/top-sellers         - Get top selling products
```

#### Authentication
```
POST   /auth/register     - Register new user
POST   /auth/login        - Login (rate limited)
POST   /auth/logout       - Logout
GET    /auth/me           - Validate token
```

#### Health & Stats
```
GET    /health            - Health check endpoint
GET    /stats/cache       - Cache statistics
```

## Prerequisites

- **Go**: 1.21 or higher
- **PostgreSQL**: 16 or higher
- **Redis**: 7 or higher
- **Node.js**: 18+ (for frontend development)
- **Make**: For build automation
- **Podman** or **Docker**: For containerized deployment

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```

### 2. Install Go Dependencies

```bash
cd backend
go mod download
```

### 3. Configure Environment

Copy the example environment file and configure your settings:

```bash
cp .env.example .env
```

### 4. Database Setup

#### Option A: Using Podman/Docker Compose (Recommended)

```bash
# Start all services (PostgreSQL, Redis, App)
podman-compose up -d
# Or with Docker
docker-compose up -d
```

#### Option B: Local Development

1. Install PostgreSQL 16+ and Redis 7+
2. Create a database and user
3. Run migrations:

```bash
cd backend
go run cmd/api/main.go migrate
# Or manually
migrate -path internal/migrations -database "postgres://user:pass@localhost:5432/ecommerce?sslmode=disable" up
```

## Configuration

Configure the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Application environment | `development` |
| `PORT` | Server port | `8080` |
| `BLUEPRINT_DB_HOST` | PostgreSQL host | `localhost` |
| `BLUEPRINT_DB_PORT` | PostgreSQL port | `5432` |
| `BLUEPRINT_DB_DATABASE` | Database name | `ecommerce` |
| `BLUEPRINT_DB_USERNAME` | Database user | - |
| `BLUEPRINT_DB_PASSWORD` | Database password | - |
| `BLUEPRINT_DB_SCHEMA` | Database schema | `public` |
| `REDIS_HOST` | Redis host | `localhost` |
| `REDIS_PORT` | Redis port | `6379` |
| `JWT_SECRET` | JWT signing secret (min 32 chars) | - |

## Running Locally (Development)

### Backend

```bash
cd backend

# Run with air (live reload)
make watch

# Or run directly
go run cmd/api/main.go

# Run tests
make test

# Run integration tests (requires database)
make itest
```

The API will be available at `http://localhost:8080`

### Frontend (Optional)

```bash
# React frontend
cd frontend
npm install
npm run dev

# Or Nuxt 3 frontend (experimental)
cd frontend-nuxt
npm install
npm run dev
```

## Running with Podman Compose

### Start Services

```bash
# Build and start all containers
podman-compose up -d --build

# View logs
podman-compose logs -f app

# Stop services
podman-compose down
```

### Services Available
- **API**: http://localhost:8080
- **PostgreSQL**: localhost:5432
- **Redis**: localhost:6379
- **Swagger UI**: http://localhost:8080/api/docs

## Running Tests

```bash
cd backend

# Run all tests
make test

# Run unit tests only
go test ./internal/test/unit/... -v

# Run integration tests (requires database)
make itest

# Run specific test
go test -run TestProduct ./... -v
```

## API Documentation

Interactive API documentation is available at:

```
http://localhost:8080/api/docs
```

This provides Swagger UI with:
- Full endpoint descriptions
- Request/response schemas
- Authentication setup
- Try-it-out functionality

## Deployment

### Prerequisites
- Server with at least 2 CPU cores, 4GB RAM
- PostgreSQL 16+ installed
- Redis 7+ installed
- Open ports: 8080 (API), 5432 (PostgreSQL), 6379 (Redis)

### Steps

1. **Prepare the server**
   ```bash
   # Install dependencies
   sudo apt update && sudo apt install -y postgresql-16 redis-server golang-go
   
   # Create database
   sudo -u postgres createdb ecommerce
   sudo -u postgres createuser ecommerce_user
   sudo -u postgres grant all privileges on database ecommerce to ecommerce_user
   ```

2. **Deploy the application**
   ```bash
   # Clone and build
   git clone <repository>
   cd ecommerce-api/backend
   go build -o main ./cmd/api/main.go
   
   # Set environment variables
   export BLUEPRINT_DB_HOST=localhost
   export BLUEPRINT_DB_DATABASE=ecommerce
   export BLUEPRINT_DB_USERNAME=ecommerce_user
   export BLUEPRINT_DB_PASSWORD=your_secure_password
   export JWT_SECRET=your-32-character-minimum-secret-key
   export REDIS_HOST=localhost
   
   # Run migrations
   ./main migrate
   
   # Start the server
   ./main
   ```

3. **Or use containers**
   ```bash
   podman-compose up -d
   ```

### Health Checks

```bash
# Check API health
curl http://localhost:8080/health

# Expected response
{"status": "ok"}
```

## Project Structure

```
ecommerce-api/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go           # Application entry point
│   ├── internal/
│   │   ├── auth/                  # JWT authentication
│   │   ├── cache/                 # Redis caching
│   │   ├── category/             # Category service
│   │   ├── controllers/           # HTTP handlers
│   │   ├── customer/              # Customer service
│   │   ├── database/              # Database connection
│   │   ├── errors/                # Error handling
│   │   ├── inventory/             # Inventory service
│   │   ├── middleware/            # Fiber middleware
│   │   ├── migrations/            # SQL migrations
│   │   ├── models/                # Data models
│   │   ├── order/                 # Order service
│   │   ├── product/               # Product service
│   │   ├── querybuilder/          # SQL query builder
│   │   ├── server/                # Server setup
│   │   ├── services/              # Business logic
│   │   └── test/                  # Tests
│   ├── Dockerfile
│   ├── Makefile
│   └── go.mod
├── frontend/                      # React + Vite
├── frontend-nuxt/                 # Nuxt 3 (experimental)
├── docs/                          # Documentation
│   ├── architecture/              # Architecture docs
│   ├── adrs/                     # ADRs
│   └── runbooks/                 # Operational runbooks
├── docker-compose.yml
├── .env.example
├── README.md
├── CHANGELOG.md
├── CONTRIBUTING.md
└── SECURITY.md
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) for details.

## Support

- Open an issue for bugs or feature requests
- Check the [documentation](docs/) for detailed guides
- Review the [API specification](http://localhost:8080/api/docs) for endpoint details