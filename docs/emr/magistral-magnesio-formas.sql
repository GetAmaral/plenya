-- Formas de magnésio — tabela comparativa da Arboretum (lida por OCR) + checagem na literatura.
--
-- O QUE A TABELA RESOLVE: ela publica a EQUIVALÊNCIA ELEMENTAR de cada forma, que é o dado que
-- faltava para o fator de correção funcionar no magnésio, a família mais prescrita do catálogo.
-- Dois valores vêm explícitos da tabela:
--   · treonato: 1.500 a 2.000 mg do sal equivalem a 117 a 156 mg de magnésio elementar (~7,8%);
--   · inositol: 500 a 1.000 mg do sal equivalem a 150 a 300 mg de elementar (~30%).
-- Para as demais formas a tabela expressa a dose JÁ em magnésio elementar ("50 a 500 mg de
-- magnésio como malato"), que é a convenção de prescrição que ela adota.
--
-- ACHADO QUE MUDA DOSE: o catálogo tinha magnésio L-treonato com dose usual de 300 mg. Os ensaios
-- clínicos de cognição e sono usam 2 g/dia do SAL (Magtein), e a tabela da Arboretum diz o mesmo:
-- 1.500 a 2.000 mg. Trezentos miligramas do sal entregam cerca de 23 mg de magnésio elementar —
-- muito abaixo do que foi testado. A faixa foi corrigida, e as fórmulas que usam 300 mg passam a
-- ser sinalizadas pela própria tela como abaixo da faixa.
--
-- Percentuais estequiométricos (óxido, citrato, cloreto, sulfato, carbonato) entram anotados como
-- tal: são cálculo de fórmula molecular, não medida de lote.

-- ---------- formas que faltavam no catálogo ----------
INSERT INTO public.magistral_components
  (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, bulk_density, density_source,
   elemental_percent, correction_note, indications, indication_bullets, dose_reference, notes,
   source, evidence_status, last_review)
SELECT * FROM (VALUES
  (uuidv7(),'Magnésio citrato','citrato de magnésio','mg',300::numeric,50::numeric,500::numeric,
   0.65::numeric,'classe',16.2::numeric,
   'Citrato tem cerca de 16% de magnésio elementar (estequiometria). A tabela da Arboretum prescreve a dose já em elementar.',
   'Alta biodisponibilidade oral, com efeito laxativo em parte dos pacientes. Usado em quadros cardiovasculares, neuromusculares e nervosos.',
   E'reposição de magnésio\nconstipação, pelo efeito laxativo',
   'Suplementação: 50 a 500 mg de magnésio (como citrato) ao dia. Como laxante: 10 a 30 g do sal.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio aspartato','aspartato de magnésio','mg',300::numeric,200::numeric,400::numeric,
   0.65::numeric,'classe',NULL,NULL,
   'Forma associada ao ácido aspártico, indicada em cardiopatias, fadiga crônica e estresse físico e mental.',
   E'fadiga crônica\ncardiopatias\nestresse físico e mental',
   '200 a 400 mg de magnésio (como aspartato) ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio ascorbato','ascorbato de magnésio','mg',300::numeric,50::numeric,500::numeric,
   0.65::numeric,'classe',NULL,NULL,
   'Une magnésio e vitamina C na mesma molécula, com uso voltado a suporte imune e cardíaco.',
   E'suporte imune\nsuporte cardiovascular',
   '50 a 500 mg de magnésio (como ascorbato) ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio inositol','magnésio com inositol','mg',750::numeric,500::numeric,1000::numeric,
   0.60::numeric,'classe',30::numeric,
   'A tabela publica a equivalência: 500 a 1.000 mg do sal correspondem a 150 a 300 mg de magnésio elementar.',
   'Une o magnésio quelado ao inositol, com uso em relaxamento mental e corporal, qualidade do sono, memória e dores musculares.',
   E'sono e relaxamento\nmemória\ndores musculares e ósseas',
   '500 a 1.000 mg do sal ao dia (150 a 300 mg de magnésio elementar).',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio óxido','óxido de magnésio','mg',600::numeric,400::numeric,800::numeric,
   0.90::numeric,'classe',60.3::numeric,
   'Óxido tem cerca de 60% de magnésio elementar (estequiometria), mas é a forma menos absorvida.',
   'Forma menos absorvida, usada sobretudo pelo efeito laxativo.',
   E'constipação\nreposição de baixo custo, com absorção limitada',
   '400 a 800 mg do sal ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio carbonato','carbonato de magnésio','mg',500::numeric,250::numeric,1000::numeric,
   0.90::numeric,'classe',28.8::numeric,
   'Carbonato tem cerca de 29% de magnésio elementar (estequiometria).',
   'Propriedade antiácida, também usado pelo efeito laxativo em doses maiores.',
   E'antiácido\nlaxativo em dose maior',
   'Antiácido: 250 a 1.000 mg do sal. Laxante: 2 a 5 g do sal ao dia.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now()),

  (uuidv7(),'Magnésio sulfato','sulfato de magnésio, sal amargo','mg',NULL::numeric,NULL::numeric,NULL::numeric,
   0.90::numeric,'classe',9.9::numeric,
   'Sulfato heptaidratado tem cerca de 10% de magnésio elementar (estequiometria).',
   'Boa absorção tópica; por via oral tem efeito laxativo.',
   E'uso tópico transdérmico\nlaxativo por via oral',
   'Oral: 5 a 10 g do sal ao dia. Tópico: 50 a 200 mg do sal em creme transdérmico.',
   'Tabela comparativa Arboretum.','parceiro','suggested',now())
) AS v(id,name,synonyms,default_unit,usual_dose,min_dose,max_dose,bulk_density,density_source,
       elemental_percent,correction_note,indications,indication_bullets,dose_reference,notes,
       source,evidence_status,last_review)
WHERE NOT EXISTS (
  SELECT 1 FROM public.magistral_components c
  WHERE lower(public.immutable_unaccent(c.name)) = lower(public.immutable_unaccent(v.name))
);

-- ---------- correções nas formas que já existiam ----------
-- Treonato: a faixa estava uma ordem de grandeza abaixo do que os ensaios usam.
UPDATE public.magistral_components SET
  usual_dose = 2000, min_dose = 1500, max_dose = 2000,
  elemental_percent = 7.8,
  correction_note = 'Treonato tem cerca de 7,8% de magnésio elementar: 2 g do sal equivalem a cerca de 156 mg de elementar (tabela Arboretum).',
  dose_reference = '1.500 a 2.000 mg do sal ao dia, em duas tomadas. Os ensaios de cognição e sono com Magtein usam 2 g/dia; a tabela da Arboretum diz o mesmo. Doses de 200 a 500 mg do sal, comuns em formulários, entregam menos de 40 mg de magnésio elementar.',
  notes = coalesce(notes,'') || ' Faixa corrigida a partir da tabela comparativa Arboretum e dos ensaios com Magtein (2 g/dia).',
  last_review = now()
WHERE lower(name) = 'magnésio l-treonato';

UPDATE public.magistral_components SET
  elemental_percent = 30,
  correction_note = 'Bisglicinato com cerca de 30% de magnésio elementar (confirmar no laudo do lote). A tabela da Arboretum prescreve a dose já em elementar: 50 a 500 mg/dia.'
WHERE lower(name) = 'magnésio quelato';

UPDATE public.magistral_components SET
  correction_note = 'A tabela da Arboretum prescreve a dose já em magnésio elementar: 50 a 500 mg/dia.',
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: 50 a 500 mg de magnésio (como malato) ao dia.'
WHERE lower(name) = 'magnésio dimalato' AND correction_note IS NULL;

UPDATE public.magistral_components SET
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: 50 a 500 mg de magnésio (como taurato) ao dia, com uso voltado à saúde cardiovascular e ao diabetes tipo 2.',
  usual_dose = coalesce(usual_dose, 300), min_dose = coalesce(min_dose, 50), max_dose = coalesce(max_dose, 500)
WHERE lower(name) = 'magnésio taurato';

UPDATE public.magistral_components SET
  elemental_percent = 11.9,
  correction_note = 'Cloreto hexaidratado tem cerca de 12% de magnésio elementar (estequiometria). A tabela indica sobretudo uso tópico, a 15%.',
  dose_reference = coalesce(dose_reference,'') || ' Tabela Arboretum: melhor absorção tópica, 15% em formulação de uso externo.'
WHERE lower(name) = 'cloreto de magnésio';
