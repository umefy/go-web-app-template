SHELL := /bin/bash

APP_NAME=webapp
MAIN_PATH=cmd/server/main.go
BUILD_FOLDER=./bin
PRODUCTION_ENTRY=$(BUILD_FOLDER)/$(APP_NAME)
AIR_TMP_FOLDER=./tmp
ENVRC_FILE ?=.envrc

.PHONY: default
default: dev

.PHONY: print
print:
	echo "ENV_FILE is $(ENVRC_FILE)"

.PHONY: docker_compose_up
docker_compose_up:
	docker compose up -d

.PHONY: docker_compose_down
docker_compose_down:
	docker compose down

.PHONY: dev
dev:
	@make docker_compose_up
	@make migration_up
	@trap 'make docker_compose_down' INT TERM; \
	air; 

.PHONY: build
build:
	@echo "⏱️ building code now..."
	@make clean
	go build -o $(PRODUCTION_ENTRY) $(ARGS) $(MAIN_PATH)
	@echo "✅ building finish"

.PHONY: start
start:
	@echo "starting running production build now..."
	source $(ENVRC_FILE) && $(PRODUCTION_ENTRY) --env=prod $(ARGS)

.PHONY: mockery
mockery:
	@echo "⏱️ generating mock now..."
	@rm -rf mocks
	@mockery
	@echo "✅ generating finish"

.PHONY: clean
clean:
	@echo "🗑️ cleaning old build now..."
	@rm -rf $(BUILD_FOLDER)
	@rm -rf $(AIR_TMP_FOLDER)
	@echo "✅ cleaning finish"

.PHONY: check
check: tidy fmt lint test

.PHONY: generate
generate: docker_compose_up regen_openapi regen_proto migration_up regen_gorm  mockery regen_graphql docker_compose_down

.PHONY: fmt
fmt:
	@echo "⏱️ formatting code now..."
	go fmt ./...
	@echo "✅ formatting finish"

.PHONY: goimports
goimports:
	@echo "⏱️ running goimports now..."
	goimports -w .
	@echo "✅ finishing goimports..."

.PHONY: test
test:
	@echo "⏱️ running tests now... "
	go test -race -parallel=4 -timeout 30s -cover $(ARGS) ./internal/...
	@echo "✅ passing all tests."

.PHONY: lint
lint:
	@echo "⏱️ running linting now..."
	golangci-lint run $(ARGS)
	@echo "✅ passing linting..."

.PHONY: tidy
tidy:
	@echo "⏱️ go mod tidy now..."
	go mod tidy
	@echo "✅ finishing tidy..."

.PHONY: regen_openapi
regen_openapi:
	@echo "⏱️ running openapi to go code now..."
	./scripts/regen_openapi.sh
	@echo "✅ finishing openapi to go code..."

.PHONY: regen_proto
regen_proto:
	@echo "⏱️ running regen proto now..."
	./scripts/regen_proto.sh
	@echo "✅ finishing regen proto..."

.PHONY: regen_gorm
regen_gorm:
	@echo "⏱️ running regen gorm now..."
	source $(ENVRC_FILE) && go run gorm/generate.go
	@echo "✅ finishing regen gorm..."

.PHONY: regen_graphql
regen_graphql:
	@echo "⏱️ running regen graphql now..."
	gqlgen generate
	@echo "✅ finishing regen graphql..."

.PHONY: migration_up
migration_up:
	source $(ENVRC_FILE) && goose up

.PHONY: migration_down
migration_down:
	source $(ENVRC_FILE) && goose down

.PHONY: migration_create
migration_create:
	@if [ -z "$(migration_name)" ]; then \
		echo "Error: migration_name is required! Usage: make migration_create migration_name=your_migration"; \
		exit 1; \
	fi
	source $(ENVRC_FILE) && goose create $(migration_name) sql

.PHONY: migration_reset
migration_reset:
	source $(ENVRC_FILE) && goose reset

.PHONY: seed_database
seed_database:
	@echo "⏱️ seeding database now..."
	@make migration_reset
	@make migration_up
	source $(ENVRC_FILE) && go run cmd/seed/database/*.go
	@echo "✅ seeding database finish"

.PHONY: help
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