#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

# Function to generate protobuf files
generate_proto() {
    local proto_dir=$1
    local output_dir=$2
    local proto_files
    proto_files=$(find $proto_dir -name "*.proto")

    # Clean and create output directory
    rm -rf $output_dir
    mkdir -p $output_dir

    # Generate files
    {
        protoc -I $proto_dir \
            --go_out=$output_dir --go_opt=paths=source_relative \
            --go-grpc_out=$output_dir --go-grpc_opt=paths=source_relative \
            $proto_files
    } 2>&1 | grep -v "warning: Import google/protobuf/.* is unused" || true
}

# GRPC directories
GRPC_PROTO_DIR=$BASE_PATH/proto/grpc
GRPC_GO_OUT_DIR=$BASE_PATH/protogen/grpc

# Set protoc flags
export PROTOC_FLAGS="--allow_unused_imports"

# Generate Go files
generate_proto $GRPC_PROTO_DIR $GRPC_GO_OUT_DIR
