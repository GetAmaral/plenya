//go:build legacy_scripts
// +build legacy_scripts

// Script ad-hoc — pra rodar use `go run -tags legacy_scripts ./cmd/batch-upload-articles`
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/services"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// Script para fazer upload em lote de todos os PDFs em /uploads/originals
// Simula o fluxo de upload normal: salva em /uploads/articles, extrai metadados, auto-queue para embedding
//
// Uso:
//   go run apps/api/cmd/batch-upload-articles/main.go

func main() {
	ctx := context.Background()

	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using environment variables")
	}

	// Conectar ao banco
	cfg := config.LoadConfig()
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Connected to database")

	// Criar serviços
	uploadFolder := "./uploads/articles"
	queueService := services.NewEmbeddingQueueService(db)
	aiService := services.NewAIService(cfg)
	articleService := services.NewArticleService(db, uploadFolder, queueService, aiService)

	// Diretório de origem (PDFs a processar)
	originalsDir := "./uploads/originals"

	// Listar todos os PDFs
	files, err := filepath.Glob(filepath.Join(originalsDir, "*.pdf"))
	if err != nil {
		log.Fatalf("❌ Failed to list PDF files: %v", err)
	}

	if len(files) == 0 {
		log.Fatalf("❌ No PDF files found in %s", originalsDir)
	}

	fmt.Printf("\n📚 Found %d PDF files in %s\n\n", len(files), originalsDir)

	// User ID real do banco (doctor2@plenya.com)
	userID := uuid.MustParse("019b8139-b683-75b2-8d27-82c3db7deedd")

	// Contadores
	successCount := 0
	failCount := 0
	skippedCount := 0

	// Processar cada PDF
	for i, filePath := range files {
		filename := filepath.Base(filePath)
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(files), filename)

		// Abrir arquivo
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("   ❌ Failed to open file: %v\n\n", err)
			failCount++
			continue
		}

		// Upload via ArticleService (simula upload HTTP)
		article, err := articleService.UploadPDF(file, filename, userID)
		file.Close()

		if err != nil {
			// Verificar se é duplicação
			if err.Error() == "este arquivo PDF já foi importado anteriormente" {
				fmt.Printf("   ⏭️  Skipped (already imported)\n\n")
				skippedCount++
			} else {
				fmt.Printf("   ❌ Failed: %v\n\n", err)
				failCount++
			}
			continue
		}

		fmt.Printf("   ✅ Uploaded successfully\n")
		fmt.Printf("      ID: %s\n", article.ID.String()[:8]+"...")
		fmt.Printf("      Title: %s\n", truncate(article.Title, 60))
		fmt.Printf("      Authors: %s\n", truncate(article.Authors, 60))

		// Mostrar status de abstract e fullContent
		hasAbstract := article.Abstract != nil && *article.Abstract != ""
		hasFullContent := article.FullContent != nil && *article.FullContent != ""
		fmt.Printf("      Abstract: %v | FullContent: %v\n", hasAbstract, hasFullContent)
		fmt.Printf("      Embedding: %s (queued automatically)\n\n", article.EmbeddingStatus)

		successCount++

		// Pequeno delay para não sobrecarregar (opcional)
		time.Sleep(500 * time.Millisecond)
	}

	// Resumo final
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("📊 BATCH UPLOAD SUMMARY\n")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("   ✅ Success:  %d\n", successCount)
	fmt.Printf("   ⏭️  Skipped:  %d (already imported)\n", skippedCount)
	fmt.Printf("   ❌ Failed:   %d\n", failCount)
	fmt.Printf("   📦 Total:    %d\n\n", len(files))

	if successCount > 0 {
		fmt.Println("🤖 Embeddings are being processed by the background worker")
		fmt.Println("   Check progress with: docker compose logs -f api")
	}

	// Mostrar estatísticas da fila
	stats, err := queueService.GetStats()
	if err == nil {
		fmt.Println("\n📋 EMBEDDING QUEUE STATUS")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Printf("   Pending:     %d\n", stats["pending"])
		fmt.Printf("   Processing:  %d\n", stats["processing"])
		fmt.Printf("   Completed:   %d\n", stats["completed"])
		fmt.Printf("   Failed:      %d\n", stats["failed"])
	}

	// Fechar conexão
	sqlDB, _ := db.DB()
	sqlDB.Close()

	fmt.Println("\n✅ Batch upload completed!")
}

// truncate trunca string para N caracteres
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
