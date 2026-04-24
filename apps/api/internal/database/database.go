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
	// Workaround GORM circular FK: User.SelectedPatient → Patient.UserID → User
	// Criamos users sozinho primeiro, sem a relação SelectedPatient, pra que
	// AutoMigrate(Patient) (que vem em seguida) não tente referenciar tabela inexistente.
	if err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (id uuid PRIMARY KEY)`).Error; err != nil {
		return err
	}
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
		&models.ConversationRead{},

		// Email ingest (worker IMAP IDLE)
		&models.EmailIngestState{},

		// Notifications (in-app sino) + Subscriptions + Workers — necessários
		// pra serviços que rodam em background (workers/notification_service)
		&models.Notification{},
		&models.SubscriptionPlan{},
		&models.PatientSubscription{},
		&models.ProcessingJob{},
		&models.EmbeddingQueue{},
		&models.ArticleEmbedding{},
		&models.ScoreItemEmbedding{},

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

		// Email ingest dedup: lookup por Message-ID (RFC 5322) em metadata.
		// Usado pra evitar duplicar LeadActivity quando o worker reprocessa o mesmo
		// email (race no resume após restart, ou múltiplas pastas espelhando).
		`CREATE INDEX IF NOT EXISTS idx_lead_activities_email_message_id
		 ON lead_activities ((metadata->>'message_id'))`,

		// Central de Conversas (Bloco A) — LeadActivity pode pertencer a Lead OU Patient.
		// GORM AutoMigrate detecta o ponteiro *uuid.UUID em LeadID e gera coluna nullable,
		// mas migrações de tabela existente que tinham NOT NULL não são revertidas
		// automaticamente — precisamos do ALTER explícito (idempotente).
		`ALTER TABLE lead_activities ALTER COLUMN lead_id DROP NOT NULL`,

		// CHECK constraint: exatamente um de lead_id|patient_id setado.
		// `ADD CONSTRAINT IF NOT EXISTS` foi adicionado no PG 9.6+ pra constraints
		// nomeadas — DO block garante idempotência cross-version.
		`DO $$ BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'lead_activities_owner_check'
			) THEN
				ALTER TABLE lead_activities ADD CONSTRAINT lead_activities_owner_check
				CHECK (
					(lead_id IS NOT NULL AND patient_id IS NULL) OR
					(lead_id IS NULL AND patient_id IS NOT NULL)
				);
			END IF;
		END $$`,

		// Composite index pra Central de Conversas listar histórico por Patient
		// filtrando por tipo (mirror do existente pra Lead).
		`CREATE INDEX IF NOT EXISTS idx_lead_activities_patient_type_created
		 ON lead_activities (patient_id, type, created_at)`,

		// Lookup case-insensitive por email em Patient (worker IMAP).
		// Substitui o índice plano gerado por GORM (que não cobre LOWER()).
		`CREATE INDEX IF NOT EXISTS idx_patients_email_lower
		 ON patients (LOWER(email))`,

		// Migração de dados (idempotente):
		// Pra cada Lead com ConvertedPatientID setado, move suas activities pro Patient.
		// Skipa as que já têm patient_id (re-runs no-op).
		// Decisão: paciente convertido só deve ter activities como Patient — Lead vira
		// "registro histórico de captura" sem conversa anexada.
		`UPDATE lead_activities la
		 SET patient_id = l.converted_patient_id, lead_id = NULL
		 FROM leads l
		 WHERE la.lead_id = l.id
		   AND l.converted_patient_id IS NOT NULL
		   AND la.patient_id IS NULL`,

		// Central de Conversas (Bloco B) — ConversationRead trackeia LastReadAt por
		// (user, conversa) pra calcular unread_count na lista de conversas.
		// CHECK constraint garante que owner_type é "lead" ou "patient".
		`DO $$ BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'conversation_reads_owner_type_check'
			) THEN
				ALTER TABLE conversation_reads ADD CONSTRAINT conversation_reads_owner_type_check
				CHECK (owner_type IN ('lead','patient'));
			END IF;
		END $$`,
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
