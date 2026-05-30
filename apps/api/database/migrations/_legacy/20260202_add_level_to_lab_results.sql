-- Migration: Add level column to lab_results
-- Description: Adiciona campo level (integer) para indicar nível/prioridade do resultado

-- Add level column
ALTER TABLE lab_results
ADD COLUMN level INTEGER;

-- Add comment
COMMENT ON COLUMN lab_results.level IS 'Nível/prioridade do resultado laboratorial';
