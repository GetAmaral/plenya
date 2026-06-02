package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// IssuedDocumentType — documento clínico GERADO e (idealmente) assinado no EMR.
// Distinto de PatientDocument, que é metadado de UPLOAD de arquivo.
type IssuedDocumentType string

const (
	IssuedDocCertificate IssuedDocumentType = "certificate" // atestado médico
	IssuedDocDeclaration IssuedDocumentType = "declaration" // declaração de comparecimento
	IssuedDocReport      IssuedDocumentType = "report"      // relatório/laudo
)

// IssuedDocumentStatus — draft (editável) → signed (imutável, PDF gerado/assinado).
type IssuedDocumentStatus string

const (
	IssuedDocDraft  IssuedDocumentStatus = "draft"
	IssuedDocSigned IssuedDocumentStatus = "signed"
)

// IssuedDocument é um documento clínico emitido pelo médico (atestado, declaração, laudo).
// No "sign" o PDF é gerado (gofpdf), assinado com ICP-Brasil quando há certificado
// (degrada para assinatura manual quando não há), publicado como PatientDocument
// (visível/baixável no portal) e os metadados de assinatura ficam aqui para validação por QR.
//
// @Description Documento clínico emitido e assinável (atestado/declaração/laudo)
type IssuedDocument struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"patientId"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"appointmentId,omitempty"`
	DoctorID      uuid.UUID  `gorm:"type:uuid;not null" json:"doctorId"`

	Type    IssuedDocumentType `gorm:"type:varchar(20);not null;check:type IN ('certificate','declaration','report')" json:"type"`
	Title   string             `gorm:"type:varchar(200);not null" json:"title"`
	Body    string             `gorm:"type:text;not null;default:''" json:"body"`
	Purpose *string            `gorm:"type:varchar(300)" json:"purpose,omitempty"`

	// DaysOff — dias de afastamento (atestado).
	DaysOff *int `gorm:"type:int" json:"daysOff,omitempty"`

	// CID só consta com consentimento do paciente (sigilo). Default: sem CID.
	IncludesCID bool    `gorm:"not null;default:false;column:includes_cid" json:"includesCid"`
	CIDCode     *string `gorm:"type:varchar(10);column:cid_code" json:"cidCode,omitempty"`
	CIDConsent  bool    `gorm:"not null;default:false;column:cid_consent" json:"cidConsent"`

	Status              IssuedDocumentStatus `gorm:"type:varchar(10);not null;default:'draft';check:status IN ('draft','signed')" json:"status"`
	HasDigitalSignature bool                 `gorm:"not null;default:false" json:"hasDigitalSignature"`
	SignedAt            *time.Time           `gorm:"type:timestamptz" json:"signedAt,omitempty"`
	SignedPDFPath       *string              `gorm:"type:varchar(500)" json:"-"` // relativo ao uploadsRoot
	SignedPDFHash       *string              `gorm:"type:varchar(64)" json:"signedPdfHash,omitempty"`
	CertificateSerial   *string              `gorm:"type:varchar(100)" json:"certificateSerial,omitempty"`
	QRCodeData          *string              `gorm:"type:text" json:"qrCodeData,omitempty"`

	// PatientDocumentID — cópia publicada no portal (download autenticado).
	PatientDocumentID *uuid.UUID `gorm:"type:uuid" json:"patientDocumentId,omitempty"`

	IssuedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"issuedByUserId"`
	IssuedAt       time.Time `gorm:"type:timestamptz;not null;default:now()" json:"issuedAt"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	Doctor  User    `gorm:"foreignKey:DoctorID;constraint:OnDelete:RESTRICT" json:"doctor,omitempty"`
}

func (IssuedDocument) TableName() string { return "issued_documents" }

func (d *IssuedDocument) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.Must(uuid.NewV7())
	}
	if d.IssuedAt.IsZero() {
		d.IssuedAt = time.Now().UTC()
	}
	return nil
}
