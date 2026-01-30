-- ================================================================
-- ENRICHMENT: DHEA-S - Mulheres (40-49 anos)
-- Score Item ID: 5745d2f7-7b96-4329-9692-6d7cd8bb92c6
-- Generated: 2026-01-28
-- ================================================================

BEGIN;

-- ================================================================
-- STEP 1: Insert Articles
-- ================================================================

-- Article 1: Meta-analysis on DHEA supplementation (2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Impact of DHEA supplementation on testosterone and estradiol levels in postmenopausal women: a meta-analysis of randomized controlled trials assessing dose and duration effects',
    'ShuYun He, Kunna Lu, Lianying Zhang, Hanwen Cao, Xinyuan Tang, Xinhuan Zhang',
    'Diabetology & Metabolic Syndrome',
    '2025-07-04',
    '10.1186/s13098-025-01770-0',
    'meta_analysis',
    'Meta-análise de 21 ensaios clínicos randomizados (n=986-1084) avaliando efeitos da suplementação de DHEA em mulheres pós-menopáusicas. DHEA ≥50 mg/dia aumentou significativamente testosterona (29,65 ng/dL) e estradiol (8,65 pg/mL). Mulheres ≥60 anos apresentaram elevação mais pronunciada de estradiol (8,92 pg/mL). Recomenda-se monitoramento para efeitos adversos andrógenos e estimulação endometrial.',
    'en',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id;

-- Store article 1 ID for linking
DO $$
DECLARE
    v_article_1_id uuid;
BEGIN
    SELECT id INTO v_article_1_id
    FROM articles
    WHERE doi = '10.1186/s13098-025-01770-0';

    -- Store in temporary table for later use
    CREATE TEMP TABLE IF NOT EXISTS temp_article_ids (article_id uuid, article_num int);
    DELETE FROM temp_article_ids WHERE article_num = 1;
    INSERT INTO temp_article_ids VALUES (v_article_1_id, 1);
END $$;

-- Article 2: Clinical review on DHEA administration in women (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Should Dehydroepiandrosterone Be Administered to Women?',
    'Margaret E Wierman, Katja Kiseljak-Vassiliades',
    'The Journal of Clinical Endocrinology and Metabolism',
    '2022-03-07',
    '10.1210/clinem/dgac130',
    'review',
    'Revisão crítica sobre administração de DHEA em mulheres. Evidências suportam pequenos benefícios na qualidade de vida e humor em mulheres com insuficiência adrenal primária/secundária ou anorexia, mas não para ansiedade ou função sexual. O efeito ósseo de DHEA é significativamente menor que terapia estrogênica ou medicações aprovadas para osteoporose. DHEA vaginal local mostra promessa para atrofia vulvovaginal.',
    'en',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id;

-- Store article 2 ID
DO $$
DECLARE
    v_article_2_id uuid;
BEGIN
    SELECT id INTO v_article_2_id
    FROM articles
    WHERE doi = '10.1210/clinem/dgac130';

    DELETE FROM temp_article_ids WHERE article_num = 2;
    INSERT INTO temp_article_ids VALUES (v_article_2_id, 2);
END $$;

-- Article 3: Androgen evaluation guideline (2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Society for Endocrinology Clinical Practice Guideline for the Evaluation of Androgen Excess in Women',
    'Yasir S Elhassan, Wiebke Arlt, Miles J Levy',
    'Clinical Endocrinology',
    '2025-01-15',
    '10.1111/cen.15265',
    'review',
    'Diretriz clínica da Society for Endocrinology para avaliação de excesso androgênico em mulheres. DHEA-S sérico é a medida mais confiável de produção adrenal de andrógenos, particularmente útil na transição perimenopáusica. Aproximadamente 25% da produção androgênica ocorre nas adrenais em mulheres pré-menopáusicas, aumentando para 100% na pós-menopausa. Declínio de DHEA-S inicia aos 30 anos devido à involução da zona reticular adrenal.',
    'en',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id;

-- Store article 3 ID
DO $$
DECLARE
    v_article_3_id uuid;
BEGIN
    SELECT id INTO v_article_3_id
    FROM articles
    WHERE doi = '10.1111/cen.15265';

    DELETE FROM temp_article_ids WHERE article_num = 3;
    INSERT INTO temp_article_ids VALUES (v_article_3_id, 3);
END $$;

-- Article 4: DHEA as sexual hormone precursor (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    abstract,
    language,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'The Utilization of Dehydroepiandrosterone as a Sexual Hormone Precursor in Premenopausal and Postmenopausal Women: An Overview',
    'Elena Tsourdi, Elizabeth Curtis, Athanasios D Anastasilakis, Lorenz C Hofbauer, Martina Rauner',
    'Endocrine Reviews',
    '2022-02-18',
    '10.1210/endrev/bnab039',
    'review',
    'Revisão sobre DHEA como precursor hormonal sexual em mulheres pré e pós-menopáusicas. DHEA-S é produzido quase exclusivamente nas adrenais e representa 75% dos andrógenos em mulheres pré-menopáusicas e 100% na pós-menopausa. Mulheres perimenopáusicas têm apenas 50% dos níveis de pico de DHEA. DHEA intravaginal (Prasterona) é o único produto aprovado pela FDA para dispareunia moderada a severa na síndrome geniturinária da menopausa.',
    'en',
    'Endocrinology',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW()
RETURNING id;

-- Store article 4 ID
DO $$
DECLARE
    v_article_4_id uuid;
BEGIN
    SELECT id INTO v_article_4_id
    FROM articles
    WHERE doi = '10.1210/endrev/bnab039';

    DELETE FROM temp_article_ids WHERE article_num = 4;
    INSERT INTO temp_article_ids VALUES (v_article_4_id, 4);
END $$;

-- ================================================================
-- STEP 2: Update Score Item with Clinical Content
-- ================================================================

UPDATE score_items
SET
    clinical_relevance = 'DHEA-S (sulfato de deidroepiandrosterona) é o andrógeno adrenal mais abundante e marcador mais confiável da produção adrenal de andrógenos em mulheres entre 40-49 anos, período crítico da transição perimenopáusica. Nesta faixa etária, mulheres apresentam apenas 50% dos níveis de pico de DHEA observados aos 20-30 anos devido à involução progressiva da zona reticular adrenal. Durante a perimenopausa, aproximadamente 25% dos andrógenos provêm das adrenais, 25% dos ovários e 50% de conversão periférica; na pós-menopausa, as adrenais tornam-se responsáveis por 100% da produção androgênica. Níveis adequados de DHEA-S são fundamentais para manutenção da função sexual (libido, lubrificação vaginal), densidade óssea, massa muscular, humor e energia. A redução de DHEA-S nesta década correlaciona-se com aumento de queixas de fadiga, disfunção sexual, ressecamento vaginal, perda de massa óssea e declínio cognitivo. DHEA-S serve como precursor para conversão periférica em testosterona e estradiol, hormônios essenciais para saúde metabólica, cardiovascular e neurocognitiva. Valores de referência para mulheres de 40 anos situam-se entre 160-240 μg/dL, com declínio gradual ao longo da década. Avaliação de DHEA-S é particularmente relevante em mulheres com sintomas de insuficiência androgênica, síndrome geniturinária da menopausa, osteopenia/osteoporose precoce ou fadiga inexplicada. Monitoramento longitudinal permite identificação precoce de insuficiência adrenal ou declínio acelerado associado a condições patológicas.',

    patient_explanation = 'DHEA-S é um hormônio produzido pelas glândulas adrenais (pequenas glândulas localizadas acima dos rins) que funciona como matéria-prima para produção de outros hormônios importantes no seu corpo, especialmente testosterona e estrogênio. Entre os 40-49 anos, período em que muitas mulheres começam a entrar na perimenopausa (transição para a menopausa), os níveis de DHEA-S naturalmente diminuem para aproximadamente metade do que eram aos 20-30 anos. Este hormônio é fundamental para manter sua energia, disposição, saúde dos ossos, massa muscular, libido (desejo sexual), lubrificação vaginal e humor equilibrado. Quando os níveis de DHEA-S estão baixos, você pode sentir cansaço constante, diminuição do interesse sexual, ressecamento vaginal, dificuldade de concentração, irritabilidade ou perda de força muscular. Valores normais para mulheres aos 40 anos ficam entre 160-240 μg/dL, diminuindo gradualmente com o passar dos anos. Este exame é especialmente importante se você está apresentando sintomas de menopausa precoce, fadiga crônica, perda de libido ou osteoporose. Em alguns casos específicos, quando os níveis estão muito baixos e causando sintomas significativos, seu médico pode considerar suplementação hormonal, embora esta decisão deva ser individualizada e cuidadosamente avaliada quanto a riscos e benefícios.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: Valores de referência para mulheres 40-49 anos: 160-240 μg/dL aos 40 anos, com redução gradual de 10-15% por década. Níveis <100 μg/dL sugerem insuficiência adrenal ou declínio acelerado; >400 μg/dL podem indicar tumor adrenal, síndrome de Cushing ou uso de DHEA exógeno. AVALIAÇÃO INICIAL: Coletar DHEA-S juntamente com cortisol matinal, testosterona total e livre, SHBG, estradiol, FSH e LH para avaliação hormonal completa da perimenopausa. Considerar teste de estimulação com ACTH se DHEA-S <50 μg/dL para descartar insuficiência adrenal. Avaliar sintomas clínicos: fadiga (escala de fadiga), função sexual (FSFI - Female Sexual Function Index), sintomas climatéricos (índice de Kupperman), densidade óssea (DEXA se osteopenia suspeitada). CONDUTA NÍVEIS BAIXOS: Investigar causas secundárias (insuficiência adrenal, hipopituitarismo, uso de corticoides, doenças crônicas). Considerar suplementação de DHEA 25-50 mg/dia apenas em casos selecionados: insuficiência adrenal documentada, sintomas significativos de deficiência androgênica refratários a outras terapias, ou osteoporose com baixa resposta a tratamento convencional. Monitorar resposta clínica após 3 meses: melhora de fadiga, libido, densidade óssea (DEXA anual). Vigilância para efeitos adversos: acne, hirsutismo, alteração de voz, irregularidade menstrual em perimenopáusicas. DHEA vaginal (Prasterona 6,5 mg/dia) é alternativa aprovada para atrofia vulvovaginal e dispareunia moderada-severa. SEGUIMENTO: Reavaliar DHEA-S a cada 6-12 meses em mulheres sintomáticas ou em suplementação. Titular dose baseado em níveis séricos (alvo: terço superior da faixa normal) e resposta clínica. Descontinuar se sem benefício após 6 meses ou se efeitos andrógenos significativos. CONTRAINDICAÇÕES à suplementação: câncer de mama, câncer endometrial, sangramento vaginal inexplicado, doença hepática ativa.',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '5745d2f7-7b96-4329-9692-6d7cd8bb92c6';

-- ================================================================
-- STEP 3: Link Articles to Score Item
-- ================================================================

-- Link all articles
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '5745d2f7-7b96-4329-9692-6d7cd8bb92c6'::uuid,
    article_id
FROM temp_article_ids
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ================================================================
-- STEP 4: Verification Query
-- ================================================================

-- Display enrichment results
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(DISTINCT asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '5745d2f7-7b96-4329-9692-6d7cd8bb92c6'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- Show linked article titles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.article_type
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
WHERE asi.score_item_id = '5745d2f7-7b96-4329-9692-6d7cd8bb92c6'
ORDER BY a.publish_date DESC;

-- Clean up temporary table
DROP TABLE IF EXISTS temp_article_ids;

COMMIT;

-- ================================================================
-- END OF ENRICHMENT
-- ================================================================
