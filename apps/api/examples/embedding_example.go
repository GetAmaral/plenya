package main

import (
	"context"
	"fmt"
	"log"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

// Exemplo de uso dos serviços de Embedding e Chunking
// Este arquivo demonstra como usar os serviços criados na Fase 2
//
// Para executar:
// go run examples/embedding_example.go
//
// IMPORTANTE: Configure OPENAI_API_KEY no .env antes de executar

func main() {
	fmt.Println("🚀 Exemplo de Embedding e Chunking - Plenya RAG System")
	fmt.Println("=" + string(make([]byte, 60)))

	// 1. Carregar configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Erro ao carregar config: %v", err)
	}

	// Verificar se OPENAI_API_KEY está configurada
	hasValidAPIKey := cfg.OpenAI.APIKey != "" && cfg.OpenAI.APIKey != "your_openai_api_key_here"
	if hasValidAPIKey {
		fmt.Printf("✅ Config carregada - Model: %s\n", cfg.OpenAI.EmbeddingModel)
		fmt.Println("✅ OPENAI_API_KEY configurada\n")
	} else {
		fmt.Println("⚠️  OPENAI_API_KEY não configurada (executando apenas exemplos de chunking)\n")
	}

	// 2. Criar serviço de chunking
	chunkingService := services.NewChunkingService()
	fmt.Println("✅ ChunkingService criado")

	// 3. Exemplo de chunking de artigo
	exampleArticle()

	// 4. Exemplo de chunking de ScoreItem
	exampleScoreItem(chunkingService)

	// 5. Exemplo de embedding (APENAS SE API KEY VÁLIDA)
	fmt.Println("\n⚠️  Pulando exemplo de embedding (requer API key válida)")
	fmt.Println("   Para testar embeddings, configure OPENAI_API_KEY e descomente exampleEmbedding()")

	// Descomentar para testar embedding real (consome tokens):
	// exampleEmbedding(cfg, nil)
}

// exampleArticle demonstra chunking de artigo científico
func exampleArticle() {
	fmt.Println("\n📄 Exemplo 1: Chunking de Artigo Científico")
	fmt.Println("-" + string(make([]byte, 60)))

	chunkingService := services.NewChunkingService()

	abstract := `Background: Hypertension is a major risk factor for cardiovascular disease.
	Objective: To evaluate the efficacy of lifestyle interventions in blood pressure control.`

	fullContent := `Introduction: Cardiovascular disease remains the leading cause of death worldwide.

	Methods: We conducted a randomized controlled trial with 500 participants.
	Participants were randomized to intervention or control groups.

	Results: The intervention group showed significant reduction in systolic BP.
	Mean reduction was 12.5 mmHg (95% CI: 10.2-14.8, p<0.001).

	Discussion: Our findings suggest that lifestyle interventions are effective.
	These results align with previous studies in the field.

	Conclusion: Lifestyle modifications should be recommended as first-line therapy.`

	chunks, err := chunkingService.ChunkArticle(fullContent, abstract)
	if err != nil {
		log.Printf("❌ Erro ao fazer chunking: %v", err)
		return
	}

	fmt.Printf("✅ Artigo dividido em %d chunks\n\n", len(chunks))

	for i, chunk := range chunks {
		section := chunk.Metadata["section"]
		wordCount := chunk.Metadata["word_count"]
		preview := chunk.Text
		if len(preview) > 100 {
			preview = preview[:100] + "..."
		}

		fmt.Printf("Chunk %d:\n", i)
		fmt.Printf("  Seção: %v\n", section)
		fmt.Printf("  Palavras: %v\n", wordCount)
		fmt.Printf("  Preview: %s\n\n", preview)
	}
}

// exampleScoreItem demonstra geração de texto para ScoreItem
func exampleScoreItem(chunkingService *services.ChunkingService) {
	fmt.Println("\n🎯 Exemplo 2: Chunking de ScoreItem")
	fmt.Println("-" + string(make([]byte, 60)))

	clinicalRelevance := "Valores baixos de hemoglobina indicam anemia, associada a fadiga e risco cardiovascular aumentado."
	patientExplanation := "Hemoglobina é a proteína que transporta oxigênio no sangue. Valores baixos podem causar cansaço."
	conduct := "Investigar causa (deficiência de ferro, B12, folato). Suplementação conforme indicado. Encaminhar ao hematologista se Hb < 10 g/dL."

	text := chunkingService.ChunkScoreItem(
		"Hemoglobina - Homens",
		&clinicalRelevance,
		&patientExplanation,
		&conduct,
	)

	fmt.Println("✅ Texto combinado para embedding:\n")
	fmt.Println(text)
	fmt.Printf("\n📊 Tamanho: %d caracteres (~%d tokens)\n", len(text), len(text)/4)
}

// exampleEmbedding demonstra geração de embedding real
// ATENÇÃO: Consome tokens da API OpenAI (~$0.0001 por execução)
func exampleEmbedding(cfg *config.Config, db *gorm.DB) {
	fmt.Println("\n🤖 Exemplo 3: Geração de Embedding Real")
	fmt.Println("-" + string(make([]byte, 60)))

	embeddingService := services.NewEmbeddingService(cfg, db)

	// Texto de exemplo
	text := "Diabetes mellitus tipo 2 é uma doença metabólica caracterizada por hiperglicemia crônica."

	// Gerar embedding
	ctx := context.Background()
	embedding, err := embeddingService.GenerateEmbedding(ctx, text)
	if err != nil {
		log.Printf("❌ Erro ao gerar embedding: %v", err)
		return
	}

	fmt.Printf("✅ Embedding gerado com sucesso!\n")
	fmt.Printf("   Dimensões: %d\n", len(embedding))
	fmt.Printf("   Primeiros 10 valores: %v\n", embedding[:10])
	fmt.Printf("   Custo estimado: $%.6f\n", embeddingService.GetEstimatedCost(text))
}
