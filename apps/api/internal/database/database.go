package database

import (
	"fmt"
	"log"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB é a instância global do banco de dados
var DB *gorm.DB

// Connect estabelece conexão com PostgreSQL
func Connect(cfg *config.Config) error {
	var err error

	// Configuração do logger GORM
	gormLogger := logger.Default
	if cfg.Server.Environment == "production" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	// Conectar ao PostgreSQL
	DB, err = gorm.Open(postgres.Open(cfg.Database.GetDSN()), &gorm.Config{
		Logger: gormLogger,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configurar pool de conexões
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Database connected successfully")

	return nil
}

// AutoMigrate executa as migrations automáticas do GORM
// NOTA: Em produção, usaremos Atlas para migrations
func AutoMigrate() error {
	if err := DB.AutoMigrate(
		// Core
		&models.User{},
		&models.Patient{},
		&models.AuditLog{},

		// Anamnesis
		&models.AnamnesisTemplate{},
		&models.AnamnesisTemplateItem{},
		&models.Anamnesis{},
		&models.AnamnesisItem{},

		// Appointments & Prescriptions
		&models.Appointment{},
		&models.Prescription{},
		&models.PrescriptionMedication{},

		// Lab Tests & Results
		&models.LabTestDefinition{},
		&models.LabTestUnitConversion{},
		&models.LabRequest{},
		&models.LabRequestTemplate{},
		&models.LabResult{},
		&models.LabResultBatch{},
		&models.LabResultValue{},
		&models.LabResultView{},
		&models.LabResultViewItem{},

		// Scores
		&models.ScoreGroup{},
		&models.ScoreSubgroup{},
		&models.ScoreItem{},
		&models.ScoreLevel{},

		// Escore Light (público anônimo)
		&models.AnonymousScoreSession{},
		&models.AnonymousScoreItem{},
		&models.AnonymousScoreSnapshot{},
		&models.AnonymousScoreGroupResult{},

		// CRM — Leads
		&models.Lead{},
		&models.LeadActivity{},

		// Articles
		&models.Article{},

		// Training
		&models.Exercise{},
		&models.WorkoutPlan{},
		&models.WorkoutPlanSession{},
		&models.WorkoutSessionExercise{},
		&models.PhysicalAssessment{},
		&models.WorkoutPeriodization{},
		&models.WorkoutMesocycle{},
		&models.FitnessTestResult{},
		&models.PosturalAssessment{},
	); err != nil {
		return err
	}

	// Índices em expressions (não suportados via tag GORM):
	//
	// (Phase 2 do CRM)
	//
	// 1. lead_activities.metadata->>'wa_message_id' — usado em RecordWhatsAppStatus pra
	//    correlacionar status updates da Meta com a mensagem outbound original. Sem
	//    índice, cada webhook status faz seq scan (3× por mensagem).
	//
	// 2. (lead_id, type, created_at) — usado em Stats() JOIN LATERAL pra calcular
	//    "tempo médio até primeiro contato" (MIN message_sent created_at por lead).
	//
	// CREATE INDEX IF NOT EXISTS é idempotente — pode rodar a cada boot sem custo.
	indexStmts := []string{
		`CREATE INDEX IF NOT EXISTS idx_lead_activities_wa_message_id
		 ON lead_activities ((metadata->>'wa_message_id'))`,
		`CREATE INDEX IF NOT EXISTS idx_lead_activities_lead_type_created
		 ON lead_activities (lead_id, type, created_at)`,
	}
	for _, stmt := range indexStmts {
		if err := DB.Exec(stmt).Error; err != nil {
			return err
		}
	}
	return nil
}

// Close fecha a conexão com o banco de dados
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
