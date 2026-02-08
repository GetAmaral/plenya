package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
)

// Palavras-chave para detecção de gênero (case-insensitive)
var maleKeywords = []string{
	"homem",
	"homens",
	"masculino",
	"masculina",
	"masculinas",
	"masculinos",
	"homem adulto",
	"sexo masculino",
	// Plural e variações
	"dos homens",
	"para homens",
	"em homens",
	"no homem",
	"nos homens",
}

var femaleKeywords = []string{
	"mulher",
	"mulheres",
	"feminino",
	"feminina",
	"femininas",
	"femininos",
	"mulher adulta",
	"sexo feminino",
	// Plural e variações
	"das mulheres",
	"para mulheres",
	"em mulheres",
	"na mulher",
	"nas mulheres",
}

// detectGender analisa o nome e retorna o gênero detectado
func detectGender(name string) string {
	nameLower := strings.ToLower(name)

	// Verificar palavras-chave masculinas
	for _, keyword := range maleKeywords {
		if strings.Contains(nameLower, keyword) {
			return "male"
		}
	}

	// Verificar palavras-chave femininas
	for _, keyword := range femaleKeywords {
		if strings.Contains(nameLower, keyword) {
			return "female"
		}
	}

	// Nenhuma palavra-chave encontrada
	return "not_applicable"
}

func main() {
	fmt.Println("=== Score Item Gender Update Script ===\n")

	// Carregar .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}

	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Conectar ao banco de dados
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()
	fmt.Println("✅ Database connected\n")

	// Buscar todos os scoreItems (incluindo soft deleted se necessário)
	var scoreItems []models.ScoreItem
	if err := database.DB.Find(&scoreItems).Error; err != nil {
		log.Fatalf("❌ Failed to fetch score items: %v", err)
	}

	if len(scoreItems) == 0 {
		fmt.Println("⚠️  No score items found in database")
		return
	}

	fmt.Printf("📊 Found %d score items\n\n", len(scoreItems))

	// Contadores
	var (
		totalProcessed = 0
		maleCount      = 0
		femaleCount    = 0
		naCount        = 0
		unchanged      = 0
		errors         = 0
	)

	// Processar cada item
	fmt.Println("Processing score items:")
	fmt.Println("----------------------------------------")

	for i, item := range scoreItems {
		totalProcessed++

		// Detectar gênero baseado no nome
		detectedGender := detectGender(item.Name)

		// Verificar se já tem o valor correto
		currentGender := "not_applicable" // valor padrão
		if item.Gender != nil {
			currentGender = *item.Gender
		}

		if currentGender == detectedGender {
			unchanged++
			continue
		}

		// Atualizar no banco de dados
		if err := database.DB.Model(&models.ScoreItem{}).
			Where("id = ?", item.ID).
			Update("gender", detectedGender).Error; err != nil {
			log.Printf("❌ Error updating item %s: %v\n", item.ID, err)
			errors++
			continue
		}

		// Contar por tipo
		switch detectedGender {
		case "male":
			maleCount++
		case "female":
			femaleCount++
		case "not_applicable":
			naCount++
		}

		// Log da mudança
		fmt.Printf("[%d/%d] ✓ %s\n", i+1, len(scoreItems), item.Name)
		fmt.Printf("        %s → %s\n", currentGender, detectedGender)
	}

	// Resumo final
	fmt.Println("\n----------------------------------------")
	fmt.Println("=== Summary ===")
	fmt.Printf("Total processed:     %d\n", totalProcessed)
	fmt.Printf("Updated to 'male':   %d\n", maleCount)
	fmt.Printf("Updated to 'female': %d\n", femaleCount)
	fmt.Printf("Updated to 'not_applicable': %d\n", naCount)
	fmt.Printf("Unchanged:           %d\n", unchanged)
	fmt.Printf("Errors:              %d\n", errors)
	fmt.Println("\n✅ Script completed successfully!")

	os.Exit(0)
}
