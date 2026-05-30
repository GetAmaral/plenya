-- Create lab_requests table
CREATE TABLE IF NOT EXISTS lab_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    exams TEXT NOT NULL,
    notes TEXT,
    doctor_id UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_lab_requests_patient_id ON lab_requests(patient_id);
CREATE INDEX idx_lab_requests_date ON lab_requests(date);
CREATE INDEX idx_lab_requests_doctor_id ON lab_requests(doctor_id);
CREATE INDEX idx_lab_requests_deleted_at ON lab_requests(deleted_at);

-- Comment
COMMENT ON TABLE lab_requests IS 'Pedidos de exames laboratoriais';
