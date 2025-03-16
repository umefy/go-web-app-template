#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

# Define output directories for different languages
OPENAPI_GO_OUT_DIR=$BASE_PATH/openapi/protogen
rm -rf $OPENAPI_GO_OUT_DIR

# Create output directories if they don't exist
mkdir -p $OPENAPI_GO_OUT_DIR

# Find all .proto files
PROTO_FILES=$(find $BASE_PATH/openapi/proto -name "*.proto")

export PROTOC_FLAGS="--allow_unused_imports"
# Generate Go files
protoc -I $BASE_PATH/openapi/proto \
--go_out=$OPENAPI_GO_OUT_DIR --go_opt=paths=source_relative \
--go-grpc_out=$OPENAPI_GO_OUT_DIR --go-grpc_opt=paths=source_relative \
$PROTO_FILES