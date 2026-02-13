package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

// ScoreItemEmbedding armazena embedding vetorial de um ScoreItem
// Usado para recomendar artigos relevantes para parâmetros clínicos
// e descobrir quais parâmetros um artigo cobre
type ScoreItemEmbedding struct {
	// Primary Key
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Foreign Key - ScoreItem (relação 1:1)
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"scoreItemId" validate:"required"`

	// Vector embedding (OpenAI text-embedding-3-large: 1024 dimensões)
	Embedding pgvector.Vector `gorm:"type:vector(1024)" json:"-"` // Não expor no JSON

	// Texto fonte usado para gerar embedding
	// Combinação: name + clinical_relevance + patient_explanation + conduct
	// @minLength 1
	TextSource string `gorm:"type:text;not null" json:"textSource" validate:"required,min=1"`

	// Relationships
	ScoreItem *ScoreItem `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:CASCADE" json:"scoreItem,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName especifica o nome da tabela no banco de dados
func (ScoreItemEmbedding) TableName() string {
	return "score_item_embeddings"
}

// BeforeCreate hook - executado antes de criar registro
func (sie *ScoreItemEmbedding) BeforeCreate(tx *gorm.DB) error {
	if sie.ID == uuid.Nil {
		sie.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// GetEmbeddingSlice retorna embedding como []float32 para uso em código Go
func (sie *ScoreItemEmbedding) GetEmbeddingSlice() []float32 {
	if sie.Embedding.Slice() == nil {
		return nil
	}
	return sie.Embedding.Slice()
}

// SetEmbeddingFromSlice define embedding a partir de []float32
func (sie *ScoreItemEmbedding) SetEmbeddingFromSlice(values []float32) error {
	sie.Embedding = pgvector.NewVector(values)
	return nil
}
