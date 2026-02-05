package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabRequest representa um pedido de exames laboratoriais
// @Description Pedido de exames laboratoriais para um paciente
type LabRequest struct {
	// ID único do pedido
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	// @required
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId" validate:"required"`

	// Data do pedido
	// @required
	Date time.Time `gorm:"type:date;not null;index" json:"date" validate:"required"`

	// Lista de exames solicitados (texto livre, um exame por linha)
	// @example "Hemograma Completo\nGlicemia de Jejum\nColesterol Total"
	Exams string `gorm:"type:text;not null" json:"exams" validate:"required"`

	// Observações do médico (opcional)
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// ID do médico solicitante (opcional)
	DoctorID *uuid.UUID `gorm:"type:uuid;index" json:"doctorId,omitempty"`

	// ID do template utilizado (opcional)
	LabRequestTemplateID *uuid.UUID `gorm:"type:uuid;index" json:"labRequestTemplateId,omitempty"`

	// URL do PDF gerado (quando gerado, o pedido fica bloqueado para edição)
	PdfURL *string `gorm:"type:text" json:"pdfUrl,omitempty"`

	// Assinatura Digital ICP-Brasil (para pedidos de exames)

	// Caminho do PDF assinado
	// @example /app/uploads/lab-requests/lab_request_550e8400_signed.pdf
	SignedPDFPath *string `gorm:"type:varchar(500)" json:"signedPdfPath,omitempty"`

	// Hash SHA-256 do PDF assinado (integridade)
	// @example a1b2c3d4e5f6...
	SignedPDFHash *string `gorm:"type:varchar(64)" json:"signedPdfHash,omitempty"`

	// Dados do QR Code (URL de validação)
	// @example https://plenya.com.br/lab-requests/validate/550e8400...
	QRCodeData *string `gorm:"type:text" json:"qrCodeData,omitempty"`

	// Data/hora da assinatura digital
	SignedAt *time.Time `gorm:"type:timestamp" json:"signedAt,omitempty"`

	// Número de série do certificado usado na assinatura
	// @example 1234567890ABCDEF
	CertificateSerial *string `gorm:"type:varchar(100)" json:"certificateSerial,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos
	Patient             *Patient             `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor              *User                `gorm:"foreignKey:DoctorID;constraint:OnDelete:SET NULL" json:"doctor,omitempty"`
	LabRequestTemplate  *LabRequestTemplate  `gorm:"foreignKey:LabRequestTemplateID;constraint:OnDelete:SET NULL" json:"labRequestTemplate,omitempty"`
}

// TableName especifica o nome da tabela
func (LabRequest) TableName() string {
	return "lab_requests"
}

// BeforeCreate hook to generate UUID v7
func (lr *LabRequest) BeforeCreate(tx *gorm.DB) error {
	if lr.ID == uuid.Nil {
		lr.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
