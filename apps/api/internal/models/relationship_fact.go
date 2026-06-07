package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Categorias da taxonomia controlada de fatos sociais (§3 do plano). NUNCA clínico.
const (
	RelationshipCategoryIdentity    = "identidade_social"       // apelido, profissão, cidade, idioma
	RelationshipCategoryFamily      = "familia_rede"            // cônjuge, filhos, quem acompanha/influencia
	RelationshipCategoryPreferences = "preferencias_atendimento" // canal, presencial/online, horário
	RelationshipCategoryArrival     = "contexto_chegada"        // como conheceu, quem indicou, motivo (não-clínico)
	RelationshipCategoryRelationship = "relacionamento"         // estágio, sensibilidades de abordagem, datas
)

// Fontes possíveis de um fato.
const (
	RelationshipSourceAI       = "ai"
	RelationshipSourceStaff    = "staff"
	RelationshipSourceForm     = "form"
	RelationshipSourceConsulta = "consulta"
)

// RelationshipFact — fato atômico SOCIAL sobre uma pessoa (memória semântica da Lívia). N por
// (owner_type, owner_id). Validade temporal: ValidUntil nil = ativo. NÃO contém dado clínico
// (LGPD/CFM): o clínico vive no prontuário. Ver docs/emr/plano-livia-memoria-dossie-360.md (§2.2/§3).
//
// @Description Fato social atômico de relacionamento de um lead/paciente
type RelationshipFact struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	OwnerType string    `gorm:"type:varchar(20);not null;index:idx_relationship_fact_owner,priority:1" json:"ownerType"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;index:idx_relationship_fact_owner,priority:2" json:"ownerId"`

	Category string `gorm:"type:varchar(40);not null" json:"category"`
	Key      string `gorm:"column:key;type:varchar(60);not null" json:"key"`
	Value    string `gorm:"type:text;not null" json:"value"`

	Source     string     `gorm:"type:varchar(20);not null;default:'ai'" json:"source"`
	AddedBy    *uuid.UUID `gorm:"type:uuid" json:"addedBy,omitempty"`
	Confidence *float32   `json:"confidence,omitempty"`

	ValidFrom  time.Time  `gorm:"autoCreateTime" json:"validFrom"`
	ValidUntil *time.Time `json:"validUntil,omitempty"` // nil = ativo

	// Restrito: aparece para a equipe no EMR, mas é EXCLUÍDO do que a IA enxerga (§12). Default false.
	Restricted bool `gorm:"column:restricted;not null;default:false" json:"restricted"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name
func (RelationshipFact) TableName() string { return "relationship_facts" }

// BeforeCreate hook generates UUID v7
func (m *RelationshipFact) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
