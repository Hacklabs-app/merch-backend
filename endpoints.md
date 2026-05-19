# API Endpoints Documentation

This document outlines the API surface for the `merch` backend. All endpoints are versioned with `/api/v1`.

---

## 1. System

### Health Check
Check the API server status.

- **URL:** `/api/v1/health`
- **Method:** `GET`
- **Auth:** None

**Response:**
```json
{
  "status": "healthy",
  "service": "merch API",
  "version": "v1"
}
```

---

## 2. Authentication

### User Registration
Register a new user account.

- **URL:** `/api/v1/auth/register`
- **Method:** `POST`
- **Auth:** None

**Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "full_name": "John Doe",
  "phone_number": "+254700000000"
}
```

**Response (201 Created):**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "email": "user@example.com",
  "full_name": "John Doe",
  "phone_number": "+254700000000",
  "role": "customer",
  "created_at": "2026-05-19T..."
}
```

### User Login
Authenticate a user and return a JWT token.

- **URL:** `/api/v1/auth/login`
- **Method:** `POST`
- **Auth:** None

**Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```

**Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```
