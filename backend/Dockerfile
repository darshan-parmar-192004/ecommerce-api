# Build stage
FROM golang:1.26-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o main ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o migrate ./cmd/migrate

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o seed ./cmd/seed

# Migrate stage
FROM alpine:3.20.1 AS migrate
WORKDIR /app
RUN apk add --no-cache postgresql-client

COPY --from=build /app/migrate .
COPY --from=build /app/internal/migrations ./internal/migrations

CMD ["./migrate", "up"]

# Seed stage
FROM alpine:3.20.1 AS seed
WORKDIR /app
RUN apk add --no-cache postgresql-client

COPY --from=build /app/seed .
COPY --from=build /app/internal/datasets ./internal/datasets

CMD ["./seed"]

# Run stage
FROM alpine:3.20.1 AS prod

WORKDIR /app

RUN apk add --no-cache wget postgresql-client

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /app/main .
COPY --from=build /app/migrate .
COPY --from=build /app/seed .
COPY --from=build /app/internal/migrations ./internal/migrations
COPY --from=build /app/internal/datasets ./internal/datasets
COPY --from=build /app/entrypoint.sh .

RUN chmod +x entrypoint.sh
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./entrypoint.sh"]
