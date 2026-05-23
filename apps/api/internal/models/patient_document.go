package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientDocumentType categoriza documentos clínicos compartilhados com o paciente.
type PatientDocumentType string

const (
	DocumentTypeCertificate PatientDocumentType = "certificate" // atestado
	DocumentTypeReport      PatientDocumentType = "report"      // relatório
	DocumentTypeReferral    PatientDocumentType = "referral"    // encaminhamento
	DocumentTypeDeclaration PatientDocumentType = "declaration" // declaração
	DocumentTypeOther       PatientDocumentType = "other"
)

// PatientDocumentSource indica a origem do documento (auditoria).
type PatientDocumentSource string

const (
	DocumentSourceStaffUpload PatientDocumentSource = "staff_upload" // upload manual pela equipe
	DocumentSourceWhatsApp    PatientDocumentSource = "whatsapp"     // recebido por WhatsApp inbound
	DocumentSourceEmail       PatientDocumentSource = "email"        // anexo recebido por e-mail
	DocumentSourcePortal      PatientDocumentSource = "portal"       // enviado pelo paciente via portal
)

// PatientDocument é um documento clínico (PDF, imagem, etc) que a equipe
// faz upload no EMR e o paciente consegue baixar via portal.
//
// FilePath é relativo ao /app/uploads (storage local). V2 pode migrar pra S3.
type PatientDocument struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// UploadedBy é nullable: documentos recebidos por WhatsApp/portal não têm
	// usuário staff responsável (Source indica a origem real).
	UploadedBy *uuid.UUID `gorm:"type:uuid" json:"uploadedBy,omitempty"`

	// Source: origem do documento (staff_upload | whatsapp | portal).
	Source PatientDocumentSource `gorm:"type:varchar(20);not null;default:'staff_upload'" json:"source"`

	// OriginWAMessageID: wa_message_id que originou o documento (quando Source=whatsapp).
	// Usado pra idempotência do webhook (Meta reentrega mensagens).
	// uniqueIndex: idempotência do webhook (Postgres trata NULLs como distintos,
	// então uploads manuais sem origem não conflitam).
	OriginWAMessageID *string `gorm:"type:varchar(120);uniqueIndex:uq_patient_documents_wa_msg" json:"-"`

	Type        PatientDocumentType `gorm:"type:varchar(30);not null;check:type IN ('certificate','report','referral','declaration','other')" json:"type"`
	Title       string              `gorm:"type:varchar(200);not null" json:"title"`
	Description string              `gorm:"type:text" json:"description"`

	FilePath    string `gorm:"type:varchar(500);not null" json:"-"` // relativo ao uploadsRoot
	FileName    string `gorm:"type:varchar(200);not null" json:"fileName"`
	ContentType string `gorm:"type:varchar(100);not null" json:"contentType"`
	SizeBytes   int64  `gorm:"not null" json:"sizeBytes"`

	IssuedAt time.Time `gorm:"type:timestamp;not null" json:"issuedAt"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UploadedByUser *User `gorm:"foreignKey:UploadedBy" json:"uploadedByUser,omitempty"`
}

func (PatientDocument) TableName() string { return "patient_documents" }

func (d *PatientDocument) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.Must(uuid.NewV7())
	}
	if d.IssuedAt.IsZero() {
		d.IssuedAt = time.Now().UTC()
	}
	return nil
}
