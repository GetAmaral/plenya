// Package handlers — MedicalRecordAggregateHandler expõe GET /patients/:id/prontuario
// com query agregada UNION sobre 8 modelos de evento.
package handlers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/services"
)

type MedicalRecordAggregateHandler struct {
	svc *services.MedicalRecordAggregateService
}

func NewMedicalRecordAggregateHandler(svc *services.MedicalRecordAggregateService) *MedicalRecordAggregateHandler {
	return &MedicalRecordAggregateHandler{svc: svc}
}

// List GET /api/v1/patients/:patientId/prontuario?types=anamnesis,lab_batch&authors=uuid1,uuid2&from=ISO&to=ISO&limit=&offset=
func (h *MedicalRecordAggregateHandler) List(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patientId"})
	}
	filter := services.AggregateFilter{
		Limit:  c.QueryInt("limit", 50),
		Offset: c.QueryInt("offset", 0),
	}
	if raw := c.Query("types"); raw != "" {
		for _, t := range strings.Split(raw, ",") {
			if t = strings.TrimSpace(t); t != "" {
				filter.Types = append(filter.Types, t)
			}
		}
	}
	if raw := c.Query("authors"); raw != "" {
		for _, a := range strings.Split(raw, ",") {
			if a = strings.TrimSpace(a); a != "" {
				if id, err := uuid.Parse(a); err == nil {
					filter.AuthorIDs = append(filter.AuthorIDs, id)
				}
			}
		}
	}
	if v := c.Query("from"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			filter.From = &t
		}
	}
	if v := c.Query("to"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			filter.To = &t
		}
	}
	rows, err := h.svc.List(patientID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}
