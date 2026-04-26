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

	// Calendar V1 — migração de dados ANTES do AutoMigrate.
	//
	// O enum AppointmentType mudou:
	//   antigo: routine | follow_up | urgent | emergency
	//   novo:   initial_assessment | follow_up | telemedicine | procedure | results_review
	//
	// AutoMigrate vai tentar reaplicar o CHECK constraint de Type — se houver
	// rows com valores antigos, o ALTER TABLE quebra. Por isso reclassificamos
	// rows residuais como 'follow_up' (mapping conservador) ANTES.
	//
	// Idempotente: roda 2x sem efeito (rows já mapeadas não satisfazem o WHERE).
	// Pre-empta também o caso de o CHECK antigo ainda estar lá: drop-if-exists
	// remove o CHECK antigo pra que o UPDATE possa rodar mesmo se o nome do
	// CHECK constraint variar entre ambientes.
	preMigrateStmts := []string{
		// Drop CHECK antigo (nome auto-gerado pode variar — usamos pattern match).
		// chk_appointments_type é o nome típico que GORM gera pra CHECK em coluna.
		`DO $$ DECLARE r record; BEGIN
			FOR r IN
				SELECT conname FROM pg_constraint
				WHERE conrelid = 'appointments'::regclass
				  AND contype = 'c'
				  AND pg_get_constraintdef(oid) LIKE '%routine%'
			LOOP
				EXECUTE 'ALTER TABLE appointments DROP CONSTRAINT ' || quote_ident(r.conname);
			END LOOP;
		EXCEPTION WHEN undefined_table THEN
			-- tabela ainda não existe (primeira migration); ignora
			NULL;
		END $$`,
		// Reclassifica valores antigos pra follow_up (conservador).
		`DO $$ BEGIN
			UPDATE appointments
			   SET type = 'follow_up'
			 WHERE type IN ('routine','urgent','emergency');
		EXCEPTION WHEN undefined_table THEN
			NULL;
		WHEN undefined_column THEN
			NULL;
		END $$`,
	}
	for _, stmt := range preMigrateStmts {
		if err := DB.Exec(stmt).Error; err != nil {
			return fmt.Errorf("calendar pre-migrate failed: %w", err)
		}
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
		&models.Campaign{},
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

		// Mobile apps
		&models.DeviceToken{},
		&models.NotificationPreferences{},

		// App paciente — execução de treino + check-ins de bem-estar
		&models.WorkoutSession{},
		&models.WorkoutSessionExerciseLog{},
		&models.HealthCheckIn{},

		// Calendar V1 — Google Calendar OAuth, working hours, ausências,
		// recursos polimórficos. Devem vir DEPOIS de User (FK DoctorID).
		&models.CalendarCredential{},
		&models.WorkingHours{},
		&models.DoctorAbsence{},
		&models.AppointmentResource{},

		// Continuum — programa de acompanhamento longitudinal.
		// Templates primeiro (referenciados por items); depois inscrição,
		// items, boxes e revisões do plano integrado.
		&models.ContinuumBoxTemplate{},
		&models.ContinuumTemplate{},
		&models.ContinuumTemplateItem{},
		&models.PatientContinuum{},
		&models.PatientContinuumItem{},
		&models.PatientContinuumBox{},
		&models.IntegratedPlanRevision{},

		// Portal do Paciente (minha.plenyasaude.com.br)
		&models.PatientPortalInvite{},
		&models.PatientMagicLink{},
		&models.TelemedLobbyToken{},
		&models.PatientDocument{},
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

		// Calendar V1 — anti double-booking via Postgres EXCLUDE constraint.
		//
		// Por quê EXCLUDE e não UNIQUE: appointments têm range temporal
		// (scheduled_at + duration_minutes), não pontual. UNIQUE não suporta
		// "no overlap of intervals". EXCLUDE USING gist + tstzrange resolve
		// nativamente — Postgres rejeita com erro 23P01 (exclusion_violation)
		// na transação, garantindo que mesmo race condition entre 2 secretárias
		// criando simultaneamente é bloqueada (1 commita, outro 409).
		//
		// btree_gist: extensão necessária pra usar `=` (btree) junto com `&&`
		// (gist) na mesma constraint. Disponível por padrão em pgvector image.
		//
		// WHERE clause: só consideramos appointments ATIVOS — cancelled/no_show
		// e soft-deleted não bloqueiam novo agendamento no mesmo horário.
		`CREATE EXTENSION IF NOT EXISTS btree_gist`,
		// Notas IMPORTANTES sobre essa constraint:
		//
		// 1. Usamos `tsrange` (timestamp without time zone) em vez de `tstzrange`
		//    porque a coluna `scheduled_at` é declarada como `timestamp` (sem TZ).
		//    `tstzrange()` exige conversão implícita timestamp→timestamptz que
		//    depende da sessão (TimeZone) → NÃO é IMMUTABLE → Postgres rejeita
		//    em index expression. `tsrange` opera direto na coluna sem cast.
		//
		// 2. Usamos `duration_minutes * interval '1 minute'` em vez do cast
		//    string→interval `(duration_minutes || ' minutes')::interval`. O
		//    cast também não é IMMUTABLE (depende de DateStyle). Multiplicação
		//    por interval literal é IMMUTABLE.
		//
		// 3. Application layer DEVE armazenar scheduled_at sempre em UTC pra
		//    consistência. O ranges são comparados timestamp-a-timestamp,
		//    então qualquer drift de TZ entre escrita e leitura cria bugs.
		//    GORM config `NowFunc: time.Now().UTC()` já garante isso.
		`DO $$ BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'appointments_no_overlap'
			) THEN
				ALTER TABLE appointments ADD CONSTRAINT appointments_no_overlap
				EXCLUDE USING gist (
					doctor_id WITH =,
					tsrange(scheduled_at, scheduled_at + (duration_minutes * interval '1 minute')) WITH &&
				) WHERE (status NOT IN ('cancelled', 'no_show') AND deleted_at IS NULL);
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
