# Build stage
FROM golang:1.25.6-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o main ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o migrate ./cmd/migrate

# Migrate stage
FROM alpine:3.20.1 AS migrate

WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=build /app/migrate .
COPY --from=build /app/internal/migrations ./migrations
COPY --from=build /app/internal/datasets/ ./datasets

CMD ["./migrate", "up"]

# Run stage
FROM alpine:3.20.1 AS prod
WORKDIR /app
RUN apk add --no-cache wget postgresql-client
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /app/main .
COPY --from=build /app/migrate .
COPY --from=build /app/internal/migrations ./migrations

# FIX: Added wildcard to prevent "not found" crash
COPY --from=build /app/internal/data* ./datasets/

RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]
