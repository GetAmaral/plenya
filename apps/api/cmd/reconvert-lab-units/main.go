// Reprocessa a unidade dos resultados de exame já gravados.
//
// Motivo: 520 resultados entraram em produção por carga em massa, direto na tabela, sem passar
// pela camada de serviço e portanto sem conversor — 520 linhas em quatro minutos, `unit_original`
// vazio. Outros 124 passaram pelo conversor mas o par (exame, unidade) não estava na tabela
// curada, e ele desistia em silêncio. Nos dois casos o número ficou gravado na unidade do laudo
// e depois foi comparado contra uma escala em outra grandeza.
//
// É idempotente: parte SEMPRE do valor original (`result_numeric_original`/`unit_original` quando
// existem, senão do valor atual), então rodar duas vezes não converte duas vezes.
//
//	go run ./cmd/reconvert-lab-units            # só relata
//	go run ./cmd/reconvert-lab-units -aplicar   # grava
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/models"
)

type linha struct {
	ID                    string
	TestName              string
	DefID                 string
	DefName               string
	DefUnit               *string
	DefCode               string
	Unit                  *string
	UnitOriginal          *string
	ResultNumeric         *float64
	ResultNumericOriginal *float64
}

func main() {
	aplicar := flag.Bool("aplicar", false, "grava as mudanças (sem isto, só relata)")
	flag.Parse()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}

	var linhas []linha
	if err := db.Raw(`
		SELECT lr.id, lr.test_name, d.id AS def_id, d.name AS def_name, d.unit AS def_unit, d.code AS def_code,
		       lr.unit, lr.unit_original, lr.result_numeric, lr.result_numeric_original
		  FROM lab_results lr
		  JOIN lab_test_definitions d ON d.id = lr.lab_test_definition_id
		 WHERE lr.deleted_at IS NULL AND lr.result_numeric IS NOT NULL AND lr.unit IS NOT NULL
		 ORDER BY d.name`).Scan(&linhas).Error; err != nil {
		panic(err)
	}

	// Faixas plausíveis por código de exame, numa consulta só.
	plausivel := carregaFaixas(db)

	var convertidos, jaOk, pendentes int
	porMotivo := map[string]int{}
	exemplos := map[string]string{}

	for _, l := range linhas {
		if l.DefUnit == nil {
			continue
		}
		valor, unidade := valorOriginalDe(l)

		def := models.LabTestDefinition{Unit: l.DefUnit}
		if err := db.First(&def, "id = ?", l.DefID).Error; err != nil {
			continue
		}
		conv := def.ConverteParaUnidadePrincipal(db, valor, unidade, plausivel[l.DefCode])

		switch conv.Status {
		case models.ConversaoAplicada:
			convertidos++
		case models.ConversaoPendente:
			pendentes++
			chave := fmt.Sprintf("%s: %s", l.DefName, conv.Motivo)
			porMotivo[chave]++
			if _, visto := exemplos[chave]; !visto {
				exemplos[chave] = fmt.Sprintf("valor %v %s", valor, unidade)
			}
		default:
			jaOk++
		}

		if !*aplicar {
			continue
		}
		campos := map[string]any{
			"result_numeric_original": valor,
			"unit_original":           unidade,
			"unit_conversion_status":  string(conv.Status),
			"unit_conversion_note":    nil,
			"result_numeric":          conv.Valor,
			"unit":                    conv.Unidade,
		}
		if conv.Status == models.ConversaoPendente && conv.Motivo != "" {
			campos["unit_conversion_note"] = conv.Motivo
		}
		if err := db.Table("lab_results").Where("id = ?", l.ID).Updates(campos).Error; err != nil {
			fmt.Printf("  ⚠️  %s (%s): %v\n", l.TestName, l.ID, err)
		}
	}

	fmt.Printf("%d resultados analisados\n", len(linhas))
	fmt.Printf("  já na unidade do exame : %d\n", jaOk)
	fmt.Printf("  convertidos            : %d\n", convertidos)
	fmt.Printf("  para revisão           : %d\n\n", pendentes)

	if pendentes > 0 {
		fmt.Println("PARA REVISÃO (unidade que o código não pode resolver sozinho):")
		chaves := make([]string, 0, len(porMotivo))
		for k := range porMotivo {
			chaves = append(chaves, k)
		}
		sort.Slice(chaves, func(i, j int) bool { return porMotivo[chaves[i]] > porMotivo[chaves[j]] })
		for _, k := range chaves {
			fmt.Printf("  %3dx  %s\n        exemplo: %s\n", porMotivo[k], k, exemplos[k])
		}
	}
	if !*aplicar {
		fmt.Println("\n(simulação — rode com -aplicar para gravar)")
	}
}

// valorOriginalDe devolve o par (valor, unidade) como veio do laudo. Partir daqui, e não do
// valor atual, é o que torna o reprocessamento repetível sem converter duas vezes.
func valorOriginalDe(l linha) (float64, string) {
	if l.UnitOriginal != nil && *l.UnitOriginal != "" && l.ResultNumericOriginal != nil {
		return *l.ResultNumericOriginal, *l.UnitOriginal
	}
	return *l.ResultNumeric, *l.Unit
}

func carregaFaixas(db *gorm.DB) map[string]func(float64) bool {
	var faixas []struct {
		Code  string
		Menor *float64
		Maior *float64
	}
	db.Raw(`
		SELECT si.lab_test_code AS code,
		       MIN(LEAST(NULLIF(sl.lower_limit,'')::double precision, NULLIF(sl.upper_limit,'')::double precision)) AS menor,
		       MAX(GREATEST(NULLIF(sl.lower_limit,'')::double precision, NULLIF(sl.upper_limit,'')::double precision)) AS maior
		  FROM score_levels sl
		  JOIN score_items si ON si.id = sl.score_item_id
		 WHERE sl.deleted_at IS NULL AND si.deleted_at IS NULL AND si.lab_test_code IS NOT NULL
		 GROUP BY 1`).Scan(&faixas)

	out := map[string]func(float64) bool{}
	for _, f := range faixas {
		if f.Menor == nil || f.Maior == nil || *f.Maior <= 0 || *f.Maior < *f.Menor {
			continue
		}
		menor, maior := *f.Menor, *f.Maior
		piso, teto := menor/10, maior*10
		if menor <= 0 {
			piso = menor*10 - 10
		}
		out[f.Code] = func(v float64) bool { return v >= piso && v <= teto }
	}
	return out
}
