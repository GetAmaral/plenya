package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

// TelemedRecordingHandler expõe o status (gravação + transcrição) de uma
// teleconsulta e o link de download sob demanda do MP4.
type TelemedRecordingHandler struct {
	svc *services.TelemedRecordingService
}

func NewTelemedRecordingHandler(svc *services.TelemedRecordingService) *TelemedRecordingHandler {
	return &TelemedRecordingHandler{svc: svc}
}

// GetForAppointment — GET /api/v1/appointments/:id/telemed-recording
func (h *TelemedRecordingHandler) GetForAppointment(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid appointment id"})
	}
	resp, err := h.svc.GetByAppointment(c.Context(), appointmentID)
	if err != nil {
		if errors.Is(err, services.ErrTelemedRecordingNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "no recording"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Download — GET /api/v1/appointments/:id/telemed-recording/download
// Gera link assinado temporário do MP4 (a gravação fica no storage do Daily).
func (h *TelemedRecordingHandler) Download(c *fiber.Ctx) error {
	appointmentID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid appointment id"})
	}
	url, err := h.svc.RecordingDownloadLink(c.Context(), appointmentID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrTelemedRecordingNotFound), errors.Is(err, services.ErrTelemedNoRecording):
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "no recording available"})
		case errors.Is(err, services.ErrDailyNotConfigured):
			return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{Error: "daily.co not configured"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
		}
	}
	return c.JSON(dto.TelemedRecordingDownloadResponse{DownloadURL: url})
}
