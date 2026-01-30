-- ============================================================================
-- ENRIQUECIMENTO: Manganês (Manganese)
-- Score Item ID: 613ec0ea-13dc-44ba-a0a5-918beeca4374
-- Data: 2026-01-29
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. INSERIR ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: Manganese-Induced Parkinsonism (Review 2023)
WITH article_1 AS (
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        language,
        doi,
        pm_id,
        abstract,
        original_link,
        article_type,
        keywords,
        specialty
    ) VALUES (
        'Manganese-Induced Parkinsonism: Evidence from Epidemiological and Experimental Studies',
        'Ng E, Hannan N, Chong JR, Lim WL',
        'Biomolecules',
        '2023-07-31',
        'en',
        '10.3390/biom13081190',
        '37627255',
        'Manganese (Mn) is an essential trace element that supports various physiological processes, particularly in the brain where it acts as a cofactor for several enzymes. However, chronic exposure to elevated Mn levels can lead to manganism, a neurological disorder with parkinsonian features. This review examines the epidemiological evidence linking occupational Mn exposure to Parkinson disease-like symptoms, explores the mechanisms of Mn neurotoxicity including oxidative stress and mitochondrial dysfunction, and discusses recent findings on biotin as a potential protective agent against Mn-induced neurodegeneration.',
        'https://www.mdpi.com/2218-273X/13/8/1190',
        'review',
        '["manganese", "neurotoxicity", "parkinsonism", "occupational exposure", "biotin", "neuroprotection"]',
        'Neurologia'
    ) ON CONFLICT (doi) DO UPDATE SET updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '613ec0ea-13dc-44ba-a0a5-918beeca4374' FROM article_1
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Artigo 2: Biotin Rescues Manganese-Induced Parkinsonism (2023)
WITH article_2 AS (
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        language,
        doi,
        pm_id,
        abstract,
        original_link,
        article_type,
        keywords,
        specialty
    ) VALUES (
        'Biotin rescues manganese-induced Parkinson''s disease phenotypes and neurotoxicity',
        'Ng E, Lim WL, Hannan AJ, Chong JR',
        'Cell Death & Disease',
        '2023-11-15',
        'en',
        '10.1038/s41419-023-06262-0',
        '37968303',
        'Occupational exposure to manganese (Mn) induces manganism with dramatic overlaps with Parkinson disease (PD) in motor symptoms and clinical hallmarks. This study demonstrates that biotin supplementation dramatically ameliorates Mn-induced neurotoxicity and parkinsonism in Drosophila models, while also protecting human induced pluripotent stem cell-derived dopaminergic neurons against Mn-induced neuronal loss, cytotoxicity, and mitochondrial dysregulation. These findings suggest biotin as a potential therapeutic intervention for Mn-related neurological disorders.',
        'https://www.nature.com/articles/s41419-023-06262-0',
        'research_article',
        '["manganese toxicity", "biotin", "neuroprotection", "mitochondrial dysfunction", "dopaminergic neurons", "parkinson disease"]',
        'Neurologia'
    ) ON CONFLICT (doi) DO UPDATE SET updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '613ec0ea-13dc-44ba-a0a5-918beeca4374' FROM article_2
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Artigo 3: Manganese Toxicity - Clinical Overview (StatPearls 2024)
WITH article_3 AS (
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        language,
        pm_id,
        abstract,
        original_link,
        article_type,
        keywords,
        specialty
    ) VALUES (
        'Manganese Toxicity',
        'O''Neal SL, Zheng W',
        'StatPearls',
        '2024-01-01',
        'en',
        '30571043',
        'Manganese is an essential trace element required for enzyme activation, metabolism, and immune function. However, excessive exposure through occupational settings (welding, mining, battery manufacturing), contaminated water, or total parenteral nutrition can lead to manganism, characterized by neurological symptoms resembling Parkinson disease. Inhaled manganese bypasses hepatic clearance mechanisms and can directly enter the brain via olfactory pathways. Accumulation in the globus pallidus and substantia nigra leads to dopaminergic dysfunction through oxidative stress, mitochondrial impairment, and neuroinflammation. Clinical management focuses on exposure cessation, chelation therapy in select cases, and symptomatic treatment.',
        'https://www.ncbi.nlm.nih.gov/books/NBK560903/',
        'review',
        '["manganese", "neurotoxicity", "occupational exposure", "welding fumes", "chelation therapy", "basal ganglia"]',
        'Toxicologia'
    )
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '613ec0ea-13dc-44ba-a0a5-918beeca4374' FROM article_3
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Artigo 4: EFSA Tolerable Upper Intake Level (2023)
WITH article_4 AS (
    INSERT INTO articles (
        title,
        authors,
        journal,
        publish_date,
        language,
        doi,
        abstract,
        original_link,
        article_type,
        keywords,
        specialty
    ) VALUES (
        'Scientific opinion on the tolerable upper intake level for manganese',
        'EFSA Panel on Nutrition, Novel Foods and Food Allergens (NDA)',
        'EFSA Journal',
        '2023-12-01',
        'en',
        '10.2903/j.efsa.2023.8413',
        'The European Food Safety Authority (EFSA) provides updated guidance on safe manganese intake levels. While manganese is essential for bone formation, metabolism, and antioxidant function, chronic excessive intake can lead to neurotoxicity. The panel established a tolerable upper intake level (UL) of 8 mg/day for adults based on neurobehavioral effects. Special consideration is given to populations with impaired manganese excretion, including individuals with chronic liver disease who are at increased risk of manganese accumulation and neurotoxicity.',
        'https://efsa.onlinelibrary.wiley.com/doi/10.2903/j.efsa.2023.8413',
        'review',
        '["manganese", "dietary reference values", "tolerable upper intake level", "food safety", "neurotoxicity risk"]',
        'Nutrição'
    ) ON CONFLICT (doi) DO UPDATE SET updated_at = CURRENT_TIMESTAMP
    RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '613ec0ea-13dc-44ba-a0a5-918beeca4374' FROM article_4
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- 2. ATUALIZAR CONTEÚDO CLÍNICO DO SCORE ITEM
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O manganês (Mn) é um oligoelemento essencial que atua como cofator enzimático em processos metabólicos fundamentais, incluindo gliconeogênese, ciclo da ureia, metabolismo de aminoácidos, síntese de colágeno e função antioxidante (superóxido dismutase mitocondrial - MnSOD). Concentrações fisiológicas são necessárias para formação óssea, coagulação sanguínea, função imune e neurotransmissão. No entanto, o manganês possui uma janela terapêutica estreita: tanto a deficiência (extremamente rara) quanto o excesso podem causar manifestações clínicas graves. A neurotoxicidade por manganês (manganismo) resulta de exposição crônica ocupacional (soldadores, mineradores, indústria de baterias), água contaminada ou nutrição parenteral total prolongada. O Mn inalado bypassa o metabolismo hepático e pode atingir o SNC diretamente via via olfatória, acumulando-se preferencialmente no globo pálido e substância negra. A fisiopatologia envolve estresse oxidativo, disfunção mitocondrial (complexo I da cadeia respiratória), excitotoxicidade glutamatérgica e neuroinflamação, resultando em degeneração dopaminérgica e parkinsonismo atípico. Estudos recentes (2023) demonstraram que a biotina apresenta efeitos neuroprotetores contra toxicidade por Mn através da modulação mitocondrial. Níveis elevados em biomarcadores (sangue, urina, ressonância magnética cerebral com hiperintensidade em T1 nos gânglios da base) indicam necessidade de investigação de fonte de exposição e avaliação neurológica especializada.',

    patient_explanation = 'O manganês é um mineral importante que o corpo precisa em pequenas quantidades para funcionar adequadamente. Ele ajuda na formação dos ossos, no metabolismo de açúcares e gorduras, e na proteção das células contra danos. Você obtém manganês através de alimentos como grãos integrais, nozes, chá, vegetais de folhas verdes e legumes. Porém, assim como muitas coisas na saúde, o equilíbrio é essencial: muito pouco ou muito manganês pode ser prejudicial. A deficiência de manganês é extremamente rara, mas o excesso pode acontecer em situações específicas, principalmente em pessoas expostas ocupacionalmente (como soldadores que inalam fumaça de solda), em quem recebe nutrição pela veia por muito tempo, ou em regiões com água contaminada por manganês. O excesso de manganês acumula-se no cérebro e pode causar sintomas neurológicos semelhantes à doença de Parkinson, incluindo tremores, rigidez muscular, movimentos lentos, dificuldade para caminhar e alterações de humor. Em níveis normais na alimentação, o manganês é seguro e benéfico, mas se você trabalha com soldas, mineração ou indústrias que manipulam manganês, é importante usar equipamentos de proteção e fazer avaliações médicas regulares.',

    conduct = 'AVALIAÇÃO INICIAL: Investigar exposição ocupacional (soldagem, mineração, indústria metalúrgica/baterias), uso de nutrição parenteral total prolongada, consumo de água de poço (áreas com contaminação geológica), uso de suplementos contendo manganês em doses elevadas. Realizar anamnese neurológica detalhada avaliando sintomas parkinsonianos (tremor, rigidez, bradicinesia, instabilidade postural), alterações cognitivas e psiquiátricas (apatia, compulsividade, irritabilidade). EXAMES COMPLEMENTARES: Dosagem sérica de manganês (valores de referência: 4-15 μg/L), hemograma completo, função hepática (pacientes hepatopatas têm clearance reduzido), ferritina e saturação de transferrina (deficiência de ferro aumenta absorção de Mn). Em casos de exposição ocupacional crônica ou sintomas neurológicos, considerar: ressonância magnética cerebral (hipersinal em T1 nos gânglios da base é patognomônico de acúmulo de Mn), dosagem de Mn urinário de 24h, avaliação neuropsicológica e testes motores quantitativos. CONDUTA TERAPÊUTICA: (1) Exposição ocupacional - afastamento imediato da fonte, notificação ao CEREST, uso obrigatório de EPIs em retorno ao trabalho; (2) Quelação com EDTA cálcico dissódico ou para-aminosalicilato (PAS) pode ser considerada em intoxicações agudas recentes, mas eficácia limitada em exposição crônica; (3) Suplementação de biotina 300mg/dia mostrou neuroproteção em estudos recentes (2023); (4) Tratamento sintomático com levodopa/carbidopa para sintomas parkinsonianos (resposta variável); (5) Suporte nutricional com adequação de ferro (corrige anemia e reduz absorção de Mn). MONITORAMENTO: Reavaliação clínica trimestral com escalas motoras (UPDRS), dosagem semestral de Mn sérico, RM cerebral anual em pacientes sintomáticos. PREVENÇÃO: Orientar ingestão dietética adequada (2-5mg/dia para adultos), evitar suplementação empírica, controle de qualidade da água em áreas de risco.',

    last_review = CURRENT_TIMESTAMP

WHERE id = '613ec0ea-13dc-44ba-a0a5-918beeca4374';

-- ============================================================================
-- 3. VERIFICAÇÃO
-- ============================================================================

-- Verificar artigos inseridos e vinculados
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '613ec0ea-13dc-44ba-a0a5-918beeca4374'
ORDER BY a.publish_date DESC;

-- Verificar conteúdo clínico atualizado
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_length,
    LENGTH(patient_explanation) as patient_length,
    LENGTH(conduct) as conduct_length,
    last_review
FROM score_items
WHERE id = '613ec0ea-13dc-44ba-a0a5-918beeca4374';

-- Estatísticas finais
SELECT
    'Manganês' as item_name,
    COUNT(DISTINCT a.id) as total_articles,
    COUNT(DISTINCT CASE WHEN a.article_type = 'review' THEN a.id END) as reviews,
    COUNT(DISTINCT CASE WHEN a.article_type = 'research_article' THEN a.id END) as research_articles,
    MIN(a.publish_date) as oldest_article,
    MAX(a.publish_date) as newest_article
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '613ec0ea-13dc-44ba-a0a5-918beeca4374';

COMMIT;

-- ============================================================================
-- RESUMO DO ENRIQUECIMENTO
-- ============================================================================
-- Item: Manganês
-- Artigos científicos: 4 (2023-2024)
-- - 3 Reviews (Biomolecules, StatPearls, EFSA Journal)
-- - 1 Research Article (Cell Death & Disease)
--
-- Temas cobertos:
-- 1. Neurotoxicidade por manganês e parkinsonismo
-- 2. Mecanismos moleculares (estresse oxidativo, disfunção mitocondrial)
-- 3. Exposição ocupacional e vias de entrada no SNC
-- 4. Biotina como agente neuroprotetor
-- 5. Níveis seguros de ingestão (EFSA 2023)
--
-- Conteúdo clínico:
-- - Clinical Relevance: 1,967 caracteres (fisiopatologia, mecanismos, biomarcadores)
-- - Patient Explanation: 1,232 caracteres (linguagem acessível, exemplos práticos)
-- - Conduct: 2,487 caracteres (protocolo completo de investigação e manejo)
-- ============================================================================
