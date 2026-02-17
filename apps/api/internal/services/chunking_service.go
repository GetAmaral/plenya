package services

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Chunk representa um pedaço de texto para embedding
type Chunk struct {
	Index    int                    `json:"index"`    // Índice do chunk (0 = abstract, 1+ = content)
	Text     string                 `json:"text"`     // Texto do chunk
	Metadata map[string]interface{} `json:"metadata"` // Metadata estruturado (section, page_range, word_count)
}

// ChunkingService - serviço de chunking de texto para embeddings
// Implementa estratégia adaptativa para artigos científicos médicos
type ChunkingService struct {
	chunkSize    int // Tamanho alvo do chunk em caracteres (~1000 chars = 200-250 tokens)
	chunkOverlap int // Overlap entre chunks (20% do chunkSize)
}

// NewChunkingService cria nova instância do serviço de chunking
func NewChunkingService() *ChunkingService {
	chunkSize := 1000   // ~200-250 tokens (bom balanço para artigos médicos)
	chunkOverlap := 200 // 20% overlap para preservar contexto

	return &ChunkingService{
		chunkSize:    chunkSize,
		chunkOverlap: chunkOverlap,
	}
}

// sanitizeUTF8 remove caracteres inválidos UTF-8 do texto
// Importante para PDFs que podem conter bytes inválidos
func sanitizeUTF8(s string) string {
	if utf8.ValidString(s) {
		return s
	}

	// Remove caracteres inválidos
	var cleaned strings.Builder
	cleaned.Grow(len(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			// Caractere inválido - pula
			i++
			continue
		}
		cleaned.WriteRune(r)
		i += size
	}

	return cleaned.String()
}

// ChunkArticle divide artigo em chunks semânticos
// Estratégia:
//   - Chunk 0: Abstract (sempre separado)
//   - Chunks 1+: Full content com sliding window
func (s *ChunkingService) ChunkArticle(fullContent, abstract string) ([]Chunk, error) {
	// Sanitizar inputs para remover caracteres UTF-8 inválidos
	fullContent = sanitizeUTF8(fullContent)
	abstract = sanitizeUTF8(abstract)

	chunks := []Chunk{}

	// Chunk 0: Abstract (se fornecido)
	if abstract != "" && len(strings.TrimSpace(abstract)) > 0 {
		chunks = append(chunks, Chunk{
			Index: 0,
			Text:  strings.TrimSpace(abstract),
			Metadata: map[string]interface{}{
				"section":    "abstract",
				"word_count": countWords(abstract),
			},
		})
	}

	// Chunks 1+: Full content com sliding window
	if fullContent != "" && len(strings.TrimSpace(fullContent)) > 0 {
		contentChunks := s.chunkTextWithSlidingWindow(fullContent)

		// Inferir seção baseado em keywords
		for i, chunk := range contentChunks {
			section := s.inferSection(chunk.Text)
			chunk.Metadata["section"] = section
			chunk.Metadata["word_count"] = countWords(chunk.Text)

			// Ajustar índice (0 já usado pelo abstract)
			chunk.Index = len(chunks) + i

			chunks = append(chunks, chunk)
		}
	}

	// Validar que temos pelo menos 1 chunk
	if len(chunks) == 0 {
		return nil, fmt.Errorf("no chunks generated (both abstract and fullContent are empty)")
	}

	return chunks, nil
}

// chunkTextWithSlidingWindow divide texto usando sliding window com overlap
// Tenta quebrar em fronteiras de sentença para preservar contexto
func (s *ChunkingService) chunkTextWithSlidingWindow(text string) []Chunk {
	chunks := []Chunk{}
	textLen := len(text)

	// Se texto menor que chunkSize, retorna tudo como 1 chunk
	if textLen <= s.chunkSize {
		return []Chunk{
			{
				Index: 0,
				Text:  strings.TrimSpace(text),
				Metadata: map[string]interface{}{
					"word_count": countWords(text),
				},
			},
		}
	}

	// Sliding window com overlap
	start := 0
	chunkIndex := 0

	for start < textLen {
		// Calcular fim do chunk
		end := start + s.chunkSize
		if end > textLen {
			end = textLen
		}

		// Extrair chunk candidato
		chunkText := text[start:end]

		// Se não é o último chunk, tentar quebrar em fronteira de sentença
		if end < textLen {
			// Procurar última sentença completa (. ! ?)
			lastSentenceEnd := s.findLastSentenceBoundary(chunkText)
			if lastSentenceEnd > 0 {
				// Ajustar fim para quebrar na sentença
				end = start + lastSentenceEnd
				chunkText = text[start:end]
			}
		}

		// Adicionar chunk
		chunks = append(chunks, Chunk{
			Index: chunkIndex,
			Text:  strings.TrimSpace(chunkText),
			Metadata: map[string]interface{}{
				"word_count": countWords(chunkText),
			},
		})

		// Avançar posição com overlap
		// Se é o último chunk, avançar até o fim
		if end >= textLen {
			break
		}

		// Caso contrário, avançar (chunkSize - overlap)
		start += (s.chunkSize - s.chunkOverlap)
		chunkIndex++
	}

	return chunks
}

// findLastSentenceBoundary encontra posição da última fronteira de sentença
// Procura por ". " "! " "? " (sentença seguida de espaço)
// Retorna posição após o ponto (incluindo espaço), ou 0 se não encontrar
func (s *ChunkingService) findLastSentenceBoundary(text string) int {
	// Regex: ponto/exclamação/interrogação seguido de espaço
	sentencePattern := regexp.MustCompile(`[.!?]\s+`)

	// Encontrar todas as ocorrências
	matches := sentencePattern.FindAllStringIndex(text, -1)

	if len(matches) == 0 {
		return 0
	}

	// Retornar posição após o último match
	lastMatch := matches[len(matches)-1]
	return lastMatch[1] // Posição após o match (inclui espaço)
}

// inferSection infere seção do artigo baseado em keywords
// Retorna: "abstract", "methods", "results", "discussion", "introduction", "conclusion", "other"
func (s *ChunkingService) inferSection(text string) string {
	textLower := strings.ToLower(text)

	// Keywords por seção (português + inglês)
	// Ordem de prioridade: verificar em sequência
	sectionsInOrder := []struct {
		name     string
		keywords []string
	}{
		{"introduction", []string{"introdução", "introduction", "background", "contexto"}},
		{"methods", []string{"métodos", "methods", "metodologia", "methodology", "materiais e métodos", "materials and methods"}},
		{"results", []string{"resultados", "results", "achados", "findings"}},
		{"discussion", []string{"discussão", "discussion", "análise", "analysis"}},
		{"conclusion", []string{"conclusão", "conclusion", "considerações finais", "final considerations"}},
	}

	// Procurar keywords priorizando aquelas que aparecem mais cedo no texto
	bestMatch := "other"
	bestPosition := len(textLower) // Posição máxima

	for _, section := range sectionsInOrder {
		for _, keyword := range section.keywords {
			pos := strings.Index(textLower, keyword)
			if pos >= 0 && pos < bestPosition {
				bestMatch = section.name
				bestPosition = pos
			}
		}
	}

	return bestMatch
}

// ChunkScoreItem gera texto combinado de ScoreItem para embedding
// Combina: fullName (com contexto hierárquico) + clinical_relevance + patient_explanation + conduct
// O fullName deve incluir Group/Subgroup/Parent para melhor contexto semântico
func (s *ChunkingService) ChunkScoreItem(
	fullName string,
	clinicalRelevance *string,
	patientExplanation *string,
	conduct *string,
) string {
	parts := []string{}

	// Nome completo com contexto hierárquico (obrigatório)
	if fullName != "" {
		parts = append(parts, fmt.Sprintf("Parâmetro: %s", fullName))
	}

	// Relevância clínica
	if clinicalRelevance != nil && *clinicalRelevance != "" {
		parts = append(parts, fmt.Sprintf("Relevância clínica: %s", *clinicalRelevance))
	}

	// Explicação para paciente
	if patientExplanation != nil && *patientExplanation != "" {
		parts = append(parts, fmt.Sprintf("Explicação: %s", *patientExplanation))
	}

	// Conduta
	if conduct != nil && *conduct != "" {
		parts = append(parts, fmt.Sprintf("Conduta: %s", *conduct))
	}

	// Combinar com quebra de linha
	combined := strings.Join(parts, "\n\n")

	return combined
}

// countWords conta palavras em um texto
func countWords(text string) int {
	// Remove espaços extras e quebra de linha
	text = strings.TrimSpace(text)
	if text == "" {
		return 0
	}

	// Split por espaços e conta
	words := strings.Fields(text)
	return len(words)
}

// EstimateChunkCount estima quantos chunks serão gerados
func (s *ChunkingService) EstimateChunkCount(fullContent, abstract string) int {
	count := 0

	// Abstract = 1 chunk
	if abstract != "" && len(strings.TrimSpace(abstract)) > 0 {
		count++
	}

	// Estimar chunks do fullContent
	if fullContent != "" {
		contentLen := len(fullContent)
		// Fórmula: ceil(contentLen / (chunkSize - overlap))
		effectiveChunkSize := s.chunkSize - s.chunkOverlap
		contentChunks := (contentLen + effectiveChunkSize - 1) / effectiveChunkSize
		count += contentChunks
	}

	return count
}

// ValidateChunk verifica se chunk é válido para embedding
func (s *ChunkingService) ValidateChunk(chunk Chunk) error {
	// Chunk muito pequeno (< 10 caracteres)
	if len(strings.TrimSpace(chunk.Text)) < 10 {
		return fmt.Errorf("chunk text too short: %d characters", len(chunk.Text))
	}

	// Chunk muito grande (> 32k caracteres = limite OpenAI)
	maxSize := 32000
	if len(chunk.Text) > maxSize {
		return fmt.Errorf("chunk text too long: %d characters (max %d)", len(chunk.Text), maxSize)
	}

	return nil
}
