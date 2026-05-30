-- Migration: Refactor subscription system with SubscriptionPlan catalog
-- Date: 2026-02-16
-- Description: Creates subscription_plans table (catalog) and refactors patient_subscriptions

-- =========================================
-- 1. DROP old patient_subscriptions table
-- =========================================
DROP TABLE IF EXISTS patient_subscriptions CASCADE;
DROP TABLE IF EXISTS notifications CASCADE; -- Will recreate with updated FK

-- =========================================
-- 2. CREATE subscription_plans table
-- =========================================
CREATE TABLE subscription_plans (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,

    -- Features and pricing
    features JSONB,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',
    billing_cycle VARCHAR(20) NOT NULL CHECK (billing_cycle IN ('monthly', 'quarterly', 'yearly', 'one_time')),

    -- Method association
    method_id UUID REFERENCES methods(id) ON DELETE SET NULL ON UPDATE CASCADE,

    -- Configuration
    is_active BOOLEAN NOT NULL DEFAULT true,
    trial_period_days INTEGER NOT NULL DEFAULT 0 CHECK (trial_period_days >= 0),
    "order" INTEGER NOT NULL DEFAULT 0,

    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes for subscription_plans
CREATE INDEX idx_plan_name ON subscription_plans(name);
CREATE INDEX idx_plan_method ON subscription_plans(method_id);
CREATE INDEX idx_plan_is_active ON subscription_plans(is_active);
CREATE INDEX idx_plan_order ON subscription_plans("order");
CREATE INDEX idx_plan_deleted_at ON subscription_plans(deleted_at);

-- =========================================
-- 3. CREATE patient_subscriptions table (refactored)
-- =========================================
CREATE TABLE patient_subscriptions (
    id UUID PRIMARY KEY,

    -- Foreign keys
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    subscription_plan_id UUID NOT NULL REFERENCES subscription_plans(id) ON DELETE RESTRICT,

    -- Plan snapshot (preserves plan data at subscription time)
    plan_snapshot JSONB NOT NULL,

    -- Status
    status VARCHAR(20) NOT NULL CHECK (status IN ('active', 'inactive', 'cancelled', 'expired', 'suspended', 'trial')),
    auto_renew BOOLEAN NOT NULL DEFAULT true,

    -- Dates
    start_date DATE NOT NULL,
    end_date DATE,
    trial_end_date DATE,
    next_billing_date DATE,
    cancelled_at TIMESTAMP,

    -- Customization (overrides plan defaults)
    discount_percent DECIMAL(5,2) NOT NULL DEFAULT 0 CHECK (discount_percent >= 0 AND discount_percent <= 100),
    discount_reason TEXT,
    custom_price DECIMAL(10,2) CHECK (custom_price >= 0),
    custom_trial_days INTEGER CHECK (custom_trial_days >= 0),

    -- Management
    cancellation_reason TEXT,
    notes TEXT,
    renewal_count INTEGER NOT NULL DEFAULT 0,

    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes for patient_subscriptions
CREATE INDEX idx_patient_subscriptions_patient ON patient_subscriptions(patient_id);
CREATE INDEX idx_patient_subscriptions_plan ON patient_subscriptions(subscription_plan_id);
CREATE INDEX idx_patient_subscriptions_status ON patient_subscriptions(status);
CREATE INDEX idx_patient_subscriptions_deleted_at ON patient_subscriptions(deleted_at);

-- =========================================
-- 4. RECREATE notifications table with updated FK
-- =========================================
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id UUID REFERENCES patients(id) ON DELETE CASCADE,
    subscription_id UUID REFERENCES patient_subscriptions(id) ON DELETE CASCADE,

    type VARCHAR(50) NOT NULL CHECK (type IN ('trial_expiring', 'renewal_upcoming', 'subscription_expired', 'payment_pending', 'general')),
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    action_url TEXT,
    action_text VARCHAR(100),

    read BOOLEAN NOT NULL DEFAULT false,
    read_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for notifications
CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_patient ON notifications(patient_id);
CREATE INDEX idx_notifications_subscription ON notifications(subscription_id);
CREATE INDEX idx_notifications_type ON notifications(type);
CREATE INDEX idx_notifications_read ON notifications(read);
CREATE INDEX idx_notifications_created_at ON notifications(created_at DESC);

-- =========================================
-- 5. SEED subscription plans
-- =========================================

-- Get default method ID (will use in plans)
DO $$
DECLARE
    default_method_id UUID;
BEGIN
    -- Get default method
    SELECT id INTO default_method_id FROM methods WHERE is_default = true AND deleted_at IS NULL LIMIT 1;

    -- If no default method, get first active method
    IF default_method_id IS NULL THEN
        SELECT id INTO default_method_id FROM methods WHERE deleted_at IS NULL LIMIT 1;
    END IF;

    -- Insert seed plans
    INSERT INTO subscription_plans (id, name, description, features, price, currency, billing_cycle, method_id, is_active, trial_period_days, "order")
    VALUES
        -- Plano Básico
        (
            gen_random_uuid(),
            'Plano Básico',
            'Acompanhamento essencial com consultas e exames básicos',
            '{"consultas": 4, "exames": "básicos", "whatsapp": true, "urgencia": false}'::jsonb,
            149.90,
            'BRL',
            'monthly',
            default_method_id,
            true,
            7,
            1
        ),
        -- Plano Premium
        (
            gen_random_uuid(),
            'Plano Premium',
            'Acompanhamento completo com consultas ilimitadas e exames avançados',
            '{"consultas": "ilimitado", "exames": "completos", "whatsapp": true, "urgencia": true, "nutricao": true}'::jsonb,
            299.90,
            'BRL',
            'monthly',
            default_method_id,
            true,
            14,
            2
        ),
        -- Plano Anual VIP
        (
            gen_random_uuid(),
            'Plano Anual VIP',
            'Acompanhamento premium anual com desconto e benefícios exclusivos',
            '{"consultas": "ilimitado", "exames": "completos", "whatsapp": true, "urgencia": true, "nutricao": true, "psicologia": true, "checkup_anual": true}'::jsonb,
            2999.00,
            'BRL',
            'yearly',
            default_method_id,
            true,
            30,
            3
        );
END $$;

-- =========================================
-- Migration complete
-- =========================================
