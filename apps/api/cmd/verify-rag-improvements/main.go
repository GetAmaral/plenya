package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Verificar melhorias do sistema RAG
func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("🔍 Verificando melhorias do sistema RAG Auto-Link\n")

	// 1. Verificar distribuição de confidence atual
	fmt.Println("📊 1. Distribuição de Confidence Scores (antes das melhorias)")
	var results []struct {
		Bucket      string
		Count       int64
		AvgConf     float64
		Percentage  float64
	}

	db.Raw(`
		WITH total AS (
			SELECT COUNT(*) as total_count
			FROM article_score_items
			WHERE auto_linked = true
		)
		SELECT
			CASE
				WHEN confidence_score >= 0.7 THEN 'high (0.7-1.0)'
				WHEN confidence_score >= 0.5 THEN 'medium (0.5-0.7)'
				WHEN confidence_score >= 0.3 THEN 'low (0.3-0.5)'
				ELSE 'very_low (<0.3)'
			END AS bucket,
			COUNT(*) as count,
			ROUND(AVG(confidence_score)::numeric, 3) as avg_conf,
			ROUND(100.0 * COUNT(*) / (SELECT total_count FROM total), 2) as percentage
		FROM article_score_items
		WHERE auto_linked = true
		GROUP BY bucket
		ORDER BY avg_conf DESC
	`).Scan(&results)

	for _, r := range results {
		fmt.Printf("   %s: %d links (%.1f%%) - avg: %.3f\n",
			r.Bucket, r.Count, r.Percentage, r.AvgConf)
	}

	// 2. Verificar tabelas de infraestrutura
	fmt.Println("\n🏗️  2. Infraestrutura de Staleness Tracking")

	var staleCount, totalEmbeddings int64
	db.Raw("SELECT COUNT(*) FROM score_item_embeddings WHERE is_stale = true").Scan(&staleCount)
	db.Raw("SELECT COUNT(*) FROM score_item_embeddings").Scan(&totalEmbeddings)
	fmt.Printf("   Embeddings stale: %d/%d (%.1f%%)\n",
		staleCount, totalEmbeddings, 100.0*float64(staleCount)/float64(totalEmbeddings))

	var queuePending, queueTotal int64
	db.Raw("SELECT COUNT(*) FROM embedding_queue WHERE status = 'pending'").Scan(&queuePending)
	db.Raw("SELECT COUNT(*) FROM embedding_queue").Scan(&queueTotal)
	fmt.Printf("   Fila de regeneração: %d pending / %d total\n", queuePending, queueTotal)

	var auditCount int64
	db.Raw("SELECT COUNT(*) FROM embedding_audit_log").Scan(&auditCount)
	fmt.Printf("   Audit log: %d entradas\n", auditCount)

	// 3. Verificar tabelas de batch processing
	fmt.Println("\n⚙️  3. Batch Processing & Checkpoints")

	var runCount int64
	db.Raw("SELECT COUNT(*) FROM auto_link_processing_state").Scan(&runCount)
	fmt.Printf("   Runs processados: %d\n", runCount)

	if runCount > 0 {
		var lastRun struct {
			RunID      string
			Status     string
			Processed  int
			Linked     int
			Failed     int
			StartedAt  string
		}
		db.Raw(`
			SELECT
				run_id::text as run_id,
				status,
				total_processed as processed,
				total_linked as linked,
				jsonb_array_length(failed_items) as failed,
				started_at::text
			FROM auto_link_processing_state
			ORDER BY started_at DESC
			LIMIT 1
		`).Scan(&lastRun)

		fmt.Printf("   Último run: %s\n", lastRun.RunID[:8]+"...")
		fmt.Printf("     Status: %s\n", lastRun.Status)
		fmt.Printf("     Processados: %d (linked: %d, failed: %d)\n",
			lastRun.Processed, lastRun.Linked, lastRun.Failed)
	}

	// 4. Verificar feedback mechanism
	fmt.Println("\n📝 4. Feedback Mechanism")

	var feedbackCount, totalLinks int64
	db.Raw("SELECT COUNT(*) FROM article_score_items WHERE user_feedback IS NOT NULL").Scan(&feedbackCount)
	db.Raw("SELECT COUNT(*) FROM article_score_items WHERE auto_linked = true").Scan(&totalLinks)

	fmt.Printf("   Links com feedback: %d/%d (%.1f%%)\n",
		feedbackCount, totalLinks, 100.0*float64(feedbackCount)/float64(totalLinks))

	if feedbackCount > 0 {
		var feedbackBreakdown []struct {
			Feedback string
			Count    int64
		}
		db.Raw(`
			SELECT user_feedback as feedback, COUNT(*) as count
			FROM article_score_items
			WHERE user_feedback IS NOT NULL
			GROUP BY user_feedback
			ORDER BY count DESC
		`).Scan(&feedbackBreakdown)

		for _, fb := range feedbackBreakdown {
			fmt.Printf("     %s: %d\n", fb.Feedback, fb.Count)
		}
	}

	// 5. Verificar funções helper
	fmt.Println("\n🛠️  5. Helper Functions")

	var functions []string
	db.Raw(`
		SELECT proname
		FROM pg_proc
		WHERE proname IN (
			'invalidate_embedding',
			'create_auto_link_run',
			'save_batch_checkpoint',
			'complete_auto_link_run',
			'submit_article_link_feedback'
		)
	`).Scan(&functions)

	for _, fn := range functions {
		fmt.Printf("   ✓ %s\n", fn)
	}

	// 6. Verificar views analytics
	fmt.Println("\n📈 6. Analytics Views")

	var views []string
	db.Raw(`
		SELECT viewname
		FROM pg_views
		WHERE viewname IN (
			'embedding_health_stats',
			'auto_link_run_summary',
			'auto_link_batch_performance',
			'auto_link_failure_analysis',
			'article_link_precision_by_confidence',
			'article_link_problematic_items',
			'article_link_feedback_stats'
		)
		ORDER BY viewname
	`).Scan(&views)

	for _, v := range views {
		fmt.Printf("   ✓ %s\n", v)
	}

	fmt.Println("\n✅ Verificação concluída!")
	fmt.Println("\n💡 Próximos passos:")
	fmt.Println("   1. Executar: docker compose exec api go run /app/cmd/auto-link-all/main.go")
	fmt.Println("   2. Comparar distribuição de confidence (espera-se redução de links <0.5)")
	fmt.Println("   3. Verificar logs de batch processing")
	fmt.Println("   4. Regenerar embeddings: docker compose exec api go run /app/cmd/regenerate-embeddings/main.go")
}
