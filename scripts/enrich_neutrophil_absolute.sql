-- ================================================================================
-- ENRIQUECIMENTO: Neutrófilos (absoluto)
-- ID: 3faeb6db-b8d6-4fb1-8740-07bfd91002c7
-- Timestamp: 2026-01-28
-- ================================================================================

-- Começar transação
BEGIN;

-- ================================================================================
-- FASE 1: Criar Artigos Científicos
-- ================================================================================

-- Artigo 1: CTCAE v6 (2025-2026)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM articles WHERE original_link = 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12745037/') THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            abstract,
            full_content,
            original_link,
            created_at,
            updated_at
        ) VALUES (
            gen_random_uuid(),
            'A paradigm shift in neutrophil adverse event grading: What now?',
            'PMC Editorial Team',
            'PMC - PubMed Central',
            '2025-01-01',
            'review',
            E'CTCAE v6 (2025) atualiza classificação de neutropenia: Grade 1 agora <1500-1000/µL (antes Grade 2), Grade 4 <100/µL.',
            E'**Key Findings:** CTCAE v6 (2025) atualiza classificação de neutropenia: Grade 1 agora <1500-1000/µL (antes Grade 2), Grade 4 <100/µL. Mudanças visam inclusão de variante Duffy null (comum em pessoas com ancestralidade africana subsaariana).\n\n**Clinical Significance:** Esta atualização reconhece a diversidade genética populacional e reduz exclusão desnecessária de pacientes em ensaios clínicos.',
            'https://pmc.ncbi.nlm.nih.gov/articles/PMC12745037/',
            NOW(),
            NOW()
        );
    ELSE
        UPDATE articles SET
            title = 'A paradigm shift in neutrophil adverse event grading: What now?',
            full_content = E'**Key Findings:** CTCAE v6 (2025) atualiza classificação de neutropenia: Grade 1 agora <1500-1000/µL (antes Grade 2), Grade 4 <100/µL. Mudanças visam inclusão de variante Duffy null (comum em pessoas com ancestralidade africana subsaariana).\n\n**Clinical Significance:** Esta atualização reconhece a diversidade genética populacional e reduz exclusão desnecessária de pacientes em ensaios clínicos.',
            updated_at = NOW()
        WHERE original_link = 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12745037/';
    END IF;
END $$;

-- Capturar o ID do artigo 1 para vinculação
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE original_link = 'https://pmc.ncbi.nlm.nih.gov/articles/PMC12745037/';

    -- Vincular artigo 1 ao score_item
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '3faeb6db-b8d6-4fb1-8740-07bfd91002c7')
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- Artigo 2: AGIHO 2024 Guidelines
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM articles WHERE original_link = 'https://www.thelancet.com/journals/lanepe/article/PIIS2666-7762(25)00006-7/fulltext') THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            abstract,
            full_content,
            original_link,
            specialty,
            created_at,
            updated_at
        ) VALUES (
            gen_random_uuid(),
            '2024 update of the AGIHO guideline on diagnosis and empirical treatment of fever of unknown origin in adult neutropenic patients',
            'AGIHO Working Group',
            'The Lancet Regional Health – Europe',
            '2025-01-15',
            'clinical_trial',
            E'Diretrizes atualizadas para manejo de neutropenia febril em pacientes com tumores sólidos e malignidades hematológicas.',
            E'**Key Findings:** Diretrizes atualizadas para manejo de neutropenia febril. Monoterapia empírica com beta-lactâmicos anti-pseudomonas é primeira linha. Estratificação de risco via índice MASCC.\n\n**Recommendations:** Antibioticoterapia empírica imediata (<60min) em neutropenia febril. Beta-lactâmicos anti-pseudomonas (piperacilina-tazobactam, cefepime, meropenem) são primeira linha baseados em evidências de alta certeza.',
            'https://www.thelancet.com/journals/lanepe/article/PIIS2666-7762(25)00006-7/fulltext',
            'Hematologia, Infectologia',
            NOW(),
            NOW()
        );
    END IF;
END $$;

-- Vincular artigo 2
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE original_link = 'https://www.thelancet.com/journals/lanepe/article/PIIS2666-7762(25)00006-7/fulltext';

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '3faeb6db-b8d6-4fb1-8740-07bfd91002c7')
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- Artigo 3: StatPearls - Febrile Neutropenia
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM articles WHERE original_link = 'https://www.ncbi.nlm.nih.gov/books/NBK541102/') THEN
        INSERT INTO articles (
            id,
            title,
            authors,
            journal,
            publish_date,
            article_type,
            abstract,
            full_content,
            original_link,
            pm_id,
            specialty,
            created_at,
            updated_at
        ) VALUES (
            gen_random_uuid(),
            'Febrile Neutropenia',
            'StatPearls Authors',
            'StatPearls Publishing - NCBI Bookshelf',
            '2025-01-01',
            'review',
            E'Neutropenia febril é emergência oncológica que requer manejo imediato e baseado em evidências.',
            E'**Key Findings:** Neutropenia febril definida como temperatura ≥38.3°C (ou ≥38°C por >1h) com ANC <1000/µL. Requer avaliação emergencial, hemoculturas (2 conjuntos: periférica + cateter) e antibioticoterapia empírica imediata.\n\n**Management Protocol:** Tempo porta-antibiótico <60min. Coleta de hemoculturas pré-antibiótico. Estratificação de risco MASCC (≥21 pontos = baixo risco). Monoterapia com beta-lactâmico anti-pseudomonas como primeira linha.',
            'https://www.ncbi.nlm.nih.gov/books/NBK541102/',
            'NBK541102',
            'Hematologia, Oncologia, Infectologia',
            NOW(),
            NOW()
        );
    END IF;
END $$;

-- Vincular artigo 3
DO $$
DECLARE
    v_article_id UUID;
BEGIN
    SELECT id INTO v_article_id FROM articles WHERE original_link = 'https://www.ncbi.nlm.nih.gov/books/NBK541102/';

    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (v_article_id, '3faeb6db-b8d6-4fb1-8740-07bfd91002c7')
    ON CONFLICT (score_item_id, article_id) DO NOTHING;
END $$;

-- ================================================================================
-- FASE 2: Atualizar Conteúdo Clínico do Score Item
-- ================================================================================

UPDATE score_items SET
    clinical_relevance = E'**Contagem Absoluta de Neutrófilos (ANC): Pilar da Avaliação Imunológica**

Os neutrófilos são a primeira linha de defesa contra infecções bacterianas e fúngicas. A contagem absoluta de neutrófilos (ANC) é calculada multiplicando o percentual de neutrófilos (segmentados + bastonetes) pela contagem total de leucócitos.

**ATUALIZAÇÃO CTCAE v6 (2025-2026):**
O Common Terminology Criteria for Adverse Events versão 6, implementado em janeiro de 2026, modernizou a classificação de neutropenia:

• **Grau 1:** ANC 1000-1500/µL (antes era Grau 2)
• **Grau 2:** ANC 500-1000/µL (antes era Grau 3)
• **Grau 3:** ANC 100-500/µL (antes era Grau 4)
• **Grau 4:** ANC <100/µL (novo limiar)

Esta mudança reconhece que pessoas com variante genética Duffy null (comum em populações de ancestralidade africana subsaariana) apresentam ANC naturalmente mais baixo sem aumento de risco infeccioso.

**Valores de Referência:**
• Normal: 1.500-8.000/µL (algumas fontes: 2.500-7.000/µL)
• Neutropenia leve: 1.000-1.500/µL
• Neutropenia moderada: 500-1.000/µL
• Neutropenia grave: <500/µL
• Neutropenia profunda: <100/µL

**Neutropenia Febril (Emergência Médica):**
Definida como temperatura ≥38,3°C (ou ≥38°C por >1h) + ANC <1.000/µL. Requer hospitalização imediata, hemoculturas e antibioticoterapia empírica.

**Neutrofilia:**
ANC >7.700/µL indica resposta inflamatória, infecção bacteriana aguda, estresse fisiológico (cirurgia, trauma), uso de corticoides ou doenças mieloproliferativas.

**Desvio à Esquerda:**
Aumento de neutrófilos jovens (bastonetes, metamielócitos) sugere infecção bacteriana grave com mobilização acelerada da medula óssea.',

    patient_explanation = E'**O que são neutrófilos e por que a contagem absoluta é importante?**

Neutrófilos são glóbulos brancos especializados em combater infecções causadas por bactérias e fungos. Eles atuam como "soldados de primeira linha" do seu sistema imunológico, chegando rapidamente ao local da infecção para destruir os invasores.

A **contagem absoluta de neutrófilos (ANC)** mede quantos desses soldados você tem circulando no sangue. Este número é calculado automaticamente pelo laboratório usando a contagem total de leucócitos e a porcentagem de neutrófilos.

**Valores normais:**
• Adultos: 1.500 a 8.000 neutrófilos por microlitro (µL)
• Crianças: acima de 1.500/µL

**O que significa ter neutrófilos baixos (neutropenia)?**

Quando seus neutrófilos estão abaixo de 1.500/µL, você tem neutropenia, o que aumenta o risco de infecções. Quanto mais baixa a contagem, maior o risco:

• **1.000-1.500/µL:** Risco levemente aumentado
• **500-1.000/µL:** Risco moderado - cuidados extras necessários
• **Abaixo de 500/µL:** Risco grave - qualquer febre é emergência médica

**IMPORTANTE - Neutropenia Febril é Emergência:**
Se você tem neutrófilos baixos E desenvolve febre (≥38,3°C), procure atendimento médico IMEDIATAMENTE. Esta combinação pode evoluir rapidamente para infecção grave, pois seu corpo não consegue se defender adequadamente.

**O que significa ter neutrófilos altos (neutrofilia)?**

Valores acima de 7.700/µL geralmente indicam:
• Infecção bacteriana em curso
• Resposta ao estresse físico (cirurgia, trauma)
• Uso de medicamentos (corticoides)
• Inflamação sistêmica
• Raramente: doenças da medula óssea

**Observação importante sobre diversidade genética:**
Algumas pessoas, especialmente aquelas com ancestralidade africana, têm contagens naturalmente mais baixas (1.000-1.500/µL) devido à variante genética Duffy null. Isso NÃO significa risco aumentado de infecção - é uma variação normal. Os novos critérios médicos (2025) já consideram isso.',

    conduct = E'**Condutas Clínicas Baseadas em ANC (Atualizadas 2025)**

**1. ANC NORMAL (1.500-8.000/µL):**
✓ Nenhuma ação específica necessária
✓ Sistema imunológico funcionando adequadamente
✓ Manter acompanhamento de rotina

**2. NEUTROPENIA LEVE (1.000-1.500/µL) - CTCAE v6 Grau 1:**
→ Investigar causas:
  • Medicamentos (anticonvulsivantes, antibióticos, quimioterapia)
  • Deficiências nutricionais (B12, folato, cobre)
  • Infecções virais recentes (influenza, HIV, hepatites)
  • Doenças autoimunes (lúpus, artrite reumatoide)
  • Variante genética Duffy null (considerar em afrodescendentes)

→ Monitoramento: Repetir hemograma em 1-2 semanas
→ Orientações ao paciente:
  • Higiene rigorosa das mãos
  • Evitar contato com pessoas doentes
  • Cozinhar bem carnes e ovos
  • Lavar frutas/vegetais adequadamente

**3. NEUTROPENIA MODERADA (500-1.000/µL) - CTCAE v6 Grau 2:**
→ Avaliação hematológica obrigatória
→ Investigação ampliada:
  • Hemograma completo seriado
  • Esfregaço periférico
  • Dosagem de vitaminas (B12, folato)
  • Sorologias virais (se indicado)
  • Considerar aspirado de medula óssea

→ Precauções:
  • Evitar aglomerações e locais fechados
  • Máscara em ambientes públicos se indicado
  • Profilaxia antibiótica em casos selecionados
  • Educação sobre sinais de infecção

**4. NEUTROPENIA GRAVE (<500/µL) - CTCAE v6 Graus 3-4:**
→ ALTO RISCO INFECCIOSO - Monitoramento intensivo
→ Hospitalização se sinais de infecção
→ Considerar:
  • Fator estimulador de colônias (G-CSF/filgrastim)
  • Profilaxia antibiótica (fluoroquinolona)
  • Profilaxia antifúngica se prolongada (>7 dias)
  • Isolamento protetor se <100/µL

**5. NEUTROPENIA FEBRIL (ANC <1.000/µL + febre ≥38,3°C):**
🚨 **EMERGÊNCIA MÉDICA - PROTOCOLO IMEDIATO:**

→ Tempo porta-antibiótico: <60 minutos
→ Colher ANTES de antibióticos:
  • 2 conjuntos de hemoculturas (periférica + cateter se presente)
  • Urinocultura
  • Culturas de sítios específicos (feridas, cateter)

→ Antibioticoterapia empírica (Guidelines AGIHO 2024):
  **Primeira linha:** Beta-lactâmico anti-pseudomonas
  • Piperacilina-tazobactam 4,5g IV 6/6h OU
  • Cefepime 2g IV 8/8h OU
  • Meropenem 1g IV 8/8h (se colonização ESBL/alta prevalência)

→ Estratificação de risco (Índice MASCC):
  • ≥21 pontos: Baixo risco → considerar alta precoce/oral
  • <21 pontos: Alto risco → hospitalização prolongada

→ Adicionar se indicado:
  • Vancomicina (suspeita de infecção por Gram+ ou cateter)
  • Antifúngico empírico (febre persistente >4-7 dias)

**6. NEUTROFILIA (>7.700/µL):**
→ Investigar causa subjacente:
  • Sinais de infecção bacteriana
  • Inflamação sistêmica (PCR, VHS)
  • Revisar medicações (corticoides)
  • Descartar leucemia mieloide/mieloproliferação (se muito elevado)

→ Avaliar "desvio à esquerda":
  • Aumento de bastonetes (>10%) sugere infecção bacteriana aguda
  • Presença de metamielócitos/mielócitos = mobilização medular intensa

**7. ELEGIBILIDADE PARA ENSAIOS CLÍNICOS (Atualização 2025):**
→ Novo critério padrão: ANC ≥1.000/µL (anteriormente ≥1.500/µL)
→ Reflete mudança CTCAE v6 e inclusão de populações diversas

**MONITORAMENTO PÓS-QUIMIOTERAPIA:**
→ Nadir esperado: 7-14 dias após infusão
→ Hemograma 2-3x/semana durante nadir
→ G-CSF profilático se risco >20% de neutropenia febril

**REFERÊNCIAS CRÍTICAS:**
• CTCAE v6 (2025) - Nova classificação de neutropenia
• AGIHO 2024 Guidelines - Manejo de febre de origem indeterminada em neutropênicos
• MASCC Index - Estratificação de risco em neutropenia febril',

    last_review = CURRENT_DATE,
    updated_at = NOW()

WHERE id = '3faeb6db-b8d6-4fb1-8740-07bfd91002c7';

-- ================================================================================
-- VERIFICAÇÃO FINAL
-- ================================================================================

-- Contar artigos vinculados
SELECT
    si.id,
    si.name,
    COUNT(asi.article_id) as total_articles
FROM score_items si
LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
WHERE si.id = '3faeb6db-b8d6-4fb1-8740-07bfd91002c7'
GROUP BY si.id, si.name;

-- Listar artigos vinculados
SELECT
    a.title,
    a.journal,
    EXTRACT(YEAR FROM a.publish_date) as year,
    a.original_link
FROM articles a
JOIN article_score_items asi ON asi.article_id = a.id
WHERE asi.score_item_id = '3faeb6db-b8d6-4fb1-8740-07bfd91002c7';

-- Verificar conteúdo atualizado
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_length,
    LENGTH(patient_explanation) as patient_explanation_length,
    LENGTH(conduct) as conduct_length,
    last_review
FROM score_items
WHERE id = '3faeb6db-b8d6-4fb1-8740-07bfd91002c7';

-- Commit da transação
COMMIT;

-- ================================================================================
-- SUMÁRIO
-- ================================================================================
-- ✅ 3 artigos científicos criados (CTCAE v6, AGIHO 2024, StatPearls)
-- ✅ 3 vinculações article-score_item criadas
-- ✅ Conteúdo clínico atualizado (PT-BR):
--    • clinical_relevance: ~2.100 caracteres
--    • patient_explanation: ~1.800 caracteres
--    • conduct: ~3.800 caracteres
-- ✅ last_review atualizado para data atual
-- ================================================================================
