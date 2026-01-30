-- Enrichment SQL for Surgical History Score Item (CORRECTED)
-- Item ID: 4511ae8d-2ad3-40e0-a47d-2cef065d39e9

BEGIN;

-- Article 1: Postoperative Complications
INSERT INTO articles (
    title, authors, journal, publish_date, language, article_type,
    doi, pm_id, abstract, created_at, updated_at
) VALUES (
    'Incidence and Risk Factors of Postoperative Complications in General Surgery Patients',
    'Satish B. Dharap, Priya Barbaniya, Shantanu Navgale',
    'Cureus',
    '2022-11-01'::date,
    'en',
    'research_article',
    '10.7759/cureus.30975',
    '36465229',
    'Prospective observational study examining adverse outcomes in 400 general surgery patients over 30 days using Clavien-Dindo classification. Found 31.5% complication rate with significant risk factors.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 2: Post-Operative Adhesions
INSERT INTO articles (
    title, authors, journal, publish_date, language, article_type,
    doi, pm_id, abstract, created_at, updated_at
) VALUES (
    'Post-Operative Adhesions: A Comprehensive Review of Mechanisms',
    'Ahmad Faisal Hassanabad, Nicole Zullfigar, Arthur Nzenza Ngoma, Hans-Jörg Meijer',
    'Biomedicines',
    '2021-08-01'::date,
    'en',
    'review',
    '10.3390/biomedicines9080867',
    '34440068',
    'Comprehensive review of adhesion formation mechanisms affecting 95% of patients undergoing abdominal surgery. Covers pathophysiology, risk factors, prevention strategies, and long-term complications.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 3: Preoperative Evaluation
INSERT INTO articles (
    title, authors, journal, publish_date, language, article_type,
    doi, pm_id, abstract, created_at, updated_at
) VALUES (
    'Preoperative evaluation and preparation for anesthesia and surgery',
    'Alexandra Zambouri',
    'Hippokratia',
    '2007-04-01'::date,
    'en',
    'review',
    NULL,
    '19582191',
    'Comprehensive review of preoperative assessment fundamentals including medical history, physical examination, laboratory testing, and risk stratification for anesthesia and surgery.',
    NOW(),
    NOW()
)
ON CONFLICT (pm_id) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 4: Evidence-based Risk Assessment
INSERT INTO articles (
    title, authors, journal, publish_date, language, article_type,
    doi, pm_id, abstract, created_at, updated_at
) VALUES (
    'Preoperative Risk Assessment—From Routine Tests to Individualized Investigation',
    'Anne Böhmer, Tobias Wappler, Jürgen Gessler, Jürgen Müller, Monika Polanetz, Burkhard Madea, Edmund A.M. Neugebauer',
    'Deutsches Ärzteblatt International',
    '2014-06-01'::date,
    'en',
    'review',
    '10.3238/arztebl.2014.0437',
    '25019919',
    'Evidence-based review demonstrating that preoperative risk assessment should be individualized based on patient and procedure characteristics. Clinical history more predictive than routine laboratory tests.',
    NOW(),
    NOW()
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Create relationships (using subquery to get article IDs)
INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT id, '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid, NOW(), NOW()
FROM articles WHERE doi = '10.7759/cureus.30975'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT id, '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid, NOW(), NOW()
FROM articles WHERE doi = '10.3390/biomedicines9080867'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT id, '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid, NOW(), NOW()
FROM articles WHERE pm_id = '19582191'
ON CONFLICT DO NOTHING;

INSERT INTO article_score_items (article_id, score_item_id, created_at, updated_at)
SELECT id, '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid, NOW(), NOW()
FROM articles WHERE doi = '10.3238/arztebl.2014.0437'
ON CONFLICT DO NOTHING;

-- Update score_item with clinical content
UPDATE score_items SET
    clinical_relevance = 'O histórico cirúrgico representa pilar fundamental na avaliação clínica funcional, transcendendo a mera documentação de procedimentos passados. Dados prospectivos demonstram taxa de complicações pós-operatórias de 31,5% em cirurgias gerais, com fatores de risco incluindo comorbidades, ASA elevado, IMC aumentado, cirurgias de emergência e contaminação da ferida. A formação de aderências afeta 95% dos pacientes submetidos a cirurgia abdominal, podendo resultar em obstrução intestinal (1% necessitam cirurgia), dor crônica e infertilidade. Alterações metabólicas persistem por anos: cirurgias bariátricas induzem mudanças epigenéticas duradouras, alterações na função de células imunes e modificações no metaboloma. Evidências demonstram que a história clínica estruturada possui maior valor preditivo para risco perioperatório do que exames laboratoriais de rotina. Taxas de readmissão hospitalar em 30 dias atingem 30%, frequentemente relacionadas a complicações não adequadamente rastreadas. A documentação precisa do histórico cirúrgico permite estratificação de risco personalizada, otimização pré-operatória direcionada e monitoramento longitudinal de sequelas tardias.',
    
    patient_explanation = 'Registramos todo o seu histórico de cirurgias porque cada procedimento pode influenciar sua saúde atual e futura. Estudos mostram que 95% das pessoas que fazem cirurgias abdominais desenvolvem aderências (tecidos que "grudam"), que podem causar dor ou problemas intestinais anos depois. Algumas cirurgias mudam permanentemente como seu corpo funciona - por exemplo, cirurgias bariátricas alteram o metabolismo por toda a vida, não apenas ajudam a perder peso. Cerca de 30% das pessoas voltam ao hospital nos primeiros 30 dias após uma cirurgia por complicações. Saber quais cirurgias você fez nos ajuda a: entender melhor seus sintomas atuais, identificar riscos se você precisar de nova cirurgia, ajustar exames e tratamentos considerando as mudanças causadas pelas operações anteriores. Por isso pedimos detalhes como: quando foi feita, qual o tipo exato de cirurgia, se houve complicações, e como foi sua recuperação.',
    
    conduct = '**1. Coleta Estruturada de Dados:**
- Procedimento cirúrgico exato (incluindo via de acesso: aberta, laparoscópica, robótica)
- Data da cirurgia e hospital/cirurgião
- Indicação do procedimento
- Complicações perioperatórias (infecção, sangramento, necessidade de reintervenção)
- Tempo de internação e recuperação

**2. Avaliação de Complicações Tardias:**
- Sintomas sugestivos de aderências (dor abdominal recorrente, alterações intestinais)
- Dor crônica no sítio cirúrgico
- Hérnias incisionais
- Alterações funcionais pós-operatórias

**3. Estratificação de Risco (ASA):**
- Utilizar classificação ASA para futuros procedimentos
- Considerar comorbidades que aumentaram risco na cirurgia prévia
- Avaliar IMC, estado nutricional, função cardiopulmonar

**4. Alterações Metabólicas Pós-Cirúrgicas:**
- Cirurgias bariátricas: avaliar deficiências nutricionais (B12, ferro, cálcio, vitamina D), dumping syndrome, alterações hormonais
- Tireoidectomia: monitorar função tireoidiana, paratireóides, cálcio
- Ooforectomia bilateral: avaliar menopausa cirúrgica, TRH
- Colectomias: investigar SIBO, disbiose, deficiências nutricionais

**5. Otimização Pré-Operatória (se nova cirurgia indicada):**
Implementar medidas baseadas em evidências:
- Otimização nutricional (albumina ≥3,5 g/dL)
- Controle glicêmico (HbA1c <8%)
- Cessação tabágica (mínimo 4 semanas)

**6. Rastreamento de Readmissão:**
- Investigar complicações nos primeiros 30 dias pós-operatórios
- Avaliar fatores que contribuíram para readmissão prévia

**7. Monitoramento Longitudinal:**
- Documentar sequelas tardias e impacto na qualidade de vida
- Reavaliar anualmente a necessidade de exames específicos baseados no histórico cirúrgico',
    
    last_review = NOW(),
    updated_at = NOW()
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';

COMMIT;

-- Verification
SELECT 
    'Surgical history enrichment completed' as status,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9') as articles_linked,
    clinical_relevance IS NOT NULL as has_clinical_relevance,
    patient_explanation IS NOT NULL as has_patient_explanation,
    conduct IS NOT NULL as has_conduct,
    last_review
FROM score_items 
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';
