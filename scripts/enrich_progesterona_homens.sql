-- ============================================================================
-- ENRICHMENT: Progesterona - Homens
-- Score Item ID: 5e465089-1492-44c4-9e7b-6696731a56a3
-- Generated: 2026-01-29
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. UPDATE SCORE_ITEMS: Add clinical content
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A progesterona em homens, embora frequentemente negligenciada na prática clínica, desempenha papéis fisiológicos fundamentais. Produzida principalmente pelas células de Leydig nos testículos e em menor quantidade pelas glândulas adrenais, a progesterona serve como precursor essencial na biossíntese de testosterona através da via esteroidogênica. Valores de referência normais variam de 0,13 a 0,97 ng/mL (0,4-3,1 nmol/L), sem variações significativas relacionadas à idade. A dosagem de progesterona em homens é clinicamente relevante em múltiplos contextos: investigação de infertilidade masculina (influencia espermiogênese, capacitação espermática e reação acrossômica), rastreamento e monitoramento de hiperplasia adrenal congênita (HAC) especialmente deficiência de 21-hidroxilase, avaliação de puberdade precoce (testículos pequenos com pênis aumentado), diagnóstico diferencial de tumores adrenais secretores de esteroides, e investigação de distúrbios do desenvolvimento sexual (DSD) ovotesticular. Níveis elevados (>1,0 ng/mL) associam-se a HAC, tumores adrenocorticais, resistência periférica aos andrógenos, e curiosamente a pré-diabetes e diabetes tipo 2 em homens. Níveis muito baixos raramente têm significado clínico isolado, mas podem indicar insuficiência adrenal primária ou hipogonadismo hipogonadotrófico severo. A dosagem de 17-hidroxiprogesterona (17-OHP) é frequentemente mais útil que progesterona isolada para diagnóstico de HAC, com valores basais >10.000 ng/dL praticamente confirmando HAC clássica. Efeitos extrarreproductivos da progesterona incluem neuroproteção via metabólitos 5α-reduzidos (neuroesteroides), modulação imunológica, efeitos cardiovasculares e renais, bloqueio da secreção de gonadotrofinas, e melhora da qualidade do sono.',

  patient_explanation = 'A progesterona não é apenas um "hormônio feminino" - os homens também produzem e precisam dela para funcionar adequadamente. Nos homens, a progesterona é fabricada principalmente nos testículos e serve como matéria-prima para produzir testosterona, o principal hormônio masculino. Seus níveis normais são baixos (0,13-0,97 ng/mL) e permanecem estáveis durante toda a vida adulta. Este exame pode ser solicitado quando há suspeita de problemas de fertilidade (a progesterona ajuda os espermatozoides a amadurecerem e funcionarem corretamente), quando meninos apresentam sinais de puberdade muito cedo, ou para investigar alterações nas glândulas adrenais. Valores elevados podem indicar hiperplasia adrenal congênita (uma condição genética onde as glândulas adrenais produzem hormônios de forma desequilibrada), tumores nas glândulas adrenais, ou curiosamente, podem estar associados a risco aumentado de pré-diabetes. A progesterona também tem efeitos importantes no cérebro (ajuda a proteger os neurônios e melhora o sono), no sistema imunológico e cardiovascular. Valores muito baixos geralmente não causam preocupação isoladamente. Se o resultado estiver alterado, seu médico pode solicitar outros exames hormonais complementares, especialmente 17-hidroxiprogesterona, cortisol, testosterona, LH e FSH, para entender melhor o que está acontecendo. O tratamento, quando necessário, depende da causa específica identificada.',

  conduct = 'VALORES ELEVADOS (>1,0 ng/mL): Investigar hiperplasia adrenal congênita - solicitar 17-hidroxiprogesterona basal (se >2.000 ng/dL, considerar HAC não clássica; se >10.000 ng/dL, sugere HAC clássica), cortisol sérico, ACTH, androstenediona, DHEA-S e testosterona total. Considerar teste de estímulo com ACTH sintético (250 mcg) se 17-OHP basal entre 200-1.000 ng/dL para distinguir HAC não clássica de portadores heterozigotos. Avaliar história clínica: puberdade precoce na infância, baixa estatura com idade óssea avançada, infertilidade, acne severa. Realizar ultrassom ou TC de adrenais para descartar tumores adrenocorticais produtores de esteroides. Se HAC confirmada, encaminhar para endocrinologista - tratamento com glicocorticoides (hidrocortisona 15-25 mg/dia dividida em 2-3 doses, ou prednisona 5-7,5 mg/dia) para suprimir ACTH e normalizar andrógenos. Monitorar regularmente 17-OHP, androstenediona e testosterona para ajuste de dose. Em homens com valores >1,4 ng/mL, rastrear glicemia de jejum, hemoglobina glicada e TOTG para avaliar risco metabólico (associação com pré-diabetes/diabetes tipo 2). Investigar função testicular se infertilidade: espermograma, FSH, LH, testosterona total e livre, prolactina. Para tumores adrenais, avaliar necessidade de cirurgia com cirurgião endócrino. VALORES NORMAIS (0,13-0,97 ng/mL): Tranquilizar paciente. Se contexto de investigação de infertilidade, prosseguir com avaliação andrológica completa. VALORES MUITO BAIXOS (<0,13 ng/mL): Raramente significativos isoladamente. Avaliar função adrenal (cortisol, ACTH) e gonadal (testosterona, LH, FSH) se sintomas compatíveis com insuficiência hormonal. CONSIDERAÇÕES ESPECIAIS: Em neonatos masculinos, valores de referência são diferentes (>630 ng/dL até 28 dias, caindo para <110 ng/dL aos 6 meses). Screening neonatal para HAC é mandatório em vários países. Coletar amostra pela manhã (8h) para comparabilidade. Interferências: uso de medroxiprogesterona, espironolactona, cetoconazol podem elevar níveis. Estresse agudo pode elevar cortisol e precursores. Sempre interpretar em contexto clínico e com outros hormônios da via esteroidogênica.',

  updated_at = CURRENT_TIMESTAMP
WHERE id = '5e465089-1492-44c4-9e7b-6696731a56a3';

-- Verify character counts
DO $$
DECLARE
  v_clinical_relevance TEXT;
  v_patient_explanation TEXT;
  v_conduct TEXT;
  v_len_clinical INTEGER;
  v_len_patient INTEGER;
  v_len_conduct INTEGER;
BEGIN
  SELECT clinical_relevance, patient_explanation, conduct
  INTO v_clinical_relevance, v_patient_explanation, v_conduct
  FROM score_items
  WHERE id = '5e465089-1492-44c4-9e7b-6696731a56a3';

  v_len_clinical := LENGTH(v_clinical_relevance);
  v_len_patient := LENGTH(v_patient_explanation);
  v_len_conduct := LENGTH(v_conduct);

  RAISE NOTICE '=== CHARACTER COUNT VALIDATION ===';
  RAISE NOTICE 'clinical_relevance: % chars (target: 1500-2000)', v_len_clinical;
  RAISE NOTICE 'patient_explanation: % chars (target: 1000-1500)', v_len_patient;
  RAISE NOTICE 'conduct: % chars (target: 1500-2500)', v_len_conduct;

  IF v_len_clinical < 1500 OR v_len_clinical > 2000 THEN
    RAISE WARNING 'clinical_relevance length outside target range!';
  END IF;

  IF v_len_patient < 1000 OR v_len_patient > 1500 THEN
    RAISE WARNING 'patient_explanation length outside target range!';
  END IF;

  IF v_len_conduct < 1500 OR v_len_conduct > 2500 THEN
    RAISE WARNING 'conduct length outside target range!';
  END IF;
END $$;

-- ----------------------------------------------------------------------------
-- 2. INSERT ARTICLES: Add scientific references
-- ----------------------------------------------------------------------------

-- Article 1: Oettel & Mukhopadhyay 2004 - Progesterone: the forgotten hormone in men
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type, created_at, updated_at
) VALUES (
  gen_random_uuid(),
  'Progesterone: the forgotten hormone in men?',
  'Oettel M, Mukhopadhyay AK',
  'Aging Male',
  '2004-09-01'::date,
  '15669543',
  '10.1080/13685530400004199',
  'review',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date,
  pm_id = EXCLUDED.pm_id,
  article_type = EXCLUDED.article_type,
  updated_at = CURRENT_TIMESTAMP
RETURNING id, title, pm_id;

-- Link Article 1 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '5e465089-1492-44c4-9e7b-6696731a56a3'
FROM articles a
WHERE a.pm_id = '15669543'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 2: San Martín et al 2021 - Classical CAH in adult males
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type, created_at, updated_at
) VALUES (
  gen_random_uuid(),
  'Classical congenital adrenal hyperplasia due to 21-hydroxylase deficiency (21-OHD) in adult males: Clinical presentation, hormone function and the detection of adrenal and testicular adrenal rest tumors (TARTs)',
  'San Martín P, Eugenio Russmann ML, Mendeluk G, Fierro MF, Marino R, Pardes E',
  'Endocrinol Diabetes Nutr (Engl Ed)',
  '2021-04-01'::date,
  '34266634',
  '10.1016/j.endien.2020.07.003',
  'research_article',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date,
  pm_id = EXCLUDED.pm_id,
  article_type = EXCLUDED.article_type,
  updated_at = CURRENT_TIMESTAMP
RETURNING id, title, pm_id;

-- Link Article 2 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '5e465089-1492-44c4-9e7b-6696731a56a3'
FROM articles a
WHERE a.pm_id = '34266634'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 3: Lei et al 2025 - LH regulation of testosterone
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type, created_at, updated_at
) VALUES (
  gen_random_uuid(),
  'Luteinizing Hormone Regulates Testosterone Production, Leydig Cell Proliferation, Differentiation, and Circadian Rhythm During Spermatogenesis',
  'Lei T, Yang Y, Yang WX',
  'International Journal of Molecular Sciences',
  '2025-04-10'::date,
  '40332028',
  '10.3390/ijms26083548',
  'review',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date,
  pm_id = EXCLUDED.pm_id,
  article_type = EXCLUDED.article_type,
  updated_at = CURRENT_TIMESTAMP
RETURNING id, title, pm_id;

-- Link Article 3 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '5e465089-1492-44c4-9e7b-6696731a56a3'
FROM articles a
WHERE a.pm_id = '40332028'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 4: Panken et al 2025 - Testosterone to Estradiol Ratios
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type, created_at, updated_at
) VALUES (
  gen_random_uuid(),
  'Testosterone to Estradiol Ratios in Fertile and Subfertile Men: A Large Cohort Analysis',
  'Panken EJ, Hayon S, Greenberg DR, Kumar SK, Brannigan RE, Halpern JA',
  'Urology',
  '2025-01-01'::date,
  '39542362',
  '10.1016/j.urology.2024.11.004',
  'research_article',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date,
  pm_id = EXCLUDED.pm_id,
  article_type = EXCLUDED.article_type,
  updated_at = CURRENT_TIMESTAMP
RETURNING id, title, pm_id;

-- Link Article 4 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '5e465089-1492-44c4-9e7b-6696731a56a3'
FROM articles a
WHERE a.pm_id = '39542362'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- 3. VERIFICATION: Check all links
-- ----------------------------------------------------------------------------

SELECT
  si.id,
  si.name,
  si.code,
  LENGTH(si.clinical_relevance) as clinical_len,
  LENGTH(si.patient_explanation) as patient_len,
  LENGTH(si.conduct) as conduct_len,
  COUNT(asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '5e465089-1492-44c4-9e7b-6696731a56a3'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct;

-- List all linked articles
SELECT
  a.pm_id,
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '5e465089-1492-44c4-9e7b-6696731a56a3'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- ENRICHMENT COMPLETE
-- ============================================================================
-- Score Item: Progesterona - Homens
-- Articles Added: 4
-- Clinical Content: Complete (PT-BR)
-- Character Counts: Validated
-- ============================================================================
