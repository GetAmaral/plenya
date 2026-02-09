package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// Regex para parsing de conversões: "1 g/dL = 10 g/L" ou "ng/mL × 2.5 = nmol/L" ou "nmol/L ÷ 1000 = μmol/L"
// Captura: (valor1) (unidade1) [operador] (valor2) (unidade2)
// Operadores suportados: = (igual), × (multiplicação), ÷ (divisão)
// Suporta vírgula como separador decimal (ex: 2,5)
var conversionRegex = regexp.MustCompile(`([0-9]+[.,]?[0-9]*)\s*([a-zA-Zµμ/]+)\s*([=×÷])\s*([0-9]+[.,]?[0-9]*)\s*([a-zA-Zµμ/]+)`)

type Stats struct {
	TotalScoreItems      int
	WithConversion       int
	Parsed               int
	Created              int
	Skipped              int
	Errors               int
	ErrorDetails         []string
	UniqueConversions    map[string]bool
	LabTestDefinitions   int
	LTDWithConversion    int
	LTDParsed            int
	LTDCreated           int
	LTDSkipped           int
}

func main() {
	// Carregar .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Conectar ao banco
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "plenya_user"),
		getEnv("DB_PASSWORD", "plenya_password"),
		getEnv("DB_NAME", "plenya_db"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	stats := &Stats{
		UniqueConversions: make(map[string]bool),
		ErrorDetails:      []string{},
	}

	log.Println("=== Migração de Conversões de Unidades ===")
	log.Println()

	// 1. Migrar de ScoreItem.unitConversion
	log.Println("📋 Fase 1: Extraindo conversões de ScoreItems...")
	migrateFromScoreItems(db, stats)

	// 2. Migrar de LabTestDefinition.unitConversion
	log.Println()
	log.Println("📋 Fase 2: Extraindo conversões de LabTestDefinitions...")
	migrateFromLabTestDefinitions(db, stats)

	// Imprimir estatísticas
	log.Println()
	log.Println("=== Estatísticas ===")
	log.Printf("ScoreItems:")
	log.Printf("  Total: %d", stats.TotalScoreItems)
	log.Printf("  Com conversão: %d", stats.WithConversion)
	log.Printf("  Parseadas com sucesso: %d", stats.Parsed)
	log.Printf("  Criadas: %d", stats.Created)
	log.Printf("  Ignoradas (duplicatas): %d", stats.Skipped)
	log.Println()
	log.Printf("LabTestDefinitions:")
	log.Printf("  Total: %d", stats.LabTestDefinitions)
	log.Printf("  Com conversão: %d", stats.LTDWithConversion)
	log.Printf("  Parseadas com sucesso: %d", stats.LTDParsed)
	log.Printf("  Criadas: %d", stats.LTDCreated)
	log.Printf("  Ignoradas (duplicatas): %d", stats.LTDSkipped)
	log.Println()
	log.Printf("Total de erros: %d", stats.Errors)
	log.Printf("Conversões únicas encontradas: %d", len(stats.UniqueConversions))

	if stats.Errors > 0 {
		log.Println()
		log.Println("=== Erros ===")
		for _, errMsg := range stats.ErrorDetails {
			log.Println(errMsg)
		}
	}

	log.Println()
	log.Println("✅ Migração concluída!")
}

func migrateFromScoreItems(db *gorm.DB, stats *Stats) {
	var scoreItems []models.ScoreItem
	if err := db.Where("unit_conversion IS NOT NULL AND unit_conversion != ''").Find(&scoreItems).Error; err != nil {
		log.Fatalf("Failed to fetch score items: %v", err)
	}

	stats.TotalScoreItems = len(scoreItems)
	log.Printf("Encontrados %d ScoreItems com conversão\n", stats.TotalScoreItems)

	for _, item := range scoreItems {
		stats.WithConversion++

		// Parse conversion string
		mainValue, mainUnit, secondaryValue, secondaryUnit, err := parseConversion(*item.UnitConversion)
		if err != nil {
			stats.Errors++
			errMsg := fmt.Sprintf("❌ ScoreItem %s (code=%s): %v", item.ID, item.LabTestCode, err)
			stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
			log.Println(errMsg)
			continue
		}

		stats.Parsed++

		// Calcular fator: secondary = main * factor → factor = secondary / main
		if mainValue == 0 {
			stats.Errors++
			errMsg := fmt.Sprintf("❌ ScoreItem %s: mainValue é zero, não pode dividir", item.ID)
			stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
			log.Println(errMsg)
			continue
		}

		factor := secondaryValue / mainValue

		// Buscar LabTestDefinition pelo código
		var labTestDef models.LabTestDefinition
		if err := db.Where("code = ?", item.LabTestCode).First(&labTestDef).Error; err != nil {
			stats.Errors++
			errMsg := fmt.Sprintf("❌ ScoreItem %s: LabTestDefinition não encontrado para code=%s", item.ID, item.LabTestCode)
			stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
			log.Println(errMsg)
			continue
		}

		// Criar LabTestUnitConversion
		conversion := models.LabTestUnitConversion{
			LabTestDefinitionID: labTestDef.ID,
			MainUnit:            mainUnit,
			SecondaryUnit:       secondaryUnit,
			ConversionFactor:    factor,
		}

		// Tentar criar (ignorar duplicatas)
		if err := db.Create(&conversion).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
				stats.Skipped++
				log.Printf("⏭️  ScoreItem %s: Conversão já existe (code=%s, %s -> %s)", item.ID, item.LabTestCode, secondaryUnit, mainUnit)
			} else {
				stats.Errors++
				errMsg := fmt.Sprintf("❌ ScoreItem %s: Erro ao criar conversão: %v", item.ID, err)
				stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
				log.Println(errMsg)
			}
			continue
		}

		stats.Created++
		conversionKey := fmt.Sprintf("%s: %s -> %s (factor=%.4f)", item.LabTestCode, secondaryUnit, mainUnit, factor)
		stats.UniqueConversions[conversionKey] = true
		log.Printf("✅ Criado: %s", conversionKey)
	}
}

func migrateFromLabTestDefinitions(db *gorm.DB, stats *Stats) {
	var labTestDefs []models.LabTestDefinition
	if err := db.Where("unit_conversion IS NOT NULL AND unit_conversion != ''").Find(&labTestDefs).Error; err != nil {
		log.Fatalf("Failed to fetch lab test definitions: %v", err)
	}

	stats.LabTestDefinitions = len(labTestDefs)
	log.Printf("Encontrados %d LabTestDefinitions com conversão\n", stats.LabTestDefinitions)

	for _, labTestDef := range labTestDefs {
		stats.LTDWithConversion++

		// Parse conversion string
		mainValue, mainUnit, secondaryValue, secondaryUnit, err := parseConversion(*labTestDef.UnitConversion)
		if err != nil {
			stats.Errors++
			errMsg := fmt.Sprintf("❌ LabTestDefinition %s (code=%s): %v", labTestDef.ID, labTestDef.Code, err)
			stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
			log.Println(errMsg)
			continue
		}

		stats.LTDParsed++

		// Calcular fator
		if mainValue == 0 {
			stats.Errors++
			errMsg := fmt.Sprintf("❌ LabTestDefinition %s: mainValue é zero", labTestDef.ID)
			stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
			log.Println(errMsg)
			continue
		}

		factor := secondaryValue / mainValue

		// Criar LabTestUnitConversion
		conversion := models.LabTestUnitConversion{
			LabTestDefinitionID: labTestDef.ID,
			MainUnit:            mainUnit,
			SecondaryUnit:       secondaryUnit,
			ConversionFactor:    factor,
		}

		// Tentar criar (ignorar duplicatas)
		if err := db.Create(&conversion).Error; err != nil {
			if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
				stats.LTDSkipped++
				log.Printf("⏭️  LabTestDef %s: Conversão já existe (code=%s, %s -> %s)", labTestDef.ID, labTestDef.Code, secondaryUnit, mainUnit)
			} else {
				stats.Errors++
				errMsg := fmt.Sprintf("❌ LabTestDef %s: Erro ao criar conversão: %v", labTestDef.ID, err)
				stats.ErrorDetails = append(stats.ErrorDetails, errMsg)
				log.Println(errMsg)
			}
			continue
		}

		stats.LTDCreated++
		conversionKey := fmt.Sprintf("%s: %s -> %s (factor=%.4f)", labTestDef.Code, secondaryUnit, mainUnit, factor)
		stats.UniqueConversions[conversionKey] = true
		log.Printf("✅ Criado: %s", conversionKey)
	}
}

// parseConversion extrai valores e unidades de string "1 g/dL = 10 g/L", "ng/mL × 2,5 = nmol/L" ou "nmol/L ÷ 1000 = μmol/L"
func parseConversion(convStr string) (mainValue float64, mainUnit string, secondaryValue float64, secondaryUnit string, err error) {
	matches := conversionRegex.FindStringSubmatch(convStr)
	if len(matches) != 6 {
		err = fmt.Errorf("formato inválido: '%s'", convStr)
		return
	}

	// Substituir vírgula por ponto para parsing
	mainValueStr := strings.Replace(matches[1], ",", ".", 1)
	mainValue, err = strconv.ParseFloat(mainValueStr, 64)
	if err != nil {
		err = fmt.Errorf("erro ao parsear valor principal '%s': %v", matches[1], err)
		return
	}

	mainUnit = strings.TrimSpace(matches[2])
	operator := matches[3]

	// Substituir vírgula por ponto para parsing
	secondaryValueStr := strings.Replace(matches[4], ",", ".", 1)
	secondaryValue, err = strconv.ParseFloat(secondaryValueStr, 64)
	if err != nil {
		err = fmt.Errorf("erro ao parsear valor secundário '%s': %v", matches[4], err)
		return
	}

	secondaryUnit = strings.TrimSpace(matches[5])

	// Se operador é ÷ (divisão), a conversão é: mainUnit / divisor = secondaryUnit
	// Exemplo: "nmol/L ÷ 1000 = μmol/L" significa nmol/L é a unidade principal
	// Para converter: 1000 nmol/L = 1 μmol/L → factor = 0.001
	if operator == "÷" {
		// Trocar main e secondary para normalizar
		mainValue, secondaryValue = secondaryValue, mainValue
		mainUnit, secondaryUnit = secondaryUnit, mainUnit
	}

	return
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
