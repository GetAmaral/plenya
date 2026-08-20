-- Catálogo de componentes magistrais + incompatibilidades.
--
-- PORQUÊ ESTE CATÁLOGO É DIFERENTE DO DE INDUSTRIALIZADOS: o da ANVISA a gente importa inteiro
-- (26 mil apresentações, migration 00066). Este NÃO tem equivalente público importável. Nome e
-- sinônimos existem na DCB da Anvisa; faixa de dose de INSUMO magistral, densidade aparente e
-- palatabilidade não existem em base aberta — bula dá dose de industrializado, que é outra coisa.
--
-- Por isso o desenho é CURADORIA OPORTUNISTA: todo campo clínico nasce NULL, a UI trata NULL como
-- "sem sugestão" (nunca como zero), e quando o médico prescreve uma dose para uma substância sem
-- faixa cadastrada a tela oferece salvar aquilo como padrão. Projeto de curadoria em lote nunca
-- termina; curadoria dentro do fluxo termina sozinha em alguns meses de uso.
--
-- `bulk_density` é o campo com consequência clínica mais direta: sem ele a calculadora de cápsula
-- NÃO OPINA (mostra "sem densidade cadastrada para X"). Silêncio ganha de número inventado — a
-- densidade aparente varia com lote, granulometria e compactação da farmácia, e o erro honesto do
-- cálculo é ±20-30%.
--
-- Cada linha carrega procedência (`source`, `reviewed_by`, `last_review`), no mesmo padrão de
-- `score_levels.last_review`: dado clínico sem dono é dado que ninguém confere.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_components (
    id               uuid NOT NULL,
    name             varchar(200) NOT NULL,
    synonyms         text NOT NULL DEFAULT '',
    cas              varchar(20),
    dcb_code         varchar(20),

    -- Unidade em que a substância é prescrita (mg, mcg, UI, %, mL)
    default_unit     varchar(20)  NOT NULL DEFAULT 'mg',

    -- Faixa de dose por unidade aviada. NULL = sem sugestão (a UI não inventa).
    usual_dose       numeric(14,4),
    min_dose         numeric(14,4),
    max_dose         numeric(14,4),

    -- Densidade aparente (g/mL) para a calculadora de cápsula. NULL = calculadora se cala.
    bulk_density     numeric(8,4),

    -- Sinalizadores que alimentam as regras de compatibilidade e de palatabilidade.
    eutectic_former     boolean NOT NULL DEFAULT false,
    hygroscopic         boolean NOT NULL DEFAULT false,
    oxidizing           boolean NOT NULL DEFAULT false,
    oxidation_sensitive boolean NOT NULL DEFAULT false,
    photosensitive      boolean NOT NULL DEFAULT false,
    -- 0 = sem amargor relevante · 1 = leve · 2 = marcante · 3 = intolerável em sachê
    bitterness       smallint,
    sachet_ok        boolean,

    notes            text,
    source           varchar(30)  NOT NULL DEFAULT 'manual',
    reviewed_by      uuid,
    last_review      timestamp with time zone,
    -- Quantas vezes já foi prescrita: ordena a busca pelo repertório real do prescritor.
    usage_count      integer      NOT NULL DEFAULT 0,
    is_active        boolean      NOT NULL DEFAULT true,

    created_at       timestamp with time zone NOT NULL DEFAULT now(),
    updated_at       timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at       timestamp with time zone,
    CONSTRAINT magistral_components_pkey PRIMARY KEY (id),
    CONSTRAINT chk_magistral_components_bitterness CHECK (bitterness IS NULL OR bitterness BETWEEN 0 AND 3),
    CONSTRAINT chk_magistral_components_density CHECK (bulk_density IS NULL OR bulk_density > 0),
    CONSTRAINT chk_magistral_components_dose_range CHECK (
        min_dose IS NULL OR max_dose IS NULL OR min_dose <= max_dose
    )
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_magistral_components_name
    ON public.magistral_components (lower(public.immutable_unaccent(name)))
    WHERE deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_components_trgm
    ON public.magistral_components USING gin (lower(public.immutable_unaccent(name)) gin_trgm_ops);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_components_deleted_at
    ON public.magistral_components (deleted_at);
-- +goose StatementEnd

-- Incompatibilidades que NÃO saem de flag (par específico, com mecanismo). Nasce pequena e
-- curada: um conjunto que nunca dá falso positivo constrói confiança; um grande e barulhento
-- vira ruído ignorado em duas semanas. Não existe base brasileira pronta para importar disto
-- (Trissel's é injetável, domínio errado).
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.magistral_incompatibilities (
    id             uuid NOT NULL,
    component_a_id uuid NOT NULL,
    component_b_id uuid NOT NULL,
    -- info = observação · warn = evitar salvo intenção explícita · avoid = não associar
    severity       varchar(10)  NOT NULL DEFAULT 'warn',
    mechanism      varchar(200) NOT NULL DEFAULT '',
    note           text,
    source         varchar(200) NOT NULL DEFAULT '',
    reviewed_by    uuid,
    last_review    timestamp with time zone,
    created_at     timestamp with time zone NOT NULL DEFAULT now(),
    updated_at     timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at     timestamp with time zone,
    CONSTRAINT magistral_incompatibilities_pkey PRIMARY KEY (id),
    CONSTRAINT fk_magistral_incompatibilities_a
        FOREIGN KEY (component_a_id) REFERENCES public.magistral_components(id) ON DELETE CASCADE,
    CONSTRAINT fk_magistral_incompatibilities_b
        FOREIGN KEY (component_b_id) REFERENCES public.magistral_components(id) ON DELETE CASCADE,
    CONSTRAINT chk_magistral_incompatibilities_severity CHECK (severity IN ('info','warn','avoid')),
    CONSTRAINT chk_magistral_incompatibilities_distinct CHECK (component_a_id <> component_b_id)
);
-- +goose StatementEnd

-- O par é simétrico: gravar sempre com o menor uuid em A evita a mesma regra cadastrada duas vezes.
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_magistral_incompatibilities_pair
    ON public.magistral_incompatibilities (component_a_id, component_b_id)
    WHERE deleted_at IS NULL;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_magistral_incompatibilities_b
    ON public.magistral_incompatibilities (component_b_id);
-- +goose StatementEnd

-- FK do componente da receita para o catálogo (a coluna já existia, sem referência).
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components
    DROP CONSTRAINT IF EXISTS fk_prescription_formula_components_magistral;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components
    ADD CONSTRAINT fk_prescription_formula_components_magistral
    FOREIGN KEY (magistral_component_id) REFERENCES public.magistral_components(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.prescription_formula_components
    DROP CONSTRAINT IF EXISTS fk_prescription_formula_components_magistral;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_incompatibilities;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.magistral_components;
-- +goose StatementEnd
