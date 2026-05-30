-- Migration: Add gender and age range fields to score_items
-- Date: 2026-02-08
-- Reason: Allow score items to be specific to gender and/or age ranges

-- Add gender column (optional - not_applicable, male, female)
ALTER TABLE score_items ADD COLUMN gender VARCHAR(20) DEFAULT 'not_applicable' CHECK (gender IN ('not_applicable', 'male', 'female'));

-- Add age range columns (optional)
ALTER TABLE score_items ADD COLUMN age_range_min INTEGER CHECK (age_range_min >= 0 AND age_range_min <= 150);
ALTER TABLE score_items ADD COLUMN age_range_max INTEGER CHECK (age_range_max >= 0 AND age_range_max <= 150);

-- Add comment to explain the fields
COMMENT ON COLUMN score_items.gender IS 'Gênero aplicável: not_applicable (padrão), male, female';
COMMENT ON COLUMN score_items.age_range_min IS 'Idade mínima aplicável em anos (0-150)';
COMMENT ON COLUMN score_items.age_range_max IS 'Idade máxima aplicável em anos (0-150)';
