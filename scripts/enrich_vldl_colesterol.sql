-- ============================================================================
-- ENRICHMENT: VLDL Colesterol (ID: 0ed8b90b-8679-4842-a7b2-68b0e3ea9154)
-- Data: 2026-01-29
-- Artigos: 4 artigos peer-reviewed (2023-2024)
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. UPDATE CLINICAL CONTENT
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
    clinical_relevance = 'O VLDL-colesterol (Very Low-Density Lipoprotein) é uma lipoproteína sintetizada no fígado que transporta triglicerídeos e colesterol endógeno para os tecidos periféricos. Representa aproximadamente 15-20% do colesterol total e é calculado pela fórmula de Friedewald (VLDL-C = TG/5 em mg/dL). Estudos recentes demonstram que o VLDL-C é aterogênico e contribui para o risco cardiovascular residual, especialmente quando avaliado como parte do colesterol não-HDL. As diretrizes de 2024 da Sociedade Polonesa de Lipidologia estabelecem que níveis elevados de VLDL-C (>30 mg/dL) indicam superprodução hepática de lipoproteínas ricas em triglicerídeos, característica da síndrome metabólica e diabetes tipo 2. O colesterol remanescente (presente em partículas VLDL e IDL) mostrou-se superior ao LDL-C na predição de eventos cardiovasculares em estudos de 2023, com associação causal demonstrada por análises de randomização mendeliana. A razão TG/HDL-C >3,5 correlaciona-se fortemente com excesso de VLDL-C e resistência insulínica. Em pacientes com doença cardiovascular estabelecida, o VLDL-C elevado (quartil superior >38 mg/dL) associou-se a risco 49% maior de eventos adversos em membros inferiores (HR 1.49, IC95% 1.16-1.93), independente do LDL-C. Valores <30 mg/dL indicam metabolismo lipídico adequado; 30-40 mg/dL sugerem dislipidemia aterogênica leve; >40 mg/dL requerem intervenção terapêutica intensiva com modificação de estilo de vida e possível farmacoterapia.',

    patient_explanation = 'O VLDL-colesterol é um tipo de gordura produzida pelo fígado que transporta triglicerídeos para seus músculos e tecidos adiposos. Diferente do LDL (colesterol "ruim" mais conhecido), o VLDL carrega principalmente triglicerídeos, mas também contribui para o acúmulo de placas nas artérias. Valores normais ficam abaixo de 30 mg/dL - quanto mais baixo, melhor. Níveis elevados (>30 mg/dL) geralmente indicam que você está comendo muitos carboidratos refinados, açúcares ou álcool, e seu fígado está produzindo excesso de gordura. Isso é comum em pessoas com síndrome metabólica, pré-diabetes ou diabetes tipo 2. O VLDL alto aumenta seu risco de entupimento das artérias (aterosclerose), especialmente nas pernas, podendo causar dor ao caminhar. Para baixar o VLDL, reduza carboidratos simples (pães brancos, doces, refrigerantes), aumente fibras (vegetais, leguminosas), pratique exercícios aeróbicos regularmente (150 min/semana), perca peso se estiver acima do ideal e limite álcool a 1-2 doses/dia. Se seus triglicerídeos estão muito altos (>200 mg/dL), seu médico pode calcular o VLDL dividindo os triglicerídeos por 5. Valores persistentemente elevados mesmo com mudanças de hábitos podem necessitar de medicamentos como fibratos ou ômega-3 em altas doses, sempre sob orientação médica.',

    conduct = 'ESTRATIFICAÇÃO DE RISCO: Calcular VLDL-C pela fórmula de Friedewald (TG/5) em jejum de 12h. Se TG >400 mg/dL, a fórmula perde acurácia - solicitar dosagem direta ou usar fórmula de Martin-Hopkins. Avaliar colesterol não-HDL (CT - HDL-C) como marcador integrado: meta <130 mg/dL para risco moderado, <100 mg/dL para alto risco, <80 mg/dL para muito alto risco. Calcular colesterol remanescente (CT - LDL-C - HDL-C): valores >38 mg/dL indicam risco aumentado de eventos cardiovasculares. VLDL-C <30 mg/dL: Metabolismo lipídico adequado. Manter dieta mediterrânea, exercícios aeróbicos 150 min/semana, IMC <25 kg/m². Reavaliar lipidograma anualmente. VLDL-C 30-40 mg/dL: Dislipidemia aterogênica leve. Intervenção intensiva no estilo de vida: reduzir carboidratos refinados <50g/dia, aumentar fibras solúveis >25g/dia, ômega-3 de peixes gordos 2-3x/semana. Exercícios aeróbicos moderados-vigorosos 200-300 min/semana. Perda ponderal de 5-10% se IMC >25. Limitar álcool a 14g/dia (1 dose). Reavaliar em 3 meses. VLDL-C >40 mg/dL: Hipertrigliceridemia significativa. Investigar causas secundárias: diabetes descompensado, hipotireoidismo, síndrome nefrótica, medicamentos (corticoides, betabloqueadores, diuréticos tiazídicos). Iniciar terapia nutricional intensiva (nutricionista) + exercícios supervisionados. Se TG >500 mg/dL, risco de pancreatite aguda - considerar fenofibrato 145-200 mg/dia para redução rápida. Para TG 200-499 mg/dL com risco cardiovascular elevado, associar ômega-3 EPA 2-4g/dia (icosapento etil) à estatina otimizada. Monitorar enzimas hepáticas (TGO/TGP) e CPK a cada 3 meses. Meta terapêutica: TG <150 mg/dL (idealmente <100 mg/dL), não-HDL-C <100 mg/dL, VLDL-C <30 mg/dL. Pacientes com doença arterial periférica e VLDL-C elevado requerem seguimento vascular semestral (ITB, doppler). Reavaliar lipidograma em 6-8 semanas após início de farmacoterapia.',

    last_review = CURRENT_DATE
WHERE id = '0ed8b90b-8679-4842-a7b2-68b0e3ea9154';

-- Verificar atualização
SELECT
    name,
    LENGTH(clinical_relevance) as relevance_chars,
    LENGTH(patient_explanation) as explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = '0ed8b90b-8679-4842-a7b2-68b0e3ea9154';

-- ----------------------------------------------------------------------------
-- 2. INSERT SCIENTIFIC ARTICLES
-- ----------------------------------------------------------------------------

-- Article 1: Remnant cholesterol and metabolic syndrome (2023)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    -- Check if article already exists
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '37045908';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            pm_id,
            article_type
        ) VALUES (
            gen_random_uuid(),
            'Remnant cholesterol can identify individuals at higher risk of metabolic syndrome in the general population',
            'Zhang Y, Jin JL, Cao YX, et al.',
            'Scientific Reports',
            '2023-04-12'::date,
            '37045908',
            'research_article'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 1 inserted: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 1 already exists: %', v_article_id;
    END IF;
END $$;

-- Article 2: TG/HDL-C ratio and cardiovascular disease (2023)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '36900073';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            pm_id,
            article_type
        ) VALUES (
            gen_random_uuid(),
            'The Triglyceride/High-Density Lipoprotein Cholesterol (TG/HDL-C) Ratio as a Risk Marker for Metabolic Syndrome and Cardiovascular Disease',
            'Karimi Fard M, Aminorroaya A, Kachuei A, et al.',
            'Diagnostics (Basel)',
            '2023-03-01'::date,
            '36900073',
            'review'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 2 inserted: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 2 already exists: %', v_article_id;
    END IF;
END $$;

-- Article 3: VLDL-cholesterol and cardiovascular events (2021)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '32810544';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            pm_id,
            article_type
        ) VALUES (
            gen_random_uuid(),
            'The relation between VLDL-cholesterol and risk of cardiovascular events in patients with manifest cardiovascular disease',
            'Marston NA, Giugliano RP, Melloni GEM, et al.',
            'International Journal of Cardiology',
            '2021-01-01'::date,
            '32810544',
            'research_article'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 3 inserted: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 3 already exists: %', v_article_id;
    END IF;
END $$;

-- Article 4: 2024 Guidelines on lipid metabolism (2024)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '38757022';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            pm_id,
            article_type
        ) VALUES (
            gen_random_uuid(),
            '2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders',
            'Banach M, Burchardt P, Chlebus K, et al.',
            'Archives of Medical Science',
            '2024-03-18'::date,
            '38757022',
            'review'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 4 inserted: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 4 already exists: %', v_article_id;
    END IF;
END $$;

-- ----------------------------------------------------------------------------
-- 3. LINK ARTICLES TO SCORE ITEM
-- ----------------------------------------------------------------------------

-- Link Article 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'::uuid
FROM articles a
WHERE a.pm_id = '37045908'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'::uuid
FROM articles a
WHERE a.pm_id = '36900073'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'::uuid
FROM articles a
WHERE a.pm_id = '32810544'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'::uuid
FROM articles a
WHERE a.pm_id = '38757022'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- 4. VERIFICATION
-- ----------------------------------------------------------------------------

-- Verify articles linked
SELECT
    si.name as score_item,
    a.title,
    a.journal,
    a.publish_date,
    a.pm_id,
    a.article_type
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'
ORDER BY a.publish_date DESC;

-- Final count
SELECT
    si.name,
    COUNT(DISTINCT a.id) as article_count,
    si.last_review
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '0ed8b90b-8679-4842-a7b2-68b0e3ea9154'
GROUP BY si.id, si.name, si.last_review;

COMMIT;

-- ============================================================================
-- ENRICHMENT SUMMARY
-- ============================================================================
-- Score Item: VLDL Colesterol
-- Clinical Relevance: 1,865 characters ✓
-- Patient Explanation: 1,390 characters ✓
-- Conduct: 2,463 characters ✓
-- Articles: 4 peer-reviewed (2021-2024)
--   1. PMID 37045908 - Remnant cholesterol/metabolic syndrome (2023)
--   2. PMID 36900073 - TG/HDL-C ratio/CVD risk (2023)
--   3. PMID 32810544 - VLDL-C/cardiovascular events (2021)
--   4. PMID 38757022 - 2024 Lipid guidelines (2024)
-- Last Review: 2026-01-29
-- ============================================================================
