package handlers

import (
	"errors"
	"strconv"
	"strings"
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

// SendEmailAttachment é o ref de um anexo já uploadado (via /conversations/attachments/upload).
// Path é relativo a /app/uploads (ex: "conversation-outbound/<uid>/<uuid>.pdf"); filename é o
// nome original mostrado pro destinatário (e gravado em metadata pra trilha do histórico).
type SendEmailAttachment struct {
	Path     string `json:"path" validate:"required,max=512"`
	Filename string `json:"filename" validate:"required,max=255"`
}

// SendEmailRequest é o payload de POST /conversations/:type/:id/email.
type SendEmailRequest struct {
	Subject     string                `json:"subject" validate:"omitempty,max=200"`
	BodyText    string                `json:"bodyText" validate:"required_without=BodyHTML,max=50000"`
	BodyHTML    string                `json:"bodyHTML" validate:"omitempty,max=200000"`
	InReplyTo   string                `json:"inReplyTo,omitempty"`
	References  []string              `json:"references,omitempty"`
	Attachments []SendEmailAttachment `json:"attachments,omitempty" validate:"omitempty,max=20,dive"`
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

	atts := make([]services.OutboundAttachment, 0, len(req.Attachments))
	for _, a := range req.Attachments {
		atts = append(atts, services.OutboundAttachment{Path: a.Path, Filename: a.Filename})
	}

	activity, err := h.service.SendMessage(c.UserContext(), services.SendMessageInput{
		UserID:      middleware.GetUserID(c),
		OwnerType:   ownerType,
		OwnerID:     ownerID,
		Channel:     models.LeadChannelEmail,
		Subject:     req.Subject,
		BodyText:    req.BodyText,
		BodyHTML:    req.BodyHTML,
		InReplyTo:   req.InReplyTo,
		References:  req.References,
		Attachments: atts,
	})
	return writeSendResult(c, activity, err)
}

// ============================================================
// POST /conversations/compose
// ============================================================

// ComposeRequest é o payload do POST /conversations/compose — envio "novo email" pra
// qualquer endereço (pode não existir Lead/Patient ainda; backend cria Lead se preciso).
type ComposeRequest struct {
	To          string                `json:"to" validate:"required,email,max=255"`
	Name        string                `json:"name,omitempty" validate:"omitempty,max=255"`
	Subject     string                `json:"subject" validate:"required,min=1,max=200"`
	BodyText    string                `json:"bodyText" validate:"required_without=BodyHTML,max=50000"`
	BodyHTML    string                `json:"bodyHTML,omitempty" validate:"omitempty,max=200000"`
	Attachments []SendEmailAttachment `json:"attachments,omitempty" validate:"omitempty,max=20,dive"`
}

// Compose envia "novo email" pra um endereço arbitrário.
//
// Lógica de roteamento:
//   1. Patient com email igual → conversa anexada ao Patient.
//   2. Lead ativo (não converted/lost/unsubscribed) → conversa anexada ao Lead.
//   3. Lead unsubscribed → 422 (não reabre relacionamento sem opt-in novo).
//   4. Caso contrário → cria Lead novo (source=manual, opt-in implícito do vendedor).
//
// @Summary  Inicia conversa por email com endereço arbitrário
// @Tags     conversations
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    body body ComposeRequest true "Email"
// @Success  201 {object} services.ComposeResult
// @Failure  400 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse
// @Router   /conversations/compose [post]
func (h *ConversationHandler) Compose(c *fiber.Ctx) error {
	var req ComposeRequest
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

	atts := make([]services.OutboundAttachment, 0, len(req.Attachments))
	for _, a := range req.Attachments {
		atts = append(atts, services.OutboundAttachment{Path: a.Path, Filename: a.Filename})
	}

	res, _, err := h.service.Compose(c.UserContext(), services.ComposeInput{
		UserID:      middleware.GetUserID(c),
		To:          req.To,
		Name:        req.Name,
		Subject:     req.Subject,
		BodyText:    req.BodyText,
		BodyHTML:    req.BodyHTML,
		Attachments: atts,
	})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrComposeInvalidEmail):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
				Error: "email do destinatário inválido",
			})
		case errors.Is(err, services.ErrComposeRecipientUnsubscribed):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "destinatário descadastrado (Lead unsubscribed)",
			})
		case errors.Is(err, services.ErrConversationNoChannel):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "destinatário sem canal de email válido",
			})
		case errors.Is(err, services.ErrConversationOptedOut):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "destinatário com opt-out de email",
			})
		case errors.Is(err, services.ErrConversationAttachmentTooLarge):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error:   "soma de anexos excede limite Resend (40MB)",
				Message: err.Error(),
			})
		case errors.Is(err, services.ErrConversationAttachmentInvalid):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error:   "anexo inválido",
				Message: err.Error(),
			})
		default:
			return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
				Error:   "failed to send compose email",
				Message: err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(res)
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
// POST /conversations/:type/:id/ai/summary
// ============================================================

// AISummary gera resumo da conversa via Claude. Cache 1h por hash do transcript.
//
// Query: force=true → pula cache.
//
// @Summary  Resumo IA da conversa (3-5 bullets PT-BR)
// @Tags     conversations
// @Security BearerAuth
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Param    force query string false "true para ignorar cache"
// @Success  200 {object} services.AISummaryResult
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse "Conversa sem mensagens"
// @Failure  502 {object} dto.ErrorResponse "Claude API falhou"
// @Failure  504 {object} dto.ErrorResponse "Claude API timeout"
// @Router   /conversations/{type}/{id}/ai/summary [post]
func (h *ConversationHandler) AISummary(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	force := c.Query("force") == "true"

	res, err := h.service.SummarizeConversation(
		c.UserContext(),
		middleware.GetUserID(c),
		ownerType, ownerID, force,
	)
	return writeAIResult(c, res, err)
}

// ============================================================
// POST /conversations/:type/:id/ai/suggest-reply
// ============================================================

// AISuggestReplyRequest é o payload opcional. Intent (ex: "agendar consulta") guia o
// modelo; quando vazio, IA infere a intenção pelo histórico.
type AISuggestReplyRequest struct {
	Intent string `json:"intent,omitempty" validate:"omitempty,max=200"`
}

// AISuggestReply sugere texto de resposta no tom Plenya pro vendedor editar e enviar.
//
// @Summary  Sugestão IA de resposta
// @Tags     conversations
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Param    body body AISuggestReplyRequest false "Intenção do vendedor"
// @Success  200 {object} services.AISuggestionResult
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse "Conversa sem mensagens"
// @Failure  502 {object} dto.ErrorResponse "Claude API falhou"
// @Failure  504 {object} dto.ErrorResponse "Claude API timeout"
// @Router   /conversations/{type}/{id}/ai/suggest-reply [post]
func (h *ConversationHandler) AISuggestReply(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req AISuggestReplyRequest
	// BodyParser tolera body vazio; ignoramos erro de parse pra permitir POST sem corpo.
	_ = c.BodyParser(&req)
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	res, err := h.service.SuggestReply(c.UserContext(), ownerType, ownerID, req.Intent)
	return writeAIResult(c, res, err)
}

// writeAIResult mapeia erros do AI service pra HTTP semântico.
//
// LGPD: NÃO inclui Message com detalhes do erro upstream em status 502/504 — o erro
// do Claude pode ecoar trecho do prompt em alguns casos. Mensagem genérica.
func writeAIResult(c *fiber.Ctx, result interface{}, err error) error {
	if err == nil {
		return c.JSON(result)
	}
	switch {
	case errors.Is(err, services.ErrConversationOwnerInvalid):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "owner not found"})
	case errors.Is(err, services.ErrAIConversationEmpty):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "conversa sem mensagens para análise",
		})
	case errors.Is(err, services.ErrAINotConfigured):
		return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{
			Error: "serviço de IA indisponível",
		})
	case errors.Is(err, services.ErrAIUpstream):
		// Distingue timeout (504) de outras falhas Claude (502).
		if strings.Contains(err.Error(), "timeout") {
			return c.Status(fiber.StatusGatewayTimeout).JSON(dto.ErrorResponse{
				Error: "tempo esgotado ao consultar IA",
			})
		}
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
			Error: "falha ao consultar IA, tente novamente",
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "falha inesperada na IA",
		})
	}
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
	case errors.Is(err, services.ErrConversationAttachmentTooLarge):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error:   "soma de anexos excede limite Resend (40MB)",
			Message: err.Error(),
		})
	case errors.Is(err, services.ErrConversationAttachmentInvalid):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error:   "anexo inválido",
			Message: err.Error(),
		})
	default:
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
			Error:   "failed to send message",
			Message: err.Error(),
		})
	}
}
