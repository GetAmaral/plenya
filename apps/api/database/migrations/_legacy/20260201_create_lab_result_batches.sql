-- Migration: Create lab_result_batches table and migrate existing lab_results
-- Description: Implements hierarchical batch system for lab results

-- Step 1: Create lab_result_batches table
CREATE TABLE IF NOT EXISTS lab_result_batches (
    id UUID PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    lab_request_id UUID REFERENCES lab_requests(id) ON DELETE SET NULL,
    requesting_doctor_id UUID REFERENCES users(id) ON DELETE SET NULL,
    laboratory_name VARCHAR(200) NOT NULL,
    collection_date TIMESTAMP NOT NULL,
    result_date TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'partial', 'completed')),
    observations TEXT,
    attachments TEXT, -- JSON array
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes for lab_result_batches
CREATE INDEX idx_batch_patient ON lab_result_batches(patient_id);
CREATE INDEX idx_batch_request ON lab_result_batches(lab_request_id);
CREATE INDEX idx_batch_doctor ON lab_result_batches(requesting_doctor_id);
CREATE INDEX idx_batch_collection_date ON lab_result_batches(collection_date);
CREATE INDEX idx_batch_result_date ON lab_result_batches(result_date);
CREATE INDEX idx_batch_deleted_at ON lab_result_batches(deleted_at);

-- Step 2: Add lab_result_batch_id column to lab_results (nullable for now)
ALTER TABLE lab_results ADD COLUMN IF NOT EXISTS lab_result_batch_id UUID;

-- Add lab_test_definition_id column to lab_results (optional FK)
ALTER TABLE lab_results ADD COLUMN IF NOT EXISTS lab_test_definition_id UUID REFERENCES lab_test_definitions(id) ON DELETE SET NULL;

-- Add result_text and result_numeric columns
ALTER TABLE lab_results ADD COLUMN IF NOT EXISTS result_text TEXT;
ALTER TABLE lab_results ADD COLUMN IF NOT EXISTS result_numeric DECIMAL(12,4);

-- Step 3: Migrate existing data - Create synthetic batches
-- Group existing results by patient + collection_date + laboratory
INSERT INTO lab_result_batches (
    id,
    patient_id,
    requesting_doctor_id,
    laboratory_name,
    collection_date,
    result_date,
    status,
    attachments,
    created_at,
    updated_at
)
SELECT DISTINCT ON (
    patient_id,
    COALESCE(collection_date, request_date),
    COALESCE(laboratory, 'Laboratório não informado')
)
    uuid_generate_v7() AS id,
    patient_id,
    requesting_doctor_id,
    COALESCE(laboratory, 'Laboratório não informado') AS laboratory_name,
    COALESCE(collection_date, request_date) AS collection_date,
    result_date,
    CASE
        WHEN status = 'completed' THEN 'completed'::VARCHAR(20)
        ELSE 'pending'::VARCHAR(20)
    END AS status,
    attachments,
    MIN(created_at) AS created_at,
    MAX(updated_at) AS updated_at
FROM lab_results
WHERE deleted_at IS NULL
GROUP BY
    patient_id,
    requesting_doctor_id,
    COALESCE(collection_date, request_date),
    result_date,
    COALESCE(laboratory, 'Laboratório não informado'),
    status,
    attachments;

-- Step 4: Link lab_results to batches
UPDATE lab_results lr
SET lab_result_batch_id = (
    SELECT lrb.id
    FROM lab_result_batches lrb
    WHERE lrb.patient_id = lr.patient_id
        AND lrb.collection_date = COALESCE(lr.collection_date, lr.request_date)
        AND lrb.laboratory_name = COALESCE(lr.laboratory, 'Laboratório não informado')
    LIMIT 1
)
WHERE lr.deleted_at IS NULL;

-- Step 5: Copy result to result_text (for backward compatibility)
UPDATE lab_results
SET result_text = result
WHERE result IS NOT NULL AND result_text IS NULL;

-- Step 6: Make lab_result_batch_id NOT NULL
ALTER TABLE lab_results ALTER COLUMN lab_result_batch_id SET NOT NULL;

-- Step 7: Add foreign key constraint
ALTER TABLE lab_results ADD CONSTRAINT fk_lab_results_batch
    FOREIGN KEY (lab_result_batch_id) REFERENCES lab_result_batches(id) ON DELETE CASCADE;

-- Create index for lab_result_batch_id
CREATE INDEX idx_result_batch ON lab_results(lab_result_batch_id);
CREATE INDEX idx_result_test_def ON lab_results(lab_test_definition_id);

-- Step 8: Remove old columns from lab_results
ALTER TABLE lab_results DROP COLUMN IF EXISTS patient_id;
ALTER TABLE lab_results DROP COLUMN IF EXISTS requesting_doctor_id;
ALTER TABLE lab_results DROP COLUMN IF EXISTS request_date;
ALTER TABLE lab_results DROP COLUMN IF EXISTS collection_date;
ALTER TABLE lab_results DROP COLUMN IF EXISTS result_date;
ALTER TABLE lab_results DROP COLUMN IF EXISTS laboratory;
ALTER TABLE lab_results DROP COLUMN IF EXISTS attachments;
ALTER TABLE lab_results DROP COLUMN IF EXISTS result;

-- Step 9: Update status check constraint
ALTER TABLE lab_results DROP CONSTRAINT IF EXISTS lab_results_status_check;
ALTER TABLE lab_results ADD CONSTRAINT lab_results_status_check
    CHECK (status IN ('pending', 'completed', 'cancelled'));

-- Migration complete
