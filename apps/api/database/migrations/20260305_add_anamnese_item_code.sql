-- Migration: Add anamnese_item_code to score_items
-- Version: 20260305
-- Description: Adiciona campo anamnese_item_code para hardcoding de ScoreItems de anamnese
--              Complemento ao lab_test_code: items sem lab_test_code são itens de anamnese

ALTER TABLE score_items
ADD COLUMN anamnese_item_code VARCHAR(200);

CREATE INDEX idx_score_items_anamnese_item_code
ON score_items(anamnese_item_code)
WHERE anamnese_item_code IS NOT NULL;
