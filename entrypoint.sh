#!/bin/bash
set -e

echo "Waiting for ports to be active: "
sh -c './wait-for-it.sh db:5432 -t 30'
sh -c './wait-for-it.sh cache:6379 -t 30'
echo "Ports 5432, 6379 are active"

echo "DB URL: $DB_DSN"
echo "Redis URL: $REDIS_ADDR"

# Run database migrations
echo "Running migrations..."
migrate -path /migrations -database "$DB_DSN" up

# Start the Go application
echo "Starting the application..."
exec /app/myapp
