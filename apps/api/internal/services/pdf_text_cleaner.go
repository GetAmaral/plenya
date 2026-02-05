package services

import (
	"regexp"
	"strings"
)

// PDFTextCleaner - serviço para limpar/processar texto extraído de PDFs de laudos
type PDFTextCleaner struct {
	// Padrões compilados para performance
	dotsPattern           *regexp.Regexp
	pageNumberPattern     *regexp.Regexp
	signaturePattern      *regexp.Regexp
	hashPattern           *regexp.Regexp
	disclaimerPattern     *regexp.Regexp
	patientHeaderPattern  *regexp.Regexp
	previousResultsPattern *regexp.Regexp
	urlPattern            *regexp.Regexp
	notePattern           *regexp.Regexp
	materialMethodPattern *regexp.Regexp
}

// NewPDFTextCleaner cria uma nova instância do cleaner
func NewPDFTextCleaner() *PDFTextCleaner {
	return &PDFTextCleaner{
		// Remove sequências de pontos usados para espaçamento (ex: "......:")
		dotsPattern: regexp.MustCompile(`\.{2,}[:.\s]*`),

		// Remove números de página (ex: "Pág.: 1 / 18")
		pageNumberPattern: regexp.MustCompile(`(?i)Pág\.?:\s*\d+\s*/\s*\d+`),

		// Remove linhas de assinatura/responsável técnico
		signaturePattern: regexp.MustCompile(`(?i)Responsável Técnico.{0,200}?(CRF|CRM|CRBM).{0,100}?Assinado eletronicamente.{0,200}$`),

		// Remove hashes de assinatura eletrônica
		hashPattern: regexp.MustCompile(`(?i)Este exame possui assinatura eletrônica avançada:\s*[a-f0-9]{60,}`),

		// Remove disclaimers legais
		disclaimerPattern: regexp.MustCompile(`(?i)A aceitação deste resultado está condicionada.{0,300}?\.br`),

		// Remove cabeçalho repetido do paciente em cada página
		patientHeaderPattern: regexp.MustCompile(`(?i)Nome\.+:\s*.+?Entrada\.+:.+?Requisiç\.+:.+?Sexo\.+:.+?Convênio\.+:.+?DN\.+:.+?Médico\.+:.+?Idade\.+:.+?RG\.+:.+?Impresso\.+:.+?CPF\.+:.+?Lic\.Sanit\.+:.+?`),

		// Remove seções de "Resultados Anteriores" completas
		previousResultsPattern: regexp.MustCompile(`(?i)Resultados Anteriores:.*?(?:\n\n|$)`),

		// Remove URLs
		urlPattern: regexp.MustCompile(`https?://[^\s]+|www\.[^\s]+`),

		// Remove notas longas explicativas (mantém apenas título da nota)
		notePattern: regexp.MustCompile(`(?i)Nota\s*\d*\s*:.*?(?:\n\n|$)`),

		// Simplifica linhas Material/Método (mantém resumido)
		materialMethodPattern: regexp.MustCompile(`(?i)Material:\s*(.+?)\s*Coletado em:\s*(\S+)\s*Método:\s*(.+?)$`),
	}
}

// CleanText - processa e limpa o texto completo extraído do PDF
// Retorna versão reduzida focada apenas em dados essenciais para IA
func (c *PDFTextCleaner) CleanText(fullText string) string {
	text := fullText

	// 1. Remove caracteres de controle (form feed, etc)
	text = strings.ReplaceAll(text, "\x0C", "\n")
	text = strings.ReplaceAll(text, "\r", "")

	// 2. Remove padrões identificados (em ordem de mais específico para mais geral)
	text = c.hashPattern.ReplaceAllString(text, "")
	text = c.signaturePattern.ReplaceAllString(text, "")
	text = c.disclaimerPattern.ReplaceAllString(text, "")
	text = c.pageNumberPattern.ReplaceAllString(text, "")
	text = c.urlPattern.ReplaceAllString(text, "")
	text = c.patientHeaderPattern.ReplaceAllString(text, "")
	text = c.previousResultsPattern.ReplaceAllString(text, "")

	// 3. Simplifica Material/Método (mantém apenas o essencial)
	text = c.materialMethodPattern.ReplaceAllString(text, "Material: $1, Método: $3")

	// 4. Remove notas longas (mantém apenas "Nota:" sem conteúdo)
	// Isso remove explicações longas mas mantém indicadores de notas importantes
	text = c.notePattern.ReplaceAllString(text, "")

	// 5. Remove sequências de pontos usados para espaçamento
	text = c.dotsPattern.ReplaceAllString(text, ": ")

	// 6. Limpa espaçamento excessivo
	// Remove múltiplos espaços
	text = regexp.MustCompile(`[ \t]+`).ReplaceAllString(text, " ")

	// Remove múltiplas linhas vazias (mantém máximo 2)
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")

	// 7. Remove linhas que são apenas espaços/símbolos
	lines := strings.Split(text, "\n")
	var cleanedLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Mantém linha se tem pelo menos 3 caracteres alfanuméricos
		if len(trimmed) > 0 && countAlphanumeric(trimmed) >= 3 {
			cleanedLines = append(cleanedLines, trimmed)
		}
	}
	text = strings.Join(cleanedLines, "\n")

	// 8. Remove símbolos especiais isolados
	text = regexp.MustCompile(`(?m)^\s*[+*-]+\s*$`).ReplaceAllString(text, "")

	// 9. Final trim
	text = strings.TrimSpace(text)

	return text
}

// countAlphanumeric conta caracteres alfanuméricos em uma string
func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return count
}

// GetCompressionStats - retorna estatísticas de compressão
func (c *PDFTextCleaner) GetCompressionStats(original, cleaned string) map[string]interface{} {
	origLines := len(strings.Split(original, "\n"))
	cleanLines := len(strings.Split(cleaned, "\n"))

	reduction := float64(len(original)-len(cleaned)) / float64(len(original)) * 100

	return map[string]interface{}{
		"originalChars":  len(original),
		"cleanedChars":   len(cleaned),
		"originalLines":  origLines,
		"cleanedLines":   cleanLines,
		"reductionPct":   reduction,
		"savedChars":     len(original) - len(cleaned),
	}
}
