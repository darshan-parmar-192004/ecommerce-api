# API Design

## RESTful Design Principles

This API follows REST (Representational State Transfer) conventions:

1. **Resource-Oriented URLs** - URLs represent resources, not actions
2. **HTTP Methods** - Use appropriate HTTP verbs (GET, POST, PUT, DELETE)
3. **Stateless** - Each request contains all information needed
4. **Standard Status Codes** - Use HTTP status codes appropriately
5. **JSON Responses** - Consistent JSON format for all responses

## URL Structure

### Base URL
```
http://localhost:8080/api (proposed)
```

For now, endpoints are at root:
```
http://localhost:8080/
```

### Resource Naming
- **Plural nouns** for collections: `/products`, `/orders`, `/categories`
- **Singular for specific items**: `/products/:id`, `/orders/:id`
- **CamelCase not used**: use-kebab-case or snake_case in URLs

### URL Hierarchy
```
/products              # Product collection
/products/:id          # Single product
/categories            # Category collection
/categories/:id        # Single category
/categories/:id/products  # Products in category
/customers/me          # Current user (relative path)
/customers/:id         # Specific customer
/customers/:id/orders  # Customer's orders
/orders/:id            # Single order
/inventory             # Inventory collection
/auth/register         # Auth action
/auth/login            # Auth action
```

## HTTP Methods

| Method | Usage | Idempotent |
|--------|-------|------------|
| GET | Retrieve resources | Yes |
| POST | Create new resources | No |
| PUT | Replace existing resource | Yes |
| PATCH | Partial update | No |
| DELETE | Remove resource | Yes |

## Authentication & Authorization

### Authentication

The API uses **JWT (JSON Web Tokens)** with Redis session storage.

**Token Flow:**
1. User submits credentials to `/auth/login`
2. Server validates and issues JWT (HS256)
3. Server stores session in Redis with TTL
4. Client includes token in subsequent requests

**Token Storage:**
- **Cookie** (default): HTTP-only, Secure, SameSite=Lax
- **Bearer Header**: Alternative for API clients

**Token Contents:**
```json
{
  "customer_id": "uuid",
  "email": "user@example.com",
  "role": "customer|admin",
  "exp": 1234567890
}
```

**Token Expiry:** 24 hours

### Authorization

Role-based access control (RBAC):

| Role | Access |
|------|--------|
| `customer` | Own data, public endpoints |
| `admin` | All customer data, inventory, analytics |

**Protected Endpoints:**
- `/products` (POST, PUT, DELETE) - Admin only
- `/customers/me` - Owner only
- `/customers/:id/orders` - Owner or admin
- `/customers/:id/lifetime-value` - Admin only
- `/orders/*` - Owner or admin
- `/inventory/*` - Admin only
- `/auth/logout` - Authenticated users
- `/auth/me` - Authenticated users

### Middleware

```go
// Authenticate - validates JWT
authMiddleware.Authenticate

// RequireAdmin - checks for admin role
middleware.RequireAdmin()

// ValidateCustomerAccess - ensures user can access data
middleware.ValidateCustomerAccess

// ValidateOrderOwnership - order access check
middleware.ValidateOrderOwnership
```

## Error Handling Strategy

### HTTP Status Codes

| Code | Usage |
|------|-------|
| 200 | Successful GET, PUT, PATCH |
| 201 | Successful POST (created) |
| 204 | Successful DELETE (no content) |
| 400 | Invalid input, validation error |
| 401 | Missing or invalid authentication |
| 403 | Authenticated but not authorized |
| 404 | Resource not found |
| 409 | Conflict (e.g., duplicate email) |
| 422 | Validation failed |
| 429 | Rate limit exceeded |
| 500 | Internal server error |

### Error Response Format

```json
{
  "error": "Human-readable message",
  "code": "MACHINE_READABLE_CODE"
}
```

**Example Errors:**

```json
// 400 Bad Request
{
  "error": "Invalid email format",
  "code": "INVALID_EMAIL"
}

// 401 Unauthorized
{
  "error": "Invalid or expired token",
  "code": "UNAUTHORIZED"
}

// 403 Forbidden
{
  "error": "Admin access required",
  "code": "FORBIDDEN"
}

// 404 Not Found
{
  "error": "Product not found",
  "code": "NOT_FOUND"
}

// 409 Conflict
{
  "error": "Email already registered",
  "code": "DUPLICATE_EMAIL"
}

// 429 Too Many Requests
{
  "error": "Rate limit exceeded. Try again later.",
  "code": "RATE_LIMITED"
}

// 500 Internal Error
{
  "error": "Internal server error",
  "code": "INTERNAL_ERROR"
}
```

### Custom Error Codes

Defined in `backend/internal/errors/errors.go`:
- `ErrCodeBadRequest`
- `ErrCodeUnauthorized`
- `ErrCodeForbidden`
- `ErrCodeNotFound`
- `ErrCodeDuplicate`
- `ErrCodeInternalError`

## Request/Response Patterns

### Pagination

List endpoints support pagination:

```
GET /products?page=1&limit=20
```

**Response:**
```json
{
  "products": [...],
  "total": 150,
  "page": 1,
  "limit": 20
}
```

### Filtering

Query parameters for filtering:

```
GET /products?category_id=cat-001&min_price=10&max_price=100&search=mouse
```

### Sorting

Not currently implemented, but pattern would be:

```
GET /products?sort=price&order=desc
```

### Field Selection

Not currently implemented.

## Versioning

Current: **v1** (no version in URL)

Future consideration: `/api/v1/products`

## Content Negotiation

- **Request**: Accept header (application/json)
- **Response**: Always application/json

## CORS Configuration

Allowed origins (configurable):
- `http://localhost:3000`
- `http://localhost:5173`
- `http://127.0.0.1:5173`

Allowed methods: GET, POST, PUT, DELETE, OPTIONS

Allowed headers: Accept, Authorization, Content-Type

Credentials: enabled

Max age: 300 seconds

## Rate Limiting

| Endpoint | Limit |
|----------|-------|
| `/auth/login` | 5 requests/minute/IP |
| All other | 100 requests/minute/IP |

Response when rate limited:
```json
{
  "error": "Rate limit exceeded. Try again later.",
  "code": "RATE_LIMITED"
}
```

## API Standards Summary

| Aspect | Standard |
|--------|----------|
| Format | JSON |
| URL | /resource or /resource/:id |
| Method | GET/POST/PUT/DELETE |
| Auth | JWT Bearer + Cookie |
| Errors | Consistent JSON with codes |
| Status | Standard HTTP codes |
| Pagination | page/limit params |
| Rate Limit | Per-endpoint limits |