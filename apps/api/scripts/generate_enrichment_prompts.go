//go:build legacy_scripts
// +build legacy_scripts

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// ChunkData estrutura de um chunk individual
type ChunkData struct {
	ArticleID    string  `json:"article_id"`
	ArticleTitle string  `json:"article_title"`
	ArticleYear  int     `json:"article_year"`
	Journal      string  `json:"journal"`
	ChunkIndex   int     `json:"chunk_index"`
	ChunkText    string  `json:"chunk_text"`
	Section      string  `json:"section"`
	Similarity   float64 `json:"similarity"`
	WordCount    int     `json:"word_count"`
	PageRange    *string `json:"page_range,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run scripts/generate_enrichment_prompts.go <item_number_or_uuid>")
	}

	// Carregar config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Conectar banco
	if err = database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := database.DB

	// Resolver ScoreItem ID
	scoreItemID, err := resolveScoreItemID(db, os.Args[1])
	if err != nil {
		log.Fatalf("Failed to resolve ScoreItem: %v", err)
	}

	// Buscar ScoreItem
	var item models.ScoreItem
	if err := db.Preload("Subgroup.Group").First(&item, scoreItemID).Error; err != nil {
		log.Fatalf("Failed to find ScoreItem: %v", err)
	}

	// Buscar preparation
	var prep models.ScoreItemEnrichmentPreparation
	if err := db.Where("score_item_id = ?", scoreItemID).First(&prep).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("❌ Preparation not found for score_item: %s\nRun: POST /api/v1/score-items/%s/prepare-enrichment", scoreItemID, scoreItemID)
		}
		log.Fatalf("Failed to fetch preparation: %v", err)
	}

	// Parse chunks JSON
	// SelectedChunks tem formato: {"items": [...]}
	chunksBytes, err := json.Marshal(prep.SelectedChunks)
	if err != nil {
		log.Fatalf("Failed to marshal chunks: %v", err)
	}

	var chunksWrapper struct {
		Items []ChunkData `json:"items"`
	}
	if err := json.Unmarshal(chunksBytes, &chunksWrapper); err != nil {
		log.Fatalf("Failed to unmarshal chunks: %v", err)
	}

	chunks := chunksWrapper.Items

	if len(chunks) == 0 {
		log.Fatalf("❌ No chunks found in preparation")
	}

	// Construir contexto de chunks
	chunksContext := buildChunksContext(chunks)

	// Gerar os 4 prompts
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Printf("📊 ScoreItem: %s\n", item.Name)
	fmt.Printf("🔬 Chunks: %d chunks de %d artigos\n", len(chunks), countUniqueArticles(chunks))
	fmt.Printf("✅ Status: %s\n", prep.Status)
	fmt.Println("═══════════════════════════════════════════════════════════════\n")

	// PROMPT 1: clinical_relevance
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("PROMPT 1: CLINICAL_RELEVANCE")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println(buildPromptClinicalRelevance(&item, chunksContext))
	fmt.Println("\n\n")

	// PROMPT 2: patient_explanation
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("PROMPT 2: PATIENT_EXPLANATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println(buildPromptPatientExplanation(&item, chunksContext))
	fmt.Println("\n\n")

	// PROMPT 3: conduct
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("PROMPT 3: CONDUCT")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println(buildPromptConduct(&item, chunksContext))
	fmt.Println("\n\n")

	// PROMPT 4: max_points
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("PROMPT 4: MAX_POINTS")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println(buildPromptMaxPoints(&item, chunksContext))
	fmt.Println("\n\n")

	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("✅ INSTRUÇÕES PARA USO:")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("1. Copie cada prompt individualmente e envie ao Claude")
	fmt.Println("2. Valide os outputs gerados")
	fmt.Println("3. Atualize o score_item via API PUT:")
	fmt.Printf("   PUT /api/v1/score-items/%s\n", scoreItemID)
	fmt.Println("   Body: {\"clinical_relevance\": \"...\", \"patient_explanation\": \"...\", \"conduct\": \"...\", \"points\": 30}")
	fmt.Println("═══════════════════════════════════════════════════════════════")
}

// buildChunksContext formata chunks agrupados por artigo
func buildChunksContext(chunks []ChunkData) string {
	// Agrupar chunks por artigo
	articleGroups := make(map[string][]ChunkData)
	for _, chunk := range chunks {
		articleGroups[chunk.ArticleID] = append(articleGroups[chunk.ArticleID], chunk)
	}

	// Ordenar artigos por similaridade média
	type ArticleGroup struct {
		ArticleID   string
		Chunks      []ChunkData
		AvgSimilarity float64
	}

	var groups []ArticleGroup
	for articleID, articleChunks := range articleGroups {
		var totalSim float64
		for _, chunk := range articleChunks {
			totalSim += chunk.Similarity
		}
		groups = append(groups, ArticleGroup{
			ArticleID:     articleID,
			Chunks:        articleChunks,
			AvgSimilarity: totalSim / float64(len(articleChunks)),
		})
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].AvgSimilarity > groups[j].AvgSimilarity
	})

	// Construir string formatada
	var sb strings.Builder
	for i, group := range groups {
		firstChunk := group.Chunks[0]
		sb.WriteString(fmt.Sprintf("\n=== ARTIGO %d: %s (%s, %d) ===\n\n",
			i+1, firstChunk.ArticleTitle, firstChunk.Journal, firstChunk.ArticleYear))

		// Ordenar chunks por seção (Results → Discussion → Methods → Introduction)
		sort.Slice(group.Chunks, func(i, j int) bool {
			return getSectionPriority(group.Chunks[i].Section) < getSectionPriority(group.Chunks[j].Section)
		})

		// Escrever cada chunk
		for _, chunk := range group.Chunks {
			sb.WriteString(fmt.Sprintf("[%s] (Relevância: %.0f%%)\n",
				strings.ToUpper(chunk.Section), chunk.Similarity*100))
			sb.WriteString(fmt.Sprintf("%s\n\n", chunk.ChunkText))
		}
	}

	return sb.String()
}

func getSectionPriority(section string) int {
	switch section {
	case "results":
		return 1
	case "discussion":
		return 2
	case "methods":
		return 3
	case "introduction":
		return 4
	case "abstract":
		return 5
	default:
		return 6
	}
}

func countUniqueArticles(chunks []ChunkData) int {
	articles := make(map[string]bool)
	for _, chunk := range chunks {
		articles[chunk.ArticleID] = true
	}
	return len(articles)
}

// buildPromptClinicalRelevance gera prompt para clinical_relevance
func buildPromptClinicalRelevance(item *models.ScoreItem, chunksContext string) string {
	unitStr := "N/A"
	if item.Unit != nil {
		unitStr = *item.Unit
	}

	genderStr := "not_applicable"
	if item.Gender != nil {
		genderStr = *item.Gender
	}

	ageRangeStr := formatAgeRange(item.AgeRangeMin, item.AgeRangeMax)

	return fmt.Sprintf(`Você é especialista médico gerando clinical_relevance de alta precisão científica.

PARÂMETRO: %s
UNIDADE: %s
DEMOGRAFIA: %s, idade %s

EVIDÊNCIAS CIENTÍFICAS (chunks dos artigos mais relevantes):
%s

TAREFA: Gere clinical_relevance (1200-1800 caracteres) com:
✓ Definição fisiológica precisa
✓ Valores normais de referência
✓ Fisiopatologia resumida
✓ Dados epidemiológicos com NÚMEROS (RR, OR, HR, NNT, CI 95%%, p-values)
✓ Anos específicos dos estudos citados
✓ Estratificação de risco quantificada
✓ Considerações por população (idade, sexo, comorbidades)

FORMATO OBRIGATÓRIO:
"[Definição clara]. Valores normais: X-Y [unidade]. [Fisiopatologia breve].
Estudos recentes (ANOS) demonstram [DADOS QUANTITATIVOS COM CI E P-VALUES].
Estratificação: [RANGES com significância clínica]. [Considerações populacionais]."

RETORNE APENAS O TEXTO (sem JSON, sem markdown).`,
		item.Name,
		unitStr,
		genderStr,
		ageRangeStr,
		chunksContext,
	)
}

// buildPromptPatientExplanation gera prompt para patient_explanation
func buildPromptPatientExplanation(item *models.ScoreItem, chunksContext string) string {
	return fmt.Sprintf(`Você é educador médico traduzindo ciência para linguagem acessível.

PARÂMETRO: %s

CHUNKS CIENTÍFICOS:
%s

TAREFA: Gere patient_explanation (600-900 caracteres) em linguagem simples:
✓ O QUE é este exame (definição leiga)
✓ POR QUE importa para a saúde
✓ O QUE significam valores alterados
✓ Flesch-Kincaid ≥ 60 (nível 8º ano)
✓ Sem jargão ou com explicação entre parênteses
✓ Tom empático e empoderador

RETORNE APENAS O TEXTO.`,
		item.Name,
		chunksContext,
	)
}

// buildPromptConduct gera prompt para conduct
func buildPromptConduct(item *models.ScoreItem, chunksContext string) string {
	return fmt.Sprintf(`Você é médico assistente gerando protocolo clínico estruturado.

PARÂMETRO: %s

CHUNKS (focus em METHODS e RESULTS):
%s

TAREFA: Gere conduct (1000-1500 caracteres, markdown):

## Interpretação de Valores
[extrair ranges de results/methods]

## Investigação Complementar
[extrair workup de discussion]

## Protocolo de Tratamento
[extrair guidelines/recommendations]

## Critérios de Referência
[extrair indications for specialist]

RETORNE EM MARKDOWN.`,
		item.Name,
		chunksContext,
	)
}

// buildPromptMaxPoints gera prompt para max_points
func buildPromptMaxPoints(item *models.ScoreItem, chunksContext string) string {
	return fmt.Sprintf(`Baseado nas evidências científicas sobre %s:

CHUNKS:
%s

TAREFA: Determine max_points (0-50) considerando:
- Impacto prognóstico (morbidade/mortalidade)
- Modificabilidade (pode ser tratado?)
- Prevalência de alterações
- Custo-efetividade do monitoramento

RETORNE: apenas o número (ex: 35) + justificativa em 1 linha.`,
		item.Name,
		chunksContext,
	)
}

// Helper functions
func resolveScoreItemID(db *gorm.DB, input string) (uuid.UUID, error) {
	if id, err := uuid.Parse(input); err == nil {
		return id, nil
	}

	seqNum, err := strconv.Atoi(input)
	if err != nil {
		return uuid.Nil, fmt.Errorf("input must be a UUID or sequence number")
	}

	var item models.ScoreItem
	offset := seqNum - 1
	if offset < 0 {
		return uuid.Nil, fmt.Errorf("sequence number must be >= 1")
	}

	err = db.Order("created_at ASC").Offset(offset).Limit(1).First(&item).Error
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to find ScoreItem #%d: %w", seqNum, err)
	}

	return item.ID, nil
}

func formatAgeRange(min, max *int) string {
	if min == nil && max == nil {
		return "all ages"
	}
	if min != nil && max == nil {
		return fmt.Sprintf("≥%d years", *min)
	}
	if min == nil && max != nil {
		return fmt.Sprintf("≤%d years", *max)
	}
	return fmt.Sprintf("%d-%d years", *min, *max)
}
