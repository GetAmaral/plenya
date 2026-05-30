-- Migration: Cleanup Lab Test System
-- Remove LabTestReferenceRange e campos desnecessários de LabResultValue
-- Created: 2026-01-25

-- ============================================================
-- 1. Remover colunas desnecessárias de lab_result_values
-- ============================================================

ALTER TABLE lab_result_values
DROP COLUMN IF EXISTS reference_range,
DROP COLUMN IF EXISTS is_abnormal,
DROP COLUMN IF EXISTS is_critical;

-- ============================================================
-- 2. Dropar tabela lab_test_reference_ranges
-- ============================================================

-- Dropar trigger primeiro
DROP TRIGGER IF EXISTS trg_lab_test_reference_ranges_updated_at ON lab_test_reference_ranges;

-- Dropar função do trigger
DROP FUNCTION IF EXISTS update_lab_test_reference_ranges_updated_at();

-- Dropar tabela
DROP TABLE IF EXISTS lab_test_reference_ranges;

-- ============================================================
-- 3. Comentários atualizados
-- ============================================================

COMMENT ON TABLE lab_result_values IS 'Valores estruturados de resultados de exames laboratoriais. Faixas de referência e classificação (anormal/crítico) são determinadas pelo sistema de Score.';
COMMENT ON COLUMN lab_result_values.unit IS 'Unidade de medida do valor (pode sobrescrever a unidade padrão do exame)';
