// Command reprocess-lab-batch — recria os LabResults de um ou mais lotes a partir do
// PDFContentJSON já armazenado (sem refazer OCR/IA) e re-classifica. One-off de
// recuperação/reprocessamento. Roda dentro do container api (acesso ao DB + chaves).
//
//	./reprocess-lab-batch <batchID> [batchID...]
package main

import (
	"flag"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/services"
)

func main() {
	flag.Parse()
	ids := flag.Args()
	if len(ids) == 0 {
		log.Fatal("uso: reprocess-lab-batch <batchID> [batchID...]")
	}

	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	_ = crypto.Init(cfg.Security.EncryptionKey)
	crypto.SetBlindIndexKey(cfg.Security.BlindIndexKey)

	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	batchSvc := services.NewLabResultBatchService(db)
	pjs := services.NewProcessingJobService(db, nil, nil, nil, batchSvc)

	for _, idStr := range ids {
		id, perr := uuid.Parse(idStr)
		if perr != nil {
			log.Printf("❌ id inválido %s: %v", idStr, perr)
			continue
		}
		m, u, rerr := pjs.ReprocessResultsFromJSON(id)
		if rerr != nil {
			log.Printf("❌ %s: %v", idStr, rerr)
			continue
		}
		log.Printf("✅ %s reprocessado: %d casados + %d não-casados", idStr, m, u)
	}
}
