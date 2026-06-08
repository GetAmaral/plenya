package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WebPushSubscription guarda uma inscrição de Web Push (protocolo VAPID) de um
// navegador/PWA da equipe. Funciona como o canal "desktop/iPhone-sem-app" do
// EMR: o backend dispara o push por aqui em paralelo ao Expo (mobile).
//
// Uma inscrição = um navegador num device. O mesmo usuário pode ter várias
// (notebook + celular). Identificada unicamente pelo endpoint que o navegador
// fornece. Quando o endpoint volta 404/410 (gone), a inscrição é apagada.
//
// @Description Inscrição de notificação Web Push de um navegador autenticado
type WebPushSubscription struct {
	// ID único da inscrição
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Usuário dono do navegador
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`

	// Endpoint do push service (FCM/Mozilla/WNS). Único por inscrição.
	Endpoint string `gorm:"type:text;not null;uniqueIndex:idx_web_push_endpoint" json:"endpoint"`

	// Chave pública P-256 do cliente (base64 url-safe), usada na cripto do payload.
	P256dh string `gorm:"type:text;not null" json:"-"`

	// Segredo de autenticação do cliente (base64 url-safe).
	Auth string `gorm:"type:text;not null" json:"-"`

	// Rótulo opcional do device (navegador/SO). Mostrado em "avisos ativos".
	// @example Chrome · macOS
	DeviceLabel *string `gorm:"type:varchar(200)" json:"deviceLabel,omitempty"`

	// User-Agent no momento do registro (diagnóstico).
	UserAgent *string `gorm:"type:varchar(400)" json:"-"`

	// Última vez que um push foi entregue com sucesso por esta inscrição.
	LastSeenAt time.Time `gorm:"not null;index" json:"lastSeenAt"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`
}

// TableName especifica o nome da tabela.
func (WebPushSubscription) TableName() string {
	return "web_push_subscriptions"
}

// BeforeCreate hook gera UUID v7.
func (w *WebPushSubscription) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.Must(uuid.NewV7())
	}
	if w.LastSeenAt.IsZero() {
		w.LastSeenAt = time.Now()
	}
	return nil
}
