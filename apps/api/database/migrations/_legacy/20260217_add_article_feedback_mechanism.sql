-- Migration: Add user feedback mechanism for article-scoreitem links
-- Version: 20260217
-- Description: Allows users to approve/reject auto-linked articles for quality tracking

-- ============================================
-- Part 1: Add feedback columns to article_score_items
-- ============================================

ALTER TABLE article_score_items
ADD COLUMN IF NOT EXISTS user_feedback VARCHAR(20) CHECK (user_feedback IN ('approved', 'rejected', 'irrelevant'));

ALTER TABLE article_score_items
ADD COLUMN IF NOT EXISTS feedback_at TIMESTAMP;

ALTER TABLE article_score_items
ADD COLUMN IF NOT EXISTS feedback_by UUID; -- References users(id) but no FK constraint (user may be deleted)

COMMENT ON COLUMN article_score_items.user_feedback IS 'User feedback on link quality: approved (good match), rejected (bad match), irrelevant (not related)';
COMMENT ON COLUMN article_score_items.feedback_at IS 'Timestamp when feedback was provided';
COMMENT ON COLUMN article_score_items.feedback_by IS 'User ID who provided feedback';

-- ============================================
-- Part 2: Add index for efficient feedback queries
-- ============================================

CREATE INDEX IF NOT EXISTS idx_article_score_items_feedback
ON article_score_items(user_feedback)
WHERE user_feedback IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_article_score_items_auto_linked_feedback
ON article_score_items(auto_linked, user_feedback)
WHERE auto_linked = true;

-- ============================================
-- Part 3: Analytics views for precision tracking
-- ============================================

-- View: Precision by confidence bucket
CREATE OR REPLACE VIEW article_link_precision_by_confidence AS
SELECT
    CASE
        WHEN confidence_score >= 0.8 THEN '0.8-1.0 (high)'
        WHEN confidence_score >= 0.6 THEN '0.6-0.8 (medium)'
        WHEN confidence_score >= 0.4 THEN '0.4-0.6 (low)'
        ELSE '0.0-0.4 (very low)'
    END AS confidence_bucket,
    COUNT(*) FILTER (WHERE user_feedback = 'approved') AS approved_count,
    COUNT(*) FILTER (WHERE user_feedback = 'rejected') AS rejected_count,
    COUNT(*) FILTER (WHERE user_feedback = 'irrelevant') AS irrelevant_count,
    COUNT(*) AS total_with_feedback,
    ROUND(
        100.0 * COUNT(*) FILTER (WHERE user_feedback = 'approved') /
        NULLIF(COUNT(*), 0),
        2
    ) AS precision_percentage,
    ROUND(AVG(confidence_score)::numeric, 3) AS avg_confidence
FROM article_score_items
WHERE user_feedback IS NOT NULL AND auto_linked = true
GROUP BY confidence_bucket
ORDER BY avg_confidence DESC;

-- View: Most rejected score items (need improvement)
CREATE OR REPLACE VIEW article_link_problematic_items AS
SELECT
    si.id,
    si.name,
    si.full_name,
    COUNT(*) FILTER (WHERE asi.user_feedback = 'rejected') AS rejected_count,
    COUNT(*) FILTER (WHERE asi.user_feedback = 'approved') AS approved_count,
    COUNT(*) AS total_feedback,
    ROUND(
        100.0 * COUNT(*) FILTER (WHERE asi.user_feedback = 'rejected') /
        NULLIF(COUNT(*), 0),
        2
    ) AS rejection_rate,
    ROUND(AVG(asi.confidence_score)::numeric, 3) AS avg_confidence
FROM score_items si
JOIN article_score_items asi ON asi.score_item_id = si.id
WHERE asi.user_feedback IS NOT NULL AND asi.auto_linked = true
GROUP BY si.id, si.name, si.full_name
HAVING COUNT(*) FILTER (WHERE asi.user_feedback = 'rejected') > 0
ORDER BY rejection_rate DESC, rejected_count DESC
LIMIT 50;

-- View: Overall feedback stats
CREATE OR REPLACE VIEW article_link_feedback_stats AS
SELECT
    COUNT(*) AS total_auto_linked,
    COUNT(*) FILTER (WHERE user_feedback IS NOT NULL) AS total_with_feedback,
    COUNT(*) FILTER (WHERE user_feedback = 'approved') AS approved,
    COUNT(*) FILTER (WHERE user_feedback = 'rejected') AS rejected,
    COUNT(*) FILTER (WHERE user_feedback = 'irrelevant') AS irrelevant,
    ROUND(
        100.0 * COUNT(*) FILTER (WHERE user_feedback IS NOT NULL) /
        NULLIF(COUNT(*), 0),
        2
    ) AS feedback_coverage_percentage,
    ROUND(
        100.0 * COUNT(*) FILTER (WHERE user_feedback = 'approved') /
        NULLIF(COUNT(*) FILTER (WHERE user_feedback IS NOT NULL), 0),
        2
    ) AS precision_percentage,
    ROUND(AVG(confidence_score)::numeric, 3) AS avg_confidence
FROM article_score_items
WHERE auto_linked = true;

-- Grant permissions
GRANT SELECT ON article_link_precision_by_confidence TO plenya_user;
GRANT SELECT ON article_link_problematic_items TO plenya_user;
GRANT SELECT ON article_link_feedback_stats TO plenya_user;

-- ============================================
-- Part 4: Helper function to submit feedback
-- ============================================

CREATE OR REPLACE FUNCTION submit_article_link_feedback(
    p_score_item_id UUID,
    p_article_id UUID,
    p_feedback VARCHAR(20),
    p_user_id UUID
) RETURNS BOOLEAN AS $$
DECLARE
    v_exists BOOLEAN;
BEGIN
    -- Validate feedback value
    IF p_feedback NOT IN ('approved', 'rejected', 'irrelevant') THEN
        RAISE EXCEPTION 'Invalid feedback value: %. Must be: approved, rejected, or irrelevant', p_feedback;
    END IF;

    -- Check if link exists
    SELECT EXISTS(
        SELECT 1 FROM article_score_items
        WHERE score_item_id = p_score_item_id AND article_id = p_article_id
    ) INTO v_exists;

    IF NOT v_exists THEN
        RAISE EXCEPTION 'Article-ScoreItem link not found';
    END IF;

    -- Update feedback
    UPDATE article_score_items
    SET
        user_feedback = p_feedback,
        feedback_at = NOW(),
        feedback_by = p_user_id
    WHERE score_item_id = p_score_item_id AND article_id = p_article_id;

    RETURN TRUE;
END;
$$ LANGUAGE plpgsql;

COMMENT ON FUNCTION submit_article_link_feedback IS 'Submit user feedback for an article-scoreitem link';

-- ============================================
-- Comments for documentation
-- ============================================

COMMENT ON VIEW article_link_precision_by_confidence IS 'Precision metrics grouped by confidence score buckets';
COMMENT ON VIEW article_link_problematic_items IS 'ScoreItems with high rejection rates that need review';
COMMENT ON VIEW article_link_feedback_stats IS 'Overall feedback statistics for auto-linked articles';
