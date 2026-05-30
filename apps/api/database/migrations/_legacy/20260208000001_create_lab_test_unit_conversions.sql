-- Create lab_test_unit_conversions table
CREATE TABLE IF NOT EXISTS lab_test_unit_conversions (
    id UUID PRIMARY KEY,
    lab_test_definition_id UUID NOT NULL,
    main_unit VARCHAR(50) NOT NULL,
    secondary_unit VARCHAR(50) NOT NULL,
    conversion_factor DOUBLE PRECISION NOT NULL CHECK (conversion_factor > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_unit_conv_test
        FOREIGN KEY (lab_test_definition_id)
        REFERENCES lab_test_definitions(id)
        ON DELETE CASCADE
);

-- Create unique constraint (apenas uma conversão por unidade secundária, ignorando soft-deleted)
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_test_secondary_unit
    ON lab_test_unit_conversions(lab_test_definition_id, secondary_unit)
    WHERE deleted_at IS NULL;

-- Create index for performance
CREATE INDEX IF NOT EXISTS idx_unit_conv_test
    ON lab_test_unit_conversions(lab_test_definition_id, deleted_at);

-- Add comments
COMMENT ON TABLE lab_test_unit_conversions IS 'Conversões de unidades para exames laboratoriais';
COMMENT ON COLUMN lab_test_unit_conversions.main_unit IS 'Unidade principal/padrão (ex: g/dL)';
COMMENT ON COLUMN lab_test_unit_conversions.secondary_unit IS 'Unidade alternativa/secundária (ex: g/L)';
COMMENT ON COLUMN lab_test_unit_conversions.conversion_factor IS 'Fator: secondaryValue = mainValue * factor';
