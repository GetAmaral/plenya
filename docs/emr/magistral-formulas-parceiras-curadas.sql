-- Fórmulas do formulário das farmácias parceiras, CURADAS antes de entrar.
--
-- O formulário traz 80 fórmulas. Não entram todas: entram as que cobrem frentes que o Dr. Getúlio
-- atende e que sobrevivem à conferência. Três coisas foram feitas em cada uma:
--
--   1. FORMAS TROCADAS pelas que ele usa: vitamina C → palmitato de ascorbila, B12 →
--      metilcobalamina, B6 → piridoxal-5-fosfato, folato → metilfolato, selênio → selenometionina,
--      coenzima Q10 → CavaQ10.
--   2. ERROS DE UNIDADE CORRIGIDOS. O formulário traz "SELÊNIO 30mg" na Anti-Ox Maxi Ultra (é
--      mcg: 30 mg seriam ~100 vezes o teto de suplemento e dose tóxica) e "VIT A 50mg" na de
--      cãimbras (50 mg de retinol seriam ~166.000 UI).
--   3. DUPLICIDADE E TETO DE B6. A fórmula "Energizante" soma piridoxal-5-fosfato 7 mg com
--      vitamina B6 100 mg — a mesma vitamina duas vezes, ~107 mg/dia, acima do teto da IN 28
--      (98,6 mg) e na faixa associada a neuropatia sensitiva em uso crônico. Aqui vai só P5P 25 mg.
--
-- Cada fórmula guarda em `notes` o que foi alterado e por quê. Todas entram para conferência.

WITH nova AS (
  INSERT INTO public.magistral_formula_templates
    (id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
     quantity_to_dispense, quantity_unit, posology, duration, notes)
  SELECT * FROM (VALUES
    (uuidv7(), 'Cortisol elevado e estresse',
     'Adaptógenos com L-teanina para modulação do eixo HPA em estresse com cortisol elevado.',
     E'estresse com cortisol alto\nansiedade diurna\ncoadjuvante do sono',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',60::numeric,'cápsulas',
     '1 cápsula 2 vezes ao dia', 60,
     'Formulário das farmácias parceiras, sem alteração: as três substâncias e as doses são as do formulário.'),

    (uuidv7(), 'Concentração e memória',
     'Suporte cognitivo com adaptógeno, colina e coenzima Q10 para foco e memória.',
     E'queixa de foco e memória\nfadiga mental',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula pela manhã, após o café', 60,
     'Formulário das parceiras, com coenzima Q10 trocada por CavaQ10.'),

    (uuidv7(), 'Esteatose e resistência insulínica',
     'Morosil com ácido alfa-lipoico e cofatores para esteatose hepática, resistência insulínica e dislipidemia.',
     E'esteatose hepática\nresistência insulínica\ndislipidemia',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',60::numeric,'cápsulas',
     '1 cápsula 2 vezes ao dia', 90,
     'Formulário das parceiras. Vitamina C trocada por palmitato de ascorbila e B6 por P5P. Morosil mantido em 200 mg 2x (400 mg/dia, que é a dose dos ensaios).'),

    (uuidv7(), 'Antioxidante amplo',
     'Associação antioxidante ampla com cofatores minerais, polifenóis e suporte mitocondrial.',
     E'estresse oxidativo\nsuporte antienvelhecimento',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário das parceiras (Anti-Ox Maxi Ultra), com três correções: selênio de 30 mg para 30 mcg de selenometionina (o formulário traz mg, que seria dose tóxica), vitamina C para palmitato de ascorbila e CoQ10 para CavaQ10.'),

    (uuidv7(), 'Hipotireoidismo, suporte de cofatores',
     'Cofatores da síntese e conversão de hormônio tireoidiano: iodo, selênio, zinco, tirosina e vitamina A.',
     E'hipotireoidismo, suporte nutricional\nconversão de T4 em T3',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 90,
     'Formulário das parceiras, com selênio quelato trocado por selenometionina e vitamina C por palmitato de ascorbila. Iodo mantido em 100 mcg: em tireoidite autoimune, avaliar antes de repor.'),

    (uuidv7(), 'Insônia fitoterápica',
     'Associação fitoterápica para indução do sono, sem melatonina.',
     E'insônia de indução\nquando se prefere evitar melatonina',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 dose antes de dormir', 30,
     'Formulário das parceiras, sem alteração.'),

    (uuidv7(), 'Fadiga pós atividade física',
     'Suporte mitocondrial e de recuperação com carnitina, D-ribose, aminoácidos e cofatores.',
     E'fadiga após esforço\nrecuperação muscular',
     'cápsula','internal','oral','Excipiente qsp 1 cápsula',30::numeric,'cápsulas',
     '1 cápsula ao dia', 60,
     'Formulário das parceiras, enxugada: das 19 substâncias originais ficaram as 10 com dose relevante. B6 virou P5P, B12 virou metilcobalamina, selênio virou selenometionina, CoQ10 virou CavaQ10 e a dose de CoQ10 subiu de 10 para 50 mg, porque 10 mg não tem efeito descrito.')
  ) AS v(id, name, indication, indication_bullets, pharmaceutical_form, usage_type, route, vehicle,
         quantity_to_dispense, quantity_unit, posology, duration, notes)
  WHERE NOT EXISTS (
    SELECT 1 FROM public.magistral_formula_templates t
    WHERE lower(t.name) = lower(v.name) AND t.deleted_at IS NULL
  )
  RETURNING id, name
)
INSERT INTO public.magistral_formula_template_components
  (id, template_id, display_order, substance, quantity, unit, category, note, as_elemental)
SELECT uuidv7(), nova.id, c.ord, c.substance, c.qty, c.unit, 'simple', c.nota, c.elem
FROM nova JOIN (VALUES
  ('Cortisol elevado e estresse',0,'Relora',150::numeric,'mg','',false),
  ('Cortisol elevado e estresse',1,'L-teanina',200::numeric,'mg','',false),
  ('Cortisol elevado e estresse',2,'Ashwagandha',200::numeric,'mg','',false),

  ('Concentração e memória',0,'Teacrina',150::numeric,'mg','',false),
  ('Concentração e memória',1,'Rhodiola rosea',150::numeric,'mg','',false),
  ('Concentração e memória',2,'Fosfatidilcolina',100::numeric,'mg','',false),
  ('Concentração e memória',3,'CavaQ10',50::numeric,'mg','troca do prescritor',false),

  ('Esteatose e resistência insulínica',0,'Morosil',200::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',1,'Ácido alfa-lipoico',75::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',2,'Piridoxal-5-fosfato',3::numeric,'mg','',false),
  ('Esteatose e resistência insulínica',3,'Palmitato de ascorbila',60::numeric,'mg','troca do prescritor',false),

  ('Antioxidante amplo',0,'Ácido alfa-lipoico',50::numeric,'mg','',false),
  ('Antioxidante amplo',1,'CavaQ10',40::numeric,'mg','troca do prescritor',false),
  ('Antioxidante amplo',2,'N-acetilcisteína',100::numeric,'mg','no lugar da cisteína do formulário',false),
  ('Antioxidante amplo',3,'Trans-resveratrol',5::numeric,'mg','',false),
  ('Antioxidante amplo',4,'Licopeno',5::numeric,'mg','',false),
  ('Antioxidante amplo',5,'PQQ',10::numeric,'mg','',false),
  ('Antioxidante amplo',6,'Selenometionina',30::numeric,'mcg','corrigido: o formulário trazia 30 mg',true),
  ('Antioxidante amplo',7,'Zinco quelato',15::numeric,'mg','',true),
  ('Antioxidante amplo',8,'Cobre',1::numeric,'mg','',true),
  ('Antioxidante amplo',9,'Riboflavina',10::numeric,'mg','',false),
  ('Antioxidante amplo',10,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false),
  ('Antioxidante amplo',11,'Vitamina E',45::numeric,'UI','',false),

  ('Hipotireoidismo, suporte de cofatores',0,'Iodo',100::numeric,'mcg','avaliar antes em tireoidite autoimune',true),
  ('Hipotireoidismo, suporte de cofatores',1,'Selenometionina',100::numeric,'mcg','',true),
  ('Hipotireoidismo, suporte de cofatores',2,'Zinco quelato',15::numeric,'mg','',true),
  ('Hipotireoidismo, suporte de cofatores',3,'L-tirosina',100::numeric,'mg','',false),
  ('Hipotireoidismo, suporte de cofatores',4,'Vitamina A',1000::numeric,'UI','',false),
  ('Hipotireoidismo, suporte de cofatores',5,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false),
  ('Hipotireoidismo, suporte de cofatores',6,'Vitamina D3',50::numeric,'UI','dose simbólica no formulário; ajustar pelo exame',false),

  ('Insônia fitoterápica',0,'Passiflora',120::numeric,'mg','',false),
  ('Insônia fitoterápica',1,'Valeriana',50::numeric,'mg','',false),
  ('Insônia fitoterápica',2,'5-HTP',75::numeric,'mg','griffonia do formulário',false),

  ('Fadiga pós atividade física',0,'L-carnitina',500::numeric,'mg','',false),
  ('Fadiga pós atividade física',1,'CavaQ10',50::numeric,'mg','formulário trazia 10 mg de CoQ10',false),
  ('Fadiga pós atividade física',2,'NADH',5::numeric,'mg','',false),
  ('Fadiga pós atividade física',3,'L-taurina',100::numeric,'mg','',false),
  ('Fadiga pós atividade física',4,'Magnésio quelato',100::numeric,'mg','',true),
  ('Fadiga pós atividade física',5,'Zinco quelato',15::numeric,'mg','',true),
  ('Fadiga pós atividade física',6,'Selenometionina',30::numeric,'mcg','',true),
  ('Fadiga pós atividade física',7,'Piridoxal-5-fosfato',10::numeric,'mg','no lugar de B6',false),
  ('Fadiga pós atividade física',8,'Metilcobalamina',100::numeric,'mcg','no lugar de B12',false),
  ('Fadiga pós atividade física',9,'Palmitato de ascorbila',100::numeric,'mg','troca do prescritor',false)
) AS c(formula, ord, substance, qty, unit, nota, elem) ON c.formula = nova.name;
