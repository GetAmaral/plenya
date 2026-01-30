-- Criar relações many-to-many entre artigos e score_item CIMT
-- Score Item ID: 9064aa05-9f21-402f-9c36-49887d2b8ec9 (CIMT Carótidas)

-- Verificar se o score_item existe
SELECT id, name FROM score_items WHERE id = '9064aa05-9f21-402f-9c36-49887d2b8ec9';

-- Inserir relações com os 4 artigos
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '9064aa05-9f21-402f-9c36-49887d2b8ec9'::uuid
FROM articles
WHERE pm_id IN ('PMC12162042', 'PMC2691892', 'PMC11429111', 'PMC11663449')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Adicionar também os artigos de Hipertensão que já existiam
INSERT INTO article_score_items (article_id, score_item_id)
SELECT id, '9064aa05-9f21-402f-9c36-49887d2b8ec9'::uuid
FROM articles
WHERE id IN ('38a12fd2-1298-4907-811d-9ae980e1bad8', 'a29e47f5-30d5-40d0-bb18-9f8502540a0d')
ON CONFLICT (article_id, score_item_id) DO NOTHING;

-- Verificar as relações criadas
SELECT
    a.title,
    a.pm_id,
    si.name as score_item
FROM article_score_items asi
JOIN articles a ON asi.article_id = a.id
JOIN score_items si ON asi.score_item_id = si.id
WHERE asi.score_item_id = '9064aa05-9f21-402f-9c36-49887d2b8ec9'
ORDER BY a.publish_date DESC;
