#!/bin/bash

export MIGRATION_DSN="host=$PG_DB_HOST port=$PG_PORT_IN dbname=$PG_DB_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"

sleep 2 && goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v