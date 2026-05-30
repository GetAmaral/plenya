-- Add original value columns to lab_results
ALTER TABLE lab_results
    ADD COLUMN IF NOT EXISTS result_numeric_original DECIMAL(12,4),
    ADD COLUMN IF NOT EXISTS unit_original VARCHAR(50);

-- Add comments
COMMENT ON COLUMN lab_results.result_numeric IS 'Valor numérico CONVERTIDO para unidade padrão (ou original se sem conversão)';
COMMENT ON COLUMN lab_results.unit IS 'Unidade CONVERTIDA para unidade padrão (ou original se sem conversão)';
COMMENT ON COLUMN lab_results.result_numeric_original IS 'Valor numérico ORIGINAL (antes da conversão)';
COMMENT ON COLUMN lab_results.unit_original IS 'Unidade ORIGINAL (antes da conversão)';
