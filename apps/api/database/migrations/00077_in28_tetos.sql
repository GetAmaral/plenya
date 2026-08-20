-- +goose Up
-- +goose StatementBegin

-- Tetos diarios do Anexo IV da IN 28/2018 (texto consolidado, com as alteracoes ate a IN
-- 373/2025), guardados como TABELA DE REFERENCIA e nao espalhados pelo codigo: sao numeros de
-- norma, mudam por instrucao normativa nova, e quem confere precisa ver a fonte ao lado.
--
-- FRONTEIRA QUE IMPORTA: este teto vale para SUPLEMENTO ALIMENTAR. Formula magistral prescrita
-- passa dele legitimamente — B12 de 1.000 mcg e vitamina D de 7.000 UI sao conduta comum e ficam
-- 100x e 3,5x acima do Anexo IV. Por isso o alerta que nasce daqui e informativo: diz que a
-- formula saiu do territorio de suplemento e virou prescricao, o que e informacao real para o
-- prescritor e para a farmacia, nao um impedimento.

CREATE TABLE IF NOT EXISTS in28_limits (
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    -- Nomes de probioticos do Anexo IV passam de 200 caracteres (cepas com ATCC no nome).
    nutrient   varchar(400) NOT NULL UNIQUE,
    unit       varchar(20)  NOT NULL,
    -- Coluna ">=19 anos". NULL quando a norma diz NE (nao estabelecido) ou NA (nao se aplica).
    max_adult  numeric(14,4),
    kind       varchar(10)  NOT NULL DEFAULT 'valor',
    source     varchar(200) NOT NULL DEFAULT 'IN 28/2018 Anexo IV, texto consolidado',
    created_at timestamptz  NOT NULL DEFAULT now(),
    updated_at timestamptz  NOT NULL DEFAULT now(),
    CONSTRAINT chk_in28_kind CHECK (kind IN ('valor', 'NE', 'NA'))
);

CREATE INDEX IF NOT EXISTS idx_in28_nutrient ON in28_limits (lower(nutrient));

-- Ligacao do catalogo com a norma. Varias substancias apontam para o MESMO nutriente de
-- proposito: e assim que piridoxal-5-fosfato e vitamina B6 na mesma formula somam antes de
-- comparar com o teto, que foi exatamente o erro encontrado no formulario das parceiras.
ALTER TABLE magistral_components
    ADD COLUMN IF NOT EXISTS in28_nutrient varchar(400),
    -- Quantas unidades do Anexo IV valem UMA unidade desta substancia: 1 quando a unidade e a
    -- mesma, 1000 de mg para µg, 0,025 de UI de vitamina D para µg.
    ADD COLUMN IF NOT EXISTS in28_factor numeric(14,6) NOT NULL DEFAULT 1;

COMMENT ON COLUMN magistral_components.in28_nutrient IS
    'Nutriente correspondente no Anexo IV da IN 28. Varias substancias podem apontar para o mesmo.';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE magistral_components
    DROP COLUMN IF EXISTS in28_nutrient,
    DROP COLUMN IF EXISTS in28_factor;

DROP TABLE IF EXISTS in28_limits;

-- +goose StatementEnd
