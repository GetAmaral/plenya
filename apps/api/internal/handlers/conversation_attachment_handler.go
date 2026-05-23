package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
)

// ConversationAttachmentHandler gerencia upload de anexos pra serem enviados em emails
// da Central de Conversas. Storage segue o padrão do projeto: /app/uploads/<bucket>/...
//
// Bucket: conversation-outbound/<userid>/<uuidv7>.<ext>
// Servido via app.Static("/uploads", "/app/uploads") — frontend baixa em
// https://api.plenyasaude.com.br/uploads/<path>.
type ConversationAttachmentHandler struct {
	uploadRoot string // "/app/uploads"
}

// NewConversationAttachmentHandler cria o handler. uploadRoot tipicamente "/app/uploads".
func NewConversationAttachmentHandler(uploadRoot string) *ConversationAttachmentHandler {
	if uploadRoot == "" {
		uploadRoot = "/app/uploads"
	}
	return &ConversationAttachmentHandler{uploadRoot: uploadRoot}
}

// Limites e whitelists ───────────────────────────────────────────────

// maxAttachmentBytes limita por arquivo. Resend permite 40MB por email TOTAL —
// validação do total fica em SendConversationReply (soma de todos paths).
const maxAttachmentBytes = 10 * 1024 * 1024

// allowedAttachmentExts é a whitelist de extensões aceitas.
// Confronto duplo (ext + content-type) defende contra rename (.exe → .pdf).
var allowedAttachmentExts = map[string]bool{
	".pdf":  true,
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".doc":  true,
	".docx": true,
	".xls":  true,
	".xlsx": true,
	// Áudio (envio de nota de voz pelo EMR via WhatsApp).
	".ogg": true,
	".mp3": true,
	".m4a": true,
	".aac": true,
}

// allowedAttachmentMimePrefixes — match por prefixo (cobre "image/jpeg",
// "application/vnd.openxmlformats-...", etc.). Vazio quando o navegador não
// envia content-type → caímos no check por extensão somente.
var allowedAttachmentMimePrefixes = []string{
	"application/pdf",
	"image/jpeg",
	"image/png",
	"image/webp",
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.ms-excel",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"audio/ogg",
	"audio/mpeg",
	"audio/mp4",
	"audio/aac",
	"application/octet-stream", // alguns browsers mandam isso pra .docx — fica permitido só se ext bater
}

// ConversationAttachmentResponse é o payload retornado ao frontend.
// Path é relativo ao uploadRoot ("conversation-outbound/<uid>/<file>"); frontend
// concatena com /uploads/ pra montar URL pública e usa em SendEmail payload.
type ConversationAttachmentResponse struct {
	Path        string `json:"path"`
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int64  `json:"size"`
}

// Upload aceita multipart/form-data (campo "file"), valida e salva em disco.
//
// @Summary  Upload de anexo pra conversa
// @Tags     conversations
// @Security BearerAuth
// @Accept   multipart/form-data
// @Produce  json
// @Param    file formData file true "Arquivo (PDF, imagem ou Office)"
// @Success  201 {object} ConversationAttachmentResponse
// @Failure  400 {object} dto.ErrorResponse
// @Failure  413 {object} dto.ErrorResponse
// @Router   /conversations/attachments/upload [post]
func (h *ConversationAttachmentHandler) Upload(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "arquivo não enviado",
			Message: err.Error(),
		})
	}

	if file.Size <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "arquivo vazio",
		})
	}
	if file.Size > maxAttachmentBytes {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(dto.ErrorResponse{
			Error: fmt.Sprintf("arquivo excede limite de %dMB", maxAttachmentBytes/(1024*1024)),
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedAttachmentExts[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "extensão não permitida (use pdf, jpg, png, webp, doc, docx, xls, xlsx)",
		})
	}

	// MIME do header — se vier preenchido, exige whitelist; se vier vazio, OK
	// (alguns browsers / mobile não setam corretamente pra Office).
	contentType := file.Header.Get("Content-Type")
	if contentType != "" && !mimeIsAllowed(contentType) {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "content-type não permitido",
			Message: contentType,
		})
	}

	// Path: conversation-outbound/<userid>/<uuidv7>.<ext>
	relDir := filepath.Join("conversation-outbound", userID.String())
	absDir := filepath.Join(h.uploadRoot, relDir)
	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "falha ao preparar storage",
			Message: err.Error(),
		})
	}

	id, err := uuid.NewV7()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "falha ao gerar id",
		})
	}
	storedName := id.String() + ext
	relPath := filepath.Join(relDir, storedName)
	absPath := filepath.Join(absDir, storedName)

	// Stream copy (não usamos SaveFile pra fechar handles explicitamente em erro)
	if err := saveMultipart(file, absPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "falha ao salvar arquivo",
			Message: err.Error(),
		})
	}

	// Sanitiza filename retornado (mantém original limpo de path traversal)
	cleanName := filepath.Base(file.Filename)
	if cleanName == "" || cleanName == "." || cleanName == "/" {
		cleanName = "anexo" + ext
	}

	return c.Status(fiber.StatusCreated).JSON(ConversationAttachmentResponse{
		Path:        filepath.ToSlash(relPath),
		Filename:    cleanName,
		ContentType: contentType,
		Size:        file.Size,
	})
}

func mimeIsAllowed(ct string) bool {
	ct = strings.ToLower(strings.TrimSpace(ct))
	// Strip parâmetros tipo "; charset=..."
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	for _, p := range allowedAttachmentMimePrefixes {
		if ct == p {
			return true
		}
	}
	return false
}

func saveMultipart(file *multipart.FileHeader, absPath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()

	dst, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer func() { _ = dst.Close() }()

	if _, err := io.Copy(dst, src); err != nil {
		_ = os.Remove(absPath)
		return err
	}
	return nil
}
