package dto

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/datatypes"
)

// PrepareEnrichmentRequest request para preparar chunks
type PrepareEnrichmentRequest struct {
	// Limite de chunks a buscar (default: 20, max: 100)
	// @example 20
	// @minimum 1
	// @maximum 100
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1,lte=100"`

	// Threshold mínimo de similaridade (default: 0.65, range: 0.0-1.0)
	// @example 0.65
	// @minimum 0
	// @maximum 1
	MinSimilarity *float64 `json:"minSimilarity,omitempty" validate:"omitempty,gte=0,lte=1"`
}

// PrepareEnrichmentResponse response da preparação
type PrepareEnrichmentResponse struct {
	// ID da preparação criada
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `json:"id"`

	// ID do score item
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ScoreItemID uuid.UUID `json:"scoreItemId"`

	// Nome do score item
	// @example Hemoglobina Glicada (HbA1c)
	ScoreItemName string `json:"scoreItemName"`

	// Chunks selecionados (JSONB array)
	// @example [{"article_id":"019c...", "article_title":"Effects of...", "section":"results", "similarity":0.89}]
	SelectedChunks datatypes.JSONMap `json:"selectedChunks"`

	// Metadata agregado
	// @example {"total_chunks":20, "articles_count":5, "avg_similarity":0.83}
	Metadata datatypes.JSONMap `json:"metadata"`

	// Status da preparação
	// @enum ready,processing,completed,failed
	// @example ready
	Status string `json:"status"`

	// Mensagem de erro (se houver)
	// @example Failed to fetch chunks: connection timeout
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// ToEnrichmentPreparationDTO converte model para DTO
func ToEnrichmentPreparationDTO(m *models.ScoreItemEnrichmentPreparation) *PrepareEnrichmentResponse {
	dto := &PrepareEnrichmentResponse{
		ID:             m.ID,
		ScoreItemID:    m.ScoreItemID,
		SelectedChunks: m.SelectedChunks,
		Metadata:       m.Metadata,
		Status:         m.Status,
		ErrorMessage:   m.ErrorMessage,
	}

	if m.ScoreItem != nil {
		dto.ScoreItemName = m.ScoreItem.Name
	}

	return dto
}

// EnrichmentPreparationStatsResponse estatísticas de preparações
type EnrichmentPreparationStatsResponse struct {
	// Total de preparações
	// @example 150
	Total int64 `json:"total"`

	// Preparações por status
	// @example {"ready":80, "processing":10, "completed":50, "failed":10}
	ByStatus map[string]int64 `json:"byStatus"`
}
