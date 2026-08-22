APP_NAME := boardgo
MAIN_PATH := ./cmd/api
BIN_DIR := ./tmp
BIN_PATH := $(BIN_DIR)/$(APP_NAME)

DB_URL ?= $(DATABASE_URL)
MIGRATIONS_DIR := ./migrations

.PHONY: help
help:
	@echo "Available commands:"
	@echo ""
	@echo "  make dev          Start Air development server"
	@echo "  make run          Run application"
	@echo "  make build        Build application"
	@echo "  make test         Run tests"
	@echo "  make clean        Remove build artifacts"
	@echo ""
	@echo "  make migrate-up   Run all migrations"
	@echo "  make migrate-down Rollback last migration"
	@echo "  make migrate-status Show migration status"
	@echo "  make migrate-reset Rollback all migrations"
	@echo "  make migrate-create name=xxx Create migration"
	@echo ""

.PHONY: dev
dev:
	air

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_PATH) $(MAIN_PATH)

.PHONY: test
test:
	go test ./...

.PHONY: test-race
test-race:
	go test -race ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# ==========================================
# Goose migrations
# ==========================================

.PHONY: migrate-up
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

.PHONY: migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

.PHONY: migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" reset

.PHONY: migrate-create
migrate-create:
ifndef name
	$(error name is required. Example: make migrate-create name=create_users)
endif
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
