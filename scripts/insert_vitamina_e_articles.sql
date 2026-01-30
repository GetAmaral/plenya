-- Insert 4 scientific articles for Vitamina E (Alfa-Tocoferol)

BEGIN;

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

-- Link all articles to the Vitamina E score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '80ffc2b2-545b-4389-891f-b6aba1f7865c'::uuid
FROM articles a
WHERE a.pm_id IN ('40065887', '32810309', '39978678', '33503988')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

COMMIT;

-- Verify insertion
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
