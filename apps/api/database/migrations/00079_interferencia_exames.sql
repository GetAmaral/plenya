-- +goose Up
-- +goose StatementBegin

-- Substancia que ATRAPALHA EXAME.
--
-- O caso que motivou: a formula capilar do material de GLP-1 leva biotina 10 mg. Biotina a partir
-- de 5 mg/dia interfere em imunoensaios biotinilados — TSH, troponina, hormonios tireoidianos,
-- hCG — e pode dar resultado falsamente alto OU falsamente baixo conforme o formato do ensaio
-- (FDA 2017/2019; documento de orientacao da AACC). O paciente segue medicado por um numero que
-- nao e dele.
--
-- Isso importa DUAS vezes neste sistema: a receita sai daqui, e o exame que ela corrompe volta
-- para ca e alimenta as regras de dose dinamica. Um EMR que prescreve e le exame precisa saber
-- que uma coisa contamina a outra.

ALTER TABLE magistral_components
    ADD COLUMN IF NOT EXISTS assay_interference text,
    -- Dose diaria, na unidade da substancia, a partir da qual a interferencia e descrita.
    -- NULL com texto preenchido = interfere em qualquer dose.
    ADD COLUMN IF NOT EXISTS assay_interference_dose numeric(14,4);

COMMENT ON COLUMN magistral_components.assay_interference IS
    'Como esta substancia atrapalha exame laboratorial, e o que fazer antes da coleta.';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE magistral_components
    DROP COLUMN IF EXISTS assay_interference,
    DROP COLUMN IF EXISTS assay_interference_dose;

-- +goose StatementEnd
