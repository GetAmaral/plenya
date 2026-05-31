#!/bin/sh
set -e

# CI=true evita ERR_PNPM_ABORTED_REMOVE_MODULES_DIR_NO_TTY quando pnpm
# precisa remover/recriar node_modules sem TTY (docker compose detached).
export CI=true

echo "🔍 Checking if dependencies need update..."

# Marcador de data não basta: os node_modules são volumes anônimos que podem
# ficar parciais (ex.: o binário do next some) sem que package.json/lockfile mudem.
# Por isso validamos também a presença real do binário antes de pular o install.
NEXT_BIN="apps/site/node_modules/next/dist/bin/next"

if [ ! -f "/tmp/.deps-installed" ] || \
   [ ! -f "$NEXT_BIN" ] || \
   [ "package.json" -nt "/tmp/.deps-installed" ] || \
   [ "pnpm-lock.yaml" -nt "/tmp/.deps-installed" ]; then
    echo "📦 Installing dependencies..."
    pnpm install --no-frozen-lockfile
    touch /tmp/.deps-installed
else
    echo "✅ Dependencies up to date"
fi

echo "🚀 Starting Plenya Site dev server..."
exec pnpm dev --filter "@plenya/site"
