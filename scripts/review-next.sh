#!/bin/bash
# Review Next ScoreItem
# Usage: ./scripts/review-next.sh [offset] [--skip-llm]
#
# Examples:
#   ./scripts/review-next.sh 7              # Auto-review offset 7
#   ./scripts/review-next.sh 7 --skip-llm   # Skip LLM, only process Tier 1
#   ./scripts/review-next.sh                # Start from offset 0

set -e

OFFSET="${1:-0}"
SKIP_LLM="${2:-}"

echo "🔄 Reviewing ScoreItem at offset $OFFSET..."
echo

if [ -n "$SKIP_LLM" ]; then
    docker compose exec -T api sh -c "cd /app && go run cmd/auto-review/main.go $OFFSET $SKIP_LLM"
else
    docker compose exec -T api sh -c "cd /app && go run cmd/auto-review/main.go $OFFSET"
fi
