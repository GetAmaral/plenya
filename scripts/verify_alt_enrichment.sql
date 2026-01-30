-- Final verification for ALT enrichment
SELECT
    'Transaminase pirúvica (ALT)' as item_name,
    LENGTH(clinical_relevance) as clinical_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
        ELSE '✗ FAILED'
    END as clinical_status,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓ OK'
        ELSE '✗ FAILED'
    END as patient_status,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
        ELSE '✗ FAILED'
    END as conduct_status,
    last_review
FROM score_items
WHERE id = '06241683-bc19-4d56-b9a8-dade736e674f';

-- Verify linked articles
SELECT
    COUNT(*) as total_articles,
    STRING_AGG(a.pm_id, ', ' ORDER BY a.publish_date DESC) as pmids
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '06241683-bc19-4d56-b9a8-dade736e674f'
    AND a.pm_id IS NOT NULL;

-- Display article details
SELECT
    a.pm_id,
    LEFT(a.title, 80) as title,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '06241683-bc19-4d56-b9a8-dade736e674f'
    AND a.pm_id IS NOT NULL
ORDER BY a.publish_date DESC;
