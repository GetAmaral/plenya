package services

import (
	"strings"
	"testing"
)

// TestNewChunkingService testa criação do serviço
func TestNewChunkingService(t *testing.T) {
	service := NewChunkingService()

	if service.chunkSize != 1000 {
		t.Errorf("Expected chunkSize 1000, got %d", service.chunkSize)
	}

	if service.chunkOverlap != 200 {
		t.Errorf("Expected chunkOverlap 200, got %d", service.chunkOverlap)
	}
}

// TestChunkArticle_OnlyAbstract testa chunking de artigo com apenas abstract
func TestChunkArticle_OnlyAbstract(t *testing.T) {
	service := NewChunkingService()

	abstract := "Este é um abstract de teste sobre diabetes mellitus tipo 2."
	chunks, err := service.ChunkArticle("", abstract)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(chunks) != 1 {
		t.Errorf("Expected 1 chunk, got %d", len(chunks))
	}

	if chunks[0].Index != 0 {
		t.Errorf("Expected index 0, got %d", chunks[0].Index)
	}

	if chunks[0].Metadata["section"] != "abstract" {
		t.Errorf("Expected section 'abstract', got '%v'", chunks[0].Metadata["section"])
	}
}

// TestChunkArticle_OnlyFullContent testa chunking com apenas full content
func TestChunkArticle_OnlyFullContent(t *testing.T) {
	service := NewChunkingService()

	// Texto curto (menor que chunkSize)
	content := "Introduction: This is a short article about hypertension treatment."
	chunks, err := service.ChunkArticle(content, "")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(chunks) < 1 {
		t.Error("Expected at least 1 chunk")
	}

	// Verificar inferência de seção
	if chunks[0].Metadata["section"] != "introduction" {
		t.Errorf("Expected section 'introduction', got '%v'", chunks[0].Metadata["section"])
	}
}

// TestChunkArticle_AbstractAndContent testa chunking completo
func TestChunkArticle_AbstractAndContent(t *testing.T) {
	service := NewChunkingService()

	abstract := "Abstract about cardiovascular disease."
	content := "Introduction: Cardiovascular disease is a major health concern. Methods: We conducted a study. Results: The results showed improvement."

	chunks, err := service.ChunkArticle(content, abstract)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Deve ter pelo menos 2 chunks (abstract + content)
	if len(chunks) < 2 {
		t.Errorf("Expected at least 2 chunks, got %d", len(chunks))
	}

	// Primeiro chunk deve ser abstract
	if chunks[0].Metadata["section"] != "abstract" {
		t.Errorf("First chunk should be abstract, got '%v'", chunks[0].Metadata["section"])
	}

	// Chunks seguintes devem ter índices sequenciais
	for i := 1; i < len(chunks); i++ {
		if chunks[i].Index != i {
			t.Errorf("Chunk %d has wrong index: %d", i, chunks[i].Index)
		}
	}
}

// TestChunkArticle_LongContent testa chunking de conteúdo longo
func TestChunkArticle_LongContent(t *testing.T) {
	service := NewChunkingService()

	// Gerar conteúdo longo (> chunkSize)
	content := strings.Repeat("This is a sentence. ", 200) // ~4000 chars

	chunks, err := service.ChunkArticle(content, "")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Deve gerar múltiplos chunks
	if len(chunks) < 2 {
		t.Errorf("Expected multiple chunks for long content, got %d", len(chunks))
	}

	// Verificar que chunks têm tamanho razoável
	for i, chunk := range chunks {
		chunkLen := len(chunk.Text)
		if chunkLen > 1500 { // chunkSize + margem
			t.Errorf("Chunk %d too large: %d chars", i, chunkLen)
		}
		if chunkLen < 10 {
			t.Errorf("Chunk %d too small: %d chars", i, chunkLen)
		}
	}
}

// TestChunkArticle_EmptyInput testa erro com input vazio
func TestChunkArticle_EmptyInput(t *testing.T) {
	service := NewChunkingService()

	_, err := service.ChunkArticle("", "")

	if err == nil {
		t.Error("Expected error for empty input")
	}

	if !strings.Contains(err.Error(), "no chunks generated") {
		t.Errorf("Expected 'no chunks generated' error, got: %v", err)
	}
}

// TestInferSection testa inferência de seção
func TestInferSection(t *testing.T) {
	service := NewChunkingService()

	tests := []struct {
		text     string
		expected string
	}{
		{"Introdução: Este estudo avalia...", "introduction"},
		{"Introduction: This study evaluates...", "introduction"},
		{"Métodos: Foram coletadas amostras...", "methods"},
		{"Methods and Materials: We collected...", "methods"},
		{"Resultados: Os pacientes apresentaram...", "results"},
		{"Results: The findings showed...", "results"},
		{"Discussão: Nossos resultados sugerem...", "discussion"},
		{"Discussion: Our findings suggest...", "discussion"},
		{"Conclusão: Este estudo demonstrou...", "conclusion"},
		{"Conclusion: This study demonstrated...", "conclusion"},
		{"Qualquer outro texto sem keywords.", "other"},
	}

	for _, test := range tests {
		result := service.inferSection(test.text)
		if result != test.expected {
			t.Errorf("For text '%s', expected section '%s', got '%s'",
				test.text, test.expected, result)
		}
	}
}

// TestChunkScoreItem testa geração de texto para ScoreItem
func TestChunkScoreItem(t *testing.T) {
	service := NewChunkingService()

	clinicalRelevance := "Valores baixos indicam anemia."
	patientExplanation := "Hemoglobina transporta oxigênio no sangue."
	conduct := "Investigar causa e suplementar se necessário."

	result := service.ChunkScoreItem(
		"Hemoglobina",
		&clinicalRelevance,
		&patientExplanation,
		&conduct,
	)

	// Verificar que contém todos os componentes
	if !strings.Contains(result, "Hemoglobina") {
		t.Error("Result should contain name")
	}
	if !strings.Contains(result, "anemia") {
		t.Error("Result should contain clinical relevance")
	}
	if !strings.Contains(result, "oxigênio") {
		t.Error("Result should contain patient explanation")
	}
	if !strings.Contains(result, "suplementar") {
		t.Error("Result should contain conduct")
	}
}

// TestChunkScoreItem_OnlyName testa com apenas nome
func TestChunkScoreItem_OnlyName(t *testing.T) {
	service := NewChunkingService()

	result := service.ChunkScoreItem("Glicose", nil, nil, nil)

	if !strings.Contains(result, "Glicose") {
		t.Error("Result should contain name")
	}

	// Não deve conter labels dos campos opcionais
	if strings.Contains(result, "Relevância clínica:") {
		t.Error("Result should not contain empty clinical relevance label")
	}
}

// TestCountWords testa contagem de palavras
func TestCountWords(t *testing.T) {
	tests := []struct {
		text     string
		expected int
	}{
		{"", 0},
		{"   ", 0},
		{"Hello", 1},
		{"Hello world", 2},
		{"  Hello   world  ", 2},
		{"This is a test sentence.", 5},
	}

	for _, test := range tests {
		result := countWords(test.text)
		if result != test.expected {
			t.Errorf("For text '%s', expected %d words, got %d",
				test.text, test.expected, result)
		}
	}
}

// TestEstimateChunkCount testa estimativa de chunks
func TestEstimateChunkCount(t *testing.T) {
	service := NewChunkingService()

	// Apenas abstract
	count := service.EstimateChunkCount("", "Short abstract")
	if count != 1 {
		t.Errorf("Expected 1 chunk for abstract only, got %d", count)
	}

	// Abstract + conteúdo curto
	count = service.EstimateChunkCount("Short content", "Abstract")
	if count < 2 {
		t.Errorf("Expected at least 2 chunks, got %d", count)
	}

	// Conteúdo longo (4000 chars)
	longContent := strings.Repeat("x", 4000)
	count = service.EstimateChunkCount(longContent, "Abstract")
	if count < 3 {
		t.Errorf("Expected at least 3 chunks for long content, got %d", count)
	}
}

// TestValidateChunk testa validação de chunks
func TestValidateChunk(t *testing.T) {
	service := NewChunkingService()

	// Chunk válido
	validChunk := Chunk{
		Index: 0,
		Text:  "This is a valid chunk with enough content for embedding.",
	}
	if err := service.ValidateChunk(validChunk); err != nil {
		t.Errorf("Valid chunk should not produce error: %v", err)
	}

	// Chunk muito pequeno
	tooSmall := Chunk{
		Index: 0,
		Text:  "Short",
	}
	if err := service.ValidateChunk(tooSmall); err == nil {
		t.Error("Too small chunk should produce error")
	}

	// Chunk muito grande
	tooLarge := Chunk{
		Index: 0,
		Text:  strings.Repeat("x", 40000),
	}
	if err := service.ValidateChunk(tooLarge); err == nil {
		t.Error("Too large chunk should produce error")
	}
}

// TestFindLastSentenceBoundary testa busca de fronteira de sentença
func TestFindLastSentenceBoundary(t *testing.T) {
	service := NewChunkingService()

	tests := []struct {
		text     string
		expected int
	}{
		{"No sentence here", 0},
		{"One sentence. And another", 14}, // Após ". "
		{"Multiple. Sentences! Here? Yes", 27}, // Após "? "
		{"End with period.", 0},               // Sem espaço após ponto
	}

	for _, test := range tests {
		result := service.findLastSentenceBoundary(test.text)
		if result != test.expected {
			t.Errorf("For text '%s', expected position %d, got %d",
				test.text, test.expected, result)
		}
	}
}
