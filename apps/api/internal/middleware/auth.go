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
	UserID string           `json:"userId"`
	Email  string           `json:"email"`
	Role   models.UserRole  `json:"role"`
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
		c.Locals("userRole", claims.Role)

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

// GetUserRole retorna a role do usuário do contexto
func GetUserRole(c *fiber.Ctx) models.UserRole {
	return c.Locals("userRole").(models.UserRole)
}
