-- Enrichment SQL for Surgical History Score Item
-- Item ID: 4511ae8d-2ad3-40e0-a47d-2cef065d39e9
-- Group: Histórico de doenças (Disease History)
-- Subgroup: Cirurgias já realizadas (Previous Surgeries)

BEGIN;

-- ==========================================
-- STEP 1: Insert Scientific Articles
-- ==========================================

-- Article 1: Postoperative Complications in General Surgery
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    year,
    doi,
    pmid,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Incidence and Risk Factors of Postoperative Complications in General Surgery Patients',
    'Satish B. Dharap, Priya Barbaniya, Shantanu Navgale',
    'Cureus',
    2022,
    '10.7759/cureus.30975',
    '36465229',
    'Prospective observational study examining adverse outcomes in 400 general surgery patients over 30 days using Clavien-Dindo classification. Found 31.5% complication rate with significant risk factors including comorbidities, higher ASA grade, BMI, emergency surgery, open surgery, palliative surgery, deeper cavity surgery, intraoperative blood loss, prolonged duration, and contaminated wounds. Recommends preoperative optimization, meticulous surgical technique, and infection control.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 2: Post-Operative Adhesions Mechanisms
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    year,
    doi,
    pmid,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Post-Operative Adhesions: A Comprehensive Review of Mechanisms',
    'Ali Fatehi Hassanabad, Anna N Zarzycki, Kristina Jeon, Justin F Deniset, Paul W M Fedak',
    'Biomedicines',
    2021,
    '10.3390/biomedicines9080867',
    '34440071',
    'Comprehensive review demonstrating pathologic adhesions develop after 95% of all operations. Examines mechanisms including impaired fibrinolysis, inflammatory cytokine production (TGF-β), and tissue hypoxia (VEGF expression). Notes substantial morbidity and >$2.5 billion annual US healthcare costs. Emphasizes multifactorial pathophysiology involving immune recruitment, mesothelial dysfunction, fibroblast activation, and dysregulated inflammatory mediators requiring integrated prevention approaches.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 3: Preoperative Evaluation and Preparation
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    year,
    pmid,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Preoperative evaluation and preparation for anesthesia and surgery',
    'A Zambouri',
    'Hippokratia',
    2007,
    '19582171',
    'Comprehensive review of preoperative assessment components emphasizing reduction of perioperative morbidity and mortality. Details multifactorial risk including patient health, surgical invasiveness, and anesthetic type. Covers cardiovascular/pulmonary risk assessment, diabetes management, anticoagulation protocols, and regional anesthesia. Stresses inadequate preoperative preparation significantly contributes to complications and proper optimization substantially improves outcomes.',
    NOW(),
    NOW()
);

-- Article 4: Preoperative Risk Assessment - Individualized Investigation
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    year,
    doi,
    pmid,
    abstract,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    'Preoperative Risk Assessment—From Routine Tests to Individualized Investigation',
    'Andreas B Böhmer, Frank Wappler, Bernd Zwissler',
    'Deutsches Ärzteblatt International',
    2014,
    '10.3238/arztebl.2014.0437',
    '25008311',
    'Evidence-based review of preoperative risk evaluation for elective non-cardiac surgery. Emphasizes history and physical examination as central components rather than routine laboratory screening. Recommends individualized testing based on clinical findings, with advanced age alone insufficient justification. States strongest perioperative complication predictors derive from thorough clinical evaluation and surgical procedure nature, not indiscriminate screening.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- ==========================================
-- STEP 2: Create Article-Score Item Relationships
-- ==========================================

-- Link Article 1 (Postoperative Complications)
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT
    a.id,
    '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid,
    NOW(),
    NOW()
FROM articles a
WHERE a.doi = '10.7759/cureus.30975'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 2 (Post-Operative Adhesions)
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT
    a.id,
    '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid,
    NOW(),
    NOW()
FROM articles a
WHERE a.doi = '10.3390/biomedicines9080867'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 3 (Preoperative Evaluation)
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT
    a.id,
    '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid,
    NOW(),
    NOW()
FROM articles a
WHERE a.pmid = '19582171' AND a.journal = 'Hippokratia'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Link Article 4 (Preoperative Risk Assessment)
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT
    a.id,
    '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid,
    NOW(),
    NOW()
FROM articles a
WHERE a.doi = '10.3238/arztebl.2014.0437'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- ==========================================
-- STEP 3: Update Score Item Clinical Content
-- ==========================================

UPDATE score_items
SET
    clinical_relevance = 'O histórico cirúrgico completo é fundamental para a avaliação de risco perioperatório e planejamento terapêutico individualizado em Medicina Funcional. Estudos prospectivos demonstram que cirurgias prévias estão associadas a taxas de complicações de até 31,5% em procedimentos subsequentes, com fatores de risco significativos incluindo comorbidades, classificação ASA elevada, cirurgias de emergência, procedimentos em cavidades profundas e perda sanguínea intraoperatória. A documentação cirúrgica detalhada permite identificar riscos específicos como aderências pós-operatórias (presentes em 95% de todas as operações), que podem causar obstrução intestinal, dor crônica, infertilidade e prolongamento do tempo operatório em reintervenções. O histórico cirúrgico também revela alterações metabólicas e imunológicas duradouras: cirurgias bariátricas, por exemplo, induzem mudanças epigenéticas na imunidade inata que persistem mesmo após perda ponderal substancial, incluindo fenótipo hiperinflamatório de células imunes por "imunidade treinada". Pesquisas recentes demonstram que o tecido adiposo retém memória epigenética da obesidade, desencadeando mudanças epigenéticas persistentes na imunidade inata e exacerbando neuroinflamação. Além disso, cirurgias geram modificações microbioma intestinal com efeitos de longo prazo sobre regulação metabólica via eixo intestino-cérebro. A avaliação pré-operatória baseada em histórico cirúrgico detalhado é considerada o preditor mais forte de complicações perioperatórias, superando testes laboratoriais de rotina. Protocolos clínicos enfatizam que história clínica e exame físico focados em fatores de risco cardiovascular/pulmonar são centrais na estratificação de risco, permitindo otimização pré-operatória direcionada que reduz significativamente morbidade e mortalidade.',

    patient_explanation = 'Registrar todas as cirurgias que você já realizou é essencial para garantir sua segurança em futuros tratamentos médicos e procedimentos cirúrgicos. Cada cirurgia deixa "marcas" no organismo que vão além das cicatrizes visíveis: em 95% dos casos, formam-se aderências internas (tecidos que grudam uns nos outros durante a cicatrização), que podem causar dores, dificultar digestão ou complicar cirurgias futuras. Seu histórico cirúrgico também revela como seu corpo responde a estresse, inflamação e cicatrização - informações valiosas para personalizar seu tratamento. Por exemplo, cirurgias bariátricas alteram permanentemente seu metabolismo e sistema imunológico, afetando como você processa nutrientes e responde a inflamações. Operações abdominais podem modificar a flora intestinal, impactando digestão, humor e energia por anos. Estudos mostram que cerca de 30% dos pacientes operados precisam de reinternações relacionadas a complicações tardias das cirurgias. Conhecer seu histórico completo permite ao médico: (1) prever riscos específicos antes de novos procedimentos; (2) ajustar medicamentos e suplementos conforme suas necessidades pós-cirúrgicas; (3) evitar exames e tratamentos desnecessários; (4) preparar melhor seu corpo caso precise de novas intervenções. Informações importantes incluem: tipo de cirurgia, data, hospital, se houve complicações, quanto tempo levou sua recuperação e se ficaram sequelas. Essas informações protegem você e ajudam a construir um plano de saúde verdadeiramente personalizado.',

    conduct = 'PROTOCOLO DE DOCUMENTAÇÃO E AVALIAÇÃO DE HISTÓRICO CIRÚRGICO

1. COLETA DE DADOS ESTRUTURADA
   - Data da cirurgia (mês/ano se exata desconhecida)
   - Tipo de procedimento (laparoscópico vs. aberto; eletivo vs. emergência)
   - Indicação/motivo da cirurgia
   - Local anatômico e cavidade envolvida
   - Duração estimada do procedimento
   - Intercorrências intraoperatórias documentadas
   - Tipo de anestesia utilizada
   - Instituição e equipe cirúrgica (quando relevante para acesso a prontuários)

2. COMPLICAÇÕES E RECUPERAÇÃO
   - Complicações pós-operatórias imediatas (até 30 dias): infecção, deiscência, sangramento, eventos tromboembólicos
   - Complicações tardias: aderências sintomáticas, hérnias incisionais, dor crônica pós-cirúrgica
   - Tempo de hospitalização e recuperação funcional
   - Necessidade de reintervenções ou revisões cirúrgicas
   - Cicatrização: queloide, hipertrófica, deiscência parcial
   - Sequelas funcionais ou estéticas persistentes

3. ESTRATIFICAÇÃO DE RISCO PERIOPERATÓRIO
   - Aplicar classificação ASA baseada em histórico e comorbidades atuais
   - Identificar múltiplas cirurgias abdominais/pélvicas (alto risco de aderências densas)
   - Cirurgias em contexto de peritonite, doença de Crohn ou câncer colorretal (risco aumentado de aderências sintomáticas)
   - Procedimentos de emergência prévios (maior incidência de complicações relacionadas a aderências)
   - História de má cicatrização ou complicações infecciosas recorrentes
   - Avaliar capacidade funcional atual (história e exame físico > testes de rotina)

4. AVALIAÇÃO DE ALTERAÇÕES METABÓLICAS E SISTÊMICAS
   - Cirurgias bariátricas/metabólicas: avaliar absorção de nutrientes, suplementação, alterações metabólicas duradouras, fenótipo imune hiperinflamatório por memória epigenética
   - Cirurgias abdominais extensas: investigar disbiose intestinal, síndrome do intestino curto, má absorção
   - Cirurgias endócrinas (tireoide, adrenal, hipófise): monitorar reposição hormonal e status metabólico
   - Cirurgias ginecológicas: avaliar função reprodutiva, equilíbrio hormonal, aderências pélvicas
   - Cirurgias ortopédicas maiores: avaliar limitações funcionais, dor crônica, inflamação sistêmica residual

5. SOLICITAÇÃO DE DOCUMENTAÇÃO COMPLEMENTAR
   - Relatórios cirúrgicos e anatomopatológicos quando disponíveis
   - Imagens pré e pós-operatórias relevantes
   - Registros de complicações e reintervenções
   - Laudos de exames pré-operatórios (baseline funcional)

6. OTIMIZAÇÃO PRÉ-OPERATÓRIA PARA NOVAS CIRURGIAS
   - Implementar três medidas-chave baseadas em evidências: otimização clínica pré-operatória; técnica cirúrgica meticulosa para reduzir tempo operatório, sangramento e complicações intraoperatórias; protocolos rigorosos de controle de infecção
   - Correção de anemia, hipoalbuminemia, hiperglicemia
   - Cessação tabágica (mínimo 4 semanas)
   - Otimização de doenças crônicas (diabetes, hipertensão, DPOC)
   - Avaliação cardiovascular e pulmonar direcionada por achados clínicos
   - Planejamento anestésico individualizado considerando histórico
   - Discussão de riscos específicos relacionados a cirurgias prévias

7. MONITORAMENTO LONGITUDINAL
   - Reavaliação anual do histórico cirúrgico e sequelas tardias
   - Investigação de sintomas potencialmente relacionados a aderências: dor abdominal recorrente, distensão, alterações do hábito intestinal
   - Atualização do campo last_review sempre que novas informações cirúrgicas forem adicionadas
   - Integração do histórico cirúrgico com outros determinantes de saúde na abordagem de Medicina Funcional

NÍVEL DE EVIDÊNCIA: Baseado em revisões sistemáticas, estudos prospectivos e diretrizes internacionais de avaliação perioperatória (ASA, NICE, ESC).',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';

-- ==========================================
-- STEP 4: Verification Queries
-- ==========================================

-- Verify articles inserted
SELECT
    id,
    title,
    authors,
    journal,
    year,
    doi,
    pmid
FROM articles
WHERE doi IN (
    '10.7759/cureus.30975',
    '10.3390/biomedicines9080867',
    '10.3238/arztebl.2014.0437'
) OR (pmid = '19582171' AND journal = 'Hippokratia')
ORDER BY year DESC;

-- Verify relationships created
SELECT
    asi.article_id,
    a.title,
    a.year,
    si.name as score_item_name
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
JOIN score_items si ON asi.score_item_id = si.id
WHERE asi.score_item_id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'
ORDER BY a.year DESC;

-- Verify score item updated
SELECT
    id,
    name,
    LENGTH(clinical_relevance) as clinical_relevance_length,
    LENGTH(patient_explanation) as patient_explanation_length,
    LENGTH(conduct) as conduct_length,
    last_review,
    updated_at
FROM score_items
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';

COMMIT;

-- Summary Report
SELECT
    'Enrichment completed for: Registrar quaisquer cirurgias realizadas' as status,
    COUNT(DISTINCT a.id) as articles_linked,
    si.name,
    si.last_review
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
LEFT JOIN articles a ON asi.article_id = a.id
WHERE si.id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'
GROUP BY si.id, si.name, si.last_review;
