-- Upload de artigos científicos: Hematócrito - Mulheres
-- Score Item ID: 32037968-f263-4e7d-ab85-ea83efd61c7b
-- Data: 2026-01-28

-- Bloco transacional para inserir artigos e criar relações
DO $$
DECLARE
  article_id_1 uuid;
  article_id_2 uuid;
  article_id_3 uuid;
  article_id_4 uuid;
  article_id_5 uuid;
  score_item_id uuid := '32037968-f263-4e7d-ab85-ea83efd61c7b';
BEGIN
  -- Artigo 1: The relationship between heavy menstrual bleeding, iron deficiency, and iron deficiency anemia
  INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    pm_id,
    abstract,
    original_link,
    article_type,
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
  ) VALUES (
    'The relationship between heavy menstrual bleeding, iron deficiency, and iron deficiency anemia',
    'Multiple authors',
    'PubMed Central',
    '2023-01-27',
    'en',
    '36706856',
    'Heavy menstrual bleeding is highly prevalent in reproductive-aged women and a major contributor to iron deficiency and iron deficiency anemia. This review examines the clinical relationship between menstrual blood loss, iron stores, and hematological parameters in women.',
    'https://pubmed.ncbi.nlm.nih.gov/36706856/',
    'review',
    '["menstrual blood loss", "iron deficiency", "anemia", "hematocrit", "women"]'::jsonb,
    'Hematology',
    false,
    NOW(),
    NOW()
  ) RETURNING id INTO article_id_1;

  -- Relacionar artigo 1 com score_item
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_1, score_item_id)
  ON CONFLICT DO NOTHING;

  -- Artigo 2: The sex difference in haemoglobin levels in adults
  INSERT INTO articles (
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
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
  ) VALUES (
    'The sex difference in haemoglobin levels in adults — Mechanisms, causes, and consequences',
    'Murphy WG',
    'Blood Reviews',
    '2014-02-01',
    'en',
    '10.1016/j.blre.2013.12.003',
    '24491804',
    'Women have mean hemoglobin levels approximately 12% lower than men. This difference is attributed to sex hormones (estrogen and androgens), menstrual blood loss, and their effects on erythropoiesis and renal microvascular function.',
    'https://pubmed.ncbi.nlm.nih.gov/24491804/',
    'review',
    '["sex differences", "hematocrit", "hemoglobin", "hormones", "erythropoiesis"]'::jsonb,
    'Hematology',
    false,
    NOW(),
    NOW()
  ) RETURNING id INTO article_id_2;

  -- Relacionar artigo 2
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_2, score_item_id)
  ON CONFLICT DO NOTHING;

  -- Artigo 3: A Review of Clinical Guidelines on the Management of Iron Deficiency
  INSERT INTO articles (
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
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
  ) VALUES (
    'A Review of Clinical Guidelines on the Management of Iron Deficiency and Iron-Deficiency Anemia in Women with Heavy Menstrual Bleeding',
    'Multiple authors',
    'Advances in Therapy',
    '2020-12-01',
    'en',
    '10.1007/s12325-020-01564-y',
    '33247314',
    'This review synthesizes current clinical guidelines for managing iron deficiency and iron-deficiency anemia in women with heavy menstrual bleeding, including diagnostic algorithms and treatment protocols.',
    'https://pmc.ncbi.nlm.nih.gov/articles/PMC7695235/',
    'review',
    '["iron deficiency anemia", "heavy menstrual bleeding", "management", "guidelines", "women"]'::jsonb,
    'Hematology',
    false,
    NOW(),
    NOW()
  ) RETURNING id INTO article_id_3;

  -- Relacionar artigo 3
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_3, score_item_id)
  ON CONFLICT DO NOTHING;

  -- Artigo 4: Severe anemia from heavy menstrual bleeding
  INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    pm_id,
    abstract,
    original_link,
    article_type,
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
  ) VALUES (
    'Severe anemia from heavy menstrual bleeding requires heightened attention',
    'Multiple authors',
    'American Journal of Obstetrics and Gynecology',
    '2015-05-01',
    'en',
    '25935784',
    '149 women had 168 episodes with hemoglobin levels below 5 g/dL attributed to chronic excessive menstrual bleeding, with mean corpuscular volume of 62.2 fL indicating severe microcytic hypochromic anemia.',
    'https://pubmed.ncbi.nlm.nih.gov/25935784/',
    'research_article',
    '["severe anemia", "menstrual bleeding", "hemoglobin", "hematocrit", "women"]'::jsonb,
    'Hematology',
    false,
    NOW(),
    NOW()
  ) RETURNING id INTO article_id_4;

  -- Relacionar artigo 4
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_4, score_item_id)
  ON CONFLICT DO NOTHING;

  -- Artigo 5: Hematocrit Test: Reference Ranges and Clinical Interpretation
  INSERT INTO articles (
    title,
    authors,
    journal,
    publish_date,
    language,
    abstract,
    original_link,
    article_type,
    keywords,
    specialty,
    favorite,
    created_at,
    updated_at
  ) VALUES (
    'Hematocrit Test: Reference Ranges and Clinical Interpretation',
    'Cleveland Clinic',
    'Cleveland Clinic Health Library',
    '2024-01-15',
    'en',
    'Comprehensive clinical guide on hematocrit testing, including normal reference ranges for women (36-48%), causes of abnormal results, and clinical decision-making for follow-up testing.',
    'https://my.clevelandclinic.org/health/diagnostics/17683-hematocrit',
    'clinical_trial',
    '["hematocrit", "reference range", "women", "clinical interpretation", "anemia"]'::jsonb,
    'Hematology',
    false,
    NOW(),
    NOW()
  ) RETURNING id INTO article_id_5;

  -- Relacionar artigo 5
  INSERT INTO article_score_items (article_id, score_item_id)
  VALUES (article_id_5, score_item_id)
  ON CONFLICT DO NOTHING;

  -- Mostrar resultados
  RAISE NOTICE '✓ 5 artigos inseridos e vinculados ao score_item: %', score_item_id;
  RAISE NOTICE '  - Artigo 1: %', article_id_1;
  RAISE NOTICE '  - Artigo 2: %', article_id_2;
  RAISE NOTICE '  - Artigo 3: %', article_id_3;
  RAISE NOTICE '  - Artigo 4: %', article_id_4;
  RAISE NOTICE '  - Artigo 5: %', article_id_5;
END $$;

-- Verificar artigos vinculados (apenas os 5 novos)
SELECT
  a.id,
  LEFT(a.title, 60) as title,
  a.authors,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '32037968-f263-4e7d-ab85-ea83efd61c7b'
  AND a.created_at > NOW() - INTERVAL '1 minute'
ORDER BY a.publish_date DESC;
