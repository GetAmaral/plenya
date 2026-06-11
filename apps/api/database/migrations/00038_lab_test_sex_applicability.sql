-- Aplicabilidade por sexo do exame: filtro automático ao carregar template de pedido de exames
-- (ex.: PSA/USG próstata só carregam p/ ♂; CA-125/mamografia/USG mamas/transvaginal só p/ ♀).
-- Templates guardam TODOS os exames; o filtro acontece no load por Patient.Gender.
-- Plano: ~/.claude/plans (templates de pedidos de exames).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_test_definitions
  ADD COLUMN IF NOT EXISTS sex_applicability varchar(10) NOT NULL DEFAULT 'all'
  CHECK (sex_applicability IN ('all', 'male', 'female'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.lab_test_definitions DROP COLUMN IF EXISTS sex_applicability;
-- +goose StatementEnd
