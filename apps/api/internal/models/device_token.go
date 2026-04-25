package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DevicePlatform identifica a plataforma nativa do device.
type DevicePlatform string

const (
	DevicePlatformIOS     DevicePlatform = "ios"
	DevicePlatformAndroid DevicePlatform = "android"
)

// AppVariant identifica qual app mobile o token pertence.
type AppVariant string

const (
	AppVariantPro AppVariant = "pro"
	AppVariantApp AppVariant = "app"
)

// DeviceToken registra um device autenticado para envio de push notifications.
// Funciona também como "sessão mobile" — listada em GET /me/sessions e
// revogável via DELETE /me/sessions/:id.
//
// @Description Token de push notification de um device autenticado
type DeviceToken struct {
	// ID único do token
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Usuário dono do device
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`

	// Plataforma nativa
	// @enum ios,android
	Platform DevicePlatform `gorm:"type:varchar(10);not null;index" json:"platform" validate:"required,oneof=ios android"`

	// Variante do app
	// @enum pro,app
	AppVariant AppVariant `gorm:"type:varchar(10);not null;index" json:"appVariant" validate:"required,oneof=pro app"`

	// Expo push token (ou FCM/APNs token nativo). Único por device+app.
	Token string `gorm:"type:text;not null;uniqueIndex:idx_device_token_unique" json:"token" validate:"required,max=512"`

	// Identificador opcional do device (modelo, hostname). Mostrado em /me/sessions.
	// @example iPhone 15 Pro · iOS 17.5
	DeviceLabel *string `gorm:"type:varchar(200)" json:"deviceLabel,omitempty"`

	// Versão do app no momento do registro
	// @example 0.1.0
	AppVersion *string `gorm:"type:varchar(40)" json:"appVersion,omitempty"`

	// Última vez que o app fez request com este token
	LastSeenAt time.Time `gorm:"not null;index" json:"lastSeenAt"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Soft delete (revogação)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName especifica o nome da tabela.
func (DeviceToken) TableName() string {
	return "device_tokens"
}

// BeforeCreate hook gera UUID v7.
func (d *DeviceToken) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.Must(uuid.NewV7())
	}
	if d.LastSeenAt.IsZero() {
		d.LastSeenAt = time.Now()
	}
	return nil
}
