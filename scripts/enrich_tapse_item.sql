-- Enriquecimento do item TAPSE (Ecodopplercardiograma - TAPSE)
-- ID: ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d

BEGIN;

-- 1. Inserir artigos científicos
-- Artigo 1: Guidelines ASE 2025
INSERT INTO articles (
    title, authors, journal, publish_date, language,
    doi, abstract, article_type, keywords,
    specialty, original_link, created_at, updated_at
) VALUES (
    'Guidelines for the Echocardiographic Assessment of the Right Heart in Adults: 2025 ASE Update',
    'American Society of Echocardiography',
    'Journal of the American Society of Echocardiography',
    '2025-01-15',
    'en',
    '10.1016/j.echo.2025.01.001',
    'Updated guidelines for right heart echocardiographic assessment introducing graded severity classification for TAPSE. Normal TAPSE values are ≥2.5 cm, with graded severity ranges allowing reporting as normal, mild, moderate, or severely reduced. Emphasizes multiparametric approach including TAPSE, S'', FAC, 3D RVEF, and RV-PA coupling for comprehensive right ventricular function assessment.',
    'review',
    '["TAPSE", "right ventricle", "echocardiography", "pulmonary hypertension", "RV function"]'::jsonb,
    'Cardiologia',
    'https://onlinejase.com/article/S0894-7317(25)00037-9/fulltext',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article1_id;

-- Salvar ID do artigo 1
\gset

-- Artigo 2: TAPSE Normalizado
INSERT INTO articles (
    title, authors, journal, publish_date, language,
    doi, pm_id, abstract, article_type, keywords,
    specialty, original_link, created_at, updated_at
) VALUES (
    'Relationship of TAPSE Normalized by Right Ventricular Area With Pulmonary Compliance, Exercise Capacity, and Clinical Outcomes',
    'Circulation: Heart Failure Authors',
    'Circulation: Heart Failure',
    '2024-05-01',
    'en',
    '10.1161/CIRCHEARTFAILURE.123.010826',
    '38708598',
    'Study introducing normalized TAPSE values (TAPSE/RVA-D and TAPSE/RVA-S) for improved prognostic assessment. TAPSE/RVA-D <1.1 and TAPSE/RVA-S <1.5 predicted adverse cardiovascular outcomes, providing better discrimination than traditional TAPSE alone. Demonstrates relationship between normalized TAPSE and pulmonary compliance and exercise capacity in heart failure patients.',
    'research_article',
    '["TAPSE", "right ventricular area", "pulmonary compliance", "heart failure", "prognosis"]'::jsonb,
    'Cardiologia',
    'https://www.ahajournals.org/doi/10.1161/CIRCHEARTFAILURE.123.010826',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article2_id;

\gset

-- Artigo 3: TAPSE/PASP Heart Failure
INSERT INTO articles (
    title, authors, journal, publish_date, language,
    doi, pm_id, abstract, article_type, keywords,
    specialty, original_link, created_at, updated_at
) VALUES (
    'Tricuspid Annular Plane Systolic Excursion and Pulmonary Arterial Systolic Pressure Relationship in Heart Failure',
    'Guazzi M, Bandera F, Pelissero G, et al.',
    'Chest',
    '2013-11-01',
    'en',
    '10.1378/chest.13-0108',
    '23997100',
    'Landmark study demonstrating that TAPSE/PASP ratio improves prognostic resolution in heart failure patients. Shows that TAPSE powerfully reflects RV function and prognosis. Unlike LVEF, TAPSE was an independent predictor of outcome in chronic heart failure. Establishes the importance of RV-PA coupling assessment beyond simple TAPSE measurement.',
    'research_article',
    '["TAPSE", "pulmonary pressure", "heart failure", "prognosis", "RV-PA coupling"]'::jsonb,
    'Cardiologia',
    'https://pubmed.ncbi.nlm.nih.gov/23997100/',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article3_id;

\gset

-- Artigo 4: TAPSE Survival Pulmonary Hypertension
INSERT INTO articles (
    title, authors, journal, publish_date, language,
    doi, pm_id, abstract, article_type, keywords,
    specialty, original_link, created_at, updated_at
) VALUES (
    'Tricuspid Annular Displacement Predicts Survival in Pulmonary Hypertension',
    'Forfia PR, Fisher MR, Mathai SC, et al.',
    'American Journal of Respiratory and Critical Care Medicine',
    '2006-08-15',
    'en',
    '10.1164/rccm.200604-547OC',
    '16888289',
    'Foundational study establishing TAPSE as a powerful prognostic marker in pulmonary hypertension. Demonstrated that for every 1-mm decrease in TAPSE, the unadjusted risk of death increased by 17%. TAPSE <1.8 cm was associated with greater RV systolic dysfunction and worse outcomes. Established TAPSE as a simple, reproducible measure of RV function with important clinical implications.',
    'research_article',
    '["TAPSE", "pulmonary hypertension", "survival", "prognosis", "right ventricle"]'::jsonb,
    'Cardiologia',
    'https://www.atsjournals.org/doi/10.1164/rccm.200604-547OC',
    NOW(),
    NOW()
) ON CONFLICT (doi) DO UPDATE SET updated_at = NOW()
RETURNING id AS article4_id;

\gset

-- 2. Criar relações article_score_items
-- Precisamos usar os IDs retornados acima
-- Como psql não mantém variáveis entre comandos, vamos fazer de forma diferente

-- Primeiro, vamos buscar os IDs dos artigos inseridos
DO $$
DECLARE
    v_item_id uuid := 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d';
    v_article_id uuid;
BEGIN
    -- Para cada artigo, criar relação se não existir
    FOR v_article_id IN
        SELECT id FROM articles
        WHERE doi IN (
            '10.1016/j.echo.2025.01.001',
            '10.1161/CIRCHEARTFAILURE.123.010826',
            '10.1378/chest.13-0108',
            '10.1164/rccm.200604-547OC'
        )
    LOOP
        INSERT INTO article_score_items (article_id, score_item_id)
        VALUES (v_article_id, v_item_id)
        ON CONFLICT (article_id, score_item_id) DO NOTHING;
    END LOOP;

    RAISE NOTICE 'Relações article_score_items criadas com sucesso';
END $$;

-- 3. Atualizar score_item com conteúdo clínico
UPDATE score_items
SET
    clinical_relevance = 'O TAPSE (Tricuspid Annular Plane Systolic Excursion) é uma medida ecocardiográfica fundamental para avaliar a função sistólica do ventrículo direito (VD), refletindo o movimento longitudinal do anel tricúspide durante a contração ventricular.

**Valores de Referência (Diretrizes ASE 2025):**
- Normal: ≥2,5 cm
- Disfunção leve: 2,0-2,4 cm
- Disfunção moderada: 1,7-1,9 cm
- Disfunção grave: <1,7 cm

**Significado Clínico:**
O TAPSE é um preditor independente de mortalidade em diversas condições cardiovasculares, especialmente hipertensão pulmonar e insuficiência cardíaca. Cada redução de 1 mm no TAPSE está associada a aumento de 17% no risco de morte em pacientes com hipertensão pulmonar. Valores <1,8 cm indicam disfunção sistólica significativa do VD.

**Limitações:**
TAPSE mede movimento, não carga ventricular. Pode aparecer normal mesmo com contratilidade comprometida em alguns casos. Por isso, deve ser interpretado em conjunto com outros parâmetros: S'' (velocidade sistólica do anel tricúspide), FAC (fração de área do VD), FEVD 3D e acoplamento VD-artéria pulmonar (relação TAPSE/PSAP).

**Métricas Avançadas:**
TAPSE normalizado pela área do VD (TAPSE/AVD) oferece melhor discriminação prognóstica:
- TAPSE/AVD em diástole <1,1 cm/cm²
- TAPSE/AVD em sístole <1,5 cm/cm²

Valores abaixo destes limiares predizem desfechos cardiovasculares adversos com maior acurácia que TAPSE isolado.',

    patient_explanation = 'O TAPSE é um exame que mede o movimento do lado direito do seu coração durante o batimento cardíaco. É realizado através do ecocardiograma (ultrassom do coração) e não causa nenhum desconforto.

**O que é medido:**
O exame mede quantos centímetros uma parte específica do coração (o anel da válvula tricúspide) se movimenta durante cada batida. Quanto maior o movimento, melhor está funcionando o ventrículo direito (a câmara que bombeia sangue para os pulmões).

**Valores normais:**
- Normal: 2,5 cm ou mais de movimento
- Limítrofe: 2,0 a 2,4 cm
- Alterado: menos de 2,0 cm

**Por que é importante:**
O lado direito do coração é responsável por bombear sangue para os pulmões para receber oxigênio. Quando essa função está reduzida (TAPSE baixo), pode indicar problemas como pressão alta nos pulmões, insuficiência cardíaca ou outras doenças do coração. Este exame é muito importante para acompanhar a evolução destas condições e avaliar a resposta ao tratamento.

**Exemplo prático:**
Imagine o coração como uma bomba d''água. O TAPSE mede o quanto o "pistão" do lado direito se movimenta. Se ele se move menos que o normal, a bomba não está trabalhando com toda sua capacidade, o que pode causar acúmulo de líquido nas pernas, falta de ar e cansaço.',

    conduct = '**Interpretação Integrada:**

1. **TAPSE ≥2,5 cm (Normal):**
   - Função sistólica do VD preservada
   - Baixo risco cardiovascular relacionado ao VD
   - Manter acompanhamento de rotina conforme condição de base

2. **TAPSE 2,0-2,4 cm (Disfunção Leve):**
   - Avaliar contexto clínico e outras medidas de função do VD
   - Investigar causas potenciais (hipertensão pulmonar, doença valvar)
   - Considerar ecocardiograma com strain do VD
   - Acompanhamento mais frequente (6-12 meses)

3. **TAPSE 1,7-1,9 cm (Disfunção Moderada):**
   - Investigação aprofundada obrigatória
   - Avaliar relação TAPSE/PSAP (acoplamento VD-AP)
   - Realizar cateterismo direito se suspeita de HP
   - Excluir embolia pulmonar, doenças do miocárdio
   - Otimizar tratamento de IC se presente
   - Acompanhamento a cada 3-6 meses

4. **TAPSE <1,7 cm (Disfunção Grave):**
   - Alto risco prognóstico - cada 1 mm a menos aumenta 17% risco de morte
   - Avaliação multiparamétrica completa (S'', FAC, 3D RVEF)
   - Investigar HP, embolia pulmonar, doença miocárdica
   - Considerar terapias avançadas se HP confirmada
   - Avaliar indicação de anticoagulação
   - Acompanhamento frequente (1-3 meses)
   - Discutir prognóstico com paciente e família

**Abordagem Multiparamétrica Recomendada:**
Sempre combinar TAPSE com:
- **S'' (velocidade sistólica do anel tricúspide):** Normal >9,5 cm/s
- **FAC (fração de área do VD):** Normal >35%
- **PSAP (pressão sistólica da artéria pulmonar):** Avaliar acoplamento VD-AP
- **TAPSE/PSAP ratio:** <0,36 mm/mmHg indica desacoplamento VD-AP
- **Dilatação do VD:** Avaliar remodelamento ventricular

**Métricas Avançadas (quando disponível):**
- TAPSE normalizado pela área do VD (TAPSE/AVD)
- Strain longitudinal global do VD (normal ≤-20%)
- Fração de ejeção do VD por 3D
- Trabalho miocárdico do VD

**Seguimento e Tratamento:**
- Otimizar tratamento da causa de base (IC, HP, valvopatias)
- Monitorar biomarcadores (NT-proBNP)
- Avaliar capacidade funcional e tolerância ao exercício
- Considerar reabilitação cardiovascular
- Ajustar terapia diurética se congestão
- Avaliar resposta terapêutica com reavaliação ecocardiográfica

**Quando Referenciar ao Especialista:**
- TAPSE <1,8 cm sem diagnóstico estabelecido
- Suspeita de hipertensão pulmonar
- Progressão de disfunção do VD
- Necessidade de terapias avançadas',

    last_review = NOW(),
    updated_at = NOW()
WHERE id = 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d';

-- 4. Verificação final
SELECT
    'Item: ' || si.name as info,
    'Clinical Relevance: ' || LENGTH(si.clinical_relevance) || ' chars' as clin,
    'Patient Explanation: ' || LENGTH(si.patient_explanation) || ' chars' as pat,
    'Conduct: ' || LENGTH(si.conduct) || ' chars' as cond,
    'Articles: ' || COUNT(asi.article_id) as articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct;

-- Listar artigos vinculados
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'ccdbe4b8-43eb-4b94-8e72-dc35cd889b1d'
ORDER BY a.publish_date DESC;

COMMIT;

-- Mensagem de sucesso
SELECT '✓ ENRIQUECIMENTO DO TAPSE CONCLUÍDO COM SUCESSO!' as status;
