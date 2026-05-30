-- Create processing_jobs table for async PDF processing
CREATE TABLE IF NOT EXISTS processing_jobs (
    id UUID PRIMARY KEY,
    lab_result_batch_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    pdf_path TEXT NOT NULL,
    extracted_text TEXT,
    ai_response TEXT,
    error_message TEXT,
    attempts INTEGER NOT NULL DEFAULT 0,
    max_attempts INTEGER NOT NULL DEFAULT 3,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_lab_result_batch
        FOREIGN KEY (lab_result_batch_id)
        REFERENCES lab_result_batches(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_status
        CHECK (status IN ('pending', 'processing', 'completed', 'failed'))
);

-- Indexes for efficient polling and querying
CREATE INDEX IF NOT EXISTS idx_job_batch ON processing_jobs(lab_result_batch_id);
CREATE INDEX IF NOT EXISTS idx_job_status ON processing_jobs(status);
CREATE INDEX IF NOT EXISTS idx_job_created ON processing_jobs(created_at);
CREATE INDEX IF NOT EXISTS idx_job_deleted ON processing_jobs(deleted_at);

-- Composite index for polling query (SELECT ... WHERE status='pending' AND deleted_at IS NULL ORDER BY created_at)
CREATE INDEX IF NOT EXISTS idx_job_poll ON processing_jobs(status, created_at)
WHERE deleted_at IS NULL AND status = 'pending';

-- Add comment
COMMENT ON TABLE processing_jobs IS 'Async processing jobs for PDF extraction and AI interpretation';
COMMENT ON COLUMN processing_jobs.type IS 'Job type: pdf_extraction, etc';
COMMENT ON COLUMN processing_jobs.status IS 'Job status: pending, processing, completed, failed';
COMMENT ON COLUMN processing_jobs.attempts IS 'Number of processing attempts';
COMMENT ON COLUMN processing_jobs.extracted_text IS 'OCR extracted text from PDF';
COMMENT ON COLUMN processing_jobs.ai_response IS 'Raw JSON response from AI service';
