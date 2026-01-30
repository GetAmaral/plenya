-- ============================================================================
-- ENRIQUECIMENTO: Densitometria - T-Score Colo Femoral
-- Item ID: 4ae093c8-b6ac-4ac1-b242-dcd64194e6ad
-- Data: 2026-01-28
-- ============================================================================

-- 1. INSERIR ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: Classification of Osteoporosis (Indian J Orthop 2023)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    abstract,
    article_type,
    keywords,
    specialty
) VALUES (
    'Classification of Osteoporosis',
    'S S Amarnath, Vishal Kumar, S Lakshmana Das',
    'Indian Journal of Orthopaedics',
    '2023-12-06',
    'en',
    '10.1007/s43465-023-01058-3',
    'The World Health Organization devised a BMD classification system utilizing T-scores for specific populations. T-score is defined as patient measured BMD value minus the reference BMD value divided by the reference standard deviation. T-scores apply to postmenopausal women and men aged 50 years and older. Conversely, Z-scores are preferred for premenopausal women, adults under 50, and children. The diagnostic approach emphasizes that bone mineral density measurement alone cannot diagnose osteoporosis in men under 50 years. The FRAX algorithm supplements BMD testing by incorporating clinical fracture risk predictors. Treatment is recommended when FRAX indicates a 10-year hip fracture probability of at least 3% or major osteoporotic fracture risk exceeding 20%.',
    'review_article',
    '["osteoporosis", "T-score", "bone mineral density", "WHO classification", "FRAX"]',
    'Ortopedia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 2: T-score cutpoints in Chinese population (QIMS 2022)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    abstract,
    article_type,
    keywords,
    specialty
) VALUES (
    'Estimations of bone mineral density defined osteoporosis prevalence and cutpoint T-score for defining osteoporosis among older Chinese population: a framework based on relative fragility fracture risks',
    'Yì Xiáng J Wáng, Ben-Heng Xiao, et al',
    'Quantitative Imaging in Medicine and Surgery',
    '2022-09-01',
    'en',
    '10.21037/qims-22-281',
    'The research proposes adjusted diagnostic thresholds for Chinese populations. The investigators estimate osteoporosis prevalence rates based on the assumption that fragility fracture risk in older Chinese individuals is approximately half that of Caucasian counterparts. Estimated osteoporosis prevalence (age ≥50 years): Chinese women: 7.5% (lumbar), 7.5% (femoral neck), 6.7% (total hip); Chinese men: 2.0% (lumbar), 3.8% (femoral neck), 3.4% (total hip). The authors recommend revised cutpoint T-scores rather than conventional WHO thresholds of −2.5 and −1.0, suggesting population-specific adjustments better reflect actual fracture risk in Chinese populations.',
    'research_article',
    '["bone mineral density", "osteoporosis", "T-score", "Chinese population", "ethnic differences", "fracture risk"]',
    'Epidemiologia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 3: Osteopenia (StatPearls 2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    abstract,
    article_type,
    keywords,
    specialty
) VALUES (
    'Osteopenia',
    'Matthew A. Varacallo, Travis J. Seaman, Jagmohan S. Jandu, Jasleen Kaur',
    'StatPearls [Internet]',
    '2025-12-01',
    'en',
    'Osteopenia refers to reduced bone mineral density below normal values without fulfilling the diagnostic threshold for osteoporosis, measured via dual-energy x-ray absorptiometry (DXA). T-Score Classification: Normal within 1 SD, Osteopenia between -1.0 and -2.5, Osteoporosis less than -2.5. Clinical significance: approximately 48-56% of fragility fractures in postmenopausal women occur in individuals with osteopenia-level bone density. Management: Low-to-moderate risk receives nonpharmacologic management (exercise, calcium/vitamin D supplementation). High risk receives pharmacologic therapy when 10-year hip fracture probability reaches 3% or major osteoporotic fracture risk exceeds 20% via FRAX assessment. Prevalence: 43.3 million American adults over 50 have osteopenia, affecting approximately 50% of women and 30% of men.',
    'book_chapter',
    '["osteopenia", "bone mineral density", "T-score", "fracture prevention", "FRAX", "osteoporosis"]',
    'Medicina Interna'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- 2. CRIAR RELAÇÕES ARTICLE-SCORE_ITEM
-- ============================================================================

-- Buscar IDs dos artigos recém-inseridos
WITH article_ids AS (
    SELECT id FROM articles WHERE doi IN (
        '10.1007/s43465-023-01058-3',
        '10.21037/qims-22-281'
    )
    UNION
    SELECT id FROM articles WHERE title = 'Osteopenia' AND journal = 'StatPearls [Internet]'
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    article_ids.id,
    '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad'::uuid
FROM article_ids
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- 3. ATUALIZAR SCORE_ITEM COM CONTEÚDO CLÍNICO
-- ============================================================================

UPDATE score_items SET
    clinical_relevance = 'O T-Score no colo femoral é um dos principais parâmetros diagnósticos da osteoporose segundo critérios da Organização Mundial da Saúde (OMS). Representa o número de desvios-padrão que a densidade mineral óssea (DMO) do paciente se afasta da média de adultos jovens saudáveis. Valores ≤ -2,5 confirmam osteoporose; entre -1,0 e -2,5 indicam osteopenia (baixa massa óssea); e > -1,0 são considerados normais. O colo femoral é um dos dois sítios anatômicos obrigatórios na densitometria (junto à coluna lombar L1-L4), pois é local de alta incidência de fraturas osteoporóticas, especialmente em idosos. A avaliação do T-Score permite estratificação de risco de fraturas por fragilidade, que causam alta morbimortalidade, especialmente a fratura de quadril (mortalidade de 20-30% em 1 ano). O algoritmo FRAX complementa o T-Score ao calcular risco de fratura em 10 anos, indicando tratamento farmacológico quando risco de fratura de quadril ≥3% ou de fraturas osteoporóticas maiores ≥20%. A densitometria deve ser repetida a cada 1-2 anos em pacientes de alto risco para monitorar progressão e resposta ao tratamento.',

    patient_explanation = 'O T-Score do colo do fêmur (osso da coxa próximo ao quadril) mede a "força" do seu osso comparando com a densidade óssea média de pessoas jovens e saudáveis. Imagine uma escala: quanto mais negativo o número, mais fraco está o osso. Um T-Score acima de -1,0 é considerado normal (osso saudável). Entre -1,0 e -2,5 indica osteopenia, que é uma "fase de alerta" onde o osso está mais fraco, mas ainda não é osteoporose. Abaixo de -2,5 já caracteriza osteoporose, aumentando muito o risco de fraturas em quedas simples. O colo do fêmur é uma região crítica porque é onde ocorrem as fraturas de quadril, que são graves e podem comprometer sua independência. Se seu T-Score estiver alterado, seu médico poderá recomendar suplementação de cálcio e vitamina D, exercícios de fortalecimento, mudanças na dieta e, em casos mais graves, medicamentos específicos para fortalecer os ossos e prevenir fraturas. É importante repetir o exame periodicamente para acompanhar a evolução.',

    conduct = 'T-Score > -1,0 (normal): Orientar estilo de vida saudável com exercícios resistidos 3x/semana, ingestão adequada de cálcio (1000-1200mg/dia) e vitamina D (800-1000UI/dia). Reavaliar em 5 anos ou antes se surgirem fatores de risco. || T-Score entre -1,0 e -2,5 (osteopenia): Calcular FRAX Score (risco de fratura em 10 anos). Se FRAX indicar baixo risco: medidas não-farmacológicas (exercício, cálcio 1200mg/dia, vitamina D 1000-2000UI/dia, evitar tabaco/álcool). Se FRAX alto (≥3% para quadril ou ≥20% para fraturas maiores): considerar tratamento farmacológico. Repetir densitometria em 1-2 anos. Investigar causas secundárias: hiperparatireoidismo, hipertireoidismo, uso de corticoides, deficiência de vitamina D. || T-Score ≤ -2,5 (osteoporose): Iniciar tratamento farmacológico imediato (bisfosfonatos, denosumabe, teriparatida ou abaloparatida conforme perfil). Suplementar cálcio e vitamina D. Solicitar dosagem sérica de vitamina D, PTH, TSH, creatinina. Avaliar risco de quedas e orientar prevenção (fisioterapia, adaptações domiciliares, revisão de medicamentos sedativos). Repetir densitometria em 1 ano para avaliar resposta ao tratamento. Se fratura prévia: investigar osteoporose secundária completa.',

    last_review = CURRENT_TIMESTAMP

WHERE id = '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad';

-- 4. VERIFICAÇÃO FINAL
-- ============================================================================

-- Verificar artigos inseridos
SELECT
    id,
    title,
    journal,
    publish_date,
    doi
FROM articles
WHERE doi IN (
    '10.1007/s43465-023-01058-3',
    '10.21037/qims-22-281'
)
OR (title = 'Osteopenia' AND journal = 'StatPearls [Internet]');

-- Verificar relações criadas
SELECT
    a.title,
    si.name
FROM article_score_items asi
JOIN articles a ON a.id = asi.article_id
JOIN score_items si ON si.id = asi.score_item_id
WHERE asi.score_item_id = '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad';

-- Verificar atualização do score_item
SELECT
    id,
    name,
    CASE
        WHEN clinical_relevance IS NOT NULL THEN 'OK'
        ELSE 'FALTA'
    END as clinical_relevance_status,
    CASE
        WHEN patient_explanation IS NOT NULL THEN 'OK'
        ELSE 'FALTA'
    END as patient_explanation_status,
    CASE
        WHEN conduct IS NOT NULL THEN 'OK'
        ELSE 'FALTA'
    END as conduct_status,
    last_review
FROM score_items
WHERE id = '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad';

-- ============================================================================
-- FIM DO SCRIPT
-- ============================================================================
