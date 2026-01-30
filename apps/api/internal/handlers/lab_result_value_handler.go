package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// LabResultValueHandler handles HTTP requests for lab result values
type LabResultValueHandler struct {
	service *services.LabResultValueService
}

// NewLabResultValueHandler creates a new lab result value handler
func NewLabResultValueHandler(service *services.LabResultValueService) *LabResultValueHandler {
	return &LabResultValueHandler{service: service}
}

// CreateLabResultValue creates a new lab result value
// @Summary Create lab result value
// @Tags LabResults
// @Accept json
// @Produce json
// @Param value body models.LabResultValue true "Lab result value data"
// @Success 201 {object} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/values [post]
func (h *LabResultValueHandler) CreateLabResultValue(c *fiber.Ctx) error {
	var value models.LabResultValue
	if err := c.BodyParser(&value); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.CreateLabResultValue(&value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(value)
}

// CreateLabResultValues creates multiple lab result values in a transaction
// @Summary Create multiple lab result values
// @Tags LabResults
// @Accept json
// @Produce json
// @Param values body []models.LabResultValue true "Array of lab result values"
// @Success 201 {array} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/values/batch [post]
func (h *LabResultValueHandler) CreateLabResultValues(c *fiber.Ctx) error {
	var values []models.LabResultValue
	if err := c.BodyParser(&values); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.CreateLabResultValues(values); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(values)
}

// GetLabResultValueByID retrieves a lab result value by ID
// @Summary Get lab result value by ID
// @Tags LabResults
// @Produce json
// @Param id path string true "Lab result value ID"
// @Success 200 {object} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/values/{id} [get]
func (h *LabResultValueHandler) GetLabResultValueByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	value, err := h.service.GetLabResultValueByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(value)
}

// GetValuesByLabResult retrieves all values for a specific lab result
// @Summary Get values for lab result
// @Tags LabResults
// @Produce json
// @Param id path string true "Lab result ID"
// @Success 200 {array} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/{id}/values [get]
func (h *LabResultValueHandler) GetValuesByLabResult(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	values, err := h.service.GetValuesByLabResult(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(values)
}

// GetValuesByPatient retrieves all lab result values for a patient
// @Summary Get lab values for patient
// @Tags LabResults
// @Produce json
// @Param patientId path string true "Patient ID"
// @Success 200 {array} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{patientId}/lab-values [get]
func (h *LabResultValueHandler) GetValuesByPatient(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	values, err := h.service.GetValuesByPatient(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(values)
}

// GetLatestValueForTest retrieves the most recent value for a specific test and patient
// @Summary Get latest value for test
// @Tags LabResults
// @Produce json
// @Param patientId path string true "Patient ID"
// @Param testId path string true "Lab test definition ID"
// @Success 200 {object} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/patients/{patientId}/lab-values/test/{testId}/latest [get]
func (h *LabResultValueHandler) GetLatestValueForTest(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid patient ID format",
		})
	}

	testID, err := uuid.Parse(c.Params("testId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid test ID format",
		})
	}

	value, err := h.service.GetLatestValueForTest(patientID, testID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(value)
}

// UpdateLabResultValue updates an existing lab result value
// @Summary Update lab result value
// @Tags LabResults
// @Accept json
// @Produce json
// @Param id path string true "Lab result value ID"
// @Param value body models.LabResultValue true "Updated lab result value data"
// @Success 200 {object} models.LabResultValue
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/values/{id} [put]
func (h *LabResultValueHandler) UpdateLabResultValue(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var value models.LabResultValue
	if err := c.BodyParser(&value); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.UpdateLabResultValue(id, &value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, _ := h.service.GetLabResultValueByID(id)
	return c.JSON(updated)
}

// DeleteLabResultValue soft deletes a lab result value
// @Summary Delete lab result value
// @Tags LabResults
// @Param id path string true "Lab result value ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/lab-results/values/{id} [delete]
func (h *LabResultValueHandler) DeleteLabResultValue(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := h.service.DeleteLabResultValue(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
