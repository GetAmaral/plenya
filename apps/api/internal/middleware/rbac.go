package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/models"
)

// RequireRole middleware verifica se o usuário tem pelo menos uma das roles permitidas
func RequireRole(allowedRoles ...models.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRoles := GetUserRoles(c)

		// Verificar se o usuário tem pelo menos uma das roles permitidas
		for _, userRole := range userRoles {
			for _, allowedRole := range allowedRoles {
				if userRole == string(allowedRole) {
					return c.Next()
				}
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "insufficient permissions",
		})
	}
}

// RequireAdmin middleware permite apenas usuários com role admin
func RequireAdmin() fiber.Handler {
	return RequireRole(models.RoleAdmin)
}

// RequireDoctor middleware permite apenas usuários com role doctor
func RequireDoctor() fiber.Handler {
	return RequireRole(models.RoleDoctor)
}

// RequireDoctorOrNurse middleware permite doctors e nurses
func RequireDoctorOrNurse() fiber.Handler {
	return RequireRole(models.RoleDoctor, models.RoleNurse)
}

// RequireMedicalStaff middleware permite admins, doctors e nurses (legado).
// Prefira RequireClinician/RequireAnyStaff/RequireDoctor pra granularidade certa.
func RequireMedicalStaff() fiber.Handler {
	return RequireRole(models.RoleAdmin, models.RoleDoctor, models.RoleNurse)
}

// RequireAnyStaff — qualquer membro da equipe da clínica (todos exceto patient).
// Usar como guard padrão em recursos compartilhados (ex: ler paciente, agenda).
// Decisão de produto: clínica única, todos staff veem todos os pacientes.
// Multi-tenant fica pra V2.
func RequireAnyStaff() fiber.Handler {
	return RequireRole(
		models.RoleAdmin,
		models.RoleManager,
		models.RoleSecretary,
		models.RoleDoctor,
		models.RoleNurse,
		models.RolePsychologist,
		models.RoleNutritionist,
		models.RolePhysicalEducator,
	)
}

// RequireClinician — profissionais clínicos que executam ações sobre paciente
// (criar/editar exame, anamnese, avaliação física). Exclui secretary/manager
// que são apenas operacionais.
func RequireClinician() fiber.Handler {
	return RequireRole(
		models.RoleAdmin,
		models.RoleDoctor,
		models.RoleNurse,
		models.RolePsychologist,
		models.RoleNutritionist,
		models.RolePhysicalEducator,
	)
}

// RequireAdminOps — operações administrativas (criar lead, agendar consulta,
// atribuir vendedor, gerenciar campanhas). Não inclui clínicos puros.
func RequireAdminOps() fiber.Handler {
	return RequireRole(models.RoleAdmin, models.RoleManager, models.RoleSecretary)
}
