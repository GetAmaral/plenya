package main

// Re-indexação de Embeddings — Plenya RAG System
//
// Uso: go run cmd/reindex-embeddings/main.go [--dry-run]
//
// O que faz:
//   1. Apaga todos os article_embeddings existentes
//   2. Apaga todos os score_item_embeddings existentes
//   3. Reseta embedding_status dos artigos para "pending"
//   4. Limpa a embedding_queue
//   5. Adiciona todos os artigos e score items à fila novamente
//
// Use após trocar o modelo de embedding (ex: OpenAI → Voyage AI).
// O EmbeddingWorker processará a fila automaticamente quando a API iniciar.

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("♻️  Reindex Embeddings — Plenya RAG System")
	fmt.Println("==========================================")

	dryRun := false
	for _, arg := range os.Args[1:] {
		if arg == "--dry-run" {
			dryRun = true
		}
	}

	if dryRun {
		fmt.Println("🔍 DRY RUN — nenhuma alteração será feita\n")
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env não encontrado (usando variáveis de ambiente)")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Falha ao carregar config: %v", err)
	}

	// Mostrar qual provider está ativo
	if cfg.VoyageAI.APIKey != "" {
		fmt.Printf("🌍 Provider: Voyage AI (%s)\n\n", cfg.VoyageAI.Model)
	} else {
		fmt.Printf("🤖 Provider: OpenAI (%s)\n\n", cfg.OpenAI.EmbeddingModel)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Falha ao conectar ao banco: %v", err)
	}
	fmt.Println("✅ Conectado ao banco\n")

	// ── Estatísticas antes ─────────────────────────────────────────────────────
	var articleEmbCount, scoreEmbCount, queueCount int64
	db.Raw("SELECT COUNT(*) FROM article_embeddings").Scan(&articleEmbCount)
	db.Raw("SELECT COUNT(*) FROM score_item_embeddings").Scan(&scoreEmbCount)
	db.Raw("SELECT COUNT(*) FROM embedding_queue").Scan(&queueCount)

	fmt.Println("📊 Estado atual:")
	fmt.Printf("   article_embeddings:    %d\n", articleEmbCount)
	fmt.Printf("   score_item_embeddings: %d\n", scoreEmbCount)
	fmt.Printf("   embedding_queue:       %d jobs\n\n", queueCount)

	if dryRun {
		fmt.Println("✅ Dry run completo. Use sem --dry-run para executar.")
		return
	}

	// ── 1. Apagar embeddings de artigos ────────────────────────────────────────
	fmt.Print("🗑️  Apagando article_embeddings... ")
	if err := db.Exec("DELETE FROM article_embeddings").Error; err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}
	fmt.Println("✅")

	// ── 2. Apagar embeddings de score items ────────────────────────────────────
	fmt.Print("🗑️  Apagando score_item_embeddings... ")
	if err := db.Exec("DELETE FROM score_item_embeddings").Error; err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}
	fmt.Println("✅")

	// ── 3. Resetar embedding_status dos artigos ────────────────────────────────
	fmt.Print("🔄 Resetando embedding_status dos artigos... ")
	result := db.Exec("UPDATE articles SET embedding_status = 'pending', chunk_count = 0, last_embedded_at = NULL")
	if result.Error != nil {
		log.Fatalf("❌ Erro: %v", result.Error)
	}
	fmt.Printf("✅ (%d artigos)\n", result.RowsAffected)

	// ── 4. Limpar embedding_queue ──────────────────────────────────────────────
	fmt.Print("🗑️  Limpando embedding_queue... ")
	if err := db.Exec("DELETE FROM embedding_queue").Error; err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}
	fmt.Println("✅")

	// ── 5. Re-enfileirar todos os artigos ─────────────────────────────────────
	fmt.Print("📄 Enfileirando artigos... ")
	enqueueResult := db.Exec(`
		INSERT INTO embedding_queue (id, entity_type, entity_id, status, retry_count, created_at)
		SELECT gen_random_uuid(), 'article', id, 'pending', 0, NOW()
		FROM articles
		WHERE deleted_at IS NULL
	`)
	if enqueueResult.Error != nil {
		log.Fatalf("❌ Erro: %v", enqueueResult.Error)
	}
	fmt.Printf("✅ (%d artigos na fila)\n", enqueueResult.RowsAffected)

	// ── 6. Re-enfileirar todos os score items ─────────────────────────────────
	fmt.Print("🎯 Enfileirando score items... ")
	enqueueScoreResult := db.Exec(`
		INSERT INTO embedding_queue (id, entity_type, entity_id, status, retry_count, created_at)
		SELECT gen_random_uuid(), 'score_item', id, 'pending', 0, NOW()
		FROM score_items
		WHERE deleted_at IS NULL
	`)
	if enqueueScoreResult.Error != nil {
		log.Fatalf("❌ Erro: %v", enqueueScoreResult.Error)
	}
	fmt.Printf("✅ (%d score items na fila)\n", enqueueScoreResult.RowsAffected)

	// ── Resumo final ───────────────────────────────────────────────────────────
	var totalQueue int64
	db.Raw("SELECT COUNT(*) FROM embedding_queue WHERE status = 'pending'").Scan(&totalQueue)

	fmt.Printf("\n✅ Re-indexação iniciada!\n")
	fmt.Printf("   Total na fila: %d jobs\n\n", totalQueue)
	fmt.Println("💡 Próximos passos:")
	fmt.Println("   1. Adicione VOYAGE_API_KEY no .env")
	fmt.Println("   2. Reinicie a API: docker compose up -d --build api")
	fmt.Println("   3. O EmbeddingWorker processará a fila automaticamente")
	fmt.Println("   4. Acompanhe: docker compose logs -f api | grep -E '(Voyage|OpenAI|✅|❌)'")
}
