package scheduler

import (
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientAgeJob é responsável por atualizar a idade de todos os pacientes
type PatientAgeJob struct {
	db *gorm.DB
}

// NewPatientAgeJob cria uma nova instância do job
func NewPatientAgeJob(db *gorm.DB) *PatientAgeJob {
	return &PatientAgeJob{db: db}
}

// Run executa o job de atualização de idade
func (j *PatientAgeJob) Run() error {
	log.Println("🎂 Starting daily patient age update job...")

	var patients []models.Patient

	// Buscar todos os pacientes (incluindo soft deleted = false)
	if err := j.db.Find(&patients).Error; err != nil {
		log.Printf("❌ Failed to fetch patients: %v", err)
		return err
	}

	successCount := 0
	errorCount := 0

	// Atualizar idade de cada paciente
	for i := range patients {
		patient := &patients[i]

		// Recalcular idade
		patient.CalculateAge()

		// Salvar (sem trigger de hooks desnecessários)
		if err := j.db.Model(patient).
			Select("age", "age_text", "updated_at").
			Updates(map[string]interface{}{
				"age":      patient.Age,
				"age_text": patient.AgeText,
			}).Error; err != nil {
			log.Printf("❌ Failed to update patient %s age: %v", patient.ID, err)
			errorCount++
			continue
		}

		successCount++
	}

	log.Printf("✅ Patient age update completed: %d success, %d errors", successCount, errorCount)
	return nil
}

// Start inicia o scheduler que roda diariamente às 00:00
func (j *PatientAgeJob) Start() {
	// Calcula o tempo até a próxima execução (meia-noite)
	now := time.Now()
	nextRun := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	durationUntilNextRun := nextRun.Sub(now)

	log.Printf("⏰ Patient age job scheduled to run at %s (in %s)", nextRun.Format("2006-01-02 15:04:05"), durationUntilNextRun)

	// Primeira execução
	time.AfterFunc(durationUntilNextRun, func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  Patient age job error: %v", err)
		}

		// Execuções subsequentes (24 horas)
		ticker := time.NewTicker(24 * time.Hour)
		go func() {
			for range ticker.C {
				if err := j.Run(); err != nil {
					log.Printf("⚠️  Patient age job error: %v", err)
				}
			}
		}()
	})
}
