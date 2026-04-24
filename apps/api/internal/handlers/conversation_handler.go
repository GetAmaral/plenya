package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// ConversationHandler expõe os endpoints da Central de Conversas (Bloco B).
type ConversationHandler struct {
	service   *services.ConversationService
	validator *validator.Validate
}

func NewConversationHandler(s *services.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		service:   s,
		validator: validator.New(),
	}
}

// ============================================================
// GET /conversations
// ============================================================
//
// Query params:
//   - assigned_to_me=true   → filtra Leads atribuídos ao userID (Patients sempre incluídos)
//   - unread_only=true      → só conversas com unread_count > 0
//   - channel=email|whatsapp
//   - search=foo            → name/email/phone (LIKE case-insensitive)
//   - cursor=<base64>       → cursor pra próxima página
//   - limit=50              → 1..200, default 50
//
// @Summary  Lista conversas (Lead + Patient unificado)
// @Tags     conversations
// @Security BearerAuth
// @Produce  json
// @Success  200 {object} services.ListConversationsResult
// @Router   /conversations [get]
func (h *ConversationHandler) List(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	limit, _ := strconv.Atoi(c.Query("limit", "50"))

	params := services.ListConversationsParams{
		UserID:        userID,
		AssignedToMe:  c.Query("assigned_to_me") == "true",
		UnreadOnly:    c.Query("unread_only") == "true",
		ChannelFilter: c.Query("channel"),
		Search:        c.Query("search"),
		Limit:         limit,
		OffsetCursor:  c.Query("cursor"),
	}
	if params.ChannelFilter != "" && params.ChannelFilter != "email" && params.ChannelFilter != "whatsapp" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "channel must be email or whatsapp",
		})
	}

	res, err := h.service.List(c.UserContext(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to list conversations",
			Message: err.Error(),
		})
	}
	return c.JSON(res)
}

// ============================================================
// GET /conversations/:type/:id/messages
// ============================================================
//
// :type = "lead" | "patient"
// :id   = uuid
// Query: before=<RFC3339>, limit=100 (max 500)
//
// Side-effect: marca conversa como lida (LastReadAt = now). UI ganha unread=0
// imediatamente após este endpoint.
//
// @Summary  Histórico de mensagens dum owner (Lead OU Patient)
// @Tags     conversations
// @Security BearerAuth
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Success  200 {array} models.LeadActivity
// @Failure  404 {object} dto.ErrorResponse
// @Router   /conversations/{type}/{id}/messages [get]
func (h *ConversationHandler) Messages(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	var before *time.Time
	if s := c.Query("before"); s != "" {
		t, perr := time.Parse(time.RFC3339, s)
		if perr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error:   "invalid before (use RFC3339)",
				Message: perr.Error(),
			})
		}
		before = &t
	}

	activities, err := h.service.Messages(c.UserContext(), services.GetMessagesParams{
		UserID:    middleware.GetUserID(c),
		OwnerType: ownerType,
		OwnerID:   ownerID,
		Limit:     limit,
		Before:    before,
	})
	if err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "owner not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to load messages",
			Message: err.Error(),
		})
	}
	return c.JSON(fiber.Map{"items": activities})
}

// ============================================================
// POST /conversations/:type/:id/read
// ============================================================
//
// Marca lido sem fetch das mensagens. Idempotente.
//
// @Summary  Marca conversa como lida
// @Tags     conversations
// @Security BearerAuth
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Success  204
// @Router   /conversations/{type}/{id}/read [post]
func (h *ConversationHandler) MarkRead(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	if err := h.service.MarkRead(c.UserContext(), middleware.GetUserID(c), ownerType, ownerID); err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid owner"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to mark read",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ============================================================
// POST /conversations/:type/:id/email
// ============================================================

// SendEmailRequest é o payload de POST /conversations/:type/:id/email.
type SendEmailRequest struct {
	Subject    string   `json:"subject" validate:"omitempty,max=200"`
	BodyText   string   `json:"bodyText" validate:"required_without=BodyHTML,max=50000"`
	BodyHTML   string   `json:"bodyHTML" validate:"omitempty,max=200000"`
	InReplyTo  string   `json:"inReplyTo,omitempty"`
	References []string `json:"references,omitempty"`
}

// SendEmail envia resposta de email pelo owner (Lead OU Patient).
//
// @Summary  Envia email pra conversa
// @Tags     conversations
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Param    body body SendEmailRequest true "Email"
// @Success  201 {object} models.LeadActivity
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse
// @Router   /conversations/{type}/{id}/email [post]
func (h *ConversationHandler) SendEmail(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req SendEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid body",
			Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	activity, err := h.service.SendMessage(c.UserContext(), services.SendMessageInput{
		UserID:     middleware.GetUserID(c),
		OwnerType:  ownerType,
		OwnerID:    ownerID,
		Channel:    models.LeadChannelEmail,
		Subject:    req.Subject,
		BodyText:   req.BodyText,
		BodyHTML:   req.BodyHTML,
		InReplyTo:  req.InReplyTo,
		References: req.References,
	})
	return writeSendResult(c, activity, err)
}

// ============================================================
// POST /conversations/:type/:id/whatsapp
// ============================================================

// SendWhatsAppRequest é o payload de POST /conversations/:type/:id/whatsapp.
type SendWhatsAppRequest struct {
	BodyText string `json:"bodyText" validate:"required,min=1,max=4096"`
}

// SendWhatsApp envia mensagem WhatsApp (session message, dentro da janela 24h).
//
// @Summary  Envia mensagem WhatsApp pra conversa
// @Tags     conversations
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Param    body body SendWhatsAppRequest true "Mensagem"
// @Success  201 {object} models.LeadActivity
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse
// @Router   /conversations/{type}/{id}/whatsapp [post]
func (h *ConversationHandler) SendWhatsApp(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req SendWhatsAppRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid body",
			Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	activity, err := h.service.SendMessage(c.UserContext(), services.SendMessageInput{
		UserID:    middleware.GetUserID(c),
		OwnerType: ownerType,
		OwnerID:   ownerID,
		Channel:   models.LeadChannelWhatsApp,
		BodyText:  req.BodyText,
	})
	return writeSendResult(c, activity, err)
}

// ============================================================
// Helpers
// ============================================================

// parseOwnerParams extrai e valida :type e :id da URL.
func parseOwnerParams(c *fiber.Ctx) (string, uuid.UUID, error) {
	t := c.Params("type")
	if t != "lead" && t != "patient" {
		return "", uuid.Nil, errors.New("type must be lead or patient")
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return "", uuid.Nil, errors.New("invalid id")
	}
	return t, id, nil
}

// writeSendResult mapeia errors do ConversationService pra HTTP semântico.
//   - ErrConversationOwnerInvalid     → 404
//   - ErrConversationNoChannel        → 422 (sem email/phone)
//   - ErrConversationOptedOut         → 422
//   - ErrConversationWindowClosed     → 422
//   - default (Resend/Meta falhou)    → 502
func writeSendResult(c *fiber.Ctx, activity *models.LeadActivity, err error) error {
	if err == nil {
		return c.Status(fiber.StatusCreated).JSON(activity)
	}
	switch {
	case errors.Is(err, services.ErrConversationOwnerInvalid):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "owner not found"})
	case errors.Is(err, services.ErrConversationNoChannel):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "owner sem email/telefone cadastrado",
		})
	case errors.Is(err, services.ErrConversationOptedOut):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "owner com opt-out de canal",
		})
	case errors.Is(err, services.ErrConversationWindowClosed):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "fora da janela de 24h — use template aprovado",
		})
	default:
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
			Error:   "failed to send message",
			Message: err.Error(),
		})
	}
}
