package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Script para re-extrair conteúdo de PDFs que falharam
// Testa o novo fallback para pdftotext

func main() {
	fmt.Println("🔄 Re-extração de PDFs Falhados")
	fmt.Println("================================\n")

	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env não encontrado")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco: %v", err)
	}

	fmt.Println("✅ Conectado ao banco\n")

	// Criar ArticleService
	uploadFolder := "./uploads/articles"
	queueService := services.NewEmbeddingQueueService(db)
	aiService := services.NewAIService(cfg)
	articleService := services.NewArticleService(db, uploadFolder, queueService, aiService)

	// Buscar artigos com status pending e sem conteúdo
	type ArticleInfo struct {
		ID           string
		Title        string
		InternalLink *string
		ContentLen   int
	}

	var articles []ArticleInfo
	err = db.Raw(`
		SELECT
			id,
			title,
			internal_link,
			LENGTH(COALESCE(full_content, '')) as content_len
		FROM articles
		WHERE embedding_status = 'pending'
		  AND (full_content IS NULL OR full_content = '')
		  AND internal_link IS NOT NULL
		  AND deleted_at IS NULL
		ORDER BY title
	`).Scan(&articles).Error

	if err != nil {
		log.Fatalf("❌ Erro ao buscar artigos: %v", err)
	}

	if len(articles) == 0 {
		fmt.Println("✅ Nenhum artigo pendente sem conteúdo encontrado")
		return
	}

	fmt.Printf("📄 Encontrados %d artigos para re-extrair\n\n", len(articles))

	// Re-extrair cada artigo
	successCount := 0
	failCount := 0

	for i, article := range articles {
		fmt.Printf("[%d/%d] %s\n", i+1, len(articles), article.Title)
		fmt.Printf("   ID: %s\n", article.ID[:8]+"...")

		if article.InternalLink == nil {
			fmt.Println("   ⚠️  Sem InternalLink, pulando...")
			failCount++
			continue
		}

		// Re-extrair metadados do PDF
		metadata, err := articleService.ExtractPDFMetadata(*article.InternalLink)
		if err != nil {
			fmt.Printf("   ❌ Erro ao extrair: %v\n\n", err)
			failCount++
			continue
		}

		// Verificar se extraiu conteúdo
		contentLen := len(metadata.FullContent)
		if contentLen < 100 {
			fmt.Printf("   ⚠️  Extração falhou (apenas %d chars)\n\n", contentLen)
			failCount++
			continue
		}

		// Atualizar no banco
		err = db.Exec(`
			UPDATE articles
			SET
				full_content = ?,
				abstract = ?,
				updated_at = NOW()
			WHERE id = ?
		`, metadata.FullContent, metadata.Abstract, article.ID).Error

		if err != nil {
			fmt.Printf("   ❌ Erro ao atualizar: %v\n\n", err)
			failCount++
			continue
		}

		fmt.Printf("   ✅ Extraído: %d caracteres\n", contentLen)
		if metadata.Abstract != "" {
			fmt.Printf("   📝 Abstract: %d caracteres\n", len(metadata.Abstract))
		}

		// Adicionar à fila de embeddings
		articleUUID, err := uuid.Parse(article.ID)
		if err != nil {
			fmt.Printf("   ⚠️  Erro ao parsear UUID: %v\n", err)
		} else if err := queueService.QueueArticle(articleUUID); err != nil {
			fmt.Printf("   ⚠️  Erro ao adicionar à fila: %v\n", err)
		} else {
			fmt.Println("   📥 Adicionado à fila de embeddings")
		}

		fmt.Println()
		successCount++
	}

	// Resumo
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("✅ Sucesso: %d\n", successCount)
	fmt.Printf("❌ Falhas:  %d\n", failCount)
	fmt.Printf("📦 Total:   %d\n", len(articles))
	fmt.Println("\n✅ Re-extração concluída!")
}
