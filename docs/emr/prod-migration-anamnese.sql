-- ============================================================================
-- MIGRAÇÃO PROD — Templates de anamnese + ajustes de escore (bundle, 2026-06-28)
-- ============================================================================
-- Ordem: (1) rename templates → (2) Tier C soft-delete → (3) recompose moves →
-- (4) ASEX /25→/30 → (5) Tratamentos composição L5 → (6) sufixos coerentes →
-- (7) Hipertensão arterial (pessoal). Todos idempotentes.
--
-- ⚠️ ACOPLAMENTO COM CÓDIGO: a Parte 6 renomeia DOENCA_CARDIOVASCULAR_2 → _FAMILIAR,
-- ATIVIDADE_FISICA → ATIVIDADE_FISICA_VIDA_ADULTA e EXERCICIO_FISICO →
-- EXERCICIO_FISICO_VIDA_ADULTA, que também são constantes em
-- apps/api/internal/models/clinical_codes.go. Aplicar o SQL e o deploy do api JUNTOS.
--
-- Aplicar:
--   cat docs/emr/prod-migration-anamnese.sql | ssh plenya \
--     "sudo docker exec -i mb511beqjtgd7nsjlnngh3m6 psql -U plenya_user -d plenya_db"
--
-- ⚠️ GERADO por concatenação (editar fontes e RE-GERAR):
--   rename-templates-anamnese.sql · escalas-tierC-soft-delete-score.sql ·
--   recompose-templates-anamnese.sql · escalas-asex-nome-fix.sql ·
--   escore-tratamentos-composicao-level5.sql · escore-sufixos-coerentes.sql ·
--   escore-hipertensao-arterial.sql
-- ============================================================================

-- ████ PARTE 1/7 — RENAME TEMPLATES ██████████████████████████████████████████
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

-- ████ PARTE 2/7 — TIER C SOFT-DELETE ███████████████████████████████████████
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

-- ████ PARTE 3/7 — RECOMPOSE (moves item-a-item) ████████████████████████████
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

-- Especificação dos moves. Duas formas de identificar o item:
--   • por nome: grupo+subgrupo (+item; item=NULL => subgrupo inteiro)
--   • por código: anamnese_item_code (robusto p/ itens c/ nome frágil, ex.: escalas com "____")
CREATE TEMP TABLE moves(src uuid, dst uuid, grupo text, subgrupo text, item text, code text) ON COMMIT DROP;
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

-- Moves por código (escalas) -------------------------------------------------
-- 2026-06-27 · inversão ASEX <-> IIEF-5 no par Continuum Médico Inicial<->Complemento
--   ASEX: Complemento (…105) -> Inicial (…104) · IIEF-5: Inicial (…104) -> Complemento (…105)
INSERT INTO moves (src, dst, code) VALUES
 ('11111111-1111-7111-8111-111111111105','11111111-1111-7111-8111-111111111104','ASEX_25'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105','IIEF_5_25');

-- Resolve item -> score_item_id (por nome OU por código, no ambiente atual) ---
CREATE TEMP TABLE resolved ON COMMIT DROP AS
  SELECT m.src, m.dst, si.id AS score_item_id          -- por nome
  FROM moves m
  JOIN score_groups    g  ON g.name  = m.grupo    AND g.deleted_at  IS NULL
  JOIN score_subgroups sg ON sg.group_id = g.id   AND sg.name = m.subgrupo AND sg.deleted_at IS NULL
  JOIN score_items     si ON si.subgroup_id = sg.id AND si.deleted_at IS NULL
                          AND (m.item IS NULL OR si.name = m.item)
  WHERE m.code IS NULL
UNION
  SELECT m.src, m.dst, si.id                            -- por código
  FROM moves m
  JOIN score_items si ON si.anamnese_item_code = m.code AND si.deleted_at IS NULL
  WHERE m.code IS NOT NULL;

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

-- ████ PARTE 4/7 — ASEX nome /25 -> /30 ██████████████████████████████████████
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

-- ████ PARTE 5/7 — ESCORE: Tratamentos composição, level 5 ███████████████████
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

-- ████ PARTE 6/7 — ESCORE: sufixos coerentes (anti _2) ███████████████████████
-- ============================================================================
-- Escore — substitui sufixos numéricos (_2/_3/...) por sufixos COERENTES
-- ============================================================================
-- Colisões de anamnese (mesmo nome em contextos diferentes) eram desambiguadas
-- por número (ridículo). Aqui cada uma vira um sufixo coerente com grupo/subgrupo/
-- pai/gênero. NÃO toca genética (códigos rsID, risco de matching) nem maxscore de
-- escala (ASEX_25, GAD_7_21...). ATIVIDADE_FISICA / EXERCICIO_FISICO viram variantes
-- por etapa de vida; a constante clínica (clinical_codes.go) passa a apontar _VIDA_ADULTA.
-- Idempotente: WHERE ancorado no padrão original; re-rodar não reprocessa.
-- ============================================================================
BEGIN;

-- A1) Doenças FAMILIARES → _FAMILIAR (CV, auto-imunes, virais, HAS) ----------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','') || '_FAMILIAR', updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico Familiar de Doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCA_CARDIOVASCULAR|DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS|HIPERTENSAO_ARTERIAL)(_[0-9]+)?$';

-- A2) Doenças PESSOAIS → código limpo (tira o _N que sobrou) ------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$',''), updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Histórico de doenças' AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(DOENCAS_AUTO_IMUNES|DOENCAS_VIRAIS_CRONICAS)_[0-9]+$';

-- A3) Álcool (hábitos/vícios) → _HABITOS (álcool da Alimentação fica limpo) ---
UPDATE score_items SET anamnese_item_code = 'ALCOOL_HABITOS', updated_at = now()
WHERE anamnese_item_code = 'ALCOOL_2' AND deleted_at IS NULL;

-- B) Timeline familiar por PAI (Mãe/Pai) -------------------------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Mãe' THEN 'MAE' WHEN 'Pai' THEN 'PAI' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Mãe','Pai')
  AND si.anamnese_item_code ~ '^(DURANTE_ADOLESCENCIA_DO_PACIENTE|DURANTE_INFANCIA_DO_PACIENTE|PRE_NATAL_PRE_CONCEPCAO)(_[0-9]+)?$';

-- C) Etapas de vida por GRUPO (Infância/Adolescência/Vida adulta/Pré-natal) ---
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE g.name WHEN 'Alimentação' THEN 'ALIMENTACAO' WHEN 'Composição corporal' THEN 'COMPOSICAO'
                          WHEN 'Movimento e atividade física' THEN 'MOVIMENTO' WHEN 'Sono' THEN 'SONO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND si.deleted_at IS NULL
  AND g.name IN ('Alimentação','Composição corporal','Movimento e atividade física','Sono')
  AND si.anamnese_item_code ~ '^(INFANCIA|ADOLESCENCIA|VIDA_ADULTA|PRE_NATAL)(_[0-9]+)?$';

-- D) Movimento histórico (atividade/exercício/ar livre/esporte) por etapa de vida (pai) -
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE p.name WHEN 'Infância' THEN 'INFANCIA' WHEN 'Adolescência' THEN 'ADOLESCENCIA' WHEN 'Vida adulta' THEN 'VIDA_ADULTA' END, updated_at = now()
FROM score_items p
WHERE si.parent_item_id = p.id AND si.deleted_at IS NULL AND p.name IN ('Infância','Adolescência','Vida adulta')
  AND si.anamnese_item_code ~ '^(ATIVIDADE_FISICA|EXERCICIO_FISICO|ATIVIDADES_AO_AR_LIVRE|ESPORTE_COMPETITIVO)(_[0-9]+)?$';

-- E) Testes práticos por GÊNERO + FAIXA ETÁRIA (colunas) ----------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE si.gender WHEN 'male' THEN 'HOMENS' WHEN 'female' THEN 'MULHERES' END
    || COALESCE('_' || si.age_range_min || '_' || si.age_range_max, ''), updated_at = now()
WHERE si.deleted_at IS NULL AND si.gender IN ('male','female')
  AND si.anamnese_item_code ~ '^(RESISTENCIA_NEUROMUSCULAR_ABDOMINAL|RESISTENCIA_NEUROMUSCULAR_FLEXAO_DE_SOLO|RESISTENCIA_NEUROMUSCULAR_PRANCHA|RESISTENCIA_CARDIORRESPIRATORIA_BURPEE)_[0-9]+$';

-- F) Social → _ATUAL / _HISTORICO (por subgrupo) -----------------------------
UPDATE score_items si SET anamnese_item_code = regexp_replace(si.anamnese_item_code,'_[0-9]+$','')
    || '_' || CASE sg.name WHEN 'Atual' THEN 'ATUAL' WHEN 'Histórico' THEN 'HISTORICO' END, updated_at = now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE si.subgroup_id = sg.id AND g.name = 'Social' AND sg.name IN ('Atual','Histórico') AND si.deleted_at IS NULL
  AND si.anamnese_item_code ~ '^(CONDICOES_DE_MORADIA|PROFISSOES|RELIGIOSIDADE|SITUACAO_CONJUGAL|SITUACAO_DE_PETS|SITUACAO_FAMILIAR|SITUACAO_FINANCEIRA_RENDA_ATIVA_E_PASSIVA)(_[0-9]+)?$';

-- Conferência: nenhum código duplicado fora da genética -----------------------
SELECT si.anamnese_item_code, COUNT(*) AS n
FROM score_items si JOIN score_subgroups sg ON sg.id=si.subgroup_id JOIN score_groups g ON g.id=sg.group_id
WHERE si.deleted_at IS NULL AND si.anamnese_item_code IS NOT NULL AND g.name <> 'Genética'
GROUP BY si.anamnese_item_code HAVING COUNT(*) > 1;

COMMIT;

-- ████ PARTE 7/7 — ESCORE: Hipertensão arterial (pessoal) ████████████████████
-- ============================================================================
-- Escore — adiciona "Hipertensão arterial" em Histórico de doenças › Doenças crônicas
-- ============================================================================
-- Lacuna: HAS não existia como doença crônica (só "Anti-hipertensivos" em Medicamentos).
-- Posicionada ANTES de "Insuficiência cardíaca" (order 7; IC é 8).
-- Níveis espelham o Diabetes (matriz duração × controle) — análogo crônico/controlável.
-- points=25 (igual Diabetes e Doença renal crônica). default_level5=true (L5 = "Não tenho").
-- Adicionada aos 5 templates que contêm os irmãos do subgrupo.
-- Idempotente: identifica por anamnese_item_code; só insere o que faltar; renumera ordem.
-- ============================================================================
BEGIN;

-- 1) score_item HAS (insere se não existir) ----------------------------------
INSERT INTO score_items (id, name, anamnese_item_code, subgroup_id, "order", points, default_level5, gender, created_at, updated_at)
SELECT uuid_generate_v7(), 'Hipertensão arterial', 'HIPERTENSAO_ARTERIAL', sg.id, 7, 25, true, 'not_applicable', now(), now()
FROM score_subgroups sg JOIN score_groups g ON g.id = sg.group_id
WHERE g.name = 'Histórico de doenças' AND sg.name LIKE 'Doenças crônicas%'
  AND NOT EXISTS (SELECT 1 FROM score_items x WHERE x.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND x.deleted_at IS NULL);

-- 2) 6 níveis (matriz duração × controle, espelhando Diabetes) ----------------
INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, v.level, '=', v.lname, v.legend, now(), now()
FROM score_items si
CROSS JOIN (VALUES
  (0, 'Tenho há mais de 5 anos, com controle inadequado', 'Hipertensão há mais de cinco anos e mal controlada, exige cuidado intensivo.'),
  (1, 'Tenho há menos de 5 anos, com controle inadequado', 'Hipertensão recente com a pressão ainda mal controlada, requer ajuste.'),
  (2, 'Tenho há mais de 5 anos, com controle adequado', 'Hipertensão de longa data, porém bem controlada.'),
  (3, 'Tenho há menos de 5 anos, com controle adequado', 'Hipertensão há menos de cinco anos e bem controlada.'),
  (4, 'Não tenho, mas tenho histórico familiar forte', 'Sem hipertensão, mas com histórico forte na família.'),
  (5, 'Não tenho', 'Não tem hipertensão, situação favorável.')
) AS v(level, lname, legend)
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM score_levels sl WHERE sl.item_id = si.id AND sl.level = v.level AND sl.deleted_at IS NULL);

-- 3) adiciona aos 5 templates que têm os irmãos (Doenças crônicas) ------------
INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), t.id, si.id, 0, now(), now()
FROM score_items si
JOIN anamnesis_templates t ON t.id IN (
  '11111111-1111-7111-8111-111111111101', -- Médico | Inicial
  '11111111-1111-7111-8111-111111111102', -- Médico | Acompanhamento
  '11111111-1111-7111-8111-111111111104', -- Continuum | Médico | Inicial
  '11111111-1111-7111-8111-111111111106', -- Continuum | Médico | Acompanhamento
  '11111111-1111-7111-8111-111111111107') -- Continuum | Médico | Reavaliação Trimestral
WHERE si.anamnese_item_code = 'HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL
  AND NOT EXISTS (
    SELECT 1 FROM anamnesis_template_items x
    WHERE x.anamnesis_template_id = t.id AND x.score_item_id = si.id AND x.deleted_at IS NULL);

-- 4) renumera "order" (natural g/sg/i) nos 5 templates tocados ----------------
WITH o AS (
  SELECT ti.id, ROW_NUMBER() OVER (PARTITION BY ti.anamnesis_template_id
                                   ORDER BY g."order", sg."order", si."order") AS rn
  FROM anamnesis_template_items ti
  JOIN score_items si ON si.id = ti.score_item_id
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups g ON g.id = sg.group_id
  WHERE ti.anamnesis_template_id IN (
    '11111111-1111-7111-8111-111111111101','11111111-1111-7111-8111-111111111102',
    '11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111106',
    '11111111-1111-7111-8111-111111111107') AND ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = o.rn, updated_at = now() FROM o WHERE t.id = o.id;

-- conferência ----------------------------------------------------------------
SELECT si.name, si.points, si."order",
  (SELECT COUNT(*) FROM score_levels sl WHERE sl.item_id=si.id AND sl.deleted_at IS NULL) AS niveis,
  (SELECT COUNT(*) FROM anamnesis_template_items ti WHERE ti.score_item_id=si.id AND ti.deleted_at IS NULL) AS em_templates
FROM score_items si WHERE si.anamnese_item_code='HIPERTENSAO_ARTERIAL' AND si.deleted_at IS NULL;

COMMIT;
