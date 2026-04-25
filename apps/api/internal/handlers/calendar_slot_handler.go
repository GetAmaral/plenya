// Package handlers — CalendarSlotHandler expõe o slot picker
// (Calendar V1, Bloco D).
package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type CalendarSlotHandler struct {
	slotSvc *services.CalendarSlotService
}

func NewCalendarSlotHandler(slotSvc *services.CalendarSlotService) *CalendarSlotHandler {
	return &CalendarSlotHandler{slotSvc: slotSvc}
}

// List retorna slots disponíveis pro médico no dia especificado.
//
// GET /api/v1/calendar/slots?doctorId=UUID&date=2026-04-25&type=initial_assessment&durationMin=90
//
// Query params:
//   - doctorId   (required) — UUID do médico
//   - date       (required) — YYYY-MM-DD (interpretado no TZ America/Sao_Paulo)
//   - type       (optional) — AppointmentType; default follow_up
//   - durationMin (optional) — sobrescreve duração default do tipo
//
// Auth: admin/secretary/manager/doctor (RBAC aplicado pelo router).
//
// Resp 200:
//
//	{
//	  "doctorId": "...",
//	  "date": "2026-04-25",
//	  "type": "initial_assessment",
//	  "durationMin": 90,
//	  "slots": [{"startUtc": "...", "endUtc": "..."}, ...]
//	}
func (h *CalendarSlotHandler) List(c *fiber.Ctx) error {
	doctorIDStr := c.Query("doctorId")
	if doctorIDStr == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "validation_failed",
			"message": "doctorId is required",
		})
	}
	doctorID, err := uuid.Parse(doctorIDStr)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "validation_failed",
			"message": "doctorId must be a valid UUID",
		})
	}

	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "validation_failed",
			"message": "date is required (YYYY-MM-DD)",
		})
	}
	// Aceita YYYY-MM-DD ou full RFC3339. Layout "2006-01-02" parseia em UTC,
	// mas o service vai re-localizar pro TZ correto.
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   "validation_failed",
				"message": "date must be YYYY-MM-DD or RFC3339",
			})
		}
	}

	apptType := models.AppointmentType(c.Query("type", string(models.AppointmentFollowUp)))
	// Validação branda — service usa fallback default.

	var overrideDuration *int
	if d := c.Query("durationMin"); d != "" {
		dur, err := strconv.Atoi(d)
		if err != nil || dur <= 0 || dur > 480 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   "validation_failed",
				"message": "durationMin must be 1..480",
			})
		}
		overrideDuration = &dur
	}

	slots, err := h.slotSvc.ListAvailable(c.Context(), services.ListSlotsParams{
		DoctorID:         doctorID,
		Date:             date,
		AppointmentType:  apptType,
		OverrideDuration: overrideDuration,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "internal_error",
			"message": err.Error(),
		})
	}

	durationMin := 0
	if overrideDuration != nil {
		durationMin = *overrideDuration
	}

	return c.JSON(fiber.Map{
		"doctorId":    doctorID,
		"date":        dateStr,
		"type":        apptType,
		"durationMin": durationMin,
		"slots":       slots,
	})
}
