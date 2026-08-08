ARG GO_VERSION=1.26.5

FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /

ARG SERVICE

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/service ./cmd/services/${SERVICE}

COPY scripts/grpc_health_probe-linux-amd64 /app/grpc_health_probe
RUN chmod +x /app/grpc_health_probe

FROM scratch

COPY --from=build /app/service /service
COPY --from=build /app/grpc_health_probe /grpc_health_probe
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/service"]