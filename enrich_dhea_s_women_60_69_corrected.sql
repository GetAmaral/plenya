-- ====================================================================
-- ENRICHMENT: DHEA-S - Mulheres (60-69 anos)
-- Score Item ID: 22952700-63e3-422d-b4b2-179c0785a20e
-- ====================================================================
-- Generated: 2026-01-28
-- This script enriches the DHEA-S score item for women aged 60-69 with:
-- - 4 peer-reviewed scientific articles (2010-2025)
-- - Comprehensive clinical content in Portuguese
-- ====================================================================

BEGIN;

-- ====================================================================
-- STEP 1: INSERT SCIENTIFIC ARTICLES
-- ====================================================================

-- Article 1: DHEA and Bone Health Review (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Dehydroepiandrosterone and Bone Health: Mechanisms and Insights',
    'Mohamad NV, Che Razali NS, Mohd Shamsuddin NA',
    'Biomedicines',
    '2024-12-01',
    '10.3390/biomedicines12122780',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC11673555/',
    'review',
    'en',
    'Endocrinologia',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Article 2: DHEA Meta-Analysis on Hormones (2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects',
    'He SY, Lu K, Zhang L, Cao H, Tang X, Zhang X',
    'Diabetology & Metabolic Syndrome',
    '2025-07-01',
    '10.1186/s13098-025-01770-0',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12231631/',
    'meta_analysis',
    'en',
    'Endocrinologia',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Article 3: DHEA and Osteoporosis/Sarcopenia (2021)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Serum concentrations of oxytocin, DHEA and follistatin are associated with osteoporosis or sarcopenia in community-dwelling postmenopausal women',
    'Du Y, Xu C, Shi H, et al',
    'BMC Geriatrics',
    '2021-10-12',
    '10.1186/s12877-021-02481-7',
    'https://bmcgeriatr.biomedcentral.com/articles/10.1186/s12877-021-02481-7',
    'research_article',
    'en',
    'Geriatria',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- Article 4: DHEA and Cardiovascular Risk/Frailty RCT (2010)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    original_link,
    article_type,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Effects of dehydroepiandrosterone (DHEA) on cardiovascular risk factors in older women with frailty characteristics',
    'Boxer RS, Kleppinger A, Brindisi J, Feinn R, Burleson JA, Kenny AM',
    'Age and Ageing',
    '2010-07-01',
    '10.1093/ageing/afq043',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC2899943/',
    'clinical_trial',
    'en',
    'Geriatria',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO NOTHING;

-- ====================================================================
-- STEP 2: LINK ARTICLES TO SCORE ITEM
-- ====================================================================

INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '22952700-63e3-422d-b4b2-179c0785a20e'::uuid,
    a.id
FROM articles a
WHERE a.doi IN (
    '10.3390/biomedicines12122780',
    '10.1186/s13098-025-01770-0',
    '10.1186/s12877-021-02481-7',
    '10.1093/ageing/afq043'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ====================================================================
-- STEP 3: UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ====================================================================

UPDATE score_items
SET
    clinical_relevance = 'O DHEA-S (sulfato de dehidroepiandrosterona) é um andrógeno adrenal fundamental para mulheres na faixa etária de 60-69 anos, particularmente após a menopausa. Seus níveis declinam significativamente com o envelhecimento, sendo este declínio associado a múltiplas condições geriátricas. Em mulheres idosas, o DHEA-S serve como precursor hormonal essencial, convertendo-se em estrogênios e testosterona perifericamente, compensando parcialmente a deficiência hormonal pós-menopausa. Evidências demonstram correlação robusta entre níveis adequados de DHEA-S e melhor força muscular (β=0.19, p=0.041), função muscular (β=0.58, p=0.004) e massa óssea (β=0.51, p=0.017), sendo fundamental na prevenção de sarcopenia e osteoporose. Mulheres idosas com sarcopenia apresentam níveis significativamente menores de DHEA-S comparadas às sem a condição. Na saúde óssea, estudos demonstram que a suplementação com 50mg/dia aumenta densidade mineral óssea no rádio ultradistal e coluna lombar, mediada pela elevação de estradiol. A avaliação de DHEA-S é particularmente relevante em mulheres com fragilidade, perda de massa muscular, osteoporose ou risco aumentado de fraturas. Apesar das evidências promissoras para saúde músculo-esquelética, estudos não demonstraram benefícios cardiovasculares consistentes, tornando sua indicação mais específica para condições músculo-esqueléticas em mulheres desta faixa etária.',

    patient_explanation = 'O DHEA-S é um hormônio produzido principalmente pelas suas glândulas suprarrenais que funciona como uma "matéria-prima" para produzir outros hormônios importantes, especialmente após a menopausa quando seus ovários produzem menos hormônios. Na faixa dos 60-69 anos, os níveis de DHEA-S naturalmente diminuem bastante, e isso pode afetar vários aspectos da sua saúde. Níveis adequados deste hormônio estão associados a músculos mais fortes, melhor equilíbrio e ossos mais resistentes. Quando o DHEA-S está baixo, você pode ter maior risco de desenvolver fragilidade óssea (osteoporose), perder massa muscular (sarcopenia) e sentir mais fraqueza geral. Seu corpo usa o DHEA-S para fabricar pequenas quantidades de estrogênio e testosterona, que são importantes para manter a força dos ossos e músculos mesmo depois da menopausa. Em alguns casos, quando os níveis estão muito baixos e há problemas de saúde relacionados, seu médico pode considerar reposição hormonal, geralmente com doses de 50mg por dia. Porém, é importante entender que a suplementação deve ser cuidadosamente avaliada, pois embora possa ajudar ossos e músculos, não demonstrou benefícios claros para o coração. O exame de DHEA-S é especialmente importante se você apresenta perda de força muscular, osteoporose ou histórico de fraturas.',

    conduct = 'INTERPRETAÇÃO DOS RESULTADOS: Em mulheres de 60-69 anos, valores de referência variam entre laboratórios, mas geralmente situam-se entre 12-154 µg/dL. Valores abaixo de 30 µg/dL são considerados deficiência significativa nesta faixa etária e requerem investigação clínica. A interpretação deve considerar idade específica, pois o declínio é progressivo e valores "normais-baixos" podem ser clinicamente relevantes se acompanhados de sintomas.

AVALIAÇÃO COMPLEMENTAR: Solicitar simultaneamente estradiol, testosterona total e livre, SHBG para avaliar conversão periférica. Densitometria óssea (DEXA) é mandatória quando DHEA-S está baixo para avaliação de osteoporose/osteopenia. Avaliar força muscular através de dinamometria de preensão manual e testes funcionais (velocidade de marcha, teste de sentar-levantar) para detectar sarcopenia. Investigar função adrenal se níveis muito baixos (<15 µg/dL) para excluir insuficiência adrenal. Avaliar critérios de fragilidade (Fried Frailty Phenotype) incluindo perda de peso não intencional, fadiga, fraqueza, baixa velocidade de marcha.

INDICAÇÕES DE SUPLEMENTAÇÃO: Meta-análises demonstram que doses de 50mg/dia aumentam significativamente testosterona (WMD: 24.31 ng/dL) e estradiol (WMD: 7.86 pg/mL), com efeitos mais pronunciados em mulheres ≥60 anos. Duração mínima recomendada é 26 semanas para otimizar benefícios. Indicações estabelecidas incluem: insuficiência adrenal documentada, osteoporose/osteopenia com DHEA-S baixo, sarcopenia confirmada com deficiência hormonal, uso crônico de glicocorticoides. Considerar tratamento em mulheres com fragilidade e DHEA-S <30 µg/dL, sempre associado a exercícios resistidos e suplementação de cálcio (1000-1200mg/dia) e vitamina D3 (1000UI/dia).

MONITORAMENTO E PRECAUÇÕES: Dosar DHEA-S, estradiol e função hepática a cada 3 meses nos primeiros 6 meses, depois semestralmente. Atenção especial ao estradiol elevado que pode estimular endométrio; realizar ultrassom transvaginal se sangramento. Contraindicações incluem câncer hormônio-dependente ativo, doença hepática grave. Efeitos colaterais possíveis: acne, hirsutismo leve, voz grave. Evidências não suportam uso para prevenção cardiovascular. Abordagem deve ser multimodal: exercícios resistidos 2-3x/semana, nutrição adequada proteica (1.2-1.5g/kg/dia), vitamina D3 e cálcio são fundamentais independente de reposição hormonal.',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '22952700-63e3-422d-b4b2-179c0785a20e';

COMMIT;

-- ====================================================================
-- VERIFICATION QUERIES
-- ====================================================================

-- Check clinical content character counts
SELECT
    'clinical_relevance' as field,
    LENGTH(clinical_relevance) as char_count,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '22952700-63e3-422d-b4b2-179c0785a20e'

UNION ALL

SELECT
    'patient_explanation' as field,
    LENGTH(patient_explanation) as char_count,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '22952700-63e3-422d-b4b2-179c0785a20e'

UNION ALL

SELECT
    'conduct' as field,
    LENGTH(conduct) as char_count,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
        ELSE '✗ OUT OF RANGE'
    END as status
FROM score_items
WHERE id = '22952700-63e3-422d-b4b2-179c0785a20e';

-- Check linked articles count
SELECT
    si.id,
    si.name,
    COUNT(asi.article_id) as articles_count,
    CASE
        WHEN COUNT(asi.article_id) >= 2 THEN '✓ OK'
        ELSE '✗ INSUFFICIENT'
    END as status
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '22952700-63e3-422d-b4b2-179c0785a20e'
GROUP BY si.id, si.name;

-- List all linked articles with details
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.article_type,
    a.doi
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '22952700-63e3-422d-b4b2-179c0785a20e'
ORDER BY a.publish_date DESC;
