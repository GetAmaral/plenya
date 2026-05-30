-- Migration: Create anamnesis templates and template items
-- Date: 2026-02-01
-- Description: Cria tabelas de templates de anamnese para facilitar criação de anamneses padronizadas por área

-- Create anamnesis_templates table
CREATE TABLE IF NOT EXISTS anamnesis_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    area VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT chk_name_not_empty CHECK (length(trim(name)) >= 2)
);

-- Create indexes for anamnesis_templates
CREATE INDEX IF NOT EXISTS idx_anamnesis_template_area ON anamnesis_templates(area);
CREATE INDEX IF NOT EXISTS idx_anamnesis_templates_deleted_at ON anamnesis_templates(deleted_at);

-- Create anamnesis_template_items table
CREATE TABLE IF NOT EXISTS anamnesis_template_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    anamnesis_template_id UUID NOT NULL,
    score_item_id UUID NOT NULL,
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_anamnesis_template_items_template FOREIGN KEY (anamnesis_template_id) REFERENCES anamnesis_templates(id) ON DELETE CASCADE,
    CONSTRAINT fk_anamnesis_template_items_score_item FOREIGN KEY (score_item_id) REFERENCES score_items(id) ON DELETE RESTRICT,

    -- Constraints
    CONSTRAINT chk_order_range CHECK ("order" >= 0 AND "order" <= 9999)
);

-- Create indexes for anamnesis_template_items
CREATE INDEX IF NOT EXISTS idx_anamnesis_template_item_template ON anamnesis_template_items(anamnesis_template_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_template_item_score_item ON anamnesis_template_items(score_item_id);
CREATE INDEX IF NOT EXISTS idx_anamnesis_template_item_order ON anamnesis_template_items("order");
CREATE INDEX IF NOT EXISTS idx_anamnesis_template_items_deleted_at ON anamnesis_template_items(deleted_at);

-- Add anamnesis_template_id to anamnesis table
ALTER TABLE anamnesis
ADD COLUMN IF NOT EXISTS anamnesis_template_id UUID;

-- Create index for anamnesis_template_id
CREATE INDEX IF NOT EXISTS idx_anamnesis_template ON anamnesis(anamnesis_template_id);

-- Add foreign key constraint
ALTER TABLE anamnesis
ADD CONSTRAINT IF NOT EXISTS fk_anamnesis_template
FOREIGN KEY (anamnesis_template_id) REFERENCES anamnesis_templates(id) ON DELETE SET NULL;

-- Add comments
COMMENT ON TABLE anamnesis_templates IS 'Templates de anamnese para diferentes áreas médicas';
COMMENT ON COLUMN anamnesis_templates.name IS 'Nome do template (ex: Anamnese Clínica Geral)';
COMMENT ON COLUMN anamnesis_templates.area IS 'Área de aplicação (ex: Clínica Médica, Psicologia, Pediatria)';

COMMENT ON TABLE anamnesis_template_items IS 'Itens que compõem um template de anamnese';
COMMENT ON COLUMN anamnesis_template_items.anamnesis_template_id IS 'Referência ao template de anamnese';
COMMENT ON COLUMN anamnesis_template_items.score_item_id IS 'Referência ao item de score a ser incluído';
COMMENT ON COLUMN anamnesis_template_items."order" IS 'Ordem de exibição do item no template';

COMMENT ON COLUMN anamnesis.anamnesis_template_id IS 'Template utilizado para criar esta anamnese (opcional)';
