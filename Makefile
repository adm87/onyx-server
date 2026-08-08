SERVICES := auth user
PROTO_DIR := proto/services
GEN_DIR := proto/gen
GO_VERSION := 1.26.5

define SERVICE_template
.PHONY: up-$(1)-service
up-$(1)-service:
	@echo "Starting $(1) service..."
	docker compose up --build $(1)-service

.PHONY: down-$(1)-service
down-$(1)-service:
	@echo "Stopping $(1) service..."
	docker compose stop $(1)-service
endef

$(foreach svc,$(SERVICES),$(eval $(call SERVICE_template,$(svc))))

.PHONY: up-postgres
up-postgres:
	@echo "Starting Postgres..."
	docker compose up postgres

.PHONY: down-postgres
down-postgres:
	@echo "Stopping Postgres..."
	docker compose stop postgres

.PHONY: up-gateway
up-gateway:
	@echo "Starting Onyx gateway..."
	docker compose up --build gateway

.PHONY: down-gateway
down-gateway:
	@echo "Stopping Onyx gateway..."
	docker compose stop gateway

.PHONY: up down
up:
	docker compose up --build

down:
	docker compose down

.PHONY: proto
proto:
	protoc \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		--proto_path=$(PROTO_DIR) \
		$(PROTO_DIR)/auth/v1/auth.proto \
		$(PROTO_DIR)/user/v1/user.proto

.PHONY: openapi
openapi:
	go generate ./openapi