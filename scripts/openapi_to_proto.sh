#!/bin/bash

BASE_PATH="$(cd "$(dirname "$0")/.." && pwd)"

INPUT_PATH=$BASE_PATH/openapi/docs/api.yaml
OUTPUT_PATH=$BASE_PATH/openapi/proto
TEMPLATES_PATH=$BASE_PATH/openapi/generator/templates

PROTO_PACKAGE_NAME=v1.models.api

MODELS_PACKAGE=v1/models

rm -rf "${OUTPUT_PATH:?}/${MODELS_PACKAGE:?}"

openapi-generator generate \
  -i "$INPUT_PATH" \
  -g protobuf-schema \
  -o "$OUTPUT_PATH" \
  -t "$TEMPLATES_PATH" \
  --global-property=models \
  --package-name=$PROTO_PACKAGE_NAME \
  --additional-properties \
  goPackagePath=$MODELS_PACKAGE\;api,modelPackage="$MODELS_PACKAGE"