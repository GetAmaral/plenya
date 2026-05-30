-- Remove status column from lab_results table
-- +goose Up
ALTER TABLE lab_results DROP COLUMN IF EXISTS status;

-- +goose Down
ALTER TABLE lab_results ADD COLUMN status VARCHAR(20) DEFAULT 'pending' NOT NULL;
ALTER TABLE lab_results ADD CONSTRAINT lab_results_status_check CHECK (status IN ('pending', 'completed', 'cancelled'));
