-- ============================================================================
-- ENRICHMENT: Vitamina E (Alfa-Tocoferol) - Score Item ID: 80ffc2b2-545b-4389-891f-b6aba1f7865c
-- ============================================================================
-- Generated: 2026-01-29
-- Purpose: Add clinical content and scientific articles for Vitamin E score item
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: Vitamin E (α-Tocopherol): Emerging Clinical Role and Adverse Risks of Supplementation in Adults (2025)
INSERT INTO articles (id, title, article_type, pm_id, publish_date, authors, journal, doi)
VALUES (
    gen_random_uuid(),
    'Vitamin E (α-Tocopherol): Emerging Clinical Role and Adverse Risks of Supplementation in Adults',
    'review',
    '40065887',
    '2025-02-07'::date,
    'Not specified',
    'Cureus',
    '10.7759/cureus.78679'
) ON CONFLICT (doi) DO NOTHING;

-- Article 2: Systematic review with meta-analysis: The effect of vitamin E supplementation in adult patients with non-alcoholic fatty liver disease (2021)
INSERT INTO articles (id, title, article_type, pm_id, publish_date, authors, journal, doi)
VALUES (
    gen_random_uuid(),
    'Systematic review with meta-analysis: The effect of vitamin E supplementation in adult patients with non-alcoholic fatty liver disease',
    'meta_analysis',
    '32810309',
    '2021-02-01'::date,
    'Not specified',
    'Journal of Gastroenterology and Hepatology',
    '10.1111/jgh.15216'
) ON CONFLICT (doi) DO NOTHING;

-- Article 3: Antioxidant-independent activities of alpha-tocopherol (2025)
INSERT INTO articles (id, title, article_type, pm_id, publish_date, authors, journal, doi)
VALUES (
    gen_random_uuid(),
    'Antioxidant-independent activities of alpha-tocopherol',
    'research_article',
    '39978678',
    '2025-04-01'::date,
    'Not specified',
    'Journal of Biological Chemistry',
    '10.1016/j.jbc.2025.108327'
) ON CONFLICT (doi) DO NOTHING;

-- Article 4: Alpha-Tocopherol Metabolites (the Vitamin E Metabolome) and Their Interindividual Variability during Supplementation (2021)
INSERT INTO articles (id, title, article_type, pm_id, publish_date, authors, journal, doi)
VALUES (
    gen_random_uuid(),
    'Alpha-Tocopherol Metabolites (the Vitamin E Metabolome) and Their Interindividual Variability during Supplementation',
    'research_article',
    '33503988',
    '2021-01-25'::date,
    'Not specified',
    'Antioxidants (Basel)',
    '10.3390/antiox10020173'
) ON CONFLICT (doi) DO NOTHING;

-- ============================================================================
-- STEP 2: Link Articles to Score Item
-- ============================================================================

-- Link all articles to the Vitamina E score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '80ffc2b2-545b-4389-891f-b6aba1f7865c'::uuid
FROM articles a
WHERE a.pm_id IN ('40065887', '32810309', '39978678', '33503988')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- STEP 3: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A vitamina E (alfa-tocoferol) é o principal antioxidante lipossolúvel do organismo, essencial para proteção celular contra estresse oxidativo e danos causados por radicais livres. Sua ação se estende além da função antioxidante clássica, incluindo modulação da expressão gênica, metabolismo lipídico e funções neuromusculares. O alfa-tocoferol é crítico para manutenção da integridade de membranas celulares, especialmente em tecidos de alta demanda metabólica como sistema nervoso, muscular e cardiovascular. Valores séricos normais situam-se entre 5-18 mg/L, sendo que níveis abaixo de 5 mg/L indicam deficiência, embora a interpretação deva considerar o perfil lipídico, utilizando a relação alfa-tocoferol/lipídios totais (<0,8 mg/g) como indicador mais preciso em pacientes com hiperlipidemia. A deficiência verdadeira é rara em indivíduos saudáveis, ocorrendo principalmente em contextos de má absorção de gorduras (doença celíaca, fibrose cística, colestase crônica, doença de Crohn), abetalipoproteinemia e síndromes de malabsorção grave. Manifestações clínicas incluem neuropatia periférica progressiva, ataxia espinocerebelar, oftalmoplegia, retinopatia, fraqueza muscular profunda e, em casos extremos, cegueira e arritmias cardíacas. O papel terapêutico do alfa-tocoferol tem sido extensivamente estudado em esteatohepatite não alcoólica (NASH), onde metanálises demonstram redução significativa de transaminases (ALT, AST), melhora histológica de esteatose e inflamação lobular, embora sem impacto consistente na fibrose hepática. Estudos recentes revelaram atividades biológicas independentes de antioxidação, incluindo modulação da biossíntese de colesterol e regulação de genes específicos, com profundas implicações clínicas e translacionais.',

    patient_explanation = 'A vitamina E é uma vitamina encontrada em alimentos como óleos vegetais, nozes, sementes e vegetais de folhas verdes. Ela atua como um escudo protetor para as células do seu corpo, defendendo-as contra danos causados por substâncias chamadas radicais livres. Essa vitamina é especialmente importante para o bom funcionamento dos nervos, músculos e para manter seu sistema imunológico forte. O exame de sangue mede quanto de vitamina E você tem no corpo. Valores normais geralmente ficam entre 5 e 18 mg/L. A deficiência de vitamina E é rara em pessoas saudáveis, mas pode acontecer se você tem problemas para absorver gorduras da alimentação, como em doenças intestinais crônicas. Quando falta vitamina E, você pode apresentar fraqueza muscular, dificuldade para coordenar movimentos, formigamentos nas mãos e pés, problemas de visão e, em casos graves, dificuldade para andar. Se seus níveis estiverem baixos, o médico pode recomendar suplementação com cápsulas de vitamina E em doses específicas. É importante não tomar suplementos por conta própria em doses muito altas, pois isso pode aumentar o risco de sangramentos. A vitamina E também tem sido estudada para tratamento de problemas no fígado relacionados ao acúmulo de gordura (esteatose hepática), mostrando benefícios quando usada sob orientação médica.',

    conduct = 'Avaliação inicial: solicitar dosagem sérica de alfa-tocoferol em jejum de 8-12 horas, preferencialmente junto com perfil lipídico completo para cálculo da relação alfa-tocoferol/lipídios totais (normal >0,8 mg/g). Valores <5 mg/L ou relação <0,8 mg/g sugerem deficiência e exigem investigação de causa subjacente. Investigação etiológica: pesquisar condições de má absorção intestinal (dosagem de elastase fecal, gordura fecal, testes celíacos, calprotectina), doenças hepatobiliares (perfil hepático, ultrassom abdominal), fibrose cística (teste do suor, genética), abetalipoproteinemia (perfil lipídico com apolipoproteínas) e uso de medicamentos quelantes de gordura. Avaliar manifestações neurológicas com exame físico detalhado: reflexos osteotendinosos, propriocepção, marcha, movimentos oculares, força muscular e sensibilidade. Em casos com sinais neurológicos, considerar ressonância magnética de crânio/medula e eletroneuromiografia. Tratamento da deficiência: corrigir causa-base sempre que possível. Suplementação oral com RRR-alfa-tocoferol (forma natural mais biodisponível): adultos 15-25 mg/dia (22,5-37,5 UI/dia) em casos leves; 100-800 UI/dia em casos moderados; 100-200 UI/kg/dia em abetalipoproteinemia ou deficiências graves com manifestações neurológicas. Associar com refeições contendo gordura para otimizar absorção. Monitoramento: dosar alfa-tocoferol sérico após 3-6 meses de suplementação, ajustar dose conforme resposta. Em pacientes com manifestações neurológicas, reavaliar clinicamente a cada 3 meses e realizar exames de imagem de controle após 6-12 meses. Uso em NASH: considerar vitamina E 800 UI/dia em pacientes adultos não diabéticos, sem cirrose, com NASH confirmada por biópsia e histologia agressiva, por período mínimo de 18-24 meses, com monitoramento trimestral de transaminases. Contraindicações e precauções: evitar doses >1000 mg/dia pelo risco de AVC hemorrágico e efeitos pró-oxidantes. Cautela em pacientes anticoagulados (aumenta risco de sangramento). Não indicar para prevenção primária de doenças cardiovasculares ou câncer conforme diretrizes USPSTF 2022. Reavaliação periódica a cada 6-12 meses mesmo após normalização dos níveis.',

    last_review = CURRENT_DATE
WHERE id = '80ffc2b2-545b-4389-891f-b6aba1f7865c';

-- ============================================================================
-- STEP 4: Verification Queries
-- ============================================================================

-- Verify articles were inserted
SELECT
    a.pm_id,
    a.title,
    a.article_type,
    a.publish_date,
    COUNT(asi.score_item_id) as linked_items
FROM articles a
LEFT JOIN article_score_items asi ON a.id = asi.article_id
WHERE a.pm_id IN ('40065887', '32810309', '39978678', '33503988')
GROUP BY a.pm_id, a.title, a.article_type, a.publish_date
ORDER BY a.publish_date DESC;

-- Verify score item enrichment
SELECT
    si.id,
    si.description,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(DISTINCT asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '80ffc2b2-545b-4389-891f-b6aba1f7865c'
GROUP BY si.id, si.description, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- List all linked articles with full details
SELECT
    si.description as score_item,
    a.title as article_title,
    a.article_type,
    a.pm_id,
    a.publish_date,
    'https://pubmed.ncbi.nlm.nih.gov/' || a.pm_id || '/' as pubmed_url
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '80ffc2b2-545b-4389-891f-b6aba1f7865c'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- Character Count Validation
-- ============================================================================
-- Expected character counts:
-- clinical_relevance: 1500-2000 chars (Target: ~1850)
-- patient_explanation: 1000-1500 chars (Target: ~1320)
-- conduct: 1500-2500 chars (Target: ~2350)
-- ============================================================================
