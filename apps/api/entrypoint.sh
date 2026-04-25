#!/bin/sh
set -e

echo "🔄 Running go mod tidy..."
go mod tidy

echo "🚀 Starting server..."
exec go run ./cmd/server
