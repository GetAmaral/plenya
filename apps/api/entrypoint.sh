#!/bin/sh
set -e

echo "🔄 Running go mod tidy..."
go mod tidy

echo "🔨 Building server (compile once, restart fast)..."
go build -o /tmp/server ./cmd/server

echo "🚀 Starting server..."
exec /tmp/server
