-- ============================================================================
-- Escore — item "Tratamentos em uso para modificar composição corporal"
-- Ajuste do LEVEL 5 (melhor nível): name + site_legend
-- ============================================================================
-- L5 deixa de ser só "não utilizando nenhum" e passa a contemplar também o
-- acompanhamento adequado com nutricionista/educador físico.
-- Identifica o nível por anamnese_item_code + level (UUIDs de score_levels
-- diferem entre dev e prod). Idempotente.
-- ============================================================================
UPDATE score_levels sl
SET name        = 'Nenhum ou Acompanhamento nutri/educador físico adequado',
    site_legend = 'Nenhum tratamento em uso para mudar a composição do corpo, ou acompanhamento adequado com nutricionista e educador físico.',
    updated_at  = now()
FROM score_items si
WHERE sl.item_id = si.id
  AND si.anamnese_item_code = 'TRATAMENTOS_EM_USO_PARA_MODIFICAR_COMPOSICAO_CORPORAL'
  AND sl.level = 5
  AND sl.deleted_at IS NULL;

SELECT sl.level, sl.name, sl.site_legend
FROM score_items si JOIN score_levels sl ON sl.item_id = si.id
WHERE si.anamnese_item_code = 'TRATAMENTOS_EM_USO_PARA_MODIFICAR_COMPOSICAO_CORPORAL'
  AND sl.level = 5 AND sl.deleted_at IS NULL;
