-- Migration: Add clinical enrichment fields to score_items
-- Created: 2026-01-27
-- Description: Adiciona campos detalhados de enriquecimento clínico aos items de escore

-- Add comprehensive clinical enrichment columns to score_items
ALTER TABLE score_items
ADD COLUMN IF NOT EXISTS clinical_explanation TEXT,
ADD COLUMN IF NOT EXISTS low_explanation TEXT,
ADD COLUMN IF NOT EXISTS high_explanation TEXT,
ADD COLUMN IF NOT EXISTS clinical_significance TEXT,
ADD COLUMN IF NOT EXISTS interpretation_guide TEXT,
ADD COLUMN IF NOT EXISTS recommendations TEXT[],
ADD COLUMN IF NOT EXISTS related_conditions TEXT[],
ADD COLUMN IF NOT EXISTS patient_friendly_explanation TEXT;

-- Create indexes for full-text search
CREATE INDEX IF NOT EXISTS idx_score_items_clinical_exp ON score_items USING gin(to_tsvector('portuguese', COALESCE(clinical_explanation, '')));
CREATE INDEX IF NOT EXISTS idx_score_items_patient_exp ON score_items USING gin(to_tsvector('portuguese', COALESCE(patient_friendly_explanation, '')));

-- Add comments to document the purpose of each field
COMMENT ON COLUMN score_items.clinical_explanation IS 'Explicação científica completa do parâmetro clínico: mecanismos bioquímicos/fisiológicos, importância clínica (3-5 parágrafos técnicos para profissionais)';
COMMENT ON COLUMN score_items.low_explanation IS 'Explicação detalhada do significado de valores baixos/diminuídos: mecanismos fisiopatológicos, causas, implicações clínicas (2-3 parágrafos técnicos)';
COMMENT ON COLUMN score_items.high_explanation IS 'Explicação detalhada do significado de valores altos/elevados: mecanismos fisiopatológicos, causas, implicações clínicas (2-3 parágrafos técnicos)';
COMMENT ON COLUMN score_items.clinical_significance IS 'Significância clínica profunda: como interpretar resultados, correlações com outros exames, importância diagnóstica, seguimento clínico (2-3 parágrafos aplicados)';
COMMENT ON COLUMN score_items.interpretation_guide IS 'Guia prático de interpretação para médicos: valores de referência, variações fisiológicas, quando solicitar, como interpretar em diferentes contextos (2-3 parágrafos práticos)';
COMMENT ON COLUMN score_items.recommendations IS 'Array de recomendações clínicas específicas e acionáveis baseadas em evidências (5-8 recomendações práticas)';
COMMENT ON COLUMN score_items.related_conditions IS 'Array de condições clínicas relacionadas ao parâmetro (5-10 condições relevantes)';
COMMENT ON COLUMN score_items.patient_friendly_explanation IS 'Explicação acessível para pacientes em linguagem clara: o que é o exame, importância, o que resultados indicam (2-3 parágrafos informativos mas compreensíveis)';
