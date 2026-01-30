# PCR Ultrassensível - Comandos de Verificação

Execute estes comandos para verificar o enriquecimento do item PCR ultrassensível.

---

## 1. Verificar Campos Enriquecidos

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    name,
    LENGTH(clinical_relevance) as len_clinical,
    LENGTH(patient_explanation) as len_patient,
    LENGTH(conduct) as len_conduct,
    last_review
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

**Resultado esperado:**
- Clinical relevance: ~1.668 caracteres
- Patient explanation: ~1.683 caracteres
- Conduct: ~2.472 caracteres
- Last review: 2026-01-29

---

## 2. Visualizar Clinical Relevance

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT clinical_relevance
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

## 3. Visualizar Patient Explanation

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT patient_explanation
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

## 4. Visualizar Conduct

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT conduct
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

## 5. Listar Artigos Vinculados

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    a.title,
    a.journal,
    EXTRACT(YEAR FROM a.publish_date) as year,
    a.doi,
    a.pm_id
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
ORDER BY a.publish_date DESC;
"
```

**Resultado esperado:** 12 artigos, incluindo:
- ACC Scientific Statement 2025
- UK Biobank Study 2024
- StatPearls 2024

---

## 6. Contar Artigos por Ano

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    EXTRACT(YEAR FROM a.publish_date) as year,
    COUNT(*) as num_articles
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
GROUP BY EXTRACT(YEAR FROM a.publish_date)
ORDER BY year DESC;
"
```

---

## 7. Verificar Hierarquia do Item

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    si.name as item_name,
    sg.name as subgroup_name,
    sgrp.name as group_name
FROM score_items si
JOIN score_subgroups sg ON si.subgroup_id = sg.id
JOIN score_groups sgrp ON sg.group_id = sgrp.id
WHERE si.id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

**Resultado esperado:**
- Item: PCR ultrassensível
- Subgroup: Laboratoriais
- Group: Exames

---

## 8. Verificar Artigos Novos Inseridos

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    id,
    title,
    doi,
    pm_id,
    created_at
FROM articles
WHERE
    (doi IN ('10.1016/j.jacc.2025.08.047', '10.1093/eurheartj/ehaf937')
    OR pm_id IN ('41020749', '30020613'))
    AND created_at::date = CURRENT_DATE
ORDER BY created_at DESC;
"
```

---

## 9. Preview de Todos os Campos

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    'Clinical Relevance' as campo,
    SUBSTRING(clinical_relevance, 1, 200) || '...' as preview
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Patient Explanation',
    SUBSTRING(patient_explanation, 1, 200) || '...'
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Conduct',
    SUBSTRING(conduct, 1, 200) || '...'
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

## 10. Estatísticas Completas

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
WITH item_stats AS (
    SELECT
        si.name,
        LENGTH(si.clinical_relevance) as len_clinical,
        LENGTH(si.patient_explanation) as len_patient,
        LENGTH(si.conduct) as len_conduct,
        si.last_review,
        COUNT(DISTINCT asi.article_id) as num_articles
    FROM score_items si
    LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
    WHERE si.id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
    GROUP BY si.id
),
article_years AS (
    SELECT
        EXTRACT(YEAR FROM a.publish_date) as year,
        COUNT(*) as articles_per_year
    FROM articles a
    JOIN article_score_items asi ON a.id = asi.article_id
    WHERE asi.score_item_id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
    GROUP BY EXTRACT(YEAR FROM a.publish_date)
    ORDER BY year DESC
    LIMIT 3
)
SELECT
    '📊 ESTATÍSTICAS DO ITEM PCR ULTRASSENSÍVEL' as info,
    '' as value
UNION ALL
SELECT '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━', ''
UNION ALL
SELECT 'Nome do Item', name FROM item_stats
UNION ALL
SELECT '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━', ''
UNION ALL
SELECT 'Clinical Relevance (caracteres)', len_clinical::text FROM item_stats
UNION ALL
SELECT 'Patient Explanation (caracteres)', len_patient::text FROM item_stats
UNION ALL
SELECT 'Conduct (caracteres)', len_conduct::text FROM item_stats
UNION ALL
SELECT 'Total de conteúdo (caracteres)', (len_clinical + len_patient + len_conduct)::text FROM item_stats
UNION ALL
SELECT '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━', ''
UNION ALL
SELECT 'Total de artigos vinculados', num_articles::text FROM item_stats
UNION ALL
SELECT 'Última revisão', to_char(last_review, 'YYYY-MM-DD HH24:MI:SS') FROM item_stats
UNION ALL
SELECT '━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━', ''
UNION ALL
SELECT 'Artigos mais recentes:', ''
UNION ALL
SELECT '  ' || year::text, articles_per_year::text || ' artigos' FROM article_years;
"
```

---

## 11. Buscar por Palavras-Chave

### Buscar "inflamação" no conteúdo

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    'Clinical Relevance' as campo,
    (LENGTH(clinical_relevance) - LENGTH(REPLACE(LOWER(clinical_relevance), 'inflamação', ''))) / LENGTH('inflamação') as ocorrencias
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Patient Explanation',
    (LENGTH(patient_explanation) - LENGTH(REPLACE(LOWER(patient_explanation), 'inflamação', ''))) / LENGTH('inflamação')
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Conduct',
    (LENGTH(conduct) - LENGTH(REPLACE(LOWER(conduct), 'inflamação', ''))) / LENGTH('inflamação')
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

### Buscar "cardiovascular" no conteúdo

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    'Clinical Relevance' as campo,
    (LENGTH(clinical_relevance) - LENGTH(REPLACE(LOWER(clinical_relevance), 'cardiovascular', ''))) / LENGTH('cardiovascular') as ocorrencias
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Patient Explanation',
    (LENGTH(patient_explanation) - LENGTH(REPLACE(LOWER(patient_explanation), 'cardiovascular', ''))) / LENGTH('cardiovascular')
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL
SELECT
    'Conduct',
    (LENGTH(conduct) - LENGTH(REPLACE(LOWER(conduct), 'cardiovascular', ''))) / LENGTH('cardiovascular')
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

## 12. Exportar Conteúdo Completo para Arquivo

```bash
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    '=== PCR ULTRASSENSÍVEL - CONTEÚDO COMPLETO ===' as content
UNION ALL SELECT ''
UNION ALL SELECT '## CLINICAL RELEVANCE'
UNION ALL SELECT ''
UNION ALL SELECT clinical_relevance FROM score_items WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL SELECT ''
UNION ALL SELECT '## PATIENT EXPLANATION'
UNION ALL SELECT ''
UNION ALL SELECT patient_explanation FROM score_items WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83'
UNION ALL SELECT ''
UNION ALL SELECT '## CONDUCT'
UNION ALL SELECT ''
UNION ALL SELECT conduct FROM score_items WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
" > /home/user/plenya/pcr-ultrassensivel-content-export.txt
```

---

## ✅ Checklist de Validação

Marque cada item após verificar:

- [ ] Campos clinical_relevance, patient_explanation e conduct preenchidos
- [ ] Total de conteúdo > 5.000 caracteres
- [ ] Mínimo de 10 artigos vinculados
- [ ] Artigos de 2024-2025 presentes
- [ ] ACC Scientific Statement 2025 vinculado
- [ ] UK Biobank Study 2024 vinculado
- [ ] Valores de referência (<1, 1-3, >3 mg/L) mencionados
- [ ] Last review atualizado (2026-01-29)
- [ ] Linguagem técnica em clinical_relevance
- [ ] Linguagem acessível em patient_explanation
- [ ] Conduta clínica completa com protocolo de solicitação
- [ ] Referências a diretrizes (ESC, ACC, AHA) em conduct

---

## 🔍 Troubleshooting

### Item não encontrado
```bash
# Verificar se o ID está correto
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT id, name FROM score_items WHERE name ILIKE '%PCR%' OR name ILIKE '%reativa%';
"
```

### Campos vazios
```bash
# Verificar NULL vs string vazia
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT
    clinical_relevance IS NULL as clinical_is_null,
    patient_explanation IS NULL as patient_is_null,
    conduct IS NULL as conduct_is_null
FROM score_items
WHERE id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

### Artigos não vinculados
```bash
# Verificar relações many-to-many
docker compose exec -T db psql -U plenya_user -d plenya_db -c "
SELECT COUNT(*) as total_relations
FROM article_score_items
WHERE score_item_id = '9b8eb642-4b9c-48e2-b15d-ee2aeba78c83';
"
```

---

**Última atualização:** 2026-01-29
