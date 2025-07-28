#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

CONFIG_PATH=$BASE_PATH/openapi/openapi_generator_config.yml

GENERATED_PATH=$BASE_PATH/internal/delivery/restful/openapi/v1/generated
rm -rf "$GENERATED_PATH"

openapi-generator generate \
  -c "$CONFIG_PATH"

goimports -w "$GENERATED_PATH"
