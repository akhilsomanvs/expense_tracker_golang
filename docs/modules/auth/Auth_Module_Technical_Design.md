# Technical Design Document — Authentication Module

## Module: Authentication
**Version:** 1.0
**Project:** Expense Tracker API
**Architecture:** Hybrid Modular Clean Architecture

## Responsibilities
- User registration
- User login
- Password hashing
- JWT generation
- JWT validation

## Module Structure
internal/auth/
- handler.go
- service.go
- repository.go
- model.go
- dto.go
- validator.go
- routes.go

## API Design
POST /auth/signup
POST /auth/login

## Security
- bcrypt password hashing
- JWT authentication

## Database
users table with id, name, email, password_hash

## Definition of Done
- Signup works
- Login works
- JWT works
- Validation works
- Tests added
