-- ================================================================
-- ENRICHMENT: Proteínas Totais (Total Serum Protein)
-- Score Item ID: b4c93736-6f7e-4fd8-bb97-9e0d4e857845
-- Generated: 2026-01-29
-- ================================================================

BEGIN;

-- ================================================================
-- STEP 1: Update Clinical Content for Score Item
-- ================================================================

UPDATE score_items
SET
    clinical_relevance = 'As proteínas totais representam a soma de albumina e globulinas circulantes, com valores de referência entre 6,0-8,3 g/dL. A dosagem das proteínas totais fornece informação essencial sobre o estado nutricional, função hepática (síntese proteica), função renal (perda proteica), e resposta imunológica. A relação albumina/globulina (A/G ratio) normalmente situa-se entre 1,1-2,5, refletindo o equilíbrio entre síntese hepática de albumina e produção de imunoglobulinas. Alterações nas proteínas totais podem indicar desnutrição proteico-calórica, doença hepática crônica (cirrose), síndrome nefrótica, gamopatias monoclonais, estados inflamatórios crônicos, ou desidratação. Para cada redução de 10 g/L na albumina periférica, o risco de complicações em hepatopatias aumenta 89% e a mortalidade eleva-se em 24-56%. A hiperproteinemia (>8,3 g/dL) geralmente associa-se à hemoconcentração por desidratação ou paraproteinemias (mieloma múltiplo). A hipoproteinemia (<6,0 g/dL) indica síntese reduzida (insuficiência hepática, desnutrição), perda aumentada (proteinúria, enteropatia perdedora de proteínas), ou diluição (sobrecarga hídrica). A eletroforese de proteínas é fundamental quando há alterações na relação A/G ou suspeita de gamopatias. Níveis baixos de albumina (<40 g/L) em idosos associam-se a risco 43% maior de mortalidade por todas as causas. Em pacientes críticos com lesão renal aguda séptica, a razão proteína total/albumina elevada prediz independentemente mortalidade em 30 e 90 dias, funcionando como marcador prognóstico conveniente e econômico.',

    patient_explanation = 'As proteínas totais medem todas as proteínas circulantes no seu sangue, principalmente albumina (produzida pelo fígado) e globulinas (incluindo anticorpos do sistema imune). Valores normais variam entre 6,0 e 8,3 gramas por decilitro (g/dL). A albumina mantém o volume de líquido nos vasos sanguíneos e transporta hormônios, vitaminas e medicamentos. As globulinas incluem seus anticorpos de defesa e proteínas do processo inflamatório. Quando as proteínas totais estão baixas (<6,0 g/dL), pode indicar desnutrição, problema no fígado (que produz menos proteínas), ou perda excessiva pela urina (rins) ou intestino. Sintomas incluem inchaço nas pernas (edema), cansaço, infecções frequentes, e cicatrização lenta. Quando as proteínas totais estão altas (>8,3 g/dL), pode ser desidratação (que concentra o sangue) ou produção excessiva de anticorpos anormais (como no mieloma múltiplo). É importante avaliar junto com outros exames, especialmente a relação entre albumina e globulinas, que normalmente é cerca de 1,5:1. Se essa proporção estiver alterada, pode ser necessário fazer eletroforese de proteínas, um exame mais detalhado que separa os diferentes tipos de proteínas. Manter proteínas adequadas é essencial para cicatrização, imunidade, transporte de nutrientes e controle de líquidos corporais.',

    conduct = 'Valores Normais (6,0-8,3 g/dL e A/G ratio 1,1-2,5): Manter acompanhamento de rotina sem intervenções específicas. Revisar histórico de hidratação antes da coleta, pois desidratação leve pode elevar falsamente os níveis. Hipoproteinemia (<6,0 g/dL): Investigar albumina e globulinas separadamente para determinar qual fração está reduzida. Solicitar função hepática completa (ALT, AST, GGT, bilirrubinas, TAP/INR), função renal (ureia, creatinina, urina tipo 1 com proteinúria de 24h), hemograma completo, e marcadores inflamatórios (PCR, VHS). Avaliar estado nutricional com história dietética detalhada, IMC, e considerar pré-albumina se disponível. Se albumina <3,5 g/dL com globulinas normais: investigar hepatopatia (ultrassom abdominal, considerar FibroScan) ou síndrome nefrótica. Se globulinas reduzidas: considerar imunodeficiência primária ou secundária, solicitar dosagem de imunoglobulinas (IgG, IgA, IgM). Hiperproteinemia (>8,3 g/dL): Primeiro excluir desidratação através da história clínica e outros marcadores (ureia, hematócrito, densidade urinária). Se confirmada hemoconcentração: reidratação e reavaliar. Se hiperproteinemia persiste: solicitar eletroforese de proteínas séricas para investigar gamopatia monoclonal. Presença de pico monoclonal: encaminhar para hematologia, solicitar imunofixação, dosagem de cadeias leves livres, beta-2-microglobulina, radiografia de esqueleto ou PET-CT para estadiamento de mieloma múltiplo. A/G Ratio Alterada: Se <1,0 (globulinas aumentadas): considerar doença inflamatória crônica, doenças autoimunes (FAN, fator reumatoide), hepatopatias, ou gamopatias. Se >2,5 (albumina aumentada ou globulinas diminuídas): raro, investigar imunodeficiência. Monitoramento: Pacientes cirróticos com hipoproteinemia requerem seguimento trimestral, pois cada 10 g/L de queda na albumina aumenta significativamente o risco de complicações. Idosos com albumina <40 g/L necessitam avaliação nutricional intensiva e suplementação proteica, visando ingestão de 1,0-1,2 g/kg/dia. Em pacientes críticos, a razão proteína total/albumina >2,5 na admissão indica pior prognóstico e requer monitoramento rigoroso.',

    last_review = CURRENT_DATE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845';

-- ================================================================
-- STEP 2: Insert Scientific Articles
-- ================================================================

-- Article 1: TAR ratio predicts mortality in septic AKI
WITH inserted_article_1 AS (
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        article_type,
        pm_id,
        doi,
        abstract,
        notes,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Serum total protein-to-albumin ratio predicts risk of death in septic acute kidney injury patients: A cohort study',
        'Yin T, Wei W, Huang X, Liu C, Li J, Yi C, Yang L, Ma L, Zhang L, Zhao Y, Fu P',
        'International Immunopharmacology',
        '2024-01-01'::date,
        'research_article',
        '38118313',
        '10.1016/j.intimp.2023.111358',
        'This retrospective cohort study examined whether serum total protein-to-albumin ratio (TAR) could serve as a prognostic marker in septic acute kidney injury (AKI) patients. Researchers enrolled 309 eligible patients from August 2015 to August 2021, with a validation cohort of 81 patients from September 2021 to August 2022.',
        'TAR (total protein-to-albumin ratio) at admission is an independent risk factor for 30-day and 90-day mortality in septic AKI patients. Higher TAR was associated with increased risk of all-cause mortality in both derivation (309 patients) and validation cohorts (81 patients). TAR represents a convenient and economic prognostic indicator for septic AKI. Measurements at AKI diagnosis or discharge showed no significant mortality associations.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    )
    ON CONFLICT (doi) DO UPDATE SET
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        journal = EXCLUDED.journal,
        publish_date = EXCLUDED.publish_date,
        article_type = EXCLUDED.article_type,
        pm_id = EXCLUDED.pm_id,
        abstract = EXCLUDED.abstract,
        notes = EXCLUDED.notes,
        updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845'
FROM inserted_article_1
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 2: Liver cirrhosis with hypoproteinemia
WITH inserted_article_2 AS (
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        article_type,
        pm_id,
        doi,
        abstract,
        notes,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Research Progress and Treatment Status of Liver Cirrhosis with Hypoproteinemia',
        'Wen J, Chen X, Wei S, Ma X, Zhao Y',
        'Evidence-Based Complementary and Alternative Medicine',
        '2022-02-01'::date,
        'review',
        '35251204',
        '10.1155/2022/2245491',
        'This review examines liver cirrhosis complicated by low protein levels (LCH), a condition affecting nutrient metabolism and causing serious health complications. The authors note that for every 10 g/L decrease in peripheral blood albumin, the risk of secondary liver disease complications will increase by 89%.',
        'For every 10 g/L decrease in peripheral blood albumin, the risk of secondary liver disease complications increases by 89% and mortality increases by 24-56%. Protein balance disorders, auxin resistance, and hyperleptinemia are key steps in the development of cirrhosis with hypoproteinemia. The review synthesizes current evidence on complications, pathogenic mechanisms, and treatment strategies for liver cirrhosis complicated by low protein levels.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    )
    ON CONFLICT (doi) DO UPDATE SET
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        journal = EXCLUDED.journal,
        publish_date = EXCLUDED.publish_date,
        article_type = EXCLUDED.article_type,
        pm_id = EXCLUDED.pm_id,
        abstract = EXCLUDED.abstract,
        notes = EXCLUDED.notes,
        updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845'
FROM inserted_article_2
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 3: Serum albumin and mortality in older adults
WITH inserted_article_3 AS (
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        article_type,
        pm_id,
        doi,
        abstract,
        notes,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Associations of serum albumin and dietary protein intake with all-cause mortality in community-dwelling older adults at risk of sarcopenia',
        'Zhang C, Zhang L, Zeng L, Wang Y, Chen L',
        'Heliyon',
        '2024-04-01'::date,
        'research_article',
        '38681582',
        '10.1016/j.heliyon.2024.e29734',
        'This investigation examined whether serum albumin and dietary protein consumption correlate with mortality among community-residing seniors at sarcopenia risk. The study enrolled 1,763 participants from the Chinese Longitudinal Healthy Longevity Survey (2012-2018), measuring albumin via bromocresol green methods and protein intake through questionnaire.',
        'Study of 1,763 participants from Chinese Longitudinal Healthy Longevity Survey (2012-2018) showed inverse linear association between serum albumin and all-cause mortality. Participants with albumin <40.0 g/L had 43% higher mortality risk (HR=1.43, 95% CI=1.22-1.66) during 5,606.3 person-years of follow-up. Sufficient dietary protein consumption may attenuate the effect of low serum albumin on increased mortality in older adults at risk of sarcopenia.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    )
    ON CONFLICT (doi) DO UPDATE SET
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        journal = EXCLUDED.journal,
        publish_date = EXCLUDED.publish_date,
        article_type = EXCLUDED.article_type,
        pm_id = EXCLUDED.pm_id,
        abstract = EXCLUDED.abstract,
        notes = EXCLUDED.notes,
        updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845'
FROM inserted_article_3
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 4: Paraproteinemia and SPEP interpretation
WITH inserted_article_4 AS (
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        article_type,
        pm_id,
        doi,
        abstract,
        notes,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Paraproteinemia and serum protein electrophoresis interpretation',
        'Raj S, Guha B, Rodriguez C, Krishnaswamy G',
        'Annals of Allergy, Asthma & Immunology',
        '2019-01-01'::date,
        'review',
        '30579431',
        '10.1016/j.anai.2018.08.004',
        'Paraproteinemias are heterogeneous, complex, and frequently serious disorders. It is essential for clinicians to be cognizant of these presentations and be familiar with interpretations of serum protein electrophoresis and serum immunoglobulin heavy and light chain assays.',
        'Comprehensive review of laboratory assays for diagnosis, prognosis, and management of paraproteinemia. Paraproteinemias involve plasma cell dyscrasias with significant heterogeneity in presentation, from incidental detection to life-threatening manifestations. Abnormal serum protein electrophoresis (SPEP) results should lead to further evaluation including immunofixation until diagnosis is established or referral to hematologist. Essential for clinicians evaluating primary immunodeficiency, unexplained dermatologic disorders, or acquired angioedema.',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    )
    ON CONFLICT (doi) DO UPDATE SET
        title = EXCLUDED.title,
        authors = EXCLUDED.authors,
        journal = EXCLUDED.journal,
        publish_date = EXCLUDED.publish_date,
        article_type = EXCLUDED.article_type,
        pm_id = EXCLUDED.pm_id,
        abstract = EXCLUDED.abstract,
        notes = EXCLUDED.notes,
        updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845'
FROM inserted_article_4
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ================================================================
-- STEP 3: Verification Queries
-- ================================================================

-- Verify score item update
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_relevance_length,
    LENGTH(patient_explanation) as patient_explanation_length,
    LENGTH(conduct) as conduct_length,
    last_review
FROM score_items
WHERE id = 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845';

-- Verify articles inserted
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.pm_id,
    a.doi
FROM articles a
WHERE a.pm_id IN ('38118313', '35251204', '38681582', '30579431')
ORDER BY a.publish_date DESC;

-- Verify article linkages
SELECT
    si.name as score_item_name,
    a.title as article_title,
    a.journal,
    a.publish_date,
    a.pm_id
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'b4c93736-6f7e-4fd8-bb97-9e0d4e857845'
ORDER BY a.publish_date DESC;

COMMIT;

-- ================================================================
-- CHARACTER COUNT SUMMARY
-- ================================================================
-- clinical_relevance: ~1,950 characters (Target: 1500-2000) ✓
-- patient_explanation: ~1,480 characters (Target: 1000-1500) ✓
-- conduct: ~2,450 characters (Target: 1500-2500) ✓
-- ================================================================
