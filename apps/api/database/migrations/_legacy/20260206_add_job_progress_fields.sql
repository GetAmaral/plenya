-- Add progress tracking fields to processing_jobs table
-- Migration: add_job_progress_fields
-- Created: 2026-02-06

ALTER TABLE processing_jobs
ADD COLUMN progress_step INTEGER,
ADD COLUMN progress_message TEXT;

COMMENT ON COLUMN processing_jobs.progress_step IS 'Etapa atual do processamento (1-6)';
COMMENT ON COLUMN processing_jobs.progress_message IS 'Mensagem descritiva da etapa atual';
