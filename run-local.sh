#!/bin/sh

export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_NAME=jams
export DATABASE_USER=jams
export DATABASE_PASSWORD=jams
export DATABASE_SSL_MODE=disable
export DATABASE_MIGRATION_PATH=file://./resources/db/migrations

go run ./server/cmd/server/main.go