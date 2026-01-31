-- Make points column optional in score_items table
-- Points field is no longer mandatory as some score items may not have point values

ALTER TABLE score_items
ALTER COLUMN points DROP NOT NULL,
ALTER COLUMN points DROP DEFAULT;

-- Add comment
COMMENT ON COLUMN score_items.points IS 'Pontos máximos atribuídos a este item (opcional). Valores entre 0 e 100';
