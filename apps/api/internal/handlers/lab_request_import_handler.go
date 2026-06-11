package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

// LabRequestImportHandler — importa um pedido de exames externo (foto/PDF), extrai e casa os exames
// no catálogo para dedup no formulário.
type LabRequestImportHandler struct {
	importService *services.LabRequestImportService
}

func NewLabRequestImportHandler(s *services.LabRequestImportService) *LabRequestImportHandler {
	return &LabRequestImportHandler{importService: s}
}

var importAllowedTypes = map[string]string{
	"application/pdf": ".pdf",
	"image/jpeg":      ".jpg",
	"image/jpg":       ".jpg",
	"image/png":       ".png",
}

// ImportExams godoc
// @Summary Extrai exames de um pedido externo (foto/PDF) e casa com o catálogo
// @Router /api/v1/lab-requests/extract-exams [post]
func (h *LabRequestImportHandler) ImportExams(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "file required", Message: "campo multipart 'file' é obrigatório"})
	}
	contentType := file.Header.Get("Content-Type")
	ext, ok := importAllowedTypes[contentType]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "invalid file type", Message: "envie PDF ou imagem (JPG/PNG)"})
	}
	if file.Size > 20*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "file too large", Message: "máximo 20 MB"})
	}

	tmpPath := fmt.Sprintf("/tmp/labreq-import-%s%s", uuid.NewString(), ext)
	if err := c.SaveFile(file, tmpPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "save failed", Message: err.Error()})
	}
	defer os.Remove(tmpPath)

	isImage := strings.HasPrefix(contentType, "image/")
	items, err := h.importService.ExtractExams(tmpPath, isImage)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "extract failed", Message: err.Error()})
	}

	return c.JSON(fiber.Map{"items": items})
}
