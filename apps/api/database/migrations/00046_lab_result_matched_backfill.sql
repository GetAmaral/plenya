-- Conserta o campo `matched` dos resultados existentes: por causa do default:true no
-- GORM, todo result entrou com matched=true (mesmo sem definição catalogada), anulando
-- a sinalização "Não catalogado". A verdade é: matched = (tem lab_test_definition_id).
-- Ver docs/emr/plano-importacao-laudos-melhorias.md (Fase 3, achado no smoke test).

-- +goose Up
-- +goose StatementBegin
UPDATE public.lab_results
SET matched = (lab_test_definition_id IS NOT NULL);
-- +goose StatementEnd

-- +goose StatementBegin
UPDATE public.lab_results
SET match_reason = 'Não encontrado no catálogo de exames'
WHERE matched = false AND match_reason IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd
