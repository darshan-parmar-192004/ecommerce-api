# ADR 003: Use Redis for Caching and Session Storage

## Status
Accepted

## Context
We needed a caching layer for improving API performance and a session store for JWT authentication.

**Requirements:**
- Fast in-memory data store
- Support for TTL/expiration
- Session storage for JWT revocation
- Query caching for database results
- Pub/sub for future features (optional)

## Decision
Use **Redis 7+** for caching and session storage.

## Rationale

### Why Redis?
1. **In-memory speed** - Sub-millisecond response times
2. **Built-in expiration** - Perfect for TTL-based caching
3. **Session support** - Native key-value with expiry
4. **Rich data types** - Strings, hashes, sets, sorted sets
5. **Pub/Sub** - For future real-time features
6. **Cluster mode** - Horizontal scaling when needed

### Use Cases Implemented

#### 1. Session Storage
```go
// Store JWT session
sessionKey := "session:" + customerID
redis.Set(ctx, sessionKey, token, 24*time.Hour)
```
- Stores active JWT sessions for revocation
- TTL matches JWT expiry (24 hours)
- Key pattern: `session:{customer_id}`

#### 2. Query Caching
```go
// Cache product list
cacheKey := "products:list:page=1:limit=20"
redis.Set(ctx, cacheKey, products, 5*time.Minute)
```
- Caches frequently accessed data
- TTL: 5 minutes for product listings
- Invalidation on data changes

### Alternative Considered
- **Memcached**: Less feature-rich than Redis, no pub/sub
- **In-memory Go map**: No persistence, no TTL, single instance only

## Consequences

### Positive
- Dramatically improved read performance (sub-millisecond vs DB milliseconds)
- Session management with automatic expiration
- Simple integration with Go (go-redis library)
- Persistence option (RDB/AOF) for development
- Data structure support beyond simple key-value

### Negative
- Additional infrastructure to manage
- Cache invalidation complexity
- Memory constraints (all data in RAM)
- Network latency (cache is separate service)

### Mitigations
- Appropriate TTL based on data volatility
- Cache-aside pattern (not write-through)
- Cache invalidation on mutations
- Monitor memory usage

## Implementation Details

### Connection Configuration
```go
// go-redis configuration
redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
```

### Connection Pool
```go
// pgx (PostgreSQL) pool settings
MaxOpenConnections:    20
MaxIdleConnections:    5
ConnMaxLifetime:       30 * time.Minute
```

### Cache Keys Pattern
| Pattern | Description | TTL |
|---------|-------------|-----|
| `products:list:*` | Product listings | 5 min |
| `products:{id}` | Single product | 10 min |
| `categories:list` | All categories | 15 min |
| `categories:tree` | Category hierarchy | 1 hour |
| `session:{customer_id}` | JWT session | 24 hours |
| `stats:*` | Statistics data | 1 min |

## Future Considerations
- **Cache warming**: Pre-populate cache on startup
- **Redis Cluster**: For high availability and scaling
- **Pub/Sub**: For real-time notifications
- **Lua scripts**: Atomic cache operations

## References
- [go-redis](https://github.com/redis/go-redis)
- [Redis persistence](https://redis.io/docs/management/persistence/)
- [Cache patterns](https://redis.io/docs/manual patterns/)