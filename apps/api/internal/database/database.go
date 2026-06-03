package database

import (
	"fmt"
	"log"
	"time"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// cryptoBlindKeyHelper — true se BlindIndexKey foi setado.
func cryptoBlindKeyHelper() bool {
	return crypto.GetBlindIndexKey() != ""
}

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

// AutoMigrate executa as migrations automáticas do GORM.
//
// M9 — Workaround GORM circular FK: User.SelectedPatient → Patient.UserID → User.
// Criamos `users` sozinho primeiro (CREATE TABLE IF NOT EXISTS), sem a relação
// SelectedPatient, pra que AutoMigrate(Patient) (que vem em seguida) não tente
// referenciar tabela inexistente.
//
// TODO M9 — migrar pra Atlas migrations completas (gera SQL declarativo a partir
// dos models, sem AutoMigrate). Em prod, AutoMigrate é arriscado: GORM não dropa
// colunas, não reconcilia FK divergente, e usa REINDEX em alguns cenários. O
// workaround acima é cheiro disso.
//
// O caller (cmd/server/main.go) gate AutoMigrate por env:
//   - DEV  : sempre roda
//   - PROD : roda apenas se MIGRATIONS_AUTO=true (default false)
//
// Esse gate é defesa em profundidade contra rollback acidental de schema.
func AutoMigrate() error {
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

		// TZ-HARDENING — converte TODAS as colunas de timestamp "naive"
		// (timestamp without time zone) para timestamptz em todo o schema.
		//
		// Por quê: colunas naive não carregam fuso. O app sempre gravou/leu em UTC
		// (container TZ=UTC, sessão Postgres TZ=UTC, GORM NowFunc=time.Now().UTC()),
		// mas o tipo naive é frágil: qualquer camada que assuma outro fuso corrompe
		// o instante silenciosamente. Foi a raiz de um bug de off-by-one no
		// agendamento. timestamptz fixa o instante de forma inequívoca.
		//
		// USING ... AT TIME ZONE 'UTC': interpreta o valor naive existente COMO UTC
		// (que é o que ele sempre foi) e produz o timestamptz no mesmo instante —
		// preserva os dados, não desloca nada.
		//
		// Guardado por "scheduled_at ainda é naive": roda uma vez; após converter,
		// vira no-op. Em DB novo (tabela ainda não existe) cai no EXCEPTION e o
		// AutoMigrate cria tudo já como timestamptz (models declaram timestamptz).
		//
		// A EXCLUDE constraint anti-overlap usa tsrange(scheduled_at,...) e bloqueia
		// o ALTER TYPE; é dropada aqui e recriada com tstzrange(scheduled_at,end_at)
		// no bloco de índices (após o AutoMigrate). end_at é coluna materializada
		// nova (ver model Appointment.EndAt) — criada e backfillada aqui.
		`DO $$
		DECLARE r record;
		BEGIN
			IF EXISTS (
				SELECT 1 FROM information_schema.columns
				 WHERE table_name = 'appointments'
				   AND column_name = 'scheduled_at'
				   AND data_type = 'timestamp without time zone'
			) THEN
				ALTER TABLE appointments DROP CONSTRAINT IF EXISTS appointments_no_overlap;

				FOR r IN
					SELECT c.table_name, c.column_name
					  FROM information_schema.columns c
					  JOIN information_schema.tables t
					    ON t.table_schema = c.table_schema
					   AND t.table_name = c.table_name
					 WHERE c.table_schema = 'public'
					   AND c.data_type = 'timestamp without time zone'
					   AND t.table_type = 'BASE TABLE'
				LOOP
					EXECUTE format(
						'ALTER TABLE %I ALTER COLUMN %I TYPE timestamptz USING %I AT TIME ZONE ''UTC''',
						r.table_name, r.column_name, r.column_name
					);
				END LOOP;

				-- end_at materializado p/ a EXCLUDE constraint (tstzrange precisa de
				-- expressão IMMUTABLE; scheduled_at + interval não é sobre timestamptz).
				ALTER TABLE appointments ADD COLUMN IF NOT EXISTS end_at timestamptz;
				UPDATE appointments
				   SET end_at = scheduled_at + (duration_minutes * interval '1 minute')
				 WHERE end_at IS NULL;
				ALTER TABLE appointments ALTER COLUMN end_at SET NOT NULL;
			END IF;
		EXCEPTION WHEN undefined_table THEN
			NULL;
		END $$`,

		// Fluxo de check-in — amplia chk_appointments_status com os status novos
		// checked_in (paciente chegou, aguardando) e in_progress (em atendimento).
		//
		// A tag `check:` do GORM só cria a constraint na criação da tabela; numa
		// tabela existente o conjunto de valores não é atualizado pelo AutoMigrate.
		// Recriamos aqui explicitamente (idempotente: só dropa se o def atual ainda
		// não contempla 'checked_in').
		//
		// A EXCLUDE anti-overlap NÃO muda: ela exclui só cancelled/no_show, então
		// checked_in/in_progress continuam contando como ocupação do horário.
		`DO $$ BEGIN
			IF EXISTS (
				SELECT 1 FROM pg_constraint
				 WHERE conname = 'chk_appointments_status'
				   AND pg_get_constraintdef(oid) NOT LIKE '%checked_in%'
			) THEN
				ALTER TABLE appointments DROP CONSTRAINT chk_appointments_status;
			END IF;
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'chk_appointments_status'
			) THEN
				ALTER TABLE appointments ADD CONSTRAINT chk_appointments_status
				CHECK (status IN ('scheduled','confirmed','checked_in','in_progress','completed','cancelled','no_show'));
			END IF;
		EXCEPTION WHEN undefined_table THEN
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
		&models.RefreshToken{},   // C3 — refresh token rotation + revocation
		&models.MagicLinkToken{}, // magic link single-use do Escore Light (sem FK de user)

		// Anamnesis
		&models.AnamnesisTemplate{},
		&models.AnamnesisTemplateItem{},
		&models.Anamnesis{},
		&models.AnamnesisItem{},

		// Appointments & Prescriptions
		&models.Appointment{},
		&models.ClinicalNote{},
		&models.PatientAllergy{},
		&models.ConsultationVitals{},
		&models.PatientProblem{},
		&models.MedicationInUse{},
		&models.CIDCode{},
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
		&models.WaitlistEntry{},
		&models.AppointmentRecall{},
		&models.ConsultationPrice{},
		&models.PaymentReceiptCounter{},
		&models.AppointmentPayment{},

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

		// Documentos clínicos emitidos/assináveis (P3 frente 2)
		&models.IssuedDocument{},

		// Plano de cuidado AGIR (P3 frente 3)
		&models.CarePlanItem{},

		// Gravação + transcrição da teleconsulta (Daily.co + Deepgram)
		&models.TelemedRecording{},
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
		// 1. Usamos `tstzrange(scheduled_at, end_at)` sobre as colunas timestamptz.
		//    end_at é coluna materializada (Appointment.EndAt, mantida pelo hook
		//    BeforeSave). Construir o fim inline — `scheduled_at + interval` — NÃO
		//    é IMMUTABLE sobre timestamptz (a soma com interval pode cruzar DST),
		//    então o Postgres rejeita em index expression. `tstzrange(timestamptz,
		//    timestamptz)` É IMMUTABLE — por isso materializamos end_at.
		//
		// 2. Application layer grava scheduled_at/end_at sempre em UTC (timestamptz
		//    é inequívoco quanto ao instante). GORM `NowFunc: time.Now().UTC()` e o
		//    parse RFC3339 garantem isso.
		`DO $$ BEGIN
			IF NOT EXISTS (
				SELECT 1 FROM pg_constraint WHERE conname = 'appointments_no_overlap'
			) THEN
				ALTER TABLE appointments ADD CONSTRAINT appointments_no_overlap
				EXCLUDE USING gist (
					doctor_id WITH =,
					tstzrange(scheduled_at, end_at) WITH &&
				) WHERE (status NOT IN ('cancelled', 'no_show') AND deleted_at IS NULL);
			END IF;
		END $$`,
	}
	for _, stmt := range indexStmts {
		if err := DB.Exec(stmt).Error; err != nil {
			return err
		}
	}

	// M4 — backfill CPFBlindIndex em rows existentes (quando vazio).
	// Idempotente: re-runs no-op (WHERE cpf_blind_index = '' filtra pendentes).
	// Só roda se BlindIndexKey configurado (caso contrário, hook BeforeSave
	// silencia e seria inútil). Best-effort: erro não aborta migration.
	if err := backfillCPFBlindIndex(); err != nil {
		fmt.Printf("⚠️  M4 backfill CPF blind index falhou (não-fatal): %v\n", err)
	}
	return nil
}

// backfillCPFBlindIndex — M4 — para cada Patient com cpf não nulo e
// cpf_blind_index vazio, lê (descriptografa via AfterFind), recalcula índice
// via hook BeforeSave (Save chama BeforeSave). Best-effort.
//
// NOTA: usa Save (não Update) pra que BeforeSave dispare. Update pulando
// hooks deixaria índice vazio. Se houver muitos Patients (>10k), considere
// batchear via paginação.
func backfillCPFBlindIndex() error {
	// Pula se chave de blind index não tá setada (hook é no-op).
	// (não temos acesso ao config aqui; checamos via crypto package).
	// Como o hook é tolerante a key-missing, podemos rodar mesmo assim:
	// ele só vai gravar idx="" e o WHERE abaixo continuará retornando rows.
	// Pra evitar loop, checamos a key:
	if cryptoBlindIndexKeySet() == false {
		return nil
	}
	type patientLite struct {
		ID models.Patient
	}
	// Carrega só IDs e CPF cifrado pra processar
	var patients []models.Patient
	if err := DB.Where("cpf IS NOT NULL AND cpf <> '' AND (cpf_blind_index IS NULL OR cpf_blind_index = '')").
		Limit(5000). // safety cap; rerun cobre o resto
		Find(&patients).Error; err != nil {
		return err
	}
	if len(patients) == 0 {
		return nil
	}
	count := 0
	for i := range patients {
		// AfterFind já decifrou CPF; Save chama BeforeSave que recalcula índice.
		if err := DB.Save(&patients[i]).Error; err != nil {
			fmt.Printf("⚠️  backfill CPF idx falhou para patient=%s: %v\n", patients[i].ID, err)
			continue
		}
		count++
	}
	if count > 0 {
		fmt.Printf("✅ M4 backfill: %d patients com CPF blind index calculado\n", count)
	}
	return nil
}

// cryptoBlindIndexKeySet — wrapper local pra evitar import cycle aparente
// (crypto já é importado, mas mantemos a checagem isolada pra documentar).
func cryptoBlindIndexKeySet() bool {
	// Retorna true se a key foi setada (não vazia).
	// Importamos crypto aqui via models? — não, importamos direto.
	return cryptoBlindKeyHelper()
}

// RevokeAuditLogMutations — H8: em produção, revoga UPDATE/DELETE da tabela
// audit_logs no nível do DB (defense in depth além dos hooks GORM).
// Não-fatal: se REVOKE quebrar (role faltando, etc.), só loga warning.
// Em dev, NÃO roda — preserva tooling de cleanup local.
func RevokeAuditLogMutations(env string) {
	if env != "production" {
		return
	}
	// CURRENT_USER funciona porque rodamos sob o role da app (não superuser).
	stmts := []string{
		`REVOKE UPDATE, DELETE, TRUNCATE ON audit_logs FROM CURRENT_USER`,
	}
	for _, s := range stmts {
		if err := DB.Exec(s).Error; err != nil {
			fmt.Printf("⚠️  H8 audit_logs REVOKE falhou (não-fatal): %v\n", err)
		}
	}
}

// Close fecha a conexão com o banco de dados
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
