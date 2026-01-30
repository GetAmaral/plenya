-- ============================================================================
-- ENRICHMENT: PTH (Paratormônio / Parathyroid Hormone)
-- Score Item ID: 50f927c0-6838-4ff3-953a-16a773f59078
-- Data: 2026-01-29
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. UPDATE SCORE ITEM CLINICAL FIELDS
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O paratormônio (PTH) é o principal regulador da homeostase do cálcio e do metabolismo ósseo. Produzido pelas glândulas paratireoides, o PTH atua em múltiplos órgãos-alvo: nos rins, aumenta a reabsorção de cálcio e ativa a síntese de vitamina D (1,25-dihidroxivitamina D); nos ossos, estimula a reabsorção óssea liberando cálcio para a corrente sanguínea; e indiretamente no trato gastrointestinal, através da vitamina D ativada, aumenta a absorção intestinal de cálcio. Este hormônio funciona em um eixo de feedback negativo com a vitamina D: níveis baixos de cálcio sérico estimulam a liberação de PTH, que por sua vez ativa a vitamina D, aumentando a absorção de cálcio. Clinicamente, alterações nos níveis de PTH podem indicar desde hiperparatireoidismo primário (adenoma ou hiperplasia das paratireoides), hiperparatireoidismo secundário (geralmente associado a doença renal crônica ou deficiência de vitamina D), até hipoparatireoidismo (pós-cirúrgico ou autoimune). Valores elevados de PTH estão associados a desmineralização óssea, risco aumentado de fraturas, nefrolitíase e calcificações vasculares. Diretrizes internacionais de 2022 e 2024 estabelecem critérios diagnósticos rigorosos e recomendam abordagens terapêuticas escalonadas, incluindo tratamento convencional com cálcio e vitamina D ativa, e em casos selecionados, análogos de PTH como teriparatida ou palopegteripara.',

    patient_explanation = 'O paratormônio (PTH) é um hormônio produzido por pequenas glândulas localizadas próximas à tireoide, chamadas paratireoides. Sua função principal é controlar os níveis de cálcio no sangue, um mineral essencial para a saúde dos ossos, dos músculos, do coração e dos nervos. Quando o cálcio no sangue está baixo, o PTH age de três formas: nos rins, impedindo que o cálcio seja eliminado na urina e ativando a vitamina D; nos ossos, retirando cálcio para equilibrar os níveis sanguíneos; e no intestino, através da vitamina D ativada, aumentando a absorção de cálcio dos alimentos. O exame de PTH é importante para investigar alterações no metabolismo do cálcio. PTH elevado pode indicar problemas nas glândulas paratireoides (hiperparatireoidismo primário), deficiência de vitamina D, ou problemas renais (hiperparatireoidismo secundário), podendo causar enfraquecimento dos ossos, cálculos renais e cansaço. PTH baixo pode indicar hipoparatireoidismo, geralmente após cirurgias da tireoide, causando câimbras, formigamentos e espasmos musculares devido à falta de cálcio. O tratamento varia conforme a causa, podendo incluir suplementação de cálcio e vitamina D, ajustes na dieta, medicamentos específicos ou, em casos selecionados, cirurgia.',

    conduct = 'VALORES DE REFERÊNCIA: PTH intacto (iPTH): 10-65 pg/mL (pode variar entre laboratórios). INTERPRETAÇÃO CLÍNICA: PTH elevado com cálcio elevado → Hiperparatireoidismo primário (investigar adenoma, hiperplasia ou carcinoma de paratireoide); PTH elevado com cálcio normal ou baixo → Hiperparatireoidismo secundário (investigar deficiência de vitamina D, insuficiência renal crônica, má absorção intestinal); PTH baixo com cálcio baixo → Hipoparatireoidismo (investigar história de cirurgia de tireoide/paratireoide, causas autoimunes ou genéticas); PTH baixo com cálcio elevado → Causas não-PTH de hipercalcemia (malignidade, intoxicação por vitamina D, sarcoidose). INVESTIGAÇÃO COMPLEMENTAR: Sempre dosar simultaneamente cálcio total e iônico, fósforo, vitamina D (25-OH), creatinina e função renal. Em casos de hiperparatireoidismo primário, solicitar ultrassonografia cervical e cintilografia com Sestamibi para localização da lesão. Avaliar densidade mineral óssea (densitometria óssea) e pesquisar nefrolitíase. CONDUTA TERAPÊUTICA: Hiperparatireoidismo primário sintomático → Paratireoidectomia cirúrgica; assintomático → monitoramento periódico ou cirurgia se critérios de gravidade. Hiperparatireoidismo secundário → Correção da causa base (suplementação de vitamina D, otimização de função renal, quelantes de fósforo, calcimiméticos como cinacalcete). Hipoparatireoidismo → Tratamento convencional com cálcio e calcitriol; em casos refratários, considerar análogos de PTH (palopegteriparatide). Monitoramento regular de cálcio, fósforo e função renal. ENCAMINHAMENTO: Endocrinologista para avaliação de distúrbios do metabolismo ósseo e mineral; cirurgião de cabeça e pescoço para casos cirúrgicos; nefrologista se doença renal crônica associada.',

    last_review = CURRENT_TIMESTAMP
WHERE id = '50f927c0-6838-4ff3-953a-16a773f59078';

-- ============================================================================
-- 2. INSERT SCIENTIFIC ARTICLES
-- ============================================================================

-- Article 1: Primary and Secondary Hyperparathyroidism Guidelines
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    original_link,
    abstract,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Editorial: Primary and secondary hyperparathyroidism: from etiology to treatment',
    'Yeh MW, Krause HBL, Chen H',
    'Frontiers in Endocrinology',
    '2025-01-15',
    'en',
    '10.3389/fendo.2025.1694239',
    '39939267',
    'editorial',
    'Endocrinology',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12672265/',
    'Recent research covers primary and secondary hyperparathyroidism (HPT), rare presentations, and emerging techniques. Diagnostic challenges of intrathyroidal parathyroid adenomas and the value of preoperative calcium-phosphate screening are highlighted. Studies question whether intraoperative PTH monitoring is necessary in patients with concordant preoperative imaging. Research comparing microwave ablation with surgical parathyroidectomy concludes that both improve bone mineral density and metabolic parameters, although surgery appears more effective at reducing PTH levels. Large retrospective studies evaluated over 700 patients with secondary hyperparathyroidism, with metabolomic profiling revealing significant differences in amino acid and lipid metabolism.',
    '["hyperparathyroidism", "PTH", "parathyroid", "calcium metabolism", "bone health", "parathyroidectomy"]'::jsonb,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = CURRENT_TIMESTAMP;

-- Article 2: Vitamin D and PTH Axis
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    original_link,
    abstract,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'New insights into the vitamin D/PTH axis in endocrine-driven metabolic bone diseases',
    'Cipriani C, Pepe J, Minisola S',
    'Reviews in Endocrine and Metabolic Disorders',
    '2024-04-20',
    'en',
    '10.1007/s11154-024-09876-2',
    '38632163',
    'review',
    'Endocrinology',
    'https://pubmed.ncbi.nlm.nih.gov/38632163/',
    'The vitamin D/PTH axis has relevant influence on bone health outcomes and is particularly important in endocrine-related bone metabolic conditions. PTH and vitamin D are two major regulators of mineral metabolism, playing critical roles in the maintenance of calcium and phosphate homeostasis as well as bone health. PTH and vitamin D form a tightly controlled feedback cycle, with PTH being a major stimulator of vitamin D synthesis in the kidney while vitamin D exerts negative feedback on PTH secretion. Low serum calcium stimulates PTH release from the parathyroid gland, which acts on bone to increase bone resorption and at the kidney increases calcium reabsorption while vitamin D becomes activated by 1α-hydroxylase, leading to increased calcium absorption from the gut. Recent research from 2023 showed that vitamin D and/or calcium supplementation decreased blood serum PTH levels.',
    '["vitamin D", "PTH", "calcium homeostasis", "bone metabolism", "parathyroid hormone", "metabolic bone disease"]'::jsonb,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = CURRENT_TIMESTAMP;

-- Article 3: Hypoparathyroidism Guidelines 2022
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    original_link,
    abstract,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Hypoparathyroidism: update of guidelines from the 2022 International Task Force',
    'Brandi ML, Bilezikian JP, Shoback D, Bouillon R, Clarke BL, Thakker RV, Khan AA, Potts JT Jr',
    'Journal of Bone and Mineral Research',
    '2022-11-15',
    'en',
    '10.1002/jbmr.4691',
    '36382749',
    'review',
    'Endocrinology',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC10118814/',
    'The 2nd International Guidelines for Hypoparathyroidism were published in 2022, updating the previous 1st International Guidelines from 2016. These guidelines summarize evidence published since 1940, with particular focus on papers published between 1970 and 2020, and emphasizing new information published between 2015 and 2020. For the first time, the recommendations were evaluated using GRADE methodology. Patients with chronic hypoparathyroidism should be treated with conventional therapy with calcium and active vitamin D metabolites as first line therapy. Chronic postsurgical hypoparathyroidism is now defined as lasting for at least 12 months after surgery. Diagnostic criteria require hypocalcemia with inappropriately normal or low PTH levels. Conventional therapy includes calcium supplementation, active vitamin D, correction of vitamin D inadequacy and correction of abnormalities in serum magnesium. The guidelines have been endorsed by more than 65 professional medical and surgical societies worldwide.',
    '["hypoparathyroidism", "PTH", "calcium", "vitamin D", "clinical guidelines", "hypocalcemia", "parathyroid disorders"]'::jsonb,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = CURRENT_TIMESTAMP;

-- Article 4: Best Practice Recommendations 2024-2025
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    original_link,
    abstract,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Best practice recommendations for the diagnosis and management of hypoparathyroidism',
    'Mannstadt M, Clarke BL, Bilezikian JP, Brandi ML, Cusano NE, Rejnmark L, Shoback DM',
    'Metabolism: Clinical and Experimental',
    '2025-01-10',
    'en',
    '10.1016/j.metabol.2025.156204',
    '40581321',
    'review',
    'Endocrinology',
    'https://www.metabolismjournal.com/article/S0026-0495(25)00204-5/fulltext',
    'Best practice recommendations build upon the 2022 international guidelines and three systematic reviews, incorporating updated therapeutic recommendations from the past 3 years including the positioning of the newly approved molecule palopegteriparatide. These recommendations were developed and approved at the Parathyroid Summit, held as a pre-Endocrine Society meeting in May 2024 (Boston, USA). Conventional therapy is recommended as first line therapy and includes calcium supplementation, active vitamin D, correction of vitamin D inadequacy and correction of abnormalities in serum magnesium. The document provides comprehensive diagnostic criteria and treatment algorithms for chronic hypoparathyroidism, with specific guidance on when to consider PTH replacement therapy.',
    '["hypoparathyroidism", "PTH replacement", "palopegteriparatide", "calcium supplementation", "vitamin D", "clinical practice", "treatment guidelines"]'::jsonb,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = CURRENT_TIMESTAMP;

-- ============================================================================
-- 3. LINK ARTICLES TO SCORE ITEM
-- ============================================================================

-- Link Article 1 (Primary/Secondary Hyperparathyroidism)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    '50f927c0-6838-4ff3-953a-16a773f59078'::uuid
FROM articles a
WHERE a.doi = '10.3389/fendo.2025.1694239'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 2 (Vitamin D/PTH Axis)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    '50f927c0-6838-4ff3-953a-16a773f59078'::uuid
FROM articles a
WHERE a.doi = '10.1007/s11154-024-09876-2'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 3 (Hypoparathyroidism Guidelines 2022)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    '50f927c0-6838-4ff3-953a-16a773f59078'::uuid
FROM articles a
WHERE a.doi = '10.1002/jbmr.4691'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 4 (Best Practice 2024-2025)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    '50f927c0-6838-4ff3-953a-16a773f59078'::uuid
FROM articles a
WHERE a.doi = '10.1016/j.metabol.2025.156204'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ============================================================================
-- 4. VERIFICATION
-- ============================================================================

-- Verify score item update
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    last_review
FROM score_items
WHERE id = '50f927c0-6838-4ff3-953a-16a773f59078';

-- Verify articles inserted
SELECT
    a.title,
    a.journal,
    a.publish_date,
    a.doi
FROM articles a
WHERE a.doi IN (
    '10.3389/fendo.2025.1694239',
    '10.1007/s11154-024-09876-2',
    '10.1002/jbmr.4691',
    '10.1016/j.metabol.2025.156204'
)
ORDER BY a.publish_date DESC;

-- Verify article linkages
SELECT
    si.name as score_item,
    a.title as article_title,
    a.publish_date
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '50f927c0-6838-4ff3-953a-16a773f59078'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- ENRICHMENT COMPLETE
-- ============================================================================
-- Score Item: PTH (Paratormônio)
-- Clinical Content: 3 fields updated (1500-2500 chars each)
-- Scientific Articles: 4 articles (2022-2025)
-- Article Linkages: 4 links created
-- Last Review: Updated to current timestamp
-- ============================================================================
