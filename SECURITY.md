# Security Policy

## Supported Versions

We support and provide security updates for the following versions:

| Version | Supported | Notes |
|---------|-----------|-------|
| 1.0.x | Yes | Current stable release |
| 0.1.x | No | Deprecated |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please report it responsibly.

### How to Report

**Do NOT create a public GitHub issue for security vulnerabilities.**

Instead, please report them via one of these methods:

1. **Email**: Send email to security@example.com (replace with your email)
2. **GitHub Security Advisories**: Use [GitHub's private vulnerability reporting](https://github.com/yourusername/ecommerce-api/security/advisories/new)

### What to Include

When reporting, please include:

- Type of vulnerability
- Full paths of source file(s) related to the vulnerability
- Location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### Response Timeline

We aim to respond to security reports within:

- **Acknowledgment**: Within 48 hours
- **Initial Assessment**: Within 7 days
- **Fix Timeline**: Depending on severity
  - Critical: 24-72 hours
  - High: 7 days
  - Medium: 30 days
  - Low: Next release cycle

## Security Best Practices

### For Users

1. **Use Strong Passwords**: Minimum 8 characters with uppercase, lowercase, digit, and special character
2. **Keep Tokens Secure**: Don't expose JWT tokens in URLs
3. **Use HTTPS**: Always use HTTPS in production
4. **Rotate Secrets**: Periodically rotate JWT_SECRET and database passwords

### For Administrators

1. **Environment Variables**: Never commit secrets to version control
   ```bash
   # Use .env files (already in .gitignore)
   JWT_SECRET=your-32-character-minimum-secret-key
   ```

2. **Database Security**
   ```sql
   -- Use strong passwords
   ALTER USER ecommerce_user WITH PASSWORD 'strong_password_here';
   
   -- Limit connections
   ALTER SYSTEM SET max_connections = 100;
   ```

3. **Redis Security**
   ```bash
   # Set Redis password in redis.conf
   requirepass your_redis_password
   
   # Or in production use protected mode
   protected-mode yes
   ```

4. **Network Security**
   - Only expose port 8080 (API) externally
   - Keep PostgreSQL (5432) and Redis (6379) internal
   - Use firewall rules to restrict access

5. **Regular Updates**
   - Keep Go version updated
   - Update dependencies: `go get -u ./...`
   - Update PostgreSQL and Redis to latest stable

## Security Features Implemented

### Authentication
- JWT tokens with HS256 signing
- 24-hour token expiry
- Redis session storage for revocation

### Password Security
- bcrypt hashing (cost factor 10)
- Password validation rules

### Rate Limiting
- Login endpoint: 5 requests/minute/IP
- General endpoints: 100 requests/minute/IP

### HTTP Security
- Security headers (X-Frame-Options, X-Content-Type-Options, etc.)
- CORS configuration
- HTTP-only cookies for auth tokens

### Data Protection
- Prepared statements (parameterized queries)
- Input validation
- SQL injection prevention (via goqu)

## Known Limitations

1. **JWT Revocation**: Tokens can only be revoked via Redis session. If Redis is unavailable, revocation won't work.
2. **Session Security**: Session data is stored in Redis. Ensure Redis is secured and backed up.
3. **No Audit Logging**: Currently no audit trail for sensitive operations. Consider adding for production.

## Security Updates

Security updates will be released as:
- **Hotfixes**: For critical vulnerabilities
- **Patch releases**: For minor issues
- **Release notes**: Clearly indicate security fixes

Subscribe to release notifications to stay updated.

## Acknowledgments

Thank you to the following for responsibly disclosing vulnerabilities:

- [List contributors or "None yet" if no vulnerabilities reported]

## Contact

For security-related questions, contact: security@example.com

---

*Last updated: 2024-01-15*
*Version: 1.0.0-alpha*