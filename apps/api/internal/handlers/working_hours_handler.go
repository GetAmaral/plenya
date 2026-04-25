// Package handlers — WorkingHoursHandler expõe CRUD direto da tabela
// working_hours pra UI Calendar V1 (Bloco G frontend).
//
// RBAC:
//   - GET: admin/secretary/manager/doctor (qualquer staff pode ler agenda).
//   - POST/PUT/DELETE: admin/manager OU o próprio doctor (dono dos blocos).
//
// Os endpoints CRUD são "burros" — apenas persistem o que a UI envia.
// A lógica do slot picker (intersecção working_hours ∩ ¬absence ∩ etc) está
// em CalendarSlotService e não passa por aqui.
package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
)

type WorkingHoursHandler struct {
	db *gorm.DB
}

func NewWorkingHoursHandler(db *gorm.DB) *WorkingHoursHandler {
	return &WorkingHoursHandler{db: db}
}

// canManageDoctor retorna true se o caller pode gerenciar configs do doctorID:
// admin/manager (sempre) ou o próprio doctor.
func canManageDoctor(c *fiber.Ctx, doctorID uuid.UUID) bool {
	roles := middleware.GetUserRoles(c)
	for _, r := range roles {
		if r == string(models.RoleAdmin) || r == string(models.RoleManager) {
			return true
		}
	}
	return middleware.GetUserID(c) == doctorID
}

// canReadDoctor retorna true pra qualquer staff (admin/secretary/manager/doctor).
// Doctor pode ver outros doctors em modo read-only.
func canReadDoctor(c *fiber.Ctx) bool {
	roles := middleware.GetUserRoles(c)
	for _, r := range roles {
		switch r {
		case string(models.RoleAdmin),
			string(models.RoleSecretary),
			string(models.RoleManager),
			string(models.RoleDoctor),
			string(models.RoleNurse):
			return true
		}
	}
	return false
}

type workingHoursPayload struct {
	Weekday      int `json:"weekday"`
	StartMinute  int `json:"startMinute"`
	EndMinute    int `json:"endMinute"`
	SlotDuration int `json:"slotDuration"`
}

func parseDoctorID(c *fiber.Ctx) (uuid.UUID, error) {
	return uuid.Parse(c.Params("doctorId"))
}

// List GET /api/v1/doctors/:doctorId/working-hours
func (h *WorkingHoursHandler) List(c *fiber.Ctx) error {
	if !canReadDoctor(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	var rows []models.WorkingHours
	if err := h.db.WithContext(c.Context()).
		Where("doctor_id = ?", doctorID).
		Order("weekday ASC, start_minute ASC").
		Find(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// Create POST /api/v1/doctors/:doctorId/working-hours
func (h *WorkingHoursHandler) Create(c *fiber.Ctx) error {
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	if !canManageDoctor(c, doctorID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	var body workingHoursPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if body.SlotDuration == 0 {
		body.SlotDuration = 30
	}
	row := models.WorkingHours{
		DoctorID:     doctorID,
		Weekday:      body.Weekday,
		StartMinute:  body.StartMinute,
		EndMinute:    body.EndMinute,
		SlotDuration: body.SlotDuration,
	}
	if err := h.db.WithContext(c.Context()).Create(&row).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(row)
}

// Update PUT /api/v1/doctors/:doctorId/working-hours/:id
func (h *WorkingHoursHandler) Update(c *fiber.Ctx) error {
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	if !canManageDoctor(c, doctorID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var row models.WorkingHours
	if err := h.db.WithContext(c.Context()).
		Where("id = ? AND doctor_id = ?", id, doctorID).First(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var body workingHoursPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if body.SlotDuration == 0 {
		body.SlotDuration = 30
	}
	row.Weekday = body.Weekday
	row.StartMinute = body.StartMinute
	row.EndMinute = body.EndMinute
	row.SlotDuration = body.SlotDuration
	if err := h.db.WithContext(c.Context()).Save(&row).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(row)
}

// Delete DELETE /api/v1/doctors/:doctorId/working-hours/:id
func (h *WorkingHoursHandler) Delete(c *fiber.Ctx) error {
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	if !canManageDoctor(c, doctorID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	res := h.db.WithContext(c.Context()).
		Where("id = ? AND doctor_id = ?", id, doctorID).
		Delete(&models.WorkingHours{})
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": res.Error.Error()})
	}
	if res.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
