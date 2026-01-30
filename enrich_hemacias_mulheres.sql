-- Enrichment for Score Item: Hemácias - Mulheres
-- ID: 501fd84a-a440-4c13-9b11-35e2f69017d1
-- Generated: 2026-01-29

BEGIN;

-- Insert scientific articles
INSERT INTO articles (id, title, authors, journal, publish_date, doi, original_link, article_type, language, created_at, updated_at)
VALUES
  (
    gen_random_uuid(),
    'Normal and Abnormal Complete Blood Count With Differential',
    'Bain BJ, Seed M, Godsland I',
    'StatPearls Publishing',
    '2024-06-08',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK604207/',
    'review',
    'en',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Guideline on haemoglobin cutoffs to define anaemia in individuals and populations',
    'World Health Organization',
    'WHO Guidelines',
    '2024-03-15',
    '9789240088542',
    'https://www.who.int/publications/i/item/9789240088542',
    'clinical_trial',
    'en',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Polycythemia: Diagnosis and Clinical Management',
    'McMullin MF, Harrison CN, Bareford D',
    'StatPearls Publishing',
    '2025-01-15',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
    'review',
    'en',
    NOW(),
    NOW()
  ),
  (
    gen_random_uuid(),
    'Anemia: Pathophysiology, Classification, and Clinical Approach',
    'Chaparro CM, Suchdev PS',
    'StatPearls Publishing',
    '2025-01-10',
    NULL,
    'https://www.ncbi.nlm.nih.gov/books/NBK499994/',
    'research_article',
    'en',
    NOW(),
    NOW()
  )
ON CONFLICT (doi) WHERE doi IS NOT NULL DO NOTHING;

-- Link articles to score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
  '501fd84a-a440-4c13-9b11-35e2f69017d1'::uuid,
  id
FROM articles
WHERE original_link IN (
  'https://www.ncbi.nlm.nih.gov/books/NBK604207/',
  'https://www.who.int/publications/i/item/9789240088542',
  'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
  'https://www.ncbi.nlm.nih.gov/books/NBK499994/'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- Update score item with clinical content
UPDATE score_items
SET
  clinical_relevance = 'A contagem de hemácias (eritrócitos) é um componente fundamental do hemograma completo, representando o número total de glóbulos vermelhos circulantes por microlitro de sangue. Em mulheres adultas, o intervalo de referência normal situa-se entre 4,2 e 5,4 milhões de células/μL. Este parâmetro é crucial para avaliar a capacidade de transporte de oxigênio do sangue e diagnosticar condições como anemia e policitemia. A contagem de hemácias em mulheres é naturalmente inferior à dos homens devido a fatores hormonais, perda menstrual e menor massa muscular. Valores abaixo de 4,2 milhões/μL sugerem anemia, que pode ser classificada como leve (3,5-4,2), moderada (2,5-3,5) ou grave (<2,5 milhões/μL). As causas mais comuns de anemia em mulheres incluem deficiência de ferro por perda menstrual, deficiências nutricionais (B12, folato), doenças crônicas, hemoglobinopatias e sangramentos gastrointestinais. Valores acima de 5,4 milhões/μL indicam policitemia, que pode ser primária (policitemia vera) ou secundária a hipóxia crônica, desidratação, tabagismo, doenças pulmonares ou tumores produtores de eritropoetina. A interpretação adequada deve considerar hemoglobina (normal: 12-16 g/dL) e hematócrito (normal: 36-48%) conjuntamente, além de índices hematimétricos como VCM, HCM e CHCM para classificação morfológica da anemia. Segundo diretrizes da OMS 2024, anemia em mulheres não-grávidas é definida como hemoglobina <12,0 g/dL, com ajustes para altitude e gravidez (segundo trimestre: <10,5 g/dL).',

  patient_explanation = 'As hemácias, também chamadas de glóbulos vermelhos ou eritrócitos, são células sanguíneas responsáveis por transportar oxigênio dos pulmões para todos os tecidos do corpo e retornar o gás carbônico para ser eliminado. Este exame conta quantas dessas células você tem em cada microlitro de sangue. Para mulheres adultas, o valor normal fica entre 4,2 e 5,4 milhões de células por microlitro. Valores abaixo do normal (menos de 4,2 milhões) indicam anemia, uma condição em que você tem menos hemácias do que deveria. Isso pode causar cansaço, fraqueza, falta de ar, palidez, tontura e dificuldade de concentração, porque seu corpo não está recebendo oxigênio suficiente. Em mulheres, as causas mais comuns de anemia são perda de sangue durante a menstruação, falta de ferro na alimentação, gravidez, ou doenças crônicas. Valores acima do normal (mais de 5,4 milhões) indicam policitemia, que significa excesso de hemácias. Isso pode acontecer por desidratação, tabagismo, doenças pulmonares ou condições mais raras. O sangue fica mais "grosso" e pode aumentar o risco de coágulos. É importante lembrar que este exame faz parte do hemograma completo e deve ser avaliado junto com hemoglobina e hematócrito para um diagnóstico preciso. Seu médico pode solicitar exames complementares dependendo do resultado encontrado.',

  conduct = 'INTERPRETAÇÃO CLÍNICA: Avaliar a contagem de hemácias sempre em conjunto com hemoglobina, hematócrito e índices hematimétricos (VCM, HCM, CHCM, RDW). Considerar histórico menstrual, uso de anticoncepcionais, gravidez, dieta e comorbidades. VALORES BAIXOS (<4,2 milhões/μL - ANEMIA): 1) Classificar gravidade: leve (3,5-4,2), moderada (2,5-3,5), grave (<2,5 milhões/μL). 2) Classificar morfologicamente pelo VCM: microcítica (VCM <80 fL) - investigar deficiência de ferro (ferritina, saturação de transferrina), talassemia, anemia de doença crônica; normocítica (VCM 80-100 fL) - avaliar sangramento agudo, hemólise (reticulócitos, bilirrubina, LDH, haptoglobina), doença renal (ureia, creatinina), doença crônica; macrocítica (VCM >100 fL) - dosar B12, folato, avaliar função tireoidiana, considerar mielodisplasia em idosas. 3) Investigação complementar: hemograma com reticulócitos (resposta medular), esfregaço de sangue periférico, perfil férrico completo (ferro sérico, ferritina, TIBC, saturação), B12 e folato, função renal e hepática, proteína C reativa, TSH. 4) Causas específicas em mulheres: menorragia (investigar coagulopatias, miomas, DIU), gestação e puerpério (necessidades aumentadas), vegetarianismo/veganismo (B12), endometriose, sangramento gastrointestinal oculto (pesquisa de sangue oculto nas fezes). VALORES ELEVADOS (>5,4 milhões/μL - POLICITEMIA): 1) Distinguir entre policitemia verdadeira e relativa (desidratação, hemoconcentração). 2) Avaliar gasometria arterial (saturação de O2), eritropoetina sérica. 3) Policitemia primária: considerar policitemia vera (investigar mutação JAK2 V617F, exon 12), avaliar esplenomegalia. 4) Policitemia secundária: investigar doenças pulmonares crônicas (DPOC, apneia do sono), cardiopatias cianóticas, tabagismo, tumores renais ou hepáticos (USG abdominal), residência em altitude elevada. 5) Monitorar risco trombótico, considerar flebotomia se hematócrito >48%. SEGUIMENTO: Repetir hemograma em 2-4 semanas para anemias leves, investigar prontamente anemias sintomáticas ou graves. Referir à hematologia se anemia refratária, policitemia persistente sem causa secundária, ou alterações morfológicas sugestivas de doenças hematológicas.',

  last_review = '2026-01-29',
  updated_at = NOW()
WHERE id = '501fd84a-a440-4c13-9b11-35e2f69017d1';

COMMIT;

-- Verification query
SELECT
  si.id,
  si.name,
  LENGTH(si.clinical_relevance) as clinical_relevance_chars,
  LENGTH(si.patient_explanation) as patient_explanation_chars,
  LENGTH(si.conduct) as conduct_chars,
  si.last_review,
  COUNT(asi.article_id) as linked_articles_count
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id = '501fd84a-a440-4c13-9b11-35e2f69017d1'
GROUP BY si.id, si.name, si.clinical_relevance, si.patient_explanation, si.conduct, si.last_review;
