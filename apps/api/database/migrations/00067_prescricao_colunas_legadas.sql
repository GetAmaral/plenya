-- Prescrições: remover as colunas do modelo FLAT que ainda travavam todo INSERT.
--
-- SINTOMA: "sempre dá erro na hora de salvar" a receita, em produção. Não era a UI nem o payload:
-- a tabela `prescriptions` guardava, do desenho antigo (1 receita = 1 medicamento), nove colunas
-- NOT NULL SEM DEFAULT — medication_name, active_ingredient, concentration, dosage, frequency,
-- route, duration, quantity, quantity_in_words. O model Go deixou de ter esses campos quando a
-- receita passou a ter N medicamentos em `prescription_medications`, então o INSERT do GORM nunca
-- mandava nada para elas e o Postgres recusava:
--
--     ERROR: null value in column "medication_name" of relation "prescriptions"
--            violates not-null constraint (SQLSTATE 23502)
--
-- Por isso `prescriptions` tem 0 linhas em produção: nenhuma receita jamais foi gravada.
--
-- SEGURANÇA DO DROP: as colunas estão órfãs de código (nenhum model, service, handler ou query as
-- referencia) e órfãs de dado (0 linhas em prod e nenhuma linha com conteúdo em dev). O dado vivo
-- de medicamento mora em `prescription_medications` desde a migração para receita com N itens.
-- `medication_definition_id`, `category` e `instructions` vêm do mesmo desenho flat e saem junto —
-- deixar meia estrutura antiga é o que produz o próximo bug silencioso.
--
-- O Down recria as colunas com DEFAULT '' (e 0 nos inteiros) em vez de NOT NULL sem default: assim
-- o rollback devolve o formato antigo sem reintroduzir o defeito que esta migration corrige.

-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.prescriptions
    DROP COLUMN IF EXISTS medication_definition_id,
    DROP COLUMN IF EXISTS medication_name,
    DROP COLUMN IF EXISTS active_ingredient,
    DROP COLUMN IF EXISTS category,
    DROP COLUMN IF EXISTS concentration,
    DROP COLUMN IF EXISTS dosage,
    DROP COLUMN IF EXISTS frequency,
    DROP COLUMN IF EXISTS route,
    DROP COLUMN IF EXISTS duration,
    DROP COLUMN IF EXISTS quantity,
    DROP COLUMN IF EXISTS quantity_in_words,
    DROP COLUMN IF EXISTS instructions;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.prescriptions
    ADD COLUMN IF NOT EXISTS medication_definition_id uuid,
    ADD COLUMN IF NOT EXISTS medication_name          varchar(200) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS active_ingredient        varchar(200) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS category                 varchar(20)  NOT NULL DEFAULT 'simple',
    ADD COLUMN IF NOT EXISTS concentration            varchar(100) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS dosage                   varchar(100) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS frequency                varchar(100) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS route                    varchar(50)  NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS duration                 integer      NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS quantity                 integer      NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS quantity_in_words        varchar(200) NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS instructions             text;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_prescriptions_category ON public.prescriptions USING btree (category);
-- +goose StatementEnd
