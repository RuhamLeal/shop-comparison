#!make
include .env

GOOS ?= $(go env GOOS)
GOARCH ?= $(go env GOARCH)

BINARY_NAME = api

ifeq ($(GOOS),windows)
	BINARY = $(BINARY_NAME).exe
	RM_FILE = if exist build\$(BINARY) del /F /Q build\$(BINARY)
else
	BINARY = $(BINARY_NAME)
	RM_FILE = rm -f ./build/$(BINARY)
endif

dbDriver=sqlite3
migrationPath=./internal/infra/sqlite/migrations

prod: prod-dependencies prod-setup prod-build

dev: start-air

dev-dependencies:
	@echo "\n🟠 Installing project dependencies..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
	@go install github.com/swaggo/swag/cmd/swag@v1.16.4
	@go install github.com/air-verse/air@v1.62.0
	@go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
	@echo "\n🟢 Development dependencies installed gracefully"

prod-dependencies:
	@echo "\n🟠 Installing project production dependencies..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.29.0
	@go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
	@echo "\n🟢 Production dependencies installed gracefully"

prod-setup: migration-up

dev-setup: migration-up

prod-build: clean sqlc
	@echo "\n🟠 Compiling api for production"
	go build -ldflags="-w -s" -o build/$(BINARY) cmd/api/main.go
	@echo "🟢 Compilation done gracefully"

dev-build: clean swag sqlc
	@echo "\n🟠 Compiling api for development"
	go build -gcflags="all=-N -l" -o build/$(BINARY) cmd/api/main.go
	@echo "🟢 Compilation done gracefully"

.PHONY: sqlc
sqlc:
	@echo "\n🟠 Generating database queries and models using sqlc..."
	@sqlc generate
	@echo "\n🟢 Queries and models generated gracefully using sqlc"

create-migration:
	@GOOSE_DRIVER=$(dbDriver) GOOSE_DBSTRING=${GOOSE_DB_DSN} goose -dir=$(migrationPath) create $(name) sql
	@echo "🟢 Migration file created gracefully"

migration-up:
	@echo "\n🟠 Starting process to run pending migrations"
	@goose $(dbDriver) ${GOOSE_DB_DSN} -dir=$(migrationPath) up
	@echo "🟢 Process to run pending migrations done gracefully"

migration-down:
	@goose $(dbDriver) ${GOOSE_DB_DSN} -dir=$(migrationPath) down

migration-reset:
	@goose $(dbDriver) ${GOOSE_DB_DSN} -dir=$(migrationPath) reset

migration-rollback:
	@echo "\n🟠 Starting process to run migrations rollback"
	@goose $(dbDriver) ${GOOSE_DB_DSN} -dir=$(migrationPath) down-to 0
	@echo "🟢 Process to run migrations rollback done gracefully"

swag:
	@echo "\n🟠 Generating api docs"
	swag init --parseInternal --parseDependency -g cmd/api/main.go
	@echo "🟢 Docs generated gracefully"

start-air:
	@echo "\n🟠 Starting server with Air..."
	air -c .air.toml
	@echo "\n🟢 Air started gracefully"

clean:
	$(RM_FILE)