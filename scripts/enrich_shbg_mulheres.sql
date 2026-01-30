-- Enriquecimento do Score Item: SHBG - Mulheres
-- ID: c21ccec2-66b2-49e3-911a-8d0944eda087
-- Data: 2026-01-29

BEGIN;

-- 1. Atualizar conteúdo clínico do score item
UPDATE score_items
SET
  clinical_relevance = 'A Globulina Ligadora de Hormônios Sexuais (SHBG) é uma proteína produzida pelo fígado que se liga com alta afinidade à testosterona e estradiol, regulando a biodisponibilidade desses hormônios. Em mulheres, níveis diminuídos de SHBG (<40 nmol/L) estão fortemente associados à Síndrome dos Ovários Policísticos (SOP), resistência à insulina, síndrome metabólica e doença hepática gordurosa não alcoólica. A produção hepática de SHBG é suprimida pela insulina, tornando-a um marcador indireto sensível de resistência insulínica. Valores baixos aumentam a fração livre de testosterona, intensificando os sinais clínicos de hiperandrogenismo (hirsutismo, acne, alopecia androgênica). O Índice de Andrógenos Livres (FAI), calculado como (Testosterona Total/SHBG) × 100, é utilizado para avaliar o status androgênico, com valores >5 indicando hiperandrogenismo. Em mulheres com SOP, a SHBG serve como biomarcador precoce de distúrbios metabólicos e pode orientar decisões terapêuticas. Contraceptivos hormonais combinados aumentam significativamente os níveis de SHBG através do componente estrogênico, reduzindo andrógenos livres e melhorando manifestações clínicas. Valores elevados de SHBG (>114 nmol/L) podem ocorrer em uso de estrogênios, hipertireoidismo, anorexia nervosa ou hepatopatias, reduzindo a biodisponibilidade hormonal. A dosagem de SHBG é essencial na avaliação de irregularidades menstruais, infertilidade, hiperandrogenismo e risco cardiometabólico em mulheres.',

  patient_explanation = 'A SHBG é uma proteína do sangue que funciona como um "transporte" para hormônios sexuais, especialmente testosterona e estrogênio. Quando a SHBG está baixa, mais testosterona fica "livre" no sangue, podendo causar sintomas como excesso de pelos no rosto e corpo, acne, queda de cabelo no padrão masculino e irregularidade menstrual. Valores baixos de SHBG são comuns em mulheres com Síndrome dos Ovários Policísticos (SOP) e estão relacionados à resistência à insulina e maior risco de diabetes tipo 2. Mulheres com SHBG abaixo de 40 nmol/L frequentemente apresentam alterações metabólicas que merecem investigação e acompanhamento. O exame de SHBG é especialmente útil quando há suspeita de excesso de hormônios masculinos, mesmo que a testosterona total esteja normal, pois o que importa é a quantidade de hormônio "livre" disponível. Anticoncepcionais hormonais costumam aumentar a SHBG, ajudando a controlar os sintomas de excesso de testosterona. Valores muito altos podem ocorrer em uso prolongado de pílulas anticoncepcionais ou problemas na tireoide. O médico pode solicitar este exame junto com dosagens hormonais para investigar alterações menstruais, dificuldade para engravidar, sintomas de excesso de pelos ou acne persistente, sempre considerando seu quadro clínico completo.',

  conduct = 'SHBG baixa (<40 nmol/L): Investigar Síndrome dos Ovários Policísticos (solicitar ultrassom transvaginal, perfil androgênico completo incluindo testosterona total e livre, androstenediona, DHEA-S), avaliar resistência insulínica (glicemia de jejum, insulina basal, HOMA-IR, hemoglobina glicada, teste oral de tolerância à glicose se indicado), perfil lipídico completo, enzimas hepáticas e marcadores de esteatose hepática. Calcular Índice de Andrógenos Livres (FAI = Testosterona Total/SHBG × 100; normal <5). Investigar síndrome metabólica (circunferência abdominal, pressão arterial, triglicerídeos, HDL). Orientar modificações de estilo de vida: perda de peso de 5-10% melhora significativamente SHBG e perfil metabólico, exercício físico regular, dieta com baixo índice glicêmico. Considerar metformina em presença de resistência insulínica (500-2000 mg/dia). Para controle de hiperandrogenismo e regularização menstrual em mulheres que não desejam engravidar: anticoncepcional hormonal combinado (etinilestradiol 20-35 mcg + progestágeno), preferencialmente com progestágenos de baixo potencial androgênico (drospirenona, ciproterona, dienogeste). Espironolactona 50-200 mg/dia pode ser associada para hirsutismo refratário. Reavaliar SHBG após 3-6 meses de tratamento. SHBG elevada (>114 nmol/L): Investigar uso de estrógenos exógenos, avaliar função tireoidiana (TSH, T4 livre), descartar anorexia nervosa ou hepatopatias. Em uso de anticoncepcionais: valores elevados são esperados e benéficos. Considerar redução da dose estrogênica se sintomas de hipoestrogenismo. Valores normais (40-114 nmol/L): Manter acompanhamento de rotina, reforçar hábitos saudáveis. Sempre interpretar SHBG no contexto clínico global, correlacionando com manifestações clínicas, perfil hormonal completo e fatores de risco cardiometabólico.',

  last_review = CURRENT_DATE
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';

-- 2. Inserir artigos científicos

-- Artigo 1: SHBG as potential drug candidate for metabolic disorders (2022)
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type
) VALUES (
  gen_random_uuid(),
  'Sex hormone binding globulin as a potential drug candidate for liver-related metabolic disorders treatment',
  'Bourebaba N, Ngo TH, Śmieszek A, Bourebaba L, Marycz K',
  'Biomedicine & Pharmacotherapy',
  '2022-09-01'::date,
  '35738176',
  '10.1016/j.biopha.2022.113261',
  'review'
) ON CONFLICT (pm_id) DO NOTHING
RETURNING id;

-- Artigo 2: Contraception for Women with PCOS (2022)
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type
) VALUES (
  gen_random_uuid(),
  'Contraception for Women with Polycystic Ovary Syndrome: Dealing with a Complex Condition',
  'Spritzer PM',
  'Revista Brasileira de Ginecologia e Obstetrícia',
  '2022-05-27'::date,
  '35623618',
  '10.1055/s-0042-1748036',
  'editorial'
) ON CONFLICT (pm_id) DO NOTHING
RETURNING id;

-- Artigo 3: SHBG as Early Biomarker in PCOS (2020)
INSERT INTO articles (
  id, title, authors, journal, publish_date, pm_id, doi, article_type
) VALUES (
  gen_random_uuid(),
  'Sex Hormone-Binding Globulin (SHBG) as an Early Biomarker and Therapeutic Target in Polycystic Ovary Syndrome',
  'Qu X, Donnelly R',
  'International Journal of Molecular Sciences',
  '2020-11-01'::date,
  '33139661',
  '10.3390/ijms21218191',
  'review'
) ON CONFLICT (pm_id) DO NOTHING
RETURNING id;

-- 3. Criar associações na tabela junction (article_score_items)

-- Associar artigo 1
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'c21ccec2-66b2-49e3-911a-8d0944eda087'::uuid
FROM articles a
WHERE a.pm_id = '35738176'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Associar artigo 2
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'c21ccec2-66b2-49e3-911a-8d0944eda087'::uuid
FROM articles a
WHERE a.pm_id = '35623618'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Associar artigo 3
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'c21ccec2-66b2-49e3-911a-8d0944eda087'::uuid
FROM articles a
WHERE a.pm_id = '33139661'
ON CONFLICT (article_id, score_item_id) DO NOTHING;

COMMIT;

-- Verificação
SELECT
  'clinical_relevance' as field,
  LENGTH(clinical_relevance) as char_count,
  CASE
    WHEN LENGTH(clinical_relevance) BETWEEN 1500 AND 2000 THEN '✓ OK'
    ELSE '✗ FORA DO RANGE'
  END as status
FROM score_items
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087'

UNION ALL

SELECT
  'patient_explanation' as field,
  LENGTH(patient_explanation) as char_count,
  CASE
    WHEN LENGTH(patient_explanation) BETWEEN 1000 AND 1500 THEN '✓ OK'
    ELSE '✗ FORA DO RANGE'
  END as status
FROM score_items
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087'

UNION ALL

SELECT
  'conduct' as field,
  LENGTH(conduct) as char_count,
  CASE
    WHEN LENGTH(conduct) BETWEEN 1500 AND 2500 THEN '✓ OK'
    ELSE '✗ FORA DO RANGE'
  END as status
FROM score_items
WHERE id = 'c21ccec2-66b2-49e3-911a-8d0944eda087';

-- Verificar artigos vinculados
SELECT
  si.name as score_item,
  a.title,
  a.journal,
  a.publish_date,
  a.pm_id,
  a.article_type
FROM score_items si
JOIN article_score_items asi ON si.id = asi.score_item_id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = 'c21ccec2-66b2-49e3-911a-8d0944eda087'
ORDER BY a.publish_date DESC;
