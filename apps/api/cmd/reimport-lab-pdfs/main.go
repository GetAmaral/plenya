// Command reimport-lab-pdfs — re-importa do ZERO um conjunto de PDFs de laudo para um
// paciente: para cada PDF cria um lote novo + copia o arquivo + enfileira um ProcessingJob.
// O worker do servidor (startProcessingWorker) processa: OCR → IA → cria results → classifica,
// já com extração de laboratório/data de coleta. One-off. Roda no container api de produção.
//
//	./reimport-lab-pdfs <patientID> <src.pdf> [src.pdf...]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
)

const uploadsDir = "/app/uploads/lab-result-batches"

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("uso: reimport-lab-pdfs <patientID> <src.pdf> [src.pdf...]")
	}
	patientID, err := uuid.Parse(args[0])
	if err != nil {
		log.Fatalf("patientID inválido: %v", err)
	}
	srcPaths := args[1:]

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

	for _, src := range srcPaths {
		data, rerr := os.ReadFile(src)
		if rerr != nil {
			log.Printf("❌ %s: ler arquivo: %v", src, rerr)
			continue
		}

		// 1) lote novo (lab/data placeholder — o worker preenche do PDF)
		batch := models.LabResultBatch{
			PatientID:      patientID,
			LaboratoryName: "Importado via PDF",
			CollectionDate: time.Now().UTC(),
			Status:         models.LabResultBatchPending,
		}
		if err := db.Create(&batch).Error; err != nil {
			log.Printf("❌ %s: criar lote: %v", src, err)
			continue
		}

		// 2) copia o PDF p/ um caminho novo (desacopla dos lotes antigos)
		dst := filepath.Join(uploadsDir, fmt.Sprintf("%s_%d.pdf", batch.ID, time.Now().Unix()))
		if err := os.WriteFile(dst, data, 0o644); err != nil {
			log.Printf("❌ %s: copiar p/ %s: %v", src, dst, err)
			continue
		}

		// 3) enfileira o job (o worker do servidor processa)
		step := models.StepUploadComplete
		msg := "Reimportação enfileirada"
		job := models.ProcessingJob{
			LabResultBatchID: batch.ID,
			Type:             models.ProcessingJobTypePDFExtraction,
			PDFPath:          dst,
			Status:           models.ProcessingJobPending,
			ProgressStep:     &step,
			ProgressMessage:  &msg,
			Attempts:         0,
			MaxAttempts:      3,
		}
		if err := db.Create(&job).Error; err != nil {
			log.Printf("❌ %s: criar job: %v", src, err)
			continue
		}

		log.Printf("✅ lote %s enfileirado (job %s, pdf %s)", batch.ID, job.ID, filepath.Base(dst))
	}
	log.Printf("Concluído. O worker do servidor vai processar os jobs pendentes.")
}
