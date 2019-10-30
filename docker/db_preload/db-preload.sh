#!/bin/sh -cexu

# Build and run migrations
go run github.com/znacol/camping/backend/cmd/db-migrator

# Install psql
apt update -qq
apt-get install -y -qq postgresql-client

# Insert dev data
export PGPASSWORD=password
psql -h camping-db -U root -d camping < /go/src/github.com/znacol/camping/docker/db_preload/dev-data.sql
