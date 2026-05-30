-- Migration: Digital Prescription Foundation (Phase 1)
-- Description: Add professional fields to users, create medication_definitions table,
--              update prescriptions table for digital signatures, add signature fields to lab_requests
-- Date: 2026-02-03

-- ============================================================================
-- 1. Add professional and certificate fields to users table
-- ============================================================================

ALTER TABLE users ADD COLUMN IF NOT EXISTS crm VARCHAR(20);
ALTER TABLE users ADD COLUMN IF NOT EXISTS crm_uf VARCHAR(2);
ALTER TABLE users ADD COLUMN IF NOT EXISTS rqe VARCHAR(20);
ALTER TABLE users ADD COLUMN IF NOT EXISTS specialty VARCHAR(100);
ALTER TABLE users ADD COLUMN IF NOT EXISTS professional_address TEXT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS professional_phone VARCHAR(20);
ALTER TABLE users ADD COLUMN IF NOT EXISTS certificate_pfx TEXT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS certificate_password TEXT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS certificate_expiry DATE;
ALTER TABLE users ADD COLUMN IF NOT EXISTS certificate_serial VARCHAR(100);

-- Add index on CRM for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_crm ON users(crm);

-- ============================================================================
-- 2. Create medication_definitions table
-- ============================================================================

CREATE TABLE IF NOT EXISTS medication_definitions (
    id UUID PRIMARY KEY,
    common_name VARCHAR(500) NOT NULL,
    active_ingredient VARCHAR(500) NOT NULL,
    category VARCHAR(20) NOT NULL CHECK (category IN ('simple', 'c1', 'c5', 'antibiotic', 'glp1')),
    validity_days INTEGER NOT NULL DEFAULT 30,
    max_per_prescription INTEGER NOT NULL DEFAULT 10,
    max_treatment_days INTEGER NOT NULL DEFAULT 60,
    requires_digital_signature BOOLEAN NOT NULL DEFAULT FALSE,
    requires_sncr BOOLEAN NOT NULL DEFAULT FALSE,
    anvisa_code VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Indexes for medication_definitions
CREATE INDEX IF NOT EXISTS idx_medication_definitions_common_name ON medication_definitions(common_name);
CREATE INDEX IF NOT EXISTS idx_medication_definitions_category ON medication_definitions(category);
CREATE INDEX IF NOT EXISTS idx_medication_definitions_deleted_at ON medication_definitions(deleted_at);

-- ============================================================================
-- 3. Update prescriptions table (drop and recreate to avoid conflicts)
-- ============================================================================

-- First, backup existing data if table exists
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'prescriptions') THEN
        -- Create backup table
        CREATE TABLE IF NOT EXISTS prescriptions_backup_20260203 AS SELECT * FROM prescriptions;
    END IF;
END $$;

-- Drop existing prescriptions table
DROP TABLE IF EXISTS prescriptions CASCADE;

-- Recreate prescriptions table with new structure
CREATE TABLE prescriptions (
    id UUID PRIMARY KEY,

    -- Basic data
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    doctor_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,

    -- Medication
    medication_definition_id UUID REFERENCES medication_definitions(id) ON DELETE SET NULL,
    medication_name VARCHAR(200) NOT NULL,
    active_ingredient VARCHAR(200) NOT NULL,
    category VARCHAR(20) NOT NULL DEFAULT 'simple' CHECK (category IN ('simple', 'c1', 'c5', 'antibiotic', 'glp1')),

    -- Prescription details
    concentration VARCHAR(100) NOT NULL,
    dosage VARCHAR(100) NOT NULL,
    frequency VARCHAR(100) NOT NULL,
    route VARCHAR(50) NOT NULL,
    duration INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    quantity_in_words VARCHAR(200) NOT NULL,
    instructions TEXT,

    -- Dates
    prescription_date TIMESTAMP NOT NULL,
    valid_until DATE NOT NULL,

    -- SNCR (Sistema Nacional de Controle de Receitas)
    sncr_number VARCHAR(50) UNIQUE,
    sncr_status VARCHAR(20),

    -- Digital Signature (ICP-Brasil)
    signed_pdf_path VARCHAR(500),
    signed_pdf_hash VARCHAR(64),
    qr_code_data TEXT,
    signed_at TIMESTAMP,
    certificate_serial VARCHAR(100),

    -- Status and control
    status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'completed', 'cancelled', 'expired')),
    is_used BOOLEAN NOT NULL DEFAULT FALSE,
    dispensed_at TIMESTAMP,

    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- Indexes for prescriptions
CREATE INDEX IF NOT EXISTS idx_prescriptions_patient_id ON prescriptions(patient_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_doctor_id ON prescriptions(doctor_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_medication_definition_id ON prescriptions(medication_definition_id);
CREATE INDEX IF NOT EXISTS idx_prescriptions_category ON prescriptions(category);
CREATE INDEX IF NOT EXISTS idx_prescriptions_prescription_date ON prescriptions(prescription_date);
CREATE INDEX IF NOT EXISTS idx_prescriptions_valid_until ON prescriptions(valid_until);
CREATE INDEX IF NOT EXISTS idx_prescriptions_sncr_number ON prescriptions(sncr_number);
CREATE INDEX IF NOT EXISTS idx_prescriptions_is_used ON prescriptions(is_used);
CREATE INDEX IF NOT EXISTS idx_prescriptions_deleted_at ON prescriptions(deleted_at);

-- ============================================================================
-- 4. Add digital signature fields to lab_requests table
-- ============================================================================

ALTER TABLE lab_requests ADD COLUMN IF NOT EXISTS signed_pdf_path VARCHAR(500);
ALTER TABLE lab_requests ADD COLUMN IF NOT EXISTS signed_pdf_hash VARCHAR(64);
ALTER TABLE lab_requests ADD COLUMN IF NOT EXISTS qr_code_data TEXT;
ALTER TABLE lab_requests ADD COLUMN IF NOT EXISTS signed_at TIMESTAMP;
ALTER TABLE lab_requests ADD COLUMN IF NOT EXISTS certificate_serial VARCHAR(100);

-- ============================================================================
-- 5. Add comments for documentation
-- ============================================================================

COMMENT ON COLUMN users.crm IS 'Número do CRM (Conselho Regional de Medicina)';
COMMENT ON COLUMN users.crm_uf IS 'UF do CRM (ex: SP, RJ, MG)';
COMMENT ON COLUMN users.certificate_pfx IS 'Certificado digital A1 (criptografado em Base64)';
COMMENT ON COLUMN users.certificate_password IS 'Senha do certificado A1 (criptografada)';

COMMENT ON TABLE medication_definitions IS 'Catálogo de medicamentos com regras regulatórias (ANVISA/CFM)';
COMMENT ON COLUMN medication_definitions.category IS 'Categoria regulatória: simple (receita simples), c1 (controle especial), c5 (psicotrópicos), antibiotic (antimicrobianos), glp1 (GLP-1 agonistas)';
COMMENT ON COLUMN medication_definitions.requires_sncr IS 'Se requer registro no SNCR (Sistema Nacional de Controle de Receitas) da ANVISA';

COMMENT ON TABLE prescriptions IS 'Prescrições médicas digitais com assinatura ICP-Brasil';
COMMENT ON COLUMN prescriptions.sncr_number IS 'Número SNCR gerado pela ANVISA (ou stub)';
COMMENT ON COLUMN prescriptions.signed_pdf_hash IS 'Hash SHA-256 do PDF assinado (integridade)';
COMMENT ON COLUMN prescriptions.qr_code_data IS 'Dados do QR Code para validação pública';
COMMENT ON COLUMN prescriptions.certificate_serial IS 'Número de série do certificado ICP-Brasil usado na assinatura';

COMMENT ON COLUMN lab_requests.signed_pdf_path IS 'Caminho do PDF de pedido de exames assinado digitalmente';
COMMENT ON COLUMN lab_requests.qr_code_data IS 'QR Code para validação pública do pedido de exames';
