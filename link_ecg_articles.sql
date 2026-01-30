-- Criar relações entre artigos e o score_item ECG - Sokolow-Lyon
-- Score Item ID: 631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb

-- Relação com meta-análise (PubMed 32745730)
INSERT INTO article_score_items (article_id, score_item_id, relevance_notes, created_at, updated_at)
SELECT
    a.id,
    '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'::uuid,
    'Meta-análise com 58.400 participantes demonstrando que Sokolow-Lyon prediz MACE com RR 1.62 (IC95% 1.40-1.89), mortalidade geral RR 1.47 e mortalidade cardiovascular RR 1.38. Evidência de alta qualidade para estratificação de risco.',
    NOW(),
    NOW()
FROM articles a
WHERE a.pm_id = '32745730'
ON CONFLICT DO NOTHING;

-- Relação com estudo de diferenças por sexo (PubMed 16708082)
INSERT INTO article_score_items (article_id, score_item_id, relevance_notes, created_at, updated_at)
SELECT
    a.id,
    '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'::uuid,
    'Estudo prospectivo de 6.668 hipertensos (11.2 anos) mostrando que cada 0.1 mV aumenta risco de morte cardiovascular em 1.4-3.9%. Mulheres têm maior risco de AVC, homens de DAC. Importante para interpretação ajustada por sexo.',
    NOW(),
    NOW()
FROM articles a
WHERE a.pm_id = '16708082'
ON CONFLICT DO NOTHING;

-- Verificar relações criadas
SELECT
    si.name as score_item,
    a.title as article,
    a.pm_id,
    a.publish_date,
    asi.relevance_notes
FROM article_score_items asi
JOIN score_items si ON asi.score_item_id = si.id
JOIN articles a ON asi.article_id = a.id
WHERE si.id = '631f2baf-49e4-4c5b-bea0-13bf3f2a4fbb'
ORDER BY a.publish_date DESC;
