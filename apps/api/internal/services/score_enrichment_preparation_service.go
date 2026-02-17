package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ScoreEnrichmentPreparationService prepara chunks para enrichment (ETAPA 1 - SEM IA)
type ScoreEnrichmentPreparationService struct {
	db         *gorm.DB
	vectorRepo *repository.ArticleVectorRepository
	prepRepo   *repository.ScoreEnrichmentPreparationRepository
}

// NewScoreEnrichmentPreparationService cria nova instância
func NewScoreEnrichmentPreparationService(db *gorm.DB) *ScoreEnrichmentPreparationService {
	return &ScoreEnrichmentPreparationService{
		db:         db,
		vectorRepo: repository.NewArticleVectorRepository(db),
		prepRepo:   repository.NewScoreEnrichmentPreparationRepository(db),
	}
}

// PrepareChunksAdaptive busca chunks com threshold adaptativo (fallback progressivo)
// Garante pelo menos 15 chunks ou usa o threshold mais baixo possível
func (s *ScoreEnrichmentPreparationService) PrepareChunksAdaptive(
	scoreItemID uuid.UUID,
) (*models.ScoreItemEnrichmentPreparation, error) {
	targetChunks := 20  // Alvo mínimo ideal
	maxLimit := 50      // Buscar até 50 chunks

	// Tentativas progressivas de threshold (do mais rigoroso ao mais permissivo)
	thresholds := []struct {
		value float64
		label string
	}{
		{0.35, "normal"},
		{0.30, "permissivo"},
		{0.25, "muito permissivo"},
		{0.20, "extremo"},
		{0.15, "fallback final"},
	}

	var bestChunks []repository.ChunkSearchResult
	var usedThreshold float64
	var thresholdLabel string

	// Validar que score_item existe
	var scoreItem models.ScoreItem
	err := s.db.Preload("Subgroup.Group").First(&scoreItem, "id = ?", scoreItemID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("score_item not found: %s", scoreItemID)
		}
		return nil, fmt.Errorf("failed to fetch score_item: %w", err)
	}

	// Tentar cada threshold até encontrar >= targetChunks
	for _, t := range thresholds {
		chunks, err := s.vectorRepo.FindTopChunksForScoreItem(scoreItemID, maxLimit, t.value)
		if err != nil {
			continue
		}

		if len(chunks) >= targetChunks {
			bestChunks = chunks
			usedThreshold = t.value
			thresholdLabel = t.label
			break
		}

		// Guardar o melhor resultado mesmo se < targetChunks
		if len(chunks) > len(bestChunks) {
			bestChunks = chunks
			usedThreshold = t.value
			thresholdLabel = t.label
		}
	}

	if len(bestChunks) == 0 {
		return nil, fmt.Errorf("no chunks found even with threshold 0.15")
	}

	// Limitar chunks ao máximo desejado (30 para qualidade)
	finalLimit := 30
	if usedThreshold >= 0.30 {
		finalLimit = 30 // Threshold bom, 30 chunks suficientes
	} else if usedThreshold >= 0.20 {
		finalLimit = 35 // Threshold mais baixo, incluir mais chunks
	} else {
		finalLimit = 40 // Fallback extremo, incluir ainda mais
	}

	if len(bestChunks) > finalLimit {
		bestChunks = bestChunks[:finalLimit]
	}

	// Buscar ScoreItem com relações para fullName
	s.db.Preload("Subgroup.Group").Preload("ParentItem").First(&scoreItem, scoreItemID)
	fullName := scoreItem.GetFullName()

	return s.savePreparationWithPrompts(scoreItemID, &scoreItem, fullName, bestChunks, usedThreshold, thresholdLabel)
}

// PrepareChunks busca chunks relevantes via RAG e salva no banco (SEM IA)
// Retorna preparation criado ou erro
// DEPRECATED: Use PrepareChunksAdaptive para melhor qualidade
func (s *ScoreEnrichmentPreparationService) PrepareChunks(
	scoreItemID uuid.UUID,
	limit int,              // Total de chunks (default: 20)
	minSimilarity float64,  // Threshold (default: 0.65)
) (*models.ScoreItemEnrichmentPreparation, error) {
	// Validar que score_item existe
	var scoreItem models.ScoreItem
	err := s.db.Preload("Subgroup.Group").First(&scoreItem, "id = ?", scoreItemID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("score_item not found: %s", scoreItemID)
		}
		return nil, fmt.Errorf("failed to fetch score_item: %w", err)
	}

	// Defaults
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.65
	}

	// 1. Buscar chunks via RAG
	chunks, err := s.vectorRepo.FindTopChunksForScoreItem(scoreItemID, limit, minSimilarity)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch chunks: %w", err)
	}

	if len(chunks) == 0 {
		return nil, fmt.Errorf("no chunks found with similarity >= %.2f", minSimilarity)
	}

	// Delegar para savePreparation com threshold usado
	return s.savePreparation(scoreItemID, chunks, minSimilarity, "manual")
}

// savePreparation extrai lógica comum de salvar preparation (com metadata enriquecido)
func (s *ScoreEnrichmentPreparationService) savePreparation(
	scoreItemID uuid.UUID,
	chunks []repository.ChunkSearchResult,
	usedThreshold float64,
	thresholdLabel string,
) (*models.ScoreItemEnrichmentPreparation, error) {
	// 1. Converter chunks para JSON structure
	chunksData := make([]map[string]interface{}, len(chunks))
	articleIDs := make(map[uuid.UUID]bool)
	sectionsCount := make(map[string]int)
	yearStats := make([]int, 0, len(chunks))

	var totalSimilarity float64
	var totalWordCount int
	minSimilarity := 1.0
	maxSimilarity := 0.0

	for i, chunk := range chunks {
		chunksData[i] = map[string]interface{}{
			"article_id":    chunk.ArticleID.String(),
			"article_title": chunk.ArticleTitle,
			"article_year":  chunk.ArticleYear,
			"journal":       chunk.Journal,
			"chunk_index":   chunk.ChunkIndex,
			"chunk_text":    chunk.ChunkText,
			"section":       chunk.Section,
			"similarity":    chunk.Similarity,
			"word_count":    chunk.WordCount,
		}

		if chunk.PageRange != nil {
			chunksData[i]["page_range"] = *chunk.PageRange
		}

		articleIDs[chunk.ArticleID] = true
		sectionsCount[chunk.Section]++
		totalSimilarity += chunk.Similarity
		totalWordCount += chunk.WordCount

		// Track min/max similarity
		if chunk.Similarity < minSimilarity {
			minSimilarity = chunk.Similarity
		}
		if chunk.Similarity > maxSimilarity {
			maxSimilarity = chunk.Similarity
		}

		// Track years for recency stats
		if chunk.ArticleYear > 0 {
			yearStats = append(yearStats, chunk.ArticleYear)
		}
	}

	// 2. Calcular metadata enriquecido
	avgSimilarity := totalSimilarity / float64(len(chunks))
	avgChunkLength := 0
	if len(chunks) > 0 {
		avgChunkLength = totalWordCount / len(chunks)
	}

	// Recency stats
	newestYear := 0
	oldestYear := 9999
	var totalYear int
	for _, year := range yearStats {
		if year > newestYear {
			newestYear = year
		}
		if year < oldestYear {
			oldestYear = year
		}
		totalYear += year
	}
	avgYear := 0
	if len(yearStats) > 0 {
		avgYear = totalYear / len(yearStats)
	}

	// Quality grade
	qualityGrade := "poor"
	if len(chunks) >= 25 && avgSimilarity >= 0.6 {
		qualityGrade = "excellent"
	} else if len(chunks) >= 20 && avgSimilarity >= 0.5 {
		qualityGrade = "good"
	} else if len(chunks) >= 15 && avgSimilarity >= 0.4 {
		qualityGrade = "fair"
	}

	// 3. Criar preparation com metadata enriquecido
	chunksInterface := make([]interface{}, len(chunksData))
	for i, chunk := range chunksData {
		chunksInterface[i] = chunk
	}

	prep := &models.ScoreItemEnrichmentPreparation{
		ID:          uuid.Must(uuid.NewV7()),
		ScoreItemID: scoreItemID,
		SelectedChunks: datatypes.JSONMap{
			"items": chunksInterface,
		},
		Metadata: datatypes.JSONMap{
			"total_chunks":          len(chunks),
			"articles_count":        len(articleIDs),
			"avg_similarity":        avgSimilarity,
			"min_similarity":        minSimilarity,
			"max_similarity":        maxSimilarity,
			"sections_distribution": sectionsCount,
			"total_word_count":      totalWordCount,
			"avg_chunk_length":      avgChunkLength,
			"recency_stats": map[string]interface{}{
				"newest_year": newestYear,
				"oldest_year": oldestYear,
				"avg_year":    avgYear,
			},
			"quality_grade":  qualityGrade,
			"threshold_used": usedThreshold,
			"threshold_type": thresholdLabel,
		},
		Status: "ready",
	}

	// 4. Salvar no banco (upsert)
	err := s.prepRepo.Create(prep)
	if err != nil {
		return nil, fmt.Errorf("failed to save preparation: %w", err)
	}

	// 5. Recarregar com ScoreItem preloaded
	prep, err = s.prepRepo.FindByScoreItemID(scoreItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to reload preparation: %w", err)
	}

	return prep, nil
}

// GetPreparation busca preparation por score_item_id
func (s *ScoreEnrichmentPreparationService) GetPreparation(scoreItemID uuid.UUID) (*models.ScoreItemEnrichmentPreparation, error) {
	prep, err := s.prepRepo.FindByScoreItemID(scoreItemID)
	if err != nil {
		return nil, fmt.Errorf("failed to get preparation: %w", err)
	}

	if prep == nil {
		return nil, fmt.Errorf("preparation not found for score_item: %s", scoreItemID)
	}

	return prep, nil
}

// DeletePreparation remove preparation por score_item_id
func (s *ScoreEnrichmentPreparationService) DeletePreparation(scoreItemID uuid.UUID) error {
	return s.prepRepo.DeleteByScoreItemID(scoreItemID)
}

// GetStats retorna estatísticas de preparações
func (s *ScoreEnrichmentPreparationService) GetStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Total
	total, err := s.prepRepo.Count()
	if err != nil {
		return nil, err
	}
	stats["total"] = total

	// Por status
	ready, _ := s.prepRepo.CountByStatus("ready")
	processing, _ := s.prepRepo.CountByStatus("processing")
	completed, _ := s.prepRepo.CountByStatus("completed")
	failed, _ := s.prepRepo.CountByStatus("failed")

	stats["by_status"] = map[string]int64{
		"ready":      ready,
		"processing": processing,
		"completed":  completed,
		"failed":     failed,
	}

	return stats, nil
}

// savePreparationWithPrompts salva preparation com prompts prontos para Claude
func (s *ScoreEnrichmentPreparationService) savePreparationWithPrompts(
	scoreItemID uuid.UUID,
	scoreItem *models.ScoreItem,
	fullName string,
	chunks []repository.ChunkSearchResult,
	usedThreshold float64,
	thresholdLabel string,
) (*models.ScoreItemEnrichmentPreparation, error) {
	if len(chunks) == 0 {
		return nil, fmt.Errorf("no chunks to process")
	}

	// 1. Processar chunks e calcular metadata
	chunksData := make([]map[string]interface{}, len(chunks))
	articleIDs := make(map[uuid.UUID]bool)
	sectionsCount := make(map[string]int)
	yearStats := make([]int, 0)

	var totalSimilarity, totalWordCount float64
	minSim, maxSim := 1.0, 0.0

	for i, chunk := range chunks {
		chunksData[i] = map[string]interface{}{
			"article_id":    chunk.ArticleID.String(),
			"article_title": chunk.ArticleTitle,
			"article_year":  chunk.ArticleYear,
			"journal":       chunk.Journal,
			"chunk_index":   chunk.ChunkIndex,
			"chunk_text":    chunk.ChunkText,
			"section":       chunk.Section,
			"similarity":    chunk.Similarity,
			"word_count":    chunk.WordCount,
		}

		if chunk.PageRange != nil {
			chunksData[i]["page_range"] = *chunk.PageRange
		}

		articleIDs[chunk.ArticleID] = true
		sectionsCount[chunk.Section]++
		totalSimilarity += chunk.Similarity
		totalWordCount += float64(chunk.WordCount)

		if chunk.Similarity < minSim {
			minSim = chunk.Similarity
		}
		if chunk.Similarity > maxSim {
			maxSim = chunk.Similarity
		}
		if chunk.ArticleYear > 0 {
			yearStats = append(yearStats, chunk.ArticleYear)
		}
	}

	avgSim := totalSimilarity / float64(len(chunks))
	avgLen := int(totalWordCount / float64(len(chunks)))

	// Recency stats
	newestYear, oldestYear, avgYear := 0, 9999, 0
	if len(yearStats) > 0 {
		total := 0
		for _, y := range yearStats {
			if y > newestYear {
				newestYear = y
			}
			if y < oldestYear {
				oldestYear = y
			}
			total += y
		}
		avgYear = total / len(yearStats)
	}

	// Quality grade
	grade := "poor"
	if len(chunks) >= 25 && avgSim >= 0.6 {
		grade = "excellent"
	} else if len(chunks) >= 20 && avgSim >= 0.5 {
		grade = "good"
	} else if len(chunks) >= 15 && avgSim >= 0.4 {
		grade = "fair"
	}

	metadata := datatypes.JSONMap{
		"total_chunks":          len(chunks),
		"articles_count":        len(articleIDs),
		"avg_similarity":        avgSim,
		"min_similarity":        minSim,
		"max_similarity":        maxSim,
		"sections_distribution": sectionsCount,
		"total_word_count":      int(totalWordCount),
		"avg_chunk_length":      avgLen,
		"recency_stats": map[string]interface{}{
			"newest_year": newestYear,
			"oldest_year": oldestYear,
			"avg_year":    avgYear,
		},
		"quality_grade":  grade,
		"threshold_used": usedThreshold,
		"threshold_type": thresholdLabel,
	}

	// 2. Gerar prompts
	prompts := s.buildPrompts(scoreItem, fullName, chunks, len(articleIDs), avgSim)

	// DEBUG
	fmt.Printf("DEBUG: Generated prompts - CR: %d chars, PE: %d chars\n",
		len(prompts.ClinicalRelevance), len(prompts.PatientExplanation))

	// 3. Criar preparation
	chunksInterface := make([]interface{}, len(chunksData))
	for i, chunk := range chunksData {
		chunksInterface[i] = chunk
	}

	prep := &models.ScoreItemEnrichmentPreparation{
		ID:          uuid.Must(uuid.NewV7()),
		ScoreItemID: scoreItemID,
		SelectedChunks: datatypes.JSONMap{
			"items": chunksInterface,
		},
		Metadata:                 metadata,
		PromptClinicalRelevance:  &prompts.ClinicalRelevance,
		PromptPatientExplanation: &prompts.PatientExplanation,
		PromptConduct:            &prompts.Conduct,
		PromptMaxPoints:          &prompts.MaxPoints,
		Status:                   "ready",
	}

	// DEBUG
	fmt.Printf("DEBUG: prep.PromptClinicalRelevance is nil? %v\n", prep.PromptClinicalRelevance == nil)
	if prep.PromptClinicalRelevance != nil {
		fmt.Printf("DEBUG: Length before save: %d\n", len(*prep.PromptClinicalRelevance))
	}

	// 4. Salvar usando GORM (que deveria incluir todos os campos)
	err := s.db.Create(prep).Error
	if err != nil {
		return nil, fmt.Errorf("failed to save preparation: %w", err)
	}

	fmt.Printf("DEBUG: Saved to DB, ID: %s\n", prep.ID)

	// 5. Recarregar para confirmar
	var reloaded models.ScoreItemEnrichmentPreparation
	err = s.db.Preload("ScoreItem.Subgroup.Group").
		Preload("ScoreItem.ParentItem").
		First(&reloaded, "score_item_id = ?", scoreItemID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to reload: %w", err)
	}

	fmt.Printf("DEBUG: After reload - CR is nil? %v\n", reloaded.PromptClinicalRelevance == nil)
	if reloaded.PromptClinicalRelevance != nil {
		fmt.Printf("DEBUG: Length after reload: %d\n", len(*reloaded.PromptClinicalRelevance))
	}

	return &reloaded, nil
}

// processChunks processa chunks e retorna chunksData + metadata
func (s *ScoreEnrichmentPreparationService) processChunks(
	chunks []repository.ChunkSearchResult,
	usedThreshold float64,
	thresholdLabel string,
) ([]map[string]interface{}, datatypes.JSONMap) {
	chunksData := make([]map[string]interface{}, len(chunks))
	articleIDs := make(map[uuid.UUID]bool)
	sectionsCount := make(map[string]int)
	yearStats := make([]int, 0)

	var totalSimilarity, totalWordCount float64
	minSim, maxSim := 1.0, 0.0

	for i, chunk := range chunks {
		chunksData[i] = map[string]interface{}{
			"article_id":    chunk.ArticleID.String(),
			"article_title": chunk.ArticleTitle,
			"article_year":  chunk.ArticleYear,
			"journal":       chunk.Journal,
			"chunk_index":   chunk.ChunkIndex,
			"chunk_text":    chunk.ChunkText,
			"section":       chunk.Section,
			"similarity":    chunk.Similarity,
			"word_count":    chunk.WordCount,
		}

		if chunk.PageRange != nil {
			chunksData[i]["page_range"] = *chunk.PageRange
		}

		articleIDs[chunk.ArticleID] = true
		sectionsCount[chunk.Section]++
		totalSimilarity += chunk.Similarity
		totalWordCount += float64(chunk.WordCount)

		if chunk.Similarity < minSim {
			minSim = chunk.Similarity
		}
		if chunk.Similarity > maxSim {
			maxSim = chunk.Similarity
		}
		if chunk.ArticleYear > 0 {
			yearStats = append(yearStats, chunk.ArticleYear)
		}
	}

	avgSim := totalSimilarity / float64(len(chunks))
	avgLen := int(totalWordCount / float64(len(chunks)))

	// Recency stats
	newestYear, oldestYear, avgYear := 0, 9999, 0
	if len(yearStats) > 0 {
		total := 0
		for _, y := range yearStats {
			if y > newestYear {
				newestYear = y
			}
			if y < oldestYear {
				oldestYear = y
			}
			total += y
		}
		avgYear = total / len(yearStats)
	}

	// Quality grade
	grade := "poor"
	if len(chunks) >= 25 && avgSim >= 0.6 {
		grade = "excellent"
	} else if len(chunks) >= 20 && avgSim >= 0.5 {
		grade = "good"
	} else if len(chunks) >= 15 && avgSim >= 0.4 {
		grade = "fair"
	}

	metadata := datatypes.JSONMap{
		"total_chunks":          len(chunks),
		"articles_count":        len(articleIDs),
		"avg_similarity":        avgSim,
		"min_similarity":        minSim,
		"max_similarity":        maxSim,
		"sections_distribution": sectionsCount,
		"total_word_count":      int(totalWordCount),
		"avg_chunk_length":      avgLen,
		"recency_stats": map[string]interface{}{
			"newest_year": newestYear,
			"oldest_year": oldestYear,
			"avg_year":    avgYear,
		},
		"quality_grade":  grade,
		"threshold_used": usedThreshold,
		"threshold_type": thresholdLabel,
	}

	return chunksData, metadata
}

// EnrichmentPrompts contém os 4 prompts gerados
type EnrichmentPrompts struct {
	ClinicalRelevance  string
	PatientExplanation string
	Conduct            string
	MaxPoints          string
}

// generateEnrichmentPrompts gera os 4 prompts prontos para Claude
func (s *ScoreEnrichmentPreparationService) generateEnrichmentPrompts(
	scoreItem *models.ScoreItem,
	fullName string,
	chunks []repository.ChunkSearchResult,
	metadata datatypes.JSONMap,
) EnrichmentPrompts {
	totalChunks := len(chunks)
	articlesCount := len(extractUniqueArticles(chunks))

	// Formatar chunks para inclusão no prompt
	chunksText := formatChunksForPrompt(chunks)

	// Contexto do ScoreItem
	siContext := fmt.Sprintf(`**ScoreItem:** %s
**Nome curto:** %s`, fullName, scoreItem.Name)

	if scoreItem.Unit != nil {
		siContext += fmt.Sprintf(`
**Unidade:** %s`, *scoreItem.Unit)
	}

	if scoreItem.Gender != nil && *scoreItem.Gender != "not_applicable" {
		siContext += fmt.Sprintf(`
**Gênero aplicável:** %s`, *scoreItem.Gender)
	}

	if scoreItem.AgeRangeMin != nil || scoreItem.AgeRangeMax != nil {
		siContext += fmt.Sprintf(`
**Faixa etária:** %s`, formatAgeRangeForPrompt(scoreItem.AgeRangeMin, scoreItem.AgeRangeMax))
	}

	sciContext := fmt.Sprintf(`

**Contexto científico disponível:**
- %d chunks de %d artigos científicos
- Avg similarity: %.3f
- Seções: results, discussion, methods, introduction, abstract

`, totalChunks, articlesCount, metadata["avg_similarity"])

	// PROMPT 1: Clinical Relevance
	promptCR := fmt.Sprintf(`Você é um médico especialista em medicina baseada em evidências.

%s%s

## CHUNKS CIENTÍFICOS

%s

## TAREFA

Analise os %d chunks científicos acima e gere um texto de **RELEVÂNCIA CLÍNICA** (clinical_relevance) para médicos:

**Requisitos:**
- Extensão: 1200-1800 caracteres
- Linguagem técnica e precisa
- Incluir: definição fisiológica, valores de referência, fisiopatologia resumida
- Citar dados epidemiológicos com NÚMEROS (prevalência, RR, OR, sensibilidade/especificidade)
- Estratificação de risco por valores
- Considerações populacionais (idade, gênero, comorbidades)
- Base em evidências dos chunks fornecidos

**Retorne APENAS o texto da relevância clínica, sem preâmbulos.**`, siContext, sciContext, chunksText, totalChunks)

	// PROMPT 2: Patient Explanation
	promptPE := fmt.Sprintf(`Você é um médico especialista em comunicação com pacientes.

%s%s

## CHUNKS CIENTÍFICOS

%s

## TAREFA

Analise os %d chunks científicos acima e gere uma **EXPLICAÇÃO PARA PACIENTE** (patient_explanation):

**Requisitos:**
- Extensão: 600-900 caracteres
- Linguagem SIMPLES (8º ano escolar)
- Explicar: O QUE é o parâmetro, POR QUE importa, O QUE significam valores alterados
- Tom empático e empoderador
- Sem jargão médico complexo
- Base em evidências dos chunks

**Retorne APENAS o texto da explicação, sem preâmbulos.**`, siContext, sciContext, chunksText, totalChunks)

	// PROMPT 3: Conduct
	promptConduct := fmt.Sprintf(`Você é um médico especialista em protocolos clínicos.

%s%s

## CHUNKS CIENTÍFICOS

%s

## TAREFA

Analise os %d chunks científicos acima e gere **CONDUTAS CLÍNICAS** (conduct) em formato Markdown:

**Requisitos:**
- Extensão: 1000-1500 caracteres
- Formato Markdown com seções:
  ## Interpretação de Valores
  ## Investigação Complementar
  ## Protocolo de Tratamento
  ## Critérios de Referência/Encaminhamento
- Condutas baseadas em evidências dos chunks
- Incluir valores de corte quando aplicável
- Especificar exames complementares
- Critérios objetivos de encaminhamento

**Retorne APENAS o texto da conduta em Markdown, sem preâmbulos.**`, siContext, sciContext, chunksText, totalChunks)

	// PROMPT 4: Max Points
	promptPoints := fmt.Sprintf(`Você é um médico especialista em estratificação de risco.

%s%s

## CHUNKS CIENTÍFICOS

%s

## TAREFA

Analise os %d chunks científicos acima e determine **MAX_POINTS** (0-50) para este parâmetro:

**Critérios de pontuação:**
- Impacto prognóstico (0-20 pontos): RR de mortalidade, eventos CV, etc
- Modificabilidade (0-15 pontos): Lifestyle, farmacológico, cirúrgico
- Prevalência (0-10 pontos): Comum vs raro
- Urgência clínica (0-5 pontos): Emergência vs rotina

**Exemplos:**
- Glicemia jejum: 35-40 pontos (alto impacto, modificável, comum)
- Hemoglobina: 30-35 pontos (moderado impacto, modificável, comum)
- Vitamina D: 15-20 pontos (baixo impacto, modificável, muito comum)
- Genética: 5-15 pontos (variável impacto, não modificável)

**Retorne APENAS o número (0-50) com 1 linha de justificativa.**`, siContext, sciContext, chunksText, totalChunks)

	return EnrichmentPrompts{
		ClinicalRelevance:  promptCR,
		PatientExplanation: promptPE,
		Conduct:            promptConduct,
		MaxPoints:          promptPoints,
	}
}

// formatChunksForPrompt formata chunks para inclusão no prompt
func formatChunksForPrompt(chunks []repository.ChunkSearchResult) string {
	var builder strings.Builder

	for i, chunk := range chunks {
		builder.WriteString(fmt.Sprintf("### Chunk %d/%d\n", i+1, len(chunks)))
		builder.WriteString(fmt.Sprintf("**Article:** %s (%d)\n", chunk.ArticleTitle, chunk.ArticleYear))
		if chunk.Journal != "" {
			builder.WriteString(fmt.Sprintf("**Journal:** %s\n", chunk.Journal))
		}
		builder.WriteString(fmt.Sprintf("**Section:** %s | **Similarity:** %.3f\n\n", chunk.Section, chunk.Similarity))
		builder.WriteString(chunk.ChunkText)
		builder.WriteString("\n\n---\n\n")
	}

	return builder.String()
}

// extractUniqueArticles extrai IDs únicos de artigos
func extractUniqueArticles(chunks []repository.ChunkSearchResult) map[uuid.UUID]bool {
	articles := make(map[uuid.UUID]bool)
	for _, chunk := range chunks {
		articles[chunk.ArticleID] = true
	}
	return articles
}

// formatAgeRangeForPrompt formata faixa etária para prompts
func formatAgeRangeForPrompt(min, max *int) string {
	if min != nil && max != nil {
		return fmt.Sprintf("%d-%d anos", *min, *max)
	}
	if min != nil {
		return fmt.Sprintf("≥%d anos", *min)
	}
	if max != nil {
		return fmt.Sprintf("≤%d anos", *max)
	}
	return "todas as idades"
}

// buildPrompts gera os 4 prompts prontos para Claude
func (s *ScoreEnrichmentPreparationService) buildPrompts(
	scoreItem *models.ScoreItem,
	fullName string,
	chunks []repository.ChunkSearchResult,
	articlesCount int,
	avgSimilarity float64,
) EnrichmentPrompts {
	// Formatar chunks
	var chunksText strings.Builder
	for i, chunk := range chunks {
		chunksText.WriteString(fmt.Sprintf("### Chunk %d/%d\n", i+1, len(chunks)))
		chunksText.WriteString(fmt.Sprintf("**Article:** %s (%d)\n", chunk.ArticleTitle, chunk.ArticleYear))
		if chunk.Journal != "" {
			chunksText.WriteString(fmt.Sprintf("**Journal:** %s\n", chunk.Journal))
		}
		chunksText.WriteString(fmt.Sprintf("**Section:** %s | **Similarity:** %.3f\n\n", chunk.Section, chunk.Similarity))
		chunksText.WriteString(chunk.ChunkText)
		chunksText.WriteString("\n\n---\n\n")
	}

	// Contexto ScoreItem
	siContext := fmt.Sprintf("**ScoreItem:** %s\n**Nome curto:** %s", fullName, scoreItem.Name)
	if scoreItem.Unit != nil {
		siContext += fmt.Sprintf("\n**Unidade:** %s", *scoreItem.Unit)
	}
	if scoreItem.Gender != nil && *scoreItem.Gender != "not_applicable" {
		siContext += fmt.Sprintf("\n**Gênero:** %s", *scoreItem.Gender)
	}

	sciContext := fmt.Sprintf("\n\n**Contexto:** %d chunks de %d artigos (avg sim: %.3f)\n\n",
		len(chunks), articlesCount, avgSimilarity)

	// PROMPT 1: Clinical Relevance
	promptCR := fmt.Sprintf(`Você é um médico especialista em medicina baseada em evidências.

%s%s## CHUNKS CIENTÍFICOS

%s
## TAREFA

Gere **RELEVÂNCIA CLÍNICA** (1200-1800 chars) para médicos:
- Definição fisiológica e valores de referência
- Fisiopatologia resumida
- Dados epidemiológicos com NÚMEROS
- Estratificação de risco

**Retorne APENAS o texto, sem preâmbulos.**`,
		siContext, sciContext, chunksText.String())

	// PROMPT 2
	promptPE := fmt.Sprintf(`Você é um médico especialista.

%s%s## CHUNKS CIENTÍFICOS

%s
## TAREFA

Gere **EXPLICAÇÃO PARA PACIENTE** (600-900 chars):
- Linguagem simples
- O QUE é, POR QUE importa, O QUE significam valores alterados

**Retorne APENAS o texto.**`,
		siContext, sciContext, chunksText.String())

	// PROMPT 3
	promptConduct := fmt.Sprintf(`Você é um médico especialista.

%s%s## CHUNKS CIENTÍFICOS

%s
## TAREFA

Gere **CONDUTAS** (1000-1500 chars) em Markdown:
## Interpretação de Valores
## Investigação Complementar
## Protocolo de Tratamento
## Critérios de Referência

**Retorne APENAS o Markdown.**`,
		siContext, sciContext, chunksText.String())

	// PROMPT 4
	promptPoints := fmt.Sprintf(`Estratificação de risco.

%s%s
Determine **MAX_POINTS** (0-50):
- Impacto prognóstico (0-20)
- Modificabilidade (0-15)
- Prevalência (0-10)
- Urgência (0-5)

**Retorne número + justificativa.**`,
		siContext, sciContext)

	return EnrichmentPrompts{
		ClinicalRelevance:  promptCR,
		PatientExplanation: promptPE,
		Conduct:            promptConduct,
		MaxPoints:          promptPoints,
	}
}
