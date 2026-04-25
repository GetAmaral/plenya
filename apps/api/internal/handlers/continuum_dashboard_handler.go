// Package handlers — ContinuumDashboardHandler expõe as 3 visões do panorama
// da equipe: per-patient, per-week, alerts.
package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/services"
)

type ContinuumDashboardHandler struct {
	svc *services.ContinuumDashboardService
}

func NewContinuumDashboardHandler(svc *services.ContinuumDashboardService) *ContinuumDashboardHandler {
	return &ContinuumDashboardHandler{svc: svc}
}

// PerPatient GET /api/v1/continuum/dashboard/patients
func (h *ContinuumDashboardHandler) PerPatient(c *fiber.Ctx) error {
	rows, err := h.svc.PerPatient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// PerWeek GET /api/v1/continuum/dashboard/week?start=YYYY-MM-DD
func (h *ContinuumDashboardHandler) PerWeek(c *fiber.Ctx) error {
	weekStart := time.Now().UTC().Truncate(24 * time.Hour)
	if v := c.Query("start"); v != "" {
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid start (use YYYY-MM-DD)"})
		}
		weekStart = t.UTC()
	} else {
		// Domingo da semana atual.
		offset := int(weekStart.Weekday())
		weekStart = weekStart.AddDate(0, 0, -offset)
	}
	rows, err := h.svc.PerWeek(weekStart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"weekStart": weekStart.Format("2006-01-02"),
		"items":     rows,
	})
}

// Alerts GET /api/v1/continuum/dashboard/alerts?dueSoonDays=7
func (h *ContinuumDashboardHandler) Alerts(c *fiber.Ctx) error {
	rows, err := h.svc.Alerts(c.QueryInt("dueSoonDays", 7))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}
