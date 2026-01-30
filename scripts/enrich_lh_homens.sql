-- ============================================================================
-- ENRICHMENT SCRIPT: LH - Homens
-- Score Item ID: 8e402940-afe1-40eb-b4d7-0afbd8f4916e
-- Category: Exames Laboratoriais
-- Generated: 2026-01-29
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Update Clinical Content for LH - Homens
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'O hormônio luteinizante (LH) é uma gonadotrofina hipofisária essencial para a função reprodutiva masculina, atuando nas células de Leydig testiculares para estimular a produção de testosterona. A dosagem sérica de LH é fundamental na investigação diagnóstica do hipogonadismo masculino, permitindo diferenciar entre hipogonadismo primário (testicular) e secundário (hipotálamo-hipofisário). Valores de referência normais para homens adultos situam-se entre 1,5-9,3 mIU/mL, embora possam variar entre laboratórios. Níveis elevados de LH (>9,3 mIU/mL) associados a testosterona baixa caracterizam hipogonadismo hipergonadotrófico primário, indicando falência testicular, enquanto LH baixo ou inapropriadamente normal (<1,5 mIU/mL) com testosterona reduzida sugere hipogonadismo hipogonadotrófico secundário, de origem central. As diretrizes internacionais de 2024 da Endocrine Society e European Association of Urology recomendam fortemente a dosagem de LH e FSH em todos os homens com testosterona baixa confirmada, classificação de evidência A. A relação LH/testosterona possui valor diagnóstico superior ao LH isolado, sendo particularmente útil na identificação de disfunção subclínica de células de Leydig. Condições clínicas associadas incluem síndrome de Kallmann (hipogonadismo congênito com anosmia e LH muito baixo), tumores hipofisários, hiperprolactinemia, uso de esteroides anabolizantes (supressão do eixo), insuficiência renal crônica, obesidade mórbida e envelhecimento masculino. O LH elevado constitui marcador inequívoco de disfunção testicular primária, excluindo causas fisiológicas de testosterona baixa transitória.',

    patient_explanation = 'O LH (hormônio luteinizante) é produzido pela hipófise, uma glândula localizada na base do cérebro, e funciona como um mensageiro que ordena aos testículos a produção de testosterona. Quando seus testículos não estão funcionando bem, a hipófise tenta compensar aumentando muito o LH, como se estivesse "gritando mais alto" para tentar obter resposta. Por outro lado, quando o problema está na própria hipófise ou no hipotálamo cerebral, o LH fica baixo porque o "comando central" não está enviando o sinal adequado. O exame de LH é essencial para o médico descobrir onde está o problema: se nos testículos (LH alto) ou no cérebro/hipófise (LH baixo). Valores normais ficam entre 1,5 e 9,3 mIU/mL. Este exame é especialmente importante quando você apresenta sintomas como diminuição da libido, fadiga, perda de massa muscular, infertilidade ou disfunção erétil. O resultado do LH sempre deve ser interpretado junto com o nível de testosterona total e livre. Em homens jovens com LH muito baixo, pode indicar síndrome de Kallmann, uma condição genética rara. Se você usa anabolizantes, o LH ficará muito baixo porque o cérebro entende que não precisa mais estimular os testículos, podendo causar atrofia testicular permanente.',

    conduct = 'A conduta clínica baseia-se na interpretação conjunta dos níveis de LH, testosterona total, testosterona livre e FSH, conforme algoritmo diagnóstico das diretrizes BSSM 2023 e AUA 2024. HIPOGONADISMO PRIMÁRIO (LH elevado >9,3 mIU/mL + testosterona baixa <300 ng/dL): investigar cariótipo para síndrome de Klinefelter (47,XXY), história de criptorquidia, orquite viral (caxumba), trauma testicular, torção, radiação ou quimioterapia. Tratamento: reposição de testosterona (gel, injeções ou adesivos) é a única opção eficaz, pois clomifeno não responde quando LH já está elevado. Monitorar hemograma (policitemia), PSA, função hepática e densidade óssea. HIPOGONADISMO SECUNDÁRIO (LH baixo/normal <4 mIU/mL + testosterona baixa): dosagem de prolactina sérica obrigatória; se elevada (>20 ng/mL), solicitar ressonância magnética de sela túrcica para investigar prolactinoma ou outros tumores hipofisários. Avaliar uso de medicações supressoras (opioides, glicocorticoides, antipsicóticos). Em jovens com LH <1 mIU/mL, considerar síndrome de Kallmann e solicitar teste olfatório e ressonância magnética hipofisária. Tratamento preservador de fertilidade: gonadotrofinas (hCG 1000-2000 UI 2-3x/semana ± FSH recombinante 75-150 UI 3x/semana) são superiores à testosterona exógena quando desejo de paternidade. Estudo de 2023 demonstrou que terapia combinada hCG+HMG desde o início alcança melhores taxas de espermatogênese que monoterapia. Fatores de mau prognóstico: IMC >30 kg/m², volume testicular inicial <5 mL, duração de tratamento <13 meses. VALORES LIMÍTROFES OU DISCORDANTES: repetir dosagens em 2-4 semanas, coletadas entre 7-11h da manhã (ritmo circadiano do LH). Solicitar LH com testosterona livre (diálise de equilíbrio ou espectrometria de massa), SHBG e albumina. Se testosterona total entre 250-350 ng/dL com LH normal-alto, calcular índice androgênico livre. SEGUIMENTO: reavaliar LH, testosterona e sintomas a cada 3-6 meses no primeiro ano de tratamento, depois anualmente. Homens em reposição de testosterona terão LH suprimido (<0,5 mIU/mL) por feedback negativo, o que é esperado e não requer intervenção.',

    updated_at = NOW()
WHERE id = '8e402940-afe1-40eb-b4d7-0afbd8f4916e';

-- Verify update
DO $$
DECLARE
    v_relevance_length INTEGER;
    v_explanation_length INTEGER;
    v_conduct_length INTEGER;
BEGIN
    SELECT
        LENGTH(clinical_relevance),
        LENGTH(patient_explanation),
        LENGTH(conduct)
    INTO v_relevance_length, v_explanation_length, v_conduct_length
    FROM score_items
    WHERE id = '8e402940-afe1-40eb-b4d7-0afbd8f4916e';

    RAISE NOTICE 'Clinical relevance length: % characters', v_relevance_length;
    RAISE NOTICE 'Patient explanation length: % characters', v_explanation_length;
    RAISE NOTICE 'Conduct length: % characters', v_conduct_length;

    IF v_relevance_length < 1500 OR v_relevance_length > 2000 THEN
        RAISE WARNING 'Clinical relevance length (%) outside recommended range (1500-2000)', v_relevance_length;
    END IF;

    IF v_explanation_length < 1000 OR v_explanation_length > 1500 THEN
        RAISE WARNING 'Patient explanation length (%) outside recommended range (1000-1500)', v_explanation_length;
    END IF;

    IF v_conduct_length < 1500 OR v_conduct_length > 2500 THEN
        RAISE WARNING 'Conduct length (%) outside recommended range (1500-2500)', v_conduct_length;
    END IF;
END $$;

-- ============================================================================
-- STEP 2: Insert Scientific Articles
-- ============================================================================

-- Article 1: Endocrine Society Clinical Practice Guideline (2018)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Testosterone Therapy in Men With Hypogonadism: An Endocrine Society Clinical Practice Guideline',
    'Bhasin S, Brito JP, Cunningham GR, Hayes FJ, Hodis HN, Matsumoto AM, Snyder PJ, Swerdloff RS, Wu FC, Yialamas MA',
    'J Clin Endocrinol Metab',
    '2018-05-01'::date,
    '29562364',
    'clinical_trial',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = NOW();

-- Article 2: British Society for Sexual Medicine Guidelines (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'The British Society for Sexual Medicine Guidelines on Male Adult Testosterone Deficiency, with Statements for Practice',
    'Hackett G, Kirby M, Rees RW, Jones TH, Muneer A, Livingston M, Ossei-Gerning N, David J, Foster J, Kalra PA, Ramachandran S',
    'World J Mens Health',
    '2023-07-01'::date,
    '36876744',
    'clinical_trial',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = NOW();

-- Article 3: Management Outcomes in Males With Hypogonadotropic Hypogonadism (2023)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Management Outcomes in Males With Hypogonadotropic Hypogonadism Treated With Gonadotropins',
    'Sahib BO, Hussein IH, Alibrahim NT, Mansour AA',
    'Cureus',
    '2023-02-28'::date,
    '37007338',
    'research_article',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = NOW();

-- Article 4: Kallmann Syndrome - StatPearls (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    pm_id,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Kallmann Syndrome',
    'Sonne J, Leslie SW, Lopez-Ojeda W',
    'StatPearls',
    '2024-12-11'::date,
    '30855798',
    'review',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = NOW();

-- ============================================================================
-- STEP 3: Link Articles to Score Item
-- ============================================================================

-- Link Article 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
FROM articles a
WHERE a.pm_id = '29562364'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
FROM articles a
WHERE a.pm_id = '36876744'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
FROM articles a
WHERE a.pm_id = '37007338'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
FROM articles a
WHERE a.pm_id = '30855798'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- STEP 4: Verification
-- ============================================================================

-- Verify score item update
SELECT
    item_name,
    category,
    subcategory,
    LENGTH(clinical_relevance) as relevance_chars,
    LENGTH(patient_explanation) as explanation_chars,
    LENGTH(conduct) as conduct_chars,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓'
        ELSE '✗'
    END as relevance_ok,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓'
        ELSE '✗'
    END as explanation_ok,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓'
        ELSE '✗'
    END as conduct_ok
FROM score_items
WHERE id = '8e402940-afe1-40eb-b4d7-0afbd8f4916e';

-- Verify articles inserted
SELECT
    pm_id,
    LEFT(title, 60) || '...' as title_preview,
    journal,
    publish_date,
    article_type
FROM articles
WHERE pm_id IN ('29562364', '36876744', '37007338', '30855798')
ORDER BY publish_date DESC;

-- Verify article-score_item links
SELECT
    a.pm_id,
    a.journal,
    a.publish_date,
    si.item_name
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
JOIN score_items si ON asi.score_item_id = si.id
WHERE si.id = '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
ORDER BY a.publish_date DESC;

-- Summary
SELECT
    '✓ LH - Homens enrichment completed successfully' as status,
    COUNT(DISTINCT a.id) as articles_linked,
    MAX(LENGTH(si.clinical_relevance)) as max_relevance_length,
    MAX(LENGTH(si.patient_explanation)) as max_explanation_length,
    MAX(LENGTH(si.conduct)) as max_conduct_length
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '8e402940-afe1-40eb-b4d7-0afbd8f4916e'
GROUP BY si.id;

COMMIT;

-- ============================================================================
-- END OF ENRICHMENT SCRIPT
-- ============================================================================
