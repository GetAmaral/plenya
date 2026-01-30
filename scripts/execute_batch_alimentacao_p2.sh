#!/bin/bash
# Execute Batch ALIMENTAÇÃO Parte 2 - 20 items enrichment

set -e

echo "============================================="
echo "BATCH ALIMENTAÇÃO PARTE 2 - 20 ITEMS"
echo "============================================="
echo ""

# Check if API key is set
if [ -z "$ANTHROPIC_API_KEY" ] && [ ! -f ~/.anthropic_key ]; then
    echo "ERROR: ANTHROPIC_API_KEY not configured"
    echo ""
    echo "Please set your API key using one of these methods:"
    echo ""
    echo "1. Environment variable:"
    echo "   export ANTHROPIC_API_KEY='sk-ant-...'"
    echo ""
    echo "2. Config file:"
    echo "   echo 'sk-ant-...' > ~/.anthropic_key"
    echo ""
    exit 1
fi

# Check if Python dependencies are installed
echo "Checking dependencies..."
python3 -c "import anthropic" 2>/dev/null || {
    echo "ERROR: anthropic package not installed"
    echo "Install it with: pip install anthropic"
    exit 1
}

echo "✓ Dependencies OK"
echo ""

# Execute enrichment script
echo "Starting enrichment process..."
echo "This will take approximately 10-12 minutes"
echo ""

python3 scripts/batch_alimentacao_parte2.py

# Check if execution was successful
if [ $? -eq 0 ]; then
    echo ""
    echo "============================================="
    echo "EXECUTION COMPLETED SUCCESSFULLY"
    echo "============================================="
    echo ""
    echo "Generated files:"
    echo "  - batch_alimentacao_parte2_results.json"
    echo "  - batch_alimentacao_parte2.sql"
    echo ""
    echo "Next steps:"
    echo "  1. Review the generated JSON and SQL files"
    echo "  2. Apply SQL to database:"
    echo "     docker compose exec db psql -U plenya_user -d plenya_db -f /path/to/batch_alimentacao_parte2.sql"
    echo "  3. Validate in database"
    echo ""
else
    echo ""
    echo "ERROR: Execution failed"
    echo "Check the error messages above"
    exit 1
fi
