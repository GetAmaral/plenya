-- VALIDAÇÃO DO BATCH FINAL 3 - EXAMES C
-- Executar APÓS aplicar batch_final_3_ALL_35_items.sql

-- 1. Verificar total de items atualizados
SELECT
  'Total de items atualizados nas últimas 2 horas' as check_name,
  COUNT(*) as result,
  '35' as expected
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', '30af9809-7316-47e6-b363-8279c7bd3a69',
  'afdd9d67-3f42-4e6e-a77d-0e75475ca72d', '3a0d7e6b-c53d-47de-a50c-d7774e542835',
  '2e3c06ce-bcb6-4649-984e-8c30a92e26f4', 'b2dd0c76-7bce-4beb-a8e2-52d70d467241',
  '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb', '6e542cb0-6982-42e8-93dc-40f139652223',
  'c8795e89-b10a-4d51-b940-463ab5e89c3e', 'fb5c484c-bb3c-4017-b4da-866d96d9cb20',
  'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d', '3e67c010-1164-4e0f-b23f-109557d6d51d',
  'b99d1e15-baa3-4bcb-956e-7314dbccfa82', '210eedab-08e7-4f47-ae6a-37aecea9bc16',
  'b87387b4-d024-4dbb-be70-84778ca2dce0', 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb',
  'a92d20ce-702f-4b15-8817-098b9539c0f0', 'c04e1997-0846-4557-8990-baffb1f7542a',
  '501fd84a-a440-4c13-9b11-35e2f69017d1', '81050f48-2b89-4da9-a480-f55af29bdb8b',
  'f1a1d106-b5ac-4483-9de9-ee90ae103d26', 'bf18b584-1f48-46aa-b345-c98a82fb6709',
  '56350a39-c589-4ed5-b757-0424399b7a61', '3faeb6db-b8d6-4fb1-8740-07bfd91002c7',
  'c8ef5fcc-01ec-4eee-b4f5-11c358392c0b', '91a9cf25-2622-4296-bf16-64b9f95b1e8d',
  'b3cb69e0-98fd-48ca-82ae-9eee62a8218e', '73411e66-c180-46ad-8f5b-7d67b26ef877',
  '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d', '039f1542-7596-4671-8ed0-049b3b41cfc4',
  'cbd75469-ca59-4b12-ab18-9b6b202f54fc', 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd',
  'd1fd128f-7f92-4498-be3f-3bfe9564ce70', '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23',
  'eba25efc-0b97-4b95-8012-0e027c6ec311'
)
AND updated_at > NOW() - INTERVAL '2 hours';

-- 2. Verificar campos preenchidos
SELECT
  'Items com todos os campos obrigatórios preenchidos' as check_name,
  COUNT(*) as result,
  '35' as expected
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', '30af9809-7316-47e6-b363-8279c7bd3a69',
  'afdd9d67-3f42-4e6e-a77d-0e75475ca72d', '3a0d7e6b-c53d-47de-a50c-d7774e542835',
  '2e3c06ce-bcb6-4649-984e-8c30a92e26f4', 'b2dd0c76-7bce-4beb-a8e2-52d70d467241',
  '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb', '6e542cb0-6982-42e8-93dc-40f139652223',
  'c8795e89-b10a-4d51-b940-463ab5e89c3e', 'fb5c484c-bb3c-4017-b4da-866d96d9cb20',
  'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d', '3e67c010-1164-4e0f-b23f-109557d6d51d',
  'b99d1e15-baa3-4bcb-956e-7314dbccfa82', '210eedab-08e7-4f47-ae6a-37aecea9bc16',
  'b87387b4-d024-4dbb-be70-84778ca2dce0', 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb',
  'a92d20ce-702f-4b15-8817-098b9539c0f0', 'c04e1997-0846-4557-8990-baffb1f7542a',
  '501fd84a-a440-4c13-9b11-35e2f69017d1', '81050f48-2b89-4da9-a480-f55af29bdb8b',
  'f1a1d106-b5ac-4483-9de9-ee90ae103d26', 'bf18b584-1f48-46aa-b345-c98a82fb6709',
  '56350a39-c589-4ed5-b757-0424399b7a61', '3faeb6db-b8d6-4fb1-8740-07bfd91002c7',
  'c8ef5fcc-01ec-4eee-b4f5-11c358392c0b', '91a9cf25-2622-4296-bf16-64b9f95b1e8d',
  'b3cb69e0-98fd-48ca-82ae-9eee62a8218e', '73411e66-c180-46ad-8f5b-7d67b26ef877',
  '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d', '039f1542-7596-4671-8ed0-049b3b41cfc4',
  'cbd75469-ca59-4b12-ab18-9b6b202f54fc', 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd',
  'd1fd128f-7f92-4498-be3f-3bfe9564ce70', '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23',
  'eba25efc-0b97-4b95-8012-0e027c6ec311'
)
AND unit IS NOT NULL
AND optimal_explanation IS NOT NULL
AND clinical_interpretation IS NOT NULL
AND medical_conduct IS NOT NULL
AND related_articles IS NOT NULL;

-- 3. Verificar tamanho médio dos campos
SELECT
  'Tamanho médio de clinical_interpretation' as metric,
  ROUND(AVG(LENGTH(clinical_interpretation))) as avg_chars,
  ROUND(AVG(LENGTH(clinical_interpretation) / 6.0)) as approx_words,
  '400-600 palavras esperado' as target
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', 'afdd9d67-3f42-4e6e-a77d-0e75475ca72d',
  'fb5c484c-bb3c-4017-b4da-866d96d9cb20', '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d'
);

-- 4. Amostras de items para validação manual
SELECT
  name,
  optimal_range_min,
  optimal_range_max,
  unit,
  SUBSTRING(optimal_explanation, 1, 80) || '...' as explanation_preview,
  LENGTH(clinical_interpretation) as interp_chars,
  LENGTH(medical_conduct) as conduct_chars,
  array_length(related_articles, 1) as num_articles
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', -- Fundoscopia RD
  'afdd9d67-3f42-4e6e-a77d-0e75475ca72d', -- Volume Ovariano
  'fb5c484c-bb3c-4017-b4da-866d96d9cb20', -- GLS
  '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d', -- Homocisteína
  'bf18b584-1f48-46aa-b345-c98a82fb6709', -- RDW
  'eba25efc-0b97-4b95-8012-0e027c6ec311'  -- IgG
)
ORDER BY name;

-- 5. Verificar artigos científicos
SELECT
  'Total de artigos científicos adicionados' as metric,
  SUM(array_length(related_articles, 1)) as total_articles,
  '105 esperado (3/item x 35)' as expected
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', '30af9809-7316-47e6-b363-8279c7bd3a69',
  'afdd9d67-3f42-4e6e-a77d-0e75475ca72d', '3a0d7e6b-c53d-47de-a50c-d7774e542835',
  '2e3c06ce-bcb6-4649-984e-8c30a92e26f4', 'b2dd0c76-7bce-4beb-a8e2-52d70d467241',
  '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb', '6e542cb0-6982-42e8-93dc-40f139652223',
  'c8795e89-b10a-4d51-b940-463ab5e89c3e', 'fb5c484c-bb3c-4017-b4da-866d96d9cb20',
  'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d', '3e67c010-1164-4e0f-b23f-109557d6d51d',
  'b99d1e15-baa3-4bcb-956e-7314dbccfa82', '210eedab-08e7-4f47-ae6a-37aecea9bc16',
  'b87387b4-d024-4dbb-be70-84778ca2dce0', 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb',
  'a92d20ce-702f-4b15-8817-098b9539c0f0', 'c04e1997-0846-4557-8990-baffb1f7542a',
  '501fd84a-a440-4c13-9b11-35e2f69017d1', '81050f48-2b89-4da9-a480-f55af29bdb8b',
  'f1a1d106-b5ac-4483-9de9-ee90ae103d26', 'bf18b584-1f48-46aa-b345-c98a82fb6709',
  '56350a39-c589-4ed5-b757-0424399b7a61', '3faeb6db-b8d6-4fb1-8740-07bfd91002c7',
  'c8ef5fcc-01ec-4eee-b4f5-11c358392c0b', '91a9cf25-2622-4296-bf16-64b9f95b1e8d',
  'b3cb69e0-98fd-48ca-82ae-9eee62a8218e', '73411e66-c180-46ad-8f5b-7d67b26ef877',
  '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d', '039f1542-7596-4671-8ed0-049b3b41cfc4',
  'cbd75469-ca59-4b12-ab18-9b6b202f54fc', 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd',
  'd1fd128f-7f92-4498-be3f-3bfe9564ce70', '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23',
  'eba25efc-0b97-4b95-8012-0e027c6ec311'
);

-- 6. Listar todos os 35 items para conferência
SELECT
  ROW_NUMBER() OVER (ORDER BY name) as num,
  name,
  CASE
    WHEN optimal_range_min IS NOT NULL AND optimal_range_max IS NOT NULL
      THEN optimal_range_min::TEXT || '-' || optimal_range_max::TEXT || ' ' || unit
    WHEN optimal_range_max IS NOT NULL
      THEN '<=' || optimal_range_max::TEXT || ' ' || unit
    ELSE unit
  END as range,
  array_length(related_articles, 1) as articles,
  DATE_TRUNC('minute', updated_at) as updated
FROM score_items
WHERE id IN (
  'dbbb0258-ebcb-4bb2-ace3-73d0d7b93a52', '30af9809-7316-47e6-b363-8279c7bd3a69',
  'afdd9d67-3f42-4e6e-a77d-0e75475ca72d', '3a0d7e6b-c53d-47de-a50c-d7774e542835',
  '2e3c06ce-bcb6-4649-984e-8c30a92e26f4', 'b2dd0c76-7bce-4beb-a8e2-52d70d467241',
  '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb', '6e542cb0-6982-42e8-93dc-40f139652223',
  'c8795e89-b10a-4d51-b940-463ab5e89c3e', 'fb5c484c-bb3c-4017-b4da-866d96d9cb20',
  'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d', '3e67c010-1164-4e0f-b23f-109557d6d51d',
  'b99d1e15-baa3-4bcb-956e-7314dbccfa82', '210eedab-08e7-4f47-ae6a-37aecea9bc16',
  'b87387b4-d024-4dbb-be70-84778ca2dce0', 'dd3bb9b1-26c1-4fc3-a85d-c83bf6d0fdfb',
  'a92d20ce-702f-4b15-8817-098b9539c0f0', 'c04e1997-0846-4557-8990-baffb1f7542a',
  '501fd84a-a440-4c13-9b11-35e2f69017d1', '81050f48-2b89-4da9-a480-f55af29bdb8b',
  'f1a1d106-b5ac-4483-9de9-ee90ae103d26', 'bf18b584-1f48-46aa-b345-c98a82fb6709',
  '56350a39-c589-4ed5-b757-0424399b7a61', '3faeb6db-b8d6-4fb1-8740-07bfd91002c7',
  'c8ef5fcc-01ec-4eee-b4f5-11c358392c0b', '91a9cf25-2622-4296-bf16-64b9f95b1e8d',
  'b3cb69e0-98fd-48ca-82ae-9eee62a8218e', '73411e66-c180-46ad-8f5b-7d67b26ef877',
  '6eb3a909-6c91-4b72-a5d6-bfa546d8de8d', '039f1542-7596-4671-8ed0-049b3b41cfc4',
  'cbd75469-ca59-4b12-ab18-9b6b202f54fc', 'bfadf0e5-4df7-4e70-9ed7-772a237a03fd',
  'd1fd128f-7f92-4498-be3f-3bfe9564ce70', '5230ca12-d4dc-4d1c-8ee8-e8b8b3b23a23',
  'eba25efc-0b97-4b95-8012-0e027c6ec311'
)
ORDER BY num;
