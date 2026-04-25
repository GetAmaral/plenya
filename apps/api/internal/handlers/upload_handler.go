package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
)

// UploadHandler aceita upload genérico (imagens + PDF) usado por features
// como avaliação física, postural e anexos de anamnese. Retorna path/url
// que o caller anexa ao recurso (PhysicalAssessment.Photos[]).
//
// Storage: <uploadRoot>/general/<userID>/<uuidv7>.<ext>
// Servido via app.Static("/uploads", "/app/uploads").
type UploadHandler struct {
	uploadRoot string
}

// NewUploadHandler instancia o handler.
func NewUploadHandler(uploadRoot string) *UploadHandler {
	if uploadRoot == "" {
		uploadRoot = "/app/uploads"
	}
	return &UploadHandler{uploadRoot: uploadRoot}
}

const maxGeneralUploadBytes = 10 * 1024 * 1024 // 10MB

var generalUploadAllowedExts = map[string]bool{
	".pdf":  true,
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".heic": true,
}

var generalUploadAllowedMimes = map[string]bool{
	"application/pdf":          true,
	"application/octet-stream": true, // mobile às vezes manda assim p/ heic
	"image/jpeg":               true,
	"image/png":                true,
	"image/webp":               true,
	"image/heic":               true,
	"image/heif":               true,
}

// UploadResponse é o payload retornado.
type UploadResponse struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	URL         string `json:"url"`
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	SizeBytes   int64  `json:"sizeBytes"`
}

// Create godoc
// @Summary Upload genérico de arquivo (imagem/PDF)
// @Description Aceita JPEG, PNG, WebP, HEIC ou PDF até 10MB. Use o `path` retornado para anexar ao recurso (avaliação física, anamnese, etc.).
// @Tags uploads
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Arquivo (imagem ou PDF)"
// @Success 201 {object} UploadResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 413 {object} dto.ErrorResponse
// @Router /uploads [post]
func (h *UploadHandler) Create(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "arquivo não enviado",
			Message: err.Error(),
		})
	}

	if file.Size <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "arquivo vazio"})
	}
	if file.Size > maxGeneralUploadBytes {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(dto.ErrorResponse{
			Error: fmt.Sprintf("arquivo excede limite de %dMB", maxGeneralUploadBytes/(1024*1024)),
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !generalUploadAllowedExts[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "extensão não permitida (use jpg, jpeg, png, webp, heic ou pdf)",
		})
	}

	contentType := strings.ToLower(strings.TrimSpace(file.Header.Get("Content-Type")))
	if i := strings.Index(contentType, ";"); i >= 0 {
		contentType = strings.TrimSpace(contentType[:i])
	}
	if contentType != "" && !generalUploadAllowedMimes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "content-type não permitido",
			Message: contentType,
		})
	}

	relDir := filepath.Join("general", userID.String())
	absDir := filepath.Join(h.uploadRoot, relDir)
	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "falha ao preparar storage",
			Message: err.Error(),
		})
	}

	id, err := uuid.NewV7()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao gerar id"})
	}
	storedName := id.String() + ext
	relPath := filepath.Join(relDir, storedName)
	absPath := filepath.Join(absDir, storedName)

	if err := saveMultipart(file, absPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "falha ao salvar arquivo",
			Message: err.Error(),
		})
	}

	cleanName := filepath.Base(file.Filename)
	if cleanName == "" || cleanName == "." || cleanName == "/" {
		cleanName = "arquivo" + ext
	}

	return c.Status(fiber.StatusCreated).JSON(UploadResponse{
		ID:          id.String(),
		Path:        filepath.ToSlash(relPath),
		URL:         "/uploads/" + filepath.ToSlash(relPath),
		Filename:    cleanName,
		ContentType: contentType,
		SizeBytes:   file.Size,
	})
}
