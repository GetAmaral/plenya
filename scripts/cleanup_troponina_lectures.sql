-- Remove incorrect lecture links (lectures shouldn't be linked to this lab test item)
DELETE FROM article_score_items
WHERE score_item_id = '247a9b99-ec59-4dd7-b7a3-b5482b1dd553'
AND article_id IN (
  SELECT id FROM articles WHERE article_type = 'lecture'
);

-- Verify final state
SELECT
  a.pm_id,
  a.title,
  a.journal,
  a.publish_date,
  a.article_type
FROM articles a
JOIN article_score_items asi ON a.id = asi.article_id
WHERE asi.score_item_id = '247a9b99-ec59-4dd7-b7a3-b5482b1dd553'
ORDER BY a.publish_date DESC;
