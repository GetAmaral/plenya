package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Anamnesis representa uma anamnese (histórico médico) do paciente
// @Description Anamnese com queixa principal, histórico da doença atual, antecedentes, etc
type Anamnesis struct {
	// ID único da anamnese
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do médico que realizou a anamnese
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"doctorId"`

	// Queixa principal
	ChiefComplaint string `gorm:"type:text;not null" json:"chiefComplaint" validate:"required"`

	// História da doença atual (HDA)
	HistoryOfPresentIllness string `gorm:"type:text" json:"historyOfPresentIllness,omitempty"`

	// Histórico médico prévio
	PastMedicalHistory *string `gorm:"type:text" json:"pastMedicalHistory,omitempty"`

	// Medicações em uso
	CurrentMedications *string `gorm:"type:text" json:"currentMedications,omitempty"`

	// Alergias
	Allergies *string `gorm:"type:text" json:"allergies,omitempty"`

	// Histórico familiar
	FamilyHistory *string `gorm:"type:text" json:"familyHistory,omitempty"`

	// Histórico social (tabagismo, etilismo, etc)
	SocialHistory *string `gorm:"type:text" json:"socialHistory,omitempty"`

	// Revisão de sistemas
	ReviewOfSystems *string `gorm:"type:text" json:"reviewOfSystems,omitempty"`

	// Exame físico
	PhysicalExamination *string `gorm:"type:text" json:"physicalExamination,omitempty"`

	// Hipótese diagnóstica
	Assessment *string `gorm:"type:text" json:"assessment,omitempty"`

	// Plano terapêutico
	Plan *string `gorm:"type:text" json:"plan,omitempty"`

	// Data da consulta
	ConsultationDate time.Time `gorm:"type:timestamp;not null;index" json:"consultationDate"`

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
func (Anamnesis) TableName() string {
	return "anamnesis"
}
