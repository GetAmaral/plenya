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
	return DB.AutoMigrate(
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

		// Lab Tests & Results
		&models.LabTestDefinition{},
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

		// Articles
		&models.Article{},
	)
}

// Close fecha a conexão com o banco de dados
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
