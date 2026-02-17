package main

import (
	"context"
	"fmt"
	"log"
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

const (
	scoreItemID     = "c77cedd3-2800-70ef-abfd-b3cdd23b2ced"
	minArticles     = 7
	ragThreshold    = 0.7
	pubmedMaxSearch = 10
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar banco
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	db := database.DB

	ctx := context.Background()

	// Inicializar serviços
	vectorRepo := repository.NewArticleVectorRepository(db)
	enrichmentService := services.NewScoreItemEnrichmentService(db)
	pubmedService := services.NewPubMedService(db)
	embeddingQueueService := services.NewEmbeddingQueueService(db)

	log.Println("=== REVISÃO COMPLETA: Adesão ao tratamento nutricional ===")
	log.Printf("ScoreItem ID: %s\n", scoreItemID)

	// 1. Buscar ScoreItem atual
	itemUUID, err := uuid.Parse(scoreItemID)
	if err != nil {
		log.Fatalf("Invalid UUID: %v", err)
	}

	var item models.ScoreItem
	if err := db.Preload("Subgroup.Group").First(&item, itemUUID).Error; err != nil {
		log.Fatalf("Failed to fetch ScoreItem: %v", err)
	}

	log.Printf("\n✓ ScoreItem encontrado: %s", item.Name)
	log.Printf("  Grupo: %s > %s", item.Subgroup.Group.Name, item.Subgroup.Name)
	log.Printf("  Pontuação atual: %.1f", getPoints(item.Points))
	log.Printf("  Última revisão: %s\n", formatLastReview(item.LastReview))

	// 2. Verificar artigos linkados atuais
	var currentArticleCount int64
	if err := db.Table("article_score_items").
		Where("score_item_id = ?", itemUUID).
		Count(&currentArticleCount).Error; err != nil {
		log.Fatalf("Failed to count articles: %v", err)
	}

	log.Printf("\n📚 Artigos linkados atuais: %d", currentArticleCount)

	// 3. Buscar artigos similares via RAG
	log.Println("\n🔍 Buscando artigos similares via RAG...")
	ragResults, err := vectorRepo.FindSimilarArticlesForScoreItem(itemUUID, 20, ragThreshold)
	if err != nil {
		log.Printf("Warning: RAG search failed: %v", err)
		ragResults = []repository.ArticleSearchResult{}
	}

	log.Printf("✓ RAG encontrou %d artigos similares (threshold %.2f)", len(ragResults), ragThreshold)

	// Printar top 5 artigos RAG
	if len(ragResults) > 0 {
		log.Println("\nTop 5 artigos RAG:")
		for i, result := range ragResults {
			if i >= 5 {
				break
			}
			log.Printf("  %d. [%.3f] %s (%s, %d)",
				i+1,
				result.Similarity,
				result.Article.Title,
				result.Article.Journal,
				result.Article.PublishDate.Year())
		}
	}

	// 4. Verificar se precisa buscar mais artigos no PubMed
	totalAvailable := int(currentArticleCount) + len(ragResults)
	needsPubMed := totalAvailable < minArticles

	log.Printf("\n📊 Total de artigos disponíveis: %d (mínimo: %d)", totalAvailable, minArticles)

	var pubmedArticles []*models.Article
	if needsPubMed {
		log.Printf("\n⚠️  Artigos insuficientes. Buscando no PubMed...")

		// Gerar query PubMed otimizada
		pubmedQuery := generatePubMedQuery(&item)
		log.Printf("Query PubMed: %s", pubmedQuery)

		// Buscar PMIDs
		pmids, err := pubmedService.RateLimitedSearch(ctx, pubmedQuery, pubmedMaxSearch)
		if err != nil {
			log.Printf("Warning: PubMed search failed: %v", err)
		} else {
			log.Printf("✓ PubMed encontrou %d artigos", len(pmids))

			// Fetch metadata de cada artigo
			if len(pmids) > 0 {
				pubmedMetadata, err := pubmedService.FetchArticleDetails(ctx, pmids)
				if err != nil {
					log.Printf("Warning: Failed to fetch PubMed details: %v", err)
				} else {
					// Criar artigos no banco
					for _, pm := range pubmedMetadata {
						// Tentar baixar PDF
						var pdfPath string
						if pm.DOI != "" {
							pdfPath, _ = pubmedService.DownloadPDF(ctx, pm.DOI, uuid.New())
							if pdfPath != "" {
								log.Printf("  ✓ PDF baixado: %s", pm.Title)
							}
						}

						// Criar artigo
						article, err := pubmedService.CreateArticleFromPubMed(pm, pdfPath)
						if err != nil {
							log.Printf("  ⚠️  Failed to create article: %v", err)
							continue
						}

						pubmedArticles = append(pubmedArticles, article)

						// Linkar ao ScoreItem
						if err := linkArticleToScoreItem(db, article.ID, itemUUID); err != nil {
							log.Printf("  ⚠️  Failed to link article: %v", err)
						}
					}

					log.Printf("✓ %d novos artigos adicionados do PubMed", len(pubmedArticles))
				}
			}
		}
	} else {
		log.Println("✓ Artigos suficientes disponíveis via RAG")
	}

	// 5. Fazer enqueue dos novos artigos para processamento de embeddings
	if len(pubmedArticles) > 0 {
		log.Println("\n🧬 Enfileirando novos artigos para processamento de embeddings...")
		for _, article := range pubmedArticles {
			if err := embeddingQueueService.QueueArticle(article.ID); err != nil {
				log.Printf("  ⚠️  Failed to queue article for embeddings: %v", err)
			} else {
				log.Printf("  ✓ Artigo enfileirado: %s", article.Title)
			}
		}
		log.Println("  ℹ️  Embeddings serão processados em background")
	}

	// 6. Determinar tier de enriquecimento
	tier := enrichmentService.DetermineTier(&item)
	log.Printf("\n📝 Tier de enriquecimento: %s", tier)
	log.Printf("   clinical_relevance: %d chars", len(stringValue(item.ClinicalRelevance)))
	log.Printf("   patient_explanation: %d chars", len(stringValue(item.PatientExplanation)))
	log.Printf("   conduct: %d chars", len(stringValue(item.Conduct)))

	// 7. Gerar enriquecimento via LLM
	var enrichmentResult *services.EnrichmentResult
	if tier == services.TierPreserve {
		log.Println("\n✓ Campos clínicos excelentes - preservando conteúdo atual")
	} else {
		log.Printf("\n🤖 Gerando enriquecimento via Claude Sonnet 4.5...")

		enrichmentResult, err = enrichmentService.GenerateEnrichment(ctx, &item, tier)
		if err != nil {
			log.Fatalf("Failed to generate enrichment: %v", err)
		}

		log.Printf("✓ Enriquecimento gerado (confidence: %.2f)", enrichmentResult.Confidence)

		// Validar resultado
		validationErrors := enrichmentService.ValidateResult(enrichmentResult)
		if len(validationErrors) > 0 {
			log.Printf("\n⚠️  Validation errors:")
			for _, verr := range validationErrors {
				log.Printf("   - %s", verr)
			}
			log.Fatal("Enrichment validation failed. Aborting.")
		}

		log.Println("✓ Validação passou")

		// Preview do resultado
		log.Println("\n📄 Preview do enriquecimento:")
		log.Printf("   clinical_relevance: %d chars", len(enrichmentResult.ClinicalRelevance))
		log.Printf("   patient_explanation: %d chars", len(enrichmentResult.PatientExplanation))
		log.Printf("   conduct: %d chars", len(enrichmentResult.Conduct))
		log.Printf("   max_points: %d (justificativa: %s)",
			enrichmentResult.MaxPoints,
			enrichmentResult.Justification)

		// 8. Salvar review history (audit trail)
		log.Println("\n💾 Salvando audit trail...")
		if err := enrichmentService.SaveReviewHistory(&item, enrichmentResult, tier); err != nil {
			log.Fatalf("Failed to save review history: %v", err)
		}
		log.Println("✓ Review history salvo")

		// 9. Atualizar ScoreItem no banco
		log.Println("\n✍️  Atualizando ScoreItem no banco...")
		now := time.Now()
		updates := map[string]interface{}{
			"clinical_relevance":  enrichmentResult.ClinicalRelevance,
			"patient_explanation": enrichmentResult.PatientExplanation,
			"conduct":             enrichmentResult.Conduct,
			"points":              float64(enrichmentResult.MaxPoints),
			"last_review":         now,
			"updated_at":          now,
		}

		if err := db.Model(&item).Updates(updates).Error; err != nil {
			log.Fatalf("Failed to update ScoreItem: %v", err)
		}

		log.Println("✓ ScoreItem atualizado com sucesso!")
	}

	// 10. Relatório final
	log.Println("\n" + strings.Repeat("=", 60))
	log.Println("REVISÃO COMPLETA FINALIZADA")
	log.Println(strings.Repeat("=", 60))

	// Re-buscar item atualizado
	var updatedItem models.ScoreItem
	db.First(&updatedItem, itemUUID)

	var finalArticleCount int64
	db.Table("article_score_items").
		Where("score_item_id = ?", itemUUID).
		Count(&finalArticleCount)

	log.Printf("\n📊 Estado final:")
	log.Printf("   Artigos linkados: %d", finalArticleCount)
	log.Printf("   Pontuação: %.1f", getPoints(updatedItem.Points))
	log.Printf("   Última revisão: %s", formatLastReview(updatedItem.LastReview))
	log.Printf("   clinical_relevance: %d chars", len(stringValue(updatedItem.ClinicalRelevance)))
	log.Printf("   patient_explanation: %d chars", len(stringValue(updatedItem.PatientExplanation)))
	log.Printf("   conduct: %d chars", len(stringValue(updatedItem.Conduct)))

	if enrichmentResult != nil {
		log.Printf("\n🎯 Mudanças aplicadas:")
		log.Printf("   Tier: %s", tier)
		log.Printf("   Confidence: %.2f", enrichmentResult.Confidence)
		log.Printf("   Justificativa: %s", enrichmentResult.Justification)
	}

	log.Println("\n✅ Script concluído com sucesso!")
}

// Helper functions
func generatePubMedQuery(item *models.ScoreItem) string {
	// Query específica para adesão nutricional
	query := "(adherence OR compliance OR concordance) AND "
	query += "(nutritional therapy OR nutrition treatment OR dietary intervention) AND "
	query += "(assessment OR measurement OR evaluation) "
	query += "AND (Review[ptyp] OR Meta-Analysis[ptyp]) "

	// Últimos 8 anos
	currentYear := time.Now().Year()
	query += fmt.Sprintf("AND %d:%d[dp]", currentYear-8, currentYear)

	return query
}

func linkArticleToScoreItem(db *gorm.DB, articleID, scoreItemID uuid.UUID) error {
	// Verificar se já existe
	var count int64
	err := db.Table("article_score_items").
		Where("article_id = ? AND score_item_id = ?", articleID, scoreItemID).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return nil // Já existe
	}

	// Criar link usando GORM Association
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, scoreItemID).Error; err != nil {
		return fmt.Errorf("score item not found: %w", err)
	}

	var article models.Article
	if err := db.First(&article, articleID).Error; err != nil {
		return fmt.Errorf("article not found: %w", err)
	}

	// Adicionar associação
	return db.Model(&scoreItem).Association("Articles").Append(&article)
}

func getPoints(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

func stringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func formatLastReview(t *time.Time) string {
	if t == nil {
		return "nunca"
	}
	return t.Format("2006-01-02")
}
