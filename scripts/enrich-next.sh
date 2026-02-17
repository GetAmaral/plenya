#!/usr/bin/env bash
# Enrich next pending ScoreItem (with last_review IS NULL)
# Usage:
#   ./scripts/enrich-next.sh        # Process item at OFFSET 5 (default)
#   ./scripts/enrich-next.sh 10     # Process item at OFFSET 10
#   ./scripts/enrich-next.sh 0 5    # Process 5 items starting at OFFSET 0

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

OFFSET="${1:-5}"
LIMIT="${2:-1}"

echo "🚀 Starting batch enrichment..."
echo "   Offset: $OFFSET"
echo "   Limit: $LIMIT"
echo ""

cd "$PROJECT_ROOT/apps/api"

# Run the Go script
go run scripts/batch_enrich_score_items.go "$OFFSET" "$LIMIT"

echo ""
echo "✅ Done! Check batch_enrichment_report.json for details."
