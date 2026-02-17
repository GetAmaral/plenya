#!/bin/bash
# Score Item Enrichment - Quick Commands
# ZERO-API workflow for enriching clinical fields

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
print_usage() {
    echo -e "${BLUE}Score Item Enrichment - Quick Commands${NC}"
    echo ""
    echo "Usage: ./scripts/enrich_score.sh [command] [options]"
    echo ""
    echo "Commands:"
    echo "  report              Generate enrichment status report"
    echo "  review <offset>     Review item without enriching (dry-run)"
    echo "  enrich <offset>     Enrich single item"
    echo "  batch <start> <end> Enrich range of items (e.g., 0 10)"
    echo ""
    echo "Examples:"
    echo "  ./scripts/enrich_score.sh report"
    echo "  ./scripts/enrich_score.sh review 0"
    echo "  ./scripts/enrich_score.sh enrich 0"
    echo "  ./scripts/enrich_score.sh batch 0 9"
}

run_report() {
    echo -e "${GREEN}📊 Generating enrichment report...${NC}"
    docker compose exec -T api go run /app/scripts/report_pending_enrichment.go
}

run_review() {
    local offset=$1
    if [ -z "$offset" ]; then
        echo -e "${RED}Error: Offset required${NC}"
        echo "Usage: ./scripts/enrich_score.sh review <offset>"
        exit 1
    fi

    echo -e "${YELLOW}👁️  Reviewing item at offset $offset (dry-run)...${NC}"
    docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset "$offset" -dry-run
}

run_enrich() {
    local offset=$1
    if [ -z "$offset" ]; then
        echo -e "${RED}Error: Offset required${NC}"
        echo "Usage: ./scripts/enrich_score.sh enrich <offset>"
        exit 1
    fi

    echo -e "${GREEN}🧠 Enriching item at offset $offset...${NC}"
    docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset "$offset"
}

run_batch() {
    local start=$1
    local end=$2

    if [ -z "$start" ] || [ -z "$end" ]; then
        echo -e "${RED}Error: Start and end offsets required${NC}"
        echo "Usage: ./scripts/enrich_score.sh batch <start> <end>"
        exit 1
    fi

    if [ "$start" -gt "$end" ]; then
        echo -e "${RED}Error: Start must be less than or equal to end${NC}"
        exit 1
    fi

    local count=$((end - start + 1))
    echo -e "${GREEN}🚀 Batch enrichment: Processing $count items (offset $start to $end)${NC}"
    echo -e "${YELLOW}⏱️  Estimated time: $((count * 10)) seconds (with delays)${NC}"
    echo ""

    # Confirm
    read -p "Continue? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Aborted."
        exit 0
    fi

    for ((i=start; i<=end; i++)); do
        echo -e "\n${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
        echo -e "${GREEN}Processing item $((i-start+1))/$count (offset $i)${NC}"
        echo -e "${BLUE}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}\n"

        docker compose exec -T api go run /app/scripts/enrich_by_offset.go -offset "$i"

        # Rate limiting (skip on last item)
        if [ "$i" -lt "$end" ]; then
            echo -e "\n${YELLOW}⏳ Waiting 5 seconds before next item...${NC}"
            sleep 5
        fi
    done

    echo -e "\n${GREEN}✅ Batch complete!${NC}"
    echo -e "${BLUE}Running final report...${NC}\n"
    run_report
}

# Main
case "${1:-}" in
    report)
        run_report
        ;;
    review)
        run_review "$2"
        ;;
    enrich)
        run_enrich "$2"
        ;;
    batch)
        run_batch "$2" "$3"
        ;;
    *)
        print_usage
        exit 1
        ;;
esac
