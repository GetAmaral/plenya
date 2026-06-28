-- ============================================================================
-- MIGRAÇÃO PROD — Templates de anamnese (bundle consolidado, 2026-06-27)
-- ============================================================================
-- Roda TUDO de uma vez, na ordem: (1) rename → (2) Tier C soft-delete →
-- (3) recompose (moves item-a-item + renumeração). Todos idempotentes:
-- re-rodar é seguro; o que já estiver aplicado vira no-op.
--
-- Como aplicar em prod (via VPS):
--   cat docs/emr/prod-migration-anamnese.sql | ssh plenya \
--     "sudo docker exec -i mb511beqjtgd7nsjlnngh3m6 psql -U plenya_user -d plenya_db"
--
-- ⚠️ GERADO por concatenação de:
--   docs/emr/rename-templates-anamnese.sql
--   docs/emr/escalas-tierC-soft-delete-score.sql
--   docs/emr/recompose-templates-anamnese.sql
-- Editar os fontes e RE-GERAR, nunca editar este arquivo à mão.
-- ============================================================================

-- ████ PARTE 1/3 — RENAME ████████████████████████████████████████████████████
-- ============================================================================
-- Renomeação dos 13 templates de anamnese — padrão "[Continuum |] Prof | Momento"
-- ============================================================================
-- Aplica só os novos nomes (keyed por UUID fixo) SEM re-rodar o seed inteiro
-- (re-rodar o seed reverteria a curadoria manual de itens). Idempotente.
-- Padrão: labels curtas (Médico/Nutri/Psico/Ed. Física), separador " | ".
-- ============================================================================
BEGIN;
UPDATE anamnesis_templates SET name='Médico | Inicial',                            updated_at=now() WHERE id='11111111-1111-7111-8111-111111111101';
UPDATE anamnesis_templates SET name='Médico | Acompanhamento',                     updated_at=now() WHERE id='11111111-1111-7111-8111-111111111102';
UPDATE anamnesis_templates SET name='Médico | Revisão de Exames',                  updated_at=now() WHERE id='11111111-1111-7111-8111-111111111103';
UPDATE anamnesis_templates SET name='Continuum | Médico | Inicial',               updated_at=now() WHERE id='11111111-1111-7111-8111-111111111104';
UPDATE anamnesis_templates SET name='Continuum | Médico | Complemento',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111105';
UPDATE anamnesis_templates SET name='Continuum | Médico | Acompanhamento',        updated_at=now() WHERE id='11111111-1111-7111-8111-111111111106';
UPDATE anamnesis_templates SET name='Continuum | Médico | Reavaliação Trimestral',updated_at=now() WHERE id='11111111-1111-7111-8111-111111111107';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111108';
UPDATE anamnesis_templates SET name='Continuum | Nutri | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111109';
UPDATE anamnesis_templates SET name='Continuum | Psico | Inicial',                updated_at=now() WHERE id='11111111-1111-7111-8111-111111111110';
UPDATE anamnesis_templates SET name='Continuum | Psico | Acompanhamento',         updated_at=now() WHERE id='11111111-1111-7111-8111-111111111111';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Inicial',           updated_at=now() WHERE id='11111111-1111-7111-8111-111111111112';
UPDATE anamnesis_templates SET name='Continuum | Ed. Física | Acompanhamento',    updated_at=now() WHERE id='11111111-1111-7111-8111-111111111113';
SELECT name, area FROM anamnesis_templates WHERE deleted_at IS NULL ORDER BY area, name;
COMMIT;

-- ████ PARTE 2/3 — TIER C SOFT-DELETE ███████████████████████████████████████
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

-- ████ PARTE 3/3 — RECOMPOSE (moves item-a-item) ████████████████████████████
-- ============================================================================
-- Recomposição (curadoria) dos templates de anamnese — DELTA pós-seed
-- ============================================================================
-- Move itens entre templates DEPOIS do seed (docs/emr/seed-anamnese-templates.sql).
-- Por que um arquivo à parte: o seed monta a composição por lógica de conjuntos
-- e re-rodá-lo REVERTERIA estas curadorias manuais. Este delta é a fonte das
-- mudanças item-a-item feitas à mão; rodar APÓS o seed (ou sozinho em prod).
--
-- Portabilidade dev↔prod: itens são identificados por NOME (grupo/subgrupo/item),
-- nunca por UUID de score_item — os UUIDs de score_items diferem entre ambientes.
-- Templates usam UUID fixo (idênticos em todo lugar), então são referenciados por id.
--
-- Idempotente: re-rodar não duplica (remove só ativos; insere só se faltante;
-- renumera por ordem natural). Rodar dentro de transação; trocar ROLLBACK->COMMIT
-- após conferir os counts.
-- ============================================================================
BEGIN;

-- Especificação dos moves (item=NULL => todos os itens do subgrupo) -----------
CREATE TEMP TABLE moves(src uuid, dst uuid, grupo text, subgrupo text, item text) ON COMMIT DROP;
INSERT INTO moves (src, dst, grupo, subgrupo, item) VALUES
 -- 2026-06-27 · Continuum | Médico | Inicial (…104) -> Complemento (…105)
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
  'Objetivos','Percepção de futuro (6m-5a-10a-30a)', NULL),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105',
  'Vida Sexual','Atual','Uso recente outros medicamentos/suplementos para libido/desempenho sexual'),
 -- 2026-06-27 · Continuum | Nutri | Inicial (…108) -> Médico | Inicial (…104)
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Frutas'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Verduras e Legumes'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Alimentação','Atual (últmos 6 meses)','Consumo de Proteínas (cálculos com base no recordatório)'),
 ('11111111-1111-7111-8111-111111111108','11111111-1111-7111-8111-111111111104',
  'Composição corporal','Atual','Tratamentos em uso para modificar composição corporal'),
 -- 2026-06-27 · Continuum | Médico | Complemento (…105) -> Médico | Inicial (…104)
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104',
  'Histórico de doenças','Cirurgias já realizadas', NULL),
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104',
  'Histórico de doenças','Hábitos e vícios nocivos (Questionar ativamente sobre uso passado ou atual):', NULL),
 -- 2026-06-27 · Continuum | Médico | Complemento (…105) -> Psico | Inicial (…110)
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111110',
  'Sono','Histórico', NULL),
 -- 2026-06-27 · Continuum | Médico | Inicial (…104) -> Psico | Inicial (…110)
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111110',
  'Cognição','Atual','Escala PHQ-9 (humor): ___/27'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111110',
  'Cognição','Atual','GAD-7 (ansiedade): ___/21'),
 -- 2026-06-27 · Continuum | Psico | Inicial (…110) -> Médico | Inicial (…104)
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','5 palavras de Dubois - imediato: ____/5'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','5 palavras de Dubois - tardio: ____/5'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','Span de dígitos - Direto:___/8'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','Span de dígitos - Inverso:___/7'),
 ('11111111-1111-7111-8111-111111111110','11111111-1111-7111-8111-111111111104',
  'Cognição','Atual','Uso atual de psicotrópicos para cognição'),
 -- 2026-06-27 · Continuum | Ed. Física | Inicial (…112) -> Médico | Inicial (…104)
 ('11111111-1111-7111-8111-111111111112','11111111-1111-7111-8111-111111111104',
  'Movimento e atividade física','Histórico','Lesões relacionadas ao exercício'),
 ('11111111-1111-7111-8111-111111111112','11111111-1111-7111-8111-111111111104',
  'Movimento e atividade física','Histórico','Cirurgias realizadas relacionadas ao exercício');

-- Resolve nomes -> score_item_id (no ambiente atual) -------------------------
CREATE TEMP TABLE resolved ON COMMIT DROP AS
SELECT m.src, m.dst, si.id AS score_item_id
FROM moves m
JOIN score_groups    g  ON g.name  = m.grupo    AND g.deleted_at  IS NULL
JOIN score_subgroups sg ON sg.group_id = g.id   AND sg.name = m.subgrupo AND sg.deleted_at IS NULL
JOIN score_items     si ON si.subgroup_id = sg.id AND si.deleted_at IS NULL
                        AND (m.item IS NULL OR si.name = m.item);

-- 1) remove do template de origem -------------------------------------------
UPDATE anamnesis_template_items t SET deleted_at = now()
FROM resolved r
WHERE t.anamnesis_template_id = r.src AND t.score_item_id = r.score_item_id
  AND t.deleted_at IS NULL;

-- 2) adiciona no template de destino (se ainda não tiver ativo) --------------
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), r.dst, r.score_item_id, 0, now(), now()
FROM resolved r
WHERE NOT EXISTS (
  SELECT 1 FROM anamnesis_template_items x
  WHERE x.anamnesis_template_id = r.dst AND x.score_item_id = r.score_item_id
    AND x.deleted_at IS NULL);

-- 3) renumera "order" (ordem natural g/sg/i) nos templates tocados -----------
WITH touched AS (
  SELECT src AS tid FROM resolved UNION SELECT dst FROM resolved
), o2 AS (
  SELECT ti.id,
         ROW_NUMBER() OVER (PARTITION BY ti.anamnesis_template_id
                            ORDER BY g."order", sg."order", si."order") AS rn
  FROM anamnesis_template_items ti
  JOIN score_items     si ON si.id = ti.score_item_id
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups    g  ON g.id  = sg.group_id
  WHERE ti.anamnesis_template_id IN (SELECT tid FROM touched)
    AND ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = o2.rn, updated_at = now()
FROM o2 WHERE t.id = o2.id;

-- Conferência ----------------------------------------------------------------
SELECT 'resolvidos' AS check, COUNT(*) FROM resolved;
SELECT at.name, COUNT(*) FILTER (WHERE ti.deleted_at IS NULL) AS itens
FROM anamnesis_templates at
LEFT JOIN anamnesis_template_items ti ON ti.anamnesis_template_id = at.id
WHERE at.id IN (SELECT src FROM resolved UNION SELECT dst FROM resolved)
GROUP BY at.name ORDER BY at.name;

COMMIT;
-- ROLLBACK;  -- use enquanto confere; troque por COMMIT para aplicar
