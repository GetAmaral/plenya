// Converte as indicações e a posologia já gravadas em TÓPICOS curtos.
//
// É uma compressão FIEL do texto que já está no catálogo: nenhuma busca nova, nenhuma fonte
// nova, nada que não esteja escrito no campo de origem. Se o texto não diz, o tópico não existe.
//
// Roda em lotes (várias substâncias por chamada) porque a tarefa é curta e o gargalo é a
// ida-e-volta da API.
//
// Uso:
//   go run ./cmd/magistral-bullets --dry-run --limit 8
//   go run ./cmd/magistral-bullets --redo
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
)

type bulletItem struct {
	Name              string   `json:"name"`
	IndicationBullets []string `json:"indicationBullets"`
	DoseBullets       []string `json:"doseBullets"`
}

func main() {
	var (
		dryRun    = flag.Bool("dry-run", false, "não grava; só mostra")
		limit     = flag.Int("limit", 0, "no máximo N substâncias (0 = todas)")
		batch     = flag.Int("batch", 8, "substâncias por chamada")
		redo      = flag.Bool("redo", false, "refaz quem já tem tópicos")
		modelName = flag.String("model", "claude-opus-5", "modelo do Claude")
	)
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := database.Connect(cfg); err != nil {
		log.Fatal(err)
	}
	db := database.DB

	var components []models.MagistralComponent
	q := db.Where("deleted_at IS NULL AND (indications IS NOT NULL OR dose_reference IS NOT NULL)")
	if !*redo {
		q = q.Where("indication_bullets IS NULL AND dose_bullets IS NULL")
	}
	if err := q.Order("usage_count DESC").Order("name").Find(&components).Error; err != nil {
		log.Fatal(err)
	}
	if *limit > 0 && len(components) > *limit {
		components = components[:*limit]
	}

	fmt.Printf("✂️  %d substância(s) para resumir em tópicos\n\n", len(components))

	var feitas int
	for i := 0; i < len(components); i += *batch {
		end := i + *batch
		if end > len(components) {
			end = len(components)
		}
		lote := components[i:end]

		items, err := summarize(cfg.Claude.APIKey, *modelName, lote)
		if err != nil {
			fmt.Printf("⚠️  lote %d-%d falhou: %v\n", i+1, end, err)
			continue
		}

		byName := map[string]bulletItem{}
		for _, it := range items {
			byName[strings.ToLower(strings.TrimSpace(it.Name))] = it
		}

		for _, c := range lote {
			it, ok := byName[strings.ToLower(c.Name)]
			if !ok {
				fmt.Printf("   %s — o modelo não devolveu tópicos\n", c.Name)
				continue
			}
			ind := strings.Join(it.IndicationBullets, "\n")
			dose := strings.Join(it.DoseBullets, "\n")
			fmt.Printf("   %s\n", c.Name)
			for _, b := range it.IndicationBullets {
				fmt.Printf("      · %s\n", b)
			}
			for _, b := range it.DoseBullets {
				fmt.Printf("      → %s\n", b)
			}
			if *dryRun {
				continue
			}
			updates := map[string]any{}
			if ind != "" {
				updates["indication_bullets"] = ind
			}
			if dose != "" {
				updates["dose_bullets"] = dose
			}
			if len(updates) == 0 {
				continue
			}
			// UpdateColumns: resumir não é revisão clínica, então não carimba last_review.
			if err := db.Model(&models.MagistralComponent{}).Where("id = ?", c.ID).
				UpdateColumns(updates).Error; err != nil {
				fmt.Printf("      ⚠️  gravação falhou: %v\n", err)
				continue
			}
			feitas++
		}
	}

	fmt.Printf("\n📊 %d substância(s) com tópicos\n", feitas)
	if *dryRun {
		fmt.Println("(dry-run: nada gravado)")
	}
	_ = gorm.ErrRecordNotFound
}

const bulletSystem = `Você resume, em tópicos, texto de catálogo de prescrição magistral.

REGRAS INVIOLÁVEIS:
1. O tópico é COMPRESSÃO FIEL do texto recebido. Não acrescente indicação, dose ou ressalva que não esteja escrita ali.
2. Indicações: no máximo 5 tópicos, cada um com no máximo 8 palavras. Comece pelo uso principal. Nomeie a condição ou o público ("uso de estatina", "insuficiência cardíaca", "fertilidade"), não o mecanismo, salvo quando o mecanismo for a única coisa que o texto diz.
3. Posologia: no máximo 4 tópicos, cada um com no máximo 12 palavras, sempre com o número e a unidade quando o texto tiver ("100 mg/dia com refeição gordurosa").
4. Português do Brasil, sem travessão, sem ponto final, sem repetir o nome da substância no começo do tópico.
5. Campo sem texto de origem devolve lista vazia.`

func summarize(apiKey, model string, lote []models.MagistralComponent) ([]bulletItem, error) {
	var b strings.Builder
	for _, c := range lote {
		fmt.Fprintf(&b, "\n### %s\n", c.Name)
		if c.Indications != nil && strings.TrimSpace(*c.Indications) != "" {
			fmt.Fprintf(&b, "INDICAÇÕES: %s\n", strings.TrimSpace(*c.Indications))
		}
		if c.DoseReference != nil && strings.TrimSpace(*c.DoseReference) != "" {
			fmt.Fprintf(&b, "POSOLOGIA: %s\n", strings.TrimSpace(*c.DoseReference))
		}
	}

	payload := map[string]any{
		"model":      model,
		"max_tokens": 4000,
		"system":     bulletSystem,
		"messages": []map[string]any{{
			"role":    "user",
			"content": "Resuma em tópicos cada substância abaixo." + b.String(),
		}},
		"tools": []map[string]any{{
			"name":        "registrar_topicos",
			"description": "Tópicos de indicação e posologia por substância.",
			"input_schema": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"items": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type": "object",
							"properties": map[string]any{
								"name":              map[string]any{"type": "string"},
								"indicationBullets": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
								"doseBullets":       map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
							},
							"required": []string{"name", "indicationBullets", "doseBullets"},
						},
					},
				},
				"required": []string{"items"},
			},
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "registrar_topicos"},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := (&http.Client{Timeout: 5 * time.Minute}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("claude %d: %s", resp.StatusCode, snippet(string(raw)))
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
		if c.Type == "tool_use" && c.Name == "registrar_topicos" {
			var out struct {
				Items []bulletItem `json:"items"`
			}
			if err := json.Unmarshal(c.Input, &out); err != nil {
				return nil, err
			}
			return out.Items, nil
		}
	}
	return nil, fmt.Errorf("resposta sem tool_use (stop_reason=%s)", parsed.StopReason)
}

func snippet(s string) string {
	if len(s) > 300 {
		return s[:300]
	}
	return s
}
