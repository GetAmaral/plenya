package handlers

import (
	"errors"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// AnonymousScoreHandler expõe o Escore Plenya Light no domínio público
// (sem autenticação). Ver registerScoreLightRoutes em main.go.
type AnonymousScoreHandler struct {
	service     *services.AnonymousScoreService
	authService *services.AuthService
	pdfService  *services.AnonymousLabPDFService
	validator   *validator.Validate
}

func NewAnonymousScoreHandler(
	service *services.AnonymousScoreService,
	authService *services.AuthService,
	pdfService *services.AnonymousLabPDFService,
) *AnonymousScoreHandler {
	return &AnonymousScoreHandler{
		service:     service,
		authService: authService,
		pdfService:  pdfService,
		validator:   validator.New(),
	}
}

// Tamanho máximo de upload: 10MB
const maxLightPDFSize = 10 * 1024 * 1024

// ExtractLabsFromPDF — endpoint público para extração de exames de PDF do Light.
// Multipart/form-data com campo "pdf". Retorna LightLabExtractionResult.
//
// Operação síncrona (~8-15s). PDF não é persistido (LGPD).
//
// @Summary  Extrai exames de PDF para preencher itens do Escore Light
// @Tags     score-light
// @Accept   multipart/form-data
// @Produce  json
// @Param    pdf formData file true "PDF dos exames laboratoriais"
// @Success  200 {object} services.LightLabExtractionResult
// @Router   /score-light/extract-labs [post]
func (h *AnonymousScoreHandler) ExtractLabsFromPDF(c *fiber.Ctx) error {
	if h.pdfService == nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{
			Error: "pdf extraction service not available",
		})
	}
	fileHeader, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "missing file",
			Message: "envie o PDF no campo 'pdf' (multipart/form-data)",
		})
	}
	if fileHeader.Size > maxLightPDFSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(dto.ErrorResponse{
			Error:   "file too large",
			Message: "PDF deve ter no máximo 10MB",
		})
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "failed to open file",
			Message: err.Error(),
		})
	}
	defer src.Close()
	data, err := io.ReadAll(src)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "failed to read file",
			Message: err.Error(),
		})
	}

	tmpPath, err := services.SaveTempPDF(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to stage upload",
			Message: err.Error(),
		})
	}

	result, err := h.pdfService.ExtractFromPDF(tmpPath)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error:   "extraction failed",
			Message: err.Error(),
		})
	}
	return c.JSON(result)
}

// GetConfig retorna a configuração completa do Escore Light (todos os items
// marcados como IsLightVersion + levels). Cacheável no edge.
//
// @Summary  Configuração pública do Escore Plenya Light
// @Tags     score-light
// @Produce  json
// @Success  200 {object} services.LightConfig
// @Router   /score-light/config [get]
func (h *AnonymousScoreHandler) GetConfig(c *fiber.Ctx) error {
	cfg, err := h.service.BuildLightConfig()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to build light config",
			Message: err.Error(),
		})
	}
	c.Set("Cache-Control", "public, max-age=300, s-maxage=3600")
	return c.JSON(cfg)
}

// CreateSession cria uma nova sessão anônima com as respostas do paciente
// e retorna o snapshot calculado.
//
// @Summary  Submete respostas do Escore Light
// @Tags     score-light
// @Accept   json
// @Produce  json
// @Param    body body services.CreateSessionRequest true "Respostas + dados demográficos"
// @Success  201 {object} models.AnonymousScoreSession
// @Router   /score-light/sessions [post]
func (h *AnonymousScoreHandler) CreateSession(c *fiber.Ctx) error {
	var req services.CreateSessionRequest
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

	session, err := h.service.CreateSession(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to create session",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}

// GetSession retorna a sessão anônima pelo PublicCode.
//
// @Summary  Recupera resultado pelo código público
// @Tags     score-light
// @Produce  json
// @Param    code path string true "Public code"
// @Success  200 {object} models.AnonymousScoreSession
// @Failure  404 {object} dto.ErrorResponse
// @Router   /score-light/sessions/{code} [get]
func (h *AnonymousScoreHandler) GetSession(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "public code is required",
		})
	}
	session, err := h.service.GetSessionByPublicCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "session not found or expired",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to load session",
			Message: err.Error(),
		})
	}
	return c.JSON(session)
}

// ClaimRequest é o payload do POST /sessions/:code/claim
type ClaimRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// RequestClaim envia o magic link por email. Retorna 200 mesmo se a sessão não
// existir (anti-enumeração).
//
// @Summary  Solicita magic link para salvar resultado
// @Tags     score-light
// @Accept   json
// @Produce  json
// @Param    code path string true "Public code da sessão"
// @Param    body body ClaimRequest true "Email do paciente"
// @Success  200 {object} map[string]string
// @Router   /score-light/sessions/{code}/claim [post]
func (h *AnonymousScoreHandler) RequestClaim(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "code is required"})
	}
	var req ClaimRequest
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

	if err := h.service.RequestClaim(code, req.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to send magic link",
			Message: err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Se a sessão existir, um link foi enviado para o email informado.",
	})
}

// ConfirmRequest é o payload do POST /sessions/claim/confirm
type ConfirmRequest struct {
	Token string `json:"token" validate:"required"`
}

// ConfirmClaim valida o magic token e devolve tokens do EMR para login automático.
//
// @Summary  Confirma o magic link e devolve tokens do EMR
// @Tags     score-light
// @Accept   json
// @Produce  json
// @Param    body body ConfirmRequest true "Magic token JWT"
// @Success  200 {object} services.ConfirmClaimResult
// @Failure  400 {object} dto.ErrorResponse
// @Router   /score-light/claim/confirm [post]
func (h *AnonymousScoreHandler) ConfirmClaim(c *fiber.Ctx) error {
	var req ConfirmRequest
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

	result, err := h.service.ConfirmClaim(req.Token, h.authService)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "claim failed",
			Message: err.Error(),
		})
	}
	return c.JSON(result)
}

// DeleteSession — direito de eliminação (LGPD art. 18 VI).
// Excluir uma sessão Light pelo publicCode. Quem tem o code é o owner.
//
// @Summary  Excluir uma sessão Light (LGPD direito de eliminação)
// @Tags     score-light
// @Produce  json
// @Param    code path string true "Código público da sessão"
// @Success  204
// @Failure  404 {object} dto.ErrorResponse
// @Router   /score-light/sessions/{code} [delete]
func (h *AnonymousScoreHandler) DeleteSession(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "code is required"})
	}
	if err := h.service.DeleteSessionByPublicCode(code); err != nil {
		if err.Error() == "session not found" {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "session not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to delete session",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// MySessions retorna todas as sessões claimed pelo paciente atual.
// Rota autenticada — usada pela área do paciente no EMR.
//
// @Summary  Sessões Light do paciente logado
// @Tags     score-light
// @Security BearerAuth
// @Produce  json
// @Success  200 {array} services.PublicSession
// @Router   /score-light/my-sessions [get]
func (h *AnonymousScoreHandler) MySessions(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Error: "unauthorized"})
	}
	sessions, err := h.service.GetSessionsByCurrentUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to load sessions",
			Message: err.Error(),
		})
	}
	return c.JSON(sessions)
}
