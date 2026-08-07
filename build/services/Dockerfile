ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /

ARG SERVICE

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/service ./cmd/services/${SERVICE}

FROM scratch

COPY --from=build /app/service /service
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/service"]