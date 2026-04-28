package middleware

import (
	"log"
	"regexp"
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

		// H7 — normaliza path em vez de mapear caso a caso (que perdia rotas).
		// Estratégia: pega path, remove /api/v1, troca UUIDs por :id e usa o
		// resultado como resource_path. Mantemos compat com modelo AuditResource
		// (constantes de conveniência), mas agora qualquer rota é audit-ável.
		resource := extractResource(c.Path())
		if resource == "" {
			// Fallback: usa o path normalizado direto (truncado em 50 chars
			// pra caber no varchar(50) do schema atual).
			resource = models.AuditResource(normalizeResourcePath(c.Path()))
		}
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

// extractResource extrai o nome do recurso conhecido da rota.
// H7 — fallback pra normalizeResourcePath quando não bate em nenhum mapeamento.
func extractResource(path string) models.AuditResource {
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// Mapear rotas conhecidas pra constantes (mantém audits agregáveis).
	switch {
	case contains(path, "patients"):
		return models.ResourcePatients
	case contains(path, "anamnesis"):
		return models.ResourceAnamnesis
	case contains(path, "appointments"):
		return models.ResourceAppointments
	case contains(path, "prescriptions"):
		return models.ResourcePrescriptions
	case contains(path, "lab-results"), contains(path, "lab-result-batches"):
		return models.ResourceLabResults
	case contains(path, "users"), contains(path, "auth"):
		return models.ResourceUsers
	default:
		return ""
	}
}

// uuidPattern — regex pra detectar UUIDs em segmentos de path.
var uuidPattern = regexp.MustCompile(`(?i)[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

// normalizeResourcePath — devolve resource_path normalizado pra uso H7.
// Ex: /api/v1/continuum/templates/abc-123-uuid → "continuum.templates.:id"
func normalizeResourcePath(path string) string {
	p := strings.TrimPrefix(path, "/api/v1/")
	p = strings.TrimPrefix(p, "/")
	p = uuidPattern.ReplaceAllString(p, ":id")
	p = strings.ReplaceAll(p, "/", ".")
	if len(p) > 50 {
		p = p[:50]
	}
	return p
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
