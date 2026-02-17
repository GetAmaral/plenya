#!/bin/bash
set -e

# Score Item Enrichment Script
# Zero-API approach: Direct DB access + LLM enrichment

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Default values
OFFSET=${1:-0}
BATCH=${2:-10}
DRY_RUN=${3:-false}

echo "=== Plenya Score Item Enrichment ==="
echo "Project root: $PROJECT_ROOT"
echo "Offset: $OFFSET"
echo "Batch size: $BATCH"
echo "Dry run: $DRY_RUN"
echo ""

# Check if ANTHROPIC_API_KEY is set
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "ERROR: ANTHROPIC_API_KEY environment variable not set"
    echo "Please export ANTHROPIC_API_KEY=your_key_here"
    exit 1
fi

# Load .env file if exists
if [ -f "$PROJECT_ROOT/.env" ]; then
    echo "Loading environment from .env..."
    set -a
    source "$PROJECT_ROOT/.env"
    set +a
fi

# Run enrichment inside Docker container (Go is available there)
echo "Installing Go dependencies in container..."
docker compose exec -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" api go get github.com/liushuangls/go-anthropic/v2

# Run enrichment
echo ""
echo "Starting enrichment..."
echo ""

if [ "$DRY_RUN" = "true" ]; then
    docker compose exec -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" api \
        go run /app/cmd/enrich_score_items.go -offset="$OFFSET" -batch="$BATCH" -dry-run
else
    docker compose exec -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" api \
        go run /app/cmd/enrich_score_items.go -offset="$OFFSET" -batch="$BATCH"
fi

echo ""
echo "=== Enrichment Complete ==="
