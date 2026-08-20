-- +goose Up
-- +goose StatementBegin

-- Incompatibilidade com a BASE (veiculo), que o modelo nao representava.
--
-- A tabela de pares cobre ativo x ativo. Mas o levantamento de 400 prescricoes magistrais da
-- farmacia-escola da UFRJ (Vigil. sanit. debate 2019;7(1):5-13) mostra a distribuicao real dos
-- erros farmacotecnicos: 63% sao ativo x FORMULACAO e 23% ativo x BASE, contra 13% ativo x ativo.
-- Ou seja: a maior parte do problema nao estava sendo representada.
--
-- Regra casa por texto porque o veiculo e campo livre na formula ("creme Lanette", "creme nao
-- ionico", "vaselina solida"). min_percent existe porque varias dessas regras sao de
-- concentracao, nao de presenca: acido acima de 10% ou ureia acima de 30% derrubam a emulsao
-- anionica, abaixo disso convivem.

CREATE TABLE IF NOT EXISTS magistral_base_incompatibilities (
    id               uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    -- Trecho a casar contra o veiculo escrito na formula, sem acento e sem caixa.
    base_pattern     varchar(120) NOT NULL,
    -- Trecho a casar contra o nome do componente. NULL = vale para qualquer componente que
    -- satisfaca as demais condicoes.
    substance_pattern varchar(120),
    -- Dispara so acima desta concentracao (em % na formula). NULL = a presenca ja basta.
    min_percent      numeric(6,2),
    severity         varchar(10)  NOT NULL DEFAULT 'warn',
    mechanism        text         NOT NULL,
    recommendation   text         NOT NULL DEFAULT '',
    source           varchar(300) NOT NULL DEFAULT '',
    is_active        boolean      NOT NULL DEFAULT true,
    created_at       timestamptz  NOT NULL DEFAULT now(),
    updated_at       timestamptz  NOT NULL DEFAULT now(),
    CONSTRAINT chk_base_incompat_sev CHECK (severity IN ('info', 'warn', 'avoid'))
);

CREATE INDEX IF NOT EXISTS idx_base_incompat_ativo
    ON magistral_base_incompatibilities (is_active);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS magistral_base_incompatibilities;

-- +goose StatementEnd
