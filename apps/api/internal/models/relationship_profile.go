package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RelationshipProfile — memória de longo prazo (SOCIAL) da Lívia por pessoa. 1 por
// (owner_type, owner_id). NÃO contém dado clínico (LGPD/CFM): só relacionamento/social/
// administrativo; o clínico vive no prontuário. O rolling_summary é mantido pelo job a cada
// +5 mensagens ou no fim do atendimento e injetado no prompt do recepcionista virtual.
// Ver docs/emr/plano-livia-memoria-dossie-360.md (§0/§3.1/Fase A).
//
// @Description Perfil de relacionamento (memória social de longo prazo) de um lead/paciente
type RelationshipProfile struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	OwnerType string    `gorm:"type:varchar(20);not null;uniqueIndex:uq_relationship_profile_owner,priority:1" json:"ownerType"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_relationship_profile_owner,priority:2" json:"ownerId"`

	// Resumo rolante (SOCIAL) da relação/conversa — injetado no prompt da Lívia.
	RollingSummary string `gorm:"type:text;not null;default:''" json:"rollingSummary"`

	// Controle de regeneração do resumo.
	SummaryUpdatedAt *time.Time `json:"summaryUpdatedAt,omitempty"`
	SummaryMsgCount  int        `gorm:"not null;default:0" json:"summaryMsgCount"`

	// Estágio do relacionamento (novo, em conversa, agendou, paciente ativo, Continuum, inativo).
	// Reservado para as fases seguintes (dossiê/360).
	RelationshipStage string `gorm:"type:varchar(40);not null;default:''" json:"relationshipStage"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name
func (RelationshipProfile) TableName() string { return "relationship_profiles" }

// BeforeCreate hook generates UUID v7
func (m *RelationshipProfile) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
