# Build stage
FROM golang:1.26-alpine AS build

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY go.mod go.sum ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o main ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o migrate ./cmd/migrate

# Migrate stage
# Migrate stage
FROM alpine:3.21 AS migrate
WORKDIR /app
RUN apk add --no-cache postgresql-client
# Copy the binary
COPY --from=build /app/migrate .
# FIX: Copy the migrations and datasets from the build stage
COPY --from=build /app/internal/migrations ./migrations
COPY --from=build /app/internal/datasets ./datasets

ENTRYPOINT ["./migrate"]
CMD ["up"]

# Run stage (prod)
FROM alpine:3.21 AS prod
WORKDIR /app
RUN apk add --no-cache wget postgresql-client
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /app/main .
# FIX: Just copy what you need into the final structure
COPY --from=build /app/internal/migrations ./migrations
COPY --from=build /app/internal/datasets ./datasets

RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Read-only root filesystem for security
VOLUME ["/app"]

CMD ["./main"]
