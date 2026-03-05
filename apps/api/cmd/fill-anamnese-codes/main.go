package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// removeAccents remove acentos de caracteres Unicode, convertendo para ASCII equivalente
func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Remove combining marks (acentos)
	}), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

var (
	// Remove conteúdo entre parênteses (unidades como "(cm)", "(%)", "(°)", "(kg/m²)")
	reParenthetical = regexp.MustCompile(`\s*\([^)]*\)\s*`)
	// Remove todos os chars que não são letras/números/espaço/hífen
	reNonAlphanumeric = regexp.MustCompile(`[^a-zA-Z0-9\s\-]+`)
	// Converte espaços, hífens e underscores múltiplos em um único underscore
	reCollapseUnderscores = regexp.MustCompile(`[_\s\-]+`)
	// Trim underscores das bordas
	reTrimUnderscores = regexp.MustCompile(`^_+|_+$`)
)

// generateCode gera um código estável a partir do nome do ScoreItem
// Ex: "Coxa (cm) - homem" → "COXA_HOMEM"
//
//	"Ângulo de fase (°) - mulher" → "ANGULO_DE_FASE_MULHER"
//	"Razão androide/ginoide - mulher" → "RAZAO_ANDROIDE_GINOIDE_MULHER"
func generateCode(name string) string {
	// 1. Remove conteúdo parentético (unidades)
	s := reParenthetical.ReplaceAllString(name, " ")

	// 2. Remove acentos
	s = removeAccents(s)

	// 3. Substitui "/" por "_" antes de remover chars especiais
	s = strings.ReplaceAll(s, "/", "_")

	// 4. Remove chars não alfanuméricos (exceto espaço e hífen)
	s = reNonAlphanumeric.ReplaceAllString(s, " ")

	// 5. Colapsa espaços/hífens em underscore
	s = reCollapseUnderscores.ReplaceAllString(s, "_")

	// 6. Remove underscores das bordas
	s = reTrimUnderscores.ReplaceAllString(s, "")

	// 7. Uppercase
	return strings.ToUpper(s)
}

func main() {
	fmt.Println("=== Fill AnamneseItemCode Script ===\n")

	// Carregar .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	if err := database.Connect(cfg); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()
	fmt.Println("✅ Database connected\n")

	// Buscar todos os ScoreItems sem lab_test_code
	var items []models.ScoreItem
	if err := database.DB.
		Where("lab_test_code IS NULL AND deleted_at IS NULL").
		Find(&items).Error; err != nil {
		log.Fatalf("❌ Failed to fetch score items: %v", err)
	}

	fmt.Printf("📊 %d score items sem lab_test_code\n\n", len(items))

	// Rastrear codes gerados para detectar duplicatas
	codeCount := make(map[string]int)
	codeToID := make(map[string]string) // code → primeiro item ID (para diagnóstico)

	// Primeira passagem: gerar todos os codes e detectar duplicatas
	type itemCode struct {
		item models.ScoreItem
		code string
	}
	planned := make([]itemCode, 0, len(items))

	for _, item := range items {
		code := generateCode(item.Name)
		codeCount[code]++
		if _, exists := codeToID[code]; !exists {
			codeToID[code] = item.ID.String()
		}
		planned = append(planned, itemCode{item: item, code: code})
	}

	// Detectar duplicatas
	hasDuplicates := false
	for code, count := range codeCount {
		if count > 1 {
			hasDuplicates = true
			fmt.Printf("⚠️  Duplicata detectada: '%s' (%d ocorrências)\n", code, count)
		}
	}
	if hasDuplicates {
		fmt.Println()
	}

	// Segunda passagem: resolver duplicatas com sufixo numérico
	dupCounter := make(map[string]int)
	var (
		updated   = 0
		unchanged = 0
		errors    = 0
	)

	fmt.Println("Processando items:")
	fmt.Println("----------------------------------------")

	for _, ic := range planned {
		finalCode := ic.code

		// Se há duplicata, adiciona sufixo _2, _3, etc. (primeiro não recebe sufixo)
		if codeCount[ic.code] > 1 {
			dupCounter[ic.code]++
			if dupCounter[ic.code] > 1 {
				finalCode = fmt.Sprintf("%s_%d", ic.code, dupCounter[ic.code])
			}
		}

		// Verifica se já tem o code correto
		if ic.item.AnamneseItemCode != nil && *ic.item.AnamneseItemCode == finalCode {
			unchanged++
			continue
		}

		// Atualiza usando Exec para não disparar hooks de embedding/review
		if err := database.DB.Exec(
			`UPDATE score_items SET anamnese_item_code = ? WHERE id = ?`,
			finalCode, ic.item.ID,
		).Error; err != nil {
			log.Printf("❌ Erro ao atualizar %s (%s): %v\n", ic.item.Name, ic.item.ID, err)
			errors++
			continue
		}

		fmt.Printf("✓ %-50s → %s\n", ic.item.Name, finalCode)
		updated++
	}

	fmt.Println("\n----------------------------------------")
	fmt.Println("=== Resumo ===")
	fmt.Printf("Atualizados:  %d\n", updated)
	fmt.Printf("Sem mudança:  %d\n", unchanged)
	fmt.Printf("Erros:        %d\n", errors)

	if errors > 0 {
		fmt.Println("\n❌ Concluído com erros")
		os.Exit(1)
	}
	fmt.Println("\n✅ Concluído com sucesso!")
}
