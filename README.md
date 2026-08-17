# MirrorMate

MirrorMate is a full-stack application built with **Go, PostgreSQL, Redis, React, TypeScript, and Vite**.

The repository is structured as a monorepo containing a Go backend and a React frontend. The backend uses PostgreSQL for persistent data, Redis for caching/rate limiting, Goose for database migrations, and a layered internal architecture.

The goal of this README is to make it possible for a developer to **fork the repository, clone it locally, configure the environment, start the infrastructure, run migrations, and begin developing** without needing project-specific instructions from the original author.

---

## Table of Contents

* [Project Structure](#project-structure)
* [Tech Stack](#tech-stack)
* [Prerequisites](#prerequisites)
* [Getting Started](#getting-started)

  * [1. Fork and Clone](#1-fork-and-clone)
  * [2. Configure Environment Variables](#2-configure-environment-variables)
  * [3. Start Infrastructure](#3-start-infrastructure)
  * [4. Verify PostgreSQL](#4-verify-postgresql)
  * [5. Verify Redis](#5-verify-redis)
  * [6. Run Database Migrations](#6-run-database-migrations)
  * [7. Seed the Database](#7-seed-the-database)
  * [8. Start the Backend](#8-start-the-backend)
  * [9. Start the Frontend](#9-start-the-frontend)
* [Environment Variables](#environment-variables)
* [Database](#database)
* [Redis](#redis)
* [Backend Development](#backend-development)
* [Frontend Development](#frontend-development)
* [Project Architecture](#project-architecture)
* [Useful Commands](#useful-commands)
* [Troubleshooting](#troubleshooting)
* [Development Workflow](#development-workflow)
* [Contributing](#contributing)

---

# Project Structure

```text
mirrormate/
│
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   │   └── main.go
│   │   │
│   │   └── migrate/
│   │       ├── migrations/
│   │       │   └── *.sql
│   │       │
│   │       └── seed/
│   │           └── seed.go
│   │
│   ├── internal/
│   │   ├── auth/
│   │   ├── db/
│   │   ├── env/
│   │   ├── mailer/
│   │   ├── ratelimiter/
│   │   └── store/
│   │
│   ├── scripts/
│   │   └── db_init.sql
│   │
│   └── go.mod
│
├── frontend/
│   ├── public/
│   ├── src/
│   ├── package.json
│   ├── package-lock.json
│   ├── vite.config.ts
│   └── tsconfig*.json
│
├── .envrc.example
├── .gitignore
├── .air.toml
├── docker-compose.yml
└── README.md
```

The backend separates executable entry points from application internals:

* `cmd/api` — starts the API server.
* `cmd/migrate` — database migration tooling.
* `cmd/migrate/migrations` — Goose migration files.
* `cmd/migrate/seed` — database seed functionality.
* `internal/auth` — authentication-related functionality.
* `internal/db` — database connection functionality.
* `internal/env` — environment configuration.
* `internal/mailer` — email functionality.
* `internal/ratelimiter` — rate-limiting functionality.
* `internal/store` — data-access/repository layer.

---

# Tech Stack

## Backend

* Go 1.25.5
* PostgreSQL 16.3
* Redis 6.2
* Goose
* Docker / Docker Compose
* REST API
* Repository/store-based data access
* JWT authentication
* Redis-backed functionality
* Email integration

## Frontend

* React 19
* TypeScript
* Vite
* ESLint

## Infrastructure

* Docker Compose
* PostgreSQL
* Redis
* Redis Commander

---

# Prerequisites

Install the following before starting development.

### Required

* Git
* Go 1.25.5
* Node.js
* npm
* Docker
* Docker Compose
* Goose

You can verify your installations with:

```bash
git --version
go version
node --version
npm --version
docker --version
docker compose version
goose -version
```

The backend currently targets Go `1.25.5`.

---

# Getting Started

Follow these steps when setting up MirrorMate for the first time.

---

## 1. Fork and Clone

Fork the repository on GitHub and then clone your fork:

```bash
git clone https://github.com/<your-username>/mirrormate.git
cd mirrormate
```

If you are working directly from the original repository:

```bash
git clone https://github.com/Ayushmangit/mirrormate.git
cd mirrormate
```

---

## 2. Configure Environment Variables

MirrorMate provides an example environment file:

```text
.envrc.example
```

Create your local environment configuration from it:

```bash
cp .envrc.example .envrc
```

Then edit `.envrc`:

```bash
nvim .envrc
```

or use your preferred editor.

### Using direnv

If you use `direnv`, allow the environment after creating `.envrc`:

```bash
direnv allow
```

Verify that the variables are available:

```bash
env | grep -E 'GOOSE|DB_|REDIS|ADDR|FRONTEND|VITE'
```

### Important

Never commit your real `.envrc` or any API keys, passwords, JWT secrets, or other credentials.

The repository intentionally provides `.envrc.example` so developers can create their own local configuration.

---

# Environment Variables

The example environment file currently contains the following configuration:

```bash
FROM_EMAIL=
SENDGRID_API_KEY=

ENV=

GOOSE_DBSTRING=
GOOSE_DRIVER=postgres

DB_ADDR=
ADDR=:8080

EXTERNAL_URL=
FRONTEND_URL=

DB_MAX_OPEN_CONNS=30
DB_MAX_IDLE_CONNS=30
DB_MAX_IDLE_TIME=15m

REDIS_ADDR=""
REDIS_DB=0
REDIS_PASSWORD=""
REDIS_ENABED=

VITE_BASE_URL=https://localhost:5173/v1
```

Fill in the values appropriate for your local environment.

For local Docker development, the important database and Redis addresses should account for the ports exposed by `docker-compose.yml`.

---

# 3. Start Infrastructure

MirrorMate uses Docker Compose for its local PostgreSQL and Redis infrastructure.

From the repository root:

```bash
docker compose up -d
```

Check that everything is running:

```bash
docker compose ps
```

You should have:

```text
PostgreSQL
Redis
Redis Commander
```

running.

To view logs:

```bash
docker compose logs
```

Or for a specific service:

```bash
docker compose logs db
docker compose logs redis
docker compose logs redis-commander
```

---

# PostgreSQL

The Docker Compose configuration creates:

```text
Database: mirrormate
User:     admin
Password: adminpassword
```

The PostgreSQL container listens on port `5432` internally but is exposed to the host on port `5433`.

Therefore, when connecting from your host machine:

```text
localhost:5433
```

is the correct PostgreSQL address.

The local connection string is:

```text
postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable
```

Do not use port `5432` from your host unless you have another PostgreSQL instance configured there.

---

# 4. Verify PostgreSQL

After starting Docker:

```bash
docker compose ps
```

The PostgreSQL health check should eventually report the service as healthy.

You can also test the connection directly:

```bash
psql "postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable"
```

If `psql` connects successfully, PostgreSQL is ready.

You can also check the container logs:

```bash
docker compose logs db
```

---

# Database Initialization

The repository contains:

```text
backend/scripts/db_init.sql
```

which contains the database initialization SQL.

It currently creates the `mirrormate` database and enables the PostgreSQL `citext` extension.

The normal development flow should use the Docker Compose PostgreSQL instance followed by the application's migration system.

---

# 5. Run Database Migrations

MirrorMate uses Goose for database migrations.

The migration files are located at:

```text
backend/cmd/migrate/migrations/
```

Before running Goose, make sure your environment contains:

```bash
GOOSE_DBSTRING=postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable
GOOSE_DRIVER=postgres
```

From the repository root:

```bash
cd backend
```

Then run:

```bash
goose -dir ./cmd/migrate/migrations postgres "$GOOSE_DBSTRING" up
```

Alternatively, if you prefer passing the connection string directly:

```bash
goose -dir ./cmd/migrate/migrations postgres \
  "postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable" \
  up
```

Check migration status:

```bash
goose -dir ./cmd/migrate/migrations postgres "$GOOSE_DBSTRING" status
```

To roll back the most recent migration:

```bash
goose -dir ./cmd/migrate/migrations postgres "$GOOSE_DBSTRING" down
```

### Important

Always run migrations against the correct database.

For the Docker Compose setup, use:

```text
localhost:5433
```

not:

```text
localhost:5432
```

---

# 6. Seed the Database

The repository contains a seed command under:

```text
backend/cmd/migrate/seed/
```

If you need development/test data, inspect the seed implementation and run the seed command according to its current requirements.

From the `backend` directory:

```bash
go run ./cmd/migrate/seed
```

If the seed command requires specific environment variables or arguments, configure them before running it.

---

# 7. Start the Backend

From:

```text
backend/
```

run:

```bash
go run ./cmd/api
```

The API address is controlled by:

```bash
ADDR=:8080
```

By default, the backend should therefore be available at:

```text
http://localhost:8080
```

You can also build the API:

```bash
go build ./cmd/api
```

---

# Backend Development with Air

The repository contains:

```text
.air.toml
```

for live-reloading development.

If Air is installed:

```bash
air
```

Run it from the directory where the Air configuration is expected.

If Air is not installed, install it using the official Air installation instructions or run the API directly with:

```bash
go run ./cmd/api
```

---

# 8. Start the Frontend

Open another terminal.

From the repository root:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the Vite development server:

```bash
npm run dev
```

Vite will display the local development URL in the terminal.

The frontend uses Vite and TypeScript.

---

# Frontend Commands

## Development

```bash
npm run dev
```

## Production build

```bash
npm run build
```

## Preview production build

```bash
npm run preview
```

## Lint

```bash
npm run lint
```

---

# Redis

Redis is included in Docker Compose.

The Redis container listens on:

```text
6379
```

internally.

It is exposed to your host machine on:

```text
6380
```

Therefore:

### From your host machine

```text
localhost:6380
```

### From another Docker Compose service

```text
redis:6379
```

Test Redis from your host:

```bash
redis-cli -p 6380 ping
```

Expected response:

```text
PONG
```

---

# Redis Commander

Redis Commander is included for development and Redis inspection.

It is exposed on:

```text
http://localhost:8082
```

Open that address in your browser after starting Docker Compose.

Redis Commander connects to the Redis Compose service using:

```text
redis:6379
```

inside the Docker network.

---

# Docker Commands

Start all services:

```bash
docker compose up -d
```

Stop services:

```bash
docker compose down
```

Stop services and remove the database volume:

```bash
docker compose down -v
```

> Warning: removing the volume deletes the local PostgreSQL data.

Restart services:

```bash
docker compose restart
```

View running containers:

```bash
docker compose ps
```

View logs:

```bash
docker compose logs -f
```

---

# Project Architecture

The backend follows a separation between application entry points, infrastructure, and domain/data-access functionality.

```text
HTTP Request
     │
     ▼
   Router
     │
     ▼
  Handlers
     │
     ▼
 Application / Business Logic
     │
     ▼
   Store Layer
     │
     ├───────────────┐
     ▼               ▼
PostgreSQL         Redis
```

The main internal packages are:

```text
internal/
├── auth/
├── db/
├── env/
├── mailer/
├── ratelimiter/
└── store/
```

### `internal/store`

Contains database access and repository/store functionality.

Database-related operations should generally be kept here rather than placing SQL/database logic directly inside HTTP handlers.

### `internal/db`

Responsible for database connection/setup functionality.

### `internal/auth`

Contains authentication-related functionality.

### `internal/ratelimiter`

Contains rate-limiting functionality and Redis-related rate-limiter behavior.

### `internal/mailer`

Contains email-related functionality.

### `internal/env`

Contains environment/configuration handling.

---

# Database Migrations Workflow

When modifying the database schema:

1. Do not modify an already-applied migration.
2. Create a new migration.
3. Put it inside:

```text
backend/cmd/migrate/migrations/
```

4. Apply it locally.
5. Test the application against the new schema.
6. Commit the migration together with the code that requires it.

Example:

```bash
goose -dir ./cmd/migrate/migrations \
  postgres "$GOOSE_DBSTRING" \
  create add_example_table sql
```

Then edit the generated migration.

Apply it:

```bash
goose -dir ./cmd/migrate/migrations \
  postgres "$GOOSE_DBSTRING" \
  up
```

---

# Working on the Backend

From the repository root:

```bash
cd backend
```

Install/download Go dependencies:

```bash
go mod tidy
```

Run the application:

```bash
go run ./cmd/api
```

Run tests:

```bash
go test ./...
```

Format the code:

```bash
gofmt -w .
```

Check the project:

```bash
go vet ./...
```

Build the application:

```bash
go build ./cmd/api
```

---

# Working on the Frontend

From the repository root:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start development:

```bash
npm run dev
```

Run linting:

```bash
npm run lint
```

Build:

```bash
npm run build
```

---

# Full Local Development Setup

For a new developer, the shortest setup path is:

### Terminal 1 — Infrastructure

```bash
git clone https://github.com/<your-username>/mirrormate.git
cd mirrormate

cp .envrc.example .envrc
```

Configure `.envrc`, then:

```bash
direnv allow
docker compose up -d
```

Verify:

```bash
docker compose ps
```

### Terminal 2 — Database

```bash
cd mirrormate/backend
```

Run:

```bash
goose -dir ./cmd/migrate/migrations \
  postgres "$GOOSE_DBSTRING" \
  up
```

### Terminal 3 — Backend

```bash
cd mirrormate/backend
go run ./cmd/api
```

### Terminal 4 — Frontend

```bash
cd mirrormate/frontend
npm install
npm run dev
```

At this point:

```text
Frontend        → Vite development server
Backend         → localhost:8080
PostgreSQL      → localhost:5433
Redis           → localhost:6380
Redis Commander → localhost:8082
```

---

# Troubleshooting

## PostgreSQL connection refused

If you see:

```text
connection refused
```

check:

```bash
docker compose ps
```

Make sure PostgreSQL is running.

Remember that this project maps:

```text
5433 → PostgreSQL 5432
```

Therefore your local connection string should use:

```text
localhost:5433
```

not:

```text
localhost:5432
```

---

## Redis connection refused

Check:

```bash
docker compose ps
```

Then:

```bash
redis-cli -p 6380 ping
```

Expected:

```text
PONG
```

If the Go application runs directly on your host:

```text
localhost:6380
```

If the application runs inside Docker Compose:

```text
redis:6379
```

---

## Goose cannot connect to PostgreSQL

First verify:

```bash
psql "postgres://admin:adminpassword@localhost:5433/mirrormate?sslmode=disable"
```

If this fails, the problem is PostgreSQL connectivity rather than Goose.

Then verify:

```bash
echo "$GOOSE_DBSTRING"
```

It should point to the correct local PostgreSQL instance.

---

## Migrations fail because the database does not exist

Check that the PostgreSQL container was initialized correctly:

```bash
docker compose logs db
```

The Compose configuration creates:

```text
mirrormate
```

with:

```text
admin
adminpassword
```

If you previously removed/recreated containers and are working with an old volume, inspect your local database state before deleting anything.

For a completely fresh development database:

```bash
docker compose down -v
docker compose up -d
```

Then run the migrations again.

> `docker compose down -v` deletes local database data.

---

## Frontend cannot reach backend

Make sure the backend is running:

```bash
cd backend
go run ./cmd/api
```

Then verify that the API is listening on:

```text
localhost:8080
```

Check your frontend API configuration and ensure it points to the backend rather than an unavailable remote URL.

---

# Development Workflow

When starting work on a new feature:

```text
1. Pull latest changes
        ↓
2. Create a feature branch
        ↓
3. Start PostgreSQL + Redis
        ↓
4. Apply migrations
        ↓
5. Start backend
        ↓
6. Start frontend
        ↓
7. Implement feature
        ↓
8. Add/update tests
        ↓
9. Run formatting/linting
        ↓
10. Test locally
        ↓
11. Commit changes
        ↓
12. Push feature branch
        ↓
13. Open Pull Request
```

Recommended branch naming:

```text
feature/<feature-name>
fix/<bug-name>
refactor/<area>
docs/<documentation-change>
```

Examples:

```text
feature/user-profile
feature/post-comments
fix/auth-token-expiry
refactor/store-layer
docs/update-readme
```

---

# Before Opening a Pull Request

Run the backend checks:

```bash
cd backend

gofmt -w .
go vet ./...
go test ./...
go build ./cmd/api
```

Run the frontend checks:

```bash
cd frontend

npm run lint
npm run build
```

Make sure database migrations work on a clean local database.

Also verify that:

* No secrets are committed.
* `.envrc` is not committed.
* New database changes have migrations.
* Existing migrations have not been modified unnecessarily.
* API changes are reflected in the frontend where required.
* Docker services start successfully.

---

# Security

Never commit:

```text
.env
.envrc
API keys
JWT secrets
database passwords
SendGrid API keys
private credentials
```

Use `.envrc.example` as the template for local configuration.

If a secret is accidentally committed, removing the file from the latest commit is not sufficient if the secret remains in Git history. Rotate/revoke the compromised credential immediately.

---

# Useful Development URLs

| Service         | Local Address                   |
| --------------- | ------------------------------- |
| Backend API     | `http://localhost:8080`         |
| PostgreSQL      | `localhost:5433`                |
| Redis           | `localhost:6380`                |
| Redis Commander | `http://localhost:8082`         |
| Frontend        | Vite URL shown by `npm run dev` |

---

# Contributing

Contributions are welcome.

A typical contribution should:

1. Fork the repository.
2. Create a feature/fix branch.
3. Make your changes.
4. Add or update tests where appropriate.
5. Run backend and frontend checks.
6. Make sure the project starts from a clean environment.
7. Commit your changes.
8. Push the branch.
9. Open a Pull Request.

Please keep changes focused and avoid mixing unrelated refactoring with feature work.

---

# License

Add the project's license information here once a license has been selected and added to the repository.

---

# Author

**Ayushman Chauhan**

GitHub: [@Ayushmangit](https://github.com/Ayushmangit)

Repository: [Ayushmangit/mirrormate](https://github.com/Ayushmangit/mirrormate)

# mirrormate
# mirrormate
