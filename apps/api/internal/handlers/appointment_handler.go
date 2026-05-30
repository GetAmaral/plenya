package handlers

import (
	"errors"
	"strconv"
	"strings"
	"time"

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

	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)
	resp, err := h.appointmentService.Create(userID, userRole, &req)
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
	userRole := middleware.GetPrimaryRole(c)

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
	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	params := services.ListAppointmentsParams{
		Limit:  limit,
		Offset: offset,
	}

	if patientIDStr := c.Query("patientId"); patientIDStr != "" {
		if pid, err := uuid.Parse(patientIDStr); err == nil {
			params.PatientID = &pid
		}
	}
	// doctorIds aceita CSV pra calendário multi-médico (ex: ?doctorIds=uuid1,uuid2).
	// Mantém compat com `doctorId` singular pra clients antigos.
	if doctorIdsStr := c.Query("doctorIds"); doctorIdsStr != "" {
		for _, s := range strings.Split(doctorIdsStr, ",") {
			if s = strings.TrimSpace(s); s == "" {
				continue
			}
			if id, err := uuid.Parse(s); err == nil {
				params.DoctorIDs = append(params.DoctorIDs, id)
			}
		}
	} else if doctorIDStr := c.Query("doctorId"); doctorIDStr != "" {
		if did, err := uuid.Parse(doctorIDStr); err == nil {
			params.DoctorIDs = []uuid.UUID{did}
		}
	}
	if statusStr := c.Query("status"); statusStr != "" {
		s := models.AppointmentStatus(statusStr)
		params.Status = &s
	}
	if dateFromStr := c.Query("dateFrom"); dateFromStr != "" {
		if t, err := time.Parse(time.RFC3339, dateFromStr); err == nil {
			params.DateFrom = &t
		}
	}
	if dateToStr := c.Query("dateTo"); dateToStr != "" {
		if t, err := time.Parse(time.RFC3339, dateToStr); err == nil {
			params.DateTo = &t
		}
	}

	resp, err := h.appointmentService.List(userID, userRole, params)
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
	userRole := middleware.GetPrimaryRole(c)

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
	userRole := middleware.GetPrimaryRole(c)

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

// Confirm confirma manualmente uma consulta (status -> confirmed).
// POST /api/v1/appointments/:id/confirm
// Auth: admin/secretary/manager OU o doctor da consulta.
func (h *AppointmentHandler) Confirm(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}
	userID := middleware.GetUserID(c)
	if err := h.appointmentService.Confirm(c.Context(), appointmentID, &userID); err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "appointment not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// CheckIn registra a chegada do paciente (status -> checked_in).
// POST /api/v1/appointments/:id/check-in
// Auth: RequireAnyStaff (recepção/médico).
func (h *AppointmentHandler) CheckIn(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}
	userID := middleware.GetUserID(c)
	resp, err := h.appointmentService.CheckIn(c.Context(), appointmentID, userID)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "appointment not found"})
		}
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
			Error:   "invalid transition",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// StartConsultation marca o início do atendimento (status -> in_progress).
// POST /api/v1/appointments/:id/start
// Auth: RequireAnyStaff (médico/recepção).
func (h *AppointmentHandler) StartConsultation(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}
	userID := middleware.GetUserID(c)
	resp, err := h.appointmentService.StartConsultation(c.Context(), appointmentID, userID)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "appointment not found"})
		}
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
			Error:   "invalid transition",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// GetTelemedToken POST /api/v1/appointments/:id/telemed-token
//
// HIGH H9 — gera meeting_token de owner pro staff (médico/secretaria) e
// devolve {joinUrl}. Sala Daily.co é privacy=private; URL crua não funciona.
//
// Não exposto como GET pra evitar caching de token. Token é fresh a cada call.
//
// Auth: RequireAnyStaff (já no router). Doctor só acessa suas consultas.
func (h *AppointmentHandler) GetTelemedToken(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid appointment id",
			Message: "appointment id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	joinURL, err := h.appointmentService.GetTelemedJoinURL(c.Context(), appointmentID, userID, userRole)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrAppointmentNotFound):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "appointment not found"})
		case errors.Is(err, services.ErrUnauthorized):
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{Error: "unauthorized"})
		case errors.Is(err, services.ErrAppointmentNotTelemed):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "not telemedicine",
			})
		case errors.Is(err, services.ErrAppointmentRoomNotReady):
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "room not ready",
				Message: "a sala está sendo provisionada — tente em alguns segundos",
			})
		case errors.Is(err, services.ErrAppointmentOutsideWindow):
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "outside window",
				Message: "a sala só pode ser acessada a partir de 3h antes até 3h após o horário agendado",
			})
		case errors.Is(err, services.ErrDailyNotConfigured):
			return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{
				Error: "daily.co not configured",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
				Error:   "internal server error",
				Message: err.Error(),
			})
		}
	}

	return c.JSON(fiber.Map{"joinUrl": joinURL})
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

	userRole := middleware.GetPrimaryRole(c)

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
