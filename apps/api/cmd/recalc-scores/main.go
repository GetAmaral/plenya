// Recalcula o escore dos pacientes direto pelo motor, sem passar pela API.
//
// Toda vez que o catálogo muda — faixa de nível, unidade de item, condição de aplicabilidade,
// conversão de unidade dos resultados — os snapshots já gravados ficam desatualizados. Fazer isso
// por HTTP exige um token de sessão válido, o que transforma uma operação de manutenção numa
// dependência de alguém estar logado. Aqui não.
//
// A operação é ADITIVA quando o último snapshot é de outro dia: o histórico fica. Se já existe
// snapshot de hoje, o motor o reaproveita e substitui os resultados dele (mesma semântica do
// endpoint).
//
//	recalc-scores                    # todos os pacientes que já têm escore
//	recalc-scores -paciente <uuid>   # só um
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
)

func main() {
	umPaciente := flag.String("paciente", "", "UUID de um paciente; vazio = todos que já têm escore")
	flag.Parse()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env("DB_HOST", "localhost"), env("DB_PORT", "5432"), env("DB_USER", "plenya_user"),
			os.Getenv("DB_PASSWORD"), env("DB_NAME", "plenya_db"))
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}

	// O model Patient descriptografa CPF/RG no AfterFind: sem a chave, carregar o paciente falha
	// e o recálculo não sai do lugar.
	if err := crypto.Init(os.Getenv("ENCRYPTION_KEY")); err != nil {
		fmt.Println("ENCRYPTION_KEY:", err)
		os.Exit(1)
	}
	blind := os.Getenv("BLIND_INDEX_KEY")
	if blind == "" {
		blind = os.Getenv("ENCRYPTION_KEY")
	}
	crypto.SetBlindIndexKey(blind)

	// Autoria do recálculo: um usuário administrativo real, para o snapshot não ficar órfão.
	var autorTexto string
	// `roles` é um array JSON, não coluna escalar. `jsonb_exists` e não o operador `?`, que o
	// GORM leria como placeholder de parâmetro.
	if err := db.Raw(`SELECT id FROM users WHERE deleted_at IS NULL AND jsonb_exists(roles::jsonb, 'admin')
	                   ORDER BY created_at LIMIT 1`).Scan(&autorTexto).Error; err != nil || autorTexto == "" {
		fmt.Println("não achei um usuário admin para assinar o recálculo")
		os.Exit(1)
	}
	autor, err := uuid.Parse(autorTexto)
	if err != nil {
		fmt.Println("id de usuário inválido:", autorTexto)
		os.Exit(1)
	}

	var pacientes []struct {
		ID   uuid.UUID
		Name string
	}
	q := db.Raw(`SELECT DISTINCT p.id, p.name FROM patient_score_snapshots s
	              JOIN patients p ON p.id = s.patient_id ORDER BY p.name`)
	if *umPaciente != "" {
		id, err := uuid.Parse(*umPaciente)
		if err != nil {
			fmt.Println("uuid inválido:", *umPaciente)
			os.Exit(1)
		}
		q = db.Raw(`SELECT id, name FROM patients WHERE id = ?`, id)
	}
	if err := q.Scan(&pacientes).Error; err != nil {
		panic(err)
	}

	svc := services.NewScoreSnapshotService(
		repository.NewScoreSnapshotRepository(db),
		repository.NewScoreRepository(db),
		repository.NewLabResultRepository(db),
		repository.NewAnamnesisRepository(db),
		db,
	)

	fmt.Printf("%-40s %8s %8s %s\n", "paciente", "antes", "depois", "itens avaliados")
	for _, p := range pacientes {
		antes, antesN := ultimoEscore(db, p.ID)

		if _, err := svc.CalculateSnapshot(services.CalculateSnapshotDTO{PatientID: p.ID}, autor); err != nil {
			fmt.Printf("%-40s  ERRO: %v\n", corta(p.Name, 40), err)
			continue
		}

		depois, depoisN := ultimoEscore(db, p.ID)
		marca := "  "
		if fmt.Sprintf("%.1f", antes) != fmt.Sprintf("%.1f", depois) || antesN != depoisN {
			marca = "->"
		}
		fmt.Printf("%-40s %7.1f%% %7.1f%%  %4d -> %-4d %s\n", corta(p.Name, 40), antes, depois, antesN, depoisN, marca)
	}
}

func ultimoEscore(db *gorm.DB, patientID uuid.UUID) (float64, int) {
	var r struct {
		Pct float64
		N   int
	}
	db.Raw(`SELECT total_score_percentage AS pct, items_evaluated_count AS n
	          FROM patient_score_snapshots WHERE patient_id = ?
	         ORDER BY calculated_at DESC LIMIT 1`, patientID).Scan(&r)
	return r.Pct, r.N
}

func corta(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

func env(k, padrao string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return padrao
}
