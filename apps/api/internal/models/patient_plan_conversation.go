package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// PatientPlanMessage — um turno da conversa que edita o rascunho.
//
// `Body` guarda prosa clínica identificável, da mesma classe de `clinical_notes`: entra em
// qualquer política de retenção junto com o resto do prontuário. Os campos `AI*` guardam só
// metadados — modelo, tokens, latência — porque prompt e resposta não vão para log.
type PatientPlanMessage struct {
	ID     uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PlanID uuid.UUID  `gorm:"type:uuid;not null;index" json:"planId"`
	Seq    int        `gorm:"not null" json:"seq"`
	Role   string     `gorm:"type:varchar(12);not null" json:"role"`
	Body   string     `gorm:"type:text;not null" json:"body"`
	UserID *uuid.UUID `gorm:"type:uuid" json:"userId,omitempty"`

	// ClientMessageID — idempotência de um POST que leva de dez a vinte segundos.
	ClientMessageID string `gorm:"type:varchar(64)" json:"clientMessageId,omitempty"`

	Status       string `gorm:"type:varchar(12);not null;default:'ok'" json:"status"`
	ErrorMessage string `gorm:"type:text" json:"errorMessage,omitempty"`

	AIModel           string `gorm:"type:varchar(60);column:ai_model" json:"aiModel,omitempty"`
	AIPromptVersion   string `gorm:"type:varchar(20);column:ai_prompt_version" json:"aiPromptVersion,omitempty"`
	AIInputTokens     int    `gorm:"column:ai_input_tokens" json:"aiInputTokens,omitempty"`
	AICacheReadTokens int    `gorm:"column:ai_cache_read_tokens" json:"aiCacheReadTokens,omitempty"`
	AIOutputTokens    int    `gorm:"column:ai_output_tokens" json:"aiOutputTokens,omitempty"`
	AIStopReason      string `gorm:"type:varchar(30);column:ai_stop_reason" json:"aiStopReason,omitempty"`
	LatencyMs         int    `gorm:"column:latency_ms" json:"latencyMs,omitempty"`

	DossierID *uuid.UUID `gorm:"type:uuid" json:"dossierId,omitempty"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt"`
}

func (PatientPlanMessage) TableName() string { return "patient_plan_messages" }

func (m *PatientPlanMessage) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// PatientPlanSuggestionStatus — o ciclo de vida da sugestão.
type PatientPlanSuggestionStatus string

const (
	SuggestionPending  PatientPlanSuggestionStatus = "pending"
	SuggestionAccepted PatientPlanSuggestionStatus = "accepted"
	SuggestionRejected PatientPlanSuggestionStatus = "rejected"
	// SuggestionStale — o slide alvo mudou depois que a sugestão nasceu. Aceitar agora
	// sobrescreveria em silêncio o que o médico escreveu à mão.
	SuggestionStale PatientPlanSuggestionStatus = "stale"
	// SuggestionSuperseded — um turno posterior propôs algo no mesmo campo do mesmo slide.
	SuggestionSuperseded PatientPlanSuggestionStatus = "superseded"
)

// PatientPlanSuggestion — uma alteração proposta que espera o aceite do médico.
type PatientPlanSuggestion struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PlanID    uuid.UUID `gorm:"type:uuid;not null;index" json:"planId"`
	MessageID uuid.UUID `gorm:"type:uuid;not null" json:"messageId"`

	Op           string `gorm:"type:varchar(12);not null" json:"op"`
	SlideID      string `gorm:"type:varchar(64)" json:"slideId,omitempty"`
	AfterSlideID string `gorm:"type:varchar(64)" json:"afterSlideId,omitempty"`
	FieldPath    string `gorm:"type:text" json:"fieldPath,omitempty"`

	OldValue datatypes.JSON `gorm:"type:jsonb" json:"oldValue,omitempty"`
	NewValue datatypes.JSON `gorm:"type:jsonb" json:"newValue,omitempty"`

	// BaseHash — o estado do slide alvo quando a sugestão nasceu.
	BaseHash string `gorm:"type:char(64)" json:"baseHash,omitempty"`

	Class      string         `gorm:"type:varchar(12);not null" json:"class"`
	Rationale  string         `gorm:"type:text;not null;default:''" json:"rationale"`
	Provenance datatypes.JSON `gorm:"type:jsonb;not null;default:'[]'" json:"provenance"`

	Status       PatientPlanSuggestionStatus `gorm:"type:varchar(12);not null;default:'pending'" json:"status"`
	ResolvedByID *uuid.UUID                  `gorm:"type:uuid" json:"resolvedById,omitempty"`
	ResolvedAt   *time.Time                  `gorm:"type:timestamptz" json:"resolvedAt,omitempty"`
	RevisionID   *uuid.UUID                  `gorm:"type:uuid" json:"revisionId,omitempty"`
	CreatedAt    time.Time                   `gorm:"not null" json:"createdAt"`
}

func (PatientPlanSuggestion) TableName() string { return "patient_plan_suggestions" }

func (s *PatientPlanSuggestion) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
