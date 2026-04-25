package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// FamilyAccessScope é o conjunto granular de áreas que o paciente concede a
// um familiar (cônjuge/filho/cuidador). Salvamos como JSONB pra evoluir sem migration.
type FamilyAccessScope struct {
	Appointments        bool `json:"appointments"`
	Continuum           bool `json:"continuum"`
	Exams               bool `json:"exams"`         // labs + prescrições + avaliações
	Scores              bool `json:"scores"`
	Documents           bool `json:"documents"`
	Boxes               bool `json:"boxes"`
	Messages            bool `json:"messages"`      // mensagens com a equipe (read-only do contato)
}

// FamilyGrantStatus enumera estados.
type FamilyGrantStatus string

const (
	FamilyGrantInvited  FamilyGrantStatus = "invited"  // convite enviado, ainda não aceito
	FamilyGrantActive   FamilyGrantStatus = "active"   // aceito, em vigor
	FamilyGrantRevoked  FamilyGrantStatus = "revoked"  // revogado pelo paciente ou expirado
)

// FamilyAccessGrant — paciente (PatientID) concede a um terceiro (GranteeUserID,
// pode ser nil enquanto convite pendente — vincula no consume) acesso parcial
// aos próprios dados.
//
// Token é usado no link de convite. ScopeJSON guarda o FamilyAccessScope acima.
type FamilyAccessGrant struct {
	ID            uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID     uuid.UUID         `gorm:"type:uuid;not null;index" json:"patientId"`

	// Email convidado (anti-enumeração: validamos contra User.Email no consume).
	GranteeEmail  string            `gorm:"type:varchar(255);not null" json:"granteeEmail"`

	// Nome amigável (ex: "Esposa Mariana"). Aparece pro paciente na lista.
	GranteeLabel  string            `gorm:"type:varchar(120);not null" json:"granteeLabel"`

	// Vinculado quando o convite é consumido (ou quando o User pré-existe).
	GranteeUserID *uuid.UUID        `gorm:"type:uuid;index" json:"granteeUserId,omitempty"`

	Token         string            `gorm:"type:varchar(128);uniqueIndex;not null" json:"-"`
	ScopeJSON     datatypes.JSON    `gorm:"type:jsonb;not null" json:"scope"`

	Status        FamilyGrantStatus `gorm:"type:varchar(20);not null;default:'invited';check:status IN ('invited','active','revoked')" json:"status"`

	ExpiresAt     time.Time         `gorm:"not null" json:"expiresAt"`
	AcceptedAt    *time.Time        `json:"acceptedAt,omitempty"`
	RevokedAt     *time.Time        `json:"revokedAt,omitempty"`

	CreatedAt     time.Time         `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time         `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt    `gorm:"index" json:"-"`
}

func (FamilyAccessGrant) TableName() string { return "family_access_grants" }

func (g *FamilyAccessGrant) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.Must(uuid.NewV7())
	}
	if g.ExpiresAt.IsZero() {
		// Convites expiram em 14 dias se não consumidos. Grants ativos não expiram
		// (ficam até paciente revogar).
		g.ExpiresAt = time.Now().UTC().Add(14 * 24 * time.Hour)
	}
	return nil
}

// IsActive responde true se grant pode ser usado pra acesso real.
func (g *FamilyAccessGrant) IsActive() bool {
	return g.Status == FamilyGrantActive && g.RevokedAt == nil
}
