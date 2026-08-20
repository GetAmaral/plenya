-- Receita de manipulado (fórmula magistral).
--
-- PORQUÊ: o modelo só sabe "1 linha = 1 industrializado". Uma receita magistral é N fórmulas ×
-- M substâncias, cada fórmula com forma farmacêutica, veículo, quantidade a aviar e posologia
-- próprias. Hoje isso não existe e o médico prescreve manipulado fora do sistema.
--
-- DESENHO: o manipulado continua sendo UMA LINHA em `prescriptions`, com `type='compounded'`.
-- Isso herda sem código novo sete integrações que já funcionam: publicação como PatientDocument,
-- download autenticado, link público por token, envio por WhatsApp, validação por QR, portal do
-- paciente e export LGPD. Entidade separada obrigaria a reimplementar as sete.
--
-- Componentes em TABELA, não em JSONB: a migration 00060 documenta uma linha JSONB malformada
-- derrubando a leitura da tabela inteira pelo serializer do GORM. Não repetir.
--
-- `highest_category` é denormalizado na fórmula por um motivo prático: é ele que decide modo de
-- assinatura (controlado => manual) e rótulo de Controle Especial, e recalcular isso percorrendo
-- componentes em toda leitura de PDF é caro e fácil de esquecer num caminho novo.
--
-- `template_id` nasce sem FK: a tabela de fórmulas-base (com faixa de dose e evidência do RAG)
-- vem numa fase seguinte. A coluna existe desde já para que a receita emitida registre de qual
-- fórmula-base ela saiu — informação que não dá para reconstruir depois.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.prescriptions
    ADD COLUMN IF NOT EXISTS type varchar(12) NOT NULL DEFAULT 'commercial';
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE public.prescriptions DROP CONSTRAINT IF EXISTS chk_prescriptions_type;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescriptions
    ADD CONSTRAINT chk_prescriptions_type CHECK (type IN ('commercial','compounded'));
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.prescription_formulas (
    id                    uuid NOT NULL,
    prescription_id       uuid NOT NULL,
    template_id           uuid,
    display_order         integer      NOT NULL DEFAULT 0,
    name                  varchar(200) NOT NULL DEFAULT '',
    pharmaceutical_form   varchar(60)  NOT NULL,
    usage_type            varchar(10)  NOT NULL DEFAULT 'internal',
    route                 varchar(40)  NOT NULL DEFAULT '',
    vehicle               varchar(200) NOT NULL DEFAULT '',
    quantity_to_dispense  numeric(12,3) NOT NULL,
    quantity_unit         varchar(30)  NOT NULL,
    quantity_in_words     varchar(200) NOT NULL DEFAULT '',
    posology              varchar(300) NOT NULL DEFAULT '',
    duration              integer      NOT NULL DEFAULT 0,
    instructions          text,
    highest_category      varchar(20)  NOT NULL DEFAULT 'simple',
    created_at            timestamp with time zone NOT NULL DEFAULT now(),
    updated_at            timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at            timestamp with time zone,
    CONSTRAINT prescription_formulas_pkey PRIMARY KEY (id),
    CONSTRAINT fk_prescription_formulas_prescription
        FOREIGN KEY (prescription_id) REFERENCES public.prescriptions(id) ON DELETE CASCADE,
    CONSTRAINT chk_prescription_formulas_usage CHECK (usage_type IN ('internal','external')),
    CONSTRAINT chk_prescription_formulas_quantity CHECK (quantity_to_dispense > 0)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescription_formulas_prescription_id
    ON public.prescription_formulas (prescription_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescription_formulas_deleted_at
    ON public.prescription_formulas (deleted_at);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.prescription_formula_components (
    id                       uuid NOT NULL,
    formula_id               uuid NOT NULL,
    magistral_component_id   uuid,
    medication_definition_id uuid,
    display_order            integer      NOT NULL DEFAULT 0,
    substance                varchar(200) NOT NULL,
    quantity                 numeric(14,4) NOT NULL,
    unit                     varchar(20)  NOT NULL,
    category                 varchar(20)  NOT NULL DEFAULT 'simple',
    note                     varchar(200) NOT NULL DEFAULT '',
    created_at               timestamp with time zone NOT NULL DEFAULT now(),
    updated_at               timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at               timestamp with time zone,
    CONSTRAINT prescription_formula_components_pkey PRIMARY KEY (id),
    CONSTRAINT fk_prescription_formula_components_formula
        FOREIGN KEY (formula_id) REFERENCES public.prescription_formulas(id) ON DELETE CASCADE,
    CONSTRAINT fk_prescription_formula_components_medication_definition
        FOREIGN KEY (medication_definition_id) REFERENCES public.medication_definitions(id),
    CONSTRAINT chk_prescription_formula_components_quantity CHECK (quantity > 0)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescription_formula_components_formula_id
    ON public.prescription_formula_components (formula_id);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescription_formula_components_deleted_at
    ON public.prescription_formula_components (deleted_at);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescription_formula_components_substance
    ON public.prescription_formula_components (lower(substance));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.prescription_formula_components;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE IF EXISTS public.prescription_formulas;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescriptions DROP CONSTRAINT IF EXISTS chk_prescriptions_type;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE public.prescriptions DROP COLUMN IF EXISTS type;
-- +goose StatementEnd
