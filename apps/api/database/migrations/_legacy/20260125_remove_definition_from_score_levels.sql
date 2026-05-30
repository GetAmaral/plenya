-- Remove redundant definition column from score_levels
-- name and definition were always identical

BEGIN;

ALTER TABLE score_levels DROP COLUMN IF EXISTS definition;

COMMIT;
