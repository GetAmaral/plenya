package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProcessingJobStatus representa o status de um job de processamento
type ProcessingJobStatus string

const (
	ProcessingJobPending    ProcessingJobStatus = "pending"
	ProcessingJobProcessing ProcessingJobStatus = "processing"
	ProcessingJobCompleted  ProcessingJobStatus = "completed"
	ProcessingJobFailed     ProcessingJobStatus = "failed"
)

// ProcessingJobType representa o tipo de job de processamento
type ProcessingJobType string

const (
	ProcessingJobTypePDFExtraction ProcessingJobType = "pdf_extraction"
)

// ProcessingJob representa um job de processamento assíncrono (ex: extração de PDF)
// @Description Job de processamento assíncrono com queue em PostgreSQL
type ProcessingJob struct {
	// ID único do job (UUID v7)
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do lote de resultados associado
	LabResultBatchID uuid.UUID `gorm:"type:uuid;not null;index:idx_job_batch" json:"labResultBatchId" validate:"required"`

	// Tipo de job
	Type ProcessingJobType `gorm:"type:varchar(50);not null" json:"type" validate:"required"`

	// Status do job
	Status ProcessingJobStatus `gorm:"type:varchar(20);not null;default:'pending';index:idx_job_status;check:status IN ('pending','processing','completed','failed')" json:"status"`

	// Caminho do arquivo PDF
	PDFPath string `gorm:"type:text;not null" json:"pdfPath" validate:"required"`

	// Texto extraído via OCR
	ExtractedText *string `gorm:"type:text" json:"extractedText,omitempty"`

	// Resposta bruta da IA (JSON)
	AIResponse *string `gorm:"type:text" json:"aiResponse,omitempty"`

	// Mensagem de erro (se falhou)
	ErrorMessage *string `gorm:"type:text" json:"errorMessage,omitempty"`

	// Número de tentativas realizadas
	Attempts int `gorm:"type:integer;not null;default:0" json:"attempts"`

	// Máximo de tentativas permitidas
	MaxAttempts int `gorm:"type:integer;not null;default:3" json:"maxAttempts"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime;index:idx_job_created" json:"createdAt"`

	// Data de início do processamento
	StartedAt *time.Time `gorm:"type:timestamp" json:"startedAt,omitempty"`

	// Data de conclusão (sucesso ou falha)
	CompletedAt *time.Time `gorm:"type:timestamp" json:"completedAt,omitempty"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relação com LabResultBatch
	LabResultBatch LabResultBatch `gorm:"foreignKey:LabResultBatchID;constraint:OnDelete:CASCADE" json:"labResultBatch,omitempty"`
}

// TableName especifica o nome da tabela
func (ProcessingJob) TableName() string {
	return "processing_jobs"
}

// BeforeCreate hook para gerar UUID v7
func (pj *ProcessingJob) BeforeCreate(tx *gorm.DB) error {
	if pj.ID == uuid.Nil {
		pj.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
