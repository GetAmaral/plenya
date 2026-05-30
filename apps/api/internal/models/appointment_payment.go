package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PaymentMethod — forma de pagamento.
type PaymentMethod string

const (
	PaymentMethodCash     PaymentMethod = "cash"
	PaymentMethodPix      PaymentMethod = "pix"
	PaymentMethodCard     PaymentMethod = "card"
	PaymentMethodTransfer PaymentMethod = "transfer"
	PaymentMethodOther    PaymentMethod = "other"
)

// PaymentStatus — estado do pagamento.
type PaymentStatus string

const (
	PaymentStatusPaid     PaymentStatus = "paid"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

// AppointmentPayment — registro de pagamento de consulta (avulsa ou vinculada a
// um agendamento). Recibo numerado sequencialmente por ano.
//
// Valores em centavos (int64). AppointmentID é opcional: consulta avulsa pode
// não ter agendamento no sistema.
type AppointmentPayment struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	PatientID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"patientId"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"appointmentId,omitempty"`

	// Valor em centavos de BRL.
	AmountCents int64 `gorm:"type:bigint;not null" json:"amountCents"`

	Method PaymentMethod `gorm:"type:varchar(20);not null;check:method IN ('cash','pix','card','transfer','other')" json:"method"`
	Status PaymentStatus `gorm:"type:varchar(20);not null;default:'paid';check:status IN ('paid','refunded')" json:"status"`

	// Quando o pagamento foi efetuado.
	PaidAt time.Time `gorm:"type:timestamptz;not null" json:"paidAt"`

	// Número do recibo (sequencial por ano, ex.: 2026-000123). Único.
	ReceiptNumber string `gorm:"type:varchar(20);not null;uniqueIndex" json:"receiptNumber"`

	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Estorno
	RefundedAt   *time.Time `gorm:"type:timestamptz" json:"refundedAt,omitempty"`
	RefundReason *string    `gorm:"type:text" json:"refundReason,omitempty"`

	// Quem registrou (staff).
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null;index" json:"createdByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient     Patient      `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Appointment *Appointment `gorm:"foreignKey:AppointmentID;constraint:OnDelete:SET NULL" json:"appointment,omitempty"`
}

func (AppointmentPayment) TableName() string { return "appointment_payments" }

func (p *AppointmentPayment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
