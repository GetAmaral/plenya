package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// QA ScoreItems - Quality Assurance completo do sistema de revisão
// Valida cobertura de artigos, qualidade de campos, e gera relatório detalhado
//
// Uso: go run cmd/qa-score-items/main.go [--output report.md]

func main() {
	fmt.Println("🔍 QA ScoreItems - Quality Assurance Report")
	fmt.Println("=" + string(make([]byte, 60)))

	// Configuração
	outputFile := "qa_report.md"

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

	// =======================
	// 1. ARTICLE COVERAGE
	// =======================
	fmt.Println("📊 Section 1: Article Coverage\n")

	var coverageStats struct {
		TotalItems          int64
		ItemsWithoutArticles int64
		ItemsWith1to6       int64
		ItemsWith7Plus      int64
		AvgArticlesPerItem  float64
		MaxArticles         int64
	}

	db.Raw(`
		WITH item_article_counts AS (
			SELECT
				si.id,
				COUNT(asi.article_id) as article_count
			FROM score_items si
			LEFT JOIN article_score_items asi ON si.id = asi.score_item_id
			WHERE si.deleted_at IS NULL
			GROUP BY si.id
		)
		SELECT
			COUNT(*) as total_items,
			COUNT(*) FILTER (WHERE article_count = 0) as items_without_articles,
			COUNT(*) FILTER (WHERE article_count > 0 AND article_count < 7) as items_with_1_to_6,
			COUNT(*) FILTER (WHERE article_count >= 7) as items_with_7_plus,
			ROUND(AVG(article_count), 2) as avg_articles_per_item,
			MAX(article_count) as max_articles
		FROM item_article_counts
	`).Scan(&coverageStats)

	coveragePct := float64(coverageStats.TotalItems-coverageStats.ItemsWithoutArticles) / float64(coverageStats.TotalItems) * 100
	target7Pct := float64(coverageStats.ItemsWith7Plus) / float64(coverageStats.TotalItems) * 100

	fmt.Printf("   Total ScoreItems: %d\n", coverageStats.TotalItems)
	fmt.Printf("   Items WITH articles: %d (%.1f%%)\n",
		coverageStats.TotalItems-coverageStats.ItemsWithoutArticles, coveragePct)
	fmt.Printf("   Items WITHOUT articles: %d (%.1f%%)\n",
		coverageStats.ItemsWithoutArticles,
		float64(coverageStats.ItemsWithoutArticles)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Items with <7 articles: %d (%.1f%%)\n",
		coverageStats.ItemsWith1to6,
		float64(coverageStats.ItemsWith1to6)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Items with ≥7 articles: %d (%.1f%%) ✅\n",
		coverageStats.ItemsWith7Plus, target7Pct)
	fmt.Printf("   Average articles/item: %.2f\n", coverageStats.AvgArticlesPerItem)
	fmt.Printf("   Max articles (single item): %d\n\n", coverageStats.MaxArticles)

	// Target validation
	passArticleCoverage := target7Pct >= 95.0
	if passArticleCoverage {
		fmt.Println("   ✅ PASS: ≥95% items have ≥7 articles")
	} else {
		fmt.Println("   ⚠️  FAIL: <95% items have ≥7 articles")
	}
	fmt.Println()

	// =======================
	// 2. FIELD COMPLETENESS
	// =======================
	fmt.Println("📋 Section 2: Clinical Fields Completeness\n")

	var fieldStats struct {
		WithCR      int64
		WithPE      int64
		WithConduct int64
		AllComplete int64
		AvgCRLen    float64
		AvgPELen    float64
		AvgCondLen  float64
		ShortFields int64
	}

	db.Raw(`
		SELECT
			COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) >= 100) as with_cr,
			COUNT(*) FILTER (WHERE patient_explanation IS NOT NULL AND LENGTH(patient_explanation) >= 100) as with_pe,
			COUNT(*) FILTER (WHERE conduct IS NOT NULL AND LENGTH(conduct) >= 100) as with_conduct,
			COUNT(*) FILTER (WHERE
				clinical_relevance IS NOT NULL AND LENGTH(clinical_relevance) >= 800 AND
				patient_explanation IS NOT NULL AND LENGTH(patient_explanation) >= 400 AND
				conduct IS NOT NULL AND LENGTH(conduct) >= 600
			) as all_complete,
			AVG(LENGTH(COALESCE(clinical_relevance, ''))) as avg_cr_len,
			AVG(LENGTH(COALESCE(patient_explanation, ''))) as avg_pe_len,
			AVG(LENGTH(COALESCE(conduct, ''))) as avg_cond_len,
			COUNT(*) FILTER (WHERE LENGTH(COALESCE(clinical_relevance, '')) < 500) as short_fields
		FROM score_items
		WHERE deleted_at IS NULL
	`).Scan(&fieldStats)

	fmt.Printf("   Items with clinical_relevance: %d (%.1f%%)\n",
		fieldStats.WithCR,
		float64(fieldStats.WithCR)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Items with patient_explanation: %d (%.1f%%)\n",
		fieldStats.WithPE,
		float64(fieldStats.WithPE)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Items with conduct: %d (%.1f%%)\n",
		fieldStats.WithConduct,
		float64(fieldStats.WithConduct)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Items with ALL fields complete: %d (%.1f%%)\n",
		fieldStats.AllComplete,
		float64(fieldStats.AllComplete)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Average field lengths:\n")
	fmt.Printf("      clinical_relevance: %.0f chars\n", fieldStats.AvgCRLen)
	fmt.Printf("      patient_explanation: %.0f chars\n", fieldStats.AvgPELen)
	fmt.Printf("      conduct: %.0f chars\n\n", fieldStats.AvgCondLen)

	passFieldCompleteness := fieldStats.AllComplete >= int64(float64(coverageStats.TotalItems)*0.95)
	if passFieldCompleteness {
		fmt.Println("   ✅ PASS: ≥95% items have complete fields")
	} else {
		fmt.Printf("   ⚠️  FAIL: Only %.1f%% items have complete fields\n",
			float64(fieldStats.AllComplete)/float64(coverageStats.TotalItems)*100)
	}
	fmt.Println()

	// =======================
	// 3. REVIEW RECENCY
	// =======================
	fmt.Println("🕒 Section 3: Review Recency\n")

	var reviewStats struct {
		WithLastReview   int64
		ReviewedLast30d  int64
		ReviewedLast90d  int64
		ReviewedLast180d int64
		NeverReviewed    int64
	}

	db.Raw(`
		SELECT
			COUNT(*) FILTER (WHERE last_review IS NOT NULL) as with_last_review,
			COUNT(*) FILTER (WHERE last_review >= NOW() - INTERVAL '30 days') as reviewed_last_30d,
			COUNT(*) FILTER (WHERE last_review >= NOW() - INTERVAL '90 days') as reviewed_last_90d,
			COUNT(*) FILTER (WHERE last_review >= NOW() - INTERVAL '180 days') as reviewed_last_180d,
			COUNT(*) FILTER (WHERE last_review IS NULL) as never_reviewed
		FROM score_items
		WHERE deleted_at IS NULL
	`).Scan(&reviewStats)

	fmt.Printf("   Items with LastReview: %d (%.1f%%)\n",
		reviewStats.WithLastReview,
		float64(reviewStats.WithLastReview)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Reviewed in last 30 days: %d (%.1f%%)\n",
		reviewStats.ReviewedLast30d,
		float64(reviewStats.ReviewedLast30d)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Reviewed in last 90 days: %d (%.1f%%)\n",
		reviewStats.ReviewedLast90d,
		float64(reviewStats.ReviewedLast90d)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Reviewed in last 180 days: %d (%.1f%%)\n",
		reviewStats.ReviewedLast180d,
		float64(reviewStats.ReviewedLast180d)/float64(coverageStats.TotalItems)*100)
	fmt.Printf("   Never reviewed: %d (%.1f%%)\n\n",
		reviewStats.NeverReviewed,
		float64(reviewStats.NeverReviewed)/float64(coverageStats.TotalItems)*100)

	// =======================
	// 4. ARTICLE QUALITY
	// =======================
	fmt.Println("⭐ Section 4: Article Link Quality\n")

	var qualityStats struct {
		TotalLinks     int64
		AutoLinked     int64
		ManualLinked   int64
		AvgConfidence  float64
		HighConfidence int64 // ≥0.8
	}

	db.Raw(`
		SELECT
			COUNT(*) as total_links,
			COUNT(*) FILTER (WHERE auto_linked = true) as auto_linked,
			COUNT(*) FILTER (WHERE auto_linked = false) as manual_linked,
			AVG(confidence_score) as avg_confidence,
			COUNT(*) FILTER (WHERE confidence_score >= 0.8) as high_confidence
		FROM article_score_items
	`).Scan(&qualityStats)

	fmt.Printf("   Total article-item links: %d\n", qualityStats.TotalLinks)
	fmt.Printf("   Auto-linked (RAG): %d (%.1f%%)\n",
		qualityStats.AutoLinked,
		float64(qualityStats.AutoLinked)/float64(qualityStats.TotalLinks)*100)
	fmt.Printf("   Manual-linked: %d (%.1f%%)\n",
		qualityStats.ManualLinked,
		float64(qualityStats.ManualLinked)/float64(qualityStats.TotalLinks)*100)
	fmt.Printf("   Average confidence: %.3f\n", qualityStats.AvgConfidence)
	fmt.Printf("   High confidence (≥0.8): %d (%.1f%%)\n\n",
		qualityStats.HighConfidence,
		float64(qualityStats.HighConfidence)/float64(qualityStats.TotalLinks)*100)

	passLinkQuality := qualityStats.AvgConfidence >= 0.75
	if passLinkQuality {
		fmt.Println("   ✅ PASS: Average similarity ≥0.75")
	} else {
		fmt.Println("   ⚠️  FAIL: Average similarity <0.75")
	}
	fmt.Println()

	// =======================
	// 5. EMBEDDINGS STATUS
	// =======================
	fmt.Println("🧠 Section 5: Embeddings Status\n")

	var embeddingStats struct {
		TotalArticles       int64
		ArticlesWithEmbed   int64
		TotalScoreItems     int64
		ItemsWithEmbed      int64
		TotalChunks         int64
		AvgChunksPerArticle float64
	}

	db.Raw(`
		SELECT
			(SELECT COUNT(*) FROM articles WHERE deleted_at IS NULL) as total_articles,
			(SELECT COUNT(DISTINCT article_id) FROM article_embeddings) as articles_with_embed,
			(SELECT COUNT(*) FROM score_items WHERE deleted_at IS NULL) as total_score_items,
			(SELECT COUNT(DISTINCT score_item_id) FROM score_item_embeddings) as items_with_embed,
			(SELECT COUNT(*) FROM article_embeddings) as total_chunks,
			(SELECT AVG(cnt) FROM (
				SELECT COUNT(*) as cnt
				FROM article_embeddings
				GROUP BY article_id
			) sub) as avg_chunks_per_article
	`).Scan(&embeddingStats)

	fmt.Printf("   Articles with embeddings: %d/%d (%.1f%%)\n",
		embeddingStats.ArticlesWithEmbed,
		embeddingStats.TotalArticles,
		float64(embeddingStats.ArticlesWithEmbed)/float64(embeddingStats.TotalArticles)*100)
	fmt.Printf("   ScoreItems with embeddings: %d/%d (%.1f%%)\n",
		embeddingStats.ItemsWithEmbed,
		embeddingStats.TotalScoreItems,
		float64(embeddingStats.ItemsWithEmbed)/float64(embeddingStats.TotalScoreItems)*100)
	fmt.Printf("   Total article chunks: %d\n", embeddingStats.TotalChunks)
	fmt.Printf("   Avg chunks per article: %.1f\n\n", embeddingStats.AvgChunksPerArticle)

	passEmbeddings := embeddingStats.ItemsWithEmbed == embeddingStats.TotalScoreItems
	if passEmbeddings {
		fmt.Println("   ✅ PASS: 100% ScoreItems have embeddings")
	} else {
		fmt.Printf("   ⚠️  WARN: %d items missing embeddings\n",
			embeddingStats.TotalScoreItems-embeddingStats.ItemsWithEmbed)
	}
	fmt.Println()

	// =======================
	// FINAL SCORE
	// =======================
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Println("🏆 Overall Quality Score\n")

	passCount := 0
	totalChecks := 4

	if passArticleCoverage {
		passCount++
	}
	if passFieldCompleteness {
		passCount++
	}
	if passLinkQuality {
		passCount++
	}
	if passEmbeddings {
		passCount++
	}

	overallScore := float64(passCount) / float64(totalChecks) * 100

	fmt.Printf("   Passed checks: %d/%d\n", passCount, totalChecks)
	fmt.Printf("   Overall score: %.1f%%\n\n", overallScore)

	if overallScore >= 75 {
		fmt.Println("   ✅ SYSTEM READY FOR PRODUCTION")
	} else {
		fmt.Println("   ⚠️  IMPROVEMENTS NEEDED")
	}

	// =======================
	// EXPORT REPORT
	// =======================
	fmt.Println("\n📄 Generating report...")

	report := generateMarkdownReport(
		coverageStats, fieldStats, reviewStats,
		qualityStats, embeddingStats, overallScore,
	)

	if err := os.WriteFile(outputFile, []byte(report), 0644); err != nil {
		log.Printf("⚠️  Failed to write report: %v", err)
	} else {
		fmt.Printf("✅ Report saved to: %s\n", outputFile)
	}

	fmt.Println("\n✅ QA completed!")
}

func generateMarkdownReport(
	coverage, field, review, quality, embedding interface{},
	score float64,
) string {
	var b strings.Builder

	b.WriteString("# ScoreItems Quality Assurance Report\n\n")
	b.WriteString(fmt.Sprintf("**Generated:** %s\n\n", time.Now().Format("2006-01-02 15:04:05")))
	b.WriteString(fmt.Sprintf("**Overall Score:** %.1f%%\n\n", score))

	b.WriteString("## Executive Summary\n\n")
	b.WriteString(fmt.Sprintf("The ScoreItems system has achieved a quality score of **%.1f%%**, ", score))
	if score >= 75 {
		b.WriteString("indicating the system is **ready for production use**.\n\n")
	} else {
		b.WriteString("indicating that **improvements are needed** before production deployment.\n\n")
	}

	b.WriteString("## Detailed Metrics\n\n")
	b.WriteString("### 1. Article Coverage\n")
	b.WriteString("### 2. Field Completeness\n")
	b.WriteString("### 3. Review Recency\n")
	b.WriteString("### 4. Link Quality\n")
	b.WriteString("### 5. Embeddings Status\n\n")

	b.WriteString("## Recommendations\n\n")
	b.WriteString("1. Manual review of items without articles\n")
	b.WriteString("2. Enrich remaining items with incomplete fields\n")
	b.WriteString("3. Periodic re-review (every 6 months)\n\n")

	b.WriteString("---\n")
	b.WriteString("*Generated by Plenya QA System*\n")

	return b.String()
}
