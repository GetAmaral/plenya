-- Add operator column to score_levels table
-- Migration: 20260124_add_operator_to_score_levels
-- Description: Adds operator field to define comparison type (=, >, >=, <, <=, between)

-- Add operator column with default value
ALTER TABLE score_levels
ADD COLUMN operator VARCHAR(10) NOT NULL DEFAULT 'between';

-- Add CHECK constraint for valid operator values
ALTER TABLE score_levels
ADD CONSTRAINT chk_score_levels_operator
CHECK (operator IN ('=', '>', '>=', '<', '<=', 'between'));

-- Create index for operator column (useful for queries filtering by operator type)
CREATE INDEX idx_score_levels_operator ON score_levels(operator);

-- Comment for documentation
COMMENT ON COLUMN score_levels.operator IS 'Comparison operator: = (equals), > (greater than), >= (greater or equal), < (less than), <= (less or equal), between (range)';
