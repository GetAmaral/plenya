-- =====================================================
-- ENRICHMENT: Transaminase pirúvica (ALT/TGP)
-- Score Item ID: 06241683-bc19-4d56-b9a8-dade736e674f
-- =====================================================
-- Generated: 2026-01-29
-- Focus: Alanine aminotransferase (ALT) as liver biomarker
-- Recent guidelines: EASL-EASD-EASO 2024, updated reference intervals
-- =====================================================

BEGIN;

-- =====================================================
-- 1. UPDATE CLINICAL CONTENT
-- =====================================================

UPDATE score_items
SET
    clinical_relevance = 'A alanina aminotransferase (ALT ou TGP) é uma enzima encontrada predominantemente nos hepatócitos, sendo o biomarcador mais sensível e específico para lesão hepatocelular. Valores elevados indicam perda da integridade da membrana celular ou necrose hepática, com liberação da enzima no plasma. A ALT é componente essencial da avaliação hepática inicial, incluída nos painéis de função hepática e perfis metabólicos abrangentes. Estudos de 2024 estabeleceram novos intervalos de referência mais restritivos: limite superior de 34 U/L para homens e 22 U/L para mulheres em população metabolicamente saudável, significativamente menores que valores convencionais. A relação AST/ALT fornece informações diagnósticas importantes: razão <1 sugere hepatite viral ou esteatose hepática não alcoólica (MASLD), enquanto razão ≥2 indica doença hepática alcoólica, e razão >1 em contexto não alcoólico sugere cirrose. Elevações de ALT >7x o limite superior têm sensibilidade e especificidade >95% para lesão hepática aguda. A presença de ALT elevada associa-se a aumento da mortalidade relacionada ao fígado. As diretrizes EASL-EASD-EASO 2024 recomendam busca ativa de MASLD em indivíduos com enzimas hepáticas anormais, especialmente na presença de diabetes tipo 2 ou obesidade, utilizando abordagem escalonada com escores não invasivos (FIB-4) e elastografia para avaliação de fibrose avançada.',

    patient_explanation = 'A ALT (ou TGP) é uma enzima que fica dentro das células do fígado. Quando o fígado está saudável, apenas pequenas quantidades dessa enzima circulam no sangue (geralmente <30 U/L). Se as células do fígado sofrem lesão ou inflamação, a ALT "vaza" para a corrente sanguínea, fazendo os níveis subirem. Valores normais atualizados são até 34 U/L para homens e até 22 U/L para mulheres - limites mais baixos que os antigos. Níveis levemente elevados (até 3x o normal) podem indicar esteatose hepática (gordura no fígado), muito comum em pessoas com sobrepeso, diabetes ou colesterol alto. Elevações moderadas a altas sugerem hepatites virais (B ou C), uso excessivo de álcool, medicamentos hepatotóxicos ou doenças autoimunes. Valores muito altos (>10x o normal) indicam lesão aguda grave. A ALT é comparada com outra enzima chamada AST: quando ALT está mais alta que AST, geralmente indica problemas relacionados ao fígado propriamente; quando AST está mais alta, pode ser doença alcoólica ou cirrose. É importante monitorar a ALT regularmente se você tem fatores de risco como obesidade, diabetes, síndrome metabólica ou usa medicamentos que afetam o fígado. Valores alterados requerem investigação complementar com ultrassom, outros exames de sangue e, às vezes, elastografia hepática para avaliar fibrose.',

    conduct = 'INVESTIGAÇÃO INICIAL: Confirmar elevação com nova dosagem em 2-4 semanas, avaliar histórico medicamentoso (estatinas, paracetamol, antibióticos, antifúngicos), consumo de álcool (questionário AUDIT), uso de suplementos e fitoterápicos. Solicitar painel hepático completo: AST, fosfatase alcalina, GGT, bilirrubinas totais e frações, albumina, TAP/INR. Calcular relação AST/ALT. ESTRATIFICAÇÃO POR NÍVEL: ALT <2x LSN com AST/ALT <1: investigar MASLD (esteatose metabólica) - solicitar ultrassom abdominal, perfil lipídico completo, glicemia de jejum, HbA1c, HOMA-IR, calcular FIB-4 para avaliar fibrose. ALT 2-5x LSN: adicionar sorologias virais (anti-HCV, HBsAg, anti-HBc total), ferritina e saturação de transferrina (hemocromatose), ceruloplasmina (doença de Wilson se <40 anos), FAN e anticorpos anti-músculo liso (hepatite autoimune). ALT >5x LSN: considerar lesão aguda - investigar hepatites virais agudas (IgM anti-HAV, IgM anti-HBc), toxicidade medicamentosa, hepatite isquêmica, obstrução biliar aguda. MASLD CONFIRMADA: Estratificar risco de fibrose com FIB-4; se FIB-4 intermediário ou alto, solicitar elastografia hepática (FibroScan). Iniciar modificação de estilo de vida: redução de peso 7-10%, dieta mediterrânea, exercícios aeróbicos 150min/semana, controle rigoroso de diabetes e dislipidemia. Evitar álcool. SEGUIMENTO: ALT normal com fatores de risco metabólicos: repetir anualmente. ALT persistentemente elevada sem fibrose: reavaliar a cada 6 meses com FIB-4. Fibrose significativa (F2-F3): acompanhamento hepatológico semestral, considerar biópsia se diagnóstico incerto, avaliar candidatura para terapias emergentes. Cirrose estabelecida: rastreamento semestral de hepatocarcinoma (alfafetoproteína + ultrassom), endoscopia digestiva alta para varizes esofágicas. Referir ao hepatologista se: ALT >10x LSN, icterícia, coagulopatia, ascite, encefalopatia, suspeita de hepatite autoimune ou doença de Wilson.',

    last_review = CURRENT_DATE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '06241683-bc19-4d56-b9a8-dade736e674f';

-- =====================================================
-- 2. INSERT SCIENTIFIC ARTICLES
-- =====================================================

-- Article 1: Updated Reference Intervals for ALT (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Updated Reference Intervals for Alanine Aminotransferase in a Metabolically and Histologically Normal Population',
    'Choi J, Jo C, Kim S, Ryu M, Choi WM, Lee D, Shim JH, Kim KM, Lim YS, Lee HC',
    'Clin Gastroenterol Hepatol',
    '38750867',
    '10.1016/j.cgh.2024.04.031',
    '2024-11-01'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 1 to Score Item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '06241683-bc19-4d56-b9a8-dade736e674f'
FROM articles a
WHERE a.doi = '10.1016/j.cgh.2024.04.031'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 2: EASL-EASD-EASO MASLD Guidelines (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'EASL-EASD-EASO Clinical Practice Guidelines on the management of metabolic dysfunction-associated steatotic liver disease (MASLD)',
    'Tacke F, Horn P, Wong VWS, Ratziu V, Bugianesi E, Francque S, Zelber-Sagi S, Valenti L, Roden M, Schick F, Yki-Järvinen H, Gastaldelli A, Vettor R, Frühbeck G, Dicker D',
    'J Hepatol',
    '38851997',
    '10.1016/j.jhep.2024.04.031',
    '2024-09-01'::date,
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 2 to Score Item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '06241683-bc19-4d56-b9a8-dade736e674f'
FROM articles a
WHERE a.doi = '10.1016/j.jhep.2024.04.031'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 3: AST/ALT Ratio in Chronic HBV (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'AST to ALT ratio as a prospective risk predictor for liver cirrhosis in patients with chronic HBV infection',
    'Lai X, Chen H, Dong X, Zhou G, Liang D, Xu F, Liu H, Luo Y, Liu H, Wan S',
    'Eur J Gastroenterol Hepatol',
    '38251454',
    '10.1097/MEG.0000000000002708',
    '2024-03-01'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 3 to Score Item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '06241683-bc19-4d56-b9a8-dade736e674f'
FROM articles a
WHERE a.doi = '10.1097/MEG.0000000000002708'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 4: ALT/AST Ratio in NAFLD (2024)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    pm_id,
    doi,
    publish_date,
    article_type,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Elevated ALT/AST ratio as a marker for NAFLD risk and severity: insights from a cross-sectional analysis in the United States',
    'Xuan Y, Wu D, Zhang Q, Yu Z, Yu J, Zhou D',
    'Front Endocrinol (Lausanne)',
    '39253584',
    '10.3389/fendo.2024.1457598',
    '2024-08-26'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
)
ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 4 to Score Item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '06241683-bc19-4d56-b9a8-dade736e674f'
FROM articles a
WHERE a.doi = '10.3389/fendo.2024.1457598'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- =====================================================
-- 3. VERIFICATION
-- =====================================================

-- Verify character counts
SELECT
    'ALT/TGP Enrichment' as item_name,
    LENGTH(clinical_relevance) as clinical_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    CASE
        WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓'
        ELSE '✗'
    END as clinical_ok,
    CASE
        WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓'
        ELSE '✗'
    END as patient_ok,
    CASE
        WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓'
        ELSE '✗'
    END as conduct_ok,
    last_review
FROM score_items
WHERE id = '06241683-bc19-4d56-b9a8-dade736e674f';

-- Verify linked articles
SELECT
    a.pm_id,
    a.title,
    a.journal,
    a.publish_date,
    a.article_type
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '06241683-bc19-4d56-b9a8-dade736e674f'
ORDER BY a.publish_date DESC;

COMMIT;

-- =====================================================
-- ENRICHMENT SUMMARY
-- =====================================================
-- Item: Transaminase pirúvica (ALT/TGP)
-- Articles Added: 4
-- - PMID 38750867: Updated ALT reference intervals (2024)
-- - PMID 38851997: EASL-EASD-EASO MASLD guidelines (2024)
-- - PMID 38251454: AST/ALT ratio in chronic HBV (2024)
-- - PMID 39253584: ALT/AST ratio in NAFLD (2024)
-- Focus: Liver enzyme biomarker, MASLD/NASH diagnostics
-- Clinical updates: 2024 reference ranges, fibrosis assessment
-- =====================================================
