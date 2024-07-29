#!/bin/bash

set -e

# Load environment variables from .env file
if [ -f .env ]; then
    export $(cat .env | sed 's/#.*//g' | xargs)
else
    echo ".env file not found"
    exit 1
fi

# Drop and recreate databases
echo "Dropping and recreating databases..."

# Drop and recreate user database
docker-compose exec -T user_db psql -U user_user -d postgres -c "DROP DATABASE IF EXISTS user_db;"
docker-compose exec -T user_db psql -U user_user -d postgres -c "CREATE DATABASE user_db;"

# Drop and recreate auth database
docker-compose exec -T auth_db psql -U auth_user -d postgres -c "DROP DATABASE IF EXISTS auth_db;"
docker-compose exec -T auth_db psql -U auth_user -d postgres -c "CREATE DATABASE auth_db;"

# Run user migrations
echo "Running user migrations..."
docker-compose exec -T app go run cmd/migrate/user/main.go

# Insert test data into user database
echo "Inserting test data into user database..."
docker-compose exec -T app go run cmd/migrate/user/insert_test_data.go

# Run auth migrations
echo "Running auth migrations..."
docker-compose exec -T app go run cmd/migrate/auth/main.go

# Insert test data into auth database
echo "Inserting test data into auth database..."
docker-compose exec -T app go run cmd/migrate/auth/insert_test_data.go

echo "All migrations and test data insertions completed successfully!"
