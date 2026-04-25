package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/crypto"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ErrLeadActivityOwnerInvalid é retornado quando exatamente um de
// LeadID/PatientID NÃO está setado. A regra refletida em CHECK constraint
// no banco (lead_activities_owner_check). Documentamos aqui pra falha precoce
// no app antes de bater no DB.
var ErrLeadActivityOwnerInvalid = errors.New("lead_activity: exatamente um de LeadID|PatientID obrigatório")

// LeadActivityType categoriza eventos no histórico do lead.
type LeadActivityType string

const (
	LeadActivityCreated              LeadActivityType = "created"
	LeadActivityMessageSent          LeadActivityType = "message_sent"     // outbound (Plenya → cliente)
	LeadActivityMessageReceived      LeadActivityType = "message_received" // inbound (cliente → Plenya)
	LeadActivityMessageStatusChanged LeadActivityType = "message_status_changed" // delivered/read/failed (WhatsApp)
	LeadActivityStatusChanged        LeadActivityType = "status_changed"
	LeadActivityNoteAdded            LeadActivityType = "note_added"
	LeadActivityConverted            LeadActivityType = "converted"
	LeadActivityAssigned             LeadActivityType = "assigned"
	LeadActivityUnsubscribed         LeadActivityType = "unsubscribed"
)

// LeadActivityChannel indica por qual canal a atividade ocorreu.
type LeadActivityChannel string

const (
	LeadChannelEmail    LeadActivityChannel = "email"
	LeadChannelWhatsApp LeadActivityChannel = "whatsapp"
	LeadChannelInternal LeadActivityChannel = "internal" // notas, mudanças de status feitas no admin
	LeadChannelPortal   LeadActivityChannel = "portal"   // mensagem enviada pelo paciente via meu.plenyasaude.com.br
)

// LeadActivity é o log imutável de eventos relacionados a uma "conversa" do CRM.
//
// Ownership: exatamente UM de LeadID|PatientID é setado.
//   - LeadID setado, PatientID nil  → conversa pertence a Lead (ainda não convertido).
//   - PatientID setado, LeadID nil  → conversa pertence a Patient (cliente cadastrado).
//
// CHECK constraint no banco (lead_activities_owner_check) garante a invariante.
// Validação aplicação espelha em BeforeSave pra falha precoce.
//
// @Description Evento no histórico de uma conversa (lead OU patient)
type LeadActivity struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Lead a que esta atividade pertence (mutuamente exclusivo com PatientID)
	LeadID *uuid.UUID `gorm:"type:uuid;index:idx_lead_activities_lead" json:"leadId,omitempty"`

	// Patient a que esta atividade pertence (mutuamente exclusivo com LeadID).
	// Setado quando inbound chega de cliente já cadastrado — evita criar Lead "fantasma".
	PatientID *uuid.UUID `gorm:"type:uuid;index:idx_lead_activities_patient" json:"patientId,omitempty"`

	// Tipo do evento
	// @enum created,message_sent,message_received,status_changed,note_added,converted,assigned,unsubscribed
	Type LeadActivityType `gorm:"type:varchar(40);not null;index:idx_lead_activities_type" json:"type" validate:"required"`

	// Canal (email, whatsapp, internal)
	// @enum email,whatsapp,internal
	Channel LeadActivityChannel `gorm:"type:varchar(20)" json:"channel,omitempty"`

	// Conteúdo livre — corpo da mensagem, nota, descrição da mudança
	Content *string `gorm:"type:text" json:"content,omitempty"`

	// Metadados estruturados (ex: { "template": "magic_link", "wa_message_id": "wamid.X..." })
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Usuário responsável pela ação (null se inbound do cliente ou ação do sistema)
	ActorUserID *uuid.UUID `gorm:"type:uuid" json:"actorUserId,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime;index:idx_lead_activities_created" json:"createdAt"`

	// Relationships
	Lead    *Lead    `gorm:"foreignKey:LeadID;constraint:OnDelete:CASCADE" json:"-"`
	Patient *Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Actor   *User    `gorm:"foreignKey:ActorUserID;constraint:OnDelete:SET NULL" json:"actor,omitempty"`
}

// TableName especifica o nome da tabela
func (LeadActivity) TableName() string {
	return "lead_activities"
}

// BeforeCreate hook gera UUID v7
func (a *LeadActivity) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave faz duas coisas:
//  1. Valida ownership (XOR LeadID/PatientID) — espelho da CHECK constraint do DB.
//  2. Criptografa Content quando Channel=email (LGPD: corpo de email pode conter
//     CPF, sintomas, dados sensíveis de paciente). Idempotente via isEncrypted
//     heurística reusada de patient.go (mesmo package).
//
// Funciona transparente pra LeadActivity vinculada a Patient — criptografia é por
// canal, não por dono.
func (a *LeadActivity) BeforeSave(tx *gorm.DB) error {
	hasLead := a.LeadID != nil && *a.LeadID != uuid.Nil
	hasPatient := a.PatientID != nil && *a.PatientID != uuid.Nil
	if hasLead == hasPatient { // ambos true ou ambos false
		return ErrLeadActivityOwnerInvalid
	}

	if a.Channel == LeadChannelEmail && a.Content != nil && *a.Content != "" && !isEncrypted(*a.Content) {
		encrypted, err := crypto.EncryptWithDefaultKey(*a.Content)
		if err != nil {
			return err
		}
		a.Content = &encrypted
	}
	return nil
}

// AfterFind descriptografa Content de email quando carregado.
func (a *LeadActivity) AfterFind(tx *gorm.DB) error {
	if a.Channel == LeadChannelEmail && a.Content != nil && *a.Content != "" && isEncrypted(*a.Content) {
		decrypted, err := crypto.DecryptWithDefaultKey(*a.Content)
		if err != nil {
			// Falha silenciosa: pode ser conteúdo plain antigo (pré-criptografia)
			// ou corrupção. Não vamos crashar a query — só loga e devolve cipher.
			tx.Logger.Warn(tx.Statement.Context, "lead_activity[%s] decrypt failed: %v", a.ID, err)
			return nil
		}
		a.Content = &decrypted
	}
	return nil
}
