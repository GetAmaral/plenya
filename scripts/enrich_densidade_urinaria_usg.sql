-- ============================================================================
-- ENRICHMENT: Densidade Urinária (USG)
-- Score Item ID: a03fa6c4-da56-4969-86b7-57a11e7d6907
-- Generated: 2026-01-28
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: Diabetes Insipidus Updates (Frontiers in Endocrinology, 2025)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    doi,
    article_type,
    language,
    abstract,
    notes,
    specialty,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Central and nephrogenic diabetes insipidus: updates on diagnosis and management',
    'Flynn K, Hatfield J, Brown K',
    'Frontiers in Endocrinology',
    '2025-01-15',
    '10.3389/fendo.2024.1479764',
    'review',
    'en',
    'Comprehensive review on central and nephrogenic diabetes insipidus with updates on diagnostic methods and therapeutic approaches. Highlights the role of copeptin as an emerging marker for diagnosis and differentiation between central and nephrogenic forms.',
    'KEY FINDINGS: (1) Water deprivation test remains gold standard - healthy kidneys concentrate urine >800 mOsm/kg while DI patients maintain hypotonic urine <300 mOsm/kg. (2) After desmopressin, central DI shows >50% increase in urine osmolality while nephrogenic DI shows minimal response (<15%). (3) Baseline copeptin >21.4 pmol/L diagnoses nephrogenic DI with 100% sensitivity and specificity. (4) Markedly low urine specific gravity (USG <1.008) indicates hyposthenuria characteristic of diabetes insipidus.',
    'Nephrology, Endocrinology',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING id AS article_1_id;

-- Store article_1_id for later use
DO $$
DECLARE
    v_article_1_id uuid;
    v_article_2_id uuid;
    v_article_3_id uuid;
    v_article_4_id uuid;
BEGIN
    -- Article 1 already inserted above
    SELECT id INTO v_article_1_id FROM articles WHERE doi = '10.3389/fendo.2024.1479764';

    -- Article 2: Hydration Criterion Values (European Journal of Clinical Nutrition, 2017)
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        language,
        abstract,
        notes,
        specialty,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Criterion values for urine-specific gravity and urine color representing adequate water intake in healthy adults',
        'Perrier ET, Bottin JH, Vecchio M, Lemetais G',
        'European Journal of Clinical Nutrition',
        '2017-02-01',
        '10.1038/ejcn.2016.269',
        'research_article',
        'en',
        'Prospective study with 817 urine samples from 82 healthy French adults establishing cutoff values for urine specific gravity and color as indicators of adequate water intake. Demonstrated high accuracy of USG ≥1.013 for identifying concentrated urine.',
        'KEY FINDINGS: (1) USG ≥1.013 identifies concentrated urine (osmolality >500 mOsm/kg) with very high accuracy (AUC 0.984). (2) Maintaining 24-hour urine osmolality ≤500 mOsm/kg is desirable target for renal health and adequate hydration. (3) Urine specific gravity demonstrates strong sensitivity and specificity as practical tool for hydration monitoring. (4) Urine color ≥4 (8-point scale) provides complementary field measure though with moderate reliability.',
        'Nephrology, Sports Medicine',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ) RETURNING id INTO v_article_2_id;

    -- Article 3: Adult Dehydration (StatPearls, 2025)
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        language,
        abstract,
        notes,
        specialty,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Adult Dehydration',
        'Taylor K, Tripathi AK',
        'StatPearls Publishing',
        '2025-03-05',
        'NBK555956',
        'review',
        'en',
        'Reference chapter from StatPearls on adult dehydration, addressing pathophysiology, laboratory diagnosis, and clinical management. Emphasizes limitations of urinary biomarkers when used alone, especially in elderly patients.',
        'KEY FINDINGS: (1) Urine specific gravity >1.020 and osmolality >450 mOsm/kg indicate concentrated urine reflecting renal compensation for fluid loss. (2) Serum osmolality >295 mOsm/kg serves as reasonable threshold for dehydration from water loss. (3) 2015 Cochrane review demonstrated that USG, urine osmolality and other biomarkers are unreliable diagnostic tools when used alone in elderly. (4) Simple clinical signs like skin turgor, dry mouth or urine color should NOT be used to assess hydration in elderly.',
        'Nephrology, Internal Medicine',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ) RETURNING id INTO v_article_3_id;

    -- Article 4: Hydration in Athletes (PMC, 2016)
    INSERT INTO articles (
        id,
        title,
        authors,
        journal,
        publish_date,
        doi,
        article_type,
        language,
        abstract,
        notes,
        specialty,
        created_at,
        updated_at
    ) VALUES (
        gen_random_uuid(),
        'Validity of Urine Specific Gravity when Compared to Plasma Osmolality as a Measure of Hydration Status in Male and Female NCAA Collegiate Athletes',
        'Munoz CX, Johnson EC, McKenzie AL, Guelinckx I, Graverholt G, Casa DJ',
        'Journal of Strength and Conditioning Research',
        '2016-06-01',
        '10.1519/JSC.0000000000001298',
        'research_article',
        'en',
        'Validation study comparing urine specific gravity with plasma osmolality as measures of hydration status in collegiate athletes. Demonstrated that USG is a sensitive, specific, and practical tool for pre-exercise hydration assessment.',
        'KEY FINDINGS: (1) USG demonstrated to be sensitive and specific indicator of hydration in athletes before exercise. (2) Sports medicine defines dehydration as concentrated urine with osmolality ≥700 mOsm/kg or USG ≥1.020. (3) USG is inexpensive, simple, fast, and accurate method to assess hydration status. (4) Digital refractometers offer more precise measurement than reagent strips for USG assessment.',
        'Sports Medicine, Nephrology',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ) RETURNING id INTO v_article_4_id;

    -- Link all articles to the score item
    INSERT INTO article_score_items (score_item_id, article_id)
    VALUES
        ('a03fa6c4-da56-4969-86b7-57a11e7d6907', v_article_1_id),
        ('a03fa6c4-da56-4969-86b7-57a11e7d6907', v_article_2_id),
        ('a03fa6c4-da56-4969-86b7-57a11e7d6907', v_article_3_id),
        ('a03fa6c4-da56-4969-86b7-57a11e7d6907', v_article_4_id);

END $$;

-- ============================================================================
-- STEP 2: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A densidade urinária específica (USG) é um biomarcador essencial que reflete a capacidade de concentração renal e o estado de hidratação do organismo. Medida através de refratometria ou fitas reagentes, a USG normal varia entre 1.003 e 1.030, representando a razão entre a densidade da urina e da água pura. Este parâmetro fornece informações valiosas sobre função tubular renal, balanço hídrico e distúrbios hidroeletrolíticos. Valores elevados (>1.020) indicam urina concentrada, sugerindo desidratação, perda de fluidos por vômitos/diarréia, ou compensação renal adequada. Valores baixos (<1.010) podem indicar hiper-hidratação, diabetes insípido, ou comprometimento da capacidade de concentração renal. Estudos demonstram que USG ≥1.013 identifica urina concentrada (osmolalidade >500 mOsm/kg) com acurácia de 98.4%, sendo ferramenta sensível e específica para avaliação de hidratação. No diabetes insípido, observa-se hipostenúria marcada (USG <1.008) devido à incapacidade de concentrar urina por deficiência ou resistência ao hormônio antidiurético (ADH). A síndrome de secreção inapropriada de ADH (SIADH) apresenta paradoxalmente urina inapropriadamente concentrada na presença de hiponatremia. A interpretação adequada da USG requer correlação com osmolalidade sérica, eletrólitos e quadro clínico, especialmente em idosos onde biomarcadores urinários isolados demonstram confiabilidade limitada. Manter osmolalidade urinária de 24 horas ≤500 mOsm/kg (correspondente a USG aproximada de 1.015) é alvo desejável para saúde renal e redução do risco de nefrolitíase.',

    patient_explanation = 'A densidade da urina (USG) mede o quanto sua urina está concentrada ou diluída, comparando-a com a água pura. Imagine que você adiciona açúcar na água: quanto mais açúcar, mais "densa" a água fica. O mesmo acontece com sua urina - ela contém diversas substâncias dissolvidas (sais minerais, uréia, toxinas) que deixam ela mais "pesada" que a água. Valores normais ficam entre 1.003 e 1.030. Quando você bebe pouca água ou sua muito, seus rins tentam poupar água concentrando a urina (valores acima de 1.020), que fica mais escura e com cheiro forte. Já quando você bebe muita água, a urina fica diluída (valores abaixo de 1.010), clara e abundante. Este exame é importante porque mostra se seus rins estão funcionando bem e se você está adequadamente hidratado. Densidade muito alta pode indicar desidratação - seu corpo precisa de mais água! Densidade muito baixa pode sugerir que você está bebendo água em excesso ou, raramente, que existe algum problema nos rins ou hormônios que controlam a eliminação de água. Seu médico usa este resultado junto com outros exames de sangue e sua história clínica para entender melhor sua saúde. Manter-se bem hidratado (densidade entre 1.010-1.020) é ideal para a saúde dos rins e prevenção de pedras nos rins. A cor da urina também ajuda: amarelo claro geralmente indica boa hidratação, enquanto amarelo escuro sugere necessidade de beber mais água.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: 1) Valores Normais (1.003-1.030): Avaliar contexto clínico e estado de hidratação. USG 1.010-1.020 geralmente indica hidratação adequada. 2) USG Elevada (>1.020-1.030): Sugere desidratação, perdas extrarrenais (vômitos, diarréia, sudorese), ou compensação renal adequada. Investigar sinais clínicos de depleção de volume, verificar osmolalidade sérica (>295 mOsm/kg sugere desidratação), eletrólitos e função renal. Orientar aumento da ingestão hídrica (alvo: osmolalidade urinária <500 mOsm/kg). 3) USG Muito Baixa (<1.008): Hipostenúria indica diabetes insípido (central ou nefrogênico), polidipsia primária, ou fase poliúrica de necrose tubular aguda. Realizar teste de privação hídrica: rins saudáveis concentram urina >800 mOsm/kg, enquanto DI mantém <300 mOsm/kg. Após desmopressina, DI central responde com aumento >50% na osmolalidade, enquanto DI nefrogênico mostra resposta mínima (<15%). Considerar dosagem de copeptina (>21.4 pmol/L diagnostica DI nefrogênico com 100% de sensibilidade). 4) USG Moderadamente Baixa (1.008-1.015): Pode indicar hiper-hidratação fisiológica, uso de diuréticos, ou comprometimento leve da concentração renal. Correlacionar com ingestão hídrica, medicações e função renal basal. 5) Situações Especiais: Em IDOSOS, interpretar com cautela - biomarcadores urinários isolados são não confiáveis; integrar com avaliação clínica completa e osmolalidade sérica. Em ATLETAS, USG ≥1.020 pré-exercício indica desidratação e necessidade de rehidratação. Em NEFROLITÍASE, alvo terapêutico é USG <1.010 (osmolalidade <500 mOsm/kg) para diluição adequada. SEGUIMENTO: Pacientes com alterações persistentes requerem avaliação nefrológica para investigação de tubulopatias, diabetes insípido ou SIADH. Monitorar eletrólitos séricos (sódio, potássio) e função renal (creatinina, uréia). Reavaliação em 1-2 semanas após intervenções (hidratação, ajuste de medicações). Documentar cor da urina como medida complementar de campo.',

    updated_at = CURRENT_TIMESTAMP
WHERE id = 'a03fa6c4-da56-4969-86b7-57a11e7d6907';

COMMIT;

-- ============================================================================
-- VERIFICATION QUERY
-- ============================================================================

SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    (
        SELECT COUNT(*)
        FROM article_score_items asi
        WHERE asi.score_item_id = si.id
    ) as linked_articles_count,
    (
        SELECT string_agg(a.title, ' | ')
        FROM article_score_items asi
        JOIN articles a ON asi.article_id = a.id
        WHERE asi.score_item_id = si.id
    ) as article_titles
FROM score_items si
WHERE si.id = 'a03fa6c4-da56-4969-86b7-57a11e7d6907';
