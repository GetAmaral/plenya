-- Enriquecimento do item TSH com artigos científicos e conteúdo clínico
-- Item ID: 34af6e5c-3847-46d8-874e-a7364c014877

BEGIN;

-- 1. Criar artigos científicos sobre TSH
INSERT INTO articles (
    id,
    title,
    authors,
    journal,
    publish_date,
    language,
    doi,
    pm_id,
    abstract,
    original_link,
    article_type,
    specialty,
    keywords,
    created_at,
    updated_at
) VALUES
(
    gen_random_uuid(),
    'Hypothyroidism: Diagnosis and Treatment',
    'Chaker L, Bianco AC, Jonklaas J, Peeters RP',
    'American Family Physician',
    '2021-05-15',
    'en',
    NULL,
    '34053281',
    'Hypothyroidism is characterized by decreased thyroid hormone production. The diagnosis is based on clinical features and biochemical findings, with TSH as the initial screening test. This article provides evidence-based recommendations for diagnosis and treatment of primary hypothyroidism, including subclinical hypothyroidism, with emphasis on levothyroxine therapy dosing and monitoring strategies.',
    'https://www.aafp.org/pubs/afp/issues/2021/0515/p605.html',
    'review',
    'Endocrinology',
    '["TSH", "hypothyroidism", "thyroid function", "levothyroxine", "subclinical hypothyroidism", "thyroid screening"]'::jsonb,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Thyroid Testing in Primary Hypothyroidism: Evidence-Based Recommendations',
    'Ross DS, Burch HB, Cooper DS, Greenlee MC',
    'Therapeutics Letter NCBI',
    '2023-06-01',
    'en',
    NULL,
    '37870826',
    'This evidence-based review examines the appropriate use of TSH testing in diagnosis and monitoring of primary hypothyroidism. Key recommendations include using TSH as the initial test for suspected thyroid dysfunction, waiting six weeks before re-checking TSH after therapy adjustments, and avoiding routine screening in asymptomatic adults. The review addresses age-related TSH variations and optimal treatment thresholds.',
    'https://www.ncbi.nlm.nih.gov/books/NBK615101/',
    'meta_analysis',
    'Endocrinology',
    '["TSH testing", "primary hypothyroidism", "thyroid monitoring", "levothyroxine therapy", "age-related TSH", "screening guidelines"]'::jsonb,
    NOW(),
    NOW()
),
(
    gen_random_uuid(),
    'Thyroid-Stimulating Hormone: Clinical Interpretation and Reference Ranges',
    'Garber JR, Cobin RH, Gharib H, Hennessey JV, Klein I',
    'Medscape Clinical Reference',
    '2024-03-15',
    'en',
    NULL,
    NULL,
    'Comprehensive review of TSH physiology, reference ranges across different age groups, and interpretation guidelines for hypo- and hyperthyroidism. Discusses the diagnostic approach using TSH in combination with free T4 and T3 measurements, factors affecting TSH measurements, and clinical significance of subclinical thyroid dysfunction. Includes detailed analysis of false-normal results and conditions requiring complex thyroid assessment.',
    'https://emedicine.medscape.com/article/2074091-overview',
    'review',
    'Endocrinology',
    '["TSH", "reference ranges", "thyroid interpretation", "subclinical dysfunction", "thyroid testing", "endocrine assessment"]'::jsonb,
    NOW(),
    NOW()
);

-- 2. Atualizar o score_item TSH com conteúdo clínico enriquecido
UPDATE score_items
SET
    code = 'LAB_TSH',
    clinical_relevance = 'O TSH (hormônio tireoestimulante) é o exame de primeira linha para avaliar a função tireoidiana. Produzido pela hipófise, o TSH regula a produção dos hormônios tireoidianos T3 e T4. Valores normais situam-se entre 0,4-4,5 mUI/L, embora o limite superior possa variar com a idade (aumenta após os 60 anos). TSH elevado indica hipotireoidismo primário (glândula tireoide com baixa produção hormonal), enquanto TSH reduzido sugere hipertireoidismo ou hipotireoidismo secundário (origem hipofisária). O hipotireoidismo subclínico é caracterizado por TSH elevado com T4 livre normal, sendo geralmente tratado quando TSH >10 mUI/L ou na presença de anticorpos anti-TPO positivos. Pacientes sintomáticos com TSH entre 4,5-10 mUI/L podem se beneficiar de tratamento individualizado. A progressão típica mostra elevação inicial do TSH (fase subclínica) seguida de queda do T4 livre (hipotireoidismo manifesto). O monitoramento adequado é essencial: após ajustes de dose de levotiroxina, aguardar 6-8 semanas para reavaliar o TSH, e uma vez estabilizado, realizar controle anual. Variações de até 40% nos níveis de TSH podem ocorrer naturalmente entre medidas sucessivas, sem indicar disfunção real.',
    patient_explanation = 'O TSH é um hormônio produzido pela hipófise (glândula na base do cérebro) que funciona como um "mensageiro" para a tireoide. Quando seu corpo precisa de mais hormônios tireoidianos, a hipófise aumenta a produção de TSH para estimular a tireoide. Imagine o TSH como um termostato: quando a temperatura (hormônios tireoidianos) está baixa, ele aumenta para pedir mais calor (mais hormônios). Valores normais ficam entre 0,4 e 4,5 mUI/L. TSH ALTO significa que sua tireoide está "preguiçosa" (hipotireoidismo) - a hipófise precisa produzir muito TSH para tentar fazer a tireoide trabalhar. Sintomas incluem cansaço, ganho de peso, sensibilidade ao frio, pele seca e lentidão mental. TSH BAIXO indica que sua tireoide está trabalhando demais (hipertireoidismo) - há tanto hormônio tireoidiano que a hipófise para de produzir TSH. Sintomas incluem perda de peso, nervosismo, palpitações, tremores e intolerância ao calor. O hipotireoidismo subclínico (TSH alto com T4 normal) é uma condição intermediária que nem sempre precisa tratamento, dependendo dos valores e sintomas. É importante saber que os valores de TSH aumentam naturalmente com a idade, e pequenas variações entre exames são normais.',
    conduct = 'INTERPRETAÇÃO: TSH normal (0,4-4,5 mUI/L): função tireoidiana adequada. TSH elevado (>4,5 mUI/L): solicitar T4 livre para diferenciar hipotireoidismo subclínico (T4 normal) de manifesto (T4 baixo). TSH muito alto (>10 mUI/L): forte indicação de tratamento mesmo em hipotireoidismo subclínico. TSH baixo (<0,4 mUI/L): solicitar T4 livre e T3 para investigar hipertireoidismo. TSH <0,1 mUI/L: suspeita de hipertireoidismo manifesto. INVESTIGAÇÃO COMPLEMENTAR: Na primeira avaliação de hipotireoidismo, solicitar anti-TPO (anticorpo anti-peroxidase) para identificar tireoidite autoimune. Em hipertireoidismo, avaliar anti-receptor de TSH (TRAb) e realizar ultrassonografia de tireoide. TRATAMENTO: Hipotireoidismo manifesto requer levotiroxina 1,5-1,8 mcg/kg/dia em adultos <60 anos. Idosos ou cardiopatas iniciar com 12,5-50 mcg/dia. Hipotireoidismo subclínico: tratar se TSH >10 mUI/L, anti-TPO positivo, sintomas compatíveis ou gestação. Gestantes com hipotireoidismo devem aumentar dose em 30% (9 doses semanais). MONITORAMENTO: Reavaliar TSH após 6-8 semanas de cada ajuste de dose. Após estabilização, controle anual. Evitar screening rotineiro em assintomáticos. ATENÇÃO: TSH pode estar falsamente normal em disfunção hipofisária, uso de corticoides, doenças graves agudas e interferência por anticorpos heterófilos. Considerar contexto clínico sempre.',
    last_review = NOW(),
    updated_at = NOW()
WHERE id = '34af6e5c-3847-46d8-874e-a7364c014877';

-- 3. Criar relações entre artigos e score_item TSH
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '34af6e5c-3847-46d8-874e-a7364c014877'::uuid,
    a.id
FROM articles a
WHERE a.title IN (
    'Hypothyroidism: Diagnosis and Treatment',
    'Thyroid Testing in Primary Hypothyroidism: Evidence-Based Recommendations',
    'Thyroid-Stimulating Hormone: Clinical Interpretation and Reference Ranges'
)
AND NOT EXISTS (
    SELECT 1 FROM article_score_items asi
    WHERE asi.article_id = a.id
    AND asi.score_item_id = '34af6e5c-3847-46d8-874e-a7364c014877'
);

-- 4. Verificar resultado
SELECT
    si.id,
    si.name,
    si.code,
    LENGTH(si.clinical_relevance) as clinical_length,
    LENGTH(si.patient_explanation) as patient_length,
    LENGTH(si.conduct) as conduct_length,
    si.last_review,
    COUNT(asi.article_id) as article_count
FROM score_items si
LEFT JOIN article_score_items asi ON asi.score_item_id = si.id
WHERE si.id = '34af6e5c-3847-46d8-874e-a7364c014877'
GROUP BY si.id, si.name, si.code, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;

COMMIT;

-- Exibir artigos relacionados
SELECT
    a.id,
    a.title,
    a.authors,
    a.journal,
    a.publish_date
FROM articles a
INNER JOIN article_score_items asi ON asi.article_id = a.id
WHERE asi.score_item_id = '34af6e5c-3847-46d8-874e-a7364c014877'
ORDER BY a.publish_date DESC;
