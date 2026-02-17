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
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env não encontrado")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}

	articleService := services.NewArticleService(db, "./uploads/articles", nil)

	// Buscar PRÓXIMO artigo com PDF mas sem fullContent
	var article models.Article
	err = db.Raw(`
		SELECT *
		FROM articles
		WHERE internal_link IS NOT NULL
		  AND internal_link LIKE '/uploads/articles/%'
		  AND (full_content IS NULL OR full_content = '')
		  AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT 1
	`).Scan(&article).Error

	if err == gorm.ErrRecordNotFound || article.ID.String() == "00000000-0000-0000-0000-000000000000" {
		fmt.Println("✅ Nenhum artigo para processar")
		return
	}

	if err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}

	fmt.Printf("📄 %s\n", article.Title)

	pdfPath := filepath.Join(".", (*article.InternalLink)[1:])
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		fmt.Printf("❌ PDF não encontrado: %s\n", pdfPath)
		os.Exit(1)
	}

	metadata, err := articleService.ExtractPDFMetadata(pdfPath)
	if err != nil {
		log.Fatalf("❌ Erro ao extrair: %v", err)
	}

	updates := map[string]interface{}{"full_content": metadata.FullContent}
	if metadata.Abstract != "" {
		updates["abstract"] = metadata.Abstract
	}

	err = db.Model(&models.Article{}).Where("id = ?", article.ID).Updates(updates).Error
	if err != nil {
		log.Fatalf("❌ Erro ao atualizar: %v", err)
	}

	db.Exec(`
		UPDATE embedding_queue
		SET status = 'pending', retry_count = 0, error_message = NULL
		WHERE entity_id = ? AND entity_type = 'article'
	`, article.ID)

	fmt.Printf("✅ Extraído: %d chars\n", len(metadata.FullContent))
	fmt.Println("✅ Enfileirado para embedding")
}
