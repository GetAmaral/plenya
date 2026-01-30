-- Enrichment for Leptina - Homens (ID: 778a3c03-9a5c-475d-9f9f-4dd424cd9862)
-- Generated: 2026-01-29
-- Scientific articles: 2022-2025

BEGIN;

-- Step 1: Update score_item with clinical content
UPDATE score_items
SET
    clinical_relevance = 'A leptina é um hormônio adipocitário fundamental na regulação do metabolismo energético, apetite e função reprodutiva masculina. Em homens, níveis de leptina são significativamente menores que em mulheres (média de 6,45 ng/mL vs 20,92 ng/mL) devido à menor massa adiposa e influência hormonal. A leptina correlaciona-se com IMC, resistência insulínica (HOMA-IR), triglicerídeos e ácido úrico em homens. Valores de referência variam: 0,33-19,85 ng/mL (população chinesa) ou 0,25-48,8 ng/mL (população omanense). Em homens com IMC 20-25, o intervalo é 0,42-12,32 ng/mL. Hiperleptinemia (>97,5º percentil) associa-se a maior HOMA-IR, LDL-C elevado, obesidade central (circunferência abdominal >90cm) e síndrome metabólica. O ponto de corte ótimo para síndrome metabólica, obesidade e diabetes tipo 2 situa-se entre 24,1-28,9 ng/mL. Na obesidade masculina, ocorre resistência à leptina: níveis elevados não suprimem apetite nem aumentam gasto energético, criando um "defeito triplo" - transporte prejudicado pela barreira hematoencefálica, sinalização celular comprometida e resistência central/periférica. A hiperleptinemia na obesidade não aumenta testosterona devido à resistência leptinínica e desregulação do eixo hipotálamo-hipófise-gonadal (HPG). A leptina afeta diretamente células de Leydig, reduzindo produção de testosterona. Adiposidade visceral elevada causa hiperleptinemia, inflamação sistêmica de baixo grau, aumento da atividade da aromatase e hipogonadismo induzido por obesidade. Resistência à leptina também media resistência insulínica, diabetes tipo 2, hipertensão, aterotrombose e lesão miocárdica. Em modelos animais, tratamento com leptina recombinante reduziu inflamação e formação de placas ateroscleróticas. Pacientes com infarto do miocárdio apresentam níveis significativamente elevados de leptina, resistina e HOMA-IR. A deposição ectópica de gordura resulta em alterações metabólicas incluindo resistência insulínica, aterosclerose e doenças cardiovasculares, com a leptina sendo um elo crítico entre obesidade e risco cardiovascular em homens.',

    patient_explanation = 'A leptina é conhecida como o "hormônio da saciedade", produzido pelas células de gordura do seu corpo. Em homens, os níveis normais são naturalmente mais baixos que em mulheres, geralmente entre 0,42 e 12,32 ng/mL em peso saudável (IMC 20-25). Quando você come, suas células de gordura liberam leptina, que avisa seu cérebro: "estamos satisfeitos, pare de comer". Mas quando há excesso de peso, algo curioso acontece: mesmo com muita leptina circulando, seu cérebro não recebe a mensagem corretamente - isso se chama "resistência à leptina". É como se alguém estivesse gritando mas você não conseguisse ouvir. Níveis muito altos (acima de 24-28 ng/mL) geralmente indicam obesidade e podem estar associados a problemas metabólicos: resistência à insulina, diabetes tipo 2, colesterol alto e pressão alta. Em homens especificamente, excesso de leptina pode reduzir a produção de testosterona, afetando energia, libido e massa muscular. A boa notícia é que perder peso e reduzir gordura abdominal ajuda a normalizar a leptina. Exercícios físicos regulares, especialmente treino de força, e uma dieta balanceada rica em proteínas, fibras e gorduras saudáveis melhoram a sensibilidade à leptina. Evitar alimentos ultraprocessados, açúcar refinado e excesso de carboidratos simples também é fundamental. Dormir bem (7-9 horas) é crucial, pois a privação de sono aumenta a resistência à leptina. Se seus níveis estão alterados, converse com seu médico sobre estratégias personalizadas para otimizar seu metabolismo e saúde hormonal.',

    conduct = 'AVALIAÇÃO INICIAL: Solicitar leptina sérica em jejum de 8-12 horas, preferencialmente pela manhã. Coletar simultaneamente: glicemia, insulina (calcular HOMA-IR), perfil lipídico completo (CT, HDL, LDL, TG), ácido úrico, testosterona total e livre, SHBG, hemoglobina glicada. Avaliar composição corporal: IMC, circunferência abdominal, bioimpedância ou DEXA se disponível. Investigar sintomas de síndrome metabólica, hipogonadismo e resistência insulínica. LEPTINA NORMAL (0,42-12,32 ng/mL em IMC 20-25): Manter acompanhamento preventivo anual. Reforçar hábitos saudáveis: dieta mediterrânea ou DASH, atividade física regular (150min/semana aeróbico + 2x/semana resistência), sono adequado (7-9h), controle de estresse. LEPTINA ELEVADA LEVE-MODERADA (12,32-24 ng/mL): Indica início de resistência leptinínica e acúmulo adiposo. Intervenção comportamental intensiva: redução de 5-10% do peso corporal em 6 meses através de déficit calórico moderado (500kcal/dia), dieta rica em proteínas (1,6-2g/kg), fibras (>30g/dia) e pobre em carboidratos refinados. Exercício combinado: 200min/semana aeróbico + 3x/semana resistência. Suplementação: ômega-3 (2-3g/dia), vitamina D se deficiente (>30ng/mL), magnésio. Reavaliar em 3 meses. LEPTINA ELEVADA MODERADA-ALTA (24-50 ng/mL): Resistência leptinínica estabelecida com alto risco metabólico. Avaliar síndrome metabólica (≥3 critérios), diabetes, esteatose hepática (ultrassom/elastografia), apneia do sono (polissonografia se indicado), função tireoidiana. Encaminhar para endocrinologista. Meta: perda de 10-15% do peso. Considerar dieta low-carb cetogênica supervisionada, jejum intermitente 16:8 se apropriado. Farmacoterapia se IMC>30 ou IMC>27 com comorbidades: metformina 1500-2000mg/dia (melhora sensibilidade à leptina e insulina), liraglutida/semaglutida (análogos GLP-1) em casos selecionados. Se testosterona baixa (<300ng/dL) e sintomas de hipogonadismo, considerar reposição hormonal após perda de peso (frequentemente testosterona normaliza com emagrecimento). LEPTINA MUITO ELEVADA (>50 ng/mL): Obesidade severa com resistência leptinínica grave. Equipe multidisciplinar: endocrinologista, nutricionista, psicólogo, educador físico. Investigar causas secundárias de obesidade (Cushing, hipotireoidismo, medicações). Avaliar candidatura para cirurgia bariátrica se IMC>40 ou IMC>35 com comorbidades graves após tentativas conservadoras. Monitoramento cardiovascular rigoroso: ECG, ecocardiograma, teste ergométrico se apropriado. Rastreio de complicações: nefropatia, retinopatia, neuropatia. Farmacoterapia agressiva: metformina + análogo GLP-1. Suporte psicológico para mudança comportamental sustentável. MONITORAMENTO: Reavaliar leptina, HOMA-IR, perfil metabólico e testosterona a cada 3-6 meses durante fase de intervenção ativa. Após estabilização, controle semestral/anual. Ajustar condutas conforme evolução clínica e laboratorial. Atenção a fatores de risco cardiovascular e função reprodutiva masculina.',

    last_review = CURRENT_DATE
WHERE id = '778a3c03-9a5c-475d-9f9f-4dd424cd9862';

-- Step 2: Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, language, doi, pm_id, abstract, specialty, article_type, created_at, updated_at)
VALUES
    -- Article 1: Chinese population reference intervals
    (
        gen_random_uuid(),
        'Sex- and body mass index-specific reference intervals for serum leptin: a population based study in China',
        'Zhang L, Wang Y, Chen D, et al.',
        'Nutrition & Metabolism',
        '2022-08-03',
        'en',
        '10.1186/s12986-022-00689-x',
        'PMC9358897',
        'Background: Leptin is an important adipokine that regulates energy balance and metabolism. However, sex- and BMI-specific reference intervals for serum leptin are lacking in the Chinese population. Methods: This population-based study included 2,876 participants (1,432 men and 1,444 women) aged 20-74 years from a health examination center. Serum leptin was measured by ELISA. Reference intervals were established using non-parametric methods stratified by sex and BMI categories. Results: The reference interval of serum leptin was 0.33-19.85 ng/mL in men and 3.00-46.89 ng/mL in women. In men with BMI 20-25 kg/m², leptin ranged 0.42-12.32 ng/mL; BMI 25-27.5: 2.17-20.22 ng/mL. Women consistently showed higher leptin levels. Multivariate analysis showed serum leptin correlated with BMI, HOMA-IR, uric acid in women, and plus triglycerides in men. In men with BMI 20-25, participants with leptin >97.5th percentile had significantly higher HOMA-IR, LDL-C, uric acid, central obesity (WC>90cm), metabolic syndrome, and lower HDL-C. Conclusion: Sex- and BMI-specific reference intervals for serum leptin were established for the Chinese population, providing clinical guidance for metabolic risk assessment.',
        'Endocrinology',
        'research_article',
        NOW(),
        NOW()
    ),

    -- Article 2: Leptin-testosterone-obesity relationship
    (
        gen_random_uuid(),
        'The role of leptin and low testosterone in obesity',
        'Cohen PG',
        'Medical Hypotheses',
        '2022-02-15',
        'en',
        '10.1016/j.mehy.2022.110831',
        '35102263',
        'Obesity is a growing global health concern associated with multiple metabolic complications including hypogonadism in men. This review examines the bidirectional relationship between leptin and testosterone in the context of obesity. In normal physiology, leptin promotes testosterone synthesis through stimulation of the hypothalamic-pituitary-gonadal (HPG) axis. However, in obesity, elevated leptin levels do not lead to increased testosterone due to leptin resistance and dysregulated signaling. Increased adiposity leads to elevated adipokines, particularly leptin, causing disruptions in the HPG axis. Hyperleptinemia is associated with low-grade systemic inflammation and metabolic dysfunction in obese individuals. Obesity-induced hypogonadism is a complex endocrine disorder primarily driven by excessive adipose tissue acting as an endocrine organ secreting hormones and inflammatory mediators. Leptin resistance impairs the normal stimulatory effects on the HPG axis. Recent evidence supports a direct role of leptin in affecting Leydig cells, reducing testosterone production and increasing appetite. Factors induced by obesity including systemic inflammation, increased aromatase activity, and leptin production interfere with testosterone production. Weight loss interventions that reduce leptin levels often restore testosterone production. Understanding the leptin-testosterone-adiposity axis is crucial for managing obesity-related hypogonadism in men.',
        'Endocrinology',
        'review',
        NOW(),
        NOW()
    ),

    -- Article 3: Leptin resistance mechanisms 2024
    (
        gen_random_uuid(),
        'Leptin and leptin resistance in obesity: current evidence, mechanisms and future directions',
        'Liu J, Yang X, Li Y, et al.',
        'Diabetes, Metabolic Syndrome and Obesity: Targets and Therapy',
        '2024-10-15',
        'en',
        '10.2147/DMSO.S482859',
        'PMC12486228',
        'Leptin resistance - a state of attenuated biological response despite hyperleptinemia - is commonly observed in most people with obesity. Elevated leptin levels cannot effectively act on hypothalamic neurons to suppress appetite and increase energy expenditure. Unlike insulin resistance, leptin resistance may manifest as a triple defect: impaired blood-brain barrier transport, disrupted leptin signaling at the receptor level, and both central and peripheral leptin resistance. This comprehensive review examines current evidence on leptin resistance mechanisms in obesity. Key pathophysiological mechanisms include: (1) Impaired leptin transport across the blood-brain barrier via saturated transporters; (2) Downregulation of leptin receptors in hypothalamic neurons; (3) Activation of suppressor of cytokine signaling proteins (SOCS3) that inhibit leptin signaling; (4) Endoplasmic reticulum stress and inflammatory pathways in the hypothalamus; (5) Mitochondrial dysfunction affecting energy sensing. Clinical consequences of leptin resistance extend beyond failed appetite regulation to include insulin resistance, dyslipidemia, hypertension, cardiovascular disease, and reproductive dysfunction. In men, leptin resistance contributes to obesity-induced hypogonadism through disruption of the HPG axis and direct effects on testicular Leydig cells. Therapeutic strategies under investigation include leptin sensitizers, SOCS3 inhibitors, anti-inflammatory agents, and bariatric surgery which can partially reverse leptin resistance through weight loss and metabolic improvements.',
        'Endocrinology',
        'review',
        NOW(),
        NOW()
    ),

    -- Article 4: Cardiovascular implications 2025
    (
        gen_random_uuid(),
        'Role of Leptin in Obesity, Cardiovascular Disease, and Type 2 Diabetes',
        'Fernández-Ruiz VE, Paredes S, Pennanen C, et al.',
        'International Journal of Molecular Sciences',
        '2024-02-19',
        'en',
        '10.3390/ijms25042338',
        'PMC10888990',
        'Leptin, a hormone primarily secreted by adipose tissue, plays crucial roles in regulating appetite, energy expenditure, and metabolism. However, in obesity, leptin resistance develops, leading to paradoxically elevated leptin levels without appropriate biological responses. This review explores leptin''s role in obesity-related complications, particularly cardiovascular disease (CVD) and type 2 diabetes mellitus (T2DM). Plasma leptin levels predict cardiovascular risk, potentially mediated through leptin resistance-related insulin resistance, chronic inflammation, T2DM, hypertension, atherothrombosis, and myocardial injury. In animal models of lipodystrophy prone to atherosclerosis, daily recombinant leptin treatment reduced inflammation and plaque formation, indicating that leptin therapy confers both metabolic and cardiovascular benefits when leptin deficiency (not resistance) is present. In men with acute myocardial infarction, leptin has been identified as an important link between obesity and cardiovascular risk factors. Patients with MI exhibit significantly higher levels of BMI, insulin, leptin, resistin, and HOMA-IR. Increased BMI is strongly associated with metabolic disturbances including dyslipidemia, diabetes, hypertension, and systemic inflammation, which collectively elevate MI risk. The pathogenesis of diabetes and obesity share common pathways involving insulin resistance, oxidative stress, and proinflammatory patterns. Ectopic fat deposition results in metabolic alterations including insulin resistance, atherosclerosis, and CVD. These findings demonstrate that leptin plays a central role in the complex relationship between obesity, insulin resistance, and cardiovascular disease, with leptin resistance being a key mechanism underlying these metabolic complications.',
        'Cardiology',
        'review',
        NOW(),
        NOW()
    )
RETURNING id, title;

-- Step 3: Link articles to score item
-- Get the article IDs and link them
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, '778a3c03-9a5c-475d-9f9f-4dd424cd9862'::uuid
FROM articles a
WHERE a.title IN (
    'Sex- and body mass index-specific reference intervals for serum leptin: a population based study in China',
    'The role of leptin and low testosterone in obesity',
    'Leptin and leptin resistance in obesity: current evidence, mechanisms and future directions',
    'Role of Leptin in Obesity, Cardiovascular Disease, and Type 2 Diabetes'
)
AND a.created_at >= NOW() - INTERVAL '1 minute';

-- Verification query
SELECT
    si.name as score_item,
    COUNT(asi.article_id) as article_count,
    si.last_review,
    LENGTH(si.clinical_relevance) as relevance_length,
    LENGTH(si.patient_explanation) as explanation_length,
    LENGTH(si.conduct) as conduct_length
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '778a3c03-9a5c-475d-9f9f-4dd424cd9862'
GROUP BY si.id, si.name, si.last_review, si.clinical_relevance, si.patient_explanation, si.conduct;

COMMIT;

-- Display linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.doi,
    a.pm_id
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '778a3c03-9a5c-475d-9f9f-4dd424cd9862'
ORDER BY a.publish_date DESC;
