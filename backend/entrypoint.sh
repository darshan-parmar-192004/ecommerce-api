#!/bin/sh
set -e

echo "Waiting for database..."
for i in $(seq 1 30); do
    if pg_isready -h psql_bp -p 5432 -U ${BLUEPRINT_DB_USERNAME:-user}; then
        break
    fi
    echo "Database not ready, waiting... ($i/30)"
    sleep 2
done

echo "Running migrations..."
./migrate up

echo "Starting application..."
exec ./main
