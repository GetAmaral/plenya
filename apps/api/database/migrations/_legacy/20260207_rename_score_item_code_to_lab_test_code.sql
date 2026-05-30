-- Migration: Rename score_items.code to score_items.lab_test_code
-- Date: 2026-02-07
-- Reason: Better semantics - this field references lab_test_definitions.code

-- Rename the column
ALTER TABLE score_items RENAME COLUMN code TO lab_test_code;

-- Drop old index and create new one with correct name
DROP INDEX IF EXISTS idx_score_items_code;
CREATE INDEX IF NOT EXISTS idx_score_items_lab_test_code ON score_items(lab_test_code);
