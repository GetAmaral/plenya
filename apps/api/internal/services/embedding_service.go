package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

// embeddingProvider define qual API de embedding usar
type embeddingProvider string

const (
	providerOpenAI  embeddingProvider = "openai"
	providerVoyage  embeddingProvider = "voyage"
	embeddingDims   int               = 1024
)

// EmbeddingService - serviço de integração com APIs de Embedding
// Suporta OpenAI (text-embedding-3-large) e Voyage AI (voyage-multilingual-2)
// Voyage AI é preferido quando VOYAGE_API_KEY está configurado (melhor suporte multilingual)
type EmbeddingService struct {
	provider     embeddingProvider
	openaiClient *openai.Client
	openaiModel  string
	voyageAPIKey string
	voyageModel  string
	voyageAPIURL string
	db           *gorm.DB
}

// NewEmbeddingService cria uma nova instância do serviço de embeddings
// Usa Voyage AI se VOYAGE_API_KEY estiver configurado, senão usa OpenAI
func NewEmbeddingService(cfg *config.Config, db *gorm.DB) *EmbeddingService {
	svc := &EmbeddingService{db: db}

	if cfg.VoyageAI.APIKey != "" {
		// Voyage AI: melhor suporte cross-lingual PT/EN
		svc.provider = providerVoyage
		svc.voyageAPIKey = cfg.VoyageAI.APIKey
		svc.voyageModel = cfg.VoyageAI.Model
		svc.voyageAPIURL = cfg.VoyageAI.APIURL
		fmt.Printf("🌍 Embedding provider: Voyage AI (%s) — multilingual mode\n", cfg.VoyageAI.Model)
	} else {
		// OpenAI: fallback
		svc.provider = providerOpenAI
		svc.openaiModel = cfg.OpenAI.EmbeddingModel

		clientConfig := openai.DefaultConfig(cfg.OpenAI.APIKey)
		if cfg.OpenAI.APIURL != "" && cfg.OpenAI.APIURL != "https://api.openai.com/v1" {
			clientConfig.BaseURL = cfg.OpenAI.APIURL
		}
		svc.openaiClient = openai.NewClientWithConfig(clientConfig)
		fmt.Printf("🤖 Embedding provider: OpenAI (%s)\n", cfg.OpenAI.EmbeddingModel)
	}

	return svc
}

// GetProvider retorna o provider ativo ("openai" ou "voyage")
func (s *EmbeddingService) GetProvider() string {
	return string(s.provider)
}

// GenerateEmbedding gera vetor de embedding para um texto
// Retorna []float32 com 1024 dimensões
func (s *EmbeddingService) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	if text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	// Truncar texto se muito longo (max ~32k chars)
	if len(text) > 32000 {
		text = text[:32000]
	}

	if s.provider == providerVoyage {
		return s.voyageGenerateEmbedding(ctx, text)
	}
	return s.openaiGenerateEmbedding(ctx, text)
}

// BatchGenerateEmbeddings gera embeddings para múltiplos textos em batch
// Retorna slice de embeddings na mesma ordem dos inputs
func (s *EmbeddingService) BatchGenerateEmbeddings(ctx context.Context, texts []string) ([][]float32, error) {
	if len(texts) == 0 {
		return nil, fmt.Errorf("texts cannot be empty")
	}

	maxBatch := 100
	if s.provider == providerVoyage {
		maxBatch = 128 // Voyage suporta até 128
	}

	if len(texts) > maxBatch {
		return nil, fmt.Errorf("too many texts: got %d, max %d per batch", len(texts), maxBatch)
	}

	// Truncar textos longos
	truncated := make([]string, len(texts))
	for i, t := range texts {
		if len(t) > 32000 {
			truncated[i] = t[:32000]
		} else {
			truncated[i] = t
		}
	}

	if s.provider == providerVoyage {
		return s.voyageBatchEmbeddings(ctx, truncated)
	}
	return s.openaiBatchEmbeddings(ctx, truncated)
}

// ─── OpenAI ──────────────────────────────────────────────────────────────────

func (s *EmbeddingService) openaiGenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	req := openai.EmbeddingRequest{
		Input:      []string{text},
		Model:      openai.EmbeddingModel(s.openaiModel),
		Dimensions: embeddingDims,
	}

	start := time.Now()
	resp, err := s.openaiClient.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("openai embedding failed: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned from OpenAI")
	}

	emb := resp.Data[0].Embedding
	if len(emb) != embeddingDims {
		return nil, fmt.Errorf("unexpected dims: got %d, expected %d", len(emb), embeddingDims)
	}

	fmt.Printf("💰 OpenAI Embedding - Model: %s, Tokens: %d, Duration: %v\n",
		resp.Model, resp.Usage.TotalTokens, time.Since(start))

	go s.logAPIUsage("openai", string(resp.Model), resp.Usage.TotalTokens)
	return emb, nil
}

func (s *EmbeddingService) openaiBatchEmbeddings(ctx context.Context, texts []string) ([][]float32, error) {
	req := openai.EmbeddingRequest{
		Input:      texts,
		Model:      openai.EmbeddingModel(s.openaiModel),
		Dimensions: embeddingDims,
	}

	start := time.Now()
	resp, err := s.openaiClient.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("openai batch embedding failed: %w", err)
	}

	if len(resp.Data) != len(texts) {
		return nil, fmt.Errorf("unexpected count: got %d, expected %d", len(resp.Data), len(texts))
	}

	embeddings := make([][]float32, len(resp.Data))
	for i, d := range resp.Data {
		if d.Index != i {
			return nil, fmt.Errorf("index mismatch: got %d, expected %d", d.Index, i)
		}
		if len(d.Embedding) != embeddingDims {
			return nil, fmt.Errorf("unexpected dims at index %d: got %d", i, len(d.Embedding))
		}
		embeddings[i] = d.Embedding
	}

	fmt.Printf("💰 OpenAI Batch - Model: %s, Texts: %d, Tokens: %d, Duration: %v\n",
		resp.Model, len(texts), resp.Usage.TotalTokens, time.Since(start))

	go s.logAPIUsage("openai", string(resp.Model), resp.Usage.TotalTokens)
	return embeddings, nil
}

// ─── Voyage AI ───────────────────────────────────────────────────────────────

type voyageRequest struct {
	Input []string `json:"input"`
	Model string   `json:"model"`
}

type voyageResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

func (s *EmbeddingService) voyageGenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	embeddings, err := s.voyageBatchEmbeddings(ctx, []string{text})
	if err != nil {
		return nil, err
	}
	return embeddings[0], nil
}

func (s *EmbeddingService) voyageBatchEmbeddings(ctx context.Context, texts []string) ([][]float32, error) {
	body, err := json.Marshal(voyageRequest{
		Input: texts,
		Model: s.voyageModel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal voyage request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST",
		s.voyageAPIURL+"/embeddings", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create voyage request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+s.voyageAPIKey)

	start := time.Now()
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("voyage API call failed: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read voyage response: %w", err)
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("voyage API error %d: %s", httpResp.StatusCode, string(respBody))
	}

	var voyageResp voyageResponse
	if err := json.Unmarshal(respBody, &voyageResp); err != nil {
		return nil, fmt.Errorf("failed to parse voyage response: %w", err)
	}

	if len(voyageResp.Data) != len(texts) {
		return nil, fmt.Errorf("voyage count mismatch: got %d, expected %d", len(voyageResp.Data), len(texts))
	}

	embeddings := make([][]float32, len(voyageResp.Data))
	for _, d := range voyageResp.Data {
		if d.Index < 0 || d.Index >= len(texts) {
			return nil, fmt.Errorf("voyage invalid index: %d", d.Index)
		}
		if len(d.Embedding) != embeddingDims {
			return nil, fmt.Errorf("voyage unexpected dims at index %d: got %d, expected %d",
				d.Index, len(d.Embedding), embeddingDims)
		}
		embeddings[d.Index] = d.Embedding
	}

	fmt.Printf("🌍 Voyage AI Batch - Model: %s, Texts: %d, Tokens: %d, Duration: %v\n",
		voyageResp.Model, len(texts), voyageResp.Usage.TotalTokens, time.Since(start))

	go s.logAPIUsage("voyage", voyageResp.Model, voyageResp.Usage.TotalTokens)
	return embeddings, nil
}

// ─── Shared ───────────────────────────────────────────────────────────────────

func (s *EmbeddingService) logAPIUsage(provider, model string, totalTokens int) {
	usageLog := map[string]interface{}{
		"provider":      provider,
		"model":         model,
		"endpoint":      "embeddings",
		"input_tokens":  totalTokens,
		"output_tokens": nil,
		"total_tokens":  totalTokens,
		"user_id":       nil,
		"metadata": map[string]interface{}{
			"operation": "embedding_generation",
		},
	}
	if err := s.db.Table("api_usage_logs").Create(usageLog).Error; err != nil {
		fmt.Printf("⚠️  Failed to log API usage: %v\n", err)
	}
}

// GetEstimatedCost calcula custo estimado em USD
func (s *EmbeddingService) GetEstimatedCost(text string) float64 {
	estimatedTokens := len(text) / 4
	var costPerToken float64
	if s.provider == providerVoyage {
		costPerToken = 0.10 / 1_000_000.0 // Voyage: $0.10/1M tokens
	} else {
		costPerToken = 0.13 / 1_000_000.0 // OpenAI: $0.13/1M tokens
	}
	return float64(estimatedTokens) * costPerToken
}

// HealthCheck verifica se o provider está acessível
func (s *EmbeddingService) HealthCheck(ctx context.Context) error {
	_, err := s.GenerateEmbedding(ctx, "health check")
	if err != nil {
		return fmt.Errorf("%s health check failed: %w", s.provider, err)
	}
	return nil
}
