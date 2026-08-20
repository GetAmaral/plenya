-- Fórmulas-base vindas dos formulários magistrais publicados (pesquisa externa).
--
-- POR QUE ESTE ARQUIVO EXISTE: o material da pós descreve fórmula completa em algumas frentes
-- (sono, mitocôndria, glicemia, cognição) e em outras discute a substância isolada sem fechar a
-- associação. Onde o RAG não fechou fórmula, estas vêm de formulários publicados por farmácias de
-- manipulação brasileiras, com a fonte anotada em `notes`.
--
-- Todas nascem PARA CONFERÊNCIA. São ponto de partida do repertório, não protocolo: as doses são
-- as do formulário consultado, e quem decide o que fica é o médico.
--
-- Nenhuma regra de dose é criada aqui. Regra tem trava de piso e teto e é decisão clínica —
-- cadastra-se na tela de fórmulas-base.
--
-- Idempotente: não duplica se rodar de novo.

-- +goose NO TRANSACTION (não é migration; é seed)

WITH nova AS (
  INSERT INTO public.magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT * FROM (VALUES
    (uuidv7(), 'Sono completo',
     'Associação noturna que cobre precursor de serotonina, relaxamento muscular e modulação gabaérgica.',
     E'insônia de indução e manutenção\nquando o precursor isolado não basta',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula ao deitar', 60,
     'Formulário de farmácia magistral (Farmácia João Falcão). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Ansiedade diurna',
     'Associação para ansiedade com componente gabaérgico e fitoterápico, sem sedação intensa.',
     E'ansiedade diurna\ncoadjuvante do sono',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula 1 a 2 vezes ao dia', 60,
     'Formulário de farmácia magistral (BioVittare). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Resistência insulínica',
     'Associação de berberina, cromo e antioxidante para suporte do controle glicêmico.',
     E'resistência insulínica\nsíndrome metabólica\nSOP com hiperinsulinemia',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 dose após o almoço e 1 após o jantar', 90,
     'Formulário de farmácia magistral (Farmacam). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Antioxidante e imunidade',
     'Associação antioxidante ampla, com cofatores e vitaminas lipossolúveis.',
     E'suporte antioxidante\nimunidade\nestresse oxidativo aumentado',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário de farmácia magistral (Invictus / Farmabotânica). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Suporte intestinal em sachê',
     'Prebióticos com glutamina para suporte de mucosa e microbiota.',
     E'permeabilidade intestinal\ndisbiose\nconstipação',
     'sachê', 'internal', 'oral', 'Veículo qsp 1 sachê', 30::numeric, 'sachês',
     '1 sachê ao dia, dissolvido em água', 60,
     'Formulário de farmácia magistral (Beleza Saúde). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Pele, cabelo e unhas',
     'Silício orgânico com colágeno, biotina e vitamina C para tecido conjuntivo e anexos.',
     E'queda de cabelo e unhas frágeis\nqualidade de pele\nsuporte de colágeno',
     'sachê', 'internal', 'oral', 'Veículo qsp 1 sachê', 30::numeric, 'sachês',
     '1 sachê ao dia', 90,
     'Formulário de farmácia magistral (Biostevi / Natuformulas). Doses do formulário; conferir antes de adotar.'),

    (uuidv7(), 'Climatério',
     'Associação fitoterápica para sintomas vasomotores e humor no climatério.',
     E'fogachos\nirritabilidade e sono no climatério',
     'cápsula', 'internal', 'oral', 'Excipiente qsp 1 cápsula', 60::numeric, 'cápsulas',
     '1 cápsula 1 a 2 vezes ao dia', 90,
     'Formulário de farmácia magistral (Pharmac / BioVittare). Doses do formulário; conferir antes de adotar.')
  ) AS v(id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
         quantity_to_dispense, quantity_unit, posology, duration, notes)
  WHERE NOT EXISTS (
    SELECT 1 FROM public.magistral_formula_templates t
    WHERE lower(t.name) = lower(v.name) AND t.deleted_at IS NULL
  )
  RETURNING id, name
)
INSERT INTO public.magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category)
SELECT uuidv7(), nova.id, c.ord, c.substance, c.qty, c.unit, 'simple'
FROM nova
JOIN (VALUES
  ('Sono completo', 0, 'Melatonina',            5::numeric,   'mg'),
  ('Sono completo', 1, '5-HTP',                 100::numeric, 'mg'),
  ('Sono completo', 2, 'Magnésio L-treonato',   300::numeric, 'mg'),
  ('Sono completo', 3, 'Glicina',               75::numeric,  'mg'),
  ('Sono completo', 4, 'L-teanina',             100::numeric, 'mg'),
  ('Sono completo', 5, 'Metilcobalamina',       100::numeric, 'mcg'),

  ('Ansiedade diurna', 0, '5-HTP',              100::numeric, 'mg'),
  ('Ansiedade diurna', 1, 'L-teanina',          100::numeric, 'mg'),
  ('Ansiedade diurna', 2, 'Magnésio quelato',   260::numeric, 'mg'),
  ('Ansiedade diurna', 3, 'Piridoxal-5-fosfato', 20::numeric, 'mg'),
  ('Ansiedade diurna', 4, 'Valeriana',          250::numeric, 'mg'),

  ('Resistência insulínica', 0, 'Berberina',            250::numeric, 'mg'),
  ('Resistência insulínica', 1, 'Ácido alfa-lipoico',   150::numeric, 'mg'),
  ('Resistência insulínica', 2, 'Picolinato de cromo',  100::numeric, 'mcg'),
  ('Resistência insulínica', 3, 'Ginseng brasileiro',   60::numeric,  'mg'),
  ('Resistência insulínica', 4, 'Gymnema silvestre',    25::numeric,  'mg'),
  ('Resistência insulínica', 5, 'Metilcobalamina',      12.5::numeric,'mcg'),

  ('Antioxidante e imunidade', 0, 'Coenzima Q10',       100::numeric, 'mg'),
  ('Antioxidante e imunidade', 1, 'Vitamina C',         100::numeric, 'mg'),
  ('Antioxidante e imunidade', 2, 'Curcumina',          100::numeric, 'mg'),
  ('Antioxidante e imunidade', 3, 'Trans-resveratrol',  100::numeric, 'mg'),
  ('Antioxidante e imunidade', 4, 'N-acetilcisteína',   200::numeric, 'mg'),
  ('Antioxidante e imunidade', 5, 'Ácido alfa-lipoico', 100::numeric, 'mg'),
  ('Antioxidante e imunidade', 6, 'Extrato de chá verde', 50::numeric,'mg'),
  ('Antioxidante e imunidade', 7, 'Vitamina D3',        2000::numeric,'UI'),

  ('Suporte intestinal em sachê', 0, 'FOS',         2::numeric, 'g'),
  ('Suporte intestinal em sachê', 1, 'Inulina',     2::numeric, 'g'),
  ('Suporte intestinal em sachê', 2, 'XOS',         2::numeric, 'g'),
  ('Suporte intestinal em sachê', 3, 'L-glutamina', 1::numeric, 'g'),

  ('Pele, cabelo e unhas', 0, 'Silício orgânico',   300::numeric, 'mg'),
  ('Pele, cabelo e unhas', 1, 'Biotina',            500::numeric, 'mcg'),
  ('Pele, cabelo e unhas', 2, 'Vitamina C',         300::numeric, 'mg'),
  ('Pele, cabelo e unhas', 3, 'Ácido hialurônico',  100::numeric, 'mg'),
  ('Pele, cabelo e unhas', 4, 'Colágeno Verisol',   2.5::numeric, 'g'),

  ('Climatério', 0, 'Isoflavona',          80::numeric,  'mg'),
  ('Climatério', 1, 'Amora',               500::numeric, 'mg'),
  ('Climatério', 2, 'Cimicifuga racemosa', 40::numeric,  'mg'),
  ('Climatério', 3, 'Vitex agnus-castus',  20::numeric,  'mg'),
  ('Climatério', 4, 'Ashwagandha',         80::numeric,  'mg'),
  ('Climatério', 5, '5-HTP',               50::numeric,  'mg')
) AS c(formula, ord, substance, qty, unit) ON c.formula = nova.name;
