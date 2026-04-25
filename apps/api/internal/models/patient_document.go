package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientDocumentType categoriza documentos clínicos compartilhados com o paciente.
type PatientDocumentType string

const (
	DocumentTypeCertificate PatientDocumentType = "certificate"  // atestado
	DocumentTypeReport      PatientDocumentType = "report"       // relatório
	DocumentTypeReferral    PatientDocumentType = "referral"     // encaminhamento
	DocumentTypeDeclaration PatientDocumentType = "declaration"  // declaração
	DocumentTypeOther       PatientDocumentType = "other"
)

// PatientDocument é um documento clínico (PDF, imagem, etc) que a equipe
// faz upload no EMR e o paciente consegue baixar via portal.
//
// FilePath é relativo ao /app/uploads (storage local). V2 pode migrar pra S3.
type PatientDocument struct {
	ID         uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID  uuid.UUID           `gorm:"type:uuid;not null;index" json:"patientId"`
	UploadedBy uuid.UUID           `gorm:"type:uuid;not null" json:"uploadedBy"`

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
