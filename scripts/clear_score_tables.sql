-- Script para limpar todas as tabelas de score antes da re-importação
-- Ordem: levels -> items -> subgroups -> groups (respeitando foreign keys)

BEGIN;

-- Truncate na ordem correta (de dependentes para principais)
TRUNCATE TABLE score_levels CASCADE;
TRUNCATE TABLE score_items CASCADE;
TRUNCATE TABLE score_subgroups CASCADE;
TRUNCATE TABLE score_groups CASCADE;

COMMIT;
