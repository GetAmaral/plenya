package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	// IDs dos 3 exemplos
	ids := []string{
		"019bf31d-2ef0-71b6-a5d1-db49a4fa62fa", // Leptina - EXCELLENT
		"019c1a2b-a36f-7b25-9a74-340e3e982bd3", // NOS3 - GOOD
		"019c5001-2633-7496-ac85-ae56c0863996", // Adolescência - FAIR
	}

	for i, idStr := range ids {
		id := uuid.MustParse(idStr)

		var scoreItem models.ScoreItem
		db.Preload("Subgroup.Group").First(&scoreItem, id)

		var prep models.ScoreItemEnrichmentPreparation
		db.First(&prep, "score_item_id = ?", id)

		fmt.Printf("\n" + strings.Repeat("=", 80) + "\n")
		fmt.Printf("EXEMPLO %d: %s\n", i+1, scoreItem.Name)
		fmt.Printf(strings.Repeat("=", 80) + "\n\n")

		printEnrichmentPrompt(&scoreItem, &prep)

		fmt.Printf("\n" + strings.Repeat("-", 80) + "\n")
	}
}

func printEnrichmentPrompt(si *models.ScoreItem, prep *models.ScoreItemEnrichmentPreparation) {
	// Metadata
	totalChunks := int(prep.Metadata["total_chunks"].(float64))
	articlesCount := int(prep.Metadata["articles_count"].(float64))
	avgSim := prep.Metadata["avg_similarity"].(float64)
	grade := prep.Metadata["quality_grade"].(string)

	fmt.Printf("📊 METADATA\n")
	fmt.Printf("   Quality Grade: %s\n", strings.ToUpper(grade))
	fmt.Printf("   Total Chunks: %d\n", totalChunks)
	fmt.Printf("   Articles: %d\n", articlesCount)
	fmt.Printf("   Avg Similarity: %.3f\n", avgSim)
	fmt.Printf("\n")

	// ScoreItem info
	fmt.Printf("📌 SCOREITEM INFO\n")
	fmt.Printf("   Name: %s\n", si.Name)
	if si.Unit != nil {
		fmt.Printf("   Unit: %s\n", *si.Unit)
	}
	if si.Gender != nil && *si.Gender != "not_applicable" {
		fmt.Printf("   Gender: %s\n", *si.Gender)
	}
	fmt.Printf("\n")

	// Clinical context (se existir)
	if si.ClinicalRelevance != nil && *si.ClinicalRelevance != "" {
		fmt.Printf("📖 CLINICAL RELEVANCE (ATUAL)\n")
		fmt.Printf("   %s\n\n", truncate(*si.ClinicalRelevance, 200))
	}

	// Chunks científicos
	fmt.Printf("🔬 CHUNKS CIENTÍFICOS (%d chunks de %d artigos)\n\n", totalChunks, articlesCount)

	// Parse chunks
	chunksData := prep.SelectedChunks["items"].([]interface{})

	// Mostrar primeiros 3 chunks
	for i := 0; i < 3 && i < len(chunksData); i++ {
		chunkMap := chunksData[i].(map[string]interface{})

		fmt.Printf("   Chunk %d/%d:\n", i+1, len(chunksData))
		fmt.Printf("   └─ Article: %s (%v)\n", chunkMap["article_title"], chunkMap["article_year"])
		fmt.Printf("   └─ Journal: %s\n", chunkMap["journal"])
		fmt.Printf("   └─ Section: %s | Similarity: %.3f\n", chunkMap["section"], chunkMap["similarity"])
		fmt.Printf("   └─ Text: %s\n", truncate(chunkMap["chunk_text"].(string), 200))
		fmt.Printf("\n")
	}

	if len(chunksData) > 3 {
		fmt.Printf("   ... e mais %d chunks científicos disponíveis\n\n", len(chunksData)-3)
	}

	// Exemplo de prompt para Claude
	fmt.Printf("💬 PROMPT PARA CLAUDE (estrutura)\n")
	fmt.Printf("   ```\n")
	fmt.Printf("   Você é um médico especialista. Analise os %d chunks científicos\n", totalChunks)
	fmt.Printf("   sobre \"%s\" e gere:\n", si.Name)
	fmt.Printf("\n")
	fmt.Printf("   1. ClinicalRelevance: Explicação técnica para médicos (300-500 palavras)\n")
	fmt.Printf("   2. PatientExplanation: Explicação simples para pacientes (200-300 palavras)\n")
	fmt.Printf("   3. Conduct: Condutas clínicas baseadas em evidências (300-400 palavras)\n")
	fmt.Printf("\n")
	fmt.Printf("   Contexto científico disponível:\n")
	fmt.Printf("   - %d chunks de %d artigos\n", totalChunks, articlesCount)
	fmt.Printf("   - Avg similarity: %.3f (%s quality)\n", avgSim, strings.ToUpper(grade))
	fmt.Printf("   - Seções: results, discussion, methods, introduction\n")
	fmt.Printf("   ```\n")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
