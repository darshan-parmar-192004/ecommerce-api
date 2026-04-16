# System Overview

## High-Level Architecture

```mermaid
graph TB
    subgraph "Client Layer"
        F[Frontend - React/Vite]
        N[Frontend - Nuxt 3]
        M[Mobile Apps]
    end

    subgraph "API Layer"
        G[Go Fiber API<br/>Port 8080]
        M1[Middleware<br/>Auth, Rate Limit, CORS]
    end

    subgraph "Service Layer"
        P[Product Service]
        C[Category Service]
        CU[Customer Service]
        O[Order Service]
        I[Inventory Service]
        A[Auth Service]
    end

    subgraph "Data Layer"
        PG[(PostgreSQL<br/>Port 5432)]
        R[(Redis<br/>Port 6379)]
    end

    F --> G
    N --> G
    M --> G
    G --> M1
    M1 --> P
    M1 --> C
    M1 --> CU
    M1 --> O
    M1 --> I
    M1 --> A
    P --> PG
    C --> PG
    CU --> PG
    O --> PG
    I --> PG
    A --> PG
    P --> R
    C --> R
    A --> R
```

## Components

### Frontend Layer
- **React + Vite**: Primary frontend (port 5173)
- **Nuxt 3**: Experimental alternative frontend
- **Tailwind CSS**: Styling framework

### API Layer (Go Fiber)
- **Framework**: Fiber v3 (Express-like API)
- **Port**: 8080 (configurable via PORT env)
- **Middleware**:
  - Authentication (JWT validation)
  - Authorization (role-based access)
  - Rate limiting (general + login-specific)
  - CORS (configurable origins)
  - Request logging
  - Security headers

### Service Layer
| Service | Description |
|---------|-------------|
| Product | CRUD operations, caching |
| Category | Hierarchy management |
| Customer | Profile, orders, CLV |
| Order | Order creation, ownership validation |
| Inventory | Stock, analytics |
| Auth | JWT issuance, validation, sessions |

### Data Layer
- **PostgreSQL**: Primary data store
- **Redis**: Session storage, query caching

## Data Flow

### Request Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant M as Middleware
    participant S as Service
    participant R as Repository
    participant DB as PostgreSQL
    participant CACHE as Redis

    C->>M: HTTP Request
    M->>M: Auth Check
    M->>M: Rate Limit
    M->>S: Valid Request
    S->>R: Business Logic
    R->>CACHE: Check Cache
    alt Cache Hit
        CACHE-->>R: Return Cached Data
    else Cache Miss
        R->>DB: Query Database
        DB-->>R: Result
        R->>CACHE: Store in Cache
    end
    R-->>S: Result
    S-->>M: Response
    M-->>C: HTTP Response
```

### Authentication Flow

```mermaid
sequenceDiagram
    participant U as User
    participant API as API Server
    participant JWT as JWT Module
    participant REDIS as Redis
    participant DB as PostgreSQL

    U->>API: POST /auth/login
    API->>DB: Validate Credentials
    DB-->>API: User Data
    API->>JWT: Generate Token
    JWT-->>API: JWT Token
    API->>REDIS: Store Session
    API-->>U: Set-Cookie: auth_token

    U->>API: GET /protected (with cookie)
    API->>JWT: Validate Token
    JWT-->>API: Claims (customer_id, role)
    API->>REDIS: Verify Session
    REDIS-->>API: Session Valid
    API-->>U: Protected Resource
```

## Technology Choices

| Component | Technology | Rationale |
|-----------|------------|-----------|
| API Framework | Fiber v3 | Fast, Express-like, mature |
| Database | PostgreSQL | ACID compliance, JSON support |
| Cache | Redis | Speed, session storage |
| Auth | JWT + Redis | Stateless, scalable |
| Query Builder | goqu | Type-safe SQL |
| Password | bcrypt | Secure, proven |
| Migrations | golang-migrate | Version control for DB |
| Containers | Podman/Docker | DevOps consistency |

## Infrastructure

### Development
- Local PostgreSQL (port 5432)
- Local Redis (port 6379)
- Go Fiber on port 8080

### Production (Containerized)
```yaml
services:
  app:
    ports: [8080:8080]
  psql_bp:
    ports: [5432:5432]
  redis:
    ports: [6379:6379]
```

### Network
- Bridge network: `blueprint`
- All services communicate via internal network
- Only API port exposed to host