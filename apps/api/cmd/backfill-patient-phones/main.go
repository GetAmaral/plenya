// backfill-patient-phones — re-salva todos os Patients pra acionar Patient.BeforeSave
// hook (NormalizePhoneBR). Necessário após Phase 2 do CRM, que adicionou normalização
// E.164 BR ao hook. Pacientes antigos ficaram com phone em formatos não-normalizados
// e não são encontrados pelo lookup de WhatsApp inbound.
//
// Uso (dentro do container api):
//
//	docker compose exec api go run cmd/backfill-patient-phones/main.go
//
// Idempotente: já-normalizados são re-escritos com mesmo valor.
package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	var patients []models.Patient
	if err := db.Where("phone IS NOT NULL AND phone != ''").Find(&patients).Error; err != nil {
		log.Fatalf("load patients: %v", err)
	}
	fmt.Printf("📋 %d patients com phone preenchido\n", len(patients))

	var normalized, unchanged, skipped int
	for i := range patients {
		p := &patients[i]
		original := *p.Phone
		// O Save() aciona BeforeSave que normaliza. Mas BeforeSave também valida
		// outros campos — se algum patient tem dado inválido (ex: gender obsoleto)
		// vamos pular sem falhar o batch.
		if err := db.Save(p).Error; err != nil {
			fmt.Printf("⚠️  patient %s skipped: %v (phone era %q)\n", p.ID, err, original)
			skipped++
			continue
		}
		// Re-fetch pra ver o resultado
		var fresh models.Patient
		if err := db.First(&fresh, "id = ?", p.ID).Error; err != nil {
			fmt.Printf("⚠️  patient %s refetch: %v\n", p.ID, err)
			skipped++
			continue
		}
		if fresh.Phone != nil && *fresh.Phone != original {
			fmt.Printf("✓ %s: %q → %q\n", p.ID, original, *fresh.Phone)
			normalized++
		} else {
			unchanged++
		}
	}

	fmt.Printf("\n✅ Done.\n  Normalizados: %d\n  Inalterados:  %d\n  Pulados:      %d\n",
		normalized, unchanged, skipped)
}
