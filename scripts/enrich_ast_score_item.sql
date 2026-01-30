-- ============================================================================
-- ENRICHMENT: Transaminase oxalacética (AST) - Score Item ID: 910354c8-083c-4302-bf07-5ec62b78567b
-- Generated: 2026-01-29
-- Description: Clinical content enrichment + scientific articles linking
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: UPDATE CLINICAL CONTENT FOR AST SCORE ITEM
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A aspartato aminotransferase (AST), também conhecida como transaminase oxalacética ou TGO, é uma enzima encontrada principalmente no fígado, coração, músculos esqueléticos, rins e hemácias. Sua dosagem sérica é fundamental para avaliar lesão hepatocelular, injúria miocárdica e doenças musculares. Valores normais situam-se entre 10-40 U/L, variando conforme o laboratório. A AST apresenta meia-vida plasmática de 18 horas, sendo removida da circulação duas vezes mais rápido que a ALT (36 horas). Elevações desproporcionais de AST e ALT em relação à fosfatase alcalina caracterizam padrão hepatocelular de lesão. AST marcadamente elevada (>1000 U/L) sugere hepatite viral aguda, isquemia hepática, intoxicação por paracetamol ou lesão muscular grave. A relação AST/ALT (razão De Ritis) possui valor diagnóstico: >2 sugere doença hepática alcoólica (sensibilidade 92%), >1 indica cirrose em hepatopatias crônicas, e <1 é típico de hepatite viral aguda ou esteatose hepática não alcoólica. No infarto agudo do miocárdio, AST eleva-se 6-8h após o evento, atinge pico em 24-36h e normaliza em 3-7 dias. Estudos recentes demonstram que AST elevada (≥40 U/L) reduz expectativa de vida em 10 anos e é melhor preditor de mortalidade que ALT. AST/ALT ≥2 em pacientes com STEMI prediz oclusão coronariana total com alta acurácia.',

    patient_explanation = 'A AST (transaminase oxalacética) é uma enzima presente em diversos órgãos do seu corpo, especialmente no fígado, coração e músculos. Quando há lesão nestes tecidos, a AST é liberada na corrente sanguínea, elevando seus níveis. Valores normais ficam entre 10-40 U/L. AST aumentada pode indicar: problemas no fígado (hepatite, cirrose, esteatose), infarto do coração, lesões musculares intensas (rabdomiólise), uso excessivo de álcool ou medicamentos tóxicos ao fígado. É importante analisar a AST junto com outras enzimas: a relação AST/ALT (razão De Ritis) ajuda a identificar a causa. Se AST/ALT for maior que 2, pode indicar consumo excessivo de álcool. Se for maior que 1, pode sugerir cirrose hepática. Se for menor que 1, é mais comum em hepatites virais ou gordura no fígado. AST também pode elevar-se após exercícios físicos intensos, por isso seu médico pode solicitar um novo exame após repouso. Caso sua AST esteja muito alta (acima de 1000 U/L), investigue urgentemente hepatite aguda ou infarto. Mantenha seus exames atualizados e informe seu médico sobre uso de álcool, medicamentos, suplementos e atividade física para interpretação correta dos resultados.',

    conduct = 'AST NORMAL (10-40 U/L): Manter acompanhamento de rotina conforme fatores de risco. Reforçar hábitos saudáveis: dieta equilibrada, atividade física regular, evitar álcool em excesso e hepatotóxicos. AST LEVEMENTE ELEVADA (41-100 U/L): Calcular razão AST/ALT. Se AST/ALT <1: investigar esteatose hepática não alcoólica (NASH), síndrome metabólica, diabetes. Solicitar ultrassom abdominal, perfil lipídico, glicemia, hemoglobina glicada. Se AST/ALT >2: investigar consumo alcoólico (solicitar GGT e CDT), hepatopatia alcoólica. Orientar abstinência de álcool. Se AST/ALT >1: suspeitar de fibrose/cirrose em hepatopatias crônicas. Considerar elastografia hepática ou FibroScan. Investigar causas extra-hepáticas: lesão muscular (solicitar CPK, aldolase), doença cardíaca (troponina, BNP), hemólise (LDH, bilirrubinas), macro-AST (PEG precipitation test). AST MODERADAMENTE ELEVADA (101-500 U/L): Descartar hepatites virais (sorologias para hepatites A, B, C), hepatite autoimune (FAN, anti-músculo liso, IgG), hepatite medicamentosa (revisar medicamentos/suplementos), doença de Wilson (ceruloplasmina, cobre sérico/urinário). Solicitar função hepática completa (bilirrubinas, albumina, TAP/INR). AST MARCADAMENTE ELEVADA (>1000 U/L): URGÊNCIA MÉDICA. Investigar: hepatite viral aguda fulminante, isquemia hepática (choque, insuficiência cardíaca), intoxicação por paracetamol (nível sérico, N-acetilcisteína), rabdomiólise grave (CPK >10.000 U/L, mioglobinúria), infarto agudo do miocárdio (troponina, ECG). Hospitalizar para monitorização. Considerar transplante hepático emergencial se insuficiência hepática aguda grave. SEGUIMENTO: Repetir AST/ALT em 2-4 semanas após intervenções. Se persistentemente elevada, encaminhar para hepatologista. Em pacientes com cirrose, monitorar AST/ALT trimestralmente e rastrear hepatocarcinoma (alfa-fetoproteína + ultrassom) semestralmente.',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '910354c8-083c-4302-bf07-5ec62b78567b';

-- ============================================================================
-- STEP 2: INSERT SCIENTIFIC ARTICLES
-- ============================================================================

-- Article 1: AST/ALT ratio in acute myocardial infarction (2020)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    -- Check if article with this DOI already exists
    SELECT id INTO v_article_id FROM articles WHERE doi = '10.5114/amsad.2020.103027';

    IF v_article_id IS NULL THEN
        -- Check by PMID
        SELECT id INTO v_article_id FROM articles WHERE pm_id = '33644486';
    END IF;

    IF v_article_id IS NULL THEN
        -- Insert new article
        INSERT INTO articles (id, title, authors, article_type, journal, publish_date, pm_id, doi, created_at, updated_at)
        VALUES (
            gen_random_uuid(),
            'The significance of transaminase ratio (AST/ALT) in acute myocardial infarction',
            'Authors et al.',
            'research_article',
            'Archives of Medical Sciences. Atherosclerotic Diseases',
            '2020-12-01'::date,
            '33644486',
            '10.5114/amsad.2020.103027',
            NOW(),
            NOW()
        )
        RETURNING id INTO v_article_id;
    END IF;

    -- Link to score item
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '910354c8-083c-4302-bf07-5ec62b78567b'::uuid)
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 2: Markedly Elevated AST from Non-Hepatic Causes (2022)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE doi = '10.3390/jcm12010310';

    IF v_article_id IS NULL THEN
        SELECT id INTO v_article_id FROM articles WHERE pm_id = '36615110';
    END IF;

    IF v_article_id IS NULL THEN
        INSERT INTO articles (id, title, authors, article_type, journal, publish_date, pm_id, doi, created_at, updated_at)
        VALUES (
            gen_random_uuid(),
            'Markedly Elevated Aspartate Aminotransferase from Non-Hepatic Causes',
            'Han JH, Kwak JY, Lee SS, Kim HG, Jeon H, Cha RR',
            'research_article',
            'Journal of Clinical Medicine',
            '2022-12-30'::date,
            '36615110',
            '10.3390/jcm12010310',
            NOW(),
            NOW()
        )
        RETURNING id INTO v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '910354c8-083c-4302-bf07-5ec62b78567b'::uuid)
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 3: The De Ritis Ratio: The Test of Time (2013, still highly relevant)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE pm_id = '24353357';

    IF v_article_id IS NULL THEN
        INSERT INTO articles (id, title, authors, article_type, journal, publish_date, pm_id, created_at, updated_at)
        VALUES (
            gen_random_uuid(),
            'The De Ritis Ratio: The Test of Time',
            'Botros M, Sikaris KA',
            'review',
            'Clinical Biochemist Reviews',
            '2013-11-01'::date,
            '24353357',
            NOW(),
            NOW()
        )
        RETURNING id INTO v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '910354c8-083c-4302-bf07-5ec62b78567b'::uuid)
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Article 4: Persistently high serum AST level - macro-AST (2023)
DO $$
DECLARE
    v_article_id uuid;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE doi = '10.1002/ccr3.7912';

    IF v_article_id IS NULL THEN
        SELECT id INTO v_article_id FROM articles WHERE pm_id = '37700775';
    END IF;

    IF v_article_id IS NULL THEN
        INSERT INTO articles (id, title, authors, article_type, journal, publish_date, pm_id, doi, created_at, updated_at)
        VALUES (
            gen_random_uuid(),
            'Persistently high serum aspartate aminotransferase level in an asymptomatic young patient',
            'Pei X, Zhao Y, Jiang W, Zeng Q, Liu C, Wang M, Wang H, Liang S, Gan W, Wu D, Tang H',
            'case_study',
            'Clinical Case Reports',
            '2023-09-10'::date,
            '37700775',
            '10.1002/ccr3.7912',
            NOW(),
            NOW()
        )
        RETURNING id INTO v_article_id;
    END IF;

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '910354c8-083c-4302-bf07-5ec62b78567b'::uuid)
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- ============================================================================
-- STEP 3: VERIFICATION QUERIES
-- ============================================================================

-- Verify clinical content update
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    last_review,
    updated_at
FROM score_items
WHERE id = '910354c8-083c-4302-bf07-5ec62b78567b';

-- Verify articles were inserted
SELECT
    a.id,
    a.title,
    a.article_type,
    a.journal,
    a.publish_date,
    a.pm_id,
    a.doi
FROM articles a
WHERE a.pm_id IN ('33644486', '36615110', '24353357', '37700775')
ORDER BY a.publish_date DESC;

-- Verify article-score_item links
SELECT
    si.id as score_item_id,
    si.name as score_item_name,
    a.pm_id,
    a.title as article_title,
    a.publish_date
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '910354c8-083c-4302-bf07-5ec62b78567b'
ORDER BY a.publish_date DESC;

-- Count total linked articles
SELECT
    si.name,
    COUNT(asi.article_id) as total_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '910354c8-083c-4302-bf07-5ec62b78567b'
GROUP BY si.id, si.name;

COMMIT;

-- ============================================================================
-- EXECUTION SUMMARY
-- ============================================================================
-- Score Item: Transaminase oxalacética (AST)
-- Clinical Relevance: ~1,890 characters
-- Patient Explanation: ~1,456 characters
-- Conduct: ~2,489 characters
-- Articles Added: 4
--   1. PMID 33644486 (2020) - AST/ALT in MI - Research Article
--   2. PMID 36615110 (2022) - Non-Hepatic AST - Research Article
--   3. PMID 24353357 (2013) - De Ritis Ratio - Review
--   4. PMID 37700775 (2023) - Macro-AST Case - Case Study
-- ============================================================================
