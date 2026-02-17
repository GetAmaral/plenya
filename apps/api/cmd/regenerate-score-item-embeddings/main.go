package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Regenerate ScoreItem Embeddings - Apaga e re-gera embeddings usando FullName
// Necessário após mudança para incluir contexto hierárquico (Group/Subgroup/Parent)
//
// Uso: go run cmd/regenerate-score-item-embeddings/main.go

func main() {
	fmt.Println("🔄 Regenerate ScoreItem Embeddings - Plenya RAG System")
	fmt.Println("=" + string(make([]byte, 70)))
	fmt.Println("⚠️  This will DELETE all existing ScoreItem embeddings and re-queue")
	fmt.Println("=" + string(make([]byte, 70)))
	fmt.Println()

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

	// Estatísticas iniciais
	var oldEmbeddingCount int64
	db.Model(&models.ScoreItemEmbedding{}).Count(&oldEmbeddingCount)
	fmt.Printf("📊 Current ScoreItem embeddings: %d\n\n", oldEmbeddingCount)

	// STEP 1: Apagar embeddings existentes
	fmt.Println("🗑️  Deleting existing ScoreItem embeddings...")
	result := db.Exec("DELETE FROM score_item_embeddings")
	if result.Error != nil {
		log.Fatalf("❌ Failed to delete embeddings: %v", result.Error)
	}
	fmt.Printf("✅ Deleted %d embeddings\n\n", result.RowsAffected)

	// STEP 2: Apagar jobs pendentes/processing de score_items na fila
	fmt.Println("🗑️  Cleaning up old queue jobs...")
	result = db.Exec(`
		DELETE FROM embedding_queue
		WHERE entity_type = 'score_item'
		AND status IN ('pending', 'processing', 'failed')
	`)
	if result.Error != nil {
		log.Fatalf("❌ Failed to clean queue: %v", result.Error)
	}
	fmt.Printf("✅ Cleaned %d queue jobs\n\n", result.RowsAffected)

	// STEP 3: Re-queue todos os ScoreItems
	fmt.Println("📥 Queueing all ScoreItems for re-embedding...")
	queueService := services.NewEmbeddingQueueService(db)
	count, err := queueService.QueueAllScoreItems()
	if err != nil {
		log.Fatalf("❌ Failed to queue score items: %v", err)
	}
	fmt.Printf("✅ Queued %d ScoreItems\n\n", count)

	// Estatísticas finais
	stats, err := queueService.GetStats()
	if err != nil {
		log.Fatalf("❌ Failed to get queue stats: %v", err)
	}

	fmt.Println("📊 Queue Statistics:")
	fmt.Printf("   Total jobs: %d\n", stats["total"])
	fmt.Printf("   Pending: %d\n", stats["pending"])
	fmt.Printf("   Score Items: %d\n\n", stats["score_items"])

	fmt.Println("✅ Regeneration queued successfully!")
	fmt.Println("\n💡 Next steps:")
	fmt.Println("   1. The API server's EmbeddingWorker will process the queue automatically")
	fmt.Println("   2. Monitor progress: docker compose logs -f api | grep ScoreItem")
	fmt.Println("   3. Check completion: SELECT COUNT(*) FROM score_item_embeddings;")
	fmt.Println("   4. After completion, re-run auto-link script:")
	fmt.Println("      docker compose exec api go run cmd/auto-link-articles-rag/main.go")
	fmt.Println()
	fmt.Println("📈 Expected improvements:")
	fmt.Println("   - Better semantic matching (embeddings now include Group/Subgroup/Parent context)")
	fmt.Println("   - More accurate article recommendations")
	fmt.Println("   - Reduced false positives in auto-linking")
}
