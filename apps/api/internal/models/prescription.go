package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrescriptionStatus define os status possíveis de uma prescrição
type PrescriptionStatus string

const (
	PrescriptionActive    PrescriptionStatus = "active"    // Ativa
	PrescriptionCompleted PrescriptionStatus = "completed" // Concluída
	PrescriptionCancelled PrescriptionStatus = "cancelled" // Cancelada
	PrescriptionExpired   PrescriptionStatus = "expired"   // Expirada
)

// Prescription representa uma prescrição médica digital
// @Description Prescrição médica digital com assinatura ICP-Brasil e SNCR
type Prescription struct {
	// ID único da prescrição
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Dados Básicos

	// ID do paciente
	// @required
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId" validate:"required"`

	// ID do médico prescritor
	// @required
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"doctorId" validate:"required"`

	// Instruções gerais da prescrição (aplicam-se a todos os medicamentos)
	// @example Tomar durante as refeições
	GeneralInstructions *string `gorm:"type:text" json:"generalInstructions,omitempty"`

	// Datas

	// Data da prescrição
	// @required
	PrescriptionDate time.Time `gorm:"type:timestamp;not null;index" json:"prescriptionDate" validate:"required"`

	// Data de validade (calculada automaticamente baseada na categoria)
	// @required
	ValidUntil time.Time `gorm:"type:date;not null;index" json:"validUntil" validate:"required"`

	// SNCR (Sistema Nacional de Controle de Receitas)

	// Número SNCR (gerado pela ANVISA ou stub)
	// @example BR-STUB-2026-00000001
	SNCRNumber *string `gorm:"type:varchar(50);unique;index" json:"sncrNumber,omitempty"`

	// Status no SNCR
	// @enum active,used,cancelled
	SNCRStatus *string `gorm:"type:varchar(20)" json:"sncrStatus,omitempty"`

	// Assinatura Digital ICP-Brasil

	// Caminho do PDF assinado
	// @example /app/uploads/prescriptions/prescription_550e8400_signed.pdf
	SignedPDFPath *string `gorm:"type:varchar(500)" json:"signedPdfPath,omitempty"`

	// Hash SHA-256 do PDF assinado (integridade)
	// @example a1b2c3d4e5f6...
	SignedPDFHash *string `gorm:"type:varchar(64)" json:"signedPdfHash,omitempty"`

	// Dados do QR Code (URL de validação)
	// @example https://plenya.com.br/prescriptions/validate/550e8400...
	QRCodeData *string `gorm:"type:text" json:"qrCodeData,omitempty"`

	// Data/hora da assinatura digital
	SignedAt *time.Time `gorm:"type:timestamp" json:"signedAt,omitempty"`

	// Número de série do certificado usado na assinatura
	// @example 1234567890ABCDEF
	CertificateSerial *string `gorm:"type:varchar(100)" json:"certificateSerial,omitempty"`

	// Status e Controle

	// Status da prescrição
	// @enum active,completed,cancelled,expired
	Status PrescriptionStatus `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','completed','cancelled','expired')" json:"status" validate:"required,oneof=active completed cancelled expired"`

	// Se a prescrição foi dispensada
	IsUsed bool `gorm:"type:boolean;not null;default:false;index" json:"isUsed"`

	// Data/hora da dispensação
	DispensedAt *time.Time `gorm:"type:timestamp" json:"dispensedAt,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient     Patient                  `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor      User                     `gorm:"foreignKey:DoctorID;constraint:OnDelete:RESTRICT" json:"doctor,omitempty"`
	Medications []PrescriptionMedication `gorm:"foreignKey:PrescriptionID;constraint:OnDelete:CASCADE" json:"medications,omitempty"`
}

// TableName especifica o nome da tabela
func (Prescription) TableName() string {
	return "prescriptions"
}

// BeforeCreate hook to generate UUID v7
func (p *Prescription) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
