#!/bin/bash

# Method Content Enrichment Script
# Enrich method letters and pillars with clinical content using Claude AI

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
TIER="letter"
OFFSET=0
DRY_RUN=false
STATUS_ONLY=false

# Parse arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    --tier)
      TIER="$2"
      shift 2
      ;;
    --offset)
      OFFSET="$2"
      shift 2
      ;;
    --dry-run)
      DRY_RUN=true
      shift
      ;;
    --status)
      STATUS_ONLY=true
      shift
      ;;
    --help)
      echo "Usage: ./scripts/enrich.sh [OPTIONS]"
      echo ""
      echo "Options:"
      echo "  --tier <letter|pillar>    Tier to enrich (default: letter)"
      echo "  --offset <N>              Pagination offset (default: 0)"
      echo "  --dry-run                 Preview without updating database"
      echo "  --status                  Show enrichment status only"
      echo "  --help                    Show this help message"
      echo ""
      echo "Examples:"
      echo "  ./scripts/enrich.sh --status"
      echo "  ./scripts/enrich.sh --tier letter --offset 0"
      echo "  ./scripts/enrich.sh --tier pillar --offset 5 --dry-run"
      exit 0
      ;;
    *)
      echo -e "${RED}Unknown option: $1${NC}"
      exit 1
      ;;
  esac
done

# Check if ANTHROPIC_API_KEY is set
if [ -z "$ANTHROPIC_API_KEY" ] && [ "$STATUS_ONLY" = false ]; then
  echo -e "${RED}Error: ANTHROPIC_API_KEY environment variable not set${NC}"
  echo "Set it with: export ANTHROPIC_API_KEY=your-api-key"
  exit 1
fi

cd "$(dirname "$0")/.."

if [ "$STATUS_ONLY" = true ]; then
  echo -e "${BLUE}Checking enrichment status...${NC}"
  cd apps/api
  go run scripts/check_enrichment_status.go --tier all
else
  echo -e "${YELLOW}=== Method Content Enrichment ===${NC}"
  echo "Tier: $TIER"
  echo "Offset: $OFFSET"

  if [ "$DRY_RUN" = true ]; then
    echo -e "${YELLOW}Mode: DRY RUN (preview only)${NC}"
  else
    echo -e "${GREEN}Mode: LIVE (will update database)${NC}"
  fi

  echo ""

  cd apps/api

  if [ "$DRY_RUN" = true ]; then
    go run scripts/enrich_method_content.go --tier "$TIER" --offset "$OFFSET" --dry-run
  else
    go run scripts/enrich_method_content.go --tier "$TIER" --offset "$OFFSET"
  fi

  echo ""
  echo -e "${GREEN}Done!${NC}"
fi
