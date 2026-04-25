package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CalendarProvider identifica o provedor de calendar conectado.
// V1 só suporta Google. V2 deve adicionar "microsoft" (Outlook 365 / Graph API).
type CalendarProvider string

const (
	CalendarProviderGoogle CalendarProvider = "google"
)

// CalendarCredential armazena as credenciais OAuth de um médico para acessar
// um calendar externo (Google Calendar V1, Outlook V2).
//
// SEGURANÇA / LGPD:
//   - EncryptedRefreshToken e EncryptedAccessToken são CIFRADOS via
//     crypto.EncryptWithDefaultKey antes de persistir. NUNCA armazenar
//     tokens em plaintext.
//   - Ambos campos têm `json:"-"` pra garantir que NUNCA vazem em response
//     da API — mesmo que alguém acidentalmente serialize o struct.
//   - DedicatedCalendarID é o ID do calendar "Plenya" criado na conexão
//     inicial. Toda escrita (createEvent/patchEvent/deleteEvent) usa esse
//     ID — nunca o calendar primário do médico, garantindo isolamento.
//
// Lifecycle:
//   - Criado no callback OAuth (`/api/v1/integrations/google/callback`)
//   - Atualizado pelo cron de refresh (a cada 30min, refresh quando access
//     token expira em <10min)
//   - Marcado RevokedAt quando refresh retorna `invalid_grant` (médico
//     revogou acesso em myaccount.google.com/permissions)
//
// @Description Credenciais OAuth (cifradas) de um médico para Google Calendar
type CalendarCredential struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Médico dono da credencial (FK pra users.id).
	// Unique composto com Provider — 1 conta por médico por provedor (V1).
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_credential_user_provider,priority:1" json:"userId"`

	// Provedor: "google" (V1). Futuramente "microsoft".
	Provider string `gorm:"type:varchar(20);not null;uniqueIndex:idx_credential_user_provider,priority:2" json:"provider"`

	// Email da conta Google conectada (mostrado na UI: "Google conectado · email@gmail.com").
	GoogleAccountEmail string `gorm:"type:varchar(255);not null" json:"googleAccountEmail"`

	// Refresh token OAuth (CIFRADO via crypto.EncryptWithDefaultKey).
	// NUNCA expor em JSON — usar service layer pra decrypt apenas dentro de
	// chamadas Google API.
	EncryptedRefreshToken string `gorm:"type:text;not null" json:"-"`

	// Access token OAuth (CIFRADO via crypto.EncryptWithDefaultKey).
	// Curto (1h). Renovado automaticamente pelo cron de token refresh.
	EncryptedAccessToken string `gorm:"type:text;not null" json:"-"`

	// Quando o access token atual expira. Cron renova antes desse limite.
	AccessTokenExpiresAt time.Time `gorm:"not null" json:"-"`

	// ID do calendar dedicado "Plenya" criado pra esse médico via
	// calendarList.insert. Toda operação de evento usa esse ID.
	DedicatedCalendarID string `gorm:"type:varchar(255);not null" json:"dedicatedCalendarId"`

	// Última sincronização bem-sucedida (qualquer operação Google API).
	// Usado pra detectar credenciais "esquecidas" no health dashboard.
	LastSyncAt *time.Time `json:"lastSyncAt,omitempty"`

	// Quando a credencial foi revogada (médico desconectou ou Google
	// retornou invalid_grant). Quando setado, services devem ignorar essa
	// credencial e mostrar badge "Reconectar Google" pro médico.
	RevokedAt *time.Time `json:"revokedAt,omitempty"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Relação — Doctor é o User que conectou.
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName especifica o nome da tabela
func (CalendarCredential) TableName() string {
	return "calendar_credentials"
}

// BeforeCreate hook gera UUID v7
func (c *CalendarCredential) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// IsActive retorna true se a credencial está utilizável (não revogada).
func (c *CalendarCredential) IsActive() bool {
	return c.RevokedAt == nil
}
