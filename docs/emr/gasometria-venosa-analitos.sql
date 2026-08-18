-- Gasometria venosa — analitos que faltavam no catálogo
-- ============================================================================
-- Contexto: o laudo traz a gasometria como "Gasometria Venosa - <analito>" (8 linhas).
-- O catálogo já tinha 4 filhas do painel (pH Venoso, pCO2 Venoso, pO2 Venoso, HCO3), mas
-- o matcher casava TODAS as linhas com a definição-container "Gasometria venosa" — que não
-- tem faixa nenhuma — e os 8 analitos ficavam sem nível ("Valor fora das faixas
-- configuradas"). O match agora resolve painel → analito (resolvePanelAnalyte, com teste em
-- processing_job_match_test.go); este script fecha o outro lado: as 4 filhas que faltavam.
--
-- Idempotente: pode rodar mais de uma vez. Não cria ScoreItems — dar faixa de risco a
-- Base Excess / CO2 total / H+ é decisão clínica, não de infraestrutura.
-- Sem ScoreItem, esses resultados aparecem como "não entra no escore", que é o correto.
--
-- Aplicar:
--   dev : docker compose exec -T db psql -U plenya_user -d plenya_db -f - < docs/emr/gasometria-venosa-analitos.sql
--   prod: sudo docker exec -i <container-db> psql -U plenya_user -d plenya_db < docs/emr/gasometria-venosa-analitos.sql
-- ============================================================================

BEGIN;

-- 1) Analitos que faltavam, como filhos do painel "Gasometria venosa" (PLNA6CA66FE).
--    alt_names refletem como os laudos escrevem (já em minúsculas, sem acento — é assim que
--    o BeforeSave do model normaliza e é assim que o matcher compara).
INSERT INTO lab_test_definitions (
    id, code, name, short_name, alt_names, category, is_requestable, parent_test_id,
    unit, result_type, specimen_type, display_order, is_active, sex_applicability,
    created_at, updated_at
)
SELECT * FROM (VALUES
    (uuidv7(), 'PLNGASCO2T', 'CO2 Total (gasometria)', 'CO2T',
     '["co2 total", "gas carbonico total", "conteudo total de co2", "tco2"]'::jsonb,
     'biochemistry', false, '019bf793-96eb-770e-ab46-43b3a80c69b6'::uuid,
     'mMol/L', 'numeric', 'Sangue', 97, true, 'all', now(), now()),

    (uuidv7(), 'PLNGASBE', 'Excesso de bases (BE)', 'BE',
     '["base excess", "excesso de bases", "excesso de base", "be ecf", "be"]'::jsonb,
     'biochemistry', false, '019bf793-96eb-770e-ab46-43b3a80c69b6'::uuid,
     'mMol/L', 'numeric', 'Sangue', 98, true, 'all', now(), now()),

    (uuidv7(), 'PLNGASH', 'Hidrogênio (H+)', 'H+',
     '["hidrogenio", "ion hidrogenio", "concentracao de hidrogenio", "h+"]'::jsonb,
     'biochemistry', false, '019bf793-96eb-770e-ab46-43b3a80c69b6'::uuid,
     'nMol/L', 'numeric', 'Sangue', 100, true, 'all', now(), now())
) AS novos(id, code, name, short_name, alt_names, category, is_requestable, parent_test_id,
           unit, result_type, specimen_type, display_order, is_active, sex_applicability,
           created_at, updated_at)
WHERE NOT EXISTS (
    SELECT 1 FROM lab_test_definitions d WHERE d.code = novos.code AND d.deleted_at IS NULL
);

-- 2) Saturação venosa JÁ existia como "SatO2 Venosa" (PLN910324DC, com ScoreItem). Não se
--    cria outra: só falta o apelido com que este laudo escreve ("Saturação de O2").
UPDATE lab_test_definitions
SET alt_names = alt_names || '["saturacao de o2", "svo2"]'::jsonb, updated_at = now()
WHERE code = 'PLN910324DC'
  AND deleted_at IS NULL
  AND NOT (alt_names @> '["saturacao de o2"]'::jsonb);

-- 3) Remove a duplicata de saturação venosa criada por engano (só se ninguém a referencia).
DELETE FROM lab_test_definitions d
WHERE d.code = 'PLNGASSATO2'
  AND NOT EXISTS (SELECT 1 FROM lab_results r WHERE r.lab_test_definition_id = d.id);

-- 4) pH e HCO3 estavam como result_type 'text' entre irmãos 'numeric'. São medidas
--    numéricas: só o tipo estava errado (a classificação já lia result_numeric).
UPDATE lab_test_definitions
SET result_type = 'numeric', updated_at = now()
WHERE code IN ('PLNC9A62759', 'PLNC2996ABD')
  AND result_type <> 'numeric'
  AND deleted_at IS NULL;

COMMIT;

-- Conferência
SELECT code, name, unit, result_type, display_order
FROM lab_test_definitions
WHERE deleted_at IS NULL
  AND (code = 'PLNA6CA66FE' OR parent_test_id = '019bf793-96eb-770e-ab46-43b3a80c69b6'::uuid)
ORDER BY display_order;
