-- Indicações do componente magistral + evidência vinda do RAG.
--
-- O CATÁLOGO SABE A DOSE, MAS NÃO SABE PARA QUÊ. Sem indicação, a busca só ajuda quem já sabe o
-- nome da substância; com indicação, ela responde "o que eu uso para sono?" — e a fórmula deixa de
-- depender da memória de quem prescreve.
--
-- DE ONDE VEM: o RAG do próprio consultório (1.183 artigos, 240 aulas da pós de Medicina Funcional
-- Integrativa, 38.277 embeddings). A regra que vale mais que o conteúdo: **evidência é ANEXADA,
-- nunca gerada**. O enriquecimento busca os trechos por similaridade, pede ao modelo que EXTRAIA o
-- que está escrito e guarda o trecho junto — quem lê confere na fonte. Trecho que não fala de dose
-- não vira dose.
--
-- `evidence_status` separa o que a máquina sugeriu do que o médico confirmou:
--   · pending   — nunca passou pelo enriquecimento
--   · suggested — veio do RAG/pesquisa e AGUARDA conferência (a tela mostra como sugestão)
--   · confirmed — o médico conferiu (a curadoria manual carimba isso)
-- Nenhum cálculo do sistema consome `article_id`: a evidência é para leitura humana.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD COLUMN IF NOT EXISTS indications      text,
    ADD COLUMN IF NOT EXISTS dose_reference   text,
    ADD COLUMN IF NOT EXISTS evidence_status  varchar(12) NOT NULL DEFAULT 'pending',
    ADD COLUMN IF NOT EXISTS enriched_at      timestamp with time zone;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS chk_magistral_components_evidence_status;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    ADD CONSTRAINT chk_magistral_components_evidence_status
    CHECK (evidence_status IN ('pending','suggested','confirmed'));
-- +goose StatementEnd

-- Espelha `article_score_items`, que é o precedente do repo para "evidência ligada a entidade
-- clínica". Guarda também o chunk exato — sem isso o médico teria que reler a aula inteira para
-- conferir uma linha.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_component_articles (
    id           uuid NOT NULL,
    component_id uuid NOT NULL,
    article_id   uuid NOT NULL,
    embedding_id uuid,
    chunk_index  integer,
    similarity   numeric(6,4),
    excerpt      text NOT NULL DEFAULT '',
    -- fixado pelo médico (a sugestão automática nasce como não fixada)
    pinned       boolean NOT NULL DEFAULT false,
    pinned_by    uuid,
    created_at   timestamp with time zone NOT NULL DEFAULT now(),
    updated_at   timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at   timestamp with time zone,
    CONSTRAINT magistral_component_articles_pkey PRIMARY KEY (id),
    CONSTRAINT fk_magistral_component_articles_component
        FOREIGN KEY (component_id) REFERENCES public.magistral_components(id) ON DELETE CASCADE,
    CONSTRAINT fk_magistral_component_articles_article
        FOREIGN KEY (article_id) REFERENCES public.articles(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_magistral_component_articles
    ON public.magistral_component_articles (component_id, article_id, chunk_index)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_component_articles_component
    ON public.magistral_component_articles (component_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_component_articles;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components DROP CONSTRAINT IF EXISTS chk_magistral_components_evidence_status;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.magistral_components
    DROP COLUMN IF EXISTS indications,
    DROP COLUMN IF EXISTS dose_reference,
    DROP COLUMN IF EXISTS evidence_status,
    DROP COLUMN IF EXISTS enriched_at;
-- +goose StatementEnd
