package services

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PreMatchingService - serviço para fazer matching de exames antes da IA
type PreMatchingService struct {
	db *gorm.DB

	// Cache de regex compilados por altName (para word boundary matching)
	altNameRegexCache map[string]*regexp.Regexp

	// Padrões regex compilados
	resultNumericPattern   *regexp.Regexp
	resultWithUnitPattern  *regexp.Regexp
	resultTextPattern      *regexp.Regexp
	unitPattern            *regexp.Regexp
	referenceRangePattern  *regexp.Regexp
	labNamePattern         *regexp.Regexp
	collectionDatePattern  *regexp.Regexp
	collectionDatePattern2 *regexp.Regexp
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
	MatchLength         int // Comprimento do match (para priorizar matches mais longos)
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
		db:                db,
		altNameRegexCache: make(map[string]*regexp.Regexp),

		// Resultado numérico com unidade na mesma linha
		// Formatos: "14,90 g/dL", "95.5 mg/dL", "1.234,56 U/L", "15.430 /mm3"
		// Captura: grupo 1 = número, grupo 2 = unidade
		resultWithUnitPattern: regexp.MustCompile(`(?m)[:.\s]\s*(\d{1,6}(?:[.,]\d{1,3})?)\s*(g/dL|mg/dL|mU/L|U/L|UI/L|UI/mL|mmol/L|umol/L|µmol/L|ng/mL|ng/dL|pg/mL|pg/dL|mcg/dL|ug/dL|µg/dL|mEq/L|mmHg|%|/mm[³3]|mil/mm[³3]|x10[³3]/[uµ]L|fL|pg|mm/h|seg|segundos|min|mL/min|mL/min/1,73m2)(?:\s|$|[,.])`),

		// Resultado numérico simples (fallback)
		// Ex: "14,90", "95.5", "1.234"
		resultNumericPattern: regexp.MustCompile(`(?m)[:.\s]\s*(\d{1,6}(?:[.,]\d{1,3})?)\s*(?:$|\n|[a-zA-Z])`),

		// Resultado texto: palavras capitalizadas (Negativo, Positivo, Normal, etc)
		resultTextPattern: regexp.MustCompile(`(?i)\b(Negativo|Positivo|Normal|Alterado|Reagente|N[aã]o[- ]?reagente|Ausente|Presente|Detectado|N[aã]o[- ]?detectado|Inferior|Superior)\b`),

		// Unidade: captura unidades comuns (expandido para labs brasileiros)
		unitPattern: regexp.MustCompile(`(?i)\b(g/dL|mg/dL|mU/L|U/L|UI/L|UI/mL|mmol/L|umol/L|µmol/L|ng/mL|ng/dL|pg/mL|pg/dL|mcg/dL|ug/dL|µg/dL|mEq/L|mmHg|%|/mm[³3]|mil/mm[³3]|x10[³3]/[uµ]L|fL|pg|mm/h|seg|segundos|min|mL/min|mL/min/1,73m2)\b`),

		// Faixa de referência: vários formatos brasileiros
		// Ex: "13,5 a 17,5", "70 - 100", "Inferior a 190", "Superior a 40", "VR: 13,5 a 17,5"
		referenceRangePattern: regexp.MustCompile(`(?i)(?:V\.?R\.?|Ref(?:er[eê]ncia)?|Valores?\s+de\s+Refer[eê]ncia)?[:\s]*(\d+[.,]?\d*)\s*(?:a|[-–])\s*(\d+[.,]?\d*)|(?:Inferior|Menor|<)\s*(?:a|que)?\s*(\d+[.,]?\d*)|(?:Superior|Maior|>)\s*(?:a|que)?\s*(\d+[.,]?\d*)`),

		// Nome do laboratório - padrões específicos conhecidos + genérico
		labNamePattern: regexp.MustCompile(`(?i)(?:Laborat[oó]rio|Lab\.?)[:\s]+([A-Za-zÀ-ÿ\s\-\.]+?)(?:\n|$)|(?:SABIN|LDB|LABMED|LABIMAGEM|OSWALDO\s*CRUZ|CLINILAB|CLINIMAGEM|UNIMED)`),

		// Data de coleta: vários formatos brasileiros
		// Ex: "08/01/2026", "Coletado em: 08/01/2026", "Data da Coleta: 08/01/2026"
		collectionDatePattern:  regexp.MustCompile(`(?i)(?:Colet(?:a|ado)|Data\s+(?:da\s+)?Coleta)[:\s]+(\d{2}[/-]\d{2}[/-]\d{4})`),
		collectionDatePattern2: regexp.MustCompile(`(?i)(?:Data|Emiss[aã]o)[:\s]+(\d{2}[/-]\d{2}[/-]\d{4})`),
	}
}

// normalizeText aplica lowercase e remove acentos (igual ao modelo LabTestDefinition)
func (s *PreMatchingService) normalizeText(text string) string {
	// Lowercase
	text = strings.ToLower(text)
	// Remove acentos
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, text)
	return result
}

// getAltNameRegex retorna regex compilado com word boundary para o altName
// Usa cache para evitar recompilação
func (s *PreMatchingService) getAltNameRegex(altName string) *regexp.Regexp {
	if regex, ok := s.altNameRegexCache[altName]; ok {
		return regex
	}

	// Escapar caracteres especiais de regex no altName
	escaped := regexp.QuoteMeta(altName)

	// Criar padrão com word boundary flexível
	// Permite match no início de linha, após espaço, após pontuação comum
	pattern := fmt.Sprintf(`(?:^|[\s\n\r,;:()\[\]|/\-])(%s)(?:$|[\s\n\r,;:()\[\]|/\-:])`, escaped)

	regex, err := regexp.Compile(pattern)
	if err != nil {
		// Fallback para busca simples se regex falhar
		regex = regexp.MustCompile(regexp.QuoteMeta(altName))
	}

	s.altNameRegexCache[altName] = regex
	return regex
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

	// 3. Normalizar texto para matching (lowercase + remove acentos)
	// Mantém texto original para extração de valores
	normalizedText := s.normalizeText(text)

	// 4. Preparar candidatos de match (altName -> testDef)
	// Ordenar por comprimento do altName (maior primeiro) para preferir matches mais específicos
	type matchCandidate struct {
		altName string
		testDef models.LabTestDefinition
	}
	var candidates []matchCandidate

	for _, testDef := range testDefs {
		var altNames []string
		if testDef.AltNames != nil && *testDef.AltNames != "" {
			json.Unmarshal([]byte(*testDef.AltNames), &altNames)
		}

		for _, altName := range altNames {
			// altNames já estão normalizados no banco (BeforeSave hook)
			candidates = append(candidates, matchCandidate{
				altName: altName,
				testDef: testDef,
			})
		}
	}

	// Ordenar por comprimento decrescente (matches mais longos primeiro)
	sort.Slice(candidates, func(i, j int) bool {
		return len(candidates[i].altName) > len(candidates[j].altName)
	})

	// 5. Buscar matches no texto
	matchedTests := []MatchedTest{}
	matchedPositions := make(map[int]bool) // Evitar matches sobrepostos

	for _, candidate := range candidates {
		// Usar regex com word boundary
		regex := s.getAltNameRegex(candidate.altName)
		matches := regex.FindAllStringSubmatchIndex(normalizedText, -1)

		for _, match := range matches {
			if len(match) < 4 {
				continue
			}

			// match[2] e match[3] são o grupo de captura (o altName em si)
			startPos := match[2]
			endPos := match[3]

			// Verificar se posição já foi usada por match mais longo
			alreadyMatched := false
			for pos := startPos; pos < endPos; pos++ {
				if matchedPositions[pos] {
					alreadyMatched = true
					break
				}
			}
			if alreadyMatched {
				continue
			}

			// Tentar extrair dados usando texto original (não normalizado)
			extracted := s.extractTestData(text, candidate.testDef, candidate.altName, startPos)
			if extracted != nil {
				extracted.MatchLength = len(candidate.altName)
				matchedTests = append(matchedTests, *extracted)

				// Marcar posições como usadas
				for pos := startPos; pos < endPos; pos++ {
					matchedPositions[pos] = true
				}

				break // Só precisa de um match por candidato
			}
		}
	}

	// 6. Remover seções matched do texto
	remainingText := s.removeMatchedSections(text, matchedTests)

	// 7. Estimar quantos exames restam
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
	// Garantir que startPos está dentro do texto
	if startPos >= len(text) {
		return nil
	}

	// Extrair linha completa onde o exame foi encontrado (+ próximas 3 linhas)
	// Alguns labs colocam o resultado em linhas separadas
	lines := strings.Split(text[startPos:], "\n")
	if len(lines) == 0 {
		return nil
	}

	// Buscar resultado na linha e próximas 3
	maxLines := 4
	if len(lines) < maxLines {
		maxLines = len(lines)
	}
	searchText := strings.Join(lines[:maxLines], "\n")

	// Tentar extrair resultado numérico COM unidade (mais preciso)
	var resultNumeric *float64
	var unit *string

	if matches := s.resultWithUnitPattern.FindStringSubmatch(searchText); len(matches) >= 3 {
		// Converter número brasileiro (1.234,56 -> 1234.56)
		numStr := s.parseNumber(matches[1])
		var val float64
		if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
			resultNumeric = &val
			unitStr := matches[2]
			unit = &unitStr
		}
	}

	// Fallback: tentar resultado numérico sem unidade específica
	if resultNumeric == nil {
		if matches := s.resultNumericPattern.FindStringSubmatch(searchText); len(matches) >= 2 {
			numStr := s.parseNumber(matches[1])
			var val float64
			if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
				resultNumeric = &val

				// Tentar encontrar unidade separadamente
				if unitMatch := s.unitPattern.FindStringSubmatch(searchText); len(unitMatch) >= 2 {
					unit = &unitMatch[1]
				}
			}
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
	if matches := s.referenceRangePattern.FindString(searchText); matches != "" {
		refRange = &matches
	}

	// Calcular endPos (final da seção)
	// Encontrar o fim da seção deste exame (próxima linha em branco ou próximo exame)
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

// parseNumber converte número brasileiro para formato parseável
// Ex: "1.234,56" -> "1234.56", "14,90" -> "14.90", "1234" -> "1234"
func (s *PreMatchingService) parseNumber(numStr string) string {
	// Detectar formato brasileiro (ponto como milhar, vírgula como decimal)
	// ou formato americano (vírgula como milhar, ponto como decimal)

	hasComma := strings.Contains(numStr, ",")
	hasDot := strings.Contains(numStr, ".")

	if hasComma && hasDot {
		// Formato brasileiro: 1.234,56
		// Remove pontos (milhares) e troca vírgula por ponto (decimal)
		numStr = strings.ReplaceAll(numStr, ".", "")
		numStr = strings.ReplaceAll(numStr, ",", ".")
	} else if hasComma {
		// Só vírgula: 14,90 (vírgula é decimal)
		numStr = strings.ReplaceAll(numStr, ",", ".")
	}
	// Se só tem ponto ou nenhum: já está ok

	return numStr
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
	// Tentar padrão genérico primeiro
	if matches := s.labNamePattern.FindStringSubmatch(text); len(matches) >= 2 {
		name := strings.TrimSpace(matches[1])
		if name != "" {
			return name
		}
	}

	// Fallback: detectar labs conhecidos por nome
	textUpper := strings.ToUpper(text)
	knownLabs := []string{
		"SABIN", "LDB", "LABMED", "LABIMAGEM", "OSWALDO CRUZ",
		"CLINILAB", "CLINIMAGEM", "UNIMED", "FLEURY", "DASA",
		"LAVOISIER", "HERMES PARDINI", "DIAGNOSON", "ALTA DIAGNOSTICOS",
	}
	for _, lab := range knownLabs {
		if strings.Contains(textUpper, lab) {
			return lab
		}
	}

	return ""
}

// extractCollectionDate - extrai data de coleta
func (s *PreMatchingService) extractCollectionDate(text string) string {
	// Tentar padrão principal (Coleta/Coletado)
	if matches := s.collectionDatePattern.FindStringSubmatch(text); len(matches) >= 2 {
		return s.formatDateISO(matches[1])
	}

	// Fallback: Data/Emissão
	if matches := s.collectionDatePattern2.FindStringSubmatch(text); len(matches) >= 2 {
		return s.formatDateISO(matches[1])
	}

	return ""
}

// formatDateISO converte data brasileira para ISO 8601
func (s *PreMatchingService) formatDateISO(dateStr string) string {
	dateStr = strings.ReplaceAll(dateStr, "-", "/")
	parts := strings.Split(dateStr, "/")
	if len(parts) == 3 {
		return fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
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

