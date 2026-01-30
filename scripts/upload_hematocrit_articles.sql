-- Upload de artigos científicos: Hematócrito - Homens
-- Data: 2026-01-28
-- Item ID: b99d1e15-baa3-4bcb-956e-7314dbccfa82

BEGIN;

-- 1. NCBI Bookshelf - Hemoglobin and Hematocrit
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Hemoglobin and Hematocrit',
  'clinical_reference',
  'NCBI Bookshelf - Clinical Methods',
  'https://www.ncbi.nlm.nih.gov/books/NBK259/',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

-- Capturar ID do artigo inserido
DO $$
DECLARE
  article_id_1 UUID;
BEGIN
  SELECT id INTO article_id_1 FROM articles WHERE url = 'https://www.ncbi.nlm.nih.gov/books/NBK259/';

  -- Criar relação com score_item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_1, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- 2. NEJM - Polycythemia Vera Treatment
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Cardiovascular Events and Intensity of Treatment in Polycythemia Vera',
  'peer_reviewed',
  'New England Journal of Medicine (2013)',
  'https://www.nejm.org/doi/full/10.1056/NEJMoa1208500',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

DO $$
DECLARE
  article_id_2 UUID;
BEGIN
  SELECT id INTO article_id_2 FROM articles WHERE url = 'https://www.nejm.org/doi/full/10.1056/NEJMoa1208500';

  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_2, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- 3. StatPearls - Polycythemia Review
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Polycythemia',
  'clinical_reference',
  'StatPearls - NCBI Bookshelf',
  'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

DO $$
DECLARE
  article_id_3 UUID;
BEGIN
  SELECT id INTO article_id_3 FROM articles WHERE url = 'https://www.ncbi.nlm.nih.gov/books/NBK526081/';

  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_3, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- 4. ScienceDirect - Blood Viscosity
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Blood Viscosity and Oxygen Transport',
  'educational',
  'ScienceDirect Topics',
  'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

DO $$
DECLARE
  article_id_4 UUID;
BEGIN
  SELECT id INTO article_id_4 FROM articles WHERE url = 'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity';

  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_4, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- 5. Haematologica - Thrombotic Risk
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Re-evaluation of Hematocrit as a Determinant of Thrombotic Risk in Erythrocytosis',
  'peer_reviewed',
  'Haematologica (2019)',
  'https://pmc.ncbi.nlm.nih.gov/articles/PMC6442963/',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

DO $$
DECLARE
  article_id_5 UUID;
BEGIN
  SELECT id INTO article_id_5 FROM articles WHERE url = 'https://pmc.ncbi.nlm.nih.gov/articles/PMC6442963/';

  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_5, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- 6. Medscape - Clinical Guideline
INSERT INTO articles (id, title, source_type, source, url, language, status)
VALUES (
  gen_random_uuid(),
  'Hematocrit: Reference Range, Interpretation, Collection and Panels',
  'clinical_guideline',
  'Medscape - eMedicine',
  'https://emedicine.medscape.com/article/2054320-overview',
  'en',
  'active'
)
ON CONFLICT (url) DO UPDATE SET
  title = EXCLUDED.title,
  source = EXCLUDED.source,
  updated_at = NOW()
RETURNING id;

DO $$
DECLARE
  article_id_6 UUID;
BEGIN
  SELECT id INTO article_id_6 FROM articles WHERE url = 'https://emedicine.medscape.com/article/2054320-overview';

  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_6, 'b99d1e15-baa3-4bcb-956e-7314dbccfa82')
  ON CONFLICT (article_id, score_item_id) DO NOTHING;
END $$;

-- Verificação final
SELECT
  a.title,
  a.source,
  a.url,
  COUNT(asi.score_item_id) as linked_items
FROM articles a
LEFT JOIN article_score_items asi ON a.id = asi.article_id
WHERE a.url IN (
  'https://www.ncbi.nlm.nih.gov/books/NBK259/',
  'https://www.nejm.org/doi/full/10.1056/NEJMoa1208500',
  'https://www.ncbi.nlm.nih.gov/books/NBK526081/',
  'https://www.sciencedirect.com/topics/agricultural-and-biological-sciences/blood-viscosity',
  'https://pmc.ncbi.nlm.nih.gov/articles/PMC6442963/',
  'https://emedicine.medscape.com/article/2054320-overview'
)
GROUP BY a.id, a.title, a.source, a.url
ORDER BY a.title;

COMMIT;
