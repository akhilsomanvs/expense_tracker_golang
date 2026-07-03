# Technical Design Document — Authentication Module

## Module: Authentication

**Version:** 2.0
**Project:** Expense Tracker API
**Architecture:** Hybrid Modular Clean Architecture
**Module Type:** AppModule

---

# 1. Overview

The Authentication module is responsible for:

* User registration
* User login
* Password hashing
* JWT generation
* JWT validation
* Route protection

This module ensures that only authenticated users can access protected APIs.

---

# 2. Responsibilities

Primary responsibilities:

* Signup new users
* Login existing users
* Validate credentials
* Generate JWT token
* Verify JWT token
* Provide middleware support for protected routes

---

# 3. Module Architecture

Authentication module follows modular clean architecture:

```text
Request
   ↓
Router
   ↓
Middleware
   ↓
Auth Handler
   ↓
Auth Service
   ↓
Auth Repository
   ↓
PostgreSQL
```

---

# 4. Module Structure

```bash
internal/auth/
├── module.go
├── handler.go
├── service.go
├── repository.go
├── model.go
├── dto.go
├── validator.go
└── errors.go
```

---

# 5. Contracts

---

## AppRoutes

```go
type AppRoutes interface {
	RegisterRoutes(router chi.Router)
}
```

---

## Named

```go
type Named interface {
	Name() string
}
```

---

## AppModule

```go
type AppModule interface {
	Named
	AppRoutes
}
```

All modules must implement:

* Name()
* RegisterRoutes()

---

# 6. Auth Module Implementation

---

## module.go

```go
type AuthModule struct {
	handler *Handler
}
```

---

## Constructor

```go
func NewModule(handler *Handler) *AuthModule {
	return &AuthModule{
		handler: handler,
	}
}
```

---

## Name

```go
func (m *AuthModule) Name() string {
	return "auth"
}
```

---

## Route Registration

```go
func (m *AuthModule) RegisterRoutes(router chi.Router) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", m.handler.Signup)
		r.Post("/login", m.handler.Login)
	})
}
```

---

## Compile-time Check

```go
var _ contracts.AppModule = (*AuthModule)(nil)
```

---

# 7. Layer Responsibilities

---

## Handler Layer

Responsibilities:

* Parse JSON request
* Call validator
* Call service
* Return HTTP response

Example:

```go
type Handler struct {
	service *Service
}
```

Methods:

```go
func (h *Handler) Signup(w http.ResponseWriter, r *http.Request)
func (h *Handler) Login(w http.ResponseWriter, r *http.Request)
```

---

## Service Layer

Responsibilities:

* Business logic
* Password hashing
* Password verification
* JWT generation

Example:

```go
type Service struct {
	repo *Repository
}
```

Methods:

```go
func (s *Service) Signup(req SignupRequest) error
func (s *Service) Login(req LoginRequest) (string, error)
```

---

## Repository Layer

Responsibilities:

* Insert user
* Fetch user by email
* Fetch user by ID

Example:

```go
type Repository struct {
	db *sql.DB
}
```

Methods:

```go
CreateUser(user User) error
GetUserByEmail(email string) (*User, error)
```

---

# 8. Domain Model

---

## User Model

```go
type User struct {
	ID           int
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
```

---

# 9. DTOs

---

## Signup Request

```go
type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
```

---

## Login Request

```go
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
```

---

## Login Response

```go
type LoginResponse struct {
	Token string `json:"token"`
}
```

---

# 10. API Design

---

# Signup

### Endpoint

```http
POST /auth/signup
```

---

### Request

```json
{
  "name": "Akhil",
  "email": "akhil@gmail.com",
  "password": "password123"
}
```

---

### Success Response

```json
{
  "message": "User registered successfully"
}
```

---

---

# Login

### Endpoint

```http
POST /auth/login
```

---

### Request

```json
{
  "email": "akhil@gmail.com",
  "password": "password123"
}
```

---

### Success Response

```json
{
  "token": "jwt_token_here"
}
```

---

# 11. Validation Rules

---

## Signup Validation

* Name required
* Email required
* Email valid
* Password minimum 8 chars

---

## Login Validation

* Email required
* Password required

---

# 12. Request Flow

---

# Signup Flow

```text
Request
→ Handler
→ Validator
→ Service
→ Check existing user
→ Hash password
→ Repository
→ Database
→ Response
```

---

# Login Flow

```text
Request
→ Handler
→ Validator
→ Service
→ Repository
→ Verify password
→ Generate JWT
→ Response
```

---

# 13. Database Design

Uses users table.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

---

# 14. Security Design

---

## Password Hashing

Use bcrypt.

Package:

```bash
golang.org/x/crypto/bcrypt
```

Hash password:

```go
hashed, err := bcrypt.GenerateFromPassword(
    []byte(password),
    bcrypt.DefaultCost,
)
```

Verify password:

```go
bcrypt.CompareHashAndPassword(
    []byte(hash),
    []byte(password),
)
```

---

## JWT Design

Token contains:

* user_id
* email
* expiry

Payload:

```json
{
  "user_id": 1,
  "email": "akhil@gmail.com",
  "exp": 1710000000
}
```

Recommended expiry:

* 24 hours

---

# 15. Middleware Integration

Protected routes use auth middleware.

Flow:

```text
Request
→ Auth Middleware
→ Extract Token
→ Validate Token
→ Extract Claims
→ Store User Context
→ Next Handler
```

---

# 16. Error Handling

---

Standard Errors:

| Error           | HTTP |
| --------------- | ---: |
| Invalid request |  400 |
| Unauthorized    |  401 |
| Duplicate email |  409 |
| Internal error  |  500 |

---

# 17. Logging

Log:

* signup success
* login success
* login failure
* token validation failure

Example:

```json
{
  "level": "INFO",
  "message": "login successful",
  "user_id": 1
}
```

---

# 18. Testing Requirements

---

## Signup Tests

* valid signup
* duplicate email
* invalid email
* short password

---

## Login Tests

* valid login
* invalid password
* unknown user

---

## JWT Tests

* valid token
* expired token
* invalid token

---

# 19. Definition of Done

Authentication module is complete when:

* Signup works
* Login works
* Password hashing works
* JWT generation works
* JWT validation works
* Middleware works
* Validation works
* Tests added
* Logs added
* Documentation updated
