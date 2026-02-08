#!/bin/bash

# Script helper para executar o update de gender em score_items

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
API_DIR="$(cd "$SCRIPT_DIR/../.." && pwd)"

echo "==================================="
echo "Score Item Gender Update Helper"
echo "==================================="
echo ""

# Verificar se está dentro do Docker ou host
if [ -f "/.dockerenv" ]; then
    echo "🐳 Running inside Docker container"
    INSIDE_DOCKER=true
else
    echo "💻 Running on host machine"
    INSIDE_DOCKER=false
fi

# Verificar se .env existe
if [ -f "$API_DIR/.env" ]; then
    echo "✓ Found .env file"
    if [ "$INSIDE_DOCKER" = false ]; then
        echo "📋 Loading environment variables from .env"
        export $(cat "$API_DIR/.env" | grep -v '^#' | xargs)
    fi
else
    echo "⚠️  No .env file found. Using system environment variables."
fi

# Verificar conectividade com PostgreSQL
echo ""
echo "🔍 Checking database connectivity..."

if [ "$INSIDE_DOCKER" = true ]; then
    # Dentro do Docker, o host é 'db'
    DB_HOST=${DB_HOST:-db}
else
    DB_HOST=${DB_HOST:-localhost}
fi

DB_PORT=${DB_PORT:-5432}

# Simples check de conectividade (requer nc/netcat)
if command -v nc &> /dev/null; then
    if nc -z "$DB_HOST" "$DB_PORT" 2>/dev/null; then
        echo "✓ Database is reachable at $DB_HOST:$DB_PORT"
    else
        echo "❌ Cannot reach database at $DB_HOST:$DB_PORT"
        echo "Please ensure PostgreSQL is running."
        exit 1
    fi
else
    echo "⚠️  'nc' not found, skipping connectivity check"
fi

# Executar script
echo ""
echo "🚀 Starting gender update script..."
echo "-----------------------------------"
echo ""

cd "$API_DIR"

if go run cmd/update-score-item-gender/main.go; then
    echo ""
    echo "-----------------------------------"
    echo "✅ Gender update completed successfully!"
    exit 0
else
    echo ""
    echo "-----------------------------------"
    echo "❌ Gender update failed!"
    exit 1
fi
