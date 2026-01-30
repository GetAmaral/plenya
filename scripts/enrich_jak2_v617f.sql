-- ============================================================================
-- JAK2 V617F (c.1849G>T p.V617F) Score Item Enrichment
-- ============================================================================
-- Score Item ID: 99985045-0d9c-4454-bd86-0518136f4675
-- Item: JAK2 - pesquisa da variante genética c.1849G>T (p.V617F)
-- Generated: 2026-01-29
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- STEP 1: Update Score Item with Clinical Content
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A variante JAK2 V617F (c.1849G>T, p.V617F) é a mutação driver mais prevalente nas neoplasias mieloproliferativas (NMP), detectada em aproximadamente 97% dos casos de policitemia vera (PV), 55-60% da trombocitemia essencial (TE) e 50-60% da mielofibrose primária (MF). Esta mutação somática de ganho de função resulta na ativação constitutiva da via JAK-STAT independente de citocinas, promovendo proliferação celular desregulada da linhagem eritroblástica e granulocítica. A presença da mutação JAK2 V617F constitui critério diagnóstico maior segundo a OMS (5ª edição, 2022) e a Classificação Internacional de Consenso (ICC 2022) para NMP. A quantificação da carga alélica (variant allele frequency - VAF) possui valor prognóstico independente: cargas alélicas superiores a 50% associam-se a maior risco de transformação para mielofibrose (OR 2.8) e leucemia mieloide aguda (OR 3.2), além de aumentar significativamente o risco trombótico venoso (OR 2.1). Valores de VAF ≥30% em TE correlacionam-se com fenótipo PV-like, maior incidência de prurido aquagênico, esplenomegalia e eventos trombóticos (2.1%/paciente/ano vs 1.4% em JAK2 selvagem). A detecção é realizada preferencialmente por PCR em tempo real ou sequenciamento de nova geração (NGS), com sensibilidade de 1-3%. A pesquisa de JAK2 exon 12 está indicada quando há suspeita de PV com JAK2 V617F negativo. O monitoramento da carga alélica durante tratamento com inibidores de JAK2 (ruxolitinibe, fedratinibe) ou interferon peguilado auxilia na avaliação de resposta molecular.',

  patient_explanation = 'Este exame identifica uma alteração genética específica (mutação V617F) no gene JAK2, que está presente em quase todas as pessoas com policitemia vera (produção excessiva de glóbulos vermelhos) e em cerca de metade dos pacientes com trombocitemia essencial (excesso de plaquetas) ou mielofibrose (fibrose da medula óssea). Essas são doenças do sangue chamadas neoplasias mieloproliferativas. A mutação JAK2 V617F faz com que as células do sangue se multipliquem descontroladamente, sem precisar dos sinais normais que regulam essa produção. O resultado positivo confirma o diagnóstico da doença e ajuda o médico a escolher o melhor tratamento. Além de detectar a presença da mutação, o exame pode medir a "carga alélica" - isto é, a porcentagem de células afetadas pela mutação. Cargas alélicas mais altas (acima de 50-60%) geralmente indicam doença mais avançada e maior risco de complicações como trombose (formação de coágulos), evolução para mielofibrose ou transformação em leucemia aguda. Pacientes com carga alélica elevada necessitam monitoramento mais rigoroso e tratamento mais agressivo. O exame é feito através de amostra de sangue simples e o resultado orienta decisões importantes sobre seu tratamento, incluindo o uso de medicações específicas como inibidores de JAK2 (ruxolitinibe) ou interferon peguilado, além da estratificação do risco de eventos cardiovasculares.',

  conduct = 'INDICAÇÕES CLÍNICAS: Solicitar na investigação diagnóstica de eritrocitose (hemoglobina >16.5 g/dL em homens ou >16 g/dL em mulheres), trombocitose persistente (plaquetas >450.000/mm³), esplenomegalia inexplicada, prurido aquagênico, trombose esplâncnica (Síndrome de Budd-Chiari, trombose portal/mesentérica) e suspeita de neoplasia mieloproliferativa. INTERPRETAÇÃO: Resultado POSITIVO confirma NMP JAK2-mutada (critério diagnóstico maior OMS 2022), devendo ser correlacionado com critérios clínicos, hematológicos e histopatológicos. Quantificar carga alélica (VAF) para estratificação prognóstica: VAF <25% sugere TE; VAF 25-50% corresponde a TE ou PV; VAF >50-75% indica PV com maior risco evolutivo. Resultado NEGATIVO em paciente com suspeita de PV requer pesquisa de mutações em JAK2 exon 12 (presente em 2-3% dos casos). NEGATIVO em suspeita de TE/MF direciona investigação para mutações em CALR (20-25%) e MPL (3-5%). ESTRATIFICAÇÃO DE RISCO: Em PV, classificar como baixo risco (idade <60 anos + ausência de trombose prévia) ou alto risco (idade ≥60 anos OU história de trombose). Em TE JAK2-positiva, considerar: idade >60 anos, VAF ≥30%, hemoglobina >15 g/dL e leucocitose como fatores de risco trombótico adicional. MONITORAMENTO: Repetir dosagem de VAF semestralmente em pacientes sob terapia citorredutora ou interferon peguilado para avaliar resposta molecular. Redução >50% da VAF correlaciona-se com melhor controle clínico e menor progressão. CONDUTA TERAPÊUTICA: Alto risco ou VAF >50%: hidroxiureia (meta Ht <45%), interferon peguilado (ropeginterferona alfa-2b) ou inibidores JAK2 (ruxolitinibe 10-25 mg 2x/dia em mielofibrose ou PV refratária). Baixo risco: flebotomia (meta Ht <45%) + AAS 100 mg/dia. Considerar anticoagulação em trombose prévia. Encaminhar para hematologista para confirmação diagnóstica, classificação OMS e definição de estratégia terapêutica individualizada.',

  last_review = CURRENT_DATE

WHERE id = '99985045-0d9c-4454-bd86-0518136f4675';

-- Verify the update
DO $$
DECLARE
  v_clinical_len INT;
  v_patient_len INT;
  v_conduct_len INT;
BEGIN
  SELECT
    LENGTH(clinical_relevance),
    LENGTH(patient_explanation),
    LENGTH(conduct)
  INTO v_clinical_len, v_patient_len, v_conduct_len
  FROM score_items
  WHERE id = '99985045-0d9c-4454-bd86-0518136f4675';

  RAISE NOTICE 'Character counts - Clinical: %, Patient: %, Conduct: %',
    v_clinical_len, v_patient_len, v_conduct_len;

  IF v_clinical_len < 1500 OR v_clinical_len > 2000 THEN
    RAISE WARNING 'clinical_relevance length (%) outside target range (1500-2000)', v_clinical_len;
  END IF;

  IF v_patient_len < 1000 OR v_patient_len > 1500 THEN
    RAISE WARNING 'patient_explanation length (%) outside target range (1000-1500)', v_patient_len;
  END IF;

  IF v_conduct_len < 1500 OR v_conduct_len > 2500 THEN
    RAISE WARNING 'conduct length (%) outside target range (1500-2500)', v_conduct_len;
  END IF;
END $$;

-- ----------------------------------------------------------------------------
-- STEP 2: Insert Scientific Articles
-- ----------------------------------------------------------------------------

-- Article 1: Polycythemia vera 2024 update (Tefferi & Barbui, 2023)
INSERT INTO articles (title, authors, journal, publish_date, pm_id, article_type, abstract)
SELECT
  'Polycythemia vera: 2024 update on diagnosis, risk-stratification, and management',
  'Tefferi A, Barbui T',
  'American Journal of Hematology',
  '2023-09-01'::date,
  '37357958',
  'review',
  'Comprehensive review of PV diagnosis including JAK2 V617F detection, risk stratification based on age and thrombosis history, and current management strategies with cytoreductive therapy and JAK inhibitors'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '37357958'
);

-- Article 2: JAK2 V617F allele burden (Moliterno et al, 2023)
INSERT INTO articles (title, authors, journal, publish_date, pm_id, article_type, abstract)
SELECT
  'JAK2 V617F allele burden in polycythemia vera: burden of proof',
  'Moliterno AR, Kaizer H, Reeves BN',
  'Blood',
  '2023-04-20'::date,
  '36745865',
  'review',
  'Critical review examining the clinical significance of JAK2 V617F variant allele frequency (VAF) as prognostic marker for thrombotic risk, disease progression to myelofibrosis, and transformation to acute leukemia in polycythemia vera'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '36745865'
);

-- Article 3: New advances in JAK2 V617F mutation role (Zhang et al, 2024)
INSERT INTO articles (title, authors, journal, publish_date, pm_id, article_type, abstract)
SELECT
  'New advances in the role of JAK2 V617F mutation in myeloproliferative neoplasms',
  'Zhang Y, Zhao Y, Liu Y, Zhang M, Zhang J',
  'Cancer',
  '2024-12-15'::date,
  '39277798',
  'review',
  'Comprehensive 2024 review exploring JAK2 V617F mutation mechanisms beyond JAK-STAT pathway including regulation of cellular methylation, DNA damage accumulation, cardiovascular effects, and implications for targeted therapy with JAK inhibitors'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '39277798'
);

-- Article 4: Immune mediated diseases in JAK2 V617F MPNs (Álvarez-Reguera et al, 2024)
INSERT INTO articles (title, authors, journal, publish_date, pm_id, article_type, abstract)
SELECT
  'Features of immune mediated diseases in JAK2 (V617F)-positive myeloproliferative neoplasms and the potential therapeutic role of JAK inhibitors',
  'Álvarez-Reguera C, Prieto-Peña D, Herrero-Morant A, Sánchez-Bilbao L, Batlle-López A, Fernández-Luis S, Paz-Gandiaga N, Blanco R',
  'European Journal of Internal Medicine',
  '2024-01-01'::date,
  '38044168',
  'research_article',
  'Study examining pro-inflammatory phenotype of JAK2 V617F mutation and high prevalence of rheumatic immune-mediated diseases in MPN patients, with evaluation of therapeutic potential of JAK inhibitors for concurrent autoimmune conditions'
WHERE NOT EXISTS (
  SELECT 1 FROM articles WHERE pm_id = '38044168'
);

-- ----------------------------------------------------------------------------
-- STEP 3: Link Articles to Score Item
-- ----------------------------------------------------------------------------

-- Link Article 1 (Tefferi & Barbui, 2023)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '99985045-0d9c-4454-bd86-0518136f4675'
FROM articles a
WHERE a.pm_id = '37357958'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2 (Moliterno et al, 2023)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '99985045-0d9c-4454-bd86-0518136f4675'
FROM articles a
WHERE a.pm_id = '36745865'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3 (Zhang et al, 2024)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '99985045-0d9c-4454-bd86-0518136f4675'
FROM articles a
WHERE a.pm_id = '39277798'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4 (Álvarez-Reguera et al, 2024)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '99985045-0d9c-4454-bd86-0518136f4675'
FROM articles a
WHERE a.pm_id = '38044168'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- STEP 4: Verification
-- ----------------------------------------------------------------------------

-- Display enriched score item
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_len,
  LENGTH(patient_explanation) as patient_len,
  LENGTH(conduct) as conduct_len,
  last_review
FROM score_items
WHERE id = '99985045-0d9c-4454-bd86-0518136f4675';

-- Display linked articles
SELECT
  a.pm_id,
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '99985045-0d9c-4454-bd86-0518136f4675'
ORDER BY a.publish_date DESC;

-- Summary
DO $$
DECLARE
  v_article_count INT;
BEGIN
  SELECT COUNT(*)
  INTO v_article_count
  FROM article_score_items
  WHERE score_item_id = '99985045-0d9c-4454-bd86-0518136f4675';

  RAISE NOTICE '=================================================================';
  RAISE NOTICE 'JAK2 V617F Enrichment Summary';
  RAISE NOTICE '=================================================================';
  RAISE NOTICE 'Score Item: JAK2 - pesquisa da variante genética c.1849G>T (p.V617F)';
  RAISE NOTICE 'Articles linked: %', v_article_count;
  RAISE NOTICE 'Status: ENRICHMENT COMPLETE';
  RAISE NOTICE '=================================================================';
END $$;

COMMIT;
