-- Add matched field to lab_results table
ALTER TABLE lab_results
ADD COLUMN IF NOT EXISTS matched BOOLEAN NOT NULL DEFAULT true;

-- Index for filtering matched/unmatched results
CREATE INDEX IF NOT EXISTS idx_lab_result_matched ON lab_results(matched);

-- Add comment
COMMENT ON COLUMN lab_results.matched IS 'Whether the result was matched with a LabTestDefinition (true) or extracted but not cataloged (false)';
