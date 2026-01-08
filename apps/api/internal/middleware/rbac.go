package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/models"
)

// RequireRole middleware verifica se o usuário tem uma das roles permitidas
func RequireRole(allowedRoles ...models.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := GetUserRole(c)

		// Verificar se a role do usuário está na lista de roles permitidas
		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "insufficient permissions",
		})
	}
}

// RequireAdmin middleware permite apenas admins
func RequireAdmin() fiber.Handler {
	return RequireRole(models.RoleAdmin)
}

// RequireDoctor middleware permite apenas doctors
func RequireDoctor() fiber.Handler {
	return RequireRole(models.RoleDoctor)
}

// RequireDoctorOrNurse middleware permite doctors e nurses
func RequireDoctorOrNurse() fiber.Handler {
	return RequireRole(models.RoleDoctor, models.RoleNurse)
}

// RequireMedicalStaff middleware permite admins, doctors e nurses
func RequireMedicalStaff() fiber.Handler {
	return RequireRole(models.RoleAdmin, models.RoleDoctor, models.RoleNurse)
}
