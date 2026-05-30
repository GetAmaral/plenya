-- Migration: Increase score_levels.name length from VARCHAR(200) to VARCHAR(500)
-- Generated from Go model change in apps/api/internal/models/score_level.go
-- Reason: Some level names in CSV exceeded 200 characters
-- Date: 2026-01-24

-- Alter column type
ALTER TABLE score_levels
  ALTER COLUMN name TYPE VARCHAR(500);

-- Update comment
COMMENT ON COLUMN score_levels.name IS 'Descrição do nível (ex: 55 a 70 (Ótimo)) - max 500 chars';
