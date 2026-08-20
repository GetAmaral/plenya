// Comando de enriquecimento do catálogo magistral pelo RAG do consultório.
//
// O QUE ELE FAZ: para cada substância do catálogo, busca por similaridade os trechos das aulas e
// artigos que falam dela, guarda os trechos como EVIDÊNCIA ANEXADA e pede ao Claude que EXTRAIA
// (não que invente) indicação e dose ditas ali.
//
// O QUE ELE NÃO FAZ, de propósito:
//   · não escreve dose que não esteja explícita no trecho — se a aula não dá número, o campo
//     numérico continua NULL e só a indicação é preenchida;
//   · não sobrescreve dose já curada à mão (curated/confirmed);
//   · não marca nada como confirmado: tudo entra como `suggested`, para o médico conferir contra
//     o trecho que fica guardado ao lado.
//
// Uso:
//   go run ./cmd/magistral-enrich --dry-run --limit 3
//   go run ./cmd/magistral-enrich --only "Melatonina"
//   go run ./cmd/magistral-enrich            (todas as pendentes)
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type chunkHit struct {
	EmbeddingID  uuid.UUID
	ArticleID    uuid.UUID
	ArticleTitle string
	Journal      string
	ArticleType  string
	ChunkIndex   int
	ChunkText    string
	Similarity   float64
}

// extraction é o que o modelo devolve. Campos numéricos são ponteiros: ausente ≠ zero.
type extraction struct {
	Indications  string   `json:"indications"`
	DoseText     string   `json:"doseText"`
	UsualDose    *float64 `json:"usualDose"`
	MinDose      *float64 `json:"minDose"`
	MaxDose      *float64 `json:"maxDose"`
	Unit         string   `json:"unit"`
	HasEvidence  bool     `json:"hasEvidence"`
	SourceTitles []string `json:"sourceTitles"`
}

func main() {
	var (
		dryRun    = flag.Bool("dry-run", false, "não grava nada; só mostra o que faria")
		limit     = flag.Int("limit", 0, "processa no máximo N componentes (0 = todos)")
		only      = flag.String("only", "", "processa só a substância com este nome")
		chunks    = flag.Int("chunks", 8, "trechos do RAG por substância")
		minSim    = flag.Float64("min-similarity", 0.35, "similaridade mínima do trecho")
		redo      = flag.Bool("redo", false, "reprocessa quem já foi enriquecido")
		modelName = flag.String("model", "claude-opus-5", "modelo do Claude para a extração")
	)
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if cfg.Claude.APIKey == "" {
		log.Fatal("CLAUDE_API_KEY não configurada")
	}
	if err := database.Connect(cfg); err != nil {
		log.Fatal(err)
	}
	db := database.DB.Session(&gorm.Session{Logger: database.DB.Logger.LogMode(3)}) // silent-ish

	embeddings := services.NewEmbeddingService(cfg, database.DB)
	ctx := context.Background()

	var components []models.MagistralComponent
	q := db.Where("is_active = true AND deleted_at IS NULL")
	if *only != "" {
		q = q.Where("lower(name) = lower(?)", *only)
	} else if !*redo {
		q = q.Where("evidence_status = 'pending'")
	}
	q = q.Order("usage_count DESC").Order("name")
	if *limit > 0 {
		q = q.Limit(*limit)
	}
	if err := q.Find(&components).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🔎 %d substância(s) a enriquecer · RAG: ", len(components))
	var totalChunks int64
	db.Table("article_embeddings").Count(&totalChunks)
	fmt.Printf("%d trechos indexados\n\n", totalChunks)

	var comEvidencia, semEvidencia, comDose int
	for i, c := range components {
		fmt.Printf("[%d/%d] %s\n", i+1, len(components), c.Name)

		hits, err := searchRAG(ctx, db, embeddings, c, *chunks, *minSim)
		if err != nil {
			fmt.Printf("   ⚠️  busca falhou: %v\n", err)
			continue
		}
		if len(hits) == 0 {
			fmt.Printf("   — sem trecho acima de %.2f de similaridade\n", *minSim)
			semEvidencia++
			if !*dryRun {
				db.Model(&models.MagistralComponent{}).Where("id = ?", c.ID).
					Updates(map[string]any{"evidence_status": "suggested", "enriched_at": time.Now()})
			}
			continue
		}

		aulas := 0
		for _, h := range hits {
			if h.ArticleType == "lecture" {
				aulas++
			}
		}
		fmt.Printf("   %d trecho(s) (%d de aula) · melhor %.3f · %s\n",
			len(hits), aulas, hits[0].Similarity, truncate(hits[0].ArticleTitle, 60))

		ext, err := extract(cfg.Claude.APIKey, *modelName, c, hits)
		if err != nil {
			fmt.Printf("   ⚠️  extração falhou: %v\n", err)
			continue
		}
		if !ext.HasEvidence {
			fmt.Printf("   — os trechos não falam da substância de forma útil\n")
			semEvidencia++
		} else {
			comEvidencia++
			fmt.Printf("   ✓ %s\n", truncate(strings.ReplaceAll(ext.Indications, "\n", " "), 110))
			if ext.DoseText != "" {
				fmt.Printf("     dose: %s\n", truncate(ext.DoseText, 100))
			}
			if ext.UsualDose != nil {
				comDose++
			}
		}

		if *dryRun {
			continue
		}
		if err := persist(db, &c, hits, ext); err != nil {
			fmt.Printf("   ⚠️  gravação falhou: %v\n", err)
		}
	}

	fmt.Printf("\n📊 %d com evidência · %d sem · %d ganharam dose numérica do RAG\n",
		comEvidencia, semEvidencia, comDose)
	if *dryRun {
		fmt.Println("(dry-run: nada gravado)")
	}
}

// searchRAG traz os melhores trechos para a substância. Prioriza AULA sobre artigo: o material da
// pós é o que reflete a prática que o médico aprendeu, e é onde a dose costuma estar escrita.
func searchRAG(ctx context.Context, db *gorm.DB, emb *services.EmbeddingService,
	c models.MagistralComponent, k int, minSim float64) ([]chunkHit, error) {

	terms := c.Name
	if c.Synonyms != "" {
		terms += ", " + c.Synonyms
	}
	query := fmt.Sprintf(
		"%s: indicação clínica, mecanismo de ação, dose e posologia em medicina funcional integrativa e nutrologia",
		terms)

	vec, err := emb.GenerateEmbedding(ctx, query)
	if err != nil {
		return nil, err
	}

	var hits []chunkHit
	err = db.Raw(`
		SELECT ae.id AS embedding_id,
		       ae.article_id,
		       a.title AS article_title,
		       a.journal,
		       a.article_type,
		       ae.chunk_index,
		       ae.chunk_text,
		       1 - (ae.embedding <=> ?) AS similarity
		FROM article_embeddings ae
		JOIN articles a ON a.id = ae.article_id
		WHERE a.deleted_at IS NULL
		  AND ae.embedding IS NOT NULL
		  AND 1 - (ae.embedding <=> ?) >= ?
		ORDER BY (1 - (ae.embedding <=> ?)) + CASE WHEN a.article_type = 'lecture' THEN 0.03 ELSE 0 END DESC
		LIMIT ?`,
		pgvector.NewVector(vec), pgvector.NewVector(vec), minSim, pgvector.NewVector(vec), k).
		Scan(&hits).Error
	return hits, err
}

const extractSystem = `Você extrai informação de material didático de medicina funcional integrativa e nutrologia para um catálogo de prescrição magistral.

REGRAS INVIOLÁVEIS:
1. Só afirme o que ESTÁ ESCRITO nos trechos fornecidos. Nunca complete com conhecimento próprio.
2. Dose numérica só quando o trecho traz o número para ESTA substância. Se o trecho fala da substância mas não dá dose, devolva os campos numéricos nulos e descreva a indicação mesmo assim.
3. Se os trechos não tratam da substância (só a citam de passagem ou falam de outra coisa), devolva hasEvidence=false e deixe o resto vazio.
4. Escreva em português do Brasil, prosa clínica direta, sem travessão, sem "não é X, é Y", sem fecho de efeito. Máximo 3 frases nas indicações.
5. A unidade da dose deve ser a mesma que o trecho usa (mg, mcg, g, UI).`

func extract(apiKey, model string, c models.MagistralComponent, hits []chunkHit) (*extraction, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "SUBSTÂNCIA: %s", c.Name)
	if c.Synonyms != "" {
		fmt.Fprintf(&b, " (sinônimos: %s)", c.Synonyms)
	}
	fmt.Fprintf(&b, "\nUnidade usada no catálogo: %s\n\nTRECHOS DO MATERIAL:\n", c.DefaultUnit)
	for i, h := range hits {
		fmt.Fprintf(&b, "\n[%d] %s (%s, similaridade %.3f)\n%s\n",
			i+1, h.ArticleTitle, h.ArticleType, h.Similarity, truncate(h.ChunkText, 2500))
	}
	b.WriteString("\nExtraia o que os trechos dizem sobre esta substância.")

	payload := map[string]any{
		"model":      model,
		"max_tokens": 2000,
		"system":     extractSystem,
		"messages":   []map[string]any{{"role": "user", "content": b.String()}},
		"tools": []map[string]any{{
			"name":         "registrar_extracao",
			"description":  "Registra o que os trechos dizem sobre a substância.",
			"input_schema": extractionSchema,
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "registrar_extracao"},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := (&http.Client{Timeout: 4 * time.Minute}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("claude %d: %s", resp.StatusCode, truncate(string(raw), 300))
	}

	var parsed struct {
		Content []struct {
			Type  string          `json:"type"`
			Name  string          `json:"name"`
			Input json.RawMessage `json:"input"`
		} `json:"content"`
		StopReason string `json:"stop_reason"`
	}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, err
	}
	for _, c := range parsed.Content {
		if c.Type == "tool_use" && c.Name == "registrar_extracao" {
			var out extraction
			if err := json.Unmarshal(c.Input, &out); err != nil {
				return nil, err
			}
			return &out, nil
		}
	}
	return nil, fmt.Errorf("resposta sem tool_use (stop_reason=%s)", parsed.StopReason)
}

var extractionSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"hasEvidence": map[string]any{
			"type":        "boolean",
			"description": "true se os trechos realmente tratam desta substância",
		},
		"indications": map[string]any{
			"type":        "string",
			"description": "Para que a substância é usada, segundo os trechos. Máximo 3 frases, PT-BR.",
		},
		"doseText": map[string]any{
			"type":        "string",
			"description": "Posologia como o trecho descreve (faixa, horário, com/sem alimento). Vazio se o trecho não der dose.",
		},
		"usualDose": map[string]any{"type": []string{"number", "null"}, "description": "Dose usual por dose/dia, só se explícita"},
		"minDose":   map[string]any{"type": []string{"number", "null"}, "description": "Limite inferior da faixa, só se explícito"},
		"maxDose":   map[string]any{"type": []string{"number", "null"}, "description": "Limite superior da faixa, só se explícito"},
		"unit":      map[string]any{"type": "string", "description": "Unidade da dose citada (mg, mcg, g, UI)"},
		"sourceTitles": map[string]any{
			"type":        "array",
			"items":       map[string]any{"type": "string"},
			"description": "Títulos dos trechos que sustentam o que foi extraído",
		},
	},
	"required": []string{"hasEvidence", "indications", "doseText", "sourceTitles"},
}

// persist grava evidência e sugestões. Não toca em dose já curada à mão.
func persist(db *gorm.DB, c *models.MagistralComponent, hits []chunkHit, ext *extraction) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, h := range hits {
			link := models.MagistralComponentArticle{
				ComponentID: c.ID,
				ArticleID:   h.ArticleID,
				EmbeddingID: &h.EmbeddingID,
				ChunkIndex:  &h.ChunkIndex,
				Similarity:  &h.Similarity,
				Excerpt:     truncate(h.ChunkText, 1200),
			}
			// Reprocessar não pode duplicar: a chave é (componente, artigo, chunk).
			if err := tx.Where("component_id = ? AND article_id = ? AND chunk_index = ?",
				c.ID, h.ArticleID, h.ChunkIndex).
				FirstOrCreate(&link, link).Error; err != nil {
				return err
			}
		}

		updates := map[string]any{
			"evidence_status": "suggested",
			"enriched_at":     time.Now(),
		}
		if ext.HasEvidence {
			if s := strings.TrimSpace(ext.Indications); s != "" {
				updates["indications"] = s
			}
			if s := strings.TrimSpace(ext.DoseText); s != "" {
				updates["dose_reference"] = s
			}
			// Dose numérica só entra em campo VAZIO e só se a unidade bater com a do catálogo:
			// 500 mcg gravado como 500 mg é o erro que este cuidado evita.
			sameUnit := strings.EqualFold(strings.TrimSpace(ext.Unit), c.DefaultUnit)
			if sameUnit && c.UsualDose == nil && ext.UsualDose != nil {
				updates["usual_dose"] = *ext.UsualDose
			}
			if sameUnit && c.MinDose == nil && ext.MinDose != nil {
				updates["min_dose"] = *ext.MinDose
			}
			if sameUnit && c.MaxDose == nil && ext.MaxDose != nil {
				updates["max_dose"] = *ext.MaxDose
			}
		}
		return tx.Model(&models.MagistralComponent{}).Where("id = ?", c.ID).Updates(updates).Error
	})
}

func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
