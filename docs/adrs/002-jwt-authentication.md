# ADR 002: JWT Authentication

## Status
Accepted

## Context
We needed to choose an authentication mechanism for the API. The main options were session-based authentication (server-side sessions) and token-based authentication (JWT).

**Requirements:**
- Secure user identification
- Scalable across multiple API instances
- Support for role-based access control
- Integration with frontend (React/Nuxt)
- Easy to implement and maintain

## Decision
Use **JWT (JSON Web Tokens)** with Redis for session storage.

## Rationale

### Why JWT over Session-based?
1. **Stateless** - No server-side session storage needed (except for revocation)
2. **Scalable** - Token contains all needed info; load balancers don't need session affinity
3. **Cross-platform** - Works with any language/framework
4. **Mobile-friendly** - Easy to use with mobile apps

### Why add Redis session storage?
1. **Token Revocation** - Ability to invalidate tokens before expiry
2. **Logout Support** - User can logout and session is invalidated
3. **Security** - Quick revocation for security incidents

### Implementation Details
- **Algorithm**: HS256 (HMAC SHA-256)
- **Token Expiry**: 24 hours
- **Claims**: customer_id, email, role, exp
- **Storage**: JWT in response + Redis for session tracking

## Consequences

### Positive
- Stateless authentication scales horizontally
- Simple to implement and use
- Works across different clients (web, mobile)
- Redis session storage provides revocation capability
- JWT is industry standard

### Negative
- **Token Revocation** - Cannot truly revoke a JWT (only can blacklist via Redis)
- **Token Size** - Larger than session ID
- **Storage** - Need Redis for session management
- **Security** - Token in URL could be logged; need careful cookie handling

### Mitigations
- Use HTTP-only, Secure, SameSite cookies for web clients
- Short token expiry (24h)
- Redis blacklisting for immediate revocation
- Rate limiting on auth endpoints

## Alternatives Considered
- **Session-based**: Rejected - requires sticky sessions, harder to scale
- **OAuth2/OIDC**: Overkill for single-application auth
- **API Keys**: Good for service-to-service, not user auth

## Implementation

### JWT Token Structure
```go
type Claims struct {
    CustomerID string `json:"customer_id"`
    Email      string `json:"email"`
    Role       string `json:"role"`
    jwt.RegisteredClaims
}
```

### Middleware Flow
```
Request -> JWT Validation -> Check Redis Session -> Allow/Deny
```

### Cookie Configuration
```go
Cookie: "auth_token"
- HttpOnly: true
- Secure: true (in production)
- SameSite: Lax
- Path: /
- MaxAge: 86400 (24 hours)
```

## References
- [JWT.io](https://jwt.io/)
- [golang-jwt](https://github.com/golang-jwt/jwt)
- [Best practices for JWT](https://auth0.com/blog/tenancy-billing-jwt-best-practices/)