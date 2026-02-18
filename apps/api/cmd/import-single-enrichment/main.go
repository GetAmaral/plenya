package main

import (
	"bufio"
	"encoding/json"
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

type EnrichmentResult struct {
	ScoreItemID        uuid.UUID
	ClinicalRelevance  string
	PatientExplanation string
	Conduct            string
	Points             float64
	PointsJustification string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename.md>")
		fmt.Println("Example: go run main.go 001-leptina-mulheres-DONE.md")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Determinar caminho completo
	var filepath string
	if strings.HasPrefix(filename, "/") {
		filepath = filename
	} else {
		filepath = findFile(filename)
		if filepath == "" {
			log.Fatalf("❌ File not found: %s", filename)
		}
	}

	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("📖 Reading file: %s\n", filepath)

	// Parse MD file
	result, err := parseMarkdownFile(filepath)
	if err != nil {
		log.Fatalf("❌ Failed to parse file: %v", err)
	}

	// Validar campos
	if err := validateResult(result); err != nil {
		log.Fatalf("❌ Validation failed: %v", err)
	}

	// Buscar ScoreItem
	var scoreItem models.ScoreItem
	if err := db.First(&scoreItem, result.ScoreItemID).Error; err != nil {
		log.Fatalf("❌ ScoreItem not found: %v", err)
	}

	fmt.Printf("✅ Parsed successfully: %s\n", scoreItem.Name)
	fmt.Printf("   - Clinical Relevance: %d chars\n", len(result.ClinicalRelevance))
	fmt.Printf("   - Patient Explanation: %d chars\n", len(result.PatientExplanation))
	fmt.Printf("   - Conduct: %d chars\n", len(result.Conduct))
	fmt.Printf("   - Points: %.0f\n", result.Points)

	// Atualizar banco
	fmt.Printf("\n💾 Updating database...\n")

	now := time.Now()
	updates := map[string]interface{}{
		"clinical_relevance":  result.ClinicalRelevance,
		"patient_explanation": result.PatientExplanation,
		"conduct":             result.Conduct,
		"points":              result.Points,
		"last_review":         &now,
	}

	if err := db.Model(&scoreItem).Updates(updates).Error; err != nil {
		log.Fatalf("❌ Database update failed: %v", err)
	}

	fmt.Printf("✅ Database updated successfully!\n")

	// Atualizar preparation status
	if err := db.Exec(`
		UPDATE score_item_enrichment_preparation
		SET status = 'completed'
		WHERE score_item_id = ?
	`, result.ScoreItemID).Error; err != nil {
		log.Printf("⚠️  Failed to update preparation status: %v", err)
	}

	// Mover arquivo para completed/ (se estiver em in-progress/)
	if strings.Contains(filepath, "/in-progress/") {
		newPath := strings.Replace(filepath, "/in-progress/", "/completed/", 1)
		if !strings.HasSuffix(newPath, "-DONE.md") {
			newPath = strings.TrimSuffix(newPath, ".md") + "-DONE.md"
		}
		if err := os.Rename(filepath, newPath); err != nil {
			log.Printf("⚠️  Failed to move file: %v", err)
		} else {
			fmt.Printf("📦 File moved to: %s\n", newPath)
		}
	}

	// Atualizar progress
	updateProgress()

	fmt.Printf("\n🎉 Import completed successfully!\n")
}

func findFile(filename string) string {
	// Procurar em in-progress/ e completed/
	dirs := []string{"in-progress", "completed"}
	for _, dir := range dirs {
		path := filepath.Join(EXPORT_DIR, dir, filename)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

func parseMarkdownFile(filepath string) (*EnrichmentResult, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Aumentar buffer para linhas longas
	const maxCapacity = 1024 * 1024 // 1MB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	result := &EnrichmentResult{}

	// Regex para extrair UUID
	uuidRegex := regexp.MustCompile(`\*\*ID:\*\*\s*` + "`" + `([a-f0-9-]+)` + "`")

	// Estados de parsing
	var (
		inClinicalRelevance  = false
		inPatientExplanation = false
		inConduct            = false
		inPoints             = false
		inCodeBlock          = false
	)

	var currentResponse strings.Builder

	for scanner.Scan() {
		line := scanner.Text()

		// Extrair UUID
		if matches := uuidRegex.FindStringSubmatch(line); len(matches) > 1 {
			result.ScoreItemID = uuid.MustParse(matches[1])
			continue
		}

		// Detectar início de seções de resposta
		if strings.Contains(line, "## 🔬 PROMPT 1") {
			inClinicalRelevance = true
			inPatientExplanation = false
			inConduct = false
			inPoints = false
			continue
		}
		if strings.Contains(line, "## 👥 PROMPT 2") {
			inClinicalRelevance = false
			inPatientExplanation = true
			inConduct = false
			inPoints = false
			continue
		}
		if strings.Contains(line, "## 📋 PROMPT 3") {
			inClinicalRelevance = false
			inPatientExplanation = false
			inConduct = true
			inPoints = false
			continue
		}
		if strings.Contains(line, "## 🎯 PROMPT 4") {
			inClinicalRelevance = false
			inPatientExplanation = false
			inConduct = false
			inPoints = true
			continue
		}

		// Detectar code blocks
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				// Fim do code block - salvar resposta
				response := strings.TrimSpace(currentResponse.String())

				if inClinicalRelevance {
					result.ClinicalRelevance = response
				} else if inPatientExplanation {
					result.PatientExplanation = response
				} else if inConduct {
					result.Conduct = response
				} else if inPoints {
					parsePointsResponse(response, result)
				}

				currentResponse.Reset()
				inCodeBlock = false
			} else {
				inCodeBlock = true
			}
			continue
		}

		// Coletar conteúdo dentro de code blocks
		if inCodeBlock {
			// Ignorar linhas de placeholder
			if strings.Contains(line, "[Claude escreve aqui") {
				continue
			}
			if currentResponse.Len() > 0 {
				currentResponse.WriteString("\n")
			}
			currentResponse.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func parsePointsResponse(response string, result *EnrichmentResult) {
	lines := strings.Split(response, "\n")

	for _, line := range lines {
		// Procurar linha com pontuação
		if strings.HasPrefix(strings.TrimSpace(line), "Pontuação:") {
			parts := strings.Split(line, ":")
			if len(parts) >= 2 {
				pointsStr := strings.TrimSpace(parts[1])
				// Remover caracteres não numéricos
				pointsStr = regexp.MustCompile(`[^0-9.]`).ReplaceAllString(pointsStr, "")
				if points, err := strconv.ParseFloat(pointsStr, 64); err == nil {
					result.Points = points
				}
			}
		}
	}

	result.PointsJustification = response
}

func validateResult(result *EnrichmentResult) error {
	if result.ScoreItemID == uuid.Nil {
		return fmt.Errorf("missing ScoreItem ID")
	}

	if len(result.ClinicalRelevance) < 500 {
		return fmt.Errorf("ClinicalRelevance too short: %d chars (min 500)", len(result.ClinicalRelevance))
	}
	if len(result.ClinicalRelevance) > 5000 {
		return fmt.Errorf("ClinicalRelevance too long: %d chars (max 5000)", len(result.ClinicalRelevance))
	}

	if len(result.PatientExplanation) < 300 {
		return fmt.Errorf("PatientExplanation too short: %d chars (min 300)", len(result.PatientExplanation))
	}
	if len(result.PatientExplanation) > 3000 {
		return fmt.Errorf("PatientExplanation too long: %d chars (max 3000)", len(result.PatientExplanation))
	}

	if len(result.Conduct) < 500 {
		return fmt.Errorf("Conduct too short: %d chars (min 500)", len(result.Conduct))
	}
	if len(result.Conduct) > 5000 {
		return fmt.Errorf("Conduct too long: %d chars (max 5000)", len(result.Conduct))
	}

	if result.Points < 0 || result.Points > 50 {
		return fmt.Errorf("Points out of range: %.0f (must be 0-50)", result.Points)
	}

	return nil
}

func updateProgress() {
	progressFile := filepath.Join(EXPORT_DIR, ".progress.json")

	data, err := os.ReadFile(progressFile)
	if err != nil {
		return
	}

	var progress map[string]interface{}
	json.Unmarshal(data, &progress)

	completed := int(progress["completed"].(float64))
	progress["completed"] = completed + 1
	progress["last_updated"] = time.Now().Format(time.RFC3339)

	// Contar arquivos completed
	files, _ := filepath.Glob(filepath.Join(EXPORT_DIR, "completed", "*-DONE.md"))
	progress["completed"] = len(files)

	newData, _ := json.MarshalIndent(progress, "", "  ")
	os.WriteFile(progressFile, newData, 0644)
}
