-- Liga cada substância do catálogo ao nutriente correspondente do Anexo IV da IN 28.
--
-- O fator diz quantas unidades do Anexo IV valem UMA unidade da substância. Onde o fator seria
-- chute, a substância fica SEM mapa: teto conferido com fator inventado é pior que teto ausente.
-- Por isso acetil-L-carnitina não entra (a fração de carnitina muda com o sal) e a vitamina C do
-- catálogo, cadastrada em %, também não.
--
-- Para minerais, o motor converte a dose para elemento pelo elemental_percent antes de comparar:
-- o fator aqui é só conversão de unidade.

BEGIN;

UPDATE magistral_components mc SET in28_nutrient = m.nutriente, in28_factor = m.fator
  FROM (VALUES
    -- minerais (a conversão para elemento é do motor)
    ('Cálcio',                'Cálciov',        1),
    ('Cobre',                 'Cobre',          1000),   -- catálogo em mg, norma em µg
    ('Ferro',                 'Ferro',          1),
    ('Iodo',                  'Iodo',           1),
    ('Selênio',               'Selênio',        1),
    ('Selenometionina',       'Selênio',        1),
    ('Zinco quelato',         'Zinco',          1),
    ('Zinco carnosina',       'Zinco',          1),
    ('Picolinato de cromo',   'Cromo',          1),
    ('Magnésio quelato',      'Magnésio',       1),
    ('Magnésio L-treonato',   'Magnésio',       1),
    ('Magnésio dimalato',     'Magnésio',       1),
    ('Magnésio citrato',      'Magnésio',       1),
    ('Magnésio taurato',      'Magnésio',       1),
    ('Magnésio aspartato',    'Magnésio',       1),
    ('Magnésio ascorbato',    'Magnésio',       1),
    ('Magnésio carbonato',    'Magnésio',       1),
    ('Magnésio óxido',        'Magnésio',       1),
    ('Magnésio sulfato',      'Magnésio',       1),
    ('Magnésio inositol',     'Magnésio',       1),
    ('Cloreto de magnésio',   'Magnésio',       1),

    -- vitaminas
    ('Biotina',               'Biotina',        1),
    ('Riboflavina',           'Riboflavina',    1),
    ('Nicotinamida',          'Niacina',        1),
    ('Piridoxal-5-fosfato',   'Vitamina B6',    1),      -- conta a massa do P5P, superestima ~30%
    ('Metilcobalamina',       'Vitamina B12',   1),
    ('Cianocobalamina',       'Vitamina B12',   1),
    ('Metilfolato',           'Ácido fólico iv', 1),
    ('Vitamina A',            'Vitamina Ai',    0.3),     -- 1 UI de retinol = 0,3 µg RE
    ('Vitamina D3',           'Vitamina Dii',   0.025),   -- 1 µg = 40 UI
    ('Vitamina E',            'Vitamina Eiii',  1),
    ('Vitamina K2 MK-7',      'Vitamina K',     1),
    ('Palmitato de ascorbila','Vitamina C',     0.43),    -- 43% do peso é ácido ascórbico
    ('Colina',                'Colina',         1),
    ('Alfa-GPC',              'Colina',         0.4),     -- alfa-GPC tem 40% de colina

    -- aminoácidos e bioativos
    ('L-teanina',             'Teanina',        1),
    ('L-taurina',             'Taurina',        1),
    ('L-carnitina',           'L-Carnitina',    1),
    ('Creatina',              'Creatina',       1),
    ('Glicina',               'Glicina',        1),
    ('L-glutamina',           'Glutamina',      1),
    ('L-tirosina',            'Tirosina',       1),
    ('L-triptofano',          'Triptofano',     1),
    ('Mio-inositol',          'Inositol',       0.001),   -- catálogo em mg, norma em g
    ('D-quiro-inositol',      'Inositol',       0.001),
    ('Coenzima Q10',          'Coenzima Q10',   1),
    ('Ubiquinol',             'Coenzima Q10',   1),
    ('CavaQ10',               'Coenzima Q10',   1),
    ('Licopeno',              'Licopeno',       1),
    ('Ácido hialurônico',     'Ácido hialurônico', 1),
    ('GABA',                  'Ácido gama aminobutírico (GABA)', 1),
    ('MSM',                   'Metilsulfonilmetano', 1),
    ('D-ribose',              'D-ribose',       1),
    ('Ômega-3',               'EPA e DHA',      1)
  ) AS m(nome, nutriente, fator)
 WHERE mc.name = m.nome;

-- Zinco carnosina tem 23% de zinco: sem isso, 100 mg do insumo virariam 100 mg de zinco.
UPDATE magistral_components
   SET elemental_percent = 23,
       correction_note = 'A carnosina de zinco (polaprezinco) tem 23% de zinco elementar.'
 WHERE name = 'Zinco carnosina' AND elemental_percent IS NULL;

COMMIT;
