// Package handlers — PatientContinuumHandler expõe endpoints de inscrição
// de paciente em programa Continuum + listagem de items + update de status.
//
// RBAC:
//   - GET: qualquer staff (incluindo professionals) pode visualizar.
//   - POST/PUT/DELETE: admin/manager/secretary (ops da clínica). Médico não
//     inscreve paciente — secretária/coordenador faz.
package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type PatientContinuumHandler struct {
	svc *services.PatientContinuumService
}

func NewPatientContinuumHandler(svc *services.PatientContinuumService) *PatientContinuumHandler {
	return &PatientContinuumHandler{svc: svc}
}

// ListByPatient GET /api/v1/patients/:patientId/continuum
func (h *PatientContinuumHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patientId"})
	}
	rows, err := h.svc.GetByPatient(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// Get GET /api/v1/continuum/enrollments/:id
func (h *PatientContinuumHandler) Get(c *fiber.Ctx) error {
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

// Enroll POST /api/v1/patients/:patientId/continuum
func (h *PatientContinuumHandler) Enroll(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patientId"})
	}
	var body dto.EnrollPatientContinuumPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	templateID, err := uuid.Parse(body.TemplateID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid templateId"})
	}
	startDate, err := time.Parse("2006-01-02", body.StartDate)
	if err != nil {
		// Try RFC3339 as fallback.
		startDate, err = time.Parse(time.RFC3339, body.StartDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid startDate (use YYYY-MM-DD)"})
		}
	}
	var coordID *uuid.UUID
	if body.CoordinatorDoctorID != nil && *body.CoordinatorDoctorID != "" {
		cid, err := uuid.Parse(*body.CoordinatorDoctorID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid coordinatorDoctorId"})
		}
		coordID = &cid
	}
	row, err := h.svc.Enroll(services.EnrollParams{
		PatientID:           patientID,
		TemplateID:          templateID,
		StartDate:           startDate,
		CoordinatorDoctorID: coordID,
		Notes:               body.Notes,
	})
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(row)
}

// UpdateItem PUT /api/v1/continuum/items/:id
func (h *PatientContinuumHandler) UpdateItem(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var body dto.UpdateContinuumItemPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	patch := services.UpdateItemPatch{}
	if body.Status != nil {
		s := models.ContinuumItemStatus(*body.Status)
		patch.Status = &s
	}
	if body.AppointmentID != nil && *body.AppointmentID != "" {
		aid, err := uuid.Parse(*body.AppointmentID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid appointmentId"})
		}
		patch.AppointmentID = &aid
	}
	if body.CompletedAt != nil && *body.CompletedAt != "" {
		t, err := time.Parse(time.RFC3339, *body.CompletedAt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid completedAt (use RFC3339)"})
		}
		patch.CompletedAt = &t
	}
	if body.CompletedRefType != nil {
		patch.CompletedRefType = body.CompletedRefType
	}
	if body.CompletedRefID != nil && *body.CompletedRefID != "" {
		rid, err := uuid.Parse(*body.CompletedRefID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid completedRefId"})
		}
		patch.CompletedRefID = &rid
	}
	row, err := h.svc.UpdateItem(itemID, patch)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(row)
}

// Cancel DELETE /api/v1/continuum/enrollments/:id
func (h *PatientContinuumHandler) Cancel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.Cancel(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateIntegratedPlan PUT /api/v1/continuum/enrollments/:id/integrated-plan
func (h *PatientContinuumHandler) UpdateIntegratedPlan(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var body dto.UpdateIntegratedPlanPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	updated, err := h.svc.UpdateIntegratedPlan(id, middleware.GetUserID(c), body.Content)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updated)
}

// ListPlanRevisions GET /api/v1/continuum/enrollments/:id/integrated-plan/revisions
func (h *PatientContinuumHandler) ListPlanRevisions(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	rows, err := h.svc.ListPlanRevisions(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}
