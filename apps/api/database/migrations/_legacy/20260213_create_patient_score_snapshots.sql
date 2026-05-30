-- Migration: Create Patient Score Snapshots System
-- Generated from Go models in apps/api/internal/models/patient_score_*.go
-- Date: 2026-02-13

-- Table: patient_score_snapshots
CREATE TABLE IF NOT EXISTS patient_score_snapshots (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    patient_id UUID NOT NULL,
    calculated_by_user_id UUID NOT NULL,
    calculated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    total_actual_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_possible_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    total_score_percentage DOUBLE PRECISION NOT NULL DEFAULT 0,
    items_evaluated_count INTEGER NOT NULL DEFAULT 0,
    items_not_evaluated_count INTEGER NOT NULL DEFAULT 0,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_patient_score_snapshots_patient FOREIGN KEY (patient_id) REFERENCES patients(id) ON DELETE CASCADE,
    CONSTRAINT fk_patient_score_snapshots_user FOREIGN KEY (calculated_by_user_id) REFERENCES users(id) ON DELETE RESTRICT,
    CONSTRAINT chk_snapshots_total_actual_points CHECK (total_actual_points >= 0),
    CONSTRAINT chk_snapshots_total_possible_points CHECK (total_possible_points >= 0),
    CONSTRAINT chk_snapshots_total_score_percentage CHECK (total_score_percentage >= 0 AND total_score_percentage <= 100),
    CONSTRAINT chk_snapshots_items_evaluated CHECK (items_evaluated_count >= 0),
    CONSTRAINT chk_snapshots_items_not_evaluated CHECK (items_not_evaluated_count >= 0)
);

-- Indexes for patient_score_snapshots
CREATE INDEX IF NOT EXISTS idx_snapshot_patient ON patient_score_snapshots(patient_id);
CREATE INDEX IF NOT EXISTS idx_snapshot_user ON patient_score_snapshots(calculated_by_user_id);
CREATE INDEX IF NOT EXISTS idx_snapshot_date ON patient_score_snapshots(calculated_at DESC);
CREATE INDEX IF NOT EXISTS idx_patient_score_snapshots_deleted_at ON patient_score_snapshots(deleted_at);

-- Table: patient_score_group_results
CREATE TABLE IF NOT EXISTS patient_score_group_results (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    snapshot_id UUID NOT NULL,
    group_id UUID NOT NULL,
    actual_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    possible_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    score_percentage DOUBLE PRECISION NOT NULL DEFAULT 0,
    items_evaluated_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_patient_score_group_results_snapshot FOREIGN KEY (snapshot_id) REFERENCES patient_score_snapshots(id) ON DELETE CASCADE,
    CONSTRAINT fk_patient_score_group_results_group FOREIGN KEY (group_id) REFERENCES score_groups(id) ON DELETE RESTRICT,
    CONSTRAINT chk_group_results_actual_points CHECK (actual_points >= 0),
    CONSTRAINT chk_group_results_possible_points CHECK (possible_points >= 0),
    CONSTRAINT chk_group_results_score_percentage CHECK (score_percentage >= 0 AND score_percentage <= 100),
    CONSTRAINT chk_group_results_items_evaluated CHECK (items_evaluated_count >= 0)
);

-- Indexes for patient_score_group_results
CREATE INDEX IF NOT EXISTS idx_group_result_snapshot ON patient_score_group_results(snapshot_id);
CREATE INDEX IF NOT EXISTS idx_group_result_group ON patient_score_group_results(group_id);
CREATE INDEX IF NOT EXISTS idx_patient_score_group_results_deleted_at ON patient_score_group_results(deleted_at);

-- Table: patient_score_item_results
CREATE TABLE IF NOT EXISTS patient_score_item_results (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    snapshot_id UUID NOT NULL,
    item_id UUID NOT NULL,
    group_id UUID NOT NULL,
    status VARCHAR(30) NOT NULL,
    data_source VARCHAR(20),
    lab_result_id UUID,
    anamnesis_item_id UUID,
    value_used DOUBLE PRECISION,
    level_matched_id UUID,
    level_number INTEGER,
    max_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    actual_points DOUBLE PRECISION NOT NULL DEFAULT 0,
    not_evaluated_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_patient_score_item_results_snapshot FOREIGN KEY (snapshot_id) REFERENCES patient_score_snapshots(id) ON DELETE CASCADE,
    CONSTRAINT fk_patient_score_item_results_item FOREIGN KEY (item_id) REFERENCES score_items(id) ON DELETE RESTRICT,
    CONSTRAINT fk_patient_score_item_results_group FOREIGN KEY (group_id) REFERENCES score_groups(id) ON DELETE RESTRICT,
    CONSTRAINT fk_patient_score_item_results_level FOREIGN KEY (level_matched_id) REFERENCES score_levels(id) ON DELETE SET NULL,
    CONSTRAINT fk_patient_score_item_results_lab_result FOREIGN KEY (lab_result_id) REFERENCES lab_results(id) ON DELETE SET NULL,
    CONSTRAINT fk_patient_score_item_results_anamnesis_item FOREIGN KEY (anamnesis_item_id) REFERENCES anamnesis_items(id) ON DELETE SET NULL,
    CONSTRAINT chk_item_results_status CHECK (status IN ('evaluated', 'not_applicable', 'no_data_available')),
    CONSTRAINT chk_item_results_data_source CHECK (data_source IN ('lab_result', 'anamnesis_item') OR data_source IS NULL),
    CONSTRAINT chk_item_results_level_number CHECK (level_number >= 0 AND level_number <= 6 OR level_number IS NULL),
    CONSTRAINT chk_item_results_max_points CHECK (max_points >= 0 AND max_points <= 100),
    CONSTRAINT chk_item_results_actual_points CHECK (actual_points >= 0 AND actual_points <= 100)
);

-- Indexes for patient_score_item_results
CREATE INDEX IF NOT EXISTS idx_item_result_snapshot ON patient_score_item_results(snapshot_id);
CREATE INDEX IF NOT EXISTS idx_item_result_item ON patient_score_item_results(item_id);
CREATE INDEX IF NOT EXISTS idx_item_result_group ON patient_score_item_results(group_id);
CREATE INDEX IF NOT EXISTS idx_item_result_level ON patient_score_item_results(level_matched_id);
CREATE INDEX IF NOT EXISTS idx_item_result_lab_result ON patient_score_item_results(lab_result_id);
CREATE INDEX IF NOT EXISTS idx_item_result_anamnesis_item ON patient_score_item_results(anamnesis_item_id);
CREATE INDEX IF NOT EXISTS idx_patient_score_item_results_deleted_at ON patient_score_item_results(deleted_at);

-- Comments
COMMENT ON TABLE patient_score_snapshots IS 'Snapshots completos de escores de saúde do paciente em pontos específicos no tempo';
COMMENT ON TABLE patient_score_group_results IS 'Resultados agregados por grupo clínico dentro de um snapshot';
COMMENT ON TABLE patient_score_item_results IS 'Avaliação individual de cada item de score para o paciente';

COMMENT ON COLUMN patient_score_snapshots.calculated_at IS 'Data e hora do cálculo do snapshot';
COMMENT ON COLUMN patient_score_snapshots.total_score_percentage IS 'Score final em percentual (0-100)';
COMMENT ON COLUMN patient_score_snapshots.items_evaluated_count IS 'Quantidade de itens que foram avaliados';
COMMENT ON COLUMN patient_score_snapshots.items_not_evaluated_count IS 'Quantidade de itens que não foram avaliados';

COMMENT ON COLUMN patient_score_item_results.status IS 'Status da avaliação: evaluated, not_applicable, no_data_available';
COMMENT ON COLUMN patient_score_item_results.data_source IS 'Fonte dos dados: lab_result ou anamnesis_item';
COMMENT ON COLUMN patient_score_item_results.level_number IS 'Número do nível atingido (0-6), denormalizado';
COMMENT ON COLUMN patient_score_item_results.actual_points IS 'Pontos reais obtidos (proporcional ao nível)';
