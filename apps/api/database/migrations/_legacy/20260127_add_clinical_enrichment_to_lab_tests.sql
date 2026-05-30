-- Migration: Add clinical enrichment fields to lab_test_definitions
-- Created: 2026-01-27
-- Description: Adiciona campos para enriquecimento clínico dos exames laboratoriais

-- Add clinical enrichment columns
ALTER TABLE lab_test_definitions
ADD COLUMN IF NOT EXISTS clinical_significance TEXT,
ADD COLUMN IF NOT EXISTS longevity_context TEXT,
ADD COLUMN IF NOT EXISTS clinical_recommendations TEXT;

-- Create indexes for full-text search (opcional, mas útil para busca futura)
CREATE INDEX IF NOT EXISTS idx_lab_test_def_clinical_sig ON lab_test_definitions USING gin(to_tsvector('portuguese', COALESCE(clinical_significance, '')));
CREATE INDEX IF NOT EXISTS idx_lab_test_def_longevity ON lab_test_definitions USING gin(to_tsvector('portuguese', COALESCE(longevity_context, '')));
CREATE INDEX IF NOT EXISTS idx_lab_test_def_clinical_rec ON lab_test_definitions USING gin(to_tsvector('portuguese', COALESCE(clinical_recommendations, '')));

-- Add comment to document purpose
COMMENT ON COLUMN lab_test_definitions.clinical_significance IS 'Significância clínica detalhada (200-400 palavras): mecanismos fisiológicos, aplicações clínicas, interpretação de valores alterados';
COMMENT ON COLUMN lab_test_definitions.longevity_context IS 'Contexto de longevidade (100-200 palavras): relação com envelhecimento saudável, marcadores de longevidade, implicações preventivas';
COMMENT ON COLUMN lab_test_definitions.clinical_recommendations IS 'Recomendações clínicas (150-300 palavras): quando solicitar, interpretação de resultados, fatores que afetam valores, intervenções baseadas em resultados';
