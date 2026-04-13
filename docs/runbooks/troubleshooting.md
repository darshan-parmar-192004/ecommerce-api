# Troubleshooting Runbook

## Common Issues and Solutions

### API Issues

#### 1. API Not Responding

**Symptoms:**
- `curl http://localhost:8080/health` fails
- Connection timeout

**Diagnosis:**
```bash
# Check if API is running
podman ps | grep backend
# Or: systemctl status ecommerce-api

# Check port is listening
ss -tlnp | grep 8080

# Check logs
podman-compose logs app
# Or: journalctl -u ecommerce-api -n 100
```

**Solutions:**
```bash
# If not running, start it
podman-compose up -d
# Or: systemctl start ecommerce-api

# If crashed, check for port conflicts
lsof -i :8080
kill -9 <pid> if conflicting process

# Check environment variables
env | grep -E 'BLUEPRINT|REDIS|JWT'
```

#### 2. 500 Internal Server Error

**Symptoms:**
- API returns 500 status
- Error in logs

**Diagnosis:**
```bash
# Check application logs
podman-compose logs app | grep -i error

# Look for panic messages
podman-compose logs app | grep -i panic

# Check database connection
psql -h localhost -U ecommerce_user -d ecommerce -c "SELECT 1;"
```

**Solutions:**
- Check database credentials in environment
- Verify database is accessible
- Check for database migrations not applied

#### 3. Authentication Failures

**Symptoms:**
- 401 Unauthorized on protected endpoints
- Token validation errors

**Diagnosis:**
```bash
# Check JWT secret
echo $JWT_SECRET
# Should be at least 32 characters

# Check Redis for session
redis-cli GET "session:<customer_id>"

# Verify token
curl -v -H "Authorization: Bearer <token>" http://localhost:8080/auth/me
```

**Solutions:**
- Ensure JWT_SECRET is set and consistent
- Clear Redis sessions: `redis-cli FLUSHDB`
- Re-login to get new token

### Database Issues

#### 1. Database Connection Refused

**Symptoms:**
- `pq: connection refused`
- Cannot connect to PostgreSQL

**Diagnosis:**
```bash
# Check PostgreSQL is running
podman-compose ps psql_bp
# Or: systemctl status postgresql

# Check port
ss -tlnp | grep 5432

# Test connection
psql -h localhost -U ecommerce_user -d ecommerce
```

**Solutions:**
```bash
# Start PostgreSQL
podman-compose up -d psql_bp
# Or: systemctl start postgresql

# Check credentials in .env
# BLUEPRINT_DB_HOST, BLUEPRINT_DB_PORT, etc.
```

#### 2. Database Authentication Failed

**Symptoms:**
- `pq: password authentication failed`

**Diagnosis:**
```bash
# Check credentials match
grep BLUEPRINT_DB /etc/ecommerce-api/.env

# Test with psql directly
psql -h localhost -U <username> -d ecommerce -W
```

**Solutions:**
- Update password in .env file
- Update PostgreSQL user password:
  ```sql
  ALTER USER ecommerce_user WITH PASSWORD 'new_password';
  ```
- Update .env to match

#### 3. Database Does Not Exist

**Symptoms:**
- `pq: database "ecommerce" does not exist`

**Solutions:**
```bash
# Create database
sudo -u postgres createdb ecommerce
sudo -u postgres createuser ecommerce_user
sudo -u postgres psql -c "GRANT ALL ON DATABASE ecommerce TO ecommerce_user;"

# Or with podman-compose (auto-creates)
podman-compose up -d psql_bp
```

#### 4. Too Many Connections

**Symptoms:**
- `pq: sorry, too many clients already`

**Diagnosis:**
```bash
# Check current connections
psql -h localhost -U ecommerce_user -d ecommerce -c "SELECT count(*) FROM pg_stat_activity;"
```

**Solutions:**
- Wait for connections to close
- Restart API to reset connection pool
- Reduce MaxOpenConnections in code/config

### Redis Issues

#### 1. Redis Connection Refused

**Symptoms:**
- Cannot connect to Redis

**Diagnosis:**
```bash
# Check Redis is running
podman-compose ps redis

# Test connection
redis-cli ping
# Expected: PONG
```

**Solutions:**
```bash
# Start Redis
podman-compose up -d redis

# Or locally
redis-server --daemonize yes
```

#### 2. Redis Out of Memory

**Symptoms:**
- Redis error: OOM command not allowed

**Solutions:**
```bash
# Check Redis memory
redis-cli INFO memory

# Clear cache (not sessions!)
redis-cli FLUSHDB

# Or set maxmemory in redis.conf
# maxmemory 256mb
```

### Container Issues

#### 1. Container Won't Start

**Diagnosis:**
```bash
# Check logs
podman-compose logs app

# Check for build errors
podman-compose build --no-cache
```

**Solutions:**
- Check .env file exists and has all required variables
- Rebuild: `podman-compose build --no-cache`
- Check port conflicts

#### 2. Volume Mount Issues

**Diagnosis:**
```bash
# Check volumes
podman volume ls

# Inspect volume
podman volume inspect ecommerce-api_psql_volume_bp
```

**Solutions:**
- Recreate volume: `podman volume rm <volume>`
- Check mount paths in docker-compose.yml

### Performance Issues

#### 1. Slow API Responses

**Diagnosis:**
```bash
# Check response time
time curl http://localhost:8080/products

# Check cache stats
curl http://localhost:8080/stats/cache

# Check database queries
podman-compose logs app | grep -i query
```

**Solutions:**
- Increase Redis cache TTL
- Add database indexes
- Check for N+1 query problems

#### 2. High CPU Usage

**Diagnosis:**
```bash
# Check process
top -p $(pgrep -f ecommerce-api)

# Check for infinite loops in logs
podman-compose logs | grep -i loop
```

**Solutions:**
- Restart API
- Check for runaway queries
- Profile the application

## Debugging Tips

### Enable Debug Logging

```bash
# Set log level via environment
export LOG_LEVEL=debug

# In docker-compose.yml add:
# environment:
#   - LOG_LEVEL=debug
```

### View Request Logs

```bash
# All requests logged
podman-compose logs app | grep "Request"

# Filter by endpoint
podman-compose logs app | grep "/products"
```

### Test API Endpoints

```bash
# Health check
curl http://localhost:8080/health

# Get products (public)
curl http://localhost:8080/products

# Register user
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test@1234","name":"Test User"}'

# Login
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"Test@1234"}'
```

### Database Debugging

```bash
# Connect to database
psql -h localhost -U ecommerce_user -d ecommerce

# List tables
\dt

# Check table sizes
SELECT relname, pg_size_pretty(pg_total_relation_size(relid))
FROM pg_catalog.pg_statio_user_tables
ORDER BY pg_total_relation_size(relid) DESC;

# Check indexes
SELECT indexrelname, idx_scan
FROM pg_stat_user_indexes
ORDER BY idx_scan DESC;

# Check slow queries (requires pg_stat_statements)
SELECT query, calls, mean_time 
FROM pg_stat_statements 
ORDER BY mean_time DESC 
LIMIT 10;
```

### Redis Debugging

```bash
# List all keys
redis-cli KEYS "*"

# Check session keys
redis-cli KEYS "session:*"

# Check cache keys
redis-cli KEYS "products:*"

# Monitor Redis
redis-cli MONITOR

# Check memory
redis-cli MEMORY STATS
```

## Emergency Procedures

### Complete System Restart

```bash
# Stop everything
podman-compose down

# Clear caches (optional)
redis-cli FLUSHDB

# Start everything fresh
podman-compose up -d --build

# Verify
curl http://localhost:8080/health
```

### Rollback to Previous Version

```bash
# Stop current
podman-compose down

# Checkout previous version
git checkout <previous-commit>

# Rebuild and start
podman-compose up -d --build
```

### Emergency Database Restore

See [Backup and Restore](backup-restore.md)

## Getting Help

If issue persists:
1. Collect all logs: `podman-compose logs > logs.txt`
2. Document steps to reproduce
3. Check GitHub issues
4. Create new issue with logs attached