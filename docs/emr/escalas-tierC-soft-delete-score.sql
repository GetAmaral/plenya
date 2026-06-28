-- ============================================================================
-- Tier C (FSS, FSFI, PSQI) — remover dos templates E soft-delete do escore
-- ============================================================================
-- Decisão (Getúlio 2026-06-22): escalas Tier C custom não entram no preenchimento.
-- Supera o antigo `escalas-tierC-remover-templates.sql` (que só fazia hard-DELETE
-- nos template_items): aqui SOFT-DELETA tanto as referências em templates quanto os
-- próprios score_items. Soft-deletar o score_item torna o SEED re-safe — o seed só
-- monta a partir de itens `deleted_at IS NULL`, então um re-seed não os traz de volta.
-- Idempotente (guards `deleted_at IS NULL`). FSS_9/PSQI_21 já estavam soft-deletados
-- desde fev/2026; na prática só o FSFI_36 ainda estava ativo + em template.
-- ============================================================================
BEGIN;

-- 1) soft-delete das referências em templates de anamnese
UPDATE anamnesis_template_items ati SET deleted_at = now()
FROM score_items si
WHERE ati.score_item_id = si.id
  AND si.anamnese_item_code IN ('FSS_9','FSFI_36','PSQI_21')
  AND ati.deleted_at IS NULL;

-- 2) soft-delete dos score_items
UPDATE score_items SET deleted_at = now()
WHERE anamnese_item_code IN ('FSS_9','FSFI_36','PSQI_21')
  AND deleted_at IS NULL;

-- conferência
SELECT anamnese_item_code,
       deleted_at IS NOT NULL AS item_soft_deletado,
       (SELECT COUNT(*) FROM anamnesis_template_items ti
        WHERE ti.score_item_id = score_items.id AND ti.deleted_at IS NULL) AS refs_ativas_em_templates
FROM score_items
WHERE anamnese_item_code IN ('FSS_9','FSFI_36','PSQI_21')
ORDER BY anamnese_item_code;

COMMIT;
