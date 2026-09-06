package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
	"github.com/plenya/api/internal/utils"
)

type IssuedDocumentHandler struct {
	service   *services.IssuedDocumentService
	validator *validator.Validate
}

func NewIssuedDocumentHandler(service *services.IssuedDocumentService) *IssuedDocumentHandler {
	return &IssuedDocumentHandler{service: service, validator: validator.New()}
}

func (h *IssuedDocumentHandler) mapErr(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, services.ErrIssuedDocumentNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "documento não encontrado", Message: err.Error()})
	case errors.Is(err, services.ErrPatientNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "paciente não encontrado", Message: err.Error()})
	case errors.Is(err, services.ErrIssuedDocAlreadySigned):
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{Error: "documento já assinado", Message: err.Error()})
	case errors.Is(err, services.ErrCIDConsentRequired):
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "consentimento de CID obrigatório", Message: err.Error()})
	case errors.Is(err, services.ErrDoctorMissingCRM):
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "médico sem CRM", Message: err.Error()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "erro interno", Message: err.Error()})
	}
}

// ListByPatient GET /api/v1/patients/:id/issued-documents
func (h *IssuedDocumentHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	docs, err := h.service.ListByPatient(patientID)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(docs)
}

// Create POST /api/v1/patients/:id/issued-documents
func (h *IssuedDocumentHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	var req dto.CreateIssuedDocumentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	doc, err := h.service.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(doc)
}

// Update PUT /api/v1/issued-documents/:docId — edita rascunho (assinado é imutável).
func (h *IssuedDocumentHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid document id", Message: err.Error()})
	}
	var req dto.CreateIssuedDocumentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	doc, err := h.service.Update(id, &req)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(doc)
}

// Sign POST /api/v1/issued-documents/:docId/sign
func (h *IssuedDocumentHandler) Sign(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid document id", Message: err.Error()})
	}
	doc, err := h.service.Sign(id, middleware.GetUserID(c))
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(doc)
}

// GetByID GET /api/v1/issued-documents/:docId
func (h *IssuedDocumentHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid document id", Message: err.Error()})
	}
	doc, err := h.service.GetByID(id)
	if err != nil {
		return h.mapErr(c, err)
	}
	return c.JSON(doc)
}

// DownloadPDF GET /api/v1/issued-documents/:docId/pdf (staff autenticado)
func (h *IssuedDocumentHandler) DownloadPDF(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid document id", Message: err.Error()})
	}
	full, fileName, contentType, err := h.service.GetForDownload(id)
	if err != nil {
		return h.mapErr(c, err)
	}
	c.Set("Content-Type", contentType)
	// filename= sozinho não carrega acento: "Ana-Cláudia" chegava como "Ana-Cldia" ou quebrava o
	// cabeçalho. utils.ContentDisposition manda as duas formas.
	c.Set("Content-Disposition", utils.ContentDisposition(fileName, "documento.pdf"))
	return c.SendFile(full)
}

// Delete DELETE /api/v1/issued-documents/:docId
func (h *IssuedDocumentHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid document id", Message: err.Error()})
	}
	if err := h.service.Delete(id); err != nil {
		return h.mapErr(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ValidatePublic GET /api/v1/documents/validate/:id (público, sem auth)
func (h *IssuedDocumentHandler) ValidatePublic(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"valid": false, "error": "invalid id"})
	}
	res, err := h.service.ValidatePublic(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"valid": false, "error": "document not found"})
	}
	return c.JSON(res)
}
