package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientPortalInvite registra o convite gerado pela equipe (admin/manager/secretary)
// pra um paciente acessar a área do paciente em minha.plenyasaude.com.br.
//
// Token é a chave usada no link enviado por email/WhatsApp. Single-use:
// quando o paciente abre o link, marcamos AcceptedAt e o token deixa de funcionar.
//
// O paciente pode:
//   - Definir senha (login email+senha tradicional dali em diante)
//   - Pular a senha (continuará entrando só por magic link)
//
// Convites expiram em 14 dias por padrão. Se vencer, equipe gera outro.
type PatientPortalInvite struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	InvitedBy uuid.UUID `gorm:"type:uuid;not null" json:"invitedBy"`

	// Token único na URL — geramos 64 chars hex (256 bits).
	Token string `gorm:"type:varchar(128);uniqueIndex;not null" json:"-"`

	// Janela de validade (default 14 dias).
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`

	// Se nil, ainda não foi aceito. Quando preenchido, token vira inutilizável.
	AcceptedAt *time.Time `json:"acceptedAt,omitempty"`

	// Por onde foi despachado (multi-canal opcional).
	ChannelEmail bool `gorm:"default:true" json:"channelEmail"`
	ChannelWA    bool `gorm:"default:false" json:"channelWA"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (PatientPortalInvite) TableName() string {
	return "patient_portal_invites"
}

func (i *PatientPortalInvite) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.Must(uuid.NewV7())
	}
	if i.ExpiresAt.IsZero() {
		i.ExpiresAt = time.Now().UTC().Add(14 * 24 * time.Hour)
	}
	return nil
}

// IsValid checa se o convite ainda pode ser usado.
func (i *PatientPortalInvite) IsValid() bool {
	return i.AcceptedAt == nil && time.Now().UTC().Before(i.ExpiresAt)
}
