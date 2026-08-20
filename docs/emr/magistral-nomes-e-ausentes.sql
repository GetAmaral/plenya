-- Nomes que as fórmulas usam e o catálogo não reconhecia, mais as substâncias que faltavam.
--
-- O achado que motivou o arquivo é um erro MEU: a regra por peso do "Sachê matinal mitocondrial"
-- foi escrita sobre "Magnésio glicina" marcado como dose do elemento — e "Magnésio glicina" não
-- casava com "Magnésio quelato" no catálogo. Sem casar, não há fator de correção: a farmácia
-- receberia 150 mg do bisglicinato (45 mg de magnésio) no lugar dos 500 mg pretendidos.
-- Idempotente.

BEGIN;

-- ---------------------------------------------------------------------------------------------
-- 1. Sinônimos que faltavam
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_components
   SET synonyms = 'magnésio bisglicinato, magnésio glicinato, magnésio glicina, bisglicinato de magnésio'
 WHERE name = 'Magnésio quelato';

UPDATE magistral_components
   SET synonyms = 'magtein, treonato de magnésio, magnésio treonato, L-treonato de magnésio'
 WHERE name = 'Magnésio L-treonato';

UPDATE magistral_components
   SET synonyms = 'cromo picolinato, cromo, cromo GTF, GTF',
       -- Picolinato de cromo tem 12,4% de cromo. As doses da literatura são do ELEMENTO, então
       -- fórmula escrita em cromo elementar precisa marcar "dose do elemento".
       elemental_percent = 12.4,
       correction_note = 'Doses da literatura são de cromo elementar; o picolinato tem 12,4% de cromo.'
 WHERE name = 'Picolinato de cromo';

UPDATE magistral_components
   SET synonyms = 'gimnema, Gymnema sylvestre, Gimnema silvestre, gurmar'
 WHERE name = 'Gymnema silvestre';

-- ---------------------------------------------------------------------------------------------
-- 2. Percentuais elementares que faltavam nas formas de magnésio (estequiometria)
-- ---------------------------------------------------------------------------------------------
UPDATE magistral_components SET elemental_percent = 15.5,
       correction_note = 'Estequiometria do malato de magnésio.'
 WHERE name = 'Magnésio dimalato' AND elemental_percent IS NULL;

UPDATE magistral_components SET elemental_percent = 8.9,
       correction_note = 'Estequiometria do taurato de magnésio.'
 WHERE name = 'Magnésio taurato' AND elemental_percent IS NULL;

UPDATE magistral_components SET elemental_percent = 8.4,
       correction_note = 'Estequiometria do aspartato de magnésio anidro.'
 WHERE name = 'Magnésio aspartato' AND elemental_percent IS NULL;

UPDATE magistral_components SET elemental_percent = 6.5,
       correction_note = 'Estequiometria do ascorbato de magnésio.'
 WHERE name = 'Magnésio ascorbato' AND elemental_percent IS NULL;

-- ---------------------------------------------------------------------------------------------
-- 3. Substâncias das fórmulas que não existiam no catálogo
--
-- Densidade entra como aproximação por classe (o mesmo critério das outras) e a faixa de dose só
-- entra quando a literatura dá uma — sem inventar, que foi o erro do ginseng brasileiro.
-- ---------------------------------------------------------------------------------------------
INSERT INTO magistral_components
    (id, name, synonyms, default_unit, usual_dose, min_dose, max_dose, dose_basis,
     bulk_density, density_source, hygroscopic, bitterness, sachet_ok, source, evidence_status,
     indications, dose_reference, is_active, created_at, updated_at)
VALUES
 (uuid_generate_v7(), 'Beta-hidroxibutirato', 'BHB, cetona exógena, sais de BHB', 'g',
  10, 3, 12, 'por_dia', 0.8000, 'classe', true, 3, true, 'pesquisa', 'suggested',
  'Corpo cetônico usado como fonte energética alternativa em protocolos cetogênicos e de desempenho cognitivo. Vem como sal de cálcio, magnésio, sódio ou potássio, e a carga desses minerais conta na fórmula.',
  'De 3 a 12 g/dia do sal, geralmente fracionados. Higroscópico e salgado: sachê é a forma prática. A carga de sódio e cálcio do sal precisa entrar na conta do dia.', true, now(), now()),

 (uuid_generate_v7(), 'Ácido hidroxicítrico', 'HCA, Citrimax, Garcinia cambogia', 'mg',
  1500, 1500, 3000, 'por_dia', 0.6000, 'classe', false, 2, true, 'pesquisa', 'suggested',
  'Extrato de Garcinia cambogia padronizado em ácido hidroxicítrico, usado para saciedade e controle de peso. As metanálises mostram efeito pequeno e inconsistente sobre o peso.',
  'De 500 a 1.000 mg três vezes ao dia, antes das refeições, padronizado a 50 a 60% de HCA. O efeito sobre o peso nas metanálises é pequeno e de qualidade baixa.', true, now(), now()),

 (uuid_generate_v7(), 'Gynostemma pentaphyllum', 'jiaogulan, ginostema, ginseng do sul', 'mg',
  300, 200, 450, 'por_dia', 0.5000, 'classe', false, 3, true, 'pesquisa', 'suggested',
  'Adaptógeno com ação sobre a AMPK, usado em resistência insulínica e esteatose. Amargo característico.',
  'De 200 a 450 mg/dia do extrato padronizado em gipenosídeos. Os ensaios em esteatose e resistência insulínica usam essa faixa por 8 a 12 semanas.', true, now(), now()),

 (uuid_generate_v7(), 'Goma cássia', 'cassia gum, goma de cássia', 'mg',
  NULL, NULL, NULL, 'por_dia', 0.7000, 'classe', true, 0, true, 'pesquisa', 'suggested',
  'Fibra solúvel usada como espessante e agente de corpo em sachês. Não tem dose terapêutica própria: a quantidade vem da textura pretendida.',
  'Sem faixa cadastrada de propósito: é excipiente de textura, não ativo com posologia estabelecida.', true, now(), now())
ON CONFLICT DO NOTHING;

COMMIT;
