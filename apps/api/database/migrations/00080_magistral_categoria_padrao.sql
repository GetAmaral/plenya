-- +goose Up
-- +goose StatementBegin

-- Categoria de receita que a substancia CARREGA.
--
-- Hoje a categoria mora no componente da formula e nasce sempre 'simple': quem digita
-- "testosterona" numa formula magistral emite receita simples por default, quando a Portaria
-- 344/98 pede Receituario de Controle Especial em duas vias (lista C5, anabolizantes). O
-- catalogo passa a dizer a categoria, e a tela preenche a partir dele.
--
-- E o mesmo desenho do resto: o catalogo sabe, a tela sugere, o medico decide.

ALTER TABLE magistral_components
    ADD COLUMN IF NOT EXISTS default_category varchar(20) NOT NULL DEFAULT 'simple',
    -- Nota regulatoria que acompanha a substancia (restricao de finalidade, exigencia de exame).
    ADD COLUMN IF NOT EXISTS regulatory_note text;

ALTER TABLE magistral_components
    DROP CONSTRAINT IF EXISTS chk_magistral_default_category;
ALTER TABLE magistral_components
    ADD CONSTRAINT chk_magistral_default_category
    CHECK (default_category IN ('simple', 'c1', 'c5', 'antibiotic', 'glp1', 'a_b'));

COMMENT ON COLUMN magistral_components.default_category IS
    'Categoria de receita da substancia. A tela preenche a categoria do componente a partir daqui.';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE magistral_components
    DROP CONSTRAINT IF EXISTS chk_magistral_default_category;
ALTER TABLE magistral_components
    DROP COLUMN IF EXISTS default_category,
    DROP COLUMN IF EXISTS regulatory_note;

-- +goose StatementEnd
