# Technical Design Document (TDD)
## Project: Expense Tracker API
**Version:** 2.0
**Tech Stack:** Go, Chi, PostgreSQL, Docker
**Architecture:** Hybrid Modular Clean Architecture

## 1. High-Level Architecture
Feature-first modular architecture with handler-service-repository inside each feature module.

## 2. Project Structure
```bash
expense-tracker/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   ├── expenses/
│   ├── categories/
│   ├── analytics/
│   ├── platform/
│   │   ├── config/
│   │   ├── database/
│   │   ├── logger/
│   │   └── auth/
│   ├── middleware/
│   └── shared/
├── migrations/
├── docs/
└── scripts/
```

## 3. Request Lifecycle
Request -> Router -> Middleware -> Handler -> Service -> Repository -> Database

## 4. Database Design
PostgreSQL with users, categories, expenses tables.

## 5. API Design
- POST /auth/signup
- POST /auth/login
- CRUD /categories
- CRUD /expenses
- GET /analytics/monthly
