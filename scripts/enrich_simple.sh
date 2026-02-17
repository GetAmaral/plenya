#!/bin/bash
# Simple enrichment script - runs inside Docker container
# Usage: docker compose exec api ./scripts/enrich_simple.sh [offset] [batch] [dry-run]

OFFSET=${1:-0}
BATCH=${2:-10}
DRY_RUN_FLAG=""

if [ "${3}" = "true" ]; then
    DRY_RUN_FLAG="-dry-run"
fi

# Check if we're inside container or need to exec into it
if [ -f "/.dockerenv" ]; then
    # Inside container
    echo "Running inside container..."
    go run /app/cmd/enrich_score_items.go -offset="$OFFSET" -batch="$BATCH" $DRY_RUN_FLAG
else
    # Outside container - need to exec
    echo "Running from host, executing inside container..."
    docker compose exec api /bin/bash -c "go run /app/cmd/enrich_score_items.go -offset=$OFFSET -batch=$BATCH $DRY_RUN_FLAG"
fi
