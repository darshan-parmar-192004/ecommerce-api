# ADR 005: Use Podman over Docker

## Status
Accepted

## Context
We needed to choose a container runtime for development and deployment.

**Requirements:**
- Run containers locally (development)
- Support Docker Compose files
- No root privileges required (security)
- Integration with CI/CD

## Decision
Use **Podman** as the container runtime (compatible with Docker).

## Rationale

### Why Podman over Docker?
1. **Rootless** - Runs without root privileges by default
2. **Docker-compatible** - Same CLI commands, docker-compose support
3. **No daemon** - No background daemon required
4. **Security** - Better isolation, no privileged daemon
5. **Podman Compose** - Drop-in replacement for docker-compose

### Comparison
| Feature | Podman | Docker |
|---------|--------|--------|
| Rootless | Yes | No |
| Daemonless | Yes | No |
| docker-compose | podman-compose | Native |
| CLI | podman | docker |
| Buildah compatible | Yes | No |

## Consequences

### Positive
- Better security (no privileged daemon)
- Works without root access
- Drop-in replacement for Docker
- Compatible with docker-compose.yml
- Modern container runtime

### Negative
- Different binary name (podman vs docker)
- Some CI/CD might need adjustment
- podman-compose slightly different

### Mitigations
- Use docker-compose.yml (works with both)
- Document both commands in docs
- Scripts can check for available tool

## Usage

### Development Commands
```bash
# Start all services
podman-compose up -d --build

# Or with Docker
docker-compose up -d --build

# View logs
podman-compose logs -f app

# Stop services
podman-compose down
```

### Without Compose
```bash
# Build
podman build -t ecommerce-api ./backend

# Run
podman run -p 8080:8080 --env-file .env ecommerce-api
```

## References
- [Podman docs](https://docs.podman.io/)
- [Podman Compose](https://github.com/containers/podman-compose)
- [Migrating from Docker](https://podman.io/getting-started/)