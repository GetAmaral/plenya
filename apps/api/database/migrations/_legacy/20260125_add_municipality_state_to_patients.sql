-- Add municipality and state columns to patients table for filtering

ALTER TABLE patients
ADD COLUMN municipality VARCHAR(100),
ADD COLUMN state VARCHAR(2);

-- Add indexes for performance on filtering
CREATE INDEX idx_patients_municipality ON patients(municipality);
CREATE INDEX idx_patients_state ON patients(state);

-- Add comments
COMMENT ON COLUMN patients.municipality IS 'City/municipality where the patient resides';
COMMENT ON COLUMN patients.state IS 'State/UF (2-letter code) where the patient resides';
