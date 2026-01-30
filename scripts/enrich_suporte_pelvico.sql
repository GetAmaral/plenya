-- ============================================================================
-- ENRIQUECIMENTO: Suporte Pélvico (ID: f54dbfaa-599a-40ee-907a-28431ce4858a)
-- ============================================================================
-- Data: 2026-01-29
-- Categoria: Composição Corporal
-- Foco: Avaliação do suporte pélvico, prolapso de órgãos pélvicos (POP),
--       sistema POP-Q, disfunção do assoalho pélvico, treinamento muscular
-- ============================================================================

BEGIN;

-- 1. ATUALIZAR CONTEÚDO CLÍNICO DO SCORE ITEM
-- ============================================================================

UPDATE score_items
SET
  clinical_relevance = 'A avaliação do suporte pélvico é fundamental na saúde da mulher, pois o prolapso de órgãos pélvicos (POP) afeta até 50% das mulheres multíparas e representa importante causa de morbidade. O sistema POP-Q (Pelvic Organ Prolapse Quantification) permanece como padrão-ouro para estadiamento clínico, mensurando pontos anatômicos específicos em relação ao hímen. Os estágios variam de 0 (ausência de prolapso) a IV (eversão vaginal completa). A avaliação identifica tipos de prolapso: cistocele (parede anterior), retocele (parede posterior) e prolapso uterino/cúpula vaginal (compartimento apical). Ressonância magnética dinâmica emerge como modalidade complementar para avaliação anatômica e funcional. A força muscular do assoalho pélvico deve ser documentada como ausente, fraca, normal ou forte. Prolapso sintomático geralmente corresponde a estágio ≥II (borda além do hímen). A International Urogynecology Consultation (2023) recomenda exame pélvico padronizado com POP-Q ou S-POP (sistema simplificado). Sintomas incluem sensação de pressão pélvica, protrusão vaginal, disfunção urinária/intestinal e desconforto sexual, impactando significativamente qualidade de vida.',

  patient_explanation = 'O suporte pélvico refere-se às estruturas (músculos, ligamentos e fáscias) que mantêm os órgãos pélvicos (bexiga, útero, reto) em suas posições normais. Quando essas estruturas enfraquecem, pode ocorrer prolapso - descida de um ou mais órgãos em direção ou através da abertura vaginal. Você pode sentir uma "bola" ou pressão na vagina, especialmente ao final do dia ou ao fazer esforço. Outros sintomas incluem dificuldade para urinar ou evacuar, incontinência urinária ou escape de fezes, e desconforto durante relações sexuais. Fatores de risco incluem partos vaginais múltiplos, idade avançada, obesidade, tosse crônica e trabalhos que exigem esforço físico intenso. A avaliação médica utiliza exame físico padronizado para medir o grau de descida dos órgãos. Muitas mulheres podem melhorar significativamente com exercícios para fortalecer o assoalho pélvico (exercícios de Kegel), perda de peso e modificação de atividades. Casos mais graves podem necessitar pessários vaginais (dispositivos de suporte) ou cirurgia corretiva.',

  conduct = 'AVALIAÇÃO INICIAL: Realize exame pélvico padronizado com sistema POP-Q, documentando pontos Aa, Ba, C, D, Ap, Bp, hiato genital, corpo perineal e comprimento vaginal total. Estadie conforme critérios: 0 (sem prolapso), I (>1cm acima hímen), II (entre -1cm e +1cm do hímen), III (>1cm além hímen, mas <2cm do comprimento vaginal total), IV (eversão completa). Avalie força muscular do assoalho pélvico (ausente/fraca/normal/forte). Investigue sintomas associados: incontinência urinária, dificuldade miccional/evacuatória, disfunção sexual. Considere ressonância magnética dinâmica em casos complexos ou pré-operatórios. CONDUTA CONSERVADORA (primeira linha para estágios I-II): Treinamento muscular do assoalho pélvico (PFMT) supervisionado por fisioterapeuta especializado - evidência nível 1A para eficácia. Protocolo: 3 séries de 8-12 contrações máximas diárias, mantendo 6-8 segundos, por mínimo 12-16 semanas. Biofeedback ou eletroestimulação podem auxiliar. Modificação de fatores de risco: perda de peso se IMC >25, tratamento de tosse crônica/constipação, evitar esforços excessivos. Pessários vaginais (anel, cubo, Gellhorn) para pacientes sintomáticas que preferem evitar cirurgia ou têm contraindicações cirúrgicas - requer seguimento a cada 3-6 meses. CONDUTA CIRÚRGICA: Considere para estágios III-IV sintomáticos ou falha conservadora. Opções incluem reparo vaginal (colporrafia anterior/posterior), sacrocolpopexia laparoscópica/robótica (padrão-ouro para prolapso apical), histerectomia vaginal com suspensão de cúpula. Evite telas sintéticas transvaginais devido complicações. SEGUIMENTO: Reavalie a cada 6-12 meses com POP-Q, ajustando conduta conforme evolução sintomática e objetiva.',

  last_review = CURRENT_DATE

WHERE id = 'f54dbfaa-599a-40ee-907a-28431ce4858a';


-- 2. INSERIR ARTIGOS CIENTÍFICOS
-- ============================================================================

-- Artigo 1: International Urogynecology Consultation (2023)
-- Guideline internacional sobre avaliação clínica de POP
INSERT INTO articles (
  pm_id,
  title,
  authors,
  journal,
  publish_date,
  doi,
  article_type
) VALUES (
  '37737436',
  'International Urogynecology consultation chapter 2 committee 3: the clinical evaluation of pelvic organ prolapse including investigations into associated morbidity/pelvic floor dysfunction',
  'Barbier H, Carberry CL, Karjalainen PK, Mahoney CK, Manríquez Galán V, Rosamilia A, Ruess E, Shaker D, Thariani K',
  'International Urogynecology Journal',
  '2023-09-22'::date,
  '10.1007/s00192-023-05629-8',
  'review'
) ON CONFLICT (doi) DO UPDATE SET
  pm_id = EXCLUDED.pm_id,
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date;

-- Artigo 2: Mecanismos do PFMT - Revisão Narrativa (2024)
-- Evidência nível 1A para eficácia do treinamento muscular
INSERT INTO articles (
  pm_id,
  title,
  authors,
  journal,
  publish_date,
  doi,
  article_type
) VALUES (
  '38979823',
  'Mechanisms for pelvic floor muscle training: Morphological changes and associations between changes in pelvic floor muscle variables and symptoms of female stress urinary incontinence and pelvic organ prolapse-A narrative review',
  'Bø K',
  'Neurourology and Urodynamics',
  '2024-07-09'::date,
  '10.1002/nau.25551',
  'review'
) ON CONFLICT (doi) DO UPDATE SET
  pm_id = EXCLUDED.pm_id,
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date;

-- Artigo 3: PFMT em Atletas - Meta-análise (2024)
-- Evidência sobre eficácia do PFMT em população especial
INSERT INTO articles (
  pm_id,
  title,
  authors,
  journal,
  publish_date,
  doi,
  article_type
) VALUES (
  '37688407',
  'Pelvic Floor Muscle Training Interventions in Female Athletes: A Systematic Review and Meta-analysis',
  'Rodríguez-Longobardo C, López-Torres O, Guadalupe-Grau A, Gómez-Ruano MÁ',
  'Sports Health',
  '2023-09-09'::date,
  '10.1177/19417381231195305',
  'meta_analysis'
) ON CONFLICT (doi) DO UPDATE SET
  pm_id = EXCLUDED.pm_id,
  title = EXCLUDED.title,
  authors = EXCLUDED.authors,
  journal = EXCLUDED.journal,
  publish_date = EXCLUDED.publish_date;


-- 3. VINCULAR ARTIGOS AO SCORE ITEM
-- ============================================================================

-- Vincular Artigo 1 (usando DOI para encontrar o UUID)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'f54dbfaa-599a-40ee-907a-28431ce4858a'
FROM articles a
WHERE a.doi = '10.1007/s00192-023-05629-8'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = 'f54dbfaa-599a-40ee-907a-28431ce4858a'
);

-- Vincular Artigo 2 (usando DOI para encontrar o UUID)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'f54dbfaa-599a-40ee-907a-28431ce4858a'
FROM articles a
WHERE a.doi = '10.1002/nau.25551'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = 'f54dbfaa-599a-40ee-907a-28431ce4858a'
);

-- Vincular Artigo 3 (usando DOI para encontrar o UUID)
INSERT INTO article_score_items (article_id, score_item_id)
SELECT a.id, 'f54dbfaa-599a-40ee-907a-28431ce4858a'
FROM articles a
WHERE a.doi = '10.1177/19417381231195305'
AND NOT EXISTS (
  SELECT 1 FROM article_score_items asi
  WHERE asi.article_id = a.id
  AND asi.score_item_id = 'f54dbfaa-599a-40ee-907a-28431ce4858a'
);


-- 4. VERIFICAÇÃO FINAL
-- ============================================================================

-- Verificar atualização do score item
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as clinical_relevance_chars,
  LENGTH(patient_explanation) as patient_explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = 'f54dbfaa-599a-40ee-907a-28431ce4858a';

-- Verificar artigos vinculados
SELECT
  a.pm_id,
  a.title,
  a.journal,
  a.publish_date,
  a.article_type,
  a.doi
FROM articles a
INNER JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = 'f54dbfaa-599a-40ee-907a-28431ce4858a'
ORDER BY a.publish_date DESC;

COMMIT;

-- ============================================================================
-- RESUMO DO ENRIQUECIMENTO
-- ============================================================================
-- Score Item: Suporte Pélvico
-- Artigos: 3 artigos científicos (2023-2024)
-- Foco: Sistema POP-Q, guideline IUG 2023, evidência PFMT nível 1A
-- Contagens esperadas:
--   - clinical_relevance: ~1500-2000 caracteres
--   - patient_explanation: ~1000-1500 caracteres
--   - conduct: ~1500-2500 caracteres
-- ============================================================================
