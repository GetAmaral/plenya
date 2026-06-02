package handlers

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type ClinicalNoteHandler struct {
	service   *services.ClinicalNoteService
	validator *validator.Validate
}

func NewClinicalNoteHandler(service *services.ClinicalNoteService) *ClinicalNoteHandler {
	return &ClinicalNoteHandler{
		service:   service,
		validator: validator.New(),
	}
}

// mapErr converte erros do service em respostas HTTP.
func (h *ClinicalNoteHandler) mapErr(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, services.ErrClinicalNoteNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "clinical note not found"})
	case errors.Is(err, services.ErrClinicalNoteSigned):
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
			Error:   "clinical note signed",
			Message: "esta nota já foi assinada e não pode ser editada — crie um adendo",
		})
	case errors.Is(err, services.ErrClinicalNoteRestricted), errors.Is(err, services.ErrUnauthorized):
		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{Error: "forbidden", Message: "sem permissão para esta nota"})
	case errors.Is(err, services.ErrPatientNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "patient not found"})
	case errors.Is(err, services.ErrAppointmentNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "appointment not found"})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
}

// Create POST /api/v1/clinical-notes
func (h *ClinicalNoteHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateClinicalNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Create(middleware.GetUserID(c), &req)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// List GET /api/v1/clinical-notes?appointmentId=&limit=&offset=
// Se appointmentId for passado, retorna a nota daquela consulta (ou 404).
func (h *ClinicalNoteHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	if apptStr := c.Query("appointmentId"); apptStr != "" {
		apptID, err := uuid.Parse(apptStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid appointment id"})
		}
		resp, err := h.service.GetByAppointment(apptID, userID, userRole)
		if err != nil {
			return h.mapErr(c, err)
		}
		return c.JSON(resp)
	}

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}
	resp, err := h.service.ListByPatient(userID, userRole, limit, offset)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(resp)
}

// GetByID GET /api/v1/clinical-notes/:id
func (h *ClinicalNoteHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid clinical note id"})
	}
	resp, err := h.service.GetByID(id, middleware.GetUserID(c), middleware.GetPrimaryRole(c))
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(resp)
}

// Update PUT /api/v1/clinical-notes/:id
func (h *ClinicalNoteHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid clinical note id"})
	}
	var req dto.UpdateClinicalNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Update(id, middleware.GetUserID(c), middleware.GetPrimaryRole(c), &req)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(resp)
}

// Sign POST /api/v1/clinical-notes/:id/sign
func (h *ClinicalNoteHandler) Sign(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid clinical note id"})
	}
	resp, err := h.service.Sign(id, middleware.GetUserID(c), middleware.GetPrimaryRole(c))
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(resp)
}

// Amend POST /api/v1/clinical-notes/:id/amend
func (h *ClinicalNoteHandler) Amend(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid clinical note id"})
	}
	var req dto.CreateClinicalNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Amend(id, middleware.GetUserID(c), middleware.GetPrimaryRole(c), &req)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Delete DELETE /api/v1/clinical-notes/:id
func (h *ClinicalNoteHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid clinical note id"})
	}
	if err := h.service.Delete(id, middleware.GetUserID(c), middleware.GetPrimaryRole(c)); err != nil {
		return h.mapErr(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
