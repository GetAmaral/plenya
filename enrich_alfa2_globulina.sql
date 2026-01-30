-- ============================================================================
-- ENRICHMENT: Alfa-2 Globulina (Alpha-2 Globulin)
-- Score Item ID: 7eb8dd18-6c21-4691-8c19-0f4d785af63e
-- Date: 2026-01-28
-- ============================================================================

BEGIN;

-- ============================================================================
-- STEP 1: Insert Scientific Articles
-- ============================================================================

-- Article 1: Alpha-2-Macroglobulin in Cancer Prognosis (2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Prognostic significance of alpha-2-macroglobulin and low-density lipoprotein receptor-related protein-1 in various cancers',
    'Mateusz Olbromski, Monika Mrozowska, Aleksandra Piotrowska, Alicja Kmiecik, Beata Smolarz, Hanna Romanowicz, Piotr Blasiak, Adam Maciejczyk, Andrzej Wojnar, Piotr Dziegiel',
    'American Journal of Cancer Research',
    '2024-06-15'::date,
    'en',
    'research_article',
    '10.62347/VUJV9180',
    '39005669',
    'This pan-cancer study investigates expression patterns of alpha-2-macroglobulin (A2M) and its receptor LRP1 across breast, lung, and colorectal cancers using immunohistochemical analysis of 909 tissue samples and multiple cancer cell lines. Results demonstrate variable A2M and LRP1 expression correlating with tumor characteristics, with stromal involvement particularly significant in lung and colorectal malignancies. siRNA experiments reveal inverse regulatory relationships between these molecules, suggesting potential as therapeutic targets for cancer treatment development.'
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 2: Alpha-2-Macroglobulin in Inflammation and Immunity (2021)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Alpha-2-Macroglobulin in Inflammation, Immunity and Infections',
    'Jennifer Vandooren, Yoshifumi Itoh',
    'Frontiers in Immunology',
    '2021-12-14'::date,
    'en',
    'review',
    '10.3389/fimmu.2021.803244',
    '34970276',
    'This comprehensive review examines alpha-2-macroglobulin (A2M), a broad-spectrum protease inhibitor functioning as an extracellular macromolecule. Beyond inhibiting proteases, A2M performs diverse roles including cytokine binding, facilitating cell migration, and enabling damaged protein removal. The authors synthesize current knowledge about A2M significance in immune cell function, inflammatory conditions, and microbial defense mechanisms.'
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 3: AI in Serum Protein Electrophoresis (2024)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Artificial intelligence in serum protein electrophoresis: history, state of the art, and perspective',
    'He He, Lingfeng Wang, Xia Wang, Mei Zhang',
    'Critical Reviews in Clinical Laboratory Sciences',
    '2024-05-01'::date,
    'en',
    'review',
    '10.1080/10408363.2023.2274325',
    '37909425',
    'This review examines how artificial intelligence enhances serum protein electrophoresis (SPEP), a blood test that separates proteins by charge and size. The authors discuss how AI technology can streamline analysis by automating protein peak identification, calculating proportions, and detecting abnormalities while reducing human error. The paper addresses AI characteristics, limitations, standardization requirements, and practical guidance for clinical implementation. Integration of AI with SPEP analysis can significantly decrease manual workload and improve result accuracy.'
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- Article 4: Acute Phase Reactants in Dermatology (2022)
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract
) VALUES (
    'Acute Phase Reactants: Relevance in Dermatology',
    'Pulpadathil Jishna, Swapna Dominic',
    'Indian Dermatology Online Journal',
    '2022-12-29'::date,
    'en',
    'review',
    '10.4103/idoj.idoj_174_21',
    '36776186',
    'Acute phase reactants (APRs) are plasma proteins that increase or decrease by at least 25% during inflammatory processes. This comprehensive review examines APR elevation across various dermatological conditions including infections, autoimmune diseases, and malignancies. Key findings include procalcitonin utility in differentiating viral from bacterial infections, ferritin elevation in inflammatory conditions and cancers, and characteristic patterns in systemic lupus erythematosus where ESR rises while CRP remains relatively normal. The authors conclude that further research is needed to establish APRs as reliable diagnostic tools in dermatology practice.'
)
ON CONFLICT (doi) DO UPDATE
SET updated_at = NOW();

-- ============================================================================
-- STEP 2: Update Score Item with Clinical Content
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A fração alfa-2 globulina é composta principalmente por proteínas de fase aguda: alfa-2 macroglobulina (A2M), haptoglobina e ceruloplasmina. Estas proteínas aumentam significativamente durante processos inflamatórios, infecções e trauma. A A2M é um inibidor de protease de amplo espectro (725 kDa) que regula coagulação, migração celular e remoção de proteínas danificadas. A haptoglobina liga-se à hemoglobina livre durante hemólise, prevenindo dano oxidativo. Valores de referência: 0,4-1,0 g/dL (7-15% das proteínas totais). Elevações ocorrem em: inflamação aguda/crônica, síndrome nefrótica (aumento até 10x da A2M pela perda glomerular seletiva), neoplasias (carcinoma renal, mama, pulmão, cólon), artrite reumatoide, lúpus eritematoso sistêmico, infecções bacterianas, gravidez, queimaduras e trauma. Reduções indicam: hemólise intravascular (consumo de haptoglobina), doença hepática (síntese prejudicada), desnutrição proteica grave e anemia hemolítica crônica. A A2M possui papel emergente como biomarcador prognóstico em câncer e alvo terapêutico potencial.',

    patient_explanation = 'A alfa-2 globulina é um grupo de proteínas importantes produzidas pelo fígado que ajudam seu corpo a responder a inflamações, infecções e lesões. Quando você está doente ou machucado, essas proteínas aumentam para proteger seu organismo. As principais proteínas deste grupo são: alfa-2 macroglobulina (que impede enzimas destruidoras de tecidos), haptoglobina (que captura hemoglobina quando glóbulos vermelhos se rompem) e ceruloplasmina (que transporta cobre). Níveis elevados geralmente indicam inflamação ativa, infecção, algumas doenças autoimunes (como artrite reumatoide ou lúpus) ou problemas renais (síndrome nefrótica). Níveis baixos podem sugerir destruição de glóbulos vermelhos (hemólise), problemas no fígado que prejudicam a produção dessas proteínas, ou desnutrição grave. Este exame faz parte do eletroforese de proteínas, que separa as proteínas do sangue em diferentes grupos para avaliar sua saúde geral. Valores normais ficam entre 0,4-1,0 g/dL, representando 7-15% das proteínas totais do sangue.',

    conduct = 'INTERPRETAÇÃO CLÍNICA: Solicitar eletroforese de proteínas séricas com fracionamento completo. Valores de referência alfa-2: 0,4-1,0 g/dL ou 7-15% das proteínas totais. Avaliar sempre em conjunto com proteínas totais, albumina e demais frações (alfa-1, beta, gama).

ALFA-2 ELEVADA (>1,0 g/dL):
1. Investigar inflamação/infecção: solicitar PCR, VHS, leucograma, procalcitonina. Identificar foco infeccioso/inflamatório com história clínica, exame físico e exames dirigidos (RX tórax, urinocultura, hemoculturas).
2. Suspeita de síndrome nefrótica (elevação >2,0 g/dL): solicitar proteinúria 24h, relação proteína/creatinina urinária, função renal, albumina sérica, lipidograma. Considerar biópsia renal se proteinúria nefrótica confirmada.
3. Investigar neoplasia oculta se elevação persistente sem causa aparente: TC tórax/abdome/pelve, marcadores tumorais específicos por idade/sexo, colonoscopia/mamografia conforme rastreamento.
4. Doenças autoimunes: solicitar FAN, anti-DNA, complemento (C3/C4), fator reumatoide, anti-CCP se artrite presente.
5. Repetir em 4-6 semanas após tratamento de condição aguda para documentar normalização.

ALFA-2 REDUZIDA (<0,4 g/dL):
1. Investigar hemólise: solicitar contagem reticulócitos, bilirrubina indireta, LDH, haptoglobina sérica (estará reduzida), teste de Coombs direto, esfregaço sanguíneo.
2. Avaliar função hepática: TGO, TGP, GGT, bilirrubinas, TAP/INR, albumina. Considerar hepatite viral, cirrose ou insuficiência hepática.
3. Desnutrição: avaliar história alimentar, IMC, albumina, pré-albumina, transferrina. Encaminhar para nutricionista.
4. Tratamento dirigido à causa base: transfusão se hemólise grave, imunossupressão se hemólise autoimune, suporte nutricional, tratamento de hepatopatia.

MONITORAMENTO: Repetir eletroforese a cada 3-6 meses em doenças crônicas. Em processos agudos, reavaliar após resolução (4-6 semanas). Documentar tendências longitudinais para avaliar resposta terapêutica.',

    updated_at = NOW()
WHERE id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e';

-- ============================================================================
-- STEP 3: Link Articles to Score Item
-- ============================================================================

-- Link Article 1 (Cancer prognosis 2024)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7eb8dd18-6c21-4691-8c19-0f4d785af63e'::uuid,
    id
FROM articles
WHERE doi = '10.62347/VUJV9180'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 2 (Inflammation review 2021)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7eb8dd18-6c21-4691-8c19-0f4d785af63e'::uuid,
    id
FROM articles
WHERE doi = '10.3389/fimmu.2021.803244'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 3 (AI in SPEP 2024)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7eb8dd18-6c21-4691-8c19-0f4d785af63e'::uuid,
    id
FROM articles
WHERE doi = '10.1080/10408363.2023.2274325'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Link Article 4 (Acute phase reactants 2022)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7eb8dd18-6c21-4691-8c19-0f4d785af63e'::uuid,
    id
FROM articles
WHERE doi = '10.4103/idoj.idoj_174_21'
ON CONFLICT (score_item_id, article_id) DO NOTHING;

COMMIT;

-- ============================================================================
-- VERIFICATION QUERIES
-- ============================================================================

-- Verify score item enrichment with character counts
SELECT
    name,
    LENGTH(clinical_relevance) as clinical_relevance_chars,
    LENGTH(patient_explanation) as patient_explanation_chars,
    LENGTH(conduct) as conduct_chars,
    updated_at
FROM score_items
WHERE id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e';

-- Verify linked articles
SELECT
    si.name as score_item_name,
    COUNT(sia.article_id) as linked_articles_count,
    STRING_AGG(a.title, ' | ' ORDER BY a.publish_date DESC) as article_titles
FROM score_items si
LEFT JOIN article_score_items sia ON si.id = sia.score_item_id
LEFT JOIN articles a ON sia.article_id = a.id
WHERE si.id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e'
GROUP BY si.id, si.name;

-- Show full article details
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id
FROM article_score_items sia
JOIN articles a ON sia.article_id = a.id
WHERE sia.score_item_id = '7eb8dd18-6c21-4691-8c19-0f4d785af63e'
ORDER BY a.publish_date DESC;
