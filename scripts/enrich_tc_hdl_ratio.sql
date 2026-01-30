-- ============================================================================
-- Enrichment Script: Relação Colesterol Total/HDL (TC/HDL Ratio)
-- Score Item ID: 63d6b0a3-a52e-4994-a3b8-29f4aa1740e5
-- Generated: 2026-01-29
-- ============================================================================
-- CRITICAL: Uses correct schema - articles table, article_score_items junction
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- STEP 1: Insert Scientific Articles (check existence by pm_id first)
-- ----------------------------------------------------------------------------

-- Article 1: PMID 36589799 (2022) - TC/HDL ratio and mortality in general population
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '36589799') THEN
    INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type)
    VALUES (
      gen_random_uuid(),
      'The effect of total cholesterol/high-density lipoprotein cholesterol ratio on mortality risk in the general population',
      'Wang Y, Xu D',
      'Frontiers in Endocrinology',
      '2022-12-15'::date,
      '36589799',
      'research_article'
    );
  END IF;
END $$;

-- Article 2: PMID 31291776 (2020) - TC/HDL ratio discordance and ASCVD in ARIC study
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '31291776') THEN
    INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type)
    VALUES (
      gen_random_uuid(),
      'Total cholesterol/HDL-cholesterol ratio discordance with LDL-cholesterol and non-HDL-cholesterol and incidence of atherosclerotic cardiovascular disease in primary prevention: The ARIC study',
      'Quispe R, Elshazly MB, Zhao D, Toth PP, Puri R, Virani SS, Blumenthal RS, Martin SS, Jones SR, Michos ED',
      'European Journal of Preventive Cardiology',
      '2020-10-01'::date,
      '31291776',
      'research_article'
    );
  END IF;
END $$;

-- Article 3: PMID 35886124 (2022) - Systematic review on cholesterol and CVD death
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '35886124') THEN
    INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type)
    VALUES (
      gen_random_uuid(),
      'Serum Cholesterol Levels and Risk of Cardiovascular Death: A Systematic Review and a Dose-Response Meta-Analysis of Prospective Cohort Studies',
      'Gao M, Wang F, Shen Y, Wei Y, Wang Z, Chen Q',
      'International Journal of Environmental Research and Public Health',
      '2022-07-06'::date,
      '35886124',
      'meta_analysis'
    );
  END IF;
END $$;

-- Article 4: PMID 36900073 (2023) - TG/HDL-C ratio as risk marker (related to lipid ratios)
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM articles WHERE pm_id = '36900073') THEN
    INSERT INTO articles (id, title, authors, journal, publish_date, pm_id, article_type)
    VALUES (
      gen_random_uuid(),
      'The Triglyceride/High-Density Lipoprotein Cholesterol (TG/HDL-C) Ratio as a Risk Marker for Metabolic Syndrome and Cardiovascular Disease',
      'Salazar MR, Carballo AR, Espeche WG, Aizpurúa M, Leiva Sisnieguez CE, March CE, Balbín E, Stavile RN, Reaven GM',
      'Diagnostics',
      '2023-03-01'::date,
      '36900073',
      'review'
    );
  END IF;
END $$;

-- ----------------------------------------------------------------------------
-- STEP 2: Link Articles to Score Item (with conflict handling)
-- ----------------------------------------------------------------------------

-- Link Article 1 (PMID 36589799)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
FROM articles a
WHERE a.pm_id = '36589799'
  AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
      AND asi.score_item_id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
  );

-- Link Article 2 (PMID 31291776)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
FROM articles a
WHERE a.pm_id = '31291776'
  AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
      AND asi.score_item_id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
  );

-- Link Article 3 (PMID 35886124)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
FROM articles a
WHERE a.pm_id = '35886124'
  AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
      AND asi.score_item_id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
  );

-- Link Article 4 (PMID 36900073)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
FROM articles a
WHERE a.pm_id = '36900073'
  AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
      AND asi.score_item_id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'::uuid
  );

-- ----------------------------------------------------------------------------
-- STEP 3: Update Score Item with Clinical Content
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A relação Colesterol Total/HDL (CT/HDL) é um indicador lipídico composto que reflete o equilíbrio entre lipoproteínas aterogênicas e protetoras, sendo amplamente utilizado na estratificação de risco cardiovascular. Estudos populacionais demonstram que valores de CT/HDL superiores a 4,22 associam-se independentemente com aumento de 13% no risco de mortalidade cardiovascular (HR 1,13; IC95% 1,02-1,25), evidenciando relação não-linear onde razões muito baixas ou excessivamente altas elevam mortalidade geral. Na coorte ARIC com 14.403 participantes em prevenção primária, indivíduos com CT/HDL elevado (≥3,3) mas LDL-colesterol e não-HDL-colesterol abaixo da mediana apresentaram risco 24-29% maior de doença aterosclerótica cardiovascular incidente, sugerindo que a razão CT/HDL captura dimensões de risco não detectadas pelos marcadores lipídicos isolados. Meta-análises recentes confirmam que cada desvio-padrão de aumento nos níveis médios de CT/HDL eleva o risco de doença cardiovascular em 26-29%, superando em poder preditivo o colesterol total isolado. O índice CT/HDL integra informações sobre colesterol aterogênico total e HDL protetor, refletindo não apenas carga lipídica mas também capacidade de transporte reverso do colesterol e atividade anti-inflamatória, sendo particularmente útil em pacientes com síndrome metabólica, diabetes tipo 2 e discordância entre marcadores lipídicos convencionais.',

  patient_explanation = 'A relação CT/HDL é um cálculo simples que divide seu colesterol total pelo HDL (colesterol "bom"), fornecendo uma medida mais completa do seu risco cardiovascular do que qualquer um desses valores isoladamente. Valores ideais ficam abaixo de 3,5, indicando equilíbrio favorável entre o colesterol que pode obstruir suas artérias e o que ajuda a limpá-las. Quando a relação ultrapassa 5,0, o risco de doenças cardíacas aumenta significativamente, mesmo que seu colesterol total pareça normal. Este índice é especialmente importante porque pesquisas mostram que algumas pessoas têm LDL aparentemente adequado, mas ainda assim apresentam risco elevado devido a baixos níveis de HDL protetor. A relação CT/HDL consegue identificar essas situações de risco oculto. Valores elevados geralmente indicam necessidade de mudanças no estilo de vida (atividade física aeróbica aumenta HDL, redução de gorduras trans e açúcares refinados melhora o perfil lipídico) e, em alguns casos, tratamento medicamentoso com estatinas ou fibratos. Acompanhar periodicamente esta relação ajuda a monitorar a efetividade das intervenções e ajustar o tratamento conforme necessário para proteger sua saúde cardiovascular a longo prazo.',

  conduct = 'ESTRATIFICAÇÃO DE RISCO: CT/HDL <3,5 indica risco cardiovascular baixo; 3,5-5,0 risco moderado; >5,0 risco elevado, com valores acima de 4,22 associados especificamente a mortalidade cardiovascular aumentada. Em pacientes com discordância entre CT/HDL elevado e LDL/não-HDL baixos, considerar avaliação adicional com escore de cálcio coronariano ou apolipoproteína B para refinar estratificação. INTERVENÇÕES LIFESTYLE: Exercícios aeróbicos moderados-vigorosos 150-300 min/semana elevam HDL em 5-10% e melhoram razão CT/HDL; redução de peso em pacientes com sobrepeso (perda de 5-10% do peso corporal) aumenta HDL e reduz triglicerídeos; dieta mediterrânea com ácidos graxos ômega-3, azeite extravirgem, oleaginosas e fibras solúveis melhora perfil lipídico global; eliminação de gorduras trans e redução de açúcares refinados/carboidratos simples. FARMACOTERAPIA: Estatinas de moderada-alta intensidade são primeira linha quando CT/HDL >5,0 com risco cardiovascular calculado ≥7,5% em 10 anos ou presença de fatores agravantes (diabetes, hipertensão, tabagismo); considerar associação de ezetimiba se LDL permanecer acima da meta (<70 mg/dL em alto risco, <55 mg/dL em muito alto risco) após 4-12 semanas de estatina; fibratos (fenofibrato) podem ser adicionados em casos com triglicerídeos persistentemente elevados (>200 mg/dL) e HDL baixo (<40 mg/dL homens, <50 mg/dL mulheres), especialmente em diabéticos; ácidos graxos ômega-3 de prescrição (icosapento etila) para triglicerídeos ≥150 mg/dL em pacientes cardiovasculares de alto risco. MONITORAMENTO: Reavaliar perfil lipídico completo após 6-12 semanas de início/ajuste terapêutico; metas incluem CT/HDL <3,5 idealmente, HDL >40 mg/dL (homens) ou >50 mg/dL (mulheres), LDL conforme risco individual; em pacientes com síndrome metabólica ou resistência insulínica, considerar também avaliação de glicemia de jejum, HbA1c e marcadores inflamatórios (PCR-us). SITUAÇÕES ESPECIAIS: Em mulheres pós-menopausa, considerar que terapia de reposição hormonal pode influenciar relação CT/HDL; em idosos >75 anos, avaliar risco-benefício individualizado de tratamento intensivo; pacientes com doença renal crônica frequentemente apresentam CT/HDL elevado e requerem manejo multidisciplinar.',

  last_review = CURRENT_DATE,
  updated_at = CURRENT_TIMESTAMP

WHERE id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5';

-- ----------------------------------------------------------------------------
-- STEP 4: Verification Queries
-- ----------------------------------------------------------------------------

-- Verify articles were inserted
SELECT
  pm_id,
  LEFT(title, 80) || '...' as title_preview,
  journal,
  publish_date,
  article_type
FROM articles
WHERE pm_id IN ('36589799', '31291776', '35886124', '36900073')
ORDER BY publish_date DESC;

-- Verify article linkages
SELECT
  a.pm_id,
  a.journal,
  a.publish_date,
  si.name as score_item_name
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
JOIN score_items si ON asi.score_item_id = si.id
WHERE si.id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5'
ORDER BY a.publish_date DESC;

-- Verify clinical content character counts
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5';

-- Show updated content preview
SELECT
  name,
  LEFT(clinical_relevance, 200) || '...' as clinical_relevance_preview,
  LEFT(patient_explanation, 200) || '...' as patient_explanation_preview,
  LEFT(conduct, 200) || '...' as conduct_preview
FROM score_items
WHERE id = '63d6b0a3-a52e-4994-a3b8-29f4aa1740e5';

COMMIT;

-- ============================================================================
-- End of Enrichment Script
-- ============================================================================
