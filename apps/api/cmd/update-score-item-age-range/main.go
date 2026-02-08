package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
)

// AgeRange representa uma faixa etária detectada
type AgeRange struct {
	Min      *int
	Max      *int
	Detected bool
	Pattern  string // Qual padrão foi detectado
}

// Padrões de idade (compilados uma vez para performance)
var (
	// Padrão: "18-65 anos" ou "18 a 65 anos" ou "18-65"
	rangePattern1 = regexp.MustCompile(`(?i)(\d+)\s*[-a]\s*(\d+)\s*anos?`)

	// Padrão: "(18-65 anos)" ou "(18 a 65)"
	rangePattern2 = regexp.MustCompile(`(?i)\((\d+)\s*[-a]\s*(\d+)\s*anos?\)`)

	// Padrão: "< 18 anos" ou "<18 anos" ou "menor que 18 anos"
	lessThanPattern = regexp.MustCompile(`(?i)(?:<\s*|menor\s+(?:que|de)\s+|abaixo\s+de\s+)(\d+)\s*anos?`)

	// Padrão: "> 65 anos" ou ">65 anos" ou "maior que 65 anos" ou "acima de 65 anos"
	greaterThanPattern = regexp.MustCompile(`(?i)(?:>\s*|maior\s+(?:que|de)\s+|acima\s+de\s+)(\d+)\s*anos?`)

	// Padrão: "18+ anos" ou "18 +" ou "65+" ou "50+"
	plusPattern = regexp.MustCompile(`(?i)(\d+)\s*\+(?:\s*anos?)?`)

	// Termos a IGNORAR (não são faixas etárias específicas)
	ignoreTerms = []string{
		"pré-menopausa",
		"pré menopausa",
		"premenopausa",
		"pós-menopausa",
		"pós menopausa",
		"posmenopausa",
		"gestação",
		"gestante",
		"grávida",
		"gravidez",
		"lactante",
		"lactação",
		"puerpério",
		"fase folicular",
		"fase lútea",
		"ovulatório",
		"menstruação",
		"ciclo menstrual",
	}
)

// detectAgeRange analisa o nome e retorna a faixa etária detectada
func detectAgeRange(name string) AgeRange {
	nameLower := strings.ToLower(name)

	// Verificar se contém termos a ignorar
	for _, term := range ignoreTerms {
		if strings.Contains(nameLower, term) {
			return AgeRange{Detected: false}
		}
	}

	// 1. Tentar padrão de range "X-Y anos" ou "X a Y anos"
	if matches := rangePattern1.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		// Validar que min < max
		if min < max && min >= 0 && max <= 150 {
			return AgeRange{
				Min:      &min,
				Max:      &max,
				Detected: true,
				Pattern:  fmt.Sprintf("%d-%d anos", min, max),
			}
		}
	}

	// 2. Tentar padrão de range entre parênteses "(X-Y anos)"
	if matches := rangePattern2.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		max, _ := strconv.Atoi(matches[2])
		if min < max && min >= 0 && max <= 150 {
			return AgeRange{
				Min:      &min,
				Max:      &max,
				Detected: true,
				Pattern:  fmt.Sprintf("(%d-%d anos)", min, max),
			}
		}
	}

	// 3. Tentar padrão "< X anos" ou "abaixo de X anos"
	if matches := lessThanPattern.FindStringSubmatch(name); matches != nil {
		max, _ := strconv.Atoi(matches[1])
		if max > 0 && max <= 150 {
			min := 0
			return AgeRange{
				Min:      &min,
				Max:      &max,
				Detected: true,
				Pattern:  fmt.Sprintf("< %d anos", max),
			}
		}
	}

	// 4. Tentar padrão "> X anos" ou "acima de X anos"
	if matches := greaterThanPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Usar 150 como limite máximo (constraint do banco)
			return AgeRange{
				Min:      &min,
				Max:      &max,
				Detected: true,
				Pattern:  fmt.Sprintf("> %d anos", min),
			}
		}
	}

	// 5. Tentar padrão "X+ anos" ou "X +"
	if matches := plusPattern.FindStringSubmatch(name); matches != nil {
		min, _ := strconv.Atoi(matches[1])
		if min >= 0 && min < 150 {
			max := 150 // Usar 150 como limite máximo (constraint do banco)
			return AgeRange{
				Min:      &min,
				Max:      &max,
				Detected: true,
				Pattern:  fmt.Sprintf("%d+ anos", min),
			}
		}
	}

	// Nenhum padrão detectado
	return AgeRange{Detected: false}
}

func main() {
	fmt.Println("=== Score Item Age Range Update Script ===")
	fmt.Println()

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
	fmt.Println("✅ Database connected")
	fmt.Println()

	// Buscar todos os scoreItems
	var scoreItems []models.ScoreItem
	if err := database.DB.Find(&scoreItems).Error; err != nil {
		log.Fatalf("❌ Failed to fetch score items: %v", err)
	}

	if len(scoreItems) == 0 {
		fmt.Println("⚠️  No score items found in database")
		return
	}

	fmt.Printf("📊 Found %d score items\n", len(scoreItems))
	fmt.Println()

	// Contadores
	var (
		totalProcessed = 0
		rangeDetected  = 0
		rangeUpdated   = 0
		unchanged      = 0
		ignored        = 0
		errors         = 0
	)

	// Processar cada item
	fmt.Println("Processing score items:")
	fmt.Println("----------------------------------------")

	for i, item := range scoreItems {
		totalProcessed++

		// Detectar age range baseado no nome
		ageRange := detectAgeRange(item.Name)

		// Se não detectou nenhum padrão
		if !ageRange.Detected {
			ignored++
			continue
		}

		rangeDetected++

		// Verificar se já tem os valores corretos
		currentMin := item.AgeRangeMin
		currentMax := item.AgeRangeMax

		if (currentMin == nil && ageRange.Min == nil || currentMin != nil && ageRange.Min != nil && *currentMin == *ageRange.Min) &&
			(currentMax == nil && ageRange.Max == nil || currentMax != nil && ageRange.Max != nil && *currentMax == *ageRange.Max) {
			unchanged++
			continue
		}

		// Atualizar no banco de dados
		updates := map[string]interface{}{
			"age_range_min": ageRange.Min,
			"age_range_max": ageRange.Max,
		}

		if err := database.DB.Model(&models.ScoreItem{}).
			Where("id = ?", item.ID).
			Updates(updates).Error; err != nil {
			log.Printf("❌ Error updating item %s: %v\n", item.ID, err)
			errors++
			continue
		}

		rangeUpdated++

		// Log da mudança
		fmt.Printf("[%d/%d] ✓ %s\n", i+1, len(scoreItems), item.Name)
		fmt.Printf("        Pattern: %s\n", ageRange.Pattern)

		var oldRange, newRange string
		if currentMin != nil && currentMax != nil {
			oldRange = fmt.Sprintf("%d-%d", *currentMin, *currentMax)
		} else if currentMin != nil {
			oldRange = fmt.Sprintf("%d-∞", *currentMin)
		} else if currentMax != nil {
			oldRange = fmt.Sprintf("0-%d", *currentMax)
		} else {
			oldRange = "null"
		}

		if ageRange.Min != nil && ageRange.Max != nil {
			newRange = fmt.Sprintf("%d-%d", *ageRange.Min, *ageRange.Max)
		} else if ageRange.Min != nil {
			newRange = fmt.Sprintf("%d-∞", *ageRange.Min)
		} else if ageRange.Max != nil {
			newRange = fmt.Sprintf("0-%d", *ageRange.Max)
		}

		fmt.Printf("        Range: %s → %s\n", oldRange, newRange)
	}

	// Resumo final
	fmt.Println("\n----------------------------------------")
	fmt.Println("=== Summary ===")
	fmt.Printf("Total processed:     %d\n", totalProcessed)
	fmt.Printf("Range detected:      %d\n", rangeDetected)
	fmt.Printf("Range updated:       %d\n", rangeUpdated)
	fmt.Printf("Unchanged:           %d\n", unchanged)
	fmt.Printf("Ignored (no range):  %d\n", ignored)
	fmt.Printf("Errors:              %d\n", errors)
	fmt.Println()
	fmt.Println("✅ Script completed successfully!")

	os.Exit(0)
}
