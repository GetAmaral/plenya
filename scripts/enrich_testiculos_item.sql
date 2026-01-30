-- Enrichment for Score Item: Testículos (ID: 933b7816-fe10-4e35-94d7-0d232cc258ce)
-- Generated: 2026-01-29
-- Focus: Testicular examination, ultrasound assessment, testicular disorders

BEGIN;

-- Update clinical content for score item "Testículos"
UPDATE score_items
SET
    clinical_relevance = 'O exame testicular é uma avaliação fundamental na saúde masculina, essencial para detecção precoce de alterações estruturais, inflamatórias e neoplásicas. A palpação testicular combinada com ultrassonografia multiparamétrica permite identificar massas testiculares, varicocele, orquite, epididimite e outras patologias. O câncer testicular é a malignidade mais comum em homens de 15 a 34 anos, com incidência estimada de 9.720 novos casos anuais nos EUA, apresentando taxa de cura superior a 95% quando detectado precocemente. A avaliação clínica deve incluir inspeção visual, palpação bilateral comparativa, avaliação de tamanho, consistência e presença de massas. O ultrassom Doppler colorido com elastografia por ondas de cisalhamento (shear wave elastography) é padrão-ouro para caracterização de lesões focais, diferenciando entre patologias benignas e malignas com maior acurácia que o ultrassom convencional. A ultrassonografia multiparamétrica (incluindo CEUS - contrast-enhanced ultrasound e SE - strain elastography) melhora significativamente o desempenho diagnóstico e pode apoiar abordagens de preservação testicular ao invés de cirurgia radical. Varicocele afeta 15-20% dos homens adultos e pode impactar negativamente a fertilidade e função testicular. Estudos demonstram correlação significativa entre epididimoorquite prévia e câncer testicular (OR=38,24; IC95%: 19,91-73,46), reforçando a importância do seguimento longitudinal. Diretrizes da European Association of Urology (EAU 2024-2025) e American Urological Association (AUA) recomendam que toda massa sólida intratesticular identificada por exame físico ou imagem seja considerada neoplasia maligna até prova em contrário. Marcadores tumorais séricos (AFP, beta-hCG e LDH) devem ser coletados antes de qualquer tratamento, incluindo orquiectomia.',

    patient_explanation = 'O exame dos testículos é uma avaliação simples e rápida que permite detectar alterações importantes na saúde masculina. O médico examina os testículos através da palpação (toque) e pode solicitar ultrassom se identificar alguma anormalidade. Este exame busca identificar: nódulos ou massas (que podem ser tumores benignos ou malignos), aumento de tamanho ou mudança de consistência, presença de líquido ao redor do testículo (hidrocele), veias dilatadas (varicocele, que pode afetar a fertilidade), sinais de inflamação ou infecção (orquite/epididimite), e dor ou sensibilidade anormal. O câncer testicular é altamente curável quando detectado cedo, com taxa de cura acima de 95%. Por isso, é importante procurar o médico se notar qualquer caroço, inchaço, dor persistente ou mudança nos testículos. O ultrassom testicular é um exame indolor que fornece imagens detalhadas e pode identificar problemas não detectáveis pela palpação. Homens jovens (15-34 anos) têm maior risco para câncer testicular e devem estar atentos a mudanças. Não há consenso sobre autoexame testicular periódico, mas conhecer a anatomia normal facilita identificar alterações precoces.',

    conduct = 'EXAME CLÍNICO: Realizar inspeção visual bilateral do escroto (assimetria, eritema, edema, lesões cutâneas), palpação bimanual comparativa de cada testículo (consistência, tamanho estimado, presença de massas), avaliação do epidídimo e cordão espermático, pesquisa de varicocele (exame em pé com manobra de Valsalva), transiluminação se hidrocele ou espermatocele suspeita. ULTRASSONOGRAFIA: Indicada para qualquer massa palpável, aumento de volume testicular assimétrico, dor testicular persistente, infertilidade masculina, trauma escrotal, suspeita de torção testicular. Ultrassom Doppler colorido avalia fluxo vascular (fundamental para varicocele e torção). Elastografia por ondas de cisalhamento melhora caracterização de lesões focais (rigidez tecidual aumentada sugere malignidade). Medida volumétrica testicular bilateral (volume normal: 12-30mL). Classificação de varicocele segundo ESUR-SPIWG com manobra de Valsalva em posição ortostática. MARCADORES TUMORAIS: Dosagem de AFP, beta-hCG e LDH antes de qualquer intervenção cirúrgica em suspeita de neoplasia. CONDUTA ESPECÍFICA: Massa sólida intratesticular → encaminhar urologia com urgência (orquiectomia radical via inguinal). Varicocele grau III com subfertilidade → avaliar varicocelectomia. Orquite/epididimite aguda → antibioticoterapia guiada por cultura (considerar Chlamydia/Gonorreia em jovens sexualmente ativos). Hidrocele/espermatocele volumosas → considerar correção cirúrgica. Seguimento longitudinal em pacientes com história de epididimoorquite prévia (associação com câncer testicular). Educação sobre reconhecimento de sinais de alerta e importância de buscar avaliação médica precoce.',

    last_review = CURRENT_DATE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = '933b7816-fe10-4e35-94d7-0d232cc258ce';

-- Verify update
DO $$
DECLARE
    v_clinical_length INTEGER;
    v_patient_length INTEGER;
    v_conduct_length INTEGER;
BEGIN
    SELECT
        LENGTH(clinical_relevance),
        LENGTH(patient_explanation),
        LENGTH(conduct)
    INTO v_clinical_length, v_patient_length, v_conduct_length
    FROM score_items
    WHERE id = '933b7816-fe10-4e35-94d7-0d232cc258ce';

    RAISE NOTICE 'Character counts:';
    RAISE NOTICE '  clinical_relevance: % chars (target: 1500-2000)', v_clinical_length;
    RAISE NOTICE '  patient_explanation: % chars (target: 1000-1500)', v_patient_length;
    RAISE NOTICE '  conduct: % chars (target: 1500-2500)', v_conduct_length;

    IF v_clinical_length < 1500 OR v_clinical_length > 2000 THEN
        RAISE WARNING 'clinical_relevance length out of target range';
    END IF;
    IF v_patient_length < 1000 OR v_patient_length > 1500 THEN
        RAISE WARNING 'patient_explanation length out of target range';
    END IF;
    IF v_conduct_length < 1500 OR v_conduct_length > 2500 THEN
        RAISE WARNING 'conduct length out of target range';
    END IF;
END $$;

-- Insert scientific articles
-- Article 1: Multiparametric Ultrasound for Focal Testicular Pathology (2024)
INSERT INTO articles (
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
    'Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review',
    'Huang DY, Alsadiq M, Yusuf GT, Deganello A, Sellars ME, Sidhu PS',
    'Cancers (Basel)',
    '39001372',
    '10.3390/cancers16132309',
    '2024-06-24'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 1 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '933b7816-fe10-4e35-94d7-0d232cc258ce'::uuid
FROM articles a
WHERE a.doi = '10.3390/cancers16132309'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 2: Ultrasound Assessment of Testicular Volume (2023)
INSERT INTO articles (
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
    'Ultrasound assessment of testicular volume - An interobserver variability study',
    'Pedersen MRV, Otto PO, Fredslund M, Smedegaard C, Jensen J, McEntee MF, Loft MK',
    'Journal of Medical Imaging and Radiation Sciences',
    '37838500',
    '10.1016/j.jmir.2023.09.001',
    '2023-12-01'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 2 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '933b7816-fe10-4e35-94d7-0d232cc258ce'::uuid
FROM articles a
WHERE a.doi = '10.1016/j.jmir.2023.09.001'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 3: Multiparametric Ultrasound for Testicular Lesions (2023)
INSERT INTO articles (
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
    'Multiparametric ultrasound for the assessment of testicular lesions with negative tumoral markers',
    'Liu H, Dong L, Xiang LH, Xu G, Wan J, Fang Y, Ding SS, Jin Y, Sun LP, Xu HX',
    'Asian Journal of Andrology',
    '35708357',
    '10.4103/aja202235',
    '2023-01-01'::date,
    'research_article',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 3 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '933b7816-fe10-4e35-94d7-0d232cc258ce'::uuid
FROM articles a
WHERE a.doi = '10.4103/aja202235'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Article 4: Call to Action for Testicular Self-Examination (2022)
INSERT INTO articles (
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
    'A Call to Action to Review the USPSTF''s Recommendation for Testicular Self-Examination',
    'Rovito MJ, Allen K, Nangia A, Craycraft M, Cary C, Lutz M, Lyon T, Fadich A, Baird B, Welch MG, Alcantara A',
    'American Journal of Men''s Health',
    '36214273',
    '10.1177/15579883221130186',
    '2022-09-01'::date,
    'review',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    journal = EXCLUDED.journal,
    pm_id = EXCLUDED.pm_id,
    publish_date = EXCLUDED.publish_date,
    article_type = EXCLUDED.article_type,
    updated_at = CURRENT_TIMESTAMP;

-- Link Article 4 to score item
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '933b7816-fe10-4e35-94d7-0d232cc258ce'::uuid
FROM articles a
WHERE a.doi = '10.1177/15579883221130186'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Final verification
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_length,
    LENGTH(si.patient_explanation) as patient_length,
    LENGTH(si.conduct) as conduct_length,
    si.last_review,
    COUNT(asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '933b7816-fe10-4e35-94d7-0d232cc258ce'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

COMMIT;
