-- Formas preferidas do Dr. Getúlio, ativos de marca das farmácias parceiras e fatores de correção.
--
-- FORMA PREFERIDA: quando ele precisa dessas vitaminas e minerais, usa sempre a forma ativa. O
-- catálogo passa a saber disso e sugere a troca ao montar a fórmula. A nota de cada par traz o
-- que a literatura sustenta — inclusive quando ela NÃO sustenta, que é o caso do palmitato de
-- ascorbila por via oral.
--
-- FATOR DE CORREÇÃO: das fichas técnicas. Selenometionina é diluída a 1% em maltodextrina, e o
-- magnésio bisglicinato tem 30% de mineral elementar; prescrever "300 mg" sem dizer se é elemento
-- ou insumo muda a fórmula em três vezes.

-- ---------- 1. novas substâncias (formas preferidas e ativos de marca) ----------
INSERT INTO public.magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
   brand, elemental_percent, correction_note, indications, indication_bullets, dose_reference,
   notes, source, evidence_status, last_review)
SELECT * FROM (VALUES
  (uuidv7(), 'Palmitato de ascorbila', 'ascorbil palmitato, vitamina C lipossolúvel', 'mg',
   500::numeric, 100::numeric, 1000::numeric, 0.55::numeric, 'classe', NULL, NULL, NULL,
   'Forma lipossolúvel da vitamina C. Por via oral é hidrolisada no trato digestivo antes da absorção, e o ascorbato liberado tem a mesma biodisponibilidade do ácido ascórbico; a vantagem real está na estabilidade da fórmula e no uso tópico.',
   E'fonte de vitamina C em fórmula com lipossolúveis\nuso tópico e estabilidade da formulação',
   'Dose equivalente à de vitamina C. Cerca de 43% do peso é ascorbato.',
   'Forma preferida do prescritor. Ressalva de literatura: Linus Pauling Institute descreve hidrólise pré-absortiva, sem ganho de biodisponibilidade oral sobre o ácido ascórbico.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Selenometionina', 'selenometionina 1%, L-selenometionina', 'mcg',
   100::numeric, 50::numeric, 200::numeric, 0.60::numeric, 'classe', NULL, 1::numeric,
   'Insumo diluído a 1% em maltodextrina: 100 mcg de selênio elementar equivalem a 10 mg do insumo. Aplicar fator de correção quando a dose prescrita for do elemento.',
   'Forma orgânica do selênio, mais biodisponível e mais retida em tecidos que o selenito de sódio. Antioxidante via glutationa peroxidase, com sinergia clássica com vitamina E.',
   E'reposição de selênio\nsuporte antioxidante e tireoidiano',
   '50 a 200 mcg/dia de selênio elementar (ficha técnica: até 319,75 mcg).',
   'Forma preferida do prescritor. Literatura sustenta: selenometionina é cerca de duas vezes mais biodisponível que o selenito.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'CavaQ10', 'coenzima Q10 gama-ciclodextrina, ubiquinona ciclodextrina', 'mg',
   100::numeric, 30::numeric, 250::numeric, 0.50::numeric, 'classe', 'CavaQ10®', NULL, NULL,
   'Coenzima Q10 (ubiquinona) complexada em gama-ciclodextrina, dispersível em água e mais estável que a CoQ10 comum.',
   E'suporte mitocondrial\nuso de estatina\nsaúde cardiovascular',
   '30 a 250 mg/dia (ficha técnica).',
   'Forma preferida do prescritor. Literatura: estudo humano cruzado com 22 voluntários mostra AUC e Cmax maiores para o complexo com gama-ciclodextrina; o ganho publicado é da ordem de 35% sobre CoQ10 com celulose. O "1800% mais biodisponível" da lâmina é claim do fabricante, sem correspondência na literatura indexada.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Morosil', 'extrato de laranja moro, Citrus sinensis Moro', 'mg',
   400::numeric, 400::numeric, 400::numeric, 0.45::numeric, 'classe', 'Morosil®', NULL, NULL,
   'Extrato padronizado de laranja moro em antocianinas, usado em redução de gordura corporal e circunferência abdominal, associado a dieta e exercício.',
   E'gordura visceral e abdominal\ncoadjuvante do emagrecimento',
   '400 mg/dia, dose usada nos ensaios.',
   'Literatura sustenta: RCT de 24 semanas com 180 participantes e meta-análise de 3 RCTs (2025) com redução de peso, IMC, circunferências e massa gorda.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Açafrão padronizado', 'Saffrin, Satiereal, Crocus sativus', 'mg',
   176::numeric, 88::numeric, 200::numeric, 0.45::numeric, 'classe', 'Saffrin®', NULL, NULL,
   'Extrato padronizado de Crocus sativus usado para reduzir beliscos e desejo por doces entre refeições.',
   E'compulsão alimentar e beliscos\ncontrole de apetite',
   '176 a 200 mg/dia nos ensaios; formulário usa 160 mg.',
   'Literatura sustenta parcialmente: RCT de 8 semanas com 60 mulheres com sobrepeso leve mostra menos beliscos e maior saciedade; efeito modesto e em população específica.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Cactinea', 'Opuntia ficus-indica, figo-da-índia', 'mg',
   1000::numeric, 500::numeric, 1000::numeric, 0.45::numeric, 'classe', 'Cactinea™', NULL, NULL,
   'Extrato de figo-da-índia usado com proposta diurética e de redução de retenção hídrica.',
   E'retenção hídrica\ncoadjuvante em celulite',
   '500 a 1.000 mg/dia (formulário).',
   'Evidência fraca: as alegações diuréticas vêm sobretudo de material do fabricante; a literatura indexada é escassa para o desfecho proposto.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Glycoxil', 'carcinina, peptidomimético de carnosina', 'mg',
   200::numeric, 50::numeric, 300::numeric, 0.45::numeric, 'classe', 'Glycoxil®', NULL, NULL,
   'Peptidomimético da carnosina (carcinina) com ação antiglicante e antiglicoxidante, usado em fórmulas de pele e antienvelhecimento.',
   E'glicação cutânea\nfórmulas antienvelhecimento',
   '50 a 300 mg/dia (formulário e material técnico).',
   'Evidência majoritariamente do fabricante (Exsymol), in vitro e in vivo de pequeno porte.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Bio Arct', 'Kappaphycus alvarezii, alga vermelha', 'mg',
   100::numeric, 100::numeric, 200::numeric, 0.45::numeric, 'classe', 'Bio Arct®', NULL, NULL,
   'Extrato de alga vermelha usado em fórmulas de pele com proposta de suporte energético celular e firmeza.',
   E'firmeza e qualidade de pele\nassociado a silício e colágeno',
   '100 a 200 mg/dia (formulário).',
   'Evidência escassa para uso oral: o que a busca encontra é sobretudo aplicação cosmética tópica e uso agrícola da alga.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Koubo', 'Cereus sylvestrii', 'mg',
   400::numeric, 300::numeric, 500::numeric, 0.45::numeric, 'classe', 'Koubo®', NULL, NULL,
   'Extrato de cacto usado com proposta de saciedade e controle de compulsão.',
   E'compulsão alimentar\nsaciedade',
   '300 a 500 mg/dia (formulário).',
   'Evidência escassa na literatura indexada; sustentação principal é material do fabricante.',
   'parceiro', 'suggested', now()),

  (uuidv7(), 'Green coffee', 'café verde, ácido clorogênico', 'mg',
   30::numeric, 30::numeric, 400::numeric, 0.45::numeric, 'classe', NULL, NULL, NULL,
   'Extrato de café verde, fonte de ácido clorogênico, usado em fórmulas de pele e de metabolismo.',
   E'antioxidante\ncoadjuvante metabólico',
   '30 mg em fórmulas de pele; doses metabólicas chegam a 400 mg.',
   'Evidência modesta e heterogênea para desfecho de peso.',
   'parceiro', 'suggested', now())
) AS v(id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
       brand, elemental_percent, correction_note, indications, indication_bullets, dose_reference,
       notes, source, evidence_status, last_review)
WHERE NOT EXISTS (
  SELECT 1 FROM public.magistral_components c
  WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.name))
);

-- ---------- 2. fator de correção nas substâncias que já existiam ----------
UPDATE public.magistral_components SET
  elemental_percent = 30,
  correction_note = 'Bisglicinato com cerca de 30% de magnésio elementar (confirmar no laudo do lote). 300 mg de magnésio elementar equivalem a cerca de 1 g do quelato.'
WHERE lower(name) = 'magnésio quelato' AND elemental_percent IS NULL;

UPDATE public.magistral_components SET
  correction_note = 'Mineral quelado: se a dose prescrita for do elemento, a farmácia aplica o fator do laudo do lote. Diga na fórmula se a dose é do elemento ou do insumo.'
WHERE lower(name) IN ('zinco quelato','ferro','cobre','picolinato de cromo','cálcio')
  AND correction_note IS NULL;

-- ---------- 3. formas preferidas ----------
UPDATE public.magistral_components AS c SET
  preferred_alternative_id = p.id,
  preference_note = v.nota
FROM (VALUES
  ('Vitamina C','Palmitato de ascorbila',
   'A literatura não mostra ganho de biodisponibilidade oral sobre o ácido ascórbico, mas ele estabiliza a fórmula com lipossolúveis.'),
  ('Selênio','Selenometionina',
   'A literatura sustenta: cerca de duas vezes mais biodisponível que o selenito.'),
  ('Coenzima Q10','CavaQ10',
   'O complexo com gama-ciclodextrina tem estudo humano com AUC maior; o ganho publicado é bem menor que o do material promocional.'),
  ('Cianocobalamina','Metilcobalamina',
   'A evidência favorece, sem ser conclusiva para reposição oral simples; o caso mais forte é neuropatia diabética.')
) AS v(generico, preferida, nota)
JOIN public.magistral_components p
  ON lower(public.immutable_unaccent(p.name)) = lower(public.immutable_unaccent(v.preferida))
WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.generico));
