-- ============================================================================
-- ENRIQUECIMENTO: Ferritina - Mulheres Pós-Menopausa
-- ID: 9d9f1270-d24d-4236-8694-5e28d8748475
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. CRIAR ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: Ferritina na transição menopausal (2025)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    doi,
    pm_id,
    original_link,
    article_type,
    language,
    abstract,
    specialty,
    created_at,
    updated_at
) VALUES (
    'Accelerated increase in ferritin levels during menopausal transition as a marker of metabolic health',
    'Kim Y, et al.',
    'Scientific Reports',
    '2025-01-15',
    '10.1038/s41598-025-14295-3',
    '40790323',
    'https://www.nature.com/articles/s41598-025-14295-3',
    'research_article',
    'en',
    'Ferritina aumenta rapidamente na transição menopausal (1 ano pós-última menstruação) e continua elevada, associada à saúde metabólica e risco de DM2.',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- Artigo 2: Ferro e menopausa (2009 - clássico)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    doi,
    pm_id,
    original_link,
    article_type,
    language,
    abstract,
    specialty,
    created_at,
    updated_at
) VALUES (
    'Iron and Menopause: Does Increased Iron Affect the Health of Postmenopausal Women?',
    'Milman N, Kirchhoff M, Jørgensen T',
    'Experimental Biology and Medicine',
    '2009-11-01',
    '10.3181/0905-MR-143',
    '19820134',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC2821138/',
    'research_article',
    'en',
    'Ferritina duplica 2-3x após menopausa (mediana 71-84 μg/L vs 37-42 μg/L pré-menopausa). Acúmulo de ferro pode aumentar risco cardiovascular e metabólico.',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- Artigo 3: Cutoffs de ferritina (2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    doi,
    pm_id,
    original_link,
    article_type,
    language,
    abstract,
    specialty,
    created_at,
    updated_at
) VALUES (
    'Ferritin Cutoffs and Diagnosis of Iron Deficiency in Primary Care',
    'Mast AE, et al.',
    'Annals of Internal Medicine',
    '2024-08-06',
    '10.7326/M23-2881',
    '39074341',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11301556/',
    'research_article',
    'en',
    'Cutoffs de 30-45 ng/mL para deficiência de ferro. Pós-menopausa requer limites superiores ajustados (>200 ng/mL sugere sobrecarga).',
    'Hematology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- ============================================================================
-- 2. VINCULAR ARTIGOS AO SCORE_ITEM
-- ============================================================================

-- Criar tabela temporária para armazenar IDs dos artigos
CREATE TEMP TABLE temp_article_ids AS
SELECT id, doi FROM articles WHERE doi IN (
    '10.1038/s41598-025-14295-3',
    '10.3181/0905-MR-143',
    '10.7326/M23-2881'
);

-- Vincular artigos (sem relevance_score, apenas relação many-to-many)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '9d9f1270-d24d-4236-8694-5e28d8748475'::uuid,
    t.id
FROM temp_article_ids t
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ============================================================================
-- 3. ATUALIZAR CONTEÚDO CLÍNICO DO SCORE_ITEM
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A ferritina reflete os estoques corporais de ferro e sofre elevação significativa após a menopausa devido à cessação da perda menstrual. Valores de referência específicos para mulheres pós-menopáusicas (mediana 71-84 μg/L) são 2-3x superiores aos de pré-menopáusicas (37-42 μg/L).

**Importância Clínica:**
- Deficiência de ferro: ferritina <30 ng/mL (requer investigação/reposição)
- Faixa normal-alta: 30-200 ng/mL (normal para pós-menopausa)
- Sobrecarga de ferro: >200 ng/mL (investigar hemocromatose, síndrome metabólica)
- Elevação rápida na transição menopausal (1 ano pós-FMP) é marcador de saúde metabólica

Ferritina elevada (>300 ng/mL) associa-se a risco aumentado de diabetes tipo 2, resistência insulínica, adiposidade central e síndrome metabólica. A ausência de menstruação permite acúmulo progressivo de ferro, tornando mulheres pós-menopáusicas suscetíveis a sobrecarga férrica, especialmente em portadoras de hemocromatose hereditária (manifestação tardia vs homens).',

    patient_explanation = 'A ferritina é uma proteína que armazena ferro no seu corpo. Após a menopausa, seus níveis de ferritina naturalmente aumentam porque você não perde mais ferro através da menstruação.

**Por que muda depois da menopausa?**
Durante a vida reprodutiva, a menstruação elimina ferro mensalmente, mantendo a ferritina mais baixa (média 37-42). Após a menopausa, esse ferro passa a se acumular, elevando a ferritina para 71-84 ou mais. Isso é normal e esperado!

**O que significam os resultados:**
- Ferritina baixa (<30): você pode estar com deficiência de ferro, mesmo sem menstruar. Causas: sangramento digestivo oculto, dieta pobre em ferro, má absorção intestinal
- Ferritina normal (30-200): seus estoques de ferro estão adequados para uma mulher pós-menopáusica
- Ferritina muito alta (>200): pode indicar excesso de ferro (hemocromatose), inflamação crônica ou problemas metabólicos

**Quando se preocupar:**
Ferritina acima de 200-300 merece investigação adicional, pois o excesso de ferro pode se depositar no fígado, coração e pâncreas, aumentando risco de diabetes, doenças cardíacas e problemas hepáticos. Sintomas de sobrecarga: cansaço extremo, dor nas articulações (especialmente mãos), escurecimento da pele, diabetes nova.',

    conduct = '**Interpretação de Resultados:**

1. **Ferritina <30 ng/mL (Deficiência de Ferro):**
   - Investigar causas: sangramento GI oculto, má absorção (doença celíaca, gastrite atrófica), dieta inadequada
   - Solicitar: hemograma completo, saturação de transferrina, ferro sérico
   - Endoscopia digestiva alta/colonoscopia se indicado (>50 anos)
   - Reposição: sulfato ferroso 300mg VO 1-2x/dia ou ferro EV se intolerância/má absorção
   - Reavaliação: ferritina após 3 meses de tratamento

2. **Ferritina 30-200 ng/mL (Normal para Pós-Menopausa):**
   - Conduta expectante
   - Manter dieta equilibrada (carnes magras, leguminosas, vegetais verde-escuros)
   - Monitoramento anual de ferritina/hemograma em check-up de rotina

3. **Ferritina 200-300 ng/mL (Limítrofe-Alta):**
   - Avaliar marcadores inflamatórios: PCR, VHS
   - Rastrear síndrome metabólica: glicemia, HbA1c, perfil lipídico, circunferência abdominal
   - Considerar teste genético se história familiar de hemocromatose
   - Repetir ferritina + saturação de transferrina em 3-6 meses

4. **Ferritina >300 ng/mL (Sobrecarga Suspeita):**
   - URGENTE: saturação de transferrina (>45% sugere hemocromatose)
   - Exames complementares: AST, ALT, ferritina sérica, ferro sérico
   - Teste genético HFE (C282Y, H63D) para hemocromatose hereditária
   - Encaminhar hematologia se saturação >45% ou ferritina >500 ng/mL
   - Tratamento: flebotomia terapêutica (retirada sangue periódica) até ferritina 50-100 ng/mL
   - Evitar: suplementos de ferro, excesso de vitamina C, álcool

**Monitoramento:**
- Pós-menopausa recente (primeiros 5 anos): ferritina anual
- Síndrome metabólica/diabetes: ferritina semestral
- Hemocromatose confirmada: ferritina a cada flebotomia (inicialmente mensal)

**Modificações de Estilo de Vida (Ferritina Alta):**
- Reduzir carnes vermelhas (preferir frango, peixe)
- Evitar alimentos fortificados com ferro
- Consumir chá/café com refeições (inibem absorção de ferro)
- Manter peso saudável e atividade física regular
- Doação de sangue regular (se elegível) ajuda a reduzir estoques férricos',

    last_review = CURRENT_DATE,
    updated_at = NOW()
WHERE id = '9d9f1270-d24d-4236-8694-5e28d8748475';

-- ============================================================================
-- 4. VERIFICAÇÃO FINAL
-- ============================================================================

\echo '============================================================================'
\echo 'RESULTADO FINAL - Ferritina - Mulheres Pós-Menopausa'
\echo '============================================================================'

SELECT
    si.name AS "Item",
    CASE
        WHEN si.clinical_relevance IS NOT NULL AND LENGTH(si.clinical_relevance) > 100 THEN '✓'
        ELSE '✗'
    END AS "Relevância",
    CASE
        WHEN si.patient_explanation IS NOT NULL AND LENGTH(si.patient_explanation) > 100 THEN '✓'
        ELSE '✗'
    END AS "Explicação",
    CASE
        WHEN si.conduct IS NOT NULL AND LENGTH(si.conduct) > 100 THEN '✓'
        ELSE '✗'
    END AS "Conduta",
    si.last_review AS "Última Revisão",
    COUNT(asi.article_id) AS "Artigos"
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '9d9f1270-d24d-4236-8694-5e28d8748475'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

\echo ''
\echo 'Artigos vinculados:'
SELECT
    a.title,
    a.journal,
    a.publish_date,
    a.doi
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '9d9f1270-d24d-4236-8694-5e28d8748475'
ORDER BY a.publish_date DESC;

COMMIT;

\echo ''
\echo '✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!'
\echo '============================================================================'
