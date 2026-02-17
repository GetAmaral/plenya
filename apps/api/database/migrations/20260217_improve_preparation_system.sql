-- Migration: Improve preparation system with stale tracking
-- Version: 20260217
-- Description: Add stale status and metadata enhancements for preparation quality tracking

-- ============================================
-- Part 1: Add 'stale' status to preparations
-- ============================================

-- Drop existing constraint
ALTER TABLE score_item_enrichment_preparation
DROP CONSTRAINT IF EXISTS score_item_enrichment_preparation_status_check;

-- Add 'stale' to allowed values
ALTER TABLE score_item_enrichment_preparation
ADD CONSTRAINT score_item_enrichment_preparation_status_check
CHECK (status IN ('ready','processing','completed','failed','stale'));

-- Add index for stale items
CREATE INDEX IF NOT EXISTS idx_preparation_stale
ON score_item_enrichment_preparation(status)
WHERE status = 'stale';

COMMENT ON COLUMN score_item_enrichment_preparation.status IS 'Status: ready (awaiting enrichment), processing (being enriched), completed (enriched), failed (enrichment failed), stale (needs re-preparation)';

-- ============================================
-- Part 2: Analytics views for preparation quality
-- ============================================

-- View: Preparations by quality grade
CREATE OR REPLACE VIEW preparation_quality_stats AS
SELECT
    metadata->>'quality_grade' as quality_grade,
    COUNT(*) as count,
    ROUND(100.0 * COUNT(*) / SUM(COUNT(*)) OVER (), 1) as percentage,
    ROUND(AVG((metadata->>'total_chunks')::int), 1) as avg_chunks,
    ROUND(AVG((metadata->>'avg_similarity')::float)::numeric, 3) as avg_similarity
FROM score_item_enrichment_preparation
WHERE metadata->>'quality_grade' IS NOT NULL
GROUP BY metadata->>'quality_grade'
ORDER BY
    CASE metadata->>'quality_grade'
        WHEN 'excellent' THEN 1
        WHEN 'good' THEN 2
        WHEN 'fair' THEN 3
        ELSE 4
    END;

-- View: Preparations needing attention
CREATE OR REPLACE VIEW preparation_needs_attention AS
SELECT
    si.id,
    si.name,
    (prep.metadata->>'total_chunks')::int as chunks,
    ROUND((prep.metadata->>'avg_similarity')::float::numeric, 3) as avg_sim,
    prep.metadata->>'quality_grade' as grade,
    prep.metadata->>'threshold_used' as threshold,
    prep.status
FROM score_item_enrichment_preparation prep
JOIN score_items si ON si.id = prep.score_item_id
WHERE (prep.metadata->>'total_chunks')::int < 15
   OR (prep.metadata->>'avg_similarity')::float < 0.4
   OR prep.status = 'stale'
ORDER BY (prep.metadata->>'total_chunks')::int ASC, (prep.metadata->>'avg_similarity')::float ASC
LIMIT 100;

-- View: Overall preparation health
CREATE OR REPLACE VIEW preparation_health_summary AS
SELECT
    COUNT(*) as total_preparations,
    COUNT(*) FILTER (WHERE status = 'ready') as ready,
    COUNT(*) FILTER (WHERE status = 'stale') as stale,
    COUNT(*) FILTER (WHERE status = 'completed') as completed,
    ROUND(AVG((metadata->>'total_chunks')::int), 1) as avg_chunks,
    ROUND(AVG((metadata->>'avg_similarity')::float)::numeric, 3) as avg_similarity,
    COUNT(*) FILTER (WHERE (metadata->>'total_chunks')::int < 10) as weak_count,
    COUNT(*) FILTER (WHERE metadata->>'quality_grade' = 'excellent') as excellent_count,
    COUNT(*) FILTER (WHERE metadata->>'quality_grade' = 'good') as good_count,
    COUNT(*) FILTER (WHERE metadata->>'quality_grade' = 'fair') as fair_count,
    COUNT(*) FILTER (WHERE metadata->>'quality_grade' = 'poor') as poor_count
FROM score_item_enrichment_preparation;

-- Grant permissions
GRANT SELECT ON preparation_quality_stats TO plenya_user;
GRANT SELECT ON preparation_needs_attention TO plenya_user;
GRANT SELECT ON preparation_health_summary TO plenya_user;

-- ============================================
-- Part 3: Helper function to invalidate preparation
-- ============================================

CREATE OR REPLACE FUNCTION invalidate_preparation(
    p_score_item_id UUID,
    p_reason TEXT DEFAULT NULL
) RETURNS VOID AS $$
BEGIN
    -- Mark as stale or delete (currently deleting for simplicity)
    DELETE FROM score_item_enrichment_preparation
    WHERE score_item_id = p_score_item_id;

    -- Could alternatively mark as stale:
    -- UPDATE score_item_enrichment_preparation
    -- SET status = 'stale'
    -- WHERE score_item_id = p_score_item_id;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION invalidate_preparation IS 'Invalidate preparation when ScoreItem changes (deletes for re-preparation)';

-- ============================================
-- Comments
-- ============================================

COMMENT ON VIEW preparation_quality_stats IS 'Preparation quality distribution by grade';
COMMENT ON VIEW preparation_needs_attention IS 'Preparations with low quality that need re-processing';
COMMENT ON VIEW preparation_health_summary IS 'Overall health metrics of preparation system';
