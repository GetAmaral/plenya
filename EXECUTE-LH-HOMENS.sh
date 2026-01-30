#!/bin/bash
# ============================================================================
# LH - Homens Enrichment Execution Script
# Score Item ID: 8e402940-afe1-40eb-b4d7-0afbd8f4916e
# Date: 2026-01-29
# ============================================================================

set -e  # Exit on any error

echo "============================================================================"
echo "LH - Homens Enrichment Execution"
echo "============================================================================"
echo ""

# Change to project directory
cd /home/user/plenya

# Step 1: Verify content
echo "Step 1: Verifying content character counts..."
echo "------------------------------------------------------------"
python3 scripts/verify_lh_content.py
echo ""

# Wait for user confirmation
echo "------------------------------------------------------------"
echo "Press ENTER to execute SQL script, or Ctrl+C to abort..."
read -r

# Step 2: Execute SQL
echo ""
echo "Step 2: Executing SQL enrichment script..."
echo "------------------------------------------------------------"
cat scripts/enrich_lh_homens.sql | docker compose exec -T db psql -U plenya_user -d plenya_db

echo ""
echo "============================================================================"
echo "Execution completed!"
echo "============================================================================"
echo ""
echo "Next steps:"
echo "  1. Review the verification output above"
echo "  2. Check that 4 articles were linked"
echo "  3. Verify character counts are within range"
echo "  4. Check the full report: LH-HOMENS-ENRICHMENT-REPORT.md"
echo ""
echo "Files generated:"
echo "  - /home/user/plenya/scripts/enrich_lh_homens.sql"
echo "  - /home/user/plenya/scripts/verify_lh_content.py"
echo "  - /home/user/plenya/LH-HOMENS-ENRICHMENT-REPORT.md"
echo "  - /home/user/plenya/EXECUTE-LH-HOMENS.sh"
echo ""
