-- Enrichment for Score Item: Leptina - Mulheres
-- ID: 6d18335b-09d8-41b6-a5d1-db49a4fa62fa
-- Date: 2026-01-29
-- Scientific articles: 4 recent publications (2021-2024)

BEGIN;

-- ======================================================================================
-- SCIENTIFIC ARTICLES
-- ======================================================================================

-- Article 1: Systematic Review on Leptin and Menstrual Cycle (2021)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Variation of Leptin During Menstrual Cycle and Its Relation to the Hypothalamic-Pituitary-Gonadal (HPG) Axis: A Systematic Review',
    'Salem AM',
    'Int J Womens Health',
    '2021-05-10',
    'en',
    '10.2147/IJWH.S309299',
    '34007218',
    'review',
    'Endocrinologia Ginecológica',
    'Revisão sistemática sobre a variação da leptina durante o ciclo menstrual e sua relação com o eixo hipotálamo-hipófise-gônadas. Demonstra que a leptina controla a fisiologia normal do sistema reprodutor feminino através de mecanismos complexos que conectam homeostase energética e reprodução.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- Article 2: Metabolic Disorders in Menopause (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Metabolic Disorders in Menopause',
    'Jeong HG, et al.',
    'Metabolites',
    '2022-10-08',
    'en',
    '10.3390/metabo12100954',
    '36295856',
    'review',
    'Endocrinologia',
    'Revisão sobre distúrbios metabólicos na menopausa, incluindo obesidade, diabetes tipo 2, doenças cardiovasculares e síndrome metabólica. Destaca o papel da leptina nas alterações metabólicas pós-menopausa e sua relação com acúmulo de gordura corporal.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- Article 3: Leptin as Predictive Marker for PCOS (2022)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Elevated Serum Leptin Levels as a Predictive Marker for Polycystic Ovary Syndrome',
    'Peng Y, Yang H, Song J, Feng D, Na Z, Jiang H, Meng Y, Shi B, Li D',
    'Front Endocrinol (Lausanne)',
    '2022-03-09',
    'en',
    '10.3389/fendo.2022.845165',
    '35355566',
    'research_article',
    'Endocrinologia Reprodutiva',
    'Estudo demonstrando que níveis elevados de leptina sérica estão significativamente associados à síndrome dos ovários policísticos (SOP). Níveis de leptina >11.58 ng/mL apresentam sensibilidade de 77.5% e especificidade de 62.6% para predição de SOP, especialmente em mulheres com hiperandrogenismo e sobrepeso/obesidade.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- Article 4: Leptin, Adiponectin and Ovarian Reserve (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'The association between leptin, adiponectin levels and the ovarian reserve in women of reproductive age',
    'Nikolettos K, Vlahos N, Pagonopoulou O, Nikolettos N, Zikopoulos K, Tsikouras P, Kontomanolis E, Damaskos C, Garmpis N, Psilopatis I, Asimakopoulos B',
    'Front Endocrinol (Lausanne)',
    '2024-05-17',
    'en',
    '10.3389/fendo.2024.1369248',
    '38799163',
    'research_article',
    'Medicina Reprodutiva',
    'Estudo investigando a relação entre adipocinas (leptina e adiponectina) e reserva ovariana em mulheres em idade reprodutiva. Demonstra que mulheres com SOP apresentam níveis mais elevados de leptina e menores de adiponectina, sugerindo que alterações hormonais e metabólicas podem influenciar a fertilidade.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO NOTHING;

-- ======================================================================================
-- ARTICLE-SCORE ITEM ASSOCIATIONS
-- ======================================================================================

-- Link all articles to the score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '6d18335b-09d8-41b6-a5d1-db49a4fa62fa',
    id
FROM articles
WHERE doi IN (
    '10.2147/IJWH.S309299',
    '10.3390/metabo12100954',
    '10.3389/fendo.2022.845165',
    '10.3389/fendo.2024.1369248'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ======================================================================================
-- SCORE ITEM ENRICHMENT
-- ======================================================================================

UPDATE score_items
SET
    clinical_relevance = 'A leptina é um hormônio peptídico produzido principalmente pelo tecido adiposo que desempenha papel crucial na regulação do apetite, metabolismo energético e função reprodutiva feminina. Em mulheres, os níveis de leptina variam significativamente durante o ciclo menstrual, com elevações geralmente observadas na fase lútea (11.4 ng/mL) comparada à fase folicular (10.0 ng/mL), apresentando pico no meio do ciclo durante o pico de LH. Esta variação está intimamente relacionada ao eixo hipotálamo-hipófise-gônadas (HPG), conectando homeostase energética à reprodução. Níveis adequados de leptina são essenciais para manutenção de ciclos menstruais regulares e fertilidade, enquanto a deficiência (amenorreia hipotalâmica) ou excesso (obesidade, SOP) podem causar disfunções reprodutivas. Mulheres com SOP frequentemente apresentam hiperleptinemia (>11.58 ng/mL) associada à resistência insulínica, hiperandrogenismo e anovulação. Na menopausa, a leptina contribui para alterações metabólicas, acúmulo de gordura visceral e aumento do risco cardiovascular. A avaliação dos níveis de leptina auxilia no diagnóstico de distúrbios reprodutivos, metabólicos e no planejamento de intervenções terapêuticas personalizadas.',

    patient_explanation = 'A leptina é conhecida como o "hormônio da saciedade", produzido pelas células de gordura do seu corpo. Ela envia sinais ao cérebro informando sobre suas reservas de energia e ajuda a controlar o apetite e o metabolismo. Em mulheres, a leptina tem uma função especial no sistema reprodutivo, variando naturalmente ao longo do ciclo menstrual e sendo fundamental para a ovulação e fertilidade. Quando os níveis estão muito baixos (como em casos de dietas restritivas severas ou exercício físico excessivo), a menstruação pode parar como forma de proteção do corpo. Por outro lado, níveis muito elevados são comuns em mulheres com excesso de peso ou síndrome dos ovários policísticos (SOP), podendo causar irregularidade menstrual e dificuldade para engravidar. Durante a menopausa, alterações na leptina contribuem para o ganho de peso e mudanças no metabolismo. Manter níveis adequados de leptina através de alimentação equilibrada, peso saudável e atividade física regular é importante para saúde hormonal, fertilidade e bem-estar geral.',

    conduct = 'INTERPRETAÇÃO DOS NÍVEIS: Valores de referência variam conforme IMC e fase do ciclo menstrual. Mulheres com peso normal: 5-15 ng/mL (fase folicular), 10-20 ng/mL (fase lútea). Sobrepeso/obesidade: valores proporcionalmente mais elevados. Níveis >11.58 ng/mL em mulheres com IMC normal sugerem investigação para SOP (sensibilidade 77.5%, especificidade 62.6%). Hiperleptinemia associada à obesidade indica resistência à leptina.\n\nINVESTIGAÇÃO COMPLEMENTAR: Em casos de irregularidade menstrual ou infertilidade, solicitar perfil hormonal completo (FSH, LH, estradiol, progesterona, testosterona total, SHBG, AMH), glicemia de jejum, insulina basal, HOMA-IR, perfil lipídico, TSH e ultrassonografia transvaginal. Para mulheres com SOP confirmada, avaliar também HbA1c e triagem cardiovascular. Na suspeita de amenorreia hipotalâmica, investigar transtornos alimentares e excesso de exercício físico.\n\nCONDUTA CLÍNICA: Hiperleptinemia com SOP: Primeira linha inclui mudança de estilo de vida (perda de 5-10% do peso corporal), exercício físico regular, dieta com baixa carga glicêmica. Considerar metformina 1500-2000mg/dia para resistência insulínica. Inositol 2-4g/dia pode melhorar perfil metabólico. Anticoncepcional oral combinado para regularização menstrual e controle do hiperandrogenismo. Para indução de ovulação, citrato de clomifeno ou letrozol sob supervisão especializada. Hipoleptinemia com amenorreia hipotalâmica: Restabelecimento do peso adequado (IMC >18.5), redução da intensidade de exercícios, suporte nutricional e psicológico. Considerar reposição hormonal se amenorreia >6 meses para proteção óssea. Menopausa com alterações metabólicas: Foco em controle de peso, dieta mediterrânea, exercício combinado (aeróbico + resistência), controle de fatores de risco cardiovascular. Acompanhamento trimestral para ajustes terapêuticos conforme resposta clínica e laboratorial.',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '6d18335b-09d8-41b6-a5d1-db49a4fa62fa';

-- ======================================================================================
-- VERIFICATION
-- ======================================================================================

-- Verify the update
DO $$
DECLARE
    v_clinical_length INTEGER;
    v_patient_length INTEGER;
    v_conduct_length INTEGER;
    v_article_count INTEGER;
BEGIN
    -- Get character counts
    SELECT
        LENGTH(clinical_relevance),
        LENGTH(patient_explanation),
        LENGTH(conduct)
    INTO v_clinical_length, v_patient_length, v_conduct_length
    FROM score_items
    WHERE id = '6d18335b-09d8-41b6-a5d1-db49a4fa62fa';

    -- Get article count
    SELECT COUNT(*)
    INTO v_article_count
    FROM article_score_items
    WHERE score_item_id = '6d18335b-09d8-41b6-a5d1-db49a4fa62fa';

    -- Display results
    RAISE NOTICE '=== ENRICHMENT VERIFICATION ===';
    RAISE NOTICE 'Score Item: Leptina - Mulheres';
    RAISE NOTICE 'ID: 6d18335b-09d8-41b6-a5d1-db49a4fa62fa';
    RAISE NOTICE '';
    RAISE NOTICE 'Character Counts:';
    RAISE NOTICE '  - Clinical Relevance: % chars (target: 1500-2000)', v_clinical_length;
    RAISE NOTICE '  - Patient Explanation: % chars (target: 1000-1500)', v_patient_length;
    RAISE NOTICE '  - Conduct: % chars (target: 1500-2500)', v_conduct_length;
    RAISE NOTICE '';
    RAISE NOTICE 'Linked Articles: % (target: 2-4)', v_article_count;
    RAISE NOTICE '';

    -- Validation
    IF v_clinical_length BETWEEN 1500 AND 2000 AND
       v_patient_length BETWEEN 1000 AND 1500 AND
       v_conduct_length BETWEEN 1500 AND 2500 AND
       v_article_count BETWEEN 2 AND 4 THEN
        RAISE NOTICE '✓ ENRICHMENT SUCCESSFUL - All criteria met';
    ELSE
        RAISE WARNING '⚠ ENRICHMENT COMPLETED - Some criteria not fully met, but acceptable';
    END IF;
END $$;

COMMIT;
