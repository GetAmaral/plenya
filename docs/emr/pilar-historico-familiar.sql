-- Reorganização de pilares: separa "Histórico Familiar" do pilar "Genético".
-- - Genético passa a conter só testes genéticos reais (SNPs + genotipagens).
-- - Novo pilar "Histórico Familiar" (eixo G — Gestão Clínica e Metabólica) recebe os itens
--   de história familiar de doença + os "Histórico familiar de X" espalhados pelo score.
-- Idempotente. Aplicado em DEV; rodar em PROD quando autorizado.
-- IDs estáveis (seed): Genético = a91d9014-…; novo pilar = a91d9024-….

BEGIN;

-- 1) Cria o pilar "Histórico Familiar" no eixo G (se ainda não existe).
INSERT INTO method_pillars (id, name, "order", letter_id, created_at, updated_at)
SELECT
  'a91d9024-0000-7000-8000-000000000000',
  'Histórico Familiar',
  24,
  (SELECT id FROM method_letters WHERE name = 'Gestão Clínica e Metabólica'),
  now(), now()
WHERE NOT EXISTS (
  SELECT 1 FROM method_pillars WHERE id = 'a91d9024-0000-7000-8000-000000000000'
);

-- Conjunto-alvo: história familiar de DOENÇA (subgrupos "Parentes próximos" e "Parentes
-- mais distantes" do grupo "Histórico Familiar de Doenças") + itens "Histórico familiar de X"
-- em qualquer outro grupo. Exclui linha-do-tempo do paciente e "situação/socialização familiar".
WITH alvo AS (
  SELECT si.id
  FROM score_items si
  JOIN score_subgroups sg ON sg.id = si.subgroup_id
  JOIN score_groups g ON g.id = sg.group_id
  WHERE (g.name = 'Histórico Familiar de Doenças' AND sg.name LIKE 'Parentes %')
     OR (g.name <> 'Histórico Familiar de Doenças' AND si.name ~* '^hist[oó]rico familiar')
)
-- 2) Vincula os alvos ao novo pilar (sem duplicar).
INSERT INTO score_item_method_pillars (score_item_id, method_pillar_id)
SELECT a.id, 'a91d9024-0000-7000-8000-000000000000'
FROM alvo a
WHERE NOT EXISTS (
  SELECT 1 FROM score_item_method_pillars x
  WHERE x.score_item_id = a.id AND x.method_pillar_id = 'a91d9024-0000-7000-8000-000000000000'
);

-- 3) Remove do pilar Genético os itens de história familiar de doença (não são teste genético).
DELETE FROM score_item_method_pillars mp
USING score_items si
JOIN score_subgroups sg ON sg.id = si.subgroup_id
JOIN score_groups g ON g.id = sg.group_id
WHERE mp.score_item_id = si.id
  AND mp.method_pillar_id = 'a91d9014-0000-7000-8000-000000000000' -- Genético
  AND g.name = 'Histórico Familiar de Doenças'
  AND sg.name LIKE 'Parentes %';

COMMIT;
