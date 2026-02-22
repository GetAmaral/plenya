package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResultBatchStatus define os status possíveis de um lote de resultados
type LabResultBatchStatus string

const (
	LabResultBatchPending   LabResultBatchStatus = "pending"   // Aguardando
	LabResultBatchPartial   LabResultBatchStatus = "partial"   // Parcial
	LabResultBatchCompleted LabResultBatchStatus = "completed" // Completo
)

// LabResultBatch representa um lote de resultados laboratoriais
// @Description Agrupa resultados de uma mesma coleta/laboratório
type LabResultBatch struct {
	// ID único do lote
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_batch_patient" json:"patientId"`

	// ID do pedido laboratorial (opcional - pode ser entrada manual)
	LabRequestID *uuid.UUID `gorm:"type:uuid;index:idx_batch_request" json:"labRequestId,omitempty"`

	// ID do médico solicitante (opcional)
	RequestingDoctorID *uuid.UUID `gorm:"type:uuid;index:idx_batch_doctor" json:"requestingDoctorId,omitempty"`

	// Nome do laboratório
	// @example Laboratório São Paulo, Fleury, Dasa
	LaboratoryName string `gorm:"type:varchar(200);not null" json:"laboratoryName" validate:"required,max=200"`

	// Data da coleta
	CollectionDate time.Time `gorm:"type:timestamp;not null;index" json:"collectionDate" validate:"required"`

	// Data de disponibilização dos resultados
	ResultDate *time.Time `gorm:"type:timestamp;index" json:"resultDate,omitempty"`

	// Status do lote
	// @enum pending,partial,completed
	Status LabResultBatchStatus `gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending','partial','completed')" json:"status" validate:"required"`

	// Observações gerais do lote
	Observations *string `gorm:"type:text" json:"observations,omitempty"`

	// Anexos (PDFs/scans do lote completo) - JSON array de URLs
	Attachments *string `gorm:"type:text" json:"attachments,omitempty"`

	// Conteúdo completo extraído do PDF via OCR (texto bruto)
	PDFContentFull *string `gorm:"type:text" json:"pdfContentFull,omitempty"`

	// Conteúdo processado/limpo do PDF (removendo ruído)
	PDFContentShort *string `gorm:"type:text" json:"pdfContentShort,omitempty"`

	// Conteúdo que precisa de IA (após remover exames já matched via pré-matching)
	PDFContentNeedAI *string `gorm:"type:text" json:"pdfContentNeedAi,omitempty"`

	// Conteúdo processado pela IA em formato JSON estruturado
	PDFContentJSON *string `gorm:"type:text" json:"pdfContentJson,omitempty"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos
	Patient          Patient      `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	LabRequest       *LabRequest  `gorm:"foreignKey:LabRequestID;constraint:OnDelete:SET NULL" json:"labRequest,omitempty"`
	RequestingDoctor *User        `gorm:"foreignKey:RequestingDoctorID;constraint:OnDelete:SET NULL" json:"requestingDoctor,omitempty"`
	LabResults       []LabResult  `gorm:"foreignKey:LabResultBatchID;constraint:OnDelete:CASCADE" json:"labResults,omitempty"`
}

// TableName especifica o nome da tabela
func (LabResultBatch) TableName() string {
	return "lab_result_batches"
}

// GetTitle retorna um título legível para o lote de resultados
func (lrb *LabResultBatch) GetTitle() string {
	return fmt.Sprintf("%s - %s", lrb.LaboratoryName, lrb.CollectionDate.Format("02/01/2006"))
}

// AfterFind popula DisplayTitle após carregar do banco
func (lrb *LabResultBatch) AfterFind(tx *gorm.DB) error {
	lrb.DisplayTitle = lrb.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7 (OBRIGATÓRIO)
func (lrb *LabResultBatch) BeforeCreate(tx *gorm.DB) error {
	if lrb.ID == uuid.Nil {
		lrb.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
