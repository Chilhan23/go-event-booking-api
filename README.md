# Event & Booking API (Go)

A backend REST API for an Event Management and Ticket Booking System built with Go, PostgreSQL, and GORM.

---

## Tech Stack

* **Go** (`1.26+`) — Core programming language
* **Gin Gonic** — High-performance HTTP web framework & routing
* **PostgreSQL** — Relational database with ACID transaction support
* **GORM** — Object-Relational Mapping (ORM) for data manipulation & queries
* **golang-migrate** — Explicit database schema versioning and DDL migrations
* **JWT (`golang-jwt/jwt/v5`)** — Stateless authentication and bearer tokens
* **bcrypt (`golang.org/x/crypto/bcrypt`)** — Secure password hashing
* **godotenv** — Environment variable management

---

## Architectural Design

The project follows a **Feature-Oriented Structure** with a **Clean 3-Tier Layered Architecture**:

```text
HTTP Request
     │
┌────▼──────────────┐
│  Handler Layer    │  HTTP parsing, DTO binding, request validation, HTTP status codes
└────┬──────────────┘
     │ (DTOs)
┌────▼──────────────┐
│  Service Layer    │  Business logic, password hashing, JWT generation, business rules
└────┬──────────────┘
     │ (Models)
┌────▼──────────────┐
│  Repository Layer │  Database persistence via GORM, context propagation
└────┬──────────────┘
     │
┌────▼──────────────┐
│  PostgreSQL DB    │  Relational storage, constraints, indexes
└───────────────────┘
```

### Layer Responsibilities

1. **DTO (`dto.go`) vs Model (`model.go`)**:
   - `DTO`: Represents the network payload contract between client and server, enforcing validation rules (`binding:"required,email,min=6"`). Prevents sensitive database fields from leaking into API responses.
   - `Model`: Represents the underlying database schema and table mappings in PostgreSQL.

2. **Handler Layer (`handler.go`)**:
   - Parses JSON request bodies into DTOs.
   - Forwards `c.Request.Context()` downstream to support request cancellation.
   - Formats HTTP responses and appropriate status codes (`200 OK`, `201 Created`, `400 Bad Request`, `401 Unauthorized`).

3. **Service Layer (`service.go`)**:
   - Pure business logic layer decoupled from the HTTP transport layer.
   - Manages domain business rules, password hashing with `bcrypt`, and token signing with `jwt`.

4. **Repository Layer (`repository.go`)**:
   - Handles database queries using GORM.
   - Every database operation is bound to `r.db.WithContext(ctx)` to respect request lifecycles.

5. **Middleware Layer (`internal/middleware/`)**:
   - Cross-cutting concerns such as JWT Bearer token authentication and request logging.
