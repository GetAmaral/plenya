-- Migration: Create Lab Test System (Definitions, Mappings, Reference Ranges, Values)
-- Created: 2026-01-25

-- ============================================================
-- 1. Add 'code' field to score_items
-- ============================================================

ALTER TABLE score_items
ADD COLUMN IF NOT EXISTS code VARCHAR(100) UNIQUE;

CREATE INDEX IF NOT EXISTS idx_score_items_code ON score_items(code);

-- ============================================================
-- 2. Create lab_test_definitions table
-- ============================================================

CREATE TABLE IF NOT EXISTS lab_test_definitions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(300) NOT NULL,
    short_name VARCHAR(50),
    tuss_code VARCHAR(20),
    loinc_code VARCHAR(20),
    category VARCHAR(30) NOT NULL,
    is_requestable BOOLEAN NOT NULL DEFAULT true,
    parent_test_id UUID,
    unit VARCHAR(50),
    unit_conversion TEXT,
    result_type VARCHAR(20) NOT NULL DEFAULT 'numeric',
    collection_method TEXT,
    fasting_hours INTEGER,
    specimen_type VARCHAR(100),
    description TEXT,
    clinical_indications TEXT,
    display_order INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_parent_test FOREIGN KEY (parent_test_id)
        REFERENCES lab_test_definitions(id) ON DELETE SET NULL,

    -- Constraints
    CONSTRAINT chk_result_type CHECK (result_type IN ('numeric', 'text', 'boolean', 'categorical')),
    CONSTRAINT chk_category CHECK (category IN (
        'hematology', 'biochemistry', 'hormones', 'immunology',
        'microbiology', 'urine', 'imaging', 'functional', 'genetics', 'other'
    ))
);

-- Indexes for lab_test_definitions
CREATE INDEX IF NOT EXISTS idx_lab_test_def_code ON lab_test_definitions(code);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_tuss ON lab_test_definitions(tuss_code);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_loinc ON lab_test_definitions(loinc_code);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_category ON lab_test_definitions(category);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_requestable ON lab_test_definitions(is_requestable);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_parent ON lab_test_definitions(parent_test_id);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_active ON lab_test_definitions(is_active);
CREATE INDEX IF NOT EXISTS idx_lab_test_def_deleted ON lab_test_definitions(deleted_at);

-- ============================================================
-- 3. Create lab_test_score_mappings table
-- ============================================================

CREATE TABLE IF NOT EXISTS lab_test_score_mappings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lab_test_id UUID NOT NULL,
    score_item_id UUID NOT NULL,
    gender VARCHAR(10),
    min_age INTEGER,
    max_age INTEGER,
    notes TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_lab_test FOREIGN KEY (lab_test_id)
        REFERENCES lab_test_definitions(id) ON DELETE CASCADE,
    CONSTRAINT fk_score_item FOREIGN KEY (score_item_id)
        REFERENCES score_items(id) ON DELETE CASCADE,

    -- Constraints
    CONSTRAINT chk_mapping_gender CHECK (gender IN ('male', 'female') OR gender IS NULL)
);

-- Indexes for lab_test_score_mappings
CREATE INDEX IF NOT EXISTS idx_lab_test_score_mapping_test ON lab_test_score_mappings(lab_test_id);
CREATE INDEX IF NOT EXISTS idx_lab_test_score_mapping_item ON lab_test_score_mappings(score_item_id);
CREATE INDEX IF NOT EXISTS idx_lab_test_score_mapping_gender ON lab_test_score_mappings(gender);
CREATE INDEX IF NOT EXISTS idx_lab_test_score_mapping_active ON lab_test_score_mappings(is_active);
CREATE INDEX IF NOT EXISTS idx_lab_test_score_mapping_deleted ON lab_test_score_mappings(deleted_at);

-- ============================================================
-- 4. Create lab_test_reference_ranges table
-- ============================================================

CREATE TABLE IF NOT EXISTS lab_test_reference_ranges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lab_test_id UUID NOT NULL,
    gender VARCHAR(10),
    min_age INTEGER,
    max_age INTEGER,
    lower_limit DOUBLE PRECISION,
    upper_limit DOUBLE PRECISION,
    text_reference VARCHAR(200),
    unit VARCHAR(50),
    description TEXT,
    source VARCHAR(200),
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_ref_range_lab_test FOREIGN KEY (lab_test_id)
        REFERENCES lab_test_definitions(id) ON DELETE CASCADE,

    -- Constraints
    CONSTRAINT chk_ref_range_gender CHECK (gender IN ('male', 'female') OR gender IS NULL)
);

-- Indexes for lab_test_reference_ranges
CREATE INDEX IF NOT EXISTS idx_lab_test_ref_range_test ON lab_test_reference_ranges(lab_test_id);
CREATE INDEX IF NOT EXISTS idx_lab_test_ref_range_gender ON lab_test_reference_ranges(gender);
CREATE INDEX IF NOT EXISTS idx_lab_test_ref_range_age ON lab_test_reference_ranges(min_age, max_age);
CREATE INDEX IF NOT EXISTS idx_lab_test_ref_range_active ON lab_test_reference_ranges(is_active);
CREATE INDEX IF NOT EXISTS idx_lab_test_ref_range_deleted ON lab_test_reference_ranges(deleted_at);

-- ============================================================
-- 5. Create lab_result_values table
-- ============================================================

CREATE TABLE IF NOT EXISTS lab_result_values (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lab_result_id UUID NOT NULL,
    lab_test_definition_id UUID NOT NULL,
    numeric_value DOUBLE PRECISION,
    text_value TEXT,
    boolean_value BOOLEAN,
    unit VARCHAR(50),
    reference_range VARCHAR(200),
    is_abnormal BOOLEAN NOT NULL DEFAULT false,
    is_critical BOOLEAN NOT NULL DEFAULT false,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    -- Foreign keys
    CONSTRAINT fk_value_lab_result FOREIGN KEY (lab_result_id)
        REFERENCES lab_results(id) ON DELETE CASCADE,
    CONSTRAINT fk_value_lab_test_def FOREIGN KEY (lab_test_definition_id)
        REFERENCES lab_test_definitions(id) ON DELETE RESTRICT
);

-- Indexes for lab_result_values
CREATE INDEX IF NOT EXISTS idx_lab_result_value_result ON lab_result_values(lab_result_id);
CREATE INDEX IF NOT EXISTS idx_lab_result_value_test ON lab_result_values(lab_test_definition_id);
CREATE INDEX IF NOT EXISTS idx_lab_result_value_abnormal ON lab_result_values(is_abnormal);
CREATE INDEX IF NOT EXISTS idx_lab_result_value_critical ON lab_result_values(is_critical);
CREATE INDEX IF NOT EXISTS idx_lab_result_value_created ON lab_result_values(created_at);
CREATE INDEX IF NOT EXISTS idx_lab_result_value_deleted ON lab_result_values(deleted_at);

-- ============================================================
-- 6. Update triggers for updated_at
-- ============================================================

-- Trigger for lab_test_definitions
CREATE OR REPLACE FUNCTION update_lab_test_definitions_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_lab_test_definitions_updated_at ON lab_test_definitions;
CREATE TRIGGER trg_lab_test_definitions_updated_at
    BEFORE UPDATE ON lab_test_definitions
    FOR EACH ROW
    EXECUTE FUNCTION update_lab_test_definitions_updated_at();

-- Trigger for lab_test_score_mappings
CREATE OR REPLACE FUNCTION update_lab_test_score_mappings_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_lab_test_score_mappings_updated_at ON lab_test_score_mappings;
CREATE TRIGGER trg_lab_test_score_mappings_updated_at
    BEFORE UPDATE ON lab_test_score_mappings
    FOR EACH ROW
    EXECUTE FUNCTION update_lab_test_score_mappings_updated_at();

-- Trigger for lab_test_reference_ranges
CREATE OR REPLACE FUNCTION update_lab_test_reference_ranges_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_lab_test_reference_ranges_updated_at ON lab_test_reference_ranges;
CREATE TRIGGER trg_lab_test_reference_ranges_updated_at
    BEFORE UPDATE ON lab_test_reference_ranges
    FOR EACH ROW
    EXECUTE FUNCTION update_lab_test_reference_ranges_updated_at();

-- Trigger for lab_result_values
CREATE OR REPLACE FUNCTION update_lab_result_values_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_lab_result_values_updated_at ON lab_result_values;
CREATE TRIGGER trg_lab_result_values_updated_at
    BEFORE UPDATE ON lab_result_values
    FOR EACH ROW
    EXECUTE FUNCTION update_lab_result_values_updated_at();

-- ============================================================
-- 7. Comments for documentation
-- ============================================================

COMMENT ON TABLE lab_test_definitions IS 'Definições estruturadas de exames laboratoriais e seus parâmetros';
COMMENT ON TABLE lab_test_score_mappings IS 'Mapeamento entre exames laboratoriais e itens do escore Plenya';
COMMENT ON TABLE lab_test_reference_ranges IS 'Faixas de referência para exames laboratoriais por gênero e idade';
COMMENT ON TABLE lab_result_values IS 'Valores estruturados de resultados de exames laboratoriais';

COMMENT ON COLUMN lab_test_definitions.is_requestable IS 'Indica se o exame pode ser solicitado individualmente (true) ou só aparece no resultado (false)';
COMMENT ON COLUMN lab_test_definitions.parent_test_id IS 'ID do exame pai para criar hierarquia (ex: Hemoglobina tem parent=Hemograma)';
COMMENT ON COLUMN lab_test_score_mappings.gender IS 'Gênero específico para o mapeamento (null = ambos)';
COMMENT ON COLUMN lab_result_values.is_abnormal IS 'Indica se o valor está fora da faixa de referência';
COMMENT ON COLUMN lab_result_values.is_critical IS 'Indica se o valor é crítico e requer atenção imediata';
