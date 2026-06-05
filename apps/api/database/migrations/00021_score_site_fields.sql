-- Campos de exibição no SITE público (leigo) para o escore — separados das mensagens
-- clínicas do EMR. Parte da Fase 1 do plano docs/emr/plano-score-versions-unificacao.md.
--   score_items:  site_render_type (tipo de input no site), site_question (pergunta leiga,
--                 migra de light_question), site_explanation (explicação leiga).
--   score_levels: site_legend (o que o nível significa em linguagem leiga).
-- light_question/light_order/is_light_version serão removidos na Fase 2 (score_versions).

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.score_items
    ADD COLUMN IF NOT EXISTS site_render_type varchar(40),
    ADD COLUMN IF NOT EXISTS site_question    text,
    ADD COLUMN IF NOT EXISTS site_explanation text;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.score_levels
    ADD COLUMN IF NOT EXISTS site_legend text;
-- +goose StatementEnd
-- +goose StatementBegin
UPDATE public.score_items
    SET site_question = light_question
    WHERE light_question IS NOT NULL AND site_question IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.score_items
    DROP COLUMN IF EXISTS site_render_type,
    DROP COLUMN IF EXISTS site_question,
    DROP COLUMN IF EXISTS site_explanation;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.score_levels
    DROP COLUMN IF EXISTS site_legend;
-- +goose StatementEnd
