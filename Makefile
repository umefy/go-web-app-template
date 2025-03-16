.PHONY: default build run clean check fmt test lint help tidy wire regen_gorm migration_up migration_down migration_create migration_reset

APP_NAME=webapp
MAIN_PATH=cmd/httpserver/main.go
BUILD_FOLDER=./bin
PRODUCTION_ENTRY=$(BUILD_FOLDER)/$(APP_NAME)
AIR_TMP_FOLDER=./tmp
TEST_EXCLUDE_PATHS=protogen|mocks|gorm
TEST_PATHS=$(shell go list ./... | grep -v -E "$(TEST_EXCLUDE_PATHS)")
ENVRC_FILE ?=.envrc

default: dev

print:
	echo "ENV_FILE is $(ENVRC_FILE)"

dev:
	# @trap "exit 0" SIGINT;source .envrc && go run -race $(MAIN_PATH) --env=dev $(ARGS) # ARGS is append for flag parse arguments
	air

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

generate: openapi_to_proto migration_up regen_gorm wire mockery

fmt:
	@echo "⏱️ formatting code now..."
	go fmt ./...
	@echo "✅ formatting finish"

test:
	@echo "⏱️ running tests now... "
	go test -race -parallel=4 -timeout 30s -cover $(ARGS) $(TEST_PATHS)
	@echo "✅ passing all tests."

lint:
	@echo "⏱️ running linting now..."
	golangci-lint run $(ARGS)
	@echo "✅ passing linting..."

tidy:
	@echo "⏱️ go mod tidy now..."
	go mod tidy
	@echo "✅ finishing tidy..."

wire:
	@echo "⏱️ running wire now..."
	wire ./...
	@echo "✅ finishing wire..."

openapi_to_proto:
	@echo "⏱️ running openapi to proto now..."
	./scripts/openapi_to_proto.sh
	@echo "✅ finishing openapi to proto..."

regen_proto:
	@echo "⏱️ running regen proto now..."
	./scripts/regen_proto.sh
	@echo "✅ finishing regen proto..."

regen_gorm:
	@echo "⏱️ running regen gorm now..."
	source $(ENVRC_FILE) && go run gorm/generate.go
	@echo "✅ finishing regen gorm..."

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

help:
	@echo "make - running go code with go run"
	@echo "make dev - running go code with dev environment"
	@echo "make build - build go code to binary"
	@echo "make start - running production build"
	@echo "make clean - clean old binary build"
	@echo "make check - generate required files, formatting, testing and running lint"
	@echo "make test - running go test"
	@echo "make fmt - formatting go code"
	@echo "make lint - running golangci lint"
	@echo "make tidy - install all dependencies"
	@echo "make wire - running wire"
	@echo "make openapi_to_proto - generating proto from openapi"
	@echo "make regen_proto - regenerating source code from proto"
	@echo "make regen_gorm - regenerating gorm models"
	@echo "make migration_up - running database migrations up"
	@echo "make migration_down - running database migrations down"
	@echo "make migration_create - creating a new database migration"
	@echo "make migration_reset - resetting all database migrations"
