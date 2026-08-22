#!/usr/bin/env bash
set -euo pipefail

OUT_DIR="./pkg/proto/gen"
mkdir -p "$OUT_DIR"

PROTO_FILES=$(find ./pkg/proto -name "*.proto")

protoc \
  --proto_path=./pkg/proto \
  --go_out="$OUT_DIR" --go_opt=paths=source_relative \
  --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative \
  $PROTO_FILES