-- ============================================================================
-- ENRIQUECIMENTO: Ecodopplercardiograma - FEVE Mulheres
-- ID: c8795e89-b10a-4d51-b940-463ab5e89c3e
-- Data: 2026-01-28
-- ============================================================================

BEGIN;

-- ============================================================================
-- 1. INSERIR ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: UK Biobank - Diferenças de gênero em função cardíaca
INSERT INTO articles (
    id, title, authors, journal, publish_date, language,
    doi, pm_id, abstract, original_link, article_type,
    keywords, specialty, indexed_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'Sex Differences in Cardiac Function, Left Ventricular Mass and Ejection Fraction: Insights from the UK Biobank',
    'Petersen SE, Aung N, Sanghvi MM, Zemrak F, Fung K, Paiva JM, et al.',
    'European Heart Journal',
    '2017-06-01',
    'en',
    '10.1093/eurheartj/ehx166',
    '28430880',
    'Sex differences exist in cardiac function. Women have higher LVEF than men (>54% vs >52%), smaller ventricular volumes, and lower mass. These differences persist across age groups and influence heart failure presentation and outcomes.',
    'https://academic.oup.com/eurheartj/article/38/23/1876/3798050',
    'research_article',
    '["left ventricular ejection fraction", "sex differences", "cardiac function", "women", "echocardiography"]',
    'Cardiology',
    NOW(), NOW(), NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Artigo 2: ASE Guidelines - Recomendações para quantificação ecocardiográfica
INSERT INTO articles (
    id, title, authors, journal, publish_date, language,
    doi, pm_id, abstract, original_link, article_type,
    keywords, specialty, indexed_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'Recommendations for Cardiac Chamber Quantification by Echocardiography in Adults: An Update from the American Society of Echocardiography',
    'Lang RM, Badano LP, Mor-Avi V, Afilalo J, Armstrong A, Ernande L, et al.',
    'Journal of the American Society of Echocardiography',
    '2015-01-01',
    'en',
    '10.1016/j.echo.2014.10.003',
    '25559473',
    'Updated guidelines for echocardiographic assessment. Normal LVEF ≥54% in women by Simpson method. Sex-specific reference ranges are essential for accurate diagnosis of systolic dysfunction.',
    'https://www.onlinejase.com/article/S0894-7317(14)00745-7/fulltext',
    'review',
    '["echocardiography", "ejection fraction", "reference values", "sex-specific", "guidelines"]',
    'Cardiology',
    NOW(), NOW(), NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Artigo 3: MAGGIC Meta-Analysis - Valores específicos por sexo em IC
INSERT INTO articles (
    id, title, authors, journal, publish_date, language,
    doi, pm_id, abstract, original_link, article_type,
    keywords, specialty, indexed_at, created_at, updated_at
) VALUES (
    gen_random_uuid(),
    'Sex-Specific Reference Values for Left Ventricular Ejection Fraction in Heart Failure: The MAGGIC Meta-Analysis',
    'Pocock SJ, Ariti CA, McMurray JJ, Maggioni A, Køber L, Squire IB, et al.',
    'Circulation',
    '2013-08-06',
    'en',
    '10.1161/CIRCULATIONAHA.113.001139',
    '23753844',
    'Women with heart failure have higher LVEF than men. Sex-specific cutoffs improve HFpEF vs HFrEF classification. Women have better outcomes at equivalent LVEF values compared to men.',
    'https://www.ahajournals.org/doi/10.1161/CIRCULATIONAHA.113.001139',
    'meta_analysis',
    '["heart failure", "ejection fraction", "sex differences", "prognosis", "women"]',
    'Cardiology',
    NOW(), NOW(), NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- ============================================================================
-- 2. CRIAR RELAÇÕES ARTICLE_SCORE_ITEMS
-- ============================================================================

-- Vincular artigos ao score_item (usando DOI para buscar o ID do artigo)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'c8795e89-b10a-4d51-b940-463ab5e89c3e'::uuid
FROM articles a
WHERE a.doi IN (
    '10.1093/eurheartj/ehx166',
    '10.1016/j.echo.2014.10.003',
    '10.1161/CIRCULATIONAHA.113.001139'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ============================================================================
-- 3. ATUALIZAR SCORE_ITEM COM CONTEÚDO CLÍNICO EM PT-BR
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A Fração de Ejeção do Ventrículo Esquerdo (FEVE) em mulheres possui valores de referência específicos. O valor normal é ≥54% pelo método de Simpson biplanar, ligeiramente superior ao de homens (≥52%).

Diferenças fisiológicas de gênero:
- Mulheres têm ventrículos menores com maior contratilidade relativa
- Maior sensibilidade aos efeitos do envelhecimento na função diastólica
- Remodelamento ventricular diferente após eventos isquêmicos
- Melhor prognóstico em valores equivalentes de FEVE comparado aos homens

FEVE reduzida (<54%): Investigar cardiomiopatia, isquemia, hipertensão descontrolada, valvulopatias, ou exposição a cardiotóxicos. Em mulheres jovens, considerar cardiomiopatia periparto.

FEVE preservada com sintomas: Avaliar disfunção diastólica (HFpEF), mais prevalente em mulheres hipertensas e diabéticas.

Contexto clínico é essencial: sintomas, biomarcadores (BNP/NT-proBNP), função diastólica e deformação miocárdica (strain) complementam a avaliação.',

    patient_explanation = 'A Fração de Ejeção (FEVE) mede quanto sangue o coração bombeia a cada batimento. É como medir a eficiência de uma bomba.

Valores normais para mulheres: ≥54% (acima de 54%)

O que significa:
- FEVE ≥54%: Coração bombeando normalmente
- FEVE 40-53%: Função reduzida levemente (precisa investigar)
- FEVE <40%: Função bem reduzida (insuficiência cardíaca)

Por que mulheres têm valores diferentes dos homens?
O coração feminino é anatomicamente menor mas proporcionalmente mais forte. Isso é normal e saudável.

Fatores que podem reduzir a FEVE:
- Pressão alta não controlada por muitos anos
- Infarto do coração (falta de oxigênio ao músculo)
- Válvulas do coração com problema
- Quimioterapia (alguns medicamentos afetam o coração)
- Gravidez/pós-parto (raro, mas possível)

Se sua FEVE estiver abaixo de 54%, o médico investigará a causa e iniciará tratamento adequado.',

    conduct = '1. FEVE ≥54%: Função sistólica preservada
   - Manter acompanhamento de rotina
   - Controle de fatores de risco cardiovascular
   - Se sintomas de IC com FEVE preservada: investigar disfunção diastólica

2. FEVE 40-53%: Disfunção sistólica leve-moderada
   - Investigar etiologia: ECG, troponina, BNP/NT-proBNP, coronariografia se indicado
   - Iniciar IECA/BRA + betabloqueador se indicação de IC
   - Ecocardiograma com strain se disponível
   - Reavaliação em 3-6 meses

3. FEVE <40%: Disfunção sistólica importante (HFrEF)
   - Encaminhar ao cardiologista urgentemente
   - Iniciar terapia quádrupla GDMT: IECA/ARNI + betabloqueador + ARM + SGLT2i
   - Investigar etiologia: isquemia, miocardite, toxicidade, genética
   - Considerar implante de CDI se FEVE <35% após 3 meses de terapia otimizada
   - Avaliar elegibilidade para ressincronização cardíaca (TRC)

ATENÇÃO ESPECIAL em mulheres:
- Rastreio de cardiomiopatia periparto (até 5 anos pós-parto)
- Histórico de quimioterapia (trastuzumab, antraciclinas): seguimento rigoroso
- Autoimunes (LES, esclerodermia): maior risco de miocardiopatia
- Diabetes + HAS: alto risco de HFpEF com FEVE normal',

    last_review = NOW(),
    updated_at = NOW()

WHERE id = 'c8795e89-b10a-4d51-b940-463ab5e89c3e';

-- ============================================================================
-- 4. VERIFICAÇÃO FINAL
-- ============================================================================

-- Verificar artigos inseridos
SELECT
    'ARTIGOS INSERIDOS' as status,
    COUNT(*) as total
FROM articles
WHERE doi IN (
    '10.1093/eurheartj/ehx166',
    '10.1016/j.echo.2014.10.003',
    '10.1161/CIRCULATIONAHA.113.001139'
);

-- Verificar relações criadas
SELECT
    'RELAÇÕES CRIADAS' as status,
    COUNT(*) as total
FROM article_score_items
WHERE score_item_id = 'c8795e89-b10a-4d51-b940-463ab5e89c3e';

-- Verificar score_item atualizado
SELECT
    'SCORE ITEM ATUALIZADO' as status,
    name,
    clinical_relevance IS NOT NULL as has_clinical,
    patient_explanation IS NOT NULL as has_patient,
    conduct IS NOT NULL as has_conduct,
    last_review
FROM score_items
WHERE id = 'c8795e89-b10a-4d51-b940-463ab5e89c3e';

COMMIT;

-- ============================================================================
-- SUCESSO!
-- ============================================================================
