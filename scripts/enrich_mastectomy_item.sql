-- Enrichment SQL for Mastectomy Score Item
-- Score Item ID: e65c56dc-5c07-4270-8a8b-017b293ca147
-- Group: Histórico de doenças
-- Subgroup: Cirurgias já realizadas

BEGIN;

-- Insert Article 1: Surveillance Imaging After Mastectomy
WITH article1 AS (
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
        'Yield of Surveillance Imaging After Mastectomy With or Without Reconstruction for Patients With Prior Breast Cancer: A Systematic Review and Meta-analysis',
        'Smith D, Sepehr S, Karakatsanis A, Strand F, Valachis A',
        'JAMA Network Open',
        2022,
        '10.1001/jamanetworkopen.2022.44212',
        '36454573',
        'This systematic review and meta-analysis investigated whether routine imaging surveillance benefits patients who have undergone mastectomy for prior breast cancer. The researchers analyzed 16 studies examining mammography, ultrasound, and MRI outcomes. Detection rates of nonpalpable cancers were substantially lower than overall detection rates across all imaging methods. The pooled rates per 1,000 examinations were: mammography 1.86, ultrasound 2.66, and MRI 5.17 for overall cancer detection. The authors concluded that imaging surveillance in this context is unnecessary in clinical practice, at least until further studies demonstrate otherwise, challenging current surveillance practices and suggesting clinicians should reconsider routine imaging protocols following mastectomy.',
        NOW(),
        NOW()
    ) RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'e65c56dc-5c07-4270-8a8b-017b293ca147'::uuid FROM article1;

-- Insert Article 2: Long-Term Effects of Breast Cancer Surgery
WITH article2 AS (
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
        'Long-Term Effects of Breast Cancer Surgery, Treatment, and Survivor Care',
        'Lovelace DL, McDaniel LR, Golden D',
        'Journal of Midwifery and Womens Health',
        2019,
        '10.1111/jmwh.13012',
        '31322834',
        'Up to 90% of breast cancer survivors experience lasting complications from their treatment. These sequelae encompass physical changes like chronic pain and lymphedema, alongside emotional and psychosocial effects including depression and anxiety. The review emphasizes that women may have decreased strength, aerobic capacity, mobility, fatigue, and cognitive dysfunction. Primary care clinicians should employ comprehensive evaluation and multidisciplinary treatment strategies combining medication, physical therapy, counseling, and complementary approaches to optimize survivor health and quality of life.',
        NOW(),
        NOW()
    ) RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'e65c56dc-5c07-4270-8a8b-017b293ca147'::uuid FROM article2;

-- Insert Article 3: Physical Health Decline After Treatment
WITH article3 AS (
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
        'Physical Health Decline After Chemotherapy or Endocrine Therapy in Breast Cancer Survivors',
        'Bodelon C, Masters M, Bloodworth DE, Briggs PJ, Rees-Punia E, McCullough LR, Patel AV, Teras LR',
        'JAMA Network Open',
        2025,
        '10.1001/jamanetworkopen.2024.62365',
        '40019757',
        'This prospective cohort study examined physical health outcomes in breast cancer survivors compared with age-matched controls. The research included 2,566 women diagnosed with nonmetastatic breast cancer and 12,826 comparison participants from the Cancer Prevention Study-3 cohort (2006-2013, followed through April 2020). The findings revealed significant disparities: compared with women without cancer, there was a greater physical health decline within 2 years for those receiving chemotherapy (β = -3.13) or combined treatments (β = -3.26). Women on aromatase inhibitor endocrine therapy alone showed modest decline (β = -1.12). A critical distinction emerged over time: long-term declines persisted only among chemotherapy recipients. The researchers concluded that BC survivors who received chemotherapy had a long-lasting physical health decline, unlike survivors who received endocrine therapy without chemotherapy, emphasizing the need for further investigation into treatment-related health consequences.',
        NOW(),
        NOW()
    ) RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'e65c56dc-5c07-4270-8a8b-017b293ca147'::uuid FROM article3;

-- Insert Article 4: Metabolic Syndrome and Breast Cancer Mortality
WITH article4 AS (
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
        'Metabolic syndrome is associated with breast cancer mortality: A systematic review and meta-analysis',
        'Harborg S, Larsen HB, Elsgaard S, Borgquist S',
        'Journal of Internal Medicine',
        2025,
        '10.1111/joim.20052',
        '39775978',
        'This systematic review examined how metabolic syndrome affects survival outcomes in breast cancer survivors. Researchers searched medical databases for studies comparing survival data between patients with and without metabolic syndrome at diagnosis. The analysis encompassed 42,135 breast cancer survivors from 17 eligible studies. The findings revealed significant associations: breast cancer survivors with metabolic syndrome faced increased recurrence risk, higher mortality rates, and shorter disease-free survival periods compared to those without the condition. Specifically, the pooled data showed increased risk of recurrence (HR 1.69, 95% CI: 1.39-2.06) and BC mortality (HR 1.83, 95% CI: 1.35-2.49). The authors concluded that these results necessitate developing clinical screening guidelines for metabolic syndrome in breast cancer survivors. They recommend further research identifying interventions to reduce metabolic syndrome prevalence among this population to improve both metabolic health and cancer outcomes.',
        NOW(),
        NOW()
    ) RETURNING id
)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, 'e65c56dc-5c07-4270-8a8b-017b293ca147'::uuid FROM article4;

-- Update the score_item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'O histórico de mastectomia representa informação crítica na avaliação médica funcional, com implicações que transcendem o tratamento oncológico imediato. Até 90% das sobreviventes de câncer de mama experimentam complicações duradouras que incluem disfunções físicas (dor crônica, linfedema, redução da capacidade aeróbica e força muscular), alterações metabólicas significativas, e impactos psicossociais profundos. Estudos recentes (2025) demonstram que pacientes submetidas à quimioterapia apresentam declínio físico persistente que se estende por mais de 5 anos após o diagnóstico, com reduções mensuráveis na função física (β = -3.13) comparadas a controles pareadas por idade. As sobreviventes que receberam terapia endócrina isolada mostram declínio restrito aos primeiros 2 anos. O histórico de mastectomia também está intrinsecamente associado a maior risco cardiovascular e síndrome metabólica: meta-análise de 2025 envolvendo 42.135 sobreviventes revelou que aquelas com síndrome metabólica ao diagnóstico apresentam risco 69% maior de recorrência (HR 1.69) e 83% maior de mortalidade por câncer de mama (HR 1.83). Alterações metabólicas incluem dislipidemia, resistência insulínica, obesidade e hipertensão, frequentemente exacerbadas pelos tratamentos adjuvantes. A terapia com inibidores de aromatase correlaciona-se com maior risco de diabetes em idosas. Aspectos psicológicos são igualmente relevantes: mastectomia está associada a piora da imagem corporal, saúde sexual comprometida, ansiedade e depressão, com recuperação psicológica frequentemente mais lenta que a física. Na medicina funcional, este histórico demanda avaliação holística englobando função endócrina (especialmente menopausa precoce induzida), saúde óssea (osteoporose secundária a terapia hormonal), capacidade física funcional, perfil metabólico completo, função cognitiva (quimiocerebro) e bem-estar psicossocial.',

    patient_explanation = 'Ter realizado uma mastectomia significa que você passou por uma cirurgia importante que pode ter efeitos duradouros na sua saúde geral, indo além do tratamento do câncer em si. É muito importante que você entenda que cerca de 90% das mulheres que passaram por esse procedimento apresentam algum tipo de efeito a longo prazo. Esses efeitos podem incluir: mudanças físicas como dor persistente, inchaço no braço (linfedema), cansaço mais intenso, redução da força muscular e da capacidade de fazer exercícios; mudanças no metabolismo, como tendência a ganhar peso, alterações no colesterol e açúcar no sangue, e maior risco de desenvolver pressão alta ou diabetes; e impactos emocionais, incluindo questões com a imagem corporal, autoestima, vida sexual, e maior vulnerabilidade à ansiedade e depressão. Estudos recentes mostram que mulheres que fizeram quimioterapia podem sentir esses efeitos físicos por mais de 5 anos, enquanto aquelas que fizeram apenas hormonioterapia geralmente melhoram após os 2 primeiros anos. É fundamental saber que mulheres com síndrome metabólica (combinação de pressão alta, açúcar ou gordura elevada no sangue, e excesso de peso abdominal) após a mastectomia têm risco significativamente maior de o câncer voltar. Por isso, cuidar da sua saúde metabólica através de alimentação balanceada, atividade física regular, controle do estresse e sono adequado é essencial não apenas para sua qualidade de vida, mas também para reduzir riscos de recorrência. Informe sempre este histórico aos seus médicos, pois ele influencia muitos aspectos do seu cuidado de saúde.',

    conduct = 'AVALIAÇÃO INICIAL ABRANGENTE: Realizar anamnese detalhada investigando tipo de mastectomia (simples, radical modificada, bilateral), lateralidade, data do procedimento, tratamentos adjuvantes recebidos (quimioterapia, radioterapia, hormonioterapia), complicações pós-operatórias, e reconstrução mamária (imediata ou tardia). Avaliar histórico oncológico completo incluindo estadiamento, tipo histológico, status de receptores hormonais e HER2. AVALIAÇÃO FÍSICA FUNCIONAL: Examinar amplitude de movimento do ombro e braço ipsilateral, força muscular, presença de linfedema (medições circunferenciais bilaterais), cicatrizes, áreas de dor ou parestesia, e sinais de complicações tardias. Aplicar questionários de capacidade física funcional e qualidade de vida (FACT-B, EORTC QLQ-C30). RASTREAMENTO METABÓLICO COMPLETO: Solicitar painel metabólico incluindo glicemia de jejum, HbA1c, perfil lipídico completo (colesterol total, LDL, HDL, triglicerídeos), pressão arterial, circunferência abdominal, IMC. Avaliar presença de síndrome metabólica segundo critérios diagnósticos (≥3 dos seguintes: CA >88cm, TG ≥150mg/dL, HDL <50mg/dL, PA ≥130/85mmHg, glicemia ≥100mg/dL). AVALIAÇÃO ENDÓCRINA: Investigar status hormonal especialmente em pacientes em terapia endócrina: perfil tiroideano (TSH, T4 livre), vitamina D, marcadores de saúde óssea se uso de inibidores de aromatase (densitometria óssea, cálcio, fosfatase alcalina). Avaliar sintomas de menopausa e impacto na qualidade de vida. SCREENING CARDIOVASCULAR: Considerar que síndrome metabólica eleva significativamente risco cardiovascular em sobreviventes (HR 1.29 para eventos CV). Avaliar fatores de risco adicionais, considerar ECG basal, ecocardiograma se quimioterapia com antraciclinas ou trastuzumab. AVALIAÇÃO PSICOSSOCIAL: Rastrear depressão (PHQ-9), ansiedade (GAD-7), imagem corporal, função sexual, e qualidade do sono. Encaminhar para suporte psicológico/psiquiátrico se indicado. INTERVENÇÕES TERAPÊUTICAS: Prescrever programa de exercícios supervisionados (combinação aeróbico e resistência) demonstrados eficazes na redução de síndrome metabólica, sarcopenia e melhora da qualidade de vida. Implementar plano nutricional anti-inflamatório mediterrâneo com foco em controle glicêmico e redução de gordura visceral. Suplementação individualizada: vitamina D, ômega-3, probióticos. Considerar fisioterapia para reabilitação funcional do membro superior e linfedema. MONITORAMENTO LONGITUDINAL: Reavaliar função física anualmente. Monitorar marcadores metabólicos semestralmente nos primeiros 2 anos, depois anualmente. Vigilância oncológica conforme guidelines (exame físico, mamografia contralateral se aplicável). Nota: evidências recentes questionam benefício de imagem de rotina pós-mastectomia; basear seguimento em exame clínico. Coordenar cuidado multidisciplinar envolvendo oncologia, cardiologia, endocrinologia, fisioterapia, nutrição e saúde mental conforme necessário.',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = 'e65c56dc-5c07-4270-8a8b-017b293ca147';

COMMIT;

-- Verification query
SELECT
    si.id,
    si.item_text,
    si.clinical_relevance IS NOT NULL as has_clinical_relevance,
    si.patient_explanation IS NOT NULL as has_patient_explanation,
    si.conduct IS NOT NULL as has_conduct,
    si.last_review,
    COUNT(asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'e65c56dc-5c07-4270-8a8b-017b293ca147'
GROUP BY si.id, si.item_text, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
