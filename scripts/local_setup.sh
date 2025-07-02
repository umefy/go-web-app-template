#! /bin/bash

set -euo pipefail

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

# Function to handle errors
error_handler() {
  echo "❌ Error occurred on line $1"
}

# Trap errors and call error_handler
trap 'error_handler $LINENO' ERR

# Function to check if a command exists and install it if missing
check_and_install() {
  local cmd="$1"
  local install_cmd="$2"
  local resolved_path
  resolved_path="$(command -v "$cmd" 2>/dev/null || true)"
  # Not found at all
  if [[ -z "$resolved_path" ]]; then
    echo "$cmd not found. Installing... ⏱️"
    eval "$install_cmd"
    echo "$cmd installed ✅"
    asdf reshim golang
    return
  fi

  # If it's an asdf shim, check if the tool is really available
  if [[ "$resolved_path" == "$HOME/.asdf/shims/"* ]]; then
    # Try to resolve real command path via `asdf which`
    if ! asdf which "$cmd" &>/dev/null; then
      echo "$cmd shim found, but no backing tool. Reinstalling... ⏱️"
      eval "$install_cmd"
      echo "$cmd reinstalled ✅"
      asdf reshim golang
      return
    fi
  fi

  echo "$cmd is already installed and usable ✅"
}

# install tools
check_and_install "lefthook" "brew install lefthook"
check_and_install "protoc" "brew install protobuf"
check_and_install "protoc-gen-go" "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
check_and_install "protoc-gen-go-grpc" "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
check_and_install "wire" "go install github.com/google/wire/cmd/wire@latest"
check_and_install "mockery" "go install github.com/vektra/mockery/v3@latest"
check_and_install "goose" "go install github.com/pressly/goose/v3/cmd/goose@latest"
check_and_install "goimports" "go install golang.org/x/tools/cmd/goimports@latest"
check_and_install "gqlgen" "go install github.com/99designs/gqlgen@latest"
check_and_install "openapi-generator" "brew install openapi-generator"
check_and_install "golangci-lint" "brew install golangci-lint"
check_and_install "air" "go install github.com/air-verse/air@latest"

# setup tools
if [ ! -d "google" ]; then
  cp -r "$(brew --prefix protobuf)/include/google" .
fi

# setup envrc
if [ ! -f ".envrc" ]; then
  echo "Creating .envrc from example file ✅"
  cp .envrc.example .envrc
else
  echo ".envrc already exists, skipping ✅"
fi

# Check if current directory is a git repository
if [ ! -d ".git" ]; then
  echo "Initializing git repository ✅"
  git init
else
  echo "Git repository already exists ✅"
fi

# lefthook
lefthook install

cd "$BASE_PATH/scripts/commitlint" && pnpm install && cd "$BASE_PATH" || return

# setup project
go mod tidy -e

make generate

echo "✅ Local setup complete."
echo -e "🚀 start the project by running\033[32m make\033[0m."
