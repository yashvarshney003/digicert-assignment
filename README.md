# DigiCert Public Library API

A RESTful API built using **Go (Gin + GORM)** to manage books in a public library system.

## Features

- Add, update, delete, and retrieve books
- PostgreSQL for persistent storage
- Dockerized setup for easy deployment
- Wait-for-it script to ensure DB readiness before app starts
- Unit tests that block app startup if any test fails

---

## Tech Stack

- **Go 1.24**
- **Gin** (web framework)
- **GORM** (ORM)
- **PostgreSQL** (database)
- **Docker + Docker Compose**

---

## Getting Started

### Clone the repository

```bash
git clone https://github.com/your-username/public-library-api.git
cd public-library-api
```

---

###  Run Tests Locally (Optional)

Ensure Go is installed and run tests manually:

```bash
cd public_library
go test -v ./...
```

> The app is configured to **run tests automatically inside Docker**. If any test fails, the app **will not start**.

---

### Run with Docker Compose

This will:
- Start PostgreSQL container
- Wait until DB is ready using `wait-for-it.sh`
- Run Go test cases
- Start the app only if tests pass

```bash
docker-compose up --build
```

---

## Postman Collection

Use the following collection to explore the API endpoints:

[Postman Link](https://web.postman.co/workspace/My-Workspace~4a725b07-e3a5-4f97-81ae-919f1f646a7f/collection/8096381-ff39f516-2951-4d2c-a410-dc6d0bd86269?action=share&source=copy-link&creator=8096381)

---

## 📂 Project Structure

```bash
public_library/
│
├── cmd/               # App entry point
├── config/            # Environment configuration
├── controllers/       # Route handlers
├── database/          # DB initialization
├── models/            # Data models
├── routes/            # Route definitions
├── wait-for-it.sh     # DB readiness script
├── Dockerfile         # App image definition
├── docker-compose.yml # Multi-container setup
└── go.mod / go.sum    # Go module dependencies
```

---

## 🧪 Test Behavior

If you're using Docker Compose, tests will be run **before** the app starts:
- **Pass**: App starts
- **Fail**: App won’t start (container exits)

---

## Notes

- `wait-for-it.sh` is used to wait for PostgreSQL container readiness.
- Make sure port `5432` is free on your machine.

---
