package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnamnesisVisibility representa os níveis de visibilidade da anamnese
type AnamnesisVisibility string

const (
	VisibilityAll         AnamnesisVisibility = "all"
	VisibilityMedicalOnly AnamnesisVisibility = "medicalOnly"
	VisibilityPsychOnly   AnamnesisVisibility = "psychOnly"
)

// Anamnesis representa uma anamnese (histórico médico) do paciente
// @Description Anamnese simplificada com itens registrados via AnamnesisItem
type Anamnesis struct {
	// ID único da anamnese
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do paciente
	// @example 550e8400-e29b-41d4-a716-446655440000
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_patient" json:"patientId" validate:"required"`

	// ID do profissional que realizou a anamnese (médico, psicólogo, etc.)
	// @example 550e8400-e29b-41d4-a716-446655440000
	AuthorID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_author" json:"authorId" validate:"required"`

	// ID do template de anamnese utilizado (opcional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	AnamnesisTemplateID *uuid.UUID `gorm:"type:uuid;index:idx_anamnesis_template" json:"anamnesisTemplateId,omitempty"`

	// Data da consulta
	// @example 2026-01-25T14:30:00Z
	ConsultationDate time.Time `gorm:"type:timestamp;not null;index:idx_anamnesis_consultation_date" json:"consultationDate" validate:"required"`

	// Conteúdo completo da anamnese
	// @example Paciente do sexo masculino, 45 anos, comparece à consulta com queixa de...
	Content *string `gorm:"type:text" json:"content,omitempty"`

	// Resumo da anamnese
	// @example Paciente com hipertensão controlada, busca avaliação para dor torácica
	Summary *string `gorm:"type:text" json:"summary,omitempty"`

	// Visibilidade da anamnese
	// @enum all,medicalOnly,psychOnly
	// @example all
	Visibility AnamnesisVisibility `gorm:"type:varchar(20);not null;default:'all';check:visibility IN ('all','medicalOnly','psychOnly')" json:"visibility" validate:"required,oneof=all medicalOnly psychOnly"`

	// Observações gerais (opcional)
	// @example Paciente colaborativo, anamnese detalhada
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Patient           Patient             `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Author            User                `gorm:"foreignKey:AuthorID;constraint:OnDelete:RESTRICT" json:"author,omitempty"`
	AnamnesisTemplate *AnamnesisTemplate  `gorm:"foreignKey:AnamnesisTemplateID;constraint:OnDelete:SET NULL" json:"anamnesisTemplate,omitempty"`
	Items             []AnamnesisItem     `gorm:"foreignKey:AnamnesisID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

// TableName especifica o nome da tabela
func (Anamnesis) TableName() string {
	return "anamnesis"
}

// BeforeCreate hook to generate UUID v7
func (a *Anamnesis) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
