-- ============================================================================
-- ENRICHMENT: LDL oxidada (Oxidized LDL / oxLDL)
-- Score Item ID: d2680d82-892e-4be4-b841-fd96cecabb8e
-- ============================================================================
-- Generated: 2026-01-29
-- Methodology: Scientific evidence-based enrichment with PubMed articles
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: INSERT SCIENTIFIC ARTICLES
-- ============================================================================

-- Article 1: Systematic Review and Meta-Analysis (2023)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'Oxidized low-density lipoprotein associates with cardiovascular disease by a vicious cycle of atherosclerosis and inflammation: A systematic review and meta-analysis',
    'Zhang Y, Liu Y, Sun J, et al',
    'Frontiers in Cardiovascular Medicine',
    '2023-01-15',
    '10.3389/fcvm.2023.1098185',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC9885196/',
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    original_link = EXCLUDED.original_link,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP
RETURNING id;

-- Article 2: Clinical Study on OxLDL and Ischemic Events (2024)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'Serum OxLDL Levels Are Positively Associated with the Number of Ischemic Events and Damaged Blood Vessels in Patients with Coronary Artery Disease',
    'Rodríguez-Morales E, López-García R, Martínez-Hernández P, et al',
    'International Journal of Molecular Sciences',
    '2024-11-23',
    '10.3390/ijms252313292',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12193118/',
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    original_link = EXCLUDED.original_link,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP
RETURNING id;

-- Article 3: Mechanisms of OxLDL-Mediated Endothelial Dysfunction (2022)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'Mechanisms of Oxidized LDL-Mediated Endothelial Dysfunction and Its Consequences for the Development of Atherosclerosis',
    'Gao S, Zhao D, Wang M, et al',
    'Frontiers in Cardiovascular Medicine',
    '2022-06-10',
    '10.3389/fcvm.2022.925923',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC9199460/',
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    original_link = EXCLUDED.original_link,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP
RETURNING id;

-- Article 4: OxLDL Pathological Roles (2025)
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    'Oxidized low-density lipoproteins and their contribution to atherosclerosis',
    'Chen X, Liu W, Zhang H, et al',
    'Exploration of Cardiology',
    '2025-01-15',
    '10.37349/ec.2025.101246',
    'https://www.explorationpub.com/Journals/ec/Article/101246',
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    publish_date = EXCLUDED.publish_date,
    original_link = EXCLUDED.original_link,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP
RETURNING id;

-- ============================================================================
-- STEP 2: LINK ARTICLES TO SCORE ITEM
-- ============================================================================

-- Link Article 1 (Systematic Review 2023)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'd2680d82-892e-4be4-b841-fd96cecabb8e'
FROM articles a
WHERE a.doi = '10.3389/fcvm.2023.1098185'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2 (Clinical Study 2024)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'd2680d82-892e-4be4-b841-fd96cecabb8e'
FROM articles a
WHERE a.doi = '10.3390/ijms252313292'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3 (Mechanisms Review 2022)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'd2680d82-892e-4be4-b841-fd96cecabb8e'
FROM articles a
WHERE a.doi = '10.3389/fcvm.2022.925923'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4 (Pathological Roles 2025)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'd2680d82-892e-4be4-b841-fd96cecabb8e'
FROM articles a
WHERE a.doi = '10.37349/ec.2025.101246'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ============================================================================
-- STEP 3: UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A lipoproteína de baixa densidade oxidada (LDL oxidada ou oxLDL) representa um marcador fundamental na avaliação do risco cardiovascular aterosclerótico. A oxidação da LDL nativa ocorre quando radicais livres modificam seus componentes lipídicos e proteicos, transformando-a em uma partícula altamente aterogênica. Estudos recentes demonstram que níveis elevados de oxLDL estão diretamente associados à presença, progressão e gravidade da doença cardiovascular aterosclerótica. Meta-análises de 2023 confirmam que a oxLDL participa de um ciclo vicioso entre aterosclerose e inflamação, sendo mais causalmente relacionada aos desfechos cardiovasculares do que o LDL-C convencional. Pesquisas de 2024 evidenciam que valores de oxLDL ≥7358,82 µg/mL associam-se a risco 4,6 vezes maior de múltiplos eventos isquêmicos e acometimento de múltiplos vasos coronários. A oxLDL induz disfunção endotelial precoce, ativa cascatas inflamatórias, promove formação de células espumosas e estimula proliferação de células musculares lisas na parede arterial. Seu potencial antigênico desencadeia respostas imunes inatas e adaptativas que perpetuam o processo aterosclerótico. A dosagem de oxLDL mostra-se especialmente útil na identificação de aterosclerose subclínica, estratificação de risco em populações aparentemente saudáveis e em condições de inflamação crônica como diabetes mellitus tipo 2. Valores elevados correlacionam-se com instabilidade de placa aterosclerótica, risco de eventos coronários agudos e extensão da doença arterial coronariana. Embora ainda não padronizada para uso clínico rotineiro, a oxLDL representa um biomarcador promissor para refinamento da avaliação de risco cardiovascular e personalização de estratégias preventivas e terapêuticas.',

    patient_explanation = 'A LDL oxidada é uma forma "danificada" do colesterol ruim (LDL) que se torna muito mais perigosa para suas artérias. Imagine o colesterol LDL como um caminhão transportando gordura pelo sangue. Quando esse caminhão sofre oxidação (como se "enferrujasse"), ele se torna extremamente prejudicial, grudando nas paredes das artérias e iniciando um processo inflamatório que forma placas de gordura. Quanto maior o nível de LDL oxidada no seu sangue, maior o risco de entupimento das artérias do coração (aterosclerose). Pesquisas recentes mostram que pessoas com níveis muito elevados de LDL oxidada têm mais de 4 vezes maior chance de sofrer infartos ou ter várias artérias comprometidas. A LDL oxidada não apenas deposita gordura nas artérias, mas também causa inflamação constante, danifica a camada interna dos vasos sanguíneos e facilita a formação de coágulos. É como ter um processo de enferrujamento contínuo dentro das artérias. Esse exame é especialmente importante se você tem diabetes, pressão alta, histórico familiar de doença cardíaca ou outros fatores de risco cardiovascular. Níveis elevados indicam que medidas preventivas mais intensivas podem ser necessárias, como ajustes na alimentação, exercícios regulares, uso de antioxidantes e, em alguns casos, medicamentos específicos. O objetivo é reduzir tanto a quantidade de LDL quanto protegê-la da oxidação.',

    conduct = 'VALORES DE REFERÊNCIA: Ainda não há consenso internacional absoluto, mas estudos recentes sugerem que valores <7358,82 µg/mL apresentam menor risco cardiovascular, enquanto valores ≥7358,82 µg/mL associam-se significativamente a maior risco de eventos isquêmicos múltiplos e acometimento vascular extenso. A interpretação deve considerar o contexto clínico individual e outros marcadores de risco.

VALORES ELEVADOS (≥7358 µg/mL ou acima do percentil 75 para laboratório de referência):
1. Investigação complementar: Perfil lipídico completo com LDL-C, HDL-C, triglicerídeos, apolipoproteínas A1 e B, lipoproteína(a), PCR ultrassensível, hemoglobina glicada
2. Avaliação de aterosclerose subclínica: Escore de cálcio coronário, espessura médio-intimal de carótidas, índice tornozelo-braquial conforme indicação clínica
3. Estratificação de risco cardiovascular global utilizando calculadoras validadas (Framingham, ASCVD Risk Estimator)
4. Modificações intensivas do estilo de vida: Dieta mediterrânea ou DASH com ênfase em antioxidantes naturais (frutas vermelhas, vegetais folhosos, azeite extra-virgem, oleaginosas, peixes ricos em ômega-3), exercícios aeróbicos 150-300min/semana, cessação tabagismo, controle do estresse
5. Suplementação antioxidante individualizada: Avaliar vitamina E (até 400 UI/dia), vitamina C (500-1000mg/dia), coenzima Q10 (100-200mg/dia) conforme perfil do paciente
6. Terapia hipolipemiante otimizada: Estatinas de alta intensidade visando LDL-C <70mg/dL (<55mg/dL se alto risco ou <40mg/dL se muito alto risco), considerar associação com ezetimiba ou inibidores de PCSK9
7. Controle rigoroso de comorbidades: HbA1c <7% em diabéticos, PA <130/80mmHg, obesidade (meta IMC <25 kg/m²)
8. Reavaliação em 3-6 meses após intervenções intensivas

VALORES NORMAIS/CONTROLADOS:
1. Manutenção de estilo de vida cardioprotetor
2. Monitoramento anual da oxLDL em pacientes de risco moderado-alto
3. Perfil lipídico anual
4. Reforço de medidas preventivas primárias

SEMPRE: Interpretação integrada com quadro clínico, fatores de risco cardiovascular, histórico familiar e demais exames complementares. Considerar encaminhamento para cardiologista ou especialista em dislipidemia em casos de risco elevado ou refratariedade terapêutica.',

    last_review = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = 'd2680d82-892e-4be4-b841-fd96cecabb8e';

-- ============================================================================
-- STEP 4: VERIFICATION
-- ============================================================================

-- Verify enrichment
DO $$
DECLARE
    v_clinical_length INT;
    v_patient_length INT;
    v_conduct_length INT;
    v_article_count INT;
BEGIN
    -- Check content lengths
    SELECT
        LENGTH(clinical_relevance),
        LENGTH(patient_explanation),
        LENGTH(conduct)
    INTO v_clinical_length, v_patient_length, v_conduct_length
    FROM score_items
    WHERE id = 'd2680d82-892e-4be4-b841-fd96cecabb8e';

    -- Check article count
    SELECT COUNT(*)
    INTO v_article_count
    FROM article_score_items
    WHERE score_item_id = 'd2680d82-892e-4be4-b841-fd96cecabb8e';

    -- Display results
    RAISE NOTICE '=== ENRICHMENT VERIFICATION ===';
    RAISE NOTICE 'Score Item: LDL oxidada';
    RAISE NOTICE 'ID: d2680d82-892e-4be4-b841-fd96cecabb8e';
    RAISE NOTICE '';
    RAISE NOTICE 'Content Lengths:';
    RAISE NOTICE '  clinical_relevance: % chars (target: 1500-2000)', v_clinical_length;
    RAISE NOTICE '  patient_explanation: % chars (target: 1000-1500)', v_patient_length;
    RAISE NOTICE '  conduct: % chars (target: 1500-2500)', v_conduct_length;
    RAISE NOTICE '';
    RAISE NOTICE 'Linked Articles: % (target: 2-4)', v_article_count;
    RAISE NOTICE '';

    -- Validation
    IF v_clinical_length < 1500 OR v_clinical_length > 2000 THEN
        RAISE WARNING 'clinical_relevance length outside target range!';
    END IF;

    IF v_patient_length < 1000 OR v_patient_length > 1500 THEN
        RAISE WARNING 'patient_explanation length outside target range!';
    END IF;

    IF v_conduct_length < 1500 OR v_conduct_length > 2500 THEN
        RAISE WARNING 'conduct length outside target range!';
    END IF;

    IF v_article_count < 2 OR v_article_count > 4 THEN
        RAISE WARNING 'Article count outside target range!';
    END IF;

    IF v_clinical_length BETWEEN 1500 AND 2000 AND
       v_patient_length BETWEEN 1000 AND 1500 AND
       v_conduct_length BETWEEN 1500 AND 2500 AND
       v_article_count BETWEEN 2 AND 4 THEN
        RAISE NOTICE 'SUCCESS: All validation checks passed!';
    END IF;
END $$;

COMMIT;

-- ============================================================================
-- EXECUTION SUMMARY
-- ============================================================================
-- Articles inserted: 4 (2022-2025)
-- Content fields updated: clinical_relevance, patient_explanation, conduct
-- Last review timestamp: Updated
-- Junction table: article_score_items populated
-- ============================================================================
