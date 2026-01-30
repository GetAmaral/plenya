#!/bin/bash

# Execute Surgical History Score Item Enrichment
# Item ID: 4511ae8d-2ad3-40e0-a47d-2cef065d39e9

echo "=========================================="
echo "Surgical History Score Item Enrichment"
echo "=========================================="
echo ""
echo "Group: Histórico de doenças (Disease History)"
echo "Subgroup: Cirurgias já realizadas (Previous Surgeries)"
echo "Item: Registrar quaisquer cirurgias realizadas"
echo ""
echo "Starting enrichment process..."
echo ""

# Execute the SQL file
docker compose exec -T db psql -U plenya_user -d plenya_db < /home/user/plenya/enrich_surgical_history_item.sql

if [ $? -eq 0 ]; then
    echo ""
    echo "=========================================="
    echo "✓ Enrichment completed successfully!"
    echo "=========================================="
    echo ""
    echo "Summary:"
    echo "- 4 scientific articles inserted/linked"
    echo "- Article-score item relationships created"
    echo "- Clinical content updated (PT-BR)"
    echo "- last_review timestamp set"
    echo ""
else
    echo ""
    echo "=========================================="
    echo "✗ Error during enrichment"
    echo "=========================================="
    echo ""
    exit 1
fi
