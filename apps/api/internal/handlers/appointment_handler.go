package handlers

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type AppointmentHandler struct {
	appointmentService *services.AppointmentService
	validator          *validator.Validate
}

func NewAppointmentHandler(appointmentService *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
		validator:          validator.New(),
	}
}

// Create cria uma nova consulta
func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	resp, err := h.appointmentService.Create(&req)
	if err != nil {
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "patient not found",
				Message: "no patient found with the given id",
			})
		}
		if errors.Is(err, services.ErrAppointmentConflict) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "appointment conflict",
				Message: "doctor already has an appointment at this time",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetByID busca uma consulta por ID
func (h *AppointmentHandler) GetByID(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.appointmentService.GetByID(appointmentID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "appointment not found",
				Message: "no appointment found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to view this appointment",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista consultas
func (h *AppointmentHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	// Parse optional filters
	var patientID, doctorID *uuid.UUID
	var status *models.AppointmentStatus

	if patientIDStr := c.Query("patientId"); patientIDStr != "" {
		pid, err := uuid.Parse(patientIDStr)
		if err == nil {
			patientID = &pid
		}
	}
	if doctorIDStr := c.Query("doctorId"); doctorIDStr != "" {
		did, err := uuid.Parse(doctorIDStr)
		if err == nil {
			doctorID = &did
		}
	}
	if statusStr := c.Query("status"); statusStr != "" {
		s := models.AppointmentStatus(statusStr)
		status = &s
	}

	resp, err := h.appointmentService.List(userID, userRole, patientID, doctorID, status, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to list appointments",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update atualiza uma consulta
func (h *AppointmentHandler) Update(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}

	var req dto.UpdateAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.appointmentService.Update(appointmentID, userID, userRole, &req)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "appointment not found",
				Message: "no appointment found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to update this appointment",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Cancel cancela uma consulta
func (h *AppointmentHandler) Cancel(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}

	var req dto.CancelAppointmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.appointmentService.Cancel(appointmentID, userID, userRole, &req)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "appointment not found",
				Message: "no appointment found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to cancel this appointment",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete deleta uma consulta
func (h *AppointmentHandler) Delete(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}

	userRole := middleware.GetUserRole(c)

	err = h.appointmentService.Delete(appointmentID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "appointment not found",
				Message: "no appointment found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "only admins can delete appointments",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
