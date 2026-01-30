#!/bin/bash
set -e

echo "=== SETUP VIRTUAL ENVIRONMENT ==="
python3 -m venv /tmp/dietary_venv
source /tmp/dietary_venv/bin/activate
pip install anthropic requests --quiet

echo ""
echo "=== EXECUTING ENRICHMENT SCRIPT ==="
python scripts/enrich_dietary_batch_20.py

echo ""
echo "=== DONE ==="
