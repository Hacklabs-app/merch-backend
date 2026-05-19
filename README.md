# Merch API

A professional, domain-driven REST API for an internal school club e-commerce platform. Instead of a traditional shopping cart, **merch** uses a campaign-based pre-order system backed by an immutable digital wallet ledger.

## 🚀 Concept & Flow

This API supports a "Lipa Mdogo Mdogo" (installment) model for internal communities:
1.  **Fund Wallet:** Students top up their internal digital wallet using Paystack.
2.  **Set Goals:** Students add items from active "Campaigns" (e.g., *Q3 Club Hoodie*) to their target goals.
3.  **Commit:** Once their wallet balance meets the item cost, they commit to the purchase.
4.  **Bulk Order:** Admins use the dashboard to aggregate all committed purchases for a single bulk supplier order.
5.  **Vouchers:** Admins can generate secure, single-use vouchers to gift wallet credits to students.

## 🛠️ Tech Stack

This project strictly adheres to Clean Architecture and uses standard libraries wherever possible to ensure long-term stability and minimize dependency bloat.

| Component | Technology |
|-----------|-----------|
| **Language** | Go 1.24+ |
| **Web Framework** | Fiber v2 |
| **Database** | PostgreSQL (Neon Serverless) |
| **Schema Migrations** | `golang-migrate` |
| **Authentication** | JWT & bcrypt |
| **Observability** | Structured JSON Logging (`log/slog`) |
| **Payments** | Paystack Webhooks |
| **Deployment** | Google Cloud Run + GitHub Actions |

## 📁 Project Structure (Clean Architecture)

```
merch-backend/
├── cmd/
│   ├── api/          # Main application entry point
│   └── migrate/      # Standalone DB migration executor for CI/CD
├── internal/
│   ├── config/       # Environment variable parsing
│   ├── domain/       # Core business entities and interfaces
│   ├── handler/      # Fiber HTTP endpoints
│   ├── middleware/   # Request logging and Auth
│   ├── repository/   # PostgreSQL database operations
│   └── service/      # Core business logic (Wallets, Campaigns)
├── migrations/       # SQL up/down files for schema versioning
└── .github/          # CI/CD pipelines (Lint, Sec, Deploy)
```

## 🚦 Getting Started

### Prerequisites
- Go 1.24+
- A PostgreSQL database (e.g., Neon or local Docker)

### 1. Environment Setup
Copy the example environment file:
```bash
cp .env.example .env
```
Update the `.env` file with your database credentials:
```env
PORT=8080
ENVIRONMENT=development
DATABASE_URL=postgres://user:pass@host/dbname?sslmode=require
JWT_SECRET=super_secret_key
PAYSTACK_SECRET_KEY=sk_test_...
```

### 2. Run the Application
The application will automatically run pending database migrations on startup if `ENVIRONMENT=development`.

```bash
go run cmd/api/main.go
```

### 3. Verify
```bash
curl http://localhost:8080/api/v1/health
```

## 🔐 CI/CD & Deployments

Deployments to Google Cloud Run are handled automatically via GitHub Actions upon merging to the `main` branch. 

**The Pipeline:**
1. Code Quality (`gosec` security scanner, `go test`).
2. Database Migrations (Executed via `cmd/migrate` against the production DB).
3. Docker Build & Push to Google Container Registry.
4. Cloud Run Deployment.

---

## ✨ Contributors

Thanks to these wonderful people for bringing this project to life! 

<a href="https://github.com/Hacklabs-app/merch-backend/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Hacklabs-app/merch-backend" />
</a>
