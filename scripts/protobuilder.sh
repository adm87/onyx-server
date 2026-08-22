#!/usr/bin/env bash
set -euo pipefail

OUT_DIR="./pkg/proto/gen"
mkdir -p "$OUT_DIR"

mapfile -t PROTO_FILES < <(find ./pkg/proto -name "*.proto" -not -path "./pkg/proto/gen/*")

protoc \
  --proto_path=./pkg/proto \
  --proto_path=/opt/googleapis \
  --go_out="$OUT_DIR" --go_opt=paths=source_relative \
  --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out="$OUT_DIR" --grpc-gateway_opt=paths=source_relative \
  "${PROTO_FILES[@]}"

echo "Proto generation complete."