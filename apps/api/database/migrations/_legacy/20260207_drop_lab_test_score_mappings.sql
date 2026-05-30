-- Migration: Drop lab_test_score_mappings table
-- Date: 2026-02-07
-- Reason: Mapping is done directly via code field (labTestDefinition.code = scoreItem.code)
-- The intermediate table is redundant since both tables share the same code values

-- Drop the table and all related objects
DROP TABLE IF EXISTS lab_test_score_mappings CASCADE;

-- Drop any lingering indexes (if they weren't dropped with CASCADE)
DROP INDEX IF EXISTS idx_lab_test_score_mapping_test;
DROP INDEX IF EXISTS idx_lab_test_score_mapping_item;
DROP INDEX IF EXISTS idx_lab_test_score_mapping_gender;
DROP INDEX IF EXISTS idx_lab_test_score_mapping_active;
DROP INDEX IF EXISTS idx_lab_test_score_mapping_deleted;

-- Drop the trigger function if it exists
DROP FUNCTION IF EXISTS update_lab_test_score_mappings_updated_at() CASCADE;
