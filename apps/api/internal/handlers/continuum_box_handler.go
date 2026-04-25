// Package handlers — ContinuumBoxHandler expõe listagem cross-paciente de
// PatientContinuumBox + update de status/tracking pela equipe de logística.
package handlers

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type ContinuumBoxHandler struct {
	svc *services.ContinuumBoxService
}

func NewContinuumBoxHandler(svc *services.ContinuumBoxService) *ContinuumBoxHandler {
	return &ContinuumBoxHandler{svc: svc}
}

// List GET /api/v1/continuum/boxes?status=planned,preparing&limit=&offset=
func (h *ContinuumBoxHandler) List(c *fiber.Ctx) error {
	filter := services.BoxFilter{
		Limit:  c.QueryInt("limit", 100),
		Offset: c.QueryInt("offset", 0),
	}
	if raw := c.Query("status"); raw != "" {
		for _, s := range strings.Split(raw, ",") {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			filter.Statuses = append(filter.Statuses, models.PatientContinuumBoxStatus(s))
		}
	}
	rows, err := h.svc.List(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// Get GET /api/v1/continuum/boxes/:id
func (h *ContinuumBoxHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	row, err := h.svc.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(row)
}

// Update PUT /api/v1/continuum/boxes/:id
func (h *ContinuumBoxHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var body dto.UpdateBoxPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	patch := services.UpdateBoxPatch{}
	if body.Status != nil {
		s := models.PatientContinuumBoxStatus(*body.Status)
		patch.Status = &s
	}
	patch.TrackingCode = body.TrackingCode
	patch.Carrier = body.Carrier
	patch.Notes = body.Notes
	patch.Contents = body.Contents
	patch.Address = body.Address

	row, err := h.svc.Update(id, patch)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(row)
}

// Counts GET /api/v1/continuum/boxes/counts
func (h *ContinuumBoxHandler) Counts(c *fiber.Ctx) error {
	out, err := h.svc.CountByStatus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}
