-- Migration: add_menopause_to_patients
-- Created at: 2026-02-08
-- Description: Adiciona campo menopause à tabela patients (apenas para gender=female)

-- Add menopause column
ALTER TABLE patients ADD COLUMN menopause BOOLEAN DEFAULT NULL;

-- Add comment
COMMENT ON COLUMN patients.menopause IS 'Indica se a paciente está na menopausa (apenas para gender=female)';
