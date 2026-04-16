# ADR 001: Use PostgreSQL as Primary Database

## Status
Accepted

## Context
We needed to choose a primary database for storing e-commerce data (products, orders, customers, inventory). The main options considered were PostgreSQL, MySQL, and MongoDB.

**Requirements:**
- ACID compliance for financial transactions (orders, payments)
- Relational data with complex joins (orders -> customers -> products)
- Support for complex queries (aggregations, filtering, sorting)
- Good indexing performance
- JSON support for flexible schemas
- Mature ecosystem and community support

## Decision
Use **PostgreSQL 16+** as the primary database.

## Rationale

### Why PostgreSQL over MySQL?
1. **Better JSON Support** - Native JSONB type with indexed queries
2. **More Advanced Indexing** - Partial, GiST, GIN indexes
3. **Better Concurrency** - MVCC implementation is superior
4. **More ANSI SQL Compliance** - Fewer edge cases
5. **Rich Ecosystem** - Extensions like PostGIS, TimescaleDB

### Why not MongoDB?
1. **ACID Transactions** - MongoDB's transactions are newer and less mature
2. **Complex Joins** - Not MongoDB's strength
3. **Data Integrity** - Relational integrity is harder to enforce
4. **Team Expertise** - More Go developer familiarity with SQL

### Why not other options?
- **SQLite**: Not suitable for concurrent multi-user e-commerce
- **CockroachDB**: Overkill for current scale
- **TimescaleDB**: Could add later for time-series needs

## Consequences

### Positive
- ACID compliance ensures data integrity for orders and transactions
- Complex queries supported (aggregations, JOINs, subqueries)
- JSONB for flexible metadata without sacrificing relational integrity
- Rich indexing strategies for performance
- Excellent Go driver (pgx) with connection pooling
- golang-migrate for version-controlled schema management

### Negative
- More complex setup than SQLite
- JSON handling slightly more verbose than MongoDB
- Horizontal scaling requires more effort than NoSQL

### Mitigations
- Use goqu for type-safe SQL query building
- Connection pooling via pgx for performance
- Redis caching for frequently accessed data

## Alternatives Considered
- MySQL: Rejected for weaker JSON support and less mature concurrency
- MongoDB: Rejected for ACID concerns and complex joins
- SQLite: Rejected for concurrent access limitations

## References
- [PostgreSQL vs MySQL](https://www.postgresql.org/about/)
- [pgx driver](https://github.com/jackc/pgx)
- [goqu query builder](https://github.com/doug-martin/goqu)