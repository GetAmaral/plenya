-- ============================================================================
-- Correção do denominador do nome da ASEX no escore: "/25" -> "/30"
-- ============================================================================
-- A ASEX (Arizona Sexual Experience Scale) tem 5 itens × 1–6 → total 5–30
-- (maxScore=30 no SCALE_REGISTRY; cutoff de disfunção ≥19). O nome do score item
-- trazia "____/25" por engano. Idempotente (só toca se ainda estiver "/25").
-- Ref.: McGahuey et al. 2000 (J Sex Marital Ther); range 5–30.
-- ============================================================================
UPDATE score_items
SET name = replace(name, '/25', '/30'), updated_at = now()
WHERE anamnese_item_code = 'ASEX_25' AND name LIKE '%/25%';

SELECT name, anamnese_item_code FROM score_items WHERE anamnese_item_code = 'ASEX_25';
