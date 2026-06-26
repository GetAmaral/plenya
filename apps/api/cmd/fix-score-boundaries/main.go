// Command fix-score-boundaries — fecha os vãos das faixas de score (não-genéticas):
// para cada faixa `between`, sobe o upper até encostar no lower da faixa seguinte (por valor);
// extremo baixo `<` → `<=`; extremo alto `>=` → `>`. Idempotente. Dry-run por padrão.
//
//	./fix-score-boundaries          # dry-run (só imprime)
//	./fix-score-boundaries --apply  # grava
package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

func pf(s *string) (float64, bool) {
	if s == nil {
		return 0, false
	}
	v, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(*s), ",", "."), 64)
	return v, err == nil
}
func fmtNum(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }

func main() {
	apply := flag.Bool("apply", false, "grava as mudanças (senão dry-run)")
	flag.Parse()
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	type row struct {
		ID         string
		Name       string
		LabCode    *string
		Level      int
		Operator   string
		LowerLimit *string
		UpperLimit *string
	}
	var rows []row
	db.Table("score_levels sl").
		Select("sl.id, si.name, si.lab_test_code as lab_code, sl.level, sl.operator, sl.lower_limit, sl.upper_limit").
		Joins("JOIN score_items si ON si.id = sl.item_id AND si.deleted_at IS NULL").
		Where("sl.deleted_at IS NULL").
		Where("si.lab_test_code IS NULL OR si.lab_test_code NOT LIKE 'GEN_%'").
		Scan(&rows)

	byItem := map[string][]*row{}
	order := []string{}
	for i := range rows {
		r := &rows[i]
		key := r.Name
		if _, ok := byItem[key]; !ok {
			order = append(order, key)
		}
		byItem[key] = append(byItem[key], r)
	}

	type change struct {
		id, name, what, before, after string
	}
	var changes []change

	// Itens RE-ESTRUTURADOS à mão (decisões em docs/emr/score-faixas-decisoes.md) — o tool
	// não toca neles. Os com nível '=' (discretas/qualitativas/dead-= manuais) são pulados abaixo.
	manualSkip := map[string]bool{
		"RDW": true, "Homocisteína": true, "Tempo de sono": true,
		"Consumo de Calorias (profissional calcula com base na taxa metabólica + atividades)": true,
		"Hora de acordar": true, "Hora de dormir": true,
		"Regularidade no acordar": true, "Regularidade no dormir": true,
	}

	for _, key := range order {
		bands := byItem[key]
		// Pula itens decididos à mão e itens com nível '=' (escalas discretas tipo digit span,
		// qualitativos, e os '=' numéricos já tratados manualmente) — não levam ajuste de fronteira.
		if manualSkip[key] {
			continue
		}
		hasEq := false
		for _, r := range bands {
			if r.Operator == "=" {
				hasEq = true
				break
			}
		}
		if hasEq {
			continue
		}
		// considera só operadores numéricos com limites parseáveis
		type b struct {
			r        *row
			pos, rnk float64
			lower    float64
			hasLower bool
			upper    float64
			hasUpper bool
		}
		var bs []b
		for _, r := range bands {
			op := r.Operator
			if op != "<" && op != "<=" && op != ">" && op != ">=" && op != "between" {
				continue
			}
			lo, okLo := pf(r.LowerLimit)
			up, okUp := pf(r.UpperLimit)
			var pos, rnk float64
			switch op {
			case "<", "<=":
				if !okUp {
					continue
				}
				pos, rnk = up, 0
			case "between":
				if !okLo || !okUp {
					continue
				}
				pos, rnk = lo, 1
			case ">", ">=":
				if !okLo {
					continue
				}
				pos, rnk = lo, 2
			}
			bs = append(bs, b{r: r, pos: pos, rnk: rnk, lower: lo, hasLower: okLo, upper: up, hasUpper: okUp})
		}
		sort.SliceStable(bs, func(i, j int) bool {
			if bs[i].pos != bs[j].pos {
				return bs[i].pos < bs[j].pos
			}
			return bs[i].rnk < bs[j].rnk
		})
		for i, band := range bs {
			r := band.r
			if r.Operator == "between" {
				// 1) upper sobe até o lower da próxima faixa (between-between e between→extremo alto)
				if i+1 < len(bs) && bs[i+1].hasLower {
					curUp, okCur := pf(r.UpperLimit)
					cur := ""
					if r.UpperLimit != nil {
						cur = *r.UpperLimit
					}
					if !okCur || curUp != bs[i+1].lower {
						newUp := fmtNum(bs[i+1].lower)
						changes = append(changes, change{r.ID, key, "upper", cur, newUp})
						if *apply {
							db.Model(&models.ScoreLevel{}).Where("id = ?", r.ID).Update("upper_limit", newUp)
						}
					}
				}
				// 2) se a faixa ANTERIOR é o extremo baixo (</<=), o EXTREMO é sagrado:
				//    baixa o lower DESTE between até encostar no upper do extremo (2.1 → 2)
				if i > 0 && (bs[i-1].r.Operator == "<" || bs[i-1].r.Operator == "<=") && bs[i-1].hasUpper {
					curLo, okCur := pf(r.LowerLimit)
					cur := ""
					if r.LowerLimit != nil {
						cur = *r.LowerLimit
					}
					if !okCur || curLo != bs[i-1].upper {
						newLo := fmtNum(bs[i-1].upper)
						changes = append(changes, change{r.ID, key, "lower", cur, newLo})
						if *apply {
							db.Model(&models.ScoreLevel{}).Where("id = ?", r.ID).Update("lower_limit", newLo)
						}
					}
				}
			}
			// 3) extremo baixo < → <= (upper NÃO muda — é sagrado)
			if r.Operator == "<" {
				changes = append(changes, change{r.ID, key, "op", "<", "<="})
				if *apply {
					db.Model(&models.ScoreLevel{}).Where("id = ?", r.ID).Update("operator", "<=")
				}
			}
			// 4) extremo alto >= → >
			if r.Operator == ">=" {
				changes = append(changes, change{r.ID, key, "op", ">=", ">"})
				if *apply {
					db.Model(&models.ScoreLevel{}).Where("id = ?", r.ID).Update("operator", ">")
				}
			}
		}
	}

	upper, ops := 0, 0
	for _, c := range changes {
		if c.what == "upper" {
			upper++
		} else {
			ops++
		}
	}
	mode := "DRY-RUN (nada gravado)"
	if *apply {
		mode = "APLICADO"
	}
	chItems := map[string]bool{}
	for _, c := range changes {
		chItems[c.name] = true
	}
	fmt.Printf("== %s ==\nuppers ajustados: %d | operadores de extremo trocados: %d | itens alterados: %d\n",
		mode, upper, ops, len(chItems))
	for _, name := range []string{"LDL Colesterol", "Anti-TPO", "Albumina"} {
		fmt.Printf("\n%s:\n", name)
		for _, c := range changes {
			if c.name == name {
				fmt.Printf("   [%s] %s → %s\n", c.what, c.before, c.after)
			}
		}
	}
}
