-- Verification query for H2S Peak enrichment
-- Score Item: b87387b4-d024-4dbb-be70-84778ca2dce0

-- 1. Content character counts
SELECT
    'Content Character Counts' as verification_step,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars
FROM score_items
WHERE id = 'b87387b4-d024-4dbb-be70-84778ca2dce0';

-- 2. Article statistics
SELECT
    'Article Statistics' as verification_step,
    COUNT(*) as total_articles,
    COUNT(CASE WHEN pm_id IS NOT NULL AND pm_id != '' THEN 1 END) as with_pmid,
    COUNT(CASE WHEN doi IS NOT NULL AND doi != '' THEN 1 END) as with_doi,
    COUNT(CASE WHEN article_type = 'research_article' THEN 1 END) as research_articles,
    COUNT(CASE WHEN article_type = 'review' THEN 1 END) as reviews,
    COUNT(CASE WHEN article_type = 'lecture' THEN 1 END) as lectures
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'b87387b4-d024-4dbb-be70-84778ca2dce0';

-- 3. Scientific articles only (with PMID)
SELECT
    'Scientific Articles' as verification_step,
    title,
    pm_id,
    EXTRACT(YEAR FROM publish_date) as year,
    article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'b87387b4-d024-4dbb-be70-84778ca2dce0'
AND pm_id IS NOT NULL
AND pm_id != ''
ORDER BY publish_date DESC;
