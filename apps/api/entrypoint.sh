#!/bin/sh
set -e

echo "🔄 Running go mod tidy..."
go mod tidy

echo "🔨 Building migrate + server..."
go build -o /tmp/migrate ./cmd/migrate
go build -o /tmp/server ./cmd/server

echo "🗄️  Applying database migrations (goose)..."
# `up` é auto-adotável: em banco vazio roda o baseline; em banco pré-existente
# (schema sem goose) marca o baseline como aplicado e segue. Schema via goose,
# não via AutoMigrate (MIGRATIONS_AUTO fica desligado por default).
/tmp/migrate up

echo "🚀 Starting server..."
exec /tmp/server
