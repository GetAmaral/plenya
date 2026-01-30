-- Verificação completa do enriquecimento SHBG - Mulheres

SELECT '=== VERIFICAÇÃO DE CONTEÚDO CLÍNICO ===' as section;

SELECT
  si.name as item_name,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review
FROM score_items si
WHERE si.id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';

SELECT '=== ARTIGOS CIENTÍFICOS VINCULADOS ===' as section;

SELECT
  a.title,
  a.authors,
  a.journal,
  a.publish_date,
  a.pm_id,
  a.doi,
  a.article_type
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'c21ccec2-66b2-49e3-911a-8d0944eda087'
  AND a.pm_id IS NOT NULL
ORDER BY a.publish_date DESC;

SELECT '=== CONTAGEM TOTAL DE ARTIGOS ===' as section;

SELECT COUNT(*) as total_articles
FROM article_score_items asi
WHERE asi.score_item_id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';
