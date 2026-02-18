package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const EXPORT_DIR = "/app/score-enrich"

func main() {
	godotenv.Load()
	cfg, _ := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Buscar TODAS preparations ordenadas por hierarquia
	var preparations []models.ScoreItemEnrichmentPreparation
	db.Joins("JOIN score_items si ON si.id = score_item_enrichment_preparation.score_item_id").
		Joins("JOIN score_subgroups ssg ON ssg.id = si.subgroup_id").
		Joins("JOIN score_groups sg ON sg.id = ssg.group_id").
		Joins("LEFT JOIN score_items parent_si ON parent_si.id = si.parent_item_id").
		Order("sg.order ASC, ssg.order ASC, COALESCE(parent_si.order, si.order) ASC, si.order ASC").
		Find(&preparations)

	fmt.Printf("🔬 Exporting %d ScoreItem preparations to Markdown...\n\n", len(preparations))

	// Carregar ScoreItems com TODOS os relacionamentos
	var scoreItemIDs []string
	for _, prep := range preparations {
		scoreItemIDs = append(scoreItemIDs, prep.ScoreItemID.String())
	}

	var scoreItems []models.ScoreItem
	db.Preload("Subgroup.Group").
		Preload("ParentItem").
		Where("id IN ?", scoreItemIDs).
		Find(&scoreItems)

	// Mapear ScoreItems por ID
	scoreItemMap := make(map[string]*models.ScoreItem)
	for i := range scoreItems {
		scoreItemMap[scoreItems[i].ID.String()] = &scoreItems[i]
	}

	// Exportar cada preparation (1 arquivo por item)
	for i, prep := range preparations {
		scoreItem := scoreItemMap[prep.ScoreItemID.String()]
		if scoreItem == nil {
			continue
		}

		baseName := fmt.Sprintf("%04d-%s", i+1, slugify(scoreItem.Name))
		exportPromptFile(scoreItem, &prep, baseName, EXPORT_DIR)

		if (i+1)%50 == 0 {
			fmt.Printf("  Exported %d/%d items...\n", i+1, len(preparations))
		}
	}

	fmt.Printf("\n✅ Exported %d items to %s/\n", len(preparations), EXPORT_DIR)
}

func exportPromptFile(si *models.ScoreItem, prep *models.ScoreItemEnrichmentPreparation, baseName, exportDir string) {
	filename := fmt.Sprintf("%s__pending.md", baseName)
	filePath := filepath.Join(exportDir, filename)

	content := generatePromptMarkdown(si, prep)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		log.Printf("❌ Failed to write %s: %v", filename, err)
	}
}

func generatePromptMarkdown(si *models.ScoreItem, prep *models.ScoreItemEnrichmentPreparation) string {
	var sb strings.Builder

	// Header
	fullName := si.GetFullName()
	sb.WriteString(fmt.Sprintf("# ScoreItem: %s\n\n", si.Name))
	sb.WriteString(fmt.Sprintf("**ID:** `%s`\n", si.ID))
	sb.WriteString(fmt.Sprintf("**FullName:** %s\n", fullName))
	if si.Unit != nil {
		sb.WriteString(fmt.Sprintf("**Unit:** %s\n", *si.Unit))
	}
	if si.Gender != nil && *si.Gender != "not_applicable" {
		sb.WriteString(fmt.Sprintf("**Gender:** %s\n", *si.Gender))
	}
	if si.AgeRangeMin != nil {
		sb.WriteString(fmt.Sprintf("**Age Min:** %d anos\n", *si.AgeRangeMin))
	}
	if si.AgeRangeMax != nil {
		sb.WriteString(fmt.Sprintf("**Age Max:** %d anos\n", *si.AgeRangeMax))
	}
	if si.PostMenopause != nil {
		sb.WriteString(fmt.Sprintf("**Post-Menopause:** %v\n", *si.PostMenopause))
	}
	sb.WriteString("\n")

	// Metadata da preparation
	totalChunks := getIntFromJSON(prep.Metadata, "total_chunks")
	articlesCount := getIntFromJSON(prep.Metadata, "articles_count")
	avgSim := getFloatFromJSON(prep.Metadata, "avg_similarity")
	grade := getStringFromJSON(prep.Metadata, "quality_grade")
	newestYear := getIntFromJSON(prep.Metadata, "newest_article_year")

	sb.WriteString("**Preparation Metadata:**\n")
	sb.WriteString(fmt.Sprintf("- Quality Grade: **%s**\n", strings.ToUpper(grade)))
	sb.WriteString(fmt.Sprintf("- Total Chunks: %d de %d artigos\n", totalChunks, articlesCount))
	sb.WriteString(fmt.Sprintf("- Avg Similarity: %.3f\n", avgSim))
	if newestYear > 0 {
		sb.WriteString(fmt.Sprintf("- Newest Article: %d\n", newestYear))
	}
	sb.WriteString("\n---\n\n")

	// Prompt combinado
	sb.WriteString(generateCombinedPrompt(si, prep))

	return sb.String()
}

func generateCombinedPrompt(si *models.ScoreItem, prep *models.ScoreItemEnrichmentPreparation) string {
	var sb strings.Builder

	hasPoints := si.Points != nil && *si.Points > 0

	sb.WriteString("## Contexto\n\n")
	sb.WriteString("Você é um especialista em medicina funcional integrativa e está contribuindo com o **Escore Plenya** — um escore completo de análise de saúde que avalia todos os aspectos da saúde, performance e longevidade humana. Cada ScoreItem representa um parâmetro clínico, laboratorial, genético, comportamental ou histórico que compõe esse escore.\n\n")
	sb.WriteString("Seu papel é gerar conteúdo clínico de alta qualidade para enriquecer cada parâmetro do escore com relevância clínica, orientação ao paciente e conduta prática.\n\n")
	sb.WriteString("**Regras inegociáveis:**\n")
	sb.WriteString("- Use **apenas** o conhecimento médico real consolidado e os dados presentes nos chunks científicos abaixo\n")
	sb.WriteString("- **Não alucine, não invente** dados, estudos, estatísticas ou referências que não estejam nos chunks ou no seu conhecimento médico estabelecido\n")
	sb.WriteString("- Se um dado específico não constar nos chunks e não for do seu conhecimento consolidado, **não o inclua**\n")
	sb.WriteString("- Seja preciso: prefira omitir a inventar\n\n")
	sb.WriteString("## Instrução\n\n")
	sb.WriteString("Com base nos chunks científicos abaixo, gere as respostas em formato JSON.\n\n")
	sb.WriteString(fmt.Sprintf("**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `%s`.**\n\n", si.ID))

	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString(fmt.Sprintf("  \"score_item_id\": \"%s\",\n", si.ID))
	sb.WriteString("  \"clinical_relevance\": \"Texto técnico para médicos (1000-5000 chars): definição fisiológica precisa, valores de referência e interpretação, fisiopatologia resumida, dados epidemiológicos com números concretos, estratificação de risco baseada em evidências.\",\n")
	if hasPoints {
		sb.WriteString("  \"points\": 1,\n")
	} else {
		sb.WriteString("  \"points\": 0,\n")
	}
	sb.WriteString("  \"patient_explanation\": \"Texto simples para pacientes (500-1000 chars): o que é este parâmetro sem jargões, por que é importante para a saúde, o que valores alterados podem significar. Tom tranquilizador e educativo.\",\n")
	sb.WriteString("  \"conduct\": \"Conduta clínica em Markdown (1000-5000 chars): investigação complementar necessária, critérios de encaminhamento a especialistas, intervenções baseadas em evidências. Use bullet points, seções e negrito.\"\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	if hasPoints {
		sb.WriteString("**Regras para `points` (1-50):**\n")
		sb.WriteString("- Baixo impacto clínico: 1-9 pts\n")
		sb.WriteString("- Alto impacto clínico: 10-19 pts\n")
		sb.WriteString("- Alto impacto em mortalidade: 20-50 pts\n")
		sb.WriteString("- Critérios: gravidade/mortalidade (40%), prevalência (30%), intervencionabilidade (30%)\n\n")
	} else {
		sb.WriteString("**Nota:** `points` deve ser `0` neste item — não calcule pontuação.\n\n")
	}

	sb.WriteString("---\n\n")
	sb.WriteString(getChunksContext(si, prep))

	return sb.String()
}

func getChunksContext(si *models.ScoreItem, prep *models.ScoreItemEnrichmentPreparation) string {
	var sb strings.Builder

	sb.WriteString("### Contexto Científico\n\n")
	sb.WriteString(fmt.Sprintf("**ScoreItem:** %s\n", si.GetFullName()))
	if si.Unit != nil {
		sb.WriteString(fmt.Sprintf("**Unidade:** %s\n", *si.Unit))
	}
	if si.Gender != nil && *si.Gender != "not_applicable" {
		sb.WriteString(fmt.Sprintf("**Gênero:** %s\n", *si.Gender))
	}
	if si.AgeRangeMin != nil || si.AgeRangeMax != nil {
		ageRange := ""
		if si.AgeRangeMin != nil {
			ageRange = fmt.Sprintf("%d", *si.AgeRangeMin)
		}
		ageRange += "-"
		if si.AgeRangeMax != nil {
			ageRange += fmt.Sprintf("%d", *si.AgeRangeMax)
		}
		sb.WriteString(fmt.Sprintf("**Faixa Etária:** %s anos\n", ageRange))
	}

	totalChunks := getIntFromJSON(prep.Metadata, "total_chunks")
	articlesCount := getIntFromJSON(prep.Metadata, "articles_count")
	avgSim := getFloatFromJSON(prep.Metadata, "avg_similarity")

	sb.WriteString(fmt.Sprintf("\n**%d chunks de %d artigos (avg similarity: %.3f)**\n\n", totalChunks, articlesCount, avgSim))

	// Parse chunks
	chunksData := prep.SelectedChunks["items"].([]interface{})

	// Incluir TODOS os chunks
	for i, chunkInterface := range chunksData {
		chunkMap := chunkInterface.(map[string]interface{})

		articleTitle := getStringFromMap(chunkMap, "article_title")
		articleYear := getIntFromMap(chunkMap, "article_year")
		journal := getStringFromMap(chunkMap, "journal")
		section := getStringFromMap(chunkMap, "section")
		similarity := getFloatFromMap(chunkMap, "similarity")
		chunkText := getStringFromMap(chunkMap, "chunk_text")

		sb.WriteString(fmt.Sprintf("### Chunk %d/%d\n", i+1, len(chunksData)))
		sb.WriteString(fmt.Sprintf("**Article:** %s", articleTitle))
		if articleYear > 0 {
			sb.WriteString(fmt.Sprintf(" (%d)", articleYear))
		}
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("**Journal:** %s\n", journal))
		sb.WriteString(fmt.Sprintf("**Section:** %s | **Similarity:** %.3f\n\n", section, similarity))
		sb.WriteString(chunkText)
		sb.WriteString("\n\n---\n\n")
	}

	return sb.String()
}

func slugify(s string) string {
	s = strings.ToLower(s)

	replacements := map[string]string{
		"á": "a", "é": "e", "í": "i", "ó": "o", "ú": "u",
		"â": "a", "ê": "e", "ô": "o",
		"ã": "a", "õ": "o", "ç": "c",
	}
	for old, new := range replacements {
		s = strings.ReplaceAll(s, old, new)
	}

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")

	reg = regexp.MustCompile(`-+`)
	s = reg.ReplaceAllString(s, "-")

	s = strings.Trim(s, "-")

	if len(s) > 50 {
		s = s[:50]
	}

	return s
}

// Helper functions for safe JSON type conversion
func getIntFromJSON(m map[string]interface{}, key string) int {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case float64:
			return int(v)
		case int:
			return v
		case json.Number:
			if i, err := v.Int64(); err == nil {
				return int(i)
			}
		}
	}
	return 0
}

func getFloatFromJSON(m map[string]interface{}, key string) float64 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case float64:
			return v
		case int:
			return float64(v)
		case json.Number:
			if f, err := v.Float64(); err == nil {
				return f
			}
		}
	}
	return 0.0
}

func getStringFromJSON(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}

func getIntFromMap(m map[string]interface{}, key string) int {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case float64:
			return int(v)
		case int:
			return v
		case json.Number:
			if i, err := v.Int64(); err == nil {
				return int(i)
			}
		}
	}
	return 0
}

func getFloatFromMap(m map[string]interface{}, key string) float64 {
	if val, ok := m[key]; ok {
		switch v := val.(type) {
		case float64:
			return v
		case int:
			return float64(v)
		case json.Number:
			if f, err := v.Float64(); err == nil {
				return f
			}
		}
	}
	return 0.0
}

func getStringFromMap(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if s, ok := val.(string); ok {
			return s
		}
	}
	return ""
}
