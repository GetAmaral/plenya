-- Add pdf_content_json field to lab_result_batches table
-- Migration: add_pdf_content_json
-- Created: 2026-02-06

ALTER TABLE lab_result_batches
ADD COLUMN pdf_content_json TEXT;

COMMENT ON COLUMN lab_result_batches.pdf_content_json IS 'Conteúdo processado pela IA em formato JSON estruturado';
