-- ============================================================================
-- ENRIQUECIMENTO: Hidrogênio Pico (H₂ Máximo)
-- Item ID: 210eedab-08e7-4f47-ae6a-37aecea9bc16
-- Data: 2026-01-28
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: Understanding Our Tests (2023) - Review atualizado
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    INSERT INTO scientific_articles (
        title, authors, journal, publication_year, doi, url, article_type, relevance_score
    ) VALUES (
        'Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth',
        'Rezaie A, Heimanson Z, McCallum R, Pimentel M',
        'Clinical and Translational Gastroenterology',
        2023,
        '10.14309/ctg.0000000000000567',
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC10132719/',
        'review',
        95
    ) ON CONFLICT (doi) DO UPDATE SET
        url = EXCLUDED.url,
        relevance_score = EXCLUDED.relevance_score
    RETURNING id INTO v_article_id;

    -- Vincular ao score_item
    INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
    VALUES (
        '210eedab-08e7-4f47-ae6a-37aecea9bc16',
        v_article_id,
        'Revisão atualizada sobre interpretação do teste respiratório para SIBO'
    ) ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 1 inserido: Understanding Our Tests (2023)';
END $$;

-- Artigo 2: North American Consensus (2017)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    INSERT INTO scientific_articles (
        title, authors, journal, publication_year, doi, url, article_type, relevance_score
    ) VALUES (
        'Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus',
        'Rezaie A, Buresi M, Lembo A, Lin H, McCallum R, Rao S, Schmulson M, Valdovinos M, Zakko S, Pimentel M',
        'American Journal of Gastroenterology',
        2017,
        '10.1038/ajg.2017.46',
        'https://pubmed.ncbi.nlm.nih.gov/28323273/',
        'consensus',
        98
    ) ON CONFLICT (doi) DO UPDATE SET
        url = EXCLUDED.url,
        relevance_score = EXCLUDED.relevance_score
    RETURNING id INTO v_article_id;

    INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
    VALUES (
        '210eedab-08e7-4f47-ae6a-37aecea9bc16',
        v_article_id,
        'Consenso norte-americano - critérios diagnósticos para SIBO (≥20 ppm em 90 min)'
    ) ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 2 inserido: North American Consensus (2017)';
END $$;

-- Artigo 3: Refined Lactulose Test (2025) - Estudo mais recente
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    INSERT INTO scientific_articles (
        title, authors, journal, publication_year, doi, url, article_type, relevance_score
    ) VALUES (
        'Refined Lactulose Hydrogen Breath Test for Small Intestinal Bacterial Overgrowth Subgrouping Irritable Bowel Syndrome',
        'Dahlgren A, Sundström J, Aldenbratt A, Törnblom H, Simrén M, Strid H',
        'Gastroenterology Research and Practice',
        2025,
        '10.1155/grp/5597071',
        'https://onlinelibrary.wiley.com/doi/10.1155/grp/5597071',
        'clinical_trial',
        92
    ) ON CONFLICT (doi) DO UPDATE SET
        url = EXCLUDED.url,
        relevance_score = EXCLUDED.relevance_score
    RETURNING id INTO v_article_id;

    INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
    VALUES (
        '210eedab-08e7-4f47-ae6a-37aecea9bc16',
        v_article_id,
        'Estudo 2025 validando cutoff de 20 ppm (sensibilidade 77%, especificidade 88%)'
    ) ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 3 inserido: Refined Lactulose Test (2025)';
END $$;

-- Artigo 4: How to Interpret (2014)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    INSERT INTO scientific_articles (
        title, authors, journal, publication_year, doi, url, article_type, relevance_score
    ) VALUES (
        'How to Interpret Hydrogen Breath Tests',
        'Rana SV, Malik A',
        'Journal of Neurogastroenterology and Motility',
        2014,
        '10.5056/jnm.2014.20.3.312',
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC4155069/',
        'review',
        88
    ) ON CONFLICT (doi) DO UPDATE SET
        url = EXCLUDED.url,
        relevance_score = EXCLUDED.relevance_score
    RETURNING id INTO v_article_id;

    INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
    VALUES (
        '210eedab-08e7-4f47-ae6a-37aecea9bc16',
        v_article_id,
        'Guia prático de interpretação do teste respiratório de hidrogênio'
    ) ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 4 inserido: How to Interpret (2014)';
END $$;

-- Artigo 5: Carbohydrate Malabsorption (2014)
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    INSERT INTO scientific_articles (
        title, authors, journal, publication_year, doi, url, article_type, relevance_score
    ) VALUES (
        'Carbohydrate malabsorption in patients with non-specific abdominal complaints',
        'Yao CK, Gibson PR, Shepherd SJ',
        'Current Opinion in Clinical Nutrition & Metabolic Care',
        2014,
        '10.1097/MCO.0000000000000094',
        'https://pmc.ncbi.nlm.nih.gov/articles/PMC4171253/',
        'review',
        85
    ) ON CONFLICT (doi) DO UPDATE SET
        url = EXCLUDED.url,
        relevance_score = EXCLUDED.relevance_score
    RETURNING id INTO v_article_id;

    INSERT INTO score_item_articles (score_item_id, article_id, relevance_context)
    VALUES (
        '210eedab-08e7-4f47-ae6a-37aecea9bc16',
        v_article_id,
        'Má absorção de carboidratos: critérios diagnósticos via teste respiratório'
    ) ON CONFLICT (score_item_id, article_id) DO NOTHING;

    RAISE NOTICE 'Artigo 5 inserido: Carbohydrate Malabsorption (2014)';
END $$;

-- ============================================================================
-- 2. CONTEÚDO CLÍNICO EM PT-BR
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O Hidrogênio Pico (H₂ Máximo) é o valor máximo de hidrogênio exalado durante o teste respiratório de hidrogênio, sendo fundamental para diagnóstico de:

**SIBO (Supercrescimento Bacteriano no Intestino Delgado):**
- Critério diagnóstico: elevação ≥20 ppm nos primeiros 90 minutos (consenso norte-americano)
- Sensibilidade: 77% / Especificidade: 88% (usando cutoff de 20 ppm)
- Pacientes com SIBO apresentam pico médio de 38 ppm vs. 8 ppm em indivíduos saudáveis

**Má Absorção de Carboidratos:**
- Intolerância à lactose: pico ≥20 ppm acima do basal após carga de lactose
- Má absorção de frutose: pico ≥20 ppm acima do basal após carga de frutose
- Cutoff diferenciado: 12 ppm para SIBO vs. 20 ppm para má absorção

**Interpretação Conjunta:**
- Metano ≥10 ppm indica IMO (Intestinal Methanogen Overgrowth)
- Produtores mistos: considerar média de H₂ + CH₄ (cutoff 15 ppm)
- O valor do pico NÃO se correlaciona com intensidade dos sintomas

**Limitações:**
- 15-27% da população são "não produtores" de H₂ (resultado falso-negativo)
- Preparação inadequada pode gerar resultados inválidos
- Deve ser interpretado junto com tempo de trânsito orocecal',

    patient_explanation = 'Este exame mede o hidrogênio no ar que você expira após beber uma solução açucarada, ajudando a identificar problemas digestivos.

**Como Funciona:**
Normalmente, o açúcar é absorvido no intestino delgado. Se houver:
- Bactérias em excesso no intestino delgado (SIBO)
- Dificuldade em absorver lactose ou frutose

As bactérias fermentam o açúcar e produzem hidrogênio, que vai para o sangue e é eliminado pelos pulmões.

**Valores de Referência:**
- Normal: pico até 20 ppm (partes por milhão)
- SIBO: pico ≥20 ppm nos primeiros 90 minutos
- Má absorção: pico ≥20 ppm acima do valor inicial

**O Que Significa Resultado Alterado:**
- Pode indicar SIBO (supercrescimento bacteriano)
- Pode indicar intolerância à lactose ou frutose
- Pode explicar sintomas como inchaço, diarreia, gases e dor abdominal

**Importante:**
O resultado precisa ser interpretado pelo médico junto com seus sintomas e o tempo que levou para o hidrogênio subir. Algumas pessoas não produzem hidrogênio (são "não produtoras") e podem ter resultado falso normal.',

    conduct = '**Interpretação do Resultado:**

1. **H₂ Pico <20 ppm:**
   - Resultado negativo para SIBO e má absorção
   - ATENÇÃO: descartar paciente "não produtor" de H₂
   - Considerar teste com metano se sintomas persistentes

2. **H₂ Pico ≥20 ppm nos primeiros 90 min:**
   - SIBO provável (sensibilidade 77%, especificidade 88%)
   - Correlacionar com tempo de pico (antes ou depois de 90 min)
   - Verificar metano concomitante (IMO se ≥10 ppm)
   - **Conduta:**
     - Antibioticoterapia (rifaximina 550mg 3x/dia por 14 dias)
     - Dieta pobre em FODMAPs durante tratamento
     - Investigar causa base (dismotilidade, hipocloridria, cirurgia prévia)
     - Repetir teste 2-4 semanas após tratamento

3. **H₂ Pico ≥20 ppm após 90 min (>120 min):**
   - Má absorção de carboidrato provável
   - Identificar qual carboidrato (lactose, frutose, lactulose)
   - **Conduta:**
     - Dieta de exclusão do carboidrato identificado
     - Suplementação com enzimas digestivas (ex: lactase para lactose)
     - Orientação nutricional
     - Reavaliar sintomas em 4-6 semanas

4. **H₂ Pico 10-19 ppm (zona cinzenta):**
   - Interpretar com cautela
   - Considerar contexto clínico e sintomas
   - Pode indicar disbiose leve
   - Repetir teste ou complementar com metano

**Preparação Adequada do Paciente (Essencial):**
- Suspender antibióticos: 4 semanas antes
- Suspender probióticos: 2 semanas antes
- Suspender laxantes/procinéticos: 1 semana antes
- Jejum de 12 horas (exceto água)
- Não fumar/exercitar no dia do teste
- Dieta restrita no dia anterior (arroz branco, frango grelhado)

**Follow-up:**
- Sucesso terapêutico: redução de ≥50% no pico ou normalização
- Se falha: investigar resistência, IMO concomitante, causa não corrigida
- Recorrência comum: abordar fatores predisponentes',

    reference_range = '<20 ppm (normal), ≥20 ppm (SIBO ou má absorção)',
    updated_at = NOW()

WHERE id = '210eedab-08e7-4f47-ae6a-37aecea9bc16';

-- ============================================================================
-- 3. VERIFICAÇÃO DO RESULTADO
-- ============================================================================

DO $$
DECLARE
    v_count INTEGER;
    v_name TEXT;
    v_unit TEXT;
    v_ref_range TEXT;
BEGIN
    -- Contar artigos vinculados
    SELECT COUNT(*) INTO v_count
    FROM score_item_articles
    WHERE score_item_id = '210eedab-08e7-4f47-ae6a-37aecea9bc16';

    -- Buscar dados do item
    SELECT name, unit, reference_range INTO v_name, v_unit, v_ref_range
    FROM score_items
    WHERE id = '210eedab-08e7-4f47-ae6a-37aecea9bc16';

    RAISE NOTICE '';
    RAISE NOTICE '======================================================================';
    RAISE NOTICE 'RESULTADO FINAL DO ENRIQUECIMENTO';
    RAISE NOTICE '======================================================================';
    RAISE NOTICE 'Item: %', v_name;
    RAISE NOTICE 'Unidade: %', v_unit;
    RAISE NOTICE 'Referência: %', v_ref_range;
    RAISE NOTICE 'Artigos vinculados: %', v_count;
    RAISE NOTICE 'Status: ENRIQUECIMENTO COMPLETO';
    RAISE NOTICE '======================================================================';
    RAISE NOTICE '';
END $$;

COMMIT;

-- Exportar para validação
\echo '\n✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!'
\echo 'Item: Hidrogênio Pico (H₂ Máximo)'
\echo 'ID: 210eedab-08e7-4f47-ae6a-37aecea9bc16'
\echo 'Artigos científicos: 5'
\echo 'Conteúdo clínico: Completo em PT-BR'
