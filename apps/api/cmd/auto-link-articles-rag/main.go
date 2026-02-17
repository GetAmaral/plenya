package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Auto-Link Articles Script v2 - Conecta artigos a ScoreItems via chunks RAG (melhorado)
//
// MELHORIAS vs v1:
// - Usa FindTopChunksForScoreItem() em vez de FindSimilarArticlesForScoreItem()
// - Threshold adaptativo baseado nas características do score item
// - Requer ≥2 chunks por artigo OU 1 chunk muito similar (≥0.85) para aumentar confiança
// - Score item embeddings JÁ usam fullName (Group - Subgroup - Parent - Name)
//
// Uso: go run cmd/auto-link-articles-rag/main.go [--dry-run]

func main() {
	fmt.Println("🔗 Auto-Link Articles via RAG v2 (Chunk-Based) - Plenya System")
	fmt.Println(string(make([]byte, 70)) + "=")

	// Configuração
	dryRun := false // Se true, apenas simula sem criar links

	// Parse flags simples (opcional)
	// TODO: Adicionar flag parsing se necessário

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Conectar ao banco
	dsn := cfg.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database\n")

	// Criar repository vetorial
	vectorRepo := repository.NewArticleVectorRepository(db)

	// Buscar todos ScoreItems ativos com relationships (para fullName)
	var scoreItems []models.ScoreItem
	if err := db.Preload("Subgroup.Group").Preload("ParentItem").Find(&scoreItems).Error; err != nil {
		log.Fatalf("❌ Failed to fetch score items: %v", err)
	}

	fmt.Printf("📊 Found %d score items to process\n", len(scoreItems))
	if dryRun {
		fmt.Println("🧪 DRY RUN MODE - No links will be created")
	}
	fmt.Println()

	// Estatísticas
	totalLinksCreated := 0
	itemsWithNewLinks := 0
	itemsSkipped := 0
	itemsWithErrors := 0
	thresholdStats := make(map[string]int) // Contar por threshold usado

	// Processar cada ScoreItem
	for i, item := range scoreItems {
		fullName := item.GetFullName()
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(scoreItems), fullName)

		// 1. Determinar threshold adaptativo baseado nas características do item
		threshold := determineThreshold(&item)
		fmt.Printf("   ⚙️  Adaptive threshold: %.2f (based on specificity)\n", threshold)

		thresholdKey := fmt.Sprintf("%.2f", threshold)
		thresholdStats[thresholdKey]++

		// 2. Buscar chunks via RAG (30 chunks, threshold adaptativo)
		chunks, err := vectorRepo.FindTopChunksForScoreItem(
			item.ID,
			30, // Buscar mais chunks para ter melhor cobertura
			threshold,
		)

		if err != nil {
			fmt.Printf("   ❌ Error finding chunks: %v\n", err)
			itemsWithErrors++
			continue
		}

		if len(chunks) == 0 {
			fmt.Printf("   ⚠️  No chunks found (threshold=%.2f)\n", threshold)
			itemsSkipped++
			continue
		}

		fmt.Printf("   🔍 Found %d chunks from articles (similarity ≥ %.2f)\n", len(chunks), threshold)

		// 3. Agrupar chunks por artigo e calcular estatísticas
		articleChunks := make(map[uuid.UUID][]repository.ChunkSearchResult)
		for _, chunk := range chunks {
			articleChunks[chunk.ArticleID] = append(articleChunks[chunk.ArticleID], chunk)
		}

		fmt.Printf("   📚 Chunks distributed across %d unique articles\n", len(articleChunks))

		// 4. Filtrar artigos com confiança suficiente:
		//    - ≥2 chunks relevantes (múltiplas evidências) OU
		//    - 1 chunk com similaridade muito alta (≥0.85)
		qualifiedArticles := make(map[uuid.UUID]float64) // articleID -> avgSimilarity

		for articleID, articleChunkList := range articleChunks {
			chunkCount := len(articleChunkList)

			// Calcular max e avg similarity
			var totalSim float64
			maxSim := 0.0
			for _, chunk := range articleChunkList {
				totalSim += chunk.Similarity
				if chunk.Similarity > maxSim {
					maxSim = chunk.Similarity
				}
			}
			avgSim := totalSim / float64(chunkCount)

			// Critério de qualificação
			if chunkCount >= 2 || maxSim >= 0.85 {
				qualifiedArticles[articleID] = avgSim
			}
		}

		if len(qualifiedArticles) == 0 {
			fmt.Printf("   ⚠️  No articles meet confidence criteria (≥2 chunks or sim≥0.85)\n")
			itemsSkipped++
			continue
		}

		fmt.Printf("   ✓ %d articles qualified (meet confidence criteria)\n", len(qualifiedArticles))

		// 5. Buscar artigos JÁ linkados a este ScoreItem
		var existingLinks []struct {
			ArticleID uuid.UUID
		}
		err = db.Raw(`
			SELECT article_id
			FROM article_score_items
			WHERE score_item_id = ?
		`, item.ID).Scan(&existingLinks).Error

		if err != nil {
			fmt.Printf("   ❌ Error checking existing links: %v\n", err)
			itemsWithErrors++
			continue
		}

		existingIDs := make(map[uuid.UUID]bool)
		for _, link := range existingLinks {
			existingIDs[link.ArticleID] = true
		}

		// 6. Filtrar artigos novos (não linkados)
		newArticles := make(map[uuid.UUID]float64)
		for articleID, avgSim := range qualifiedArticles {
			if !existingIDs[articleID] {
				newArticles[articleID] = avgSim
			}
		}

		if len(newArticles) == 0 {
			fmt.Printf("   ✓ All qualified articles already linked (%d existing)\n", len(existingLinks))
			itemsSkipped++
			continue
		}

		fmt.Printf("   📎 Will create %d new links (%d already exist)\n", len(newArticles), len(existingLinks))

		// 7. Criar links (ou apenas simular se dry-run)
		if !dryRun {
			linksCreated := 0
			for articleID, avgSim := range newArticles {
				// Insert direto SQL (tabela join GORM)
				err := db.Exec(`
					INSERT INTO article_score_items
					(score_item_id, article_id, confidence_score, auto_linked, linked_at, linked_by)
					VALUES (?, ?, ?, true, NOW(), NULL)
					ON CONFLICT (score_item_id, article_id) DO NOTHING
				`, item.ID, articleID, avgSim).Error

				if err != nil {
					fmt.Printf("   ⚠️  Failed to link article %s: %v\n",
						articleID.String()[:8], err)
					continue
				}

				linksCreated++

				// Mostrar detalhes do link criado
				chunkCount := len(articleChunks[articleID])
				fmt.Printf("      ✓ Linked article %s (chunks: %d, avg_sim: %.3f)\n",
					articleID.String()[:8], chunkCount, avgSim)
			}

			if linksCreated > 0 {
				fmt.Printf("   ✅ Created %d links\n", linksCreated)
				totalLinksCreated += linksCreated
				itemsWithNewLinks++
			}
		} else {
			// Dry run - apenas mostrar
			for articleID, avgSim := range newArticles {
				chunkCount := len(articleChunks[articleID])
				fmt.Printf("      [DRY RUN] Would link article %s (chunks: %d, avg_sim: %.3f)\n",
					articleID.String()[:8], chunkCount, avgSim)
			}
			totalLinksCreated += len(newArticles)
			itemsWithNewLinks++
		}

		fmt.Println()
	}

	// Relatório final
	fmt.Println(string(make([]byte, 70)) + "=")
	fmt.Println("📊 Auto-Linking Summary:")
	fmt.Printf("   Total ScoreItems processed: %d\n", len(scoreItems))
	fmt.Printf("   Items with new links: %d\n", itemsWithNewLinks)
	fmt.Printf("   Items skipped (no new articles): %d\n", itemsSkipped)
	fmt.Printf("   Items with errors: %d\n", itemsWithErrors)
	fmt.Printf("   Total links created: %d\n", totalLinksCreated)
	if dryRun {
		fmt.Println("   ⚠️  DRY RUN - No changes were saved")
	}
	fmt.Println()

	// Mostrar distribuição de thresholds usados
	fmt.Println("📈 Threshold Distribution:")
	for threshold, count := range thresholdStats {
		fmt.Printf("   %.2s: %d items (%.1f%%)\n",
			threshold, count, float64(count)/float64(len(scoreItems))*100)
	}
	fmt.Println()

	// Estatísticas finais do banco
	if !dryRun {
		var stats struct {
			TotalLinks     int64
			AutoLinked     int64
			ManualLinked   int64
			AvgConfidence  float64
			ItemsWithLinks int64
		}

		db.Raw(`
			SELECT
				COUNT(*) as total_links,
				COUNT(*) FILTER (WHERE auto_linked = true) as auto_linked,
				COUNT(*) FILTER (WHERE auto_linked = false) as manual_linked,
				AVG(confidence_score) as avg_confidence,
				COUNT(DISTINCT score_item_id) as items_with_links
			FROM article_score_items
		`).Scan(&stats)

		fmt.Println("📈 Database Statistics:")
		fmt.Printf("   Total article-item links: %d\n", stats.TotalLinks)
		fmt.Printf("   Auto-linked: %d (%.1f%%)\n",
			stats.AutoLinked,
			float64(stats.AutoLinked)/float64(stats.TotalLinks)*100)
		fmt.Printf("   Manual-linked: %d (%.1f%%)\n",
			stats.ManualLinked,
			float64(stats.ManualLinked)/float64(stats.TotalLinks)*100)
		fmt.Printf("   Average confidence: %.3f\n", stats.AvgConfidence)
		fmt.Printf("   ScoreItems with articles: %d/%d (%.1f%%)\n",
			stats.ItemsWithLinks,
			len(scoreItems),
			float64(stats.ItemsWithLinks)/float64(len(scoreItems))*100)
		fmt.Println()
	}

	fmt.Println("✅ Auto-linking completed!")
	if !dryRun {
		fmt.Println("\n💡 Next steps:")
		fmt.Println("   1. Review coverage: SELECT COUNT(*) FROM article_score_items WHERE auto_linked = true;")
		fmt.Println("   2. Check confidence distribution: SELECT AVG(confidence_score), MIN(confidence_score) FROM article_score_items;")
		fmt.Println("   3. Validate items with <3 articles (may need PubMed search)")
	}
	fmt.Println("\n🎯 Key improvements in v2:")
	fmt.Println("   - Chunk-based matching (multiple evidences per article)")
	fmt.Println("   - Adaptive thresholds (0.65-0.75 based on specificity)")
	fmt.Println("   - Higher confidence (≥2 chunks or sim≥0.85 required)")
	fmt.Println("   - Uses fullName (Group-Subgroup-Parent-Name) for better context")
}

// determineThreshold calcula threshold adaptativo baseado nas características do score item
// Lógica:
// - Score items ESPECÍFICOS (nome curto + tem unidade + bem definido) → threshold ALTO (0.75)
// - Score items GENÉRICOS (nome longo + sem unidade) → threshold MODERADO (0.65)
// - Padrão → 0.70
func determineThreshold(item *models.ScoreItem) float64 {
	nameLen := len(item.Name)
	hasUnit := item.Unit != nil && *item.Unit != ""
	hasClinicalRelevance := item.ClinicalRelevance != nil && len(*item.ClinicalRelevance) > 500

	// CATEGORIA 1: Score items muito específicos (ex: "Hemoglobina Glicada (HbA1c)", "Vitamina D")
	// - Nome curto (<= 30 chars)
	// - Tem unidade definida
	// - Tem clinical relevance bem desenvolvido
	// → Threshold ALTO para evitar falsos positivos
	if nameLen <= 30 && hasUnit && hasClinicalRelevance {
		return 0.75
	}

	// CATEGORIA 2: Score items genéricos ou procedimentos (ex: "Endoscopia digestiva alta", "Exame Físico Completo")
	// - Nome longo (> 50 chars) ou sem unidade
	// → Threshold MODERADO para capturar mais contextos
	if nameLen > 50 || !hasUnit {
		return 0.65
	}

	// CATEGORIA 3: Padrão (maioria dos casos)
	return 0.70
}
