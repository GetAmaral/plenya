-- Migration: Add selected_level to anamnesis_items
-- Version: 20260305
-- Description: Separa numericValue (medida real) de selectedLevel (nível classificado 0-6)
--              Para dados existentes: numeric_value continha o nível (0-6), migrar para selected_level

ALTER TABLE anamnesis_items
ADD COLUMN selected_level INTEGER;

-- Migrar dados existentes: numeric_value atual = nível (0-6), mover para selected_level
UPDATE anamnesis_items
SET selected_level = ROUND(numeric_value)::int,
    numeric_value = NULL
WHERE numeric_value IS NOT NULL
  AND deleted_at IS NULL;
