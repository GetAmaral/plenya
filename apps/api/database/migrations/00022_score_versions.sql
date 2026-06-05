-- score_versions + score_version_items — versões públicas do escore (Triagem/Light/...).
-- A version SELECIONA e ORDENA um subset de score_items; grupo/subgrupo/pilares/levels/campos de
-- site vêm do próprio item. Substitui o flag is_light_version (removido na 00023).
-- Fase 2 do plano docs/emr/plano-score-versions-unificacao.md. Seed: Light (is_light_version) + Triagem (36 curados).

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.score_versions (
    id          uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    name        varchar(120) NOT NULL,
    slug        varchar(60)  NOT NULL,
    description text,
    site_intro  text,
    "order"     integer NOT NULL DEFAULT 0,
    active      boolean NOT NULL DEFAULT true,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now(),
    deleted_at  timestamptz
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS idx_score_version_slug ON public.score_versions (slug) WHERE deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_version_order ON public.score_versions ("order");
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_versions_deleted_at ON public.score_versions (deleted_at);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.score_version_items (
    id            uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    version_id    uuid NOT NULL REFERENCES public.score_versions(id) ON DELETE CASCADE,
    score_item_id uuid NOT NULL REFERENCES public.score_items(id) ON DELETE CASCADE,
    display_order integer NOT NULL DEFAULT 0,
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now()
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS idx_score_version_item_uniq ON public.score_version_items (version_id, score_item_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_version_item_version ON public.score_version_items (version_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_version_item_item ON public.score_version_items (score_item_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_score_version_item_order ON public.score_version_items (display_order);
-- +goose StatementEnd

-- Seed das versões (UUIDs fixos para reprodutibilidade dev≡prod)
-- +goose StatementBegin
INSERT INTO public.score_versions (id, name, slug, description, "order", active, created_at, updated_at) VALUES
  ('5c012e51-0000-7000-8000-000000000001', 'Light',   'light',   'Escore Plenya Light — versão pública ampliada.',   2, true, now(), now()),
  ('5c012e51-0000-7000-8000-000000000002', 'Triagem', 'triagem', 'Escore Plenya Triagem — avaliação pública rápida.', 1, true, now(), now())
ON CONFLICT (id) DO NOTHING;
-- +goose StatementEnd

-- Light: todos os itens com is_light_version=true (ordem = light_order)
-- +goose StatementBegin
INSERT INTO public.score_version_items (version_id, score_item_id, display_order)
SELECT '5c012e51-0000-7000-8000-000000000001', si.id, COALESCE(si.light_order, 0)
FROM public.score_items si
WHERE si.is_light_version = true AND si.deleted_at IS NULL
ON CONFLICT (version_id, score_item_id) DO NOTHING;
-- +goose StatementEnd

-- Triagem: os 36 itens curados (subconjunto do light)
-- +goose StatementBegin
INSERT INTO public.score_version_items (version_id, score_item_id, display_order)
SELECT '5c012e51-0000-7000-8000-000000000002', si.id, COALESCE(si.light_order, 0)
FROM public.score_items si
WHERE si.deleted_at IS NULL AND si.id IN (
        '019bf31d-2ef0-7683-8814-9f18e88dea60', '019c5376-ca0e-7b9b-b9e7-e3cd66486705', '019c535e-d7b0-7f20-8c8c-b3c2dee441fc', 'c77cedd3-2800-7de9-ad64-b75ed1037b5e',
        'c77cedd3-2800-7bb3-9e74-b76061206583', '019cbe7a-213e-7bcb-9770-cbda22cf5164', 'c77cedd3-2800-735f-bf28-c5d07d7d7092', 'c77cedd3-2800-74fb-9aca-432654d7e0d8',
        '019c53a6-f1a3-7704-9868-354859c750cd', '019c164e-cb2c-747a-bfc7-ffb3e787229f', 'c77cedd3-2800-7ca1-9dd2-ac6e478cd0e5', 'c77cedd3-2800-7bec-942e-c1eed243a6a0',
        'c77cedd3-2800-721b-94d8-b5515b895753', 'c77cedd3-2800-7360-bf3c-5b4f28c660ef', '019bf31d-2ef0-7bb8-8305-e0e0285aeb80', 'c77cedd3-2800-7aac-985c-c075f020e9e0',
        'c77cedd3-2800-7a74-99ad-56ca4b6dddc1', '019bf31d-2ef0-7171-a061-3bcc500957f7', '019bf31d-2ef0-767a-a87e-b4587192005b', '019bf31d-2ef0-7f0a-8188-216affee192f',
        '019bf31d-2ef0-7f17-a053-7f45b7162dd2', '019bf31d-2ef0-7c5c-bc5b-00fbf09daf98', '019bf31d-2ef0-7e03-bb24-dbdc747e5fd4', '019bf31d-2ef0-7de7-a98d-603c41d12ae6',
        '019bf31d-2ef0-71ae-8f72-8330e3072647', '019bf31d-2ef0-73fc-b353-73eadb4940d5', '019bf31d-2ef0-7193-b496-afe7ce51ed91', '019bf31d-2ef0-7297-bd6d-4c83d8081d2d',
        'c77cedd3-2800-7ce0-a970-351333e28cae', '019bf31d-2ef0-7c2d-b263-5ba6bc208d28', '019bf31d-2ef0-7a08-ab7e-d9c06d8d2103', '019bf31d-2ef0-7eff-b67d-00e5eec8c2d0',
        'c77cedd3-2800-75fe-b483-ea0183478225', 'c77cedd3-2800-7524-9049-f4559170db14', 'c77cedd3-2800-7e0c-b44f-ccf9c3b926ef', 'c77cedd3-2800-70c7-942e-a1134e3aa05e'
)
ON CONFLICT (version_id, score_item_id) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.score_version_items;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.score_versions;
-- +goose StatementEnd
