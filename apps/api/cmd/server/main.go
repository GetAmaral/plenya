package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/handlers"
	"github.com/plenya/api/internal/middleware"
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
	})

	// Middlewares globais
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Server.CORSOrigin,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

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

	// Inicializar services
	authService := services.NewAuthService(database.DB, cfg)
	patientService := services.NewPatientService(database.DB)
	anamnesisService := services.NewAnamnesisService(database.DB)
	appointmentService := services.NewAppointmentService(database.DB)
	prescriptionService := services.NewPrescriptionService(database.DB)
	labResultService := services.NewLabResultService(database.DB)

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(authService)
	patientHandler := handlers.NewPatientHandler(patientService)
	anamnesisHandler := handlers.NewAnamnesisHandler(anamnesisService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService)
	labResultHandler := handlers.NewLabResultHandler(labResultService)

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

	// Patients routes (protegidas)
	patients := v1.Group("/patients")
	patients.Use(middleware.Auth(cfg))
	patients.Use(middleware.AuditLog(database.DB))

	patients.Get("/", patientHandler.List)
	patients.Post("/", patientHandler.Create)
	patients.Get("/:id", patientHandler.GetByID)
	patients.Put("/:id", patientHandler.Update)
	patients.Delete("/:id", middleware.RequireMedicalStaff(), patientHandler.Delete)

	// Anamnesis routes (protegidas - apenas doctors)
	anamnesis := v1.Group("/anamnesis")
	anamnesis.Use(middleware.Auth(cfg))
	anamnesis.Use(middleware.AuditLog(database.DB))

	anamnesis.Get("/", anamnesisHandler.List)
	anamnesis.Post("/", middleware.RequireDoctor(), anamnesisHandler.Create)
	anamnesis.Get("/:id", anamnesisHandler.GetByID)
	anamnesis.Put("/:id", anamnesisHandler.Update)
	anamnesis.Delete("/:id", anamnesisHandler.Delete)

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
