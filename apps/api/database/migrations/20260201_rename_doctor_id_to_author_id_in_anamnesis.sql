-- Migration: Rename doctor_id to author_id in anamnesis table
-- Date: 2026-02-01
-- Description: Renomeia doctor_id para author_id para permitir que outros profissionais (psicólogos, enfermeiros) também possam criar anamneses

-- Rename column
ALTER TABLE anamnesis RENAME COLUMN doctor_id TO author_id;

-- Update comment
COMMENT ON COLUMN anamnesis.author_id IS 'ID do profissional que realizou a anamnese (médico, psicólogo, enfermeiro, etc.)';

-- Drop old indexes
DROP INDEX IF EXISTS idx_anamnesis_doctor;
DROP INDEX IF EXISTS idx_anamnesis_doctor_id;

-- Create new index
CREATE INDEX IF NOT EXISTS idx_anamnesis_author ON anamnesis(author_id);

-- Drop old foreign key constraint
ALTER TABLE anamnesis DROP CONSTRAINT IF EXISTS fk_anamnesis_doctor;

-- Create new foreign key constraint
ALTER TABLE anamnesis
ADD CONSTRAINT fk_anamnesis_author
FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE RESTRICT;
