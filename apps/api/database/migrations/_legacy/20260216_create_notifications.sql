-- Create notifications table
-- Notificações do sistema para usuários

CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    patient_id UUID REFERENCES patients(id) ON DELETE CASCADE,
    subscription_id UUID REFERENCES patient_subscriptions(id) ON DELETE CASCADE,

    type VARCHAR(50) NOT NULL CHECK (type IN ('trial_expiring','renewal_upcoming','subscription_expired','payment_pending','general')),
    title VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    action_url VARCHAR(500),
    action_text VARCHAR(50),

    read BOOLEAN NOT NULL DEFAULT false,
    read_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_patient ON notifications(patient_id);
CREATE INDEX idx_notifications_subscription ON notifications(subscription_id);
CREATE INDEX idx_notifications_type ON notifications(type);
CREATE INDEX idx_notifications_read ON notifications(read);
CREATE INDEX idx_notifications_deleted_at ON notifications(deleted_at);

-- Composite index for unread notifications per user
CREATE INDEX idx_notifications_user_read ON notifications(user_id, read) WHERE deleted_at IS NULL;

-- Comments
COMMENT ON TABLE notifications IS 'Notificações do sistema para usuários';
COMMENT ON COLUMN notifications.type IS 'Tipo da notificação: trial_expiring, renewal_upcoming, subscription_expired, payment_pending, general';
COMMENT ON COLUMN notifications.action_url IS 'URL para navegação ao clicar na notificação';
COMMENT ON COLUMN notifications.action_text IS 'Texto do botão de ação (ex: Ver Subscription)';
COMMENT ON COLUMN notifications.read IS 'Se a notificação foi lida pelo usuário';
