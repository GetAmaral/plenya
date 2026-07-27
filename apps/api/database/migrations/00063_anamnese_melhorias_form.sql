-- Melhorias no form de anamnese e na aplicação de templates (pontos levantados pelo Dr. Getúlio
-- em 2026-07-27). Racional completo em docs/emr/plano-melhorias-anamnese-2026-07.md.
--
-- (1) UNIDADES em "Composição corporal > Medidas Objetivas". Os 47 itens estavam com
--     score_items.unit NULL — e o form só desenha o campo numérico quando o item tem unidade.
--     Resultado: peso, altura, quadril e companhia apareciam sem nenhuma forma de digitar o
--     valor, e nada era calculado a partir deles. Com a unidade preenchida, o EMR passa a
--     mostrar o input e a calcular sozinho os derivados (IMC, BRI, razões, FMI, ASMI, MME/peso,
--     percentuais de água) — o motor está em packages/domain/src/anthropometry.ts.
--
-- (2) ÓRFÃOS DE parent_item_id. Alguns filhos ficaram com parent_item_id NULL e, como o
--     renderer promove a raiz todo filho sem pai presente, saíam soltos no meio da lista:
--       Cirurgias já realizadas → "Amputação de membro", "Adrenalectomia"
--       Medicamentos           → 15 classes de "Analgésicos…" a "Antivirais…" + "Inibidores
--                                de bomba de prótons…"
--       Vida Sexual > Atual    → ASEX (o IIEF-5, irmão, já apontava para "Escalas de desempenho:")
--
-- (3) VÍCIOS > SEXO com N5 default, como Tabaco/Álcool/Drogas/Jogos já tinham.
--     ("Outros vícios" fica de fora de propósito — pendente de confirmação do Getúlio.)
--
-- (4) DUBOIS — SCORE TOTAL PONDERADO. O código contava qualquer evocação como 1, então lembrar
--     COM DICA pontuava igual a lembrar sem dica. A cotação canônica (Dubois B. et al., Presse
--     Med 2002;31(36):1696-9; Cowppli-Bony P. et al., Rev Neurol 2005;161(12 Pt 1):1205-12;
--     Croisile B. et al., Rev Neurol 2010;166(8-9):711-20; ficha de administração INESSS/HAS)
--     dá 2 pontos ao rappel libre, 1 ao rappel indicé e 0 à palavra não evocada — cada fase
--     vale /10 e o Score Total Ponderado (imediato + tardio) vale /20, com corte ≤18 para
--     investigar. É o escore mais discriminante que o "score total" antigo /5 por fase.
--     Aqui reescalamos as faixas /5 → /10 e recalculamos o que já estava gravado (as respostas
--     individuais em scale_responses.answers já usavam os pesos 0/1/2, então a conversão é exata).
--
-- (5) ASEX e "Épocas de melhor/pior libido/desempenho" saem do Continuum | Médico | Inicial e
--     vão para o Continuum | Médico | Complemento, onde IIEF-5 e o pai "Escalas de desempenho:"
--     já estão. Reverte de propósito o swap ASEX↔IIEF-5 feito na 00058. O template completo
--     Médico | Inicial (superset) mantém tudo.
--
-- (6) ORDEM DOS TEMPLATES. anamnesis_template_items."order" foi gravado embaralhado em relação
--     à árvore do escore, então o form não seguia a ordem do grupo original (ex.: em
--     Médico | Inicial, Cognição/Atual saía 107 PHQ-9, 108 Dubois, 110 "Capacidade da memória
--     percebida", quando a árvore diz Capacidade(1) → Testes rápidos de memória(2) → filhos).
--     Renormalizamos TODOS os templates por (grupo, subgrupo, item), mantendo cada filho logo
--     abaixo do seu pai — que é exatamente a ordem em que o renderer desenha.

-- +goose Up

-- ███ (1) UNIDADES — Composição corporal > Medidas Objetivas ████████████████████████████████
UPDATE score_items si SET unit = u.unit, updated_at = now()
FROM (VALUES
  ('PESO','kg'),
  ('ALTURA','cm'),
  ('IMC','kg/m²'),
  ('BRI','índice'),
  ('TAXA_METABOLICA_BASAL','kcal'),
  ('ABDOMINAL_HOMEM','cm'),
  ('ABDOMINAL_MULHER','cm'),
  ('QUADRIL','cm'),
  ('RAZAO_CINTURA_QUADRIL_HOMEM','razão'),
  ('RAZAO_CINTURA_QUADRIL_MULHER','razão'),
  ('RAZAO_CINTURA_ALTURA','cm/cm'),
  ('PESCOCO_HOMEM','cm'),
  ('PESCOCO_MULHER','cm'),
  ('RELACAO_PESCOCO_ALTURA_HOMEM','cm/m'),
  ('RELACAO_PESCOCO_ALTURA_MULHER','cm/m'),
  ('BRACO_DIREITO_CONTRAIDO','cm'),
  ('BRACO_ESQUERDO_CONTRAIDO','cm'),
  ('BRACO_NAO_DOMINANTE_RELAXADO_HOMEM','cm'),
  ('BRACO_NAO_DOMINANTE_RELAXADO_MULHER','cm'),
  ('COXA_HOMEM','cm'),
  ('COXA_MULHER','cm'),
  ('PANTURRILHA_HOMEM','cm'),
  ('PANTURRILHA_MULHER','cm'),
  ('GORDURA_CORPORAL_HOMEM','%'),
  ('GORDURA_CORPORAL_MULHER','%'),
  ('MASSA_GORDA_TOTAL','kg'),
  ('FMI_FAT_MASS_INDEX_HOMEM','kg/m²'),
  ('FMI_FAT_MASS_INDEX_MULHER','kg/m²'),
  ('GORDURA_VISCERAL_HOMEM','cm²'),
  ('GORDURA_VISCERAL_MULHER','cm²'),
  ('GORDURA_DE_TRONCO','%'),
  ('RAZAO_ANDROIDE_GINOIDE_HOMEM','razão'),
  ('RAZAO_ANDROIDE_GINOIDE_MULHER','razão'),
  ('MASSA_MUSCULAR_ESQUELETICA','kg'),
  ('MME_PESO','%'),
  ('INDICE_MME','kg/m²'),
  ('MASSA_APENDICULAR','kg'),
  ('ASMI_HOMEM','kg/m²'),
  ('ASMI_MULHER','kg/m²'),
  ('AGUA_CORPORAL_TOTAL','L'),
  ('AGUA_CORPORAL_TOTAL_HOMEM','%'),
  ('AGUA_CORPORAL_TOTAL_MULHER','%'),
  ('AGUA_INTRACELULAR','L'),
  ('AGUA_EXTRACELULAR','L'),
  ('RAZAO_AEC_ACT','%'),
  ('ANGULO_DE_FASE_HOMEM','°'),
  ('ANGULO_DE_FASE_MULHER','°')
) AS u(code, unit)
WHERE si.anamnese_item_code = u.code
  AND si.deleted_at IS NULL
  AND (si.unit IS NULL OR si.unit = '');

-- ███ (2) ÓRFÃOS DE parent_item_id ██████████████████████████████████████████████████████████
-- Cirurgias: tudo que não é o cabeçalho do subgrupo pendura em "Cirurgias que interferem
-- diretamente no escore".
UPDATE score_items c
SET parent_item_id = p.id, updated_at = now()
FROM score_items p
WHERE p.anamnese_item_code = 'CIRURGIAS_QUE_INTERFEREM_DIRETAMENTE_NO_ESCORE'
  AND p.deleted_at IS NULL
  AND c.subgroup_id = p.subgroup_id
  AND c.deleted_at IS NULL
  AND c.parent_item_id IS NULL
  AND c.id <> p.id
  AND c.anamnese_item_code <> 'REGISTRAR_QUAISQUER_CIRURGIAS_REALIZADAS';

-- Medicamentos: as classes penduram em "Uso atual de medicamentos".
UPDATE score_items c
SET parent_item_id = p.id, updated_at = now()
FROM score_items p
WHERE p.anamnese_item_code = 'USO_ATUAL_DE_MEDICAMENTOS'
  AND p.deleted_at IS NULL
  AND c.subgroup_id = p.subgroup_id
  AND c.deleted_at IS NULL
  AND c.parent_item_id IS NULL
  AND c.id <> p.id
  AND c.anamnese_item_code <> 'HISTORICO_DE_MEDICAMENTOS_UTILIZADOS_REACOES_ADVERSAS_RESPOSTAS_TERAPEUTICAS';

-- ASEX pendura em "Escalas de desempenho:", como o IIEF-5.
UPDATE score_items c
SET parent_item_id = p.id, updated_at = now()
FROM score_items p
WHERE p.anamnese_item_code = 'ESCALAS_DE_DESEMPENHO' AND p.deleted_at IS NULL
  AND c.anamnese_item_code = 'ASEX_25' AND c.deleted_at IS NULL
  AND c.parent_item_id IS NULL;

-- ███ (3) VÍCIOS > SEXO — N5 default ████████████████████████████████████████████████████████
UPDATE score_items SET default_level5 = true, updated_at = now()
WHERE anamnese_item_code = 'SEXO' AND deleted_at IS NULL AND default_level5 = false;

-- ███ (4) DUBOIS — faixas /5 → /10 (score total ponderado) ██████████████████████████████████
-- imediato: N5 = 10 (todas espontâneas) · N3 = 9 (uma com dica) · N0 ≤ 8
UPDATE score_levels SET operator='<=', lower_limit=NULL, upper_limit='8',  name='≤8',  updated_at=now() WHERE id='019bf7af-ba01-7681-99e7-a9bdbeb76df7';
UPDATE score_levels SET operator='=',  lower_limit='9',  upper_limit='9',  name='9',   updated_at=now() WHERE id='019bf7af-ba01-7bab-ae6f-0ae7f8cc80dc';
UPDATE score_levels SET operator='=',  lower_limit='10', upper_limit='10', name='10',  updated_at=now() WHERE id='019bf7af-ba01-7b86-a639-d1ee15894e87';
-- tardio: N5 = 10 · N4 = 9 · N2 = 8 · N0 ≤ 7
UPDATE score_levels SET operator='<=', lower_limit=NULL, upper_limit='7',  name='≤7',  updated_at=now() WHERE id='019bf7af-d85e-7c93-bdfc-f6bfb2b7fdb6';
UPDATE score_levels SET operator='=',  lower_limit='8',  upper_limit='8',  name='8',   updated_at=now() WHERE id='019bf7af-d85e-7b7b-a04d-f5e4c4530087';
UPDATE score_levels SET operator='=',  lower_limit='9',  upper_limit='9',  name='9',   updated_at=now() WHERE id='019bf7af-d85e-74ef-8196-e5e3b889046c';
UPDATE score_levels SET operator='=',  lower_limit='10', upper_limit='10', name='10',  updated_at=now() WHERE id='019bf7af-d85e-70fb-a94b-d451189ce2a8';

-- Recalcula o que já estava gravado: total = Σ dos pesos das respostas (0/1/2, já persistidos),
-- numeric_value = total. O nível é reclassificado logo abaixo pelas faixas novas.
UPDATE anamnesis_items ai
SET scale_responses = jsonb_set(
      ai.scale_responses,
      '{total}',
      to_jsonb((SELECT COALESCE(SUM((v.value)::int), 0)
                FROM jsonb_each(ai.scale_responses -> 'answers') AS v(key, value)))
    ),
    numeric_value = (SELECT COALESCE(SUM((v.value)::int), 0)
                     FROM jsonb_each(ai.scale_responses -> 'answers') AS v(key, value)),
    updated_at    = now()
WHERE ai.score_item_id IN (
        SELECT id FROM score_items
        WHERE anamnese_item_code IN ('5_PALAVRAS_DE_DUBOIS_IMEDIATO_5','5_PALAVRAS_DE_DUBOIS_TARDIO_5')
          AND deleted_at IS NULL)
  AND ai.deleted_at IS NULL
  AND ai.scale_responses IS NOT NULL
  AND ai.scale_responses ? 'answers';

-- Reclassifica o selected_level pelas faixas novas (mesma semântica do motor Go:
-- "<=" lê upper_limit, "=" lê lower_limit).
UPDATE anamnesis_items ai
SET selected_level = l.level, updated_at = now()
FROM score_items si, score_levels l
WHERE ai.score_item_id = si.id
  AND l.item_id = si.id AND l.deleted_at IS NULL
  AND si.anamnese_item_code IN ('5_PALAVRAS_DE_DUBOIS_IMEDIATO_5','5_PALAVRAS_DE_DUBOIS_TARDIO_5')
  AND ai.deleted_at IS NULL
  AND ai.numeric_value IS NOT NULL
  AND (
        (l.operator = '<=' AND l.upper_limit IS NOT NULL AND ai.numeric_value <= l.upper_limit::numeric)
     OR (l.operator = '='  AND l.lower_limit IS NOT NULL AND ai.numeric_value  = l.lower_limit::numeric)
      );

-- ███ (5) ASEX + épocas de libido → Continuum | Médico | Complemento ████████████████████████
CREATE TEMP TABLE mov_063(src uuid, dst uuid, code text) ON COMMIT DROP;
INSERT INTO mov_063 (src, dst, code) VALUES
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105','ASEX_25'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105','EPOCAS_DE_MELHOR_LIBIDO_DESEMPENHO'),
 ('11111111-1111-7111-8111-111111111104','11111111-1111-7111-8111-111111111105','EPOCAS_DE_PIOR_LIBIDO_DESEMPENHO');

CREATE TEMP TABLE res_063 ON COMMIT DROP AS
  SELECT m.src, m.dst, si.id AS score_item_id
  FROM mov_063 m
  JOIN score_items si ON si.anamnese_item_code = m.code AND si.deleted_at IS NULL;

UPDATE anamnesis_template_items t SET deleted_at = now()
FROM res_063 r
WHERE t.anamnesis_template_id = r.src AND t.score_item_id = r.score_item_id AND t.deleted_at IS NULL;

INSERT INTO anamnesis_template_items (id, anamnesis_template_id, score_item_id, "order", created_at, updated_at)
SELECT uuid_generate_v7(), r.dst, r.score_item_id, 0, now(), now()
FROM res_063 r
WHERE NOT EXISTS (
  SELECT 1 FROM anamnesis_template_items x
  WHERE x.anamnesis_template_id = r.dst AND x.score_item_id = r.score_item_id AND x.deleted_at IS NULL);

-- ███ (6) ORDEM CANÔNICA EM TODOS OS TEMPLATES ██████████████████████████████████████████████
-- Chave de ordenação = a ordem em que o form desenha: grupo, subgrupo, e dentro do subgrupo
-- cada pai seguido dos seus filhos (COALESCE(pai."order", si."order") agrupa o filho com o pai;
-- o flag "é filho" garante o pai na frente; si."order" ordena os irmãos).
WITH canonica AS (
  SELECT ti.id,
         ROW_NUMBER() OVER (
           PARTITION BY ti.anamnesis_template_id
           ORDER BY g."order", g.name,
                    sg."order", sg.name,
                    COALESCE(pai."order", si."order"),
                    COALESCE(pai.name, si.name),
                    (si.parent_item_id IS NOT NULL),
                    si."order", si.name
         ) AS rn
  FROM anamnesis_template_items ti
  JOIN score_items     si  ON si.id  = ti.score_item_id
  LEFT JOIN score_items pai ON pai.id = si.parent_item_id AND pai.deleted_at IS NULL
  JOIN score_subgroups sg  ON sg.id  = si.subgroup_id
  JOIN score_groups    g   ON g.id   = sg.group_id
  WHERE ti.deleted_at IS NULL
)
UPDATE anamnesis_template_items t SET "order" = c.rn, updated_at = now()
FROM canonica c
WHERE t.id = c.id AND t."order" IS DISTINCT FROM c.rn;

-- +goose Down

-- Faixas do Dubois — volta ao score total antigo (/5 por fase)
UPDATE score_levels SET operator='<=', lower_limit=NULL, upper_limit='3', name='<=3', updated_at=now() WHERE id='019bf7af-ba01-7681-99e7-a9bdbeb76df7';
UPDATE score_levels SET operator='=',  lower_limit='4',  upper_limit='4', name='4',   updated_at=now() WHERE id='019bf7af-ba01-7bab-ae6f-0ae7f8cc80dc';
UPDATE score_levels SET operator='=',  lower_limit='5',  upper_limit='5', name='5',   updated_at=now() WHERE id='019bf7af-ba01-7b86-a639-d1ee15894e87';
UPDATE score_levels SET operator='<=', lower_limit=NULL, upper_limit='2', name='<=2', updated_at=now() WHERE id='019bf7af-d85e-7c93-bdfc-f6bfb2b7fdb6';
UPDATE score_levels SET operator='=',  lower_limit='3',  upper_limit='3', name='3',   updated_at=now() WHERE id='019bf7af-d85e-7b7b-a04d-f5e4c4530087';
UPDATE score_levels SET operator='=',  lower_limit='4',  upper_limit='4', name='4',   updated_at=now() WHERE id='019bf7af-d85e-74ef-8196-e5e3b889046c';
UPDATE score_levels SET operator='=',  lower_limit='5',  upper_limit='5', name='5',   updated_at=now() WHERE id='019bf7af-d85e-70fb-a94b-d451189ce2a8';

-- Vícios > Sexo
UPDATE score_items SET default_level5 = false, updated_at = now()
WHERE anamnese_item_code = 'SEXO' AND deleted_at IS NULL;

-- Unidades de Composição corporal > Medidas Objetivas
UPDATE score_items si SET unit = NULL, updated_at = now()
FROM score_subgroups sg, score_groups g
WHERE si.subgroup_id = sg.id AND sg.group_id = g.id
  AND g.name = 'Composição corporal' AND sg.name = 'Medidas Objetivas'
  AND si.deleted_at IS NULL;

-- Órfãos de parent_item_id, moves entre templates, ordem dos templates e os totais recalculados
-- do Dubois: reversão fiel exigiria snapshot do estado anterior (não capturado). No-op
-- proposital — restaurar do backup se preciso. Mesmo critério da 00058.
SELECT 1;
