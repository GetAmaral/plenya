package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Backfill Script - Adiciona artigos e score items existentes à fila de embeddings
// Uso: go run cmd/backfill-embeddings/main.go [--articles] [--score-items] [--all]
//
// Flags:
//   --articles     Queue apenas artigos
//   --score-items  Queue apenas score items
//   --all          Queue ambos (default)

func main() {
	fmt.Println("🔄 Backfill Embeddings - Plenya RAG System")
	fmt.Println("=" + string(make([]byte, 60)))

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
		Logger: logger.Default.LogMode(logger.Silent), // Silenciar logs SQL
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database\n")

	// Criar serviço de fila
	queueService := services.NewEmbeddingQueueService(db)

	// Parse flags
	args := os.Args[1:]
	queueArticles := true
	queueScoreItems := true

	if len(args) > 0 {
		queueArticles = false
		queueScoreItems = false

		for _, arg := range args {
			switch arg {
			case "--articles":
				queueArticles = true
			case "--score-items":
				queueScoreItems = true
			case "--all":
				queueArticles = true
				queueScoreItems = true
			default:
				fmt.Printf("⚠️  Unknown flag: %s (ignoring)\n", arg)
			}
		}
	}

	// Queue articles
	if queueArticles {
		fmt.Println("📄 Queueing articles for embedding...")
		count, err := queueService.QueueAllArticles()
		if err != nil {
			log.Fatalf("❌ Failed to queue articles: %v", err)
		}
		fmt.Printf("✅ Queued %d articles\n\n", count)
	}

	// Queue score items
	if queueScoreItems {
		fmt.Println("🎯 Queueing score items for embedding...")
		count, err := queueService.QueueAllScoreItems()
		if err != nil {
			log.Fatalf("❌ Failed to queue score items: %v", err)
		}
		fmt.Printf("✅ Queued %d score items\n\n", count)
	}

	// Mostrar estatísticas finais
	stats, err := queueService.GetStats()
	if err != nil {
		log.Fatalf("❌ Failed to get queue stats: %v", err)
	}

	fmt.Println("📊 Queue Statistics:")
	fmt.Printf("   Total jobs: %d\n", stats["total"])
	fmt.Printf("   Pending: %d\n", stats["pending"])
	fmt.Printf("   Processing: %d\n", stats["processing"])
	fmt.Printf("   Completed: %d\n", stats["completed"])
	fmt.Printf("   Failed: %d\n", stats["failed"])
	fmt.Printf("   Articles: %d\n", stats["articles"])
	fmt.Printf("   Score Items: %d\n\n", stats["score_items"])

	fmt.Println("✅ Backfill completed!")
	fmt.Println("\n💡 Next steps:")
	fmt.Println("   1. Start the API server (EmbeddingWorker will process the queue)")
	fmt.Println("   2. Monitor progress in the logs")
	fmt.Println("   3. Check queue stats: SELECT status, COUNT(*) FROM embedding_queue GROUP BY status;")
}
