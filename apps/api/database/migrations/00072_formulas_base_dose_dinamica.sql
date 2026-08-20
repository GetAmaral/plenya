-- Fórmulas-base (templates) e regras de dose dinâmica.
--
-- FÓRMULA-BASE: hoje o médico remonta "a fórmula do sono" componente a componente toda vez.
-- O template guarda a fórmula pronta com sua própria indicação, e a receita passa a nascer dela.
-- `prescription_formulas.template_id` (criada na 00068 sem FK, à espera desta tabela) ganha a
-- referência agora: a receita emitida registra de qual base saiu, informação impossível de
-- reconstruir depois.
--
-- DOSE DINÂMICA: as regras vivem NO TEMPLATE, nunca na receita emitida. O motor devolve
-- SUGESTÃO — um número com a base escrita ao lado ("25-OH-vitamina D = 22 ng/mL de 12/07 ·
-- regra: <30 ⇒ 5.000 UI/dia") que o médico aceita ou ignora. Consequências de projeto,
-- deliberadas:
--
--   1. `min_dose`/`max_dose` são NOT NULL na regra. Toda regra é travada nos dois extremos; uma
--      regra por kg com peso errado no prontuário não consegue sugerir dose absurda.
--   2. O payload de criação de receita não tem noção de regra. Um bug no motor é INCAPAZ de
--      produzir receita assinada — ele só preenche um campo que o médico ainda vai olhar.
--   3. `prescription_formula_components.suggested_quantity` guarda o que foi sugerido ao lado do
--      que foi prescrito. Sem isso não há como saber, depois, quando o médico discordou.
--
-- O exame da regra é uma referência a `lab_test_definitions.code`, escolhida pelo médico na tela:
-- nenhum código clínico entra hardcoded no Go.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_formula_templates (
    id                   uuid NOT NULL,
    name                 varchar(200) NOT NULL,
    indication           text,
    indication_bullets   text,
    pharmaceutical_form  varchar(60)  NOT NULL,
    usage_type           varchar(10)  NOT NULL DEFAULT 'internal',
    route                varchar(40)  NOT NULL DEFAULT '',
    vehicle              varchar(200) NOT NULL DEFAULT '',
    quantity_to_dispense numeric(12,3) NOT NULL DEFAULT 60,
    quantity_unit        varchar(30)  NOT NULL DEFAULT 'cápsulas',
    posology             varchar(300) NOT NULL DEFAULT '',
    duration             integer      NOT NULL DEFAULT 0,
    instructions         text,
    notes                text,
    usage_count          integer      NOT NULL DEFAULT 0,
    is_active            boolean      NOT NULL DEFAULT true,
    created_by           uuid,
    reviewed_by          uuid,
    last_review          timestamp with time zone,
    created_at           timestamp with time zone NOT NULL DEFAULT now(),
    updated_at           timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at           timestamp with time zone,
    CONSTRAINT magistral_formula_templates_pkey PRIMARY KEY (id),
    CONSTRAINT chk_magistral_formula_templates_usage CHECK (usage_type IN ('internal','external')),
    CONSTRAINT chk_magistral_formula_templates_qty CHECK (quantity_to_dispense > 0)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_formula_templates_active
    ON public.magistral_formula_templates (is_active, usage_count DESC);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_formula_templates_trgm
    ON public.magistral_formula_templates
    USING gin ((lower(public.immutable_unaccent(name)) || ' ' || lower(public.immutable_unaccent(coalesce(indication,'')))) gin_trgm_ops);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_formula_template_components (
    id                     uuid NOT NULL,
    template_id            uuid NOT NULL,
    magistral_component_id uuid,
    display_order          integer      NOT NULL DEFAULT 0,
    substance              varchar(200) NOT NULL,
    quantity               numeric(14,4) NOT NULL,
    unit                   varchar(20)  NOT NULL,
    category               varchar(20)  NOT NULL DEFAULT 'simple',
    note                   varchar(200) NOT NULL DEFAULT '',
    created_at             timestamp with time zone NOT NULL DEFAULT now(),
    updated_at             timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at             timestamp with time zone,
    CONSTRAINT magistral_formula_template_components_pkey PRIMARY KEY (id),
    CONSTRAINT fk_mftc_template FOREIGN KEY (template_id)
        REFERENCES public.magistral_formula_templates(id) ON DELETE CASCADE,
    CONSTRAINT fk_mftc_component FOREIGN KEY (magistral_component_id)
        REFERENCES public.magistral_components(id),
    CONSTRAINT chk_mftc_quantity CHECK (quantity > 0)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_mftc_template ON public.magistral_formula_template_components (template_id);
-- +goose StatementEnd

-- Regra de dose de UM componente do template. Piso e teto são obrigatórios de propósito.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_formula_template_rules (
    id                     uuid NOT NULL,
    template_component_id  uuid NOT NULL,
    -- fixed: dose fixa · per_kg: dose por kg de peso · lab_threshold: dose conforme um exame
    kind                   varchar(16) NOT NULL,
    -- per_kg
    per_kg                 numeric(12,4),
    -- lab_threshold
    lab_code               varchar(64),
    lab_operator           varchar(4),
    lab_threshold          numeric(14,4),
    dose_if_true           numeric(14,4),
    dose_if_false          numeric(14,4),
    -- fixed
    fixed_dose             numeric(14,4),
    -- trava obrigatória nos dois extremos
    min_dose               numeric(14,4) NOT NULL,
    max_dose               numeric(14,4) NOT NULL,
    -- dado mais velho que isto não sugere nada
    max_data_age_days      integer NOT NULL DEFAULT 365,
    note                   varchar(300) NOT NULL DEFAULT '',
    created_at             timestamp with time zone NOT NULL DEFAULT now(),
    updated_at             timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at             timestamp with time zone,
    CONSTRAINT magistral_formula_template_rules_pkey PRIMARY KEY (id),
    CONSTRAINT fk_mftr_component FOREIGN KEY (template_component_id)
        REFERENCES public.magistral_formula_template_components(id) ON DELETE CASCADE,
    CONSTRAINT chk_mftr_kind CHECK (kind IN ('fixed','per_kg','lab_threshold')),
    CONSTRAINT chk_mftr_operator CHECK (lab_operator IS NULL OR lab_operator IN ('lt','lte','gt','gte')),
    CONSTRAINT chk_mftr_range CHECK (min_dose > 0 AND max_dose >= min_dose),
    CONSTRAINT chk_mftr_shape CHECK (
        (kind = 'fixed'         AND fixed_dose IS NOT NULL) OR
        (kind = 'per_kg'        AND per_kg     IS NOT NULL) OR
        (kind = 'lab_threshold' AND lab_code IS NOT NULL AND lab_operator IS NOT NULL
                                AND lab_threshold IS NOT NULL AND dose_if_true IS NOT NULL)
    )
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_mftr_component
    ON public.magistral_formula_template_rules (template_component_id) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- Evidência do RAG ligada à fórmula-base, espelhando magistral_component_articles.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_formula_template_articles (
    id           uuid NOT NULL,
    template_id  uuid NOT NULL,
    article_id   uuid NOT NULL,
    embedding_id uuid,
    chunk_index  integer,
    similarity   numeric(6,4),
    excerpt      text NOT NULL DEFAULT '',
    pinned       boolean NOT NULL DEFAULT false,
    pinned_by    uuid,
    created_at   timestamp with time zone NOT NULL DEFAULT now(),
    updated_at   timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at   timestamp with time zone,
    CONSTRAINT magistral_formula_template_articles_pkey PRIMARY KEY (id),
    CONSTRAINT fk_mfta_template FOREIGN KEY (template_id)
        REFERENCES public.magistral_formula_templates(id) ON DELETE CASCADE,
    CONSTRAINT fk_mfta_article FOREIGN KEY (article_id)
        REFERENCES public.articles(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_mfta
    ON public.magistral_formula_template_articles (template_id, article_id, chunk_index)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- A receita emitida registra de qual base saiu.
-- +goose StatementBegin
ALTER TABLE public.prescription_formulas DROP CONSTRAINT IF EXISTS fk_prescription_formulas_template;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescription_formulas
    ADD CONSTRAINT fk_prescription_formulas_template
    FOREIGN KEY (template_id) REFERENCES public.magistral_formula_templates(id);
-- +goose StatementEnd

-- O que a máquina sugeriu, ao lado do que foi prescrito.
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components
    ADD COLUMN IF NOT EXISTS suggested_quantity numeric(14,4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components DROP COLUMN IF EXISTS suggested_quantity;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescription_formulas DROP CONSTRAINT IF EXISTS fk_prescription_formulas_template;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_formula_template_articles;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_formula_template_rules;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_formula_template_components;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_formula_templates;
-- +goose StatementEnd
