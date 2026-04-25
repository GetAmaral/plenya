// Package handlers — DoctorAbsenceHandler expõe CRUD direto de doctor_absences
// pra UI Calendar V1 (Bloco G frontend).
//
// RBAC mesma regra de WorkingHoursHandler (admin/manager OU próprio doctor).
package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type DoctorAbsenceHandler struct {
	db *gorm.DB
}

func NewDoctorAbsenceHandler(db *gorm.DB) *DoctorAbsenceHandler {
	return &DoctorAbsenceHandler{db: db}
}

type absencePayload struct {
	StartAt string `json:"startAt"` // RFC3339
	EndAt   string `json:"endAt"`   // RFC3339
	Reason  string `json:"reason"`
}

// List GET /api/v1/doctors/:doctorId/absences
func (h *DoctorAbsenceHandler) List(c *fiber.Ctx) error {
	if !canReadDoctor(c) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	var rows []models.DoctorAbsence
	if err := h.db.WithContext(c.Context()).
		Where("doctor_id = ?", doctorID).
		Order("start_at ASC").
		Find(&rows).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// Create POST /api/v1/doctors/:doctorId/absences
func (h *DoctorAbsenceHandler) Create(c *fiber.Ctx) error {
	doctorID, err := parseDoctorID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid doctorId"})
	}
	if !canManageDoctor(c, doctorID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden"})
	}
	var body absencePayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	startAt, err := time.Parse(time.RFC3339, body.StartAt)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "startAt must be RFC3339"})
	}
	endAt, err := time.Parse(time.RFC3339, body.EndAt)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "endAt must be RFC3339"})
	}
	row := models.DoctorAbsence{
		DoctorID: doctorID,
		StartAt:  startAt,
		EndAt:    endAt,
		Reason:   body.Reason,
	}
	if err := h.db.WithContext(c.Context()).Create(&row).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(row)
}

// Delete DELETE /api/v1/doctors/:doctorId/absences/:id
func (h *DoctorAbsenceHandler) Delete(c *fiber.Ctx) error {
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
		Delete(&models.DoctorAbsence{})
	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": res.Error.Error()})
	}
	if res.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
