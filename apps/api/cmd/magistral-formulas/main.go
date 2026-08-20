// Monta a biblioteca de fórmulas-base a partir das aulas da pós.
//
// O QUE FAZ: para cada frente clínica, busca no RAG os trechos que descrevem ASSOCIAÇÕES e
// fórmulas compostas, e pede ao Claude que extraia as fórmulas COMPLETAS ali descritas — nome,
// indicação e componentes com dose e unidade.
//
// O QUE NÃO FAZ:
//   · não inventa fórmula: componente sem dose no trecho não entra, e trecho que não descreve
//     associação não vira fórmula;
//   · não cria REGRA de dose. Regra é decisão clínica com trava de piso e teto; quem cadastra é
//     o médico, na tela de fórmulas-base;
//   · não confirma nada: tudo nasce para conferência, com os trechos de origem anexados.
//
// Uso:
//   go run ./cmd/magistral-formulas --dry-run
//   go run ./cmd/magistral-formulas --frente sono
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
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// frentes clínicas, com a consulta que traz os trechos certos do material.
var frentes = []struct {
	Chave   string
	Nome    string
	Consulta string
}{
	{"sono", "Sono e ansiedade",
		"fórmula manipulada composta para insônia e ansiedade: melatonina, 5-HTP, L-teanina, magnésio, associação e doses"},
	{"mitocondria", "Mitocôndria e energia",
		"fórmula manipulada de suporte mitocondrial e fadiga: coenzima Q10, PQQ, acetil-L-carnitina, ácido alfa-lipoico, associação e doses"},
	{"insulina", "Resistência insulínica",
		"fórmula manipulada para resistência insulínica e síndrome metabólica: berberina, inositol, cromo, ácido alfa-lipoico, associação e doses"},
	{"intestino", "Intestino e microbiota",
		"fórmula manipulada para permeabilidade intestinal e disbiose: glutamina, probiótico, prebiótico, zinco carnosina, associação e doses"},
	{"cognicao", "Cognição e foco",
		"fórmula manipulada para memória, foco e cognição: alfa-GPC, fosfatidilserina, citicolina, bacopa, associação e doses"},
	{"imunidade", "Imunidade e antioxidante",
		"fórmula manipulada antioxidante e de imunidade: vitamina C, zinco, quercetina, NAC, vitamina D, associação e doses"},
	{"menopausa", "Menopausa e hormonal",
		"fórmula manipulada para climatério e menopausa: isoflavona, cimicifuga, ashwagandha, associação e doses"},
	{"pele", "Pele, cabelo e unhas",
		"fórmula manipulada para pele, cabelos e unhas: silício orgânico, colágeno, biotina, ácido hialurônico, associação e doses"},
	// Frentes alinhadas aos módulos com mais aulas: é onde o material costuma trazer a fórmula
	// escrita, com dose, em vez de discutir a substância isolada.
	{"intestino5r", "Modulação intestinal",
		"protocolo de modulação intestinal e disbiose: prescrição com glutamina, butirato, zinco carnosina, probiótico, doses por dose e por sachê"},
	{"tdah", "TDAH e foco",
		"prescrição para TDAH: associação de nutrientes e fitoterápicos com doses, fórmula manipulada, ômega 3, magnésio, zinco, ferro"},
	{"emagrecimento", "Emagrecimento",
		"fórmula manipulada para emagrecimento e saciedade: associação com doses por cápsula, termogênico, controle de apetite"},
	{"hormonal", "Reposição hormonal",
		"prescrição de reposição hormonal feminina e masculina: fórmula manipulada com doses, creme transdérmico, cápsula"},
	{"humor", "Humor e psiquiatria",
		"prescrição em psiquiatria metabólica: associação de nutrientes com doses para depressão e ansiedade, fórmula manipulada"},
}

type formulaOut struct {
	Name       string `json:"name"`
	Indication string `json:"indication"`
	Bullets    []string `json:"indicationBullets"`
	Form       string `json:"pharmaceuticalForm"`
	Usage      string `json:"usageType"`
	Vehicle    string `json:"vehicle"`
	Dispense   float64 `json:"quantityToDispense"`
	Unit       string  `json:"quantityUnit"`
	Posology   string  `json:"posology"`
	Duration   int     `json:"duration"`
	Components []struct {
		Substance string  `json:"substance"`
		Quantity  float64 `json:"quantity"`
		Unit      string  `json:"unit"`
		Note      string  `json:"note"`
	} `json:"components"`
	SourceTitles []string `json:"sourceTitles"`
}

type chunk struct {
	EmbeddingID  uuid.UUID
	ArticleID    uuid.UUID
	ArticleTitle string
	ArticleType  string
	ChunkIndex   int
	ChunkText    string
	Similarity   float64
}

func main() {
	var (
		dryRun = flag.Bool("dry-run", false, "não grava; só mostra")
		frente = flag.String("frente", "", "processa só esta frente (chave)")
		chunks = flag.Int("chunks", 10, "trechos por frente")
		minSim = flag.Float64("min-similarity", 0.40, "similaridade mínima")
		model  = flag.String("model", "claude-opus-5", "modelo do Claude")
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
	embeddings := services.NewEmbeddingService(cfg, db)
	templates := services.NewMagistralTemplateService(db)
	ctx := context.Background()

	var doctor models.User
	if err := db.Where("email = ?", "admin@plenya.com").First(&doctor).Error; err != nil {
		log.Fatal(err)
	}

	var criadas int
	for _, f := range frentes {
		if *frente != "" && *frente != f.Chave {
			continue
		}
		fmt.Printf("\n=== %s\n", f.Nome)

		hits, err := buscar(ctx, db, embeddings, f.Consulta, *chunks, *minSim)
		if err != nil {
			fmt.Printf("   ⚠️  busca falhou: %v\n", err)
			continue
		}
		if len(hits) == 0 {
			fmt.Println("   — sem trecho relevante")
			continue
		}
		fmt.Printf("   %d trecho(s), melhor %.3f · %s\n", len(hits), hits[0].Similarity, hits[0].ArticleTitle)

		formulas, err := extrair(cfg.Claude.APIKey, *model, f.Nome, hits)
		if err != nil {
			fmt.Printf("   ⚠️  extração falhou: %v\n", err)
			continue
		}
		if len(formulas) == 0 {
			fmt.Println("   — os trechos não descrevem fórmula completa")
			continue
		}

		for _, fo := range formulas {
			if len(fo.Components) == 0 {
				continue
			}
			fmt.Printf("   • %s (%s, %d componentes)\n", fo.Name, fo.Form, len(fo.Components))
			for _, c := range fo.Components {
				fmt.Printf("       %s %.4g %s\n", c.Substance, c.Quantity, c.Unit)
			}
			if *dryRun {
				continue
			}

			req := montar(fo)
			tpl, err := templates.Save(nil, req, doctor.ID)
			if err != nil {
				fmt.Printf("       ⚠️  não gravou: %v\n", err)
				continue
			}
			if err := anexarEvidencia(db, tpl.ID, hits); err != nil {
				fmt.Printf("       ⚠️  evidência não anexada: %v\n", err)
			}
			criadas++
		}
	}

	fmt.Printf("\n📊 %d fórmula(s)-base criada(s)\n", criadas)
	if *dryRun {
		fmt.Println("(dry-run: nada gravado)")
	}
}

func montar(fo formulaOut) *dto.FormulaTemplateRequest {
	usage := models.FormulaUsageInternal
	if strings.EqualFold(fo.Usage, "external") {
		usage = models.FormulaUsageExternal
	}
	form := strings.TrimSpace(fo.Form)
	if form == "" {
		form = "cápsula"
	}
	unit := strings.TrimSpace(fo.Unit)
	if unit == "" {
		unit = "cápsulas"
	}
	dispense := fo.Dispense
	if dispense <= 0 {
		dispense = 60
	}

	req := &dto.FormulaTemplateRequest{
		Name:               strings.TrimSpace(fo.Name),
		PharmaceuticalForm: form,
		UsageType:          usage,
		Vehicle:            strings.TrimSpace(fo.Vehicle),
		QuantityToDispense: dispense,
		QuantityUnit:       unit,
		Posology:           strings.TrimSpace(fo.Posology),
		Duration:           fo.Duration,
	}
	if s := strings.TrimSpace(fo.Indication); s != "" {
		req.Indication = &s
	}
	if len(fo.Bullets) > 0 {
		b := strings.Join(fo.Bullets, "\n")
		req.IndicationBullets = &b
	}
	for _, c := range fo.Components {
		if strings.TrimSpace(c.Substance) == "" || c.Quantity <= 0 {
			continue
		}
		u := strings.TrimSpace(c.Unit)
		if u == "" {
			u = "mg"
		}
		req.Components = append(req.Components, dto.FormulaTemplateComponentRequest{
			Substance: strings.TrimSpace(c.Substance),
			Quantity:  c.Quantity,
			Unit:      u,
			Category:  models.MedCategorySimple,
			Note:      strings.TrimSpace(c.Note),
		})
	}
	return req
}

func buscar(ctx context.Context, db *gorm.DB, emb *services.EmbeddingService,
	consulta string, k int, minSim float64) ([]chunk, error) {

	vec, err := emb.GenerateEmbedding(ctx, consulta)
	if err != nil {
		return nil, err
	}
	var hits []chunk
	err = db.Raw(`
		SELECT ae.id AS embedding_id, ae.article_id, a.title AS article_title, a.article_type,
		       ae.chunk_index, ae.chunk_text, 1 - (ae.embedding <=> ?) AS similarity
		FROM article_embeddings ae
		JOIN articles a ON a.id = ae.article_id
		WHERE a.deleted_at IS NULL AND ae.embedding IS NOT NULL
		  AND 1 - (ae.embedding <=> ?) >= ?
		ORDER BY (1 - (ae.embedding <=> ?)) + CASE WHEN a.article_type = 'lecture' THEN 0.03 ELSE 0 END DESC
		LIMIT ?`,
		pgvector.NewVector(vec), pgvector.NewVector(vec), minSim, pgvector.NewVector(vec), k).
		Scan(&hits).Error
	return hits, err
}

const sistema = `Você extrai FÓRMULAS MAGISTRAIS COMPLETAS de material didático de medicina funcional integrativa.

REGRAS INVIOLÁVEIS:
1. Só extraia fórmula que esteja DESCRITA nos trechos, com componentes e doses. Nunca componha uma fórmula juntando substâncias que aparecem separadas em contextos diferentes.
2. Componente sem dose explícita não entra. Fórmula que ficaria com menos de dois componentes com dose não entra.
3. Não invente veículo, posologia ou quantidade a aviar: deixe vazio quando o trecho não disser.
4. Nome curto e clínico, em português ("Fórmula do sono", "Suporte mitocondrial"). Indicação em até 2 frases e em até 4 tópicos de no máximo 8 palavras.
5. Se os trechos não trouxerem nenhuma fórmula completa, devolva lista vazia. Isso é resposta legítima.
6. Português do Brasil, sem travessão e sem fecho de efeito.`

func extrair(apiKey, model, frente string, hits []chunk) ([]formulaOut, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "FRENTE CLÍNICA: %s\n\nTRECHOS DO MATERIAL:\n", frente)
	for i, h := range hits {
		fmt.Fprintf(&b, "\n[%d] %s (%s, similaridade %.3f)\n%s\n", i+1, h.ArticleTitle, h.ArticleType, h.Similarity, corta(h.ChunkText, 3000))
	}
	b.WriteString("\nExtraia as fórmulas completas descritas nos trechos.")

	payload := map[string]any{
		"model":      model,
		"max_tokens": 6000,
		"system":     sistema,
		"messages":   []map[string]any{{"role": "user", "content": b.String()}},
		"tools": []map[string]any{{
			"name":        "registrar_formulas",
			"description": "Fórmulas magistrais completas descritas nos trechos.",
			"input_schema": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"formulas": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type": "object",
							"properties": map[string]any{
								"name":               map[string]any{"type": "string"},
								"indication":         map[string]any{"type": "string"},
								"indicationBullets":  map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
								"pharmaceuticalForm": map[string]any{"type": "string", "description": "cápsula, sachê, creme, gel, solução oral"},
								"usageType":          map[string]any{"type": "string", "enum": []string{"internal", "external"}},
								"vehicle":            map[string]any{"type": "string"},
								"quantityToDispense": map[string]any{"type": "number"},
								"quantityUnit":       map[string]any{"type": "string"},
								"posology":           map[string]any{"type": "string"},
								"duration":           map[string]any{"type": "integer"},
								"components": map[string]any{
									"type": "array",
									"items": map[string]any{
										"type": "object",
										"properties": map[string]any{
											"substance": map[string]any{"type": "string"},
											"quantity":  map[string]any{"type": "number"},
											"unit":      map[string]any{"type": "string"},
											"note":      map[string]any{"type": "string"},
										},
										"required": []string{"substance", "quantity", "unit"},
									},
								},
								"sourceTitles": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
							},
							"required": []string{"name", "components"},
						},
					},
				},
				"required": []string{"formulas"},
			},
		}},
		"tool_choice": map[string]any{"type": "tool", "name": "registrar_formulas"},
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := (&http.Client{Timeout: 6 * time.Minute}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("claude %d: %s", resp.StatusCode, corta(string(raw), 300))
	}

	var parsed struct {
		Content []struct {
			Type  string          `json:"type"`
			Name  string          `json:"name"`
			Input json.RawMessage `json:"input"`
		} `json:"content"`
	}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, err
	}
	for _, c := range parsed.Content {
		if c.Type == "tool_use" && c.Name == "registrar_formulas" {
			var out struct {
				Formulas []formulaOut `json:"formulas"`
			}
			if err := json.Unmarshal(c.Input, &out); err != nil {
				return nil, err
			}
			return out.Formulas, nil
		}
	}
	return nil, fmt.Errorf("resposta sem tool_use")
}

// anexarEvidencia guarda os trechos que sustentaram a extração, como no catálogo de substâncias.
func anexarEvidencia(db *gorm.DB, templateID uuid.UUID, hits []chunk) error {
	for _, h := range hits {
		rec := map[string]any{
			"id": uuid.Must(uuid.NewV7()), "template_id": templateID, "article_id": h.ArticleID,
			"embedding_id": h.EmbeddingID, "chunk_index": h.ChunkIndex, "similarity": h.Similarity,
			"excerpt": corta(h.ChunkText, 1200),
		}
		if err := db.Table("magistral_formula_template_articles").Create(rec).Error; err != nil {
			return err
		}
	}
	return nil
}

func corta(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
