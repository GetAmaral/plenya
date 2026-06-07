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
	service     *services.ConversationService
	autoSvc     *services.ConversationAutomationService
	settingsSvc *services.ReceptionSettingsService
	validator   *validator.Validate
}

func NewConversationHandler(s *services.ConversationService, autoSvc *services.ConversationAutomationService, settingsSvc *services.ReceptionSettingsService) *ConversationHandler {
	return &ConversationHandler{
		service:     s,
		autoSvc:     autoSvc,
		settingsSvc: settingsSvc,
		validator:   validator.New(),
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
//
// Media — GET /conversations/:type/:id/media/:activityId
// Serve a mídia de uma mensagem (foto/áudio/documento) após validar ownership.
// Arquivo de prontuário é lido do disco; mídia de conversa é decifrada em memória.
func (h *ConversationHandler) Media(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	activityID, perr := uuid.Parse(c.Params("activityId"))
	if perr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid activityId"})
	}

	res, err := h.service.GetActivityMedia(ownerType, ownerID, activityID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "media not found"})
	}
	return serveMedia(c, res)
}

// serveMedia escreve a mídia com headers seguros: MIME já clampado pra allowlist,
// nosniff (impede sniffing pra text/html), filename já sanitizado no service.
func serveMedia(c *fiber.Ctx, res *services.ActivityMediaResult) error {
	c.Set("X-Content-Type-Options", "nosniff")
	if res.MIME != "" {
		c.Set("Content-Type", res.MIME)
	}
	if res.Filename != "" {
		c.Set("Content-Disposition", "inline; filename=\""+res.Filename+"\"")
	}
	return c.Send(res.Bytes)
}

// InterpretExam — POST /conversations/:type/:id/media/:activityId/interpret-exam
// Submete o documento de prontuário da mensagem ao interpretador de exames (sob demanda).
//
// @Summary  Interpreta a mídia da mensagem como exame
// @Tags     conversations
// @Security BearerAuth
// @Router   /conversations/{type}/{id}/media/{activityId}/interpret-exam [post]
// mediaActionBody é o corpo opcional de save-to-prontuário / interpret-exam.
// attachmentIndex aponta um anexo de e-mail; ausente = mídia WhatsApp da atividade.
type mediaActionBody struct {
	AttachmentIndex *int `json:"attachmentIndex"`
}

func (h *ConversationHandler) InterpretExam(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	activityID, perr := uuid.Parse(c.Params("activityId"))
	if perr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid activityId"})
	}
	var body mediaActionBody
	_ = c.BodyParser(&body) // corpo é opcional

	res, err := h.service.InterpretActivityAsExam(ownerType, ownerID, activityID, body.AttachmentIndex)
	if err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "not found"})
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusAccepted).JSON(res)
}

// SaveToProntuario — POST /conversations/:type/:id/messages/:activityId/save-to-prontuario
// Salva a mídia/anexo da mensagem (e-mail ou WhatsApp) nas mídias do prontuário do paciente.
func (h *ConversationHandler) SaveToProntuario(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	if ownerType != "patient" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: "só conversas de paciente têm prontuário"})
	}
	activityID, perr := uuid.Parse(c.Params("activityId"))
	if perr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid activityId"})
	}
	var body mediaActionBody
	_ = c.BodyParser(&body)

	doc, err := h.service.SaveActivityMediaToProntuario(ownerID, activityID, body.AttachmentIndex)
	if err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "not found"})
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"documentId": doc.ID})
}

// Attachment — GET /conversations/:type/:id/messages/:activityId/attachments/:idx
// Serve um anexo de e-mail (autenticado). Substitui o /uploads estático removido.
func (h *ConversationHandler) Attachment(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	activityID, perr := uuid.Parse(c.Params("activityId"))
	if perr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid activityId"})
	}
	idx, ierr := strconv.Atoi(c.Params("idx"))
	if ierr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid idx"})
	}

	res, err := h.service.GetAttachmentFile(ownerType, ownerID, activityID, idx)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "attachment not found"})
	}
	return serveMedia(c, res)
}

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
		Channel:   c.Query("channel"),
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
//  1. Patient com email igual → conversa anexada ao Patient.
//  2. Lead ativo (não converted/lost/unsubscribed) → conversa anexada ao Lead.
//  3. Lead unsubscribed → 422 (não reabre relacionamento sem opt-in novo).
//  4. Caso contrário → cria Lead novo (source=manual, opt-in implícito do vendedor).
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
// Aceita texto e/ou uma mídia (anexo já enviado via /conversations/attachments/upload).
type SendWhatsAppRequest struct {
	BodyText    string                `json:"bodyText" validate:"required_without=Attachments,max=4096"`
	Attachments []SendEmailAttachment `json:"attachments,omitempty" validate:"omitempty,max=1,dive"`
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

	atts := make([]services.OutboundAttachment, 0, len(req.Attachments))
	for _, a := range req.Attachments {
		atts = append(atts, services.OutboundAttachment{Path: a.Path, Filename: a.Filename})
	}
	activity, err := h.service.SendMessage(c.UserContext(), services.SendMessageInput{
		UserID:      middleware.GetUserID(c),
		OwnerType:   ownerType,
		OwnerID:     ownerID,
		Channel:     models.LeadChannelWhatsApp,
		BodyText:    req.BodyText,
		Attachments: atts,
	})
	return writeSendResult(c, activity, err)
}

// SendWhatsAppTemplateRequest é o payload de POST /conversations/:type/:id/whatsapp/template.
type SendWhatsAppTemplateRequest struct {
	Name     string   `json:"name" validate:"required,max=200"`
	Language string   `json:"language" validate:"omitempty,max=10"`
	Params   []string `json:"params" validate:"omitempty,max=10,dive,max=1024"`
}

// SendWhatsAppTemplate envia um template aprovado (reabre conversa fora da janela 24h).
//
// @Summary  Envia template WhatsApp
// @Tags     conversations
// @Security BearerAuth
// @Router   /conversations/{type}/{id}/whatsapp/template [post]
func (h *ConversationHandler) SendWhatsAppTemplate(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req SendWhatsAppTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	activity, err := h.service.SendWhatsAppTemplate(c.UserContext(), services.SendWhatsAppTemplateInput{
		UserID:       middleware.GetUserID(c),
		OwnerType:    ownerType,
		OwnerID:      ownerID,
		TemplateName: req.Name,
		Language:     req.Language,
		Params:       req.Params,
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

// AIReceptionReply gera a próxima mensagem do recepcionista virtual (ancorada no script +
// objeções + guardrails). No Copiloto, o atendente revisa o reply e envia.
//
// @Summary  Resposta do recepcionista virtual
// @Tags     conversations
// @Security BearerAuth
// @Produce  json
// @Param    type path string true "lead | patient"
// @Param    id   path string true "Owner ID"
// @Success  200 {object} services.ReceptionReplyResult
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse "Conversa sem mensagens"
// @Failure  502 {object} dto.ErrorResponse "Claude API falhou"
// @Router   /conversations/{type}/{id}/ai/reception-reply [post]
func (h *ConversationHandler) AIReceptionReply(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	res, err := h.service.GenerateReceptionReply(c.UserContext(), ownerType, ownerID)
	return writeAIResult(c, res, err)
}

// GetAutomation GET /conversations/:type/:id/automation
// Retorna o modo efetivo do recepcionista virtual para a conversa.
func (h *ConversationHandler) GetAutomation(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	eff, err := h.autoSvc.Get(c.UserContext(), ownerType, ownerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	// Sem override explícito, o modo efetivo é o baseline GLOBAL do Atendimento IA
	// (reception_settings + janela de horário), não o default de ambiente. Mantém o toggle
	// por conversa coerente com o que o job realmente faz (resolveMode).
	if st, sErr := h.settingsSvc.Get(c.UserContext()); sErr == nil {
		globalEff := h.settingsSvc.EffectiveGlobalMode(st, time.Now())
		if mode, _, rErr := h.autoSvc.ResolveEffectiveMode(c.UserContext(), ownerType, ownerID, globalEff); rErr == nil {
			eff.Mode = mode
		}
	}
	return c.JSON(eff)
}

// SuggestedReplyResponse — rascunho salvo + modo efetivo da conversa (p/ o preview X/3 do Auto).
type SuggestedReplyResponse struct {
	*services.SuggestedReplyView
	EffectiveMode string `json:"effectiveMode"` // off|copilot|auto resolvido p/ esta conversa
}

// GetSuggestedReply GET /conversations/:type/:id/suggested-reply
// Rascunho atual da IA (Copiloto/Auto), produzido pelo job após o debounce X.
// 204 No Content quando não há rascunho pendente.
func (h *ConversationHandler) GetSuggestedReply(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	view, err := h.autoSvc.GetSuggestedReply(c.UserContext(), ownerType, ownerID)
	if err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "owner not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	if view == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}
	// Modo efetivo (override da conversa → global com janela) p/ o frontend decidir entre
	// só mostrar o rascunho (Copiloto) ou mostrar o countdown X/3 (Auto).
	st, sErr := h.settingsSvc.Get(c.UserContext())
	mode := ""
	if sErr == nil {
		globalEff := h.settingsSvc.EffectiveGlobalMode(st, time.Now())
		mode, _, _ = h.autoSvc.ResolveEffectiveMode(c.UserContext(), ownerType, ownerID, globalEff)
	}
	return c.JSON(SuggestedReplyResponse{SuggestedReplyView: view, EffectiveMode: mode})
}

// SetAutomationRequest é o payload do toggle de automação.
type SetAutomationRequest struct {
	Mode            string `json:"mode" validate:"required,oneof=off copilot auto"`
	FallbackMinutes int    `json:"fallbackMinutes" validate:"omitempty,gte=0,lte=1440"`
}

// SetAutomation PUT /conversations/:type/:id/automation
// Define o modo (off|copilot|auto) e o fallback da conversa. Reativar limpa a pausa.
func (h *ConversationHandler) SetAutomation(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req SetAutomationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}
	eff, err := h.autoSvc.Set(c.UserContext(), ownerType, ownerID, req.Mode, req.FallbackMinutes, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(eff)
}

// ReceptionMetrics GET /conversations/ai/metrics?days=30
// Métricas de atuação do recepcionista virtual.
func (h *ConversationHandler) ReceptionMetrics(c *fiber.Ctx) error {
	days := c.QueryInt("days", 30)
	m, err := h.autoSvc.ComputeMetrics(c.UserContext(), days)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(m)
}

// GlobalAutomationResponse — settings globais + estado computado (banner).
type GlobalAutomationResponse struct {
	*models.ReceptionSettings
	EffectiveMode string `json:"effectiveMode"` // modo global resolvido agora (janela aplicada)
	AutoActiveNow bool   `json:"autoActiveNow"` // janela do Auto ativa neste momento
}

// GetGlobalAutomation GET /reception/settings — config global do Atendimento IA.
func (h *ConversationHandler) GetGlobalAutomation(c *fiber.Ctx) error {
	st, err := h.settingsSvc.Get(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(GlobalAutomationResponse{
		ReceptionSettings: st,
		EffectiveMode:     h.settingsSvc.EffectiveGlobalMode(st, time.Now()),
		AutoActiveNow:     h.settingsSvc.AutoActiveNow(st),
	})
}

// SetGlobalAutomationRequest — atualização parcial (ponteiro = mexer; nil = manter).
type SetGlobalAutomationRequest struct {
	BaselineMode           *string                `json:"baselineMode" validate:"omitempty,oneof=off copilot auto"`
	ScheduleEnabled        *bool                  `json:"scheduleEnabled"`
	Schedule               *models.WeeklySchedule `json:"schedule"`
	DebounceSeconds        *int                   `json:"debounceSeconds" validate:"omitempty,gte=0,lte=3600"`
	PreviewSeconds         *int                   `json:"previewSeconds" validate:"omitempty,gte=1,lte=600"`
	CopilotFallbackMinutes *int                   `json:"copilotFallbackMinutes" validate:"omitempty,gte=1,lte=1440"`
	OffAlertMinutes        *int                   `json:"offAlertMinutes" validate:"omitempty,gte=1,lte=1440"`
}

// SetGlobalAutomation PUT /reception/settings — atualiza a config global.
func (h *ConversationHandler) SetGlobalAutomation(c *fiber.Ctx) error {
	var req SetGlobalAutomationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}
	st, err := h.settingsSvc.Update(c.UserContext(), services.UpdateReceptionSettingsInput{
		BaselineMode:           req.BaselineMode,
		ScheduleEnabled:        req.ScheduleEnabled,
		Schedule:               req.Schedule,
		DebounceSeconds:        req.DebounceSeconds,
		PreviewSeconds:         req.PreviewSeconds,
		CopilotFallbackMinutes: req.CopilotFallbackMinutes,
		OffAlertMinutes:        req.OffAlertMinutes,
	}, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(GlobalAutomationResponse{
		ReceptionSettings: st,
		EffectiveMode:     h.settingsSvc.EffectiveGlobalMode(st, time.Now()),
		AutoActiveNow:     h.settingsSvc.AutoActiveNow(st),
	})
}

// ============================================================
// Dossiê 360 (social) — leitura + edição manual pela equipe
// ============================================================

// GetDossier GET /conversations/:type/:id/dossier
// Visão 360 social (resumo + fatos) da pessoa. Nada clínico (LGPD/CFM).
func (h *ConversationHandler) GetDossier(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	view, err := h.service.GetDossier(c.UserContext(), ownerType, ownerID)
	if err != nil {
		if errors.Is(err, services.ErrConversationOwnerInvalid) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "owner not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(view)
}

// DossierFactRequest é o payload para adicionar/editar um fato social.
type DossierFactRequest struct {
	Category string `json:"category" validate:"omitempty,max=40"`
	Key      string `json:"key" validate:"omitempty,max=60"`
	Value    string `json:"value" validate:"required,max=2000"`
}

// AddDossierFact POST /conversations/:type/:id/dossier/facts
func (h *ConversationHandler) AddDossierFact(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req DossierFactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	view, err := h.service.AddDossierFact(c.UserContext(), ownerType, ownerID, req.Category, req.Key, req.Value, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(view)
}

// UpdateDossierFact PATCH /conversations/:type/:id/dossier/facts/:factId
func (h *ConversationHandler) UpdateDossierFact(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	factID, err := uuid.Parse(c.Params("factId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid fact id"})
	}
	var req DossierFactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	view, err := h.service.UpdateDossierFact(c.UserContext(), ownerType, ownerID, factID, req.Value, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(view)
}

// DeleteDossierFact DELETE /conversations/:type/:id/dossier/facts/:factId
func (h *ConversationHandler) DeleteDossierFact(c *fiber.Ctx) error {
	ownerType, ownerID, err := parseOwnerParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	factID, err := uuid.Parse(c.Params("factId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid fact id"})
	}
	view, err := h.service.DeleteDossierFact(c.UserContext(), ownerType, ownerID, factID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(view)
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
