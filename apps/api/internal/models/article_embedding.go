package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

// ArticleEmbedding armazena embeddings vetoriais de chunks de artigos
// Cada artigo é dividido em chunks (abstract + sliding window)
// Cada chunk recebe um vetor de embedding para busca semântica
type ArticleEmbedding struct {
	// Primary Key
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Foreign Key - Artigo original
	ArticleID uuid.UUID `gorm:"type:uuid;not null;index" json:"articleId" validate:"required"`

	// Chunk metadata
	// @minimum 0
	// @example 0
	ChunkIndex int `gorm:"type:integer;not null" json:"chunkIndex" validate:"gte=0"`

	// @minLength 1
	ChunkText string `gorm:"type:text;not null" json:"chunkText" validate:"required,min=1"`

	// Metadata estruturado: {section: "abstract|methods|results", page_range: [1,3], word_count: 250}
	ChunkMetadata map[string]interface{} `gorm:"type:jsonb;default:'{}'" json:"chunkMetadata,omitempty"`

	// Vector embedding (OpenAI text-embedding-3-large: 1024 dimensões)
	// Stored as pgvector type for efficient similarity search
	Embedding pgvector.Vector `gorm:"type:vector(1024)" json:"-"` // Não expor no JSON (muito grande)

	// Relationships
	Article *Article `gorm:"foreignKey:ArticleID;constraint:OnDelete:CASCADE" json:"article,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

// TableName especifica o nome da tabela no banco de dados
func (ArticleEmbedding) TableName() string {
	return "article_embeddings"
}

// BeforeCreate hook - executado antes de criar registro
func (ae *ArticleEmbedding) BeforeCreate(tx *gorm.DB) error {
	if ae.ID == uuid.Nil {
		ae.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// GetEmbeddingSlice retorna embedding como []float32 para uso em código Go
func (ae *ArticleEmbedding) GetEmbeddingSlice() []float32 {
	if ae.Embedding.Slice() == nil {
		return nil
	}
	return ae.Embedding.Slice()
}

// SetEmbeddingFromSlice define embedding a partir de []float32
func (ae *ArticleEmbedding) SetEmbeddingFromSlice(values []float32) error {
	ae.Embedding = pgvector.NewVector(values)
	return nil
}
