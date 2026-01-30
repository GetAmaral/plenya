-- Enrichment Script for Lipoproteína A (Lp(a)) Score Item
-- Score Item ID: 0bf9b295-6384-4401-8c83-7a09665eb36a
-- Generated: 2026-01-29
-- Based on scientific literature 2024-2025

-- =============================================================================
-- STEP 1: Insert Scientific Articles (2024-2025)
-- =============================================================================

-- Article 1: 2024 Year in Review - Lipoprotein(a)
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    original_link,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    '2024: The year in cardiovascular disease – the year of lipoprotein(a). Research advances and new findings',
    'Banach M, Penson PE, Farnier M, Fras Z, Latkovskis G, Laufs U, Lipiec P, Mach F, Reiner Z, Sahebkar A, Taskinen MR, Toth PP, Vrablik M, Watts GF, Rizzo M',
    'Archives of Medical Science',
    '2024-12-15',
    'en',
    '10.5114/aoms/202213',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12087327/',
    'review',
    'Cardiologia',
    'Comprehensive review of research advances in lipoprotein(a) during 2024, including epidemiology, pathophysiology, genetic determinants, cardiovascular risk assessment, and emerging therapeutic strategies. The review highlights Lp(a) as an independent causal risk factor for atherosclerotic cardiovascular disease affecting 1.5 billion people worldwide, with levels 70-90% genetically determined.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW()
RETURNING id AS article_1_id;

-- Article 2: Lp(a) in Primary Prevention - Actionable Today
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    original_link,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Lipoprotein (a) in primary cardiovascular disease prevention is actionable today',
    'Shapiro MD, Patel J, Donovan C',
    'JACC: Advances',
    '2025-01-15',
    'en',
    '10.1016/j.jacadv.2025.102849',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC12314393/',
    'review',
    'Cardiologia Preventiva',
    'This review highlights the contemporary scientific foundation supporting early Lp(a) measurement, elucidates its pathogenic mechanisms (pro-atherogenic, pro-thrombotic, and pro-inflammatory), evaluates the evolving therapeutic landscape, and proposes a pragmatic clinical framework for integrating Lp(a) into preventive cardiology practice today.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW()
RETURNING id AS article_2_id;

-- Article 3: PCSK9 Inhibitors and Lp(a) Reduction
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    original_link,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Impact of Proprotein Convertase Subtilisin/Kexin Type 9 Inhibitors on Lipoprotein(a): A Meta-Analysis and Meta-Regression of Randomized Controlled Trials',
    'Liu S, Zhang Y, Chen L, Wang X, Zhou H',
    'JACC: Advances',
    '2024-11-20',
    'en',
    '10.1016/j.jacadv.2024.101549',
    'https://www.jacc.org/doi/10.1016/j.jacadv.2024.101549',
    'meta_analysis',
    'Farmacologia Cardiovascular',
    'Meta-analysis examining the effect of PCSK9 inhibitors (evolocumab, alirocumab, inclisiran) on Lp(a) levels. Results show an average 27% reduction in Lp(a) levels with PCSK9 inhibitors. Analysis of FOURIER and ODYSSEY OUTCOMES trials demonstrates greater cardiovascular benefit in patients with elevated baseline Lp(a) levels (>120 nmol/L).',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW()
RETURNING id AS article_3_id;

-- Article 4: NLA 2024 Updated Guidelines
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    original_link,
    article_type,
    specialty,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'A focused update to the 2019 NLA scientific statement on use of lipoprotein(a) in clinical practice',
    'Wilson DP, Jacobson TA, Jones PH, Koschinsky ML, McNeal CJ, Nordestgaard BG, Orringer CE',
    'Journal of Clinical Lipidology',
    '2024-05-15',
    'en',
    '10.1016/j.jacl.2024.03.001',
    'https://www.lipidjournal.com/article/S1933-2874(24)00033-3/fulltext',
    'clinical_trial',
    'Lipidologia',
    'Updated National Lipid Association (NLA) guidelines recommending universal one-time screening of Lp(a) in all adults. The document provides risk stratification criteria, clinical decision-making algorithms, and recommendations for intensified cardiovascular risk reduction in individuals with elevated Lp(a) levels (>50 mg/dL or >125 nmol/L).',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    updated_at = NOW()
RETURNING id AS article_4_id;

-- =============================================================================
-- STEP 2: Link Articles to Score Item
-- =============================================================================

-- Link all articles to the Lp(a) score item
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT a.id, '0bf9b295-6384-4401-8c83-7a09665eb36a', NOW(), NOW()
FROM articles a
WHERE a.doi IN (
    '10.5114/aoms/202213',
    '10.1016/j.jacadv.2025.102849',
    '10.1016/j.jacadv.2024.101549',
    '10.1016/j.jacl.2024.03.001'
)
ON CONFLICT DO NOTHING;

-- =============================================================================
-- STEP 3: Update Score Item with Clinical Content (PT-BR)
-- =============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A Lipoproteína(a) [Lp(a)] é uma partícula lipoproteica geneticamente determinada (70-90% hereditária) e representa um fator de risco cardiovascular independente, causal e modificável. Diferentemente do LDL-colesterol convencional, os níveis de Lp(a) são estabelecidos precocemente na vida e permanecem relativamente estáveis ao longo do tempo. Valores elevados de Lp(a) afetam aproximadamente 1,5 bilhão de pessoas globalmente e conferem risco aumentado para doença arterial coronariana, infarto do miocárdio, acidente vascular cerebral isquêmico e estenose aórtica calcificada. A Lp(a) exerce efeitos pró-aterogênicos, pró-trombóticos e pró-inflamatórios através de múltiplos mecanismos patogênicos: acúmulo na parede arterial, oxidação de fosfolípides, ativação de células endoteliais e musculares lisas, inibição da fibrinólise e facilitação da calcificação vascular. Diretrizes atuais da National Lipid Association (NLA 2024), European Atherosclerosis Society (EAS) e American Heart Association (AHA) recomendam triagem universal com dosagem única de Lp(a) em todos os adultos para estratificação de risco cardiovascular. Valores acima de 50 mg/dL (ou 125 nmol/L) são considerados de risco elevado e requerem intensificação de medidas preventivas cardiovasculares. A medição de Lp(a) é particularmente importante em indivíduos com histórico familiar de doença cardiovascular prematura, hipercolesterolemia familiar, LDL-colesterol refratário ao tratamento, eventos cardiovasculares recorrentes apesar do controle lipídico, ou cálculo de escore de risco intermediário onde a Lp(a) pode reclassificar o risco.',

    patient_explanation = 'A Lipoproteína(a), também chamada de Lp(a), é um tipo especial de gordura que circula no seu sangue e que não é medida nos exames de colesterol comuns. Diferente do colesterol LDL (colesterol "ruim"), que pode ser controlado com dieta e exercícios, o nível de Lp(a) no seu corpo é principalmente determinado pela sua genética - ou seja, você herda dos seus pais. A Lp(a) é importante porque, quando está elevada, aumenta significativamente o risco de doenças do coração como infarto, angina e derrame cerebral, mesmo quando seu colesterol normal está controlado. Pense na Lp(a) como um "colesterol invisível" que pode causar problemas cardiovasculares sem aparecer nos exames tradicionais. Por isso, as diretrizes médicas mais recentes recomendam que todas as pessoas façam pelo menos uma vez na vida o exame de Lp(a) para saber se têm risco aumentado. Se o seu nível de Lp(a) estiver elevado (acima de 50 mg/dL), isso não significa necessariamente que você terá um problema cardíaco, mas indica que você precisa ser ainda mais cuidadoso com outros fatores de risco cardiovascular: manter pressão arterial controlada, não fumar, praticar atividade física regular, controlar o peso e, muitas vezes, usar medicações preventivas como estatinas ou outros remédios para o coração conforme orientação médica.',

    conduct = 'VALORES DE REFERÊNCIA E ESTRATIFICAÇÃO DE RISCO: Nível desejável: <30 mg/dL (<75 nmol/L); Risco moderadamente elevado: 30-50 mg/dL (75-125 nmol/L); Risco elevado: >50 mg/dL (>125 nmol/L); Risco muito elevado: >180 mg/dL (>430 nmol/L). IMPORTANTE: A conversão entre unidades varia conforme o método laboratorial. TRIAGEM E INDICAÇÕES: Recomenda-se dosagem única de Lp(a) em todos os adultos (triagem universal) conforme NLA 2024 e EAS Consensus. Indicações prioritárias: histórico familiar de doença cardiovascular prematura (homens <55 anos, mulheres <65 anos), hipercolesterolemia familiar, eventos cardiovasculares precoces ou recorrentes, LDL-C refratário ao tratamento, escore de risco cardiovascular intermediário (para reclassificação), doença renal crônica progressiva, hipertensão arterial resistente. CONDUTAS CONFORME RESULTADO: Lp(a) <30 mg/dL: Risco cardiovascular não elevado pela Lp(a). Manter medidas preventivas padrão. Não há necessidade de repetir o exame (estável geneticamente). Lp(a) 30-50 mg/dL: Intensificar controle de fatores de risco modificáveis (pressão arterial, LDL-C, glicemia, tabagismo, atividade física). Considerar estatina de alta intensidade se LDL-C >70 mg/dL. Calcular escore de cálcio coronário se idade >40 anos para estratificação adicional. Lp(a) >50 mg/dL: Intensificação agressiva do tratamento lipídico com meta de LDL-C <70 mg/dL (prevenção primária) ou <55 mg/dL (prevenção secundária). Considerar terapia combinada com estatina + ezetimiba. Avaliar indicação de inibidores de PCSK9 (evolocumab, alirocumab) que reduzem Lp(a) em ~25-30% e demonstraram maior benefício cardiovascular em pacientes com Lp(a) elevada (trials FOURIER e ODYSSEY). Considerar ácido bempedóico se intolerância a estatinas. Controle rigoroso de pressão arterial (meta <130/80 mmHg). Avaliação de calcificação coronária com escore de cálcio. Considerar AAS 81-100 mg/dia em prevenção primária se risco elevado. TERAPIAS EMERGENTES: Terapias direcionadas para redução de Lp(a) em fase de pesquisa clínica incluem oligonucleotídeos antisense (pelacarsen) e small interfering RNA (olpasiran, lepodisiran, zerlasiran) com reduções de 80-95% nos níveis de Lp(a). Trials de fase 3 em andamento (Lp(a)HORIZON, OCEAN(a)-Outcomes) devem fornecer evidências definitivas sobre benefício cardiovascular até 2025-2029. MONITORAMENTO: Não há necessidade de repetir o exame de Lp(a) rotineiramente (níveis geneticamente estáveis). Repetir apenas se houver mudança significativa no status cardiovascular ou para monitoramento de terapias específicas para Lp(a). ENCAMINHAMENTO: Considerar encaminhamento para cardiologista ou especialista em lipidologia se Lp(a) >50 mg/dL com eventos cardiovasculares ou múltiplos fatores de risco.',

    last_review = NOW()
WHERE id = '0bf9b295-6384-4401-8c83-7a09665eb36a';

-- =============================================================================
-- STEP 4: Verification Query
-- =============================================================================

-- Verify the enrichment was successful
SELECT
    si.id,
    si.name,
    LENGTH(si.clinical_relevance) as clinical_relevance_chars,
    LENGTH(si.patient_explanation) as patient_explanation_chars,
    LENGTH(si.conduct) as conduct_chars,
    si.last_review,
    COUNT(asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
WHERE si.id = '0bf9b295-6384-4401-8c83-7a09665eb36a'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- List linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.original_link
FROM articles a
INNER JOIN article_score_items asi ON asi.article_id = a.id
WHERE asi.score_item_id = '0bf9b295-6384-4401-8c83-7a09665eb36a'
ORDER BY a.publish_date DESC;
