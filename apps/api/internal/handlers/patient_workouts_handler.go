// Package handlers — métodos de PatientPortalHandler que cobrem a área de
// treino + check-in + preferências de notificação. Mantidos em arquivo separado
// só por organização; struct e construtor estão em patient_portal_handler.go.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// ListWorkoutPlans GET /api/v1/patient/me/workout-plans
func (h *PatientPortalHandler) ListWorkoutPlans(c *fiber.Ctx) error {
	rows, err := h.workouts.ListPlans(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// GetWorkoutPlan GET /api/v1/patient/me/workout-plans/:id
func (h *PatientPortalHandler) GetWorkoutPlan(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	plan, err := h.workouts.GetPlan(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(plan)
}

// ListWorkoutSessions GET /api/v1/patient/me/workout-sessions
func (h *PatientPortalHandler) ListWorkoutSessions(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 30)
	rows, err := h.workouts.ListSessions(middleware.GetPatientID(c), limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// GetWorkoutSession GET /api/v1/patient/me/workout-sessions/:id
func (h *PatientPortalHandler) GetWorkoutSession(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	sess, err := h.workouts.GetSession(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(sess)
}

// StartWorkoutSession POST /api/v1/patient/me/workout-sessions
func (h *PatientPortalHandler) StartWorkoutSession(c *fiber.Ctx) error {
	var in services.CreateSessionInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "json inválido"})
	}
	sess, err := h.workouts.StartSession(middleware.GetPatientID(c), in)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(sess)
}

// LogWorkoutSet POST /api/v1/patient/me/workout-sessions/:id/logs
func (h *PatientPortalHandler) LogWorkoutSet(c *fiber.Ctx) error {
	sessionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	var in services.LogSetInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "json inválido"})
	}
	log, err := h.workouts.LogSet(middleware.GetPatientID(c), sessionID, in)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(log)
}

// CompleteWorkoutSession POST /api/v1/patient/me/workout-sessions/:id/complete
func (h *PatientPortalHandler) CompleteWorkoutSession(c *fiber.Ctx) error {
	sessionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	var body struct {
		Notes *string `json:"notes,omitempty"`
	}
	_ = c.BodyParser(&body)
	sess, err := h.workouts.CompleteSession(middleware.GetPatientID(c), sessionID, body.Notes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(sess)
}

// ListCheckIns GET /api/v1/patient/me/check-ins
func (h *PatientPortalHandler) ListCheckIns(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 30)
	rows, err := h.checkIns.List(middleware.GetPatientID(c), limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// LatestCheckInToday GET /api/v1/patient/me/check-ins/today
func (h *PatientPortalHandler) LatestCheckInToday(c *fiber.Ctx) error {
	row, err := h.checkIns.LatestToday(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if row == nil {
		return c.JSON(fiber.Map{"checkIn": nil})
	}
	return c.JSON(fiber.Map{"checkIn": row})
}

// CreateCheckIn POST /api/v1/patient/me/check-ins
func (h *PatientPortalHandler) CreateCheckIn(c *fiber.Ctx) error {
	var in services.CreateCheckInInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "json inválido"})
	}
	row, err := h.checkIns.Create(middleware.GetPatientID(c), in)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(row)
}

// GetNotificationPreferences GET /api/v1/patient/me/notification-preferences
// também útil pro app pro (mesmo handler, escopo userID via Auth)
func (h *PatientPortalHandler) GetNotificationPreferences(c *fiber.Ctx) error {
	prefs, err := h.notifPrefs.Get(middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(prefs)
}

// StaffListCheckIns GET /api/v1/patients/:id/check-ins (escopo staff)
func (h *PatientPortalHandler) StaffListCheckIns(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	limit := c.QueryInt("limit", 90)
	rows, err := h.checkIns.ListForStaff(id, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// PatchNotificationPreferences PATCH /api/v1/patient/me/notification-preferences
func (h *PatientPortalHandler) PatchNotificationPreferences(c *fiber.Ctx) error {
	var in services.PatchInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "json inválido"})
	}
	prefs, err := h.notifPrefs.Patch(middleware.GetUserID(c), in)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(prefs)
}
