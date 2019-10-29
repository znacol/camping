#!/bin/sh -cexu

# Build and run migrations
go run github.com/znacol/camping/backend/cmd/db-migrator

# Install psql
# apt-get update -qq
# apt-get install -y -qq postgresql-client

# Insert dev data
# export PGPASSWORD=password
# psql -h seeds-db -U root -d camping < /go/src/github.com/znacol/camping/docker/database_preload/dev-data.sql
