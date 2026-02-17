#!/bin/bash
# Plenya Score Item Enrichment - Quick Access Script
# Zero-API workflow: Database → LLM → Automatic UPDATE

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print usage
usage() {
    echo -e "${BLUE}Plenya Score Item Enrichment${NC}"
    echo ""
    echo "Usage: ./enrich.sh <command> [options]"
    echo ""
    echo "Commands:"
    echo "  report              Show pending enrichment statistics (read-only)"
    echo "  single <offset>     Enrich single item at offset (0-based)"
    echo "  batch <offset> <n>  Enrich N items starting from offset"
    echo "  dry <offset>        Dry-run for single item (no changes)"
    echo "  count               Quick count of pending vs reviewed items"
    echo ""
    echo "Examples:"
    echo "  ./enrich.sh report                  # Show statistics"
    echo "  ./enrich.sh dry 0                   # Test item at offset 0"
    echo "  ./enrich.sh single 0                # Enrich item at offset 0"
    echo "  ./enrich.sh single 4                # Enrich item at offset 4 (ZERO-API: LIMIT 1 OFFSET 4)"
    echo "  ./enrich.sh batch 0 5               # Enrich 5 items starting from 0"
    echo "  ./enrich.sh count                   # Quick count"
    echo ""
    echo "See ENRICHMENT_GUIDE.md for detailed documentation."
}

# Check Docker
check_docker() {
    if ! docker compose ps api | grep -q "Up"; then
        echo -e "${RED}ERROR: API container is not running${NC}"
        echo "Start it with: docker compose up -d"
        exit 1
    fi
}

# Report command
cmd_report() {
    echo -e "${BLUE}Generating enrichment report...${NC}"
    docker compose exec api go run scripts/report_pending_enrichment.go
}

# Count command
cmd_count() {
    echo -e "${BLUE}Counting items...${NC}"
    docker compose exec api go run scripts/count_review_status.go
}

# Single item enrichment
cmd_single() {
    local offset=$1
    if [ -z "$offset" ]; then
        echo -e "${RED}ERROR: Offset required${NC}"
        echo "Usage: ./enrich.sh single <offset>"
        exit 1
    fi

    echo -e "${BLUE}Enriching item at offset $offset...${NC}"
    echo ""
    docker compose exec api go run scripts/enrich_by_offset.go -offset="$offset"
}

# Dry run
cmd_dry() {
    local offset=$1
    if [ -z "$offset" ]; then
        echo -e "${RED}ERROR: Offset required${NC}"
        echo "Usage: ./enrich.sh dry <offset>"
        exit 1
    fi

    echo -e "${YELLOW}DRY RUN: Item at offset $offset (no changes will be made)${NC}"
    echo ""
    docker compose exec api go run scripts/enrich_by_offset.go -offset="$offset" -dry-run
}

# Batch enrichment
cmd_batch() {
    local offset=$1
    local batch=$2

    if [ -z "$offset" ] || [ -z "$batch" ]; then
        echo -e "${RED}ERROR: Both offset and batch size required${NC}"
        echo "Usage: ./enrich.sh batch <offset> <batch_size>"
        exit 1
    fi

    echo -e "${BLUE}Batch enrichment: $batch items starting from offset $offset${NC}"
    echo -e "${YELLOW}This will include RAG + PubMed integration (15-30s per item)${NC}"
    echo ""
    read -p "Continue? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Cancelled"
        exit 0
    fi

    docker compose exec api go run scripts/batch_enrich_score_items.go "$offset" "$batch"
}

# Main
main() {
    cd "$SCRIPT_DIR"

    if [ $# -eq 0 ]; then
        usage
        exit 0
    fi

    local command=$1
    shift

    check_docker

    case "$command" in
        report)
            cmd_report
            ;;
        count)
            cmd_count
            ;;
        single)
            cmd_single "$@"
            ;;
        dry)
            cmd_dry "$@"
            ;;
        batch)
            cmd_batch "$@"
            ;;
        help|--help|-h)
            usage
            ;;
        *)
            echo -e "${RED}ERROR: Unknown command: $command${NC}"
            echo ""
            usage
            exit 1
            ;;
    esac
}

main "$@"
