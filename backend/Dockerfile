# Build stage
FROM golang:1.26-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from the backend context
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

# FIX: Copy to ./migrations so the binary can find it at "file://migrations"
COPY --from=build /app/internal/migrations ./migrations
COPY --from=build /app/internal/data* ./internal/datasets/

CMD ["./migrate", "up"]

# Prod stage
FROM alpine:3.20.1 AS prod
WORKDIR /app
RUN apk add --no-cache wget postgresql-client
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /app/main .
COPY --from=build /app/migrate .

# FIX: Do the same here for the prod container
COPY --from=build /app/internal/migrations ./migrations
COPY --from=build /app/internal/data* ./internal/datasets/

RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080
CMD ["./main"]