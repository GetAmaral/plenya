package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/pdfdoc"
)

// PatientPlanAuthorKind — quem escreveu a revisão.
//
// `assistant` não substitui o autor humano: `CreatedByID` continua sendo o clínico logado, porque
// alguém responde pelo que o paciente lê. O que este campo registra é se a mão que escreveu foi a
// dele ou a da ferramenta, que é a pergunta que se faz meses depois olhando uma frase.
type PatientPlanAuthorKind string

const (
	PlanAuthorHuman     PatientPlanAuthorKind = "human"
	PlanAuthorAssistant PatientPlanAuthorKind = "assistant"
	PlanAuthorSystem    PatientPlanAuthorKind = "system"
)

// PatientPlanRevisionReason — o que provocou a revisão. Só `edit` coalesce (ver
// PatientPlanRevisionService.Record); os outros são eventos auditáveis e cada um vira uma linha.
type PatientPlanRevisionReason string

const (
	PlanRevisionEdit             PatientPlanRevisionReason = "edit"
	PlanRevisionAIApply          PatientPlanRevisionReason = "ai_apply"
	PlanRevisionSuggestionAccept PatientPlanRevisionReason = "suggestion_accept"
	PlanRevisionRestore          PatientPlanRevisionReason = "restore"
	PlanRevisionPublish          PatientPlanRevisionReason = "publish"
)

// PatientPlanRevision — o estado do rascunho depois de uma gravação.
//
// Guarda o conteúdo RESULTANTE, não o delta: restaurar é copiar uma linha, e uma cadeia de patches
// corrompida no meio mataria todo o histórico a partir dali. `Ops` é só para explicar o que mudou
// na tela sem diffar dois JSONB grandes.
//
// @Description Revisão do rascunho do plano de devolutiva
type PatientPlanRevision struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PlanID uuid.UUID `gorm:"type:uuid;not null;index" json:"planId"`

	// Seq conta EDIÇÕES. PatientPlan.Version conta PUBLICAÇÕES. Nunca somar os dois.
	Seq int `gorm:"not null" json:"seq"`
	// PlanVersion é a publicação a que esta edição pertence, para responder "o que aconteceu
	// entre a v1 e a v2".
	PlanVersion int `gorm:"not null" json:"planVersion"`

	Title   string             `gorm:"type:varchar(300);not null" json:"title"`
	Content []pdfdoc.DeckSlide `gorm:"type:jsonb;serializer:json;not null" json:"content"`
	// ContentHash — sha256 do conteúdo canônico. Gravação que não mudou nada não vira linha.
	ContentHash string `gorm:"type:char(64);not null" json:"contentHash"`

	AuthorKind  PatientPlanAuthorKind     `gorm:"type:varchar(12);not null" json:"authorKind"`
	CreatedByID uuid.UUID                 `gorm:"type:uuid;not null" json:"createdById"`
	Reason      PatientPlanRevisionReason `gorm:"type:varchar(24);not null" json:"reason"`

	Ops             any        `gorm:"type:jsonb;serializer:json" json:"ops,omitempty"`
	MessageID       *uuid.UUID `gorm:"type:uuid" json:"messageId,omitempty"`
	DossierID       *uuid.UUID `gorm:"type:uuid" json:"dossierId,omitempty"`
	AIModel         string     `gorm:"type:varchar(60);column:ai_model" json:"aiModel,omitempty"`
	AIPromptVersion string     `gorm:"type:varchar(20);column:ai_prompt_version" json:"aiPromptVersion,omitempty"`

	IsPublication bool      `gorm:"not null;default:false" json:"isPublication"`
	CreatedAt     time.Time `gorm:"not null" json:"createdAt"`

	CreatedBy User `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
}

func (PatientPlanRevision) TableName() string { return "patient_plan_revisions" }

func (r *PatientPlanRevision) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
