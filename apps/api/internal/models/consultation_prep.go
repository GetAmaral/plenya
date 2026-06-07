package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status de uma ConsultationPrep.
const (
	ConsultationPrepDraft     = "draft"
	ConsultationPrepSubmitted = "submitted"
)

// ConsultationPrep é a preparação pré-consulta preenchida pelo paciente DEPOIS de agendar e pagar.
// Espelha AnonymousScoreSession (Escore Light), porém vinculada a um Patient (e opcionalmente a um
// Appointment). Reúne as respostas do formulário curado (subset de ScoreItems com PrepOrder != nil),
// uma queixa/objetivo livre, e serve de base para o Dr. revisar antes da consulta e semear a Anamnese.
//
// @Description Preparação pré-consulta do paciente (formulário curado pós-agendamento)
type ConsultationPrep struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Paciente dono da preparação
	// @example 550e8400-e29b-41d4-a716-446655440000
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_consultation_prep_patient" json:"patientId" validate:"required"`

	// Consulta à qual esta preparação se refere (opcional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	AppointmentID *uuid.UUID `gorm:"type:uuid;index:idx_consultation_prep_appointment" json:"appointmentId,omitempty"`

	// Formulário (ScoreVersion context=patient_prep) que originou esta preparação. Define
	// quais itens o paciente respondeu (A1/B1/B2). Copiado da consulta no momento do submit.
	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreVersionID *uuid.UUID `gorm:"type:uuid;index:idx_consultation_prep_version" json:"scoreVersionId,omitempty"`

	// Status: draft (em preenchimento) | submitted (enviado ao Dr.)
	// @enum draft,submitted
	// @example draft
	Status string `gorm:"type:varchar(20);not null;default:'draft';check:status IN ('draft','submitted')" json:"status"`

	// Queixa / objetivo principal em texto livre (não é um ScoreItem)
	// @example Cansaço persistente e dificuldade de perder peso mesmo treinando
	ChiefComplaint *string `gorm:"type:text" json:"chiefComplaint,omitempty"`

	// Timestamp do envio (preenchido quando Status vira submitted)
	SubmittedAt *time.Time `gorm:"type:timestamptz" json:"submittedAt,omitempty"`

	// Consentimento LGPD
	ConsentVersion   *string    `gorm:"type:varchar(20)" json:"consentVersion,omitempty"`
	ConsentTimestamp *time.Time `gorm:"type:timestamptz" json:"consentTimestamp,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Patient   *Patient                   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	Responses []ConsultationPrepResponse `gorm:"foreignKey:PrepID;constraint:OnDelete:CASCADE" json:"responses,omitempty"`
}

// TableName specifies the table name
func (ConsultationPrep) TableName() string {
	return "consultation_preps"
}

// BeforeCreate hook generates UUID v7
func (p *ConsultationPrep) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// ConsultationPrepResponse é a resposta do paciente a um único item do formulário de preparação.
// Espelha AnonymousScoreItem / AnamnesisItem (mesmos 3 campos de valor).
//
// @Description Resposta a um item do formulário de preparação pré-consulta
type ConsultationPrepResponse struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Preparação a que esta resposta pertence
	// @example 550e8400-e29b-41d4-a716-446655440000
	PrepID uuid.UUID `gorm:"type:uuid;not null;index:idx_consultation_prep_response_prep" json:"prepId" validate:"required"`

	// ID do ScoreItem ao qual esta resposta se refere
	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_consultation_prep_response_item" json:"scoreItemId" validate:"required"`

	// Valor numérico (medida objetiva ou nível 0-6 de anamnese)
	NumericValue *float64 `gorm:"type:double precision" json:"numericValue,omitempty"`

	// Valor textual (itens de free-text)
	TextValue *string `gorm:"type:text" json:"textValue,omitempty"`

	// Nível selecionado diretamente (0-6) — múltipla escolha
	// @minimum 0
	// @maximum 6
	SelectedLevel *int `gorm:"type:integer;check:selected_level >= 0 AND selected_level <= 6" json:"selectedLevel,omitempty" validate:"omitempty,gte=0,lte=6"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Prep      *ConsultationPrep `gorm:"foreignKey:PrepID;constraint:OnDelete:CASCADE" json:"-"`
	ScoreItem *ScoreItem        `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:RESTRICT" json:"scoreItem,omitempty"`
}

// TableName specifies the table name
func (ConsultationPrepResponse) TableName() string {
	return "consultation_prep_responses"
}

// BeforeCreate hook generates UUID v7
func (r *ConsultationPrepResponse) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
