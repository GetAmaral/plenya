-- ============================================================================
-- CORREÇÃO: Inserir 2 artigos restantes com tipos válidos
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
    'review',
    '["osteoporosis", "T-score", "bone mineral density", "WHO classification", "FRAX"]',
    'Ortopedia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Artigo 3: Osteopenia (StatPearls 2025)
-- Como não tem DOI, vamos usar pm_id para evitar duplicatas
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    pm_id,
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
    'NBK499878',
    'Osteopenia refers to reduced bone mineral density below normal values without fulfilling the diagnostic threshold for osteoporosis, measured via dual-energy x-ray absorptiometry (DXA). T-Score Classification: Normal within 1 SD, Osteopenia between -1.0 and -2.5, Osteoporosis less than -2.5. Clinical significance: approximately 48-56% of fragility fractures in postmenopausal women occur in individuals with osteopenia-level bone density. Management: Low-to-moderate risk receives nonpharmacologic management (exercise, calcium/vitamin D supplementation). High risk receives pharmacologic therapy when 10-year hip fracture probability reaches 3% or major osteoporotic fracture risk exceeds 20% via FRAX assessment. Prevalence: 43.3 million American adults over 50 have osteopenia, affecting approximately 50% of women and 30% of men.',
    'review',
    '["osteopenia", "bone mineral density", "T-score", "fracture prevention", "FRAX", "osteoporosis"]',
    'Medicina Interna'
)
ON CONFLICT (pm_id) DO NOTHING
RETURNING id;

-- Criar relações para os 2 novos artigos
WITH article_ids AS (
    SELECT id FROM articles WHERE doi = '10.1007/s43465-023-01058-3'
    UNION
    SELECT id FROM articles WHERE pm_id = 'NBK499878'
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    article_ids.id,
    '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad'::uuid
FROM article_ids
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- VERIFICAÇÃO FINAL COMPLETA
SELECT
    '=== ARTIGOS INSERIDOS ===' as section,
    COUNT(*) as total
FROM articles
WHERE doi IN ('10.1007/s43465-023-01058-3', '10.21037/qims-22-281')
   OR pm_id = 'NBK499878';

SELECT
    id,
    title,
    journal,
    publish_date,
    COALESCE(doi, pm_id) as identifier
FROM articles
WHERE doi IN ('10.1007/s43465-023-01058-3', '10.21037/qims-22-281')
   OR pm_id = 'NBK499878';

SELECT
    '=== RELAÇÕES CRIADAS ===' as section,
    COUNT(*) as total
FROM article_score_items
WHERE score_item_id = '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad';

SELECT
    a.title as article_title,
    si.name as score_item_name
FROM article_score_items asi
JOIN articles a ON a.id = asi.article_id
JOIN score_items si ON si.id = asi.score_item_id
WHERE asi.score_item_id = '4ae093c8-b6ac-4ac1-b242-dcd64194e6ad'
ORDER BY a.publish_date DESC;
