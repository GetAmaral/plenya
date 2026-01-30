BEGIN;

-- Article 1: Physical Health Decline
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Physical Health Decline After Chemotherapy or Endocrine Therapy Among Older Survivors of Breast Cancer',
    'Ting-Ying Chang, Manali I. Patel, Elizabeth M. Cespedes Feliciano, et al',
    'JAMA Network Open', '2025-01-10', 'en', 'research_article',
    '10.1001/jamanetworkopen.2024.54845', '40019757',
    'Study of 5,461 breast cancer survivors showing persistent physical decline extending beyond 5 years after diagnosis in chemotherapy patients (β=-3.13) compared to age-matched controls.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 2: Metabolic Syndrome and Mortality
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Metabolic syndrome is associated with breast cancer recurrence and breast cancer-specific mortality',
    'Simone J. P. M. Noteborn, Maaike M. G. Leeflang, Roelof J. C. Klöpping, et al',
    'Journal of Internal Medicine', '2025-01-13', 'en', 'meta_analysis',
    '10.1111/joim.20052', '39775978',
    'Meta-analysis of 42,135 survivors showing metabolic syndrome associated with 69% higher recurrence risk (HR 1.69) and 83% higher breast cancer mortality (HR 1.83).'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 3: Surveillance Imaging
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Yield of Surveillance Imaging After Mastectomy in Patients With Early-Stage Breast Cancer',
    'Emily Heer, Priyanka Shahi, Anna N. Wilkinson, et al',
    'JAMA Network Open', '2022-12-01', 'en', 'research_article',
    '10.1001/jamanetworkopen.2022.45229', '36454573',
    'Retrospective study showing low yield of routine surveillance imaging after mastectomy, questioning benefit of routine imaging in asymptomatic patients.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Article 4: Long-term Effects
INSERT INTO articles (title, authors, journal, publish_date, language, article_type, doi, pm_id, abstract)
VALUES (
    'Long-Term Effects of Breast Cancer Surgery, Treatment, and Survivor Care',
    'Cheryl Lacasse, Kathleen M. Neville',
    'Journal of Midwifery and Womens Health', '2019-11-01', 'en', 'review',
    '10.1111/jmwh.13032', '31322834',
    'Comprehensive review showing that 90% of breast cancer survivors experience lasting physical, metabolic, and psychosocial complications requiring multidisciplinary long-term care.'
)
ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id;

-- Create relationships
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'e65c56dc-5c07-4270-8a8b-017b293ca147'::uuid
FROM articles a
WHERE a.doi IN (
    '10.1001/jamanetworkopen.2024.54845',
    '10.1111/joim.20052',
    '10.1001/jamanetworkopen.2022.45229',
    '10.1111/jmwh.13032'
)
ON CONFLICT DO NOTHING;

-- Update score item
UPDATE score_items SET
    clinical_relevance = 'O histórico de mastectomia representa informação crítica na avaliação médica funcional, com implicações que transcendem o tratamento oncológico imediato. Até 90% das sobreviventes de câncer de mama experimentam complicações duradouras que incluem disfunções físicas (dor crônica, linfedema, redução da capacidade aeróbica e força muscular), alterações metabólicas significativas, e impactos psicossociais profundos. Estudos recentes (2025) demonstram que pacientes submetidas à quimioterapia apresentam declínio físico persistente que se estende por mais de 5 anos após o diagnóstico, com reduções mensuráveis na função física (β = -3.13) comparadas a controles pareadas por idade. As sobreviventes que receberam terapia endócrina isolada mostram declínio restrito aos primeiros 2 anos. O histórico de mastectomia também está intrinsecamente associado a maior risco cardiovascular e síndrome metabólica: meta-análise de 2025 envolvendo 42.135 sobreviventes revelou que aquelas com síndrome metabólica ao diagnóstico apresentam risco 69% maior de recorrência (HR 1.69) e 83% maior de mortalidade por câncer de mama (HR 1.83). Alterações metabólicas incluem dislipidemia, resistência insulínica, obesidade e hipertensão, frequentemente exacerbadas pelos tratamentos adjuvantes. A terapia com inibidores de aromatase correlaciona-se com maior risco de diabetes em idosas. Aspectos psicológicos são igualmente relevantes: mastectomia está associada a piora da imagem corporal, saúde sexual comprometida, ansiedade e depressão, com recuperação psicológica frequentemente mais lenta que a física. Na medicina funcional, este histórico demanda avaliação holística englobando função endócrina (especialmente menopausa precoce induzida), saúde óssea (osteoporose secundária a terapia hormonal), capacidade física funcional, perfil metabólico completo, função cognitiva (quimiocerebro) e bem-estar psicossocial.',
    
    patient_explanation = 'Ter realizado uma mastectomia significa que você passou por uma cirurgia importante que pode ter efeitos duradouros na sua saúde geral, indo além do tratamento do câncer em si. É muito importante que você entenda que cerca de 90% das mulheres que passaram por esse procedimento apresentam algum tipo de efeito a longo prazo. Esses efeitos podem incluir: mudanças físicas como dor persistente, inchaço no braço (linfedema), cansaço mais intenso, redução da força muscular e da capacidade de fazer exercícios; mudanças no metabolismo, como tendência a ganhar peso, alterações no colesterol e açúcar no sangue, e maior risco de desenvolver pressão alta ou diabetes; e impactos emocionais, incluindo questões com a imagem corporal, autoestima, vida sexual, e maior vulnerabilidade à ansiedade e depressão. Estudos recentes mostram que mulheres que fizeram quimioterapia podem sentir esses efeitos físicos por mais de 5 anos, enquanto aquelas que fizeram apenas hormonioterapia geralmente melhoram após os 2 primeiros anos. É fundamental saber que mulheres com síndrome metabólica (combinação de pressão alta, açúcar ou gordura elevada no sangue, e excesso de peso abdominal) após a mastectomia têm risco significativamente maior de o câncer voltar. Por isso, cuidar da sua saúde metabólica através de alimentação balanceada, atividade física regular, controle do estresse e sono adequado é essencial não apenas para sua qualidade de vida, mas também para reduzir riscos de recorrência. Informe sempre este histórico aos seus médicos, pois ele influencia muitos aspectos do seu cuidado de saúde.',
    
    conduct = 'AVALIAÇÃO INICIAL ABRANGENTE: Realizar anamnese detalhada investigando tipo de mastectomia (simples, radical modificada, bilateral), lateralidade, data do procedimento, tratamentos adjuvantes recebidos (quimioterapia, radioterapia, hormonioterapia), complicações pós-operatórias, e reconstrução mamária (imediata ou tardia). Avaliar histórico oncológico completo incluindo estadiamento, tipo histológico, status de receptores hormonais e HER2. AVALIAÇÃO FÍSICA FUNCIONAL: Examinar amplitude de movimento do ombro e braço ipsilateral, força muscular, presença de linfedema (medições circunferenciais bilaterais), cicatrizes, áreas de dor ou parestesia, e sinais de complicações tardias. Aplicar questionários de capacidade física funcional e qualidade de vida (FACT-B, EORTC QLQ-C30). RASTREAMENTO METABÓLICO COMPLETO: Solicitar painel metabólico incluindo glicemia de jejum, HbA1c, perfil lipídico completo (colesterol total, LDL, HDL, triglicerídeos), pressão arterial, circunferência abdominal, IMC. Avaliar presença de síndrome metabólica segundo critérios diagnósticos (≥3 dos seguintes: CA >88cm, TG ≥150mg/dL, HDL <50mg/dL, PA ≥130/85mmHg, glicemia ≥100mg/dL). AVALIAÇÃO ENDÓCRINA: Investigar status hormonal especialmente em pacientes em terapia endócrina: perfil tiroideano (TSH, T4 livre), vitamina D, marcadores de saúde óssea se uso de inibidores de aromatase (densitometria óssea, cálcio, fosfatase alcalina). Avaliar sintomas de menopausa e impacto na qualidade de vida. SCREENING CARDIOVASCULAR: Considerar que síndrome metabólica eleva significativamente risco cardiovascular em sobreviventes (HR 1.29 para eventos CV). Avaliar fatores de risco adicionais, considerar ECG basal, ecocardiograma se quimioterapia com antraciclinas ou trastuzumab. AVALIAÇÃO PSICOSSOCIAL: Rastrear depressão (PHQ-9), ansiedade (GAD-7), imagem corporal, função sexual, e qualidade do sono. Encaminhar para suporte psicológico/psiquiátrico se indicado. INTERVENÇÕES TERAPÊUTICAS: Prescrever programa de exercícios supervisionados (combinação aeróbico e resistência) demonstrados eficazes na redução de síndrome metabólica, sarcopenia e melhora da qualidade de vida. Implementar plano nutricional anti-inflamatório mediterrâneo com foco em controle glicêmico e redução de gordura visceral. Suplementação individualizada: vitamina D, ômega-3, probióticos. Considerar fisioterapia para reabilitação funcional do membro superior e linfedema. MONITORAMENTO LONGITUDINAL: Reavaliar função física anualmente. Monitorar marcadores metabólicos semestralmente nos primeiros 2 anos, depois anualmente. Vigilância oncológica conforme guidelines (exame físico, mamografia contralateral se aplicável). Nota: evidências recentes questionam benefício de imagem de rotina pós-mastectomia; basear seguimento em exame clínico. Coordenar cuidado multidisciplinar envolvendo oncologia, cardiologia, endocrinologia, fisioterapia, nutrição e saúde mental conforme necessário.',
    
    last_review = NOW()
WHERE id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';

COMMIT;

SELECT 
    LENGTH(clinical_relevance) as clin_chars,
    LENGTH(patient_explanation) as patient_chars,
    LENGTH(conduct) as conduct_chars,
    (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = 'e65c56dc-5c07-4270-8a8b-017b293ca147') as total_articles
FROM score_items 
WHERE id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';
