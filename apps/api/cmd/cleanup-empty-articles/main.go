package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Cleanup de artigos vazios (sem PDF, sem abstract, sem fullContent)
// Faz soft delete de artigos que falharam embedding por estarem vazios

func main() {
	fmt.Println("🗑️  Cleanup de Artigos Vazios")
	fmt.Println("=============================\n")

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

	// Buscar artigos vazios SEM PDF
	type EmptyArticle struct {
		ID           string
		Title        string
		InternalLink *string
	}

	var emptyArticles []EmptyArticle
	err = db.Raw(`
		SELECT a.id, a.title, a.internal_link
		FROM articles a
		JOIN embedding_queue eq ON eq.entity_id = a.id
		WHERE eq.status = 'failed'
		  AND eq.entity_type = 'article'
		  AND eq.error_message LIKE '%no chunks generated%'
		  AND (a.internal_link IS NULL OR a.internal_link = '')
		  AND a.deleted_at IS NULL
	`).Scan(&emptyArticles).Error

	if err != nil {
		log.Fatalf("❌ Erro ao buscar artigos: %v", err)
	}

	fmt.Printf("📋 Encontrados %d artigos vazios SEM PDF:\n\n", len(emptyArticles))

	if len(emptyArticles) == 0 {
		fmt.Println("✅ Nenhum artigo para deletar.")
		return
	}

	// Mostrar lista
	for i, a := range emptyArticles {
		title := a.Title
		if len(title) > 80 {
			title = title[:80] + "..."
		}
		fmt.Printf("%d. %s\n", i+1, title)
	}

	fmt.Printf("\n⚠️  Esses artigos serão marcados como deletados (soft delete).\n")
	fmt.Printf("Continue? (y/N): ")

	var confirm string
	fmt.Scanln(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("❌ Cancelado.")
		return
	}

	// Soft delete
	fmt.Printf("\n🗑️  Deletando %d artigos...\n", len(emptyArticles))

	articleIDs := make([]string, len(emptyArticles))
	for i, a := range emptyArticles {
		articleIDs[i] = a.ID
	}

	result := db.Exec(`
		UPDATE articles
		SET deleted_at = NOW()
		WHERE id = ANY($1)
	`, articleIDs)

	if result.Error != nil {
		log.Fatalf("❌ Erro ao deletar: %v", result.Error)
	}

	fmt.Printf("✅ %d artigos deletados (soft delete)\n", result.RowsAffected)

	// Remover da fila
	db.Exec(`
		DELETE FROM embedding_queue
		WHERE entity_id = ANY($1) AND entity_type = 'article'
	`, articleIDs)

	fmt.Println("✅ Removidos da fila de embeddings\n")
	fmt.Println("🎉 Cleanup completo!")
}
