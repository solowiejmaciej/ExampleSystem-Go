#!/bin/bash
set -e

POSTGRES="psql --username ${POSTGRES_USER} --dbname postgres --tuples-only --no-align"

databases=("users-api" "notificator" "root")
echo "Creating databases..."
for dbname in "${databases[@]}"; do
    echo "Checking if database $dbname exists..."

    dbExists=$($POSTGRES -c "SELECT 1 FROM pg_database WHERE datname = '${dbname}';")

    if [ "$dbExists" = "1" ]; then
        echo "Database $dbname exists."
    else
        echo "Database $dbname does not exist. Creating..."
        $POSTGRES -c "CREATE DATABASE ${dbname} OWNER ${POSTGRES_USER};"
        echo "Database $dbname created."
    fi
done