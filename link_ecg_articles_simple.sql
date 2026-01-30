-- Criar relações entre artigos e o score_item ECG - Sokolow-Lyon
-- Score Item ID: 631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb

-- Relação com meta-análise (PubMed 32745730)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'::uuid,
    a.id
FROM articles a
WHERE a.pm_id = '32745730'
ON CONFLICT DO NOTHING;

-- Relação com estudo de diferenças por sexo (PubMed 16708082)
INSERT INTO article_score_items (score_item_id, article_id)
SELECT
    '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'::uuid,
    a.id
FROM articles a
WHERE a.pm_id = '16708082'
ON CONFLICT DO NOTHING;

-- Verificar relações criadas
SELECT
    si.name as score_item,
    a.title as article,
    a.pm_id,
    a.publish_date
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'
ORDER BY a.publish_date DESC;
