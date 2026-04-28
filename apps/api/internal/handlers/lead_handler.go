package handlers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// LeadHandler expõe endpoints públicos (criação via form) e autenticados (admin CRUD).
type LeadHandler struct {
	service      *services.LeadService
	emailService *services.EmailService
	turnstile    *services.TurnstileService // M10 — pode ser nil em dev
	validator    *validator.Validate
}

func NewLeadHandler(service *services.LeadService, emailService *services.EmailService) *LeadHandler {
	return &LeadHandler{
		service:      service,
		emailService: emailService,
		validator:    validator.New(),
	}
}

// WithTurnstile — M10 — injeta o validator. Se nil ou Secret vazio, PublicCreate
// pula a checagem de CAPTCHA com warning.
func (h *LeadHandler) WithTurnstile(t *services.TurnstileService) *LeadHandler {
	h.turnstile = t
	return h
}

// ============================================================
// Público — POST /leads (formulário de contato do site)
// ============================================================

// PublicCreateRequest é o payload do formulário de contato no site público.
type PublicCreateRequest struct {
	Name            string         `json:"name" validate:"required,min=2,max=200"`
	Email           string         `json:"email,omitempty" validate:"omitempty,email"`
	Phone           string         `json:"phone,omitempty"` // formato E.164 ou BR livre — normalizado no service
	Message         string         `json:"message,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
	ConsentVersion  string         `json:"consentVersion" validate:"required,min=5,max=20"`
	NewsletterOptIn bool           `json:"newsletterOptIn,omitempty"`
	// M10 — Cloudflare Turnstile token (cf-turnstile-response). Em prod, obrigatório.
	// Em dev (Secret vazio), validação pulada.
	TurnstileToken string `json:"turnstileToken,omitempty"`
}

// PublicCreate — endpoint público rate-limited. Cria Lead com source=contact_form.
//
// @Summary  Criar lead (formulário de contato no site público)
// @Tags     leads
// @Accept   json
// @Produce  json
// @Param    body body PublicCreateRequest true "Dados do formulário"
// @Success  201 {object} models.Lead
// @Failure  400 {object} dto.ErrorResponse
// @Router   /leads [post]
func (h *LeadHandler) PublicCreate(c *fiber.Ctx) error {
	var req PublicCreateRequest
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
	if strings.TrimSpace(req.Email) == "" && strings.TrimSpace(req.Phone) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "email or phone is required",
		})
	}

	// M10 — valida Turnstile (CAPTCHA Cloudflare). Em dev (Secret vazio),
	// h.turnstile.Verify retorna nil sem fazer chamada HTTP.
	if h.turnstile != nil {
		// Header tem precedência sobre body (frontend pode mandar via header).
		token := c.Get("cf-turnstile-response")
		if token == "" {
			token = req.TurnstileToken
		}
		if err := h.turnstile.Verify(c.Context(), token, c.IP()); err != nil {
			if errors.Is(err, services.ErrTurnstileMissing) || errors.Is(err, services.ErrTurnstileInvalid) {
				return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
					Error: "captcha falhou — recarregue a página e tente novamente",
				})
			}
			// Erro de rede com Cloudflare: retorna 503 (não 500) — comum sob picos.
			return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{
				Error:   "captcha service indisponível",
				Message: err.Error(),
			})
		}
	}

	lead, err := h.service.CreateFromContactForm(services.CreateFromContactFormInput{
		Name:            req.Name,
		Email:           req.Email,
		Phone:           req.Phone,
		Message:         req.Message,
		Metadata:        req.Metadata,
		ConsentVersion:  req.ConsentVersion,
		ConsentIPHash:   services.HashIP(c.IP()),
		NewsletterOptIn: req.NewsletterOptIn,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "failed to create lead",
			Message: err.Error(),
		})
	}
	// Notificação interna disparada async pelo service (CreateFromContactForm)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":      lead.ID,
		"status":  lead.Status,
		"message": "Recebido. A equipe Plenya entrará em contato.",
	})
}

// ============================================================
// Admin — listagem, detalhe, update, note, convert
// ============================================================

// List — GET /leads (admin/staff)
//
// @Summary  Listar leads (admin)
// @Tags     leads
// @Security BearerAuth
// @Produce  json
// @Param    source query string false "Filtro: light_claim | contact_form | whatsapp_inbound | newsletter | manual"
// @Param    status query string false "Filtro: new | contacted | qualified | converted | lost | unsubscribed"
// @Param    search query string false "Busca por name/email/phone"
// @Param    page query int false "Página (0-indexed, default 0)"
// @Param    pageSize query int false "Tamanho (default 25)"
// @Success  200 {object} services.LeadListResult
// @Router   /leads [get]
func (h *LeadHandler) List(c *fiber.Ctx) error {
	filter := services.LeadFilter{
		Search: c.Query("search"),
	}
	if s := c.Query("source"); s != "" {
		src := models.LeadSource(s)
		filter.Source = &src
	}
	if s := c.Query("status"); s != "" {
		st := models.LeadStatus(s)
		filter.Status = &st
	}
	if s := c.Query("hasEmailOptIn"); s != "" {
		if b, err := strconv.ParseBool(s); err == nil {
			filter.HasEmailOptIn = &b
		}
	}
	if s := c.Query("hasWhatsAppOptIn"); s != "" {
		if b, err := strconv.ParseBool(s); err == nil {
			filter.HasWhatsAppOptIn = &b
		}
	}
	if s := c.Query("assignedToUserId"); s != "" {
		if id, err := uuid.Parse(s); err == nil {
			filter.AssignedToUserID = &id
		}
	}
	if s := c.Query("utmCampaign"); s != "" {
		filter.UTMCampaign = &s
	}

	page, _ := strconv.Atoi(c.Query("page", "0"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "25"))

	result, err := h.service.List(filter, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to list leads",
			Message: err.Error(),
		})
	}
	return c.JSON(result)
}

// Get — GET /leads/:id
//
// @Summary  Detalhe de um lead + atividades
// @Tags     leads
// @Security BearerAuth
// @Produce  json
// @Param    id path string true "Lead ID"
// @Success  200 {object} models.Lead
// @Failure  404 {object} dto.ErrorResponse
// @Router   /leads/{id} [get]
func (h *LeadHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	lead, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to load lead",
			Message: err.Error(),
		})
	}
	return c.JSON(lead)
}

// UpdateRequest — campos atualizáveis no PATCH /leads/:id
type UpdateRequest struct {
	Status           *models.LeadStatus `json:"status,omitempty"`
	AssignedToUserID *uuid.UUID         `json:"assignedToUserId,omitempty"`
	Name             *string            `json:"name,omitempty"`
}

// Update — PATCH /leads/:id
//
// @Summary  Atualiza status/atribuição/nome do lead
// @Tags     leads
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    id path string true "Lead ID"
// @Param    body body UpdateRequest true "Patch"
// @Success  200 {object} models.Lead
// @Router   /leads/{id} [patch]
func (h *LeadHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid body",
			Message: err.Error(),
		})
	}
	actor := middleware.GetUserID(c)
	lead, err := h.service.Update(id, services.LeadPatch{
		Status:           req.Status,
		AssignedToUserID: req.AssignedToUserID,
		Name:             req.Name,
	}, actor)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to update lead",
			Message: err.Error(),
		})
	}
	return c.JSON(lead)
}

// ConvertRequest — POST /leads/:id/convert
// birthDate é obrigatório (não aceitamos placeholder) — admin precisa coletar.
type ConvertRequest struct {
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty" validate:"omitempty,email"`
	Phone     *string    `json:"phone,omitempty"`
	BirthDate *time.Time `json:"birthDate" validate:"required"`
	Gender    *string    `json:"gender,omitempty" validate:"omitempty,oneof=male female other"`
}

// Convert — converte Lead em Patient (criação manual via admin)
//
// @Summary  Converter lead em paciente (manual)
// @Tags     leads
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    id path string true "Lead ID"
// @Param    body body ConvertRequest true "Overrides opcionais (campos vêm do Lead se não fornecidos)"
// @Success  201 {object} models.Patient
// @Router   /leads/{id}/convert [post]
func (h *LeadHandler) Convert(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req ConvertRequest
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
	actor := middleware.GetUserID(c)
	patient, err := h.service.ConvertToPatient(id, actor, services.ConvertToPatientInput{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		BirthDate: req.BirthDate,
		Gender:    req.Gender,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "failed to convert lead",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(patient)
}

// Stats — GET /leads/stats?period=30d (ou 7d/90d/custom&from=YYYY-MM-DD&to=YYYY-MM-DD)
//
// @Summary  Estatísticas do funil de leads
// @Tags     leads
// @Security BearerAuth
// @Produce  json
// @Param    period query string false "7d | 30d (default) | 90d | custom"
// @Param    from query string false "YYYY-MM-DD (custom)"
// @Param    to query string false "YYYY-MM-DD (custom)"
// @Success  200 {object} services.StatsResponse
// @Router   /leads/stats [get]
func (h *LeadHandler) Stats(c *fiber.Ctx) error {
	period, err := services.ResolveStatsPeriod(c.Query("period"), c.Query("from"), c.Query("to"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid period",
			Message: err.Error(),
		})
	}
	stats, err := h.service.Stats(period)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to compute stats",
			Message: err.Error(),
		})
	}
	return c.JSON(stats)
}

// SendMessageRequest — POST /leads/:id/messages
type SendMessageRequest struct {
	Content string `json:"content" validate:"required,min=1,max=4096"`
}

// SendMessage — envia mensagem de texto (session message) via WhatsApp.
// Só funciona dentro da janela de 24h após inbound.
//
// @Summary  Enviar mensagem WhatsApp inline (session message)
// @Tags     leads
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    id path string true "Lead ID"
// @Param    body body SendMessageRequest true "Texto da mensagem"
// @Success  201 {object} models.LeadActivity
// @Router   /leads/{id}/messages [post]
func (h *LeadHandler) SendMessage(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req SendMessageRequest
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
	actor := middleware.GetUserID(c)
	activity, err := h.service.SendWhatsAppText(id, actor, req.Content)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "failed to send message",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(activity)
}

// SendEmailReplyRequest — POST /leads/:id/email-reply
type SendEmailReplyRequest struct {
	Subject    string   `json:"subject" validate:"omitempty,max=200"`
	BodyText   string   `json:"bodyText" validate:"required,min=1,max=50000"`
	BodyHTML   string   `json:"bodyHTML" validate:"omitempty,max=200000"`
	InReplyTo  string   `json:"inReplyTo,omitempty"`
	References []string `json:"references,omitempty"`
}

// SendEmailReply — envia resposta por email pro lead via Resend, espelha em Sent (IMAP),
// e registra LeadActivity{message_sent, channel=email}.
//
// @Summary  Responder email do lead
// @Tags     leads
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    id path string true "Lead ID"
// @Param    body body SendEmailReplyRequest true "Email"
// @Success  200 {object} fiber.Map
// @Failure  404 {object} dto.ErrorResponse
// @Failure  422 {object} dto.ErrorResponse
// @Router   /leads/{id}/email-reply [post]
func (h *LeadHandler) SendEmailReply(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req SendEmailReplyRequest
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

	lead, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to load lead",
			Message: err.Error(),
		})
	}

	actor := middleware.GetUserID(c)
	in := services.SendLeadReplyInput{
		Lead:        lead,
		ActorUserID: actor,
		Subject:     strings.TrimSpace(req.Subject),
		BodyText:    req.BodyText,
		BodyHTML:    req.BodyHTML, // só usado se UI enviar HTML pré-renderizado; default vazio
		InReplyTo:   strings.TrimSpace(req.InReplyTo),
		References:  req.References,
	}

	if err := h.emailService.SendLeadReply(c.UserContext(), in); err != nil {
		switch {
		case errors.Is(err, services.ErrLeadOptedOut):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "lead opt-out de email",
			})
		case errors.Is(err, services.ErrLeadNoEmail):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "lead sem email cadastrado",
			})
		default:
			return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{
				Error:   "failed to send email",
				Message: err.Error(),
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Email enviado",
		"to":      lead.Email,
	})
}

// Delete — DELETE /leads/:id (admin only via middleware)
//
// @Summary  Excluir lead (LGPD direito de eliminação, art. 18 VI)
// @Tags     leads
// @Security BearerAuth
// @Param    id path string true "Lead ID"
// @Success  204
// @Failure  404 {object} dto.ErrorResponse
// @Router   /leads/{id} [delete]
func (h *LeadHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "lead not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to delete lead",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// AddNoteRequest — POST /leads/:id/notes
type AddNoteRequest struct {
	Note string `json:"note" validate:"required,min=1,max=2000"`
}

// AddNote — adiciona uma nota livre ao lead
//
// @Summary  Adicionar nota ao lead
// @Tags     leads
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Param    id path string true "Lead ID"
// @Param    body body AddNoteRequest true "Nota"
// @Success  201
// @Router   /leads/{id}/notes [post]
func (h *LeadHandler) AddNote(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req AddNoteRequest
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
	actor := middleware.GetUserID(c)
	if err := h.service.AddNote(id, actor, req.Note); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to add note",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}
