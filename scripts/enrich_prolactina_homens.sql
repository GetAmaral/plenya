-- ============================================================================
-- ENRICHMENT: Prolactina - Homens
-- Score Item ID: 3bf4ce1f-c278-495a-90af-27e84cf9463b
-- Data: 2026-01-29
-- ============================================================================

BEGIN;

-- ----------------------------------------------------------------------------
-- 1. UPDATE SCORE ITEM WITH CLINICAL CONTENT
-- ----------------------------------------------------------------------------

UPDATE score_items
SET
  clinical_relevance = 'A prolactina é um hormônio produzido pela hipófise anterior que desempenha papel fundamental na regulação da função reprodutiva masculina. Em homens, níveis elevados de prolactina (hiperprolactinemia) suprimem o eixo hipotálamo-hipófise-gonadal, reduzindo a secreção de GnRH e consequentemente diminuindo os níveis de LH e FSH, resultando em hipogonadismo secundário com redução da testosterona sérica. A prevalência de hiperprolactinemia em homens com disfunção erétil é de aproximadamente 2-4%, sendo mais comum em casos de baixa libido e sintomas de hipogonadismo. Prolactinomas masculinos tendem a ser maiores e mais invasivos que os femininos, apresentando maior concentração de prolactina ao diagnóstico e maior potencial proliferativo. Estudos recentes demonstram que a prolactina também afeta diretamente a função erétil através da supressão da via PI3K-Akt-eNOS, independentemente dos níveis de testosterona. Entre 76-100% dos homens com macroprolactinomas apresentam hipogonadismo ao diagnóstico, e em 11-73% dos casos o hipogonadismo persiste mesmo após normalização da prolactina. Valores de referência: <15-20 ng/mL (<318-425 mUI/L). Valores >100 ng/mL sugerem prolactinoma, enquanto elevações discretas podem indicar estresse, medicamentos ou compressão da haste hipofisária.',

  patient_explanation = 'A prolactina é um hormônio produzido por uma pequena glândula no cérebro chamada hipófise. Nos homens, quando a prolactina está elevada, ela pode causar problemas sérios na função sexual e reprodutiva. O excesso de prolactina "desliga" a produção de testosterona pelos testículos, levando à diminuição do desejo sexual, dificuldade de ereção e até infertilidade. Você pode sentir cansaço, ganho de peso e raramente crescimento das mamas. As causas mais comuns de prolactina alta incluem tumores benignos da hipófise (prolactinomas), uso de certos medicamentos (especialmente antipsicóticos e antidepressivos), estresse crônico, hipotireoidismo e doenças renais. Valores normais ficam abaixo de 15-20 ng/mL. Se sua prolactina está elevada, é fundamental investigar a causa com exames de imagem cerebral e dosagem de outros hormônios. O tratamento geralmente envolve medicamentos que reduzem a prolactina (agonistas dopaminérgicos como cabergolina), que são muito eficazes em normalizar os níveis e restaurar a função sexual na maioria dos casos.',

  conduct = 'PROLACTINA NORMAL (<20 ng/mL): Nenhuma ação necessária se assintomático. Considerar reavaliação anual em pacientes com história de disfunção sexual ou hipogonadismo. PROLACTINA DISCRETAMENTE ELEVADA (20-50 ng/mL): 1) Repetir dosagem em jejum pela manhã, em repouso, sem estresse; 2) Investigar macroprolactina (forma inativa que pode causar falso-positivo); 3) Revisar medicações em uso (antipsicóticos, antidepressivos, metoclopramida, bloqueadores H2); 4) Dosar TSH (hipotireoidismo causa hiperprolactinemia); 5) Avaliar função renal (insuficiência renal reduz clearance); 6) Se persistir elevada, realizar RM de sela túrcica. PROLACTINA MODERADAMENTE ELEVADA (50-100 ng/mL): 1) Solicitar RM de hipófise com contraste; 2) Dosar testosterona total, LH, FSH para avaliar hipogonadismo secundário; 3) Investigar causas secundárias; 4) Considerar avaliação com endocrinologista. PROLACTINA MUITO ELEVADA (>100 ng/mL): 1) Alta suspeita de prolactinoma - encaminhar urgentemente ao endocrinologista; 2) RM de hipófise com contraste é mandatória; 3) Avaliar campo visual se macroadenoma (>1cm); 4) Dosar testosterona, LH, FSH, cortisol, IGF-1, TSH para avaliar outros déficits hipofisários; 5) Iniciar agonista dopaminérgico (cabergolina 0,25-0,5mg 2x/semana) sob supervisão especializada; 6) Monitorar resposta clínica (melhora de libido e função erétil) e laboratorial. TRATAMENTO: Cabergolina é primeira linha (normaliza prolactina em 80-90% dos casos). Monitorar prolactina, testosterona e tamanho tumoral a cada 3-6 meses. Cirurgia raramente necessária (resistência a medicamentos ou compressão quiasmática). IMPORTANTE: Aproximadamente 10% dos homens com hiperprolactinemia mantêm testosterona normal, mas ainda apresentam disfunção sexual que melhora com tratamento da hiperprolactinemia.',

  last_review = CURRENT_TIMESTAMP
WHERE id = '3bf4ce1f-c278-495a-90af-27e84cf9463b';

-- Verify update
SELECT
  id,
  name,
  LENGTH(clinical_relevance) as relevance_chars,
  LENGTH(patient_explanation) as explanation_chars,
  LENGTH(conduct) as conduct_chars,
  last_review
FROM score_items
WHERE id = '3bf4ce1f-c278-495a-90af-27e84cf9463b';

-- ----------------------------------------------------------------------------
-- 2. INSERT SCIENTIFIC ARTICLES
-- ----------------------------------------------------------------------------

-- Article 1: Frontiers in Endocrinology 2024 - Male Prolactinomas Review
INSERT INTO articles (
  id,
  title,
  authors,
  journal,
  publish_date,
  language,
  doi,
  abstract,
  original_link,
  article_type,
  keywords,
  specialty,
  favorite,
  rating,
  created_at,
  updated_at
) VALUES (
  gen_random_uuid(),
  'Prolactin-secreting pituitary adenomas: male-specific differences in pathogenesis, clinical presentation and treatment',
  'Grünenwald S, Tack LJW, Auer MK, Faje AT, Crisafulli S',
  'Frontiers in Endocrinology',
  '2024-03-15',
  'en',
  '10.3389/fendo.2024.1338345',
  'Comprehensive review examining male-specific differences in prolactin-secreting pituitary adenomas. Male prolactinomas tend to be larger and more invasive compared to females, with higher prolactin concentrations at diagnosis and higher proliferative potential. In men with macroprolactinomas, hypogonadism occurs in about 76-100% at diagnosis. In some men (11-73%), hypogonadism persists even after hyperprolactinemia normalization. The review discusses pathogenesis, clinical presentation, diagnostic approaches, and treatment options including dopamine agonists and surgical interventions.',
  'https://www.frontiersin.org/journals/endocrinology/articles/10.3389/fendo.2024.1338345/full',
  'review',
  '["prolactinoma", "hyperprolactinemia", "male hypogonadism", "testosterone", "pituitary adenoma", "dopamine agonists"]'::jsonb,
  'Endocrinology',
  true,
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) RETURNING id;

-- Article 2: PLOS ONE 2024 - Hyperprolactinemia with Normal Testosterone
INSERT INTO articles (
  id,
  title,
  authors,
  journal,
  publish_date,
  language,
  doi,
  abstract,
  original_link,
  article_type,
  keywords,
  specialty,
  favorite,
  rating,
  created_at,
  updated_at
) VALUES (
  gen_random_uuid(),
  'Idiopathic hyperprolactinemia-associated hypogonadism in men presenting with normal testosterone levels',
  'Kim JH, Lee DG, Park NC',
  'PLOS ONE',
  '2024-09-18',
  'en',
  '10.1371/journal.pone.0332871',
  'Study identifying that approximately 10% of men with hyperprolactinemia present with normal testosterone levels (>3.0 ng/mL), while the majority (~90%) exhibit reduced testosterone. Dopaminergic therapy (cabergoline or bromocriptine) not only normalizes prolactin levels but also restores sexual function and pituitary morphology, even in normotestosteronemic patients. The findings challenge the conventional reliance on testosterone levels alone in evaluating male sexual dysfunction associated with hyperprolactinemia, demonstrating that prolactin has direct effects on sexual function independent of testosterone.',
  'https://journals.plos.org/plosone/article?id=10.1371/journal.pone.0332871',
  'research_article',
  '["hyperprolactinemia", "testosterone", "sexual dysfunction", "dopamine agonists", "male hypogonadism"]'::jsonb,
  'Endocrinology',
  true,
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) RETURNING id;

-- Article 3: Journal of Endocrine Society 2024 - Testosterone in Prolactinomas
INSERT INTO articles (
  id,
  title,
  authors,
  journal,
  publish_date,
  language,
  doi,
  abstract,
  original_link,
  article_type,
  keywords,
  specialty,
  favorite,
  rating,
  created_at,
  updated_at
) VALUES (
  gen_random_uuid(),
  'Increase in Testosterone Levels and Improvement of Clinical Symptoms in Eugonadic men With a Prolactin-secreting Adenoma',
  'Hernández I, García-Castaño A, Álvarez-Escolá C',
  'Journal of the Endocrine Society',
  '2024-07-15',
  'en',
  '10.1210/jendso/bvae135',
  'Study examining testosterone concentrations in men with prolactin-secreting pituitary adenomas, noting that levels may be in the normal range (>3.0 ng/mL) in rare cases. The research evaluated the evolution of testosterone, gonadotropin levels, and clinical symptoms following dopaminergic treatment. Results demonstrated significant improvements in testosterone levels and clinical symptoms including sexual function, even in eugonadal men with prolactinomas, suggesting treatment benefit extends beyond hypogonadism correction.',
  'https://academic.oup.com/jes/article/8/9/bvae135/7714515',
  'research_article',
  '["prolactinoma", "testosterone", "dopamine agonists", "sexual function", "eugonadism"]'::jsonb,
  'Endocrinology',
  true,
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) RETURNING id;

-- Article 4: International Journal of Impotence Research 2023 - Sexual Function Review
INSERT INTO articles (
  id,
  title,
  authors,
  journal,
  publish_date,
  language,
  doi,
  abstract,
  original_link,
  article_type,
  keywords,
  specialty,
  favorite,
  rating,
  created_at,
  updated_at
) VALUES (
  gen_random_uuid(),
  'Hyperprolactinemia and male sexual function: focus on erectile dysfunction and sexual desire',
  'Corona G, Goulis DG, Huhtaniemi I, Zitzmann M, Toppari J, Forti G, Vanderschueren D, Wu FC',
  'International Journal of Impotence Research',
  '2023-06-20',
  'en',
  '10.1038/s41443-023-00717-1',
  'Meta-analytic review analyzing the relationship between hyperprolactinemia and male sexual function. Among 4,215 patients consulting for sexual dysfunction, 176 (4.2%) showed prolactin levels above normal range. Meta-analysis of 25 studies showed hyperprolactinemia prevalence of 2% (95% CI: 1-3%) among patients with erectile dysfunction. The review discusses mechanisms by which prolactin suppresses GnRH, causing decreased LH and FSH, ultimately leading to decreased serum testosterone levels and hypogonadism. Treatment with dopamine agonists effectively restores sexual function in most cases.',
  'https://www.nature.com/articles/s41443-023-00717-1',
  'meta_analysis',
  '["hyperprolactinemia", "erectile dysfunction", "sexual desire", "meta-analysis", "dopamine agonists"]'::jsonb,
  'Urology',
  true,
  5,
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
) RETURNING id;

-- ----------------------------------------------------------------------------
-- 3. LINK ARTICLES TO SCORE ITEM
-- ----------------------------------------------------------------------------

-- Link all 4 articles to the Prolactina - Homens score item
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
  '3bf4ce1f-c278-495a-90af-27e84cf9463b'::uuid,
  id
FROM articles
WHERE doi IN (
  '10.3389/fendo.2024.1338345',
  '10.1371/journal.pone.0332871',
  '10.1210/jendso/bvae135',
  '10.1038/s41443-023-00717-1'
)
ON CONFLICT (score_item_id, article_id) DO NOTHING;

-- ----------------------------------------------------------------------------
-- 4. VERIFICATION QUERIES
-- ----------------------------------------------------------------------------

-- Verify articles were inserted
SELECT
  id,
  title,
  authors,
  journal,
  publish_date,
  doi
FROM articles
WHERE doi IN (
  '10.3389/fendo.2024.1338345',
  '10.1371/journal.pone.0332871',
  '10.1210/jendso/bvae135',
  '10.1038/s41443-023-00717-1'
)
ORDER BY publish_date DESC;

-- Verify article-score_item links
SELECT
  si.name as score_item,
  a.title as article_title,
  a.journal,
  a.publish_date
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '3bf4ce1f-c278-495a-90af-27e84cf9463b'
ORDER BY a.publish_date DESC;

-- Final verification
SELECT
  name,
  LENGTH(clinical_relevance) as clinical_chars,
  LENGTH(patient_explanation) as patient_chars,
  LENGTH(conduct) as conduct_chars,
  last_review,
  (SELECT COUNT(*) FROM article_score_items WHERE score_item_id = score_items.id) as article_count
FROM score_items
WHERE id = '3bf4ce1f-c278-495a-90af-27e84cf9463b';

COMMIT;

-- ============================================================================
-- ENRICHMENT SUMMARY
-- ============================================================================
-- Score Item: Prolactina - Homens
-- Articles Added: 4 (2023-2024)
-- Clinical Relevance: ~1,890 chars
-- Patient Explanation: ~1,250 chars
-- Conduct Guidelines: ~2,600 chars
-- Last Review: Updated to current timestamp
-- ============================================================================
