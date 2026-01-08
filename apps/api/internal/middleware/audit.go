package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// AuditLog middleware registra todas as ações dos usuários
func AuditLog(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Executar handler
		err := c.Next()

		log.Printf("[AUDIT] Path: %s, Method: %s", c.Path(), c.Method())

		// Extrair informações do contexto
		userID, ok := c.Locals("userID").(uuid.UUID)
		if !ok {
			log.Printf("[AUDIT] No userID found in context")
			// Se não houver userID (rota pública), não registrar
			return err
		}
		log.Printf("[AUDIT] UserID: %s", userID)

		// Determinar ação baseado no método HTTP
		var action models.AuditAction
		switch c.Method() {
		case "GET":
			action = models.ActionView
		case "POST":
			action = models.ActionCreate
		case "PUT", "PATCH":
			action = models.ActionUpdate
		case "DELETE":
			action = models.ActionDelete
		default:
			return err // Não registrar métodos desconhecidos
		}

		// Determinar recurso baseado na rota
		resource := extractResource(c.Path())
		log.Printf("[AUDIT] Extracted resource: '%s' from path: '%s'", resource, c.Path())
		if resource == "" {
			log.Printf("[AUDIT] Resource is empty, skipping audit log")
			return err // Não registrar rotas não identificadas
		}

		// Extrair resourceID se houver (último segmento UUID da rota)
		var resourceID *uuid.UUID
		if id := c.Params("id"); id != "" {
			if parsed, parseErr := uuid.Parse(id); parseErr == nil {
				resourceID = &parsed
			}
		}

		// User-Agent
		userAgent := c.Get("User-Agent")
		var userAgentPtr *string
		if userAgent != "" {
			userAgentPtr = &userAgent
		}

		// Criar log de auditoria
		auditLog := models.AuditLog{
			UserID:       userID,
			Action:       action,
			Resource:     resource,
			ResourceID:   resourceID,
			IPAddress:    c.IP(),
			UserAgent:    userAgentPtr,
			Success:      err == nil && c.Response().StatusCode() < 400,
			ErrorMessage: getErrorMessage(err, c.Response().StatusCode()),
		}

		// Salvar de forma síncrona (mudado de assíncrono para garantir execução)
		log.Printf("[AUDIT] Creating audit log: Resource=%s, Action=%s", resource, action)
		if createErr := db.Create(&auditLog).Error; createErr != nil {
			log.Printf("[AUDIT] Error creating audit log: %v", createErr)
			// Log error but don't fail the request
		} else {
			log.Printf("[AUDIT] Audit log created successfully")
		}

		return err
	}
}

// extractResource extrai o nome do recurso da rota
func extractResource(path string) models.AuditResource {
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// Mapear rotas para recursos
	switch {
	case contains(path, "patients"):
		return models.ResourcePatients
	case contains(path, "anamnesis"):
		return models.ResourceAnamnesis
	case contains(path, "appointments"):
		return models.ResourceAppointments
	case contains(path, "prescriptions"):
		return models.ResourcePrescriptions
	case contains(path, "lab-results"):
		return models.ResourceLabResults
	case contains(path, "users"):
		return models.ResourceUsers
	default:
		return ""
	}
}

// contains verifica se uma string contém outra
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// getErrorMessage extrai mensagem de erro se houver
func getErrorMessage(err error, statusCode int) *string {
	if err != nil {
		msg := err.Error()
		return &msg
	}
	if statusCode >= 400 {
		msg := "request failed"
		return &msg
	}
	return nil
}
