#!/bin/bash
# Quick SQL-based enrichment report generator

echo "📊 Score Items & Levels Enrichment Status Report"
echo "=================================================="
echo ""

docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
-- Score Items Summary
SELECT
  'ScoreItems' AS tier,
  COUNT(*) AS total,
  COUNT(last_review) AS enriched,
  COUNT(*) - COUNT(last_review) AS pending
FROM score_items
WHERE deleted_at IS NULL

UNION ALL

-- Score Levels Summary
SELECT
  'ScoreLevels' AS tier,
  COUNT(*) AS total,
  COUNT(last_review) AS enriched,
  COUNT(*) - COUNT(last_review) AS pending
FROM score_levels
WHERE deleted_at IS NULL

UNION ALL

-- Total Summary
SELECT
  'TOTAL' AS tier,
  si_total + sl_total AS total,
  si_enriched + sl_enriched AS enriched,
  si_pending + sl_pending AS pending
FROM (
  SELECT
    COUNT(*) AS si_total,
    COUNT(last_review) AS si_enriched,
    COUNT(*) - COUNT(last_review) AS si_pending
  FROM score_items WHERE deleted_at IS NULL
) si,
(
  SELECT
    COUNT(*) AS sl_total,
    COUNT(last_review) AS sl_enriched,
    COUNT(*) - COUNT(last_review) AS sl_pending
  FROM score_levels WHERE deleted_at IS NULL
) sl;
EOF

echo ""
echo "---------------------------------------------------"
echo ""
echo "📋 Pending ScoreItems (last_review IS NULL):"
echo ""

docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
SELECT
  si.name,
  COALESCE(si.gender, 'not_applicable') AS gender,
  sg.name AS subgroup,
  CASE
    WHEN si.clinical_relevance IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_clinical,
  CASE
    WHEN si.patient_explanation IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_patient,
  CASE
    WHEN si.conduct IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_conduct
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
WHERE si.last_review IS NULL
  AND si.deleted_at IS NULL
ORDER BY si."order" ASC
LIMIT 10;
EOF

echo ""
echo "---------------------------------------------------"
echo ""
echo "📋 Pending ScoreLevels (last_review IS NULL):"
echo ""

docker compose exec db psql -U plenya_user -d plenya_db << 'EOF'
SELECT
  sl.name,
  sl.level,
  si.name AS parent_item,
  CASE
    WHEN sl.clinical_relevance IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_clinical,
  CASE
    WHEN sl.patient_explanation IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_patient,
  CASE
    WHEN sl.conduct IS NOT NULL THEN '✓'
    ELSE '✗'
  END AS has_conduct
FROM score_levels sl
JOIN score_items si ON sl.item_id = si.id
WHERE sl.last_review IS NULL
  AND sl.deleted_at IS NULL
ORDER BY sl.item_id, sl.level DESC
LIMIT 10;
EOF

echo ""
echo "=================================================="
echo "✅ Report complete"
