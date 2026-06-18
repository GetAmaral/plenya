-- Layout do template no próprio dado (opção C): ordem dos exames + quebra de página.
-- O template passa a ser a fonte da verdade do layout; o pedido (lab_requests.exams) segue texto
-- livre. Regra única de paginação: linha em branco = nova página. Nenhuma regra de agrupamento no
-- código — quem divide página é o flag page_break_before, configurado por template.
-- Plano: docs/emr/plano-templates-layout-estruturado.md

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.lab_request_template_tests
  ADD COLUMN IF NOT EXISTS display_order integer NOT NULL DEFAULT 0,
  ADD COLUMN IF NOT EXISTS page_break_before boolean NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_lab_request_template_tests_order
  ON public.lab_request_template_tests (lab_request_template_id, display_order);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS public.idx_lab_request_template_tests_order;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.lab_request_template_tests
  DROP COLUMN IF EXISTS display_order,
  DROP COLUMN IF EXISTS page_break_before;
-- +goose StatementEnd
