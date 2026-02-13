package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type ScoreSnapshotHandler struct {
	snapshotService *services.ScoreSnapshotService
	validator       *validator.Validate
}

func NewScoreSnapshotHandler(snapshotService *services.ScoreSnapshotService) *ScoreSnapshotHandler {
	return &ScoreSnapshotHandler{
		snapshotService: snapshotService,
		validator:       validator.New(),
	}
}

// CalculateSnapshot calculates a new score snapshot for a patient
// @Summary Calculate patient score snapshot
// @Description Calculates a new health score snapshot based on patient's lab results and anamnesis history
// @Tags score-snapshots
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Param body body services.CalculateSnapshotDTO true "Calculation request"
// @Success 201 {object} models.PatientScoreSnapshot
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id}/score-snapshots [post]
func (h *ScoreSnapshotHandler) CalculateSnapshot(c *fiber.Ctx) error {
	// Get patient ID from URL (already validated by WithSelectedPatient middleware)
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid patient id",
			Message: "patient id must be a valid UUID",
		})
	}

	// Parse request body
	var req services.CalculateSnapshotDTO
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	// Override patientID from URL (security: user can't calculate for different patient)
	req.PatientID = patientID

	// Validate request
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	// Get user ID from JWT
	calculatedByUserID := middleware.GetUserID(c)

	// Calculate snapshot
	snapshot, err := h.snapshotService.CalculateSnapshot(req, calculatedByUserID)
	if err != nil {
		if err.Error() == "patient not found" {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "patient not found",
				Message: "no patient found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to calculate snapshot",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(snapshot)
}

// GetSnapshotsByPatientID lists all snapshots for a patient
// @Summary List patient snapshots
// @Description Retrieves all score snapshots for a patient, ordered by date DESC
// @Tags score-snapshots
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {array} models.PatientScoreSnapshot
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id}/score-snapshots [get]
func (h *ScoreSnapshotHandler) GetSnapshotsByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid patient id",
			Message: "patient id must be a valid UUID",
		})
	}

	snapshots, err := h.snapshotService.GetSnapshotsByPatientID(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to retrieve snapshots",
			Message: err.Error(),
		})
	}

	return c.JSON(snapshots)
}

// GetLatestSnapshotByPatientID retrieves the most recent snapshot for a patient
// @Summary Get latest patient snapshot
// @Description Retrieves the most recent score snapshot for a patient with all relations
// @Tags score-snapshots
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} models.PatientScoreSnapshot
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/patients/{id}/score-snapshots/latest [get]
func (h *ScoreSnapshotHandler) GetLatestSnapshotByPatientID(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid patient id",
			Message: "patient id must be a valid UUID",
		})
	}

	snapshot, err := h.snapshotService.GetLatestSnapshotByPatientID(patientID)
	if err != nil {
		if err.Error() == "no snapshots found for patient" {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "no snapshots found",
				Message: "no score snapshots found for this patient",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to retrieve snapshot",
			Message: err.Error(),
		})
	}

	return c.JSON(snapshot)
}

// GetSnapshotByID retrieves a specific snapshot by ID
// @Summary Get snapshot by ID
// @Description Retrieves a specific score snapshot with all relations
// @Tags score-snapshots
// @Produce json
// @Param id path string true "Snapshot ID"
// @Success 200 {object} models.PatientScoreSnapshot
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/score-snapshots/{id} [get]
func (h *ScoreSnapshotHandler) GetSnapshotByID(c *fiber.Ctx) error {
	snapshotID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid snapshot id",
			Message: "snapshot id must be a valid UUID",
		})
	}

	snapshot, err := h.snapshotService.GetSnapshotByID(snapshotID)
	if err != nil {
		if err.Error() == "snapshot not found" {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "snapshot not found",
				Message: "no snapshot found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to retrieve snapshot",
			Message: err.Error(),
		})
	}

	return c.JSON(snapshot)
}

// DeleteSnapshot soft deletes a snapshot
// @Summary Delete snapshot
// @Description Soft deletes a score snapshot (admin/doctor only)
// @Tags score-snapshots
// @Param id path string true "Snapshot ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/score-snapshots/{id} [delete]
func (h *ScoreSnapshotHandler) DeleteSnapshot(c *fiber.Ctx) error {
	snapshotID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid snapshot id",
			Message: "snapshot id must be a valid UUID",
		})
	}

	err = h.snapshotService.DeleteSnapshot(snapshotID)
	if err != nil {
		if err.Error() == "snapshot not found" {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "snapshot not found",
				Message: "no snapshot found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to delete snapshot",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
