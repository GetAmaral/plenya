-- Migration: Create Score System Tables
-- Generated from Go models in apps/api/internal/models/score_*.go
-- Date: 2026-01-21

-- Enable UUID extension if not exists
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: score_groups
CREATE TABLE IF NOT EXISTS score_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    "order" INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Indexes for score_groups
CREATE UNIQUE INDEX IF NOT EXISTS idx_score_group_name ON score_groups(name) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_score_group_order ON score_groups("order");
CREATE INDEX IF NOT EXISTS idx_score_groups_deleted_at ON score_groups(deleted_at);

-- Table: score_subgroups
CREATE TABLE IF NOT EXISTS score_subgroups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    "order" INTEGER NOT NULL DEFAULT 0,
    group_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_score_subgroups_group FOREIGN KEY (group_id) REFERENCES score_groups(id) ON DELETE CASCADE
);

-- Indexes for score_subgroups
CREATE INDEX IF NOT EXISTS idx_score_subgroup_group ON score_subgroups(group_id);
CREATE INDEX IF NOT EXISTS idx_score_subgroup_order ON score_subgroups("order");
CREATE INDEX IF NOT EXISTS idx_score_subgroups_deleted_at ON score_subgroups(deleted_at);

-- Table: score_items
CREATE TABLE IF NOT EXISTS score_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(300) NOT NULL,
    unit VARCHAR(50),
    unit_conversion TEXT,
    points DOUBLE PRECISION NOT NULL DEFAULT 0,
    "order" INTEGER NOT NULL DEFAULT 0,
    subgroup_id UUID NOT NULL,
    parent_item_id UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_score_items_subgroup FOREIGN KEY (subgroup_id) REFERENCES score_subgroups(id) ON DELETE CASCADE,
    CONSTRAINT fk_score_items_parent FOREIGN KEY (parent_item_id) REFERENCES score_items(id) ON DELETE SET NULL,
    CONSTRAINT chk_score_items_points CHECK (points >= 0 AND points <= 100)
);

-- Indexes for score_items
CREATE INDEX IF NOT EXISTS idx_score_item_subgroup ON score_items(subgroup_id);
CREATE INDEX IF NOT EXISTS idx_score_item_parent ON score_items(parent_item_id);
CREATE INDEX IF NOT EXISTS idx_score_item_order ON score_items("order");
CREATE INDEX IF NOT EXISTS idx_score_items_deleted_at ON score_items(deleted_at);

-- Table: score_levels
CREATE TABLE IF NOT EXISTS score_levels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    level INTEGER NOT NULL,
    name VARCHAR(200) NOT NULL,
    definition TEXT,
    lower_limit VARCHAR(50),
    upper_limit VARCHAR(50),
    item_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_score_levels_item FOREIGN KEY (item_id) REFERENCES score_items(id) ON DELETE CASCADE,
    CONSTRAINT chk_score_levels_level CHECK (level >= 0 AND level <= 6)
);

-- Indexes for score_levels
CREATE INDEX IF NOT EXISTS idx_score_level_item ON score_levels(item_id);
CREATE INDEX IF NOT EXISTS idx_score_level_level ON score_levels(level);
CREATE INDEX IF NOT EXISTS idx_score_levels_deleted_at ON score_levels(deleted_at);

-- Comments
COMMENT ON TABLE score_groups IS 'Grupos principais de escores clínicos (ex: Hemograma, Lipídeos)';
COMMENT ON TABLE score_subgroups IS 'Subgrupos dentro de um grupo de escores (ex: Série Vermelha, Série Branca)';
COMMENT ON TABLE score_items IS 'Itens individuais de escores (ex: Hemoglobina - Homens, FEVE)';
COMMENT ON TABLE score_levels IS 'Níveis de estratificação de risco para cada item (0-6)';

COMMENT ON COLUMN score_groups.name IS 'Nome do grupo de escores';
COMMENT ON COLUMN score_groups."order" IS 'Ordem de exibição do grupo';

COMMENT ON COLUMN score_subgroups.name IS 'Nome do subgrupo';
COMMENT ON COLUMN score_subgroups."order" IS 'Ordem de exibição do subgrupo dentro do grupo';
COMMENT ON COLUMN score_subgroups.group_id IS 'Referência ao grupo pai';

COMMENT ON COLUMN score_items.name IS 'Nome do parâmetro clínico';
COMMENT ON COLUMN score_items.unit IS 'Unidade de medida (ex: g/dL, %, mm)';
COMMENT ON COLUMN score_items.unit_conversion IS 'Fórmula de conversão de unidades (ex: 1 g/dL = 10 g/L)';
COMMENT ON COLUMN score_items.points IS 'Pontos máximos para este item no escore (0-100)';
COMMENT ON COLUMN score_items."order" IS 'Ordem de exibição do item dentro do subgrupo';
COMMENT ON COLUMN score_items.subgroup_id IS 'Referência ao subgrupo pai';
COMMENT ON COLUMN score_items.parent_item_id IS 'Referência a item pai para hierarquia (opcional)';

COMMENT ON COLUMN score_levels.level IS 'Nível de risco (0=crítico, 1-4=intermediários, 5=ótimo, 6=reservado)';
COMMENT ON COLUMN score_levels.name IS 'Descrição do nível (ex: 55 a 70 (Ótimo))';
COMMENT ON COLUMN score_levels.definition IS 'Definição clínica do nível';
COMMENT ON COLUMN score_levels.lower_limit IS 'Limite inferior do intervalo';
COMMENT ON COLUMN score_levels.upper_limit IS 'Limite superior do intervalo';
COMMENT ON COLUMN score_levels.item_id IS 'Referência ao item de escore';
