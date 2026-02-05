package handlers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type LabResultBatchHandler struct {
	labResultBatchService *services.LabResultBatchService
	processingJobService  *services.ProcessingJobService
	validator             *validator.Validate
}

func NewLabResultBatchHandler(
	labResultBatchService *services.LabResultBatchService,
	processingJobService *services.ProcessingJobService,
) *LabResultBatchHandler {
	return &LabResultBatchHandler{
		labResultBatchService: labResultBatchService,
		processingJobService:  processingJobService,
		validator:             validator.New(),
	}
}

// Create cria um novo lote de resultados
// @Summary Criar lote de resultados
// @Description Cria um lote de resultados laboratoriais com múltiplos resultados
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param batch body dto.CreateLabResultBatchRequest true "Lote de resultados"
// @Success 201 {object} dto.LabResultBatchDetailResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches [post]
func (h *LabResultBatchHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateLabResultBatchRequest
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

	resp, err := h.labResultBatchService.Create(userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "patient not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to create lab result batch",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// List lista lotes de resultados do selectedPatient
// @Summary Listar lotes de resultados
// @Description Lista lotes de resultados do paciente selecionado
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param status query string false "Filtro por status"
// @Param limit query int false "Limite de resultados" default(50)
// @Param offset query int false "Offset de resultados" default(0)
// @Success 200 {array} dto.LabResultBatchResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches [get]
func (h *LabResultBatchHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	status := c.Query("status")
	var statusPtr *string
	if status != "" {
		statusPtr = &status
	}

	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	batches, err := h.labResultBatchService.List(userID, statusPtr, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to list lab result batches",
			Message: err.Error(),
		})
	}

	return c.JSON(batches)
}

// GetByID busca um lote por ID
// @Summary Buscar lote por ID
// @Description Busca um lote de resultados por ID com todos os resultados
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote"
// @Success 200 {object} dto.LabResultBatchDetailResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{id} [get]
func (h *LabResultBatchHandler) GetByID(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)

	batch, err := h.labResultBatchService.GetByID(batchID, userID)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result batch not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrPatientMismatch) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "access denied",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to get lab result batch",
			Message: err.Error(),
		})
	}

	return c.JSON(batch)
}

// Update atualiza um lote
// @Summary Atualizar lote
// @Description Atualiza metadados de um lote de resultados
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote"
// @Param batch body dto.UpdateLabResultBatchRequest true "Dados para atualização"
// @Success 200 {object} dto.LabResultBatchResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{id} [put]
func (h *LabResultBatchHandler) Update(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	var req dto.UpdateLabResultBatchRequest
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

	resp, err := h.labResultBatchService.Update(batchID, userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result batch not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrPatientMismatch) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "access denied",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to update lab result batch",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete deleta um lote (soft delete)
// @Summary Deletar lote
// @Description Soft delete de um lote de resultados (admin only)
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{id} [delete]
func (h *LabResultBatchHandler) Delete(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)
	userRole := string(middleware.GetPrimaryRole(c))

	err = h.labResultBatchService.Delete(batchID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result batch not found",
				Message: err.Error(),
			})
		}
		if err.Error() == "only admins can delete lab result batches" {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "forbidden",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to delete lab result batch",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// AddResult adiciona um resultado a um lote existente
// @Summary Adicionar resultado ao lote
// @Description Adiciona um novo resultado a um lote existente
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param id path string true "ID do lote"
// @Param result body dto.CreateLabResultInBatchRequest true "Resultado"
// @Success 201 {object} dto.LabResultInBatchResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{id}/results [post]
func (h *LabResultBatchHandler) AddResult(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	var req dto.CreateLabResultInBatchRequest
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

	resp, err := h.labResultBatchService.AddResult(batchID, userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result batch not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrPatientMismatch) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "access denied",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to add result to batch",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// UpdateResult atualiza um resultado individual
// @Summary Atualizar resultado
// @Description Atualiza um resultado individual dentro de um lote
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param batchId path string true "ID do lote"
// @Param resultId path string true "ID do resultado"
// @Param result body dto.UpdateLabResultInBatchRequest true "Dados para atualização"
// @Success 200 {object} dto.LabResultInBatchResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{batchId}/results/{resultId} [put]
func (h *LabResultBatchHandler) UpdateResult(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("batchId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	resultID, err := uuid.Parse(c.Params("resultId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid result id",
			Message: err.Error(),
		})
	}

	var req dto.UpdateLabResultInBatchRequest
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

	resp, err := h.labResultBatchService.UpdateResult(batchID, resultID, userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result batch not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrLabResultNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result not found",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrNoPatientSelected) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "no patient selected",
				Message: err.Error(),
			})
		}
		if errors.Is(err, services.ErrPatientMismatch) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "access denied",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to update result",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// DeleteResult deleta um resultado individual (soft delete)
// @Summary Deletar resultado
// @Description Soft delete de um resultado (admin only)
// @Tags LabResultBatch
// @Accept json
// @Produce json
// @Param batchId path string true "ID do lote"
// @Param resultId path string true "ID do resultado"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/lab-result-batches/{batchId}/results/{resultId} [delete]
func (h *LabResultBatchHandler) DeleteResult(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("batchId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	resultID, err := uuid.Parse(c.Params("resultId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid result id",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)
	userRole := string(middleware.GetPrimaryRole(c))

	err = h.labResultBatchService.DeleteResult(batchID, resultID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrLabResultNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result not found",
				Message: err.Error(),
			})
		}
		if err.Error() == "only admins can delete lab results" {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "forbidden",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to delete result",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UploadPDF - upload de PDF e criação de job de processamento
// @Summary Upload de PDF de laudo médico
// @Description Upload de PDF de laudo médico para interpretação automática via IA
// @Tags LabResultBatch
// @Accept multipart/form-data
// @Produce json
// @Param batchId path string true "Batch ID (UUID)"
// @Param file formData file true "PDF file (max 20MB)"
// @Success 201 {object} map[string]interface{} "PDF uploaded and job created"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/lab-result-batches/{batchId}/upload-pdf [post]
func (h *LabResultBatchHandler) UploadPDF(c *fiber.Ctx) error {
	batchID, err := uuid.Parse(c.Params("batchId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid batch id",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)

	// Validar ownership do batch
	_, err = h.labResultBatchService.GetByID(batchID, userID)
	if err != nil {
		if errors.Is(err, services.ErrLabResultBatchNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "batch not found",
				Message: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to get batch",
			Message: err.Error(),
		})
	}

	// Parse multipart
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "file required",
			Message: "multipart field 'file' is required",
		})
	}

	// Validar tipo
	contentType := file.Header.Get("Content-Type")
	if contentType != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid file type",
			Message: "only PDF files are allowed",
		})
	}

	// Validar tamanho (20MB)
	maxSize := int64(20 * 1024 * 1024) // 20 MB
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "file too large",
			Message: "maximum file size is 20MB",
		})
	}

	// Criar diretório de uploads se não existir
	uploadDir := "/app/uploads/lab-result-batches"
	if err := ensureDir(uploadDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to create upload directory",
			Message: err.Error(),
		})
	}

	// Salvar arquivo com nome único
	filename := fmt.Sprintf("%s_%d.pdf", batchID, time.Now().Unix())
	filepath := filepath.Join(uploadDir, filename)

	if err := c.SaveFile(file, filepath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to save file",
			Message: err.Error(),
		})
	}

	// Criar job de processamento
	job, err := h.processingJobService.Create(batchID, filepath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to create processing job",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "PDF uploaded successfully",
		"jobId":    job.ID,
		"batchId":  batchID,
		"filename": filename,
	})
}

// ensureDir cria diretório se não existir
func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}
