package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	// Create repositories and services
	vectorRepo := repository.NewArticleVectorRepository(db)
	embeddingService := services.NewEmbeddingService(cfg, db)
	semanticService := services.NewArticleSemanticService(vectorRepo, embeddingService)

	// Get query from command line
	query := "nutrition"
	if len(os.Args) > 1 {
		query = os.Args[1]
	}

	// Perform semantic search
	fmt.Printf("🔍 Searching for: %s\n\n", query)

	results, err := semanticService.SemanticSearch(context.Background(), services.SemanticSearchDTO{
		Query:         query,
		Limit:         10,
		MinSimilarity: 0, // Usa o padrão definido no serviço (0.4)
	})

	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	// Display results
	fmt.Printf("Found %d results:\n\n", len(results))

	for i, result := range results {
		fmt.Printf("%d. [%.3f] %s\n", i+1, result.Similarity, result.Article.Title)
		if result.Article.Abstract != nil && len(*result.Article.Abstract) > 100 {
			fmt.Printf("   %s...\n", (*result.Article.Abstract)[:100])
		} else if result.Article.Abstract != nil {
			fmt.Printf("   %s\n", *result.Article.Abstract)
		}
		if len(result.ChunkText) > 100 {
			fmt.Printf("   Chunk: %s...\n\n", result.ChunkText[:100])
		} else {
			fmt.Printf("   Chunk: %s\n\n", result.ChunkText)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
