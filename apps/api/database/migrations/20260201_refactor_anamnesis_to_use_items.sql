-- Migration: Refactor anamnesis to use anamnesis_items
-- Date: 2026-02-01
-- Description: Simplifica a tabela anamnesis removendo campos específicos e cria anamnesis_items para registro flexível baseado em score_items

-- Create anamnesis_items table
CREATE TABLE IF NOT EXISTS anamnesis_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    anamnesis_id UUID NOT NULL,
    score_item_id UUID NOT NULL,
    text_value TEXT,
    numeric_value DOUBLE PRECISION,
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_anamnesis_items_anamnesis FOREIGN KEY (anamnesis_id) REFERENCES anamnesis(id) ON DELETE CASCADE,
    CONSTRAINT fk_anamnesis_items_score_item FOREIGN KEY (score_item_id) REFERENCES score_items(id) ON DELETE RESTRICT
);

-- Create indexes for anamnesis_items
CREATE INDEX IF NOT EXISTS idx_anamnesis_item_anamnesis ON anamnesis_items(anamnesis_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_item_score_item ON anamnesis_items(score_item_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_item_order ON anamnesis_items("order");
CREATE INDEX IF NOT EXISTS idx_anamnesis_items_deleted_at ON anamnesis_items(deleted_at);

-- Drop old columns from anamnesis table (migrate data first if needed)
ALTER TABLE anamnesis DROP COLUMN IF EXISTS chief_complaint;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS history_of_present_illness;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS past_medical_history;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS current_medications;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS allergies;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS family_history;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS social_history;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS review_of_systems;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS physical_examination;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS assessment;
ALTER TABLE anamnesis DROP COLUMN IF EXISTS plan;

-- Rename indexes on anamnesis table for consistency
DROP INDEX IF EXISTS anamnesis_patient_id_idx;
DROP INDEX IF EXISTS anamnesis_doctor_id_idx;
DROP INDEX IF EXISTS anamnesis_consultation_date_idx;

CREATE INDEX IF NOT EXISTS idx_anamnesis_patient ON anamnesis(patient_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_doctor ON anamnesis(doctor_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_consultation_date ON anamnesis(consultation_date);

-- Add comments
COMMENT ON TABLE anamnesis_items IS 'Itens de anamnese vinculados a score_items para registro flexível de dados clínicos';
COMMENT ON COLUMN anamnesis_items.anamnesis_id IS 'Referência à anamnese principal';
COMMENT ON COLUMN anamnesis_items.score_item_id IS 'Referência ao parâmetro clínico (score_item)';
COMMENT ON COLUMN anamnesis_items.text_value IS 'Valor textual/qualitativo do registro';
COMMENT ON COLUMN anamnesis_items.numeric_value IS 'Valor numérico/quantitativo do registro';
COMMENT ON COLUMN anamnesis_items."order" IS 'Ordem de exibição do item na anamnese';
