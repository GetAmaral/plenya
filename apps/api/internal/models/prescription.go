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

// Prescription representa uma prescrição médica
// @Description Prescrição médica com medicamentos, dosagem e instruções
type Prescription struct {
	// ID único da prescrição
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do médico prescritor
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"doctorId"`

	// Nome do medicamento
	MedicationName string `gorm:"type:varchar(200);not null" json:"medicationName" validate:"required"`

	// Princípio ativo
	ActiveIngredient *string `gorm:"type:varchar(200)" json:"activeIngredient,omitempty"`

	// Dosagem (ex: 500mg, 10ml, etc)
	Dosage string `gorm:"type:varchar(100);not null" json:"dosage" validate:"required"`

	// Frequência (ex: 8/8h, 12/12h, etc)
	Frequency string `gorm:"type:varchar(100);not null" json:"frequency" validate:"required"`

	// Via de administração (ex: oral, intravenosa, tópica, etc)
	Route string `gorm:"type:varchar(50);not null" json:"route" validate:"required"`

	// Duração do tratamento (ex: 7 dias, 30 dias, contínuo)
	Duration *string `gorm:"type:varchar(100)" json:"duration,omitempty"`

	// Quantidade prescrita
	Quantity *string `gorm:"type:varchar(100)" json:"quantity,omitempty"`

	// Instruções especiais
	Instructions *string `gorm:"type:text" json:"instructions,omitempty"`

	// Status da prescrição
	Status PrescriptionStatus `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','completed','cancelled','expired')" json:"status"`

	// Data da prescrição
	PrescriptionDate time.Time `gorm:"type:timestamp;not null;index" json:"prescriptionDate"`

	// Data de início do tratamento
	StartDate *time.Time `gorm:"type:date" json:"startDate,omitempty"`

	// Data de término do tratamento
	EndDate *time.Time `gorm:"type:date" json:"endDate,omitempty"`

	// Observações adicionais
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor  User    `gorm:"foreignKey:DoctorID;constraint:OnDelete:RESTRICT" json:"doctor,omitempty"`
}

// TableName especifica o nome da tabela
func (Prescription) TableName() string {
	return "prescriptions"
}
