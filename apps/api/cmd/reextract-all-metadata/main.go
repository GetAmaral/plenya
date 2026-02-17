package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Script para re-extrair metadados de artigos usando Claude
// Processa artigos com metadados incompletos (Autor desconhecido, etc)

func main() {
	fmt.Println("🔄 Re-extração de Metadados com Claude Haiku")
	fmt.Println("============================================\n")

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

	fmt.Println("✅ Conectado ao banco\n")

	// Criar services
	aiService := services.NewAIService(cfg)
	articleService := services.NewArticleService(db, "./uploads/articles", nil, aiService)

	// Buscar artigos com metadados incompletos
	type ArticleInfo struct {
		ID           string
		Title        string
		Authors      string
		Journal      string
		InternalLink *string
	}

	var articles []ArticleInfo
	err = db.Raw(`
		SELECT
			id,
			title,
			authors,
			journal,
			internal_link
		FROM articles
		WHERE (
			authors = 'Autor desconhecido'
			OR journal = 'Revista não identificada'
		)
		AND internal_link IS NOT NULL
		AND deleted_at IS NULL
		ORDER BY created_at DESC
	`).Scan(&articles).Error

	if err != nil {
		log.Fatalf("❌ Erro ao buscar artigos: %v", err)
	}

	if len(articles) == 0 {
		fmt.Println("✅ Nenhum artigo com metadados incompletos encontrado")
		return
	}

	fmt.Printf("📄 Encontrados %d artigos para re-extrair\n\n", len(articles))

	// Processar cada artigo
	successCount := 0
	failCount := 0
	skippedCount := 0

	for i, article := range articles {
		fmt.Printf("[%d/%d] %s\n", i+1, len(articles), truncate(article.Title, 70))
		fmt.Printf("   ID: %s\n", article.ID[:8]+"...")
		fmt.Printf("   Authors atual: %s\n", truncate(article.Authors, 50))

		if article.InternalLink == nil {
			fmt.Println("   ⚠️  Sem PDF, pulando...\n")
			skippedCount++
			continue
		}

		// Re-extrair metadados
		metadata, err := articleService.ExtractPDFMetadata(*article.InternalLink)
		if err != nil {
			fmt.Printf("   ❌ Erro: %v\n\n", err)
			failCount++
			continue
		}

		// Verificar se extraiu algo útil
		if metadata.Title == "" && metadata.Authors == "" {
			fmt.Println("   ⚠️  Extração não retornou metadados úteis\n")
			skippedCount++
			continue
		}

		// Atualizar no banco
		updates := map[string]interface{}{}

		if metadata.Title != "" {
			updates["title"] = metadata.Title
		}
		if metadata.Authors != "" {
			updates["authors"] = metadata.Authors
		}
		if metadata.DOI != "" {
			updates["doi"] = metadata.DOI
		}
		if metadata.PMID != "" {
			updates["pm_id"] = metadata.PMID
		}
		if metadata.Abstract != "" {
			updates["abstract"] = metadata.Abstract
		}
		if len(metadata.Keywords) > 0 {
			updates["keywords"] = metadata.Keywords
		}

		if len(updates) == 0 {
			fmt.Println("   ⚠️  Nenhum metadado novo para atualizar\n")
			skippedCount++
			continue
		}

		updates["updated_at"] = time.Now()

		err = db.Table("articles").Where("id = ?", article.ID).Updates(updates).Error
		if err != nil {
			fmt.Printf("   ❌ Erro ao atualizar: %v\n\n", err)
			failCount++
			continue
		}

		fmt.Println("   ✅ Metadados atualizados:")
		if metadata.Title != "" {
			fmt.Printf("      Title: %s\n", truncate(metadata.Title, 70))
		}
		if metadata.Authors != "" {
			fmt.Printf("      Authors: %s\n", truncate(metadata.Authors, 70))
		}
		if metadata.DOI != "" {
			fmt.Printf("      DOI: %s\n", metadata.DOI)
		}
		fmt.Println()

		successCount++

		// Delay para não sobrecarregar API do Claude
		time.Sleep(2 * time.Second)
	}

	// Resumo
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("✅ Sucesso:  %d\n", successCount)
	fmt.Printf("⏭️  Pulados:  %d\n", skippedCount)
	fmt.Printf("❌ Falhas:   %d\n", failCount)
	fmt.Printf("📦 Total:    %d\n", len(articles))
	fmt.Println("\n✅ Re-extração concluída!")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
