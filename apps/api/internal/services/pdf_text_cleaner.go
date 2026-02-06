package services

import (
	"regexp"
	"strings"
)

// PDFTextCleaner - serviço para limpar/processar texto extraído de PDFs de laudos
// Contém padrões específicos para laboratórios conhecidos + padrões genéricos para outros
type PDFTextCleaner struct {
	// === CABEÇALHOS DE PÁGINA - ESPECÍFICOS ===
	sabinHeaderPattern   *regexp.Regexp // Sabin/LDB
	labmedHeaderPattern  *regexp.Regexp // LabMed/LabImagem
	oswaldoHeaderPattern *regexp.Regexp // Oswaldo Cruz

	// === CABEÇALHOS DE PÁGINA - GENÉRICOS ===
	genericPatientLinePattern *regexp.Regexp // Linhas com Nome/Paciente + dados
	genericHeaderDotsPattern  *regexp.Regexp // Campos com pontos (Nome.....: VALOR)
	genericPageHeaderPattern  *regexp.Regexp // Linhas com RG, CPF, DN, Convênio

	// === ASSINATURAS DIGITAIS - GENÉRICAS ===
	digitalSignaturePattern *regexp.Regexp // Qualquer "ASSINATURA DIGITAL" + hash
	electronicSignPattern   *regexp.Regexp // "assinatura eletrônica" + hash
	hexHashPattern          *regexp.Regexp // Hashes hexadecimais isolados (40+ chars)

	// === LIBERAÇÕES - GENÉRICAS ===
	releasePattern          *regexp.Regexp // "LIBERADO ELETRONICAMENTE" (qualquer formato)
	responsavelTecPattern   *regexp.Regexp // "Responsável Técnico" (qualquer formato)
	professionalCredsPattern *regexp.Regexp // Credenciais CRM/CRF/CRBM isoladas

	// === METADADOS ADMINISTRATIVOS - GENÉRICOS ===
	cnesPattern            *regexp.Regexp // CNES (qualquer formato)
	labRegistrationPattern *regexp.Regexp // Registro de laboratório
	labToLabPattern        *regexp.Regexp // Referências inter-laboratoriais
	processingNotePattern  *regexp.Regexp // Exame processado externamente
	addressPattern         *regexp.Regexp // Endereços (Rua, Av, Avenida)

	// === DISCLAIMERS E NOTAS LEGAIS - GENÉRICOS ===
	disclaimerPattern        *regexp.Regexp // Disclaimers de autenticidade
	accreditationPattern     *regexp.Regexp // Acreditações (CAP, PALC, etc)
	methodologyChangePattern *regexp.Regexp // Avisos de mudança de metodologia
	referenceChangePattern   *regexp.Regexp // Avisos de mudança de valores de referência

	// === RESULTADOS ANTERIORES E GRÁFICOS ===
	previousResultsPattern *regexp.Regexp // "Resultados Anteriores" + dados
	asciiGraphPattern      *regexp.Regexp // Gráficos ASCII
	graphNotesPattern      *regexp.Regexp // Notas sobre gráficos
	histogramLabelPattern  *regexp.Regexp // Labels de histogramas

	// === TIMESTAMPS - GENÉRICOS ===
	coletaLiberacaoPattern *regexp.Regexp // Coleta/Liberação (vários formatos)
	timestampLinePattern   *regexp.Regexp // Linhas com apenas data/hora

	// === PAGINAÇÃO - GENÉRICOS ===
	pageNumberPattern *regexp.Regexp // Pág/Página/Folha + número

	// === NOTAS EXPLICATIVAS ===
	// DISABLED: notaBlockPattern and referenciasPattern were removing 95%+ of document content
	// because the patterns matched from NOTAS:/REFERÊNCIAS: to end of file when no blank line followed
	// These sections contain clinically relevant information for AI interpretation
	notaInlinePattern *regexp.Regexp // Notas inline numeradas (kept - only removes single lines)

	// === FORMATAÇÃO E LIMPEZA ===
	dotsPattern           *regexp.Regexp // Sequências de pontos
	separatorLinePattern  *regexp.Regexp // Linhas separadoras (_, -, =)
	urlPattern            *regexp.Regexp // URLs
	materialMethodPattern *regexp.Regexp // Material/Método/Tecnologia
}

// NewPDFTextCleaner cria uma nova instância do cleaner com todos os padrões compilados
func NewPDFTextCleaner() *PDFTextCleaner {
	return &PDFTextCleaner{
		// =============================================================
		// CABEÇALHOS DE PÁGINA - ESPECÍFICOS (laboratórios conhecidos)
		// =============================================================

		// Sabin/LDB: Nome, RG, DN, CPF, Médico, Convênio, Unidade, Página
		sabinHeaderPattern: regexp.MustCompile(`(?m)^Nome\s*:\s*.+$\n^RG\s*:.+Código da OS.+$\n^DN\s*:.+CPF\s*:.+$\n^Médico\s*:.+Atendimento.+$\n^Convênio:.+Qnt de exames.+$\n^Unidade\s*:.+Página:.+$`),

		// LabMed: Sr(a), Dr(a), Local Coleta, Convênio com Pág
		labmedHeaderPattern: regexp.MustCompile(`(?m)^Sr\s*\(a\)\.+:\s*.+Idade\.+:.+$\n^Dr\s*\(a\)\.+:.+Entrada\.+:.+$\n^Local Coleta:.+Emissão\.+:.+$\n^Convênio\.+:.+Requisição\.+:.+$`),

		// Oswaldo Cruz: Nome, Requisição, Convênio, Médico, Destino, CPF, RG
		oswaldoHeaderPattern: regexp.MustCompile(`(?m)^Nome\.+:\s*.+Entrada\.+:.+$\n^Requisição\.+:.+Guia\.+:.+$\n^Convênio\.+:.+Sexo\.+:.+$\n^Médico\.+:.+Idade\.+:.+$\n^Destino\.+:.+Impresso\.+:.+$\n^CPF\.+:.+RG\.+:.+$`),

		// =============================================================
		// CABEÇALHOS DE PÁGINA - GENÉRICOS (qualquer laboratório)
		// =============================================================

		// Linhas com "Nome" ou "Paciente" seguido de dados pessoais
		genericPatientLinePattern: regexp.MustCompile(`(?mi)^(?:Nome|Paciente|Sr\.?\s*\(a\))\s*[:.].*(?:RG|CPF|DN|Nasc|Data\s+de\s+Nasc).*$`),

		// Campos com sequência de pontos para alinhamento (Campo......: Valor)
		genericHeaderDotsPattern: regexp.MustCompile(`(?m)^(?:Nome|Paciente|RG|CPF|DN|Sexo|Idade|Médico|Convênio|Requisição|Guia|Entrada|Destino|Unidade|Local)\.{2,}:\s*.+$`),

		// Linhas de cabeçalho com múltiplos campos de identificação
		genericPageHeaderPattern: regexp.MustCompile(`(?mi)^.*(?:(?:RG|CPF|DN|Convênio|Médico|Requisição)\s*[:.].*){2,}$`),

		// =============================================================
		// ASSINATURAS DIGITAIS - GENÉRICAS
		// =============================================================

		// "ASSINATURA DIGITAL" ou "Assinatura Digital" (qualquer capitalização)
		digitalSignaturePattern: regexp.MustCompile(`(?mi)^\s*ASSINATURA\s+DIGITAL\s*$`),

		// Assinatura eletrônica com hash
		electronicSignPattern: regexp.MustCompile(`(?mi)^.*(?:assinatura\s+eletr[oô]nica|digitalmente\s+assinado).*[:\s]+[A-Fa-f0-9]{32,}\s*$`),

		// Hashes hexadecimais isolados (linhas com apenas hash de 40+ caracteres)
		hexHashPattern: regexp.MustCompile(`(?m)^\s*[A-Fa-f0-9]{40,}\s*$`),

		// =============================================================
		// LIBERAÇÕES E RESPONSÁVEIS - GENÉRICAS
		// =============================================================

		// "LIBERADO ELETRONICAMENTE" em qualquer formato
		releasePattern: regexp.MustCompile(`(?mi)^.*(?:liberado|aprovado)\s+(?:eletronicamente|digitalmente).*$`),

		// "Responsável Técnico" em qualquer formato
		responsavelTecPattern: regexp.MustCompile(`(?mi)^.*Respons[aá]vel\s+T[eé]cnico.*$`),

		// Credenciais profissionais isoladas ou em contexto de liberação
		// CRM, CRF, CRBM, CREMESP, etc + número
		professionalCredsPattern: regexp.MustCompile(`(?mi)^.*(?:CRM|CRF|CRBM|CREMESP|CRBio|COREN)[/-]?\s*(?:[A-Z]{2}\s*)?\d{3,}.*(?:liberado|assinado|aprovado).*$`),

		// =============================================================
		// METADADOS ADMINISTRATIVOS - GENÉRICOS
		// =============================================================

		// CNES em qualquer contexto
		cnesPattern: regexp.MustCompile(`(?mi)^.*CNES\s*[:.]?\s*\d{5,}.*$`),

		// Registro de laboratório (CRM, ANVISA, Vigilância Sanitária)
		labRegistrationPattern: regexp.MustCompile(`(?mi)^.*(?:Laborat[oó]rio\s+registrado|Registro\s+(?:CRM|ANVISA|Lic\.?\s*Sanit)).*$`),

		// Referências inter-laboratoriais
		labToLabPattern: regexp.MustCompile(`(?mi)^.*(?:Lab-to-Lab|Apoio\s+Laboratorial|Terceirizado).*CNES.*$`),

		// Exame processado externamente
		processingNotePattern: regexp.MustCompile(`(?mi)^.*(?:Exame\s+processado|Análise\s+realizada)\s+(?:por|em|na|pelo).*$`),

		// Endereços (genérico para qualquer formato brasileiro)
		addressPattern: regexp.MustCompile(`(?mi)^.*(?:Endere[çc]o|Rua|Av\.?|Avenida|Alameda|Pra[çc]a)\s+.+(?:CEP|n[ºo°]|\d{5}-?\d{3}).*$`),

		// =============================================================
		// DISCLAIMERS E NOTAS LEGAIS - GENÉRICOS
		// =============================================================

		// Disclaimers de autenticidade/aceitação
		disclaimerPattern: regexp.MustCompile(`(?mi)^.*(?:aceita[çc][aã]o\s+(?:deste|do)\s+resultado|verifica[çc][aã]o\s+de\s+(?:sua\s+)?autenticidade|resultado.*condicionad[ao]).*$`),

		// Acreditações (CAP, PALC, ONA, ISO, DICQ, etc)
		accreditationPattern: regexp.MustCompile(`(?mi)^.*(?:ACREDITAD[OA]|CERTIFICAD[OA]).*(?:CAP|PALC|ONA|ISO|DICQ|INMETRO|SBPC|SBAC).*$`),

		// Avisos de mudança de metodologia
		methodologyChangePattern: regexp.MustCompile(`(?mi)^[\*\-\s]*(?:Aten[çc][aã]o|Aviso|Nota)?\s*(?:para\s+)?(?:mudan[çc]a\s+de\s+(?:metodologia|tecnologia|equipamento)|nova\s+metodologia).*$`),

		// Avisos de mudança de valores de referência
		referenceChangePattern: regexp.MustCompile(`(?mi)^[\*\-\s]*(?:Aten[çc][aã]o|Aviso|Nota)?\s*(?:para\s+)?nov[oa]s?\s+valores?\s+de\s+refer[eê]ncia.*$`),

		// =============================================================
		// RESULTADOS ANTERIORES E GRÁFICOS
		// =============================================================

		// "Resultados Anteriores" seguido de datas/valores (multilinhas)
		// Note: Go regex doesn't support lookahead, so we match until double newline or end
		previousResultsPattern: regexp.MustCompile(`(?msi)^Resultados?\s+Anteriore?s?\s*:.*?(?:\n\s*\n|$)`),

		// Gráficos ASCII (linhas com |, barras, eixos)
		// Note: │ is Unicode U+2502 (BOX DRAWINGS LIGHT VERTICAL)
		asciiGraphPattern: regexp.MustCompile(`(?m)(?:^\s*[\d.,]*\s*[|│].*$\n){2,}`),

		// Notas sobre gráficos não exibidos
		graphNotesPattern: regexp.MustCompile(`(?mi)^.*(?:n[aã]o\s+(?:apresentado|exibido|mostrado)\s+no\s+gr[aá]fico|primeiro\s+exame\s+do\s+paciente).*$`),

		// Labels de histogramas (equipamentos hematológicos)
		histogramLabelPattern: regexp.MustCompile(`(?mi)^Histogram[a]?\s+(?:DIFF|RBC|PLT|WBC|WNR|WDF|WPC|RET|IMI)\s*$`),

		// =============================================================
		// TIMESTAMPS - GENÉRICOS
		// =============================================================

		// Coleta/Liberação/Emissão em vários formatos
		coletaLiberacaoPattern: regexp.MustCompile(`(?mi)^(?:Coleta|Libera[çc][aã]o|Emiss[aã]o|Data\s+(?:da\s+)?Coleta|Data\s+(?:de\s+)?Libera[çc][aã]o)\s*[:.]?\s*\d{2}[/-]\d{2}[/-]\d{2,4}(?:\s*[-\s]\s*\d{2}:\d{2}(?::\d{2})?)?.*$`),

		// Linhas com apenas timestamps (sem nome de exame)
		timestampLinePattern: regexp.MustCompile(`(?m)^\s*\d{2}[/-]\d{2}[/-]\d{2,4}\s+\d{2}:\d{2}(?::\d{2})?\s*$`),

		// =============================================================
		// PAGINAÇÃO - GENÉRICOS
		// =============================================================

		// Pág/Página/Folha + número (vários formatos)
		pageNumberPattern: regexp.MustCompile(`(?mi)(?:P[aá]g(?:ina)?|Folha|Fl)\s*\.?\s*:?\s*\d+\s*(?:/|de)\s*\d+`),

		// =============================================================
		// NOTAS EXPLICATIVAS
		// =============================================================

		// DISABLED: notaBlockPattern - was removing 95%+ of content
		// Pattern `(?msi)^NOTAS?\s*:?\s*\n(?:.*\n)*?(?:\n\s*\n|$)` with 's' flag
		// caused `.` to match newlines, consuming from NOTAS: to end of file

		// Notas inline numeradas (Nota 1:, Nota 2:, etc) - safe single-line pattern
		notaInlinePattern: regexp.MustCompile(`(?mi)^Nota\s*\d+\s*:\s*.+$`),

		// DISABLED: referenciasPattern - was removing 85%+ of content
		// Same issue as notaBlockPattern

		// =============================================================
		// FORMATAÇÃO E LIMPEZA
		// =============================================================

		// Sequências de pontos para alinhamento/espaçamento
		dotsPattern: regexp.MustCompile(`\.{3,}[:.\s]*`),

		// Linhas separadoras (underscores, hífens, iguais)
		separatorLinePattern: regexp.MustCompile(`(?m)^[\s]*[_\-=]{10,}[\s]*$`),

		// URLs (http, https, www)
		urlPattern: regexp.MustCompile(`(?i)https?://[^\s]+|www\.[^\s]+\.(?:com|org|net|br)[^\s]*`),

		// Material/Método/Tecnologia - simplifica mantendo essencial
		materialMethodPattern: regexp.MustCompile(`(?mi)^Material:\s*(.+?)\s+Coletado\s+em:\s*(\d{2}/\d{2}/\d{4})(?:\s+\d{2}:\d{2})?\s+M[eé]todo:\s*(.+?)(?:\s+Tecnologia:.+)?$`),
	}
}

// CleanText - processa e limpa o texto completo extraído do PDF
// Retorna versão reduzida focada apenas em dados essenciais para IA
// Aplica padrões específicos primeiro (mais precisos), depois genéricos (fallback)
func (c *PDFTextCleaner) CleanText(fullText string) string {
	text := fullText

	// =================================================================
	// FASE 1: Normalização de caracteres
	// =================================================================
	text = strings.ReplaceAll(text, "\x0C", "\n") // form feed
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")

	// =================================================================
	// FASE 2: Cabeçalhos de página (ALTA PRIORIDADE - maior volume de ruído)
	// Primeiro tenta padrões específicos, depois genéricos
	// =================================================================

	// Específicos (mais precisos)
	text = c.sabinHeaderPattern.ReplaceAllString(text, "")
	text = c.labmedHeaderPattern.ReplaceAllString(text, "")
	text = c.oswaldoHeaderPattern.ReplaceAllString(text, "")

	// Genéricos (fallback para outros laboratórios)
	text = c.genericPatientLinePattern.ReplaceAllString(text, "")
	text = c.genericHeaderDotsPattern.ReplaceAllString(text, "")
	text = c.genericPageHeaderPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 3: Assinaturas digitais e hashes (ALTA PRIORIDADE)
	// =================================================================
	text = c.digitalSignaturePattern.ReplaceAllString(text, "")
	text = c.electronicSignPattern.ReplaceAllString(text, "")
	text = c.hexHashPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 4: Liberações e responsáveis técnicos (ALTA PRIORIDADE)
	// =================================================================
	text = c.releasePattern.ReplaceAllString(text, "")
	text = c.responsavelTecPattern.ReplaceAllString(text, "")
	text = c.professionalCredsPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 5: Disclaimers e notas legais (ALTA PRIORIDADE)
	// =================================================================
	text = c.disclaimerPattern.ReplaceAllString(text, "")
	text = c.accreditationPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 6: Resultados anteriores e gráficos (ALTA PRIORIDADE)
	// =================================================================
	text = c.previousResultsPattern.ReplaceAllString(text, "")
	text = c.asciiGraphPattern.ReplaceAllString(text, "")
	text = c.graphNotesPattern.ReplaceAllString(text, "")
	text = c.histogramLabelPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 7: Metadados administrativos (MÉDIA PRIORIDADE)
	// =================================================================
	text = c.cnesPattern.ReplaceAllString(text, "")
	text = c.labRegistrationPattern.ReplaceAllString(text, "")
	text = c.labToLabPattern.ReplaceAllString(text, "")
	text = c.processingNotePattern.ReplaceAllString(text, "")
	text = c.addressPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 8: Avisos de mudanças (MÉDIA PRIORIDADE)
	// =================================================================
	text = c.methodologyChangePattern.ReplaceAllString(text, "")
	text = c.referenceChangePattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 9: Timestamps (MÉDIA PRIORIDADE)
	// =================================================================
	text = c.coletaLiberacaoPattern.ReplaceAllString(text, "")
	text = c.timestampLinePattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 10: Notas explicativas (MÉDIA PRIORIDADE)
	// =================================================================
	// DISABLED: notaBlockPattern and referenciasPattern (removed 95%+ of content)
	// These patterns had dangerous multiline matching that consumed entire documents
	text = c.notaInlinePattern.ReplaceAllString(text, "") // Safe: only removes single lines

	// =================================================================
	// FASE 11: Paginação, separadores, URLs (BAIXA PRIORIDADE)
	// =================================================================
	text = c.pageNumberPattern.ReplaceAllString(text, "")
	text = c.separatorLinePattern.ReplaceAllString(text, "")
	text = c.urlPattern.ReplaceAllString(text, "")

	// =================================================================
	// FASE 12: Simplificação de Material/Método
	// =================================================================
	text = c.materialMethodPattern.ReplaceAllString(text, "Material: $1, Coletado: $2, Método: $3")

	// =================================================================
	// FASE 13: Limpeza de formatação
	// =================================================================

	// Remove sequências de pontos usados para alinhamento
	text = c.dotsPattern.ReplaceAllString(text, ": ")

	// Remove múltiplos espaços/tabs
	text = regexp.MustCompile(`[ \t]+`).ReplaceAllString(text, " ")

	// Remove múltiplas linhas vazias (mantém máximo 1)
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")

	// =================================================================
	// FASE 14: Filtragem de linhas sem conteúdo útil
	// =================================================================
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

	// Remove linhas que são apenas símbolos/caracteres especiais
	text = regexp.MustCompile(`(?m)^\s*[+*\-|│┌┐└┘├┤┬┴┼]+\s*$`).ReplaceAllString(text, "")

	// =================================================================
	// FASE 15: Limpeza final
	// =================================================================
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
	origLen := len(original)
	cleanLen := len(cleaned)
	origLines := len(strings.Split(original, "\n"))
	cleanLines := len(strings.Split(cleaned, "\n"))

	var reduction float64
	if origLen > 0 {
		reduction = float64(origLen-cleanLen) / float64(origLen) * 100
	}

	return map[string]interface{}{
		"originalChars":  origLen,
		"cleanedChars":   cleanLen,
		"originalLines":  origLines,
		"cleanedLines":   cleanLines,
		"reductionPct":   reduction,
		"savedChars":     origLen - cleanLen,
	}
}
