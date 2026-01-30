-- VALIDAÇÃO BATCH FINAL 1 - EXAMES LABORATORIAIS PARTE A
-- Execute após aplicar batch_final_1_exames_A.sql

\echo '==================================================================='
\echo 'VALIDAÇÃO BATCH FINAL 1 - EXAMES LABORATORIAIS PARTE A'
\echo '==================================================================='
\echo ''

\echo '1. Total de items atualizados hoje (esperado: 45)'
\echo '-------------------------------------------------------------------'
SELECT COUNT(*) as total_updated
FROM score_items
WHERE last_review >= CURRENT_DATE
  AND interpretation IS NOT NULL
  AND articles IS NOT NULL;

\echo ''
\echo '2. Items com enrichment específico detalhado (esperado: 3)'
\echo '-------------------------------------------------------------------'
SELECT
  name,
  LENGTH(interpretation) as interp_chars,
  LENGTH(low_level_description) + LENGTH(medium_level_description) + LENGTH(high_level_description) as desc_chars,
  jsonb_array_length(articles) as artigos,
  last_review::date
FROM score_items
WHERE id IN (
  '341946e7-5833-48bc-b316-71e29954eedd',
  '348fc460-9959-4648-9d0d-6acafd2f9700',
  '579a961c-e160-417f-9371-418284386f35'
)
ORDER BY name;

\echo ''
\echo '3. Distribuição por tipo de enrichment'
\echo '-------------------------------------------------------------------'
SELECT
  CASE
    WHEN LENGTH(interpretation) > 500 AND jsonb_array_length(articles) > 3 THEN 'ESPECÍFICO'
    ELSE 'PADRÃO'
  END as enrichment_type,
  COUNT(*) as total
FROM score_items
WHERE last_review >= CURRENT_DATE
GROUP BY enrichment_type
ORDER BY enrichment_type DESC;

\echo ''
\echo '4. Estatísticas de tamanho do conteúdo'
\echo '-------------------------------------------------------------------'
SELECT
  ROUND(AVG(LENGTH(interpretation))) as avg_interpretation_chars,
  ROUND(AVG(LENGTH(low_level_description))) as avg_low_desc_chars,
  ROUND(AVG(LENGTH(medium_level_description))) as avg_medium_desc_chars,
  ROUND(AVG(LENGTH(high_level_description))) as avg_high_desc_chars,
  ROUND(AVG(jsonb_array_length(articles))) as avg_articles_count
FROM score_items
WHERE last_review >= CURRENT_DATE;

\echo ''
\echo '5. Verificar integridade dos artigos (JSON válido)'
\echo '-------------------------------------------------------------------'
SELECT
  id,
  name,
  jsonb_array_length(articles) as total_artigos,
  articles->0->>'title' as primeiro_artigo_titulo
FROM score_items
WHERE id = '341946e7-5833-48bc-b316-71e29954eedd';

\echo ''
\echo '6. Items por subgrupo'
\echo '-------------------------------------------------------------------'
SELECT
  sg.name as subgrupo,
  COUNT(*) as total_items
FROM score_items si
JOIN score_groups sg ON si.group_id = sg.id
WHERE si.last_review >= CURRENT_DATE
GROUP BY sg.name
ORDER BY total_items DESC;

\echo ''
\echo '7. Lista completa dos 45 items processados'
\echo '-------------------------------------------------------------------'
SELECT
  ROW_NUMBER() OVER (ORDER BY name) as num,
  name,
  CASE
    WHEN LENGTH(interpretation) > 500 THEN '✓ ESPECÍFICO'
    ELSE '◦ PADRÃO'
  END as tipo,
  LENGTH(interpretation) as chars,
  last_review::date as atualizado
FROM score_items
WHERE last_review >= CURRENT_DATE
ORDER BY name;

\echo ''
\echo '8. Verificar campos obrigatórios preenchidos (esperado: 0 nulls)'
\echo '-------------------------------------------------------------------'
SELECT
  COUNT(*) FILTER (WHERE interpretation IS NULL) as interpretation_null,
  COUNT(*) FILTER (WHERE low_level_description IS NULL) as low_desc_null,
  COUNT(*) FILTER (WHERE medium_level_description IS NULL) as medium_desc_null,
  COUNT(*) FILTER (WHERE high_level_description IS NULL) as high_desc_null,
  COUNT(*) FILTER (WHERE low_level_recommendation IS NULL) as low_rec_null,
  COUNT(*) FILTER (WHERE medium_level_recommendation IS NULL) as medium_rec_null,
  COUNT(*) FILTER (WHERE high_level_recommendation IS NULL) as high_rec_null,
  COUNT(*) FILTER (WHERE articles IS NULL) as articles_null
FROM score_items
WHERE last_review >= CURRENT_DATE;

\echo ''
\echo '9. Sample de conteúdo específico (Mamografia - Densidade Mamária)'
\echo '-------------------------------------------------------------------'
SELECT
  name,
  LEFT(interpretation, 150) || '...' as interpretation_sample,
  LEFT(high_level_recommendation, 150) || '...' as high_rec_sample
FROM score_items
WHERE id = '341946e7-5833-48bc-b316-71e29954eedd';

\echo ''
\echo '==================================================================='
\echo 'FIM DA VALIDAÇÃO'
\echo '==================================================================='
