package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// LabRequestHandler handles HTTP requests for lab requests
type LabRequestHandler struct {
	service *services.LabRequestService
}

// NewLabRequestHandler creates a new lab request handler
func NewLabRequestHandler(service *services.LabRequestService) *LabRequestHandler {
	return &LabRequestHandler{service: service}
}

// CreateLabRequest creates a new lab request
// @Summary Create lab request
// @Tags LabRequests
// @Accept json
// @Produce json
// @Param request body dto.CreateLabRequestRequest true "Lab request data"
// @Success 201 {object} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests [post]
func (h *LabRequestHandler) CreateLabRequest(c *fiber.Ctx) error {
	var reqDTO dto.CreateLabRequestRequest
	if err := c.BodyParser(&reqDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Parse date from string to time.Time
	date, err := time.Parse("2006-01-02", reqDTO.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid date format. Expected YYYY-MM-DD",
		})
	}

	// Get user ID from JWT (middleware.AuthMiddleware sets this)
	userID := middleware.GetUserID(c)

	// Build lab request model
	req := &models.LabRequest{
		Date:  date,
		Exams: reqDTO.Exams,
		Notes: reqDTO.Notes,
	}

	// Parse template ID if provided
	if reqDTO.LabRequestTemplateID != nil && *reqDTO.LabRequestTemplateID != "" {
		templateID, err := uuid.Parse(*reqDTO.LabRequestTemplateID)
		if err == nil {
			req.LabRequestTemplateID = &templateID
		}
	}

	// Create lab request (service will auto-fill patientID from selectedPatient)
	if err := h.service.CreateLabRequest(userID, req); err != nil {
		// Check for specific errors
		if err.Error() == "no patient selected - please select a patient first" {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Nenhum paciente selecionado. Por favor, selecione um paciente primeiro.",
			})
		}
		if err.Error() == "patient id does not match selected patient" {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error: "O paciente especificado não corresponde ao paciente selecionado.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(req)
}

// GetLabRequestByID retrieves a lab request by ID
// @Summary Get lab request by ID
// @Tags LabRequests
// @Produce json
// @Param id path string true "Lab request ID"
// @Success 200 {object} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/{id} [get]
func (h *LabRequestHandler) GetLabRequestByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	req, err := h.service.GetLabRequestByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(req)
}

// GetLabRequestsByPatientID retrieves all lab requests for a patient
// @Summary Get lab requests by patient ID
// @Tags LabRequests
// @Produce json
// @Param patientId path string true "Patient ID"
// @Success 200 {array} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{patientId}/lab-requests [get]
func (h *LabRequestHandler) GetLabRequestsByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid patient ID format",
		})
	}

	reqs, err := h.service.GetLabRequestsByPatientID(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(reqs)
}

// GetLabRequestsByDate retrieves all lab requests for a specific date
// @Summary Get lab requests by date
// @Tags LabRequests
// @Produce json
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {array} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/by-date [get]
func (h *LabRequestHandler) GetLabRequestsByDate(c *fiber.Ctx) error {
	dateStr := c.Query("date")
	if dateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Date parameter is required",
		})
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid date format. Use YYYY-MM-DD",
		})
	}

	reqs, err := h.service.GetLabRequestsByDate(date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(reqs)
}

// GetLabRequestsByDateRange retrieves lab requests within a date range
// @Summary Get lab requests by date range
// @Tags LabRequests
// @Produce json
// @Param startDate query string true "Start date (YYYY-MM-DD)"
// @Param endDate query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/by-date-range [get]
func (h *LabRequestHandler) GetLabRequestsByDateRange(c *fiber.Ctx) error {
	startDateStr := c.Query("startDate")
	endDateStr := c.Query("endDate")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Both startDate and endDate parameters are required",
		})
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid startDate format. Use YYYY-MM-DD",
		})
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid endDate format. Use YYYY-MM-DD",
		})
	}

	reqs, err := h.service.GetLabRequestsByDateRange(startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(reqs)
}

// GetAllLabRequests retrieves all lab requests with pagination
// @Summary Get all lab requests
// @Tags LabRequests
// @Produce json
// @Param limit query int false "Limit" default(50)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} dto.PaginatedResponse{data=[]models.LabRequest}
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests [get]
func (h *LabRequestHandler) GetAllLabRequests(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 50)
	offset := c.QueryInt("offset", 0)

	reqs, total, err := h.service.GetAllLabRequests(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(dto.PaginatedResponse{
		Data:   reqs,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	})
}

// UpdateLabRequest updates an existing lab request
// @Summary Update lab request
// @Tags LabRequests
// @Accept json
// @Produce json
// @Param id path string true "Lab request ID"
// @Param request body dto.UpdateLabRequestRequest true "Updated lab request data"
// @Success 200 {object} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/{id} [put]
func (h *LabRequestHandler) UpdateLabRequest(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var reqDTO dto.UpdateLabRequestRequest
	if err := c.BodyParser(&reqDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Get existing lab request
	existing, err := h.service.GetLabRequestByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: "Lab request not found",
		})
	}

	// Block editing if PDF was already generated
	if existing.PdfURL != nil && *existing.PdfURL != "" {
		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Error: "Não é possível editar um pedido que já teve PDF gerado",
		})
	}

	// Update fields if provided
	if reqDTO.Date != nil {
		date, err := time.Parse("2006-01-02", *reqDTO.Date)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "Invalid date format. Expected YYYY-MM-DD",
			})
		}
		existing.Date = date
	}

	if reqDTO.Exams != nil {
		existing.Exams = *reqDTO.Exams
	}

	if reqDTO.Notes != nil {
		existing.Notes = reqDTO.Notes
	}

	if reqDTO.LabRequestTemplateID != nil {
		if *reqDTO.LabRequestTemplateID == "" {
			// Clear template
			existing.LabRequestTemplateID = nil
		} else {
			// Update template
			templateID, err := uuid.Parse(*reqDTO.LabRequestTemplateID)
			if err == nil {
				existing.LabRequestTemplateID = &templateID
			}
		}
	}

	if err := h.service.UpdateLabRequest(id, existing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, err := h.service.GetLabRequestByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}

// DeleteLabRequest soft deletes a lab request
// @Summary Delete lab request
// @Tags LabRequests
// @Param id path string true "Lab request ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/{id} [delete]
func (h *LabRequestHandler) DeleteLabRequest(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := h.service.DeleteLabRequest(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GeneratePDF generates a PDF for the lab request
// @Summary Generate PDF for lab request
// @Tags LabRequests
// @Produce json
// @Param id path string true "Lab request ID"
// @Success 200 {object} models.LabRequest
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-requests/{id}/generate-pdf [post]
func (h *LabRequestHandler) GeneratePDF(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	// Get lab request
	req, err := h.service.GetLabRequestByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: "Lab request not found",
		})
	}

	// Check if PDF already exists
	if req.PdfURL != nil && *req.PdfURL != "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "PDF já foi gerado para este pedido",
		})
	}

	// Generate PDF
	pdfURL, err := h.service.GenerateLabRequestPDF(id)
	if err != nil {
		// Log the error for debugging
		fmt.Printf("[ERROR] Failed to generate PDF for lab request %s: %v\n", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: fmt.Sprintf("Erro ao gerar PDF: %v", err),
		})
	}

	// Update lab request with PDF URL
	req.PdfURL = &pdfURL
	if err := h.service.UpdateLabRequest(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Return updated request
	updated, err := h.service.GetLabRequestByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}
