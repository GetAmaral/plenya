-- +goose Up
-- +goose StatementBegin

-- Qualidade das regras de dose dinâmica. Três buracos, medidos no dado real:
--
--  1. UNIDADE. O motor comparava `lab_results.result_numeric` com o limiar da regra sem nunca
--     olhar em que unidade o resultado foi gravado — e 390 dos 1.243 resultados numéricos do
--     banco (31%) estão numa unidade diferente da definição do exame. Há cortisol em nmol/L
--     numa definição em µg/dL (fator 27,6) e uma 25-OH-vitamina D gravada em pg/mL. A regra
--     passa a guardar a unidade em que foi escrita, e o motor recusa sugerir quando a unidade
--     do resultado não bate.
--
--  2. FAIXA. Limiar binário não é como se dosa: vitamina D, ferritina e B12 têm conduta por
--     faixa. Uma regra por componente impedia escrever as três condutas. Entra o tipo
--     `lab_band`, com a MESMA convenção de faixa já canônica no escore: intervalo meio-aberto
--     (lower, upper], lower NULL = -infinito, upper NULL = +infinito.
--
--  3. ARREDONDAMENTO. Regra por peso devolvia 1.234,5678 mg, que ninguém manipula. `round_to`
--     é o passo prático da substância (100 UI, 50 mg, 5 mg).

ALTER TABLE magistral_formula_template_rules
    ADD COLUMN IF NOT EXISTS lab_unit  varchar(50),
    ADD COLUMN IF NOT EXISTS round_to  numeric(14,4);

COMMENT ON COLUMN magistral_formula_template_rules.lab_unit IS
    'Unidade em que o limiar/faixa foi escrito. Resultado em outra unidade nao vira sugestao.';
COMMENT ON COLUMN magistral_formula_template_rules.round_to IS
    'Passo pratico de arredondamento da dose sugerida (100 UI, 50 mg). NULL = sem arredondar.';

-- CHECKs do 00072: drop + recreate, mesmo padrao das migrations anteriores. Sao DOIS — o do
-- tipo e o do formato, que exige os campos certos para cada tipo.
ALTER TABLE magistral_formula_template_rules
    DROP CONSTRAINT IF EXISTS chk_mftr_kind;
ALTER TABLE magistral_formula_template_rules
    ADD CONSTRAINT chk_mftr_kind
    CHECK (kind IN ('fixed', 'per_kg', 'lab_threshold', 'lab_band'));

ALTER TABLE magistral_formula_template_rules
    DROP CONSTRAINT IF EXISTS chk_mftr_shape;
ALTER TABLE magistral_formula_template_rules
    ADD CONSTRAINT chk_mftr_shape
    CHECK (
        (kind = 'fixed'         AND fixed_dose IS NOT NULL)
     OR (kind = 'per_kg'        AND per_kg IS NOT NULL)
     OR (kind = 'lab_threshold' AND lab_code IS NOT NULL AND lab_operator IS NOT NULL
                                AND lab_threshold IS NOT NULL AND dose_if_true IS NOT NULL)
        -- lab_band: as faixas moram na tabela filha, entao o CHECK so alcanca o exame.
     OR (kind = 'lab_band'      AND lab_code IS NOT NULL)
    );

CREATE TABLE IF NOT EXISTS magistral_formula_template_rule_bands (
    id            uuid PRIMARY KEY,
    rule_id       uuid NOT NULL REFERENCES magistral_formula_template_rules(id) ON DELETE CASCADE,
    display_order integer NOT NULL DEFAULT 0,

    -- Intervalo meio-aberto (lower, upper], igual as faixas do escore.
    -- NULL em lower = sem piso; NULL em upper = sem teto.
    lower_bound   numeric(14,4),
    upper_bound   numeric(14,4),

    dose          numeric(14,4) NOT NULL,
    -- Como a faixa aparece na frase que o medico le ("deficiencia grave").
    label         varchar(120) NOT NULL DEFAULT '',

    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now(),
    deleted_at    timestamptz,

    CONSTRAINT chk_band_dose_positiva  CHECK (dose > 0),
    CONSTRAINT chk_band_intervalo      CHECK (lower_bound IS NULL OR upper_bound IS NULL OR lower_bound < upper_bound),
    -- Faixa sem piso E sem teto pegaria tudo: e regra fixa disfarcada.
    CONSTRAINT chk_band_nao_universal  CHECK (lower_bound IS NOT NULL OR upper_bound IS NOT NULL)
);

CREATE INDEX IF NOT EXISTS idx_rule_bands_rule
    ON magistral_formula_template_rule_bands (rule_id, display_order)
    WHERE deleted_at IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS magistral_formula_template_rule_bands;

-- Regras em faixa perderiam o sentido sem as faixas: saem junto, antes de restaurar os CHECKs.
DELETE FROM magistral_formula_template_rules WHERE kind = 'lab_band';

ALTER TABLE magistral_formula_template_rules
    DROP CONSTRAINT IF EXISTS chk_mftr_kind;
ALTER TABLE magistral_formula_template_rules
    ADD CONSTRAINT chk_mftr_kind
    CHECK (kind IN ('fixed', 'per_kg', 'lab_threshold'));

ALTER TABLE magistral_formula_template_rules
    DROP CONSTRAINT IF EXISTS chk_mftr_shape;
ALTER TABLE magistral_formula_template_rules
    ADD CONSTRAINT chk_mftr_shape
    CHECK (
        (kind = 'fixed'         AND fixed_dose IS NOT NULL)
     OR (kind = 'per_kg'        AND per_kg IS NOT NULL)
     OR (kind = 'lab_threshold' AND lab_code IS NOT NULL AND lab_operator IS NOT NULL
                                AND lab_threshold IS NOT NULL AND dose_if_true IS NOT NULL)
    );

ALTER TABLE magistral_formula_template_rules
    DROP COLUMN IF EXISTS lab_unit,
    DROP COLUMN IF EXISTS round_to;

-- +goose StatementEnd
