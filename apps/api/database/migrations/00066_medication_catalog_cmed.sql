-- Catálogo de medicamentos: preparar medication_definitions para a lista CMED/ANVISA.
--
-- PORQUÊ: a tabela está VAZIA (0 linhas em produção) e o médico digita nome, princípio ativo,
-- concentração, forma e via à mão em CADA receita. A CMED publica mensalmente a Lista de Preços
-- de Medicamentos (PMC) com ~25.570 apresentações — a única base pública, sem chave de API, que
-- cobre o mercado brasileiro inteiro. Esta migration acrescenta o que essa lista traz SEM
-- descaracterizar o que a tabela já é: um catálogo de REGRAS DE PRESCRIÇÃO (validade, retenção,
-- assinatura, SNCR).
--
-- O QUE A CMED NÃO TRAZ, e por isso o import NÃO preenche:
--   · listas da Portaria SVS/MS 344/98 (A1/A2/A3/B1/B2/C1..C5) → `control_list` nasce NULL e só
--     curadoria manual preenche. A TARJA é o proxy mais próximo e vai gravada crua em `stripe`.
--   · forma farmacêutica, via e concentração normalizadas → derivadas do texto de APRESENTAÇÃO
--     pelo importador, com o grau de confiança registrado em `derivation_confidence`.
--   · posologia, bula, interações → fora de escopo.
--
-- Aditiva: nenhuma coluna nova é NOT NULL sem default, então as linhas criadas à mão pelo admin
-- continuam válidas e imunes ao import (ficam com ggrem NULL e source='manual').

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    ADD COLUMN IF NOT EXISTS ggrem            varchar(20),
    ADD COLUMN IF NOT EXISTS source           varchar(10) NOT NULL DEFAULT 'manual',
    ADD COLUMN IF NOT EXISTS source_version   varchar(8),
    ADD COLUMN IF NOT EXISTS last_imported_at timestamptz,
    ADD COLUMN IF NOT EXISTS is_active        boolean     NOT NULL DEFAULT true;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    ADD COLUMN IF NOT EXISTS presentation           text,
    ADD COLUMN IF NOT EXISTS laboratory             varchar(200),
    ADD COLUMN IF NOT EXISTS product_type           varchar(60),
    ADD COLUMN IF NOT EXISTS ean13                  varchar(14),
    ADD COLUMN IF NOT EXISTS therapeutic_class      varchar(200),
    ADD COLUMN IF NOT EXISTS therapeutic_class_code varchar(10),
    ADD COLUMN IF NOT EXISTS stripe                 varchar(20),
    ADD COLUMN IF NOT EXISTS pmc_price              numeric(12,2);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    ADD COLUMN IF NOT EXISTS concentration         varchar(120),
    ADD COLUMN IF NOT EXISTS pharmaceutical_form   varchar(60),
    ADD COLUMN IF NOT EXISTS route                 varchar(40),
    ADD COLUMN IF NOT EXISTS package_quantity      integer,
    ADD COLUMN IF NOT EXISTS derivation_confidence varchar(10);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    ADD COLUMN IF NOT EXISTS control_list    varchar(4),
    ADD COLUMN IF NOT EXISTS category_source varchar(14) NOT NULL DEFAULT 'manual',
    ADD COLUMN IF NOT EXISTS needs_review    boolean     NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS is_prescribable boolean     NOT NULL DEFAULT true,
    ADD COLUMN IF NOT EXISTS curated_at      timestamptz,
    ADD COLUMN IF NOT EXISTS curated_by      uuid;
-- +goose StatementEnd

-- A coluna SUBSTÂNCIA da CMED chega a 1.419 caracteres em associações (ativos separados por ";").
-- varchar(500) estouraria no import.
-- +goose StatementBegin
ALTER TABLE public.medication_definitions ALTER COLUMN active_ingredient TYPE text;
-- +goose StatementEnd

-- 'a_b' = tarja preta (Notificação de Receita A amarela / B azul da Portaria 344). O EMR NÃO
-- emite Notificação de Receita; essas apresentações entram no catálogo (o paciente as USA, e
-- isso serve à reconciliação de medicação em uso) mas com is_prescribable = false.
-- +goose StatementBegin
ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_category;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_category
    CHECK (category IN ('simple','c1','c5','antibiotic','glp1','a_b'));
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_source;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_source
    CHECK (source IN ('manual','cmed'));

ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_category_source;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_category_source
    CHECK (category_source IN ('manual','cmed_derived','cmed_fallback'));

ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_stripe;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_stripe
    CHECK (stripe IS NULL OR stripe IN ('vermelha','vermelha_restrita','preta','isento'));

ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_control_list;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_control_list
    CHECK (control_list IS NULL OR control_list IN ('A1','A2','A3','B1','B2','C1','C2','C3','C4','C5'));

ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_derivation;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_derivation
    CHECK (derivation_confidence IS NULL OR derivation_confidence IN ('high','medium','none'));
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS fk_medication_definitions_curated_by;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT fk_medication_definitions_curated_by
    FOREIGN KEY (curated_by) REFERENCES public.users(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- Chave de idempotência do reimport mensal. Parcial: linhas manuais (ggrem NULL) não colidem.
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS uq_medication_definitions_ggrem
    ON public.medication_definitions (ggrem) WHERE ggrem IS NOT NULL;
-- +goose StatementEnd

-- Busca: mesmo padrão já usado em `articles` — GIN trigram sobre lower(immutable_unaccent(col))
-- para '%termo%', e btree text_pattern_ops para 'termo%' (o que um autocomplete realmente quer).
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_meddef_common_name_trgm
    ON public.medication_definitions
    USING gin (lower(public.immutable_unaccent((common_name)::text)) public.gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_meddef_active_ingredient_trgm
    ON public.medication_definitions
    USING gin (lower(public.immutable_unaccent(active_ingredient)) public.gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_meddef_common_name_prefix
    ON public.medication_definitions
    (lower(public.immutable_unaccent((common_name)::text)) text_pattern_ops);

CREATE INDEX IF NOT EXISTS idx_meddef_class_code
    ON public.medication_definitions (therapeutic_class_code);

CREATE INDEX IF NOT EXISTS idx_meddef_review
    ON public.medication_definitions (needs_review) WHERE needs_review;
-- +goose StatementEnd

-- +goose StatementBegin
COMMENT ON COLUMN public.medication_definitions.ggrem IS
  'CÓDIGO GGREM da CMED — chave natural do reimport mensal. NULL em linhas criadas à mão.';
COMMENT ON COLUMN public.medication_definitions.stripe IS
  'TARJA como a CMED publica. NÃO é a Portaria 344 — é o proxy mais próximo que a fonte oferece. NULL quando a CMED publica "- (*)" (~18% das linhas).';
COMMENT ON COLUMN public.medication_definitions.control_list IS
  'Lista da Portaria SVS/MS 344/98 (A1..C5). A CMED NÃO traz este dado: fica NULL até curadoria manual. Quando preenchida, manda sobre category.';
COMMENT ON COLUMN public.medication_definitions.category_source IS
  'manual = curado por humano · cmed_derived = derivado com regra defensável · cmed_fallback = chute conservador (a fonte não permitia afirmar).';
COMMENT ON COLUMN public.medication_definitions.needs_review IS
  'Classificação derivada de forma imperfeita. A UI deve avisar antes de prescrever.';
COMMENT ON COLUMN public.medication_definitions.is_prescribable IS
  'false = aparece no catálogo (reconciliação de medicação em uso) mas NÃO no autocomplete de receita. Usado em tarja preta: o EMR não emite Notificação de Receita A/B.';
COMMENT ON COLUMN public.medication_definitions.curated_at IS
  'Marca a edição manual do médico. Enquanto NOT NULL, o reimport mensal atualiza só os campos de FONTE e não encosta nos clínicos.';
COMMENT ON COLUMN public.medication_definitions.source_version IS
  'Edição da CMED que produziu a linha (YYYYMM). Serve para detectar quem sumiu da lista nova e para datar o preço.';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS public.idx_meddef_review;
DROP INDEX IF EXISTS public.idx_meddef_class_code;
DROP INDEX IF EXISTS public.idx_meddef_common_name_prefix;
DROP INDEX IF EXISTS public.idx_meddef_active_ingredient_trgm;
DROP INDEX IF EXISTS public.idx_meddef_common_name_trgm;
DROP INDEX IF EXISTS public.uq_medication_definitions_ggrem;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    DROP CONSTRAINT IF EXISTS fk_medication_definitions_curated_by,
    DROP CONSTRAINT IF EXISTS chk_medication_definitions_derivation,
    DROP CONSTRAINT IF EXISTS chk_medication_definitions_control_list,
    DROP CONSTRAINT IF EXISTS chk_medication_definitions_stripe,
    DROP CONSTRAINT IF EXISTS chk_medication_definitions_category_source,
    DROP CONSTRAINT IF EXISTS chk_medication_definitions_source;
-- +goose StatementEnd

-- Down é destrutivo por natureza: apaga o catálogo importado e trunca princípios ativos longos.
-- Só faz sentido rodar antes do primeiro import valer alguma coisa.
-- +goose StatementBegin
DELETE FROM public.medication_definitions WHERE source = 'cmed';

ALTER TABLE public.medication_definitions DROP CONSTRAINT IF EXISTS chk_medication_definitions_category;
ALTER TABLE public.medication_definitions
    ADD CONSTRAINT chk_medication_definitions_category
    CHECK (category IN ('simple','c1','c5','antibiotic','glp1'));

ALTER TABLE public.medication_definitions
    ALTER COLUMN active_ingredient TYPE varchar(500) USING left(active_ingredient, 500);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.medication_definitions
    DROP COLUMN IF EXISTS curated_by,            DROP COLUMN IF EXISTS curated_at,
    DROP COLUMN IF EXISTS is_prescribable,       DROP COLUMN IF EXISTS needs_review,
    DROP COLUMN IF EXISTS category_source,       DROP COLUMN IF EXISTS control_list,
    DROP COLUMN IF EXISTS derivation_confidence, DROP COLUMN IF EXISTS package_quantity,
    DROP COLUMN IF EXISTS route,                 DROP COLUMN IF EXISTS pharmaceutical_form,
    DROP COLUMN IF EXISTS concentration,         DROP COLUMN IF EXISTS pmc_price,
    DROP COLUMN IF EXISTS stripe,                DROP COLUMN IF EXISTS therapeutic_class_code,
    DROP COLUMN IF EXISTS therapeutic_class,     DROP COLUMN IF EXISTS ean13,
    DROP COLUMN IF EXISTS product_type,          DROP COLUMN IF EXISTS laboratory,
    DROP COLUMN IF EXISTS presentation,          DROP COLUMN IF EXISTS is_active,
    DROP COLUMN IF EXISTS last_imported_at,      DROP COLUMN IF EXISTS source_version,
    DROP COLUMN IF EXISTS source,                DROP COLUMN IF EXISTS ggrem;
-- +goose StatementEnd
