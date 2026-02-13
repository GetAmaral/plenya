package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("🔍 Extracting text from downloaded PDFs")
	fmt.Println("========================================")

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

	// Criar article service (reutilizar ExtractPDFMetadata)
	articleService := services.NewArticleService(db, "./uploads/articles", nil)

	// Buscar artigos com PDF mas sem fullContent/abstract
	var articles []models.Article
	err = db.Raw(`
		SELECT *
		FROM articles
		WHERE internal_link IS NOT NULL
		  AND internal_link LIKE '/uploads/articles/%'
		  AND (full_content IS NULL OR full_content = '' OR abstract IS NULL OR abstract = '')
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
	`).Scan(&articles).Error

	if err != nil {
		log.Fatalf("❌ Failed to fetch articles: %v", err)
	}

	fmt.Printf("📚 Found %d articles with PDFs but no text content\n\n", len(articles))

	successCount := 0
	errorCount := 0

	// Processar cada artigo
	for i, article := range articles {
		fmt.Printf("[%d/%d] Processing: %s\n", i+1, len(articles), truncate(article.Title, 80))

		// Construir caminho do PDF
		if article.InternalLink == nil {
			fmt.Printf("   ❌ No internal_link set\n\n")
			errorCount++
			continue
		}
		pdfPath := filepath.Join(".", (*article.InternalLink)[1:]) // Remove leading /

		// Verificar se arquivo existe
		if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
			fmt.Printf("   ❌ PDF not found: %s\n\n", pdfPath)
			errorCount++
			continue
		}

		// Extrair metadados do PDF
		metadata, err := articleService.ExtractPDFMetadata(pdfPath)
		if err != nil {
			fmt.Printf("   ❌ Failed to extract PDF: %v\n\n", err)
			errorCount++
			continue
		}

		// Atualizar artigo com texto extraído
		updates := map[string]interface{}{
			"full_content": metadata.FullContent,
		}

		// Se abstract foi extraído, atualizar também
		if metadata.Abstract != "" {
			updates["abstract"] = metadata.Abstract
		}

		err = db.Model(&models.Article{}).
			Where("id = ?", article.ID).
			Updates(updates).Error

		if err != nil {
			fmt.Printf("   ❌ Failed to update article: %v\n\n", err)
			errorCount++
			continue
		}

		// Resetar embedding_queue para pending
		err = db.Exec(`
			UPDATE embedding_queue
			SET status = 'pending', retry_count = 0, error_message = NULL
			WHERE entity_id = ? AND entity_type = 'article'
		`, article.ID).Error

		if err != nil {
			fmt.Printf("   ⚠️  Failed to reset queue: %v\n", err)
		}

		successCount++
		fmt.Printf("   ✅ Extracted %d chars (abstract: %d chars)\n", len(metadata.FullContent), len(metadata.Abstract))
		fmt.Printf("   ✅ Queued for embedding\n\n")
	}

	// Resumo final
	fmt.Println("========================================")
	fmt.Println("📊 Summary:")
	fmt.Printf("   ✅ Extracted: %d\n", successCount)
	fmt.Printf("   ❌ Errors: %d\n", errorCount)
	fmt.Println("========================================")
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
