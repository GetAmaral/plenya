package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/handlers"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/scheduler"
	"github.com/plenya/api/internal/services"
	"github.com/plenya/api/internal/workers"
	"gorm.io/gorm"
)

// @title Plenya EMR API
// @version 1.0
// @description Sistema de Prontuário Médico Eletrônico
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email suporte@plenya.com.br

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3001
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite 'Bearer' seguido do token JWT

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	// M6 — customErrorHandler precisa do env pra decidir entre msg detalhada/opaca.
	appConfig = cfg

	// Inicializar criptografia
	if err := crypto.Init(cfg.Security.EncryptionKey); err != nil {
		log.Fatalf("❌ Failed to initialize crypto: %v", err)
	}
	// M4 — blind index pra CPF (HMAC-SHA256). Chave separada da encryption key.
	crypto.SetBlindIndexKey(cfg.Security.BlindIndexKey)
	log.Println("🔐 Encryption initialized")

	// Conectar ao banco de dados
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Schema é gerenciado por migrations goose (cmd/migrate), aplicadas como passo
	// do deploy (dev: entrypoint.sh; prod: prod-entrypoint.sh) — NUNCA no boot da app.
	// GORM AutoMigrate fica como escape hatch manual de prototipagem: só roda quando
	// MIGRATIONS_AUTO=true é setado explicitamente. Default = desligado (dev e prod).
	// Ver docs/emr/migrations-decisao.md.
	migrationsAutoEnv := os.Getenv("MIGRATIONS_AUTO")
	shouldRun := false
	switch migrationsAutoEnv {
	case "true", "1", "yes":
		shouldRun = true
	default:
		shouldRun = false
	}
	if shouldRun {
		log.Println("🔄 AutoMigrate manual (MIGRATIONS_AUTO=true) — escape hatch de dev...")
		if err := database.AutoMigrate(); err != nil {
			log.Fatalf("❌ Failed to auto migrate: %v", err)
		}
		log.Println("✅ AutoMigrate concluído")
	} else {
		log.Println("⏭️  AutoMigrate desligado — schema via migrations goose (cmd/migrate)")
	}

	// H8 — em produção, revoga UPDATE/DELETE de audit_logs no DB (não-fatal).
	database.RevokeAuditLogMutations(cfg.Server.Environment)

	// DEV BYPASS AUTH: Carregar admin user após DB init.
	// (config.Validate já bloqueia bypass=true em production — defense in depth.)
	if cfg.Dev.BypassAuth {
		// ANSI yellow + bold pra warning chamativo no terminal.
		log.Println("\x1b[1;33m" + strings.Repeat("=", 70) + "\x1b[0m")
		log.Println("\x1b[1;33m⚠️  DEV_BYPASS_AUTH HABILITADO — AUTENTICAÇÃO DESATIVADA\x1b[0m")
		log.Println("\x1b[1;33m   Esta flag NUNCA deve estar em produção. config.Validate aborta\x1b[0m")
		log.Println("\x1b[1;33m   o boot se ENVIRONMENT=production. Use apenas em dev local.\x1b[0m")
		log.Println("\x1b[1;33m" + strings.Repeat("=", 70) + "\x1b[0m")
		var adminUser models.User
		if err := database.DB.Where("email = ?", "admin@plenya.com").First(&adminUser).Error; err != nil {
			log.Fatalf("DEV_BYPASS_AUTH requer usuário admin@plenya.com no banco: %v", err)
		}
		cfg.Dev.AdminUserID = adminUser.ID
		cfg.Dev.AdminEmail = adminUser.Email
		cfg.Dev.AdminRoles = adminUser.GetRoles()
		log.Printf("\x1b[1;33m⚠️  DEV bypass ativo como: %s (%s)\x1b[0m", adminUser.Email, adminUser.ID)
	}

	// Initialize processing services early (needed by background worker)
	labTestDefRepo := repository.NewLabTestDefinitionRepository(database.DB)
	labTestDefService := services.NewLabTestDefinitionService(labTestDefRepo)
	labResultBatchService := services.NewLabResultBatchService(database.DB)
	ocrService := services.NewOCRService()
	aiService := services.NewAIService(cfg)
	processingJobService := services.NewProcessingJobService(
		database.DB,
		ocrService,
		aiService,
		labTestDefService,
		labResultBatchService,
	)

	// Initialize RAG embedding services (Phase 3)
	embeddingService := services.NewEmbeddingService(cfg, database.DB)
	chunkingService := services.NewChunkingService()
	embeddingQueueService := services.NewEmbeddingQueueService(database.DB)
	embeddingWorker := workers.NewEmbeddingWorker(database.DB, embeddingService, chunkingService)

	// Initialize RAG semantic search services (Phase 4)
	vectorRepo := repository.NewArticleVectorRepository(database.DB)
	semanticService := services.NewArticleSemanticService(vectorRepo, embeddingService)

	// Criar app Fiber
	app := fiber.New(fiber.Config{
		AppName:      "Plenya EMR API",
		ServerHeader: "Plenya",
		ErrorHandler: customErrorHandler,
		BodyLimit:    200 * 1024 * 1024, // 200MB para permitir upload de livros grandes
		// CRITICAL C5 — Trusted proxy enforcement.
		// Sem isso, qualquer request pode setar X-Forwarded-For e burlar
		// rate limit / audit log de IP. Com TrustedProxies setado, Fiber
		// só lê o header se o conexão chegou de IP da lista.
		EnableTrustedProxyCheck: true,
		TrustedProxies:          cfg.Server.TrustedProxies,
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	// Middlewares globais
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	// M1 — multi-origin CORS. Lista vinda de CORS_ORIGINS (CSV).
	// AllowOriginsFunc valida case-insensitive e exato (sem wildcard).
	// AllowCredentials=false (cookies não cross-domain — JWT vai em header).
	allowedOrigins := make(map[string]bool, len(cfg.Server.CORSOrigins))
	for _, o := range cfg.Server.CORSOrigins {
		o = strings.TrimSpace(o)
		if o != "" {
			allowedOrigins[strings.ToLower(o)] = true
		}
	}
	// CORS_ORIGIN (singular) também aceita CSV — splita por vírgula em vez de
	// tratar a string inteira como uma origin (corrige allowlist em dev).
	if cfg.Server.CORSOrigin != "" {
		for _, o := range strings.Split(cfg.Server.CORSOrigin, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				allowedOrigins[strings.ToLower(o)] = true
			}
		}
	}
	log.Printf("🌐 CORS allowed origins: %v", keysOf(allowedOrigins))
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return allowedOrigins[strings.ToLower(origin)]
		},
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	// HIGH H1 — REMOVIDO app.Static("/uploads", "/app/uploads").
	// Servir uploads sem auth vazava PDFs de prescrição/atestado/exames.
	// Cada recurso agora usa handler autenticado (DownloadDocument, etc.)
	// que valida permissão antes de servir o arquivo.
	// Para casos públicos validáveis (QR code de prescrição), criar rota
	// /public/* específica com token assinado de curta duração.

	// Rotas
	setupRoutes(app, cfg, processingJobService, labTestDefService, labResultBatchService, embeddingQueueService, semanticService, aiService)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("🛑 Shutting down server...")
		_ = app.Shutdown()
	}()

	// Iniciar worker de processamento (background)
	go startProcessingWorker(processingJobService)

	// Iniciar embedding worker (RAG background processing)
	go startEmbeddingWorker(embeddingWorker)

	// Iniciar scheduler de atualização de idade dos pacientes
	patientAgeJob := scheduler.NewPatientAgeJob(database.DB)
	patientAgeJob.Start()

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s (environment: %s)", cfg.Server.Port, cfg.Server.Environment)
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// setupRoutes configura as rotas da API
func setupRoutes(
	app *fiber.App,
	cfg *config.Config,
	processingJobService *services.ProcessingJobService,
	labTestDefService *services.LabTestDefinitionService,
	labResultBatchService *services.LabResultBatchService,
	embeddingQueueService *services.EmbeddingQueueService,
	semanticService *services.ArticleSemanticService,
	aiService *services.AIService,
) {
	// Health check
	app.Get("/health", healthCheck)

	// RAG/Embeddings health check
	app.Get("/health/rag", func(c *fiber.Ctx) error {
		return ragHealthCheck(c, database.DB)
	})

	// Inicializar repositories
	scoreRepo := repository.NewScoreRepository(database.DB)
	methodRepo := repository.NewMethodRepository(database.DB)
	labResultValueRepo := repository.NewLabResultValueRepository(database.DB)
	labRequestRepo := repository.NewLabRequestRepository(database.DB)
	labRequestTemplateRepo := repository.NewLabRequestTemplateRepository(database.DB)
	anamnesisTemplateRepo := repository.NewAnamnesisTemplateRepository(database.DB)
	labResultViewRepo := repository.NewLabResultViewRepository(database.DB)
	scoreSnapshotRepo := repository.NewScoreSnapshotRepository(database.DB)
	labResultRepo := repository.NewLabResultRepository(database.DB)
	anamnesisRepo := repository.NewAnamnesisRepository(database.DB)
	subscriptionRepo := repository.NewSubscriptionRepository(database.DB)
	subscriptionPlanRepo := repository.NewSubscriptionPlanRepository(database.DB)
	notificationRepo := repository.NewNotificationRepository(database.DB)

	// Inicializar services
	authService := services.NewAuthService(database.DB, cfg)
	patientService := services.NewPatientService(database.DB)
	anamnesisService := services.NewAnamnesisService(database.DB)
	clinicalNoteService := services.NewClinicalNoteService(database.DB)
	allergyService := services.NewAllergyService(database.DB)
	vitalsService := services.NewVitalsService(database.DB)
	problemService := services.NewProblemService(database.DB)
	medicationInUseService := services.NewMedicationInUseService(database.DB)
	cidService := services.NewCIDService(database.DB)
	anamnesisTemplateService := services.NewAnamnesisTemplateService(anamnesisTemplateRepo)
	// Calendar V1 (Bloco E) — AppointmentService depende de google/daily/notif.
	googleCalendarService := services.NewGoogleCalendarService(cfg, database.DB)
	dailyCoService := services.NewDailyCoService(cfg)
	emailServiceForAppt := services.NewEmailService(cfg).WithDB(database.DB)
	whatsappServiceForAppt := services.NewWhatsAppService(cfg)
	appointmentNotificationService := services.NewAppointmentNotificationService(
		database.DB, emailServiceForAppt, whatsappServiceForAppt, cfg,
	)
	appointmentService := services.NewAppointmentService(
		database.DB, googleCalendarService, dailyCoService, appointmentNotificationService, cfg,
	)
	waitlistService := services.NewWaitlistService(database.DB)
	recallService := services.NewRecallService(database.DB)
	paymentService := services.NewPaymentService(database.DB)
	prescriptionService := services.NewPrescriptionService(database.DB)
	labResultService := services.NewLabResultService(database.DB)
	scoreService := services.NewScoreService(scoreRepo)
	methodService := services.NewMethodService(methodRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	subscriptionPlanService := services.NewSubscriptionPlanService(subscriptionPlanRepo)
	notificationService := services.NewNotificationService(notificationRepo, subscriptionRepo)
	scoreSnapshotService := services.NewScoreSnapshotService(scoreSnapshotRepo, scoreRepo, labResultRepo, anamnesisRepo, database.DB)
	emailService := services.NewEmailService(cfg).WithDB(database.DB)
	whatsappService := services.NewWhatsAppService(cfg)
	leadService := services.NewLeadService(database.DB, notificationService, whatsappService, emailService, cfg)
	// Email ingest worker (IMAP IDLE → LeadActivity). No-op se MAIL_INGEST_ENABLED=false.
	notificationEmailService := services.NewNotificationEmailService(database.DB)
	emailIngestService := services.NewEmailIngestService(database.DB, cfg, leadService, notificationEmailService)
	emailIngestService.Start(context.Background())
	anonymousScoreService := services.NewAnonymousScoreService(database.DB, scoreRepo, cfg, emailService, whatsappService, leadService)
	labResultValueService := services.NewLabResultValueService(labResultValueRepo)
	labRequestService := services.NewLabRequestService(labRequestRepo, database.DB)
	labRequestTemplateService := services.NewLabRequestTemplateService(labRequestTemplateRepo)
	labResultViewService := services.NewLabResultViewService(labResultViewRepo)
	articleService := services.NewArticleService(database.DB, "./uploads/articles", embeddingQueueService, aiService)
	userService := services.NewUserService(database.DB)
	profileService := services.NewProfileService(database.DB)
	medicationDefinitionService := services.NewMedicationDefinitionService(database.DB)

	// Digital prescription services (Phase 4)
	// Assinatura em nuvem (e-CPF em nuvem / PSC) — gated off por default (ICP_CLOUD_ENABLED).
	cloudSignatureProvider := services.NewCloudSignatureProvider(services.CloudSignatureConfig{
		Enabled:  cfg.Signature.CloudEnabled,
		Provider: cfg.Signature.CloudProvider,
		BaseURL:  cfg.Signature.CloudBaseURL,
		APIKey:   cfg.Signature.CloudAPIKey,
	})
	certificateService := services.NewCertificateService(database.DB, cfg.Security.EncryptionKey, cloudSignatureProvider)
	signatureService := services.NewSignatureService(certificateService, services.SignatureOptions{
		TSAURL:      cfg.Signature.TSAURL,
		TSAUsername: cfg.Signature.TSAUsername,
		TSAPassword: cfg.Signature.TSAPassword,
	})
	sncrConfig := services.SNCRConfig{
		Enabled:        cfg.SNCR.Enabled,
		ProductionMode: cfg.SNCR.ProductionMode,
		APIURL:         cfg.SNCR.APIURL,
		APIKey:         cfg.SNCR.APIKey,
	}
	sncrService := services.NewSNCRService(sncrConfig, database.DB)
	patientDocumentsService := services.NewPatientDocumentsService(database.DB, "/app/uploads")
	prescriptionPDFService := services.NewPrescriptionPDFService(
		database.DB,
		signatureService,
		sncrService,
		patientDocumentsService,
		"/app/uploads",
	)
	// TODO: labRequestPDFService will be integrated later for signed lab requests

	// Inicializar validator
	validate := validator.New()

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(authService)
	oauthService := services.NewOAuthService(database.DB, cfg, authService)
	oauthHandler := handlers.NewOAuthHandler(oauthService)
	patientHandler := handlers.NewPatientHandler(patientService)
	anamnesisHandler := handlers.NewAnamnesisHandler(anamnesisService)
	clinicalNoteHandler := handlers.NewClinicalNoteHandler(clinicalNoteService)
	allergyHandler := handlers.NewAllergyHandler(allergyService)
	vitalsHandler := handlers.NewVitalsHandler(vitalsService)
	problemHandler := handlers.NewProblemHandler(problemService)
	medicationInUseHandler := handlers.NewMedicationInUseHandler(medicationInUseService)
	cidHandler := handlers.NewCIDHandler(cidService)
	anamnesisTemplateHandler := handlers.NewAnamnesisTemplateHandler(anamnesisTemplateService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	waitlistHandler := handlers.NewWaitlistHandler(waitlistService)
	recallHandler := handlers.NewRecallHandler(recallService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService, prescriptionPDFService)
	labResultHandler := handlers.NewLabResultHandler(labResultService)
	labResultBatchHandler := handlers.NewLabResultBatchHandler(labResultBatchService, processingJobService)
	processingJobHandler := handlers.NewProcessingJobHandler(processingJobService)
	scoreHandler := handlers.NewScoreHandler(scoreService, validate)
	methodHandler := handlers.NewMethodHandler(methodService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	subscriptionPlanHandler := handlers.NewSubscriptionPlanHandler(subscriptionPlanService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	scoreSnapshotHandler := handlers.NewScoreSnapshotHandler(scoreSnapshotService)
	anonymousLabPDFService := services.NewAnonymousLabPDFService(database.DB, services.NewOCRService(), services.NewPDFTextCleaner(), aiService)
	anonymousScoreHandler := handlers.NewAnonymousScoreHandler(anonymousScoreService, authService, anonymousLabPDFService)
	labTestDefHandler := handlers.NewLabTestDefinitionHandler(labTestDefService)
	labResultValueHandler := handlers.NewLabResultValueHandler(labResultValueService)
	labRequestHandler := handlers.NewLabRequestHandler(labRequestService, certificateService)
	labRequestTemplateHandler := handlers.NewLabRequestTemplateHandler(labRequestTemplateService)
	labResultViewHandler := handlers.NewLabResultViewHandler(labResultViewService)
	articleHandler := handlers.NewArticleHandler(articleService)
	articleSemanticHandler := handlers.NewArticleSemanticHandler(semanticService, validate)
	userHandler := handlers.NewUserHandler(userService)
	profileHandler := handlers.NewProfileHandler(profileService)
	certificateHandler := handlers.NewCertificateHandler(database.DB, certificateService)
	medicationDefHandler := handlers.NewMedicationDefinitionHandler(medicationDefinitionService)
	enrichmentPrepHandler := handlers.NewScoreEnrichmentPreparationHandler(database.DB)
	// M10 — Turnstile CAPTCHA pra rota pública /leads. Em dev (Secret vazio),
	// validação é pulada com warning (não quebra fluxo local).
	turnstileSvc := services.NewTurnstileService(cfg.Turnstile.Secret)
	if !turnstileSvc.Enabled() {
		log.Println("⚠️  TURNSTILE_SECRET vazio — CAPTCHA desabilitado (ok em dev)")
	}
	leadHandler := handlers.NewLeadHandler(leadService, emailService).WithTurnstile(turnstileSvc)
	campaignService := services.NewCampaignService(database.DB, cfg)
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	whatsappWebhookHandler := handlers.NewWhatsAppWebhookHandler(cfg, whatsappService, leadService)
	conversationService := services.NewConversationService(database.DB, leadService, emailService, whatsappService, notificationService, aiService)
	conversationAutomationService := services.NewConversationAutomationService(database.DB, cfg)
	conversationHandler := handlers.NewConversationHandler(conversationService, conversationAutomationService)
	notificationEmailHandler := handlers.NewNotificationEmailHandler(notificationEmailService)
	conversationAttachmentHandler := handlers.NewConversationAttachmentHandler("/app/uploads")

	// Portal do Paciente (minha.plenyasaude.com.br) — auth + dashboard + appointments.
	// Continuum é injetado mais tarde via WithContinuum (depende de patientContinuumService).
	patientPortalService := services.NewPatientPortalService(database.DB, cfg, authService, emailService, whatsappService)
	patientDashboardService := services.NewPatientDashboardService(database.DB)
	patientAppointmentService := services.NewPatientAppointmentService(database.DB, notificationService)
	patientMessagesService := services.NewPatientMessagesService(database.DB, notificationService)
	patientLabsService := services.NewPatientLabsService(database.DB)
	patientScoresService := services.NewPatientScoresService(database.DB)
	patientProfileService := services.NewPatientProfileService(database.DB)
	// patientDocumentsService criado acima (junto com os serviços de prescrição).
	whatsappMediaService := services.NewWhatsAppMediaService("/app/uploads")
	// WhatsApp Fase 2: roteia mídia inbound (arquivo de paciente → prontuário;
	// áudio/lead → bucket cifrado da conversa).
	leadService.SetMediaServices(patientDocumentsService, whatsappMediaService)
	leadService.SetTranscriber(services.NewTranscriptionService(cfg)) // 2.2: transcrição de áudio (Whisper)
	conversationService.SetMediaServices(whatsappMediaService, patientDocumentsService)
	// 2.0b: interpretar documento de prontuário como exame (reusa pipeline OCR→IA).
	conversationService.SetExamInterpreter(labResultBatchService, processingJobService)
	patientWorkoutsService := services.NewPatientWorkoutsService(database.DB)
	patientCheckInsService := services.NewPatientCheckInsService(database.DB)
	notificationPreferencesService := services.NewNotificationPreferencesService(database.DB)
	telemedLobbyService := services.NewTelemedLobbyService(database.DB, cfg).WithDailyCo(dailyCoService)
	telemedLobbyHandler := handlers.NewTelemedLobbyHandler(telemedLobbyService)

	// Gravação + transcrição da teleconsulta (Daily.co cloud recording + Deepgram nova-3).
	// O webhook do Daily entrega os artefatos; o handler de gravação expõe status +
	// link de download sob demanda do MP4 (referência, não armazenado).
	telemedRecordingService := services.NewTelemedRecordingService(database.DB, dailyCoService, aiService, "/app/uploads")
	telemedRecordingHandler := handlers.NewTelemedRecordingHandler(telemedRecordingService)
	dailyWebhookHandler := handlers.NewDailyWebhookHandler(cfg, dailyCoService, telemedRecordingService)

	// HIGH H9 — patientAppointmentService precisa do DailyCo pra emitir
	// meeting_token quando paciente abre detalhe da consulta no portal.
	patientAppointmentService.WithDailyCo(dailyCoService)

	// Plugar lobby no AppointmentService — quando appointment de telemedicina
	// é criado, AppointmentService.createDailyRoom também ensure-token gera o
	// token público pra link /sala/[token] dos emails/WA. TelemedLobbyService
	// também é injetado no notificationService pra montar o link /sala/<token>
	// nos emails/WhatsApp em vez da URL crua da Daily.
	appointmentService.WithTelemedLobby(telemedLobbyService)
	appointmentNotificationService.WithTelemedLobby(telemedLobbyService)
	patientPortalHandler := handlers.NewPatientPortalHandler(patientPortalService, patientDashboardService, nil, patientAppointmentService, patientMessagesService, patientLabsService, patientScoresService, patientProfileService, patientDocumentsService, patientWorkoutsService, patientCheckInsService, notificationPreferencesService, notificationService)

	// Preparação pré-consulta (Fase 2): formulário curado pós-agendamento + upload de exames pelo paciente.
	consultationPrepService := services.NewConsultationPrepService(database.DB)
	consultationPrepHandler := handlers.NewConsultationPrepHandler(consultationPrepService, anonymousScoreService, patientDocumentsService)

	// Documentos clínicos emitidos/assináveis (P3 frente 2) — reusa signature_service (ICP-Brasil)
	// + patient_documents_service (publica no portal). DocumentPDFService gera o PDF (gofpdf).
	documentPDFService := services.NewDocumentPDFService()
	issuedDocumentService := services.NewIssuedDocumentService(database.DB, documentPDFService, signatureService, patientDocumentsService, "/app/uploads")
	issuedDocumentHandler := handlers.NewIssuedDocumentHandler(issuedDocumentService)

	// Plano de cuidado AGIR (P3 frente 3) + relatório longitudinal (3b) reusando go-rod + assinatura.
	carePlanService := services.NewCarePlanService(database.DB)
	carePlanReportService := services.NewCarePlanReportService(database.DB, scoreSnapshotRepo, carePlanService, services.NewScorePDFService(), signatureService, patientDocumentsService)
	carePlanHandler := handlers.NewCarePlanHandler(carePlanService, carePlanReportService)

	// Calendar V1 (Bloco F): IA detecta intent (CONFIRM/CANCEL/RESCHEDULE) em
	// mensagens WhatsApp inbound de pacientes. Hook é registrado no LeadService
	// pra fechar o ciclo sem dep direta entre services.
	leadService.SetPatientInboundWAHook(buildAppointmentIntentHook(
		database.DB, conversationService, appointmentService, notificationService, cfg,
	))

	// Mobile app services + handlers
	deviceTokenService := services.NewDeviceTokenService(database.DB)
	lgpdConsentService := services.NewLGPDConsentService(database.DB)
	mobileConfigService := services.NewMobileConfigService(cfg)
	pushService := services.NewPushService(database.DB)
	notificationService.SetPushSender(pushService)
	dataExportService := services.NewDataExportService(database.DB)
	meMobileHandler := handlers.NewMeMobileHandler(deviceTokenService, lgpdConsentService, authService, dataExportService)
	mobileConfigHandler := handlers.NewMobileConfigHandler(mobileConfigService)
	uploadHandler := handlers.NewUploadHandler("/app/uploads")

	// API v1
	v1 := app.Group("/api/v1")

	// Rotas públicas
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Plenya EMR API v1",
			"status":  "ok",
		})
	})

	// Score Light (públicas — site público de avaliação anônima)
	scoreLight := v1.Group("/score-light")
	scoreLight.Get("/config", anonymousScoreHandler.GetConfig)
	scoreLight.Post("/sessions", anonymousScoreHandler.CreateSession)
	scoreLight.Get("/sessions/:code", anonymousScoreHandler.GetSession)
	scoreLight.Post("/sessions/:code/claim", anonymousScoreHandler.RequestClaim)
	scoreLight.Post("/claim/confirm", anonymousScoreHandler.ConfirmClaim)
	// Rate limit para extração de PDF: 10 uploads/hora por IP (DPIA recomendação)
	pdfUploadLimiter := middleware.NewIPRateLimiter(10, time.Hour)
	scoreLight.Post("/extract-labs", pdfUploadLimiter.Middleware(), anonymousScoreHandler.ExtractLabsFromPDF)
	scoreLight.Delete("/sessions/:code", anonymousScoreHandler.DeleteSession)
	scoreLight.Get("/sessions/:code/export", anonymousScoreHandler.ExportSession)
	// Autenticada — área do paciente no EMR
	scoreLight.Get("/my-sessions", middleware.Auth(cfg), anonymousScoreHandler.MySessions)

	// WhatsApp Business webhook (público — Meta valida via HMAC X-Hub-Signature-256)
	v1.Get("/webhooks/whatsapp", whatsappWebhookHandler.Verify)
	v1.Post("/webhooks/whatsapp", whatsappWebhookHandler.Receive)

	// Daily.co webhook (público — valida via HMAC X-Webhook-Signature). Entrega
	// gravação + transcrição da teleconsulta (recording.* / transcript.*).
	v1.Post("/webhooks/daily", dailyWebhookHandler.Receive)

	// Leads — POST público (form contato com rate limit), GET/PATCH autenticado (admin/staff)
	leadCreateLimiter := middleware.NewIPRateLimiter(5, time.Hour)
	v1.Post("/leads", leadCreateLimiter.Middleware(), leadHandler.PublicCreate)
	leads := v1.Group("/leads")
	leads.Use(middleware.Auth(cfg))
	leads.Use(middleware.RequireRole(models.RoleAdmin, models.RoleSecretary, models.RoleManager))
	leads.Use(middleware.AuditLog(database.DB))
	leads.Get("/", leadHandler.List)
	leads.Get("/stats", leadHandler.Stats)
	leads.Get("/:id", leadHandler.Get)
	leads.Patch("/:id", leadHandler.Update)
	leads.Post("/:id/notes", leadHandler.AddNote)
	leads.Post("/:id/messages", leadHandler.SendMessage)
	leads.Post("/:id/email-reply", leadHandler.SendEmailReply)
	leads.Post("/:id/convert", leadHandler.Convert)
	// Delete: admin only (destrutivo / LGPD)
	leads.Delete("/:id", middleware.RequireAdmin(), leadHandler.Delete)

	// Campaigns — autenticado (admin/manager). QR code endpoint é GET (cacheável,
	// pode ser embedado em <img>) mas mantém auth pra não vazar URLs internas.
	campaigns := v1.Group("/campaigns")
	campaigns.Use(middleware.Auth(cfg))
	campaigns.Use(middleware.RequireRole(models.RoleAdmin, models.RoleManager, models.RoleSecretary))
	campaigns.Use(middleware.AuditLog(database.DB))
	campaigns.Get("/", campaignHandler.List)
	campaigns.Post("/", campaignHandler.Create)
	campaigns.Get("/:id", campaignHandler.Get)
	campaigns.Patch("/:id", campaignHandler.Update)
	campaigns.Post("/:id/archive", campaignHandler.Archive)
	campaigns.Get("/:id/qrcode", campaignHandler.QRCode)
	campaigns.Delete("/:id", middleware.RequireAdmin(), campaignHandler.Delete)

	// Central de Conversas (Bloco B) — Lead+Patient unificados
	conv := v1.Group("/conversations")
	conv.Use(middleware.Auth(cfg))
	conv.Use(middleware.RequireRole(models.RoleAdmin, models.RoleSecretary, models.RoleManager))
	conv.Use(middleware.AuditLog(database.DB))
	conv.Get("/", conversationHandler.List)
	conv.Get("/ai/metrics", conversationHandler.ReceptionMetrics)
	// Bucket "Notificações" (e-mails automáticos). Rotas estáticas — antes das paramétricas
	// /:type/:id pra não colidir.
	conv.Get("/notifications", notificationEmailHandler.List)
	conv.Post("/notifications/:id/read", notificationEmailHandler.MarkRead)
	conv.Post("/attachments/upload", conversationAttachmentHandler.Upload)
	conv.Post("/compose", conversationHandler.Compose)
	conv.Get("/:type/:id/messages", conversationHandler.Messages)
	conv.Get("/:type/:id/media/:activityId", conversationHandler.Media)
	conv.Post("/:type/:id/media/:activityId/interpret-exam", conversationHandler.InterpretExam)
	conv.Get("/:type/:id/messages/:activityId/attachments/:idx", conversationHandler.Attachment)
	conv.Post("/:type/:id/messages/:activityId/save-to-prontuario", conversationHandler.SaveToProntuario)
	conv.Post("/:type/:id/read", conversationHandler.MarkRead)
	conv.Post("/:type/:id/email", conversationHandler.SendEmail)
	conv.Post("/:type/:id/whatsapp", conversationHandler.SendWhatsApp)
	conv.Post("/:type/:id/whatsapp/template", conversationHandler.SendWhatsAppTemplate)
	// IA: rate-limited 10 req/min/IP (resumo + sugestão são caros).
	conversationAILimiter := middleware.NewRateLimiter(10, time.Minute)
	conv.Post("/:type/:id/ai/summary", conversationAILimiter.Middleware(), conversationHandler.AISummary)
	conv.Post("/:type/:id/ai/suggest-reply", conversationAILimiter.Middleware(), conversationHandler.AISuggestReply)
	conv.Post("/:type/:id/ai/reception-reply", conversationAILimiter.Middleware(), conversationHandler.AIReceptionReply)
	conv.Get("/:type/:id/automation", conversationHandler.GetAutomation)
	conv.Put("/:type/:id/automation", conversationHandler.SetAutomation)

	// Auth routes (públicas).
	// CRITICAL C4 — rate limit em login/register (anti brute-force).
	// 2 camadas: por IP (10/min) e por email (5/15min).
	loginIPLimiter := middleware.NewRateLimiter(10, time.Minute)
	loginEmailLimiter := middleware.NewEmailRateLimiter(5, 15*time.Minute)
	auth := v1.Group("/auth")
	auth.Post("/register", loginIPLimiter.Middleware(), authHandler.Register)
	auth.Post("/login", loginIPLimiter.Middleware(), loginEmailLimiter.Middleware(), authHandler.Login)
	auth.Post("/login/verify-2fa", loginIPLimiter.Middleware(), authHandler.VerifyMFA)
	auth.Post("/refresh", loginIPLimiter.Middleware(), authHandler.Refresh)
	auth.Post("/logout", middleware.Auth(cfg), authHandler.Logout)
	// CRITICAL C6 — 2FA enrollment endpoints (autenticados).
	auth.Post("/2fa/enable", middleware.Auth(cfg), authHandler.Enable2FA)
	auth.Post("/2fa/verify", middleware.Auth(cfg), authHandler.Verify2FA)

	// OAuth routes (públicas, com rate limiting)
	oauthLimiter := middleware.NewRateLimiter(5, time.Minute) // 5 req/min
	auth.Post("/oauth/google", oauthLimiter.Middleware(), oauthHandler.GoogleCallback)
	auth.Post("/oauth/apple", oauthLimiter.Middleware(), oauthHandler.AppleCallback)

	// Lobby telemedicina — públicos (sem auth, token na URL).
	// Rate limit liberal: paciente pode dar reload na página.
	lobbyLimiter := middleware.NewRateLimiter(60, time.Minute)
	v1.Get("/sala/:token", lobbyLimiter.Middleware(), telemedLobbyHandler.Resolve)
	v1.Post("/sala/:token/used", lobbyLimiter.Middleware(), telemedLobbyHandler.MarkUsed)

	// Portal do Paciente — auth público (rate-limited)
	portalAuthLimiter := middleware.NewRateLimiter(10, time.Minute)
	patientAuth := auth.Group("/patient", portalAuthLimiter.Middleware())
	patientAuth.Post("/login", patientPortalHandler.LoginPassword)
	patientAuth.Post("/magic-link", patientPortalHandler.RequestMagicLink)
	patientAuth.Post("/magic-link/consume", patientPortalHandler.ConsumeMagicLink)
	patientAuth.Post("/forgot", patientPortalHandler.RequestMagicLink) // alias
	patientAuth.Post("/invite/consume", patientPortalHandler.ConsumeInvite)

	// User routes (protegidas)
	users := v1.Group("/users")
	users.Use(middleware.Auth(cfg))
	users.Get("/me", authHandler.GetMe)
	users.Put("/me/selected-patient", authHandler.UpdateSelectedPatient)
	users.Patch("/me/preferences", authHandler.UpdatePreferences)
	// Staff list (admin/secretary/manager) — pra dropdown de atribuição de leads no CRM
	users.Get("/staff", middleware.RequireRole(models.RoleAdmin, models.RoleSecretary, models.RoleManager), userHandler.ListStaff)
	// Lista de doctors — pra dropdown de agendamento (Calendar V1).
	users.Get("/doctors", middleware.RequireRole(models.RoleAdmin, models.RoleSecretary, models.RoleManager, models.RoleDoctor, models.RoleNurse), userHandler.ListDoctors)

	// Mobile app endpoints — escopo /me, autenticados
	me := v1.Group("/me")
	me.Use(middleware.Auth(cfg))
	me.Post("/device-tokens", meMobileHandler.RegisterDeviceToken)
	me.Delete("/device-tokens/:id", meMobileHandler.RemoveDeviceToken)
	me.Get("/sessions", meMobileHandler.ListSessions)
	me.Delete("/sessions/:id", meMobileHandler.RevokeSession)
	me.Get("/consent/lgpd", meMobileHandler.GetLGPDConsent)
	me.Post("/consent/lgpd", meMobileHandler.AcceptLGPDConsent)
	me.Post("/password", meMobileHandler.ChangePassword)
	me.Post("/2fa/disable", meMobileHandler.Disable2FA)
	me.Get("/export", meMobileHandler.ExportData)

	// Mobile app config (público — consultado em cold start)
	v1.Get("/mobile/config", mobileConfigHandler.Get)

	// Upload genérico (autenticado) — usado por avaliação física, anamnese, etc.
	v1.Post("/uploads", middleware.Auth(cfg), uploadHandler.Create)

	// Profile routes (protegidas)
	profile := v1.Group("/profile")
	profile.Use(middleware.Auth(cfg))
	profile.Get("/", profileHandler.GetProfile)
	profile.Put("/", profileHandler.UpdateProfile)

	// Admin users management (admin only)
	adminUsers := v1.Group("/admin/users")
	adminUsers.Use(middleware.Auth(cfg))
	adminUsers.Use(middleware.RequireAdmin())
	adminUsers.Use(middleware.AuditLog(database.DB))
	adminUsers.Get("/", userHandler.GetUsers)
	adminUsers.Get("/:id", userHandler.GetUser)
	adminUsers.Post("/", userHandler.CreateUser)
	adminUsers.Put("/:id", userHandler.UpdateUser)
	adminUsers.Delete("/:id", userHandler.DeleteUser)

	// Patients routes (protegidas).
	// CRITICAL C2 — RequireAnyStaff garante que role 'patient' NÃO atinge
	// handlers /patients/*. Pacientes acessam dados próprios via /patient/me/*.
	// Todos staff (admin/manager/secretary/doctor/nurse/etc.) podem ver todos
	// pacientes da clínica (decisão de produto: clínica única, multi-tenant em V2).
	patients := v1.Group("/patients")
	patients.Use(middleware.Auth(cfg))
	patients.Use(middleware.RequireAnyStaff())
	patients.Use(middleware.AuditLog(database.DB))

	patients.Get("/", patientHandler.List)
	patients.Post("/", patientHandler.Create)
	patients.Get("/:id", patientHandler.GetByID)
	patients.Put("/:id", patientHandler.Update)
	patients.Delete("/:id", middleware.RequireRole(models.RoleAdmin, models.RoleDoctor, models.RoleManager), patientHandler.Delete)

	// Convite pra área do paciente — admin/manager/secretary disparam.
	portalInviteRoles := middleware.RequireRole(models.RoleAdmin, models.RoleManager, models.RoleSecretary)
	patients.Get("/:id/portal-invite", portalInviteRoles, patientPortalHandler.GetInviteStatus)
	patients.Post("/:id/portal-invite", portalInviteRoles, patientPortalHandler.CreateInvite)

	// Check-ins de bem-estar — leitura staff pra continuum
	patients.Get("/:id/check-ins",
		middleware.RequireRole(models.RoleAdmin, models.RoleManager, models.RoleSecretary, models.RoleDoctor, models.RoleNurse, models.RoleNutritionist, models.RolePsychologist, models.RolePhysicalEducator),
		patientPortalHandler.StaffListCheckIns)

	// Documentos clínicos (V2): staff faz upload, paciente baixa pelo portal
	patients.Get("/:id/documents", patientPortalHandler.StaffListDocuments)
	patients.Post("/:id/documents", middleware.RequireRole(models.RoleAdmin, models.RoleManager, models.RoleSecretary, models.RoleDoctor, models.RoleNurse), patientPortalHandler.StaffUploadDocument)
	v1.Delete("/patient-documents/:id", middleware.Auth(cfg), middleware.RequireRole(models.RoleAdmin, models.RoleManager), patientPortalHandler.StaffDeleteDocument)

	// Endpoints autenticados como paciente (minha.plenyasaude.com.br)
	// Rate limit geral pra área do paciente — proteção contra conta comprometida
	// ou paciente fazendo polling excessivo. 120 req/min cobre uso legítimo.
	patientMeLimiter := middleware.NewRateLimiter(120, time.Minute)
	patientMe := v1.Group("/patient/me", middleware.Auth(cfg), middleware.RequirePatient(database.DB), patientMeLimiter.Middleware(), middleware.AuditLog(database.DB))
	patientMe.Get("/", patientPortalHandler.Me)
	// M5 — limite agressivo na troca de senha (5 tentativas / 15min por user).
	// Mitiga: paciente roubado → atacante muda senha repetidas vezes pra trolar
	// notificações; bot tentando força-bruta na rota.
	patientPasswordLimiter := middleware.NewEmailRateLimiter(5, 15*time.Minute)
	patientMe.Post("/password", patientPasswordLimiter.MiddlewareByUser(), patientPortalHandler.SetPassword)
	patientMe.Get("/dashboard", patientPortalHandler.Dashboard)
	patientMe.Get("/continuum", patientPortalHandler.MyContinuum)
	patientMe.Get("/appointments", patientPortalHandler.ListAppointments)
	patientMe.Get("/appointments/:id", patientPortalHandler.GetAppointment)
	patientMe.Post("/appointments/:id/confirm", patientPortalHandler.ConfirmAppointment)
	patientMe.Post("/appointments/:id/request-cancel", patientPortalHandler.RequestCancelAppointment)
	patientMe.Post("/appointments/:id/request-reschedule", patientPortalHandler.RequestRescheduleAppointment)
	patientMe.Get("/messages", patientPortalHandler.ListMessages)
	patientMe.Post("/messages", patientPortalHandler.SendMessage)
	patientMe.Get("/messages/unread-count", patientPortalHandler.MessagesUnreadCount)
	patientMe.Post("/messages/read", patientPortalHandler.MessagesMarkRead)
	patientMe.Get("/lab-batches", patientPortalHandler.ListLabBatches)
	patientMe.Get("/lab-batches/:id", patientPortalHandler.GetLabBatch)
	patientMe.Get("/lab-requests", patientPortalHandler.ListLabRequests)
	patientMe.Get("/prescriptions", patientPortalHandler.ListPrescriptions)
	patientMe.Get("/physical-assessments", patientPortalHandler.ListPhysicalAssessments)
	patientMe.Get("/physical-assessments/:id/html", patientPortalHandler.GetPhysicalAssessmentHTML)
	patientMe.Get("/scores", patientPortalHandler.ListScores)
	patientMe.Get("/score-snapshots/latest", patientPortalHandler.GetLatestSnapshot)
	patientMe.Get("/score-snapshots/:id", patientPortalHandler.GetCompleteSnapshot)
	patientMe.Put("/profile", patientPortalHandler.UpdateProfile)
	patientMe.Get("/boxes", patientPortalHandler.ListBoxes)
	patientMe.Get("/documents", patientPortalHandler.ListDocuments)
	patientMe.Get("/documents/:id/download", patientPortalHandler.DownloadDocument)
	patientMe.Get("/lgpd/export", patientPortalHandler.LGPDExport)
	patientMe.Post("/lgpd/account-delete-request", patientPortalHandler.LGPDRequestDelete)

	// App paciente — treino + check-ins + prefs de notificação
	patientMe.Get("/workout-plans", patientPortalHandler.ListWorkoutPlans)
	patientMe.Get("/workout-plans/:id", patientPortalHandler.GetWorkoutPlan)
	patientMe.Get("/workout-sessions", patientPortalHandler.ListWorkoutSessions)
	patientMe.Post("/workout-sessions", patientPortalHandler.StartWorkoutSession)
	patientMe.Get("/workout-sessions/:id", patientPortalHandler.GetWorkoutSession)
	patientMe.Post("/workout-sessions/:id/logs", patientPortalHandler.LogWorkoutSet)
	patientMe.Post("/workout-sessions/:id/complete", patientPortalHandler.CompleteWorkoutSession)
	patientMe.Get("/check-ins", patientPortalHandler.ListCheckIns)
	patientMe.Get("/check-ins/today", patientPortalHandler.LatestCheckInToday)
	patientMe.Post("/check-ins", patientPortalHandler.CreateCheckIn)
	patientMe.Get("/notification-preferences", patientPortalHandler.GetNotificationPreferences)
	patientMe.Patch("/notification-preferences", patientPortalHandler.PatchNotificationPreferences)

	// Preparação pré-consulta (formulário curado + upload de exames pelo paciente)
	patientMe.Get("/prep/config", consultationPrepHandler.GetConfig)
	patientMe.Get("/prep", consultationPrepHandler.GetPrep)
	patientMe.Post("/prep", consultationPrepHandler.SubmitPrep)
	patientMe.Post("/documents", consultationPrepHandler.UploadExam)

	// Anamnesis routes (protegidas - profissionais autenticados).
	// C2 — bloqueia patient role.
	anamnesis := v1.Group("/anamnesis")
	anamnesis.Use(middleware.Auth(cfg))
	anamnesis.Use(middleware.RequireAnyStaff())
	anamnesis.Use(middleware.AuditLog(database.DB))

	anamnesis.Get("/", anamnesisHandler.List)
	anamnesis.Post("/", anamnesisHandler.Create) // Qualquer profissional autenticado pode criar
	anamnesis.Get("/:id", anamnesisHandler.GetByID)
	anamnesis.Put("/:id", anamnesisHandler.Update)
	anamnesis.Delete("/:id", anamnesisHandler.Delete)

	// Anamnesis Templates routes (protegidas)
	anamnesisTemplates := v1.Group("/anamnesis-templates")
	anamnesisTemplates.Use(middleware.Auth(cfg))
	anamnesisTemplates.Use(middleware.AuditLog(database.DB))

	anamnesisTemplates.Get("/", anamnesisTemplateHandler.GetAll)
	anamnesisTemplates.Get("/search", anamnesisTemplateHandler.Search)
	anamnesisTemplates.Get("/:id", anamnesisTemplateHandler.GetByID)
	anamnesisTemplates.Post("/", middleware.RequireClinician(), anamnesisTemplateHandler.Create)
	anamnesisTemplates.Put("/:id", middleware.RequireClinician(), anamnesisTemplateHandler.Update)
	anamnesisTemplates.Put("/:id/items", middleware.RequireClinician(), anamnesisTemplateHandler.UpdateItems)
	anamnesisTemplates.Delete("/:id", middleware.RequireAdmin(), anamnesisTemplateHandler.Delete)

	// Appointments routes (protegidas).
	// C2 — bloqueia role patient. Pacientes usam /patient/me/appointments.
	appointments := v1.Group("/appointments")
	appointments.Use(middleware.Auth(cfg))
	appointments.Use(middleware.RequireAnyStaff())
	appointments.Use(middleware.AuditLog(database.DB))

	appointments.Get("/", appointmentHandler.List)
	appointments.Post("/", appointmentHandler.Create)
	// Termo de consentimento de telemedicina (texto canônico) — estático, ANTES de /:id.
	appointments.Get("/telemed-consent-term", appointmentHandler.GetTelemedConsentTerm)
	appointments.Get("/:id", appointmentHandler.GetByID)
	appointments.Put("/:id", appointmentHandler.Update)
	appointments.Post("/:id/cancel", appointmentHandler.Cancel)
	appointments.Post("/:id/confirm", appointmentHandler.Confirm)
	// Fluxo de balcão: recepção registra chegada (check-in) e início do atendimento.
	appointments.Post("/:id/check-in", appointmentHandler.CheckIn)
	appointments.Post("/:id/start", appointmentHandler.StartConsultation)
	// Consentimento de telemedicina (CFM 2.314/2022) — clínico registra no início da sala. Auditado.
	appointments.Post("/:id/telemed-consent", middleware.RequireClinician(), appointmentHandler.RegisterTelemedConsent)
	// HIGH H9 — emite meeting_token de owner pro staff. POST porque cada
	// chamada gera token novo (sem cache) — inviável como GET.
	appointments.Post("/:id/telemed-token", appointmentHandler.GetTelemedToken)
	// Gravação + transcrição da teleconsulta — clínico vê status e baixa o MP4
	// (link assinado sob demanda). Conteúdo clínico → RequireClinician.
	appointments.Get("/:id/telemed-recording", middleware.RequireClinician(), telemedRecordingHandler.GetForAppointment)
	appointments.Get("/:id/telemed-recording/download", middleware.RequireClinician(), telemedRecordingHandler.Download)
	// AI scribe: gera anamnese/SOAP a partir do transcript (rascunho revisável).
	appointments.Post("/:id/telemed-recording/generate-note", middleware.RequireClinician(), telemedRecordingHandler.GenerateNote)
	appointments.Delete("/:id", middleware.RequireAdmin(), appointmentHandler.Delete)

	// Notas de evolução clínica (SOAP/APSO) — documentação por consulta.
	// C2 — bloqueia patient role. Profissional autenticado documenta a visita;
	// nota assinada é imutável (correção via adendo).
	clinicalNotes := v1.Group("/clinical-notes")
	clinicalNotes.Use(middleware.Auth(cfg))
	clinicalNotes.Use(middleware.RequireAnyStaff())
	clinicalNotes.Use(middleware.AuditLog(database.DB))
	clinicalNotes.Get("/", clinicalNoteHandler.List) // ?appointmentId= retorna a nota da consulta
	clinicalNotes.Post("/", clinicalNoteHandler.Create)
	clinicalNotes.Get("/:id", clinicalNoteHandler.GetByID)
	clinicalNotes.Put("/:id", clinicalNoteHandler.Update)
	clinicalNotes.Post("/:id/sign", clinicalNoteHandler.Sign)
	clinicalNotes.Post("/:id/amend", clinicalNoteHandler.Amend)
	clinicalNotes.Delete("/:id", clinicalNoteHandler.Delete)

	// Lista de espera (encaixe) — operacional do balcão. Qualquer staff opera.
	waitlist := v1.Group("/waitlist")
	waitlist.Use(middleware.Auth(cfg))
	waitlist.Use(middleware.RequireAnyStaff())
	waitlist.Use(middleware.AuditLog(database.DB))
	waitlist.Get("/", waitlistHandler.List)
	waitlist.Post("/", waitlistHandler.Create)
	waitlist.Put("/:id", waitlistHandler.Update)
	waitlist.Delete("/:id", waitlistHandler.Delete)

	// Retornos previstos (recall) — fila interna da recepção, sem auto-msg.
	recalls := v1.Group("/recalls")
	recalls.Use(middleware.Auth(cfg))
	recalls.Use(middleware.RequireAnyStaff())
	recalls.Use(middleware.AuditLog(database.DB))
	recalls.Get("/", recallHandler.List)
	recalls.Post("/", recallHandler.Create)
	recalls.Put("/:id", recallHandler.Update)
	recalls.Delete("/:id", recallHandler.Delete)

	// Pagamentos de consulta + recibo — operações de recepção/financeiro.
	// RequireAdminOps = admin/manager/secretary (não-clínicos).
	payments := v1.Group("/payments")
	payments.Use(middleware.Auth(cfg))
	payments.Use(middleware.RequireAdminOps())
	payments.Use(middleware.AuditLog(database.DB))
	payments.Get("/", paymentHandler.List)
	payments.Post("/", paymentHandler.Create)
	// Estorno é mais restrito que o resto do grupo: secretária cobra mas não
	// estorna (integridade financeira). Só admin/gerente — espelha a UI, que
	// já esconde o botão de estorno da secretária.
	payments.Post("/:id/refund", middleware.RequireRole(models.RoleAdmin, models.RoleManager), paymentHandler.Refund)
	payments.Get("/:id/receipt", paymentHandler.Receipt)

	// Preços por tipo de consulta — leitura para recepção (AdminOps),
	// escrita só admin.
	prices := v1.Group("/consultation-prices")
	prices.Use(middleware.Auth(cfg))
	prices.Use(middleware.AuditLog(database.DB))
	prices.Get("/", middleware.RequireAdminOps(), paymentHandler.ListPrices)
	prices.Put("/", middleware.RequireAdmin(), paymentHandler.UpsertPrice)

	// Prescriptions routes (protegidas).
	// C2 — bloqueia patient role. Pacientes usam /patient/me/prescriptions.
	prescriptions := v1.Group("/prescriptions")
	prescriptions.Use(middleware.Auth(cfg))
	prescriptions.Use(middleware.RequireAnyStaff())
	prescriptions.Use(middleware.AuditLog(database.DB))

	prescriptions.Get("/", prescriptionHandler.List)
	prescriptions.Post("/", middleware.RequireDoctor(), prescriptionHandler.Create)
	prescriptions.Get("/:id", prescriptionHandler.GetByID)
	prescriptions.Put("/:id", prescriptionHandler.Update)
	prescriptions.Delete("/:id", prescriptionHandler.Delete)
	prescriptions.Post("/:id/sign", middleware.RequireDoctor(), prescriptionHandler.SignAndGenerate)
	prescriptions.Get("/:id/download", prescriptionHandler.Download)

	// Documentos clínicos emitidos (P3 frente 2): atestado/declaração/laudo assináveis.
	// Emitir/assinar é ato médico → RequireDoctor. Leitura/download = qualquer staff. Writes auditados.
	issuedDocs := v1.Group("/issued-documents")
	issuedDocs.Use(middleware.Auth(cfg))
	issuedDocs.Use(middleware.RequireAnyStaff())
	issuedDocs.Use(middleware.AuditLog(database.DB))
	issuedDocs.Get("/:docId", issuedDocumentHandler.GetByID)
	issuedDocs.Get("/:docId/pdf", issuedDocumentHandler.DownloadPDF)
	issuedDocs.Post("/:docId/sign", middleware.RequireDoctor(), issuedDocumentHandler.Sign)
	issuedDocs.Delete("/:docId", middleware.RequireDoctor(), issuedDocumentHandler.Delete)

	// Validation routes (public - no auth)
	v1.Get("/prescriptions/validate/:id", prescriptionHandler.ValidatePublic)
	v1.Get("/lab-requests/validate/:id", labRequestHandler.ValidatePublic)
	v1.Get("/documents/validate/:id", issuedDocumentHandler.ValidatePublic)

	// Certificates routes (admin only)
	adminCertificates := v1.Group("/admin/certificates")
	adminCertificates.Use(middleware.Auth(cfg))
	adminCertificates.Use(middleware.RequireAdmin())
	adminCertificates.Use(middleware.AuditLog(database.DB))
	adminCertificates.Post("/upload", certificateHandler.UploadCertificate)
	adminCertificates.Patch("/:userId/toggle", certificateHandler.ToggleCertificateActive)

	// Certificate routes (authenticated users)
	certificates := v1.Group("/certificates")
	certificates.Use(middleware.Auth(cfg))
	certificates.Get("/", certificateHandler.ListCertificates)
	certificates.Get("/status", certificateHandler.GetCertificateStatus)
	// e-CPF em nuvem (o próprio médico vincula/autoriza o seu certificado)
	certificates.Post("/cloud/link", middleware.RequireDoctor(), certificateHandler.LinkCloudCertificate)
	certificates.Post("/cloud/authorize", middleware.RequireDoctor(), certificateHandler.AuthorizeCloudSigning)
	certificates.Delete("/:userId", certificateHandler.DeleteCertificate)

	// Medication Definitions routes
	medicationDefs := v1.Group("/medication-definitions")
	medicationDefs.Use(middleware.Auth(cfg))

	// Read routes (all authenticated users)
	medicationDefs.Get("/", medicationDefHandler.List)
	medicationDefs.Get("/search", medicationDefHandler.Search)
	medicationDefs.Get("/:id", medicationDefHandler.GetByID)

	// Write routes (admin only)
	medicationDefs.Post("/", middleware.RequireAdmin(), middleware.AuditLog(database.DB), medicationDefHandler.Create)
	medicationDefs.Put("/:id", middleware.RequireAdmin(), middleware.AuditLog(database.DB), medicationDefHandler.Update)
	medicationDefs.Delete("/:id", middleware.RequireAdmin(), middleware.AuditLog(database.DB), medicationDefHandler.Delete)

	// Lab Results routes (protegidas).
	// C2 — bloqueia patient role. Pacientes usam /patient/me/lab-batches.
	labResults := v1.Group("/lab-results")
	labResults.Use(middleware.Auth(cfg))
	labResults.Use(middleware.RequireAnyStaff())
	labResults.Use(middleware.AuditLog(database.DB))

	labResults.Get("/", labResultHandler.List)
	labResults.Post("/", middleware.RequireDoctor(), labResultHandler.Create)
	labResults.Get("/:id", labResultHandler.GetByID)
	labResults.Put("/:id", labResultHandler.Update)
	labResults.Delete("/:id", middleware.RequireAdmin(), labResultHandler.Delete)

	// Lab Result Batches routes (protegidas).
	// C2 — bloqueia patient role.
	labResultBatches := v1.Group("/lab-result-batches")
	labResultBatches.Use(middleware.Auth(cfg))
	labResultBatches.Use(middleware.RequireAnyStaff())
	labResultBatches.Use(middleware.AuditLog(database.DB))

	labResultBatches.Get("/", labResultBatchHandler.List)
	labResultBatches.Post("/", middleware.RequireClinician(), labResultBatchHandler.Create)
	labResultBatches.Get("/:id", labResultBatchHandler.GetByID)
	labResultBatches.Put("/:id", middleware.RequireClinician(), labResultBatchHandler.Update)
	labResultBatches.Delete("/:id", middleware.RequireAdmin(), labResultBatchHandler.Delete)

	// Nested routes para results dentro de batch
	labResultBatches.Post("/:id/results", middleware.RequireClinician(), labResultBatchHandler.AddResult)
	labResultBatches.Put("/:batchId/results/:resultId", middleware.RequireClinician(), labResultBatchHandler.UpdateResult)
	labResultBatches.Delete("/:batchId/results/:resultId", middleware.RequireAdmin(), labResultBatchHandler.DeleteResult)

	// PDF upload route
	labResultBatches.Post("/:batchId/upload-pdf", middleware.RequireClinician(), labResultBatchHandler.UploadPDF)

	// Classification route - re-classifica resultados do batch
	labResultBatches.Post("/:id/classify", middleware.RequireClinician(), labResultBatchHandler.Classify)

	// Results inbox (P3) — fila cross-patient de exames a revisar (NÃO selectedPatient-scoped).
	// Revisar exame é ato clínico → RequireClinician (exclui secretary/manager). Writes auditados.
	labInbox := v1.Group("/lab-result-inbox")
	labInbox.Use(middleware.Auth(cfg))
	labInbox.Use(middleware.RequireClinician())
	labInbox.Use(middleware.AuditLog(database.DB))
	labInbox.Get("/", labResultBatchHandler.Inbox)
	labInbox.Get("/count", labResultBatchHandler.InboxCount)
	labInbox.Get("/:id", labResultBatchHandler.GetForReview)
	labInbox.Post("/:id/acknowledge", labResultBatchHandler.Acknowledge)

	// Processing Jobs routes (protegidas)
	processingJobs := v1.Group("/processing-jobs")
	processingJobs.Use(middleware.Auth(cfg))
	processingJobs.Get("/:id", processingJobHandler.GetByID)

	// Lab Requests routes (protegidas).
	// C2 — bloqueia patient role.
	labRequests := v1.Group("/lab-requests")
	labRequests.Use(middleware.Auth(cfg))
	labRequests.Use(middleware.RequireAnyStaff())
	labRequests.Use(middleware.AuditLog(database.DB))

	labRequests.Post("/", middleware.RequireClinician(), labRequestHandler.CreateLabRequest)
	labRequests.Get("/", labRequestHandler.GetAllLabRequests)
	labRequests.Get("/:id", labRequestHandler.GetLabRequestByID)
	labRequests.Get("/by-date", labRequestHandler.GetLabRequestsByDate)
	labRequests.Get("/by-date-range", labRequestHandler.GetLabRequestsByDateRange)
	labRequests.Put("/:id", middleware.RequireClinician(), labRequestHandler.UpdateLabRequest)
	labRequests.Delete("/:id", middleware.RequireAdmin(), labRequestHandler.DeleteLabRequest)
	labRequests.Post("/:id/generate-pdf", middleware.RequireClinician(), labRequestHandler.GeneratePDF)

	// Lab Requests routes dentro de patients
	patients.Get("/:patientId/lab-requests", labRequestHandler.GetLabRequestsByPatientID)

	// Lab Request Templates routes (protegidas - medical staff)
	labRequestTemplates := v1.Group("/lab-request-templates")
	labRequestTemplates.Use(middleware.Auth(cfg))
	labRequestTemplates.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	labRequestTemplates.Get("/", labRequestTemplateHandler.GetAllLabRequestTemplates)
	labRequestTemplates.Get("/search", labRequestTemplateHandler.SearchLabRequestTemplates)
	labRequestTemplates.Get("/:id", labRequestTemplateHandler.GetLabRequestTemplateByID)

	// Rotas de escrita (medical staff)
	labRequestTemplates.Post("/", middleware.RequireClinician(), labRequestTemplateHandler.CreateLabRequestTemplate)
	labRequestTemplates.Put("/:id", middleware.RequireClinician(), labRequestTemplateHandler.UpdateLabRequestTemplate)
	labRequestTemplates.Put("/:id/tests", middleware.RequireClinician(), labRequestTemplateHandler.UpdateLabRequestTemplateTests)
	labRequestTemplates.Post("/:id/tests", middleware.RequireClinician(), labRequestTemplateHandler.AddLabTestToTemplate)
	labRequestTemplates.Delete("/:id/tests/:testId", middleware.RequireClinician(), labRequestTemplateHandler.RemoveLabTestFromTemplate)
	labRequestTemplates.Delete("/:id", middleware.RequireAdmin(), labRequestTemplateHandler.DeleteLabRequestTemplate)

	// Lab Result Views routes (protegidas - medical staff)
	labResultViews := v1.Group("/lab-result-views")
	labResultViews.Use(middleware.Auth(cfg))
	labResultViews.Use(middleware.AuditLog(database.DB))

	// Leitura (todos autenticados)
	labResultViews.Get("/", labResultViewHandler.GetAll)
	labResultViews.Get("/search", labResultViewHandler.Search)
	labResultViews.Get("/:id", labResultViewHandler.GetByID)

	// Escrita (medical staff)
	labResultViews.Post("/", middleware.RequireClinician(), labResultViewHandler.Create)
	labResultViews.Put("/:id", middleware.RequireClinician(), labResultViewHandler.Update)
	labResultViews.Put("/:id/items", middleware.RequireClinician(), labResultViewHandler.UpdateItems)

	// Deleção (admin)
	labResultViews.Delete("/:id", middleware.RequireAdmin(), labResultViewHandler.Delete)

	// Score Groups routes (todas protegidas - dados proprietários)
	scoreGroups := v1.Group("/score-groups")
	scoreGroups.Use(middleware.Auth(cfg))
	scoreGroups.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	scoreGroups.Get("/", scoreHandler.GetAllScoreGroups)
	scoreGroups.Get("/tree", scoreHandler.GetAllScoreGroupTrees)
	scoreGroups.Get("/poster-pdf", scoreHandler.GeneratePosterPDF)
	scoreGroups.Get("/:id", scoreHandler.GetScoreGroupByID)
	scoreGroups.Get("/:id/tree", scoreHandler.GetScoreGroupTree)
	scoreGroups.Get("/:groupId/subgroups", scoreHandler.GetSubgroupsByGroupID)

	// Rotas de escrita (admin only)
	scoreGroups.Post("/", middleware.RequireAdmin(), scoreHandler.CreateScoreGroup)
	scoreGroups.Put("/:id", middleware.RequireAdmin(), scoreHandler.UpdateScoreGroup)
	scoreGroups.Delete("/:id", middleware.RequireAdmin(), scoreHandler.DeleteScoreGroup)

	// Score Subgroups routes (todas protegidas)
	scoreSubgroups := v1.Group("/score-subgroups")
	scoreSubgroups.Use(middleware.Auth(cfg))
	scoreSubgroups.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	scoreSubgroups.Get("/:id", scoreHandler.GetScoreSubgroupByID)
	scoreSubgroups.Get("/:subgroupId/items", scoreHandler.GetItemsBySubgroupID)

	// Rotas de escrita (admin only)
	scoreSubgroups.Post("/", middleware.RequireAdmin(), scoreHandler.CreateScoreSubgroup)
	scoreSubgroups.Put("/:id", middleware.RequireAdmin(), scoreHandler.UpdateScoreSubgroup)
	scoreSubgroups.Delete("/:id", middleware.RequireAdmin(), scoreHandler.DeleteScoreSubgroup)

	// Score Items routes (todas protegidas)
	scoreItems := v1.Group("/score-items")
	scoreItems.Use(middleware.Auth(cfg))
	scoreItems.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	scoreItems.Get("/", scoreHandler.GetAllScoreItems)

	// Score Items with Method associations (must be before /:id routes)
	scoreItems.Get("/unassigned", methodHandler.GetUnassignedScoreItems)
	scoreItems.Get("/with-pillars", methodHandler.GetAllScoreItemsWithPillars)

	scoreItems.Get("/:id", scoreHandler.GetScoreItemByID)
	scoreItems.Get("/:itemId/levels", scoreHandler.GetLevelsByItemID)

	// Inverse relationship: ScoreItems -> Articles
	scoreItems.Get("/:id/articles", articleHandler.GetArticlesForScoreItem)

	// Rotas de escrita (admin only)
	scoreItems.Post("/", middleware.RequireAdmin(), scoreHandler.CreateScoreItem)
	scoreItems.Put("/:id", middleware.RequireAdmin(), scoreHandler.UpdateScoreItem)
	scoreItems.Delete("/:id", middleware.RequireAdmin(), scoreHandler.DeleteScoreItem)

	// Enrichment Preparation routes (admin only - ETAPA 1 do plano científico)
	scoreItems.Post("/:id/prepare-enrichment", middleware.RequireAdmin(), enrichmentPrepHandler.PrepareEnrichment)
	scoreItems.Get("/:id/prepare-enrichment", middleware.RequireAdmin(), enrichmentPrepHandler.GetPreparation)
	scoreItems.Delete("/:id/prepare-enrichment", middleware.RequireAdmin(), enrichmentPrepHandler.DeletePreparation)
	v1.Get("/score-items/enrichment/stats", middleware.Auth(cfg), enrichmentPrepHandler.GetStats)

	// Score Levels routes (todas protegidas)
	scoreLevels := v1.Group("/score-levels")
	scoreLevels.Use(middleware.Auth(cfg))
	scoreLevels.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	scoreLevels.Get("/:id", scoreHandler.GetScoreLevelByID)

	// Rotas de escrita (admin only)
	scoreLevels.Post("/", middleware.RequireAdmin(), scoreHandler.CreateScoreLevel)
	scoreLevels.Put("/:id", middleware.RequireAdmin(), scoreHandler.UpdateScoreLevel)
	scoreLevels.Delete("/:id", middleware.RequireAdmin(), scoreHandler.DeleteScoreLevel)

	// Methods routes (clinical methodologies like AGIR)
	methods := v1.Group("/methods")
	methods.Use(middleware.Auth(cfg))
	methods.Use(middleware.AuditLog(database.DB))

	// Read routes (all authenticated users)
	methods.Get("/", methodHandler.GetAllMethods)
	methods.Get("/tree", methodHandler.GetAllMethodsWithTree)
	methods.Get("/:id", methodHandler.GetMethodByID)
	methods.Get("/:id/tree", methodHandler.GetMethodTree)
	methods.Get("/:id/letters", methodHandler.GetMethodLetters)

	// Write routes (admin only)
	methods.Post("/", middleware.RequireAdmin(), methodHandler.CreateMethod)
	methods.Put("/:id", middleware.RequireAdmin(), methodHandler.UpdateMethod)
	methods.Delete("/:id", middleware.RequireAdmin(), methodHandler.DeleteMethod)
	methods.Post("/:id/letters", middleware.RequireAdmin(), methodHandler.CreateMethodLetter)

	// Method Letters routes
	methodLetters := v1.Group("/method-letters")
	methodLetters.Use(middleware.Auth(cfg))
	methodLetters.Use(middleware.AuditLog(database.DB))

	// Read routes (all authenticated users)
	methodLetters.Get("/:id", methodHandler.GetMethodLetterByID)
	methodLetters.Get("/:id/pillars", methodHandler.GetLetterPillars)

	// Write routes (admin only)
	methodLetters.Post("/:id/pillars", middleware.RequireAdmin(), methodHandler.CreateMethodPillar)
	methodLetters.Put("/:id", middleware.RequireAdmin(), methodHandler.UpdateMethodLetter)
	methodLetters.Delete("/:id", middleware.RequireAdmin(), methodHandler.DeleteMethodLetter)

	// Method Pillars routes
	methodPillars := v1.Group("/method-pillars")
	methodPillars.Use(middleware.Auth(cfg))
	methodPillars.Use(middleware.AuditLog(database.DB))

	// Read routes (all authenticated users)
	methodPillars.Get("/:id", methodHandler.GetMethodPillarByID)

	// Write routes (admin only)
	methodPillars.Put("/:id", middleware.RequireAdmin(), methodHandler.UpdateMethodPillar)
	methodPillars.Delete("/:id", middleware.RequireAdmin(), methodHandler.DeleteMethodPillar)

	// Score Item assignment routes (admin only)
	methodPillars.Post("/:id/assign-item", middleware.RequireAdmin(), methodHandler.AssignScoreItemToPillar)
	methodPillars.Delete("/:id/unassign-item", middleware.RequireAdmin(), methodHandler.UnassignScoreItemFromPillar)

	// Articles routes (todas protegidas - compartilhadas entre médicos)
	articles := v1.Group("/articles")
	articles.Use(middleware.Auth(cfg))
	articles.Use(middleware.AuditLog(database.DB))

	// CRUD básico
	articles.Post("/", middleware.RequireClinician(), articleHandler.CreateArticle)
	articles.Get("/", articleHandler.ListArticles)
	articles.Get("/search", articleHandler.SearchArticles)
	articles.Get("/favorites", articleHandler.GetFavorites)
	articles.Get("/stats", articleHandler.GetStatistics)

	// Upload de artigo (PDF, EPUB, TXT, MD)
	articles.Post("/upload", middleware.RequireClinician(), articleHandler.UploadFile)

	// RAG Semantic Search routes (Phase 4) - MUST BE BEFORE /:id
	articles.Get("/semantic-search", articleSemanticHandler.SemanticSearch)
	articles.Get("/recommend", articleSemanticHandler.RecommendForScoreItem)
	articles.Get("/embedding-stats", articleSemanticHandler.GetEmbeddingStats)

	// Generic ID routes - MUST BE LAST to avoid capturing specific paths
	articles.Get("/:id", articleHandler.GetArticle)
	articles.Get("/:id/chapters", articleHandler.GetChapters)
	articles.Get("/:id/related-score-items", articleSemanticHandler.DiscoverRelatedScoreItems)
	articles.Put("/:id", middleware.RequireClinician(), articleHandler.UpdateArticle)
	articles.Delete("/:id", middleware.RequireClinician(), articleHandler.DeleteArticle)

	// Lab Test Definitions routes (todas protegidas)
	labTests := v1.Group("/lab-tests")
	labTests.Use(middleware.Auth(cfg))
	labTests.Use(middleware.AuditLog(database.DB))

	// Rotas de leitura (todos usuários autenticados)
	labTests.Get("/requestable", labTestDefHandler.GetRequestableLabTests)
	labTests.Get("/definitions", labTestDefHandler.GetAllLabTestDefinitions)
	labTests.Get("/definitions/search", labTestDefHandler.SearchLabTestDefinitions)
	labTests.Get("/definitions/:id", labTestDefHandler.GetLabTestDefinitionByID)
	labTests.Get("/definitions/code/:code", labTestDefHandler.GetLabTestDefinitionByCode)
	labTests.Get("/definitions/:id/sub-tests", labTestDefHandler.GetSubTests)

	// Rotas de escrita (admin only)
	labTests.Post("/definitions", middleware.RequireAdmin(), labTestDefHandler.CreateLabTestDefinition)
	labTests.Put("/definitions/:id", middleware.RequireAdmin(), labTestDefHandler.UpdateLabTestDefinition)
	labTests.Delete("/definitions/:id", middleware.RequireAdmin(), labTestDefHandler.DeleteLabTestDefinition)

	// Lab Result Values routes (protegidas - doctors)
	labResultValues := v1.Group("/lab-results")
	labResultValues.Use(middleware.Auth(cfg))
	labResultValues.Use(middleware.AuditLog(database.DB))

	// Rotas de valores (doctors podem criar/editar)
	labResultValues.Post("/values", middleware.RequireClinician(), labResultValueHandler.CreateLabResultValue)
	labResultValues.Post("/values/batch", middleware.RequireClinician(), labResultValueHandler.CreateLabResultValues)
	labResultValues.Get("/values/:id", labResultValueHandler.GetLabResultValueByID)
	labResultValues.Get("/:id/values", labResultValueHandler.GetValuesByLabResult)
	labResultValues.Put("/values/:id", middleware.RequireClinician(), labResultValueHandler.UpdateLabResultValue)
	labResultValues.Delete("/values/:id", middleware.RequireAdmin(), labResultValueHandler.DeleteLabResultValue)

	// Rotas específicas do paciente (dentro de /patients/:patientId)
	patients.Get("/:patientId/lab-values", labResultValueHandler.GetValuesByPatient)
	patients.Get("/:patientId/lab-values/test/:testId/latest", labResultValueHandler.GetLatestValueForTest)

	// Score Snapshots routes (patient-scoped)
	patients.Post("/:id/score-snapshots", scoreSnapshotHandler.CalculateSnapshot)
	patients.Get("/:id/score-snapshots", scoreSnapshotHandler.GetSnapshotsByPatientID)
	patients.Get("/:id/score-snapshots/latest", scoreSnapshotHandler.GetLatestSnapshotByPatientID)

	// Esqueleto clínico P2a — alergias + sinais vitais por consulta, escopados
	// por paciente (path). RequireAnyStaff (já no grupo patients) + AuditLog.
	patients.Get("/:id/allergies", allergyHandler.ListByPatient)
	patients.Post("/:id/allergies", allergyHandler.Create)
	patients.Put("/:id/allergies/:allergyId", allergyHandler.Update)
	patients.Delete("/:id/allergies/:allergyId", allergyHandler.Delete)
	patients.Get("/:id/vitals", vitalsHandler.ListByPatient)
	patients.Post("/:id/vitals", vitalsHandler.Create)

	// Esqueleto clínico P2b — lista de problemas (CID-10) + medicações em uso.
	patients.Get("/:id/problems", problemHandler.ListByPatient)
	patients.Post("/:id/problems", problemHandler.Create)
	patients.Put("/:id/problems/:problemId", problemHandler.Update)
	patients.Delete("/:id/problems/:problemId", problemHandler.Delete)
	patients.Get("/:id/medications", medicationInUseHandler.ListByPatient)
	patients.Post("/:id/medications", medicationInUseHandler.Create)
	patients.Put("/:id/medications/:medId", medicationInUseHandler.Update)
	patients.Delete("/:id/medications/:medId", medicationInUseHandler.Delete)

	// Documentos emitidos por paciente (P3 frente 2). Emitir = ato médico (RequireDoctor).
	patients.Get("/:id/issued-documents", issuedDocumentHandler.ListByPatient)
	patients.Post("/:id/issued-documents", middleware.RequireDoctor(), issuedDocumentHandler.Create)

	// Plano de cuidado AGIR (P3 frente 3). Writes = clínicos (multidisciplinar).
	patients.Get("/:id/care-plan-items", carePlanHandler.ListByPatient)
	patients.Post("/:id/care-plan-items", middleware.RequireClinician(), carePlanHandler.Create)
	patients.Put("/:id/care-plan-items/:itemId", middleware.RequireClinician(), carePlanHandler.Update)
	patients.Delete("/:id/care-plan-items/:itemId", middleware.RequireClinician(), carePlanHandler.Delete)
	// Relatório longitudinal AGIR — gera/assina/publica no portal (RequireDoctor).
	patients.Post("/:id/care-plan-report", middleware.RequireDoctor(), carePlanHandler.GenerateReport)

	// Catálogo CID-10 (curado) — busca para autocomplete do problem list.
	cidCodes := v1.Group("/cid-codes")
	cidCodes.Use(middleware.Auth(cfg))
	cidCodes.Use(middleware.RequireAnyStaff())
	cidCodes.Get("/", cidHandler.Search)

	// Score Snapshots routes (global).
	// C2 — bloqueia patient role. Pacientes usam /patient/me/score-snapshots/:id.
	scoreSnapshots := v1.Group("/score-snapshots")
	scoreSnapshots.Use(middleware.Auth(cfg))
	scoreSnapshots.Use(middleware.RequireAnyStaff())
	scoreSnapshots.Get("/:id", scoreSnapshotHandler.GetSnapshotByID)
	scoreSnapshots.Delete("/:id", middleware.RequireRole("admin", "doctor"), middleware.AuditLog(database.DB), scoreSnapshotHandler.DeleteSnapshot)

	// Patient Subscriptions routes (patient-scoped)
	patients.Get("/:id/subscriptions", subscriptionHandler.GetPatientSubscriptions)
	patients.Get("/:id/subscriptions/active", subscriptionHandler.GetActivePatientSubscription)
	patients.Post("/:id/subscriptions", middleware.RequireClinician(), middleware.AuditLog(database.DB), subscriptionHandler.CreateSubscription)

	// Subscriptions routes (global)
	subscriptions := v1.Group("/subscriptions")
	subscriptions.Use(middleware.Auth(cfg))
	subscriptions.Get("/", middleware.RequireRole("admin", "doctor"), subscriptionHandler.GetAllSubscriptions)
	subscriptions.Get("/active", middleware.RequireClinician(), subscriptionHandler.GetActiveSubscriptions)
	subscriptions.Get("/:id", subscriptionHandler.GetSubscriptionByID)
	subscriptions.Put("/:id", middleware.RequireClinician(), middleware.AuditLog(database.DB), subscriptionHandler.UpdateSubscription)
	subscriptions.Delete("/:id", middleware.RequireRole("admin", "doctor"), middleware.AuditLog(database.DB), subscriptionHandler.DeleteSubscription)
	subscriptions.Post("/:id/cancel", middleware.AuditLog(database.DB), subscriptionHandler.CancelSubscription) // Paciente pode auto-cancelar
	subscriptions.Post("/:id/renew", middleware.RequireClinician(), middleware.AuditLog(database.DB), subscriptionHandler.RenewSubscription)
	subscriptions.Post("/:id/suspend", middleware.RequireClinician(), middleware.AuditLog(database.DB), subscriptionHandler.SuspendSubscription)
	subscriptions.Post("/:id/activate", middleware.RequireClinician(), middleware.AuditLog(database.DB), subscriptionHandler.ActivateSubscription)

	// Subscription Plans routes
	subscriptionPlans := v1.Group("/subscription-plans")
	subscriptionPlans.Use(middleware.Auth(cfg))

	// Read routes (all authenticated users can view active plans)
	subscriptionPlans.Get("/active", subscriptionPlanHandler.GetActivePlans)
	subscriptionPlans.Get("/:id", subscriptionPlanHandler.GetPlanByID)

	// Admin routes (admin only)
	subscriptionPlans.Get("/", middleware.RequireAdmin(), subscriptionPlanHandler.GetAllPlans)
	subscriptionPlans.Post("/", middleware.RequireAdmin(), middleware.AuditLog(database.DB), subscriptionPlanHandler.CreatePlan)
	subscriptionPlans.Put("/:id", middleware.RequireAdmin(), middleware.AuditLog(database.DB), subscriptionPlanHandler.UpdatePlan)
	subscriptionPlans.Delete("/:id", middleware.RequireAdmin(), middleware.AuditLog(database.DB), subscriptionPlanHandler.DeletePlan)
	subscriptionPlans.Post("/:id/activate", middleware.RequireAdmin(), middleware.AuditLog(database.DB), subscriptionPlanHandler.ActivatePlan)
	subscriptionPlans.Post("/:id/deactivate", middleware.RequireAdmin(), middleware.AuditLog(database.DB), subscriptionPlanHandler.DeactivatePlan)

	// Notifications routes (protegidas - usuários autenticados)
	notifications := v1.Group("/notifications")
	notifications.Use(middleware.Auth(cfg))
	notifications.Get("/", notificationHandler.GetUserNotifications)
	notifications.Get("/unread", notificationHandler.GetUnreadNotifications)
	notifications.Get("/unread/count", notificationHandler.GetUnreadCount)
	notifications.Post("/read-all", notificationHandler.MarkAllAsRead)
	notifications.Post("/:id/read", notificationHandler.MarkAsRead)
	notifications.Delete("/:id", notificationHandler.DeleteNotification)
	notifications.Delete("/", notificationHandler.DeleteAllNotifications)

	// Favorito e rating (todos usuários autenticados podem usar)
	articles.Patch("/:id/favorite", articleHandler.ToggleFavorite)
	articles.Patch("/:id/rating", articleHandler.SetRating)

	// Download de PDF
	articles.Get("/:id/download", articleHandler.DownloadPDF)

	// Many-to-many relationship with ScoreItems
	articles.Post("/:id/score-items", middleware.RequireClinician(), articleHandler.AddScoreItemsToArticle)
	articles.Delete("/:id/score-items", middleware.RequireClinician(), articleHandler.RemoveScoreItemsFromArticle)
	articles.Get("/:id/score-items", articleHandler.GetScoreItemsForArticle)

	// === Calendar V1 (Blocos B + C + D + E) ===
	// Google Calendar OAuth + integração + Daily.co + slot picker.
	// googleCalendarService/dailyCoService/appointmentNotificationService já
	// foram inicializados no topo (são deps do AppointmentService).
	calendarSlotService := services.NewCalendarSlotService(database.DB, googleCalendarService)
	googleOAuthHandler := handlers.NewGoogleOAuthHandler(cfg, googleCalendarService, database.DB)
	calendarSlotHandler := handlers.NewCalendarSlotHandler(calendarSlotService)

	// Recepcionista virtual Fase 3: provê horários reais ao cérebro do bot (best-effort).
	conversationService.SetReceptionSlotsProvider(func(ctx context.Context) string {
		return services.BuildUpcomingSlotsText(ctx, database.DB, calendarSlotService, cfg.ReceptionBot.ConsultDoctorID)
	})

	// /api/v1/integrations/google
	// Importante: callback NÃO usa middleware Auth (Google redirect não traz JWT;
	// state JWT valida CSRF dentro do handler).
	v1.Get("/integrations/google/callback", googleOAuthHandler.Callback)

	googleAuthed := v1.Group("/integrations/google", middleware.Auth(cfg))
	googleAuthed.Get("/auth-url", googleOAuthHandler.AuthURL)
	googleAuthed.Post("/disconnect", googleOAuthHandler.Disconnect)
	googleAuthed.Get("/status", googleOAuthHandler.Status)

	// /api/v1/calendar/slots — slot picker.
	// RBAC: admin/secretary/manager/doctor podem listar.
	calendar := v1.Group("/calendar", middleware.Auth(cfg))
	calendar.Get("/slots",
		middleware.RequireRole(models.RoleAdmin, models.RoleSecretary, models.RoleManager, models.RoleDoctor),
		calendarSlotHandler.List,
	)

	// /api/v1/doctors/:doctorId/working-hours + /absences — CRUD pra UI da agenda.
	// RBAC interno (handler): admin/manager OU o próprio doctor pra mutações;
	// qualquer staff pra leitura.
	workingHoursHandler := handlers.NewWorkingHoursHandler(database.DB)
	doctorAbsenceHandler := handlers.NewDoctorAbsenceHandler(database.DB)
	doctors := v1.Group("/doctors", middleware.Auth(cfg))
	doctors.Get("/:doctorId/working-hours", workingHoursHandler.List)
	doctors.Post("/:doctorId/working-hours", workingHoursHandler.Create)
	doctors.Put("/:doctorId/working-hours/:id", workingHoursHandler.Update)
	doctors.Delete("/:doctorId/working-hours/:id", workingHoursHandler.Delete)
	doctors.Get("/:doctorId/absences", doctorAbsenceHandler.List)
	doctors.Post("/:doctorId/absences", doctorAbsenceHandler.Create)
	doctors.Delete("/:doctorId/absences/:id", doctorAbsenceHandler.Delete)

	// === Continuum (Fase 1) — Templates CRUD + clone ===
	// Programa de acompanhamento longitudinal (Semestral/Anual). Admin/manager
	// edita templates oficiais e cria customizações (ex: Trimestral via clone).
	// Inscrição de paciente (Fase 2) faz snapshot do template no momento.
	continuumTemplateService := services.NewContinuumTemplateService(database.DB)
	continuumBoxTemplateService := services.NewContinuumBoxTemplateService(database.DB)
	continuumTemplateHandler := handlers.NewContinuumTemplateHandler(continuumTemplateService)
	continuumBoxTemplateHandler := handlers.NewContinuumBoxTemplateHandler(continuumBoxTemplateService)

	allStaff := []models.Role{
		models.RoleAdmin, models.RoleManager, models.RoleSecretary,
		models.RoleDoctor, models.RoleNutritionist,
		models.RolePsychologist, models.RolePhysicalEducator,
	}
	continuum := v1.Group("/continuum", middleware.Auth(cfg))
	// Leitura aberta a qualquer staff (médico precisa ver pra entender plano).
	continuum.Get("/templates", middleware.RequireRole(allStaff...), continuumTemplateHandler.List)
	continuum.Get("/templates/:id", middleware.RequireRole(allStaff...), continuumTemplateHandler.Get)
	continuum.Get("/box-templates", middleware.RequireRole(allStaff...), continuumBoxTemplateHandler.List)
	continuum.Get("/box-templates/:id", middleware.RequireRole(allStaff...), continuumBoxTemplateHandler.Get)
	// Mutações restritas a admin/manager.
	continuum.Post("/templates", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumTemplateHandler.Create)
	continuum.Put("/templates/:id", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumTemplateHandler.Update)
	continuum.Delete("/templates/:id", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumTemplateHandler.Delete)
	continuum.Post("/templates/:id/clone", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumTemplateHandler.Clone)
	continuum.Post("/box-templates", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumBoxTemplateHandler.Create)
	continuum.Put("/box-templates/:id", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumBoxTemplateHandler.Update)
	continuum.Delete("/box-templates/:id", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumBoxTemplateHandler.Delete)
	continuum.Post("/box-templates/:id/clone", middleware.RequireRole(models.RoleAdmin, models.RoleManager), middleware.AuditLog(database.DB), continuumBoxTemplateHandler.Clone)

	// === Continuum (Fase 2) — Inscrição de paciente + timeline ===
	patientContinuumService := services.NewPatientContinuumService(database.DB)
	patientContinuumHandler := handlers.NewPatientContinuumHandler(patientContinuumService)

	// Portal do Paciente — injeta continuum service que acabou de ser criado
	// (handler foi instanciado mais cedo no setup, sem essa dep).
	patientPortalHandler.WithContinuum(patientContinuumService)

	enrollManageRoles := []models.Role{models.RoleAdmin, models.RoleManager, models.RoleSecretary}

	// Listar/ver inscrição: qualquer staff (médico precisa ver timeline do paciente).
	v1.Get("/patients/:patientId/continuum",
		middleware.Auth(cfg),
		middleware.RequireRole(allStaff...),
		patientContinuumHandler.ListByPatient,
	)
	continuum.Get("/enrollments/:id",
		middleware.RequireRole(allStaff...),
		patientContinuumHandler.Get,
	)

	// Mutações: inscrever, atualizar item, cancelar — admin/manager/secretary.
	v1.Post("/patients/:patientId/continuum",
		middleware.Auth(cfg),
		middleware.RequireRole(enrollManageRoles...),
		middleware.AuditLog(database.DB),
		patientContinuumHandler.Enroll,
	)
	continuum.Put("/items/:id",
		middleware.RequireRole(enrollManageRoles...),
		middleware.AuditLog(database.DB),
		patientContinuumHandler.UpdateItem,
	)
	continuum.Delete("/enrollments/:id",
		middleware.RequireRole(enrollManageRoles...),
		middleware.AuditLog(database.DB),
		patientContinuumHandler.Cancel,
	)

	// Plano integrado (Fase 4): qualquer profissional + secretário pode editar.
	// Histórico é leitura aberta a todo staff.
	planEditRoles := []models.Role{
		models.RoleAdmin, models.RoleManager, models.RoleSecretary,
		models.RoleDoctor, models.RoleNutritionist,
		models.RolePsychologist, models.RolePhysicalEducator,
	}
	continuum.Put("/enrollments/:id/integrated-plan",
		middleware.RequireRole(planEditRoles...),
		middleware.AuditLog(database.DB),
		patientContinuumHandler.UpdateIntegratedPlan,
	)
	continuum.Get("/enrollments/:id/integrated-plan/revisions",
		middleware.RequireRole(allStaff...),
		patientContinuumHandler.ListPlanRevisions,
	)

	// === Continuum (Fase 7) — Panorama equipe (3 visões) ===
	continuumDashboardService := services.NewContinuumDashboardService(database.DB)
	continuumDashboardHandler := handlers.NewContinuumDashboardHandler(continuumDashboardService)
	continuum.Get("/dashboard/patients",
		middleware.RequireRole(allStaff...),
		continuumDashboardHandler.PerPatient,
	)
	continuum.Get("/dashboard/week",
		middleware.RequireRole(allStaff...),
		continuumDashboardHandler.PerWeek,
	)
	continuum.Get("/dashboard/alerts",
		middleware.RequireRole(allStaff...),
		continuumDashboardHandler.Alerts,
	)

	// === Continuum (Fase 6) — Prontuário agregado ===
	mrAggregateService := services.NewMedicalRecordAggregateService(database.DB)
	mrAggregateHandler := handlers.NewMedicalRecordAggregateHandler(mrAggregateService)
	v1.Get("/patients/:patientId/prontuario",
		middleware.Auth(cfg),
		middleware.RequireRole(allStaff...),
		mrAggregateHandler.List,
	)

	// === Continuum (Fase 5) — Box logístico ===
	continuumBoxService := services.NewContinuumBoxService(database.DB)
	continuumBoxHandler := handlers.NewContinuumBoxHandler(continuumBoxService)
	// Lista/get/contagem: qualquer staff visualiza.
	continuum.Get("/boxes", middleware.RequireRole(allStaff...), continuumBoxHandler.List)
	continuum.Get("/boxes/counts", middleware.RequireRole(allStaff...), continuumBoxHandler.Counts)
	continuum.Get("/boxes/:id", middleware.RequireRole(allStaff...), continuumBoxHandler.Get)
	// Update (mudar status, registrar tracking) — admin/manager/secretary opera logística.
	continuum.Put("/boxes/:id",
		middleware.RequireRole(enrollManageRoles...),
		middleware.AuditLog(database.DB),
		continuumBoxHandler.Update,
	)

	// Cron de atrasos (1×/dia) — marca pending → missed e notifica coordenador.
	continuumLateJob := scheduler.NewContinuumLateJob(database.DB, notificationService)
	continuumLateJob.Start()

	// Seed dos templates oficiais — idempotente, roda no boot.
	// Falha de seed não derruba o servidor (loga e segue).
	if err := services.NewContinuumSeedService(database.DB).SeedOfficialTemplates(); err != nil {
		log.Printf("⚠️  Continuum seed falhou: %v", err)
	}

	// === Training Module ===
	registerTrainingRoutes(v1, cfg, semanticService)

	// HIGH H1 — REMOVIDO segundo app.Static("/uploads", "./uploads").
	// Mesmo motivo: sem auth, vaza arquivos. Use handlers autenticados.

	// Iniciar notification worker (subscription notifications)
	notificationWorker := workers.NewNotificationWorker(notificationService, 24*time.Hour) // Check every 24 hours
	notificationWorker.Start()

	// === Calendar V1 (Bloco E) — cron jobs ===
	// 1) Reminder T-24h: ticker 1h, envia template WA pra appointments na janela.
	appointmentReminderJob := scheduler.NewAppointmentReminderJob(database.DB, appointmentNotificationService)
	appointmentReminderJob.Start()
	// 2) Google token refresh: ticker 30min, renova access tokens cifrados.
	googleTokenRefreshJob := scheduler.NewGoogleTokenRefreshJob(database.DB, googleCalendarService, notificationService)
	googleTokenRefreshJob.Start()

	// 3) Push reminder T-1h pra app paciente (consulta começa em ~1h).
	appointmentPushReminderJob := scheduler.NewAppointmentPushReminderJob(database.DB, pushService, notificationPreferencesService)
	appointmentPushReminderJob.Start()

	// 4) Workout reminder diário no horário configurado pelo paciente.
	workoutReminderJob := scheduler.NewWorkoutReminderJob(database.DB, pushService, notificationPreferencesService)
	workoutReminderJob.Start()

	// 5) Recepcionista virtual (Fase 2): auto-resposta por fallback de tempo.
	//    Só inicia se RECEPTION_BOT_ENABLED=true (kill switch global).
	conversationAutoReplyJob := scheduler.NewConversationAutoReplyJob(database.DB, conversationService, conversationAutomationService, notificationService, cfg)
	conversationAutoReplyJob.Start()
}

// registerTrainingRoutes registra as rotas do módulo de treinamento
func registerTrainingRoutes(v1 fiber.Router, cfg *config.Config, semanticService *services.ArticleSemanticService) {
	// Services
	clinicalDataService := services.NewClinicalDataService(database.DB)
	exerciseService := services.NewExerciseService(database.DB)
	physicalAssessmentService := services.NewPhysicalAssessmentService(database.DB, clinicalDataService)
	assessmentHTMLService := services.NewAssessmentHTMLService(database.DB)
	fitnessTestService := services.NewFitnessTestService(database.DB)
	posturalAssessmentService := services.NewPosturalAssessmentService(database.DB)
	workoutPlanService := services.NewWorkoutPlanService(database.DB)
	workoutHtmlService := services.NewWorkoutHtmlService(database.DB, "./uploads")
	trainingAIService := services.NewTrainingAIService(database.DB, cfg, semanticService)
	periodizationService := services.NewPeriodizationService(database.DB, trainingAIService)

	// Handlers
	exerciseHandler := handlers.NewExerciseHandler(exerciseService)
	physicalAssessmentHandler := handlers.NewPhysicalAssessmentHandler(physicalAssessmentService, assessmentHTMLService)
	fitnessTestHandler := handlers.NewFitnessTestHandler(fitnessTestService)
	posturalAssessmentHandler := handlers.NewPosturalAssessmentHandler(posturalAssessmentService)
	workoutPlanHandler := handlers.NewWorkoutPlanHandler(workoutPlanService, workoutHtmlService)
	periodizationHandler := handlers.NewPeriodizationHandler(periodizationService)
	trainingAIHandler := handlers.NewTrainingAIHandler(trainingAIService)

	// Public routes (no auth)
	v1.Get("/workout-plans/public/:code", workoutPlanHandler.GetByPublicCode)

	// NSCA reference data (no auth needed)
	v1.Get("/training/nsca-reference", func(c *fiber.Ctx) error {
		return c.JSON(services.NSCAReferenceData)
	})

	// Exercises (authenticated)
	exercises := v1.Group("/exercises")
	exercises.Use(middleware.Auth(cfg))
	exercises.Get("/", exerciseHandler.List)
	exercises.Get("/:id", exerciseHandler.GetByID)

	// Physical Assessments (authenticated, patient-scoped)
	assessments := v1.Group("/physical-assessments")
	assessments.Use(middleware.Auth(cfg))
	assessments.Use(middleware.AuditLog(database.DB))
	assessments.Get("/", physicalAssessmentHandler.List)
	assessments.Post("/", middleware.RequireClinician(), physicalAssessmentHandler.Create)
	assessments.Get("/:id", physicalAssessmentHandler.GetByID)
	assessments.Post("/:id/generate-html", middleware.RequireClinician(), physicalAssessmentHandler.GenerateHTML)
	assessments.Delete("/:id", middleware.RequireAdmin(), physicalAssessmentHandler.Delete)

	// Fitness Tests (authenticated, patient-scoped)
	fitnessTests := v1.Group("/fitness-tests")
	fitnessTests.Use(middleware.Auth(cfg))
	fitnessTests.Use(middleware.AuditLog(database.DB))
	fitnessTests.Get("/", fitnessTestHandler.List)
	fitnessTests.Post("/", middleware.RequireClinician(), fitnessTestHandler.Create)
	fitnessTests.Get("/:id", fitnessTestHandler.GetByID)
	fitnessTests.Delete("/:id", middleware.RequireAdmin(), fitnessTestHandler.Delete)

	// Postural Assessments (authenticated, patient-scoped)
	posturalAssessments := v1.Group("/postural-assessments")
	posturalAssessments.Use(middleware.Auth(cfg))
	posturalAssessments.Use(middleware.AuditLog(database.DB))
	posturalAssessments.Get("/", posturalAssessmentHandler.List)
	posturalAssessments.Post("/", middleware.RequireClinician(), posturalAssessmentHandler.Create)
	posturalAssessments.Get("/:id", posturalAssessmentHandler.GetByID)
	posturalAssessments.Delete("/:id", middleware.RequireAdmin(), posturalAssessmentHandler.Delete)

	// Workout Plans (authenticated, patient-scoped)
	workoutPlans := v1.Group("/workout-plans")
	workoutPlans.Use(middleware.Auth(cfg))
	workoutPlans.Use(middleware.AuditLog(database.DB))
	workoutPlans.Get("/", workoutPlanHandler.List)
	workoutPlans.Post("/", middleware.RequireClinician(), workoutPlanHandler.Create)
	workoutPlans.Get("/:id", workoutPlanHandler.GetByID)
	workoutPlans.Put("/:id", middleware.RequireClinician(), workoutPlanHandler.Update)
	workoutPlans.Post("/:id/generate-html", middleware.RequireClinician(), workoutPlanHandler.GenerateHTML)
	workoutPlans.Delete("/:id", middleware.RequireAdmin(), workoutPlanHandler.Delete)

	// Periodizations (authenticated, patient-scoped)
	periodizations := v1.Group("/periodizations")
	periodizations.Use(middleware.Auth(cfg))
	periodizations.Use(middleware.AuditLog(database.DB))
	periodizations.Get("/", periodizationHandler.List)
	periodizations.Post("/", middleware.RequireClinician(), periodizationHandler.Create)
	periodizations.Post("/generate", middleware.RequireClinician(), periodizationHandler.Generate)
	periodizations.Get("/:id", periodizationHandler.GetByID)
	periodizations.Delete("/:id", middleware.RequireAdmin(), periodizationHandler.Delete)

	// Training AI (authenticated)
	trainingAI := v1.Group("/training/ai")
	trainingAI.Use(middleware.Auth(cfg))
	trainingAI.Post("/chat", trainingAIHandler.Chat)
	trainingAI.Post("/recommendations", trainingAIHandler.Recommendations)
}

// healthCheck verifica se a API está funcionando
// @Summary Health Check
// @Description Verifica se a API e o banco de dados estão funcionais
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func healthCheck(c *fiber.Ctx) error {
	// Verificar conexão com banco
	sqlDB, err := database.DB.DB()
	if err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status": "error",
			"error":  "database connection failed",
		})
	}

	if err := sqlDB.Ping(); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status": "error",
			"error":  "database ping failed",
		})
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"database": "connected",
	})
}

// ragHealthCheck verifica status do sistema RAG/Embeddings
// @Summary RAG Health Check
// @Description Verifica status do sistema de embeddings (queue, OpenAI, pgvector)
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health/rag [get]
func ragHealthCheck(c *fiber.Ctx, db *gorm.DB) error {
	type QueueStats struct {
		Pending    int64 `json:"pending"`
		Processing int64 `json:"processing"`
		Completed  int64 `json:"completed"`
		Failed     int64 `json:"failed"`
	}

	stats := QueueStats{}

	// Count queue by status
	db.Model(&models.EmbeddingQueue{}).Where("status = 'pending'").Count(&stats.Pending)
	db.Model(&models.EmbeddingQueue{}).Where("status = 'processing'").Count(&stats.Processing)
	db.Model(&models.EmbeddingQueue{}).Where("status = 'completed'").Count(&stats.Completed)
	db.Model(&models.EmbeddingQueue{}).Where("status = 'failed'").Count(&stats.Failed)

	// Check pgvector extension
	var pgvectorInstalled bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_extension WHERE extname = 'vector')").Scan(&pgvectorInstalled)

	// Count embeddings
	var embeddingCount int64
	db.Model(&models.ArticleEmbedding{}).Count(&embeddingCount)

	// Determine overall status
	status := "healthy"
	if stats.Failed > 100 {
		status = "degraded"
	}
	if !pgvectorInstalled {
		status = "error"
	}

	return c.JSON(fiber.Map{
		"status":           status,
		"pgvector":         pgvectorInstalled,
		"embeddings_count": embeddingCount,
		"queue":            stats,
		"queue_depth":      stats.Pending + stats.Processing,
	})
}

// customErrorHandler trata erros globais.
//
// M6 — em production, NÃO expomos err.Error() pra client em respostas 5xx.
// O erro real é logado via log.Printf (com path e método pra triagem). Erros
// 4xx (*fiber.Error) preservam mensagem porque são semânticos (404 Not Found,
// 422 Validation, etc.). Em dev mantemos detalhe pra DX.
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := err.Error()

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		// 4xx fiber.Error: preserva msg (semântica) mesmo em prod.
		return c.Status(code).JSON(fiber.Map{
			"error": msg,
			"code":  code,
		})
	}

	// Não-fiber error → 500. Em production, opaco; em dev, detalhe.
	if appConfig != nil && appConfig.Server.Environment == "production" {
		log.Printf("[error 500] %s %s | err=%v", c.Method(), c.Path(), err)
		return c.Status(code).JSON(fiber.Map{
			"error": "internal server error",
			"code":  code,
		})
	}
	return c.Status(code).JSON(fiber.Map{
		"error": msg,
		"code":  code,
	})
}

// keysOf — utilitário pra debug log de mapas string→bool (CORS allowlist).
func keysOf(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// appConfig — referência global pra customErrorHandler decidir verbosidade
// em prod. Setado no início do main() antes do app.Listen.
var appConfig *config.Config

// startProcessingWorker inicia goroutine de processamento de jobs
func startProcessingWorker(service *services.ProcessingJobService) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	log.Println("🤖 Processing worker started (polling every 3s)")

	for range ticker.C {
		if err := service.PollAndProcess(); err != nil {
			log.Printf("⚠️  Worker error: %v", err)
		}
	}
}

// startEmbeddingWorker inicia goroutine de processamento de embeddings (RAG)
func startEmbeddingWorker(worker *workers.EmbeddingWorker) {
	ctx := context.Background()

	if err := worker.Start(ctx); err != nil {
		log.Printf("❌ Embedding worker stopped: %v", err)
	}
}
