-- Enrichment for Score Item: Corpos Cetônicos
-- ID: 7ef0ca1b-e7d8-4ac5-82fa-dbc495c54fd0
-- Generated: 2026-01-28

BEGIN;

-- ============================================================================
-- INSERT SCIENTIFIC ARTICLES
-- ============================================================================

-- Article 1: 2025 Review on β-Hydroxybutyrate Testing in Ketogenic Metabolic Therapies
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'The role of β-hydroxybutyrate testing in ketogenic metabolic therapies',
    'Multiple authors from international ketogenic therapy consortium',
    'Frontiers in Nutrition',
    '2025-01-15'::date,
    'en',
    'review',
    '10.3389/fnut.2025.1629921',
    '40959698',
    'This narrative review investigates the role of ketone testing as an integral component of ketogenic metabolic therapies (KMTs), which provide a unique advantage by inducing nutritional ketosis and enabling the use of ketone bodies as biomarkers of metabolic state. Capillary blood beta-hydroxybutyrate (BHB) testing plays multiple roles in KMTs by enabling objective monitoring of dietary adherence, supporting interpretation of clinical outcomes, and informing personalized treatment adjustments. The review covers diverse therapeutic areas including diabetes management, cancer treatment, neurological disorders, and metabolic conditions, presenting a comprehensive overview of current evidence for BHB monitoring across clinical applications.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Article 2: 2023 Comprehensive Update on Measuring Ketones
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Update on Measuring Ketones',
    'Jingtong Huang, Andrea M Yeung, Richard M Bergenstal, Kristin Castorino, Eda Cengiz, Ketan Dhatariya, Isabella Niu, Jennifer L Sherr, Guillermo E Umpierrez, David C Klonoff',
    'Journal of Diabetes Science and Technology',
    '2023-02-16'::date,
    'en',
    'review',
    '10.1177/19322968231152236',
    '36794812',
    'This comprehensive review examines ketone measurement methodologies for diabetes management. Ketone bodies serve as energy substrates during low carbohydrate availability and become pathologically elevated in diabetic ketoacidosis (DKA). The article compares measurement approaches: blood beta-hydroxybutyrate testing provides superior accuracy to urine acetoacetate testing, though both have limitations. Blood ketone measurement is a valuable tool to prevent DKA, given that the rise in blood ketones may precede comparable urine indicators. Emerging technologies, particularly continuous ketone monitoring (CKM) measuring interstitial fluid beta-hydroxybutyrate, represent significant advancement potential. The review addresses multiple clinical scenarios—SGLT2 inhibitor use, low-carbohydrate diets, and automated insulin delivery integration—where enhanced ketone monitoring could improve outcomes and reduce hospitalizations.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- Article 3: 2025 StatPearls - Adult Diabetic Ketoacidosis
-- Note: Using doi placeholder since pm_id doesn't have unique constraint
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    pm_id,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Adult Diabetic Ketoacidosis',
    'Jenna M Lizzo, Amandeep Goyal, Jasleen Kaur',
    'StatPearls [Internet]',
    '2025-11-30'::date,
    'en',
    'lecture',
    '10.32388/statpearls.NBK560723',
    'NBK560723',
    'This comprehensive clinical resource examines diabetic ketoacidosis (DKA) as a serious metabolic emergency characterized by hyperglycemia, acidosis, and ketonemia. While most commonly associated with type 1 diabetes, DKA can also develop in type 2 diabetes patients. The article covers precipitating factors including new-onset diabetes, infections, medication non-adherence, and acute illness, plus emerging concerns like euglycemic DKA linked to SGLT-2 inhibitors, GLP-1 agonists, and immune checkpoint inhibitors. The content addresses pathophysiology, diagnostic criteria (blood glucose >250 mg/dL, arterial pH <7.3, serum bicarbonate <15 mEq/L, presence of ketonemia or ketonuria), clinical presentations, and evidence-based management strategies including fluid resuscitation, insulin therapy, and electrolyte replacement, alongside complications and interprofessional care coordination approaches to optimize patient outcomes.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    pm_id = EXCLUDED.pm_id,
    updated_at = NOW();

-- Article 4: 2024 Hyperglycemic Crises Consensus Report
INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    article_type,
    doi,
    abstract,
    created_at,
    updated_at
) VALUES (
    'Hyperglycemic crises in adults: A look at the 2024 consensus report',
    'International panel representing ADA, AACE, EASD, JBDS-IC, and Diabetes Technology Society',
    'Cleveland Clinic Journal of Medicine',
    '2024-03-01'::date,
    'en',
    'review',
    '10.3949/ccjm.92a.24089',
    'The 2024 consensus report represents an international collaboration to update guidelines for diagnosis, treatment, and prevention of diabetic ketoacidosis (DKA) and hyperglycemic hyperosmolar state (HHS) in adults. Major updates from the 2009 consensus include: (1) Revised diagnostic criteria reducing the glucose cutoff to 200 mg/dL or greater, adding history of diabetes as alternative to glucose values to include euglycemic DKA patients; (2) Ketosis criterion defined as urine ketone strip ≥2+ or beta-hydroxybutyrate ≥3 mmol/L; (3) Resolution criteria specifying blood glucose <200 mg/dL, serum bicarbonate ≥18 mEq/L, and venous pH >7.3. The consensus emphasizes blood beta-hydroxybutyrate measurement over urine ketone testing for superior accuracy, noting that direct beta-hydroxybutyrate measurement is associated with reduced time to recovery and greater cost-effectiveness. Patient education on ketone monitoring before discharge is strongly recommended.',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET
    title = EXCLUDED.title,
    authors = EXCLUDED.authors,
    abstract = EXCLUDED.abstract,
    updated_at = NOW();

-- ============================================================================
-- LINK ARTICLES TO SCORE ITEM
-- ============================================================================

-- Link all articles to the Corpos Cetônicos score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '7ef0ca1b-e7d8-4ac5-82fa-dbc495c54fd0'::uuid,
    a.id
FROM articles a
WHERE a.doi IN (
    '10.3389/fnut.2025.1629921',
    '10.1177/19322968231152236',
    '10.32388/statpearls.NBK560723',
    '10.3949/ccjm.92a.24089'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ============================================================================
-- UPDATE SCORE ITEM WITH ENRICHED CLINICAL CONTENT
-- ============================================================================

UPDATE score_items
SET
    clinical_relevance = 'A presença de corpos cetônicos na urina (cetonúria) ou no sangue (cetonemia) reflete um estado metabólico onde o organismo utiliza gorduras como fonte primária de energia ao invés de glicose. Em condições fisiológicas, a cetose pode ocorrer durante jejum prolongado (>12-16 horas), exercício físico intenso, dietas cetogênicas, gravidez ou lactação. No entanto, a cetonúria patológica representa uma urgência médica potencial, especialmente em pacientes diabéticos. A cetoacidose diabética (CAD) é uma complicação grave caracterizada por hiperglicemia (>200 mg/dL), acidose metabólica (pH <7,3), e cetonemia significativa (beta-hidroxibutirato ≥3,0 mmol/L), com mortalidade de 1-5% mesmo com tratamento adequado. A fisiopatologia envolve deficiência insulínica absoluta ou relativa combinada com excesso de hormônios contrarreguladores (glucagon, cortisol, catecolaminas), levando à lipólise acelerada e produção hepática excessiva de corpos cetônicos (acetoacetato, beta-hidroxibutirato, acetona). Fatores precipitantes incluem infecções (30-40% dos casos), não aderência à insulinoterapia, debut diabético, uso de inibidores SGLT-2 (causando CAD euglicêmica), e estresse fisiológico agudo. A monitorização de cetonas é essencial para prevenção, diagnóstico precoce e acompanhamento terapêutico, sendo a medição sanguínea de beta-hidroxibutirato superior à urinária de acetoacetato em termos de acurácia, precocidade diagnóstica e custo-efetividade. Tecnologias emergentes como monitorização contínua de cetonas (CKM) prometem revolucionar o manejo preventivo em populações de risco.',

    patient_explanation = 'Corpos cetônicos são substâncias produzidas pelo fígado quando seu corpo queima gordura para obter energia. Normalmente, você não deve ter cetonas na urina. Elas aparecem em duas situações principais: situações normais (jejum prolongado de mais de 12 horas, exercício muito intenso, dietas com pouca carboidrato como a dieta cetogênica, ou durante gravidez) e situações de alerta médico (especialmente se você tem diabetes). Para pessoas com diabetes, a presença de cetonas na urina pode indicar um problema sério chamado cetoacidose diabética, que é uma emergência médica. Isso acontece quando falta insulina no corpo e ele começa a quebrar gordura rapidamente demais, produzindo ácidos prejudiciais. Os sinais de alerta incluem: muita sede, urinar frequentemente, cansaço extremo, náusea ou vômitos, dor abdominal, respiração rápida e profunda, hálito com cheiro de frutas (acetona), e confusão mental. Se você tem diabetes e detectar cetonas na urina, especialmente se a glicemia estiver alta (acima de 250 mg/dL), procure atendimento médico imediatamente. O teste pode ser feito em casa com fitas reagentes de urina (que medem acetoacetato) ou com medidores de sangue (que medem beta-hidroxibutirato, sendo mais preciso). Valores de referência: urina negativo a traços; sangue <0,6 mmol/L (normal), 0,6-1,5 mmol/L (cetose leve), 1,6-3,0 mmol/L (atenção), >3,0 mmol/L (emergência médica).',

    conduct = 'AVALIAÇÃO DIAGNÓSTICA: A detecção de cetonúria exige investigação contextualizada. Utilizar preferencialmente medição sanguínea de beta-hidroxibutirato (BHB) quando disponível, pois é mais acurada que a urinária. Fitas urinárias detectam principalmente acetoacetato (resultado semiquantitativo: negativo, traços, +1 a +4), mas podem subestimar gravidade na CAD aguda onde BHB predomina. Critérios diagnósticos da CAD segundo consenso 2024: (1) glicemia ≥200 mg/dL ou história de diabetes, (2) pH arterial <7,3 OU bicarbonato sérico <18 mEq/L, (3) BHB ≥3,0 mmol/L OU cetonúria ≥2+. Classificação de gravidade: CAD leve (pH 7,25-7,30, bicarbonato 15-18 mEq/L), moderada (pH 7,00-7,24, bicarbonato 10-14 mEq/L), grave (pH <7,00, bicarbonato <10 mEq/L). MANEJO DIFERENCIADO POR CONTEXTO: (A) Cetonúria fisiológica (jejum, exercício, dieta cetogênica): Se paciente assintomático, glicemia normal (<140 mg/dL), sem acidose, orientar hidratação e alimentação adequada; BHB 0,5-3,0 mmol/L é esperado em dieta cetogênica terapêutica; monitorizar evolução. (B) CAD estabelecida (emergência médica): Internação hospitalar imediata, preferencialmente UTI se CAD grave; hidratação vigorosa com SF 0,9% (15-20 mL/kg na 1ª hora, depois 250-500 mL/h ajustado); insulinoterapia EV contínua (0,1 UI/kg/h após bolus opcional de 0,1 UI/kg); reposição de potássio guiada por níveis séricos (manter K+ 4-5 mEq/L); monitorização horária de glicemia, eletrólitos, gasometria; pesquisar e tratar fator precipitante (infecção, IAM, AVC); critérios de resolução: glicemia <200 mg/dL, bicarbonato ≥18 mEq/L, pH >7,3, BHB <1,0 mmol/L. (C) CAD euglicêmica (associada a SGLT-2i, GLP-1): Suspeitar em pacientes com cetonemia/cetonúria mas glicemia <250 mg/dL; manejo similar à CAD clássica mas com atenção à administração precoce de glicose EV. EDUCAÇÃO E PREVENÇÃO: Todo paciente diabético deve receber treinamento em monitorização de cetonas (quando testar: glicemia >250 mg/dL, doença intercorrente, náusea/vômito, uso de SGLT-2i); técnica de teste com fitas urinárias ou medidor capilar; reconhecimento de sintomas; plano de ação escrito ("sick day rules"); orientação pré-alta após episódio de CAD é mandatória conforme diretrizes 2024.',

    updated_at = NOW()
WHERE id = '7ef0ca1b-e7d8-4ac5-82fa-dbc495c54fd0';

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
    COUNT(DISTINCT asi.article_id) as linked_articles_count,
    si.updated_at
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '7ef0ca1b-e7d8-4ac5-82fa-dbc495c54fd0'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.updated_at;

-- Show linked articles details
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id,
    LENGTH(a.abstract) as abstract_length
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '7ef0ca1b-e7d8-4ac5-82fa-dbc495c54fd0'
ORDER BY a.publish_date DESC;
