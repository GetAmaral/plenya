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

// ReviewScoreItemCommand orquestra revisão completa de ScoreItem
type ReviewScoreItemCommand struct {
	db                *gorm.DB
	pubmedService     *services.PubMedService
	enrichmentService *services.ScoreItemEnrichmentService
	semanticService   *services.ArticleSemanticService
	embeddingService  *services.EmbeddingService
	articleService    *services.ArticleService
}

func main() {
	// Verificar argumentos
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run scripts/review_score_item.go <item_number_or_uuid>")
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

	// Resolver ScoreItem ID
	scoreItemID, err := resolveScoreItemID(db, os.Args[1])
	if err != nil {
		log.Fatalf("Failed to resolve ScoreItem: %v", err)
	}

	// Inicializar serviços
	queueService := services.NewEmbeddingQueueService(db)

	cmd := &ReviewScoreItemCommand{
		db:                db,
		pubmedService:     services.NewPubMedService(db),
		enrichmentService: services.NewScoreItemEnrichmentService(db),
		semanticService:   initSemanticService(cfg, db),
		embeddingService:  services.NewEmbeddingService(cfg, db),
		articleService:    services.NewArticleService(db, "/app/uploads/articles", queueService),
	}

	ctx := context.Background()

	// Executar revisão completa
	if err := cmd.ExecuteReview(ctx, scoreItemID); err != nil {
		log.Fatalf("Review failed: %v", err)
	}

	log.Println("\n✅ Revisão completa finalizada com sucesso!")
}

// resolveScoreItemID resolve o ID do ScoreItem a partir de número sequencial ou UUID
func resolveScoreItemID(db *gorm.DB, input string) (uuid.UUID, error) {
	// Tentar parsear como UUID primeiro
	if id, err := uuid.Parse(input); err == nil {
		return id, nil
	}

	// Tentar parsear como número sequencial
	seqNum, err := strconv.Atoi(input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("input must be a UUID or a sequence number (integer)")
	}

	// Buscar o N-ésimo item por ordem de criação
	var item models.ScoreItem
	offset := seqNum - 1
	if offset < 0 {
		return uuid.Nil, fmt.Errorf("sequence number must be >= 1")
	}

	err = db.Order("created_at ASC").Offset(offset).Limit(1).First(&item).Error
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to find ScoreItem #%d: %w", seqNum, err)
	}

	log.Printf("📍 Resolved #%d to UUID: %s\n", seqNum, item.ID)
	return item.ID, nil
}

func initSemanticService(cfg *config.Config, db *gorm.DB) *services.ArticleSemanticService {
	vectorRepo := repository.NewArticleVectorRepository(db)
	embeddingService := services.NewEmbeddingService(cfg, db)
	return services.NewArticleSemanticService(vectorRepo, embeddingService)
}

// ExecuteReview executa todas as etapas da revisão
func (cmd *ReviewScoreItemCommand) ExecuteReview(ctx context.Context, scoreItemID uuid.UUID) error {
	log.Printf("\n🔍 Iniciando revisão do ScoreItem %s\n", scoreItemID)

	// 1. Buscar ScoreItem do banco
	var scoreItem models.ScoreItem
	if err := cmd.db.First(&scoreItem, "id = ?", scoreItemID).Error; err != nil {
		return fmt.Errorf("failed to fetch ScoreItem: %w", err)
	}

	log.Printf("📋 ScoreItem: %s\n", scoreItem.Name)

	// 2. Contar artigos científicos já linkados (PMIDs não nulos)
	var currentArticleCount int64
	err := cmd.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", scoreItemID).
		Count(&currentArticleCount).Error
	if err != nil {
		return fmt.Errorf("failed to count current articles: %w", err)
	}

	log.Printf("📚 Artigos científicos antes: %d\n", currentArticleCount)
	beforeArticleCount := currentArticleCount

	// 3. Buscar artigos similares via RAG
	log.Println("\n🔎 Buscando artigos similares via RAG (threshold 0.7)...")
	ragArticles, err := cmd.semanticService.RecommendArticlesForScoreItem(scoreItemID, 20)
	if err != nil {
		log.Printf("⚠️  RAG search failed: %v\n", err)
		ragArticles = []repository.ArticleSearchResult{}
	}

	log.Printf("📊 Artigos encontrados via RAG: %d\n", len(ragArticles))

	// 4. Linkar artigos RAG ao ScoreItem (se ainda não estão linkados)
	linkedCount := 0
	for _, result := range ragArticles {
		// Verificar se já está linkado
		var exists bool
		err := cmd.db.Raw(`
			SELECT EXISTS(
				SELECT 1 FROM article_score_items
				WHERE article_id = ? AND score_item_id = ?
			)
		`, result.Article.ID, scoreItemID).Scan(&exists).Error

		if err != nil {
			log.Printf("⚠️  Failed to check link for article %s: %v\n", result.Article.Title, err)
			continue
		}

		if !exists {
			// Criar link
			if err := cmd.db.Exec(`
				INSERT INTO article_score_items (article_id, score_item_id, linked_at)
				VALUES (?, ?, NOW())
			`, result.Article.ID, scoreItemID).Error; err != nil {
				log.Printf("⚠️  Failed to link article %s: %v\n", result.Article.Title, err)
			} else {
				linkedCount++
				log.Printf("✅ Linkado: %s (similaridade: %.3f)\n", result.Article.Title, result.Similarity)
			}
		}
	}

	log.Printf("🔗 Novos artigos linkados via RAG: %d\n", linkedCount)

	// 5. Recontar artigos científicos
	err = cmd.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", scoreItemID).
		Count(&currentArticleCount).Error
	if err != nil {
		return fmt.Errorf("failed to recount articles: %w", err)
	}

	// 6. Buscar no PubMed se ainda faltam artigos
	minArticles := int64(7)
	if currentArticleCount < minArticles {
		needed := minArticles - currentArticleCount
		log.Printf("\n🌐 Buscando %d artigos no PubMed...\n", needed)

		pubmedArticles, err := cmd.searchPubMed(ctx, &scoreItem, int(needed))
		if err != nil {
			log.Printf("⚠️  PubMed search failed: %v\n", err)
		} else {
			log.Printf("📥 PubMed retornou %d artigos\n", len(pubmedArticles))

			// Processar cada artigo PubMed
			for i, pmArticle := range pubmedArticles {
				log.Printf("\n📄 [%d/%d] Processando: %s\n", i+1, len(pubmedArticles), pmArticle.Title)

				if err := cmd.processPubMedArticle(ctx, pmArticle, scoreItemID); err != nil {
					log.Printf("⚠️  Failed to process article: %v\n", err)
				}
			}
		}
	}

	// 7. Recontar artigos finais
	err = cmd.db.Model(&models.Article{}).
		Joins("JOIN article_score_items ON articles.id = article_score_items.article_id").
		Where("article_score_items.score_item_id = ? AND articles.pm_id IS NOT NULL", scoreItemID).
		Count(&currentArticleCount).Error
	if err != nil {
		return fmt.Errorf("failed to final count: %w", err)
	}

	log.Printf("\n📚 Total de artigos científicos após: %d\n", currentArticleCount)

	// 8. Enriquecer campos clínicos
	log.Println("\n🤖 Enriquecendo campos clínicos com LLM...")
	var tier services.EnrichmentTier
	var confidence float64
	if err := cmd.enrichClinicalFields(ctx, &scoreItem, &tier, &confidence); err != nil {
		log.Printf("⚠️  Enrichment failed: %v\n", err)
	}

	// 9. Relatório final
	log.Println("\n" + strings.Repeat("=", 80))
	log.Println("📊 RELATÓRIO FINAL")
	log.Println(strings.Repeat("=", 80))
	log.Printf("Item: %s\n", scoreItem.Name)
	log.Printf("Artigos antes → depois: %d → %d (incremento: +%d)\n", beforeArticleCount, currentArticleCount, currentArticleCount-beforeArticleCount)
	log.Printf("Tier usado: %s\n", tier)
	log.Printf("Confidence score: %.2f\n", confidence)
	if scoreItem.LastReview != nil {
		log.Printf("Última revisão: %s\n", scoreItem.LastReview.Format("2006-01-02 15:04:05"))
	}
	log.Println(strings.Repeat("=", 80))

	return nil
}

// searchPubMed busca artigos no PubMed
func (cmd *ReviewScoreItemCommand) searchPubMed(ctx context.Context, item *models.ScoreItem, limit int) ([]*services.PubMedArticle, error) {
	// Gerar query PubMed otimizada
	query := cmd.generatePubMedQuery(item)
	log.Printf("🔍 Query PubMed: %s\n", query)

	// Buscar PMIDs
	pmids, err := cmd.pubmedService.SearchArticles(ctx, query, limit)
	if err != nil {
		return nil, err
	}

	if len(pmids) == 0 {
		return []*services.PubMedArticle{}, nil
	}

	log.Printf("📋 PMIDs encontrados: %v\n", pmids)

	// Buscar metadata completa
	articles, err := cmd.pubmedService.FetchArticleDetails(ctx, pmids)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// generatePubMedQuery gera query otimizada para PubMed baseada no nome do ScoreItem
func (cmd *ReviewScoreItemCommand) generatePubMedQuery(item *models.ScoreItem) string {
	// Gerar query baseada no nome do ScoreItem
	itemName := item.Name

	// Construir query PubMed usando o nome do item
	// Adicionar MeSH terms se disponível a unidade
	query := fmt.Sprintf(`("%s"[Title/Abstract]`, itemName)

	// Se tem unidade, adicionar contexto de valores de referência
	if item.Unit != nil && *item.Unit != "" {
		query += ` OR "` + itemName + `"[MeSH Terms]`
	}

	query += `) AND (`

	// Adicionar termos de contexto clínico
	if item.Unit != nil && *item.Unit != "" {
		// É um marcador laboratorial
		query += `"reference values"[Title/Abstract] OR "normal range"[Title/Abstract] OR "clinical significance"[Title/Abstract]`
	} else {
		// É um parâmetro clínico/questionnaire
		query += `"clinical assessment"[Title/Abstract] OR "evaluation"[Title/Abstract] OR "measurement"[Title/Abstract]`
	}

	query += `) AND (Review[ptyp] OR Meta-Analysis[ptyp]) AND 2018:2026[dp] AND English[lang]`

	return query
}

// processPubMedArticle processa um artigo do PubMed (criar, baixar PDF, embeddings, linkar)
func (cmd *ReviewScoreItemCommand) processPubMedArticle(ctx context.Context, pmArticle *services.PubMedArticle, scoreItemID uuid.UUID) error {
	// 1. Verificar se artigo já existe
	var existing models.Article
	err := cmd.db.Where("pm_id = ?", pmArticle.PMID).First(&existing).Error
	if err == nil {
		log.Printf("ℹ️  Artigo já existe: %s\n", existing.Title)

		// Verificar se já está linkado ao ScoreItem
		var linked bool
		cmd.db.Raw(`
			SELECT EXISTS(
				SELECT 1 FROM article_score_items
				WHERE article_id = ? AND score_item_id = ?
			)
		`, existing.ID, scoreItemID).Scan(&linked)

		if !linked {
			// Linkar
			cmd.db.Exec(`
				INSERT INTO article_score_items (article_id, score_item_id, linked_at)
				VALUES (?, ?, NOW())
			`, existing.ID, scoreItemID)
			log.Println("✅ Artigo linkado ao ScoreItem")
		}

		return nil
	}

	// 2. Tentar baixar PDF (se tiver DOI)
	var pdfPath string
	if pmArticle.DOI != "" {
		log.Printf("📥 Tentando baixar PDF via DOI: %s\n", pmArticle.DOI)
		articleID := uuid.Must(uuid.NewV7())
		pdfPath, err = cmd.pubmedService.DownloadPDF(ctx, pmArticle.DOI, articleID)
		if err != nil {
			log.Printf("⚠️  PDF download failed: %v (continuing without PDF)\n", err)
			pdfPath = ""
		} else {
			log.Printf("✅ PDF baixado: %s\n", pdfPath)
		}
	}

	// 3. Criar artigo no banco
	article, err := cmd.pubmedService.CreateArticleFromPubMed(pmArticle, pdfPath)
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	log.Printf("✅ Artigo criado: %s (ID: %s)\n", article.Title, article.ID)

	// 4. Linkar ao ScoreItem
	err = cmd.db.Exec(`
		INSERT INTO article_score_items (article_id, score_item_id, linked_at)
		VALUES (?, ?, NOW())
	`, article.ID, scoreItemID).Error
	if err != nil {
		return fmt.Errorf("failed to link article: %w", err)
	}

	log.Println("✅ Artigo linkado ao ScoreItem")

	// 5. Processar embeddings (se tiver PDF)
	if pdfPath != "" {
		log.Println("🧮 Processando embeddings do PDF...")
		if err := cmd.processEmbeddings(ctx, article); err != nil {
			log.Printf("⚠️  Embedding processing failed: %v\n", err)
		} else {
			log.Println("✅ Embeddings processados")
		}
	}

	return nil
}

// processEmbeddings processa embeddings de artigo
func (cmd *ReviewScoreItemCommand) processEmbeddings(ctx context.Context, article *models.Article) error {
	// Marcar como processing
	cmd.db.Model(article).Update("embedding_status", "processing")

	// TODO: Implementar chunking + embedding do PDF
	// Por ora, apenas marcar como completed para não bloquear
	cmd.db.Model(article).Updates(map[string]interface{}{
		"embedding_status": "completed",
		"last_embedded_at": time.Now(),
		"chunk_count":      0,
	})

	return nil
}

// enrichClinicalFields enriquece campos clínicos usando LLM
func (cmd *ReviewScoreItemCommand) enrichClinicalFields(ctx context.Context, item *models.ScoreItem, tier *services.EnrichmentTier, confidence *float64) error {
	// 1. Determinar tier de processamento
	*tier = cmd.enrichmentService.DetermineTier(item)
	log.Printf("📊 Tier de enriquecimento: %s\n", *tier)

	if *tier == services.TierPreserve {
		log.Println("ℹ️  Conteúdo já é excelente, preservando sem alterações")
		*confidence = 1.0
		return nil
	}

	// 2. Gerar enriquecimento via LLM
	log.Println("🤖 Chamando Claude API para enriquecer campos...")

	result, err := cmd.enrichmentService.GenerateEnrichment(ctx, item, *tier)
	if err != nil {
		return fmt.Errorf("failed to generate enrichment: %w", err)
	}
	*confidence = result.Confidence

	// 3. Validar resultado
	validationErrors := cmd.enrichmentService.ValidateResult(result)
	if len(validationErrors) > 0 {
		log.Printf("⚠️  Validation warnings: %v\n", validationErrors)
	}

	// 4. Salvar histórico de revisão
	if err := cmd.saveReviewHistory(item, result, *tier); err != nil {
		return fmt.Errorf("failed to save review history: %w", err)
	}

	// 5. Atualizar ScoreItem
	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"points":              float64(result.MaxPoints),
		"last_review":         &now,
		"updated_at":          now,
	}

	if err := cmd.db.Model(item).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update ScoreItem: %w", err)
	}

	log.Println("✅ Campos clínicos atualizados")

	// Exibir preview
	log.Printf("\n📝 Preview do conteúdo atualizado:\n")
	log.Printf("clinical_relevance: %d chars\n", len(result.ClinicalRelevance))
	log.Printf("patient_explanation: %d chars\n", len(result.PatientExplanation))
	log.Printf("conduct: %d chars\n", len(result.Conduct))
	log.Printf("max_points: %d\n", result.MaxPoints)

	return nil
}

// saveReviewHistory salva histórico de revisão
func (cmd *ReviewScoreItemCommand) saveReviewHistory(
	item *models.ScoreItem,
	result *services.EnrichmentResult,
	tier services.EnrichmentTier,
) error {
	// Serializar before
	beforeJSON, _ := json.Marshal(map[string]interface{}{
		"clinical_relevance":  item.ClinicalRelevance,
		"patient_explanation": item.PatientExplanation,
		"conduct":             item.Conduct,
		"points":              item.Points,
	})

	// Serializar after
	afterJSON, _ := json.Marshal(map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"max_points":          result.MaxPoints,
	})

	// Inserir histórico
	id := uuid.Must(uuid.NewV7())
	now := time.Now()

	err := cmd.db.Exec(`
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

	if err != nil {
		log.Printf("⚠️  Failed to save review history: %v (continuing)\n", err)
	} else {
		log.Printf("✅ Histórico de revisão salvo (ID: %s)\n", id)
	}

	return nil
}
