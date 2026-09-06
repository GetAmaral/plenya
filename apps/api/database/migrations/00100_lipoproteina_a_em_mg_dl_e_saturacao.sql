-- +goose Up
-- 1) LIPOPROTEÍNA(a): DUAS ESCALAS, PORQUE NÃO EXISTE CONVERSÃO VÁLIDA.
--
-- O catálogo só tinha a escala em nmol/L. O laudo brasileiro reporta em mg/dL, o conversor marcou
-- `revisar` por grandezas diferentes e o motor do escore recusou classificar, com o motivo certo:
-- "a escala está em nmol/L e o exame veio em mg/dL". Resultado prático: toda lipoproteína(a) de
-- laboratório brasileiro fica sem pontuar.
--
-- A saída ÓBVIA seria cadastrar um fator em `lab_test_unit_conversions`, e ela está errada. O
-- consenso da European Atherosclerosis Society de 2022 diz explicitamente que não se deve usar
-- fator padrão entre mg/dL e nmol/L: a apo(a) varia de tamanho entre pessoas pelo número de
-- repetições do domínio kringle IV tipo 2, e é esse tamanho que governa a relação entre massa e
-- concentração molar. O fator pragmático de 2 a 2,5 que circula é assumidamente impreciso e
-- específico da amostra. Cadastrá-lo faria o conversor produzir, em silêncio, um número
-- confiantemente errado — exatamente o modo de falha que a rede de plausibilidade existe para
-- evitar, e que ela não pegaria aqui, porque 24,1 × 2,5 = 60 cai dentro da faixa plausível.
--
-- Então a resposta não é converter, é ter as duas escalas e deixar cada laudo ser lido na sua.
-- As fronteiras em massa vêm dos pares que o próprio consenso publica (30 mg/dL ≈ 75 nmol/L e
-- 50 mg/dL ≈ 125 nmol/L); as duas de cima espelham a escala molar pelo fator pragmático de 2,5, e
-- ficam documentadas aqui como o que são, aproximação e não medida.
INSERT INTO public.score_items
  (id, name, unit, points, "order", subgroup_id, lab_test_code, gender,
   clinical_relevance, created_at, updated_at)
SELECT uuid_generate_v7(), 'Lipoproteína A (mg/dL)', 'mg/dL', si.points, si."order",
       si.subgroup_id, si.lab_test_code, 'not_applicable',
       'Mesma escala da Lipoproteína A em nmol/L, em unidade de massa, para os laudos que reportam '
       'em mg/dL. Não é conversão: o consenso EAS 2022 desaconselha fator fixo entre as duas '
       'unidades porque o tamanho da isoforma da apo(a) varia entre indivíduos. As fronteiras de '
       '30 e 50 mg/dL são as que o consenso publica em massa; 100 e 180 mg/dL espelham 250 e 450 '
       'nmol/L pelo fator aproximado de 2,5.',
       now(), now()
FROM public.score_items si
WHERE si.lab_test_code = 'PLNA31F0501' AND si.unit = 'nmol/L' AND si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM public.score_items x
                  WHERE x.lab_test_code = 'PLNA31F0501' AND x.unit = 'mg/dL' AND x.deleted_at IS NULL);

INSERT INTO public.score_levels (id, item_id, level, name, operator, lower_limit, upper_limit, created_at, updated_at)
SELECT uuid_generate_v7(), si.id, v.nivel, v.rotulo, v.op, v.inf, v.sup, now(), now()
FROM public.score_items si
CROSS JOIN (VALUES
  (5, '≤12',     '<=',      NULL,  '12'),
  (4, '12-30',   'between', '12',  '30'),
  (3, '30-50',   'between', '30',  '50'),
  (2, '50-100',  'between', '50',  '100'),
  (1, '100-180', 'between', '100', '180'),
  (0, '>180',    '>',       '180', NULL)
) AS v(nivel, rotulo, op, inf, sup)
WHERE si.lab_test_code = 'PLNA31F0501' AND si.unit = 'mg/dL' AND si.deleted_at IS NULL
  AND NOT EXISTS (SELECT 1 FROM public.score_levels sl WHERE sl.item_id = si.id AND sl.deleted_at IS NULL);

-- 2) SATURAÇÃO DA TRANSFERRINA: A UNIDADE QUE FALTOU NA MIGRATION ANTERIOR.
--
-- A 00099 deixou `PLND7C2752F` de fora com o argumento de que guardava duas grandezas, porque há
-- resultados em µg/dL e em % apontando para ela. O argumento estava errado, e o item de escore
-- prova: "Capacidade de fixação de ferro - IST" tem unidade `%` e níveis em 16, 20, 24, 35 e 45,
-- que são as faixas do ÍNDICE DE SATURAÇÃO, não da capacidade de ligação. A definição é a
-- saturação; os resultados em µg/dL apontando para ela é que estão arquivados no exame errado.
--
-- Preencher a unidade é o que faz esses resultados passarem a ser marcados `revisar` em vez de
-- entrarem crus, e é o que permite a saturação ser avaliada, que é o que a paciente que revelou
-- isto precisava: 41% de saturação com ferritina de 337,6 é a leitura inteira do ferro dela.
UPDATE public.lab_test_definitions
   SET unit = '%'
 WHERE code = 'PLND7C2752F' AND coalesce(btrim(unit), '') = '' AND deleted_at IS NULL;

-- `IST` é uma segunda definição para a mesma coisa ("Índice de Saturação de Transferrina", %) e
-- não tem item de escore nenhum: quem lança ali some do escore sem aviso. Fica marcada, para a
-- curadoria decidir entre fundir com PLND7C2752F ou dar-lhe um item próprio.
COMMENT ON TABLE public.lab_test_definitions IS
  'Catálogo de exames. PENDENTE DE CURADORIA: os códigos IST e PLND7C2752F descrevem ambos o índice de saturação da transferrina; só PLND7C2752F tem item de escore.';

-- +goose Down
COMMENT ON TABLE public.lab_test_definitions IS NULL;

UPDATE public.lab_test_definitions SET unit = '' WHERE code = 'PLND7C2752F' AND unit = '%';

DELETE FROM public.score_levels sl
USING public.score_items si
WHERE sl.item_id = si.id AND si.lab_test_code = 'PLNA31F0501' AND si.unit = 'mg/dL';

DELETE FROM public.score_items
 WHERE lab_test_code = 'PLNA31F0501' AND unit = 'mg/dL' AND name = 'Lipoproteína A (mg/dL)';
