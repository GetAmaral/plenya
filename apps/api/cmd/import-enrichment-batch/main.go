package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const EXPORT_DIR = "/home/user/plenya/enrichment-batch"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <item-prefix>")
		fmt.Println("Example: go run main.go 0001-controle-de-diabetes-mellitus")
		os.Exit(1)
	}

	itemPrefix := os.Args[1]

	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Buscar os 4 arquivos __processed
	clinicalFile := filepath.Join(EXPORT_DIR, itemPrefix+"__clinical__processed.md")
	patientFile := filepath.Join(EXPORT_DIR, itemPrefix+"__patient__processed.md")
	conductFile := filepath.Join(EXPORT_DIR, itemPrefix+"__conduct__processed.md")
	pointsFile := filepath.Join(EXPORT_DIR, itemPrefix+"__points__processed.md")

	// Verificar se todos existem
	for _, file := range []string{clinicalFile, patientFile, conductFile, pointsFile} {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Fatalf("❌ File not found: %s", file)
		}
	}

	fmt.Printf("📖 Reading 4 processed files for: %s\n\n", itemPrefix)

	// Parse os 4 arquivos
	scoreItemID, clinical := parseMarkdownFile(clinicalFile)
	_, patient := parseMarkdownFile(patientFile)
	_, conduct := parseMarkdownFile(conductFile)
	_, pointsText := parseMarkdownFile(pointsFile)

	points := parsePoints(pointsText)

	// Validar
	if scoreItemID == uuid.Nil {
		log.Fatal("❌ Failed to extract ScoreItem ID")
	}

	fmt.Printf("✅ Parsed successfully\n")
	fmt.Printf("   - Clinical Relevance: %d chars\n", len(clinical))
	fmt.Printf("   - Patient Explanation: %d chars\n", len(patient))
	fmt.Printf("   - Conduct: %d chars\n", len(conduct))
	fmt.Printf("   - Points: %.0f\n\n", points)

	// Buscar ScoreItem
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, scoreItemID).Error; err != nil {
		log.Fatalf("❌ ScoreItem not found: %v", err)
	}

	// Atualizar banco
	fmt.Printf("💾 Updating database for: %s\n", scoreItem.Name)

	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  clinical,
		"patient_explanation": patient,
		"conduct":             conduct,
		"points":              points,
		"last_review":         &now,
	}

	if err := db.Model(&scoreItem).Updates(updates).Error; err != nil {
		log.Fatalf("❌ Database update failed: %v", err)
	}

	fmt.Printf("✅ Database updated successfully!\n\n")

	// Atualizar preparation status
	db.Exec(`UPDATE score_item_enrichment_preparation SET status = 'completed' WHERE score_item_id = ?`, scoreItemID)

	// Renomear os 4 arquivos para __imported
	for _, file := range []string{clinicalFile, patientFile, conductFile, pointsFile} {
		newFile := strings.Replace(file, "__processed.md", "__imported.md", 1)
		if err := os.Rename(file, newFile); err != nil {
			log.Printf("⚠️  Failed to rename %s: %v", file, err)
		}
	}

	fmt.Printf("📦 Files renamed to __imported\n")
	fmt.Printf("🎉 Import completed successfully!\n")
}

func parseMarkdownFile(filepath string) (uuid.UUID, string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("❌ Failed to open %s: %v", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var scoreItemID uuid.UUID
	var response strings.Builder
	inCodeBlock := false

	uuidRegex := regexp.MustCompile(`\*\*ID:\*\*\s*` + "`" + `([a-f0-9-]+)` + "`")

	for scanner.Scan() {
		line := scanner.Text()

		// Extrair UUID
		if matches := uuidRegex.FindStringSubmatch(line); len(matches) > 1 {
			scoreItemID = uuid.MustParse(matches[1])
			continue
		}

		// Detectar code blocks
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				inCodeBlock = false
			} else {
				inCodeBlock = true
			}
			continue
		}

		// Coletar resposta dentro de code block
		if inCodeBlock {
			if strings.Contains(line, "[Gerar aqui]") {
				continue
			}
			if response.Len() > 0 {
				response.WriteString("\n")
			}
			response.WriteString(line)
		}
	}

	return scoreItemID, strings.TrimSpace(response.String())
}

func parsePoints(text string) float64 {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "Pontuação:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				pointsStr := strings.TrimSpace(parts[1])
				pointsStr = regexp.MustCompile(`[^0-9.]`).ReplaceAllString(pointsStr, "")
				if points, err := strconv.ParseFloat(pointsStr, 64); err == nil {
					return points
				}
			}
		}
	}
	return 0
}
