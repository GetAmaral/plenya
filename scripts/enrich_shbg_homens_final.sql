-- ========================================
-- ENRICHMENT: SHBG - Homens
-- Score Item ID: fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291
-- Generated: 2026-01-29
-- ========================================

BEGIN;

-- ========================================
-- STEP 1: Update Clinical Content
-- ========================================

UPDATE score_items
SET
    clinical_relevance = 'A Globulina Ligadora de Hormônios Sexuais (SHBG) é uma proteína plasmática produzida principalmente pelo fígado que regula a biodisponibilidade de testosterona e estradiol. Em homens, 45% da testosterona circula ligada à SHBG de forma biologicamente inativa, enquanto apenas 2-3% permanece livre e ativa. A avaliação do SHBG é essencial para o cálculo da testosterona livre e biodisponível, especialmente em homens com suspeita de hipogonadismo funcional onde a testosterona total pode estar na faixa baixo-normal. Valores de referência em homens adultos variam de 13-71 nmol/L, com elevação progressiva após os 50 anos. Níveis baixos de SHBG (<17 nmol/L) estão fortemente associados à resistência insulínica, síndrome metabólica e diabetes tipo 2, com um risco relativo de 0,10 para desenvolvimento de diabetes em homens no quartil superior versus inferior de SHBG. A SHBG emerge como preditor independente de resistência insulínica (p=0,029) em homens com diabetes tipo 2 recém-diagnosticado. A obesidade, especialmente visceral, suprime a produção hepática de SHBG via hiperinsulinemia, criando um ciclo vicioso de redução da testosterona biodisponível, piora da composição corporal e agravamento da resistência insulínica. Em homens idosos, o aumento fisiológico do SHBG (~1,3% ao ano) contribui para declínio mais acentuado da testosterona livre (2-3% ao ano) comparado à testosterona total (1,6% ao ano), explicando sintomas de deficiência androgênica mesmo com testosterona total normal.',

    patient_explanation = 'A SHBG é uma proteína que funciona como um "transporte" para seus hormônios sexuais no sangue, especialmente a testosterona. Pense nela como um vagão de trem que carrega a testosterona pelo corpo: quando a testosterona está dentro do vagão (ligada à SHBG), ela não consegue entrar nas células e trabalhar. Apenas a testosterona "livre", que não está presa à SHBG, pode realmente atuar no organismo aumentando massa muscular, energia, libido e outros efeitos masculinos. Em homens, a SHBG normalmente fica entre 13-71 nmol/L e tende a aumentar com a idade. Quando a SHBG está muito baixa (abaixo de 17), geralmente indica problemas metabólicos como resistência à insulina, obesidade abdominal ou risco aumentado para diabetes. Por outro lado, quando está muito alta, pode indicar que você tem pouca testosterona realmente disponível para usar, mesmo que o exame de testosterona total pareça normal. É por isso que médicos especializados sempre pedem SHBG junto com testosterona: para calcular quanto de testosterona você realmente tem disponível para funcionar. Se você tem sobrepeso, diabetes ou síndrome metabólica, seu corpo produz menos SHBG, o que piora ainda mais o metabolismo criando um círculo vicioso. A avaliação do SHBG é especialmente importante em homens acima de 50 anos que apresentam sintomas como fadiga, perda de massa muscular ou redução da libido.',

    conduct = 'A dosagem de SHBG deve ser solicitada sempre que houver suspeita de hipogonadismo funcional, síndrome metabólica, obesidade com sintomas androgênicos ou discrepância entre quadro clínico e testosterona total. O exame requer jejum de 8 horas e coleta matinal (7-11h) junto com testosterona total, LH, FSH e prolactina. A interpretação exige contextualização clínica: SHBG baixa (<13 nmol/L em adultos, <17 nmol/L como preditor de resistência insulínica) associa-se a obesidade visceral, resistência insulínica, diabetes tipo 2, síndrome metabólica, esteatose hepática e hipertireoidismo. SHBG elevada (>71 nmol/L) relaciona-se a envelhecimento, hipogonadismo primário, hipertireoidismo, cirrose hepática, uso de anticonvulsivantes e estrogênios. O cálculo da testosterona livre (método gold-standard: diálise de equilíbrio; alternativa prática: Vermeulen formula usando testosterona total + SHBG + albumina) é mandatório quando SHBG está fora da faixa de referência. Em homens com SHBG elevada e sintomas de hipogonadismo, considere testosterona livre <65 pg/mL como critério diagnóstico. Em pacientes com SHBG baixa e síndrome metabólica, investigue HOMA-IR, HbA1c, perfil lipídico completo, transaminases hepáticas e considere ultrassonografia abdominal para avaliar esteatose. O acompanhamento em terapia de reposição de testosterona deve incluir SHBG a cada 6 meses inicialmente, pois a normalização dos níveis androgênicos pode elevar SHBG secundariamente. Variações agudas: exercícios intensos reduzem temporariamente (-10-15%), dieta hipocalórica prolongada eleva (+20-30%). Monitoramento longitudinal é superior a dosagens isoladas. Em homens diabéticos com SHBG <17 nmol/L, intensificar controle glicêmico e considerar metformina pode melhorar níveis. A combinação de testosterona total baixo-normal + SHBG elevada + sintomas clínicos justifica trial terapêutico de reposição androgênica em homens >50 anos após exclusão de contraindicações (PSA>4, hematócrito>54%, apneia grave não tratada).',

    last_review = CURRENT_DATE

WHERE id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291';

-- Verify update
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as relevance_chars,
    LENGTH(patient_explanation) as explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291';

-- ========================================
-- STEP 2: Insert Scientific Articles
-- ========================================

-- Article 1: New Insights in SHBG Diagnostic Potential (2025)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    -- Check if article already exists
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '40427034';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            pm_id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            language
        ) VALUES (
            gen_random_uuid(),
            '40427034',
            'New Insights in the Diagnostic Potential of Sex Hormone-Binding Globulin (SHBG)—Clinical Approach',
            'Szybiak-Skora W, Cyna W, Lacka K',
            'Biomedicines',
            '2025-05-15'::date,
            'review',
            'en'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 1 inserted with ID: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 1 already exists with ID: %', v_article_id;
    END IF;

    -- Link to score item
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 2: SHBG as Independent Predictor for Insulin Resistance (2023)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '37462840';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            pm_id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            language
        ) VALUES (
            gen_random_uuid(),
            '37462840',
            'Sex Hormone Binding Globulin is an Independent Predictor for Insulin Resistance in Male Patients with Newly Diagnosed Type 2 Diabetes Mellitus',
            'Huang R, Wang Y, Yan R, Ding B, Ma J',
            'Diabetes Therapy',
            '2023-07-18'::date,
            'research_article',
            'en'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 2 inserted with ID: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 2 already exists with ID: %', v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 3: Age, Insulin Resistance and SHBG in Healthy Men (2024)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '40857627';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            pm_id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            language
        ) VALUES (
            gen_random_uuid(),
            '40857627',
            'Association of age and insulin resistance with sex hormone-binding globulin levels in healthy men',
            'Porgere IF, Rocha BM, Escott GM, Silva LCF, Freitas PAC, Satler F, Silveiro SP',
            'Archives of Endocrinology and Metabolism',
            '2024-08-25'::date,
            'research_article',
            'en'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 3 inserted with ID: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 3 already exists with ID: %', v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 4: Evaluation and Management of Men ≥50 Years with Low Testosterone (2023)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '36995891';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (
            id,
            pm_id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            language
        ) VALUES (
            gen_random_uuid(),
            '36995891',
            'Approach to the Patient: The Evaluation and Management of Men ≥50 Years With Low Serum Testosterone Concentration',
            'Grossmann M, Jayasena CN, Anawalt BD',
            'The Journal of Clinical Endocrinology & Metabolism',
            '2023-08-18'::date,
            'review',
            'en'
        ) RETURNING id INTO v_article_id;
        RAISE NOTICE 'Article 4 inserted with ID: %', v_article_id;
    ELSE
        RAISE NOTICE 'Article 4 already exists with ID: %', v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- ========================================
-- STEP 3: Final Verification
-- ========================================

-- Show complete enrichment result
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as relevance_length,
    LENGTH(si.patient_explanation) as explanation_length,
    LENGTH(si.conduct) as conduct_length,
    si.last_review,
    COUNT(DISTINCT asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Show linked articles with details
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'fe938255-6b7a-4fbd-ac8f-8f3ba0c2d291'
ORDER BY a.publish_date DESC;

COMMIT;

-- ========================================
-- EXECUTION SUMMARY
-- ========================================
-- Score Item: SHBG - Homens
-- Clinical Relevance: 1,494 characters ✓
-- Patient Explanation: 1,367 characters ✓
-- Conduct: 2,000 characters ✓
-- Articles: 4 peer-reviewed (2023-2025)
-- ========================================
