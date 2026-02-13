package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EmbeddingQueue gerencia fila de processamento assíncrono de embeddings
// Worker processa jobs pendentes em background
type EmbeddingQueue struct {
	// Primary Key
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Entidade a processar
	// @enum article,score_item
	EntityType string `gorm:"type:varchar(50);not null;check:entity_type IN ('article','score_item')" json:"entityType" validate:"required,oneof=article score_item"`

	EntityID uuid.UUID `gorm:"type:uuid;not null" json:"entityId" validate:"required"`

	// Status do job
	// @enum pending,processing,completed,failed
	Status string `gorm:"type:varchar(20);default:'pending';check:status IN ('pending','processing','completed','failed')" json:"status" validate:"oneof=pending processing completed failed"`

	// @minimum 0
	// @maximum 5
	RetryCount int `gorm:"type:integer;default:0;check:retry_count >= 0 AND retry_count <= 5" json:"retryCount" validate:"gte=0,lte=5"`

	ErrorMessage *string `gorm:"type:text" json:"errorMessage,omitempty"`

	// Timestamps
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	ProcessedAt *time.Time `gorm:"type:timestamp" json:"processedAt,omitempty"`
}

// TableName especifica o nome da tabela no banco de dados
func (EmbeddingQueue) TableName() string {
	return "embedding_queue"
}

// BeforeCreate hook - executado antes de criar registro
func (eq *EmbeddingQueue) BeforeCreate(tx *gorm.DB) error {
	if eq.ID == uuid.Nil {
		eq.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// MarkAsProcessing atualiza status para "processing"
func (eq *EmbeddingQueue) MarkAsProcessing(tx *gorm.DB) error {
	eq.Status = "processing"
	return tx.Save(eq).Error
}

// MarkAsCompleted atualiza status para "completed" e registra timestamp
func (eq *EmbeddingQueue) MarkAsCompleted(tx *gorm.DB) error {
	now := time.Now()
	eq.Status = "completed"
	eq.ProcessedAt = &now
	return tx.Save(eq).Error
}

// MarkAsFailed atualiza status para "failed", incrementa retry e registra erro
func (eq *EmbeddingQueue) MarkAsFailed(tx *gorm.DB, errorMsg string) error {
	now := time.Now()
	eq.Status = "failed"
	eq.RetryCount++
	eq.ErrorMessage = &errorMsg
	eq.ProcessedAt = &now
	return tx.Save(eq).Error
}

// CanRetry verifica se job pode ser reprocessado
func (eq *EmbeddingQueue) CanRetry() bool {
	return eq.RetryCount < 5 && eq.Status == "failed"
}
