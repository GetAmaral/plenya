#!/bin/sh
set -e

echo "🔍 Checking if dependencies need update..."

# Verificar se package.json ou pnpm-lock.yaml mudaram
if [ ! -f "/tmp/.deps-installed" ] || \
   [ "package.json" -nt "/tmp/.deps-installed" ] || \
   [ "pnpm-lock.yaml" -nt "/tmp/.deps-installed" ]; then
    echo "📦 Installing dependencies..."
    pnpm install --no-frozen-lockfile
    touch /tmp/.deps-installed
else
    echo "✅ Dependencies up to date"
fi

echo "🚀 Starting Next.js dev server..."
exec pnpm dev --filter "@plenya/web"
