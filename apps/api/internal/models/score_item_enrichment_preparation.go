package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ScoreItemEnrichmentPreparation armazena chunks selecionados para enrichment
// ETAPA 1 do plano de enrichment científico (preparação sem IA)
// @Description Preparação de chunks científicos para enrichment via Claude
type ScoreItemEnrichmentPreparation struct {
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Foreign Key - ScoreItem a ser enriquecido
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ScoreItemID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"scoreItemId" validate:"required"`

	// Chunks selecionados via RAG (15-20 chunks)
	// Estrutura: [{article_id, article_title, article_year, journal, chunk_index, chunk_text, section, similarity, word_count}, ...]
	// @example [{"article_id":"019c...", "article_title":"Effects of...", "section":"results", "similarity":0.89, ...}]
	SelectedChunks datatypes.JSONMap `gorm:"type:jsonb;not null;default:'[]'::jsonb" json:"selectedChunks"`

	// Metadata agregado da seleção
	// Estrutura: {total_chunks: 20, articles_count: 5, avg_similarity: 0.83, sections_distribution: {results: 8, discussion: 6, ...}}
	// @example {"total_chunks":20, "articles_count":5, "avg_similarity":0.83}
	Metadata datatypes.JSONMap `gorm:"type:jsonb;default:'{}'" json:"metadata,omitempty"`

	// Prompts prontos para Claude (gerados a partir dos chunks)
	// Prompt para gerar clinical_relevance (1200-1800 chars)
	PromptClinicalRelevance *string `gorm:"type:text" json:"promptClinicalRelevance,omitempty"`

	// Prompt para gerar patient_explanation (600-900 chars)
	PromptPatientExplanation *string `gorm:"type:text" json:"promptPatientExplanation,omitempty"`

	// Prompt para gerar conduct (1000-1500 chars)
	PromptConduct *string `gorm:"type:text" json:"promptConduct,omitempty"`

	// Prompt para gerar max_points (0-50)
	PromptMaxPoints *string `gorm:"type:text" json:"promptMaxPoints,omitempty"`

	// Status do processamento
	// @enum ready,processing,completed,failed,stale
	// @example ready
	Status string `gorm:"type:varchar(20);not null;default:'ready';check:status IN ('ready','processing','completed','failed','stale')" json:"status" validate:"required,oneof=ready processing completed failed stale"`

	// Mensagem de erro (se status = failed)
	// @example Failed to fetch chunks: connection timeout
	ErrorMessage *string `gorm:"type:text" json:"errorMessage,omitempty"`

	// Relationships
	ScoreItem *ScoreItem `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:CASCADE" json:"scoreItem,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName especifica o nome da tabela no banco de dados
func (ScoreItemEnrichmentPreparation) TableName() string {
	return "score_item_enrichment_preparation"
}

// BeforeCreate hook para gerar UUID v7
func (prep *ScoreItemEnrichmentPreparation) BeforeCreate(tx *gorm.DB) error {
	if prep.ID == uuid.Nil {
		prep.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// ChunkData representa a estrutura de cada chunk em SelectedChunks
type ChunkData struct {
	ArticleID    string  `json:"article_id"`
	ArticleTitle string  `json:"article_title"`
	ArticleYear  int     `json:"article_year"`
	Journal      string  `json:"journal"`
	ChunkIndex   int     `json:"chunk_index"`
	ChunkText    string  `json:"chunk_text"`
	Section      string  `json:"section"`       // results, discussion, methods, introduction, etc
	Similarity   float64 `json:"similarity"`    // 0.0-1.0
	WordCount    int     `json:"word_count"`
	PageRange    *string `json:"page_range,omitempty"`
}

// PreparationMetadata representa a estrutura de Metadata
type PreparationMetadata struct {
	TotalChunks          int                `json:"total_chunks"`
	ArticlesCount        int                `json:"articles_count"`
	AvgSimilarity        float64            `json:"avg_similarity"`
	SectionsDistribution map[string]int     `json:"sections_distribution"` // {results: 8, discussion: 6, ...}
}
