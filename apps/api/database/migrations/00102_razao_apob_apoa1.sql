-- +goose Up
-- RAZÃO apoB/apoA-1: o quociente vale mais que a apoB isolada, e faltava no escore.
--
-- O catálogo tinha `Relação Colesterol Total/HDL` e `Relação Triglicerídeos/HDL`, e não tinha esta.
-- A apoB conta as partículas aterogênicas e a apoA-1 conta as protetoras; o quociente é o balanço
-- entre as duas, e foi a variável de maior poder preditivo para infarto entre os nove fatores de
-- risco convencionais no INTERHEART, com a coorte sueca AMORIS (137.100 pessoas, 17,8 anos de
-- seguimento médio, 22.473 eventos) mostrando a mesma direção em homens e mulheres, em todas as
-- idades.
--
-- AS FAIXAS, E POR QUE SÃO ESTAS
--
-- Não existe tabela consensual, e a literatura diz isso na cara: os próprios autores do AMORIS
-- chamam os seus três níveis de "tentativos" e pedem que diretrizes futuras estabeleçam cortes
-- baseados em evidência. Então as faixas aqui são as três do AMORIS, que é o número publicado, com
-- o ajuste por sexo que a literatura sustenta separadamente.
--
--   AMORIS, figura 4: baixo risco 0,2 a 0,6 · risco médio 0,61 a 0,9 · alto risco 0,91 a 5,0.
--   Corte de risco aumentado, geral: 0,8 (sensibilidade 90%, especificidade 70%).
--   Desfavorável por sexo: acima de 0,90 em homens e acima de 0,80 em mulheres.
--   Médias do AMORIS no basal: cerca de 1,0 em homens e 0,85 em mulheres.
--
-- Daí: o nível 5 é a faixa verde do AMORIS, igual para os dois sexos; o nível 0 é a vermelha, com o
-- limiar deslocado para 0,90 na mulher, que é onde o corte feminino já foi ultrapassado e a média
-- da coorte ficou para trás. No meio, dois níveis que separam "ainda abaixo do corte do seu sexo"
-- de "acima dele". Sem nível 2 e sem nível 4: não há número publicado que sustente essa granulação,
-- e inventar fronteira é pior que ter menos níveis.
--
-- Sem esta migration a razão não pontuaria de todo modo, porque razão não vem em laudo. Quem a
-- calcula é o motor, em internal/services/score_razoes_derivadas.go.
-- TUDO OU NADA, num bloco só.
--
-- A primeira versão pendurava `subgroup_id` e `order` num SELECT sobre o item da relação
-- colesterol/HDL. Numa base onde ele não existe (CI, ambiente recriado — o catálogo de exames NÃO
-- é semeado por migration), o SELECT não devolve linha, zero itens entram, os níveis também não, e
-- o goose marca a migration como aplicada deixando para trás uma definição de exame órfã e nenhum
-- item de escore. Falha em silêncio, que é a pior forma de falhar numa migration.
--
-- E o CROSS JOIN daquela versão tinha um segundo risco: se a relação colesterol/HDL um dia for
-- dividida por sexo, como 93 outros itens do catálogo já são, ele devolveria duas linhas por sexo e
-- o NOT EXISTS não enxerga linhas da própria instrução. Sairiam itens duplicados, valendo 36 pontos
-- em vez de 18. Por isso o subgrupo vem de um SELECT com LIMIT 1, guardado numa variável.
-- +goose StatementBegin
DO $apo$
DECLARE
  v_subgrupo uuid;
  v_ordem    int;
  v_item     uuid;
  v_sexo     text;
BEGIN
  SELECT si.subgroup_id, si."order" INTO v_subgrupo, v_ordem
    FROM public.score_items si
   WHERE si.lab_test_code = 'PLN09DBBB62' AND si.deleted_at IS NULL
   ORDER BY si."order"
   LIMIT 1;

  IF v_subgrupo IS NULL THEN
    RAISE NOTICE 'catálogo de exames ainda não semeado (PLN09DBBB62 ausente): 00102 não tem onde pendurar a razão apoB/apoA-1 e não fez nada';
    RETURN;
  END IF;

  INSERT INTO public.lab_test_definitions
    (id, code, name, short_name, alt_names, category, is_requestable, unit, result_type,
     sex_applicability, is_active, display_order, clinical_significance, created_at, updated_at)
  SELECT uuid_generate_v7(), 'PLNAPOBA1', 'Relação Apolipoproteína B/A1', 'ApoB/ApoA1',
         '["apob/apoa1", "apob/apoa-1", "relacao apob apoa1", "razao apob apoa1", "apo b/apo a1"]'::jsonb,
         'biochemistry', false, 'ratio', 'numeric', 'all', true,
         coalesce((SELECT display_order FROM public.lab_test_definitions
                    WHERE code = 'PLN09DBBB62' AND deleted_at IS NULL), 0),
         'Balanço entre as partículas aterogênicas (apoB) e as protetoras (apoA-1). Não vem em laudo: '
         'é calculada pelo motor do escore a partir das duas dosagens do MESMO lote e na MESMA unidade.',
         now(), now()
  WHERE NOT EXISTS (SELECT 1 FROM public.lab_test_definitions WHERE code = 'PLNAPOBA1');

  -- Um item por sexo, como o catálogo já faz em DHEA-S, ferritina e testosterona.
  FOREACH v_sexo IN ARRAY ARRAY['female', 'male'] LOOP
    SELECT id INTO v_item FROM public.score_items
     WHERE lab_test_code = 'PLNAPOBA1' AND gender = v_sexo AND deleted_at IS NULL
     LIMIT 1;

    IF v_item IS NULL THEN
      INSERT INTO public.score_items
        (id, name, unit, points, "order", subgroup_id, lab_test_code, gender, clinical_relevance,
         created_at, updated_at)
      VALUES (uuid_generate_v7(),
        CASE v_sexo WHEN 'female' THEN 'Relação ApoB/ApoA1 - Mulheres' ELSE 'Relação ApoB/ApoA1 - Homens' END,
        'ratio', 18, v_ordem, v_subgrupo, 'PLNAPOBA1', v_sexo,
        'Razão entre apolipoproteína B e apolipoproteína A1. Foi a variável de maior poder preditivo '
        'para infarto entre os nove fatores de risco convencionais do INTERHEART, e prediz eventos '
        'melhor que a apoB isolada e que as razões de colesterol. As faixas seguem os três níveis '
        'publicados no AMORIS (baixo até 0,6; médio de 0,61 a 0,9; alto acima de 0,91), com o limiar '
        'superior ajustado ao corte desfavorável de cada sexo: 0,80 na mulher e 0,90 no homem. Os '
        'próprios autores classificam esses cortes como tentativos.',
        now(), now())
      RETURNING id INTO v_item;
    END IF;

    INSERT INTO public.score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, created_at, updated_at)
    SELECT uuid_generate_v7(), v_item, v.nivel, v.rotulo, v.op, v.inf, v.sup, now(), now()
    FROM (VALUES
      ('female', 5, '≤0,6',    '<=',      NULL,   '0.6'),
      ('female', 3, '0,6-0,8', 'between', '0.6',  '0.8'),
      ('female', 1, '0,8-0,9', 'between', '0.8',  '0.9'),
      ('female', 0, '>0,9',    '>',       '0.9',  NULL),
      ('male',   5, '≤0,6',    '<=',      NULL,   '0.6'),
      ('male',   3, '0,6-0,9', 'between', '0.6',  '0.9'),
      ('male',   1, '0,9-1,0', 'between', '0.9',  '1.0'),
      ('male',   0, '>1,0',    '>',       '1.0',  NULL)
    ) AS v(sexo, nivel, rotulo, op, inf, sup)
    WHERE v.sexo = v_sexo
      AND NOT EXISTS (SELECT 1 FROM public.score_levels sl
                       WHERE sl.item_id = v_item AND sl.deleted_at IS NULL);
  END LOOP;
END
$apo$;
-- +goose StatementEnd

-- +goose Down
DELETE FROM public.score_levels sl USING public.score_items si
 WHERE sl.item_id = si.id AND si.lab_test_code = 'PLNAPOBA1';
DELETE FROM public.score_items WHERE lab_test_code = 'PLNAPOBA1';
DELETE FROM public.lab_test_definitions WHERE code = 'PLNAPOBA1';
