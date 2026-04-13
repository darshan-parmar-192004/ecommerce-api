# Scaling Runbook

## Overview

This runbook covers strategies for scaling the Ecommerce API to handle increased load.

## Scaling Dimensions

| Dimension | Strategy | Complexity |
|-----------|----------|------------|
| Horizontal | More API instances | Medium |
| Database | Read replicas + partitioning | High |
| Cache | Redis Cluster | Medium |
| CDN | Static assets | Low |

## Horizontal Scaling

### Current Architecture (Single Instance)

```
[Client] -> [API:8080] -> [PostgreSQL:5432]
                |
                v
              [Redis:6379]
```

### Scaled Architecture

```
                    [Load Balancer]
                           |
      +--------------------+--------------------+
      |                    |                    |
[API:8080:1]          [API:8080:2]         [API:8080:n]
      |                    |                    |
      +--------------------+--------------------+
              |            |            |
              v            v            v
        [PostgreSQL:5432] [Redis:6379]
```

### Implementation

```yaml
# docker-compose.yml with multiple API instances
services:
  api:
    build: ./backend
    deploy:
      replicas: 3
    # ... other config

  # Or use Kubernetes
  # api-deployment.yaml
  api:
    replicas: 3
```

### Session Sharing

When running multiple API instances, sessions must be shared via Redis:

```go
// Current: Each instance uses local Redis (works for single instance)
// For horizontal scaling: Ensure all instances use same Redis
// Already configured - uses REDIS_HOST env var
```

### Load Balancer Options

| Option | Description |
|--------|-------------|
| Nginx | Simple, proven |
| HAProxy | High performance |
| Traefik | Automatic service discovery |
| Cloud LB | AWS ALB, GCP Cloud LB |

### Nginx Load Balancer Config

```nginx
upstream api_backend {
    server api_1:8080;
    server api_2:8080;
    server api_3:8080;
}

server {
    listen 80;
    server_name api.example.com;

    location / {
        proxy_pass http://api_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## Database Scaling

### Read Replicas

For read-heavy workloads:

```sql
-- Create read replica (PostgreSQL)
-- In postgres.conf of replica:
# primary_conninfo = 'host=master port=5432 user=repl password=repl_password'
```

### Connection Pooling

Use PgBouncer for connection pooling:

```yaml
# docker-compose.yml
services:
  pgbouncer:
    image: edoburu/pgbouncer
    environment:
      DATABASE_URL: postgres://postgres:password@psql_bp:5432/ecommerce
      POOL_MODE: transaction
      MAX_CLIENT_CONN: 100
      DEFAULT_POOL_SIZE: 20
```

### Query Optimization

```sql
-- Add indexes for common queries
CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_orders_customer ON orders(customer_id);
CREATE INDEX idx_order_items_product ON order_items(product_id);

-- Use EXPLAIN to analyze slow queries
EXPLAIN ANALYZE SELECT * FROM orders WHERE customer_id = 'cust-123';
```

## Caching Strategy

### Current Caching

| Key Pattern | Data | TTL |
|-------------|------|-----|
| `products:list:*` | Product listings | 5 min |
| `products:{id}` | Single product | 10 min |
| `categories:*` | Categories | 15 min |
| `session:*` | JWT sessions | 24 hours |

### Scaling Redis

```yaml
# Redis Cluster (future)
# redis-cluster:
#   image: redis:7-alpine
#   command: redis-server --cluster-enabled yes
```

### Cache Invalidation

```go
// On product update
func (s *ProductService) Update(id string, req *UpdateProductRequest) error {
    // Update in DB
    err := s.repo.Update(id, req)
    if err != nil {
        return err
    }

    // Invalidate cache
    s.cache.Del(ctx, "products:"+id)
    s.cache.Del(ctx, "products:list:*")
    // Categories also contain products
    s.cache.Del(ctx, "categories:*")

    return nil
}
```

### Cache Warming

```bash
#!/bin/bash
# warm-cache.sh

# Warm product cache
curl -s http://localhost:8080/products?page=1 > /dev/null

# Warm categories
curl -s http://localhost:8080/categories > /dev/null

# Warm category hierarchy
curl -s http://localhost:8080/categories/hierarchy > /dev/null
```

## Scaling to 10K+ Users

### Architecture

```mermaid
graph TB
    subgraph "Edge"
        CDN[CDN - Static Assets]
        LB[Load Balancer]
    end

    subgraph "App Tier"
        API1[API Instance 1]
        API2[API Instance 2]
        API3[API Instance 3]
    end

    subgraph "Data Tier"
        PG_Master[(PostgreSQL Master)]
        PG_Replica[(PostgreSQL Replica)]
        Redis[(Redis Cluster)]
        PGB[PgBouncer]
    end

    CDN --> LB
    LB --> API1
    LB --> API2
    LB --> API3
    API1 --> PGB
    API2 --> PGB
    API3 --> PGB
    PGB --> PG_Master
    PG_Master --> PG_Replica
    API1 --> Redis
    API2 --> Redis
    API3 --> Redis
```

### Requirements at Scale

| Component | 1K Users | 10K Users | 100K Users |
|-----------|-----------|-----------|------------|
| API Instances | 1 | 3-5 | 10+ |
| PostgreSQL | Single | Master + Replica | Sharded |
| Redis | Single | Cluster | Cluster |
| RAM | 4GB | 16GB | 64GB+ |

## Performance Optimization

### Database

1. **Index Strategy**
   ```sql
   -- Composite index for filtered queries
   CREATE INDEX idx_products_cat_price 
   ON products(category_id, price);

   -- Partial index for active records
   CREATE INDEX idx_orders_active 
   ON orders(customer_id, order_date) 
   WHERE status != 'cancelled';
   ```

2. **Query Optimization**
   ```go
   // Use goqu for efficient queries
   query := goqu.Select(
       "p.product_id", "p.name", "p.price",
   ).From("products p").Where(
       goqu.Ex{"p.category_id": categoryID},
   ).Limit(20)

   // Avoid N+1: Use JOIN or batch queries
   ```

### Caching

1. **Increase TTL for stable data**
   ```go
   // Categories change rarely - longer TTL
   cache.Set(ctx, "categories", data, 1*time.Hour)
   ```

2. **Use cache-aside pattern**
   ```go
   func (s *ProductService) GetByID(id string) (*Product, error) {
       // Check cache first
       cached, err := s.cache.Get(ctx, "products:"+id)
       if err == nil {
           return cached, nil
       }

       // Fetch from DB
       product, err := s.repo.GetByID(id)
       if err != nil {
           return nil, err
       }

       // Store in cache
       s.cache.Set(ctx, "products:"+id, product, 10*time.Minute)
       return product, nil
   }
   ```

### API

1. **Response Compression**
   ```go
   // Already enabled in Fiber
   app := fiber.New(fiber.Config{
       Compress: true,
   })
   ```

2. **HTTP/2 Support**
   ```go
   // Enable HTTP/2
   app.Listen(":8080", fiber.ListenConfig{
       EnablePrefork: true,
   })
   ```

## Monitoring for Scaling

### Key Metrics

| Metric | Tool | Alert Threshold |
|--------|------|----------------|
| CPU Usage | Prometheus | > 80% |
| Memory | Prometheus | > 80% |
| Request Latency | Prometheus | > 500ms |
| DB Connections | pg_stat_activity | > 80% max |
| Cache Hit Rate | Redis INFO | < 80% |
| Error Rate | Prometheus | > 1% |

### Setup Prometheus

```yaml
# prometheus.yml
scrape_configs:
  - job_name: 'ecommerce-api'
    static_configs:
      - targets: ['api:8080']
```

### Dashboard

Use Grafana to visualize:
- Request rate
- Latency percentiles (p50, p95, p99)
- Error rate
- Database metrics
- Cache metrics

## Capacity Planning

### Estimate

For 10,000 daily active users:
- ~100 requests/user/day = 1M requests/day
- ~12 requests/second average
- ~50 requests/second peak

### Headroom

Plan for:
- 3x peak for safety
- Add instances before hitting 80% CPU
- Monitor and scale proactively

### Auto-Scaling (Kubernetes)

```yaml
# api-hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: api-hpa
spec:
  scaleTargetRef:
    api: api-deployment
  minReplicas: 3
  maxReplicas: 20
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
```

## Summary Checklist

- [ ] Implement load balancer for multi-instance
- [ ] Configure Redis for session sharing
- [ ] Set up PgBouncer for connection pooling
- [ ] Add database read replica
- [ ] Implement cache invalidation
- [ ] Set up monitoring and alerts
- [ ] Define auto-scaling rules
- [ ] Test failover scenarios

## Future Considerations

1. **Microservices** - Split into product, order, user services
2. **Message Queue** - Async processing for orders, emails
3. **GraphQL** - Flexible frontend queries
4. **CDN** - Static asset delivery