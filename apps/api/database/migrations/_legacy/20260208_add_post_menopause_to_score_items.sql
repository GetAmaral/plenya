-- Migration: add_post_menopause_to_score_items
-- Created at: 2026-02-08
-- Description: Adiciona campo post_menopause à tabela score_items

-- Add post_menopause column
ALTER TABLE score_items ADD COLUMN post_menopause BOOLEAN DEFAULT NULL;

-- Add comment
COMMENT ON COLUMN score_items.post_menopause IS 'Indica se o score_item é aplicável apenas para mulheres pós-menopausa';
