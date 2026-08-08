#!/bin/bash
set -euo pipefail

psql "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" \
  -c "CREATE SCHEMA IF NOT EXISTS \"${SCHEMA}\";"

migrate -path /migrations \
  -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable&search_path=${SCHEMA},public&x-migrations-table=${SCHEMA}_schema_migrations" \
  up