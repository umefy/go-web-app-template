#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

CONFIG_PATH=$BASE_PATH/openapi/openapi_generator_config.yml

rm -rf "$BASE_PATH/openapi/proto"

openapi-generator generate \
  -c "$CONFIG_PATH"
