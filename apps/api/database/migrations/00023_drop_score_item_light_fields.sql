-- Remove as colunas legadas do "Escore Light" de score_items. O conjunto Light passou a ser
-- dirigido pela score_version slug="light" (score_version_items), e a pergunta leiga vive em
-- site_question (migrada na 00021). Fase 2e do plano docs/emr/plano-score-versions-unificacao.md.
-- Pré-requisito: 00022 (cria score_versions e semeia a version "light" a partir de is_light_version)
-- já aplicada — por isso esta migration vem DEPOIS. DROP COLUMN remove os índices dependentes.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.score_items
    DROP COLUMN IF EXISTS is_light_version,
    DROP COLUMN IF EXISTS light_order,
    DROP COLUMN IF EXISTS light_question;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.score_items
    ADD COLUMN IF NOT EXISTS is_light_version boolean NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS light_order integer,
    ADD COLUMN IF NOT EXISTS light_question text;
-- +goose StatementEnd
-- +goose StatementBegin
-- Restaura o estado legado a partir das novas fontes (best-effort).
UPDATE public.score_items si
   SET is_light_version = true
  FROM public.score_version_items svi
  JOIN public.score_versions sv ON sv.id = svi.version_id
 WHERE svi.score_item_id = si.id AND sv.slug = 'light' AND sv.deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
UPDATE public.score_items si
   SET light_order = svi.display_order
  FROM public.score_version_items svi
  JOIN public.score_versions sv ON sv.id = svi.version_id
 WHERE svi.score_item_id = si.id AND sv.slug = 'light' AND sv.deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
UPDATE public.score_items SET light_question = site_question WHERE site_question IS NOT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_item_light ON public.score_items (is_light_version);
CREATE INDEX IF NOT EXISTS idx_score_item_light_order ON public.score_items (light_order);
-- +goose StatementEnd
