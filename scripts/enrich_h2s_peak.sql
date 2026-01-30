-- Enrichment for Score Item: Sulfeto de Hidrogênio Pico (H₂S Máximo)
-- ID: b87387b4-d024-4dbb-be70-84778ca2dce0
-- Generated: 2026-01-29

BEGIN;

-- 1. Update clinical content for H2S Peak score item
UPDATE score_items
SET
    clinical_relevance = 'O sulfeto de hidrogênio (H₂S) é um gás produzido por bactérias redutoras de sulfato no intestino delgado, representando um marcador diagnóstico crucial para subtipos específicos de supercrescimento bacteriano (SIBO). O valor de pico reflete a máxima produção bacteriana durante o teste respiratório com lactulose ou glicose, indicando a intensidade da colonização patológica. Níveis elevados (≥3 ppm) correlacionam-se fortemente com quadros de diarréia crônica, distensão abdominal e odor sulfuroso característico. Estudos recentes (2025) demonstraram que valores ≥1,5 ppm correlacionam-se com alterações significativas no microbioma duodenal, especialmente aumento de 2,08-log2fold no filo Thermodesulfobacteriota. A detecção de H₂S é particularmente importante em pacientes com testes "flat-line" (sem elevação de hidrogênio/metano), onde bactérias hidrogenotróficas consomem o hidrogênio produzido pela fermentação, convertendo-o em H₂S. O pico máximo auxilia na diferenciação entre SIBO-H₂, SIBO-CH₄ (IMO) e SIBO-H₂S, cada um com perfil clínico, microbiológico e terapêutico distintos. Valores ≥62,5 ppb aos 90 minutos apresentam sensibilidade de 66,4% e especificidade de 79,1% para diagnóstico de SIBO. A monitorização do H₂S amplia significativamente a capacidade diagnóstica, permitindo identificação de casos não detectáveis por testes convencionais de dois gases.',

    patient_explanation = 'O H₂S (sulfeto de hidrogênio) é um gás produzido por bactérias específicas no seu intestino delgado. Durante o teste respiratório, você ingere uma solução açucarada (lactulose ou glicose) e sopra em um aparelho especial em intervalos regulares por 3 horas. O "pico" representa o momento em que o aparelho detecta a maior quantidade de H₂S na sua respiração. Se você tem bactérias redutoras de sulfato em excesso no intestino delgado (onde não deveriam estar em grande quantidade), elas transformam o açúcar em H₂S, que é absorvido pela corrente sanguínea, chega aos pulmões e sai na sua respiração. Valores elevados (acima de 3 partes por milhão) indicam que você pode ter um tipo específico de SIBO chamado "SIBO-H₂S". Esse tipo geralmente causa diarréia mais frequente, gases com odor forte de "ovo podre" e distensão abdominal após refeições. O teste de H₂S é especialmente importante se você teve testes anteriores de SIBO normais mas continua com sintomas intestinais, pois algumas pessoas produzem principalmente H₂S ao invés de hidrogênio ou metano. O resultado ajuda seu médico a escolher o tratamento mais adequado para o seu tipo específico de problema.',

    conduct = 'INTERPRETAÇÃO: Valores <3 ppm: normal (ausência de SIBO-H₂S); ≥3 ppm: positivo para SIBO-H₂S (associado a diarréia); ≥1,5 ppm: alteração limítrofe (considerar contexto clínico); ≥62,5 ppb aos 90min: critério diagnóstico alternativo (sensibilidade 66,4%, especificidade 79,1%). INVESTIGAÇÃO COMPLEMENTAR: Realizar teste trio-smart completo (H₂+CH₄+H₂S) para identificar padrões combinados; solicitar cultura quantitativa de aspirado jejunal se disponível (padrão-ouro); avaliar fatores de risco (hipocloridria, dismotilidade, cirurgias prévias, diabetes, uso crônico de IBP); investigar causas subjacentes (esclerodermia, doença de Crohn, divertículos duodenais). PERFIL CLÍNICO: SIBO-H₂S tipicamente cursa com diarréia (vs. constipação no IMO), flatulência sulfurosa, distensão pós-prandial e má absorção. Pacientes podem apresentar teste "flat-line" para H₂/CH₄ devido ao consumo de hidrogênio por bactérias redutoras de sulfato. TRATAMENTO: Rifaximina 550mg 3x/dia por 14 dias ASSOCIADA a subsalicilato de bismuto 524mg 4x/dia (o bismuto inibe bactérias redutoras de sulfato); dieta elementar por 2-3 semanas como alternativa; dieta low-FODMAP associada a baixo teor de enxofre (reduzir ovos, carnes processadas, brócolis, alho, cebola); probióticos selecionados (evitar Lactobacillus que podem piorar H₂S). MONITORIZAÇÃO: Repetir teste respiratório 4-6 semanas após tratamento; avaliar resposta clínica (frequência evacuatória, distensão, odor fecal); considerar teste de permeabilidade intestinal. DIAGNÓSTICO DIFERENCIAL: Má absorção de bile, insuficiência pancreática exócrina, doença celíaca, colite microscópica. ENCAMINHAMENTO: Considerar gastroenterologista se refratariedade a tratamento padrão ou suspeita de doença estrutural subjacente.',

    updated_at = NOW()
WHERE id = 'b87387b4-d024-4dbb-be70-84778ca2dce0';

-- 2. Insert scientific articles

-- Article 1: Villanueva-Millan et al. 2025 (Digestive Diseases and Sciences)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
)
VALUES (
    'Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens',
    'Villanueva-Millan MJ, Leite G, Mathur R, Rezaie A, Moreno Fajardo C, de Freitas Germano J, Morales W, Sanchez M, Rivera I, Parodi G, Weitsman S, Rashid M, Hosseini A, Brimberry D, Barlow GM, Pimentel M',
    'Digestive Diseases and Sciences',
    '40569514',
    '10.1007/s10620-025-09156-y',
    '2025-01-01'::date,
    'research_article',
    'en',
    'Gastroenterology',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    specialty = EXCLUDED.specialty,
    updated_at = NOW();

-- Link Article 1 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    (SELECT id FROM articles WHERE doi = '10.1007/s10620-025-09156-y'),
    'b87387b4-d024-4dbb-be70-84778ca2dce0'::uuid
ON CONFLICT DO NOTHING;

-- Article 2: Tansel & Levinthal 2023 (Clinical and Translational Gastroenterology)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
)
VALUES (
    'Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth',
    'Tansel A, Levinthal DJ',
    'Clinical and Translational Gastroenterology',
    '36744854',
    '10.14309/ctg.0000000000000567',
    '2023-01-01'::date,
    'review',
    'en',
    'Gastroenterology',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    specialty = EXCLUDED.specialty,
    updated_at = NOW();

-- Link Article 2 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    (SELECT id FROM articles WHERE doi = '10.14309/ctg.0000000000000567'),
    'b87387b4-d024-4dbb-be70-84778ca2dce0'::uuid
ON CONFLICT DO NOTHING;

-- Article 3: Birg et al. 2019 (JGH Open)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
)
VALUES (
    'Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements',
    'Birg A, Hu S, Lin HC',
    'JGH Open: An Open Access Journal of Gastroenterology and Hepatology',
    '31276041',
    '10.1002/jgh3.12145',
    '2019-01-01'::date,
    'research_article',
    'en',
    'Gastroenterology',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    specialty = EXCLUDED.specialty,
    updated_at = NOW();

-- Link Article 3 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    (SELECT id FROM articles WHERE doi = '10.1002/jgh3.12145'),
    'b87387b4-d024-4dbb-be70-84778ca2dce0'::uuid
ON CONFLICT DO NOTHING;

-- Article 4: Guo et al. 2021 (Chinese Journal of Internal Medicine)
INSERT INTO articles (
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
)
VALUES (
    'The diagnostic value of hydrogen sulfide breath test for small intestinal bacterial overgrowth',
    'Guo HZ, Dong WX, Zhang X, Zhu SW, Liu ZJ, Duan LP',
    'Zhonghua Nei Ke Za Zhi (Chinese Journal of Internal Medicine)',
    '33765706',
    '10.3760/cma.j.cn112138-20200731-00725',
    '2021-01-01'::date,
    'research_article',
    'en',
    'Gastroenterology',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE
SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    specialty = EXCLUDED.specialty,
    updated_at = NOW();

-- Link Article 4 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    (SELECT id FROM articles WHERE doi = '10.3760/cma.j.cn112138-20200731-00725'),
    'b87387b4-d024-4dbb-be70-84778ca2dce0'::uuid
ON CONFLICT DO NOTHING;

COMMIT;

-- Verification queries
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.updated_at
FROM score_items si
WHERE si.id = 'b87387b4-d024-4dbb-be70-84778ca2dce0';

SELECT
    si.name as score_item,
    a.title as article_title,
    a.pm_id,
    a.doi,
    a.publish_date,
    a.article_type
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'b87387b4-d024-4dbb-be70-84778ca2dce0'
ORDER BY a.publish_date DESC;
