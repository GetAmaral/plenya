-- Add is_default column to methods table
-- Ensures only one method can be default at a time

-- Add column
ALTER TABLE methods
ADD COLUMN is_default BOOLEAN NOT NULL DEFAULT false;

-- Add index for performance
CREATE INDEX idx_method_is_default ON methods(is_default);

-- Comment
COMMENT ON COLUMN methods.is_default IS 'Indicates if this is the default method (only one can be default at a time)';
