-- Corrige os níveis do ScoreItem "Sexo" (vício) — subgrupo "Hábitos e vícios nocivos".
-- Bug: as 6 descrições de nível foram concatenadas no NOME de um único nível 0, deixando o
-- item com 1 nível só (todos os outros vícios do grupo têm 6 níveis: Tabaco, Álcool, Drogas,
-- Jogos/apostas). Padrão do grupo: operator '=', level_choice, gradiente 0=grave … 5=saudável,
-- só name + site_legend preenchidos.
-- Item "Sexo": 019bf31d-2ef0-7e2d-8a49-94a50eae4437 ; nível 0 existente: 019bf31d-2ef0-7026-a296-166da9082d2d
-- Idempotente. Aplicado em DEV; rodar em PROD quando autorizado.

BEGIN;

-- 1) Corrige o nome do nível 0 (estava com os 6 nomes grudados). Legenda do L0 já estava certa.
UPDATE score_levels
SET name = 'Compulsão sexual com prejuízo funcional grave', updated_at = now()
WHERE id = '019bf31d-2ef0-7026-a296-166da9082d2d';

-- 2) Cria os níveis 1..5 que faltavam (guardado por item+level p/ idempotência).
INSERT INTO score_levels (id, item_id, level, operator, name, site_legend, created_at, updated_at)
SELECT uuid_generate_v7(), '019bf31d-2ef0-7e2d-8a49-94a50eae4437', v.level, '=', v.name, v.legend, now(), now()
FROM (VALUES
  (1, 'Comportamento hipersexual frequente com impacto', 'Comportamento sexual frequente, com perda de controle e impacto na vida, merece atenção.'),
  (2, 'Padrão sexual de risco ou excessivo',             'Padrão sexual de risco ou excessivo, merece atenção.'),
  (3, 'Comportamento elevado sem prejuízo',              'Comportamento sexual elevado, sem prejuízo na vida.'),
  (4, 'Histórico de compulsão, em controle',             'Você se recuperou de um padrão de compulsão sexual, o que protege sua saúde.'),
  (5, 'Padrão saudável/sem queixas',                     'Vida sexual saudável, sem queixas.')
) AS v(level, name, legend)
WHERE NOT EXISTS (
  SELECT 1 FROM score_levels sl
  WHERE sl.item_id = '019bf31d-2ef0-7e2d-8a49-94a50eae4437' AND sl.level = v.level
);

COMMIT;
