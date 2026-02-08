package services

import (
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
		for _, altName := range testDef.AltNames {
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

// findOriginalPosition encontra a posição real no texto original
// dado um altName e uma posição aproximada do texto normalizado
func (s *PreMatchingService) findOriginalPosition(text string, altName string, approxPos int) int {
	// Buscar em uma janela ao redor da posição aproximada
	windowSize := 200 // caracteres antes e depois
	searchStart := approxPos - windowSize
	if searchStart < 0 {
		searchStart = 0
	}
	searchEnd := approxPos + windowSize + len(altName)
	if searchEnd > len(text) {
		searchEnd = len(text)
	}

	searchRegion := text[searchStart:searchEnd]
	normalizedRegion := s.normalizeText(searchRegion)

	// Encontrar altName na região normalizada
	idx := strings.Index(normalizedRegion, altName)
	if idx >= 0 {
		// Mapear de volta para o texto original
		// Contar caracteres no texto original até chegar na mesma posição lógica
		originalIdx := 0
		normalizedIdx := 0
		for originalIdx < len(searchRegion) && normalizedIdx < idx {
			// Avançar um rune no original
			_, size := getRuneAt(searchRegion, originalIdx)
			originalIdx += size
			// O normalizado avança 1 char (sem acentos)
			normalizedIdx++
		}
		return searchStart + originalIdx
	}

	return approxPos // fallback
}

// getRuneAt retorna o rune e seu tamanho em bytes na posição dada
func getRuneAt(s string, pos int) (rune, int) {
	for i, r := range s {
		if i == pos {
			return r, len(string(r))
		}
		if i > pos {
			break
		}
	}
	return 0, 1
}

// extractTestData - extrai dados do exame a partir da posição do match
// ESTRATÉGIA: Priorizar padrões mais específicos primeiro
// 1. RESULTADO: valor unidade (linha dedicada) - MAIOR PRIORIDADE
// 2. NOME: valor unidade (inline com : logo após nome)
// 3. NOME valor unidade refs (tabular hemograma)
func (s *PreMatchingService) extractTestData(
	text string,
	testDef models.LabTestDefinition,
	matchedName string,
	startPos int,
) *MatchedTest {
	// IMPORTANTE: startPos vem do texto normalizado, precisamos encontrar
	// a posição real no texto original
	originalPos := s.findOriginalPosition(text, matchedName, startPos)
	if originalPos >= len(text) {
		return nil
	}

	// Encontrar início da linha onde o match ocorreu
	lineStart := originalPos
	for lineStart > 0 && text[lineStart-1] != '\n' {
		lineStart--
	}

	// Extrair as próximas 8 linhas a partir do início da linha do match
	lines := strings.Split(text[lineStart:], "\n")
	if len(lines) == 0 {
		return nil
	}

	// Limitar a 8 linhas para busca
	maxLines := 8
	if len(lines) < maxLines {
		maxLines = len(lines)
	}

	var resultNumeric *float64
	var unit *string
	var resultText *string
	endPos := originalPos + len(matchedName)

	// ============================================================
	// PADRÃO 1: Buscar "Resultado: valor unidade" nas próximas linhas
	// Formato: Resultado: 78 µg/dL ou Resultado: 86 mg/dL Jejum:
	// IMPORTANTE: Parar no primeiro valor após "Resultado:", antes de qualquer outro ":"
	// ============================================================
	resultadoPattern := regexp.MustCompile(`(?i)Resultado:\s*(\d+[.,]?\d*)\s*([a-zA-Zµμ/%³²]+(?:/[a-zA-Z0-9µμ.,³²m]+)?)`)
	for i := 0; i < maxLines; i++ {
		line := lines[i]
		if matches := resultadoPattern.FindStringSubmatch(line); len(matches) >= 3 {
			numStr := s.parseNumber(matches[1])
			var val float64
			if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
				resultNumeric = &val
				unitStr := matches[2]
				unit = &unitStr
				endPos = lineStart + s.findLineEnd(lines, i)
				break
			}
		}
	}

	// ============================================================
	// PADRÃO 1B: "Resultado: Inferior a X" ou "Resultado: Acima de X"
	// Formato: Resultado: Inferior a 0,10 mg/dL ou Resultado: Acima de 60 mL/min
	// ============================================================
	if resultNumeric == nil {
		resultadoQualPattern := regexp.MustCompile(`(?i)Resultado:\s*(Inferior|Acima|Menor|Maior)\s*(?:a|de|que)?\s*(\d+[.,]?\d*)\s*([a-zA-Zµμ/%³²]+(?:/[a-zA-Z0-9µμ.,³²m]+)?)`)
		for i := 0; i < maxLines; i++ {
			line := lines[i]
			if matches := resultadoQualPattern.FindStringSubmatch(line); len(matches) >= 4 {
				qualifier := matches[1]
				numStr := s.parseNumber(matches[2])
				var val float64
				if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
					resultNumeric = &val
					unitStr := matches[3]
					unit = &unitStr
					// Guardar o qualificador no resultText
					qualText := strings.ToLower(qualifier)
					if qualText == "inferior" || qualText == "menor" {
						qualText = "<"
					} else {
						qualText = ">"
					}
					resultText = &qualText
					endPos = lineStart + s.findLineEnd(lines, i)
					break
				}
			}
		}
	}

	// ============================================================
	// PADRÃO 2: "NOME: valor unidade" na mesma linha do match
	// Formato: Hemácias: 4,13 milhões/mm ou COLESTEROL TOTAL: 205 mg/dL
	// Também suporta: Hemoglobina.........: 12,30 g/dL
	// IMPORTANTE: EXIGE ":" (possivelmente precedido de pontos) após o nome
	// ============================================================
	if resultNumeric == nil {
		firstLine := lines[0]
		normalizedLine := s.normalizeText(firstLine)
		normalizedName := s.normalizeText(matchedName)
		nameIdx := strings.Index(normalizedLine, normalizedName)

		if nameIdx >= 0 {
			// Calcular posição após o nome no texto original
			afterNamePos := nameIdx + len(normalizedName)
			if afterNamePos < len(firstLine) {
				afterName := firstLine[afterNamePos:]
				// Padrão: permite pontos antes de ":" seguido de espaços e número
				// ^[.]*:\s* aceita ":" ou ".....:" após o nome
				inlinePattern := regexp.MustCompile(`^[.]*:\s+(\d+[.,]?\d*)\s*([a-zA-Zµμ/%³²]+(?:/[a-zA-Z0-9µμ.,³²m]+)?)`)
				if matches := inlinePattern.FindStringSubmatch(afterName); len(matches) >= 3 {
					numStr := s.parseNumber(matches[1])
					var val float64
					if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
						resultNumeric = &val
						unitStr := matches[2]
						unit = &unitStr
						endPos = lineStart + len(firstLine)
					}
				}
			}
		}
	}

	// ============================================================
	// PADRÃO 3: Tabular (hemograma PDF1) - "NOME valor unidade refs..."
	// Formato: HEMOGLOBINA 12,2 g/dL 13,0 a 16,5...
	// O valor está LOGO APÓS o nome do exame na mesma linha (SEM :)
	// Verifica que não há ":" entre o nome e o valor
	// ============================================================
	if resultNumeric == nil {
		firstLine := lines[0]
		normalizedLine := s.normalizeText(firstLine)
		normalizedName := s.normalizeText(matchedName)
		nameIdx := strings.Index(normalizedLine, normalizedName)
		if nameIdx >= 0 {
			// Pegar texto após o nome
			afterNamePos := nameIdx + len(normalizedName)
			if afterNamePos < len(firstLine) {
				afterName := firstLine[afterNamePos:]
				// Verificar que NÃO começa com ":" (senão seria Padrão 2)
				if len(afterName) > 0 && afterName[0] != ':' {
					// Buscar espaço(s) seguido de número com unidade
					tabularPattern := regexp.MustCompile(`^\s+(\d+[.,]?\d*)\s*([a-zA-Zµμ/%³²]+(?:/[a-zA-Z0-9µμ.,³²m]+)?)`)
					if matches := tabularPattern.FindStringSubmatch(afterName); len(matches) >= 3 {
						numStr := s.parseNumber(matches[1])
						var val float64
						if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
							resultNumeric = &val
							unitStr := matches[2]
							unit = &unitStr
							endPos = lineStart + len(firstLine)
						}
					}
				}
			}
		}
	}

	// ============================================================
	// PADRÃO 4: HOMA e similares - "NOME: valor" ou "NOME valor" em linha curta
	// Formato: HOMA IR: 0,8 ou Glicemia média estimada: 105 mg/dL
	// IMPORTANTE: Só aplica se a linha for curta (sem referências longas)
	// ============================================================
	if resultNumeric == nil {
		firstLine := lines[0]
		normalizedLine := s.normalizeText(firstLine)
		normalizedName := s.normalizeText(matchedName)
		nameIdx := strings.Index(normalizedLine, normalizedName)

		if nameIdx >= 0 {
			// Pegar texto após o nome
			afterNamePos := nameIdx + len(normalizedName)
			if afterNamePos < len(firstLine) {
				afterName := firstLine[afterNamePos:]
				// Só aplicar se afterName for curto (< 50 chars) - evita pegar de linhas com referências
				if len(strings.TrimSpace(afterName)) < 50 {
					// Buscar número próximo do início do afterName (não no final da linha)
					homaPattern := regexp.MustCompile(`^[:\s]+(\d+[.,]?\d*)\s*([a-zA-Zµμ/%³²]*(?:/[a-zA-Z0-9µμ.,³²m]*)?)`)
					if matches := homaPattern.FindStringSubmatch(afterName); len(matches) >= 2 {
						numStr := s.parseNumber(matches[1])
						var val float64
						if _, err := fmt.Sscanf(numStr, "%f", &val); err == nil {
							resultNumeric = &val
							if len(matches) >= 3 && matches[2] != "" {
								unitStr := matches[2]
								unit = &unitStr
							}
							endPos = lineStart + len(firstLine)
						}
					}
				}
			}
		}
	}

	// Se não encontrou numérico, tentar texto qualitativo
	if resultNumeric == nil && resultText == nil {
		for i := 0; i < min(3, maxLines); i++ {
			if matches := s.resultTextPattern.FindStringSubmatch(lines[i]); len(matches) >= 2 {
				resultText = &matches[1]
				endPos = lineStart + s.findLineEnd(lines, i)
				break
			}
		}
	}

	// Se não encontrou nada, retornar nil
	if resultNumeric == nil && resultText == nil {
		return nil
	}

	return &MatchedTest{
		LabTestDefinitionID: testDef.ID,
		TestName:            matchedName,
		ResultNumeric:       resultNumeric,
		ResultText:          resultText,
		Unit:                unit,
		StartPos:            originalPos,
		EndPos:              endPos,
	}
}

// findLineEnd retorna a posição final da linha i no slice de linhas
func (s *PreMatchingService) findLineEnd(lines []string, lineIdx int) int {
	pos := 0
	for i := 0; i <= lineIdx && i < len(lines); i++ {
		pos += len(lines[i])
		if i < lineIdx {
			pos++ // +1 para o \n
		}
	}
	return pos
}

// isGraphLine detecta linhas que são dados de gráficos históricos (a ignorar)
// Exemplos: "250", "97 98 78", "07/02/24 27/03/25", "08:45:26 09:54:38"
func (s *PreMatchingService) isGraphLine(line string) bool {
	line = strings.TrimSpace(line)
	if line == "" {
		return true
	}

	// Linha contém apenas números e espaços (valores de gráfico)
	onlyNumbersSpaces := regexp.MustCompile(`^[\d\s.,]+$`)
	if onlyNumbersSpaces.MatchString(line) {
		return true
	}

	// Linha contém datas no formato dd/mm/yy ou dd/mm/yyyy
	datePattern := regexp.MustCompile(`^\d{2}/\d{2}/\d{2,4}(\s+\d{2}/\d{2}/\d{2,4})*\s*$`)
	if datePattern.MatchString(line) {
		return true
	}

	// Linha contém horários no formato hh:mm:ss
	timePattern := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}(\s+\d{2}:\d{2}:\d{2})*\s*$`)
	if timePattern.MatchString(line) {
		return true
	}

	return false
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

