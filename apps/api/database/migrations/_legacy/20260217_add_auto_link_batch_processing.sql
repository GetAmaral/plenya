-- Migration: Add batch processing state tracking for auto-link operations
-- Version: 20260217
-- Description: Implements checkpoint/rollback system for resilient batch processing

-- ============================================
-- Part 1: Processing state tracking table
-- ============================================

CREATE TABLE auto_link_processing_state (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    run_id UUID NOT NULL UNIQUE,
    batch_size INT NOT NULL DEFAULT 50,
    last_processed_item_id UUID,
    last_processed_index INT DEFAULT 0,
    total_items INT DEFAULT 0,
    total_processed INT DEFAULT 0,
    total_linked INT DEFAULT 0,
    total_skipped INT DEFAULT 0,
    failed_items JSONB DEFAULT '[]'::jsonb,
    started_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_checkpoint_at TIMESTAMP,
    completed_at TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'running' CHECK (status IN ('running', 'completed', 'failed', 'cancelled')),
    error_message TEXT,
    metadata JSONB DEFAULT '{}'::jsonb
);

-- Indexes for efficient querying
CREATE INDEX idx_auto_link_state_run_id ON auto_link_processing_state(run_id);
CREATE INDEX idx_auto_link_state_status ON auto_link_processing_state(status);
CREATE INDEX idx_auto_link_state_started_at ON auto_link_processing_state(started_at DESC);

-- ============================================
-- Part 2: Batch checkpoint details
-- ============================================

CREATE TABLE auto_link_batch_checkpoints (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    run_id UUID NOT NULL REFERENCES auto_link_processing_state(run_id) ON DELETE CASCADE,
    batch_number INT NOT NULL,
    batch_start_index INT NOT NULL,
    batch_end_index INT NOT NULL,
    items_processed INT NOT NULL,
    links_created INT NOT NULL,
    items_skipped INT NOT NULL,
    items_failed INT NOT NULL,
    avg_confidence DOUBLE PRECISION,
    processing_time_ms INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(run_id, batch_number)
);

-- Indexes
CREATE INDEX idx_batch_checkpoints_run_id ON auto_link_batch_checkpoints(run_id);

-- ============================================
-- Part 3: Item-level processing log (for retry)
-- ============================================

CREATE TABLE auto_link_item_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    run_id UUID NOT NULL REFERENCES auto_link_processing_state(run_id) ON DELETE CASCADE,
    score_item_id UUID NOT NULL,
    score_item_name VARCHAR(300),
    status VARCHAR(20) NOT NULL CHECK (status IN ('success', 'skipped', 'failed')),
    links_created INT DEFAULT 0,
    avg_confidence DOUBLE PRECISION,
    error_message TEXT,
    processing_time_ms INT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_item_log_run_id ON auto_link_item_log(run_id);
CREATE INDEX idx_item_log_score_item ON auto_link_item_log(score_item_id);
CREATE INDEX idx_item_log_status ON auto_link_item_log(status);

-- ============================================
-- Part 4: Statistics views
-- ============================================

-- Run summary view
CREATE VIEW auto_link_run_summary AS
SELECT
    run_id,
    status,
    total_items,
    total_processed,
    total_linked,
    total_skipped,
    ROUND(100.0 * total_processed / NULLIF(total_items, 0), 2) as completion_percentage,
    ROUND(100.0 * total_linked / NULLIF(total_processed, 0), 2) as link_success_rate,
    EXTRACT(EPOCH FROM (COALESCE(completed_at, NOW()) - started_at)) as duration_seconds,
    started_at,
    completed_at,
    jsonb_array_length(failed_items) as failed_count
FROM auto_link_processing_state
ORDER BY started_at DESC;

-- Batch performance view
CREATE VIEW auto_link_batch_performance AS
SELECT
    bc.run_id,
    bc.batch_number,
    bc.items_processed,
    bc.links_created,
    bc.items_failed,
    ROUND(bc.avg_confidence, 3) as avg_confidence,
    bc.processing_time_ms,
    ROUND(bc.processing_time_ms::numeric / NULLIF(bc.items_processed, 0), 2) as ms_per_item,
    bc.created_at
FROM auto_link_batch_checkpoints bc
ORDER BY bc.run_id, bc.batch_number;

-- Item failure analysis view
CREATE VIEW auto_link_failure_analysis AS
SELECT
    score_item_id,
    score_item_name,
    COUNT(*) as failure_count,
    array_agg(DISTINCT error_message) as error_messages,
    MAX(created_at) as last_failed_at
FROM auto_link_item_log
WHERE status = 'failed'
GROUP BY score_item_id, score_item_name
HAVING COUNT(*) > 1 -- Items that failed multiple times
ORDER BY failure_count DESC;

-- Grant permissions
GRANT SELECT ON auto_link_run_summary TO plenya_user;
GRANT SELECT ON auto_link_batch_performance TO plenya_user;
GRANT SELECT ON auto_link_failure_analysis TO plenya_user;

-- ============================================
-- Part 5: Helper functions
-- ============================================

-- Function to create a new processing run
CREATE OR REPLACE FUNCTION create_auto_link_run(
    p_total_items INT,
    p_batch_size INT DEFAULT 50
) RETURNS UUID AS $$
DECLARE
    v_run_id UUID;
BEGIN
    v_run_id := uuid_generate_v7();

    INSERT INTO auto_link_processing_state (run_id, total_items, batch_size, status)
    VALUES (v_run_id, p_total_items, p_batch_size, 'running');

    RETURN v_run_id;
END;
$$ LANGUAGE plpgsql;

-- Function to save a batch checkpoint
CREATE OR REPLACE FUNCTION save_batch_checkpoint(
    p_run_id UUID,
    p_batch_number INT,
    p_start_index INT,
    p_end_index INT,
    p_items_processed INT,
    p_links_created INT,
    p_items_skipped INT,
    p_items_failed INT,
    p_avg_confidence DOUBLE PRECISION,
    p_processing_time_ms INT,
    p_last_item_id UUID
) RETURNS VOID AS $$
BEGIN
    -- Insert batch checkpoint
    INSERT INTO auto_link_batch_checkpoints (
        run_id, batch_number, batch_start_index, batch_end_index,
        items_processed, links_created, items_skipped, items_failed,
        avg_confidence, processing_time_ms
    ) VALUES (
        p_run_id, p_batch_number, p_start_index, p_end_index,
        p_items_processed, p_links_created, p_items_skipped, p_items_failed,
        p_avg_confidence, p_processing_time_ms
    );

    -- Update run state
    UPDATE auto_link_processing_state
    SET
        last_processed_item_id = p_last_item_id,
        last_processed_index = p_end_index,
        total_processed = total_processed + p_items_processed,
        total_linked = total_linked + p_links_created,
        total_skipped = total_skipped + p_items_skipped,
        last_checkpoint_at = NOW()
    WHERE run_id = p_run_id;
END;
$$ LANGUAGE plpgsql;

-- Function to complete a run
CREATE OR REPLACE FUNCTION complete_auto_link_run(
    p_run_id UUID,
    p_status VARCHAR(20),
    p_error_message TEXT DEFAULT NULL
) RETURNS VOID AS $$
BEGIN
    UPDATE auto_link_processing_state
    SET
        status = p_status,
        completed_at = NOW(),
        error_message = p_error_message
    WHERE run_id = p_run_id;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- Comments for documentation
-- ============================================

COMMENT ON TABLE auto_link_processing_state IS 'Tracks state of auto-link batch processing runs for crash recovery';
COMMENT ON TABLE auto_link_batch_checkpoints IS 'Detailed checkpoint data for each batch processed';
COMMENT ON TABLE auto_link_item_log IS 'Item-level processing log for debugging and retry logic';
COMMENT ON COLUMN auto_link_processing_state.run_id IS 'Unique identifier for this processing run';
COMMENT ON COLUMN auto_link_processing_state.last_processed_item_id IS 'UUID of last successfully processed ScoreItem';
COMMENT ON COLUMN auto_link_processing_state.failed_items IS 'JSONB array of failed item IDs for retry';
COMMENT ON FUNCTION create_auto_link_run IS 'Creates a new auto-link processing run with initial state';
COMMENT ON FUNCTION save_batch_checkpoint IS 'Saves checkpoint after each batch completes';
COMMENT ON FUNCTION complete_auto_link_run IS 'Marks a run as completed or failed';
