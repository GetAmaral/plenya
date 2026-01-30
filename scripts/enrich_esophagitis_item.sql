-- ===============================================
-- ENRICHMENT: Endoscopia Alta - Esofagite (LA Classification)
-- Item ID: 4f6e007b-dcc2-4e51-aaad-b7359717f297
-- ===============================================

BEGIN;

-- 1. Inserir Articles Científicos
-- ===============================================

-- Article 1: Lyon Consensus 2.0 (2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    article_type,
    specialty,
    keywords,
    mesh_terms
) VALUES (
    gen_random_uuid(),
    'The Los Angeles-B esophagitis is a conclusive diagnostic evidence for gastroesophageal reflux disease: the validation of Lyon Consensus 2.0',
    'Jing Chen, Peiwen Dong, Songfeng Chen, Qianjun Zhuang, Mengyu Zhang, Kaidi Sun, Feng Tang, Qiong Wang, Yinglian Xiao',
    'Gastroenterology Report',
    '2025-03-12',
    'en',
    '10.1093/gastro/goaf004',
    '40083682',
    'This study evaluated whether LA-B esophagitis qualifies as definitive diagnostic evidence for GERD per Lyon Consensus 2.0. Among 401 patients with reflux symptoms, those with LA-B or LA-C/D esophagitis showed significantly higher acid-suppressive therapy response rates (75.0% and 82.7% respectively) compared to non-RE patients (52.4%). LA-A esophagitis showed no significant difference from non-RE patients, but when combined with specific supportive evidence—such as reflux episodes exceeding 80 daily or hypotensive esophagogastric junction—response rates improved substantially.',
    'research_article',
    'Gastroenterology',
    '["GERD", "esophagitis", "Los Angeles classification", "Lyon Consensus", "endoscopy", "acid reflux"]'::jsonb,
    '["Gastroesophageal Reflux", "Esophagitis", "Endoscopy", "Proton Pump Inhibitors"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id AS article1_id;

-- Article 2: LA-D Pathogenesis (2019)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    article_type,
    specialty,
    keywords,
    mesh_terms
) VALUES (
    gen_random_uuid(),
    'Unique Clinical Features of Los Angeles Grade D Esophagitis Suggest that Factors Other than Gastroesophageal Reflux Contribute to Its Pathogenesis',
    'Anh D Nguyen, Stuart J Spechler, Monique N Shuler, Rhonda F Souza, Kerry B Dunbar',
    'Journal of Clinical Gastroenterology',
    '2019-01-01',
    'en',
    '10.1097/MCG.0000000000000870',
    '28644313',
    'LA-D esophagitis primarily affects hospitalized older patients with lower BMIs and serious comorbidities, contrasting with LA-A patients who are younger, obese outpatients with GERD histories. The study compared 100 LA-D and 100 LA-A patients, finding LA-D patients were older (65 vs 56 years), had lower BMIs (25.9 vs 29.4), were more frequently hospitalized (70% vs 3%), and had more cardiopulmonary disorders. Conversely, GERD history was more common in LA-A patients (67% vs 45%). These differences suggest factors beyond typical GERD contribute to LA-D esophagitis pathogenesis, particularly in hospitalized patients.',
    'research_article',
    'Gastroenterology',
    '["esophagitis", "Los Angeles classification", "GERD", "pathogenesis", "hospitalized patients"]'::jsonb,
    '["Esophagitis", "Gastroesophageal Reflux", "Risk Factors", "Comorbidity"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id AS article2_id;

-- Article 3: LA Classification Validation (1999 - Estudo Seminal)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    article_type,
    specialty,
    keywords,
    mesh_terms
) VALUES (
    gen_random_uuid(),
    'Endoscopic assessment of oesophagitis: clinical and functional correlates and further validation of the Los Angeles classification',
    'L R Lundell, J Dent, J R Bennett, A L Blum, D Armstrong, J P Galmiche, F Johnson, M Hongo, J E Richter, S J Spechler, G N Tytgat, L Wallin',
    'Gut',
    '1999-08-01',
    'en',
    '10.1136/gut.45.2.172',
    '10403727',
    'Study evaluating the reliability of endoscopic criteria for describing reflux oesophagitis using the Los Angeles classification system. Forty-six endoscopists reviewed video recordings from 22 patients. Los Angeles gradings were correlated with pH monitoring data from 178 patients and omeprazole treatment outcomes from 277 patients. Results demonstrated that the Los Angeles system''s criterion for assessing circumferential extent yielded acceptable agreement (mean kappa value 0.4) among observers, while alternative percentage-based approaches showed unacceptably high variation. Oesophageal acid exposure severity and oesophagitis grades showed significant correlation with heartburn severity and treatment outcomes, supporting the clinical utility of the Los Angeles classification system.',
    'research_article',
    'Gastroenterology',
    '["Los Angeles classification", "esophagitis", "endoscopy", "GERD", "validation", "inter-observer agreement"]'::jsonb,
    '["Esophagitis", "Gastroesophageal Reflux", "Endoscopy", "Reproducibility of Results"]'::jsonb
) ON CONFLICT (doi) DO NOTHING
RETURNING id AS article3_id;

-- 2. UPDATE Score Item: Esofagite (LA Classification)
-- ===============================================

UPDATE score_items
SET
    clinical_relevance = 'A Classificação de Los Angeles (LA) é o sistema endoscópico padronizado internacionalmente para graduar a gravidade da esofagite erosiva secundária à doença do refluxo gastroesofágico (DRGE). O sistema classifica as lesões mucosas em quatro graus (A a D) com base na extensão das erosões: Grau A (erosões <5mm confinadas às pregas mucosas), Grau B (erosões >5mm nas pregas), Grau C (erosões entre 2+ pregas, <75% da circunferência) e Grau D (erosões ≥75% da circunferência). O Lyon Consensus 2.0 (2025) estabelece que LA-B é evidência diagnóstica conclusiva de DRGE, com taxa de resposta a IBP de 75-82%, enquanto LA-A requer confirmação com pHmetria. LA-D apresenta perfil clínico distinto, afetando predominantemente pacientes hospitalizados idosos com múltiplas comorbidades, sugerindo mecanismos patogênicos adicionais além do refluxo ácido típico. A classificação LA possui concordância interobservador aceitável (kappa 0.4) e correlação significativa com exposição ácida esofágica, sintomas de pirose e resposta ao tratamento.',

    patient_explanation = 'Este exame avalia inflamação e feridas no esôfago (tubo que leva comida da boca ao estômago) causadas pelo refluxo de ácido do estômago. O médico usa uma câmera (endoscopia) para visualizar o grau de lesão, classificado de A a D. **Grau A** são feridas pequenas (<5mm), geralmente leves. **Grau B** tem feridas maiores (>5mm) e confirma refluxo significativo. **Grau C** apresenta lesões que se espalham entre as pregas do esôfago. **Grau D** é o mais grave, com lesões extensas cobrindo mais de 75% da área - comum em pacientes hospitalizados com outras doenças. O resultado orienta o tempo de tratamento: casos leves (A/B) tratam 4 semanas com remédios que bloqueiam ácido (IBP), casos graves (C/D) precisam 8 semanas. Se você tem grau A, pode precisar de exames adicionais para confirmar refluxo.',

    conduct = 'Avaliar sintomas associados (pirose, regurgitação, disfagia, dor torácica). Grau LA-A: considerar pHmetria/pHimpedanciometria 24h para confirmação diagnóstica de DRGE, especialmente se sintomas atípicos. Grau LA-B: iniciar tratamento empírico com IBP dose padrão (omeprazol 20mg, pantoprazol 40mg ou esomeprazol 40mg) 1x/dia por 4 semanas. Grau LA-C/D: IBP dose padrão por 8 semanas seguido de endoscopia de controle; investigar fatores agravantes (hérnia hiatal grande, Barrett, tabagismo, obesidade). Em LA-D, avaliar comorbidades sistêmicas (cardiopulmonares, renais) e fatores não-refluxo (medicamentos ulcerogênicos, isquemia). Orientar medidas comportamentais: elevação cabeceira, evitar decúbito pós-prandial, dieta anti-refluxo, perda ponderal se IMC >25. Refratariedade após 8 semanas: considerar dose dobrada de IBP, manometria esofágica, pHmetria em uso de IBP e avaliação para fundoplicatura laparoscópica. Monitorar complicações (estenose, Barrett, hemorragia digestiva).',

    updated_at = NOW()
WHERE id = '4f6e007b-dcc2-4e51-aaad-b7359717f297';

-- 3. Criar relações Many-to-Many (article_score_items)
-- ===============================================

-- Buscar IDs dos articles recém-criados
WITH article_ids AS (
    SELECT id FROM articles WHERE doi IN (
        '10.1093/gastro/goaf004',
        '10.1097/MCG.0000000000000870',
        '10.1136/gut.45.2.172'
    )
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    article_ids.id,
    '4f6e007b-dcc2-4e51-aaad-b7359717f297'::uuid
FROM article_ids
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- 4. Verificar resultado
-- ===============================================

SELECT
    si.id,
    si.name,
    si.code,
    LENGTH(si.clinical_relevance) as relevance_chars,
    LENGTH(si.patient_explanation) as explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    COUNT(asi.article_id) as num_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '4f6e007b-dcc2-4e51-aaad-b7359717f297'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct;

-- Listar articles vinculados
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '4f6e007b-dcc2-4e51-aaad-b7359717f297'
ORDER BY a.publish_date DESC;

COMMIT;

-- ===============================================
-- SUMMARY
-- ===============================================
-- Item: Endoscopia Alta - Esofagite (LA Classification)
-- Articles: 3 (Lyon Consensus 2025, LA-D Pathogenesis 2019, Original Validation 1999)
-- Clinical Relevance: ~1400 chars
-- Patient Explanation: ~900 chars
-- Conduct: ~1200 chars
-- ===============================================
