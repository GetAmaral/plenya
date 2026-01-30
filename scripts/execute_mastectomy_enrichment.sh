#!/bin/bash

# Mastectomy Score Item Enrichment Execution Script
# Score Item ID: e65c56dc-5c07-4270-8a8b-017b293ca147
# Date: 2026-01-28

set -e

echo "=========================================="
echo "Mastectomy Score Item Enrichment"
echo "=========================================="
echo ""
echo "Score Item ID: e65c56dc-5c07-4270-8a8b-017b293ca147"
echo "Group: Histórico de doenças"
echo "Subgroup: Cirurgias já realizadas"
echo "Item: Mastectomia"
echo ""
echo "This script will:"
echo "  1. Insert 4 scientific articles"
echo "  2. Create article-score_item relationships"
echo "  3. Update clinical content (relevance, explanation, conduct)"
echo "  4. Set last_review timestamp"
echo ""
echo "Press ENTER to continue or Ctrl+C to cancel..."
read

echo "Executing enrichment SQL..."
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/scripts/enrich_mastectomy_item.sql

echo ""
echo "=========================================="
echo "Enrichment completed successfully!"
echo "=========================================="
echo ""
echo "Verification query results shown above."
echo "Expected: 4 articles linked, all content fields populated"
echo ""
echo "To verify in database:"
echo "  docker compose exec -T db psql -U plenya_user -d plenya_db -c \\"
echo "    \"SELECT id, item_text, LENGTH(clinical_relevance) as relevance_len, \\"
echo "     LENGTH(patient_explanation) as explanation_len, LENGTH(conduct) as conduct_len, \\"
echo "     last_review FROM score_items WHERE id='e65c56dc-5c07-4270-8a8b-017b293ca147';\""
echo ""
echo "See MASTECTOMY-ENRICHMENT-REPORT.md for full documentation."
