package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResultStatus define os status possíveis de um resultado laboratorial
type LabResultStatus string

const (
	LabResultPending    LabResultStatus = "pending"    // Aguardando resultado
	LabResultCompleted  LabResultStatus = "completed"  // Resultado disponível
	LabResultCancelled  LabResultStatus = "cancelled"  // Cancelado
	LabResultInProgress LabResultStatus = "in_progress" // Em andamento
)

// LabResult representa um resultado de exame laboratorial
// @Description Resultado de exame laboratorial com dados do exame e interpretação
type LabResult struct {
	// ID único do resultado
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do médico solicitante
	RequestingDoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"requestingDoctorId"`

	// Nome do exame
	TestName string `gorm:"type:varchar(200);not null" json:"testName" validate:"required"`

	// Tipo de exame (ex: hemograma, glicemia, etc)
	TestType string `gorm:"type:varchar(100);not null;index" json:"testType" validate:"required"`

	// Status do resultado
	Status LabResultStatus `gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending','completed','cancelled','in_progress')" json:"status"`

	// Data da solicitação
	RequestDate time.Time `gorm:"type:timestamp;not null;index" json:"requestDate"`

	// Data da coleta
	CollectionDate *time.Time `gorm:"type:timestamp" json:"collectionDate,omitempty"`

	// Data do resultado
	ResultDate *time.Time `gorm:"type:timestamp;index" json:"resultDate,omitempty"`

	// Resultado (pode ser JSON estruturado ou texto)
	Result *string `gorm:"type:text" json:"result,omitempty"`

	// Valores de referência
	ReferenceRange *string `gorm:"type:text" json:"referenceRange,omitempty"`

	// Unidade de medida
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Interpretação/Observações do médico
	Interpretation *string `gorm:"type:text" json:"interpretation,omitempty"`

	// Laboratório que realizou o exame
	Laboratory *string `gorm:"type:varchar(200)" json:"laboratory,omitempty"`

	// Anexos (PDFs, imagens) - armazenar URLs ou paths
	Attachments *string `gorm:"type:text" json:"attachments,omitempty"` // JSON array de URLs

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient          Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	RequestingDoctor User    `gorm:"foreignKey:RequestingDoctorID;constraint:OnDelete:RESTRICT" json:"requestingDoctor,omitempty"`
}

// TableName especifica o nome da tabela
func (LabResult) TableName() string {
	return "lab_results"
}
