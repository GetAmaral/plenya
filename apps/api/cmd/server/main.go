package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
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

	// Inicializar criptografia
	if err := crypto.Init(cfg.Security.EncryptionKey); err != nil {
		log.Fatalf("❌ Failed to initialize crypto: %v", err)
	}
	log.Println("🔐 Encryption initialized")

	// Conectar ao banco de dados
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Auto migrate (desenvolvimento - em produção usaremos Atlas)
	if cfg.Server.Environment == "development" {
		log.Println("🔄 Running auto migrations...")
		if err := database.AutoMigrate(); err != nil {
			log.Fatalf("❌ Failed to auto migrate: %v", err)
		}
		log.Println("✅ Auto migrations completed")
	}

	// Criar app Fiber
	app := fiber.New(fiber.Config{
		AppName:      "Plenya EMR API",
		ServerHeader: "Plenya",
		ErrorHandler: customErrorHandler,
		BodyLimit:    50 * 1024 * 1024, // 50MB para permitir upload de PDFs grandes
	})

	// Middlewares globais
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Server.CORSOrigin,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	// Serve static files (PDFs)
	app.Static("/uploads", "/app/uploads")

	// Rotas
	setupRoutes(app, cfg)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("🛑 Shutting down server...")
		_ = app.Shutdown()
	}()

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s (environment: %s)", cfg.Server.Port, cfg.Server.Environment)
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

// setupRoutes configura as rotas da API
func setupRoutes(app *fiber.App, cfg *config.Config) {
	// Health check
	app.Get("/health", healthCheck)

	// Inicializar repositories
	scoreRepo := repository.NewScoreRepository(database.DB)
	labTestDefRepo := repository.NewLabTestDefinitionRepository(database.DB)
	labResultValueRepo := repository.NewLabResultValueRepository(database.DB)
	labRequestRepo := repository.NewLabRequestRepository(database.DB)
	labRequestTemplateRepo := repository.NewLabRequestTemplateRepository(database.DB)
	anamnesisTemplateRepo := repository.NewAnamnesisTemplateRepository(database.DB)

	// Inicializar services
	authService := services.NewAuthService(database.DB, cfg)
	patientService := services.NewPatientService(database.DB)
	anamnesisService := services.NewAnamnesisService(database.DB)
	anamnesisTemplateService := services.NewAnamnesisTemplateService(anamnesisTemplateRepo)
	appointmentService := services.NewAppointmentService(database.DB)
	prescriptionService := services.NewPrescriptionService(database.DB)
	labResultService := services.NewLabResultService(database.DB)
	scoreService := services.NewScoreService(scoreRepo)
	labTestDefService := services.NewLabTestDefinitionService(labTestDefRepo)
	labResultValueService := services.NewLabResultValueService(labResultValueRepo)
	labRequestService := services.NewLabRequestService(labRequestRepo, database.DB)
	labRequestTemplateService := services.NewLabRequestTemplateService(labRequestTemplateRepo)
	articleService := services.NewArticleService(database.DB, "./uploads/articles")

	// Inicializar validator
	validate := validator.New()

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(authService)
	patientHandler := handlers.NewPatientHandler(patientService)
	anamnesisHandler := handlers.NewAnamnesisHandler(anamnesisService)
	anamnesisTemplateHandler := handlers.NewAnamnesisTemplateHandler(anamnesisTemplateService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService)
	labResultHandler := handlers.NewLabResultHandler(labResultService)
	scoreHandler := handlers.NewScoreHandler(scoreService, validate)
	labTestDefHandler := handlers.NewLabTestDefinitionHandler(labTestDefService)
	labResultValueHandler := handlers.NewLabResultValueHandler(labResultValueService)
	labRequestHandler := handlers.NewLabRequestHandler(labRequestService)
	labRequestTemplateHandler := handlers.NewLabRequestTemplateHandler(labRequestTemplateService)
	articleHandler := handlers.NewArticleHandler(articleService)

	// API v1
	v1 := app.Group("/api/v1")

	// Rotas públicas
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Plenya EMR API v1",
			"status":  "ok",
		})
	})

	// Auth routes (públicas)
	auth := v1.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/logout", middleware.Auth(cfg), authHandler.Logout)

	// User routes (protegidas)
	users := v1.Group("/users")
	users.Use(middleware.Auth(cfg))
	users.Get("/me", authHandler.GetMe)
	users.Put("/me/selected-patient", authHandler.UpdateSelectedPatient)
	users.Patch("/me/preferences", authHandler.UpdatePreferences)

	// Patients routes (protegidas)
	patients := v1.Group("/patients")
	patients.Use(middleware.Auth(cfg))
	patients.Use(middleware.AuditLog(database.DB))

	patients.Get("/", patientHandler.List)
	patients.Post("/", patientHandler.Create)
	patients.Get("/:id", patientHandler.GetByID)
	patients.Put("/:id", patientHandler.Update)
	patients.Delete("/:id", middleware.RequireMedicalStaff(), patientHandler.Delete)

	// Anamnesis routes (protegidas - profissionais autenticados)
	anamnesis := v1.Group("/anamnesis")
	anamnesis.Use(middleware.Auth(cfg))
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
	anamnesisTemplates.Post("/", middleware.RequireMedicalStaff(), anamnesisTemplateHandler.Create)
	anamnesisTemplates.Put("/:id", middleware.RequireMedicalStaff(), anamnesisTemplateHandler.Update)
	anamnesisTemplates.Put("/:id/items", middleware.RequireMedicalStaff(), anamnesisTemplateHandler.UpdateItems)
	anamnesisTemplates.Delete("/:id", middleware.RequireAdmin(), anamnesisTemplateHandler.Delete)

	// Appointments routes (protegidas)
	appointments := v1.Group("/appointments")
	appointments.Use(middleware.Auth(cfg))
	appointments.Use(middleware.AuditLog(database.DB))

	appointments.Get("/", appointmentHandler.List)
	appointments.Post("/", appointmentHandler.Create)
	appointments.Get("/:id", appointmentHandler.GetByID)
	appointments.Put("/:id", appointmentHandler.Update)
	appointments.Post("/:id/cancel", appointmentHandler.Cancel)
	appointments.Delete("/:id", middleware.RequireAdmin(), appointmentHandler.Delete)

	// Prescriptions routes (protegidas - apenas doctors)
	prescriptions := v1.Group("/prescriptions")
	prescriptions.Use(middleware.Auth(cfg))
	prescriptions.Use(middleware.AuditLog(database.DB))

	prescriptions.Get("/", prescriptionHandler.List)
	prescriptions.Post("/", middleware.RequireDoctor(), prescriptionHandler.Create)
	prescriptions.Get("/:id", prescriptionHandler.GetByID)
	prescriptions.Put("/:id", prescriptionHandler.Update)
	prescriptions.Delete("/:id", prescriptionHandler.Delete)

	// Lab Results routes (protegidas - apenas doctors)
	labResults := v1.Group("/lab-results")
	labResults.Use(middleware.Auth(cfg))
	labResults.Use(middleware.AuditLog(database.DB))

	labResults.Get("/", labResultHandler.List)
	labResults.Post("/", middleware.RequireDoctor(), labResultHandler.Create)
	labResults.Get("/:id", labResultHandler.GetByID)
	labResults.Put("/:id", labResultHandler.Update)
	labResults.Delete("/:id", middleware.RequireAdmin(), labResultHandler.Delete)

	// Lab Requests routes (protegidas - medical staff)
	labRequests := v1.Group("/lab-requests")
	labRequests.Use(middleware.Auth(cfg))
	labRequests.Use(middleware.AuditLog(database.DB))

	labRequests.Post("/", middleware.RequireMedicalStaff(), labRequestHandler.CreateLabRequest)
	labRequests.Get("/", labRequestHandler.GetAllLabRequests)
	labRequests.Get("/:id", labRequestHandler.GetLabRequestByID)
	labRequests.Get("/by-date", labRequestHandler.GetLabRequestsByDate)
	labRequests.Get("/by-date-range", labRequestHandler.GetLabRequestsByDateRange)
	labRequests.Put("/:id", middleware.RequireMedicalStaff(), labRequestHandler.UpdateLabRequest)
	labRequests.Delete("/:id", middleware.RequireAdmin(), labRequestHandler.DeleteLabRequest)
	labRequests.Post("/:id/generate-pdf", middleware.RequireMedicalStaff(), labRequestHandler.GeneratePDF)

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
	labRequestTemplates.Post("/", middleware.RequireMedicalStaff(), labRequestTemplateHandler.CreateLabRequestTemplate)
	labRequestTemplates.Put("/:id", middleware.RequireMedicalStaff(), labRequestTemplateHandler.UpdateLabRequestTemplate)
	labRequestTemplates.Put("/:id/tests", middleware.RequireMedicalStaff(), labRequestTemplateHandler.UpdateLabRequestTemplateTests)
	labRequestTemplates.Post("/:id/tests", middleware.RequireMedicalStaff(), labRequestTemplateHandler.AddLabTestToTemplate)
	labRequestTemplates.Delete("/:id/tests/:testId", middleware.RequireMedicalStaff(), labRequestTemplateHandler.RemoveLabTestFromTemplate)
	labRequestTemplates.Delete("/:id", middleware.RequireAdmin(), labRequestTemplateHandler.DeleteLabRequestTemplate)

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
	scoreItems.Get("/:id", scoreHandler.GetScoreItemByID)
	scoreItems.Get("/:itemId/levels", scoreHandler.GetLevelsByItemID)

	// Rotas de escrita (admin only)
	scoreItems.Post("/", middleware.RequireAdmin(), scoreHandler.CreateScoreItem)
	scoreItems.Put("/:id", middleware.RequireAdmin(), scoreHandler.UpdateScoreItem)
	scoreItems.Delete("/:id", middleware.RequireAdmin(), scoreHandler.DeleteScoreItem)

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

	// Articles routes (todas protegidas - compartilhadas entre médicos)
	articles := v1.Group("/articles")
	articles.Use(middleware.Auth(cfg))
	articles.Use(middleware.AuditLog(database.DB))

	// CRUD básico
	articles.Post("/", middleware.RequireMedicalStaff(), articleHandler.CreateArticle)
	articles.Get("/", articleHandler.ListArticles)
	articles.Get("/search", articleHandler.SearchArticles)
	articles.Get("/favorites", articleHandler.GetFavorites)
	articles.Get("/stats", articleHandler.GetStatistics)
	articles.Get("/:id", articleHandler.GetArticle)
	articles.Put("/:id", middleware.RequireMedicalStaff(), articleHandler.UpdateArticle)
	articles.Delete("/:id", middleware.RequireMedicalStaff(), articleHandler.DeleteArticle)

	// Upload de PDF
	articles.Post("/upload", middleware.RequireMedicalStaff(), articleHandler.UploadPDF)

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
	labTests.Get("/definitions/:id/score-mappings", labTestDefHandler.GetMappingsForLabTest)

	// Rotas de escrita (admin only)
	labTests.Post("/definitions", middleware.RequireAdmin(), labTestDefHandler.CreateLabTestDefinition)
	labTests.Put("/definitions/:id", middleware.RequireAdmin(), labTestDefHandler.UpdateLabTestDefinition)
	labTests.Delete("/definitions/:id", middleware.RequireAdmin(), labTestDefHandler.DeleteLabTestDefinition)

	// Score Mappings routes (admin only)
	labTests.Post("/score-mappings", middleware.RequireAdmin(), labTestDefHandler.CreateLabTestScoreMapping)
	labTests.Get("/score-mappings/:id", labTestDefHandler.GetLabTestScoreMappingByID)
	labTests.Put("/score-mappings/:id", middleware.RequireAdmin(), labTestDefHandler.UpdateLabTestScoreMapping)
	labTests.Delete("/score-mappings/:id", middleware.RequireAdmin(), labTestDefHandler.DeleteLabTestScoreMapping)

	// Lab Result Values routes (protegidas - doctors)
	labResultValues := v1.Group("/lab-results")
	labResultValues.Use(middleware.Auth(cfg))
	labResultValues.Use(middleware.AuditLog(database.DB))

	// Rotas de valores (doctors podem criar/editar)
	labResultValues.Post("/values", middleware.RequireMedicalStaff(), labResultValueHandler.CreateLabResultValue)
	labResultValues.Post("/values/batch", middleware.RequireMedicalStaff(), labResultValueHandler.CreateLabResultValues)
	labResultValues.Get("/values/:id", labResultValueHandler.GetLabResultValueByID)
	labResultValues.Get("/:id/values", labResultValueHandler.GetValuesByLabResult)
	labResultValues.Put("/values/:id", middleware.RequireMedicalStaff(), labResultValueHandler.UpdateLabResultValue)
	labResultValues.Delete("/values/:id", middleware.RequireAdmin(), labResultValueHandler.DeleteLabResultValue)

	// Rotas específicas do paciente (dentro de /patients/:patientId)
	patients.Get("/:patientId/lab-values", labResultValueHandler.GetValuesByPatient)
	patients.Get("/:patientId/lab-values/test/:testId/latest", labResultValueHandler.GetLatestValueForTest)

	// Favorito e rating (todos usuários autenticados podem usar)
	articles.Patch("/:id/favorite", articleHandler.ToggleFavorite)
	articles.Patch("/:id/rating", articleHandler.SetRating)

	// Download de PDF
	articles.Get("/:id/download", articleHandler.DownloadPDF)

	// Many-to-many relationship with ScoreItems
	articles.Post("/:id/score-items", middleware.RequireMedicalStaff(), articleHandler.AddScoreItemsToArticle)
	articles.Delete("/:id/score-items", middleware.RequireMedicalStaff(), articleHandler.RemoveScoreItemsFromArticle)
	articles.Get("/:id/score-items", articleHandler.GetScoreItemsForArticle)

	// Servir arquivos estáticos (uploads)
	app.Static("/uploads", "./uploads")
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

// customErrorHandler trata erros globais
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
