# Technical Design Document (TDD)
## Project: Expense Tracker API
**Version:** 1.0  
**Tech Stack:** Go, Chi, PostgreSQL, Docker  
**Architecture:** Layered / Clean-ish Architecture  

---

# 1. High-Level Architecture

The application follows a layered architecture.

```text
Client
   ↓
Router
   ↓
Middleware
   ↓
Handlers
   ↓
Service Layer
   ↓
Repository Layer
   ↓
PostgreSQL
```

## Responsibilities

### Router
- Route registration
- Route grouping
- Middleware binding

### Middleware
- Authentication
- Logging
- Recovery
- Request ID

### Handlers
- Parse request
- Validate input
- Return HTTP response

### Service Layer
- Business logic
- Validation beyond request schema
- Application rules

### Repository Layer
- Database queries
- CRUD operations

---

# 2. Project Structure

```bash
expense-tracker/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── service/
│   ├── auth/
│   └── database/
├── migrations/
├── docs/
├── scripts/
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

# 3. Request Lifecycle

```text
Request
   ↓
Router
   ↓
RequestID Middleware
   ↓
Logger Middleware
   ↓
Auth Middleware
   ↓
Expense Handler
   ↓
Expense Service
   ↓
Expense Repository
   ↓
Database
```

---

# 4. Database Design

Database: PostgreSQL

## Users Table

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

## Categories Table

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## Expenses Table

```sql
CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id),
    title VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    note TEXT,
    expense_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## Recommended Indexes

```sql
CREATE INDEX idx_expenses_user_id ON expenses(user_id);
CREATE INDEX idx_expenses_category_id ON expenses(category_id);
CREATE INDEX idx_expenses_date ON expenses(expense_date);
```

---

# 5. API Design

## Authentication
- POST /auth/signup
- POST /auth/login

## Category APIs
- POST /categories
- GET /categories
- PUT /categories/{id}
- DELETE /categories/{id}

## Expense APIs
- POST /expenses
- GET /expenses
- GET /expenses/{id}
- PUT /expenses/{id}
- DELETE /expenses/{id}

## Analytics APIs
- GET /analytics/monthly
