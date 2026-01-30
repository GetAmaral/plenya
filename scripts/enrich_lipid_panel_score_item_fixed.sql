-- ============================================================================
-- ENRICHMENT: Perfil lipídico / Lipidograma
-- Score Item ID: d17cea8f-2935-40bc-b2e0-bcd9ac886ade
-- ============================================================================
-- Author: AI Enrichment System
-- Date: 2026-01-29
-- Database: PostgreSQL 17
-- Schema Version: Current production schema
--
-- CRITICAL SCHEMA NOTES:
-- - Table: "articles" (NOT scientific_articles)
-- - Junction: "article_score_items" (NOT score_item_articles)
-- - Column: "pm_id" (NOT pmid)
-- - Column: "publish_date" as date type (NOT year)
-- - NO url column exists
-- - NO created_at in junction table
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: UPDATE CLINICAL CONTENT FOR SCORE ITEM
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O perfil lipídico completo é um exame fundamental na estratificação de risco cardiovascular, permitindo avaliação abrangente do metabolismo de lipoproteínas e identificação precoce de dislipidemia. Inclui dosagens de colesterol total (CT), HDL-colesterol, LDL-colesterol, triglicerídeos e colesterol não-HDL. As diretrizes ACC/AHA 2018-2024 estabelecem que o LDL-C permanece como alvo primário de terapia, com metas agressivas (<55 mg/dL para muito alto risco, <70 mg/dL para alto risco, <100 mg/dL para risco intermediário). O colesterol não-HDL representa a carga aterogênica total (LDL + VLDL + IDL + remanescentes de quilomícrons) e é especialmente útil quando triglicerídeos >200 mg/dL ou em pacientes com diabetes mellitus, obesidade abdominal e síndrome metabólica. A relação CT/HDL e triglicerídeos/HDL auxiliam na identificação de partículas LDL pequenas e densas (fenotípico B, mais aterogênicas). Estudos recentes demonstram que apolipoproteína B (apoB) e colesterol não-HDL são preditores superiores ao LDL-C isolado em pacientes em uso de estatinas, podendo prevenir 300.000-500.000 eventos isquêmicos cardíacos adicionais em 10 anos. O rastreamento pode ser realizado com amostra em jejum ou não-jejum, sendo jejum de 8-12h preferível quando triglicerídeos >400 mg/dL. Reavaliação a cada 4-8 semanas após início ou ajuste de terapia hipolipemiante é recomendada para avaliar resposta terapêutica e aderência medicamentosa. A adição de lipoproteína(a) ao painel básico é crescentemente recomendada pela National Lipid Association 2024, afetando aproximadamente 1,4 bilhão de pessoas globalmente como fator de risco genético independente.',

    patient_explanation = 'O exame de perfil lipídico (também chamado lipidograma) mede diferentes tipos de gorduras no seu sangue que podem aumentar o risco de problemas no coração e vasos sanguíneos. Ele inclui várias medidas: colesterol total (a soma de todos os tipos), HDL (o "colesterol bom" que protege o coração), LDL (o "colesterol ruim" que pode entupir as artérias), triglicerídeos (outro tipo de gordura) e colesterol não-HDL (que mostra toda a gordura prejudicial). Valores elevados de LDL e triglicerídeos ou valores baixos de HDL aumentam muito o risco de infarto e derrame. O exame geralmente não precisa de jejum, mas seu médico pode solicitar 8-12 horas de jejum em algumas situações. É recomendado para todos os adultos a partir dos 20 anos e deve ser repetido periodicamente conforme seu risco cardiovascular. Se os resultados estiverem alterados, o médico pode recomendar mudanças na alimentação, exercícios físicos e, quando necessário, medicamentos como estatinas. O objetivo é manter o LDL bem baixo (especialmente se você tem diabetes, pressão alta ou já teve problemas no coração) e o HDL elevado. Triglicerídeos muito altos também precisam de tratamento específico. Controlar essas gorduras no sangue é uma das formas mais eficazes de prevenir doenças cardiovasculares.',

    conduct = 'RASTREAMENTO: Realizar perfil lipídico completo em todos adultos ≥20 anos (ACC/AHA 2019). Repetir a cada 4-6 anos se normal em baixo risco; anualmente em pacientes com fatores de risco (diabetes, hipertensão, tabagismo, história familiar). Pode ser realizado sem jejum para rastreamento inicial; jejum de 8-12h preferível se triglicerídeos >400 mg/dL ou para seguimento de hipertrigliceridemia. INTERPRETAÇÃO: Calcular risco ASCVD em 10 anos (pooled cohort equations) em adultos 40-75 anos. Avaliar fatores intensificadores de risco: história familiar prematura de DCV, LDL persistente ≥160 mg/dL, síndrome metabólica, doença renal crônica, pré-eclâmpsia, menopausa precoce, doenças inflamatórias crônicas, etnia sul-asiática. Considerar colesterol não-HDL (meta: 30 mg/dL acima da meta de LDL) e apoB se disponível, especialmente em triglicerídeos elevados, diabetes ou obesidade. METAS TERAPÊUTICAS (ACC/AHA 2022): Prevenção primária baixo risco (<5%): LDL <130 mg/dL; risco intermediário (5-19,9%): LDL <100 mg/dL, considerar <70 mg/dL com intensificadores; alto risco (≥20%): LDL <70 mg/dL. Prevenção secundária (DCV estabelecida): LDL <55 mg/dL (muito alto risco) ou <70 mg/dL (alto risco), redução ≥50% do basal. TRATAMENTO: Iniciar estatina de alta intensidade (atorvastatina 40-80mg ou rosuvastatina 20-40mg) em prevenção secundária e primária de alto risco. Intensidade moderada em risco intermediário após discussão médico-paciente. Se meta não atingida com estatina máxima tolerada, adicionar ezetimiba 10mg (redução adicional de 15-20% LDL). Considerar inibidor PCSK9 se LDL persistente ≥70 mg/dL em muito alto risco. Para triglicerídeos 200-499 mg/dL: otimizar controle glicêmico, perda ponderal, reduzir álcool e carboidratos refinados; considerar fenofibrato ou ômega-3 em casos selecionados. Triglicerídeos ≥500 mg/dL: risco de pancreatite aguda, iniciar fibrato imediatamente. MONITORAMENTO: Perfil lipídico 4-12 semanas após início/ajuste de terapia, depois a cada 3-12 meses. Avaliar enzimas hepáticas basalmente e se sintomas. Investigar miopatia se CPK >10x LSN ou sintomas musculares significativos. Reavaliar anualmente fatores de risco e aderência medicamentosa.',

    last_review = CURRENT_DATE
WHERE id = 'd17cea8f-2935-40bc-b2e0-bcd9ac886ade';

-- Verify update
DO $$
DECLARE
    v_clinical_len INTEGER;
    v_patient_len INTEGER;
    v_conduct_len INTEGER;
BEGIN
    SELECT
        LENGTH(clinical_relevance),
        LENGTH(patient_explanation),
        LENGTH(conduct)
    INTO v_clinical_len, v_patient_len, v_conduct_len
    FROM score_items
    WHERE id = 'd17cea8f-2935-40bc-b2e0-bcd9ac886ade';

    RAISE NOTICE 'Content lengths - Clinical: %, Patient: %, Conduct: %',
        v_clinical_len, v_patient_len, v_conduct_len;

    IF v_clinical_len < 1500 OR v_patient_len < 1000 OR v_conduct_len < 1500 THEN
        RAISE EXCEPTION 'Content too short - Clinical: % (min 1500), Patient: % (min 1000), Conduct: % (min 1500)',
            v_clinical_len, v_patient_len, v_conduct_len;
    END IF;

    RAISE NOTICE '✓ Score item content updated successfully';
END $$;

-- ============================================================================
-- STEP 2: INSERT SCIENTIFIC ARTICLES
-- ============================================================================

-- Article 1: National Lipid Association Expert Consensus on ApoB (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    article_type,
    publish_date,
    abstract,
    doi,
    language
) VALUES (
    gen_random_uuid(),
    'Role of apolipoprotein B in the clinical management of cardiovascular risk in adults: An Expert Clinical Consensus from the National Lipid Association',
    'Soffer DE, Marston NA, Maki KC, Jacobson TA, Bittner VA, Peña JM, Thanassoulis G, Martin SS, Kirkpatrick CF, Virani SS, Dixon DL, Ballantyne CM, Remaley AT',
    'Journal of Clinical Lipidology',
    '39256087',
    'review',
    '2024-09-05'::date,
    'Consenso da National Lipid Association estabelecendo que apolipoproteína B (apoB) é medida clínica validada que complementa painel lipídico padrão. ApoB quantifica diretamente número de partículas aterogênicas (cada partícula LDL, VLDL, IDL contém uma molécula apoB), sendo preditor superior ao LDL-C isolado especialmente em pacientes com diabetes, síndrome metabólica, triglicerídeos elevados ou em uso de estatinas. Recomenda dosagem de apoB para estratificação de risco mais precisa e ajuste terapêutico em pacientes de risco intermediário a alto. Metas sugeridas: <80 mg/dL (risco moderado), <70 mg/dL (alto risco), <55 mg/dL (muito alto risco). Estudos de discordância demonstram que quando LDL-C, colesterol não-HDL e apoB não estão alinhados, risco cardiovascular segue mais proximamente apoB e não-HDL, refletindo melhor potencial aterogênico real.',
    '10.1016/j.jacl.2024.08.013',
    'en'
) ON CONFLICT (pm_id) DO NOTHING;

-- Article 2: Present and Future of Lipid Testing (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    article_type,
    publish_date,
    abstract,
    doi,
    language
) VALUES (
    gen_random_uuid(),
    'The Present and Future of Lipid Testing in Cardiovascular Risk Assessment',
    'White-Al Habeeb NMA, Higgins V, Wolska A, Delaney SR, Remaley AT, Beriault DR',
    'Clinical Chemistry',
    '37000150',
    'review',
    '2023-04-28'::date,
    'Revisão abrangente em Clinical Chemistry sobre métodos atuais e emergentes de avaliação lipídica. Discute diretrizes nacionais para cálculo e medição de LDL-C (equação de Friedewald, Martin-Hopkins, medição direta), vantagens do colesterol não-HDL (não requer jejum, inclui todas partículas aterogênicas), e papel crescente de biomarcadores avançados (apoB, Lp(a), partículas LDL pequenas e densas). Enfatiza que perfil lipídico pode ser realizado sem jejum para rastreamento inicial na maioria dos pacientes, facilitando adesão ao screening. Apresenta algoritmos de decisão clínica baseados em metas lipídicas específicas por categoria de risco ASCVD. Projeta futuro da avaliação lipídica incluindo análise de subfrações lipoproteicas por ressonância magnética nuclear, medição rotineira de Lp(a), e integração com escores poligênicos de risco para medicina de precisão.',
    '10.1093/clinchem/hvad012',
    'en'
) ON CONFLICT (pm_id) DO NOTHING;

-- Article 3: Serum Cholesterol and CVD Death Meta-Analysis (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    article_type,
    publish_date,
    abstract,
    doi,
    language
) VALUES (
    gen_random_uuid(),
    'Serum Cholesterol Levels and Risk of Cardiovascular Death: A Systematic Review and a Dose-Response Meta-Analysis of Prospective Cohort Studies',
    'Jung E, Kong SY, Ro YS, Ryu HH, Shin SD',
    'International Journal of Environmental Research and Public Health',
    '35886124',
    'meta_analysis',
    '2022-07-06'::date,
    'Meta-análise dose-resposta incluindo 14 estudos de coorte prospectivos com 1.055.309 participantes avaliando associação entre níveis de colesterol sérico e mortalidade cardiovascular. Demonstrou hazard ratio pooled de 1,27 para colesterol total elevado, 1,21 para LDL-C elevado (risco 21% maior de morte cardiovascular), e 0,60 para HDL-C elevado (efeito protetor com 40% redução de risco). Análise dose-resposta revelou relação linear entre LDL-C e mortalidade CV sem threshold inferior evidente, suportando conceito lower is better. Níveis de HDL-C mostraram relação inversa consistente. Triglicerídeos elevados associaram-se independentemente com risco aumentado mesmo após ajuste para outros lipídios. Resultados reforçam importância de avaliação completa do perfil lipídico além de LDL-C isolado, incluindo colesterol não-HDL e relação CT/HDL para estratificação de risco mais acurada.',
    '10.3390/ijerph19148272',
    'en'
) ON CONFLICT (pm_id) DO NOTHING;

-- Article 4: PEER Simplified Lipid Guideline 2023
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    article_type,
    publish_date,
    abstract,
    doi,
    language
) VALUES (
    gen_random_uuid(),
    'PEER simplified lipid guideline 2023 update: Prevention and management of cardiovascular disease in primary care',
    'Kolber MR, Klarenbach S, Cauchon M, Cotterill M, Regier L, Marceau RD, Duggan N, Whitley R, Halme AS, Poshtar T, Allan GM, Korownyk CS, Ton J, Froentjes L, Moe SS, Perry D, Thomas BS, McCormack JP, Falk J, Dugré N, Garrison SR, Kirkwood JEM, Young J, Braschi É, Paige A, Potter J, Weresch J, Lindblad AJ',
    'Canadian Family Physician',
    '37833089',
    'review',
    '2023-10-01'::date,
    'Diretriz canadense PEER 2023 fornecendo algoritmo simplificado para manejo de dislipidemia em cuidados primários. Recomenda rastreamento lipídico universal em adultos ≥40 anos e em adultos mais jovens com fatores de risco. Perfil lipídico pode ser não-jejum exceto se triglicerídeos prévios >4,5 mmol/L (400 mg/dL). Estratifica risco usando calculadora Framingham ou ASCVD e identifica condições de muito alto risco (DCV estabelecida, diabetes ≥15 anos, DRC estágio 3+). Metas: LDL <2,0 mmol/L (77 mg/dL) ou redução ≥50% para alto risco; <1,8 mmol/L (70 mg/dL) para muito alto risco. Enfatiza abordagem não-HDL como alvo secundário (meta 30 mg/dL acima de LDL). Estatina de intensidade moderada a alta como terapia inicial; adicionar ezetimiba se meta não atingida. Reforça importância de modificações estilo de vida (dieta mediterrânea, exercício 150 min/semana, cessação tabagismo) como base do tratamento. Intervalo de seguimento: 3 meses após início terapia, depois anualmente.',
    '10.46747/cfp.6910675',
    'en'
) ON CONFLICT (pm_id) DO NOTHING;

-- ============================================================================
-- STEP 3: LINK ARTICLES TO SCORE ITEM
-- ============================================================================

-- Link all articles to the lipid panel score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT
    a.id,
    'd17cea8f-2935-40bc-b2e0-bcd9ac886ade'::uuid
FROM articles a
WHERE a.pm_id IN ('39256087', '37000150', '35886124', '37833089')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- STEP 4: VERIFICATION
-- ============================================================================

-- Verify all articles were inserted and linked
DO $$
DECLARE
    v_article_count INTEGER;
    v_link_count INTEGER;
    v_item_name TEXT;
BEGIN
    -- Check article insertion
    SELECT COUNT(*) INTO v_article_count
    FROM articles
    WHERE pm_id IN ('39256087', '37000150', '35886124', '37833089');

    -- Check article-item links
    SELECT COUNT(*) INTO v_link_count
    FROM article_score_items
    WHERE score_item_id = 'd17cea8f-2935-40bc-b2e0-bcd9ac886ade';

    -- Get item name
    SELECT name INTO v_item_name
    FROM score_items
    WHERE id = 'd17cea8f-2935-40bc-b2e0-bcd9ac886ade';

    RAISE NOTICE '================================================';
    RAISE NOTICE 'ENRICHMENT VERIFICATION REPORT';
    RAISE NOTICE '================================================';
    RAISE NOTICE 'Score Item: %', v_item_name;
    RAISE NOTICE 'Score Item ID: d17cea8f-2935-40bc-b2e0-bcd9ac886ade';
    RAISE NOTICE 'Articles inserted: %/4', v_article_count;
    RAISE NOTICE 'Article links created: %', v_link_count;
    RAISE NOTICE '================================================';

    IF v_article_count < 4 THEN
        RAISE WARNING 'Expected 4 articles, found only %', v_article_count;
    END IF;

    IF v_link_count < 4 THEN
        RAISE WARNING 'Expected 4 article links, found only %', v_link_count;
    END IF;

    IF v_article_count = 4 AND v_link_count >= 4 THEN
        RAISE NOTICE '✓ Enrichment completed successfully!';
    END IF;
END $$;

-- Show final linked articles
SELECT
    si.name AS score_item_name,
    a.title,
    a.pm_id,
    a.article_type,
    a.publish_date,
    LEFT(a.abstract, 100) || '...' AS abstract_preview
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'd17cea8f-2935-40bc-b2e0-bcd9ac886ade'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- END OF ENRICHMENT
-- ============================================================================
