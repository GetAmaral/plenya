package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// RequirePatient é o middleware da área do paciente (minha.plenyasaude.com.br).
//
// Diferente do RequireRole(patient): além de validar que o JWT tem role=patient,
// resolve o Patient vinculado ao UserID e injeta em c.Locals("patient").
// Falha 403 se o User não tiver Patient vinculado (caso impossível em produção,
// mas defensivo).
//
// Use depois de Auth() pra que userID esteja em c.Locals.
func RequirePatient(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Confirma role
		if !HasRole(c, models.RolePatient) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "esta área é exclusiva para pacientes",
			})
		}

		// 2. Resolve patient pelo userID
		userID, ok := c.Locals("userID").(uuid.UUID)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user not authenticated",
			})
		}

		var patient models.Patient
		if err := db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "nenhum paciente vinculado a este usuário",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to load patient",
			})
		}

		c.Locals("patient", &patient)
		c.Locals("patientID", patient.ID)
		return c.Next()
	}
}

// GetPatient retorna o Patient resolvido pelo middleware RequirePatient.
// Panic se chamado sem o middleware.
func GetPatient(c *fiber.Ctx) *models.Patient {
	return c.Locals("patient").(*models.Patient)
}

// GetPatientID retorna o ID do paciente resolvido pelo middleware RequirePatient.
func GetPatientID(c *fiber.Ctx) uuid.UUID {
	return c.Locals("patientID").(uuid.UUID)
}
