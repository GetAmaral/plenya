// Command migrate aplica as migrations de schema do EMR via goose.
//
// Uso:
//
//	migrate up       aplica as migrations pendentes (caminho do deploy)
//	migrate status   mostra o estado de cada migration
//	migrate version  mostra a versão atual do schema
//	migrate down     reverte a última migration (apenas dev)
//	migrate stamp    marca a baseline (00001) como aplicada SEM executá-la,
//	                 para adotar o goose num banco que JÁ tem o schema (dev/prod existentes)
//
// O schema é fonte dos Go models; quando um model muda, cria-se uma nova
// migration goose (00002_*.sql, ...). Em produção isto roda como passo do
// deploy, antes de a aplicação subir — nunca via AutoMigrate no boot.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/plenya/api/database/migrations"
	"github.com/plenya/api/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("uso: migrate <up|status|version|down|stamp>")
	}
	command := os.Args[1]

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ config: %v", err)
	}

	db, err := sql.Open("pgx", cfg.Database.GetDSN())
	if err != nil {
		log.Fatalf("❌ abrir conexão: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ ping no banco: %v", err)
	}

	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("❌ dialect: %v", err)
	}

	switch command {
	case "up":
		if err = adoptIfLegacy(db); err != nil {
			log.Fatalf("❌ adoção do baseline: %v", err)
		}
		err = goose.Up(db, ".")
	case "status":
		err = goose.Status(db, ".")
	case "version":
		err = goose.Version(db, ".")
	case "down":
		err = goose.Down(db, ".")
	case "stamp":
		err = stampBaseline(db)
	default:
		log.Fatalf("❌ comando desconhecido: %q", command)
	}
	if err != nil {
		log.Fatalf("❌ migrate %s: %v", command, err)
	}
	log.Printf("✅ migrate %s concluído", command)
}

// adoptIfLegacy torna o `up` seguro tanto em banco vazio quanto em banco
// pré-existente: se o schema já existe (tabela `patients` presente) mas o goose
// ainda não foi adotado (sem `goose_db_version`), marca a baseline como aplicada
// antes do `up`. Em banco vazio não faz nada e o `up` roda a baseline normalmente.
func adoptIfLegacy(db *sql.DB) error {
	var gooseTable, patientsTable *string
	if err := db.QueryRow("SELECT to_regclass('public.goose_db_version')::text").Scan(&gooseTable); err != nil {
		return err
	}
	if gooseTable != nil {
		return nil // goose já adotado
	}
	if err := db.QueryRow("SELECT to_regclass('public.patients')::text").Scan(&patientsTable); err != nil {
		return err
	}
	if patientsTable == nil {
		return nil // banco vazio: deixa o `up` rodar a baseline
	}
	log.Println("schema pré-existente detectado sem goose; adotando baseline (stamp)")
	return stampBaseline(db)
}

// stampBaseline registra a migration 1 (baseline) como aplicada SEM rodá-la.
// Necessário ao adotar o goose num banco que já tem o schema (construído
// historicamente pelo AutoMigrate), evitando que o `up` tente recriar tudo.
func stampBaseline(db *sql.DB) error {
	cur, err := goose.EnsureDBVersion(db) // cria goose_db_version + versão 0
	if err != nil {
		return fmt.Errorf("garantir tabela de versão: %w", err)
	}
	if cur >= 1 {
		log.Printf("baseline já registrada (versão atual %d), nada a fazer", cur)
		return nil
	}
	if _, err := db.Exec(
		"INSERT INTO goose_db_version (version_id, is_applied) VALUES (1, true)",
	); err != nil {
		return fmt.Errorf("registrar baseline: %w", err)
	}
	log.Println("baseline (versão 1) marcada como aplicada")
	return nil
}
