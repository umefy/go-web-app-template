SHELL := /bin/bash
.PHONY: default build run clean check fmt test lint help tidy regen_gorm migration_up migration_down migration_create migration_reset generate regen_openapi regen_proto regen_graphql docker_compose_up docker_compose_down seed_database

APP_NAME=webapp
MAIN_PATH=cmd/server/main.go
BUILD_FOLDER=./bin
PRODUCTION_ENTRY=$(BUILD_FOLDER)/$(APP_NAME)
AIR_TMP_FOLDER=./tmp
ENVRC_FILE ?=.envrc

default: dev

print:
	echo "ENV_FILE is $(ENVRC_FILE)"

docker_compose_up:
	docker compose up -d

docker_compose_down:
	docker compose down

dev:
	@make docker_compose_up
	@make migration_up
	@trap 'make docker_compose_down' INT TERM; \
	air; 

build:
	@echo "⏱️ building code now..."
	@make clean
	go build -o $(PRODUCTION_ENTRY) $(ARGS) $(MAIN_PATH)
	@echo "✅ building finish"

start:
	@echo "starting running production build now..."
	$(PRODUCTION_ENTRY) --env=prod $(ARGS)

mockery:
	@echo "⏱️ generating mock now..."
	@rm -rf mocks
	@mockery
	@echo "✅ generating finish"

clean:
	@echo "🗑️ cleaning old build now..."
	@rm -rf $(BUILD_FOLDER)
	@rm -rf $(AIR_TMP_FOLDER)
	@echo "✅ cleaning finish"

check: tidy fmt lint test

generate: docker_compose_up regen_openapi regen_proto migration_up regen_gorm  mockery regen_graphql docker_compose_down

fmt:
	@echo "⏱️ formatting code now..."
	go fmt ./...
	@echo "✅ formatting finish"

goimports:
	@echo "⏱️ running goimports now..."
	goimports -w .
	@echo "✅ finishing goimports..."

test:
	@echo "⏱️ running tests now... "
	go test -race -parallel=4 -timeout 30s -cover $(ARGS) ./internal/...
	@echo "✅ passing all tests."

lint:
	@echo "⏱️ running linting now..."
	golangci-lint run $(ARGS)
	@echo "✅ passing linting..."

tidy:
	@echo "⏱️ go mod tidy now..."
	go mod tidy
	@echo "✅ finishing tidy..."

regen_openapi:
	@echo "⏱️ running openapi to go code now..."
	./scripts/regen_openapi.sh
	@echo "✅ finishing openapi to go code..."

regen_proto:
	@echo "⏱️ running regen proto now..."
	./scripts/regen_proto.sh
	@echo "✅ finishing regen proto..."

regen_gorm:
	@echo "⏱️ running regen gorm now..."
	source $(ENVRC_FILE) && go run gorm/generate.go
	@echo "✅ finishing regen gorm..."

regen_graphql:
	@echo "⏱️ running regen graphql now..."
	gqlgen generate
	@echo "✅ finishing regen graphql..."

migration_up:
	source $(ENVRC_FILE) && goose up

migration_down:
	source $(ENVRC_FILE) && goose down

migration_create:
	@if [ -z "$(migration_name)" ]; then \
		echo "Error: migration_name is required! Usage: make migration_create migration_name=your_migration"; \
		exit 1; \
	fi
	source $(ENVRC_FILE) && goose create $(migration_name) sql

migration_reset:
	source $(ENVRC_FILE) && goose reset

seed_database:
	@echo "⏱️ seeding database now..."
	@make migration_reset
	@make migration_up
	source $(ENVRC_FILE) && go run cmd/seed/database/*.go
	@echo "✅ seeding database finish"

help:
	@echo "make - running go code with go run"
	@echo "make dev - running go code with dev environment"
	@echo "make build - build go code to binary"
	@echo "make start - running production build"
	@echo "make clean - clean old binary build"
	@echo "make check - generate required files, formatting, testing and running lint"
	@echo "make generate - generate all required files"
	@echo "make test - running go test"
	@echo "make fmt - formatting go code"
	@echo "make lint - running golangci lint"
	@echo "make tidy - install all dependencies"
	@echo "make regen_openapi - generating Go models from OpenAPI specification"
	@echo "make regen_proto - regenerating source code from proto"
	@echo "make regen_gorm - regenerating gorm models"
	@echo "make regen_graphql - regenerating graphql models and resolvers"
	@echo "make migration_up - running database migrations up"
	@echo "make migration_down - running database migrations down"
	@echo "make migration_create - creating a new database migration"
	@echo "make migration_reset - resetting all database migrations"
	@echo "make docker_compose_up - starting docker compose"
	@echo "make docker_compose_down - stopping docker compose"
	@echo "make seed_database - seeding database"