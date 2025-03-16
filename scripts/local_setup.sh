#! /bin/bash

set -e

# Function to handle errors
error_handler() {
    echo "❌ Error occurred on line $1"
}

# Trap errors and call error_handler
trap 'error_handler $LINENO' ERR

# Function to check if a command exists and install it if missing
check_and_install() {
    local cmd="$1"         # First argument: the command to check
    local install_cmd="$2" # Second argument: the command to install it

    if ! command -v "$cmd" &>/dev/null; then
        echo "$cmd is not installed. Installing now...⏱️"
        eval "$install_cmd"
        echo "$cmd finished installing ✅"
    else
        echo "$cmd is already installed. ✅"
    fi
}

# install tools
check_and_install "lefthook" "brew install lefthook"
check_and_install "protoc" "brew install protobuf"
check_and_install "protoc-gen-go" "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
check_and_install "protoc-gen-go-grpc" "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
check_and_install "wire" "go install github.com/google/wire/cmd/wire@latest"
check_and_install "mockery" "go install github.com/vektra/mockery/v2@v2.46.0"
check_and_install "goose" "go install github.com/pressly/goose/v3/cmd/goose@latest"
check_and_install "openapi-generator-cli" "brew install openapi-generator"

# setup tools
if [ ! -d "google" ]; then
    ln -s "$(brew --prefix protobuf)/include/google" .
fi

# setup envrc
if [ ! -f ".envrc" ]; then
    echo "Creating .envrc from example file ✅"
    cp .envrc.example .envrc
else
    echo ".envrc already exists, skipping ✅"
fi

# lefthook
lefthook install

# setup project
go mod tidy -e

make check