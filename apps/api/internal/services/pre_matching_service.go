package services

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PreMatchingService - serviço para fazer matching de exames antes da IA
type PreMatchingService struct {
	db *gorm.DB

	// Padrões regex compilados
	resultNumericPattern  *regexp.Regexp
	resultTextPattern     *regexp.Regexp
	unitPattern           *regexp.Regexp
	referenceRangePattern *regexp.Regexp
	labNamePattern        *regexp.Regexp
	collectionDatePattern *regexp.Regexp
}

// MatchedTest - representa um exame identificado no texto
type MatchedTest struct {
	LabTestDefinitionID uuid.UUID
	TestName            string
	ResultNumeric       *float64
	ResultText          *string
	Unit                *string
	ReferenceRange      *string
	StartPos            int // Posição inicial no texto
	EndPos              int // Posição final no texto
}

// PreMatchingResult - resultado do pré-matching
type PreMatchingResult struct {
	MatchedTests      []MatchedTest
	LaboratoryName    string
	CollectionDate    string
	RemainingText     string
	MatchedCount      int
	RemainingTestsEst int
}

// NewPreMatchingService cria nova instância
func NewPreMatchingService(db *gorm.DB) *PreMatchingService {
	return &PreMatchingService{
		db: db,

		// Resultado numérico: captura números com vírgula ou ponto decimal
		// Ex: "14,90", "95.5", "1.234", "15.430"
		resultNumericPattern: regexp.MustCompile(`(\d{1,3}(?:[.,]\d{3})*(?:[.,]\d{1,2})?)\s*([a-zA-Z/%]+(?:/[a-zA-Z0-9]+)?)`),

		// Resultado texto: palavras capitalizadas (Negativo, Positivo, Normal, etc)
		resultTextPattern: regexp.MustCompile(`\b(Negativo|Positivo|Normal|Alterado|Reagente|Não reagente|Ausente|Presente)\b`),

		// Unidade: captura unidades comuns
		unitPattern: regexp.MustCompile(`\b(g/dL|mg/dL|mU/L|U/L|mmol/L|ng/mL|pg/mL|%|/mm3|mil/mm3|fL|pg)\b`),

		// Faixa de referência: vários formatos
		// Ex: "13,5 a 17,5 g/dL", "70-100 mg/dL", "Inferior a 190", "Superior a 40"
		referenceRangePattern: regexp.MustCompile(`(?i)(?:Ref(?:erência)?[:\s]*)?(?:(\d+[.,]?\d*)\s*(?:a|-)\s*(\d+[.,]?\d*)|(?:Inferior|Superior)\s+a\s+(\d+[.,]?\d*))`),

		// Nome do laboratório
		labNamePattern: regexp.MustCompile(`(?i)(?:Laboratório|Lab\.?):\s*([A-Za-zÀ-ÿ\s\-\.]+?)(?:\n|$)`),

		// Data de coleta: vários formatos brasileiros
		// Ex: "08/01/2026", "08-01-2026", "Coletado em: 08/01/2026"
		collectionDatePattern: regexp.MustCompile(`(?i)Colet(?:a|ado)\s+em:\s*(\d{2}[/-]\d{2}[/-]\d{4})`),
	}
}

// PreMatch - faz pré-matching de exames no texto
func (s *PreMatchingService) PreMatch(text string) (*PreMatchingResult, error) {
	// 1. Carregar definições de exames com altNames
	testDefs, err := s.loadTestDefinitions()
	if err != nil {
		return nil, fmt.Errorf("failed to load test definitions: %v", err)
	}

	// 2. Extrair metadata (laboratório, datas)
	labName := s.extractLabName(text)
	collectionDate := s.extractCollectionDate(text)

	// 3. Buscar matches no texto
	matchedTests := []MatchedTest{}
	textLower := strings.ToLower(text)

	for _, testDef := range testDefs {
		// Parse altNames
		var altNames []string
		if testDef.AltNames != nil && *testDef.AltNames != "" {
			json.Unmarshal([]byte(*testDef.AltNames), &altNames)
		}

		// Buscar cada altName no texto
		for _, altName := range altNames {
			altNameLower := strings.ToLower(altName)
			idx := strings.Index(textLower, altNameLower)

			if idx != -1 {
				// Encontrou! Tentar extrair dados
				matched := s.extractTestData(text, testDef, altName, idx)
				if matched != nil {
					matchedTests = append(matchedTests, *matched)
					break // Já achou, não precisa testar outros altNames
				}
			}
		}
	}

	// 4. Remover seções matched do texto
	remainingText := s.removeMatchedSections(text, matchedTests)

	// 5. Estimar quantos exames restam
	remainingEst := s.estimateRemainingTests(remainingText)

	return &PreMatchingResult{
		MatchedTests:      matchedTests,
		LaboratoryName:    labName,
		CollectionDate:    collectionDate,
		RemainingText:     remainingText,
		MatchedCount:      len(matchedTests),
		RemainingTestsEst: remainingEst,
	}, nil
}

// loadTestDefinitions - carrega definições ativas com altNames
func (s *PreMatchingService) loadTestDefinitions() ([]models.LabTestDefinition, error) {
	var defs []models.LabTestDefinition
	err := s.db.
		Where("is_active = ? AND alt_names IS NOT NULL", true).
		Find(&defs).Error
	return defs, err
}

// extractTestData - extrai dados do exame a partir da posição do match
func (s *PreMatchingService) extractTestData(
	text string,
	testDef models.LabTestDefinition,
	matchedName string,
	startPos int,
) *MatchedTest {
	// Extrair linha completa onde o exame foi encontrado (+ próximas 2 linhas)
	lines := strings.Split(text[startPos:], "\n")
	if len(lines) == 0 {
		return nil
	}

	// Buscar resultado na linha e próximas 2
	maxLines := 3
	if len(lines) < maxLines {
		maxLines = len(lines)
	}
	searchText := strings.Join(lines[:maxLines], "\n")

	// Tentar extrair resultado numérico
	var resultNumeric *float64
	var unit *string
	if matches := s.resultNumericPattern.FindStringSubmatch(searchText); len(matches) >= 3 {
		// Converter número (substituir vírgula por ponto)
		numStr := strings.ReplaceAll(matches[1], ".", "")
		numStr = strings.ReplaceAll(numStr, ",", ".")
		var val float64
		if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
			resultNumeric = &val
			unit = &matches[2]
		}
	}

	// Se não achou numérico, tentar texto
	var resultText *string
	if resultNumeric == nil {
		if matches := s.resultTextPattern.FindStringSubmatch(searchText); len(matches) >= 2 {
			resultText = &matches[1]
		}
	}

	// Se não achou nem numérico nem texto, pular
	if resultNumeric == nil && resultText == nil {
		return nil
	}

	// Tentar extrair faixa de referência
	var refRange *string
	if matches := s.referenceRangePattern.FindStringSubmatch(searchText); len(matches) >= 2 {
		refRange = &matches[0]
	}

	// Calcular endPos (final da seção) - até próximo exame ou 3 linhas
	endPos := startPos + len(searchText)

	return &MatchedTest{
		LabTestDefinitionID: testDef.ID,
		TestName:            matchedName,
		ResultNumeric:       resultNumeric,
		ResultText:          resultText,
		Unit:                unit,
		ReferenceRange:      refRange,
		StartPos:            startPos,
		EndPos:              endPos,
	}
}

// removeMatchedSections - remove seções matched do texto
func (s *PreMatchingService) removeMatchedSections(text string, matches []MatchedTest) string {
	// Ordenar matches por posição (do último para o primeiro para não afetar índices)
	// Usar mapa para marcar caracteres a remover
	toRemove := make(map[int]bool)

	for _, match := range matches {
		for i := match.StartPos; i < match.EndPos && i < len(text); i++ {
			toRemove[i] = true
		}
	}

	// Reconstruir texto sem as seções removidas
	var result strings.Builder
	for i, char := range text {
		if !toRemove[i] {
			result.WriteRune(char)
		}
	}

	return result.String()
}

// extractLabName - extrai nome do laboratório
func (s *PreMatchingService) extractLabName(text string) string {
	if matches := s.labNamePattern.FindStringSubmatch(text); len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

// extractCollectionDate - extrai data de coleta
func (s *PreMatchingService) extractCollectionDate(text string) string {
	if matches := s.collectionDatePattern.FindStringSubmatch(text); len(matches) >= 2 {
		// Converter para formato ISO 8601 (YYYY-MM-DD)
		dateStr := matches[1]
		dateStr = strings.ReplaceAll(dateStr, "-", "/")
		parts := strings.Split(dateStr, "/")
		if len(parts) == 3 {
			return fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
		}
	}
	return ""
}

// estimateRemainingTests - estima quantos exames ainda restam no texto
func (s *PreMatchingService) estimateRemainingTests(text string) int {
	// Heurística: conta linhas que parecem ter resultados
	lines := strings.Split(text, "\n")
	count := 0
	for _, line := range lines {
		if s.resultNumericPattern.MatchString(line) {
			count++
		}
	}
	return count
}

