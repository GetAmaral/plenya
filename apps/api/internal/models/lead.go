package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// LeadSource indica de onde o lead foi capturado.
type LeadSource string

const (
	LeadSourceLightClaim      LeadSource = "light_claim"      // captura no claim do Escore Light
	LeadSourceContactForm     LeadSource = "contact_form"     // formulário /contato no site público
	LeadSourceWhatsAppInbound LeadSource = "whatsapp_inbound" // primeira mensagem inbound do cliente no WhatsApp
	LeadSourceNewsletter      LeadSource = "newsletter"       // opt-in newsletter sem contexto de claim
	LeadSourceManual          LeadSource = "manual"           // criado manualmente por staff
)

// LeadStatus representa o estágio do funil em que o lead se encontra.
type LeadStatus string

const (
	LeadStatusNew          LeadStatus = "new"
	LeadStatusContacted    LeadStatus = "contacted"
	LeadStatusQualified    LeadStatus = "qualified"
	LeadStatusConverted    LeadStatus = "converted" // virou Patient
	LeadStatusLost         LeadStatus = "lost"
	LeadStatusUnsubscribed LeadStatus = "unsubscribed" // pediu pra parar de receber
)

// Lead representa um contato capturado por qualquer canal antes (ou em paralelo) à conversão em Patient.
//
// @Description Lead capturado em algum dos canais públicos da Plenya
type Lead struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Origem do lead (canal/contexto de captura)
	// @enum light_claim,contact_form,whatsapp_inbound,newsletter,manual
	// @example light_claim
	Source LeadSource `gorm:"type:varchar(40);not null;index:idx_leads_source" json:"source" validate:"required,oneof=light_claim contact_form whatsapp_inbound newsletter manual"`

	// Status atual no pipeline
	// @enum new,contacted,qualified,converted,lost,unsubscribed
	// @example new
	Status LeadStatus `gorm:"type:varchar(40);not null;default:'new';index:idx_leads_status" json:"status" validate:"required,oneof=new contacted qualified converted lost unsubscribed"`

	// Nome (opcional — pode ser preenchido depois)
	// @example Maria Silva
	Name *string `gorm:"type:varchar(255)" json:"name,omitempty"`

	// Email de contato (opcional, mas pelo menos um de Email/Phone obrigatório por validação aplicação)
	// @example maria@exemplo.com
	Email *string `gorm:"type:varchar(255);index:idx_leads_email" json:"email,omitempty"`

	// Telefone E.164 (ex: +5511999998888)
	// @example +5511999998888
	Phone *string `gorm:"type:varchar(20);index:idx_leads_phone" json:"phone,omitempty"`

	// Mensagem livre — primeiro contato (motivo do form, primeira msg do WA)
	Message *string `gorm:"type:text" json:"message,omitempty"`

	// Metadados estruturados (ex: { "reason": "cansaco", "window": "tarde" })
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Opt-ins granulares (LGPD — base legal por canal)
	EmailOptIn      bool `gorm:"default:false" json:"emailOptIn"`
	WhatsAppOptIn   bool `gorm:"default:false" json:"whatsAppOptIn"`
	NewsletterOptIn bool `gorm:"default:false" json:"newsletterOptIn"`

	// LGPD — registro de consentimento (art. 8º §6º: ônus da prova é do controlador)
	ConsentVersion   *string    `gorm:"type:varchar(20)" json:"consentVersion,omitempty"`
	ConsentTimestamp *time.Time `gorm:"type:timestamp" json:"consentTimestamp,omitempty"`
	ConsentIPHash    *string    `gorm:"type:varchar(64)" json:"-"`

	// Vinculação opcional com sessão Light (lead originado de claim).
	// uniqueIndex pra garantir 1 Lead por sessão Light — previne duplicatas em race
	// (cliente clica "Receber meu link" 2× rapidamente).
	AnonymousScoreSessionID *uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_leads_session_unique" json:"anonymousScoreSessionId,omitempty"`

	// Conversão em Patient
	ConvertedPatientID *uuid.UUID `gorm:"type:uuid;index:idx_leads_patient" json:"convertedPatientId,omitempty"`
	ConvertedAt        *time.Time `gorm:"type:timestamp" json:"convertedAt,omitempty"`
	ConvertedByUserID  *uuid.UUID `gorm:"type:uuid" json:"convertedByUserId,omitempty"`

	// Atribuição manual a um membro do time (Fase 2 vai automatizar)
	AssignedToUserID *uuid.UUID `gorm:"type:uuid;index:idx_leads_assigned" json:"assignedToUserId,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Activities       []LeadActivity         `gorm:"foreignKey:LeadID;constraint:OnDelete:CASCADE" json:"activities,omitempty"`
	ConvertedPatient *Patient               `gorm:"foreignKey:ConvertedPatientID;constraint:OnDelete:SET NULL" json:"convertedPatient,omitempty"`
	AssignedTo       *User                  `gorm:"foreignKey:AssignedToUserID;constraint:OnDelete:SET NULL" json:"assignedTo,omitempty"`
	Session          *AnonymousScoreSession `gorm:"foreignKey:AnonymousScoreSessionID;constraint:OnDelete:SET NULL" json:"-"`
}

// TableName especifica o nome da tabela
func (Lead) TableName() string {
	return "leads"
}

// BeforeCreate hook gera UUID v7 e default de status
func (l *Lead) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.Must(uuid.NewV7())
	}
	if l.Status == "" {
		l.Status = LeadStatusNew
	}
	return nil
}
