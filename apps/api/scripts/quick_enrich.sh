#!/bin/bash
# Quick enrichment of a single ScoreItem via SQL (ZERO-API approach)
# Usage: ./quick_enrich.sh [OFFSET]

OFFSET=${1:-0}

echo "🔍 Fetching next pending ScoreItem (OFFSET: $OFFSET)..."
echo ""

# Fetch and display the item
ITEM_INFO=$(docker compose exec db psql -U plenya_user -d plenya_db -t << EOF
SELECT
  id::text || '|' ||
  name || '|' ||
  COALESCE(unit, 'N/A') || '|' ||
  COALESCE(gender, 'not_applicable') || '|' ||
  COALESCE(clinical_relevance, '[empty]') || '|' ||
  COALESCE(patient_explanation, '[empty]') || '|' ||
  COALESCE(conduct, '[empty]')
FROM score_items
WHERE last_review IS NULL
  AND deleted_at IS NULL
ORDER BY "order" ASC
LIMIT 1 OFFSET $OFFSET;
EOF
)

if [ -z "$ITEM_INFO" ]; then
  echo "✅ No more pending items!"
  exit 0
fi

IFS='|' read -r ITEM_ID ITEM_NAME ITEM_UNIT ITEM_GENDER ITEM_CLINICAL ITEM_PATIENT ITEM_CONDUCT <<< "$ITEM_INFO"

echo "=================================================="
echo "📊 ScoreItem: $ITEM_NAME"
echo "   ID: $ITEM_ID"
echo "   Unit: $ITEM_UNIT"
echo "   Gender: $ITEM_GENDER"
echo ""
echo "📝 Current State:"
echo "   Clinical Relevance: ${ITEM_CLINICAL:0:60}..."
echo "   Patient Explanation: ${ITEM_PATIENT:0:60}..."
echo "   Conduct: ${ITEM_CONDUCT:0:60}..."
echo "=================================================="
echo ""

# Interactive enrichment
read -p "🩺 Clinical Relevance (or 'skip'): " CLINICAL
read -p "👤 Patient Explanation (or 'skip'): " PATIENT
read -p "💊 Conduct (or 'skip'): " CONDUCT

# Build UPDATE statement
UPDATE_FIELDS=""
if [ "$CLINICAL" != "skip" ] && [ -n "$CLINICAL" ]; then
  UPDATE_FIELDS="${UPDATE_FIELDS}clinical_relevance = '$CLINICAL', "
fi

if [ "$PATIENT" != "skip" ] && [ -n "$PATIENT" ]; then
  UPDATE_FIELDS="${UPDATE_FIELDS}patient_explanation = '$PATIENT', "
fi

if [ "$CONDUCT" != "skip" ] && [ -n "$CONDUCT" ]; then
  UPDATE_FIELDS="${UPDATE_FIELDS}conduct = '$CONDUCT', "
fi

# Check if any fields were updated
if [ -z "$UPDATE_FIELDS" ]; then
  echo ""
  echo "⏭️  No changes - skipped"
  exit 0
fi

# Add last_review timestamp
UPDATE_FIELDS="${UPDATE_FIELDS}last_review = NOW(), updated_at = NOW()"

# Confirm
read -p "💾 Save changes? (y/n): " CONFIRM

if [ "$CONFIRM" != "y" ]; then
  echo "❌ Cancelled"
  exit 0
fi

# Execute UPDATE
docker compose exec db psql -U plenya_user -d plenya_db << EOF
UPDATE score_items
SET $UPDATE_FIELDS
WHERE id = '$ITEM_ID';
EOF

echo ""
echo "✅ Updated successfully!"
echo ""
echo "📋 Report:"
echo "   Name: $ITEM_NAME"
echo "   Tier: ScoreItem"
echo "   Status: enriched"
echo "   ID: $ITEM_ID"
