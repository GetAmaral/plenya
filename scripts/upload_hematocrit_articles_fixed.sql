-- Upload de artigos científicos: Hematócrito - Homens
-- Data: 2026-01-28
-- Item ID: b99d1e15-baa3-4bcb-956e-7314dbccfa82

BEGIN;

-- 1. NCBI Bookshelf - Hemoglobin and Hematocrit
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty
)
VALUES (
  gen_random_uuid(),
  'Hemoglobin and Hematocrit',
  'Clinical Methods - NCBI Bookshelf',
  'NCBI Bookshelf',
  '2024-01-01',
  'en',
  'https://www.ncbi.nlm.nih.gov/books/NBK259/',
  'clinical_reference',
  'Hematologia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

-- Capturar ID e criar relação
DO $$
DECLARE
  article_id_1 UUID;
BEGIN
  SELECT id INTO article_id_1 FROM articles
  WHERE original_link = 'https://www.ncbi.nlm.nih.gov/books/NBK259/'
  LIMIT 1;

  IF article_id_1 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_1, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- 2. NEJM - Polycythemia Vera Treatment
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty, doi
)
VALUES (
  gen_random_uuid(),
  'Cardiovascular Events and Intensity of Treatment in Polycythemia Vera',
  'Marchioli R, Finazzi G, Specchia G, et al.',
  'New England Journal of Medicine',
  '2013-01-03',
  'en',
  'https://www.nejm.org/doi/full/10.1056/NEJMoa1208500',
  'research_article',
  'Hematologia',
  '10.1056/NEJMoa1208500'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

DO $$
DECLARE
  article_id_2 UUID;
BEGIN
  SELECT id INTO article_id_2 FROM articles
  WHERE doi = '10.1056/NEJMoa1208500'
  LIMIT 1;

  IF article_id_2 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_2, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- 3. StatPearls - Polycythemia Review
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty
)
VALUES (
  gen_random_uuid(),
  'Polycythemia',
  'StatPearls Publishing',
  'StatPearls',
  '2024-01-01',
  'en',
  'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
  'review',
  'Hematologia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

DO $$
DECLARE
  article_id_3 UUID;
BEGIN
  SELECT id INTO article_id_3 FROM articles
  WHERE original_link = 'https://www.ncbi.nlm.nih.gov/books/NBK526081/'
  LIMIT 1;

  IF article_id_3 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_3, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- 4. ScienceDirect - Blood Viscosity
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty
)
VALUES (
  gen_random_uuid(),
  'Blood Viscosity and Oxygen Transport',
  'ScienceDirect Topics',
  'ScienceDirect',
  '2024-01-01',
  'en',
  'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity',
  'educational',
  'Fisiologia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

DO $$
DECLARE
  article_id_4 UUID;
BEGIN
  SELECT id INTO article_id_4 FROM articles
  WHERE original_link = 'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity'
  LIMIT 1;

  IF article_id_4 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_4, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- 5. Haematologica - Thrombotic Risk
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty, pm_id
)
VALUES (
  gen_random_uuid(),
  'Re-evaluation of Hematocrit as a Determinant of Thrombotic Risk in Erythrocytosis',
  'Barbui T, Thiele J, Gisslinger H, et al.',
  'Haematologica',
  '2019-04-01',
  'en',
  'https://pmc.ncbi.nlm.nih.gov/articles/PMC6442963/',
  'review',
  'Hematologia',
  '30872370'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

DO $$
DECLARE
  article_id_5 UUID;
BEGIN
  SELECT id INTO article_id_5 FROM articles
  WHERE pm_id = '30872370'
  LIMIT 1;

  IF article_id_5 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_5, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- 6. Medscape - Clinical Guideline
INSERT INTO articles (
  id, title, authors, journal, publish_date, language,
  original_link, article_type, specialty
)
VALUES (
  gen_random_uuid(),
  'Hematocrit: Reference Range, Interpretation, Collection and Panels',
  'Medscape - eMedicine',
  'Medscape',
  '2024-01-01',
  'en',
  'https://emedicine.medscape.com/article/2054320-overview',
  'clinical_guideline',
  'Hematologia'
)
ON CONFLICT (doi) DO NOTHING
RETURNING id;

DO $$
DECLARE
  article_id_6 UUID;
BEGIN
  SELECT id INTO article_id_6 FROM articles
  WHERE original_link = 'https://emedicine.medscape.com/article/2054320-overview'
  LIMIT 1;

  IF article_id_6 IS NOT NULL THEN
    INSERT INTO article_score_items (article_id, score_item_id)
    VALUES (article_id_6, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
    ON CONFLICT (article_id, score_item_id) DO NOTHING;
  END IF;
END $$;

-- Verificação final
SELECT
  a.title,
  a.journal,
  a.original_link,
  COUNT(asi.score_item_id) as linked_items
FROM articles a
LEFT JOIN article_score_items asi ON a.id = asi.article_id
WHERE a.original_link IN (
  'https://www.ncbi.nlm.nih.gov/books/NBK259/',
  'https://www.nejm.org/doi/full/10.1056/NEJMoa1208500',
  'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
  'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity',
  'https://pmc.ncbi.nlm.nih.gov/articles/PMC6442963/',
  'https://emedicine.medscape.com/article/2054320-overview'
)
GROUP BY a.id, a.title, a.journal, a.original_link
ORDER BY a.title;

COMMIT;
