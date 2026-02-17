package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Enrich ScoreItems - Enriquece campos clínicos usando LLM baseado em artigos
// Implementa 3-tier preservation strategy + audit trail completo
//
// Uso: go run cmd/enrich-score-items/main.go [--model sonnet|haiku] [--tier preserve|enrich|rewrite|all] [--limit N] [--dry-run]

func main() {
	fmt.Println("✨ Enrich ScoreItems - LLM-Powered Clinical Fields")
	fmt.Println("=" + string(make([]byte, 60)))

	// Configuração
	modelChoice := "sonnet"  // ou "haiku" para economizar
	targetTier := "all"      // ou "enrich", "rewrite"
	limit := 0               // 0 = sem limite
	dryRun := false

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env file not found (using environment variables)")
	}

	// Verificar API key
	apiKey := godotenv.Get("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ ANTHROPIC_API_KEY environment variable is required")
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

	fmt.Println("✅ Connected to database")

	// Criar serviço
	enrichmentService := services.NewScoreItemEnrichmentService(db)

	// Configurar modelo
	if modelChoice == "haiku" {
		enrichmentService.SetModel("claude-haiku-4-5-20251001")
		fmt.Println("⚙️  Using: Claude Haiku (cost-optimized)")
	} else {
		fmt.Println("⚙️  Using: Claude Sonnet 4.5 (high-quality)")
	}

	fmt.Printf("🎯 Target tier: %s\n", targetTier)
	if limit > 0 {
		fmt.Printf("⚙️  Limit: %d items\n", limit)
	}
	if dryRun {
		fmt.Println("🧪 DRY RUN MODE - No changes will be made")
	}
	fmt.Println()

	// Buscar todos ScoreItems
	var allItems []models.ScoreItem
	if err := db.Find(&allItems).Error; err != nil {
		log.Fatalf("❌ Failed to fetch items: %v", err)
	}

	fmt.Printf("📊 Total items: %d\n", len(allItems))

	// Classificar por tier
	tierCounts := map[services.EnrichmentTier]int{
		services.TierPreserve: 0,
		services.TierEnrich:   0,
		services.TierRewrite:  0,
	}

	itemsByTier := make(map[services.EnrichmentTier][]models.ScoreItem)

	for _, item := range allItems {
		tier := enrichmentService.DetermineTier(&item)
		tierCounts[tier]++
		itemsByTier[tier] = append(itemsByTier[tier], item)
	}

	fmt.Println("\n📋 Tier Classification:")
	fmt.Printf("   Tier 1 (Preserve):  %d items - Skip (already excellent)\n", tierCounts[services.TierPreserve])
	fmt.Printf("   Tier 2 (Enrich):    %d items - Improve incrementally\n", tierCounts[services.TierEnrich])
	fmt.Printf("   Tier 3 (Rewrite):   %d items - Generate from scratch\n", tierCounts[services.TierRewrite])
	fmt.Println()

	// Selecionar items a processar
	var itemsToProcess []models.ScoreItem

	switch targetTier {
	case "preserve":
		itemsToProcess = itemsByTier[services.TierPreserve]
	case "enrich":
		itemsToProcess = itemsByTier[services.TierEnrich]
	case "rewrite":
		itemsToProcess = itemsByTier[services.TierRewrite]
	case "all":
		// Processar Enrich + Rewrite (skip Preserve)
		itemsToProcess = append(itemsByTier[services.TierEnrich], itemsByTier[services.TierRewrite]...)
	default:
		log.Fatalf("❌ Invalid tier: %s", targetTier)
	}

	if limit > 0 && len(itemsToProcess) > limit {
		itemsToProcess = itemsToProcess[:limit]
	}

	if len(itemsToProcess) == 0 {
		fmt.Println("✅ No items to process!")
		return
	}

	fmt.Printf("🚀 Processing %d items...\n\n", len(itemsToProcess))

	// Estatísticas
	stats := struct {
		Processed       int
		Success         int
		Failed          int
		Skipped         int
		ValidationFails int
		TotalCost       float64 // Estimativa
	}{}

	ctx := context.Background()

	// Processar cada item
	for i, item := range itemsToProcess {
		tier := enrichmentService.DetermineTier(&item)

		fmt.Printf("[%d/%d] %s\n", i+1, len(itemsToProcess), item.Name)
		fmt.Printf("   Tier: %s\n", tier)

		// Skip se Tier Preserve
		if tier == services.TierPreserve {
			fmt.Println("   ✓ Preserved (already excellent)\n")
			stats.Skipped++
			continue
		}

		if dryRun {
			fmt.Printf("   [DRY RUN] Would enrich with tier=%s\n\n", tier)
			stats.Processed++
			continue
		}

		// Gerar enriquecimento
		startTime := time.Now()
		result, err := enrichmentService.GenerateEnrichment(ctx, &item, tier)
		duration := time.Since(startTime)

		if err != nil {
			fmt.Printf("   ❌ Enrichment failed: %v\n\n", err)
			stats.Failed++
			continue
		}

		fmt.Printf("   ✓ Generated in %.1fs (confidence: %.2f)\n", duration.Seconds(), result.Confidence)

		// Validar resultado
		validationErrors := enrichmentService.ValidateResult(result)
		if len(validationErrors) > 0 {
			fmt.Printf("   ⚠️  Validation failed:\n")
			for _, verr := range validationErrors {
				fmt.Printf("      - %s\n", verr)
			}
			stats.ValidationFails++
			stats.Failed++
			fmt.Println()
			continue
		}

		// Salvar auditoria
		if err := enrichmentService.SaveReviewHistory(&item, result, tier); err != nil {
			fmt.Printf("   ⚠️  Failed to save audit: %v\n", err)
			// Continuar mesmo assim
		}

		// Atualizar item (GORM hook atualizará LastReview automaticamente)
		item.ClinicalRelevance = &result.ClinicalRelevance
		item.PatientExplanation = &result.PatientExplanation
		item.Conduct = &result.Conduct
		points := float64(result.MaxPoints)
		item.Points = &points

		if err := db.Save(&item).Error; err != nil {
			fmt.Printf("   ❌ Failed to save item: %v\n\n", err)
			stats.Failed++
			continue
		}

		fmt.Printf("   ✅ Updated successfully\n")
		fmt.Printf("      CR: %d chars, PE: %d chars, Conduct: %d chars\n\n",
			len(result.ClinicalRelevance),
			len(result.PatientExplanation),
			len(result.Conduct))

		stats.Success++
		stats.Processed++

		// Estimativa de custo (Sonnet: $3/1M in, $15/1M out)
		// Assumindo ~2500 tokens input, ~800 tokens output
		if modelChoice == "sonnet" {
			stats.TotalCost += (2500 * 0.000003) + (800 * 0.000015)
		} else { // haiku
			stats.TotalCost += (2500 * 0.0000008) + (800 * 0.000004)
		}

		// Rate limiting (evitar throttling)
		time.Sleep(500 * time.Millisecond)
	}

	// Relatório final
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("📊 Enrichment Summary:")
	fmt.Printf("   Items processed: %d\n", stats.Processed)
	fmt.Printf("   Success: %d\n", stats.Success)
	fmt.Printf("   Failed: %d\n", stats.Failed)
	fmt.Printf("   Skipped (preserved): %d\n", stats.Skipped)
	fmt.Printf("   Validation failures: %d\n", stats.ValidationFails)

	if !dryRun && stats.Success > 0 {
		fmt.Printf("\n💰 Estimated cost: $%.2f\n", stats.TotalCost)

		// Estatísticas finais do banco
		var dbStats struct {
			TotalReviewed  int64
			ReviewedToday  int64
			AvgCRLength    float64
			AvgPELength    float64
			AvgCondLength  float64
		}

		db.Raw(`
			SELECT
				COUNT(*) FILTER (WHERE last_review IS NOT NULL) as total_reviewed,
				COUNT(*) FILTER (WHERE last_review >= CURRENT_DATE) as reviewed_today,
				AVG(LENGTH(clinical_relevance)) as avg_cr_length,
				AVG(LENGTH(patient_explanation)) as avg_pe_length,
				AVG(LENGTH(conduct)) as avg_cond_length
			FROM score_items
		`).Scan(&dbStats)

		fmt.Println("\n📈 Database Statistics:")
		fmt.Printf("   Items with LastReview: %d/878 (%.1f%%)\n",
			dbStats.TotalReviewed,
			float64(dbStats.TotalReviewed)/878*100)
		fmt.Printf("   Reviewed today: %d\n", dbStats.ReviewedToday)
		fmt.Printf("   Avg field lengths: CR=%.0f, PE=%.0f, Conduct=%.0f\n",
			dbStats.AvgCRLength, dbStats.AvgPELength, dbStats.AvgCondLength)
	}

	fmt.Println("\n✅ Enrichment completed!")

	if !dryRun && stats.Success > 0 {
		fmt.Println("\n💡 Next steps:")
		fmt.Println("   1. Review audit trail: SELECT * FROM score_item_review_history ORDER BY reviewed_at DESC LIMIT 20;")
		fmt.Println("   2. Validate random sample manually")
		fmt.Println("   3. If quality is good, approve and continue with remaining items")
		fmt.Println("   4. Rollback if needed: Use before_snapshot from audit table")
	}
}
