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
	LeadActivityMessageSent          LeadActivityType = "message_sent"           // outbound (Plenya → cliente)
	LeadActivityMessageReceived      LeadActivityType = "message_received"       // inbound (cliente → Plenya)
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
	LeadChannelPortal   LeadActivityChannel = "portal"   // mensagem enviada pelo paciente via minha.plenyasaude.com.br
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

	// ContentHTML — corpo HTML original (e-mail). Content guarda o texto plano
	// (preview/busca/IA); ContentHTML guarda o HTML pra renderização fiel no viewer
	// (iframe sandbox + sanitização no front). Cifrado igual a Content (mesmo canal).
	ContentHTML *string `gorm:"type:text;column:content_html" json:"contentHtml,omitempty"`

	// Metadados estruturados (ex: { "template": "magic_link", "wa_message_id": "wamid.X..." })
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Mídia (WhatsApp Fase 2). Setados quando a atividade carrega um arquivo
	// (foto, áudio, documento, sticker) ou payload estruturado (localização/contato).
	//
	// Dois caminhos de storage:
	//   - Arquivo de PACIENTE que vira documento de prontuário → PatientDocumentID
	//     referencia o PatientDocument; o binário vive em patient-docs/ (não cifrado,
	//     pois o interpretador de exames lê do disco). MediaStorageKey fica nil.
	//   - Mídia de Lead / áudio / não-arquivo → MediaStorageKey aponta pro bucket
	//     cifrado da conversa (whatsapp-media/...).
	MediaType       *string `gorm:"type:varchar(20)" json:"mediaType,omitempty"` // image|audio|voice|video|document|sticker|location|contacts
	MediaStorageKey *string `gorm:"type:varchar(500)" json:"-"`
	MediaMIME       *string `gorm:"type:varchar(120)" json:"mediaMime,omitempty"`
	MediaFilename   *string `gorm:"type:varchar(255)" json:"mediaFilename,omitempty"`
	MediaSizeBytes  *int64  `json:"mediaSizeBytes,omitempty"`

	// PatientDocumentID referencia o documento de prontuário quando o arquivo
	// recebido foi salvo nas mídias do paciente (foto/PDF de paciente cadastrado).
	PatientDocumentID *uuid.UUID `gorm:"type:uuid;index:idx_lead_activities_patient_doc" json:"patientDocumentId,omitempty"`

	// Transcrição de áudio (WhatsApp Fase 2.2). Texto sensível → cifrado em repouso.
	Transcription *string `gorm:"type:text" json:"transcription,omitempty"`

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

	if a.encryptedChannel() {
		if err := encryptField(&a.Content); err != nil {
			return err
		}
		if err := encryptField(&a.ContentHTML); err != nil {
			return err
		}
		if err := encryptField(&a.Transcription); err != nil {
			return err
		}
	}
	return nil
}

// encryptedChannel indica se o conteúdo textual deste canal é criptografado em
// repouso (LGPD): email e WhatsApp podem conter CPF, sintomas e dados sensíveis.
func (a *LeadActivity) encryptedChannel() bool {
	return a.Channel == LeadChannelEmail || a.Channel == LeadChannelWhatsApp
}

// encryptField cifra um *string in-place se não-vazio e ainda não cifrado.
func encryptField(field **string) error {
	v := *field
	if v == nil || *v == "" || isEncrypted(*v) {
		return nil
	}
	enc, err := crypto.EncryptWithDefaultKey(*v)
	if err != nil {
		return err
	}
	*field = &enc
	return nil
}

// AfterFind descriptografa Content/Transcription de canais cifrados (email, whatsapp).
func (a *LeadActivity) AfterFind(tx *gorm.DB) error {
	if !a.encryptedChannel() {
		return nil
	}
	// Falha silenciosa por campo: conteúdo plain antigo (pré-criptografia) ou
	// corrupção não devem crashar a query — loga e mantém o valor cru.
	decryptField(&a.Content, tx, a.ID, "content")
	decryptField(&a.ContentHTML, tx, a.ID, "content_html")
	decryptField(&a.Transcription, tx, a.ID, "transcription")
	return nil
}

func decryptField(field **string, tx *gorm.DB, id uuid.UUID, label string) {
	v := *field
	if v == nil || *v == "" || !isEncrypted(*v) {
		return
	}
	dec, err := crypto.DecryptWithDefaultKey(*v)
	if err != nil {
		tx.Logger.Warn(tx.Statement.Context, "lead_activity[%s] %s decrypt failed: %v", id, label, err)
		return
	}
	*field = &dec
}
