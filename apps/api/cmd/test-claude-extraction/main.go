package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Script de teste para extração inteligente com Claude
func main() {
	fmt.Println("🧪 Teste de Extração Inteligente com Claude")
	fmt.Println("=============================================\n")

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

	// Arquivo de teste
	testPDF := "./uploads/articles/475a2cff-5fe8-4053-a2b9-7f432c2a827c.pdf"

	fmt.Printf("📄 Extraindo metadados de: %s\n\n", testPDF)

	// Extrair com Claude
	metadata, err := articleService.ExtractPDFMetadata(testPDF)
	if err != nil {
		log.Fatalf("❌ Erro: %v", err)
	}

	// Exibir resultados
	fmt.Println("\n📊 METADADOS EXTRAÍDOS")
	fmt.Println("=" + string(make([]byte, 60)))
	fmt.Printf("Title:    %s\n", metadata.Title)
	fmt.Printf("Authors:  %s\n", metadata.Authors)
	fmt.Printf("DOI:      %s\n", metadata.DOI)
	fmt.Printf("PMID:     %s\n", metadata.PMID)
	fmt.Printf("Abstract: %d chars\n", len(metadata.Abstract))
	fmt.Printf("Content:  %d chars\n", len(metadata.FullContent))

	if len(metadata.Keywords) > 0 {
		fmt.Printf("Keywords: %v\n", metadata.Keywords)
	}

	fmt.Println("\n✅ Extração concluída!")
}
