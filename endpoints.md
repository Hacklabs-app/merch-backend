# API Endpoints Documentation

This document outlines the available REST API endpoints for the merch backend. All endpoints are prefixed with the current API version (e.g., `/api/v1`).

---

## System

### Health Check

Verifies that the API server is running and responding to requests.

*   **URL:** `/api/v1/health`
*   **Method:** `GET`
*   **Authentication Required:** No

#### Success Response

*   **Code:** `200 OK`
*   **Content-Type:** `application/json`
*   **Body:**

```json
{
  "service": "merch API",
  "status": "healthy",
  "version": "v1"
}
```

#### Error Responses

*   **Code:** `404 Not Found` (If accessing an undefined route like `/api/v1/invalid`)
*   **Body:**

```json
{
  "error": "Route not found"
}
```

*Note: This document will be updated as Phase 2 (Identity & Authentication) and Phase 3 (Wallets & Campaigns) are implemented.*