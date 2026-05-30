package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConsultationPrice — preço sugerido por tipo de consulta. Admin gerencia; a
// recepção lê para pré-preencher o valor ao registrar um pagamento.
//
// Valores em centavos (int64) para evitar imprecisão de ponto flutuante.
type ConsultationPrice struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Tipo de consulta (único — um preço corrente por tipo).
	Type AppointmentType `gorm:"type:varchar(20);not null;uniqueIndex" json:"type"`

	// Valor em centavos de BRL.
	AmountCents int64 `gorm:"type:bigint;not null;default:0" json:"amountCents"`

	// Ativo — preços inativos não aparecem como sugestão.
	Active bool `gorm:"not null;default:true" json:"active"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ConsultationPrice) TableName() string { return "consultation_prices" }

func (c *ConsultationPrice) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// PaymentReceiptCounter — contador de recibos por ano, para numeração sequencial
// contínua (ex.: 2026-000123). Incrementado via UPSERT atômico no service.
type PaymentReceiptCounter struct {
	Year    int `gorm:"primaryKey" json:"year"`
	LastSeq int `gorm:"not null;default:0" json:"lastSeq"`
}

func (PaymentReceiptCounter) TableName() string { return "payment_receipt_counters" }
