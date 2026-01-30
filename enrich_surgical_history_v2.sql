BEGIN;

-- Article 1
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Incidence and Risk Factors of Postoperative Complications in General Surgery Patients',
    'Satish B. Dharap, Priya Barbaniya, Shantanu Navgale',
    'Cureus', '2022-11-01', 'en', 'research_article',
    '10.7759/cureus.30975', '36465229',
    'Prospective study of 400 general surgery patients showing 31.5% complication rate with risk factors including comorbidities, ASA grade, BMI, emergency surgery.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 2
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Post-Operative Adhesions: A Comprehensive Review of Mechanisms',
    'Ahmad Faisal Hassanabad, Nicole Zullfigar, Arthur Nzenza Ngoma, Hans-Jörg Meijer',
    'Biomedicines', '2021-08-01', 'en', 'review',
    '10.3390/biomedicines9080867', '34440068',
    'Review of adhesion formation affecting 95% of patients after abdominal surgery, covering pathophysiology and prevention.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 3 (no DOI, use different approach)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '19582191') THEN
        INSERT INTO articles (title, authors, journal, publish_date, language, article_type, pm_id, abstract)
        VALUES (
            'Preoperative evaluation and preparation for anesthesia and surgery',
            'Alexandra Zambouri',
            'Hippokratia', '2007-04-01', 'en', 'review',
            '19582191',
            'Review of preoperative assessment including medical history, physical examination, and risk stratification.'
        );
    END IF;
END $$;

-- Article 4
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Preoperative Risk Assessment—From Routine Tests to Individualized Investigation',
    'Anne Böhmer, Tobias Wappler, Jürgen Gessler, et al',
    'Deutsches Ärzteblatt International', '2014-06-01', 'en', 'review',
    '10.3238/arztebl.2014.0437', '25019919',
    'Evidence that preoperative assessment should be individualized. Clinical history more predictive than routine labs.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Create relationships
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '4511ae8d-2ad3-40e0-a47d-2cef065d39e9'::uuid
FROM articles a
WHERE a.doi IN ('10.7759/cureus.30975', '10.3390/biomedicines9080867', '10.3238/arztebl.2014.0437')
   OR a.pm_id = '19582191'
ON CONFLICT DO NOTHING;

-- Update score item
UPDATE score_items SET
    clinical_relevance = 'O histórico cirúrgico representa pilar fundamental na avaliação clínica funcional, transcendendo a mera documentação de procedimentos passados. Dados prospectivos demonstram taxa de complicações pós-operatórias de 31,5% em cirurgias gerais, com fatores de risco incluindo comorbidades, ASA elevado, IMC aumentado, cirurgias de emergência e contaminação da ferida. A formação de aderências afeta 95% dos pacientes submetidos a cirurgia abdominal, podendo resultar em obstrução intestinal (1% necessitam cirurgia), dor crônica e infertilidade. Alterações metabólicas persistem por anos: cirurgias bariátricas induzem mudanças epigenéticas duradouras, alterações na função de células imunes e modificações no metaboloma. Evidências demonstram que a história clínica estruturada possui maior valor preditivo para risco perioperatório do que exames laboratoriais de rotina. Taxas de readmissão hospitalar em 30 dias atingem 30%, frequentemente relacionadas a complicações não adequadamente rastreadas.',
    
    patient_explanation = 'Registramos todo o seu histórico de cirurgias porque cada procedimento pode influenciar sua saúde atual e futura. Estudos mostram que 95% das pessoas que fazem cirurgias abdominais desenvolvem aderências (tecidos que "grudam"), que podem causar dor ou problemas intestinais anos depois. Algumas cirurgias mudam permanentemente como seu corpo funciona - por exemplo, cirurgias bariátricas alteram o metabolismo por toda a vida. Cerca de 30% das pessoas voltam ao hospital nos primeiros 30 dias após uma cirurgia. Saber quais cirurgias você fez nos ajuda a entender sintomas atuais, identificar riscos para novas cirurgias, e ajustar tratamentos considerando as mudanças causadas pelas operações anteriores.',
    
    conduct = '**1. Coleta Estruturada:**
Procedimento exato (via de acesso), data, hospital, indicação, complicações, tempo de recuperação.

**2. Complicações Tardias:**
Aderências (dor, alterações intestinais), dor crônica, hérnias, alterações funcionais.

**3. Estratificação de Risco:**
Classificação ASA, comorbidades, IMC, função cardiopulmonar.

**4. Alterações Metabólicas:**
- Bariátricas: deficiências B12/ferro/cálcio/vitamina D, dumping, alterações hormonais
- Tireoidectomia: TSH, paratireóides, cálcio
- Ooforectomia: menopausa cirúrgica, TRH
- Colectomias: SIBO, disbiose

**5. Otimização Pré-Operatória:**
Nutrição (albumina ≥3,5), controle glicêmico (HbA1c <8%), cessação tabágica (4+ semanas).

**6. Monitoramento:**
Readmissões 30 dias, sequelas tardias, impacto na qualidade de vida.',
    
    last_review = NOW()
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';

COMMIT;

SELECT 
    LENGTH(clinical_relevance) as clin_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9') as total_articles
FROM score_items 
WHERE id = '4511ae8d-2ad3-40e0-a47d-2cef065d39e9';
