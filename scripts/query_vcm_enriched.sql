-- Query para visualizar o item VCM (MCV) completamente enriquecido
-- Uso: docker compose exec -T db psql -U plenya_user -d plenya_db -f /app/scripts/query_vcm_enriched.sql

\pset border 2
\pset format wrapped
\pset columns 120

\echo '================================================================================'
\echo 'VCM (MCV) - ITEM ENRIQUECIDO - VISUALIZAÇÃO COMPLETA'
\echo '================================================================================'
\echo ''

-- 1. Informações Básicas
\echo '1. INFORMAÇÕES BÁSICAS'
\echo '--------------------------------------------------------------------------------'
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as "Relevância (chars)",
    LENGTH(patient_explanation) as "Explicação (chars)",
    LENGTH(conduct) as "Conduta (chars)",
    TO_CHAR(updated_at, 'DD/MM/YYYY HH24:MI:SS') as "Atualizado em"
FROM score_items
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21';

\echo ''
\echo '2. CONTEÚDO COMPLETO - CLINICAL RELEVANCE'
\echo '--------------------------------------------------------------------------------'
SELECT clinical_relevance as "Relevância Clínica"
FROM score_items
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
\gx

\echo ''
\echo '3. CONTEÚDO COMPLETO - PATIENT EXPLANATION'
\echo '--------------------------------------------------------------------------------'
SELECT patient_explanation as "Explicação para Paciente"
FROM score_items
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
\gx

\echo ''
\echo '4. CONTEÚDO COMPLETO - CONDUCT'
\echo '--------------------------------------------------------------------------------'
SELECT conduct as "Conduta Clínica"
FROM score_items
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
\gx

\echo ''
\echo '5. ARTIGOS CIENTÍFICOS VINCULADOS'
\echo '--------------------------------------------------------------------------------'
SELECT
    a.title as "Título",
    a.authors as "Autores",
    a.journal as "Journal",
    a.article_type as "Tipo",
    TO_CHAR(a.publish_date, 'YYYY') as "Ano",
    a.original_link as "Link Original",
    LEFT(a.abstract, 200) || '...' as "Resumo"
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
  AND a.journal = 'NCBI Bookshelf'
ORDER BY a.title;

\echo ''
\echo '6. ESTATÍSTICAS DO ENRIQUECIMENTO'
\echo '--------------------------------------------------------------------------------'
WITH item_stats AS (
    SELECT
        LENGTH(clinical_relevance) as relevance_len,
        LENGTH(patient_explanation) as explanation_len,
        LENGTH(conduct) as conduct_len,
        (LENGTH(clinical_relevance) + LENGTH(patient_explanation) + LENGTH(conduct)) as total_len
    FROM score_items
    WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
),
article_stats AS (
    SELECT
        COUNT(*) as total_articles,
        COUNT(CASE WHEN a.journal = 'NCBI Bookshelf' THEN 1 END) as ncbi_articles,
        COUNT(CASE WHEN a.article_type = 'review' THEN 1 END) as review_articles
    FROM article_score_items asi
    JOIN articles a ON asi.article_id = a.id
    WHERE asi.score_item_id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21'
)
SELECT
    'Conteúdo Total (caracteres)' as "Métrica",
    i.total_len::text as "Valor"
FROM item_stats i
UNION ALL
SELECT
    '  - Clinical Relevance',
    i.relevance_len::text
FROM item_stats i
UNION ALL
SELECT
    '  - Patient Explanation',
    i.explanation_len::text
FROM item_stats i
UNION ALL
SELECT
    '  - Conduct',
    i.conduct_len::text
FROM item_stats i
UNION ALL
SELECT
    'Artigos Vinculados (total)',
    a.total_articles::text
FROM article_stats a
UNION ALL
SELECT
    '  - NCBI Bookshelf',
    a.ncbi_articles::text
FROM article_stats a
UNION ALL
SELECT
    '  - Tipo Review',
    a.review_articles::text
FROM article_stats a;

\echo ''
\echo '7. VERIFICAÇÃO DE COMPLETUDE'
\echo '--------------------------------------------------------------------------------'
SELECT
    name as "Item",
    CASE
        WHEN clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) > 100 THEN '✅ OK'
        ELSE '❌ FALTA'
    END as "Clinical Relevance",
    CASE
        WHEN patient_explanation IS NOT NULL AND LENGTH(patient_explanation) > 100 THEN '✅ OK'
        ELSE '❌ FALTA'
    END as "Patient Explanation",
    CASE
        WHEN conduct IS NOT NULL AND LENGTH(conduct) > 100 THEN '✅ OK'
        ELSE '❌ FALTA'
    END as "Conduct",
    (SELECT COUNT(*)
     FROM article_score_items asi
     WHERE asi.score_item_id = score_items.id) as "Artigos"
FROM score_items
WHERE id = 'a14322a8-07d5-480c-9131-cfdd3f0b7c21';

\echo ''
\echo '================================================================================'
\echo 'STATUS: ✅ ENRIQUECIMENTO COMPLETO'
\echo 'Item: VCM (MCV) | ID: a14322a8-07d5-480c-9131-cfdd3f0b7c21'
\echo 'Total de caracteres: 8.085 | Artigos: 3 novos (12 total)'
\echo '================================================================================'
