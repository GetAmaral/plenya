package services

import (
	"context"
	"fmt"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

// EmbeddingService - serviço de integração com OpenAI Embeddings API
// Gera vetores de embedding para busca semântica (RAG)
type EmbeddingService struct {
	client *openai.Client
	model  string
	apiURL string
	db     *gorm.DB
}

// NewEmbeddingService cria uma nova instância do serviço de embeddings
func NewEmbeddingService(cfg *config.Config, db *gorm.DB) *EmbeddingService {
	// Configurar cliente OpenAI
	clientConfig := openai.DefaultConfig(cfg.OpenAI.APIKey)

	// Se API URL customizada fornecida, usar
	if cfg.OpenAI.APIURL != "" && cfg.OpenAI.APIURL != "https://api.openai.com/v1" {
		clientConfig.BaseURL = cfg.OpenAI.APIURL
	}

	client := openai.NewClientWithConfig(clientConfig)

	return &EmbeddingService{
		client: client,
		model:  cfg.OpenAI.EmbeddingModel,
		apiURL: cfg.OpenAI.APIURL,
		db:     db,
	}
}

// GenerateEmbedding gera vetor de embedding para um texto
// Retorna []float32 com 1024 dimensões (text-embedding-3-large)
func (s *EmbeddingService) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	// Validar entrada
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	// Truncar texto se muito longo (max 8191 tokens ~= 32k chars)
	maxChars := 32000
	if len(text) > maxChars {
		text = text[:maxChars]
		fmt.Printf("⚠️  Text truncated to %d chars for embedding\n", maxChars)
	}

	// Preparar request
	req := openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.EmbeddingModel(s.model),
		// Para text-embedding-3-large, especificar dimensões (default 3072, usamos 1024)
		Dimensions: 1024,
	}

	// Chamar API OpenAI
	startTime := time.Now()
	resp, err := s.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}
	duration := time.Since(startTime)

	// Extrair embedding
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned from API")
	}

	embedding := resp.Data[0].Embedding

	// Validar dimensões
	if len(embedding) != 1024 {
		return nil, fmt.Errorf("unexpected embedding dimensions: got %d, expected 1024", len(embedding))
	}

	// Log token usage para tracking de custos
	fmt.Printf("💰 OpenAI Embedding - Model: %s, Tokens: %d, Duration: %v\n",
		resp.Model, resp.Usage.TotalTokens, duration)

	// Registrar uso em api_usage_logs (assíncrono)
	go s.logAPIUsage(string(resp.Model), resp.Usage.TotalTokens, nil)

	return embedding, nil
}

// BatchGenerateEmbeddings gera embeddings para múltiplos textos em batch
// Limite OpenAI: 100 textos por request (recomendado: 10-20 para melhor performance)
// Retorna slice de embeddings na mesma ordem dos inputs
func (s *EmbeddingService) BatchGenerateEmbeddings(ctx context.Context, texts []string) ([][]float32, error) {
	// Validar entrada
	if len(texts) == 0 {
		return nil, fmt.Errorf("texts cannot be empty")
	}

	// Limite OpenAI: 100 inputs por request
	batchSize := 100
	if len(texts) > batchSize {
		return nil, fmt.Errorf("too many texts: got %d, max %d per batch", len(texts), batchSize)
	}

	// Truncar textos longos
	maxChars := 32000
	truncatedTexts := make([]string, len(texts))
	for i, text := range texts {
		if len(text) > maxChars {
			truncatedTexts[i] = text[:maxChars]
		} else {
			truncatedTexts[i] = text
		}
	}

	// Preparar request
	req := openai.EmbeddingRequest{
		Input:      truncatedTexts,
		Model:      openai.EmbeddingModel(s.model),
		Dimensions: 1024,
	}

	// Chamar API OpenAI
	startTime := time.Now()
	resp, err := s.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create embeddings: %w", err)
	}
	duration := time.Since(startTime)

	// Extrair embeddings
	if len(resp.Data) != len(texts) {
		return nil, fmt.Errorf("unexpected number of embeddings: got %d, expected %d", len(resp.Data), len(texts))
	}

	embeddings := make([][]float32, len(resp.Data))
	for i, data := range resp.Data {
		// Validar ordem (OpenAI retorna na mesma ordem do input)
		if data.Index != i {
			return nil, fmt.Errorf("embedding index mismatch: got %d, expected %d", data.Index, i)
		}

		// Validar dimensões
		if len(data.Embedding) != 1024 {
			return nil, fmt.Errorf("unexpected embedding dimensions at index %d: got %d, expected 1024", i, len(data.Embedding))
		}

		embeddings[i] = data.Embedding
	}

	// Log token usage
	fmt.Printf("💰 OpenAI Batch Embeddings - Model: %s, Texts: %d, Tokens: %d, Duration: %v\n",
		resp.Model, len(texts), resp.Usage.TotalTokens, duration)

	// Registrar uso em api_usage_logs (assíncrono)
	go s.logAPIUsage(string(resp.Model), resp.Usage.TotalTokens, nil)

	return embeddings, nil
}

// logAPIUsage registra uso da API OpenAI para tracking de custos
// Executado de forma assíncrona (não bloqueia request principal)
func (s *EmbeddingService) logAPIUsage(model string, totalTokens int, userID *string) {
	// Criar registro de uso
	usageLog := map[string]interface{}{
		"provider":      "openai",
		"model":         model,
		"endpoint":      "embeddings",
		"input_tokens":  totalTokens, // Embeddings só têm input tokens
		"output_tokens": nil,
		"total_tokens":  totalTokens,
		"user_id":       userID,
		"metadata": map[string]interface{}{
			"operation": "embedding_generation",
		},
	}

	// Inserir no banco (ignora erro para não quebrar request principal)
	if err := s.db.Table("api_usage_logs").Create(usageLog).Error; err != nil {
		fmt.Printf("⚠️  Failed to log API usage: %v\n", err)
	}
}

// GetEstimatedCost calcula custo estimado em USD para um texto
// Baseado em pricing de fevereiro 2026: $0.13 / 1M tokens
func (s *EmbeddingService) GetEstimatedCost(text string) float64 {
	// Estimativa: ~4 chars = 1 token (média para português)
	estimatedTokens := len(text) / 4

	// Pricing: $0.13 / 1M tokens para text-embedding-3-large
	costPerToken := 0.13 / 1_000_000.0

	return float64(estimatedTokens) * costPerToken
}

// HealthCheck verifica se a API OpenAI está acessível
func (s *EmbeddingService) HealthCheck(ctx context.Context) error {
	// Tentar gerar embedding de teste
	testText := "health check"
	_, err := s.GenerateEmbedding(ctx, testText)
	if err != nil {
		return fmt.Errorf("openai api health check failed: %w", err)
	}
	return nil
}
