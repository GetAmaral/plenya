-- Create patient_subscriptions table
-- Pacotes de acompanhamento assinados pelos pacientes

CREATE TABLE patient_subscriptions (
    id UUID PRIMARY KEY,
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    method_id UUID REFERENCES methods(id) ON DELETE SET NULL,

    plan_name VARCHAR(200) NOT NULL,
    plan_description TEXT,
    features TEXT, -- JSON

    status VARCHAR(20) NOT NULL CHECK (status IN ('active','inactive','cancelled','expired','suspended','trial')),
    billing_cycle VARCHAR(20) NOT NULL CHECK (billing_cycle IN ('monthly','quarterly','yearly','one_time')),
    auto_renew BOOLEAN NOT NULL DEFAULT true,

    start_date DATE NOT NULL,
    end_date DATE,
    trial_end_date DATE,
    next_billing_date DATE,
    cancelled_at TIMESTAMP,

    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',
    discount DECIMAL(5,2) NOT NULL DEFAULT 0 CHECK (discount >= 0 AND discount <= 100),
    discount_reason TEXT,

    cancellation_reason TEXT,
    notes TEXT,
    renewal_count INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_patient_subscriptions_patient ON patient_subscriptions(patient_id);
CREATE INDEX idx_patient_subscriptions_method ON patient_subscriptions(method_id);
CREATE INDEX idx_patient_subscriptions_status ON patient_subscriptions(status);
CREATE INDEX idx_patient_subscriptions_deleted_at ON patient_subscriptions(deleted_at);

-- Comments
COMMENT ON TABLE patient_subscriptions IS 'Pacotes de acompanhamento assinados pelos pacientes';
COMMENT ON COLUMN patient_subscriptions.features IS 'JSON com features do plano (ex: {"consultas": 4, "exames": "ilimitado"})';
COMMENT ON COLUMN patient_subscriptions.discount IS 'Desconto em percentual (0-100)';
COMMENT ON COLUMN patient_subscriptions.status IS 'Status da assinatura: active, inactive, cancelled, expired, suspended, trial';
COMMENT ON COLUMN patient_subscriptions.billing_cycle IS 'Ciclo de cobrança: monthly, quarterly, yearly, one_time';
COMMENT ON COLUMN patient_subscriptions.auto_renew IS 'Se true, renova automaticamente no fim do período';
COMMENT ON COLUMN patient_subscriptions.renewal_count IS 'Número de renovações realizadas';
