-- =====================================================
-- ADICIONAR MAIS ARTIGOS CIENTÍFICOS: ScoreItem "6 meses"
-- ID: c77cedd3-2800-7db3-aa91-68e188fa8864
-- Data: 2026-02-16
-- =====================================================

BEGIN;

-- Artigo 3: Enhancing Therapy Adherence
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
    keywords,
    specialty,
    embedding_status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Enhancing Therapy Adherence: Impact on Clinical Outcomes, Healthcare Costs, and Patient Quality of Life',
    'Urszula Religioni, Rocío Barrios-Rodríguez, Pilar Requena, Mariola Borowska, Janusz Ostrowski',
    'Medicina (Kaunas)',
    '2025-01-01',
    'en',
    '10.3390/medicina61010153',
    '39859135',
    'Medication adherence substantially improves disease management across chronic conditions. In diabetes, adherence enhances glycemic control and reduces retinopathy and nephropathy complications. For hypertension, adherence achieves target BP levels, decreasing stroke and cardiac event risk. Non-adherence costs USD 100-300 billion annually in the US through avoidable hospitalizations. Adherent patients report enhanced quality of life through improved physical functioning, reduced psychological stress, and better social participation.',
    'research_article',
    ARRAY['medication adherence', 'chronic disease', 'clinical outcomes', 'healthcare costs', 'quality of life'],
    'Internal Medicine',
    'pending',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW();

INSERT INTO article_score_items (
    score_item_id,
    article_id,
    confidence_score,
    auto_linked,
    linked_at
)
SELECT
    'c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid,
    id,
    0.88,
    true,
    NOW()
FROM articles
WHERE doi = '10.3390/medicina61010153'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Artigo 4: Long-Term Adherence to Health Behavior Change
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
    keywords,
    specialty,
    embedding_status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Long-Term Adherence to Health Behavior Change',
    'Kathryn R Middleton, Stephen D Anton, Michal G Perri',
    'American Journal of Lifestyle Medicine',
    '2013-11-01',
    'en',
    '10.1177/1559827613488867',
    '27547170',
    'Poor adherence to behaviors recommended in lifestyle interventions is widespread, particularly over the long-term. This review examines adherence difficulties in weight management, dietary reduction, and increased physical activity. Strategies for improving adherence include extended care provision, skills training programs, enhanced social support systems, and targeted approaches for maintaining dietary and physical activity changes. Practical provider-focused solutions address the adherence problem that significantly constrains intervention success.',
    'review',
    ARRAY['health behavior', 'adherence', 'lifestyle intervention', 'weight management', 'behavioral change'],
    'Preventive Medicine',
    'pending',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW();

INSERT INTO article_score_items (
    score_item_id,
    article_id,
    confidence_score,
    auto_linked,
    linked_at
)
SELECT
    'c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid,
    id,
    0.85,
    true,
    NOW()
FROM articles
WHERE doi = '10.1177/1559827613488867'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Artigo 5: Motivational Interviewing for Medication Adherence
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
    keywords,
    specialty,
    embedding_status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Motivational interviewing to support medication adherence in adults with chronic conditions: Systematic review of randomized controlled trials',
    'Marlène Papus, Alexandra L Dima, Marie Viprey, Anne-Marie Schott, Marie Paule Schneider, Teddy Novais',
    'Patient Education and Counseling',
    '2022-11-01',
    'en',
    '10.1016/j.pec.2022.06.013',
    '35779984',
    'Systematic review of 54 RCTs examining motivational interviewing effectiveness for supporting medication adherence in adults with chronic diseases. Most interventions were in infectious diseases (16), cardiology (14), psychiatry (8), and endocrinology (7). Medication adherence showed significant improvement in 23 RCTs, and other clinical outcomes improved in 19 RCTs including behavioral and symptom improvements. MI demonstrates increasing evidence base across clinical domains.',
    'meta_analysis',
    ARRAY['motivational interviewing', 'medication adherence', 'chronic disease', 'systematic review', 'patient engagement'],
    'Patient Education',
    'pending',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW();

INSERT INTO article_score_items (
    score_item_id,
    article_id,
    confidence_score,
    auto_linked,
    linked_at
)
SELECT
    'c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid,
    id,
    0.93,
    true,
    NOW()
FROM articles
WHERE doi = '10.1016/j.pec.2022.06.013'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Artigo 6: Technology-delivered MI
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
    keywords,
    specialty,
    embedding_status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Technology-delivered motivational interviewing to improve health outcomes in patients with chronic conditions: a systematic review of the literature',
    'Sara T Ahmed, Victoria Sanchez, Emily Bacon, Terrill Bravender, Lisa M Shook',
    'British Journal of Health Psychology',
    '2022-11-01',
    'en',
    '10.1111/bjhp.12620',
    '35943381',
    'Systematic review examining technology-delivered adaptations of motivational interviewing (TAMIs) for chronic disease management. 34 studies reported TAMI use, with sample sizes 10-2069 participants aged 13-70 years. Remote MI demonstrated promise for improving depression in chronic disease patients. Technology-delivered approaches enable scalable patient engagement interventions with demonstrated effectiveness across multiple chronic conditions.',
    'review',
    ARRAY['motivational interviewing', 'technology', 'chronic disease', 'telemedicine', 'remote care'],
    'Health Psychology',
    'pending',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW();

INSERT INTO article_score_items (
    score_item_id,
    article_id,
    confidence_score,
    auto_linked,
    linked_at
)
SELECT
    'c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid,
    id,
    0.87,
    true,
    NOW()
FROM articles
WHERE doi = '10.1111/bjhp.12620'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Artigo 7: MI Evidence-Based Approach
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
    keywords,
    specialty,
    embedding_status,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Motivational Interviewing: An Evidence-Based Approach for Use in Medical Practice',
    'Carmen A Taveras Alam, Gregory D Brown, Yancey R Hahn',
    'International Journal of General Medicine',
    '2021-06-01',
    'en',
    '10.2147/IJGM.S314874',
    '34135616',
    'Motivational interviewing (MI) is an evidence-based counseling approach with demonstrated efficacy across medical specialties. Meta-analyses reveal statistically significant mean intervention effects (OR: 1.55; 95% CI: 1.40-1.71) compared to standard treatment. Significant effects reported for substance consumption, physical activity, dental hygiene, body weight, treatment adherence, willingness to change behavior, and mortality. MI integrates with shared decision-making to encourage motivation for change in ambivalent patients.',
    'review',
    ARRAY['motivational interviewing', 'patient counseling', 'behavior change', 'medical practice', 'evidence-based'],
    'General Medicine',
    'pending',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW();

INSERT INTO article_score_items (
    score_item_id,
    article_id,
    confidence_score,
    auto_linked,
    linked_at
)
SELECT
    'c77cedd3-2800-7db3-aa91-68e188fa8864'::uuid,
    id,
    0.90,
    true,
    NOW()
FROM articles
WHERE doi = '10.2147/IJGM.S314874'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- =====================================================
-- VERIFICAÇÃO FINAL
-- =====================================================

SELECT
    '=== ARTIGOS LINKADOS (ATUALIZADO) ===' as status,
    COUNT(*) as total_articles,
    COUNT(CASE WHEN a.doi IS NOT NULL OR a.pm_id IS NOT NULL THEN 1 END) as peer_reviewed
FROM article_score_items asi
INNER JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = 'c77cedd3-2800-7db3-aa91-68e188fa8864';

SELECT
    '=== TOP 7 ARTIGOS POR CONFIDENCE ===' as info,
    a.title,
    SUBSTRING(a.authors, 1, 50) || '...' as first_author,
    EXTRACT(YEAR FROM a.publish_date) as year,
    a.journal,
    asi.confidence_score
FROM article_score_items asi
INNER JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = 'c77cedd3-2800-7db3-aa91-68e188fa8864'
ORDER BY asi.confidence_score DESC
LIMIT 7;

COMMIT;
