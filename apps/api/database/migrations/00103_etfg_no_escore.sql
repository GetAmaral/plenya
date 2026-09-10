-- +goose Up
-- eTFG NO ESCORE: o número renal que decide conduta, e que não tinha item nenhum.
--
-- O catálogo tem `ETFG` desde sempre, e `score_items` não tinha uma única linha apontando para ele:
-- toda eTFG que entrou em prontuário entrou como resultado sem nível, fora do escore e fora da
-- régua. A função renal pontuava por creatinina, cistatina C, ureia e microalbuminúria, que são os
-- insumos, e não pelo número que a diretriz usa para estagiar.
--
-- Quem calcula é o motor, em internal/services/score_tfg_derivada.go, porque eTFG não vem em laudo
-- como medida: é conta. O laudo imprime uma, mas é SEMPRE a de creatinina isolada, mesmo quando a
-- cistatina C foi dosada na mesma coleta — e é justamente aí que as duas divergem.
--
-- AS FAIXAS
--
-- São os estágios da KDIGO (G1 a G5), que é também exatamente o que o laboratório imprime ao lado
-- do resultado: >90 normal · 89-60 redução discreta · 59-45 discreta a moderada · 44-30 moderada a
-- severa · 29-15 severa · <15 falência renal.
--
-- Os limites seguem a convenção meio-aberta do projeto, `(inferior, superior]`, então um valor
-- exatamente igual à fronteira cai na faixa DE BAIXO: 60 fica em G3a e não em G2. É um ponto de
-- diferença contra a leitura literal da KDIGO, e é o preço de manter uma convenção só no sistema
-- inteiro em vez de uma exceção aqui. Na prática a eTFG calculada é contínua e cai em cima da
-- fronteira raramente; quando cair, erra para o lado conservador, que é o lado certo de errar em
-- doença renal.
--
-- SOBREPOSIÇÃO DE PONTOS: creatinina valia 20 e cistatina C valia 18, e a eTFG é conta sobre as
-- duas — com este item a mesma fisiologia passaria a pontuar 60 pontos. A decisão de peso é clínica
-- e não cabia aqui; foi tomada e está na 00104, que baixa as duas para 15 cada (total renal 52).
-- +goose StatementBegin
DO $tfg$
DECLARE
  v_subgrupo uuid;
  v_ordem    int;
  v_item     uuid;
BEGIN
  -- Pendura no mesmo subgrupo e perto da ordem da creatinina, que é o vizinho natural na tela.
  SELECT si.subgroup_id, si."order" INTO v_subgrupo, v_ordem
    FROM public.score_items si
   WHERE si.lab_test_code = 'PLN357DC859' AND si.deleted_at IS NULL
   ORDER BY si."order"
   LIMIT 1;

  IF v_subgrupo IS NULL THEN
    RAISE NOTICE 'catálogo de exames ainda não semeado (creatinina ausente): 00103 não tem onde pendurar a eTFG e não fez nada';
    RETURN;
  END IF;

  -- A definição de catálogo já existe em dev e prod; o INSERT é a rede para base recriada.
  INSERT INTO public.lab_test_definitions
    (id, code, name, short_name, alt_names, category, is_requestable, unit, result_type,
     sex_applicability, is_active, display_order, clinical_significance, created_at, updated_at)
  SELECT uuid_generate_v7(), 'ETFG', 'Taxa de Filtração Glomerular Estimada', 'eTFG',
         '["etfg", "tfg", "egfr", "gfr", "filtracao glomerular", "taxa de filtracao glomerular", "ckd-epi"]'::jsonb,
         'biochemistry', false, 'mL/min/1.73m²', 'numeric', 'all', true, 0,
         'Estágio da doença renal crônica. Não é medida: é calculada pelo motor do escore por '
         'CKD-EPI 2021, preferindo a equação combinada creatinina + cistatina C quando as duas '
         'existem na mesma coleta, porque a creatinina isolada superestima a perda de função em '
         'quem tem massa muscular alta ou suplementa creatina.',
         now(), now()
  WHERE NOT EXISTS (SELECT 1 FROM public.lab_test_definitions WHERE code = 'ETFG' AND deleted_at IS NULL);

  SELECT id INTO v_item FROM public.score_items
   WHERE lab_test_code = 'ETFG' AND deleted_at IS NULL
   LIMIT 1;

  IF v_item IS NULL THEN
    INSERT INTO public.score_items
      (id, name, unit, points, "order", subgroup_id, lab_test_code, clinical_relevance,
       patient_explanation, conduct, created_at, updated_at)
    VALUES (uuid_generate_v7(),
      'Taxa de Filtração Glomerular Estimada (CKD-EPI)', 'mL/min/1.73m²', 22, v_ordem, v_subgrupo, 'ETFG',
      'Estagiamento da doença renal crônica pela KDIGO. É o desfecho que organiza toda a conduta '
      'renal: meta de pressão, ajuste de dose de drogas com eliminação renal, momento de encaminhar '
      'ao nefrologista e elegibilidade para contraste. Calculada por CKD-EPI 2021, com preferência '
      'pela equação combinada creatinina + cistatina C quando as duas existem na mesma coleta — é a '
      'que a KDIGO recomenda quando a creatinina é suspeita, e a creatinina é suspeita sempre que '
      'há massa muscular alta, suplementação de creatina, amputação, desnutrição ou dieta '
      'vegetariana.',
      'A filtração glomerular é o quanto os seus rins conseguem limpar o sangue por minuto. É o '
      'número que diz em que estágio está a função renal, e ele vale mais do que a creatinina '
      'isolada porque leva em conta idade e sexo, e porque pode ser calculado também pela cistatina '
      'C, que não depende da sua massa muscular.',
      'Confirmar toda queda com nova coleta antes de estagiar: filtração é tendência, não ponto '
      'isolado. Havendo massa muscular alta, suplementação de creatina ou discordância entre '
      'creatinina e o quadro clínico, dosar cistatina C na MESMA coleta e usar a equação combinada.',
      now(), now())
    RETURNING id INTO v_item;
  END IF;

  -- Faixas KDIGO G1-G5. Sem sobreposição: a ordem de avaliação vai do nível 0 para cima e a
  -- primeira faixa que bate vence, então limites que se cruzam entregariam sempre o nível pior.
  INSERT INTO public.score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, created_at, updated_at)
  SELECT uuid_generate_v7(), v_item, v.nivel, v.rotulo, v.op, v.inf, v.sup, now(), now()
  FROM (VALUES
    (5, 'G1 · >90',      '>',       '90',  NULL),
    (4, 'G2 · 60-90',    'between', '60',  '90'),
    (3, 'G3a · 45-60',   'between', '45',  '60'),
    (2, 'G3b · 30-45',   'between', '30',  '45'),
    (1, 'G4 · 15-30',    'between', '15',  '30'),
    (0, 'G5 · ≤15',      '<=',      NULL,  '15')
  ) AS v(nivel, rotulo, op, inf, sup)
  WHERE NOT EXISTS (SELECT 1 FROM public.score_levels sl
                     WHERE sl.item_id = v_item AND sl.deleted_at IS NULL);

  -- Pilares AGIR: Renal e Cardiovascular, os mesmos da cistatina C. Filtração baixa é fator de
  -- risco cardiovascular independente, e não só um marcador de rim.
  INSERT INTO public.score_item_method_pillars (score_item_id, method_pillar_id)
  SELECT v_item, p.id
    FROM public.method_pillars p
   WHERE p.name IN ('Renal', 'Cardiovascular')
     AND NOT EXISTS (SELECT 1 FROM public.score_item_method_pillars x
                      WHERE x.score_item_id = v_item AND x.method_pillar_id = p.id);
END
$tfg$;
-- +goose StatementEnd

-- +goose Down
-- SOFT delete, não DELETE. Assim que um paciente tem snapshot com este item,
-- `patient_score_item_results` o referencia com RESTRICT e o DELETE falha:
--   violates RESTRICT setting of foreign key constraint "fk_patient_score_item_results_item"
-- Apagar as linhas dependentes primeiro passaria, ao custo de destruir o histórico de escore de
-- quem já foi pontuado — o preço errado para desfazer uma migration. Como todo o sistema filtra por
-- `deleted_at IS NULL`, marcar apagado tira o item do escore e preserva o que já foi mostrado ao
-- paciente. Um `up` posterior cria um item novo, que é o comportamento certo.
DELETE FROM public.score_item_method_pillars p USING public.score_items si
 WHERE p.score_item_id = si.id AND si.lab_test_code = 'ETFG';
UPDATE public.score_levels sl SET deleted_at = now()
  FROM public.score_items si
 WHERE sl.item_id = si.id AND si.lab_test_code = 'ETFG' AND sl.deleted_at IS NULL;
UPDATE public.score_items SET deleted_at = now()
 WHERE lab_test_code = 'ETFG' AND deleted_at IS NULL;
-- A definição de catálogo NÃO é removida: ela é anterior a esta migration e há resultados de
-- paciente apontando para ela.
