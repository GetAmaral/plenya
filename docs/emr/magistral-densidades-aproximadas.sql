-- Densidades aparentes APROXIMADAS, para a calculadora de cápsula sair do silêncio.
--
-- A DECISÃO E O QUE ELA CUSTA: não existe tabela pública de densidade aparente por ativo, e o
-- valor real muda com lote, granulometria e compactação da farmácia (Anfarmag; Quallitá). Em vez
-- de esperar o número exato de cada insumo, entram aproximações POR CLASSE DE PÓ, declaradas como
-- tal em `density_source`. A calculadora já devolve faixa de ±25% e diz que é auxílio de
-- consultório; com estes valores ela passa a opinar, e o resultado continua sendo estimativa.
--
-- ÂNCORAS PUBLICADAS que calibram as classes:
--   · cápsula 0 (0,68 mL) comporta 400 a 450 mg de pó comum  → ~0,59 a 0,66 g/mL
--   · cápsula 00 (0,95 mL) comporta 500 a 600 mg             → ~0,53 a 0,63 g/mL
--   · a 0,70 g/mL, a cápsula 00 comporta ~665 mg             → confirma a ordem de grandeza
--   · creatina monoidratada ~0,55 g/mL · maltodextrina 0,70 · sulfato de condroitina 0,60
--
-- CLASSES USADAS (g/mL):
--   0,90  sal mineral inorgânico denso (óxido, carbonato, cloreto)
--   0,75  vitamina hidrossolúvel cristalina (C, complexo B)
--   0,65  mineral quelado, citrato, dimalato, treonato
--   0,60  aminoácido cristalino e derivado
--   0,55  fibra e polissacarídeo
--   0,50  probiótico veiculado
--   0,45  extrato seco vegetal, ativo lipossolúvel ou oleoso em pó, cristais aromáticos
--   0,35  colágeno e proteína hidrolisada
--
-- Reversível: `UPDATE magistral_components SET bulk_density=NULL, density_source=NULL
--             WHERE density_source = 'classe'` desfaz tudo o que este arquivo fez.
-- Idempotente: só preenche quem está sem densidade.

UPDATE public.magistral_components AS c
SET bulk_density = v.d, density_source = v.origem
FROM (VALUES
  -- medidas publicadas
  ('Creatina',0.55,'medida'), ('Maltodextrina',0.70,'medida'), ('Sulfato de condroitina',0.60,'medida'),

  -- minerais inorgânicos densos
  ('Cloreto de magnésio',0.90,'classe'), ('Cálcio',0.90,'classe'), ('Iodo',0.90,'classe'),

  -- minerais quelados e sais orgânicos
  ('Magnésio quelato',0.65,'classe'), ('Magnésio dimalato',0.65,'classe'),
  ('Magnésio L-treonato',0.65,'classe'), ('Magnésio taurato',0.65,'classe'),
  ('Zinco quelato',0.65,'classe'), ('Zinco carnosina',0.65,'classe'),
  ('Picolinato de cromo',0.65,'classe'), ('Selênio',0.65,'classe'),
  ('Cobre',0.65,'classe'), ('Ferro',0.65,'classe'), ('Silício orgânico',0.65,'classe'),

  -- vitaminas hidrossolúveis cristalinas
  ('Vitamina C',0.75,'classe'), ('Nicotinamida',0.75,'classe'), ('Riboflavina',0.75,'classe'),
  ('Piridoxal-5-fosfato',0.75,'classe'), ('Metilcobalamina',0.75,'classe'),
  ('Cianocobalamina',0.75,'classe'), ('Metilfolato',0.75,'classe'), ('Biotina',0.75,'classe'),
  ('Colina',0.75,'classe'),

  -- aminoácidos e derivados
  ('L-glutamina',0.60,'classe'), ('Glicina',0.60,'classe'), ('L-teanina',0.60,'classe'),
  ('L-taurina',0.60,'classe'), ('L-triptofano',0.60,'classe'), ('5-HTP',0.60,'classe'),
  ('L-carnitina',0.60,'classe'), ('Acetil-L-carnitina',0.60,'classe'),
  ('N-acetilcisteína',0.60,'classe'), ('GABA',0.60,'classe'), ('Mio-inositol',0.60,'classe'),
  ('D-quiro-inositol',0.60,'classe'), ('MSM',0.60,'classe'), ('Ácido málico',0.60,'classe'),
  ('Ácido alfa-lipoico',0.60,'classe'), ('Ácido R-alfa-lipoico',0.60,'classe'),
  ('Butirato de sódio',0.60,'classe'), ('Ácido hialurônico',0.60,'classe'),
  ('Alfa-GPC',0.60,'classe'), ('Citicolina',0.60,'classe'), ('NADH',0.60,'classe'),
  ('PQQ',0.60,'classe'), ('Melatonina',0.60,'classe'),

  -- fibras e polissacarídeos
  ('FOS',0.55,'classe'), ('Inulina',0.55,'classe'), ('XOS',0.55,'classe'),

  -- probióticos
  ('Probiótico multicepas',0.50,'classe'),

  -- extratos secos, lipossolúveis e cristais aromáticos
  ('Coenzima Q10',0.45,'classe'), ('Ubiquinol',0.45,'classe'), ('Vitamina D3',0.45,'classe'),
  ('Vitamina K2 MK-7',0.45,'classe'), ('Vitamina E',0.45,'classe'), ('Vitamina A',0.45,'classe'),
  ('Ômega-3',0.45,'classe'), ('Astaxantina',0.45,'classe'), ('Licopeno',0.45,'classe'),
  ('Curcumina',0.45,'classe'), ('Trans-resveratrol',0.45,'classe'), ('Quercetina',0.45,'classe'),
  ('Berberina',0.45,'classe'), ('Silimarina',0.45,'classe'), ('Extrato de chá verde',0.45,'classe'),
  ('Gymnema silvestre',0.45,'classe'), ('Ginseng brasileiro',0.45,'classe'),
  ('Valeriana',0.45,'classe'), ('Passiflora',0.45,'classe'), ('Amora',0.45,'classe'),
  ('Cimicifuga racemosa',0.45,'classe'), ('Vitex agnus-castus',0.45,'classe'),
  ('Ashwagandha',0.45,'classe'), ('Mucuna pruriens',0.45,'classe'), ('Yam mexicano',0.45,'classe'),
  ('Relora',0.45,'classe'), ('Tongkat ali',0.45,'classe'), ('Isoflavona',0.45,'classe'),
  ('Bacopa monnieri',0.45,'classe'), ('Ginkgo biloba',0.45,'classe'), ('Huperzina A',0.45,'classe'),
  ('Fosfatidilserina',0.45,'classe'), ('Fosfatidilcolina',0.45,'classe'),
  ('Teacrina',0.45,'classe'), ('Mentol',0.45,'classe'), ('Cânfora',0.45,'classe'),
  ('Timol',0.45,'classe'), ('Fenol',0.45,'classe'), ('Resorcina',0.45,'classe'),
  ('Talco',0.45,'classe'),

  -- proteínas hidrolisadas
  ('Colágeno Verisol',0.35,'classe'), ('Colágeno hidrolisado',0.35,'classe')
) AS v(nome, d, origem)
WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.nome))
  AND c.bulk_density IS NULL;
