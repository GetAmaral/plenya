package main

import (
	"fmt"
	"log"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// Report-only script: Shows all pending enrichments without modifying data
// Usage: go run report_pending_enrichment.go

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	db := database.DB
	enrichmentService := services.NewScoreItemEnrichmentService(db)

	// ============================================================
	// STATISTICS
	// ============================================================
	var totalItems int64
	var pendingItems int64
	var reviewedItems int64

	db.Model(&models.ScoreItem{}).Count(&totalItems)
	db.Model(&models.ScoreItem{}).Where("last_review IS NULL").Count(&pendingItems)
	db.Model(&models.ScoreItem{}).Where("last_review IS NOT NULL").Count(&reviewedItems)

	fmt.Printf("\n")
	fmt.Printf("═══════════════════════════════════════════════════════\n")
	fmt.Printf("📊 SCORE ITEM ENRICHMENT REPORT\n")
	fmt.Printf("═══════════════════════════════════════════════════════\n")
	fmt.Printf("Total Items:       %d\n", totalItems)
	fmt.Printf("Reviewed:          %d (%.1f%%)\n", reviewedItems, float64(reviewedItems)/float64(totalItems)*100)
	fmt.Printf("Pending:           %d (%.1f%%)\n", pendingItems, float64(pendingItems)/float64(totalItems)*100)
	fmt.Printf("Generated:         %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("═══════════════════════════════════════════════════════\n\n")

	// ============================================================
	// TIER BREAKDOWN
	// ============================================================
	fmt.Printf("🎯 ENRICHMENT TIER BREAKDOWN (Pending Items)\n")
	fmt.Printf("───────────────────────────────────────────────────────\n")

	var items []models.ScoreItem
	db.Preload("Subgroup.Group").
		Where("last_review IS NULL").
		Order("name ASC").
		Find(&items)

	tierCounts := map[services.EnrichmentTier]int{
		services.TierPreserve: 0,
		services.TierEnrich:   0,
		services.TierRewrite:  0,
	}

	for _, item := range items {
		tier := enrichmentService.DetermineTier(&item)
		tierCounts[tier]++
	}

	fmt.Printf("PRESERVE (excellent):  %4d items\n", tierCounts[services.TierPreserve])
	fmt.Printf("ENRICH (improve):      %4d items\n", tierCounts[services.TierEnrich])
	fmt.Printf("REWRITE (from scratch):%4d items\n", tierCounts[services.TierRewrite])
	fmt.Printf("\n")

	// ============================================================
	// BY GROUP
	// ============================================================
	fmt.Printf("📁 PENDING ITEMS BY GROUP\n")
	fmt.Printf("───────────────────────────────────────────────────────\n")

	type GroupStats struct {
		Name    string
		Pending int
	}

	var groupStats []GroupStats
	err = db.Raw(`
		SELECT
			sg.name,
			COUNT(si.id) as pending
		FROM score_items si
		JOIN score_subgroups ss ON si.subgroup_id = ss.id
		JOIN score_groups sg ON ss.group_id = sg.id
		WHERE si.last_review IS NULL
		GROUP BY sg.name
		ORDER BY pending DESC, sg.name ASC
	`).Scan(&groupStats).Error

	if err != nil {
		log.Fatalf("Failed to fetch group stats: %v", err)
	}

	for _, stat := range groupStats {
		bar := progressBar(stat.Pending, int(pendingItems))
		fmt.Printf("%-30s %4d items %s\n", truncate(stat.Name, 30), stat.Pending, bar)
	}

	// ============================================================
	// DETAILED LIST (First 20)
	// ============================================================
	fmt.Printf("\n📋 FIRST 20 PENDING ITEMS (sorted alphabetically)\n")
	fmt.Printf("───────────────────────────────────────────────────────\n")
	fmt.Printf("%-4s %-35s %-10s %-20s\n", "#", "Name", "Tier", "Group")
	fmt.Printf("───────────────────────────────────────────────────────\n")

	limit := 20
	if len(items) < limit {
		limit = len(items)
	}

	for i := 0; i < limit; i++ {
		item := items[i]
		tier := enrichmentService.DetermineTier(&item)
		groupName := ""
		if item.Subgroup != nil && item.Subgroup.Group != nil {
			groupName = item.Subgroup.Group.Name
		}

		tierIcon := getTierIcon(tier)
		fmt.Printf("%-4d %-35s %s %-9s %-20s\n",
			i,
			truncate(item.Name, 35),
			tierIcon,
			tier,
			truncate(groupName, 20))
	}

	// ============================================================
	// USAGE INSTRUCTIONS
	// ============================================================
	fmt.Printf("\n")
	fmt.Printf("═══════════════════════════════════════════════════════\n")
	fmt.Printf("💡 NEXT STEPS\n")
	fmt.Printf("═══════════════════════════════════════════════════════\n")
	fmt.Printf("\n1️⃣  Review single item (dry-run):\n")
	fmt.Printf("   go run enrich_by_offset.go -offset 0 -dry-run\n\n")
	fmt.Printf("2️⃣  Enrich single item:\n")
	fmt.Printf("   go run enrich_by_offset.go -offset 0\n\n")
	fmt.Printf("3️⃣  Batch process (manual loop):\n")
	fmt.Printf("   for i in {0..9}; do\n")
	fmt.Printf("     go run enrich_by_offset.go -offset $i\n")
	fmt.Printf("     sleep 5\n")
	fmt.Printf("   done\n\n")
	fmt.Printf("4️⃣  Re-run this report:\n")
	fmt.Printf("   go run report_pending_enrichment.go\n")
	fmt.Printf("═══════════════════════════════════════════════════════\n\n")
}

// Helper functions

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func progressBar(current, total int) string {
	if total == 0 {
		return ""
	}
	percentage := float64(current) / float64(total)
	barLength := 20
	filled := int(percentage * float64(barLength))

	bar := "["
	for i := 0; i < barLength; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	bar += fmt.Sprintf("] %.1f%%", percentage*100)
	return bar
}

func getTierIcon(tier services.EnrichmentTier) string {
	switch tier {
	case services.TierPreserve:
		return "✨"
	case services.TierEnrich:
		return "📝"
	case services.TierRewrite:
		return "🔨"
	default:
		return "❓"
	}
}
