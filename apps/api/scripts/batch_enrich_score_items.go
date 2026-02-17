package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

// BatchEnrichmentReport representa o resultado do processamento
type BatchEnrichmentReport struct {
	ItemName           string                     `json:"item_name"`
	ItemID             uuid.UUID                  `json:"item_id"`
	Tier               services.EnrichmentTier    `json:"tier"`
	Status             string                     `json:"status"` // success, failed, skipped
	ConfidenceScore    float64                    `json:"confidence_score,omitempty"`
	ArticlesBefore     int64                      `json:"articles_before"`
	ArticlesAfter      int64                      `json:"articles_after"`
	ErrorMessage       string                     `json:"error_message,omitempty"`
	ProcessingDuration time.Duration              `json:"processing_duration"`
	Timestamp          time.Time                  `json:"timestamp"`
}

// BatchEnrichmentService orquestra enriquecimento em lote
type BatchEnrichmentService struct {
	db                *gorm.DB
	pubmedService     *services.PubMedService
	enrichmentService *services.ScoreItemEnrichmentService
	semanticService   *services.ArticleSemanticService
	embeddingService  *services.EmbeddingService
	articleService    *services.ArticleService
}

func main() {
	// Parse argumentos
	// Usage: go run scripts/batch_enrich_score_items.go [offset] [limit]
	// Exemplo: go run scripts/batch_enrich_score_items.go 5 1  (processa item #6)
	offset := 5
	limit := 1

	if len(os.Args) >= 2 {
		if val, err := strconv.Atoi(os.Args[1]); err == nil {
			offset = val
		}
	}
	if len(os.Args) >= 3 {
		if val, err := strconv.Atoi(os.Args[2]); err == nil {
			limit = val
		}
	}

	// Carregar config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar banco
	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	// Inicializar serviços
	queueService := services.NewEmbeddingQueueService(db)
	service := &BatchEnrichmentService{
		db:                db,
		pubmedService:     services.NewPubMedService(db),
		enrichmentService: services.NewScoreItemEnrichmentService(db),
		semanticService:   initSemanticService(cfg, db),
		embeddingService:  services.NewEmbeddingService(cfg, db),
		articleService:    services.NewArticleService(db, "/app/uploads/articles", queueService),
	}

	ctx := context.Background()

	// Executar batch enrichment
	reports, err := service.ProcessBatch(ctx, offset, limit)
	if err != nil {
		log.Fatalf("Batch processing failed: %v", err)
	}

	// Gerar relatório final
	service.PrintFinalReport(reports)

	// Salvar relatório em JSON
	if err := service.SaveReportToFile(reports, "batch_enrichment_report.json"); err != nil {
		log.Printf("⚠️  Failed to save report to file: %v", err)
	}
}

func initSemanticService(cfg *config.Config, db *gorm.DB) *services.ArticleSemanticService {
	vectorRepo := repository.NewArticleVectorRepository(db)
	embeddingService := services.NewEmbeddingService(cfg, db)
	return services.NewArticleSemanticService(vectorRepo, embeddingService)
}

// ProcessBatch processa múltiplos ScoreItems em sequência
func (s *BatchEnrichmentService) ProcessBatch(ctx context.Context, offset, limit int) ([]BatchEnrichmentReport, error) {
	log.Printf("\n🚀 Iniciando batch enrichment (offset=%d, limit=%d)\n", offset, limit)
	log.Println(strings.Repeat("=", 80))

	// Buscar ScoreItems com last_review IS NULL
	var items []models.ScoreItem
	query := s.db.Where("last_review IS NULL").
		Order("created_at ASC").
		Offset(offset).
		Limit(limit)

	if err := query.Find(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch score items: %w", err)
	}

	if len(items) == 0 {
		log.Println("✅ Nenhum ScoreItem pendente de revisão encontrado")
		return []BatchEnrichmentReport{}, nil
	}

	log.Printf("📋 Encontrados %d ScoreItems para processar\n\n", len(items))

	// Processar cada item
	reports := make([]BatchEnrichmentReport, 0, len(items))
	for i, item := range items {
		log.Printf("\n%s\n", strings.Repeat("-", 80))
		log.Printf("🔄 Processando item %d/%d: %s\n", i+1, len(items), item.Name)
		log.Printf("%s\n", strings.Repeat("-", 80))

		startTime := time.Now()
		report := s.ProcessSingleItem(ctx, &item)
		report.ProcessingDuration = time.Since(startTime)
		report.Timestamp = time.Now()

		reports = append(reports, report)

		// Log resumo do item
		s.PrintItemSummary(report, i+1, len(items))

		// Delay entre requisições (rate limiting para Claude API)
		if i < len(items)-1 {
			log.Println("\n⏳ Aguardando 5 segundos antes do próximo item...")
			time.Sleep(5 * time.Second)
		}
	}

	return reports, nil
}

// ProcessSingleItem processa um único ScoreItem
func (s *BatchEnrichmentService) ProcessSingleItem(ctx context.Context, item *models.ScoreItem) BatchEnrichmentReport {
	report := BatchEnrichmentReport{
		ItemName: item.Name,
		ItemID:   item.ID,
		Status:   "processing",
	}

	// 1. Contar artigos antes
	var articleCount int64
	err := s.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", item.ID).
		Count(&articleCount).Error
	if err != nil {
		report.Status = "failed"
		report.ErrorMessage = fmt.Sprintf("failed to count articles: %v", err)
		return report
	}
	report.ArticlesBefore = articleCount

	// 2. Buscar artigos similares via RAG
	log.Println("🔎 Buscando artigos via RAG...")
	ragArticles, err := s.semanticService.RecommendArticlesForScoreItem(item.ID, 20)
	if err != nil {
		log.Printf("⚠️  RAG search failed: %v (continuing)\n", err)
		ragArticles = []repository.ArticleSearchResult{}
	} else {
		log.Printf("📊 RAG encontrou %d artigos\n", len(ragArticles))
	}

	// 3. Linkar artigos RAG
	linkedCount := s.linkRagArticles(ragArticles, item.ID)
	log.Printf("🔗 Linkados %d artigos via RAG\n", linkedCount)

	// 4. Recontar artigos
	s.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", item.ID).
		Count(&articleCount)

	// 5. Buscar no PubMed se necessário
	minArticles := int64(7)
	if articleCount < minArticles {
		needed := minArticles - articleCount
		log.Printf("🌐 Buscando %d artigos no PubMed...\n", needed)

		if err := s.searchAndLinkPubMed(ctx, item, int(needed)); err != nil {
			log.Printf("⚠️  PubMed search failed: %v (continuing)\n", err)
		}
	}

	// 6. Recontar artigos finais
	s.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", item.ID).
		Count(&articleCount)
	report.ArticlesAfter = articleCount

	// 7. Determinar tier
	tier := s.enrichmentService.DetermineTier(item)
	report.Tier = tier
	log.Printf("📊 Tier: %s\n", tier)

	// 8. Enriquecer campos clínicos
	if tier == services.TierPreserve {
		log.Println("✅ Conteúdo já excelente, preservando")
		report.Status = "skipped"
		report.ConfidenceScore = 1.0
		return report
	}

	log.Println("🤖 Enriquecendo com LLM...")
	result, err := s.enrichmentService.GenerateEnrichment(ctx, item, tier)
	if err != nil {
		report.Status = "failed"
		report.ErrorMessage = fmt.Sprintf("LLM enrichment failed: %v", err)
		return report
	}

	report.ConfidenceScore = result.Confidence

	// 9. Validar resultado
	validationErrors := s.enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		log.Printf("⚠️  Validation warnings: %v\n", validationErrors)
	}

	// 10. Salvar histórico
	if err := s.saveReviewHistory(item, result, tier); err != nil {
		log.Printf("⚠️  Failed to save history: %v (continuing)\n", err)
	}

	// 11. Atualizar ScoreItem
	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"points":              float64(result.MaxPoints),
		"last_review":         &now,
		"updated_at":          now,
	}

	if err := s.db.Model(item).Updates(updates).Error; err != nil {
		report.Status = "failed"
		report.ErrorMessage = fmt.Sprintf("failed to update item: %v", err)
		return report
	}

	report.Status = "success"
	log.Println("✅ Item enriquecido com sucesso")

	return report
}

// linkRagArticles linka artigos RAG ao ScoreItem
func (s *BatchEnrichmentService) linkRagArticles(ragArticles []repository.ArticleSearchResult, scoreItemID uuid.UUID) int {
	linkedCount := 0
	for _, result := range ragArticles {
		var exists bool
		err := s.db.Raw(`
			SELECT EXISTS(
				SELECT 1 FROM article_score_items
				WHERE article_id = ? AND score_item_id = ?
			)
		`, result.Article.ID, scoreItemID).Scan(&exists).Error

		if err != nil || exists {
			continue
		}

		if err := s.db.Exec(`
			INSERT INTO article_score_items (article_id, score_item_id, linked_at)
			VALUES (?, ?, NOW())
		`, result.Article.ID, scoreItemID).Error; err == nil {
			linkedCount++
		}
	}
	return linkedCount
}

// searchAndLinkPubMed busca e linka artigos do PubMed
func (s *BatchEnrichmentService) searchAndLinkPubMed(ctx context.Context, item *models.ScoreItem, limit int) error {
	// Gerar query
	query := s.generatePubMedQuery(item)
	log.Printf("🔍 Query: %s\n", query)

	// Buscar PMIDs
	pmids, err := s.pubmedService.SearchArticles(ctx, query, limit)
	if err != nil {
		return err
	}

	if len(pmids) == 0 {
		return nil
	}

	log.Printf("📋 PMIDs: %v\n", pmids)

	// Buscar metadata
	articles, err := s.pubmedService.FetchArticleDetails(ctx, pmids)
	if err != nil {
		return err
	}

	// Processar cada artigo
	for _, pmArticle := range articles {
		if err := s.processPubMedArticle(ctx, pmArticle, item.ID); err != nil {
			log.Printf("⚠️  Failed to process %s: %v\n", pmArticle.Title, err)
		}
	}

	return nil
}

// generatePubMedQuery gera query PubMed
func (s *BatchEnrichmentService) generatePubMedQuery(item *models.ScoreItem) string {
	query := fmt.Sprintf(`("%s"[Title/Abstract]`, item.Name)

	if item.Unit != nil && *item.Unit != "" {
		query += ` OR "` + item.Name + `"[MeSH Terms]`
	}

	query += `) AND (`

	if item.Unit != nil && *item.Unit != "" {
		query += `"reference values"[Title/Abstract] OR "normal range"[Title/Abstract] OR "clinical significance"[Title/Abstract]`
	} else {
		query += `"clinical assessment"[Title/Abstract] OR "evaluation"[Title/Abstract] OR "measurement"[Title/Abstract]`
	}

	query += `) AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2018:2026[dp] AND English[lang]`

	return query
}

// processPubMedArticle processa artigo do PubMed
func (s *BatchEnrichmentService) processPubMedArticle(ctx context.Context, pmArticle *services.PubMedArticle, scoreItemID uuid.UUID) error {
	// Verificar se já existe
	var existing models.Article
	err := s.db.Where("pm_id = ?", pmArticle.PMID).First(&existing).Error
	if err == nil {
		// Linkar se necessário
		var linked bool
		s.db.Raw(`
			SELECT EXISTS(
				SELECT 1 FROM article_score_items
				WHERE article_id = ? AND score_item_id = ?
			)
		`, existing.ID, scoreItemID).Scan(&linked)

		if !linked {
			s.db.Exec(`
				INSERT INTO article_score_items (article_id, score_item_id, linked_at)
				VALUES (?, ?, NOW())
			`, existing.ID, scoreItemID)
		}
		return nil
	}

	// Tentar baixar PDF
	var pdfPath string
	if pmArticle.DOI != "" {
		articleID := uuid.Must(uuid.NewV7())
		pdfPath, err = s.pubmedService.DownloadPDF(ctx, pmArticle.DOI, articleID)
		if err != nil {
			pdfPath = ""
		}
	}

	// Criar artigo
	article, err := s.pubmedService.CreateArticleFromPubMed(pmArticle, pdfPath)
	if err != nil {
		return err
	}

	// Linkar
	return s.db.Exec(`
		INSERT INTO article_score_items (article_id, score_item_id, linked_at)
		VALUES (?, ?, NOW())
	`, article.ID, scoreItemID).Error
}

// saveReviewHistory salva histórico
func (s *BatchEnrichmentService) saveReviewHistory(
	item *models.ScoreItem,
	result *services.EnrichmentResult,
	tier services.EnrichmentTier,
) error {
	beforeJSON, _ := json.Marshal(map[string]interface{}{
		"clinical_relevance":  item.ClinicalRelevance,
		"patient_explanation": item.PatientExplanation,
		"conduct":             item.Conduct,
		"points":              item.Points,
	})

	afterJSON, _ := json.Marshal(map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"max_points":          result.MaxPoints,
	})

	id := uuid.Must(uuid.NewV7())
	now := time.Now()

	return s.db.Exec(`
		INSERT INTO score_item_review_history
		(id, score_item_id, review_type, before_snapshot, after_snapshot, tier, confidence_score, model_used, reviewed_at, created_at)
		VALUES (?, ?, 'llm_enrichment', ?, ?, ?, ?, 'claude-sonnet-4-5-20250929', ?, ?)
	`,
		id,
		item.ID,
		string(beforeJSON),
		string(afterJSON),
		string(tier),
		result.Confidence,
		now,
		now,
	).Error
}

// PrintItemSummary imprime resumo de um item
func (s *BatchEnrichmentService) PrintItemSummary(report BatchEnrichmentReport, current, total int) {
	log.Printf("\n📊 Resumo [%d/%d]:\n", current, total)
	log.Printf("   Nome: %s\n", report.ItemName)
	log.Printf("   Tier: %s\n", report.Tier)
	log.Printf("   Status: %s\n", report.Status)
	if report.Status == "success" {
		log.Printf("   Confidence: %.2f\n", report.ConfidenceScore)
	}
	log.Printf("   Artigos: %d → %d\n", report.ArticlesBefore, report.ArticlesAfter)
	if report.ErrorMessage != "" {
		log.Printf("   Erro: %s\n", report.ErrorMessage)
	}
	log.Printf("   Duração: %s\n", report.ProcessingDuration.Round(time.Second))
}

// PrintFinalReport imprime relatório final consolidado
func (s *BatchEnrichmentService) PrintFinalReport(reports []BatchEnrichmentReport) {
	log.Println("\n" + strings.Repeat("=", 80))
	log.Println("📊 RELATÓRIO FINAL DO BATCH")
	log.Println(strings.Repeat("=", 80))

	if len(reports) == 0 {
		log.Println("Nenhum item processado")
		return
	}

	// Estatísticas
	totalSuccess := 0
	totalFailed := 0
	totalSkipped := 0
	totalDuration := time.Duration(0)
	avgConfidence := 0.0

	for _, r := range reports {
		switch r.Status {
		case "success":
			totalSuccess++
			avgConfidence += r.ConfidenceScore
		case "failed":
			totalFailed++
		case "skipped":
			totalSkipped++
		}
		totalDuration += r.ProcessingDuration
	}

	if totalSuccess > 0 {
		avgConfidence /= float64(totalSuccess)
	}

	log.Printf("\n📈 Estatísticas:\n")
	log.Printf("   Total processado: %d\n", len(reports))
	log.Printf("   ✅ Sucesso: %d\n", totalSuccess)
	log.Printf("   ⏭️  Skipped: %d\n", totalSkipped)
	log.Printf("   ❌ Falhas: %d\n", totalFailed)
	log.Printf("   📊 Confidence média: %.2f\n", avgConfidence)
	log.Printf("   ⏱️  Tempo total: %s\n", totalDuration.Round(time.Second))

	// Detalhes por item
	log.Printf("\n📋 Detalhes:\n")
	log.Println(strings.Repeat("-", 80))
	log.Printf("%-50s %-10s %-10s %s\n", "NOME", "TIER", "STATUS", "CONFIDENCE")
	log.Println(strings.Repeat("-", 80))

	for _, r := range reports {
		confidenceStr := "-"
		if r.Status == "success" {
			confidenceStr = fmt.Sprintf("%.2f", r.ConfidenceScore)
		}

		// Truncar nome se muito longo
		name := r.ItemName
		if len(name) > 47 {
			name = name[:44] + "..."
		}

		log.Printf("%-50s %-10s %-10s %s\n",
			name,
			r.Tier,
			r.Status,
			confidenceStr,
		)
	}

	log.Println(strings.Repeat("=", 80))
}

// SaveReportToFile salva relatório em JSON
func (s *BatchEnrichmentService) SaveReportToFile(reports []BatchEnrichmentReport, filename string) error {
	data, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
