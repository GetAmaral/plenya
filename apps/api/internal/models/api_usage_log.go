package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// APIUsageLog rastreia uso de APIs externas (OpenAI, Anthropic)
// Usado para monitorar custos e debugging
type APIUsageLog struct {
	// Primary Key
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Provedor da API
	// @enum openai,anthropic
	Provider string `gorm:"type:varchar(50);not null;check:provider IN ('openai','anthropic')" json:"provider" validate:"required,oneof=openai anthropic"`

	// @example text-embedding-3-large
	// @maxLength 100
	Model string `gorm:"type:varchar(100);not null" json:"model" validate:"required,max=100"`

	// @example embeddings
	// @maxLength 100
	Endpoint string `gorm:"type:varchar(100);not null" json:"endpoint" validate:"required,max=100"`

	// Uso de tokens
	// @minimum 0
	InputTokens  *int `gorm:"type:integer" json:"inputTokens,omitempty" validate:"omitempty,gte=0"`
	OutputTokens *int `gorm:"type:integer" json:"outputTokens,omitempty" validate:"omitempty,gte=0"`
	TotalTokens  int  `gorm:"type:integer;not null" json:"totalTokens" validate:"gte=0"`

	// Custo em USD (calculado automaticamente)
	// @example 0.000026
	CostUSD *float64 `gorm:"type:decimal(10,6)" json:"costUsd,omitempty"`

	// User que fez a requisição
	UserID *uuid.UUID `gorm:"type:uuid" json:"userId,omitempty"`

	// Metadata adicional: {entity_type: "article", entity_id: "...", operation: "chunk_embedding"}
	Metadata map[string]interface{} `gorm:"type:jsonb;default:'{}'" json:"metadata,omitempty"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL" json:"user,omitempty"`

	// Timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

// TableName especifica o nome da tabela no banco de dados
func (APIUsageLog) TableName() string {
	return "api_usage_logs"
}

// BeforeCreate hook - executado antes de criar registro
func (aul *APIUsageLog) BeforeCreate(tx *gorm.DB) error {
	if aul.ID == uuid.Nil {
		aul.ID = uuid.Must(uuid.NewV7())
	}

	// Calcular custo automaticamente se não fornecido
	if aul.CostUSD == nil {
		cost := aul.CalculateCost()
		aul.CostUSD = &cost
	}

	return nil
}

// CalculateCost calcula custo em USD baseado no provider e modelo
// Preços de fevereiro 2026
func (aul *APIUsageLog) CalculateCost() float64 {
	if aul.TotalTokens == 0 {
		return 0.0
	}

	tokensInMillions := float64(aul.TotalTokens) / 1_000_000.0

	switch aul.Provider {
	case "openai":
		switch aul.Model {
		case "text-embedding-3-large":
			return tokensInMillions * 0.13 // $0.13 / 1M tokens
		case "text-embedding-3-small":
			return tokensInMillions * 0.02 // $0.02 / 1M tokens
		default:
			return 0.0
		}

	case "anthropic":
		switch aul.Model {
		case "claude-3-5-haiku-20241022":
			// Haiku: $0.80 / 1M input tokens, $4.00 / 1M output tokens
			inputCost := 0.0
			outputCost := 0.0
			if aul.InputTokens != nil {
				inputCost = (float64(*aul.InputTokens) / 1_000_000.0) * 0.80
			}
			if aul.OutputTokens != nil {
				outputCost = (float64(*aul.OutputTokens) / 1_000_000.0) * 4.00
			}
			return inputCost + outputCost
		default:
			return 0.0
		}

	default:
		return 0.0
	}
}
