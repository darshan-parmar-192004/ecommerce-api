# Deployment Runbook

## Prerequisites

### Server Requirements
| Resource | Minimum | Recommended |
|----------|---------|-------------|
| CPU | 2 cores | 4+ cores |
| RAM | 4 GB | 8+ GB |
| Disk | 20 GB | 50+ GB |
| OS | Ubuntu 22.04+ | Ubuntu 22.04+ |

### Required Ports
| Port | Service | Purpose |
|------|---------|---------|
| 8080 | API | HTTP API |
| 5432 | PostgreSQL | Database (internal) |
| 6379 | Redis | Cache (internal) |

### Required Software
- Go 1.21+ (for building)
- PostgreSQL 16+
- Redis 7+
- Podman or Docker

## Environment Setup

### 1. Install Dependencies

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install -y postgresql-16 redis-server golang-go podman docker.io

# Verify installations
go version
psql --version
redis-server --version
podman --version
```

### 2. Configure PostgreSQL

```bash
# Start PostgreSQL
sudo systemctl enable postgresql
sudo systemctl start postgresql

# Create database and user
sudo -u postgres createuser -s ecommerce_user
sudo -u postgres createdb ecommerce
sudo -u postgres psql -c "ALTER USER ecommerce_user WITH PASSWORD 'your_secure_password'"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE ecommerce TO ecommerce_user"
```

### 3. Configure Redis

```bash
# Start Redis
sudo systemctl enable redis-server
sudo systemctl start redis-server

# Verify
redis-cli ping
# Expected: PONG
```

## Application Deployment

### Option A: Binary Deployment

```bash
# 1. Clone repository
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api/backend

# 2. Build binary
go build -o ecommerce-api ./cmd/api/main.go

# 3. Set environment variables
export APP_ENV=production
export PORT=8080
export BLUEPRINT_DB_HOST=localhost
export BLUEPRINT_DB_PORT=5432
export BLUEPRINT_DB_DATABASE=ecommerce
export BLUEPRINT_DB_USERNAME=ecommerce_user
export BLUEPRINT_DB_PASSWORD=your_secure_password
export JWT_SECRET=your-32-character-minimum-secret-key
export REDIS_HOST=localhost
export REDIS_PORT=6379

# 4. Run migrations
./ecommerce-api migrate

# 5. Start the server
./ecommerce-api

# Or run in background with systemd
sudo tee /etc/systemd/system/ecommerce-api.service > /dev/null <<EOF
[Unit]
Description=Ecommerce API
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/ecommerce-api
EnvironmentFile=/opt/ecommerce-api/.env
ExecStart=/opt/ecommerce-api/ecommerce-api
Restart=always

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable ecommerce-api
sudo systemctl start ecommerce-api
```

### Option B: Container Deployment (Recommended)

```bash
# 1. Configure environment
cp .env.example .env
# Edit .env with production values

# 2. Build and start containers
podman-compose up -d --build

# Or with Docker
docker-compose up -d --build

# 3. Verify containers are running
podman-compose ps
# Expected: app, psql_bp, redis - all "Up"

# 4. Check logs
podman-compose logs -f app

# 5. Verify health
curl http://localhost:8080/health
# Expected: {"status": "ok"}
```

## Database Migrations

### Running Migrations

```bash
# Using the binary
./ecommerce-api migrate

# Using golang-migrate directly
migrate -path internal/migrations \
  -database "postgres://ecommerce_user:password@localhost:5432/ecommerce?sslmode=disable" up

# Using Docker
podman-compose run migrate
```

### Migration Commands
| Command | Description |
|---------|-------------|
| `migrate up` | Apply all pending migrations |
| `migrate down` | Rollback last migration |
| `migrate force VERSION` | Force specific version |

## Health Checks

### API Health
```bash
curl http://localhost:8080/health
# Expected: {"status": "ok"}
```

### Database Connection
```bash
# PostgreSQL
psql -h localhost -U ecommerce_user -d ecommerce -c "SELECT 1;"

# Redis
redis-cli ping
# Expected: PONG
```

### Cache Stats
```bash
curl http://localhost:8080/stats/cache
```

## Verification Checklist

- [ ] API responds on port 8080
- [ ] Health endpoint returns 200
- [ ] Database connection works
- [ ] Redis connection works
- [ ] Can register a new user
- [ ] Can login and receive JWT
- [ ] Can access protected endpoints with token

## Troubleshooting

If deployment fails:

1. **Check logs**: `podman-compose logs -f app`
2. **Verify environment**: Ensure all .env variables are set
3. **Check ports**: `ss -tlnp | grep -E '8080|5432|6379'`
4. **Database connectivity**: `psql -h localhost -U ecommerce_user -d ecommerce`
5. **Redis connectivity**: `redis-cli ping`

## Rollback Procedure

If issues occur after deployment:

```bash
# Binary: Stop service and revert to previous version
sudo systemctl stop ecommerce-api
# Revert binary and restart

# Container: Rollback to previous image
podman-compose down
git checkout previous-tag
podman-compose up -d --build
```