-- Enrichment for: Testosterona Total - Mulheres Pré-Menopausa
-- Score Item ID: b6944c49-1b25-45c6-bad3-8dff164c977a
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, article_type, pm_id, doi)
VALUES
-- Article 1: 2023 International PCOS Guideline (Most comprehensive, recent)
(
    gen_random_uuid(),
    'Recommendations from the 2023 international evidence-based guideline for the assessment and management of polycystic ovary syndrome',
    'Teede HJ, Tay CT, Laven JJE, Dokras A, Moran LJ, Piltonen TT, Costello MF, Boivin J, Redman LM, Boyle JA, Norman RJ, Mousa A, Joham AE; International PCOS Network',
    'European Journal of Endocrinology',
    '2023-08-02'::date,
    'review',
    '37580861',
    '10.1093/ejendo/lvad096'
),

-- Article 2: Endocrine Society Clinical Practice Guideline on Hirsutism (Gold standard for hyperandrogenism evaluation)
(
    gen_random_uuid(),
    'Evaluation and Treatment of Hirsutism in Premenopausal Women: An Endocrine Society Clinical Practice Guideline',
    'Martin KA, Anderson RR, Chang RJ, Ehrmann DA, Lobo RA, Murad MH, Pugeat MM, Rosenfield RL',
    'The Journal of Clinical Endocrinology & Metabolism',
    '2018-04-01'::date,
    'clinical_trial',
    '29522147',
    '10.1210/jc.2018-00241'
),

-- Article 3: Hirsutism, Normal Androgens and PCOS Diagnosis (Important diagnostic challenge)
(
    gen_random_uuid(),
    'Hirsutism, Normal Androgens and Diagnosis of PCOS',
    'Spritzer PM, Marchesan LB, Santos BR, Fighera TM',
    'Diagnostics',
    '2022-08-09'::date,
    'review',
    '36010272',
    '10.3390/diagnostics12081922'
),

-- Article 4: Hyperandrogenism and Metabolic Syndrome in PCOS (microRNA biomarkers)
(
    gen_random_uuid(),
    'Hyperandrogenism and Metabolic Syndrome Are Associated With Changes in Serum-Derived microRNAs in Women With Polycystic Ovary Syndrome',
    'Murri M, Insenser M, Fernández-Durán E, San-Millán JL, Luque-Ramírez M, Escobar-Morreale HF',
    'Scientific Reports',
    '2019-11-18'::date,
    'research_article',
    '31737638',
    '10.1038/s41598-019-53848-3'
);

-- Link articles to score item using article_score_items junction table
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'b6944c49-1b25-45c6-bad3-8dff164c977a'::uuid
FROM articles a
WHERE a.pm_id IN ('37580861', '29522147', '36010272', '31737638');

-- Update score item with comprehensive clinical content
UPDATE score_items
SET
    clinical_relevance = 'A testosterona total é um marcador fundamental na avaliação do hiperandrogenismo em mulheres pré-menopáusicas, sendo a investigação hormonal de primeira linha em casos de irregularidade menstrual, hirsutismo e infertilidade. Valores de referência para mulheres pré-menopáusicas variam entre 10-55 ng/dL (0,35-1,9 nmol/L), embora métodos específicos por espectrometria de massa ofereçam maior precisão diagnóstica. A Síndrome dos Ovários Policísticos (SOP), presente em 10-13% das mulheres em idade reprodutiva, é a causa mais comum de elevação da testosterona, com valores acima de 70 ng/dL indicando hiperandrogenismo significativo. As diretrizes internacionais de 2023 reforçam que o diagnóstico de SOP pelos critérios de Rotterdam requer 2 de 3 características: disfunção ovulatória, morfologia ovariana policística (detectável por ultrassonografia ou AMH elevado), e evidências clínicas ou bioquímicas de hiperandrogenismo. Importante ressaltar que até 20% das pacientes com SOP não apresentam elevação detectável da testosterona total, sendo a testosterona livre o teste mais sensível (70% de anormalidade vs. 33% para testosterona total). O hiperandrogenismo bioquímico, representado por testosterona livre elevada (≥0,034 nmol/L), correlaciona-se com maior resistência insulínica, maior prevalência de síndrome metabólica e risco cardiometabólico aumentado. Estudos recentes demonstram que níveis elevados de testosterona (>3,0 nmol/L) associam-se a maior risco de doença hepática gordurosa metabólica (MASLD). A avaliação adequada requer dosagem matinal em fase folicular precoce, sendo que estudos com espectrometria de massa mostram que os níveis não variam significativamente ao longo do ciclo menstrual, simplificando a coleta.',

    patient_explanation = 'A testosterona é um hormônio naturalmente produzido pelos ovários e glândulas suprarrenais em quantidades pequenas, mas importantes para sua saúde. Nas mulheres antes da menopausa, níveis normais ficam entre 10-55 ng/dL. Quando a testosterona está elevada (acima de 70 ng/dL), pode indicar Síndrome dos Ovários Policísticos (SOP), condição que afeta 1 em cada 8 mulheres em idade reprodutiva. Sintomas de testosterona elevada incluem: menstruação irregular ou ausente, crescimento excessivo de pelos no rosto e corpo (hirsutismo), acne persistente, queda de cabelo no couro cabeludo, dificuldade para engravidar e ganho de peso na região abdominal. É importante entender que você pode ter SOP mesmo com testosterona normal no exame - cerca de 20% das mulheres com a síndrome apresentam níveis hormonais dentro da faixa de referência, mas têm os sintomas clínicos. Nesses casos, pode ser necessário dosar a testosterona livre, que é mais sensível para detectar o problema. A testosterona elevada não causa apenas sintomas estéticos - está associada a riscos importantes como resistência à insulina, diabetes tipo 2, pressão alta, colesterol alterado e doença hepática gordurosa. O diagnóstico e tratamento precoces ajudam a prevenir essas complicações a longo prazo. O exame deve ser coletado pela manhã, preferencialmente nos primeiros dias do ciclo menstrual, embora estudos recentes mostrem que o momento da coleta tem menos impacto do que se pensava anteriormente.',

    conduct = 'A conduta clínica para testosterona total elevada em mulheres pré-menopáusicas deve seguir protocolos estabelecidos pelas diretrizes internacionais de 2023 e pela Endocrine Society (2018). AVALIAÇÃO INICIAL: Confirmar hiperandrogenismo com testosterona total matinal em fase folicular; se normal mas suspeita clínica persiste, dosar testosterona livre (teste mais sensível com 70% de detecção vs. 33% da total); avaliar hirsutismo objetivamente usando o escore modificado de Ferriman-Gallwey (mFG ≥6 indica hirsutismo significativo). INVESTIGAÇÃO DIAGNÓSTICA: Aplicar critérios de Rotterdam para SOP (2 de 3: disfunção ovulatória, hiperandrogenismo clínico/bioquímico, morfologia ovariana policística por USG ou AMH elevado); excluir diagnósticos diferenciais como hiperplasia adrenal congênita não-clássica (17-hidroxiprogesterona), tumores ovarianos ou adrenais (testosterona >150 ng/dL sugere tumor), síndrome de Cushing e hiperprolactinemia; solicitar perfil metabólico completo (glicemia de jejum, insulina, HOMA-IR, lipidograma, função hepática) devido à alta associação com síndrome metabólica (49-54% das pacientes com SOP apresentam obesidade). TRATAMENTO FARMACOLÓGICO: Para maioria das mulheres com hirsutismo clinicamente significativo, iniciar com contraceptivos orais combinados estrogênio-progestina como primeira linha; adicionar antiandrogênio (espironolactona 50-200 mg/dia, acetato de ciproterona, finasterida) após 6 meses se resposta suboptimal; evitar flutamida devido a hepatotoxicidade; considerar metformina 1500-2000 mg/dia se resistência insulínica documentada (HOMA-IR >2,5) ou intolerância à glicose; sensibilizadores de insulina melhoram perfil metabólico e regularizam ciclos menstruais. MODIFICAÇÕES DE ESTILO DE VIDA: Perda de peso de 5-10% melhora significativamente hiperandrogenismo e função ovulatória; dieta com restrição calórica moderada e baixo índice glicêmico; exercício aeróbico 150 min/semana mais treinamento resistido. TRATAMENTO COSMÉTICO ADJUVANTE: Métodos de remoção direta de pelos (fotodepilação, eletrólise) para benefício cosmético adicional; creme de eflornitina tópico pode reduzir crescimento de pelos faciais. MONITORAMENTO: Reavaliar testosterona, perfil metabólico e função hepática a cada 3-6 meses; rastrear diabetes tipo 2 anualmente com glicemia de jejum ou HbA1c; avaliar risco cardiovascular com pressão arterial, lipidograma e circunferência abdominal; investigar apneia do sono se sintomas presentes; oferecer suporte psicológico considerando impacto na qualidade de vida e autoestima. PLANEJAMENTO REPRODUTIVO: Se desejo de gravidez, suspender contraceptivos; considerar citrato de clomifeno ou letrozol para indução de ovulação; encaminhar para reprodução assistida se falha após 6 ciclos de indução.',

    last_review = CURRENT_TIMESTAMP
WHERE id = 'b6944c49-1b25-45c6-bad3-8dff164c977a';

COMMIT;

-- Verification queries
SELECT
    si.name,
    si.clinical_relevance IS NOT NULL as has_clinical_relevance,
    LENGTH(si.clinical_relevance) as clinical_relevance_length,
    si.patient_explanation IS NOT NULL as has_patient_explanation,
    LENGTH(si.patient_explanation) as patient_explanation_length,
    si.conduct IS NOT NULL as has_conduct,
    LENGTH(si.conduct) as conduct_length,
    si.last_review,
    COUNT(asi.article_id) as linked_articles
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = 'b6944c49-1b25-45c6-bad3-8dff164c977a'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

-- List linked articles
SELECT
    a.title,
    a.authors,
    a.journal,
    a.publish_date,
    a.article_type,
    a.pm_id,
    a.doi
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'b6944c49-1b25-45c6-bad3-8dff164c977a'
ORDER BY a.publish_date DESC;
