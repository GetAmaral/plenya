// Command normalize-patient-names põe os nomes já cadastrados no padrão "João da Silva".
//
// O hook do model cuida de quem entra a partir de agora; este comando é para o que já está lá.
// Roda com --dry-run por padrão: mostra o que mudaria e não grava. Só com --apply escreve.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
)

func main() {
	apply := flag.Bool("apply", false, "grava as mudanças (sem isto, só mostra)")
	limite := flag.Int("limit", 0, "processa no máximo N pacientes (0 = todos)")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("banco: %v", err)
	}
	db := database.DB

	type linha struct {
		ID   string
		Name string
	}
	var linhas []linha
	q := db.Table("patients").Select("id, name").Where("deleted_at IS NULL").Order("name")
	if *limite > 0 {
		q = q.Limit(*limite)
	}
	if err := q.Scan(&linhas).Error; err != nil {
		log.Fatalf("leitura: %v", err)
	}

	var mudam int
	for _, l := range linhas {
		novo := models.NormalizePersonName(l.Name)
		if novo == l.Name {
			continue
		}
		mudam++
		fmt.Printf("  %-40q → %q\n", l.Name, novo)
		if !*apply {
			continue
		}
		// UpdateColumn direto: o hook do model recalcula idade e mexe em cripto, e aqui só o
		// nome deve mudar.
		if err := db.Table("patients").Where("id = ?", l.ID).
			UpdateColumn("name", novo).Error; err != nil {
			log.Printf("falhou em %s: %v", l.ID, err)
		}
	}

	fmt.Printf("\n%d pacientes · %d fora do padrão\n", len(linhas), mudam)
	if !*apply && mudam > 0 {
		fmt.Println("nada foi gravado. rode de novo com --apply para aplicar.")
		os.Exit(0)
	}
}
