# ADR 004: Use Go Fiber as Web Framework

## Status
Accepted

## Context
We needed to choose a web framework for building the REST API in Go.

**Requirements:**
- High performance (low latency, high throughput)
- Easy to use (Express-like API)
- Good middleware ecosystem
- HTTP/2 support
- Active maintenance

## Decision
Use **Fiber v3** as the web framework.

## Rationale

### Why Fiber over other options?

#### Comparison with Gin
| Feature | Fiber | Gin |
|---------|-------|-----|
| Performance | Faster (fasthttp) | Very fast |
| HTTP/2 | Native | Via std lib |
| API Style | Express-like | gin-like |
| Middleware | Compatible | Native |
| Actively Maintained | Yes | Yes |

#### Comparison with Echo
| Feature | Fiber | Echo |
|---------|-------|------|
| Performance | Slightly faster | Fast |
| API Style | Express-like | Custom |
| Fiber v3 | Migration in progress | Stable |

#### Why not use stdlib net/http?
- More verbose code
- No middleware system
- No routing helpers

### Why Fiber specifically?
1. **Express.js compatibility** - Easy for web developers to adapt
2. **Zero allocation** - Optimized for performance
3. **Built on fasthttp** - Faster than net/http
4. **Rich middleware** - CORS, rate limiting, logging, etc.
5. **Fiber v3** - Latest version with improved stability

## Consequences

### Positive
- High performance (comparable tofasthttp standalone)
- Express-like API familiar to web developers
- Rich built-in middleware
- Good error handling
- Active community and maintenance

### Negative
- Fiber v3 migration - Some breaking changes from v2
- Less "idiomatic" Go than stdlib
- Additional dependency

### Mitigations
- Follow Go best practices within Fiber
- Use go.mod for version pinning
- Test thoroughly after updates

## Implementation Notes

### Basic Server Setup
```go
app := fiber.New(fiber.Config{
    ServerHeader: "backend",
    AppName:      "ecommerce-api",
})
```

### Middleware Usage
```go
app.Use(middleware.Logger())
app.Use(middleware.Recovery())
app.Use(middleware.CORS())
```

### Route Definition
```go
app.Get("/products", productController.GetAll)
app.Post("/products", authMiddleware.Authenticate, productController.Create)
```

## References
- [Fiber v3 docs](https://docs.gofiber.io/)
- [Fiber GitHub](https://github.com/gofiber/fiber)
- [Benchmarks](https://github.com/gofiber/benchmarks)