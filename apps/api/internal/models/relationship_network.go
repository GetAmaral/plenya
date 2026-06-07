package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Tipos de evento de relacionamento (§2.4 do plano).
const (
	RelationshipEventBirthday   = "birthday"
	RelationshipEventGraduation = "graduation"
	RelationshipEventBirth      = "birth"
	RelationshipEventWedding    = "wedding"
	RelationshipEventLoss       = "loss"
	RelationshipEventMilestone  = "milestone"
	RelationshipEventFollowup   = "followup"
	RelationshipEventCustom     = "custom"
)

// Status do ciclo de um evento/aviso.
const (
	RelationshipEventPending      = "pending"
	RelationshipEventAcknowledged = "acknowledged"
	RelationshipEventDone         = "done"
	RelationshipEventDismissed    = "dismissed"
)

// ImportantPerson — pessoa da rede de relacionamento de um lead/paciente (cônjuge, filhos, quem
// indicou…). Só social (LGPD/CFM). Ver docs/emr/plano-livia-memoria-dossie-360.md §2.3.
//
// @Description Pessoa importante na rede de relacionamento de um lead/paciente
type ImportantPerson struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	OwnerType string    `gorm:"type:varchar(20);not null;index:idx_important_people_owner,priority:1" json:"ownerType"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;index:idx_important_people_owner,priority:2" json:"ownerId"`

	Name     string     `gorm:"type:varchar(160);not null" json:"name"`
	Relation string     `gorm:"type:varchar(60);not null;default:''" json:"relation"`
	Birthday *time.Time `gorm:"type:date" json:"birthday,omitempty"`
	Notes    string     `gorm:"type:text;not null;default:''" json:"notes"`

	Source  string     `gorm:"type:varchar(20);not null;default:'ai'" json:"source"`
	AddedBy *uuid.UUID `gorm:"type:uuid" json:"addedBy,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (ImportantPerson) TableName() string { return "important_people" }

func (m *ImportantPerson) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// RelationshipEvent — data/evento de relacionamento para o time interagir proativamente
// (aniversário, formatura, nascimento, follow-up…). lead_time_days = antecedência do aviso.
// Só social (LGPD/CFM). Ver docs/emr/plano-livia-memoria-dossie-360.md §2.4/§5.
//
// @Description Evento/data de relacionamento (alimenta os avisos do time)
type RelationshipEvent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	OwnerType string    `gorm:"type:varchar(20);not null;index:idx_relationship_events_owner,priority:1" json:"ownerType"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;index:idx_relationship_events_owner,priority:2" json:"ownerId"`

	// Evento da própria pessoa (nil) ou de alguém da rede.
	RelatedPersonID *uuid.UUID `gorm:"type:uuid" json:"relatedPersonId,omitempty"`

	Type         string    `gorm:"type:varchar(20);not null" json:"type"`
	Title        string    `gorm:"type:varchar(200);not null" json:"title"`
	EventDate    time.Time `gorm:"type:date;not null;index:idx_relationship_events_date" json:"eventDate"`
	Recurring    bool      `gorm:"not null;default:false" json:"recurring"`
	LeadTimeDays int       `gorm:"not null;default:7" json:"leadTimeDays"`
	Status       string    `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`

	Source  string     `gorm:"type:varchar(20);not null;default:'ai'" json:"source"`
	AddedBy *uuid.UUID `gorm:"type:uuid" json:"addedBy,omitempty"`
	Notes   string     `gorm:"type:text;not null;default:''" json:"notes"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (RelationshipEvent) TableName() string { return "relationship_events" }

func (m *RelationshipEvent) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
