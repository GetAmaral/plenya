-- =================================================================
-- ENRIQUECIMENTO: Histórico Familiar de Doenças - Batch 1
-- Total: 3 items
-- Data: 2026-01-27
-- Baseado em: Literatura científica, Medicina Funcional Integrativa
-- =================================================================

-- =================================================================
-- ITEM 1: Perguntar ativamente sobre principais doenças crônicas
-- ID: 56fd5913-4b41-4d56-976a-b6339b4482a6
-- Subgrupo: Parentes próximos (pais, irmãos, tios, avós, filhos, netos)
-- =================================================================

UPDATE score_items
SET
  clinical_relevance = 'A investigação sistemática de doenças crônicas em parentes de primeiro e segundo grau é fundamental para estratificação de risco cardiovascular, metabólico, oncológico e neurodegenerativo. Estudos demonstram que o histórico familiar de diabetes tipo 2 aumenta o risco em 2-6 vezes, enquanto histórico de doença cardiovascular prematura (homens <55 anos, mulheres <65 anos) eleva o risco de eventos coronarianos em 50-100%. A agregação familiar de hipertensão arterial, dislipidemia, obesidade, câncer de mama/cólon/próstata, doenças autoimunes e neurodegenerativas permite identificar padrões hereditários que justificam rastreamento precoce e intervenções preventivas personalizadas. A medicina funcional integrativa utiliza essa informação para modular fatores epigenéticos através de nutrição antiinflamatória, suplementação direcionada, exercício físico individualizado e manejo do estresse crônico. O mapeamento detalhado deve incluir idade de diagnóstico, gravidade, comorbidades e resposta terapêutica dos familiares, permitindo estimar penetrância e expressividade variável de predisposições genéticas.',

  patient_explanation = 'Conhecer as doenças que seus pais, irmãos, avós, tios e filhos tiveram é como ter um mapa do seu próprio futuro de saúde. Se várias pessoas da sua família desenvolveram diabetes, pressão alta, problemas cardíacos ou câncer, você pode ter herdado uma tendência maior para essas condições. Isso não significa que você obrigatoriamente terá essas doenças, mas sim que precisa de cuidados preventivos especiais. Por exemplo, se seu pai teve infarto aos 50 anos, você deve começar a cuidar do coração desde jovem, com alimentação adequada, exercícios e exames regulares. A medicina funcional usa essa informação para criar um plano personalizado que pode "desligar" genes ruins através de hábitos saudáveis, evitando que doenças familiares se manifestem em você.',

  conduct = 'Realizar anamnese familiar estruturada investigando: doenças cardiovasculares (IAM, AVC, IC), metabólicas (DM2, obesidade, dislipidemia), oncológicas (mama, cólon, próstata, pulmão), autoimunes (tireoidite, AR, LES), neurodegenerativas (Alzheimer, Parkinson) e psiquiátricas (depressão, ansiedade). Registrar idade de início, gravidade e desfechos. Aplicar escores de risco familiar validados (FINDRISC para diabetes, escore de Framingham modificado para DCV). Implementar rastreamento precoce conforme diretrizes (ex: colonoscopia aos 40 anos se familiar com câncer colorretal <60 anos). Prescrever intervenções preventivas personalizadas baseadas em nutrigenômica, modulação inflamatória e otimização metabólica.',

  updated_at = NOW()
WHERE id = '56fd5913-4b41-4d56-976a-b6339b4482a6';

-- Linkar artigos científicos ao Item 1
INSERT INTO article_score_items (article_id, score_item_id, created_at)
VALUES
  ('e17f93ee-ed1d-4a2c-9be6-a58db82397a9', '56fd5913-4b41-4d56-976a-b6339b4482a6', NOW()),
  ('6c333566-1f8f-4e60-bbf5-5099f786a869', '56fd5913-4b41-4d56-976a-b6339b4482a6', NOW()),
  ('3957a6ce-dfd4-4027-b5a2-a46391cd6827', '56fd5913-4b41-4d56-976a-b6339b4482a6', NOW()),
  ('c8e0ccf9-0bf8-4815-b74b-585b3c0ab799', '56fd5913-4b41-4d56-976a-b6339b4482a6', NOW())
ON CONFLICT DO NOTHING;


-- =================================================================
-- ITEM 2: Perguntar sobre quaisquer outras doenças que possam ter correlação hereditária
-- ID: 4dc130ae-9c84-4f5f-9813-561389776254
-- Subgrupo: Parentes próximos (pais, irmãos, tios, avós, filhos, netos)
-- =================================================================

UPDATE score_items
SET
  clinical_relevance = 'Além das doenças crônicas clássicas, inúmeras condições apresentam componente hereditário significativo que impacta a prática clínica. Doenças autoimunes (tireoidite de Hashimoto, doença celíaca, vitiligo, psoríase) compartilham susceptibilidade genética com penetrância familiar de 20-40%. Transtornos psiquiátricos como depressão maior, transtorno bipolar e esquizofrenia apresentam herdabilidade de 40-80%, justificando vigilância e intervenções precoces em descendentes. Condições oftalmológicas (glaucoma, degeneração macular), auditivas (perda auditiva hereditária), renais (doença policística, nefropatias hereditárias), hematológicas (trombofilias, anemias hereditárias) e dermatológicas (melanoma, dermatite atópica) requerem investigação dirigida. A medicina funcional valoriza também padrões de disbiose intestinal familiar, intolerâncias alimentares, déficits enzimáticos (MTHFR, DAO, lactase) e variantes genéticas que afetam detoxificação (GST, NAT2, CYP450). A identificação precoce permite modulação epigenética através de nutrientes específicos, probióticos personalizados e redução de exposição a gatilhos ambientais.',

  patient_explanation = 'Muitas doenças menos conhecidas também "passam de pai para filho". Se alguém na sua família tem problemas de tireoide, alergias graves, psoríase, depressão, glaucoma ou até intolerância ao glúten, você pode herdar essa tendência. Mesmo condições que parecem não relacionadas podem ter ligação genética. Por exemplo, se sua mãe tem hipotireoidismo e sua avó tinha vitiligo, pode haver uma predisposição autoimune na família que precisa ser monitorada. Conhecer essas informações permite que você tome cuidados específicos desde cedo. A medicina funcional consegue identificar essas tendências e criar um plano para "proteger" seu corpo antes que os sintomas apareçam, usando alimentação especial, suplementos naturais e mudanças no estilo de vida que compensam essas fragilidades genéticas.',

  conduct = 'Expandir investigação familiar para: doenças autoimunes (função tireoidiana, anticorpos anti-TPO, anti-transglutaminase), transtornos psiquiátricos (avaliar PANSS, HAM-D em familiares), condições oftalmológicas (tonometria se história de glaucoma), trombofilias (pesquisar fator V Leiden, protrombina G20210A se história de TEV em <50 anos). Investigar padrões de hipersensibilidade alimentar familiar (solicitar IgG alimentar, teste respiratório lactose/frutose). Realizar genotipagem de polimorfismos relevantes (MTHFR C677T, COMT, VDR, GST) quando história sugestiva. Implementar nutrição personalizada baseada em farmacogenômica e necessidades individuais de metilação, detoxificação e modulação imune.',

  updated_at = NOW()
WHERE id = '4dc130ae-9c84-4f5f-9813-561389776254';

-- Linkar artigos científicos ao Item 2
INSERT INTO article_score_items (article_id, score_item_id, created_at)
VALUES
  ('c739b1e5-9e4e-4a4f-bc54-23b753a388bd', '4dc130ae-9c84-4f5f-9813-561389776254', NOW()),
  ('6c333566-1f8f-4e60-bbf5-5099f786a869', '4dc130ae-9c84-4f5f-9813-561389776254', NOW()),
  ('e17f93ee-ed1d-4a2c-9be6-a58db82397a9', '4dc130ae-9c84-4f5f-9813-561389776254', NOW())
ON CONFLICT DO NOTHING;


-- =================================================================
-- ITEM 3: Registrar se houver alguma doença familiar importante
-- ID: 6a95547f-bb0c-4093-ad0d-3e4b7709e99f
-- Subgrupo: Parentes mais distantes
-- =================================================================

UPDATE score_items
SET
  clinical_relevance = 'A investigação de parentes de terceiro grau (primos, tios-avós, bisavós) complementa a avaliação de risco genético, especialmente em famílias pequenas ou com história de consanguinidade. Embora o risco relativo seja menor que parentes próximos (odds ratio 1.2-2.0 vs 2.0-6.0), a agregação de múltiplos casos na família extensa pode indicar síndromes de predisposição hereditária que requerem aconselhamento genético formal. Condições raras como síndrome de Lynch (câncer colorretal hereditário não-polipóide), HBOC (câncer de mama e ovário hereditário), polipose adenomatosa familiar e hipercolesterolemia familiar homozigótica apresentam padrões de herança autossômica dominante com expressividade variável que podem ser identificados através da história familiar estendida. Em medicina funcional, padrões de longevidade excepcional (centenários) em parentes distantes sugerem perfil genético protetor (polimorfismos FOXO3, APOE e2, CETP) que podem ser potencializados através de intervenções epigenéticas baseadas em restrição calórica, jejum intermitente, exercício de alta intensidade e compostos miméticos de restrição calórica (resveratrol, metformina, rapamicina).',

  patient_explanation = 'Mesmo doenças em parentes mais distantes (primos, tios distantes, bisavós) podem dar pistas importantes sobre sua saúde. Se vários membros da família extensa tiveram o mesmo tipo de câncer, por exemplo, pode haver uma "síndrome genética" que precisa ser investigada com testes específicos. Por outro lado, se você tem muitos parentes que viveram mais de 90-100 anos com saúde, isso é um sinal positivo de que sua família tem genes "bons" para longevidade que podem ser aproveitados. A medicina funcional usa essas informações para entender melhor seu potencial genético, tanto os riscos quanto as vantagens. Com base nisso, podemos criar estratégias personalizadas que imitam os hábitos dos seus antepassados longevos e protegem você das doenças que apareceram em outros ramos da família.',

  conduct = 'Documentar história familiar estendida em genograma detalhado (mínimo 3 gerações). Identificar padrões sugestivos de herança mendeliana: transmissão vertical, antecipação, segregação autossômica dominante/recessiva. Encaminhar para aconselhamento genético se ≥3 casos de mesmo câncer em parentes de qualquer grau, ou ≥2 casos em parentes de 1º grau. Investigar longevidade familiar excepcional através de questionário de saúde dos centenários e biomarcadores de envelhecimento (telômeros, metilação do DNA, AGEs). Implementar estratégias de modulação epigenética favorável: dieta mediterrânea/Okinawa, jejum 16:8, HIIT 3x/semana, suplementação senolítica (quercetina, fisetina) se história de longevidade. Se história de doenças graves, intensificar rastreamento preventivo e biomarcadores precoces.',

  updated_at = NOW()
WHERE id = '6a95547f-bb0c-4093-ad0d-3e4b7709e99f';

-- Linkar artigos científicos ao Item 3
INSERT INTO article_score_items (article_id, score_item_id, created_at)
VALUES
  ('e17f93ee-ed1d-4a2c-9be6-a58db82397a9', '6a95547f-bb0c-4093-ad0d-3e4b7709e99f', NOW()),
  ('c8e0ccf9-0bf8-4815-b74b-585b3c0ab799', '6a95547f-bb0c-4093-ad0d-3e4b7709e99f', NOW()),
  ('6c333566-1f8f-4e60-bbf5-5099f786a869', '6a95547f-bb0c-4093-ad0d-3e4b7709e99f', NOW())
ON CONFLICT DO NOTHING;


-- =================================================================
-- VERIFICAÇÃO DE SUCESSO
-- =================================================================

-- Verificar items atualizados
SELECT
  id,
  name,
  subgroup,
  CASE
    WHEN clinical_relevance IS NOT NULL THEN '✓'
    ELSE '✗'
  END as has_clinical,
  CASE
    WHEN patient_explanation IS NOT NULL THEN '✓'
    ELSE '✗'
  END as has_patient,
  CASE
    WHEN conduct IS NOT NULL THEN '✓'
    ELSE '✗'
  END as has_conduct,
  LENGTH(clinical_relevance) as clinical_length,
  LENGTH(patient_explanation) as patient_length,
  LENGTH(conduct) as conduct_length
FROM score_items
WHERE id IN (
  '56fd5913-4b41-4d56-976a-b6339b4482a6',
  '4dc130ae-9c84-4f5f-9813-561389776254',
  '6a95547f-bb0c-4093-ad0d-3e4b7709e99f'
);

-- Verificar artigos linkados
SELECT
  si.name as item_name,
  COUNT(asi.article_id) as articles_linked
FROM score_items si
LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
WHERE si.id IN (
  '56fd5913-4b41-4d56-976a-b6339b4482a6',
  '4dc130ae-9c84-4f5f-9813-561389776254',
  '6a95547f-bb0c-4093-ad0d-3e4b7709e99f'
)
GROUP BY si.name
ORDER BY si.name;
