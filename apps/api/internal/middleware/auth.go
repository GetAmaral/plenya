package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// AuthClaims representa os claims do JWT
type AuthClaims struct {
	UserID string   `json:"userId"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// Auth middleware valida JWT token
func Auth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extrair token do header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		// Formato esperado: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization format",
			})
		}

		tokenString := parts[1]

		// Validar token
		token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		claims, ok := token.Claims.(*AuthClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		// Converter UserID para UUID
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid user id",
			})
		}

		// Armazenar informações do usuário no contexto
		c.Locals("userID", userID)
		c.Locals("userEmail", claims.Email)
		c.Locals("userRoles", claims.Roles)

		return c.Next()
	}
}

// GetUserID retorna o ID do usuário do contexto
func GetUserID(c *fiber.Ctx) uuid.UUID {
	return c.Locals("userID").(uuid.UUID)
}

// GetUserEmail retorna o email do usuário do contexto
func GetUserEmail(c *fiber.Ctx) string {
	return c.Locals("userEmail").(string)
}

// GetUserRoles retorna as roles do usuário do contexto
func GetUserRoles(c *fiber.Ctx) []string {
	return c.Locals("userRoles").([]string)
}

// HasRole verifica se o usuário tem uma role específica
func HasRole(c *fiber.Ctx, role models.Role) bool {
	roles := GetUserRoles(c)
	for _, r := range roles {
		if r == string(role) {
			return true
		}
	}
	return false
}

// GetPrimaryRole retorna o role mais privilegiado do usuário (para compatibilidade)
// Ordem de privilégio: admin > manager > doctor > psychologist > nutritionist > physicalEducator > nurse > secretary > patient
func GetPrimaryRole(c *fiber.Ctx) models.Role {
	roles := GetUserRoles(c)

	// Ordem de prioridade
	priorityOrder := []models.Role{
		models.RoleAdmin,
		models.RoleManager,
		models.RoleDoctor,
		models.RolePsychologist,
		models.RoleNutritionist,
		models.RolePhysicalEducator,
		models.RoleNurse,
		models.RoleSecretary,
		models.RolePatient,
	}

	for _, priority := range priorityOrder {
		for _, userRole := range roles {
			if userRole == string(priority) {
				return priority
			}
		}
	}

	// Fallback (nunca deve acontecer)
	return models.RolePatient
}
