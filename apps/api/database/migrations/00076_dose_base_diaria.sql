-- +goose Up
-- +goose StatementBegin

-- A faixa de dose do catalogo nunca disse se e POR TOMADA ou POR DIA, e as duas leituras
-- convivem hoje. Medido: em 7 das 71 substancias com faixa numerica e texto de posologia, a
-- faixa numerica nao encosta no que o proprio texto da linha diz — gimnema cadastrada em 12,5 a
-- 25 mg com o texto dizendo 300 mg, alfa-lipoico em 25 a 200 mg com o texto dizendo 300 a 600.
--
-- A origem e conhecida: a faixa numerica veio das formulas das parceiras (dose de UMA capsula) e
-- o texto veio da literatura (dose do DIA). Sem declarar a base, o painel comparava a dose de uma
-- capsula contra uma faixa diaria em algumas substancias e contra faixa por tomada em outras.
--
-- Default 'por_dia' porque e como a literatura fala e como a maioria das linhas ja estava.

ALTER TABLE magistral_components
    ADD COLUMN IF NOT EXISTS dose_basis varchar(10) NOT NULL DEFAULT 'por_dia';

ALTER TABLE magistral_components
    DROP CONSTRAINT IF EXISTS chk_magistral_dose_basis;
ALTER TABLE magistral_components
    ADD CONSTRAINT chk_magistral_dose_basis
    CHECK (dose_basis IN ('por_dia', 'por_dose'));

COMMENT ON COLUMN magistral_components.dose_basis IS
    'Se min_dose/max_dose/usual_dose sao por dia ou por tomada. Sem isto o painel compara capsula com dose diaria.';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE magistral_components
    DROP CONSTRAINT IF EXISTS chk_magistral_dose_basis;
ALTER TABLE magistral_components
    DROP COLUMN IF EXISTS dose_basis;

-- +goose StatementEnd
