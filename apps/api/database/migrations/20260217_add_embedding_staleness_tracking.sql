-- Migration: Add embedding staleness tracking and regeneration queue
-- Version: 20260217
-- Description: Implements automatic invalidation of embeddings when source data changes

-- ============================================
-- Part 1: Add staleness flag to embeddings
-- ============================================

-- Add is_stale column to score_item_embeddings
ALTER TABLE score_item_embeddings
ADD COLUMN is_stale BOOLEAN DEFAULT FALSE NOT NULL;

-- Add index for efficient querying of stale embeddings
CREATE INDEX idx_score_item_embeddings_stale
ON score_item_embeddings(is_stale)
WHERE is_stale = true;

-- Add is_stale column to article_embeddings (for future use)
ALTER TABLE article_embeddings
ADD COLUMN is_stale BOOLEAN DEFAULT FALSE NOT NULL;

-- Add index for efficient querying of stale article embeddings
CREATE INDEX idx_article_embeddings_stale
ON article_embeddings(is_stale)
WHERE is_stale = true;

-- ============================================
-- Part 2: Create embedding regeneration queue
-- ============================================

CREATE TABLE embedding_queue (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    entity_type VARCHAR(50) NOT NULL CHECK (entity_type IN ('score_item', 'article')),
    entity_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'processing', 'completed', 'failed')),
    priority INT DEFAULT 0, -- Higher priority = process first (useful for critical items)
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMP,
    error_message TEXT,
    retry_count INT DEFAULT 0,
    max_retries INT DEFAULT 3,
    metadata JSONB DEFAULT '{}'::jsonb, -- For storing extra context if needed
    UNIQUE(entity_type, entity_id) -- Prevent duplicate queue entries
);

-- Indexes for efficient queue processing
CREATE INDEX idx_embedding_queue_status ON embedding_queue(status);
CREATE INDEX idx_embedding_queue_priority ON embedding_queue(priority DESC, created_at ASC)
WHERE status = 'pending';

-- ============================================
-- Part 3: Audit trail for embedding changes
-- ============================================

CREATE TABLE embedding_audit_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID NOT NULL,
    action VARCHAR(50) NOT NULL CHECK (action IN ('invalidated', 'regenerated', 'failed')),
    reason TEXT,
    old_text_source TEXT,
    new_text_source TEXT,
    triggered_by VARCHAR(100), -- e.g., 'BeforeUpdate hook', 'manual', 'batch job'
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Index for querying audit logs
CREATE INDEX idx_embedding_audit_log_entity ON embedding_audit_log(entity_type, entity_id);
CREATE INDEX idx_embedding_audit_log_created_at ON embedding_audit_log(created_at DESC);

-- ============================================
-- Part 4: Helper function to invalidate embeddings
-- ============================================

-- Function to mark embedding as stale and queue for regeneration
CREATE OR REPLACE FUNCTION invalidate_embedding(
    p_entity_type VARCHAR(50),
    p_entity_id UUID,
    p_reason TEXT DEFAULT NULL,
    p_priority INT DEFAULT 0
) RETURNS VOID AS $$
BEGIN
    -- Mark embedding as stale
    IF p_entity_type = 'score_item' THEN
        UPDATE score_item_embeddings
        SET is_stale = true
        WHERE score_item_id = p_entity_id;
    ELSIF p_entity_type = 'article' THEN
        UPDATE article_embeddings
        SET is_stale = true
        WHERE article_id = p_entity_id;
    END IF;

    -- Queue for regeneration (upsert)
    INSERT INTO embedding_queue (entity_type, entity_id, status, priority, metadata)
    VALUES (
        p_entity_type,
        p_entity_id,
        'pending',
        p_priority,
        jsonb_build_object('reason', p_reason, 'invalidated_at', NOW())
    )
    ON CONFLICT (entity_type, entity_id)
    DO UPDATE SET
        status = 'pending',
        priority = GREATEST(embedding_queue.priority, p_priority),
        created_at = NOW(),
        retry_count = 0,
        metadata = embedding_queue.metadata || jsonb_build_object('reason', p_reason, 'invalidated_at', NOW());

    -- Log to audit trail
    INSERT INTO embedding_audit_log (entity_type, entity_id, action, reason, triggered_by)
    VALUES (p_entity_type, p_entity_id, 'invalidated', p_reason, 'invalidate_embedding function');
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- Part 5: Statistics view for monitoring
-- ============================================

CREATE VIEW embedding_health_stats AS
SELECT
    'score_items' as entity_type,
    COUNT(*) FILTER (WHERE is_stale = false) as healthy_count,
    COUNT(*) FILTER (WHERE is_stale = true) as stale_count,
    ROUND(100.0 * COUNT(*) FILTER (WHERE is_stale = true) / NULLIF(COUNT(*), 0), 2) as stale_percentage
FROM score_item_embeddings
UNION ALL
SELECT
    'articles' as entity_type,
    COUNT(*) FILTER (WHERE is_stale = false) as healthy_count,
    COUNT(*) FILTER (WHERE is_stale = true) as stale_count,
    ROUND(100.0 * COUNT(*) FILTER (WHERE is_stale = true) / NULLIF(COUNT(*), 0), 2) as stale_percentage
FROM article_embeddings;

-- Grant permissions
GRANT SELECT ON embedding_health_stats TO plenya_user;

-- ============================================
-- Comments for documentation
-- ============================================

COMMENT ON TABLE embedding_queue IS 'Queue for asynchronous regeneration of stale embeddings';
COMMENT ON COLUMN embedding_queue.priority IS 'Higher priority items are processed first (0 = normal, 10 = high, 100 = critical)';
COMMENT ON COLUMN embedding_queue.retry_count IS 'Number of times this item has failed regeneration';
COMMENT ON COLUMN score_item_embeddings.is_stale IS 'True when source data changed and embedding needs regeneration';
COMMENT ON TABLE embedding_audit_log IS 'Audit trail for all embedding invalidations and regenerations';
COMMENT ON FUNCTION invalidate_embedding IS 'Helper function to mark embedding as stale and queue for regeneration';
