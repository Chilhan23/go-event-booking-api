# Event Booking REST API

A high-performance, production-ready RESTful Event Management and Ticket Booking API built with Go, Gin Framework, GORM, and PostgreSQL. Designed following Clean Architecture principles and a Feature-Oriented package structure (Handler, Service, Repository, DTO, and Model).

---

## Features

- **Authentication & Authorization:** Secure user registration and login with bcrypt password hashing and JWT (JSON Web Tokens) bearer authentication.
- **Event Management:** Create, retrieve, and inspect events with start/end schedules, venue locations, and seat quotas.
- **Concurrency-Safe Ticket Booking:** ACID-compliant ticket reservation powered by PostgreSQL Row-Level Locking (`SELECT FOR UPDATE`) and database transactions to strictly prevent overbooking during concurrent traffic spikes.
- **Duplicate Booking Prevention:** Enforced at the database engine level via composite `UNIQUE(event_id, user_id)` constraints.
- **Context & Cancellation Propagation:** Request contexts (`c.Request.Context()`) are forwarded across handlers, services, and repositories to prevent orphan queries upon client disconnection.
- **Clean Layer Separation:** Strict decoupling between network contracts (DTOs) and database schemas (Models) to prevent sensitive data leakage.
- **Database Migrations:** Automated schema versioning and DDL migrations managed via `golang-migrate` and custom helper scripts.

---

## Tech Stack

- **Language:** Go 1.26+
- **Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Database:** PostgreSQL
- **ORM:** [GORM](https://gorm.io)
- **Authentication:** JWT (`github.com/golang-jwt/jwt/v5`) & bcrypt (`golang.org/x/crypto/bcrypt`)
- **Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **Hot Reload:** [Air](https://github.com/air-verse/air)

---

## Project Structure

```
.
├── cmd/
│   └── main.go                      # Application entrypoint & dependency wiring
├── internal/
│   ├── config/
│   │   └── config.go                # Environment variable configuration
│   ├── database/
│   │   └── postgres.go              # GORM PostgreSQL connection setup
│   ├── middleware/
│   │   ├── auth.go                  # JWT Bearer token authentication middleware
│   │   └── logger.go                # HTTP request logger middleware
│   ├── auth/
│   │   ├── handler.go               # Auth HTTP handlers & request binding
│   │   ├── service.go               # Business logic, bcrypt hashing & token generation
│   │   ├── repository.go            # Database queries via GORM
│   │   ├── model.go                 # User database model
│   │   └── dto.go                   # User request & response DTOs
│   ├── events/
│   │   ├── handler.go               # Event HTTP handlers
│   │   ├── service.go               # Event business rules & validation
│   │   ├── repository.go            # Event database queries
│   │   ├── model.go                 # Event database model
│   │   └── dto.go                   # Event request & response DTOs
│   └── booking/
│       ├── handler.go               # Booking HTTP handlers
│       ├── service.go               # Transactional booking logic & locking
│       ├── repository.go            # Booking database queries & row-level locks
│       ├── model.go                 # Booking database model
│       └── dto.go                   # Booking request & response DTOs
├── migrations/                      # Versioned SQL migration files (.up.sql & .down.sql)
├── scripts/
│   └── migrate.sh                   # Helper script for migration operations
├── .air.toml                        # Air hot reload configuration
├── .env                             # Environment variables
├── go.mod
└── go.sum
```

---

## Environment Variables

Create a `.env` file in the root directory:

```env
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/event_booking_db?sslmode=disable
JWT_SECRET=your_secret_jwt_key
```

---

## Database Migrations

Manage database schemas using the migration script:

```bash
# Apply all pending migrations
./scripts/migrate.sh up

# Rollback last migration
./scripts/migrate.sh down

# Check migration version status
./scripts/migrate.sh status

# Create a new migration file
./scripts/migrate.sh create <migration_name>
```

---

## API Documentation

### Authentication

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| `POST` | `/auth/register` | None | Register a new user account |
| `POST` | `/auth/login`    | None | Authenticate user and receive signed JWT |

#### `POST /auth/register`
```json
// Request Body
{
  "username": "rehan",
  "email": "rehan@example.com",
  "password": "password123"
}

// Response (201 Created)
{
  "message": "user registered successfully",
  "data": {
    "id": "5b676171-999a-4d62-8506-c0841707b118",
    "username": "rehan",
    "email": "rehan@example.com"
  }
}
```

#### `POST /auth/login`
```json
// Request Body
{
  "username": "rehan",
  "password": "password123"
}

// Response (200 OK)
{
  "message": "login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "5b676171-999a-4d62-8506-c0841707b118",
      "username": "rehan",
      "email": "rehan@example.com"
    }
  }
}
```

---

### Events

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| `GET`  | `/events` | None | Retrieve all events |
| `GET`  | `/events/:id` | None | Retrieve a single event by ID |
| `POST` | `/events/create` | JWT | Create a new event (Authenticated users) |

#### `POST /events/create`
```json
// Request Body
{
  "title": "Go Concurrency & Backend Masterclass",
  "description": "Deep dive into goroutines, channels, and row-level locking.",
  "location": "Jakarta Convention Center",
  "starts_at": "2026-10-01T09:00:00Z",
  "ends_at": "2026-10-01T17:00:00Z",
  "quota": 100
}

// Response (201 Created)
{
  "message": "event created successfully",
  "data": {
    "id": "8f2b3e41-012a-43e8-a321-c0841707b119",
    "title": "Go Concurrency & Backend Masterclass",
    "description": "Deep dive into goroutines, channels, and row-level locking.",
    "location": "Jakarta Convention Center",
    "starts_at": "2026-10-01T09:00:00Z",
    "ends_at": "2026-10-01T17:00:00Z",
    "quota": 100,
    "created_at": "2026-09-21T09:00:00Z"
  }
}
```

---

### Bookings (Upcoming)

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| `POST` | `/events/:id/book` | JWT | Book a seat for an event (Race-condition safe) |
| `GET`  | `/me/bookings`     | JWT | Retrieve all bookings for the logged-in user |

---

## Roadmap Status

- [x] Clean Feature-Oriented 3-Tier Layered Architecture
- [x] Environment & Config Management
- [x] GORM + PostgreSQL Connection & Pooling
- [x] Automated Database Migrations (`golang-migrate`)
- [x] User Authentication & Authorization (Bcrypt + JWT)
- [x] JWT Auth Middleware Protection
- [x] Event Management Domain (CRUD Operations)
- [ ] Concurrency-Safe Booking Engine (`SELECT FOR UPDATE`)
- [ ] User Booking History & Cancellation
- [ ] Goroutine & Channel Async Processing

---

## License

This project is open-source software licensed under the [MIT License](LICENSE).
