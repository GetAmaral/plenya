-- Migration: Add content, summary and visibility fields to anamnesis
-- Date: 2026-02-01
-- Description: Adiciona campos de conteúdo, resumo e visibilidade à tabela anamnesis

-- Add new columns
ALTER TABLE anamnesis
ADD COLUMN IF NOT EXISTS content TEXT,
ADD COLUMN IF NOT EXISTS summary TEXT,
ADD COLUMN IF NOT EXISTS visibility VARCHAR(20) NOT NULL DEFAULT 'all' CHECK (visibility IN ('all', 'medicalOnly', 'psychOnly'));

-- Create index for visibility filtering
CREATE INDEX IF NOT EXISTS idx_anamnesis_visibility ON anamnesis(visibility);

-- Add comments
COMMENT ON COLUMN anamnesis.content IS 'Conteúdo completo da anamnese';
COMMENT ON COLUMN anamnesis.summary IS 'Resumo da anamnese';
COMMENT ON COLUMN anamnesis.visibility IS 'Visibilidade da anamnese: all (todos), medicalOnly (apenas médicos), psychOnly (apenas psicólogos)';
