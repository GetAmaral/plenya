-- ========================================
-- ENRIQUECIMENTO: Metano Basal (CH₄ Jejum)
-- Item ID: a92d20ce-702f-4b15-8817-098b9539c0f0
-- Grupo: Exames > Imagem
-- Data: 2026-01-28
-- ========================================

BEGIN;

-- ========================================
-- 1. INSERIR ARTIGOS CIENTÍFICOS
-- ========================================

-- Artigo 1: Consensus Guidelines (2017)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus',
    'Rezaie A, Buresi M, Lembo A, Lin H, McCallum R, Rao S, Schmulson M, Valdovinos M, Zakko S, Pimentel M',
    'American Journal of Gastroenterology',
    '2017-05-01',
    'en',
    '10.1038/ajg.2017.46',
    '28323273',
    'Este consenso norte-americano estabelece diretrizes para testes respiratórios com hidrogênio e metano em distúrbios gastrointestinais. Define que níveis de metano ≥10 ppm devem ser considerados positivos, estabelecendo a correlação entre metano e constipação. Pacientes com supercrescimento bacteriano predominante em metano têm 5 vezes mais probabilidade de apresentar constipação comparado aos casos predominantes em hidrogênio, com a gravidade correlacionando-se diretamente aos níveis de metano.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC5418558/',
    'editorial',
    'Gastroenterologia',
    '["breath test", "methane", "SIBO", "IMO", "constipation", "gastrointestinal disorders", "diagnostic guidelines"]'::jsonb,
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    updated_at = NOW()
RETURNING id AS article1_id \gset

-- Artigo 2: Understanding Hydrogen-Methane Breath Testing (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth',
    'Tansel A, Levinthal DJ',
    'Clinical and Translational Gastroenterology',
    '2023-02-03',
    'en',
    '10.14309/ctg.0000000000000567',
    '36744854',
    'Revisão detalhada sobre testes respiratórios hidrogênio-metano para diagnóstico de SIBO/IMO. Introduz o conceito de supercrescimento de metanógenos intestinais (IMO) para distinguir padrões predominantes em metano do SIBO clássico. Múltiplos estudos identificaram que níveis elevados de metano estão positivamente associados à constipação, com o metano demonstrando inibir diretamente o trânsito intestinal em 59% em modelos animais. Um nível de metano ≥10 ppm em jejum ou em qualquer momento durante o teste define IMO positivo. Abordagem simplificada com medição única de metano em jejum (SMM) ≥10 ppm demonstrou alta performance diagnóstica comparável aos protocolos padrão de 2 horas.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC10132719/',
    'review',
    'Gastroenterologia',
    '["breath test", "methane", "SIBO", "IMO", "intestinal methanogen overgrowth", "constipation", "diagnostic test", "fasting methane"]'::jsonb,
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    updated_at = NOW()
RETURNING id AS article2_id \gset

-- Artigo 3: Methane and Constipation Meta-analysis (2011)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Methane on breath testing is associated with constipation: a systematic review and meta-analysis',
    'Kunkel D, Basseri RJ, Makhani MD, Chong K, Chang C, Pimentel M',
    'Digestive Diseases and Sciences',
    '2011-06-01',
    'en',
    '10.1007/s10620-010-1455-4',
    '21286935',
    'Revisão sistemática e meta-análise demonstrando associação significativa entre metano no teste respiratório e constipação. Encontrou razão de chances (OR) de 3,51 (IC 95%: 2,00-6,16) para a associação entre metano e constipação. Os autores concluem que um fenótipo clínico diferente está associado a níveis mais elevados de produção de metano (constipação) em comparação à produção de hidrogênio (tipicamente diarreia). A constipação está associada a níveis elevados de metano expirado e Methanobrevibacter smithii nas fezes, e o direcionamento terapêutico aos metanógenos pode reduzir a produção de metano e melhorar a constipação.',
    'https://pubmed.ncbi.nlm.nih.gov/21286935/',
    'meta_analysis',
    'Gastroenterologia',
    '["methane", "breath test", "constipation", "meta-analysis", "Methanobrevibacter smithii", "gastrointestinal motility"]'::jsonb,
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    updated_at = NOW()
RETURNING id AS article3_id \gset

-- ========================================
-- 2. ATUALIZAR SCORE_ITEM COM CONTEÚDO CLÍNICO
-- ========================================

UPDATE score_items
SET
    clinical_relevance = 'O metano basal (CH₄) no teste respiratório em jejum é um biomarcador fundamental para diagnóstico de supercrescimento de metanógenos intestinais (IMO - Intestinal Methanogen Overgrowth). Valores ≥10 ppm definem IMO positivo e estão fortemente associados à constipação crônica. Metanógenos, especialmente Methanobrevibacter smithii, são arqueias (não bactérias) que produzem metano e inibem diretamente o trânsito intestinal em até 59%. Pacientes com IMO têm 5 vezes mais probabilidade de apresentar constipação comparado ao SIBO clássico (hidrogênio-predominante). A medição em jejum tem sensibilidade de 86,4% e especificidade de 100% comparada aos testes de 2 horas, permitindo diagnóstico simplificado. A severidade da constipação correlaciona-se diretamente com os níveis de metano. IMO também está associado a sintomas inespecíficos como distensão abdominal e dor, mas a constipação é o fenótipo clínico dominante.',

    patient_explanation = 'O metano basal é medido no ar que você expira após jejum de 8-12 horas. Esse gás é produzido por micro-organismos chamados metanógenos que vivem no seu intestino. Quando o nível de metano está alto (≥10 ppm), indica que há um crescimento excessivo desses organismos, condição chamada IMO. O metano tem um efeito especial: ele deixa o intestino mais lento, dificultando a passagem das fezes e causando prisão de ventre. Por isso, pessoas com IMO frequentemente apresentam constipação crônica, além de sintomas como inchaço abdominal, gases e desconforto. O teste é simples - você apenas sopra em um tubo especial enquanto está em jejum. Se o resultado for positivo, existem tratamentos específicos com antibióticos que agem nos metanógenos (diferentes dos usados para bactérias comuns) e mudanças na alimentação que podem ajudar a melhorar o funcionamento do intestino.',

    conduct = 'Valores de metano basal ≥10 ppm em jejum confirmam IMO e requerem abordagem terapêutica específica. PRIMEIRA LINHA: Rifaximina 550mg 3x/dia por 14 dias combinada com neomicina 500mg 2x/dia por 14 dias (neomicina é essencial pois rifaximina sozinha tem eficácia limitada contra metanógenos). ALTERNATIVA: Metronidazol 500mg 3x/dia por 14 dias se neomicina não disponível. Methanobrevibacter smithii é resistente a muitos antibióticos convencionais, exigindo seleção cuidadosa. SUPORTE NUTRICIONAL: Dieta FODMAP baixa por 4-6 semanas durante e após antibioticoterapia. Evitar carboidratos fermentáveis (lactose, frutose, sorbitol, oligossacarídeos). PROCINÉTICOS: Considerar prucaloprida 2mg/dia ou linaclotida 290mcg/dia para melhorar trânsito intestinal lentificado. PROBIÓTICOS: Evitar Bifidobacterium spp durante tratamento (podem produzir gases); preferir Saccharomyces boulardii ou Lactobacillus plantarum após antibioticoterapia. REAVALIAÇÃO: Repetir teste respiratório 4 semanas após término do tratamento. Taxa de recorrência é alta (40-50%), podendo necessitar ciclos adicionais ou terapia de manutenção. INVESTIGAR CAUSAS: Avaliar dismotilidade intestinal (pseudo-obstrução, diabetes), uso crônico de IBP (suprimir ácido gástrico facilita colonização), divertículos, estenoses. SEGUIMENTO: Se persistir sintomático apesar de metano normalizado, considerar outros diagnósticos (SII, intolerâncias alimentares, disbiose bacteriana concomitante).',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = 'a92d20ce-702f-4b15-8817-098b9539c0f0';

-- Verificar se atualização foi bem sucedida
DO $$
DECLARE
    updated_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO updated_count
    FROM score_items
    WHERE id = 'a92d20ce-702f-4b15-8817-098b9539c0f0'
    AND clinical_relevance IS NOT NULL
    AND patient_explanation IS NOT NULL
    AND conduct IS NOT NULL;

    IF updated_count = 0 THEN
        RAISE EXCEPTION 'Falha ao atualizar score_item a92d20ce-702f-4b15-8817-098b9539c0f0';
    END IF;

    RAISE NOTICE 'Score item atualizado com sucesso!';
END $$;

-- ========================================
-- 3. CRIAR RELAÇÕES ARTICLE-SCORE_ITEM
-- ========================================

-- Relacionar todos os 3 artigos ao score_item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'a92d20ce-702f-4b15-8817-098b9539c0f0'
FROM articles
WHERE doi IN (
    '10.1038/ajg.2017.46',
    '10.14309/ctg.0000000000000567',
    '10.1007/s10620-010-1455-4'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ========================================
-- 4. VERIFICAÇÃO FINAL
-- ========================================

SELECT
    '=== VERIFICAÇÃO FINAL ===' AS status,
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) AS clinical_length,
    LENGTH(si.patient_explanation) AS patient_length,
    LENGTH(si.conduct) AS conduct_length,
    si.last_review,
    COUNT(DISTINCT a.id) AS articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
LEFT JOIN articles a ON a.id = asi.article_id
WHERE si.id = 'a92d20ce-702f-4b15-8817-098b9539c0f0'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Listar artigos vinculados
SELECT
    '=== ARTIGOS VINCULADOS ===' AS status,
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi
FROM articles a
JOIN article_score_items asi ON asi.article_id = a.id
WHERE asi.score_item_id = 'a92d20ce-702f-4b15-8817-098b9539c0f0'
ORDER BY a.publish_date DESC;

COMMIT;

-- ========================================
-- FIM DO SCRIPT
-- ========================================
